package block

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/opsidian/ocl/ocl"
	"github.com/opsidian/ocl/util"
)

// ParseFields parses all fields of a given go struct
func ParseFields(str *ast.StructType, file *ast.File) ([]*Field, error) {
	fields := make([]*Field, 0, len(str.Fields.List))

	var idField string
	var valueField string

	for _, field := range str.Fields.List {
		field, err := parseField(field)
		if err != nil {
			return nil, err
		}

		if field.IsID {
			if idField != "" {
				return nil, fmt.Errorf("multiple id fields were found: %s, %s", idField, field.Name)
			}
			idField = field.Name
		}

		if field.IsValue {
			if valueField != "" {
				return nil, fmt.Errorf("multiple value fields were found: %s, %s", valueField, field.Name)
			}
			valueField = field.Name
		}

		if field != nil {
			fields = append(fields, field)
		}
	}

	if idField == "" {
		return nil, fmt.Errorf("no fields were found with the \"id\" tag")
	}

	if valueField != "" {
		for _, field := range fields {
			if !field.IsValue && field.Required {
				return nil, errors.New("when setting a value field then no other fields can be required")
			}
		}
	}

	return fields, nil
}

func parseField(field *ast.Field) (*Field, error) {
	var tag string
	name := field.Names[0].String()
	if field.Tag != nil {
		var err error
		tag, err = strconv.Unquote(field.Tag.Value)
		if err != nil {
			return nil, fmt.Errorf("tag is invalid for %s", name)
		}
	}

	tags := util.ParseFieldTag(reflect.StructTag(tag), "ocl", name)

	for _, key := range tags.Keys() {
		if _, validTag := ocl.BlockTags[strings.ToLower(key)]; !validTag {
			return nil, fmt.Errorf("invalid tag %q on field %q", key, name)
		}
	}

	if tags.GetBool(ocl.BlockTagIgnore) {
		return nil, nil
	}

	return &Field{
		Name:        name,
		ParamName:   tags.GetWithDefault(ocl.BlockTagName, generateParamName(name)),
		Required:    tags.GetBool(ocl.BlockTagRequired),
		Type:        getFieldType(field.Type),
		Stage:       tags.GetWithDefault(ocl.BlockTagStage, "default"),
		IsID:        tags.GetBool(ocl.BlockTagID),
		IsValue:     tags.GetBool(ocl.BlockTagValue),
		IsReference: tags.GetBool(ocl.BlockTagReference),
		IsBlock:     tags.GetBool(ocl.BlockTagBlock),
		IsFactory:   tags.GetBool(ocl.BlockTagFactory),
	}, nil
}

func getFieldType(typeNode ast.Expr) string {
	switch t := typeNode.(type) {
	case *ast.Ident:
		return t.String()
	default:
		b := &bytes.Buffer{}
		format.Node(b, token.NewFileSet(), t)
		return b.String()
	}
}

func generateParamName(name string) string {
	re := regexp.MustCompile("[A-Z][a-z0-9]*")
	name = re.ReplaceAllStringFunc(name, func(str string) string {
		return "_" + strings.ToLower(str)
	})
	return strings.TrimLeft(name, "_")
}

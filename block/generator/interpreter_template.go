package generator

type InterpreterTemplateParams struct {
	Package                string
	Type                   string
	Name                   string
	HasForeignID           bool
	Stages                 []string
	BlockTypes             map[string]string
	NodeTypes              map[string]string
	IDField                *Field
	ValueField             *Field
	Fields                 []*Field
	NodeValueFunctionNames map[string]string
	EvalFieldsCnt          int
}

const interpreterTemplate = `
// Code generated by Basil. DO NOT EDIT.
package {{.Package}}

import (
	"errors"
	"fmt"
	"strings"

	"github.com/opsidian/basil/basil"
	basilblock "github.com/opsidian/basil/block"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

{{ $root := .}}

type {{$root.Name}}Interpreter struct {}

func (i {{$root.Name}}Interpreter) StaticCheck(ctx interface{}, node basil.BlockNode) (string, parsley.Error) {
	validParamNames := map[variable.ID]struct{}{
		{{- range $root.Fields -}}{{- if and (not .IsID) (not .IsBlock)}}
		"{{.ParamName}}": struct{}{},
		{{- end -}}{{- end }}
	}

	for paramName, paramNode := range node.ParamNodes() {
		if _, valid := validParamNames[paramName]; !valid {
			return "", parsley.NewError(paramNode.Pos(), fmt.Errorf("%q parameter does not exist", paramName))
		}
	}

	{{ range $root.Fields }}{{ if and (not .IsID) (not .IsNode) (not .IsBlock) }}
	if paramNode, ok := node.ParamNodes()["{{.ParamName}}"]; ok {
		if err := variable.CheckNodeType(paramNode, "{{.Type}}"); err != nil {
			return "", err
		}
	}
	{{ end }}{{ end }}

	requiredParamNames := []variable.ID{
		{{- range $root.Fields -}}{{- if and (.Required) (not .IsID) (not .IsNode) (not .IsBlock)}}
		"{{.ParamName}}",
		{{- end -}}{{- end }}
	}

	for _, paramName := range requiredParamNames {
		if _, set := node.ParamNodes()[paramName]; !set {
			return "", parsley.NewError(node.Pos(), fmt.Errorf("%s parameter is required", paramName))
		}
	}

	return "*{{$root.Name}}", nil
}

// CreateBlock creates a new {{$root.Name}} block
func (i {{$root.Name}}Interpreter) Eval(parentCtx interface{}, node basil.BlockNode) (basil.Block, parsley.Error) {
	block := &{{$root.Name}}{
		{{.IDField.Name}}: node.ID(),
	}

	ctx := block.Context(parentCtx)

	{{ if len .NodeTypes }}
	for _, blockNode := range node.BlockNodes() {
		switch b := blockNode.(type) {
		{{- range $type, $fieldName := .NodeTypes -}}
		case {{trimPrefix $type "[]" }}:
			block.{{$fieldName}} = append(block.{{$fieldName}}, b)
		{{- end}}
		}
	}
	{{ end}}

	if err := i.EvalBlock(ctx, node, "default", block); err != nil {
		return nil, err
	}

	return block, nil
}

// EvalBlock evaluates all fields belonging to the given stage on a {{$root.Name}} block
func (i {{$root.Name}}Interpreter) EvalBlock(ctx interface{}, node basil.BlockNode, stage string, res basil.Block) parsley.Error {
	var err parsley.Error

	if preInterpreter, ok := res.(basil.BlockPreInterpreter); ok {
		if err = preInterpreter.PreEval(ctx, stage); err != nil {
			return err
		}
	}

	{{ if .EvalFieldsCnt -}}
	block, ok := res.(*{{$root.Name}})
	if !ok {
		panic("result must be a type of *{{$root.Name}}")
	}

	switch stage {
	{{- range $stage := $root.Stages}}
	case "{{$stage}}":
		{{- range $root.Fields -}}{{- if and (eq .Stage $stage) (not .IsID) (not .IsNode) (not .IsBlock)}}
		if valueNode, ok := node.ParamNodes()["{{.ParamName}}"]; ok {
			if block.{{.Name}}, err = variable.{{index $root.NodeValueFunctionNames .Type}}(valueNode, ctx); err != nil {
				return err
			}
		}
		{{ end }}{{ end -}}
	{{end -}}
	default:
		panic(fmt.Sprintf("unknown stage: %s", stage))
	}

	switch stage {
	{{- range $stage := $root.Stages}}
	case "{{$stage}}":
		{{- range $root.Fields }}{{ if and (eq .Stage $stage) .IsNode}}
		var childBlock interface{}
		for _, childBlockNode := range node.BlockNodes() {
			{{ if $root.BlockTypes -}}
			if childBlock, err = childBlockNode.Value(ctx); err != nil {
				return err
			}

			switch b := childBlock.(type) {
			{{- range $type, $fieldName := $root.BlockTypes -}}
			case {{trimPrefix $type "[]" }}:
				block.{{$fieldName}} = append(block.{{$fieldName}}, b)
			{{- end}}
			default:
				panic(fmt.Sprintf("block type %T is not supported in {{$root.Name}}, please open a bug ticket", childBlock))
			}
			{{ else }}
			panic(fmt.Sprintf("block type %T is not supported in {{$root.Name}}, please open a bug ticket", childBlock))
			{{- end }}
		}
		{{ end }}{{ end -}}
	{{- end -}}
	default:
		panic(fmt.Sprintf("unknown stage: %s", stage))
	}
	{{- end }}

	if postInterpreter, ok := res.(basil.BlockPostInterpreter); ok {
		if err = postInterpreter.PostEval(ctx, stage); err != nil {
			return err
		}
	}

	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i {{$root.Name}}Interpreter) HasForeignID() bool {
	return {{.HasForeignID}}
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i {{$root.Name}}Interpreter) ValueParamName() variable.ID {
	return {{ if .ValueField }}variable.ID("{{.ValueField.ParamName}}"){{ else }}""{{ end }}
}

func (i {{$root.Name}}Interpreter) BlockRegistry() parsley.NodeTransformerRegistry {
	var block basil.Block = &{{$root.Name}}{}
	if b, ok := block.(basil.BlockRegistryAware); ok {
		return b.BlockRegistry()
	}

	return nil
}
`
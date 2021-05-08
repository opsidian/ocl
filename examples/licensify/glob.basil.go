// Code generated by Basil. DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// GlobInterpreter is the basil interpreter for the Glob block
type GlobInterpreter struct {
	s schema.Schema
}

func (i GlobInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Glob",
			Properties: map[string]schema.Schema{
				"exclude": &schema.Array{
					Items: &schema.String{},
				},
				"file": &schema.Reference{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"eval_stage": "init", "generated": "true"},
						Pointer:     true,
					},
					Ref: "http://basil.schema/github.com/opsidian/basil/examples/licensify/File",
				},
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
				"include": &schema.Array{
					Items: &schema.String{},
				},
				"path": &schema.String{},
			},
			Required: []string{"path", "file"},
		}
	}
	return i.s
}

// Create creates a new Glob block
func (i GlobInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Glob{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i GlobInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i GlobInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Glob
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i GlobInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "exclude":
		return b.(*Glob).exclude
	case "id":
		return b.(*Glob).id
	case "include":
		return b.(*Glob).include
	case "path":
		return b.(*Glob).path
	default:
		panic(fmt.Errorf("unexpected parameter %q in Glob", name))
	}
}

func (i GlobInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Glob)
	switch name {
	case "exclude":
		b.exclude = make([]string, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.exclude[valuek] = valuev.(string)
		}
	case "include":
		b.include = make([]string, len(value.([]interface{})))
		for valuek, valuev := range value.([]interface{}) {
			b.include[valuek] = valuev.(string)
		}
	case "path":
		b.path = value.(string)
	}
	return nil
}

func (i GlobInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Glob)
	switch name {
	case "file":
		b.file = value.(*File)
	}
	return nil
}

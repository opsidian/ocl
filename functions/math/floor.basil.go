// Code generated by Basil. DO NOT EDIT.

package math

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
	"github.com/opsidian/parsley/parsley"
)

// FloorInterpreter is the basil interpreter for the Floor function
type FloorInterpreter struct {
	s schema.Schema
}

func (i FloorInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It returns the greatest integer value less than or equal to x.",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name: "number",
					Schema: &schema.Untyped{
						Types: []string{"integer", "number"},
					},
				},
			},
			Result: &schema.Integer{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i FloorInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	parameters := i.Schema().(*schema.Function).GetParameters()
	arguments := node.ArgumentNodes()

	arg0, evalErr := parsley.EvaluateNode(ctx, arguments[0])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[0].Schema.ValidateValue(arg0); err != nil {
		return nil, parsley.NewError(arguments[0].Pos(), err)
	}
	var val0 = arg0

	return Floor(val0), nil
}

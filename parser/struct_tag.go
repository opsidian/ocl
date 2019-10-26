// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/text"
	"github.com/opsidian/parsley/text/terminal"
)

func StructTag() *combinator.Sequence {
	var value parser.Func

	value = combinator.Choice(
		terminal.TimeDuration(),
		terminal.Float(),
		terminal.Integer(),
		terminal.String(true),
		terminal.Bool("true", "false"),
		Array(&value, text.WsSpaces),
		ID(basil.IDRegExpPattern),
	).Name("value")

	keyValue := combinator.SeqFirstOrAll(
		ID(basil.IDRegExpPattern),
		text.LeftTrim(terminal.Rune('='), text.WsSpaces),
		text.LeftTrim(&value, text.WsSpaces),
	).Name("parameter name or parameter name and value pair")

	return SepByComma(keyValue, text.WsSpaces).Bind(structTagInterpreter{})
}

type structTagInterpreter struct {
}

func (s structTagInterpreter) Eval(userCtx interface{}, node parsley.NonTerminalNode) (interface{}, parsley.Error) {
	nodes := node.(parsley.NonTerminalNode).Children()
	res := make(map[basil.ID]interface{}, (len(nodes)+1)/2)
	for i := range nodes {
		if i%2 == 0 {
			parts := nodes[i].(parsley.NonTerminalNode).Children()
			idNode := parts[0].(*basil.IDNode)

			if _, exists := res[idNode.ID()]; exists {
				return nil, parsley.NewError(idNode.Pos(), fmt.Errorf("parameter %q was already defined", idNode.ID()))
			}
			if len(parts) >= 3 {
				val, err := parts[2].Value(userCtx)
				if err != nil {
					return nil, err
				}
				res[idNode.ID()] = val
			} else {
				res[idNode.ID()] = true
			}
		}
	}
	return res, nil
}

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"fmt"

	"github.com/opsidian/parsley/parser"

	"github.com/opsidian/basil/basil/block"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/text"
	"github.com/opsidian/parsley/text/terminal"
)

// Block returns a parser for parsing blocks
//   S     -> ID ID? {
//              (ATTR|S)*
//            }
//         -> ID ID? VALUE
//   ID    -> /[a-z][a-z0-9]*(?:_[a-z0-9]+)*/
//   ATTR  -> ID ("="|":=") P
//   VALUE -> EXPRESSION
//         -> ARRAY
//         -> MAP
func Block(expr parsley.Parser) *combinator.Sequence {
	return blockWithOptions(expr, true, true, true)
}

func blockWithOptions(
	expr parsley.Parser,
	allowID bool,
	allowCustomParameters bool,
	allowDirectives bool,
) *combinator.Sequence {
	var p combinator.Sequence

	var directives parsley.Parser
	if allowDirectives {
		directives = combinator.Many(text.RightTrim(Directive(expr), text.WsSpacesForceNl))
	} else {
		directives = parser.Empty()
	}

	paramOrBlock := combinator.Choice(
		Parameter(expr, allowCustomParameters),
		&p,
	).Name("parameter or block")

	emptyBody := combinator.SeqOf(
		terminal.Rune('{'),
		text.LeftTrim(terminal.Rune('}'), text.WsSpacesNl),
	).Token(block.TokenBody)

	body := combinator.SeqOf(
		terminal.Rune('{'),
		combinator.Many(text.LeftTrim(paramOrBlock, text.WsSpacesForceNl)),
		text.LeftTrim(terminal.Rune('}'), text.WsSpacesForceNl),
	).Token(block.TokenBody)

	blockValue := combinator.Choice(
		emptyBody,
		body,
		expr,
		MultilineText(),
		Array(expr, text.WsSpacesNl),
		Map(expr),
		parser.Empty(),
	).Name("block value")

	var id parsley.Parser
	if allowID {
		id = combinator.SeqTry(
			ID(basil.IDRegExpPattern),
			text.LeftTrim(ID(basil.IDRegExpPattern), text.WsSpaces),
		).ReturnSingle()
	} else {
		id = ID(basil.IDRegExpPattern)
	}

	p = *combinator.SeqOf(
		directives,
		id,
		text.LeftTrim(blockValue, text.WsSpaces),
	).Name("block").Token(block.Token).Bind(blockInterpreter{})

	return &p
}

type blockInterpreter struct{}

func (b blockInterpreter) Eval(userCtx interface{}, node parsley.NonTerminalNode) (interface{}, parsley.Error) {
	panic("Eval should not be called on a raw block node")
}

func (b blockInterpreter) TransformNode(userCtx interface{}, node parsley.Node) (parsley.Node, parsley.Error) {
	registry := userCtx.(basil.BlockTransformerRegistryAware).BlockTransformerRegistry()

	nodes := node.(parsley.NonTerminalNode).Children()

	var typeNode *basil.IDNode
	switch n := nodes[1].(type) {
	case parsley.NonTerminalNode:
		typeNode = n.Children()[0].(*basil.IDNode)
	case *basil.IDNode:
		typeNode = n
	default:
		panic(fmt.Errorf("unexpected node type: %T", nodes[1]))
	}

	transformer, exists := registry.NodeTransformer(string(typeNode.ID()))
	if !exists {
		return nil, parsley.NewError(typeNode.Pos(), fmt.Errorf("%q block is unknown or not allowed", typeNode.ID()))
	}

	return transformer.TransformNode(userCtx, node)
}

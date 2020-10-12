// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parsley

// Node represents an AST node
//go:generate counterfeiter . Node
type Node interface {
	Token() string
	Type() string
	Value(userCtx interface{}) (interface{}, Error)
	Pos() Pos
	ReaderPos() Pos
}

// NonTerminalNode represents a nonterminal AST node
//go:generate counterfeiter . NonTerminalNode
type NonTerminalNode interface {
	Node
	Children() []Node
}
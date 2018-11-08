// Copyright (c) 2018 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package function

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

// Node is a function node
type Node struct {
	nameNode      parsley.Node
	argumentNodes []parsley.Node
	readerPos     parsley.Pos
	interpreter   Interpreter
	resultType    string
}

// Name returns with the function name
func (n *Node) Name() basil.ID {
	name, _ := n.nameNode.Value(nil)
	return name.(basil.ID)
}

// Token returns with the node's token
func (n *Node) Token() string {
	return "FUNC"
}

// Type returns with the node's type
func (n *Node) Type() string {
	return n.resultType
}

// StaticCheck runs static analysis on the node
func (n *Node) StaticCheck(ctx interface{}) parsley.Error {
	_, err := n.interpreter.StaticCheck(ctx, n)
	if err != nil {
		return err
	}

	return nil
}

// Value returns with the result of the function
func (n *Node) Value(ctx interface{}) (interface{}, parsley.Error) {
	return n.interpreter.Eval(ctx, n)
}

// Pos returns with the node's position
func (n *Node) Pos() parsley.Pos {
	return n.nameNode.Pos()
}

// ReaderPos returns with the reader's position
func (n *Node) ReaderPos() parsley.Pos {
	return n.readerPos
}

// SetReaderPos amends the reader position using the given function
func (n *Node) SetReaderPos(f func(parsley.Pos) parsley.Pos) {
	n.readerPos = f(n.readerPos)
}

// ArgumentNodes returns with the function argument nodes
func (n *Node) ArgumentNodes() []parsley.Node {
	return n.argumentNodes
}

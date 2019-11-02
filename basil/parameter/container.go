// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parameter

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

var _ basil.ParameterContainer = &Container{}

// Container is a parameter container
type Container struct {
	evalCtx *basil.EvalContext
	node    basil.ParameterNode
	parent  basil.BlockContainer
	value   interface{}
	err     parsley.Error
	jobID   basil.ID
	wgs     []basil.WaitGroup
}

// NewContainer creates a new parameter container
func NewContainer(
	evalCtx *basil.EvalContext,
	node basil.ParameterNode,
	parent basil.BlockContainer,
	value interface{},
	wgs []basil.WaitGroup,
) *Container {
	return &Container{
		evalCtx: evalCtx,
		node:    node,
		parent:  parent,
		value:   value,
		wgs:     wgs,
	}
}

// ID returns with the parameter id
func (c *Container) ID() basil.ID {
	return c.node.ID()
}

// JobName returns with the job name
func (c *Container) JobName() basil.ID {
	return c.node.ID()
}

// ID returns with the block id
func (c *Container) JobID() basil.ID {
	return c.jobID
}

// SetJobID sets the job id
func (c *Container) SetJobID(id basil.ID) {
	c.jobID = id
}

// Node returns with the parameter node
func (c *Container) Node() basil.Node {
	return c.node
}

// BlockContainer returns with the parent block container
func (c *Container) BlockContainer() basil.BlockContainer {
	return c.parent
}

// Value returns with the parameter value or an evaluation error
func (c *Container) Value() (interface{}, parsley.Error) {
	if c.err != nil {
		return nil, c.err
	}

	return c.value, nil
}

// Run evaluates the parameter
func (c *Container) Run() {
	if !c.evalCtx.Run() {
		return
	}

	c.value, c.err = c.node.Value(c.evalCtx)
	c.parent.SetChild(c)
}

func (c *Container) Cancel() bool {
	return c.evalCtx.Cancel()
}

func (c *Container) Lightweight() bool {
	return true
}

// Close notifies all wait groups
func (c *Container) Close() {
	for _, wg := range c.wgs {
		wg.Done(c.err)
	}
}

// WaitGroups returns nil
func (c *Container) WaitGroups() []basil.WaitGroup {
	return c.wgs
}

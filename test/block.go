// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package test

import (
	"fmt"
	"time"

	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/block"
	"github.com/opsidian/basil/basil/schema"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . BlockWithInit
type BlockWithInit interface {
	basil.Block
	basil.BlockInitialiser
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . BlockWithMain
type BlockWithMain interface {
	basil.Block
	basil.BlockRunner
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . BlockWithClose
type BlockWithClose interface {
	basil.Block
	basil.BlockCloser
}

// @block
type Block struct {
	// @id
	IDField basil.ID
	// @value
	Value             interface{}
	FieldString       string
	FieldInt          int64
	FieldFloat        float64
	FieldBool         bool
	FieldArray        []interface{}
	FieldMap          map[string]interface{}
	FieldTimeDuration time.Duration
	// @name "custom_field"
	FieldCustomName string

	// @name "testblock"
	Blocks []*Block
}

func (b *Block) ID() basil.ID {
	return b.IDField
}

func (b *Block) ParseContextOverride() basil.ParseContextOverride {
	return basil.ParseContextOverride{
		BlockTransformerRegistry: block.InterpreterRegistry{
			"testblock": BlockInterpreter{},
		},
	}
}

func (b *Block) Compare(b2 *Block, input string) {
	interpreter := BlockInterpreter{}
	compareBlocks(b, b2, interpreter, input)
	Expect(len(b.Blocks)).To(Equal(len(b2.Blocks)), "child block count does not match, input: %s", input)
	for i2 := range b2.Blocks {
		found := false
		for i := range b.Blocks {
			if b.Blocks[i].ID() == b2.Blocks[i2].ID() {
				compareBlocks(b.Blocks[i], b2.Blocks[i2], interpreter, input)
				found = true
				break
			}
		}
		if !found {
			ginkgo.Fail(fmt.Sprintf("block not found with id %s", b2.Blocks[i2].ID()))
		}
	}
}

func compareBlocks(b1, b2 basil.Identifiable, interpreter basil.BlockInterpreter, input string) {
	Expect(b1.ID()).To(Equal(b2.ID()), "id does not match, input: %s", input)

	for propertyName, p := range interpreter.Schema().(schema.ObjectKind).GetProperties() {
		if block.IsBlockSchema(p) {
			continue
		}

		v1 := interpreter.Param(b1, basil.ID(propertyName))
		v2 := interpreter.Param(b2, basil.ID(propertyName))
		if v2 != nil {
			Expect(v1).To(Equal(v2), "%s does not match, input: %s", propertyName, input)
		} else {
			Expect(v1).To(BeNil(), "%s does not match, input: %s", propertyName, input)
		}
	}
}

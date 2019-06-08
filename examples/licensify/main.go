// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/block"
	"github.com/opsidian/basil/examples/common"
	"github.com/opsidian/basil/parser"
	"github.com/opsidian/basil/util"
)

//go:generate basil generate
type Main struct {
	id basil.ID `basil:"id"`
}

func (m *Main) ParseContextOverride() basil.ParseContextOverride {
	return basil.ParseContextOverride{
		BlockTransformerRegistry: block.InterpreterRegistry{
			"glob":      GlobInterpreter{},
			"licensify": LicensifyInterpreter{},
			"print":     common.PrintInterpreter{},
			"println":   common.PrintlnInterpreter{},
		},
	}
}

func main() {
	ctx, cancel := util.CreateDefaultContext()
	defer cancel()

	parseCtx := basil.NewParseContext(basil.NewIDRegistry(8, 16))

	p := parser.NewMain("main", MainInterpreter{})

	if err := p.ParseFile(
		parseCtx,
		"main.basil",
	); err != nil {
		panic(err)
	}

	common.Main(ctx, parseCtx)
}
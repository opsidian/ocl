// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package block

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

type mainInterpreter struct {
	basil.BlockInterpreter
	schema schema.Schema
}

func (m *mainInterpreter) Schema() schema.Schema {
	return m.schema
}

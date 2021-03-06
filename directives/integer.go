// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// @block {
//   eval_stage = "parse"
// }
type Integer struct {
	schema.Integer
}

func (i *Integer) ApplyToParameterConfig(config *basil.ParameterConfig) {
	config.Schema = &i.Integer
}

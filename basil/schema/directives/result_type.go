// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import (
	"errors"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// @block
type ResultType struct {
	// @id
	id basil.ID
}

func (r *ResultType) ID() basil.ID {
	return r.id
}

func (r *ResultType) ApplyToSchema(s schema.Schema) error {
	if s.Type() != schema.TypeUntyped {
		return errors.New("@result_type can only be set on an interface{} parameter")
	}
	return nil
}

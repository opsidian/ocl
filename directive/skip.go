// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directive

import "github.com/opsidian/basil/basil"

//go:generate basil generate
type Skip struct {
	id   basil.ID `basil:"id"`
	cond bool     `basil:"value"`
}

func (s Skip) ID() basil.ID {
	return s.id
}

func (s Skip) ApplyDirective(blockCtx basil.BlockContext, container basil.BlockContainer) error {
	if s.cond {
		container.Skip()
	}
	return nil
}

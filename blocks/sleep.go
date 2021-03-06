// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package blocks

import (
	"errors"
	"time"

	"github.com/opsidian/basil/basil"
)

// @block
type Sleep struct {
	// @id
	id basil.ID
	// @value
	// @required
	duration time.Duration
}

func (s *Sleep) ID() basil.ID {
	return s.id
}

func (s *Sleep) Main(blockCtx basil.BlockContext) error {
	select {
	case <-time.After(s.duration):
		return nil
	case <-blockCtx.Done():
		return errors.New("aborted")
	}
}

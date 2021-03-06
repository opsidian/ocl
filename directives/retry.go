// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import (
	"time"

	"github.com/opsidian/basil/basil"
)

// @block {
//   eval_stage = "init"
// }
type Retry struct {
	// @id
	id basil.ID
	// @value
	// @required
	count int64
}

func (r *Retry) ID() basil.ID {
	return r.id
}

func (r *Retry) ApplyToRuntimeConfig(config *basil.RuntimeConfig) {
	config.Retry = r
}

func (r *Retry) RetryCount() int {
	return int(r.count)
}

func (r *Retry) RetryDelay(int) time.Duration {
	return 0
}

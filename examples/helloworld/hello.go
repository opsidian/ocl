// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/opsidian/basil/basil"
)

// Hello is capable to print some greetings
// @block
type Hello struct {
	// @id
	id basil.ID
	// @required
	to string
	// @read_only
	greeting string
	// @ignore
	r *rand.Rand
}

func (h *Hello) ID() basil.ID {
	return h.id
}

// Init will initialise the random generator
func (h *Hello) Init(ctx basil.BlockContext) (bool, error) {
	h.r = rand.New(rand.NewSource(time.Now().Unix()))
	return false, nil
}

// Main will generate a random greeting
func (h *Hello) Main(ctx basil.BlockContext) error {
	greetings := []string{"Hello", "Hi", "Hey", "Yo", "Sup"}

	h.greeting = fmt.Sprintf("%s %s!", greetings[h.r.Intn(len(greetings))], h.to)

	return nil
}

// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package functions_test

import (
	"errors"

	"github.com/opsidian/basil/basil/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	basilfunction "github.com/opsidian/basil/basil/function"
	"github.com/opsidian/basil/functions"
	"github.com/opsidian/basil/parsers"
	"github.com/opsidian/basil/test"
	"github.com/opsidian/parsley/parsley"
)

var _ = Describe("String", func() {

	registry := basilfunction.InterpreterRegistry{
		"test": functions.StringInterpreter{},
	}

	DescribeTable("it evaluates the input correctly",
		func(input string, expected interface{}) {
			test.ExpectFunctionToEvaluate(parsers.Expression(), registry)(input, expected)
		},
		test.TableEntry(`test(1)`, "1"),
		test.TableEntry(`test(1.1)`, "1.1"),
		test.TableEntry(`test(false)`, "false"),
		test.TableEntry(`test(true)`, "true"),
		test.TableEntry(`test("foo")`, "foo"),
		test.TableEntry(`test(1m30s)`, "1m30s"),
		test.TableEntry(`test([1, 2])`, "[1, 2]"),
		test.TableEntry(`test([1.1, 2.2])`, "[1.1, 2.2]"),
		test.TableEntry(`test([false, true])`, "[false, true]"),
		test.TableEntry(`test(["foo", "bar"])`, "[foo, bar]"),
		test.TableEntry(`test([1s, 2s])`, "[1s, 2s]"),
		test.TableEntry(`test([[1, 2], [3, 4]])`, "[[1, 2], [3, 4]]"),
	)

	DescribeTable("it will have a parse error",
		func(input string, expectedErr error) {
			test.ExpectFunctionToHaveParseError(parsers.Expression(), registry)(input, expectedErr)
		},
		test.TableEntry(`test()`, errors.New("test requires exactly 1 argument, but got 0 at testfile:1:1")),
		test.TableEntry(`test(1, 2)`, errors.New("test requires exactly 1 argument, but got 2 at testfile:1:9")),
		test.TableEntry(`test([1, "foo"])`, errors.New("items must have the same type, but found integer and string at testfile:1:6")),
	)

	It("should return with string type", func() {
		test.ExpectFunctionNode(parsers.Expression(), registry)(
			"test(1)",
			func(userCtx interface{}, node parsley.Node) {
				Expect(node.Schema().(schema.Schema).Type()).To(Equal(schema.TypeString))
			},
		)
	})

})

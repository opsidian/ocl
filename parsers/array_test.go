// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parsers_test

import (
	"errors"

	"github.com/opsidian/basil/basil/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opsidian/basil/parsers"
	"github.com/opsidian/basil/test"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/text/terminal"
)

var _ = Describe("Array", func() {

	var p combinator.Sequence
	q := combinator.Choice(
		terminal.String(schema.StringValue(), false),
		terminal.Float(schema.NumberValue()),
		terminal.Integer(schema.IntegerValue()),
		terminal.Nil(schema.NullValue(), "NULL"),
		test.EvalErrorParser(schema.IntegerValue(), "ERR"),
		&p,
	).Name("value")

	p = *parsers.Array(q)

	Describe("when new lines are allowed", func() {

		DescribeTable("it evaluates the input correctly",
			func(input string, expected interface{}) {
				test.ExpectParserToEvaluate(&p)(input, expected)
			},
			test.TableEntry("[]", []interface{}{}),
			test.TableEntry("[NULL]", []interface{}{nil}),
			test.TableEntry("[1]", []interface{}{int64(1)}),
			test.TableEntry("[\n]", []interface{}{}),
			test.TableEntry("[\nNULL,\n]", []interface{}{nil}),
			test.TableEntry("[\n1\n]", []interface{}{int64(1)}),
			test.TableEntry("[\n1,\n]", []interface{}{int64(1)}),
			test.TableEntry(`[1, 2.1]`, []interface{}{int64(1), 2.1}),
		)

		DescribeTable("it returns a static check error",
			func(input string, expectedErr error) {
				test.ExpectParserToHaveStaticCheckError(&p)(input, expectedErr)
			},
			test.TableEntry(`[1, "foo"]`, errors.New("items must have the same type, but found integer and string at testfile:1:1")),
			test.TableEntry(`[[1, 2], [3, "foo"]]`, errors.New("items must have the same type, but found integer and string at testfile:1:10")),
		)

		DescribeTable("it returns a parse error",
			func(input string, expectedErr error) {
				test.ExpectParserToHaveParseError(&p)(input, expectedErr)
			},
			test.TableEntry("[", errors.New("was expecting \"]\" at testfile:1:2")),
			test.TableEntry("[1", errors.New("was expecting \"]\" at testfile:1:3")),
			test.TableEntry("[1,", errors.New("was expecting \"]\" at testfile:1:4")),
		)

		DescribeTable("it returns an eval error",
			func(input string, expectedErr error) {
				test.ExpectParserToHaveEvalError(&p)(input, expectedErr)
			},
			test.TableEntry("[ERR, 1]", errors.New("ERR at testfile:1:2")),
			test.TableEntry("[1, ERR]", errors.New("ERR at testfile:1:5")),
		)
	})

})

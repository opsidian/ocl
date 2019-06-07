package function

import (
	"github.com/opsidian/basil/basil/function"
	"github.com/opsidian/basil/function/json"
	"github.com/opsidian/basil/function/math"
	"github.com/opsidian/basil/function/strings"
)

func Registry() function.InterpreterRegistry {
	return function.InterpreterRegistry{
		"array_contains":          ArrayContainsInterpreter{},
		"is_empty":                IsEmptyInterpreter{},
		"len":                     LenInterpreter{},
		"string":                  StringInterpreter{},
		"math.abs":                math.AbsInterpreter{},
		"math.ceil":               math.CeilInterpreter{},
		"math.floor":              math.FloorInterpreter{},
		"strings.has_prefix":      strings.HasPrefixInterpreter{},
		"strings.has_suffix":      strings.HasSuffixInterpreter{},
		"strings.join":            strings.JoinInterpreter{},
		"strings.lower":           strings.LowerInterpreter{},
		"strings.replace":         strings.ReplaceInterpreter{},
		"strings.split":           strings.SplitInterpreter{},
		"strings.string_contains": strings.ContainsInterpreter{},
		"strings.title":           strings.TitleInterpreter{},
		"strings.trim_prefix":     strings.TrimPrefixInterpreter{},
		"strings.trim_space":      strings.TrimSpaceInterpreter{},
		"strings.trim_suffix":     strings.TrimSuffixInterpreter{},
		"strings.upper":           strings.UpperInterpreter{},
		"json_decode":             json.DecodeInterpreter{},
		"json_encode":             json.EncodeInterpreter{},
	}
}

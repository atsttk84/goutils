package trace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrace(t *testing.T) {
	assert.Equal(t, "message=[test string] call=[github.com/atsttk84/goutils/trace.TestTrace] parameters=[string]", Trace("test string", "string"))
	assert.Equal(t, "message=[test int] call=[github.com/atsttk84/goutils/trace.TestTrace] parameters=[1]", Trace("test int", 1))
	assert.Equal(t, "message=[test map] call=[github.com/atsttk84/goutils/trace.TestTrace] parameters=[map[k:1]]", Trace("test map", map[string]int{"k": 1}))
	assert.Equal(t, "message=[test slice] call=[github.com/atsttk84/goutils/trace.TestTrace] parameters=[[1 2 3]]", Trace("test slice", []int{1, 2, 3}))
	assert.Equal(t, "message=[test variadic argument] call=[github.com/atsttk84/goutils/trace.TestTrace] parameters=[1 2 3]", Trace("test variadic argument", 1, 2, 3))
}

func TestTraceLine(t *testing.T) {
	assert.Equal(t, "message=[test string] call=[github.com/atsttk84/goutils/trace.TestTraceLine] parameters=[string] line=[/go/src/github.com/atsttk84/goutils/trace/trace_test.go:17]", TraceLine("test string", "string"))
}

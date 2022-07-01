package text

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {
	m := map[string]interface{}{}
	m["S"] = "string"
	m["I"] = 1
	m["F"] = 1.2
	m["B"] = true
	m["A"] = []int{1, 2}
	m["M"] = map[int]int{1: 11, 2: 22}

	actual := ToString(m)
	fmt.Println(actual)
	assert.Contains(t, actual, `"S":"string"`)
	assert.Contains(t, actual, `"I":1`)
	assert.Contains(t, actual, `"F":1.2`)
	assert.Contains(t, actual, `"B":true`)
	assert.Contains(t, actual, `"A":[1,2]`)
	assert.Contains(t, actual, `"M":{"1":11,"2":22}`)

	type Sample struct {
		S string
		I int
		F float64
		B bool
		A []int
		M map[int]int
	}
	s := Sample{
		S: "string",
		I: 1,
		F: 1.2,
		B: true,
		A: []int{1, 2},
		M: map[int]int{1: 11, 2: 22},
	}
	actual = ToString(s)
	fmt.Println(actual)
	assert.Contains(t, actual, `"S":"string"`)
	assert.Contains(t, actual, `"I":1`)
	assert.Contains(t, actual, `"F":1.2`)
	assert.Contains(t, actual, `"B":true`)
	assert.Contains(t, actual, `"A":[1,2]`)
	assert.Contains(t, actual, `"M":{"1":11,"2":22}`)
}

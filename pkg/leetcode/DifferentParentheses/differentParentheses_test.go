package diffways

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testOne = "test one leetCode task"
	testTwo = "test two leetCode task"
)

func TestDiffWaysToCompute(t *testing.T) {
	var (
		str1 string = "2-1-1"
		str2 string = "2*3-4*5"
	)
	obj := NewHandler()
	t.Run(testOne, func(t *testing.T) {
		want := []int{2, 0}
		got := obj.DiffWaysToCompute(str1)
		if !assert.EqualValues(t, got, want) {
			t.Errorf("error test one want %v got %v", want, got)
		}
	})

	t.Run(testTwo, func(t *testing.T) {
		want := []int{-34, -10, -14, -10, 10}
		got := obj.DiffWaysToCompute(str2)
		if !assert.EqualValues(t, got, want) {
			t.Errorf("error test two want %v, got %v", want, got)
		}
	})
}

package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testVisitorOne = "Test visitor one test"
)

func TestVisitor(t *testing.T) {
	t.Run(testVisitorOne, func(t *testing.T) {
		x := NewVisitor()
		want := []string{"Point A", "Point B", "Point C"}
		points := []Pointer{NewPointA(), NewPointB(), NewPointC()}
		for i, p := range points {
			got := p.Accept(x)
			if !assert.EqualValues(t, want[i], got) {
				t.Errorf("error test visitor want %v got %v", want[i], got)
			}
		}
	})
}

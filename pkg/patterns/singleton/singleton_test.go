package singelton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testSingleton = "test  Singleton"
)

func TestSingleton(t *testing.T) {
	t.Run(testSingleton, func(t *testing.T) {
		wg := sync.WaitGroup{}
		x := NewSingleton()
		for i := 0; i < 100; i++ {
			x = NewSingleton()
			x.Add()
			x.Done()
			wg.Add(1)
			go func(x Singleton, wg *sync.WaitGroup) {
				defer wg.Done()
				y := NewSingleton()
				y.Add()
				if !assert.EqualValues(t, x, y, "") {
					t.Errorf("got %p want  %p", x, y)
				}
			}(x, &wg)
		}
	})
}

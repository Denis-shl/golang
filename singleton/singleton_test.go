package Singelton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	t.Run("test  Singleton", func(t *testing.T) {
		wg := sync.WaitGroup{}
		x := GetInstance()
		for i := 0; i < 100; i++ {
			x = GetInstance()
			x.Add()
			x.Done()
			wg.Add(1)
			go func(x Singleton, wg *sync.WaitGroup) {
				defer wg.Done()
				y := GetInstance()
				y.Add()
				if !assert.EqualValues(t, x,y, "") {
					t.Errorf("got %p want  %p", x, y)
				}
			}(x, &wg)
		}
	})
}

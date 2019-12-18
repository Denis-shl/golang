package singleton

import (
	"fmt"
	"testing"
)

func TestSingelton(t *testing.T) {
	for i := 0; i < 2; i++ {
		x := GetInstance()
		x.Add()
		x.Done()
		go func(x Singleton) {
			y := GetInstance()
			y.Add()
			if &x != &y {
				t.Errorf("got %p want  %p", &x, &y)
			}
		}(x)
		fmt.Println(i)
	}
}

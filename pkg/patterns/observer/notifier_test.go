package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testNewObserve = "Test new observeA and observerB"
	testObserve    = "Test observe"
)

func TestObserve(t *testing.T) {
	notify := NewConcreteNotify()
	observeA := NewObserveA()
	observeB := NewObserveB()
	t.Run(testNewObserve, func(t *testing.T) {
		if !assert.NotNil(t, notify, observeA, observeB) {
			t.Errorf("Error new notify %v, obsA %v, obsB %v", notify, observeA, observeB)
		}
	})
	t.Run(testObserve, func(t *testing.T) {
		notify.Register(observeA)
		notify.Register(observeB)
		notify.Notify(12)
		want := 12
		gotA := observeA.GetData()
		gotB := observeB.GetData()
		if !assert.EqualValues(t, want, gotA, gotB) {
			t.Errorf("error test observe want %v gotA %v gotB %v", want, gotA, gotB)
		}
	})
}

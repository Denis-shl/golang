package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	newMediatorAndSender = "Test func new mediator and Sender"
	testFuncSenderA = "Test mediator and func senderA"
	testFuncSenderB = "Test mediator and func senderB"
	testFuncNotSender = "Test mediator and func not sender"
)
func TestMediator(t *testing.T) {
	mediator := NewMediator()
	senderA := NewSenderA()
	senderB := NewSenderB()
	senderA.SetMediator(mediator)
	senderB.SetMediator(mediator)
	mediator.NewConcrete(senderA, senderB)

	t.Run(newMediatorAndSender, func(t *testing.T) {
		if !assert.NotNil(t, senderA, senderB, mediator){
			t.Errorf("Test func mediator and Sender error  nil %v or nil %v or nil %v", mediator, senderA, senderB)
		}
	})
	t.Run(testFuncSenderA, func (t *testing.T){
		wantOne := "i senderB"
		gotOne := senderA.Send("Hello who are you")
		if !assert.EqualValues(t, gotOne, wantOne){
			t.Errorf("Test mediator error gotA %v, wantA %v", gotOne, wantOne)
		}
	})
	t.Run(testFuncSenderB, func (t *testing.T){
		wantTwo := "i senderA"
		gotTwo := senderB.Send("hello")
		if !assert.EqualValues(t, wantTwo, gotTwo){
			t.Errorf("Test mediator error gotA %v, wantA %v", gotTwo, wantTwo)
		}
	})
	t.Run(testFuncNotSender, func (t *testing.T){
		wantSenderNot := "not sender"
		gotSenderNot := senderB.Send("Hi")
		if !assert.EqualValues(t, wantSenderNot, gotSenderNot){
			t.Errorf("Test mediator error gotA %v, wantA %v", gotSenderNot, wantSenderNot)
		}
	})
}

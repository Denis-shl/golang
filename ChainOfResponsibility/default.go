package chainOfRes

import "fmt"

type Default struct {
}

func (p *Default) Treat(name string) bool {
	fmt.Println("couldn't find a doctor")
	return false
}

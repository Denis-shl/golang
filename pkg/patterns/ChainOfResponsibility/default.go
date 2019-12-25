package chainOfRes

import "fmt"

// Default ...
type Default struct {
}

// Treat ...
func (p *Default) Treat(name string) bool {
	fmt.Println("couldn't find a doctor")
	return false
}

package server

type (
	response = string
	request  = string
)

// RealSubject ...
type Server interface {
	Request(str string) response
}

// server implements a real subject
type server struct {
}

// SetName put a name in a real object
func (r *server) Request(str request) response {
	if str == "nginx" {
		return "open code"
	}
	return "closed source code"
}

// NewServer ...
func NewServer() Server {
	return &server{}
}

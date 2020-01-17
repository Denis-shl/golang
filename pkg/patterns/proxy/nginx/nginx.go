package nginx

type (
	response = string
	request  = string
)

// Nginx ...
type Nginx interface {
	Request(request) response
}

// server implements a real subject
type server struct {
}

// SetName put a name in a real object
func (r *server) Request(_ request) response {
	return "nginx open source code"
}

// NewServer ...
func NewServer() Nginx {
	return &server{}
}

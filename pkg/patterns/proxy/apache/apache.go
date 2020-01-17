package apache

type (
	request  = string
	response = string
)

// Apache ...
type Apache interface {
	Request(request) response
}

type server struct {
}

func (s *server) Request(_ request) response {
	return "apache open source code"
}

// NewServer
func NewServer() Apache {
	return &server{}
}

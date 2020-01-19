package handlerb

type (
	response = string
	request  = string
)

type server interface {
	Request(request) response
}

// Handler ...
type Handler interface {
	ProcessRequest(response) request
}

type handler struct {
	name   string
	server server
	next   Handler
}

// Treat ...
func (h *handler) ProcessRequest(str response) request {
	if str == "apache" {
		return h.server.Request(str)
	}
	return h.next.ProcessRequest(str)
}

// NewHandlerA ...
func NewHandlerB(next Handler, server server) Handler {
	return &handler{next: next, server: server}
}

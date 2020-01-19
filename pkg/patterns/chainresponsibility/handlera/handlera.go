package handlera

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

func (h *handler) ProcessRequest(str response) request {
	if str == "nginx" {
		return h.server.Request(str)
	}
	return h.next.ProcessRequest(str)
}

// NewHandlerA ...
func NewHandlerA(next Handler, server server) Handler {
	return &handler{next: next, server: server}
}

package lasthandler

type (
	response = string
	request  = string
)

type server interface {
	Request(request) response
}

// LastHandler ...
type LastHandler interface {
	ProcessRequest(response) request
}

type lastHandler struct {
	name         string
	serverApache server
	serverNginx  server
}

func (l *lastHandler) ProcessRequest(response) request {
	return "400 Bad Request"
}

// NewLastHandler ...
func NewLastHandler(serverApache server, serverNginx server) LastHandler {
	return &lastHandler{
		serverApache: serverApache,
		serverNginx:  serverNginx,
	}
}

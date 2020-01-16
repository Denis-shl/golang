package proxy

type (
	queryName      = string
	requestHistory = []string
	response       = string
)

// Proxy ...
type Proxy interface {
	Request(str string) response
	GetRequestHistory() requestHistory
}

type server interface {
	Request(str string) response
}

type proxy struct {
	server         server
	requestHistory []queryName
}

// SetName put a name in a  proxy object
func (p *proxy) Request(str queryName) response {
	p.requestHistory = append(p.requestHistory, str)
	return p.server.Request(str)
}

// GetName getting proxy name
func (p *proxy) GetRequestHistory() requestHistory {
	return p.requestHistory
}

// NewProxy ...
func NewProxy(server server) Proxy {
	return &proxy{server: server}
}

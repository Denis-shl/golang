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

// Request ...
func (p *proxy) Request(str queryName) response {
	p.saveHistoryRequest(str)
	return p.server.Request(str)
}

// GetRequestHistory ...
func (p *proxy) GetRequestHistory() requestHistory {
	return p.requestHistory
}

func (p *proxy) saveHistoryRequest(str queryName ) {
	p.requestHistory = append(p.requestHistory, str)
}

// NewProxy ...
func NewProxy(server server) Proxy {
	return &proxy{server: server}
}

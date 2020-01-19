package proxy

import "strings"

type (
	requestHistory = []string
	response       = string
	request        = string
)

type handlers interface {

}

type server interface {
	Request(request) response
}

// Proxy ...
type Proxy interface {
	Request(str request) response
	GetRequestHistory() requestHistory
}

type proxy struct {
	apache         server
	nginx          server
	requestHistory []request
	chainHandlers		handler
}



// Request ...
func (p *proxy) Request(str request) (response response) {
	p.requestHistory = append(p.requestHistory, str)

	if strings.Contains(str, "nginx") == true && strings.Contains(str, "apache") == true {
		return p.nginx.Request(str) + "," + p.apache.Request(str)
	}
	if strings.Contains(str, "nginx") == true {
		return p.nginx.Request(str)
	}
	if strings.Contains(str, "apache") == true {
		return p.apache.Request(str)
	}
	return "400"
}

// GetRequestHistory ...
func (p *proxy) GetRequestHistory() requestHistory {
	return p.requestHistory
}

// NewProxy ...
func NewProxy(apache, nginx server) Proxy {
	return &proxy{apache: apache, nginx: nginx}
}

package client

type target interface {
	Request()
}

// Client ...
type Client interface {
	Request(target target)
}

type client struct {
}

func (c *client) Request(target target) {
	target.Request()
}

// NewClient ...
func NewClient() Client {
	return &client{}
}

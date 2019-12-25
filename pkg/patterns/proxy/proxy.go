package proxy

// Subject provides an interface for a real subject and its surrogate.
type Subject interface {
	putName(string)
	getName() string
}

// Proxy ...
type Proxy struct {
	realSubject Subject
}

// PutName put a name in a  proxy object
func (p *Proxy) PutName(x string) {
	p.getRealObj()
	p.realSubject.putName(x)
}

// GetName getting proxy name
func (p *Proxy) GetName() string {
	p.getRealObj()
	return p.realSubject.getName()
}

// getRealObj getting a real object
func (p *Proxy) getRealObj() {
	if p.realSubject == nil {
		p.realSubject = &realSubject{}
	}
}

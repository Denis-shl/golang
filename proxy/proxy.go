package proxy

// Subject provides an interface for a real subject and its surrogate.
type Subject interface {
	PutName(string)
	GetName() string
}

// Proxy ...
type Proxy struct {
	realSubject Subject
}

// PutName put a name in a  proxy object
func (p *Proxy) PutName(x string) {
	p.getRealObj()
	p.realSubject.PutName(x)
}

// GetName getting proxy name
func (p *Proxy) GetName() string {
	p.getRealObj()
	return p.realSubject.GetName()
}

// RealSubject implements a real subject
type realSubject struct {
	name string
}

// SetName put a name in a real object
func (r *realSubject) PutName(str string) {
	r.name = str
}

// GetName getting name
func (r *realSubject) GetName() string {
	return r.name
}

// getRealObj getting a real object
func (p *Proxy) getRealObj() {
	if p.realSubject == nil {
		p.realSubject = &realSubject{}
	}
}

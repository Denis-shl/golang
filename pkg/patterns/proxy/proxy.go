package proxy

// Proxer ...
type Proxer interface {
	SetName(str string)
	GetName() string
}

type proxy struct {
	realSubject Objective
}

// SetName put a name in a  proxy object
func (p *proxy) SetName(str string) {
	p.getRealObj()
	p.realSubject.SetName(str)
}

// GetName getting proxy name
func (p *proxy) GetName() string {
	p.getRealObj()
	return p.realSubject.GetName()
}

// getRealObj getting a real object
func (p *proxy) getRealObj() {
	if p.realSubject == nil {
		p.realSubject = NewObjective()
	}
}

// NewProxer ...
func NewProxer() Proxer {
	return &proxy{}
}

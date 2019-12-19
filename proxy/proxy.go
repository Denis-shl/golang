package proxy

// Subject provides an interface for a real subject and its surrogate.
type Subject interface {
	PutName(string)
	GetName() string
}

type Proxy struct {
	realSubject Subject
}

// PutName put a name in a  proxy object
func (s *Proxy) PutName(x string) {
	s.getRealObj()
	s.realSubject.PutName(x)
}

// GetName getting proxy name
func (s *Proxy) GetName() string {
	s.getRealObj()
	return s.realSubject.GetName()
}

// RealSubject implements a real subject
type RealSubject struct {
	name string
}

// SetName put a name in a real object
func (s *RealSubject) PutName(x string) {
	s.name = x
}

// GetName getting name
func (s *RealSubject) GetName() string {
	return s.name
}

// getRealObj getting a real object
func (s *Proxy) getRealObj() {
	if s.realSubject == nil {
		s.realSubject = &RealSubject{}
	}
}

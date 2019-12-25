package proxy

// RealSubject implements a real subject
type realSubject struct {
	name string
}

// putName put a name in a real object
func (r *realSubject) putName(str string) {
	r.name = str
}

// getName getting name
func (r *realSubject) getName() string {
	return r.name
}

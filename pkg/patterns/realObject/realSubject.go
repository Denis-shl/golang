package object

// Objective ...
type Objective interface {
	SetName(str string)
	GetName() string
}

// realSubject implements a real subject
type realSubject struct {
	name string
}

// putName put a name in a real object
func (r *realSubject) SetName(str string) {
	r.name = str
}

// getName getting name
func (r *realSubject) GetName() string {
	return r.name
}

// NewObjective ...
func NewObjective() Objective {
	return &realSubject{}
}

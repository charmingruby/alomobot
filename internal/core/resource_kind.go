package core

type ResourceKind string

const (
	TypeResouceKind ResourceKind = "type"
	CupResouceKind  ResourceKind = "cup"
)

func (k ResourceKind) Valid() bool {
	switch k {
	case TypeResouceKind, CupResouceKind:
		return true
	default:
		return false
	}
}

func (k ResourceKind) String() string {
	return string(k)
}

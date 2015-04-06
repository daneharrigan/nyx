package nameserver

type Nameserver interface {
}

func New() Nameserver {
	return new(nameserver)
}

type nameserver struct{}

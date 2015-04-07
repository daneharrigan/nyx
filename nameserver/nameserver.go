package nameserver

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("record not found")
)

type Protocol int

const (
	HTTP Protocol = iota
	HTTPS
)

type Nameserver interface {
	Lookup(string) (*Record, error)
	Add(string, *Node)
}

type Record struct {
	Nodes     []*Node
	Paths     map[string][]*Node
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Node struct {
	Host      string
	Port      int
	Protocol  Protocol
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New() Nameserver {
	return &nameserver{
		records: make(map[string]*Record),
	}
}

type nameserver struct {
	records map[string]*Record
}

func (ns *nameserver) Lookup(host string) (*Record, error) {
	if record, ok := ns.records[host]; ok {
		return record, nil
	}

	return nil, ErrNotFound
}

func (ns *nameserver) Add(host string, n *Node) {
	if _, ok := ns.records[host]; !ok {
		ns.records[host] = &Record{
			Paths:     make(map[string][]*Node),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	ns.records[host].Nodes = append(ns.records[host].Nodes, n)
}

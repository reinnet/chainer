package chain

import (
	"math/rand"

	"github.com/reinnet/chainer/vnf"
)

// Node represents a single VNF in chain
type Node struct {
	Type string
}

// Link connects two node on a chain
type Link struct {
	Source      int
	Destination int
	Bandwidth   int
}

// Chain is a linear form of VNFs
type Chain struct {
	Cost  int
	Nodes []Node
	Links []Link
}

// Bandwidth chain's links bandwidth
const Bandwidth = 250

// LengthLB is a inclusive chain's length lower bound
const LengthLB = 4

// LengthUB is a exclusive chain's length upper bound
const LengthUB = 7

// Cost is a per instance cost
const Cost = 100

// New creates a chain
func New() Chain {
	types := vnf.Types()

	// chain length
	l := rand.Intn(LengthUB-LengthLB) + LengthLB

	c := Chain{
		Cost:  l * Cost,
		Nodes: make([]Node, 0),
		Links: make([]Link, 0),
	}

	// ingress node
	c.Nodes = append(c.Nodes, Node{
		Type: "ingress",
	})

	for j := 1; j < l; j++ {
		// place VNF types into chain
		t := types[rand.Intn(len(types))]

		c.Nodes = append(c.Nodes, Node{
			Type: t.Name,
		})

		// place links between VNFs
		c.Links = append(c.Links, Link{
			Source:      j - 1,
			Destination: j,
			Bandwidth:   Bandwidth,
		})
	}

	// egress node
	c.Nodes = append(c.Nodes, Node{
		Type: "egress",
	})

	c.Links = append(c.Links, Link{
		Source:      l - 1,
		Destination: l,
		Bandwidth:   Bandwidth,
	})

	return c
}

/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 09-07-2019
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/roadtomsc/chainer/vnf"
	"gopkg.in/yaml.v2"
)

// Node represents a single VNF in chain
type Node struct {
	Type string
	ID   string
}

// Link connects two node on a chain
type Link struct {
	Source      string
	Destination string
	Bandwidth   int
}

// Chain is a linear form of VNFs
type Chain struct {
	Cost  int
	Nodes []Node
	Links []Link
}

// Config is a generated result of chainer
type Config struct {
	Chains []Chain
}

func main() {
	var number = flag.Int("n", 100, "number of chains")
	flag.Parse()

	var cs []Chain

	for i := 0; i < *number; i++ {
		bw := math.MinInt32

		c := Chain{
			Cost:  rand.Intn(100) + 100,
			Nodes: make([]Node, 0),
			Links: make([]Link, 0),
		}

		l := rand.Intn(3) + 4 // chain length

		// ingress node
		c.Nodes = append(c.Nodes, Node{
			Type: "ingress",
			ID:   "0",
		})

		// place VNF types into chain
		for j := 1; j < l; j++ {
			t := vnf.Get(rand.Intn(vnf.Len()))
			if t.Capacity < bw {
				bw = t.Capacity
			}
			c.Nodes = append(c.Nodes, Node{
				Type: t.Name,
				ID:   strconv.Itoa(j),
			})
		}

		// place links between VNFs
		for j := 1; j < l; j++ {
			c.Links = append(c.Links, Link{
				Source:      strconv.Itoa(j - 1),
				Destination: strconv.Itoa(j),
				Bandwidth:   bw,
			})
		}

		// egress node
		c.Nodes = append(c.Nodes, Node{
			Type: "egress",
			ID:   strconv.Itoa(l),
		})

		c.Links = append(c.Links, Link{
			Source:      strconv.Itoa(l - 1),
			Destination: strconv.Itoa(l),
			Bandwidth:   250,
		})

		cs = append(cs, c)
	}

	b, err := yaml.Marshal(Config{
		Chains: cs,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := f.Write(b); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.Close(); err != nil {
		return
	}
	log.Println(string(b))
}

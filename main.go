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
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"

	"github.com/roadtomsc/chainer/vnf"
	"gopkg.in/yaml.v2"
)

type Node struct {
	Type string
	ID   string
}

type Link struct {
	Source      string
	Destination string
	Bandwidth   int
}

type Chain struct {
	Cost  int
	Nodes []Node
	Links []Link
}

type Config struct {
	Chains []Chain
}

func main() {
	var cs []Chain

	for i := 0; i < 100; i++ {
		bw := math.MinInt32

		c := Chain{
			Cost:  rand.Intn(100) + 100,
			Nodes: make([]Node, 0),
			Links: make([]Link, 0),
		}

		l := rand.Intn(3) + 4
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

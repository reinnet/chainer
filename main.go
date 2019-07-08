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
	"math/rand"
	"os"
	"strconv"

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

func VNFType(i int) string {
	return []string{
		"vIDS",
		"vNAT",
		"vFW",
		"vDPI",
	}[i]
}

func main() {
	var cs []Chain

	for i := 0; i < 100; i++ {
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

		for j := 1; j < l; j++ {
			c.Nodes = append(c.Nodes, Node{
				Type: VNFType(rand.Intn(4)),
				ID:   strconv.Itoa(j),
			})

			c.Links = append(c.Links, Link{
				Source:      strconv.Itoa(j - 1),
				Destination: strconv.Itoa(j),
				Bandwidth:   250,
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

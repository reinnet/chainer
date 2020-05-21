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
	"math/rand"
	"os"
	"time"

	"github.com/reinnet/chainer/chain"

	"gopkg.in/yaml.v3"
)

// Config is a generated result of chainer
type Config struct {
	Chains []chain.Chain
}

func main() {
	rand.Seed(time.Now().Unix())

	var number = flag.Int("n", 100, "number of chains")

	flag.Parse()

	var cs []chain.Chain

	for i := 0; i < *number; i++ {
		cs = append(cs, chain.New())
	}

	if err := store(cs, "chains.yaml"); err != nil {
		fmt.Println(err)
		return
	}
}

// store write the configuration into given YAML file
func store(cs []chain.Chain, name string) error {
	b, err := yaml.Marshal(Config{
		Chains: cs,
	})
	if err != nil {
		return err
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

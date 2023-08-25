package uniqid

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(0)
	if err != nil {
		log.Fatalf("uniqid: unable to create new node: %s", err)
	}
	node = n
}

func Generate() int64 {
	return int64(node.Generate())
}

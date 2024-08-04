package snowflake

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Failed to initialize snowflake node: %v", err)
	}
}

func GenerateID() uint64 {
	return uint64(node.Generate().Int64())
}

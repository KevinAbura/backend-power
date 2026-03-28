package main

import (
	"fmt"
	"os"

	"github.com/KevinAbura/backend-power/internal/api/router"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: backend-power <command>")
		os.Exit(1)
	}

	command, ok := router.Route(os.Args[1])
	if !ok {
		fmt.Printf("unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Printf("routed command: %s\n", command)
}

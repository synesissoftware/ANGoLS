package main

import (
	angols "github.com/synesissoftware/ANGoLS"
	"github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("angols v%s\n", angols.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}

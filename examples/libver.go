package main

import (
	angols "github.com/synesissoftware/ANGoLS"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("angols v%s\n", ver2go.CalcVersionString(angols.VersionMajor, angols.VersionMinor, angols.VersionPatch, angols.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}

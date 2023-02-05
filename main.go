package main

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok-rest-template/internal"
	"github.com/gowok/gowok-rest-template/pkg"
	"github.com/gowok/gowok-rest-template/rest"
)

func init() {
	pkg.PrepareAll()
	internal.PrepareAll()
	rest.PrepareAll()
}

func main() {
	rest.Run()

	go gowok.StartPProf()
	gowok.GracefulStop(func() {
		fmt.Println()
		fmt.Println("Stopping...")
	})
}

package main

import (
	"github.com/golang-source-engine/vmt"
	"log"
	"os"
)

func main() {
	stream,err := os.Open("./sample.vmt")
	if err != nil {
		panic(err)
	}

	var mat vmt.Material
	mat = vmt.NewProperties()
	mat,err = vmt.FromStream(stream, mat)
	if err != nil {
		panic(err)
	}
	log.Println(mat)
}

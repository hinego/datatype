package main

import (
	"github.com/hinego/datatype"
	"log"
)

func main() {
	dd := datatype.NewInt64(100)
	log.Println(dd.Load())
	dd.Store(50)
	log.Println(dd.Load())
	log.Println(dd.Swap(60))
	log.Println(dd.Load())
	log.Println(dd.Add(1))
	log.Println(dd.Load())
}

package main

import (
	"log"

	"github.com/Pointdexter37/kin/cmd"
)

func main(){
	if err := cmd.Execute(); err != nil{
		log.Fatal(err)
	}
}
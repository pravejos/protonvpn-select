package main

import (
	"log"
)

func debug(v ...interface{}) {
	if *verbose {
		log.Println(v...)
	}
}

func fatal(v ...interface{}) {
	log.Fatal(v...)
}

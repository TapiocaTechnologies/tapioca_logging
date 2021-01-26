package main

import (
	"fmt"
	"log"
	"os"
)
func Logging(message string){

	f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
	fmt.Println(message)
}

func main(){
	Logging("An error has occured")
}
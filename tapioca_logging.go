package tapioca_logging

import (
	"fmt"
	"log"
	"os"
	
)
var file_name = ""
func SetLogFile(file string){

	file_name = file
	
}
func Logging(message string){
	
	f, err := os.OpenFile(file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("[GENERAL] %s", message)
	fmt.Println(message)
}

func LoggingError(e error) {
	f, err := os.OpenFile(file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("[ERROR] %s", e)
	fmt.Println(e)
}

func LoggingNetwork(message string){
	
	f, err := os.OpenFile(file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	
	log.SetOutput(f)
	log.Printf("[NETWORK] %s", message)
	fmt.Println(message)
}

func LoggingNetworkError(e error){
	
	f, err := os.OpenFile(file_name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	
	log.SetOutput(f)
	log.Printf("[NETWORK] [Error] %s", e)
	fmt.Println(e)
}


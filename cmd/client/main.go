package main

import "log"

var config = map[string]interface{}{
	"server": "localhost",
	"port":   24888,
	"debug":  true,
}

func main() {
	log.Println("Starting Client...")
}

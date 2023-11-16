package main

import (
	"flag"

	"github.com/yousefzinsazk78/fiber_post_second_version/internal/routes"
)

func main() {
	///set port in flag of program
	port := flag.String("default port", ":8000", "you can set your custom font")
	flag.Parse()

	//todo : implement custom route
	routes.Run(*port)

}

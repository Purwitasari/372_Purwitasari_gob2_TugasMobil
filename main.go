package main

import "tugas1/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}

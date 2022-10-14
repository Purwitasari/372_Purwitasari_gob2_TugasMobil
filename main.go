package main

import "sesi-06-tugas/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}

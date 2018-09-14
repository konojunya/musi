package main

import "github.com/konojunya/musi/router"

func main() {
	r := router.GetRouter()
	r.Run(":3000")
}

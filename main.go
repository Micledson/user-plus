package main

import "user-plus/tests/endpoints/routes"

func main() {
	r := routes.LoadRoutes()
	r.Logger.Fatal(r.Start(":8081"))
}

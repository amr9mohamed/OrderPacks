package main

import (
	"orderPacks/cmd/http"
	"orderPacks/pkg/orderPacks/repository"
	"orderPacks/pkg/orderPacks/service"
)

func main() {
	r := repository.NewMemDB([]int{250, 500, 1000, 2000, 5000})
	s := service.NewPacksManager(r)
	server := http.NewServer(s)
	engine := server.InitServer()
	if err := engine.Run(":8080"); err != nil {
		return
	}
}

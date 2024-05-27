package main

import (
	"fmt"
	"pagi/config"
	"pagi/handler"
	"pagi/repo"
)

func main() {
	db := config.Connect()

	GameStoreRepo := &repo.MysqlRepo{DB: db}
	GameStoreHandler := &handler.GameStoreHandler{Repo: GameStoreRepo}
	router, server := config.InitServer()
	router.GET("/branches", GameStoreHandler.GetBranches)
	router.GET("/branch/:id", GameStoreHandler.GetBranchBYId)
	router.POST("/branch", GameStoreHandler.AddBranch)
	router.PUT("/branch/:id", GameStoreHandler.UpdateBranch)
	router.DELETE("/branch/:id", GameStoreHandler.DeleteBranch)

	fmt.Println("server running on localhost:8080")
	panic(server.ListenAndServe())
}

package main

import (
	"app/main/repository"
	"app/main/route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベース
	repository.Init()
	defer repository.Close()

	// ルーティング
	route.Init()
}

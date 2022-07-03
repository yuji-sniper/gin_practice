package main

import (
	"app/main/handler"
	"app/main/infra"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベース
	infra.Init()
	defer infra.Close()
	// ルーティング
	handler.Init()
}

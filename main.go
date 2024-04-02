package main

import (
	"myapp/db"
	"myapp/route"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	//dbInit()
	db.DbInit()

	//ハンドラ関数の割り当て部分
	//Index
	//最初のページ(ルート(/))にアクセスしたときにGETリクエストが自動でなされる
	router.GET("/", route.Index)
	//Create
	router.POST("/new", route.Create)
	//Detail
	router.GET("/detail/:id", route.Detail)
	//Update
	router.POST("/update/:id", route.Update)
	//削除確認
	router.GET("/delete_check/:id", route.DelConf)
	//Delete
	router.POST("/delete/:id", route.Delete)

	//デフォルトのHTTPポート（:8080）でWebサーバを起動
	router.Run()

}

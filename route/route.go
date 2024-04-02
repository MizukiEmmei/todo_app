package route

import (
	"myapp/db"
	"strconv"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {//ctx変数にリクエスト情報を格納
	todos := db.DbGetAll()
	//テンプレートhtmlをレンダリングして表示させる
	ctx.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
	//ctx.HTMLの引数
	// 1. HTTPステータスコード（この場合は200、つまり成功）
	// 2. レンダリングするHTMLテンプレートの名前（"index.html"）
	// 3. テンプレートに渡すデータ（この例では、"todos"キーにToDoアイテムのリストを含むマップ）
}

func Create(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	db.DbInsert(text, status)
	ctx.Redirect(302, "/") //ルートページにリダイレクト
}

func Detail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n) //文字列を数値に変換
	if err != nil {
		panic(err)
	}
	todo := db.DbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
}

func Update(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	db.DbUpdate(id, text, status)
	ctx.Redirect(302, "/")
}

func DelConf(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := db.DbGetOne(id)
	ctx.HTML(200, "delete.html", gin.H{"todo": todo})
}

func Delete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	db.DbDelete(id)
	ctx.Redirect(302, "/")

}

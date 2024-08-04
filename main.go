package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	temp := template.Must(template.ParseFiles("index.html"))
	// POST 요청 처리
	r.POST("/", func(ctx *gin.Context) {
		// 폼 데이터 수신
		name := ctx.PostForm("name")
		content := ctx.PostForm("content")
		imgUrl := ctx.PostForm("img")

		// 익명 구조체를 사용하여 데이터를 저장
		responseData := struct {
			Name    string `json:"name"`
			Content string `json:"content"`
			ImgURL  string `json:"imgUrl"`
		}{
			Name:    name,
			Content: content,
			ImgURL:  imgUrl,
		}
		temp.Execute(ctx.Writer, responseData)
	})

	// 서버 실행
	r.Run(":8080") // 기본 포트는 8080입니다.
}

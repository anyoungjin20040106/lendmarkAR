package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageData struct {
	Name string
}

func main() {
	r := gin.Default()
	r.Static("/Build", "./Build")

	// TemplateData 디렉토리 마운트
	r.Static("/TemplateData", "./TemplateData")
	// HTML 템플릿 파일 로드
	r.LoadHTMLFiles("index.html")

	// POST 요청 처리
	r.GET("/", func(ctx *gin.Context) {

		// 익명 구조체를 사용하여 데이터를 저장
		responseData := struct {
			Name string `json:"name"`
		}{
			Name: "수종사",
		}
		ctx.HTML(http.StatusOK, "index.html", responseData)
	})
	// POST 요청 처리
	r.POST("/", func(ctx *gin.Context) {
		// 예시 폼 데이터 수신
		name := ctx.PostForm("name")

		// 템플릿에 전달할 데이터 구조체
		pageData := PageData{
			Name: name,
		}

		// HTML 템플릿 렌더링
		ctx.HTML(http.StatusOK, "index.html", pageData)
	})

	// 서버 실행
	r.Run(":18012")
}

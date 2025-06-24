package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JoaoRafa19/portifolio/assets"
	"github.com/JoaoRafa19/portifolio/internal/utils"
	"github.com/JoaoRafa19/portifolio/ui/components/dialog"
	"github.com/JoaoRafa19/portifolio/ui/pages"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	InitDotEnv()
	mux := gin.New()
	mux.Use(gin.Recovery())
	mux.Use(utils.StdLogger(gin.LoggerWithConfig(gin.LoggerConfig{})))
	mux.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s %s %d %s %s\n",
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				param.ClientIP,
				param.Method,
				param.StatusCode,
				param.Latency,
			)
		}}))

	SetupAssetsRoutes(mux)
	mux.GET("/", func(ctx *gin.Context) {
		pages.HomePage(false).Render(ctx.Request.Context(), ctx.Writer)
	})
	mux.GET("/contact/dialog", func(ctx *gin.Context) {
		showDialog := ctx.Query("show") == "true"
		dialog.ShowDialogForm(showDialog).Render(ctx.Request.Context(), ctx.Writer)
	})

	mux.GET("/teste", func(ctx *gin.Context) {

		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "index.html")
		}).ServeHTTP(ctx.Writer, ctx.Request)
	})

	fmt.Println("Server is running on http://localhost:8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Server stopped")
}

func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func SetupAssetsRoutes(mux *gin.Engine) {
	var isDevelopment = os.Getenv("GO_ENV") != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.GET("/assets/*filepath", gin.WrapH(http.StripPrefix("/assets/", assetHandler)))
}

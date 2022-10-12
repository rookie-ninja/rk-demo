package main

import (
	"context"
	_ "embed"
	"github.com/gin-gonic/gin"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
	"net/http"
)

//go:embed serverB.yaml
var rawBootB []byte

func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(rawBootB))

	// Register handler
	entry := rkgin.GetGinEntry("serverB")
	entry.Router.GET("/v1/second", Second)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

func Second(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &SecondResponse{
		Message: "Second response",
	})
}

type SecondResponse struct {
	Message string
}

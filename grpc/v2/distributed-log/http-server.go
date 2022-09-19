package main

import (
	"context"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-gin/v2/boot"
	"github.com/rookie-ninja/rk-gin/v2/middleware/context"
	"github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

//go:embed http-server.yaml
var httpBoot []byte

func main() {
	boot := rkboot.NewBoot(rkboot.WithBootConfigRaw(httpBoot))

	// register gin
	ginEntry := rkgin.GetGinEntry("httpServer")
	ginEntry.Router.GET("/v1/hello", func(ctx *gin.Context) {
		// create grpc client
		opts := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithInsecure(),
		}
		// create connection with grpc-serverA
		connA, _ := grpc.Dial("localhost:2008", opts...)
		defer connA.Close()
		clientA := greeter.NewServerAClient(connA)
		// create connection with grpc-serverB
		connB, _ := grpc.Dial("localhost:2022", opts...)
		defer connA.Close()
		clientB := greeter.NewServerBClient(connB)

		// eject span context from gin context and inject into grpc ctx
		grpcCtx := trace.ContextWithRemoteSpanContext(context.Background(), rkginctx.GetTraceSpan(ctx).SpanContext())
		md := metadata.Pairs()
		rkginctx.GetTracerPropagator(ctx).Inject(grpcCtx, &rkgrpcctx.GrpcMetadataCarrier{Md: &md})
		grpcCtx = metadata.NewOutgoingContext(grpcCtx, md)

		// call gRPC server A
		clientA.Login(grpcCtx, &greeter.LoginReq{})
		clientA.CallA(grpcCtx, &greeter.CallAReq{})

		// call gRPC server B
		clientB.CallB(grpcCtx, &greeter.CallBReq{})

		ctx.JSON(http.StatusOK, map[string]string{
			"traceId": rkginctx.GetTraceId(ctx),
		})
	})

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
}

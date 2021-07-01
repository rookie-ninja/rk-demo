package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestGreeter_HappyCase(t *testing.T) {
	ctx, _ := gin.CreateTestContext(&mockResponseWriter{})
	ctx.Params = []gin.Param{
		{
			Key:   "name",
			Value: "ut",
		},
	}

	Greeter(ctx)
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

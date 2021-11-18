package api

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGreeter_HappyCase(t *testing.T) {
	recorder := httptest.NewRecorder()

	ctx := echo.New().NewContext(
		httptest.NewRequest(http.MethodGet, "/ut-path", nil),
		recorder)

	ctx.Request().URL = &url.URL{
		RawQuery: "name=ut",
	}

	// Call echo handler
	assert.Nil(t, Greeter(ctx))

	// Assert expected result
	assert.Contains(t, recorder.Body.String(), "Hello ut!")
}

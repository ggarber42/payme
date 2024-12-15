package http_controllerTest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	http_controller "github.com/ggarber42/payme/internal/infra/input/http_controller"
	"github.com/ggarber42/payme/tests/mock"
)

func TestHealthChecker(t *testing.T) {
	t.Run("return status 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
		response := httptest.NewRecorder()

		cs := mock.NewCardService()
		controller := http_controller.NewHttpController(cs)
		controller.HealthCheckerHandler(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}


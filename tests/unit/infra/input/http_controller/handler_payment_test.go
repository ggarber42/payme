package http_controllerTest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ggarber42/payme/internal/domain/entity"
	vendor_service "github.com/ggarber42/payme/internal/domain/services/mock_vendor"
	http_controller "github.com/ggarber42/payme/internal/infra/input/http_controller"
	"github.com/ggarber42/payme/tests/mock"
)

func TestPaymentHandler(t *testing.T) {
	t.Run("returns status 202 if payload is correct", func(t *testing.T) {

		requestObject := mock.NewRequestObject()

		requestBody, err := json.Marshal(requestObject.ValidPayload())
		if err != nil {
			t.Fatalf("failed to marshal payload: %v", err)
		}

		request, _ := http.NewRequest(http.MethodPost, "/payment/{vendor}", bytes.NewReader(requestBody))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		cs := mock.NewCardService()
		controller := http_controller.NewHttpController(cs)
		controller.PaymentHandler(response, AddChiURLParams(request, map[string]string{entity.VENDOR: entity.STONE}))

		gotCode := response.Code
		wantCode := http.StatusAccepted

		if gotCode != wantCode {
			t.Errorf("got %d, want %d", gotCode, wantCode)
		}

		var got vendor_service.ServiceResponse
		err = json.Unmarshal(response.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}

		want := "payment accepted"
		if got.Message!= want{
			t.Errorf("got %q, want %q", got, want)
		}
		
	})
}

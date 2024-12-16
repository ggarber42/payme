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
		if got.Message != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("return status 400 if payload is incorrect", func(t *testing.T) {
		t.Run("Vendor", func(t *testing.T) {
			t.Run("urlParam", func(t *testing.T) {
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
				controller.PaymentHandler(response, AddChiURLParams(request, map[string]string{entity.VENDOR: "xxx"}))

				gotCode := response.Code
				wantCode := http.StatusBadRequest

				if gotCode != wantCode {
					t.Errorf("got %d, want %d", gotCode, wantCode)
				}

				var got http_controller.ErrorResponse
				err = json.Unmarshal(response.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}

				watStatusText := "Invalid request"
				if got.StatusText != watStatusText {
					t.Errorf("got %q, want %q", got, watStatusText)
				}

				wantErrorText := "vendor values suported are: stone, cielo and rede"
				if got.ErrorText != wantErrorText {
					t.Errorf("got %q, want %q", got, wantErrorText)
				}
			})
			t.Run("payload", func(t *testing.T) {
				requestObject := mock.NewRequestObject()
				requestBody, err := json.Marshal(requestObject.InvalidVendor())
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
				wantCode := http.StatusBadRequest

				if gotCode != wantCode {
					t.Errorf("got %d, want %d", gotCode, wantCode)
				}

				var got vendor_service.ServiceResponse
				err = json.Unmarshal(response.Body.Bytes(), &got)
				if err != nil {
					t.Fatal(err)
				}

			})
		})

		t.Run("CardData", func(t *testing.T) {
			requestObject := mock.NewRequestObject()
			requestBody, err := json.Marshal(requestObject.InvalidCardData())
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
			wantCode := http.StatusBadRequest

			if gotCode != wantCode {
				t.Errorf("got %d, want %d", gotCode, wantCode)
			}

			var got http_controller.ErrorResponse
			err = json.Unmarshal(response.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}

			watStatusText := "Invalid request"
			if got.StatusText != watStatusText {
				t.Errorf("got %q, want %q", got, watStatusText)
			}

			wantErrorText := "{\"invalid_field\":\"token: has an invalid value\",\"missing_field\":\"name: missing required field:\"}"
			if got.ErrorText != wantErrorText {
				t.Errorf("got %q, want %q", got, wantErrorText)
			}
		})
	})

	t.Run("returns 403 if payment is refused", func(t *testing.T) {

		t.Run("stone", func(t *testing.T) {
			requestObject := mock.NewRequestObject()
			requestObject.Payload["purchase"] = map[string]interface{}{
				"purchaseId":  "12345",
				"date":        "2024-11-04T12:30:00Z",
				"customerId":  "cust_67890",
				"totalAmount": 1200.0,
				"currency":    "USD",
				"products": []map[string]interface{}{
					{
						"id":       "12345",
						"name":     "Product 1",
						"quantity": 1,
						"price":    100.00,
					},
					{
						"id":       "67890",
						"name":     "Product 2",
						"quantity": 2,
						"price":    150.00,
					},
				},
				"installments": map[string]float64{
					"numberOfInstallments":   2,
					"installmentAmount":      125.00,
					"totalInstallmentAmount": 250.00,
				},
			}
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
			wantCode := http.StatusForbidden

			if gotCode != wantCode {
				t.Errorf("got %d, want %d", gotCode, wantCode)
			}

			var got vendor_service.ServiceResponse
			err = json.Unmarshal(response.Body.Bytes(), &got)
			if err != nil {
				t.Fatal(err)
			}

			want := "the minimal required value is 1500"
			if got.Message != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})

	})
}

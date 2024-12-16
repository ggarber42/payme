package mock

type Payload map[string]interface{}

type RequestPayload struct {
	Payload
}

func NewRequestPayload() *RequestPayload {
	payload := map[string]interface{}{
		"cardData": map[string]string{
			"name":       "John Doe",
			"token":      "7zIzcfXQJId44/NCv1o8hQ==",
			"expireDate": "12/32",
			"flag":       "mishacard",
		},
		"purchase": map[string]interface{}{
			"purchaseId":  "12345",
			"date":        "2024-11-04T12:30:00Z",
			"customerId":  "cust_67890",
			"totalAmount": 1500.0,
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
		},
		"store": map[string]interface{}{
			"name": "Raimunda e Sebastião Construções ME",
			"cnpj": "32.442.717/0001-80",
			"address": map[string]string{
				"street": "Rua R-Trinta e Nove",
				"number": "376",
				"cep":    "13279-105",
			},
		},
		"vendor": "stone",
	}
	return &RequestPayload{payload}
}

func (p *RequestPayload) InvalidCardData() *RequestPayload {
	p.Payload["cardData"] = map[string]string{}
	return p
}

func (p *RequestPayload) InvalidVendor() *RequestPayload{
	p.Payload["vendor"] = "invalid"
	return p
}

func (p *RequestPayload) Build() Payload {
	return p.Payload
}

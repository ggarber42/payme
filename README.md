

### Testing locally

```bash
curl -i localhost:3000/healthcheck | jq .

curl -i -X POST localhost:3000/payment \
  -H "Content-Type: application/json" \
  -d ' {
    "cardData": {
        "name": "John Doe",
        "token": "7zIzcfXQJId44/NCv1o8hQ==",
        "expireDate": "1234567812345678",
			  "flag": "mishacard"
    },
    "purchase": {
        "purchaseId": "12345",
        "date": "2024-11-04T12:30:00Z",
        "customerId": "cust_67890",
        "totalAmount": 1200.0,
        "currency": "USD",
        "products": [
            {
                "id": "12345",
                "name": "Product 1",
                "quantity": 1,
                "price": 100.00
            },
            {
                "id": "67890",
                "name": "Product 2",
                "quantity": 2,
                "price": 150.00
            }
        ],
        "installments": {
            "numberOfInstallments": 2,
            "installmentAmount": 125.00,
            "totalInstallmentAmount": 250.00
        }
    },
    "store": {
        "name": "Raimunda e Sebastião Construções ME",
        "cnpj": "32.442.717/0001-80",
        "address": {
            "street": "Rua R-Trinta e Nove",
            "number": "376",
            "cep": "13279-105"
        }
    },
    "purchaser": "cielo"
}'
```
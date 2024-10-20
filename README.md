

### Testing locally

```bash
curl -i localhost:3000/healthcheck | jq .

curl -i -X POST localhost:3000/payment \
  -H "Content-Type: application/json" \
  -d '{
    "article": {
      "title": "My New Article",
      "content": "This is the content of the article."
    },
    "user": {
      "name": "John Doe",
      "email": "johndoe@example.com"
    },
    "id": "123456"
  }'

curl -i -X POST localhost:3000/payment \
-H "Content-Type: application/json" \
-d '{
  "cardData": {
    "cardName": "John Doe"
  },
  "cardToken": "your_card_token_here",
  "shoppingData": {
    "products": [
      {
        "id": "12345",
        "name": "Product 1"
      },
      {
        "id": "67890",
        "name": "Product 2"
      }
    ]
  }
}'

curl -i -X POST localhost:3000/payment \
-H "Content-Type: application/json" \
-d '{
  "cardData": {
    "cardName": "John Doe",
    "cardToken": "your_card_token_here"
  },
  "shoppingData": {
    "products": [
      {
        "id": "12345",
        "name": "Product 1"
      },
      {
        "id": "67890",
        "name": "Product 2"
      }
    ]
  }
}'

```
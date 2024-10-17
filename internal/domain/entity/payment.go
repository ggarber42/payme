package entity

type PaymentRequest struct {
	CardData     *CardData     `json:"cardData"`
	CardToken    string        `json:"cardToken"`
	ShoppingData *ShoppingData `json:"shoppingData"`
}

type CardData struct {
	CardName string `json:"cardName"`
}

type ShoppingData struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

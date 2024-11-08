package entity

type Purchase struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

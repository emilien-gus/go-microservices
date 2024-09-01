package data

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Discription string  `json:"discription,omitempty"`
}

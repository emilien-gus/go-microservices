package data

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Discription string  `json:"discription,omitempty"`
}

var productsList = []*Product{
	&Product{
		ID:          0,
		Name:        "meat",
		Price:       500,
		Discription: "Boneless meat",
	},
}

func GetProducts() []*Product {
	return productsList
}

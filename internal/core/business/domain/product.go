package domain

type Product struct {
	Name        string
	Description string
	Price       float64
	Pictures    []Picture
	Discount
}

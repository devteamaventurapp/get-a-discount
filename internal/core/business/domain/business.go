package domain

type Business struct {
	ID   string
	Name string
	Category
	Location
	Products []Product
}

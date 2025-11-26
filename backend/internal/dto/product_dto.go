package dto

type CreateProductDTO struct {
	Name        string `form:"name"`
	Price       int    `form:"price"`
	Description string `form:"description"`
	Url         string //`default:"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTU073gNaeN3kAc7dWjS4ifOurPfBxK9X8rEwTTVIwp_WIq2WBKQHjot5r95ZL0GLlV_eZmAMXccU8Keg-8uOU9x2AcOlIMLAv-g8bHgH6z&s=10"`
}

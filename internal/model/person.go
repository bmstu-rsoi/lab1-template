package model

type Person struct {
	ID      int     `json:"id"`
	Name    string  `json:"name" validate:"required"`
	Age     *int    `json:"age"`
	Address *string `json:"address"`
	Work    *string `json:"work"`
}

package web

type CakeRequest struct {
	Title       string  `validate:"required" json:"title"`
	Description string  `validate:"required" json:"description"`
	Rating      float64 `validate:"required,numeric,min=0,max=10" json:"rating"`
	Image       string  `validate:"required,url" json:"image"`
}

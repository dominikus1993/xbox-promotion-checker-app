package games

type PromotionPrice = float64

type RegularPrice = float64

type Title = string
type Link = string

type XboxStoreGame struct {
	ID       string
	Title    Title
	Link     Link
	Price    PromotionPrice
	OldPrice RegularPrice
}

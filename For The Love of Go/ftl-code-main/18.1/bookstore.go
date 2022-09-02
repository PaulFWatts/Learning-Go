package bookstore

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
	ID int
	PriceCents int
	DiscountPercent int
}

func NetPriceCents(b Book) int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}
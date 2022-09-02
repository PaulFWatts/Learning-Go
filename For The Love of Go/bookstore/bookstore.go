package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
}

type Catalog map[int]Book

// Buy is a function that takes a Book, decrements its number of copies, and returns the modified
// book.
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

// GetAllBooks is a method that returns a slice of all books in the catalog.
func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

// GetBook is a method that returns a book from the catalog.
func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

// NetPriceCents is a method that returns the price of a book, taking into account any discount.
func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("bad price %d (must not be negative)", price)
	}
	b.PriceCents = price
	return nil
}
func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("unknown category %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}

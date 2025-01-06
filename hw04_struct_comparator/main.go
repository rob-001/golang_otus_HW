package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

func NewBook(id int, title string, author string, year, size int, rate float32) Book {
	return Book{
		ID:     id,
		Title:  title,
		Author: author,
		Year:   year,
		Size:   size,
		Rate:   rate,
	}
}

func (b Book) GetId() int {
	return b.ID
}

func (b Book) GetTitle() string {
	return b.Title
}

func (b Book) GetAuthor() string {
	return b.Author
}

func (b Book) GetYear() int {
	return b.Year
}

func (b Book) GetSize() int {
	return b.Size
}

func (b Book) GetRate() float32 {
	return b.Rate
}

func (b *Book) SetId(id int) {
	b.ID = id
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) SetAuthor(author string) {
	b.Author = author
}

func (b *Book) SetYear(year int) {
	b.Year = year
}

func (b *Book) SetSize(size int) {
	b.Size = size
}

func (b *Book) SetRate(rate float32) {
	b.Rate = rate
}

const (
	CompareByYear int = iota
	CompareBySize
	CompareByRate
)

func Compare(b1, b2 Book, CompareBook int) bool {
    switch CompareBook {
		case CompareByYear:
			return b1.Year > b2.Year
        case CompareBySize:
			return b1.Size > b2.Size
		case CompareByRate:
			return b1.Rate > b2.Rate
        default:
			return false
	}
}
func main() {
	Book1 := Book{}
	Book2 := Book{}
	Book1.SetId(10)
	Book1.SetTitle("Book 1")
	Book1.SetAuthor("Author 1")
	Book1.SetYear(2000)
	Book1.SetSize(1000)
	Book1.SetRate(4.5)
	Book2.SetId(20)
	Book2.SetTitle("Book 2")
	Book2.SetAuthor("Author 2")
	Book2.SetYear(210)
	Book2.SetSize(2001)
	Book2.SetRate(4.7)
	fmt.Print(Book1, Book2, "\n")
	fmt.Println(Compare(Book1, Book2, CompareByRate))
	fmt.Println(Compare(Book1, Book2, CompareBySize))
	fmt.Println(Compare(Book1, Book2, CompareByYear))

}

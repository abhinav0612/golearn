package data

var BooksData = make(map[string]interface{})

type Book struct {
	Name   string
	Author string
}

func init() {
	BooksData["1"] = Book{
		Name:   "Book 1",
		Author: "Author 1",
	}
	BooksData["2"] = Book{
		Name:   "Book 2",
		Author: "Author 2",
	}
	BooksData["3"] = Book{
		Name:   "Book 3",
		Author: "Author 3",
	}
	BooksData["4"] = Book{
		Name:   "Book 4",
		Author: "Author 4",
	}
}

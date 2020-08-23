package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go Lang Programming"
	Book1.author = "Rajesh Kumar"
	Book1.subject = "Go Lang Programming Tutorial"
	Book1.book_id = 6495434

	Book2.title = "Telecom Cloud Billing"
	Book2.author = "Ammit Kumar"
	Book2.subject = "Telecom Cloud Tutorial"
	Book2.book_id = 6495712

	printBook(&Book1)

	printBook(&Book2)
}
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

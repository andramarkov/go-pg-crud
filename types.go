package main

import "time"

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all books stored on db
type IndexPage struct {
	AllBooks []Book
}

//BookPage represents the content of the book page, available on "/book.html"
//The book page shows info about a given book
type BookPage struct {
	TargetBook Book
}

//Book represents a book object
type Book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}
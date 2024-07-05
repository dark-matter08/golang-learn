//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int
	lended int
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[Not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}

		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMembersAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBook(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()

}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not available in the library")
		return false
	}

	if book.lended == book.total {
		fmt.Println("All copies of this book have been lended out")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}

	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book is not part of this Libraries collection")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("This member did not checkout this book")
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit

	return true
}

func main() {

	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["WebApps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}
	library.books["The Little Go Book"] = BookEntry{
		total:  3,
		lended: 0,
	}
	library.books["Let's learn Go"] = BookEntry{
		total:  2,
		lended: 0,
	}
	library.books["Go BootCamp"] = BookEntry{
		total:  1,
		lended: 0,
	}

	library.members["Nde"] = Member{"Nde", make(map[Title]LendAudit)}
	library.members["Che"] = Member{"Che", make(map[Title]LendAudit)}
	library.members["Lucien"] = Member{"Lucien", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBook(&library)
	printMembersAudit(&library)

	member := library.members["Nde"]
	checkedOut := checkoutBook(&library, "Go BootCamp", &member)
	fmt.Println("\nChecked out a book:")

	if checkedOut {
		printLibraryBook(&library)
		printMembersAudit(&library)
	}

	returned := returnBook(&library, "Go BootCamp", &member)
	fmt.Println("\nChecked in book:")

	if returned {
		printLibraryBook(&library)
		printMembersAudit(&library)
	}

}

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

// create struct / model
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct{
	FName string `json:"fname"`
	LName string `json:"lname"`
}

// Initialize book variables as slices
var books  []Book = []Book{
	{ID: "1", Isbn: "1779221096",Title:"Hairdresser of Harare",Author:&Author{FName:"Tendai",LName:"Huchu"}},
	{ID: "2", Isbn: "0399588191",Title:"Born A Crime",Author:&Author{FName:"Trevor",LName:"Noah"}},
	{ID: "3", Isbn: "035698191",Title:"We Need New Names",Author:&Author{FName:"Bulawayo",LName:"Noviolet"}},
}




// create route handlers
// get all books
func getBooks(w http.ResponseWriter, r *http.Request)  {
	//set content type to application/json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)	
}

// get a book
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	// loop through all books anf find id
	for _,item := range books{
		if item.ID == parameters["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
		
	}
	json.NewEncoder(w).Encode(&Book{})
}
// create a book
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
	
}
// update a book's details
func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index,item := range books{
		if item.ID == parameters["id"]{
			books = append(books[:index],books[index+1:]... )
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = parameters["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}
//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for index,item := range books{
		if item.ID == parameters["id"]{
			books = append(books[:index],books[index+1:]... )
			break
			
		}
	}
	json.NewEncoder(w).Encode(books)
	
}



func main()  {
	// initialize the router
	r := mux.NewRouter()

	// create the route endpoint
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8000",r))
	


	
}
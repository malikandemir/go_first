package main

import(
    "encoding/json"
    "log"
    "net/http"
    "math/rand"
    "strconv"
    "github.com/gorilla/mux"
    "database/sql"
    "time"
    _ "github.com/go-sql-driver/mysql"

)

// Book Struct (Model)
type Book struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

// Author Struct
type Author struct{
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

// Inıt books var as as slice Book struct
var books []Book

// Get All Book
func getBooks(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(books)
}
// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  params := mux.Vars(r) // Get params

  for _, item :=range books{
    if item.ID == params["id"]{
        json.NewEncoder(w).Encode(item)
        return
    }
  }
  json.NewEncoder(w).Encode(&Book{})
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa( rand.Intn(100000000) )// Mock ID
  books = append(books,book)
  json.NewEncoder(w).Encode(book)
}

//Update a Book
func updateBook(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa( rand.Intn(100000000) )// Mock ID
  books = append(books,book)
  json.NewEncoder(w).Encode(book)

}

//Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  params := mux.Vars(r) // Get params

    for index, item :=range books{
      if item.ID == params["id"]{
          books = append(books[:index], books[index+1:]...)
          json.NewEncoder(w).Encode(item)
          return
      }
    }

  var book Book
  _ = json.NewDecoder(r.Body).Decode(&book)
  book.ID = strconv.Itoa( rand.Intn(100000000) )// Mock ID
  books = append(books,book)
  json.NewEncoder(w).Encode(book)
}

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func DbConnect(){
    db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.2:3306)/mehmetalikandemir.com")
    if err != nil {
        panic(err)
    }
    // See "Important settings" section.
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)

     // Execute the query
        results, err := db.Query("SELECT id, name FROM users")
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        for results.Next() {
                var user User
                // for each row, scan the result into our tag composite object
                err = results.Scan(&user.ID, &user.Name)
                if err != nil {
                    panic(err.Error()) // proper error handling instead of panic in your app
                }
                        // and then print out the tag's Name attribute
                log.Printf(user.Name)
            }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
}


func main(){

    DbConnect()

    // Init Router
    r := mux.NewRouter()

    // Mock Data - @todo - implement DB
    books = append(books, Book{ID: "1", Isbn: "5466" , Title: "Book One",
    Author: &Author{Firstname:"Ali", Lastname:"Kan"}})
    books = append(books, Book{ID: "2", Isbn: "5466" , Title: "Book Two",
    Author: &Author{Firstname:"Ali2", Lastname:"Kan"}})
    books = append(books, Book{ID: "3", Isbn: "5466" , Title: "Book Three",
    Author: &Author{Firstname:"Ali3", Lastname:"Kan"}})


    // Route Handlers / Endpoints
    r.HandleFunc("/api/books",getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
    r.HandleFunc("/api/books",createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000",r))


}
package entities

type Book struct {
    ID          int
    UserID      int
    Title       string
    Author      string
    Description string
    Status      string
    CreatedAt   string
    UpdatedAt   string
}

type BookRequest struct {
    ID           int
    RequesterID  int
    BookID       int
    Status       string
    CreatedAt    string
    UpdatedAt    string
}
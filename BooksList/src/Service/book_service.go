package service

import (
	"BooksList/src/Model"
	"errors"
)

//Criação do Service 
type BookService interface {
	AddBook(book *model.Book) error
	GetBook(id int) (*model.Book, error)
	UpdateBook(book *model.Book) error
	DeleteBook(id int) error
}

//Utilizando o map para armazenamento em memória (futuramente usar um BD?)
type bookService struct {
	books map[int]*model.Book
}

func NewBookService() BookService {
	return &bookService{
		books: make(map[int]*model.Book),
	}
}

//informa se o livro já existe no sistema
func (s *bookService) AddBook(book *model.Book) error {
	if _, exists := s.books[book.ID]; exists {
		return errors.New("book already exists")
	}
	s.books[book.ID] = book
	return nil
}

//Informa se o livro não foi encontrado no sistema
func (s *bookService) GetBook(id int) (*model.Book, error) {
	book, exists := s.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

//
func (s *bookService) UpdateBook(book *model.Book) error {
	if _, exists := s.books[book.ID]; !exists {
		return errors.New("book not found")
	}
	s.books[book.ID] = book
	return nil
}


//
func (s *bookService) DeleteBook(id int) error {
	if _, exists := s.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(s.books, id)
	return nil
}

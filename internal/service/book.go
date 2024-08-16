package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
)

type BookService struct {
	bookRepository domain.BookRepository
}

func NewBook(br domain.BookRepository) domain.BookService {
	return &BookService{
		bookRepository: br,
	}
}

// Index implements domain.BookService.
func (b *BookService) Index(ctx context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	booksData := make([]dto.BookData, 0, len(books))

	for _, v := range books {
		booksData = append(booksData, dto.BookData{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Isbn:        v.Isbn,
		})
	}

	return booksData, nil
}

// Create implements domain.BookService.
func (b *BookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		Id:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		Isbn:        req.Isbn,
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}

	return b.bookRepository.Save(ctx, &book)
}

// Delete implements domain.BookService.
func (b *BookService) Delete(ctx context.Context, id string) error {
	persisted, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("Book not found")
	}

	return b.bookRepository.Delete(ctx, id)
}

// Show implements domain.BookService.
func (b *BookService) Show(ctx context.Context, id string) (bd *dto.BookData, er error) {
	book, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return &dto.BookData{}, err
	}

	return &dto.BookData{Id: book.Id, Title: book.Title, Description: book.Description, Isbn: book.Isbn}, nil
}

// Update implements domain.BookService.
func (b *BookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return errors.New("Book not found")
	}

	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.Isbn = req.Isbn
	persisted.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return b.bookRepository.Update(ctx, &persisted)
}

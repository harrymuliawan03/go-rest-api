package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
)

type BookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(br domain.BookRepository, bsr domain.BookStockRepository) domain.BookService {
	return &BookService{
		bookRepository:      br,
		bookStockRepository: bsr,
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
		return domain.NewNotFoundError("Book")
	}

	return b.bookRepository.Delete(ctx, id)
}

// Show implements domain.BookService.
func (b *BookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	book, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return dto.BookShowData{}, err
	}

	stocks, err := b.bookStockRepository.FindByBookId(ctx, id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	stocksData := make([]dto.BookStockData, 0)

	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}

	return dto.BookShowData{
		BookData: dto.BookData{Id: book.Id, Title: book.Title, Description: book.Description, Isbn: book.Isbn},
		Stocks:   stocksData,
	}, nil
}

// Update implements domain.BookService.
func (b *BookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)

	if err != nil {
		return err
	}

	if persisted.Id == "" {
		return domain.NewNotFoundError("Book")
	}

	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.Isbn = req.Isbn
	persisted.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return b.bookRepository.Update(ctx, &persisted)
}

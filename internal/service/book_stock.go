package service

import (
	"context"

	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
)

type BookStockService struct {
	bookStockRepository domain.BookStockRepository
	bookRepository      domain.BookRepository
}

func NewBookStock(bsr domain.BookStockRepository, br domain.BookRepository) domain.BookStockService {
	return &BookStockService{
		bookStockRepository: bsr,
		bookRepository:      br,
	}
}

// Create implements domain.BookStockService.
func (b BookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	book, err := b.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}

	if book.Id == "" {
		return domain.NewNotFoundError("Book")
	}

	stocks := make([]domain.BookStock, 0)
	for _, v := range req.Codes {
		stocks = append(stocks, domain.BookStock{
			Code:   v,
			BookId: req.BookId,
			Status: domain.BookStockStatusAvailable,
		})
	}

	return b.bookStockRepository.Save(ctx, stocks)
}

// Delete implements domain.BookStockService.
func (b *BookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	return b.bookStockRepository.DeleteByCodes(ctx, req.Codes)
}

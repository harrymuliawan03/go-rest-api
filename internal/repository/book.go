package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/harrymuliawan03/go-rest-api/domain"
)

type BookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &BookRepository{
		db: goqu.New("default", con),
	}
}

// Delete implements domain.BookRepository.
func (br *BookRepository) Delete(ctx context.Context, id string) error {
	executor := br.db.Update("books").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()

	_, err := executor.ExecContext(ctx)

	return err
}

// FindAll implements domain.BookRepository.
func (br *BookRepository) FindAll(ctx context.Context) (result []domain.Book, err error) {
	dataset := br.db.From("books").Where(goqu.C("deleted_at").IsNull())

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements domain.BookRepository.
func (br *BookRepository) FindById(ctx context.Context, id string) (result domain.Book, err error) {
	dataset := br.db.From("books").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))

	_, err = dataset.ScanStructContext(ctx, &result)

	return
}

// Save implements domain.BookRepository.
func (br *BookRepository) Save(ctx context.Context, b *domain.Book) error {
	executor := br.db.Insert("books").Rows(b).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

// Update implements domain.BookRepository.
func (br *BookRepository) Update(ctx context.Context, b *domain.Book) error {
	executor := br.db.Update("books").Where(goqu.C("id").Eq(b.Id)).Set(b).Executor()

	_, err := executor.ExecContext(ctx)

	return err
}

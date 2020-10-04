package domain

import (
	"context"
	"time"
)

// Article ...
type Article struct {
	ID           int64           `json:"id"`
	Title        string          `json:"title" validate:"required"`
	Content      string          `json:"content" validate:"required"`
	Author       Author          `json:"author"`
	Token        string          `json:"token" validate:"required"`
	FireLocation FireLocationDTo `json:"fire_location"`
	CloudData    []byte          `json:"data"`
	UpdatedAt    time.Time       `json:"updated_at"`
	CreatedAt    time.Time       `json:"created_at"`
}

// FireLocationDTo - fire locations data transfer object
type FireLocationDTo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"string"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Article, string, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetFireLocation(ctx context.Context, ar *Article) (res []Article, err error)
	Update(ctx context.Context, ar *Article) error
	GetByTitle(ctx context.Context, title string) (Article, error)
	Store(context.Context, *Article) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	Update(ctx context.Context, ar *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	context "context"

	domain "github.com/NASA-Look-For-DUD/domain"

	mock "github.com/stretchr/testify/mock"
)

// ArticleUsecase is an autogenerated mock type for the ArticleUsecase type
type ArticleUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ArticleUsecase) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Article, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.Article); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Article)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ArticleUsecase) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Article); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *ArticleUsecase) GetByTitle(ctx context.Context, title string) (domain.Article, error) {
	ret := _m.Called(ctx, title)

	var r0 domain.Article
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Article); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(domain.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *ArticleUsecase) Store(_a0 context.Context, _a1 *domain.Article) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Article) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *ArticleUsecase) Update(ctx context.Context, ar *domain.Article) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Article) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

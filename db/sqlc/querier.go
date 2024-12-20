// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteStore(ctx context.Context, arg DeleteStoreParams) error
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
	ListStoresOfOwner(ctx context.Context, owner string) ([]Store, error)
	UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)

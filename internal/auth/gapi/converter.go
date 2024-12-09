package gapi

import (
	db "VOU-Server/db/sqlc"
	"VOU-Server/proto/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *gen.User {
	return &gen.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

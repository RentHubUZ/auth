package storage

import (
	pb "auth/genproto/user"
	"auth/internal/models"
	"context"
)

type IStorage interface {
	User() IUserStorage
	Close()
}

type IUserStorage interface {
	Add(ctx context.Context, u *models.RegisterReq) (*models.RegisterResp, error)
	Read(ctx context.Context, id string) (*pb.Profile, error)
	Update(ctx context.Context, u *pb.NewData) error
	Delete(ctx context.Context, id string) error
	Validate(ctx context.Context, id string) (*pb.Status, error)
	GetByEmail(ctx context.Context, email string) (*models.UserDetails, error)
	UpdatePassword(ctx context.Context, id, password string) error
	GetRole(ctx context.Context, id string) (string, error)
	GetPassword(ctx context.Context, id string) (string, error)
}

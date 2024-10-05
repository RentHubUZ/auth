package service

import (
	pb "auth/genproto/user"
	logger "auth/internal/logs"
	"auth/internal/storage"
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	pb.UnimplementedUserServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewUserService(s storage.IStorage) *UserService {
	return &UserService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *UserService) GetProfile(ctx context.Context, req *pb.ID) (*pb.Profile, error) {
	s.logger.Info("GetProfile is invoked", slog.Any("request", req))

	resp, err := s.storage.User().Read(ctx, req.Id)
	if err != nil {
		return nil, handleError(err, "failed to get profile", s.logger)
	}

	s.logger.Info("GetProfile is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, req *pb.NewData) (*pb.Void, error) {
	s.logger.Info("UpdateProfile is invoked", slog.Any("request", req))

	err := s.storage.User().Update(ctx, req)
	if err != nil {
		return nil, handleError(err, "failed to update profile", s.logger)
	}

	s.logger.Info("UpdateProfile is completed")
	return &pb.Void{}, nil
}

func (s *UserService) DeleteProfile(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.logger.Info("DeleteProfile is invoked", slog.Any("request", req))

	err := s.storage.User().Delete(ctx, req.Id)
	if err != nil {
		return nil, handleError(err, "failed to delete profile", s.logger)
	}

	s.logger.Info("DeleteProfile is completed")
	return &pb.Void{}, nil
}

func (s *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	s.logger.Info("ValidateUser is invoked", slog.Any("request", req))

	resp, err := s.storage.User().Validate(ctx, req.Id)
	if err != nil {
		return nil, handleError(err, "failed to validate user", s.logger)
	}

	s.logger.Info("ValidateUser is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *pb.NewPass) (*pb.Void, error) {
	s.logger.Info("UpdatePassword is invoked")

	password, err := s.storage.User().GetPassword(ctx, req.Id)
	if err != nil {
		return nil, handleError(err, "failed to get current password", s.logger)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.OldPassword)); err != nil {
		return nil, handleError(err, "invalid current password", s.logger)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, handleError(err, "failed to hash new password", s.logger)
	}

	err = s.storage.User().UpdatePassword(ctx, req.Id, string(hashedPassword))
	if err != nil {
		return nil, handleError(err, "failed to update password", s.logger)
	}

	s.logger.Info("UpdatePassword is completed")
	return &pb.Void{}, nil
}

func handleError(err error, msg string, logger *slog.Logger) error {
	er := errors.Wrap(err, msg)
	logger.Error(er.Error())
	return er
}

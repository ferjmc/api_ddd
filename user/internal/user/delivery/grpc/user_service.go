package grpc

import (
	"context"

	"github.com/ferjmc/api_ddd/user/internal/models"
	"github.com/ferjmc/api_ddd/user/internal/user"
	"github.com/ferjmc/api_ddd/user/pkg/grpc_errors"
	"github.com/ferjmc/api_ddd/user/pkg/logger"
	userService "github.com/ferjmc/api_ddd/user/proto/user"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	userUC user.UseCase
	logger logger.Logger
}

func NewUserService(userUC user.UseCase, logger logger.Logger) *UserService {
	return &UserService{userUC: userUC, logger: logger}
}

func (u *UserService) GetUserByID(ctx context.Context, r *userService.GetByIDRequest) (*userService.GetByIDResponse, error) {

	userUUID, err := uuid.FromString(r.GetUserID())
	if err != nil {
		u.logger.Errorf("uuid.FromString: %v", err)
		return nil, grpc_errors.ErrorResponse(err, "uuid.FromString")
	}

	foundUser, err := u.userUC.GetByID(ctx, userUUID)
	if err != nil {
		u.logger.Errorf("uuid.FromString: %v", err)
		return nil, grpc_errors.ErrorResponse(err, "userUC.GetByID")
	}

	return &userService.GetByIDResponse{User: foundUser.ToProto()}, nil
}

func (u *UserService) GetUsersByIDs(ctx context.Context, req *userService.GetByIDsReq) (*userService.GetByIDsRes, error) {
	usersByIDs, err := u.userUC.GetUsersByIDs(ctx, req.GetUsersIDs())
	if err != nil {
		u.logger.Errorf("userUC.GetUsersByIDs: %v", err)
		return nil, grpc_errors.ErrorResponse(err, "userUC.GetUsersByIDs")
	}

	u.logger.Infof("USERS LIST RESPONSE: %v", u.idsToUUID(usersByIDs))

	return &userService.GetByIDsRes{Users: u.idsToUUID(usersByIDs)}, nil
}

func (u *UserService) idsToUUID(users []*models.UserResponse) []*userService.User {
	usersList := make([]*userService.User, 0, len(users))
	for _, val := range users {
		usersList = append(usersList, val.ToProto())
	}

	return usersList
}

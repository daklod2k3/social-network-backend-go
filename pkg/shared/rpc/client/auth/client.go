package authRpcClient

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shared/entity"
	authEntity "shared/entity/auth"
	"shared/global"
	"shared/rpc/pb"
)

type authRpc struct {
	client pb.AuthClient
}

type AuthRpcService interface {
	GetSession(form *authEntity.SessionRequest) (*authEntity.AuthResponse, error)
}

func NewClient() *authRpc {
	var cfg = global.Config
	var (
		url = cfg.Auth.Url
	)

	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("could not connect to auth service:" + err.Error())
	}
	auth := pb.NewAuthClient(conn)

	return &authRpc{
		client: auth,
	}
}

func (c *authRpc) GetSession(form *authEntity.SessionRequest) (*authEntity.AuthResponse, error) {

	req, err := c.client.GetSession(context.Background(), &pb.SessionReq{
		AccessToken:  form.AccessToken,
		RefreshToken: form.RefreshToken,
	})

	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(req.UserId)

	if err != nil {
		return nil, err
	}

	var user *entity.User = nil
	if req.User != nil {
		id, _ := primitive.ObjectIDFromHex(req.User.Id)
		userId, _ := uuid.Parse(req.User.UserId)
		user = &entity.User{
			ID:          id,
			DisplayName: req.User.DisplayName,
			AvatarPath:  req.User.AvatarPath,
			Status:      req.User.Status,
			UserId:      userId,
		}
	}

	return &authEntity.AuthResponse{
		req.AccessToken,
		req.RefreshToken,
		user,
		&userId,
	}, nil
}

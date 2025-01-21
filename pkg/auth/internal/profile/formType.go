package profile

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	DisplayName *string `json:"display_name,required"`
	AvatarPath  *string `json:"avatar_path,omitempty,string"`
}

type SecureProfile struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DisplayName   string             `bson:"display_name" json:"display_name"`
	AvatarPath    string             `bson:"avt_path" json:"avt_path"`
	Status        string             `bson:"status" json:"status"`
	TotalFollower int                `bson:"total_follower" json:"total_follower"`
	UserId        uuid.UUID          `bson:"user_id" json:"user_id"`
}

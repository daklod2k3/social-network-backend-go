package profile

import "github.com/google/uuid"

type Post struct {
	DisplayName string    `json:"display_name,required"`
	AvatarPath  string    `json:"avatar_path"`
	Status      string    `json:"status"`
	UserId      uuid.UUID `json:"user_id,required"`
}

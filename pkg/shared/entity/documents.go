package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the Users collection
type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DisplayName   string             `bson:"display_name" json:"display_name"`
	AvatarPath    string             `bson:"avt_path" json:"avt_path"`
	Status        string             `bson:"status" json:"status"`
	TotalFollower int                `bson:"total_follower" json:"total_follower"`
	UserId        uuid.UUID          `bson:"user_id" json:"user_id"`
}

// Follow represents the Follows collection
type Follow struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	FollowersOf primitive.ObjectID   `bson:"followers_of" json:"followers_of"`
	Followers   []primitive.ObjectID `bson:"followers" json:"followers"`
}

// Post represents the Posts collection
type Post struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedBy     primitive.ObjectID `bson:"created_by" json:"created_by"`
	Content       string             `bson:"content" json:"content"`
	CreatedAt     primitive.DateTime `bson:"created_at" json:"created_at"`
	TotalLike     int                `bson:"total_like" json:"total_like"`
	TotalComments int                `bson:"total_comments" json:"total_comments"`
	TotalShare    int                `bson:"total_share" json:"total_share"`
	Videos        []Resource         `bson:"videos" json:"videos"`
	Images        []Resource         `bson:"images" json:"images"`
}

// Like represents the Likes collection
type Like struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type      string             `bson:"type" json:"type"`
	CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
	ParentID  primitive.ObjectID `bson:"parent_id" json:"parent_id"`
}

// Interaction represents the Interactions collection
type Interaction struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ParentID primitive.ObjectID `bson:"parent_id" json:"parent_id"`
	Comments int                `bson:"comments" json:"comments"`
	Likes    int                `bson:"likes" json:"likes"`
	Shares   int                `bson:"shares" json:"shares"`
}

// CommentPost represents the Comments : Posts collection
type CommentPost struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ParentID primitive.ObjectID `bson:"parent_id" json:"parent_id"`
}

// NewsFeed represents the NewsFeeds collection
type NewsFeed struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID primitive.ObjectID   `bson:"user_id" json:"user_id"`
	Posts  []primitive.ObjectID `bson:"posts" json:"posts"`
}

type Resource struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Path string             `bson:"path" json:"path"`
}

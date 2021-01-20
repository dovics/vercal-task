package service

import "time"

// ResponseCommentInfoSerializer -
type ResponseCommentInfoSerializer struct {
	Meta struct {
		ActiveReview string `json:"activeReview"`
	} `json:"meta"`
	Data []CommentSerializer `json:"data"`
}

// ResponseAddCommentSerializer -
type ResponseAddCommentSerializer struct {
	Data CommentSerializer
}

// CommentSerializer -
type CommentSerializer struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	User   struct {
		ID             int    `json:"id"`
		Login          string `json:"login"`
		Name           string `json:"name"`
		AvatarURL      string `json:"avatar_url"`
		FollowersCount int    `json:"followers_count"`
		FollowingCount int    `json:"following_count"`
		IsPaid         bool   `json:"isPaid"`
		Serializer     string `json:"_serializer"`
	} `json:"user"`
	ParentID      string    `json:"parent_id"`
	Format        string    `json:"format"`
	BodyAsl       string    `json:"body_asl"`
	LikesCount    int       `json:"likes_count"`
	Mood          int       `json:"mood"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Status        int       `json:"status"`
	ToUserID      string    `json:"to_user_id"`
	Type          string    `json:"type"`
	SelectionID   string    `json:"selection_id"`
	SelectionType string    `json:"selection_type"`
	Reactions     []string  `json:"reactions"`
	SourceType    string    `json:"source_type"`
	SourceID      string    `json:"source_id"`
	Serializer    string    `json:"_serializer"`
}

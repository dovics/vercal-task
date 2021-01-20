package service

import "time"

type ResponseCommentSerializer struct {
	Meta struct {
		ActiveReview interface{} `json:"activeReview"`
	} `json:"meta"`
	Data []CommentSerializer `json:"data"`
}

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
	ParentID      interface{}   `json:"parent_id"`
	Format        string        `json:"format"`
	BodyAsl       string        `json:"body_asl"`
	LikesCount    int           `json:"likes_count"`
	Mood          int           `json:"mood"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Status        int           `json:"status"`
	ToUserID      interface{}   `json:"to_user_id"`
	Type          interface{}   `json:"type"`
	SelectionID   interface{}   `json:"selection_id"`
	SelectionType interface{}   `json:"selection_type"`
	Reactions     []interface{} `json:"reactions"`
	SourceType    interface{}   `json:"source_type"`
	SourceID      interface{}   `json:"source_id"`
	Serializer    string        `json:"_serializer"`
}

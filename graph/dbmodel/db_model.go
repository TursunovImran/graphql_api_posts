package dbmodel

import (
	"time"
)

type Comment struct {
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PostID    uint       `json:"postId"`
	AuthorID  uint       `json:"authorId"`
	Content   string     `json:"content" gorm:"type:varchar(2000);"`
	CreatedAt time.Time  `json:"createdAt"`
	ParentID  *int       `json:"parentId,omitempty"`
	Replies   []*Comment `json:"replies,omitempty"`
}

type CommentInput struct {
	Content  string `json:"content"`
	ParentID *int   `json:"parentId,omitempty"`
}

type PaginationInput struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Post struct {
	ID         int        `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	AuthorID   int        `json:"authorId"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CreatedAt  time.Time  `json:"createdAt"`
	Comments   []*Comment `json:"comments,omitempty"`
	CanComment bool       `json:"canComment"`
}

type PostInput struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CanComment bool   `json:"canComment"`
}

type User struct {
	ID       int        `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username string     `json:"username"`
	Posts    []*Post    `json:"posts,omitempty"`
	Comments []*Comment `json:"comments,omitempty"`
}

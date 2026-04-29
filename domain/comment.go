package domain

type Comment struct {
	ID        int    `json:"id" db:"id"`
	Content   string `json:"content" db:"content"`
	PostID    int    `json:"postId" db:"post_id"`
	UserID    int    `json:"userId" db:"user_id"`
	Username  string `json:"username,omitempty" db:"username"` // For joins
	ParentID  *int   `json:"parentId,omitempty" db:"parent_id"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
}

package domain

type Post struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Slug        string `json:"slug" db:"slug"`
	Content     string `json:"content" db:"content"`
	Summary     string `json:"summary" db:"summary"`
	Thumbnail   string `json:"thumbnail,omitempty" db:"thumbnail"`
	AuthorID    int    `json:"authorId" db:"author_id"`
	AuthorName  string `json:"authorName,omitempty" db:"author_name"` // For joins
	Status      string `json:"status" db:"status"`                    // draft, published, archived
	ViewCount   int    `json:"viewCount" db:"view_count"`
	PublishedAt *int64 `json:"publishedAt,omitempty" db:"published_at"`
	CreatedAt   int64  `json:"createdAt" db:"created_at"`
	UpdatedAt   int64  `json:"updatedAt" db:"updated_at"`
}

package domain

type Tag struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Slug      string `json:"slug" db:"slug"`
	CreatedAt int64  `json:"createdAt" db:"created_at"`
	UpdatedAt int64  `json:"updatedAt" db:"updated_at"`
}

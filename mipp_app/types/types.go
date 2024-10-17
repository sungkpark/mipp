package types

import "time"

type Idea struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserName    string    `json:"userName"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Ideas struct {
	IDs []string `json:"ids" validate:"dive"`
}

type CreateIdeaPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserName    string `json:"userName"`
}

type IdeaStore interface {
	GetIdeaByID(ideaID int) (*Idea, error)
	GetIdeasByID(ideaIDs []int) ([]Idea, error)
	GetIdeas(offset int, limit int) ([]*Idea, error)
	CreateIdea(CreateIdeaPayload) error
}

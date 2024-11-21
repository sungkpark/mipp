package types

import (
	"database/sql"
	"time"
)

type Idea struct {
	ID          int       `json:"ideaId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserName    string    `json:"userName"`
	CapturedUrl string    `json:"capturedUrl"`
	DomainId    int       `json:"domainId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Ideas struct {
	IDs []string `json:"ids" validate:"dive"`
}

type CreateIdeaPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	UserName    string `json:"userName"`
	CapturedUrl string `json:"capturedUrl"`
	DomainName  string `json:"domainName"`
}

type IdeaStore interface {
	GetIdeaByID(ideaID int) (*Idea, error)
	GetIdeasByID(ideaIDs []int) ([]Idea, error)
	GetIdeas(offset int, limit int) ([]*Idea, error)
	CreateIdea(CreateIdeaPayload, int) error
}

type Domain struct {
	ID                 int            `json:"domainId"`
	DomainName         string         `json:"domainName"`
	CompanyInformation sql.NullString `json:"companyInformation"`
	Verified           bool           `json:"verified"`
}

type CreateDomainPayload struct {
	DomainName string `json:"domainName" validate:"required"`
}

type DomainStore interface {
	GetDomainByID(domainID int) (*Domain, error)
	GetDomainByName(domainName string) (*Domain, error)
	GetDomains() ([]*Domain, error)
}

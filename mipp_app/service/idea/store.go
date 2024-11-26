package idea

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"mipp.com/app/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetIdeaByID(ideaID int) (*types.Idea, error) {
	rows, err := s.db.Query("SELECT * FROM ideas WHERE ideaId = ?", ideaID)
	if err != nil {
		return nil, err
	}

	idea := new(types.Idea)
	for rows.Next() {
		idea, err = rowMapper(rows)
		if err != nil || idea.ID != ideaID {
			return nil, err
		}
	}

	return idea, nil
}

func (s *Store) GetIdeasByID(ideaIDs []int) ([]types.Idea, error) {
	placeholders := strings.Repeat(",?", len(ideaIDs)-1)
	query := fmt.Sprintf("SELECT * FROM ideas WHERE ideaId IN (?%s)", placeholders)

	args := make([]interface{}, len(ideaIDs))
	for i, v := range ideaIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	ideas := []types.Idea{}
	for rows.Next() {
		idea, err := rowMapper(rows)
		if err != nil {
			return nil, err
		}
		ideas = append(ideas, *idea)
	}
	return ideas, nil
}

func (s *Store) GetIdeasByDomainID(domainID int) ([]types.Idea, error) {
	rows, err := s.db.Query("SELECT * FROM ideas WHERE domainId = ? ORDER BY rand() LIMIT 3", domainID)
	if err != nil {
		return nil, err
	}

	ideas := []types.Idea{}
	for rows.Next() {
		idea, err := rowMapper(rows)
		if err != nil || idea.DomainId != domainID {
			return nil, err
		}
		ideas = append(ideas, *idea)
	}
	
	return ideas, nil
}

func (s *Store) GetIdeas(offset int, limit int) ([]*types.Idea, error) {
	query := fmt.Sprintf("SELECT * FROM ideas LIMIT %v OFFSET %v", limit, offset)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	ideas := []*types.Idea{}
	for rows.Next() {
		idea, err := rowMapper(rows)
		if err != nil {
			return nil, err
		}

		ideas = append(ideas, idea)
	}

	return ideas, nil
}

func (s *Store) CreateIdea(idea types.CreateIdeaPayload, domainId int) error {
	_, err := s.db.Exec("INSERT INTO ideas (title, description, username, capturedUrl, domainId, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?, ?)", idea.Title, idea.Description, idea.UserName, idea.CapturedUrl, domainId, time.Now(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateIdea(idea types.Idea) error {
	_, err := s.db.Exec("UPDATE ideas SET title = ?, description = ?, username = ?, updatedAt = ? WHERE ideaId = ?", idea.Title, idea.Description, idea.UserName, time.Now(), idea.ID)
	if err != nil {
		return err
	}
	return nil
}

func rowMapper(rows *sql.Rows) (*types.Idea, error) {
	idea := new(types.Idea)

	err := rows.Scan(
		&idea.ID,
		&idea.Title,
		&idea.Description,
		&idea.UserName,
		&idea.CapturedUrl,
		&idea.DomainId,
		&idea.CreatedAt,
		&idea.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return idea, nil
}

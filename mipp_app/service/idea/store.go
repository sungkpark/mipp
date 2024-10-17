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
	rows, err := s.db.Query("SELECT * FROM ideas WHERE id = ?", ideaID)
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
	query := fmt.Sprintf("SELECT * FROM ideas WHERE id IN (?%s)", placeholders)

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

func (s *Store) CreateIdea(idea types.CreateIdeaPayload) error {
	_, err := s.db.Exec("INSERT INTO ideas (title, description, username, createdAt) VALUES (?, ?, ?, ?)", idea.Title, idea.Description, idea.UserName, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateIdea(idea types.Idea) error {
	_, err := s.db.Exec("UPDATE ideas SET title = ?, description = ?, username = ? WHERE id = ?", idea.Title, idea.Description, idea.UserName, idea.ID)
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
		&idea.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return idea, nil
}

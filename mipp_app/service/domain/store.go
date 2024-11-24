package domain

import (
	"database/sql"
	"fmt"

	"mipp.com/app/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetDomainByID(domainID int) (*types.Domain, error) {
	rows, err := s.db.Query("SELECT * FROM domains WHERE domainId = ?", domainID)
	if err != nil {
		return nil, err
	}

	domain := new(types.Domain)
	for rows.Next() {
		domain, err = rowMapper(rows)
		if err != nil || domain.ID != domainID {
			return nil, err
		}
	}

	return domain, nil
}

func (s *Store) GetDomainByName(domainName string) (*types.Domain, error) {
	domain := new(types.Domain)
	err := s.db.QueryRow("SELECT * FROM domains WHERE domainName = ?", domainName).Scan(&domain.ID, &domain.DomainName, &domain.CompanyInformation, &domain.Verified)
	if err != nil {
		if err == sql.ErrNoRows {
			err := s.createDomain(domainName)
			if err != nil {
				return nil, err
			}

			domainId := new(sql.NullInt64)
			err = s.db.QueryRow("SELECT domainId FROM domains WHERE domainName = ?", domainName).Scan(&domainId)
			if err != nil {
				return nil, err
			}
			if domainId.Valid {
				domain, err = s.GetDomainByID(int(domainId.Int64))
			}
			if err != nil {
				return nil, err
			}

			return domain, nil
		} else {
			return nil, err
		}
	}

	return domain, nil
}

func (s *Store) GetDomains(limit int) ([]*types.Domain, error) {
	// give domains, some sort of time, so that we query 20 by most relevant
	query := fmt.Sprintf("SELECT * FROM domains LIMIT %v", limit)
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	domains := []*types.Domain{}
	for rows.Next() {
		domain, err := rowMapper(rows)
		if err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}
	
	return domains, nil
}

func (s *Store) createDomain(domainName string) error {
	_, err := s.db.Exec(`INSERT INTO domains (domainName) VALUES (?)`, domainName)
	if err != nil {
		return err
	}
	return nil
}

func rowMapper(rows *sql.Rows) (*types.Domain, error) {
	domain := new(types.Domain)

	err := rows.Scan(
		&domain.ID,
		&domain.DomainName,
		&domain.CompanyInformation,
		&domain.Verified,
	)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

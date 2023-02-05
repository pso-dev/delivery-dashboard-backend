package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/pso-dev/delivery-dashboard/backend/internal/data"
)

type Resource struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	JobTitle       string   `json:"jobTitle"`
	Manager        string   `json:"manager"`
	Location       string   `json:"location"`
	WorkGroup      string   `json:"workGroup"`
	Clearance      string   `json:"clearance"`
	Specialties    []string `json:"specialties"`
	Certifications []string `json:"certifications"`
	Active         bool     `json:"active"`
}

type ResourceRepository struct {
	DB *sql.DB
}

func (repo *ResourceRepository) Insert(r *Resource) error {
	query := ``

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []interface{}{
		r.ID,
		r.Name,
		r.Email,
		r.JobTitle,
		r.Manager,
		r.Location,
		r.WorkGroup,
		r.Clearance,
		r.Specialties,
		r.Certifications,
		r.Active,
	}

	return repo.DB.QueryRowContext(ctx, query, args...).Scan()
}

func (repo *ResourceRepository) Get(id int64) (*Resource, error) {
	query := ``

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var res Resource
	res.ID = id

	err := repo.DB.QueryRowContext(ctx, query, id).Scan(
		&res.Name,
		&res.Email,
		&res.JobTitle,
		&res.Manager,
		&res.Location,
		&res.WorkGroup,
		&res.Clearance,
		&res.Specialties,
		&res.Certifications,
		&res.Active,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, data.ErrNotFound
		default:
			return nil, err
		}
	}

	return &res, nil
}

func (repo *ResourceRepository) GetAll(name string, workgroups []string, clearance string, specialties []string,
	certifications []string, manager string, active bool) ([]*Resource, error) {
	query := ``

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []interface{}{
		name,
		pq.Array(workgroups),
		clearance,
		pq.Array(specialties),
		pq.Array(certifications),
		active,
	}

	rows, err := repo.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecords := 0
	resources := []*Resource{}

	for rows.Next() {
		var resource Resource
		err := rows.Scan(
			&totalRecords,
			&resource.ID,
			&resource.Name,
			&resource.Email,
			&resource.JobTitle,
			&resource.Manager,
			&resource.Location,
			&resource.WorkGroup,
			&resource.Clearance,
			pq.Array(&resource.Specialties),
			pq.Array(&resource.Certifications),
			&resource.Active,
		)
		if err != nil {
			return nil, err
		}
		resources = append(resources, &resource)
	}
	return resources, nil
}

func (repo *ResourceRepository) Update(r *Resource) error {
	query := ``

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args := []interface{}{
		r.Name,
		r.Email,
		r.JobTitle,
		r.Manager,
		r.Location,
		r.WorkGroup,
		r.Clearance,
		pq.Array(r.Specialties),
		pq.Array(r.Certifications),
		r.Active,
		r.ID,
	}

	err := repo.DB.QueryRowContext(ctx, query, args...).Scan()
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return data.ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (repo *ResourceRepository) Delete(id int64) error {
	query := ``

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo.DB.QueryRowContext(ctx, query, id)

	return nil
}

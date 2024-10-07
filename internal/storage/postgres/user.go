package postgres

import (
	pb "auth/genproto/user"
	"auth/internal/models"
	"auth/internal/storage"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.IUserStorage {
	return &UserRepo{db: db}
}

func (r *UserRepo) Add(ctx context.Context, u *models.RegisterReq) (*models.RegisterResp, error) {
	query := `
	insert into users
		(full_name, email, phone_number, hashed_password)
	values
		($1, $2, $3, $4)
	returning id
	`

	var id string
	err := r.db.QueryRowContext(ctx, query, u.FullName, u.Email, u.PhoneNumber, u.HashedPassword).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.RegisterResp{ID: id}, nil
}

func (r *UserRepo) Read(ctx context.Context, id string) (*pb.Profile, error) {
	query := `
	select
		full_name, email, phone_number, created_at, updated_at
	from
		users
	where
		deleted_at is null and id = $1
	`

	var p pb.Profile
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.FullName,
		&p.Email, &p.PhoneNumber, &p.CreatedAt, &p.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &p, nil
}

func (r *UserRepo) Update(ctx context.Context, u *pb.NewData) error {
	query := `
	update
		users
	set
		updated_at = now()
	`

	var params []interface{}
	if u.FullName != "" {
		query += fmt.Sprintf(", full_name = $%d", len(params)+1)
		params = append(params, u.FullName)
	}
	if u.Email != "" {
		query += fmt.Sprintf(", email = $%d", len(params)+1)
		params = append(params, u.Email)
	}
	if u.PhoneNumber != "" {
		query += fmt.Sprintf(", phone_number = $%d", len(params)+1)
		params = append(params, u.PhoneNumber)
	}
	query += fmt.Sprintf(" where deleted_at is null and id = $%d", len(params)+1)
	params = append(params, u.Id)

	res, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	query := `
	update
		users
	set
		deleted_at = now()
	where
		deleted_at is null and id = $1
	`

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepo) Validate(ctx context.Context, id string) (*pb.Status, error) {
	if id == "" {
		log.Print("empty id provided to validate user")
		return &pb.Status{Valid: false}, nil
	}

	query := `
	select
		exists (
			select 1
			from users
			where deleted_at is null and id = $1
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return nil, err
	}

	return &pb.Status{Valid: exists}, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*models.UserDetails, error) {
	query := `
	select
		id, hashed_password, role
	from
		users
	where
		deleted_at is null and email = $1
	`

	var u models.UserDetails
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.HashedPassword, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepo) UpdatePassword(ctx context.Context, id, password string) error {
	query := `
	update
		users
	set
		hashed_password = $1, updated_at = now()
	where
		deleted_at is null and id = $2
	`

	res, err := r.db.ExecContext(ctx, query, password, id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepo) GetRole(ctx context.Context, id string) (string, error) {
	query := `
	select
		role
	from
		users
	where
		deleted_at is null and id = $1
	`

	var role string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	return role, nil
}

func (r *UserRepo) GetPassword(ctx context.Context, id string) (string, error) {
	query := `
	select
		hashed_password
	from
		users
	where
		deleted_at is null and id = $1
	`

	var password string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	return password, nil
}

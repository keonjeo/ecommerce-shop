package postgres

import (
	"fmt"
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
)

// PgUserStore ...
type PgUserStore struct {
	PgStore
}

func newPgUserStore(pgst *PgStore) store.UserStore {
	return &PgUserStore{*pgst}
}

// Save ...
func (s PgUserStore) Save(user *model.User) (*model.User, *model.AppError) {
	q := `INSERT INTO public.user(first_name, last_name, username, email, password, gender, locale, avatar_url, active, email_verified, failed_attempts, last_login_at, created_at, updated_at, deleted_at) 
	VALUES(:first_name, :last_name, :username, :email, :password, :gender, :locale, :avatar_url, :active, :email_verified, :failed_attempts, :last_login_at, :created_at, :updated_at, :deleted_at)`
	_, err := s.db.NamedExec(q, user)

	if err != nil {
		if IsUniqueConstraintError(err) {
			return nil, model.NewAppError("PgUserStore.Save", "store.postgres.user.save.unique_constraint.app_error", nil, fmt.Sprintf("userID: %d, %v", user.ID, err.Error()), http.StatusInternalServerError)
		}
		return nil, model.NewAppError("PgUserStore.Save", "store.postgres.user.save.app_error", nil, fmt.Sprintf("userID: %d, %v", user.ID, err.Error()), http.StatusInternalServerError)
	}

	return user, nil
}

// Get ...
func (s PgUserStore) Get(id int) (*model.User, *model.AppError) {
	return &model.User{}, nil
}

// GetAll ...
func (s PgUserStore) GetAll() ([]*model.User, *model.AppError) {
	return []*model.User{}, nil
}

// GetByEmail ...
func (s PgUserStore) GetByEmail(email string) (*model.User, *model.AppError) {
	return &model.User{}, nil
}

// Update ...
func (s PgUserStore) Update(id int, u *model.User) (*model.User, *model.AppError) {
	return &model.User{}, nil
}

// Delete ...
func (s PgUserStore) Delete(id int) (*model.User, *model.AppError) {
	return &model.User{}, nil
}

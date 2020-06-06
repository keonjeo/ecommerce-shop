package postgres

import (
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// PgUserStore is the postgres implementation
type PgUserStore struct {
	PgStore
}

// NewPgUserStore creates the new user store
func NewPgUserStore(pgst *PgStore) store.UserStore {
	return &PgUserStore{*pgst}
}

var (
	msgSave             = &i18n.Message{ID: "store.postgres.user.save.app_error", Other: "could not save user to db"}
	msgGetByEmail       = &i18n.Message{ID: "store.postgres.user.login.app_error", Other: "could not get user by email"}
	msgUniqueConstraint = &i18n.Message{ID: "store.postgres.user.save.unique_constraint.app_error", Other: "invalid credentials"}
)

// Save inserts the new user in the db
func (s PgUserStore) Save(user *model.User) (*model.User, *model.AppErr) {
	q := `INSERT INTO public.user(first_name, last_name, username, email, password, gender, locale, avatar_url, active, email_verified, failed_attempts, last_login_at, created_at, updated_at, deleted_at) 
	VALUES(:first_name, :last_name, :username, :email, :password, :gender, :locale, :avatar_url, :active, :email_verified, :failed_attempts, :last_login_at, :created_at, :updated_at, :deleted_at) RETURNING id`

	var id int64
	rows, err := s.db.NamedQuery(q, user)
	defer rows.Close()
	if err != nil {
		return nil, model.NewAppErr("PgUserStore.Save", model.ErrInternal, locale.GetUserLocalizer("en"), msgSave, http.StatusInternalServerError, nil)
	}
	for rows.Next() {
		rows.Scan(&id)
	}
	if err := rows.Err(); err != nil {
		if IsUniqueConstraintError(err) {
			return nil, model.NewAppErr("PgUserStore.Save", model.ErrConflict, locale.GetUserLocalizer("en"), msgUniqueConstraint, http.StatusInternalServerError, nil)
		}
		return nil, model.NewAppErr("PgUserStore.Save", model.ErrInternal, locale.GetUserLocalizer("en"), msgSave, http.StatusInternalServerError, nil)
	}
	user.ID = id
	return user, nil
}

// Get ...
func (s PgUserStore) Get(id int) (*model.User, *model.AppErr) {
	return &model.User{}, nil
}

// GetAll ...
func (s PgUserStore) GetAll() ([]*model.User, *model.AppErr) {
	return []*model.User{}, nil
}

// GetByEmail gets one user by email
func (s PgUserStore) GetByEmail(email string) (*model.User, *model.AppErr) {
	var user model.User
	if err := s.db.Get(&user, "SELECT * FROM public.user where email = $1", email); err != nil {
		return nil, model.NewAppErr("PgUserStore.GetByEmail", model.ErrInternal, locale.GetUserLocalizer("en"), msgGetByEmail, http.StatusInternalServerError, nil)
	}
	return &user, nil
}

// Update ...
func (s PgUserStore) Update(id int, u *model.User) (*model.User, *model.AppErr) {
	return &model.User{}, nil
}

// Delete ...
func (s PgUserStore) Delete(id int) (*model.User, *model.AppErr) {
	return &model.User{}, nil
}

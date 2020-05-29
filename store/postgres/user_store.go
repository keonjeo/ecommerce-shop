package postgres

import (
	"net/http"

	"github.com/dankobgd/ecommerce-shop/model"
	"github.com/dankobgd/ecommerce-shop/store"
	"github.com/dankobgd/ecommerce-shop/utils/locale"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// PgUserStore ...
type PgUserStore struct {
	PgStore
}

func newPgUserStore(pgst *PgStore) store.UserStore {
	return &PgUserStore{*pgst}
}

var (
	msgSaveUser         = &i18n.Message{ID: "store.postgres.user.save.app_error", Other: "could not save user to db"}
	msgUniqueConstraint = &i18n.Message{ID: "store.postgres.user.save.unique_constraint.app_error", Other: "unique constraint error"}
)

// Save ...
func (s PgUserStore) Save(user *model.User) (*model.User, *model.AppErr) {
	q := `INSERT INTO public.user(first_name, last_name, username, email, password, gender, locale, avatar_url, active, email_verified, failed_attempts, last_login_at, created_at, updated_at, deleted_at) 
	VALUES(:first_name, :last_name, :username, :email, :password, :gender, :locale, :avatar_url, :active, :email_verified, :failed_attempts, :last_login_at, :created_at, :updated_at, :deleted_at)`
	_, err := s.db.NamedExec(q, user)

	if err != nil {
		if IsUniqueConstraintError(err) {
			return nil, model.NewAppErr("PgUserStore.Save", model.ErrConflict, locale.GetUserLocalizer("en"), msgUniqueConstraint, http.StatusInternalServerError, nil)
		}
		return nil, model.NewAppErr("PgUserStore.Save", model.ErrInternal, locale.GetUserLocalizer("en"), msgSaveUser, http.StatusInternalServerError, nil)
	}

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

// GetByEmail ...
func (s PgUserStore) GetByEmail(email string) (*model.User, *model.AppErr) {
	return &model.User{}, nil
}

// Update ...
func (s PgUserStore) Update(id int, u *model.User) (*model.User, *model.AppErr) {
	return &model.User{}, nil
}

// Delete ...
func (s PgUserStore) Delete(id int) (*model.User, *model.AppErr) {
	return &model.User{}, nil
}

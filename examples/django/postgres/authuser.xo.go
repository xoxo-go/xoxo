// Package postgres contains the types for schema 'public'.
package postgres

// GENERATED BY XOXO. DO NOT EDIT.

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

// AuthUser represents a row from 'public.auth_user'.
type AuthUser struct {
	ID          int         `json:"id"`           // id
	Password    string      `json:"password"`     // password
	LastLogin   pq.NullTime `json:"last_login"`   // last_login
	IsSuperuser bool        `json:"is_superuser"` // is_superuser
	Username    string      `json:"username"`     // username
	FirstName   string      `json:"first_name"`   // first_name
	LastName    string      `json:"last_name"`    // last_name
	Email       string      `json:"email"`        // email
	IsStaff     bool        `json:"is_staff"`     // is_staff
	IsActive    bool        `json:"is_active"`    // is_active
	DateJoined  *time.Time  `json:"date_joined"`  // date_joined

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthUser exists in the database.
func (au *AuthUser) Exists() bool {
	return au._exists
}

// Deleted provides information if the AuthUser has been deleted from the database.
func (au *AuthUser) Deleted() bool {
	return au._deleted
}

// Insert inserts the AuthUser to the database.
func (au *AuthUser) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if au._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.auth_user (` +
		`password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined)
	err = db.QueryRow(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined).Scan(&au.ID)
	if err != nil {
		return err
	}

	// set existence
	au._exists = true

	return nil
}

// Update updates the AuthUser in the database.
func (au *AuthUser) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !au._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if au._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.auth_user SET (` +
		`password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10` +
		`) WHERE id = $11`

	// run query
	XOLog(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.ID)
	_, err = db.Exec(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.ID)
	return err
}

// Save saves the AuthUser to the database.
func (au *AuthUser) Save(db XODB) error {
	if au.Exists() {
		return au.Update(db)
	}

	return au.Insert(db)
}

// Upsert performs an upsert for AuthUser.
//
// NOTE: PostgreSQL 9.5+ only
func (au *AuthUser) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if au._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.auth_user (` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.password, EXCLUDED.last_login, EXCLUDED.is_superuser, EXCLUDED.username, EXCLUDED.first_name, EXCLUDED.last_name, EXCLUDED.email, EXCLUDED.is_staff, EXCLUDED.is_active, EXCLUDED.date_joined` +
		`)`

	// run query
	XOLog(sqlstr, au.ID, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined)
	_, err = db.Exec(sqlstr, au.ID, au.Password, au.LastLogin, au.IsSuperuser, au.Username, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined)
	if err != nil {
		return err
	}

	// set existence
	au._exists = true

	return nil
}

// Delete deletes the AuthUser from the database.
func (au *AuthUser) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !au._exists {
		return nil
	}

	// if deleted, bail
	if au._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.auth_user WHERE id = $1`

	// run query
	XOLog(sqlstr, au.ID)
	_, err = db.Exec(sqlstr, au.ID)
	if err != nil {
		return err
	}

	// set deleted
	au._deleted = true

	return nil
}

// AuthUserByID retrieves a row from 'public.auth_user' as a AuthUser.
//
// Generated from index 'auth_user_pkey'.
func AuthUserByID(db XODB, id int) (*AuthUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined ` +
		`FROM public.auth_user ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	au := AuthUser{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.Username, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined)
	if err != nil {
		return nil, err
	}

	return &au, nil
}

// AuthUsersByUsername retrieves a row from 'public.auth_user' as a AuthUser.
//
// Generated from index 'auth_user_username_6821ab7c_like'.
func AuthUsersByUsername(db XODB, username string) ([]*AuthUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined ` +
		`FROM public.auth_user ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	q, err := db.Query(sqlstr, username)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthUser{}
	for q.Next() {
		au := AuthUser{
			_exists: true,
		}

		// scan
		err = q.Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.Username, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined)
		if err != nil {
			return nil, err
		}

		res = append(res, &au)
	}

	return res, nil
}

// AuthUserByUsername retrieves a row from 'public.auth_user' as a AuthUser.
//
// Generated from index 'auth_user_username_key'.
func AuthUserByUsername(db XODB, username string) (*AuthUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, username, first_name, last_name, email, is_staff, is_active, date_joined ` +
		`FROM public.auth_user ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	au := AuthUser{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username).Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.Username, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined)
	if err != nil {
		return nil, err
	}

	return &au, nil
}

package models

import "time"

type User struct {
	ID           int64      `db:"id" json:"id"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	Password     string     `db:"password" json:"-"`
	FirstName    string     `db:"first_name" json:"first_name"`
	LastName     string     `db:"last_name" json:"last_name"`
	Phone        string     `db:"phone" json:"phone"`
	ProfileImage string     `db:"profile_image" json:"profile_image"`
	Bio          string     `db:"bio" json:"bio"`
	DateOfBirth  *time.Time `db:"date_of_birth" json:"date_of_birth"`
	IsActive     bool       `db:"is_active" json:"is_active"`
	IsVerified   bool       `db:"is_verified" json:"is_verified"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

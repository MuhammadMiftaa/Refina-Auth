package model

import "database/sql"

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

type Users struct {
	Base
	Name           string       `gorm:"type:varchar(100);not null"`
	Email          string       `gorm:"type:varchar(100);unique;not null"`
	Password       string       `gorm:"type:varchar(100);not null"`
	Role           string       `gorm:"type:varchar(100);not null;default:'user'"`
	EmailVerfiedAt sql.NullTime `gorm:"type:timestamp"`
}
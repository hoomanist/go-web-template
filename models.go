package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Token    string
}

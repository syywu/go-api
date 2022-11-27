package handlers

import "gorm.io/gorm"

// dependency injection- initialse db connection once and use it everywhere (prevent too many db connection)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

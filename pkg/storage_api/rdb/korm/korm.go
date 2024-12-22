package korm

import (
	_ "github.com/mattn/go-sqlite3"
)

type I interface {
	Create(any) error
	Delete(any) error
	Update(any) error
	Query(any) error
	Close() error
}

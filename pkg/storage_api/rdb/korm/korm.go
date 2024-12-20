package korm

type I interface {
	Create(any) error
	Delete(any) error
	Update(any) error
	Query(any) error
	Close() error
}

package item

import "errors"

type Item interface {
	Name() string
	SetName(string)
	Id() string
	Species() string
	Tags() []Tag
	SetTag(Tag)
}

type Tag string

var (
	ErrAlreadyExists = errors.New("Item already exists")
	ErrDoesNotExist  = errors.New("Item does not exist")
)

type ItemRepository interface {
	GetAll() ([]Item, error)
	GetByName(string) (Item, error)
	GetById(string) (Item, error)
	NewItem(string, string, string) (Item, error)
}

package Map

import "errors"

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	name, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return name, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

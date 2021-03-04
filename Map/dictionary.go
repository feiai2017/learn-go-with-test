package Map

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	name, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return name, nil
}

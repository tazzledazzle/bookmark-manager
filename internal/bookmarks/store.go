package bookmarks

import (
	"errors"
)

// In-memory store and CRUD logic
var (
	bookmarks = make(map[int]Bookmark)
	nextID    = 1
)

func Create(b Bookmark) Bookmark {
	b.ID = nextID
	nextID++
	bookmarks[b.ID] = b
	return b
}

func GetAll() []Bookmark {
	values := make([]Bookmark, 0, len(bookmarks))
	for _, b := range bookmarks {
		values = append(values, b)
	}

	return values
}

func Get(id int) (Bookmark, error) {
	b, ok := bookmarks[id]
	if !ok {
		return Bookmark{}, errors.New("bookmark not found")
	}

	return b, nil
}

func Update(id int, b Bookmark) error {
	if _, ok := bookmarks[id]; !ok {
		return errors.New("not found")
	}
	b.ID = id
	bookmarks[id] = b
	return nil
}

func Delete(id int) error {
	if _, ok := bookmarks[id]; !ok {
		return errors.New("not found")
	}
	delete(bookmarks, id)
	return nil
}

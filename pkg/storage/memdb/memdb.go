package memdb

import "got_news/pkg/storage"

// Хранилище данных.
type Store struct{}

// Конструктор объекта хранилища.
func New() *Store {
	return new(Store)
}

func (s *Store) AddPost(storage.Post) error {
	return nil
}

func (s *Store) DeletePost(storage.Post) error {
	return nil
}

func (s *Store) PostsNItems(storage.Post) ([]storage.Post, error) {
	return posts, nil
}

var posts = []storage.Post{
	{
		ID:      1,
		Title:   "Статья 1",
		Content: "Содержание статьи 1",
		PubTime: 1,
		Link:    "http://http1",
	},
	{
		ID:      2,
		Title:   "Статья 2",
		Content: "Содержание статьи 2",
		PubTime: 2,
		Link:    "http://http2",
	},
	{
		ID:      3,
		Title:   "Статья 3",
		Content: "Содержание статьи 3",
		PubTime: 3,
		Link:    "http://http3",
	},
}

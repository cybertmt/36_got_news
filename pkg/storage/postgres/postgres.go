package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"got_news/pkg/storage"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// AddPost создает статью.
func (s *Storage) AddPost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		INSERT INTO posts (title, content, pubTime, link)
       	SELECT $1, $2, $3, $4
       	WHERE NOT EXISTS (SELECT 1 FROM posts WHERE title=$1);
	`,
		p.Title, p.Content, p.PubTime, p.Link,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	// ВАЖНО не забыть проверить rows.Err()
	return rows.Err()
}

// DeletePost удаляет статью по id.
func (s *Storage) DeletePost(p storage.Post) error {
	rows, err := s.db.Query(context.Background(), `
		DELETE FROM posts
		WHERE posts.id = $1;
	`,
		p.ID,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	// ВАЖНО не забыть проверить rows.Err()
	return rows.Err()
}

// PostById возвращает статьи, отсортированные по времени создания, в количестве = n.
func (s *Storage) PostsNItems(p storage.Post) ([]storage.Post, error) {
	n := p.ID
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			*
		FROM posts
		ORDER BY pubTime DESC
		LIMIT $1;
	`, n,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PubTime,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

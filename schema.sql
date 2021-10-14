DROP TABLE IF EXISTS posts;

CREATE TABLE posts
(
    id      SERIAL PRIMARY KEY,
    title   TEXT   NOT NULL,
    content TEXT   NOT NULL,
    pubTime BIGINT NOT NULL,
    link    TEXT   NOT NULL

);

-- INSERT INTO posts (title, content, pubtime, link) VALUES ('Статья 1', 'Содержание статьи 1', 1,'http//http1');
-- INSERT INTO posts (title, content, pubtime, link) VALUES ('Статья 2', 'Содержание статьи 2', 0,'http//http2');
-- INSERT INTO posts (title, content, pubtime, link) VALUES ('Статья 3', 'Содержание статьи 3', 2,'http//http3');
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (1, 'Статья 2', 'Содержание статьи 2', 0);
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (1, 'Статья 3', 'Содержание статьи 3', 0);
-- INSERT INTO posts (author_id, title, content, created_at) VALUES (0, 'Статья 4', 'Содержание статьи 4', 11);
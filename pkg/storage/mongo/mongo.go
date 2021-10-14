package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"got_news/pkg/storage"
	"log"
)

type Storage struct {
	Client *mongo.Client
	DB     *mongo.Database
}

const (
	databaseName   = "gotnews" // имя БД
	collectionName = "posts"   // имя коллекции в БД
)

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрывать ресурсы
	//defer client.Disconnect(context.Background())
	// проверка связи с БД
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	s := Storage{
		Client: client,
		DB:     client.Database(databaseName),
	}
	return &s, nil
}

// AddPost создает статью.
func (s *Storage) AddPost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	result := collection.FindOne(context.Background(), bson.M{"title": p.Title})
	if result.Err() == mongo.ErrNoDocuments {
		_, err := collection.InsertOne(context.Background(), p)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeletePost удаляет статью по id.
func (s *Storage) DeletePost(p storage.Post) error {
	collection := s.Client.Database(databaseName).Collection(collectionName)
	_, err := collection.DeleteOne(context.Background(), bson.M{"id": p.ID})
	if err != nil {
		return err
	}
	return nil
}

// PostById возвращает статьи, отсортированные по времени создания, в количестве = n.
func (s *Storage) PostsNItems(p storage.Post) ([]storage.Post, error) {
	n := p.ID
	collection := s.Client.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	options := options.Find()

	// Sort by `_id` field descending
	options.SetSort(bson.D{{"pubtime", -1}})

	// Limit by 10 documents only
	options.SetLimit(int64(n))

	cur, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	//defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var t storage.Post
		err := cur.Decode(&t)
		if err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	return data, cur.Err()
}

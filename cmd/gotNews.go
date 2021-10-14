package main

import (
	"encoding/json"
	"got_news/pkg/api"
	"got_news/pkg/rss"
	"got_news/pkg/storage"
	"got_news/pkg/storage/postgres"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	// канал для чтения новостей rss-каналов и записи в бд
	PostChannel := make(chan storage.Post)
	// канал сбора ошибок rss обработчика
	ErrorChannel := make(chan error)

	// структура config.jason для rss обработчика
	type rssConfig struct {
		Rss     []string `json:"rss"`
		RPeriod int64      `json:"request_period"`
	}

	// Читаем файл конфигурации в последовательность байт
	jsonFile, err := os.Open("cmd/config.json")
	if err != nil {
		ErrorChannel <- err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Конвертируем байты в структуру rss конфигурации
	var rssCfg rssConfig
	_ = json.Unmarshal(byteValue, &rssCfg)

	// Запускаем rss обработчики в отдельных горутинах с конфигурациями
	for i := 0; i < len(rssCfg.Rss); i++ {
		go rss.FetchRss(rssCfg.Rss[i], rssCfg.RPeriod, PostChannel, ErrorChannel)
	}

	// Сервер GotNews.
	type server struct {
		db  storage.Interface
		api *api.API
	}
	// Создаём объект сервера.
	var srv server

	// используем переменные окружения для адреса бд, имени бд и пароля
	host := os.Getenv("host")
	bdName := os.Getenv("bdName")
	pwd := os.Getenv("pwd")

	//Создаём объекты баз данных.
	//
	//БД в памяти.
	//db1 := memdb.New()

	// Реляционная БД PostgreSQL.
	db2, err := postgres.New("postgres://cyber:" + pwd + "@" + host + "/" + bdName)
	if err != nil {
		ErrorChannel <- err
	}
	//Документная БД MongoDB.
	//db3, err := mongo.New("mongodb://" + host + ":27017/")
	//if err != nil {
	//	ErrorChannel <- err
	//}
	//_, _, _ = db1, db2, db3

	//Инициализируем хранилище сервера конкретной БД.
	srv.db = db2

	// Запускаем запись новостей из канал PostChannel в бд в отдельной горутине
	go func() {
		for {
			err = srv.db.AddPost(<- PostChannel)
			if err != nil {
				ErrorChannel <- err
			}
		}
	}()

	// Запускаем логирование ошибок из канала ErrorChannel
	go func() {
		for {
			log.Println(<-ErrorChannel)
		}
	}()

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	err = http.ListenAndServe(":80", srv.api.Router())
	if err != nil {
		ErrorChannel <- err
	}

}

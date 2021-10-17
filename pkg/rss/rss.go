package rss

import (
	"github.com/mmcdole/gofeed"
	"got_news/pkg/storage"
	"time"
)

// функция парсит rss feed и отправляет объект storage.Post в канал
func FetchRss(link string, rPeriod int64, postChan chan<- storage.Post, errChan chan<- error) {
	for {

		var p storage.Post

		// создаем парсер для rss канала
		fp := gofeed.NewParser()

		// парсим указанную ссылку
		feed, err := fp.ParseURL(link)

		// в случае неудачи отправляем ошибку в канал errChan
		if err != nil {
			errChan <- err
		}

		// для каждого элемента массива feed.items, который является списком новостей
		// производим сопоставление с основной структурой новости storage.Post
		// и отправляем в канал postChan
		for i := 0; i < len(feed.Items); i++ {
			p.Link = feed.Items[i].Link
			p.Title = feed.Items[i].Title
			p.Content = feed.Items[i].Description
			t := feed.Items[i].PublishedParsed
			tUnix := t.Unix()
			p.PubTime = tUnix

			postChan <- p

		}

		// повторяем проход парсера через rPeriod минут
		time.Sleep(time.Duration(rPeriod) * time.Minute)
	}
}

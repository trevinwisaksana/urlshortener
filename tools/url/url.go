package urltools

import (
	"time"

	"github.com/trevinwisaksana/trevin-urlshortener/model"
	"github.com/trevinwisaksana/trevin-urlshortener/tools"
)

func DummyURL(owner string) (url model.Url) {
	randomID := tools.RandomAlphanumericString(5)

	url = model.Url{
		LongUrl:   "https://www.notion.so/stockbit/Backend-Engineering-Challenge-Link-Shortener-82bf71375701427c9cdd54a10a775ba6?pvs=4",
		ShortUrl:  randomID,
		CreatedAt: time.Now(),
		Owner:     owner,
	}

	return
}

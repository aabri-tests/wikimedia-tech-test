package wikimedia

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/wikimedia/internal/application/ports"
	"github.com/wikimedia/internal/infrastructure/adapters/wikimedia/http"
	"github.com/wikimedia/pkg/config"
)

type WikiMediaFactory struct {
	Cfg *config.Config
}

func New(c *config.Config) ports.WikiMediaFactory {
	return &WikiMediaFactory{
		Cfg: c,
	}
}

func (w WikiMediaFactory) GetWikiMediaAdapter() (ports.WikiMedia, error) {
	client := resty.New()
	client.SetRetryCount(3)
	client.SetRetryWaitTime(5 * time.Second)
	client.SetRetryMaxWaitTime(20 * time.Second)

	return &http.WikiMedia{
		Client: client,
	}, nil
}

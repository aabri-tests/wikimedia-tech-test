package usecases

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/wikimedia/internal/application/ports"
)

// Response represents the response from the Wikipedia API
type Response struct {
	Query struct {
		Pages []struct {
			Title     string
			Missing   bool
			Revisions []struct {
				Content string `json:"content"`
			} `json:"revisions"`
		} `json:"pages"`
	} `json:"query"`
}

type searchUseCase struct {
	wikiMedia ports.WikiMedia
	Parser    ports.Parser
	Cache     ports.Cache
	Log       ports.LogInfoFormat
}

func New(wikiMedia ports.WikiMedia, parser ports.Parser, cache ports.Cache, log ports.LogInfoFormat) ports.WikiMediaUseCase {
	return &searchUseCase{
		wikiMedia: wikiMedia,
		Parser:    parser,
		Cache:     cache,
		Log:       log,
	}
}

func (s *searchUseCase) Search(query, language string) (interface{}, error) {
	cacheKey := s.generateCacheKey(query, language)
	shortDescription, err := s.Cache.Get(cacheKey)
	if string(shortDescription) == "" || err != nil {
		body, searchErr := s.wikiMedia.Search(query, language)
		if searchErr != nil {
			return nil, searchErr
		}

		var response Response
		marshalErr := json.Unmarshal(body.([]byte), &response)
		if marshalErr != nil {
			return nil, marshalErr
		}

		if response.Query.Pages[0].Missing {
			return "", nil // Person not found
		}

		content := response.Query.Pages[0].Revisions[0].Content
		parsed, err := s.Parser.Parse(content)
		if err != nil {
			return nil, err
		}
		shortDescription = []byte(parsed)
		// cache the response
		err = s.Cache.Set(cacheKey, shortDescription)
		if err != nil {
			s.Log.Warnf("Error caching response for key %s: %s", cacheKey, err)
		}
	}

	return string(shortDescription), nil
}

func (s *searchUseCase) generateCacheKey(query, language string) string {
	return fmt.Sprintf("%s-%s", language, strings.ReplaceAll(query, " ", ""))
}

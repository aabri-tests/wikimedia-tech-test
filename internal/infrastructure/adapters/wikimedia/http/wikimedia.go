package http

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

var baseURL = "https://%s.wikipedia.org/w/api.php"

type WikiMedia struct {
	Client *resty.Client
}

func (w *WikiMedia) Search(query, language string) (interface{}, error) {
	formattedName := strings.ReplaceAll(strings.Title(query), " ", "_")
	params := w.generateSearchParams(formattedName)

	response, err := w.sendRequest(params, language)
	if err != nil {
		return nil, err
	}

	return response.Body(), nil
}
func (w *WikiMedia) generateSearchParams(name string) map[string]string {
	return map[string]string{
		"action":        "query",
		"prop":          "revisions",
		"titles":        name,
		"rvlimit":       "1",
		"formatversion": "2",
		"format":        "json",
		"rvprop":        "content",
	}
}
func (w *WikiMedia) sendRequest(params map[string]string, language string) (*resty.Response, error) {
	uri := fmt.Sprintf(baseURL, language)
	response, err := w.Client.R().SetQueryParams(params).Get(uri)
	if err != nil {
		return nil, err
	}
	return response, nil
}

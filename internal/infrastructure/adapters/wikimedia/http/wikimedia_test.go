package http_test

import (
	http2 "net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wikimedia/internal/infrastructure/adapters/wikimedia/http"
)

func TestWikiMedia_Search_Success(t *testing.T) {
	// Create a new WikiMedia client with a mocked REST client
	client := http.WikiMedia{
		Client: resty.New(),
	}
	// Create a test server that always returns a 200 OK response with a JSON body
	testServer := httptest.NewServer(http2.HandlerFunc(func(w http2.ResponseWriter, r *http2.Request) {
		w.WriteHeader(http2.StatusOK)
		w.Write([]byte(`{"query": {"pages": [{"title": "Test Page"}]}}`))
	}))
	defer testServer.Close()
	client.Client.SetBaseURL(testServer.URL)

	// Call the Search method with a valid query
	_, err := client.Search("test page", "en")
	assert.NoError(t, err)
}

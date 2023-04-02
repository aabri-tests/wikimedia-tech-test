package usecases_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/wikimedia/internal/application/usecases"
	"github.com/wikimedia/internal/mocks"
)

type expectedResult *string

func stringPtr(s string) *string {
	return &s
}

var testCases = []struct {
	name           string
	query          string
	language       string
	rawData        string
	expectedResult expectedResult
	expectedError  error
	cacheError     error
}{
	{
		name:           "cache miss",
		query:          "Isaac Newton",
		language:       "en",
		rawData:        `{"query": {"pages": [{"title": "Isaac Newton","revisions": [{"content": "{{Short description|%s}}"}]}]}}`,
		expectedResult: stringPtr("Isaac Newton was an English mathematician, physicist, and astronomer. He is known for discovering the laws of motion and universal gravitation"),
		expectedError:  nil,
		cacheError:     errors.New("cache miss"),
	},
	{
		name:           "Internal error",
		query:          "Isaac Newton",
		language:       "en",
		rawData:        `{"query": {"pages": [{"title": "Isaac Newton","revisions": [{"content": "{{Short description|%s}}"}]}]}}`,
		expectedResult: nil,
		expectedError:  errors.New("internal error"),
		cacheError:     nil,
	},
	{
		name:           "unmarshal error",
		query:          "Isaac Newton",
		language:       "en",
		rawData:        `{[{"content": "{{Short description|%s}}"}]}]}}`,
		expectedResult: stringPtr("Isaac Newton was an English mathematician, physicist, and astronomer. He is known for discovering the laws of motion and universal gravitation"),
		expectedError:  errors.New("invalid character '[' looking for beginning of object key string"),
		cacheError:     nil,
	},
}

func TestSearchUseCase_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Loop through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wikiMedia := mocks.NewMockWikiMedia(ctrl)
			parser := mocks.NewMockParser(ctrl)
			cache := mocks.NewMockCache(ctrl)
			log := mocks.NewMockLogInfoFormat(ctrl)

			cacheKey := fmt.Sprintf("%s-%s", tc.language, strings.ReplaceAll(tc.query, " ", ""))

			if tc.expectedError == nil {
				rawData := fmt.Sprintf(tc.rawData, *tc.expectedResult)
				cache.EXPECT().Set(cacheKey, []byte(*tc.expectedResult)).AnyTimes().Return(nil)
				parser.EXPECT().Parse(fmt.Sprintf(`{{Short description|%s}}`, *tc.expectedResult)).AnyTimes().Return(*tc.expectedResult, nil)
				wikiMedia.EXPECT().Search(tc.query, tc.language).AnyTimes().Return([]byte(rawData), tc.expectedError)
			} else if tc.expectedError.Error() == "internal error" {
				wikiMedia.EXPECT().Search(tc.query, tc.language).AnyTimes().Return(nil, tc.expectedError)
			} else {
				rawData := fmt.Sprintf(tc.rawData, *tc.expectedResult)
				wikiMedia.EXPECT().Search(tc.query, tc.language).AnyTimes().Return([]byte(rawData), nil)
			}

			cache.EXPECT().Get(cacheKey).AnyTimes().Return([]byte(""), tc.cacheError)

			service := usecases.New(wikiMedia, parser, cache, log)
			search, err := service.Search(tc.query, tc.language)

			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, *tc.expectedResult, search)
			}
		})
	}
}

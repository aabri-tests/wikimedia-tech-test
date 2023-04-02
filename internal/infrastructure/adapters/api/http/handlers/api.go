package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wikimedia/internal/application/ports"
	"github.com/wikimedia/internal/domain/entities"
	"github.com/wikimedia/internal/infrastructure/adapters/api/http/responses"
	"github.com/wikimedia/internal/infrastructure/adapters/api/http/validator"
	"github.com/wikimedia/internal/utils"
)

var (
	defaultLanguage = "en"
	notFound        = errors.New("This Person is not found in the specified language edition of Wikipedia")
)

type apiHandler struct {
	wikiMediaUseCase ports.WikiMediaUseCase
}

func New(wikiMediaUseCase ports.WikiMediaUseCase) ports.WikiMediaHTTPHandler {
	return &apiHandler{
		wikiMediaUseCase: wikiMediaUseCase,
	}
}

// Search
// PingExample godoc
// @Summary Search by full name
// @Schemes
// @Description Search by full name
// @Tags search
// @Accept json
// @Produce json
// @Param   keyword     query    string     true        "yoshua bengio"
// @Param   language     query    string     false        "en"
// @Success      200  {object}  entities.Search
// @Failure      400,404,500  {object}  entities.Status
// @Router /search [get]
func (handler *apiHandler) Search(c *gin.Context) {
	response := make(chan interface{})
	errChan := make(chan entities.Status, 1)
	query := getFiltersFromContext(c)

	// Validate the user request
	if err := validator.ValidateSearchRequest(query); err != nil {
		errChan <- entities.Status{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if !validator.IsLanguageAllowed(query.Language) {
		errChan <- entities.Status{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Only %#v are supported", validator.AllowedLanguages),
		}
		return
	}

	go func() {
		result, err := handler.wikiMediaUseCase.Search(query.Keyword, query.Language)
		if err != nil {
			errChan <- entities.Status{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}
		if utils.IsEmpty(result.(string)) {
			errChan <- entities.Status{
				Code:    http.StatusNotFound,
				Message: notFound.Error(),
			}
		}
		response <- result
	}()

	select {
	case res := <-response:
		resp := res.(string)
		c.JSON(http.StatusOK, responses.SearchResponse{
			Name:             strings.Title(query.Keyword),
			ShortDescription: resp,
			Language:         query.Language,
			Status: responses.Status{
				Code:    http.StatusOK,
				Message: "Ok",
			},
		})
	case err := <-errChan:
		c.AbortWithStatusJSON(err.Code, err.Message)
	}
}

// HealthCheck
// PingExample godoc
// @Summary Health check endpoint
// @Schemes
// @Description Health check endpoint
// @Tags health
// @Accept json
// @Produce json
// @Success      200  {object}  responses.HealthResponse
// @Router /health [get]
func (handler *apiHandler) HealthCheck(c *gin.Context) {
	response := make(chan bool)
	go func() {
		response <- true
	}()
	select {
	case _ = <-response:
		c.JSON(http.StatusOK, responses.HealthResponse{Status: "ok"})
	}
}

// Metrics
// PingExample godoc
// @Summary Prometheus Metrics
// @Schemes
// @Description Prometheus Metrics
// @Tags metrics
// @Accept json
// @Produce json
// @Success 200 {string} string "Prometheus metrics in plain text format"
// @Router /metrics [get]
func (handler *apiHandler) Metrics() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}

func getFiltersFromContext(c *gin.Context) validator.SearchQuery {
	keyword := c.Query("keyword")
	language := c.Query("language")
	if language == "" {
		language = defaultLanguage
	}
	return validator.SearchQuery{Keyword: keyword, Language: language}
}

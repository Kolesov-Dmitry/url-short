package url_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-short/internal/repos/urls"
	"url-short/internal/url"
	"url-short/mocks"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeUrlStoreMock() *mocks.UrlStore {
	storeMock := &mocks.UrlStore{}

	storeMock.On("Create", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, u *urls.Url) error {
			return nil
		},
	)

	storeMock.On("Read", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, hash urls.UrlHash) *urls.Url {
			// return nil for http://www.test.org
			if hash == urls.UrlHash("217962200") {
				return nil
			}

			return &urls.Url{
				Hash: hash,
				URL:  "http://www.test_2.org",
			}
		},

		func(ctx context.Context, hash urls.UrlHash) error {
			return nil
		},
	)

	return storeMock
}

func makeLinkingStoreMock() *mocks.LinkingStore {
	linkingMock := &mocks.LinkingStore{}

	linkingMock.On("Create", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, hash urls.UrlHash) error {
			return nil
		},
	)

	linkingMock.On("Read", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, hash urls.UrlHash) chan string {
			result := make(chan string, 2)
			result <- "test_1"
			result <- "test_2"
			close(result)

			return result
		},
	)

	return linkingMock
}

func Test_Shorten_Post(t *testing.T) {
	// Mocking
	urlStoreMock := makeUrlStoreMock()
	linkingStore := makeLinkingStoreMock()

	// Setup PostsService
	router := mux.NewRouter()
	shortenService := url.NewService(urlStoreMock, linkingStore)
	shortenService.Register(router)

	// Setup server
	server := httptest.NewServer(router)
	defer server.Close()

	// Do request
	requestString := "{\"url\": \"http://www.test.org\"}"
	res, err := http.Post(server.URL+"/api/v1/shorten", "application/json", strings.NewReader(requestString))

	// Check
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	urlStoreMock.AssertNumberOfCalls(t, "Create", 1)
	urlStoreMock.AssertNumberOfCalls(t, "Read", 1)
	linkingStore.AssertNumberOfCalls(t, "Create", 0)
	linkingStore.AssertNumberOfCalls(t, "Read", 0)
}

func Test_Expand_Get(t *testing.T) {
	// Mocking
	urlStoreMock := makeUrlStoreMock()
	linkingStore := makeLinkingStoreMock()

	// Setup PostsService
	router := mux.NewRouter()
	shortenService := url.NewService(urlStoreMock, linkingStore)
	shortenService.Register(router)

	// Setup server
	server := httptest.NewServer(router)
	defer server.Close()

	// Do request
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Get(server.URL + "/12345678")

	// Check
	assert.Equal(t, "http://www.test_2.org", res.Header.Get("Location"))

	urlStoreMock.AssertNumberOfCalls(t, "Create", 0)
	urlStoreMock.AssertNumberOfCalls(t, "Read", 1)
	linkingStore.AssertNumberOfCalls(t, "Read", 0)
}

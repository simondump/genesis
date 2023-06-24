package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

type UnauthorizedPostConfig struct {
	Url     string
	Body    string
	Handler func(*httptest.ResponseRecorder)
}

type AuthorizedGetConfig struct {
	Url     string
	Token   string
	Handler func(*httptest.ResponseRecorder)
}

type AuthorizedBodyConfig struct {
	Url     string
	Body    string
	Token   string
	Handler func(*httptest.ResponseRecorder)
}

type AuthorizedDeleteConfig struct {
	Url     string
	Token   string
	Handler func(*httptest.ResponseRecorder)
}

func tryUnauthorizedPost(config UnauthorizedPostConfig) {
	router := SetupRoutes()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", config.Url, strings.NewReader(config.Body))

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)
	config.Handler(response)
}

func tryAuthorizedPut(config AuthorizedBodyConfig) {
	router := SetupRoutes()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("PUT", config.Url, strings.NewReader(config.Body))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+config.Token)

	router.ServeHTTP(response, request)
	config.Handler(response)
}

func tryAuthorizedDelete(config AuthorizedDeleteConfig) {
	router := SetupRoutes()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("DELETE", config.Url, nil)

	request.Header.Set("Authorization", "Bearer "+config.Token)

	router.ServeHTTP(response, request)
	config.Handler(response)
}

func tryAuthorizedGet(config AuthorizedGetConfig) {
	router := SetupRoutes()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", config.Url, nil)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+config.Token)

	router.ServeHTTP(response, request)
	config.Handler(response)
}

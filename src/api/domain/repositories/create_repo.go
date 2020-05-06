package repositories

import (
	"golang-microservices/src/api/utils/errors"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}


type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int                         `json:"status"`
	Results    []CreateRespositoriesResult `json:"results"`
}

type CreateRespositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.ApiError     `json:"error"`
}
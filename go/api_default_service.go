/*
 * Message API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	infra "github.com/GIT_USER_ID/GIT_REPO_ID/go/infra/mysql"
	"github.com/GIT_USER_ID/GIT_REPO_ID/go/model"
	"github.com/GIT_USER_ID/GIT_REPO_ID/go/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// CreateMessage - Create a new message
// Saveを実行する
func (s *DefaultApiService) CreateMessage(ctx context.Context, newMessage NewMessage) (ImplResponse, error) {
	db, err := sqlx.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/go_practice?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var f repository.MessageRepository = &infra.MessageRepository{DB: db}

	message := &model.Message{Name: newMessage.Name, Message: newMessage.Message, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	f.Save(message)

	return Response(http.StatusOK, nil), nil
}

// GetMessages - Get all messages
// Listを実行する
func (s *DefaultApiService) GetMessages(ctx context.Context) (ImplResponse, error) {
	db, err := sqlx.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/go_practice?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var f repository.MessageRepository = &infra.MessageRepository{DB: db}

	messages, err := f.List()
	if err != nil {
		fmt.Println(err)
	}

	return Response(http.StatusOK, messages), nil
}

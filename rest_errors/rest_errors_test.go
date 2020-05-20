package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("test message", errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, "message: test message - status: 500 - error: internal_server_error - causes: [database error]", err.Error())
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("test message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, "message: test message - status: 400 - error: bad_request - causes: []", err.Error())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("test message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, "message: test message - status: 404 - error: not_found - causes: []", err.Error())
}

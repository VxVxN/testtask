package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/VxVxN/testtask/pkg/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateEventHandler(t *testing.T) {
	reqBody := "user_id=1&date=2025-10-22"
	req, err := http.NewRequest("POST", "/event", bytes.NewBufferString(reqBody))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createEventHandler)

	handler.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, "{\"result\":\"event created\"}\n", rr.Body.String())
}

func TestParseEventParams(t *testing.T) {
	reqBody := "user_id=1&date=2025-10-22"
	req, err := http.NewRequest("POST", "/event", bytes.NewBufferString(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	event, err := parseEventParams(req)
	require.NoError(t, err)

	expectedDate, _ := time.Parse("2006-01-02", "2025-10-22")
	assert.Equal(t, expectedDate, event.Date)
}

func TestCreateEventHandlerInvalidUserID(t *testing.T) {
	reqBody := "user_id=abc&date=2025-10-22"
	req, err := http.NewRequest("POST", "/event", bytes.NewBufferString(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createEventHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, `{"error": "invalid user_id"}`, rr.Body.String())
}

func TestGetEventsHandler(t *testing.T) {
	events = []event.Event{
		{UserID: 1, Date: time.Now()},
		{UserID: 2, Date: time.Now()},
		{UserID: 3, Date: time.Date(2024, 10, 23, 0, 0, 0, 0, time.UTC)},
	}

	req, err := http.NewRequest("GET", "/events?period=month", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getEventsHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	resp, err := io.ReadAll(rr.Result().Body)
	require.NoError(t, err)

	var response struct {
		Result []EventResponse `json:"result"`
	}
	err = json.Unmarshal(resp, &response)
	require.NoError(t, err)

	assert.Len(t, response.Result, 2)

	// todo add more testcases
}

func TestUpdateEventHandler(t *testing.T) {
	events = []event.Event{
		{UserID: 1, Date: time.Date(2025, 11, 22, 0, 0, 0, 0, time.UTC)},
	}

	reqBody := "user_id=1&date=2025-11-22"
	req, err := http.NewRequest("PUT", "/event", bytes.NewBufferString(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateEventHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"result\":\"event updated\"}\n", rr.Body.String())
}

func TestDeleteEventHandler(t *testing.T) {
	events = []event.Event{
		{UserID: 1, Date: time.Date(2025, 11, 22, 0, 0, 0, 0, time.UTC)},
	}

	reqBody := "user_id=1&date=2025-11-22"
	req, err := http.NewRequest("DELETE", "/event", bytes.NewBufferString(reqBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteEventHandler)

	handler.ServeHTTP(rr, req)

	resp, err := io.ReadAll(rr.Result().Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "{\"result\":\"event deleted\"}\n", string(resp))
}

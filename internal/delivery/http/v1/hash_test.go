package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHashHandler(t *testing.T) {
	tests := []struct {
		name          string
		inp           Body
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "Generate Hash - Status OK",
			inp:  Body{Text: "test"},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Generate Hash - Status Bad Request",
			inp:  Body{},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1")
			handler := NewHashHandler()
			handler.LoadRoutes(api)

			body, err := json.Marshal(tt.inp)
			require.NoError(t, err)

			req, err := http.NewRequest("POST", "/api/v1/hash/calc", bytes.NewReader(body))
			require.NoError(t, err)
			router.ServeHTTP(recorder, req)

			// Проверка ответа
			tt.checkResponse(t, recorder)
		})
	}
}

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/begenov/tsarka-task/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestEmailHandler_checkEmails(t *testing.T) {
	emails := util.RandomEmailsAndWords(10)

	check_emails := util.EmailsCheck(emails)

	tests := []struct {
		name          string
		inp           emailRequest
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "status ok",
			inp: emailRequest{
				Emails: emails,
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				var res responseEmails
				bytes, err := io.ReadAll(recoder.Body)
				require.NoError(t, err)
				err = json.Unmarshal(bytes, &res)
				require.NoError(t, err)
				require.Equal(t, check_emails, res.Emails)
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "status bad request",
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &InnEmailHandler{}
			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/check/email", server.URL)
			body, err := json.Marshal(tt.inp)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

func TestInnEmailHandler_checkInn(t *testing.T) {

	tests := []struct {
		name          string
		inp           innRequest
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder, inp []string)
	}{
		{
			name: "status ok",
			inp: innRequest{
				Inn: []string{
					"011203551253",
					"010403513242",
					"010403153242",
					"010401553242",
					"010403553242",
					"012403553242",
				},
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder, inp []string) {
				var res innResponse
				bytes, err := io.ReadAll(recoder.Body)
				require.NoError(t, err)
				require.NotEmpty(t, bytes)
				err = json.Unmarshal(bytes, &res)
				check_inp := util.InnCheck(inp)
				require.NoError(t, err)
				require.Equal(t, check_inp, res.Inn)
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "status bad request",
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder, inp []string) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &InnEmailHandler{}
			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/check/inn", server.URL)
			body, err := json.Marshal(tt.inp)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder, tt.inp.Inn)
		})
	}
}

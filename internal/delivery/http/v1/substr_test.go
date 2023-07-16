package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/begenov/tsarka-task/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestSubstrHandler_findSubstr(t *testing.T) {
	inp := util.RandomString(20)
	tests := []struct {
		name          string
		inp           requestSubstr
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "status ok",
			inp: requestSubstr{
				Text: inp,
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "status ok",
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
			handler := &SubstrHandler{}
			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/substr/find", server.URL)
			body, err := json.Marshal(tt.inp)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

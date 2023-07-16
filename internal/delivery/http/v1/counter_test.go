package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/begenov/tsarka-task/internal/service"

	mock_store "github.com/begenov/tsarka-task/internal/repository/mocks"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCouterHandler_addCounter(t *testing.T) {

	tests := []struct {
		name          string
		inp           counterRequest
		buildStubs    func(store *mock_store.MockCounters)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "status ok",
			inp: counterRequest{
				I: 1,
			},
			buildStubs: func(store *mock_store.MockCounters) {
				var i int64 = 1
				store.EXPECT().Add(gomock.Any(), gomock.Any()).Times(1).Return(i, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "status bad request",
			buildStubs: func(store *mock_store.MockCounters) {
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockCounters(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &CounterHandler{
				service: service.NewCounterervice(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/counter/add/%d", server.URL, tt.inp.I)
			request, err := http.NewRequest(http.MethodPost, url, nil)
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

func TestCounterHandler_subCounter(t *testing.T) {

	tests := []struct {
		name          string
		inp           counterRequest
		buildStubs    func(store *mock_store.MockCounters)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "status ok",
			inp: counterRequest{
				I: 1,
			},
			buildStubs: func(store *mock_store.MockCounters) {
				var i int64 = 1
				store.EXPECT().Sub(key, gomock.Any()).Times(1).Return(i, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "status bad request",
			inp: counterRequest{
				I: -1,
			},
			buildStubs: func(store *mock_store.MockCounters) {
				store.EXPECT().Sub(key, gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockCounters(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &CounterHandler{
				service: service.NewCounterervice(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/counter/sub/%d", server.URL, tt.inp.I)
			request, err := http.NewRequest(http.MethodPost, url, nil)
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

func TestCounterHandler_valCounter(t *testing.T) {
	tests := []struct {
		name string

		buildStubs    func(store *mock_store.MockCounters)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "status ok",
			buildStubs: func(store *mock_store.MockCounters) {
				var i int64 = 1
				store.EXPECT().Get(key).Times(1).Return(i, nil)

			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockCounters(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &CounterHandler{
				service: service.NewCounterervice(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/counter/val", server.URL)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

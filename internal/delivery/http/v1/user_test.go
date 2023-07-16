package v1

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/begenov/tsarka-task/internal/domain"
	mock_store "github.com/begenov/tsarka-task/internal/repository/mocks"
	"github.com/begenov/tsarka-task/internal/service"
	"github.com/begenov/tsarka-task/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_createUser(t *testing.T) {
	user := randUser()
	tests := []struct {
		name          string
		inp           domain.User
		buildStubs    func(store *mock_store.MockUsers)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		// TODO: Add test cases.
		{
			name: "status ok",
			inp:  user,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Eq(user)).Return(1, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "status bad request",
			inp: domain.User{
				LastName:  "",
				FirstName: "",
			},
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "status internal server",
			inp:  user,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(0, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockUsers(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &UserHandler{
				service: service.NewUserService(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/user", server.URL)
			body, err := json.Marshal(tt.inp)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

func randUser() domain.User {
	return domain.User{
		ID:        1,
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
	}
}

func TestUserHandler_getUser(t *testing.T) {
	user := randUser()
	tests := []struct {
		name          string
		inp           int
		buildStubs    func(store *mock_store.MockUsers)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		// TODO: Add test cases.
		{
			name: "status ok",
			inp:  1,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().GetUser(gomock.Any(), 1).Times(1).Return(user, nil)
			},

			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				var inp domain.User
				bytes, err := io.ReadAll(recoder.Body)
				require.NoError(t, err)
				err = json.Unmarshal(bytes, &inp)
				require.NoError(t, err)

				require.Equal(t, user, inp)
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name:       "invalid user ID",
			inp:        0,
			buildStubs: func(store *mock_store.MockUsers) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "user not found",
			inp:  2,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().GetUser(gomock.Any(), 2).Times(1).Return(domain.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "error retrieving user",
			inp:  3,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().GetUser(gomock.Any(), 3).Times(1).Return(domain.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockUsers(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &UserHandler{
				service: service.NewUserService(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/user/%d", server.URL, tt.inp)

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

func TestUserHandler_updateUser(t *testing.T) {
	user := randUser()
	tests := []struct {
		name          string
		id            int
		inp           domain.User
		buildStubs    func(store *mock_store.MockUsers)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		// TODO: Add test cases.
		{
			name: "successful update",
			id:   1,
			inp:  user,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().UpdateUser(gomock.Any(), user).Times(1).Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				var response domain.User
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				require.NoError(t, err)

				require.Equal(t, user, response)
			},
		},
		{
			name:       "invalid request",
			id:         0,
			inp:        user,
			buildStubs: func(store *mock_store.MockUsers) {},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
		{
			name: "user not found",
			id:   1,
			inp:  user,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().UpdateUser(gomock.Any(), user).Times(1).Return(domain.User{}, domain.NotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "invalid user ID",
			id:         0,
			buildStubs: func(store *mock_store.MockUsers) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "error updating user",
			id:   1,
			inp:  user,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().UpdateUser(gomock.Any(), user).Times(1).Return(domain.User{}, errors.New("database error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockUsers(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &UserHandler{
				service: service.NewUserService(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/user/%d", server.URL, tt.id)

			jsonBody, err := json.Marshal(tt.inp)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(jsonBody))
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

func TestUserHandler_deleteUser(t *testing.T) {
	tests := []struct {
		name          string
		id            int
		buildStubs    func(store *mock_store.MockUsers)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "successful deletion",
			id:   1,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().DeleteUser(gomock.Any(), 1).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},

		{
			name:       "invalid user ID",
			id:         0,
			buildStubs: func(store *mock_store.MockUsers) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},

		{
			name: "user not found",
			id:   2,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().DeleteUser(gomock.Any(), 2).Times(1).Return(domain.NotFound)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},

		{
			name: "error deleting user",
			id:   3,
			buildStubs: func(store *mock_store.MockUsers) {
				store.EXPECT().DeleteUser(gomock.Any(), 3).Times(1).Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mock_store.NewMockUsers(ctrl)
			recorder := httptest.NewRecorder()
			router := gin.Default()
			api := router.Group("/api/v1/rest")
			handler := &UserHandler{
				service: service.NewUserService(store),
			}
			tt.buildStubs(store)

			handler.LoadRoutes(api)
			server := httptest.NewServer(router)
			defer server.Close()
			url := fmt.Sprintf("%s/api/v1/rest/user/%d", server.URL, tt.id)

			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			router.ServeHTTP(recorder, request)
			tt.checkResponse(t, recorder)
		})
	}
}

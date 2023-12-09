package handlers_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"

	"github.com/ogniloud/madr/internal/handlers"
	handlerMocks "github.com/ogniloud/madr/internal/handlers/mocks"
	"github.com/ogniloud/madr/internal/ioutil"
	"github.com/ogniloud/madr/internal/models"
)

func TestEndpoints_GetUserInfo(t *testing.T) {

	type mocks struct {
		data   *handlerMocks.Datalayer
		logger *handlerMocks.Logger
	}

	prepareRequest := func(r *http.Request, userID string) *http.Request {
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", userID)

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
		return r
	}

	tests := []struct {
		name         string
		mock         mocks
		prepareMocks func(mocks mocks)
		request      *http.Request
		wantBody     string
		wantStatus   int
	}{
		{
			name: "User id is less than 1",
			mock: mocks{
				data: handlerMocks.NewDatalayer(t),
			},
			prepareMocks: func(mocks mocks) {
				return
			},
			request: prepareRequest(httptest.NewRequest(http.MethodGet, "/api/user/{id}", nil), "0"),
			wantBody: `{"message":"User id cannot be less than 1"}
`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "User id is not an integer",
			mock: mocks{
				data:   handlerMocks.NewDatalayer(t),
				logger: handlerMocks.NewLogger(t),
			},
			prepareMocks: func(mocks mocks) {
				mocks.logger.EXPECT().Error("Unable to convert user id to integer", "error", mock.Anything)
			},
			request: prepareRequest(httptest.NewRequest(http.MethodGet, "/api/user/{id}", nil), "buba"),
			wantBody: `{"message":"don't worry, we are working on it"}
`,
			wantStatus: http.StatusInternalServerError,
		},
		{
			name: "User not found",
			mock: mocks{
				data:   handlerMocks.NewDatalayer(t),
				logger: handlerMocks.NewLogger(t),
			},
			prepareMocks: func(mocks mocks) {
				mocks.data.EXPECT().GetUserById(mock.Anything, 15).Return(models.UserInfo{}, models.ErrUserNotFound)
			},
			request: prepareRequest(httptest.NewRequest(http.MethodGet, "/api/user/{id}", nil), "15"),
			wantBody: `{"message":"user with specified id not found"}
`,
			wantStatus: http.StatusNotFound,
		},
		{
			name: "User found, happy path",
			mock: mocks{
				data:   handlerMocks.NewDatalayer(t),
				logger: handlerMocks.NewLogger(t),
			},
			prepareMocks: func(mocks mocks) {
				mocks.data.EXPECT().GetUserById(mock.Anything, 228).Return(models.UserInfo{
					ID:       228,
					Username: "idiot",
					Email:    "sobaka@govno.com",
				}, nil)
			},
			request: prepareRequest(httptest.NewRequest(http.MethodGet, "/api/user/{id}", nil), "228"),
			wantBody: `{"id":228,"username":"idiot","email":"sobaka@govno.com"}
`,
			wantStatus: http.StatusOK,
		},
		{
			name: "Error from datalayer",
			mock: mocks{
				data:   handlerMocks.NewDatalayer(t),
				logger: handlerMocks.NewLogger(t),
			},
			prepareMocks: func(mocks mocks) {
				mocks.data.EXPECT().GetUserById(mock.Anything, 144).Return(models.UserInfo{}, fmt.Errorf("some kind of different error"))
				mocks.logger.EXPECT().Error("unable to get user info in GetUserInfo", "error", fmt.Errorf("some kind of different error"))
			},
			request: prepareRequest(httptest.NewRequest(http.MethodGet, "/api/user/{id}", nil), "144"),
			wantBody: `{"message":"don't worry, we are working on it"}
`,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := handlers.New(tt.mock.data, ioutil.JSONErrorWriter{Logger: log.Default()}, tt.mock.logger)

			tt.prepareMocks(tt.mock)

			response := httptest.NewRecorder()
			s.GetUserInfo(response, tt.request)

			// response body
			got := response.Body.String()
			want := tt.wantBody

			// response status code
			gotStatus := response.Code
			wantStatus := tt.wantStatus

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}

			if gotStatus != wantStatus {
				t.Errorf("got %d, want %d", gotStatus, wantStatus)
			}
		})
	}

}

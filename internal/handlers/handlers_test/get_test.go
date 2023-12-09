package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/ogniloud/madr/internal/handlers"
)

func TestEndpoints_GetUserInfo(t *testing.T) {

	s := handlers.New(nil, nil)

	prepareRequest := func(r *http.Request, userID string) *http.Request {
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", userID)

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
		return r
	}

	tests := []struct {
		name       string
		request    *http.Request
		wantBody   string
		wantStatus int
	}{
		{
			name:    "User id is less than 1",
			request: prepareRequest(httptest.NewRequest(http.MethodGet, "/api/user/{id}", nil), "0"),
			wantBody: `{"message":"User id cannot be less than 1"}
`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

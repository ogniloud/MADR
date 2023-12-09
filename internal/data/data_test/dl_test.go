package data_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ogniloud/madr/internal/data"
	dataLayerMocks "github.com/ogniloud/madr/internal/data/mocks"
)

func Test_GetUserById(t *testing.T) {
	type mocks struct {
		userCredentials *dataLayerMocks.UserCredentials
	}

	type args struct {
		userID int
	}

	type want struct {
		id       int
		username string
		email    string
		err      error
	}

	test := []struct {
		name    string
		mocks   mocks
		args    args
		want    want
		prepare func(mocks mocks)
	}{
		{
			name: "error from database",
			mocks: mocks{
				userCredentials: dataLayerMocks.NewUserCredentials(t),
			},
			args: args{
				userID: 15,
			},
			want: want{
				id:       0,
				username: "",
				email:    "",
				err:      fmt.Errorf("unable to get user info in GetUserById: some kind of error"),
			},
			prepare: func(mocks mocks) {
				mocks.userCredentials.EXPECT().GetUserInfo(context.Background(), 15).
					Return("", "", fmt.Errorf("some kind of error"))
			},
		},
		{
			name: "happy path",
			mocks: mocks{
				userCredentials: dataLayerMocks.NewUserCredentials(t),
			},
			args: args{
				userID: 17,
			},
			want: want{
				id:       17,
				username: "koko",
				email:    "kuku@mail.ru",
				err:      nil,
			},
			prepare: func(mocks mocks) {
				mocks.userCredentials.EXPECT().GetUserInfo(context.Background(), 17).Return("koko", "kuku@mail.ru", nil)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(tt.mocks)

			dl := data.New(tt.mocks.userCredentials, 1, time.Minute, []byte{})

			userInfo, err := dl.GetUserById(context.Background(), tt.args.userID)

			if err != tt.want.err {
				if err == nil || tt.want.err == nil {
					t.Errorf("got %v, want %v", err, tt.want.err)
				} else if err.Error() != tt.want.err.Error() {
					t.Errorf("got %v, want %v", err, tt.want.err)
				}
			}

			if userInfo.Username != tt.want.username {
				t.Errorf("got %v, want %v", err, tt.want.username)
			}

			if userInfo.Email != tt.want.email {
				t.Errorf("got %v, want %v", err, tt.want.email)
			}

		})
	}
}

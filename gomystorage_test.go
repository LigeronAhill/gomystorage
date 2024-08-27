package gomystorage

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		token string
	}
	token := os.Getenv("MS_TOKEN")
	t.Log(token)
	wantUrl, err := url.Parse("https://api.moysklad.ru/api/remap/1.2")
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		name string
		args args
		want *ApiClient
	}{
		{
			name: "wrong token",
			args: args{token: "someToken"},
			want: &ApiClient{
				baseUrl: *wantUrl,
				token:   "Bearer someToken",
				client:  http.Client{},
			},
		},
		{
			name: "correct token",
			args: args{token: token},
			want: &ApiClient{
				baseUrl: *wantUrl,
				token:   "Bearer " + token,
				client:  http.Client{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

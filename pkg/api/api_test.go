package api

import (
	"github.com/gorilla/mux"
	"got_news/pkg/storage"
	"net/http"
	"reflect"
	"testing"
)

func TestAPI_Router(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
		want   *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			if got := api.Router(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_endpoints(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestAPI_postsHandlerNItems(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		db storage.Interface
	}
	tests := []struct {
		name string
		args args
		want *API
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

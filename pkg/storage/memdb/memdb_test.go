package memdb

import (
	"got_news/pkg/storage"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Store
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_AddPost(t *testing.T) {
	type args struct {
		in0 storage.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.AddPost(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_DeletePost(t *testing.T) {
	type args struct {
		in0 storage.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.DeletePost(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_PostsNItems(t *testing.T) {
	type args struct {
		in0 storage.Post
	}
	tests := []struct {
		name    string
		args    args
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.PostsNItems(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostsNItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostsNItems() got = %v, want %v", got, tt.want)
			}
		})
	}
}
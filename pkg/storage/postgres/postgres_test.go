package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"got_news/pkg/storage"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		constr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Storage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.constr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddPost(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.AddPost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeletePost(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeletePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_PostsNItems(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.PostsNItems(tt.args.p)
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

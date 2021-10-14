package rss

import (
	"got_news/pkg/storage"
	"testing"
)

func TestFetchRss(t *testing.T) {
	type args struct {
		link     string
		rPeriod  int64
		postChan chan<- storage.Post
		errChan  chan<- error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
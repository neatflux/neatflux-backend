package v1_test

import (
	"context"
	"testing"

	v1 "github.com/neatflux/neatflux-backend/pkg/api/v1"
)

func TestGetVideoInfo(t *testing.T) {
	s := server{}

	// TESTS CASES
	tests := []struct {
		apiVersion string
		vidID      int64
		wantTitle  string
	}{
		{
			apiVersion: "v1",
			vidID:      1,
			wantTitle:  "Cute cat video",
		},
	}

	for _, tt := range tests {
		req := &v1.VideoInfoRequest{ApiVersion: tt.apiVersion, Id: tt.vidID}
		resp, err := s.GetVideoInfo(context.Background(), req)
		if err != nil {
			t.Errorf("GetVideoInfo(%v) got unexpected error")
		}
		if resp.Message != tt.wantTitle {
			t.Errorf("GetVideoInfo(%v)=%v, wanted %v", tt.vidID, resp.Message, tt.wantTitle)
		}
	}
}

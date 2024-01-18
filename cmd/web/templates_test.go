package main

import (
	"testing"
	"time"

	"snippetbox.pawelkrml.com/internal/assert"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 3, 24, 12, 15, 0, 0, time.UTC),
			want: "24 Mar 2023 at 12:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 3, 24, 12, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "24 Mar 2023 at 11:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}

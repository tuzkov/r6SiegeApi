package r6_test

import (
	"testing"

	r6 "github.com/tuzkov/r6SiegeApi"
)

func TestRankFromMMR(t *testing.T) {
	tests := []struct {
		name string
		mmr  float32
		want int
	}{
		{
			name: "copper 4-1",
			mmr:  1000,
			want: 1,
		},
		{
			name: "copper 4-2",
			mmr:  1399,
			want: 1,
		},
		{
			name: "copper 3-1",
			mmr:  1400,
			want: 2,
		},
		{
			name: "copper 3-2",
			mmr:  1440,
			want: 2,
		},
		{
			name: "copper 3-3",
			mmr:  1499,
			want: 2,
		},
		{
			name: "copper 2",
			mmr:  1500,
			want: 3,
		},
		{
			name: "bronze 1",
			mmr:  2050,
			want: 8,
		},
		{
			name: "silver 2",
			mmr:  2350,
			want: 11,
		},
		{
			name: "silver 1",
			mmr:  2499,
			want: 12,
		},
		{
			name: "gold 4-1",
			mmr:  2500,
			want: 13,
		},
		{
			name: "gold 4-2",
			mmr:  2650,
			want: 13,
		},
		{
			name: "gold 4-3",
			mmr:  2699,
			want: 13,
		},
		{
			name: "gold 3-1",
			mmr:  2700,
			want: 14,
		},
		{
			name: "gold 3-2",
			mmr:  2800,
			want: 14,
		},
		{
			name: "gold 3-3",
			mmr:  2899,
			want: 14,
		},
		{
			name: "gold 2-1",
			mmr:  2900,
			want: 15,
		},
		{
			name: "gold 2-2",
			mmr:  2950,
			want: 15,
		},
		{
			name: "gold 1",
			mmr:  3299,
			want: 16,
		},
		{
			name: "plat 3-1",
			mmr:  3300,
			want: 17,
		},
		{
			name: "plat 3-2",
			mmr:  3600,
			want: 17,
		},
		{
			name: "plat 3-3",
			mmr:  3699,
			want: 17,
		},
		{
			name: "plat 2-1",
			mmr:  3700,
			want: 18,
		},
		{
			name: "plat 1",
			mmr:  4499,
			want: 19,
		},
		{
			name: "diamond 1",
			mmr:  4500,
			want: 20,
		},
		{
			name: "diamond 2",
			mmr:  5000,
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := r6.RankFromMMR(tt.mmr); got != tt.want {
				t.Errorf("RankFromMMR() = %v, want %v", got, tt.want)
			}
		})
	}
}

package r6

import (
	"testing"
)

func TestRankBracketNew(t *testing.T) {
	tests := []struct {
		name     string
		rank     int
		expected string
	}{
		{
			"unranked",
			0,
			"Unranked",
		},
		{
			"copper 5",
			1,
			"Copper 5",
		},
		{
			"copper 1",
			5,
			"Copper 1",
		},
		{
			"Bronze 5",
			6,
			"Bronze 5",
		},
		{
			"Bronze 1",
			10,
			"Bronze 1",
		},
		{
			"Silver 5",
			11,
			"Silver 5",
		},
		{
			"Silver 1",
			15,
			"Silver 1",
		},
		{
			"Gold 3",
			16,
			"Gold 3",
		},
		{
			"Gold 1",
			18,
			"Gold 1",
		},
		{
			"Platinum 3",
			19,
			"Platinum 3",
		},
		{
			"Platinum 1",
			21,
			"Platinum 1",
		},
		{
			"Diamond",
			22,
			"Diamond",
		},
		{
			"Champion",
			23,
			"Champion",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if act := RankBracketNew(test.rank); act != test.expected {
				t.Errorf("rankBracketNew %s != %s", act, test.expected)
			}
		})
	}
}

func TestRankBracketNewEmoji(t *testing.T) {
	tests := []struct {
		name     string
		rank     int
		expected string
	}{
		{
			"unranked",
			0,
			"\U0000274C Unranked",
		},
		{
			"copper 5",
			1,
			"\U0001F4A9 Copper 5",
		},
		{
			"copper 1",
			5,
			"\U0001F4A9 Copper 1",
		},
		{
			"Bronze 5",
			6,
			"\U0001F949 Bronze 5",
		},
		{
			"Bronze 1",
			10,
			"\U0001F949 Bronze 1",
		},
		{
			"Silver 5",
			11,
			"\U0001F948 Silver 5",
		},
		{
			"Silver 1",
			15,
			"\U0001F948 Silver 1",
		},
		{
			"Gold 3",
			16,
			"\U0001F947 Gold 3",
		},
		{
			"Gold 1",
			18,
			"\U0001F947 Gold 1",
		},
		{
			"Platinum 3",
			19,
			"\U0001F3C6 Platinum 3",
		},
		{
			"Platinum 1",
			21,
			"\U0001F3C6 Platinum 1",
		},
		{
			"Diamond",
			22,
			"\U0001F48E Diamond",
		},
		{
			"Champion",
			23,
			"\U0001F480 Champion",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if act := RankBracketNewEmoji(test.rank); act != test.expected {
				t.Errorf("rankBracketNewEmoji %s != %s", act, test.expected)
			}
		})
	}
}

func TestRankFromMMRNew(t *testing.T) {
	tests := []struct {
		name     string
		mmr      float32
		expected int
	}{
		{
			"copper 5",
			1000,
			1,
		},
		{
			"copper 4",
			1200,
			2,
		},
		{
			"copper 4",
			1299,
			2,
		},
		{
			"copper 3",
			1300,
			3,
		},
		{
			"copper 3",
			1350,
			3,
		},
		{
			"copper 3",
			1399,
			3,
		},
		{
			"copper 2",
			1400,
			4,
		},
		{
			"copper 2",
			1499,
			4,
		},
		{
			"copper 1",
			1500,
			5,
		},
		{
			"copper 1",
			1599,
			5,
		},
		{
			"bronze 5",
			1600,
			6,
		},
		{
			"bronze 3",
			1850,
			8,
		},
		{
			"bronze 1",
			2099,
			10,
		},
		{
			"silver 5",
			2100,
			11,
		},
		{
			"silver 3",
			2350,
			13,
		},
		{
			"silver 1",
			2599,
			15,
		},
		{
			"gold 3",
			2600,
			16,
		},
		{
			"gold 3",
			2799,
			16,
		},
		{
			"gold 2",
			2800,
			17,
		},
		{
			"gold 2",
			2999,
			17,
		},
		{
			"gold 1",
			3000,
			18,
		},
		{
			"gold 1",
			3199,
			18,
		},
		{
			"plat 3",
			3200,
			19,
		},
		{
			"plat 3",
			3599,
			19,
		},
		{
			"plat 2",
			3600,
			20,
		},
		{
			"plat 2",
			3999,
			20,
		},
		{
			"plat 1",
			4000,
			21,
		},
		{
			"plat 1",
			4399,
			21,
		},
		{
			"diamond",
			4400,
			22,
		},
		{
			"diamond",
			4499,
			22,
		},
		{
			"champion",
			5000,
			23,
		},
		{
			"champion",
			8000,
			23,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if act := RankFromMMRNew(test.mmr); act != test.expected {
				t.Errorf("RankFromMMRNew() = %d, expected %d", act, test.expected)
			}
		})
	}
}

func TestRankFromMMROld(t *testing.T) {
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
			if got := RankFromMMROld(tt.mmr); got != tt.want {
				t.Errorf("RankFromMMROld() = %v, want %v", got, tt.want)
			}
		})
	}
}

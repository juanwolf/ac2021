package main

import (
	"testing"
)

func TestGamma(t *testing.T) {
	tests := []struct {
		i    []string
		want string
	}{
		{i: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want: "10110"},
	}

	for _, test := range tests {
		got := Gamma(test.i)
		if test.want != got {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}

func TestEpsilon(t *testing.T) {
	tests := []struct {
		i    []string
		want string
	}{
		{i: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want: "01001"},
	}

	for _, test := range tests {
		got := Epsilon(test.i)
		if test.want != got {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}

func TestConvertStringBytesToInt(t *testing.T) {
	tests := []struct {
		i    string
		want int64
	}{
		{i: "10110", want: 22},
		{i: "01001", want: 9},
	}
	for _, test := range tests {
		got := convertStringBytesToInt(test.i)
		if test.want != got {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}
}

func TestOxygenGeneratorRating(t *testing.T) {
	tests := []struct {
		i    []string
		want string
	}{
		{i: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want: "10111"},
	}

	for _, test := range tests {
		got := OxygenGeneratorRating(test.i)
		if test.want != got {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}

}

func TestCO2Scrubber(t *testing.T) {
	tests := []struct {
		i    []string
		want string
	}{
		{i: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want: "01010"},
	}

	for _, test := range tests {
		got := CO2Scrubber(test.i)
		if test.want != got {
			t.Errorf("Expected %+v, got: %+v", test.want, got)
		}
	}

}

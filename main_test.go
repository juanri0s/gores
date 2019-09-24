package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrint(t *testing.T) {
	type St struct {
		Name            string `json:"name"`
		IsLookingForJob bool   `json:"isLookingForJob"`
	}
	tests := []struct {
		name     string
		expected string
		actual   St
	}{
		{"prettyprint from struct test 1", prettyPrint(St{"Joe", true}), St{"Joe", true}},
		{"prettyprint from struct test 2", prettyPrint(St{"Janice", false}), St{"Janice", false}},
		{"prettyprint from struct test 3", prettyPrint(St{"Jerry", true}), St{"Jerry", true}},
	}
	for _, tt := range tests {

		res, _ := json.MarshalIndent(tt.actual, "", "  ")
		actual := string(res)

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, actual, tt.expected, "The prettyprint should be the same as the marshallindent")
		})
	}
}

func TestParseYearToDate(t *testing.T) {
	type args struct {
		s string
	}

	t1 := args{"2020"}
	t2 := args{"2006"}
	t3 := args{"1999"}

	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"string 2020 to datetime", t1, time.Date(2020, time.January, 01, 00, 00, 00, 00, time.UTC)},
		{"string 2006 to datetime", t2, time.Date(2006, time.January, 01, 00, 00, 00, 00, time.UTC)},
		{"string 1999 to datetime", t3, time.Date(1999, time.January, 01, 00, 00, 00, 00, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := parseYearToDate(tt.args.s)
			assert.Equal(t, want, actual, "The string year should be the same as the datetime.")
		})
	}
}

func TestParseMonthYearToDate(t *testing.T) {
	type args struct {
		s string
	}

	t1 := args{"January 2006"}
	t2 := args{"December 2010"}
	t3 := args{"March 2019"}

	tests := []struct {
		name string
		args args
		want time.Time
	}{

		{"string March 2020 to datetime", t1, time.Date(2006, time.January, 01, 00, 00, 00, 00, time.UTC)},
		{"string February2006 to datetime", t2, time.Date(2010, time.December, 01, 0, 0, 0, 0, time.UTC)},
		{"string December2020 to datetime", t3, time.Date(2019, time.March, 01, 0, 0, 0, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			actual := parseMonthYearToDate(tt.args.s)
			assert.Equal(t, want, actual, "The string monthYear should be the same as the datetime.")
		})
	}
}

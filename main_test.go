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
		name string
		expected string
		actual St
	}{
		{
			"prettyprint from struct test 1",
			PrettyPrint( St {
					"Joe",
					true,
				},
			),
			St {
				"Joe",
				true,
			},
		},
		{
			"prettyprint from struct test 2",
			PrettyPrint( St {
					"Janice",
					false,
				},
			),
			St {
				"Janice",
				false,
			},
		},
		{
			"prettyprint from struct test 3",
			PrettyPrint( St {
					"Jerry",
					true,
				},
			),
			St {
				"Jerry",
				true,
			},
		},
	}
	for _, test := range tests {

		res, _ := json.MarshalIndent(test.actual, "", "  ")
		actual := string(res)

		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, actual, test.expected, "The prettyprint should be the same as the marshallindent")
		})
	}
}

func TestParseYearToDate(t *testing.T) {
	t1, _ := time.Parse("2006", "2020")
	t2, _ := time.Parse("2006", "2006")
	t3, _ := time.Parse("2006", "1999")

	tests := []struct {
		name     string
		actual   time.Time
		expected time.Time
	}{
		{
			"string 2020 to datetime",
			ParseYearToDate("2020"),
			t1,
		},
		{
			"string 2005 to datetime",
			ParseYearToDate("2006"),
			t2,
		},
		{
			"string 1999 to datetime",
			ParseYearToDate("1999"),
			t3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.actual, test.expected, "The string year should be the same as the datetime.")
		})
	}
}

func TestParseMonthYearToDate(t *testing.T) {
	t1, _ := time.Parse("January 2006", "March 2020")
	t2, _ := time.Parse("January 2006", "February 2006")
	t3, _ := time.Parse("January 2006", "December 1999")

	tests := []struct {
		name     string
		actual   time.Time
		expected time.Time
	}{
		{
			"string March 2020 to datetime",
			ParseMonthYearToDate("March 2020"),
			t1,
		},
		{
			"string February2006 to datetime",
			ParseMonthYearToDate("February 2006"),
			t2,
		},
		{
			"string December2020 to datetime",
			ParseMonthYearToDate("December 1999"),
			t3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.actual, test.expected, "The string monthYear should be the same as the datetime.")
		})
	}
}


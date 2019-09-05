package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

type PP struct {
	Name            string `json:"name"`
	IsLookingForJob bool   `json:"isLookingForJob"`
}

var tests = []PP{
	{"Joe", true},
	{"Janice", false},
	{"Jerry", true},
}

func TestPrettyPrint(t *testing.T) {

	for _, test  := range tests {

		res, _ := json.MarshalIndent(test, "", "  ")
		want := string(res)

		got := PrettyPrint(test)

		if !reflect.DeepEqual(got, want) {
			t.Error(
				"For", test,
				"expected", want,
				"got", got,
			)
		}
	}
}

package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	input := []http.Header{
		{},
		{},
		{},
		{},
		{},
	}

	input[1].Add("Authorization", "")
	input[2].Add("Authorization", "APIKEY")
	input[3].Add("Authorization", "ApiKey fenaugi neuiang")
	input[4].Add("Authorization", "ApiKey dgsauigufehuw")

	wantVal := []string{
		"",
		"",
		"",
		"fenaugi",
		"dgsauigufehuw",
	}

	wantErr := []error{
		errors.New("no authorization header included"),
		errors.New("no authorization header included"),
		errors.New("malformed authorization header"),
		nil,
		nil,
	}

	for i := range input {
		gotVal, gotErr := GetAPIKey(input[i])
		if gotErr == nil || wantErr[i] == nil {
			if gotVal != wantVal[i] || gotErr != wantErr[i] {
				t.Fatalf("%d: wanted %v, %v but got %v, %v", i, wantVal[i], wantErr[i], gotVal, gotErr)
			}
			continue
		}
		if gotVal != wantVal[i] || gotErr.Error() != wantErr[i].Error() {
			t.Fatalf("%d: wanted %v, %v but got %v, %v", i, wantVal[i], wantErr[i], gotVal, gotErr)
		}
	}
}

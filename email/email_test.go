package validators

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		{
			"test.", false,
		},
		{
			"test@test.com", true,
		},
		{
			"test.com", false,
		},
		{
			"test@test", false,
		},
		{
			"test@test.", false,
		},
		{
			"test@test.com.", false,
		},
		{
			"", false,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("email %v", tc.email), func(t *testing.T) {
			got := IsValid(tc.email)
			if got != tc.want {
				t.Fatalf("IsValid(%v) = %v; want %v", tc.email, got, tc.want)
			}
		})
	}
}

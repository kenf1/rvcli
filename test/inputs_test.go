package test

import (
	"strings"
	"testing"

	"github.com/kenf1/rvcli/logic"
)

func TestCheckInputs(t *testing.T) {
	validUser := logic.UserConfig{
		Username: "user",
		Password: "pass",
		Fullname: "Full Name",
		Email:    "email@example.com",
	}
	validApi := logic.ApiConfig{
		Host: "localhost",
	}

	tests := []struct {
		name       string
		user       logic.UserConfig
		api        logic.ApiConfig
		wantErr    bool
		wantSubstr string
	}{
		{
			name:    "All valid",
			user:    validUser,
			api:     validApi,
			wantErr: false,
		},
		{
			name:       "Missing Username",
			user:       logic.UserConfig{Password: "pass", Fullname: "Full Name", Email: "email@example.com"},
			api:        validApi,
			wantErr:    true,
			wantSubstr: "Username is missing",
		},
		{
			name:       "Missing Password",
			user:       logic.UserConfig{Username: "user", Fullname: "Full Name", Email: "email@example.com"},
			api:        validApi,
			wantErr:    true,
			wantSubstr: "Password is missing",
		},
		{
			name:       "Missing Fullname",
			user:       logic.UserConfig{Username: "user", Password: "pass", Email: "email@example.com"},
			api:        validApi,
			wantErr:    true,
			wantSubstr: "Fullname is missing",
		},
		{
			name:       "Missing Email",
			user:       logic.UserConfig{Username: "user", Password: "pass", Fullname: "Full Name"},
			api:        validApi,
			wantErr:    true,
			wantSubstr: "Email is missing",
		},
		{
			name:       "Missing Host",
			user:       validUser,
			api:        logic.ApiConfig{},
			wantErr:    true,
			wantSubstr: "Host is missing",
		},
		{
			name:       "Multiple missing",
			user:       logic.UserConfig{},
			api:        logic.ApiConfig{},
			wantErr:    true,
			wantSubstr: "Username is missing; Password is missing; Fullname is missing; Email is missing; Host is missing",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.CheckInputs(tt.user, tt.api)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if !strings.Contains(err.Error(), tt.wantSubstr) {
					t.Errorf("expected error to contain %q, got %q", tt.wantSubstr, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}

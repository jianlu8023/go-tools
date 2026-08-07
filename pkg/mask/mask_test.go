package mask

import "testing"

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{"normal email", "user@example.com", "***@example.com"},
		{"empty string", "", "***masked***"},
		{"no @ symbol", "userexample.com", "***masked***"},
		{"complex email", "john.doe+tag@mail.example.co.uk", "***@mail.example.co.uk"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskEmail(tt.email); got != tt.want {
				t.Errorf("MaskEmail(%q) = %q, want %q", tt.email, got, tt.want)
			}
		})
	}
}

func TestMaskSensitiveInfo(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "simple url",
			str:  "http://example.com",
			want: "http://***.com",
		},
		{
			name: "url with path and query",
			str:  "https://api.test.org/v1/users/123?key=secret",
			want: "https://***.org/***/***/***?key=***",
		},
		{
			name: "country code tld url",
			str:  "https://sub.domain.co.uk/path/to/resource",
			want: "https://***.***.co.uk/***/***/***",
		},
		{
			name: "plain ip",
			str:  "192.168.1.1",
			want: "***.***.***.***",
		},
		{
			name: "plain domain",
			str:  "openai.com",
			want: "***.com",
		},
		{
			name: "www domain",
			str:  "www.openai.com",
			want: "***.***.com",
		},
		{
			name: "api subdomain",
			str:  "api.openai.com",
			want: "***.***.com",
		},
		{
			name: "api key in single quotes",
			str:  "'api_key:AIzaSyAAAaUooTUni8AdaOkSRMda30n_Q4vrV70'",
			want: "'api_key:***'",
		},
		{
			name: "api key in double quotes",
			str:  `"api_key:AIzaSyAAAaUooTUni8AdaOkSRMda30n_Q4vrV70"`,
			want: `"api_key:***"`,
		},
		{
			name: "plain text without sensitive info",
			str:  "this is a normal text",
			want: "this is a normal text",
		},
		{
			name: "empty string",
			str:  "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskSensitiveInfo(tt.str); got != tt.want {
				t.Errorf("MaskSensitiveInfo(%q) = %q, want %q", tt.str, got, tt.want)
			}
		})
	}
}

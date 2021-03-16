package config

import "testing"

func TestParse(t *testing.T) {
	got := parseYaml([]byte("id: 1\ntoken: a_token_string"))
	if got.ID != 1 {
		t.Errorf("parseYaml(): Config.ID = %d, expected %d", got.ID, 1)
	}
	if got.Token != "a_token_string" {
		t.Errorf("parseYaml(): Config.Token = %s, want 'a_token_string'", got.Token)
	}
}

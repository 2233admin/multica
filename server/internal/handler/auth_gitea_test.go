package handler

import (
	"testing"
)

func TestGiteaRedirectURIUsesConfiguredValue(t *testing.T) {
	t.Setenv("GITEA_REDIRECT_URI", "http://multica.local/auth/callback")

	got, err := giteaRedirectURI("http://multica.local/auth/callback")
	if err != nil {
		t.Fatalf("giteaRedirectURI returned error: %v", err)
	}
	if got != "http://multica.local/auth/callback" {
		t.Fatalf("redirect URI = %q", got)
	}
}

func TestGiteaRedirectURIRejectsMismatch(t *testing.T) {
	t.Setenv("GITEA_REDIRECT_URI", "http://multica.local/auth/callback")

	if _, err := giteaRedirectURI("http://evil.local/auth/callback"); err == nil {
		t.Fatal("expected redirect URI mismatch to fail")
	}
}

func TestGiteaBackendURLUsesIssuer(t *testing.T) {
	t.Setenv("GITEA_ISSUER_URL", "http://gitea.local/")
	t.Setenv("GITEA_BACKEND_HOST_OVERRIDE", "")

	if got := giteaBackendURL("/api/v1/user"); got != "http://gitea.local/api/v1/user" {
		t.Fatalf("backend URL = %q", got)
	}
}

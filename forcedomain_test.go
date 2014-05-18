package forcedomain

import (
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_NoRedirectWithoutDomainEnv(t *testing.T) {
	recorder := httptest.NewRecorder()
	m := martini.New()
	m.Use(ForceDomainRedirect())
	r, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(recorder, r)
	if recorder.Code != http.StatusOK {
		t.Error("There was redirect without DOMAIN env")
	}
}

func Test_RedirectWithDomainEnvAndUriScheme(t *testing.T) {
	os.Setenv("DOMAIN", "https://example.com")
	recorder := httptest.NewRecorder()
	m := martini.New()
	m.Use(ForceDomainRedirect())
	r, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(recorder, r)
	if recorder.Code != http.StatusTemporaryRedirect {
		t.Error("There was no redirect")
	}
	if recorder.HeaderMap["Location"][0] != "https://example.com" {
		t.Error("There was redirect but to wrong url")
	}
}

func Test_RedirectWithDomainEnvAndNoURIScheme(t *testing.T) {
	os.Setenv("DOMAIN", "example.com")
	recorder := httptest.NewRecorder()
	m := martini.New()
	m.Use(ForceDomainRedirect())
	r, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(recorder, r)
	if recorder.Code != http.StatusTemporaryRedirect {
		t.Error("There was no redirect")
	}
	if recorder.HeaderMap["Location"][0] != "http://example.com" {
		t.Error("There was redirect but to wrong url")
	}
}

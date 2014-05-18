package forcedomain

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-martini/martini"
)

func ForceDomainRedirect() martini.Handler {

	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		if domain := os.Getenv("DOMAIN"); len(domain) != 0 {
			domain = addOptionalPrefix(domain)
			http.Redirect(w, r, domain, http.StatusTemporaryRedirect)
		}
		return
	}
}

func addOptionalPrefix(domain string) string {
	if strings.HasPrefix(domain, "http://") || strings.HasPrefix(domain, "https://") {
		return domain
	} else {
		prefixed_domain := strings.Join([]string{"http://", domain}, "")
		return prefixed_domain
	}

}

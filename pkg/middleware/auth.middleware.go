package middleware

import (
	"log"
	"net/http"

	"github.com/KenethSandoval/fvexpress/internal/auth"
	"github.com/KenethSandoval/fvexpress/pkg/listening"
)

// ValidateMiddleware function, which will be called for each request
func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		jwtW := auth.JwtWrapper{
			SecretKey: "ponerenv",
			Issuer:    "AuthService",
		}

		claim, err := jwtW.ValidaJWT(token)
		if err != nil {
			log.Printf("%v", err)
		}

		if claim != nil {
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI, string(listening.ColorGreen), r.Method)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

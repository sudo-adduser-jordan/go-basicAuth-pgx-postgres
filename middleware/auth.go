package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"main/database"
	"main/utils"
	"net/http"
)

// Basic auth with bycrypt store
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))

			expectedUsernameHash := sha256.Sum256([]byte(database.ReadUser(username)))
			var expectedPasswordHash [32]byte
			if utils.CheckPasswordHash(password, database.ReadPassword(username)) {
				expectedPasswordHash = sha256.Sum256([]byte(password))
			}

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

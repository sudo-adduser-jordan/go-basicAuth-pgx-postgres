package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (record *statusRecorder) WriteHeader(code int) {
	record.status = code
	record.ResponseWriter.WriteHeader(code)
}

func Logger(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Initialize the status to 500 in case WriteHeader is not called
		record := statusRecorder{w, 500}
		defer func() {
			log.Printf(
				"| %s |	%s|	%s %s %s",
				statusColor(record.status),
				time.Since(start),
				styles.GreenText(r.Host),
				methodColor(r.Method),
				r.URL.Path,
			)

		}()
		f.ServeHTTP(&record, r)
	})
}

func statusColor(status int) lipgloss.Style {
	s := strconv.Itoa(status)
	s = fmt.Sprintf(" " + s + " ")

	if 100 <= status && status <= 199 {
		return styles.YellowText(s)
	}
	if 200 <= status && status <= 299 {
		return styles.GreenText(s)
	}
	if 300 <= status && status <= 399 {
		return styles.YellowText(s)
	}
	if 400 <= status && status <= 499 {
		return styles.RedText(s)
	}
	if 500 <= status && status <= 599 {
		return styles.RedText(s)
	}

	return styles.RedText(s)
}

func methodColor(method string) lipgloss.Style {
	s := fmt.Sprintf(" " + method + " ")
	switch method {
	case "GET":
		return styles.BlueLabel(s)
	case "POST":
		return styles.GreenLabel(s)
	case "PUT":
		return styles.PurpleLabel(s)
	case "DELETE":
		return styles.RedLabel(s)
	default:
		return styles.RedText(s)
	}
}

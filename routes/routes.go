package routes

import (
	"fmt"
	"main/handlers"
	"main/middleware"
	"net/http"

	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

func SetupRoutes() {

	mux := http.NewServeMux()
	mux.HandleFunc("/register", middleware.Logger(handlers.Register))

	mux.HandleFunc("/login", middleware.Logger((middleware.BasicAuth(handlers.Login))))
	mux.HandleFunc("/update", middleware.Logger((middleware.BasicAuth(handlers.Update))))
	mux.HandleFunc("/delete", middleware.Logger((middleware.BasicAuth(handlers.Delete))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// IdleTimeout:  time.Minute,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 30 * time.Second,
	}

	fmt.Print("\n(∩｀-´)⊃━ ☆ﾟ . * ･ ｡ﾟ => http server started on ")
	fmt.Println(styles.GreenText(server.Addr))
	fmt.Println()

	server.ListenAndServe()
}

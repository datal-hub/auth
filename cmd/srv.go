package cmd

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/urfave/cli"

	"github.com/datal-hub/auth/handlers/actions"
	"github.com/datal-hub/auth/handlers/middleware"
	log "github.com/datal-hub/auth/pkg/logger"
	"github.com/datal-hub/auth/pkg/settings"
)

//description run the application from the command line in server mode
var Srv = cli.Command{
	Name:        "srv",
	Usage:       "Start auth",
	Description: `Auth service serves register and login requests.`,
	Action:      runSrv,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "listen, l", Value: "127.0.0.1:8181",
			Usage: "Start server at specified address, port",
		},
	},
}

func updateSettings(c *cli.Context) {
	if c.IsSet("listen") {
		settings.ListenAddr = c.String("listen")
	}
}

func robots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("User-agent: *\nDisallow: /\n"))
}

func runSrv(c *cli.Context) error {
	updateSettings(c)

	r := mux.NewRouter()

	logger := alice.New(middleware.LoggingHandler)
	r.Methods("GET").Path("/robots.txt").Handler(
		logger.Then(http.HandlerFunc(robots)))

	limiter := tollbooth.NewLimiter(1, nil)
	limiter.SetMessage(`{"message": "You have reached maximum request limit."}`)
	limiter.SetMessageContentType("application/json; charset=utf-8")

	chain := alice.New(middleware.LoggingHandler, middleware.DatabaseHandler)

	r.Methods("POST").Path("/register").Handler(tollbooth.LimitHandler(limiter,
		chain.Then(http.HandlerFunc(actions.Register))))

	logDetails := log.Fields{
		"listenAddr": settings.ListenAddr,
	}
	log.InfoF("runSrv: Starting auth service.", logDetails)
	srv := &http.Server{
		Handler:      r,
		Addr:         settings.ListenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		logDetails = log.Fields{"message": err.Error()}
		log.ErrorF("runSrv: ListenAndServe error.", logDetails)
	}

	return nil
}

package http

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/earlofurl/pxthc/token"
	_ "github.com/lib/pq"
	"github.com/rogpeppe/go-internal/modfile"
	"math"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"

	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/config"
	"github.com/earlofurl/pxthc/jsonlog"
	"github.com/earlofurl/pxthc/postgres"
	"github.com/earlofurl/pxthc/sqlc"
)

const (
	swaggerDocsAssetPath = "./docs/"
)

// Server represents an HTTP server. It is meant to wrap all HTTP functionality
// used by the application so that dependent packages (such as cmd/pxthc) do not
// need to reference the "net/http" package at all.
type Server struct {
	Version string
	cfg     *config.Config

	logger *jsonlog.Logger
	store  sqlc.Store
	wg     sync.WaitGroup

	ln         net.Listener
	httpServer *http.Server
	router     *chi.Mux
	tokenMaker token.Maker
	//sc     *securecookie.SecureCookie

	// Bind address & domain for the server's listener.
	// If domain is specified, server is run on TLS using acme/autocert.
	Addr   string
	Domain string

	// Keys used for secure cookie encryption.
	HashKey  string
	BlockKey string

	// Servics used by the various HTTP routes.
	//AuthService           wtf.AuthService
	//EventService          wtf.EventService
	userService   pxthc.UserService
	uomService    pxthc.UomService
	strainService pxthc.StrainService
}

type Options func(opts *Server) error

// NewServer creates a new HTTP server with db store connection.
func NewServer(opts ...Options) *Server {
	s := defaultServer()

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to apply server option at startup")
		}
	}

	return s
}

func defaultServer() *Server {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Init(version string) {
	s.Version = version
	//s.runDBMigration(s.cfg.MigrationURL, s.cfg.DBSource)
	s.newStore()
	s.newTokenMaker()
	//s.newValidator()
	s.newRouter()
	s.newServices()
	s.setGlobalMiddleware()
	s.InitRoutes()
}

func (s *Server) newStore() {
	conn, err := openDbConn(s.cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}
	s.store = sqlc.NewStore(conn)
}

func (s *Server) newTokenMaker() {
	tokenMaker, err := token.NewPasetoMaker(s.cfg.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create token maker")
	}
	s.tokenMaker = tokenMaker
}

func openDbConn(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DBMaxOpenConnections)
	db.SetMaxIdleConns(cfg.DBMaxIdleConnections)

	duration, err := time.ParseDuration(cfg.DBMaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Info().Msg("database connection pool and sqlc.Store established")

	return db, nil
}

func (s *Server) newRouter() {
	s.router = chi.NewRouter()
}

func (s *Server) newServices() {
	s.userService = postgres.NewUserService(&s.store)
	s.uomService = postgres.NewUomService(&s.store)
	s.strainService = postgres.NewStrainService(&s.store)
}

// setGlobalMiddleware sets the global middleware for the chi router to apply to all routes.
func (s *Server) setGlobalMiddleware() {
	// Global NotFound response.
	s.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "endpoint not found"}`))
	})
	//s.router.Use(chiMiddleware.CORS)
	if s.cfg.RequestLog {
		s.router.Use(chiMiddleware.Logger)
	}
	s.router.Use(chiMiddleware.Recoverer)
}

func (s *Server) runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func (s *Server) Run() {
	s.httpServer = &http.Server{
		Addr:              s.cfg.HTTPServerAddress + ":" + s.cfg.HTTPServerPort,
		Handler:           s.router,
		ReadHeaderTimeout: s.cfg.ReadHeaderTimeout,
	}

	fmt.Println(` 
▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄       ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄                 ▄▄▄▄▄▄▄▄▄▄▄  ▄         ▄  ▄▄▄▄▄▄▄▄▄▄▄ 
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌     ▐░▌▐░░░░░░░░░░░▌▐░▌               ▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌
▐░█▀▀▀▀▀▀▀█░▌ ▀▀▀▀█░█▀▀▀▀  ▐░▌   ▐░▌ ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌                ▀▀▀▀█░█▀▀▀▀ ▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀▀▀ 
▐░▌       ▐░▌     ▐░▌       ▐░▌ ▐░▌  ▐░▌          ▐░▌                    ▐░▌     ▐░▌       ▐░▌▐░▌          
▐░█▄▄▄▄▄▄▄█░▌     ▐░▌        ▐░▐░▌   ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌                    ▐░▌     ▐░█▄▄▄▄▄▄▄█░▌▐░▌          
▐░░░░░░░░░░░▌     ▐░▌         ▐░▌    ▐░░░░░░░░░░░▌▐░▌                    ▐░▌     ▐░░░░░░░░░░░▌▐░▌          
▐░█▀▀▀▀▀▀▀▀▀      ▐░▌        ▐░▌░▌   ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌                    ▐░▌     ▐░█▀▀▀▀▀▀▀█░▌▐░▌          
▐░▌               ▐░▌       ▐░▌ ▐░▌  ▐░▌          ▐░▌                    ▐░▌     ▐░▌       ▐░▌▐░▌          
▐░▌           ▄▄▄▄█░█▄▄▄▄  ▐░▌   ▐░▌ ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄           ▐░▌     ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄▄▄ 
▐░▌          ▐░░░░░░░░░░░▌▐░▌     ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌          ▐░▌     ▐░▌       ▐░▌▐░░░░░░░░░░░▌
 ▀            ▀▀▀▀▀▀▀▀▀▀▀  ▀       ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀            ▀       ▀         ▀  ▀▀▀▀▀▀▀▀▀▀▀ `)

	go func() {
		start(s)
	}()

	_ = gracefulShutdown(context.Background(), s)
}

func (s *Server) Config() *config.Config {
	return s.cfg
}

func start(s *Server) {
	log.Printf("Serving at %s:%s\n", s.cfg.HTTPServerAddress, s.cfg.HTTPServerPort)
	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}

func gracefulShutdown(ctx context.Context, s *Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-quit

	log.Info().Msg("Shutting down...")

	ctx, shutdown := context.WithTimeout(ctx, s.Config().GracefulShutdownTimeout*time.Second)
	defer shutdown()

	// Close all opened resources
	//_ = s.store.Close() TODO: implement this as a method on the store interface

	return s.httpServer.Shutdown(ctx)
}

// PrintAllRegisteredRoutes prints all registered routes from Chi router.
// definitely can be an extension to the router instead.
func (s *Server) PrintAllRegisteredRoutes(exceptions ...string) {
	exceptions = append(exceptions, "/swagger")

	walkFunc := func(method string, path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {

		for _, val := range exceptions {
			if strings.HasPrefix(path, val) {
				return nil
			}
		}

		switch method {
		case "GET":
			fmt.Printf("%s", gchalk.Green(fmt.Sprintf("%-8s", method)))
		case "POST", "PUT", "PATCH":
			fmt.Printf("%s", gchalk.Yellow(fmt.Sprintf("%-8s", method)))
		case "DELETE":
			fmt.Printf("%s", gchalk.Red(fmt.Sprintf("%-8s", method)))
		default:
			fmt.Printf("%s", gchalk.White(fmt.Sprintf("%-8s", method)))
		}

		//fmt.Printf("%-25s %60s\n", path, getHandler(getModName(), handler))
		fmt.Printf("%s", strPad(path, 25, "-", "RIGHT"))
		fmt.Printf("%s\n", strPad(getHandler(getModName(), handler), 60, "-", "LEFT"))

		return nil
	}
	if err := chi.Walk(s.router, walkFunc); err != nil {
		fmt.Print(err)
	}

	if s.cfg.RunSwagger {
		fmt.Printf("%s", gchalk.Green(fmt.Sprintf("%-8s", "GET")))
		fmt.Printf("/swagger\n")
	}
}

// StrPad returns the input string padded on the left, right or both sides using padType to the specified padding length padLength.
//
// Example:
// input := "Codes";
// StrPad(input, 10, " ", "RIGHT")        // produces "Codes     "
// StrPad(input, 10, "-=", "LEFT")        // produces "=-=-=Codes"
// StrPad(input, 10, "_", "BOTH")         // produces "__Codes___"
// StrPad(input, 6, "___", "RIGHT")       // produces "Codes_"
// StrPad(input, 3, "*", "RIGHT")         // produces "Codes"
// taken from // https://gist.github.com/asessa/3aaec43d93044fc42b7c6d5f728cb039
func strPad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

func getHandler(projectName string, handler http.Handler) (funcName string) {
	// https://github.com/go-chi/chi/issues/424
	funcName = runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
	base := filepath.Base(funcName)

	nameSplit := strings.Split(funcName, "")
	names := nameSplit[len(projectName):]
	path := strings.Join(names, "")

	pathSplit := strings.Split(path, "/")
	path = strings.Join(pathSplit[:len(pathSplit)-1], "/")

	sFull := strings.Split(base, ".")
	s := sFull[len(sFull)-1:]

	s = strings.Split(s[0], "")
	if len(s) <= 4 && len(sFull) >= 3 {
		s = sFull[len(sFull)-3 : len(sFull)-2]
		return "@" + gchalk.Blue(strings.Join(s, ""))
	}
	s = s[:len(s)-3]
	funcName = strings.Join(s, "")

	return path + "@" + gchalk.Blue(funcName)
}

// adapted from https://stackoverflow.com/a/63393712/1033134
func getModName() string {
	goModBytes, err := os.ReadFile("go.mod")
	if err != nil {
		os.Exit(0)
	}
	return modfile.ModulePath(goModBytes)
}

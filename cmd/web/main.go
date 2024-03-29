package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ueetim/snippetbox/internal/models"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

// for application-wide dependencies
type application struct {
	errorLog 		*log.Logger
	infoLog 		*log.Logger
	snippets		*models.SnippetModel
	templateCache	map[string]*template.Template
	formDecoder 	*form.Decoder
	sessionManager	*scs.SessionManager
}

func main() {

	// flag for server start
	addr := flag.String("addr", ":4000", "HTTP network address")

	// db conn
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	// for logging info msgs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ltime)

	// for logging error msgs
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// init a new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// for getting form data
	formDecoder := form.NewDecoder()

	// init a new session manager, config to use mysql db, and set lifetime of 12 hrs for sessions
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	// init a new instance of application struct
	app := &application{
		errorLog: 		errorLog,
		infoLog: 		infoLog,
		snippets:		&models.SnippetModel{DB: db},
		templateCache: 	templateCache,
		formDecoder:	formDecoder,
		sessionManager: sessionManager,
	}

	srv := &http.Server {
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
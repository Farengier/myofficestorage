package www

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	log     *log.Logger
	listen  string
	storage storage
}
type storage interface {
	Store(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close()
}

func Server(logger *log.Logger, listen string, storage storage) *server {
	if logger == nil {
		logger = log.New(io.Discard, "", 0)
	}
	return &server{
		log:     logger,
		listen:  listen,
		storage: storage,
	}
}

func (s *server) Start() error {
	s.log.Println("starting server at", s.listen)

	r := mux.NewRouter()
	r.HandleFunc("/", s.homePageHandler)
	r.HandleFunc("/{key}", s.storeHandler)

	srv := http.Server{Addr: s.listen, Handler: r}
	return srv.ListenAndServe()
}

func (s *server) homePageHandler(w http.ResponseWriter, r *http.Request) {
	response := `
Simple key-value storage for MyOffice company
* to store value use POST on url /key with value in body JSON-encoded string
* to get value use GET on url /key, response empty if no such value stored
* to delete value use POST on url /key with empty body
`
	s.writeToResponse(w, response)
}

func (s *server) storeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.writeToResponse(w, "Failed parsing url")
		return
	}

	vars := mux.Vars(r)
	key, ok := vars["key"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		s.writeToResponse(w, "No key present in url")
		return
	}

	if r.Method == "GET" {
		s.read(key, w)
		return
	}
	if r.Method == "POST" {
		s.write(key, w, r)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	s.writeToResponse(w, "Unsupported request type")
}

func (s *server) read(key string, w http.ResponseWriter) {
	s.log.Println("Reading from", key)
	value, err := s.storage.Get(key)
	s.log.Println("Readed", value, err)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write(value)
	if err != nil {
		s.log.Println("Error writing response", err)
	}
}

func (s *server) write(key string, w http.ResponseWriter, r *http.Request) {
	s.log.Println("Writing to", key)
	defer r.Body.Close()
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error reading body content")
		return
	}

	if len(value) == 0 {
		err = s.storage.Delete(key)
	} else {
		err = s.storage.Store(key, value)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error processing request", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *server) writeToResponse(w http.ResponseWriter, msg ...any) {
	_, err := fmt.Fprint(w, msg...)
	if err != nil {
		log.Println("Error writing response", err)
	}
}

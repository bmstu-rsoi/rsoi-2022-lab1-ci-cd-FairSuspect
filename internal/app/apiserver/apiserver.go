package apiserver

import (
	"encoding/json"
	"fmt"
	"http-rest-api/internal/app/apiserver/model"
	"http-rest-api/store"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}

}

// Start ..
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.router.Use(commonMiddleware)
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil

}

func (s *APIServer) configureRouter() {

	s.router.HandleFunc("/error", s.handleForbidden())
	s.router.HandleFunc("/persons", s.handlePersons())
	s.router.HandleFunc("/persons/{id:[0-9]+}", s.handlePersonsId())

}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err

	}
	s.store = st

	return nil
}
func (s *APIServer) handlePersons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			persons, err := s.store.Person().GetAll()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			j, err := json.MarshalIndent(persons, "", "\t")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			io.WriteString(w, string(j)+"\n")
			break
		case "POST":
			var p *model.Person

			// Try to decode the request body into the struct. If there is an error,
			// respond to the client with the error message and a 400 status code.
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				io.WriteString(w, "Failed to parse model: "+err.Error())
				return
			}
			person, err := s.store.Person().Create(p)
			if err != nil {
				io.WriteString(w, "Failed to create model \n"+err.Error())
				return
			}
			w.WriteHeader(http.StatusCreated)
			parsed, err := json.Marshal(person)
			if err != nil {
				io.WriteString(w, "Failed to evaulate model \n"+err.Error())
			}
			io.WriteString(w, "Created \n"+string(parsed))
			break
		default:
			refuseMethod(w)

		}

	}
}
func (s *APIServer) handlePersonsId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "GET":

			person, err := s.store.Person().GetById(id)
			if err != nil {
				http.Error(w, err.Error(), 404)
				return
			}
			j, err := json.Marshal(person)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, string(j)+"\n")
			break
		case "DELETE":

			_, err := s.store.Person().DeleteById(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)

			break

		case http.MethodPatch:

			var p *model.Person

			// Try to decode the request body into the struct. If there is an error,
			// respond to the client with the error message and a 400 status code.
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			person, err := s.store.Person().Patch(p)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			j, err := json.MarshalIndent(person, "", "\t")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, string(j))
			break

		}

	}
}

func (s *APIServer) handleForbidden() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}
		apiParam := r.URL.Query().Get("api")
		if len(apiParam) == 0 {
			http.Error(w, "api is required", 400)
			return
		}
		io.WriteString(w, fmt.Sprintf("Your api is %s", apiParam))
	}
}

func refuseMethod(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", 405)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

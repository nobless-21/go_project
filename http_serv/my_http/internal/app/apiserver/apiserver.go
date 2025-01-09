package apiserver

import (
	"io"
	"my_http/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct{
	config *Config
	loglevel *logrus.Logger
	router *mux.Router
	store *store.Store
}

func New(config *Config) *APIserver{
	return &APIserver{
		config: config,
		loglevel: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error{
	if err := s.configureLogger(); err != nil{
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil{
		return err
	}

	s.loglevel.Info("starting api server")

	return http.ListenAndServe(s.config.Port, s.router)
}

func (s *APIserver) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil{
		return err
	}

	s.loglevel.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter(){
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *APIserver) configureStore() error{
	st := store.NewStore(s.config.Store)

	if err := st.Open(); err != nil{
		return err
	}

	s.store = st

	return nil
}

func (s *APIserver) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Hello World")
	}
}
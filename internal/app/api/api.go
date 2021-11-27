package api

import (
	"forum/pkg/logger"
	"forum/storage"
	"net/http"
)

var (
	prefix = "/api"
)

type API struct {
	config  *Config
	router  *http.ServeMux
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		router: http.NewServeMux(),
	}
}

// Start http server/configure loggers, router, database connection and etc
func (api *API) Start() error {

	logger.InfoLogger.Println("Starting the application at port:", api.config.Port)

	api.configureRouter()
	if err := api.configureStore(); err != nil {
		return err
	}
	return http.ListenAndServe(":"+api.config.Port, api.router)
}

func (api *API) configureRouter() {
	api.router.HandleFunc(prefix+"/user/register", api.PostUserRegister)
	api.router.HandleFunc(prefix+"/users", api.GetAllUsers)
}

//configureStore method
func (api *API) configureStore() error {
	st := storage.New(api.config.Storage)
	if err := st.Open(); err != nil {
		return err
	}
	api.storage = st
	st.AddTables()
	return nil
}

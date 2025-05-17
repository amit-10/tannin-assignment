package ports

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tanninio/home-assignment/internal/app"
	"github.com/tanninio/home-assignment/internal/metrics"
)

func HttpCreateServiceHandler(svc app.PetService, path string, configureRouters func(root, svcrouter *mux.Router)) http.Handler {
	root := mux.NewRouter()
	AddBaseRoutesHandlers(root)
	svcrouter := root.PathPrefix(path).Subrouter()
	configureRouters(root, svcrouter)
	svcsrv := NewStrictHandlerWithOptions(NewHttpServer(svc), nil, StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  HttpRespondWithHttpError,
		ResponseErrorHandlerFunc: HttpRespondWithHttpError,
	})
	svcrouter.Handle("/", HandlerFromMux(svcsrv, svcrouter))
	return root
}

func HttpServeHandler(addr string, handler http.Handler) {
	logrus.Info("Starting HTTP server")
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func AddBaseRoutesHandlers(root *mux.Router) {
	root.HandleFunc("/metrics", metrics.MetricsHandler).Methods("GET")
}

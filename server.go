package delivery

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pavank830/delivery/bestroute"
)

func init() {
	tr := &http.Transport{
		DisableCompression: true,
		Proxy:              http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost:   3,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       configTLS(""),
	}

	httpClient = &http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}
}

//Start -- func that starts http server
func Start(port string) error {
	return startHTTPSvr(port)
}

func configTLS(serverName string) *tls.Config {
	TLSConfig := &tls.Config{
		ServerName:             serverName,
		Certificates:           []tls.Certificate{{}},
		Rand:                   rand.Reader,
		SessionTicketsDisabled: false,
		MinVersion:             tls.VersionTLS12,
		InsecureSkipVerify:     true,
	}
	return TLSConfig
}

func getRouter() *mux.Router {
	// gorilla/mux for routing
	router := mux.NewRouter()
	router.Use(starterMiddleware)
	router.HandleFunc(bestroute.BestRoutePath, bestroute.GetBestRoute).Methods(http.MethodPost)
	router.NotFoundHandler = http.HandlerFunc(defaultHandler)
	return router
}

func starterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Credentials", "true")
		if req.Method == "OPTIONS" {
			resp.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(resp, req)
		log.Printf("Duration: %s %s %s\n", req.Method, req.RequestURI, time.Since(start).String())
	})
}

func startHTTPSvr(port string) error {
	server := &http.Server{
		Addr:         listenAddr + ":" + port,
		Handler:      getRouter(),
		ReadTimeout:  time.Duration(HTTPReadTimeout) * time.Second,
		WriteTimeout: time.Duration(HTTPWriteTimeout) * time.Second,
	}
	log.Printf("## HTTP Server listening on %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Failed to start http server %+v\n", err)
		return err
	}
	return nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("defaultHandler Duration: %s %s ", r.Method, r.RequestURI)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not found"))
}

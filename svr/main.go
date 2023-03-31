package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"openai-svr/handler"
	openaiwrapper "openai-svr/openai_wrapper"
	"openai-svr/utils"
	"os"
	"time"

	"flag"

	"github.com/gorilla/mux"
)

var configuration = utils.Configuration{}

func parseCmd() error {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "server config file")
	flag.Parse()

	configuration.ReadConfig(configFile)
	fmt.Printf("%v\n", configuration)

	openaiwrapper.SetupToken(configuration.OpenaiApiToken)
	return nil
}

func main() {
	if err := parseCmd(); err != nil {
		fmt.Println("parseCmd failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("main start")
	r := mux.NewRouter()
	r.HandleFunc("/chat/text-completion", handler.HandleSingleCompletion).Methods("POST")
	r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })

	caClient, _ := ioutil.ReadFile(configuration.CertFile)
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caClient)
	svr := &http.Server{
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Minute * 3,
		Handler:      r,
		TLSConfig: &tls.Config{
			ClientCAs:  caPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	err := svr.ListenAndServeTLS(configuration.CertFile, configuration.KeyFile)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
		os.Exit(1)
	}
}

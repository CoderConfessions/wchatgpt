package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"openai-svr/handler"
	mysqlwrapper "openai-svr/mysql_wrapper"
	openaiwrapper "openai-svr/openai_wrapper"
	"openai-svr/utils"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var configuration = utils.Configuration{}

func main() {
	if err := utils.ParseCmd(&configuration); err != nil {
		fmt.Println("parseCmd failed:", err.Error())
		os.Exit(1)
	}
	openaiwrapper.SetupOpenAIClientConfig(configuration.OpenaiApiToken, "")

	if err := mysqlwrapper.InitPool(); err != nil {
		fmt.Println("init mysql connetion poll failed:", err.Error())
		os.Exit(1)
	}
	defer mysqlwrapper.ReleasePool()

	fmt.Println("main start")
	r := mux.NewRouter()
	handler.Register(r)

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

	// err := svr.ListenAndServeTLS(configuration.CertFile, configuration.KeyFile)
	if err := svr.ListenAndServe(); err != nil {
		fmt.Println("ListenAndServe error: ", err.Error())
		os.Exit(1)
	}
}

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

	"k8s.io/klog/v2"

	"github.com/gorilla/mux"
)

var configuration = utils.NewConfiguration()

func main() {
	if err := utils.ParseCmd(&configuration); err != nil {
		klog.Fatalf("ParseCmd failed: %s", err.Error())
		os.Exit(1)
	}
	utils.InitLocks()
	openaiwrapper.SetupOpenAIClientConfig(configuration.OpenaiApiToken, "")

	if err := mysqlwrapper.InitPool(configuration.DB); err != nil {
		klog.Fatalf("init mysql connection pool failed: %s", err.Error())
	}
	defer mysqlwrapper.ReleasePool()

	r := mux.NewRouter()
	handler.Register(r)

	caClient, _ := ioutil.ReadFile(configuration.CertFile)
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caClient)
	svr := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", configuration.IP, configuration.Port),
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

package hlogger

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	lg "github.com/xinnige/gologger/logger"
)

func TestLogRequest(t *testing.T) {
	logger := lg.NewLogger(false, os.Stdout)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)
	}))
	defer ts.Close()

	const body = "Payload..."
	r, err := http.NewRequest("POST", ts.URL, strings.NewReader(body))
	assert.Nil(t, err)
	logRequest(logger, r)
}

func TestLogRequestFile(t *testing.T) {
	logfile, err := os.OpenFile("request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer func() {
		cerr := logfile.Close()
		if cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	logger := lg.NewLogger(true, logfile)

	const body = "Payload..."
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%q", dump)
	}))
	defer ts.Close()
	r, err := http.NewRequest("POST", ts.URL, strings.NewReader(body))
	r.Header.Add("Authorization", "some-token-here")
	assert.Nil(t, err)
	logRequest(logger, r)
}

func TestLogRespFile(t *testing.T) {
	logfile, err := os.OpenFile("resp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer func() {
		cerr := logfile.Close()
		if cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	logger := lg.NewLogger(true, logfile)

	const body = "Payload..."
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Date", "Wed, 19 Jul 1972 19:00:00 GMT")
		fmt.Fprintln(w, body)
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	logResponse(logger, resp)
}

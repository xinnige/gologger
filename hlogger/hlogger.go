package hlogger

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"

	lg "github.com/xinnige/gologger/logger"
)

func warning(logger *lg.Logger, obj interface{}, format string, v ...interface{}) {
	logger.Warning(fmt.Sprintf("%T", obj), format, v...)
}
func fatal(logger *lg.Logger, obj interface{}, format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf("%T", obj), format, v...)
}
func info(logger *lg.Logger, obj interface{}, format string, v ...interface{}) {
	logger.Info(fmt.Sprintf("%T", obj), format, v...)
}
func debug(logger *lg.Logger, obj interface{}, format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf("%T", obj), format, v...)
}

func valueOrDefault(value, def string) string {
	if value != "" {
		return value
	}
	return def
}

var reqWriteExcludeHeaderDump = map[string]bool{
	"Host":              true, // not in Header map anyway
	"Transfer-Encoding": true,
	"Trailer":           true,
}

func logRequest(logger *lg.Logger, r *http.Request) {
	if logger.IsDebug {
		text, err := httputil.DumpRequest(r, true)
		if err != nil {
			warning(logger, *logger, "Failed to log request, err: %v", err)
			return
		}
		debug(logger, *r, "Request %s", text)
		return
	}

	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %s HTTP/%d.%d\n", valueOrDefault(r.Method, "GET"), r.URL.String(), r.ProtoMajor, r.ProtoMinor)
	if err := r.Header.WriteSubset(&b, reqWriteExcludeHeaderDump); err != nil {
		fatal(logger, *logger, "Failed to log request, err: %v", err)
		return
	}
	info(logger, *r, "Request %s", b.Bytes())
}

func logResponse(logger *lg.Logger, resp *http.Response) {
	text, err := httputil.DumpResponse(resp, logger.IsDebug)
	if err != nil {
		warning(logger, *logger, "Failed to log response, err: %v", err)
		return
	}
	info(logger, *resp, "Response %s", text)
}

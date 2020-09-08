# Golang logger example

A simple example of golang logger

## Logger

A common logger

### Log Level

* INFO
* DEBUG
* WARNING
* FATAL
* ERROR

### Modes & Formats

2 types of predefined formats are available for now:

* standard simple mode
* datadog compatible mode

### Standard Simple Mode

#### Usage

##### 1. An example of logging to OS Stdout/Stderr

```
debugMode := true
logger := NewLogger(debugMode, os.Stdout)
logger.Debug("Test", "test write to file")
```

##### 2. An example of logging to file

```
logfile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
    log.Printf("error opening file: %v", err)
}
defer func() {
    cerr := logfile.Close()
    if cerr != nil {
        log.Printf("error closing file: %v", cerr)
    }
}()

debugMode := false
logger := NewLogger(debugMode, logfile)
logger.Debug("Test", "test write to file")
```

##### 3. An example of logging struct info

```
debugMode := true
logger := NewLogger(debugMode, os.Stdout)

func info(logger *Logger, obj interface{}, format string, v ...interface{}) {
	logger.Warning(fmt.Sprintf("%T", obj), format, v...)
}

func debug(logger *Logger, obj interface{}, format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf("%T", obj), format, v...)
}

func error(logger *Logger, obj interface{}, format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf("%T", obj), format, v...)
}

mystruct := &MyStruct{}
info(logger, mystruct, "Process w/ %v", mystruct)

err := mystruct.Do()
if err != nil {
    error(logger, *mystruct, "Failed to process, err: %v", err)
}

```

#### Log Output Example

```
2020/03/25 12:47:54 [DEBUG][Test] test write to file
2020/03/25 12:47:54 [INFO][Test] test write to file
2020/03/25 12:47:54 [FATAL][Test] test write to file
2020/03/25 12:47:54 [ERROR][Test] test write to file
2020/03/25 12:47:54 [WARN][Test] test write to file
```

### Datadog Log Explorer Compatible Mode

#### Usage

##### An example of logging to OS Stdout/Stderr

```
debugMode := true
logger := NewDdgLogger(debugMode, os.Stdout)
logger.Debug("Test", "test write to file")
```

#### Log Output Example

```
[INFO]	2020-03-25T07:52:16.788Z	[Test] test write to file
[DEBUG]	2020-03-25T07:52:16.788Z	[Test] test write to file
[FATAL]	2020-03-25T07:52:16.788Z	[Test] test write to file
[ERROR]	2020-03-25T07:52:16.788Z	[Test] test write to file
[WARN]	2020-03-25T07:52:16.788Z	[Test] test write to file
```

### Options

* SetFlags  
  set flags of logger, e.g. `log.LstdFlags | log.Lshortfile`
* SetLabels  
  set label of log level, defaults are `DEBUG`, `INFO`, `WARN`, `FATAL`, `ERROR`
* SetDdgMode  
  switch to logging in datadog-compatible format
* SetStdMode  
  switch to logging in a simple format

## HTTP Request/Response Logger

### Log Level

* INFO
* DEBUG

### Request logs

#### Usage

```

const body = "Payload..."
r, err := http.NewRequest("POST", "http://127.0.0.1:50132", strings.NewReader(body))

logger := NewLogger(false, os.Stdout)
logRequest(logger, r)
```

#### Log Output Example

* DEBUG

```
2020/03/25 12:11:49 [DEBUG][http.Request] Request POST / HTTP/1.1
Host: 127.0.0.1:50132
Authorization: some-token-here

Payload...
```

* INFO

```
2020/03/25 12:11:49 [INFO][http.Request] Request POST http://127.0.0.1:50132 HTTP/1.1
Authorization: some-token-here
```

### Response logs

#### Usage

```
resp, _ := http.Get("http://127.0.0.1:8080")

logger := NewLogger(false, os.Stdout)
logResponse(logger, resp)
```

#### Log Output Example

* DEBUG

```
2020/03/25 17:08:56 [INFO][http.Response] Response HTTP/1.1 200 OK
Content-Length: 11
Content-Type: text/plain; charset=utf-8
Date: Wed, 19 Jul 1972 19:00:00 GMT

Payload...
```

* INFO

```
2020/03/25 17:08:56 [INFO][http.Response] Response HTTP/1.1 200 OK
Content-Length: 11
Content-Type: text/plain; charset=utf-8
Date: Wed, 19 Jul 1972 19:00:00 GMT
```

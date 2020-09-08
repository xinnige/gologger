package logger

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	assert.NotNil(t, logger)
}

func TestNewFileLogger(t *testing.T) {
	logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer func() {
		cerr := logfile.Close()
		if cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	logger := NewLogger(true, logfile)
	logger.Debug("Test", "test write to file")
	logger.Info("Test", "test write to file")
	logger.Fatal("Test", "test write to file")
	logger.Error("Test", "test write to file")
	logger.Warning("Test", "test write to file")
	assert.NotNil(t, logger)

	assert.Equal(t, StdMode, logger.GetMode())
}

func TestNewDdgLogger(t *testing.T) {
	logfile, err := os.OpenFile("ddg.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer func() {
		cerr := logfile.Close()
		if cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	logger := NewDdgLogger(true, logfile)
	logger.Debug("Test", "test write to file")
	logger.Info("Test", "test write to file")
	logger.Fatal("Test", "test write to file")
	logger.Error("Test", "test write to file")
	logger.Warning("Test", "test write to file")
	assert.NotNil(t, logger)

	assert.Equal(t, DdgMode, logger.GetMode())
}
func TestDebug(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	logger.Debug("Test", "test")

	logger = NewLogger(false, os.Stdout)
	logger.Debug("Test", "test")
}

func TestSwitchMode(t *testing.T) {
	logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer func() {
		cerr := logfile.Close()
		if cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	logger := NewDdgLogger(true, logfile)
	logger.Debug("Test", "test write to file")
	assert.Equal(t, DdgMode, logger.GetMode())

	logger.SetStdMode()
	logger.Debug("Test", "test write to file")
	assert.Equal(t, StdMode, logger.GetMode())

	logger.SetDdgMode()
	logger.Debug("Test", "test write to file")
	assert.Equal(t, DdgMode, logger.GetMode())
}

func TestSetLabels(t *testing.T) {
	logfile, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer func() {
		cerr := logfile.Close()
		if cerr != nil {
			log.Printf("error closing file: %v", cerr)
		}
	}()

	logger := NewLogger(true, logfile)
	logger.SetLabels("info", "debug", "warn", "fatal", "error")
	logger.Debug("Test", "test write to file")
	logger.Info("Test", "test write to file")
	logger.Fatal("Test", "test write to file")
	logger.Error("Test", "test write to file")
	logger.Warning("Test", "test write to file")

}

func TestSetOutput(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	logger.SetOutput(os.Stderr)
}

func TestSetFlags(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	logger.SetFlags(log.Lshortfile)
}
func TestLogging(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	assert.NotNil(t, logger)
	logger.Debug("Test", "test write to file")
	logger.Info("Test", "test write to file")
	logger.Fatal("Test", "test write to file")
	logger.Error("Test", "test write to file")
	logger.Warning("Test", "test write to file")
}

func TestSetLogger(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	logger.SetLogger(log.New(os.Stdout, "", log.Lshortfile))
}

func TestPrint(t *testing.T) {
	logger := NewLogger(true, os.Stdout)
	logger.Print("ok")

	logger.Printf("%d", 1)

	logger.Println("ok")

	assert.Nil(t, logger.Output(1, "ok"))
}

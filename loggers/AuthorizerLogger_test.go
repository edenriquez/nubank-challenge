package loggers

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	nubankModels "github.com/edenriquez/nubank-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestLogError(t *testing.T) {
	errorList := []error{}

	out := make(chan string)
	r, w, old := listenPipe()

	LogError(errorList...)

	go func() { out <- bufferReader(r) }()
	closeBuffer(w, old)
	assert.Equal(t, <-out, "")
}
func TestPrintErrorList(t *testing.T) {
	errorList := []error{
		errors.New("first"),
		errors.New("second"),
	}

	out := make(chan string)
	r, w, old := listenPipe()

	printErrorList(errorList...)

	go func() { out <- bufferReader(r) }()
	closeBuffer(w, old)
	output := <-out
	assert.Contains(t, output, errorList[0].Error())
	assert.Contains(t, output, errorList[1].Error())
}

func TestLogAction(t *testing.T) {
	account := nubankModels.Account{}
	account.Mock()

	out := make(chan string)
	r, w, old := listenPipe()

	LogAction(account, "creation")

	go func() { out <- bufferReader(r) }()
	closeBuffer(w, old)

	assert.Contains(t, <-out, "creation	-	{\"account\":{\"active-card\":true,\"available-limit\":100}}")
}

// helpers to recover stdout from logger
func listenPipe() (*os.File, *os.File, *os.File) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Stderr = w
	log.SetOutput(os.Stderr)
	return r, w, old
}

func bufferReader(r *os.File) string {
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func closeBuffer(w *os.File, old *os.File) {
	w.Close()
	os.Stdout = old
}

package executor

import (
	"bytes"
	"context"
	"testing"
	"time"
)

func TestShortCommand(t *testing.T) {
	l := LinuxAdapter{}

	outputChan := make(chan []byte, 1024)
	err := l.Run(context.Background(), "echo 123", outputChan)
	assertNil(t, err)

	expected := "123\n"
	actual := readFromChan(outputChan)

	assert(t, expected, actual)
}

func TestLongCommandAndStop(t *testing.T) {
	l := LinuxAdapter{}

	outputChan := make(chan []byte, 1024)
	command := "counter=1\n" +
		"while true; do\n" +
		"echo $counter\n" +
		"counter=$((counter + 1))\n" +
		"sleep 2\n" +
		"done"
	runCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go l.Run(runCtx, command, outputChan)

	expected := "1\n2\n3\n4\n5\n"
	actual := readFromChan(outputChan)

	assert(t, expected, actual)
}

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Should not produce an error: %s", err.Error())
	}
}

func assert(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Result was incorrect, \nexpected: %s \nactual: %s", expected, actual)
	}
}

func readFromChan(outputChan <-chan []byte) string {
	var buffer bytes.Buffer
	for {
		bytes, ok := <-outputChan
		if !ok {
			break
		}
		buffer.Write(bytes)
	}
	return buffer.String()
}

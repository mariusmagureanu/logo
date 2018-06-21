package tests

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/mariusmagureanu/logo"
)

func TestBasic(t *testing.T) {
	buf := new(bytes.Buffer)

	logo.InitNewLogger(buf, 0)
	logo.InfoSync("Foo Bar")

	if buf.Len() == 0 {
		t.Errorf("The buffer was not supposed to be empty.")
	}
}

func TestInitStdout(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logo.InitNewLogger(os.Stdout, 0)
	logo.InfoSync("Foo Bar")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if !strings.Contains(string(out), "Foo Bar") {
		t.Errorf("Expected Foo Bar, got %s instead.", string(out))
	}
}

func TestInitTmp(t *testing.T) {
}

func TestInitToFile(t *testing.T) {
}

func TestLogError(t *testing.T) {
}

func TestLogErrorAsync(t *testing.T) {
}

func TestLogWarning(t *testing.T) {
}

func TestLogWarningAsync(t *testing.T) {

}

package herror

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

var outBuffer *bytes.Buffer

func TestMain(m *testing.M) {
	outBytes := make([]byte, 0, 1024)
	outBuffer = bytes.NewBuffer(outBytes)
	log.Logger = log.Output(outBuffer)
	os.Exit(m.Run())
}

var ErrTest = errors.New("one err")

func testResultErrFunc() error {
	return ErrTest
}

func TestResultErr(t *testing.T) {
	t.Run("add msg", func(t *testing.T) {
		he := HandleErr(testResultErrFunc()).Msg("test msg")
		assert.Equal(t, he.msg, "test msg")
	})
	t.Run("ignore err", func(t *testing.T) {
		defer outBuffer.Reset()
		HandleErr(
			testResultErrFunc(),
		).Msg("ignore err").Ignore()
		//t.Log(outBuffer.String())
		//{"level":"debug","error":"one err","Ignore":"HandleErr err","time":"2022-06-18T13:29:32+08:00","message":"ignore err"}
		json := gjson.Parse(outBuffer.String())
		assert.Equal(t, json.Get("level").String(), "debug")
		assert.Equal(t, json.Get("error").String(), ErrTest.Error())
		assert.Equal(t, json.Get("Ignore").String(), "HandleErr err")
		assert.Equal(t, json.Get("message").String(), "ignore err")
	})
}

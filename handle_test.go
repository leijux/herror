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

var (
	outBuffer *bytes.Buffer

	ErrTest = errors.New("test err")
)

func TestMain(m *testing.M) {
	outBytes := make([]byte, 0, 1024)
	outBuffer = bytes.NewBuffer(outBytes)
	zlog = log.Output(outBuffer)
	os.Exit(m.Run())
}

func testHandleErrFunc() error {
	return ErrTest
}

func TestHandleErr(t *testing.T) {
	t.Run("add msg", func(t *testing.T) {
		he := HandleErr(testHandleErrFunc()).Msg("test msg")
		assert.Equal(t, he.msg, "test msg")
	})
	t.Run("must err", func(t *testing.T) {
		defer outBuffer.Reset()
		defer func() {
			if err := recover(); err != nil {
				assert.Equal(t, err.(string), "must err")

				json := gjson.ParseBytes(outBuffer.Bytes())
				assert.Equal(t, json.Get("level").String(), "panic")
				assert.Equal(t, json.Get("error").String(), ErrTest.Error())
				assert.Equal(t, json.Get("Must").String(), "HandleErr err")
				assert.Equal(t, json.Get("message").String(), "must err")
			}
		}()
		HandleErr(testHandleErrFunc()).
			Msg("must err").
			Must()
	})
	t.Run("ignore err", func(t *testing.T) {
		defer outBuffer.Reset()
		HandleErr(
			testHandleErrFunc(),
		).Msg("ignore err").Ignore()
		//t.Log(outBuffer.String())
		//{"level":"debug","error":"one err","Ignore":"HandleErr err","time":"2022-06-18T13:29:32+08:00","message":"ignore err"}
		json := gjson.ParseBytes(outBuffer.Bytes())
		assert.Equal(t, json.Get("level").String(), "debug")
		assert.Equal(t, json.Get("error").String(), ErrTest.Error())
		assert.Equal(t, json.Get("Ignore").String(), "HandleErr err")
		assert.Equal(t, json.Get("message").String(), "ignore err")
	})
}

func testResultErrFunc(i int) (int, error) {
	return i, ErrTest
}

func TestResultErr(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				n: 0,
			},
		},
		{
			name: "2",
			args: args{
				n: 100,
			},
		},
		{
			name: "3",
			args: args{
				n: -50,
			},
		},
	}
	for _, tt := range tests {
		t.Run("add msg", func(t *testing.T) {
			he := ResultErr(testResultErrFunc(tt.args.n)).Msg("test msg")

			assert.Equal(t, he.msg, "test msg")
			assert.Equal(t, he.result, tt.args.n)
		})
		t.Run("must err", func(t *testing.T) {
			defer outBuffer.Reset()
			defer func() {
				if err := recover(); err != nil {
					assert.Equal(t, err.(string), "must err")

					json := gjson.ParseBytes(outBuffer.Bytes())
					assert.Equal(t, json.Get("level").String(), "panic")
					assert.Equal(t, json.Get("error").String(), ErrTest.Error())
					assert.Equal(t, json.Get("Must").String(), "ResultErr err")
					assert.Equal(t, json.Get("message").String(), "must err")
				}
			}()
			ResultErr(testResultErrFunc(tt.args.n)).
				Msg("must err").
				Must()
		})
		t.Run("ignore err", func(t *testing.T) {
			defer outBuffer.Reset()
			r := ResultErr(
				testResultErrFunc(tt.args.n),
			).Msg("ignore err").Ignore()

			assert.Equal(t, r, tt.args.n)

			json := gjson.ParseBytes(outBuffer.Bytes())
			assert.Equal(t, json.Get("level").String(), "debug")
			assert.Equal(t, json.Get("error").String(), ErrTest.Error())
			assert.Equal(t, json.Get("Ignore").String(), "ResultErr err")
			assert.Equal(t, json.Get("message").String(), "ignore err")
		})
	}
}

func testResultsErrFunc(i int, s string) (int, string, error) {
	return i, s, ErrTest
}

func TestResultsErr(t *testing.T) {
	type args struct {
		n   int
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				n:   0,
				str: "",
			},
		},
		{
			name: "2",
			args: args{
				n:   100,
				str: "abcdefg",
			},
		},
		{
			name: "3",
			args: args{
				n:   -50,
				str: "123584%$#@^&",
			},
		},
	}
	for _, tt := range tests {
		t.Run("add msg", func(t *testing.T) {
			he := ResultsErr(testResultsErrFunc(tt.args.n, tt.args.str)).Msg("test msg")

			assert.Equal(t, he.msg, "test msg")
			assert.Equal(t, he.result1, tt.args.n)
			assert.Equal(t, he.result2, tt.args.str)
		})
		t.Run("must err", func(t *testing.T) {
			defer outBuffer.Reset()
			defer func() {
				if err := recover(); err != nil {
					assert.Equal(t, err.(string), "must err")

					json := gjson.ParseBytes(outBuffer.Bytes())
					assert.Equal(t, json.Get("level").String(), "panic")
					assert.Equal(t, json.Get("error").String(), ErrTest.Error())
					assert.Equal(t, json.Get("Must").String(), "ResultsErr err")
					assert.Equal(t, json.Get("message").String(), "must err")
				}
			}()
			ResultsErr(testResultsErrFunc(tt.args.n, tt.args.str)).
				Msg("must err").
				Must()
		})
		t.Run("ignore err", func(t *testing.T) {
			defer outBuffer.Reset()
			r1, r2 := ResultsErr(
				testResultsErrFunc(tt.args.n, tt.args.str),
			).Msg("ignore err").Ignore()

			assert.Equal(t, r1, tt.args.n)
			assert.Equal(t, r2, tt.args.str)

			json := gjson.ParseBytes(outBuffer.Bytes())
			assert.Equal(t, json.Get("level").String(), "debug")
			assert.Equal(t, json.Get("error").String(), ErrTest.Error())
			assert.Equal(t, json.Get("Ignore").String(), "ResultsErr err")
			assert.Equal(t, json.Get("message").String(), "ignore err")
		})
	}
}

package result_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Meduzz/helper/fp/result"
)

var errBoom = errors.New("poff")

type message struct {
	Text string
}

func okFunc() (string, error) {
	return "Hello world", nil
}

func failFunc() (string, error) {
	return "", errBoom
}

func StructFunc(msg string) (*message, error) {
	if msg != "error" {
		return &message{msg}, nil
	} else {
		return nil, fmt.Errorf(msg)
	}
}

func Test_Then(t *testing.T) {
	subject := result.Execute(okFunc())

	result, err := result.Then(subject, func(s string) (string, error) {
		return fmt.Sprintf("%s!", s), nil
	}).Get()

	if err != nil {
		t.Errorf("then execution failed")
	}

	if result != "Hello world!" {
		t.Errorf("result was not the expected one, but: %s", result)
	}
}

func Test_AdvancedThen(t *testing.T) {
	subject := result.Execute(StructFunc("Hello world"))

	subject2 := result.Then(subject, func(msg *message) (string, error) {
		return msg.Text, nil
	})

	result, err := result.Then(subject2, func(s string) (*message, error) {
		return &message{s}, nil
	}).Get()

	if err != nil {
		t.Errorf("there was an error: %v", err)
	}

	if result.Text != "Hello world" {
		t.Errorf("execution did not do what I expected: %s", result.Text)
	}
}

func Test_PlainError(t *testing.T) {
	subject := result.Execute(failFunc())

	_, err := subject.Get()

	if !errors.Is(err, errBoom) {
		t.Errorf("error was not boom but: %v", err)
	}
}

func Test_Recover(t *testing.T) {
	subject := result.Failure[string](errBoom)

	result, err := result.Recover(subject, func(err error) string {
		return err.Error()
	}).Get()

	if err != nil {
		t.Errorf("err was not nil, but: %v", err)
	}

	if result != errBoom.Error() {
		t.Errorf("result was not boom, but: %s", result)
	}
}

func Test_GetOrElseString(t *testing.T) {
	subject := result.Execute(failFunc())

	result := result.GetOrElse(subject, func() string {
		return "tada"
	})

	if result != "tada" {
		t.Errorf("result was not tada, but: '%s'", result)
	}
}

func Test_GetOrElseStruct(t *testing.T) {
	subject := result.Execute[*message](nil, errBoom)

	result := result.GetOrElse(subject, func() *message {
		return &message{"tada"}
	})

	if result.Text != "tada" {
		t.Errorf("result was not tada, but: '%s'", result.Text)
	}
}

func Test_GetOrElseHappyString(t *testing.T) {
	subject := result.Execute(okFunc())

	result := result.GetOrElse(subject, func() string {
		return "tada"
	})

	if result != "Hello world" {
		t.Errorf("result was not 'Hello world', but: '%s'", result)
	}
}

func Test_GetOrElseHappyStruct(t *testing.T) {
	subject := result.Success(&message{"Hello world"})

	result := result.GetOrElse(subject, func() *message {
		return &message{"tada"}
	})

	if result.Text != "Hello world" {
		t.Errorf("result was not 'Hello world', but: '%s'", result.Text)
	}
}

func Test_ThenOnFail(t *testing.T) {
	subject := result.Execute(failFunc())

	result, err := result.Then(subject, func(it string) (string, error) {
		return fmt.Sprintf("%s!", it), nil
	}).Get()

	if err == nil {
		t.Error("there was no error")
	}

	if result != "" {
		t.Errorf("there was a result: %s", result)
	}
}

func Test_RecoverOnSuccess(t *testing.T) {
	subject := result.Execute(okFunc())

	result, err := result.Recover(subject, func(err error) string {
		return err.Error()
	}).Get()

	if err != nil {
		t.Errorf("there was an unexpected error: %v", err)
	}

	if result != "Hello world" {
		t.Errorf("the result was not 'Hello world', but: %s", result)
	}
}

func Test_MapHappyCase(t *testing.T) {
	subject := result.Execute[string](okFunc())
	data := result.Map[string, *message](subject, func(in string) *message {
		return &message{in}
	})

	if data == nil {
		t.Error("data is nil")
	}

	if data.IsFailure() {
		t.Error("data is failure")
	}

	msg, err := data.Get()

	if err != nil {
		t.Error("error was set")
	}

	if msg.Text != "Hello world" {
		t.Errorf("message was not '%s' but '%s", "Hello world", msg.Text)
	}
}

func Test_MapUnhappyCase(t *testing.T) {
	subject := result.Execute[string](failFunc())
	data := result.Map(subject, func(in string) *message {
		return &message{in}
	})

	if data == nil {
		t.Error("data was nil")
	}

	if data.IsSuccess() {
		t.Error("data was a success")
	}

	val, err := data.Get()

	if val != nil {
		t.Error("value was set")
	}

	if err != errBoom {
		t.Error("error was not boom")
	}
}

func Test_FlatMapHappyCase(t *testing.T) {
	subject := result.Execute(okFunc())
	data := result.FlatMap(subject, func(in string) *result.Operation[*message] {
		return result.Execute(StructFunc(in))
	})

	if data == nil {
		t.Error("data was nil")
	}

	if data.IsFailure() {
		t.Error("data was a failure")
	}

	msg, err := data.Get()

	if err != nil {
		t.Error("err was set")
	}

	if msg.Text != "Hello world" {
		t.Error("msg was not correct")
	}
}

func Test_FlatMapUnhappyCase(t *testing.T) {
	subject := result.Execute(failFunc())
	data := result.FlatMap(subject, func(in string) *result.Operation[*message] {
		return result.Execute(StructFunc(in))
	})

	if data == nil {
		t.Error("data was nil")
	}

	if data.IsSuccess() {
		t.Error("data was a success")
	}

	msg, err := data.Get()

	if msg != nil {
		t.Error("msg was set")
	}

	if err != errBoom {
		t.Error("error was not boom")
	}
}

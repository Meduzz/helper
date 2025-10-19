package result

import (
	"errors"
	"fmt"
	"testing"
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
	subject := Execute(okFunc())

	result, err := Then(subject, func(s string) (string, error) {
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
	subject := Execute(StructFunc("Hello world"))

	subject2 := Then(subject, func(msg *message) (string, error) {
		return msg.Text, nil
	})

	result, err := Then(subject2, func(s string) (*message, error) {
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
	subject := Execute(failFunc())

	_, err := subject.Get()

	if !errors.Is(err, errBoom) {
		t.Errorf("error was not boom but: %v", err)
	}
}

func Test_Recover(t *testing.T) {
	subject := Failure[string](errBoom)

	result, err := Recover(subject, func(err error) string {
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
	subject := Execute(failFunc())

	result := GetOrElse(subject, func() string {
		return "tada"
	})

	if result != "tada" {
		t.Errorf("result was not tada, but: '%s'", result)
	}
}

func Test_GetOrElseStruct(t *testing.T) {
	subject := Execute[*message](nil, errBoom)

	result := GetOrElse(subject, func() *message {
		return &message{"tada"}
	})

	if result.Text != "tada" {
		t.Errorf("result was not tada, but: '%s'", result.Text)
	}
}

func Test_GetOrElseHappyString(t *testing.T) {
	subject := Execute(okFunc())

	result := GetOrElse(subject, func() string {
		return "tada"
	})

	if result != "Hello world" {
		t.Errorf("result was not 'Hello world', but: '%s'", result)
	}
}

func Test_GetOrElseHappyStruct(t *testing.T) {
	subject := Success(&message{"Hello world"})

	result := GetOrElse(subject, func() *message {
		return &message{"tada"}
	})

	if result.Text != "Hello world" {
		t.Errorf("result was not 'Hello world', but: '%s'", result.Text)
	}
}

func Test_ThenOnFail(t *testing.T) {
	subject := Execute(failFunc())

	result, err := Then(subject, func(it string) (string, error) {
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
	subject := Execute(okFunc())

	result, err := Recover(subject, func(err error) string {
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
	subject := Execute(okFunc())
	data := Map(subject, func(in string) *message {
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
	subject := Execute(failFunc())
	data := Map(subject, func(in string) *message {
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
	subject := Execute(okFunc())
	data := FlatMap(subject, func(in string) *Operation[*message] {
		return Execute(StructFunc(in))
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
	subject := Execute(failFunc())
	data := FlatMap(subject, func(in string) *Operation[*message] {
		return Execute(StructFunc(in))
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

func Test_Batch(t *testing.T) {
	happy := []string{
		"hello",
		"bai",
	}

	unhappy := []string{
		"hello",
		"bai",
		"error",
	}

	t.Run("happy case", func(t *testing.T) {
		subject := Batch(happy, StructFunc)

		if subject.IsFailure() {
			t.Error("subject was a failure")
		}

		items, _ := subject.Get()

		if len(items) != 2 {
			t.Errorf("items length was not 2 but %d", len(items))
		}
	})

	t.Run("unhappy case", func(t *testing.T) {
		subject := Batch(unhappy, StructFunc)

		if subject.IsSuccess() {
			t.Error("subject was a success")
		}

		_, err := subject.Get()

		if err.Error() != "error" {
			t.Errorf("error was not 'error' but '%s'", err.Error())
		}
	})
}

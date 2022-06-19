package fp

import (
	"errors"
	"fmt"
	"testing"
)

var errBoom = errors.New("poff")

type Message struct {
	Text string
}

func OkFunc() (string, error) {
	return "Hello world", nil
}

func FailFunc() (string, error) {
	return "", errBoom
}

func Test_Then(t *testing.T) {
	subject := Execute(OkFunc())

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
	subject := Execute(OkFunc())

	result, err := Then(subject, func(s string) (*Message, error) {
		return &Message{s}, nil
	}).Get()

	if err != nil {
		t.Errorf("there was an error: %v", err)
	}

	if result.Text != "Hello world" {
		t.Errorf("execution did not do what I expected: %s", result.Text)
	}
}

func Test_PlainError(t *testing.T) {
	subject := Execute(FailFunc())

	_, err := subject.Get()

	if !errors.Is(err, errBoom) {
		t.Errorf("error was not boom but: %v", err)
	}
}

func Test_Recover(t *testing.T) {
	subject := Execute(FailFunc())

	result, err := subject.Recover(func(err error) string {
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
	subject := Execute(FailFunc())

	result := subject.GetOrElse(func() string {
		return "tada"
	})

	if result != "tada" {
		t.Errorf("result was not tada, but: '%s'", result)
	}
}

func Test_GetOrElseStruct(t *testing.T) {
	subject := Execute[*Message](nil, nil)

	result := subject.GetOrElse(func() *Message {
		return &Message{"tada"}
	})

	if result.Text != "tada" {
		t.Errorf("result was not tada, but: '%s'", result.Text)
	}
}

func Test_GetOrElseHappyString(t *testing.T) {
	subject := Execute(OkFunc())

	result := subject.GetOrElse(func() string {
		return "tada"
	})

	if result != "Hello world" {
		t.Errorf("result was not 'Hello world', but: '%s'", result)
	}
}

func Test_GetOrElseHappyStruct(t *testing.T) {
	subject := Execute(&Message{"Hello world"}, nil)

	result := subject.GetOrElse(func() *Message {
		return &Message{"tada"}
	})

	if result.Text != "Hello world" {
		t.Errorf("result was not 'Hello world', but: '%s'", result.Text)
	}
}

func Test_ThenOnFail(t *testing.T) {
	subject := Execute(FailFunc())

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
	subject := Execute(OkFunc())

	result, err := subject.Recover(func(err error) string {
		return err.Error()
	}).Get()

	if err != nil {
		t.Errorf("there was an unexpected error: %v", err)
	}

	if result != "Hello world" {
		t.Errorf("the result was not 'Hello world', but: %s", result)
	}
}

func Test_OpFilterToEmpty(t *testing.T) {
	subject := Execute(OkFunc())

	result, err := subject.Filter(func(s string) bool {
		return s == "fail"
	}).Get()

	if err != nil {
		t.Errorf("there was an error: %v", err)
	}

	if result != "" {
		t.Error("the result was not empty/null")
	}
}

func Test_OpFilterToKeep(t *testing.T) {
	subject := Execute(OkFunc())

	result, err := subject.Filter(func(s string) bool {
		return true
	}).Get()

	if err != nil {
		t.Errorf("there was an unexpected error: %v", err)
	}

	if result != "Hello world" {
		t.Errorf("the result was not 'Hello world', but: %s", result)
	}
}

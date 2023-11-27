package test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var (
	AttachmentsRoot = "test_attachments"
)

func AddAttachment(t *testing.T, filename string, a any) {
	err := os.MkdirAll(AttachmentsRoot, 0700)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Create(filepath.Join(AttachmentsRoot, filename))
	if err != nil {
		t.Fatal(err)
	}

	switch v := a.(type) {
	case io.Reader:
		io.Copy(f, v)
	case []byte:
		f.Write(v)
	default:
		fmt.Fprint(f, v)
	}
}

func CheckEqual[T comparable](t *testing.T, testcase int, name string, out, expect T) {
	if out != expect {
		t.Errorf("Testcase %d, %s: out %v, expect %v.", testcase, name, out, expect)
	}
}

func CheckErrorIs(t *testing.T, testcase int, name string, out, expect error) {
	if !errors.Is(out, expect) {
		t.Errorf("Testcase %d, %s: out %v, expect %v.", testcase, name, out, expect)
	}
}

func CheckDeepEqual(t *testing.T, testcase int, name string, out, expect any) {
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("Testcase %d, %s: out %v, expect %v.", testcase, name, out, expect)
	}
}

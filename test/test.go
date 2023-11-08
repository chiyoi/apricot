package test

import (
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
	default:
		fmt.Fprint(f, v)
	}
}

func CheckEqual[T comparable](t *testing.T, testcase int, out, expect T) {
	if out != expect {
		t.Errorf("Testcase %d: out %v, expect %v.", testcase, out, expect)
	}
}

func CheckDeepEqual(t *testing.T, testcase int, out, expect any) {
	if !reflect.DeepEqual(out, expect) {
		t.Errorf("Testcase %d: out %v, expect %v.", testcase, out, expect)
	}
}

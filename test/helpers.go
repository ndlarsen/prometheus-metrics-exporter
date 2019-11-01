package test

import (
	"fmt"
	"testing"
)

func Compare(t *testing.T, err error, errStr string) {

	if err.Error() != errStr {
		_errStr := fmt.Sprintf("%s was not equal to %s", errStr, err.Error())
		t.Fatal(_errStr)
	}

}

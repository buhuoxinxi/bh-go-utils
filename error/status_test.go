package bherror

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	t.Log(err.Error())

	t.Log(err)
}

func TestNewWithError(t *testing.T) {
	oldErr := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	newErr := NewWithError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), oldErr)

	t.Log(newErr.Error())

	t.Log(newErr)

}

func TestFromError(t *testing.T) {
	err := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))

	status, ok := FromError(err)
	if !ok {
		t.Errorf("FromError(err) error : err isn't produced from this package")
		return
	}

	if status.Code != http.StatusNotFound {
		t.Errorf("FromError(err) error : status.Code != input.Code")
	}
	t.Logf("\n%v", err)
}

func TestForward(t *testing.T) {
	err := tNew()

	newErr := Forward(err)

	t.Logf("\n%v", newErr)
}

func tNew() error {
	return New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

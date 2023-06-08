package errors

// Custom Error package to help w errors

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type Error struct {
	Title  string   `json:"title"`
	Status int      `json:"status,omitempty"`
	Traces []string `json:"traces,omitempty"`
}

// New Errs
func New(args ...interface{}) *Error {

	err := &Error{}

	for _, arg := range args {
		switch arg.(type) {
		case Error:
			errTemp := arg.(Error)
			err = &errTemp
		case *Error:
			err = arg.(*Error)
		case string:
			err.Title = arg.(string)
		case error:
			err.Title = arg.(error).Error()
		default:
			err.Title = "unknown error"
		}
	}

	if err.Title == "" {
		err.Title = "Unknown error"
	}

	return err
}

// Error returns the error message
func (e Error) Error() string {
	strErr, _ := jsoniter.MarshalToString(e)
	return strErr
}

// Add line of code for easy error trace
func AddTrace(err interface{}) *Error {

	errs := New(err)
	errs.Traces = append(errs.Traces, getLineOfCode(2))

	return errs
}

// getLineOfCode gets the line of code where the error was created
func getLineOfCode(skip int) string {
	_, filepath, line, _ := runtime.Caller(skip)
	fpath := strings.Split(filepath, "/")
	if len(fpath) > 4 {
		start := len(fpath) - 4
		filepath = "/" + strings.Join(fpath[start:], "/")
	}

	details := fmt.Sprintf(
		"%s[%d]",
		filepath,
		line,
	)

	return details
}

// Check if err2 contains err1 title
func IsContainsTitle(err1 error, err2 error) bool {

	if err1 == nil || err2 == nil {
		return false
	}

	var (
		errs1 string
		errs2 string
	)

	if err1 != nil {
		switch err1.(type) {
		case Error:
			err, ok := err1.(Error)
			if !ok {
				return false
			}
			errs1 = err.Title
		case *Error:
			temp, ok := err1.(*Error)
			if !ok {
				return false
			}
			errs1 = temp.Title
		default:
			errs1 = err1.Error()
		}
	}

	if err2 != nil {
		switch err2.(type) {
		case Error:
			err, ok := err2.(Error)
			if !ok {
				return false
			}
			errs2 = err.Title
		case *Error:
			temp, ok := err2.(*Error)
			if !ok {
				return false
			}
			errs2 = temp.Title
		default:
			errs2 = err2.Error()
		}
	}

	return strings.Contains(strings.ToLower(errs1), strings.ToLower(errs2))
}

func GetHTTPStatus(err error) int {

	errCustom, ok := err.(Error)
	if !ok {
		return http.StatusInternalServerError
	}

	return errCustom.Status
}

func GetErrorMessage(err error) string {

	errCustom, ok := err.(Error)
	if !ok || errCustom.Title == "" {
		return http.StatusText(http.StatusInternalServerError)
	}

	return errCustom.Title
}

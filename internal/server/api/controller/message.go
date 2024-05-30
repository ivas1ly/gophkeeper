package controller

import (
	"errors"
	"strings"

	pv "github.com/bufbuild/protovalidate-go"
)

const (
	MsgInternalServerError = "internal server error"
)

func MsgFromValidationError(err error) string {
	var valErr *pv.ValidationError

	var sb strings.Builder
	if ok := errors.As(err, &valErr); ok {
		valMsg := valErr.ToProto()

		for _, v := range valMsg.GetViolations() {
			sb.WriteString(v.GetFieldPath())
			sb.WriteString(": ")
			sb.WriteString(v.GetMessage())
			sb.WriteString("; ")
		}
	}

	return sb.String()
}

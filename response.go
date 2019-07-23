package webresponse

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	StatusSuccess = "success"
	StatusInfo    = "info"
	StatusError   = "error"
	StatusWarning = "warning"
)

var (
	Successful    error = nil
	ErrBadRequest       = errors.New(
		"We're sorry, but there was an error in communication.  Please try again.",
	)
	ErrForbidden = errors.New(
		"We're sorry, but you are not logged in.  Please login to continue.",
	)
	ErrInternalServer = errors.New(
		"We're sorry, but something went wrong on our end.  Please try again later.",
	)
)

func Error(err error) []byte {
	return Status(err.Error(), StatusError)
}

func Success(message string) []byte {
	return Status(message, StatusSuccess)
}

func Status(status string, statusType string) []byte {
	return General(map[string]string{
		"status":     status,
		"statusType": statusType,
	})
}

func General(payload map[string]string) []byte {
	var response bytes.Buffer
	response.WriteByte('{')

	i := 1
	keyValueTemplate := `"%s": "%s"`
	for key, val := range payload {
		response.WriteString(fmt.Sprintf(keyValueTemplate, key, val))
		if i < len(payload) {
			response.WriteByte(',')
		}
		i += 1
	}

	response.WriteByte('}')
	return response.Bytes()
}

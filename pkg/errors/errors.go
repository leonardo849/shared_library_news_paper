package errors

import (
	"errors"
	"fmt"
	"strings"

	errorsUfb "github.com/leonardo849/utils_for_backend/pkg/errors"
	"gorm.io/gorm"
)

func HandleErrors(err error, model string) (status int, message string) {
	if err == nil {
		return 200, ""
	}
	errMessage := err.Error()
	message = fmt.Sprintf("err: %s, model: %s", errMessage, model)
	if strings.Contains(errMessage, errorsUfb.NOTFOUND) || errors.Is(err, gorm.ErrRecordNotFound) {
		return 404, message
	} else if strings.Contains(errMessage, errorsUfb.CONFLICT) {
		return 409, message
	} else {
		return 500, message
	}
}
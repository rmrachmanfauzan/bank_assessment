package utilities
import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)


func ResponseMessage( err error, message string, data interface{}, ) map[string]interface{} {

	var errorMessages []string
	log := logrus.New()
	
	if err != nil {
		
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be no more than %s characters long", err.Field(), err.Param()))
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", err.Field()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("%s: %s", err.Field(), err.Tag()))
			}
		}
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Validation Errors")	
	}

	return map[string]interface{}{
		"message": message,
		"data":    data,
		"error":  errorMessages,
	}
}

package definition

import (
	"fmt"
	"regexp"
	"strconv"
)

type ErrorDefinition struct {
	ErrorNumber  int64
	ErrorType    string
	Symbol       string
	SQLState     string
	Message      string
	Description  string
	MySQLVersion string
}

// FromError convert from error data
func FromError(inErr error) ErrorDefinition {
	ed := ErrorDefinition{}

	rgx := regexp.MustCompile("^.*: Error +(?P<ErrorNumber>[0-9]+): +(?P<Message>.+)$")
	rawErrMes := inErr.Error()
	if !rgx.MatchString(rawErrMes) {
		return ed
	}

	submatch := rgx.FindStringSubmatch(rawErrMes)
	keys := rgx.SubexpNames()
	for i, key := range keys {
		value := submatch[i]
		switch key {
		case "ErrorNumber":
			en, err := strconv.Atoi(value)
			if err != nil {
				fmt.Printf("%#v\n", err)
				return ed
			}
			errNum := int64(en)
			ed.ErrorNumber = errNum
			ed.ErrorType = ResolveErrorType(errNum)
		case "Message":
			ed.Message = value
		}
	}
	return ed
}

// ResolveErrorType return error type string from ErrorNumber
func ResolveErrorType(errNum int64) string {
	if errNum < 1000 {
		return "GlobalError"
	}
	if errNum >= 2000 && errNum < 3000 {
		return "ClientError"
	}
	return "ServerError"
}

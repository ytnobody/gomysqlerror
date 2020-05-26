package scraper

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/antchfx/htmlquery"
	"github.com/ytnobody/gomysqlerror/definition"
	"golang.org/x/net/html"
)

var (
	URLFormat                = "https://dev.mysql.com/doc/refman/%s/en/%s"
	GlobalErrorReferenceHTML = "global-error-reference.html"
	ClientErrorReferenceHTML = "client-error-reference.html"
	ServerErrorReferenceHTML = "server-error-reference.html"
)

// FetchAllErrorReference fetch all error reference from dev.mysql.com
func FetchAllErrorReference(mysqlVersion string) ([]definition.ErrorDefinition, error) {
	sv, err := FetchServerErrorReference(mysqlVersion)
	if err != nil {
		return nil, err
	}

	cl, err := FetchClientErrorReference(mysqlVersion)
	if err != nil {
		return nil, err
	}

	gl, err := FetchGlobalErrorReference(mysqlVersion)
	if err != nil {
		return nil, err
	}

	errDefList := []definition.ErrorDefinition{}
	for _, errDef := range sv {
		errDefList = append(errDefList, errDef)
	}
	for _, errDef := range cl {
		errDefList = append(errDefList, errDef)
	}
	for _, errDef := range gl {
		errDefList = append(errDefList, errDef)
	}

	return errDefList, nil
}

// FetchGlobalErrorReference fetch global error reference from dev.mysql.com
func FetchGlobalErrorReference(mysqlVersion string) ([]definition.ErrorDefinition, error) {
	errDefList, err := fetchErrorReference(mysqlVersion, GlobalErrorReferenceHTML)
	return errDefList, err
}

// FetchServerErrorReference fetch server error reference from dev.mysql.com
func FetchServerErrorReference(mysqlVersion string) ([]definition.ErrorDefinition, error) {
	errDefList, err := fetchErrorReference(mysqlVersion, ServerErrorReferenceHTML)
	return errDefList, err
}

// FetchClientErrorReference fetch server error reference from dev.mysql.com
func FetchClientErrorReference(mysqlVersion string) ([]definition.ErrorDefinition, error) {
	errDefList, err := fetchErrorReference(mysqlVersion, ClientErrorReferenceHTML)
	return errDefList, err
}

func removeVoidDefinition(errDefList []definition.ErrorDefinition) []definition.ErrorDefinition {
	rtnList := []definition.ErrorDefinition{}
	for _, errDef := range errDefList {
		if errDef.ErrorNumber > 0 {
			rtnList = append(rtnList, errDef)
		}
	}
	return rtnList
}

func fetchErrorReference(mysqlVersion string, document string) ([]definition.ErrorDefinition, error) {
	url := fmt.Sprintf(URLFormat, mysqlVersion, document)

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return nil, err
	}

	errDefList, err := extractErrorDefinitionList(doc, mysqlVersion)
	errDefList = removeVoidDefinition(errDefList)
	return errDefList, err
}

func extractErrorDefinitionList(doc *html.Node, mysqlVersion string) ([]definition.ErrorDefinition, error) {
	items, err := htmlquery.QueryAll(doc, "//*/li[@class=\"listitem\"]")
	if err != nil {
		return nil, err
	}

	errDefList := []definition.ErrorDefinition{}
	for _, item := range items {
		errDef := parseItem(item)
		errDef.MySQLVersion = mysqlVersion
		errDefList = append(errDefList, errDef)
	}

	return errDefList, nil
}

func parseItem(item *html.Node) definition.ErrorDefinition {
	ed := definition.ErrorDefinition{}

	innerText := htmlquery.InnerText(item)
	rgx := regexp.MustCompile(";|\n +\n +")
	messages := rgx.Split(innerText, -1)

	for i, message := range messages {
		rgx = regexp.MustCompile("\n +")
		message = rgx.ReplaceAllString(message, " ")
		rgx = regexp.MustCompile("^ ")
		message = rgx.ReplaceAllString(message, "")
		rgx = regexp.MustCompile("\n$")
		message = rgx.ReplaceAllString(message, "")
		messages[i] = message
	}

	rgx = regexp.MustCompile("^Error number: ")
	if !rgx.Match([]byte(messages[0])) {
		return ed
	}

	for _, message := range messages {
		if len(message) < 1 {
			continue
		}
		rgx = regexp.MustCompile(": ")
		parts := rgx.Split(message, 2)
		switch parts[0] {
		case "Error number":
			errNum, _ := strconv.Atoi(parts[1])
			ed.ErrorNumber = errNum
			ed.ErrorType = resolveErrorType(errNum)
		case "Symbol":
			ed.Symbol = parts[1]
		case "SQLSTATE":
			ed.SQLState = parts[1]
		case "Message":
			ed.Message = parts[1]
		default:
			ed.Description = message
		}
	}

	return ed
}

func resolveErrorType(errNum int) string {
	if errNum < 1000 {
		return "GlobalError"
	}
	if errNum >= 2000 && errNum < 3000 {
		return "ClientError"
	}
	return "ServerError"
}

package scraper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/ytnobody/gomysqlerror/definition"
	"golang.org/x/net/html"
)

var (
	URLFormat                = "https://dev.mysql.com/doc/mysql-errors/%s/en/%s"
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
		ps, err := htmlquery.QueryAll(item, "/p")
		if err != nil {
			return nil, err
		}
		errDef := parseItem(ps)
		errDef.MySQLVersion = mysqlVersion
		errDefList = append(errDefList, errDef)
	}

	return errDefList, nil
}

func parseItem(items []*html.Node) definition.ErrorDefinition {
	ed := definition.ErrorDefinition{}

	innerTexts := []string{}
	for _, item := range items {
		innerTexts = append(innerTexts, htmlquery.InnerText(item))
	}
	innerText := strings.Join(innerTexts, ";\n")
	rgx := regexp.MustCompile(";|\n +\n +")
	messages := rgx.Split(innerText, -1)

	for i, message := range messages {
		rgx = regexp.MustCompile(" *Message:")
		message = rgx.ReplaceAllString(message, "Message:")
		rgx = regexp.MustCompile("\n *")
		message = rgx.ReplaceAllString(message, "")
		rgx = regexp.MustCompile("^ ")
		message = rgx.ReplaceAllString(message, "")
		rgx = regexp.MustCompile("\n$")
		message = rgx.ReplaceAllString(message, "")
		rgx = regexp.MustCompile("^\n+")
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
		rgx = regexp.MustCompile(": *")
		parts := rgx.Split(message, 2)
		switch parts[0] {
		case "Error number":
			en, _ := strconv.Atoi(parts[1])
			errNum := int64(en)
			ed.ErrorNumber = errNum
			ed.ErrorType = definition.ResolveErrorType(errNum)
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

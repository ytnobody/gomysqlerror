package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/ytnobody/gomysqlerror/definition"
	"github.com/ytnobody/gomysqlerror/scraper"
)

func main() {
	flag.Parse()
	args := flag.Args()
	mysqlVersion := args[0]
	generateErrorDefinitions(mysqlVersion)
}

func generateErrorDefinitions(mysqlVersion string) {
	errDefList, err := scraper.FetchAllErrorReference(mysqlVersion)
	if err != nil {
		panic(err)
	}
	src := generateCode(mysqlVersion, errDefList)

	fmt.Println(src)
}

func generateCode(mysqlVersion string, errDefList []definition.ErrorDefinition) string {
	srcFormat := `package v%serror

import (
    "github.com/ytnobody/gomysqlerror/definition"
    "github.com/imdario/mergo"
)

var ErrorMap = map[int]definition.ErrorDefinition{
    %s
}

const (
%s
)

// DefinitionByError fetch definition struct by error data
func DefinitionByError(err error) definition.ErrorDefinition {
    ed := definition.FromError(err)
    defined := ErrorMap[ed.ErrorNumber]
    _ = mergo.Merge(&ed, defined)
    return ed
}

// Isa check a error data that isa specified error number
func Isa(err error, errNum int) bool {
    ed := DefinitionByError(err)
    if ed.ErrorNumber == errNum {
        return true
    }
    return false
}

%s

`
	varCode := generateVarCode(errDefList)
	constCode := generateConstCode(errDefList)
	fnCode := generateFnCode(errDefList)
	packageVersion := strings.ReplaceAll(mysqlVersion, ".", "")
	src := fmt.Sprintf(srcFormat, packageVersion, varCode, constCode, fnCode)

	return src
}

func generateVarCode(errDefList []definition.ErrorDefinition) string {
	varList := []string{}
	for _, errDef := range errDefList {
		varLine := fmt.Sprintf("%d: %#v,", errDef.ErrorNumber, errDef)
		varList = append(varList, varLine)
	}
	varCode := strings.Join(varList, "\n    ")
	return varCode
}

func generateConstCode(errDefList []definition.ErrorDefinition) string {
	constList := []string{}
	for _, errDef := range errDefList {
		constLine := fmt.Sprintf("%s int = %d", errDef.Symbol, errDef.ErrorNumber)
		constList = append(constList, constLine)
	}
	constCode := strings.Join(constList, "\n    ")
	return constCode
}

func generateFnCode(errDefList []definition.ErrorDefinition) string {
	fnList := []string{}
	fnFormat := `
// Is%s check mysql error is "%s" 
func Is%s(err error) bool {
    result := Isa(err, %s)
    return result
}
`
	for _, errDef := range errDefList {
		mes := errDef.Message
		mes = strings.ReplaceAll(mes, "\n", " ")

		symbol := camelizeSymbol(errDef.Symbol)
		fnLine := fmt.Sprintf(fnFormat, symbol, mes, symbol, errDef.Symbol)
		fnList = append(fnList, fnLine)
	}
	fnCode := strings.Join(fnList, "\n    ")
	return fnCode
}

func camelizeSymbol(symbol string) string {
	res := symbol

	rgx := regexp.MustCompile("^ER_")
	res = rgx.ReplaceAllString(res, "SERVER_ERROR_")

	rgx = regexp.MustCompile("^CR_")
	res = rgx.ReplaceAllString(res, "CLIENT_ERROR_")

	rgx = regexp.MustCompile("^EE_")
	res = rgx.ReplaceAllString(res, "GLOBAL_ERROR_")

	res = strings.ToLower(res)
	res = strcase.ToCamel(res)
	return res
}

package generator

import (
	"fmt"
	"strings"

	"github.com/ytnobody/gomysqlerror/definition"
	"github.com/ytnobody/gomysqlerror/scraper"
)

func GenerateErrorDefinitions(mysqlVersion string) {
	errDefList, err := scraper.FetchAllErrorReference(mysqlVersion)
	if err != nil {
		panic(err)
	}
	generateCode(mysqlVersion, errDefList)
}

func generateCode(mysqlVersion string, errDefList []definition.ErrorDefinition) string {
	srcFormat := `
package v%serror

import (
	"github.com/ytnobody/gomysqlerror/definition"
)

var ErrorMap = map[int]definition.ErrorDefinition{
	%s
}

const (
	%s
)

`
	varCode := generateVarCode(errDefList)
	constCode := generateConstCode(errDefList)
	packageVersion := strings.ReplaceAll(mysqlVersion, ".", "")
	src := fmt.Sprintf(srcFormat, packageVersion, varCode, constCode)

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

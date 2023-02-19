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

var V56RetiredErrorSymbols = []string{
	"ER_FOREIGN_DUPLICATE_KEY",
	"ER_CANT_CHANGE_TX_ISOLATION",
	"ER_CANT_LOCK_RPL_INFO_TABLE",
	"ER_NO_SUCH_PARTITION",
	"ER_GTID_MODE_2_OR_3_REQUIRES_DISABLE_GTID_UNSAFE_STATEMENTS_ON",
}

var V57RetiredErrorSymbols = []string{
	"ER_VIEW_SELECT_DERIVED",
	"ER_LOAD_DATA_INVALID_COLUMN",
	"WARN_PLUGIN_DELETE_BUILTIN",
	"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL",
	"ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON",
	"ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON",
	"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF",
	"ER_FK_CANNOT_DELETE_PARENT",
	"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE",
	"ER_KEY_BASED_ON_GENERATED_COLUMN",
	"ER_WARN_DEPRECATED_SQLMODE_UNSET",
	"CR_SECURE_AUTH",
}

var V80RetiredErrorSymbols = []string{
	"ER_WRONG_OUTER_JOIN",
	"ER_GTID_NEXT_TYPE_UNDEFINED_GROUP",
	"ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP",
	"ER_WINDOW_NO_GROUP_ORDER",
	"ER_WINDOW_SE_NOT_ACCEPTABLE",
	"ER_DDL_IN_PROGRESS",
	"ER_TOO_MANY_CONCURRENT_CLONES",
	"ER_REGEXP_STRING_NOT_TERMINATED",
	"ER_INNODB_CANNOT_BE_IGNORED",
	"WARN_DATA_TRUNCATED_FUNCTIONAL_INDEX",
	"CR_COMPRESSION_NOT_SUPPORTED",
}

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
	errDefList = removeRetiredErrorDefinition(mysqlVersion, errDefList)
	src := generateCode(mysqlVersion, errDefList)

	fmt.Println(src)
}

func removeRetiredErrorDefinition(mysqlVersion string, errDefList []definition.ErrorDefinition) []definition.ErrorDefinition {

	res := []definition.ErrorDefinition{}
	retiredSymbols := []string{}
	switch mysqlVersion {
	case "5.6":
		retiredSymbols = V56RetiredErrorSymbols
	case "5.7":
		retiredSymbols = V57RetiredErrorSymbols
	case "8.0":
		retiredSymbols = V80RetiredErrorSymbols
	}

	for _, ed := range errDefList {
		isRetired := false
		for _, retired := range retiredSymbols {
			if ed.Symbol == retired {
				isRetired = true
				break
			}
		}
		if isRetired == false {
			res = append(res, ed)
		}
	}

	return res
}

func generateCode(mysqlVersion string, errDefList []definition.ErrorDefinition) string {
	srcFormat := `package v%serror

import (
    "github.com/ytnobody/gomysqlerror/definition"
    "github.com/imdario/mergo"
)

var ErrorMap = map[int64]definition.ErrorDefinition{
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
func Isa(err error, errNum int64) bool {
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
		varLine = strings.ReplaceAll(varLine, "definition.ErrorDefinition", "")
		varList = append(varList, varLine)
	}
	varCode := strings.Join(varList, "\n    ")
	return varCode
}

func generateConstCode(errDefList []definition.ErrorDefinition) string {
	constList := []string{}
	for _, errDef := range errDefList {
		constLine := fmt.Sprintf("%s int64 = %d", errDef.Symbol, errDef.ErrorNumber)
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

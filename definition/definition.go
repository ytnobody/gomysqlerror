package definition

type ErrorDefinition struct {
	ErrorNumber  int
	ErrorType    string
	Symbol       string
	SQLState     string
	Message      string
	Description  string
	MySQLVersion string
}

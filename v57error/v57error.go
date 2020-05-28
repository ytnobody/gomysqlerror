package v57error

import (
    "github.com/ytnobody/gomysqlerror/definition"
    "github.com/imdario/mergo"
)

var ErrorMap = map[int]definition.ErrorDefinition{
    1000: definition.ErrorDefinition{ErrorNumber:1000, ErrorType:"ServerError", Symbol:"ER_HASHCHK", SQLState:"HY000", Message:"hashchk", Description:"Unused. ", MySQLVersion:"5.7"},
    1001: definition.ErrorDefinition{ErrorNumber:1001, ErrorType:"ServerError", Symbol:"ER_NISAMCHK", SQLState:"HY000", Message:"isamchk", Description:"Unused. ", MySQLVersion:"5.7"},
    1002: definition.ErrorDefinition{ErrorNumber:1002, ErrorType:"ServerError", Symbol:"ER_NO", SQLState:"HY000", Message:"NO", Description:"Used in the construction of other messages. ", MySQLVersion:"5.7"},
    1003: definition.ErrorDefinition{ErrorNumber:1003, ErrorType:"ServerError", Symbol:"ER_YES", SQLState:"HY000", Message:"YES", Description:"Extended EXPLAIN format generates Note messages. ER_YES is used in the Code column for these messages in subsequent SHOW WARNINGS output. ", MySQLVersion:"5.7"},
    1004: definition.ErrorDefinition{ErrorNumber:1004, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_FILE", SQLState:"HY000", Message:"Can't create file '%s' (errno: %d - %s)", Description:"destination file already exists but is not writeable. ", MySQLVersion:"5.7"},
    1005: definition.ErrorDefinition{ErrorNumber:1005, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_TABLE", SQLState:"HY000", Message:"Can't create table '%s' (errno: %d)", Description:"InnoDB reports this error when a table cannot be created. If the error message refers to error 150, table creation failed because a foreign key constraint was not correctly formed. If the error message refers to error −1, table creation probably failed because the table includes a column name that matched the name of an internal InnoDB table. ", MySQLVersion:"5.7"},
    1006: definition.ErrorDefinition{ErrorNumber:1006, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_DB", SQLState:"HY000", Message:"Can't create database '%s' (errno: %d) ", Description:"", MySQLVersion:"5.7"},
    1007: definition.ErrorDefinition{ErrorNumber:1007, ErrorType:"ServerError", Symbol:"ER_DB_CREATE_EXISTS", SQLState:"HY000", Message:"Can't create database '%s'", Description:"Drop the database first if you really want to replace an existing database, or add an IF NOT EXISTS clause to the CREATE DATABASE statement if to retain an existing database without having the statement produce an error. ", MySQLVersion:"5.7"},
    1008: definition.ErrorDefinition{ErrorNumber:1008, ErrorType:"ServerError", Symbol:"ER_DB_DROP_EXISTS", SQLState:"HY000", Message:"Can't drop database '%s'", Description:"database doesn't exist ", MySQLVersion:"5.7"},
    1009: definition.ErrorDefinition{ErrorNumber:1009, ErrorType:"ServerError", Symbol:"ER_DB_DROP_DELETE", SQLState:"HY000", Message:"Error dropping database (can't delete '%s', errno: %d) ", Description:"", MySQLVersion:"5.7"},
    1010: definition.ErrorDefinition{ErrorNumber:1010, ErrorType:"ServerError", Symbol:"ER_DB_DROP_RMDIR", SQLState:"HY000", Message:"Error dropping database (can't rmdir '%s', errno: %d) ", Description:"", MySQLVersion:"5.7"},
    1011: definition.ErrorDefinition{ErrorNumber:1011, ErrorType:"ServerError", Symbol:"ER_CANT_DELETE_FILE", SQLState:"HY000", Message:"Error on delete of '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1012: definition.ErrorDefinition{ErrorNumber:1012, ErrorType:"ServerError", Symbol:"ER_CANT_FIND_SYSTEM_REC", SQLState:"HY000", Message:"Can't read record in system table", Description:"Returned by InnoDB for attempts to access InnoDB INFORMATION_SCHEMA tables when InnoDB is unavailable. ", MySQLVersion:"5.7"},
    1013: definition.ErrorDefinition{ErrorNumber:1013, ErrorType:"ServerError", Symbol:"ER_CANT_GET_STAT", SQLState:"HY000", Message:"Can't get status of '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1014: definition.ErrorDefinition{ErrorNumber:1014, ErrorType:"ServerError", Symbol:"ER_CANT_GET_WD", SQLState:"HY000", Message:"Can't get working directory (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1015: definition.ErrorDefinition{ErrorNumber:1015, ErrorType:"ServerError", Symbol:"ER_CANT_LOCK", SQLState:"HY000", Message:"Can't lock file (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1016: definition.ErrorDefinition{ErrorNumber:1016, ErrorType:"ServerError", Symbol:"ER_CANT_OPEN_FILE", SQLState:"HY000", Message:"Can't open file: '%s' (errno: %d - %s)", Description:"InnoDB reports this error when the table from the InnoDB data files cannot be found, even though the .frm file for the table exists. See Section\u00a014.22.3, “Troubleshooting InnoDB Data Dictionary Operations”. ", MySQLVersion:"5.7"},
    1017: definition.ErrorDefinition{ErrorNumber:1017, ErrorType:"ServerError", Symbol:"ER_FILE_NOT_FOUND", SQLState:"HY000", Message:"Can't find file: '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1018: definition.ErrorDefinition{ErrorNumber:1018, ErrorType:"ServerError", Symbol:"ER_CANT_READ_DIR", SQLState:"HY000", Message:"Can't read dir of '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1019: definition.ErrorDefinition{ErrorNumber:1019, ErrorType:"ServerError", Symbol:"ER_CANT_SET_WD", SQLState:"HY000", Message:"Can't change dir to '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1020: definition.ErrorDefinition{ErrorNumber:1020, ErrorType:"ServerError", Symbol:"ER_CHECKREAD", SQLState:"HY000", Message:"Record has changed since last read in table '%s' ", Description:"", MySQLVersion:"5.7"},
    1021: definition.ErrorDefinition{ErrorNumber:1021, ErrorType:"ServerError", Symbol:"ER_DISK_FULL", SQLState:"HY000", Message:"Disk full (%s)", Description:"waiting for someone to free some space... (errno: %d - %s) ", MySQLVersion:"5.7"},
    1022: definition.ErrorDefinition{ErrorNumber:1022, ErrorType:"ServerError", Symbol:"ER_DUP_KEY", SQLState:"23000", Message:"Can't write", Description:"duplicate key in table '%s' ", MySQLVersion:"5.7"},
    1023: definition.ErrorDefinition{ErrorNumber:1023, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_CLOSE", SQLState:"HY000", Message:"Error on close of '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1024: definition.ErrorDefinition{ErrorNumber:1024, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_READ", SQLState:"HY000", Message:"Error reading file '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1025: definition.ErrorDefinition{ErrorNumber:1025, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_RENAME", SQLState:"HY000", Message:"Error on rename of '%s' to '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1026: definition.ErrorDefinition{ErrorNumber:1026, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_WRITE", SQLState:"HY000", Message:"Error writing file '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1027: definition.ErrorDefinition{ErrorNumber:1027, ErrorType:"ServerError", Symbol:"ER_FILE_USED", SQLState:"HY000", Message:"'%s' is locked against change ", Description:"", MySQLVersion:"5.7"},
    1028: definition.ErrorDefinition{ErrorNumber:1028, ErrorType:"ServerError", Symbol:"ER_FILSORT_ABORT", SQLState:"HY000", Message:"Sort aborted ", Description:"", MySQLVersion:"5.7"},
    1029: definition.ErrorDefinition{ErrorNumber:1029, ErrorType:"ServerError", Symbol:"ER_FORM_NOT_FOUND", SQLState:"HY000", Message:"View '%s' doesn't exist for '%s' ", Description:"", MySQLVersion:"5.7"},
    1030: definition.ErrorDefinition{ErrorNumber:1030, ErrorType:"ServerError", Symbol:"ER_GET_ERRNO", SQLState:"HY000", Message:"Got error %d from storage engine", Description:"Check the %d value to see what the OS error means. For example, 28 indicates that you have run out of disk space. ", MySQLVersion:"5.7"},
    1031: definition.ErrorDefinition{ErrorNumber:1031, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_HA", SQLState:"HY000", Message:"Table storage engine for '%s' doesn't have this option ", Description:"", MySQLVersion:"5.7"},
    1032: definition.ErrorDefinition{ErrorNumber:1032, ErrorType:"ServerError", Symbol:"ER_KEY_NOT_FOUND", SQLState:"HY000", Message:"Can't find record in '%s' ", Description:"", MySQLVersion:"5.7"},
    1033: definition.ErrorDefinition{ErrorNumber:1033, ErrorType:"ServerError", Symbol:"ER_NOT_FORM_FILE", SQLState:"HY000", Message:"Incorrect information in file: '%s' ", Description:"", MySQLVersion:"5.7"},
    1034: definition.ErrorDefinition{ErrorNumber:1034, ErrorType:"ServerError", Symbol:"ER_NOT_KEYFILE", SQLState:"HY000", Message:"Incorrect key file for table '%s'", Description:"try to repair it ", MySQLVersion:"5.7"},
    1035: definition.ErrorDefinition{ErrorNumber:1035, ErrorType:"ServerError", Symbol:"ER_OLD_KEYFILE", SQLState:"HY000", Message:"Old key file for table '%s'", Description:"repair it! ", MySQLVersion:"5.7"},
    1036: definition.ErrorDefinition{ErrorNumber:1036, ErrorType:"ServerError", Symbol:"ER_OPEN_AS_READONLY", SQLState:"HY000", Message:"Table '%s' is read only ", Description:"", MySQLVersion:"5.7"},
    1037: definition.ErrorDefinition{ErrorNumber:1037, ErrorType:"ServerError", Symbol:"ER_OUTOFMEMORY", SQLState:"HY001", Message:"Out of memory", Description:"restart server and try again (needed %d bytes) ", MySQLVersion:"5.7"},
    1038: definition.ErrorDefinition{ErrorNumber:1038, ErrorType:"ServerError", Symbol:"ER_OUT_OF_SORTMEMORY", SQLState:"HY001", Message:"Out of sort memory, consider increasing server sort buffer size ", Description:"", MySQLVersion:"5.7"},
    1039: definition.ErrorDefinition{ErrorNumber:1039, ErrorType:"ServerError", Symbol:"ER_UNEXPECTED_EOF", SQLState:"HY000", Message:"Unexpected EOF found when reading file '%s' (errno: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    1040: definition.ErrorDefinition{ErrorNumber:1040, ErrorType:"ServerError", Symbol:"ER_CON_COUNT_ERROR", SQLState:"08004", Message:"Too many connections ", Description:"", MySQLVersion:"5.7"},
    1041: definition.ErrorDefinition{ErrorNumber:1041, ErrorType:"ServerError", Symbol:"ER_OUT_OF_RESOURCES", SQLState:"HY000", Message:"Out of memory", Description:"if not, you may have to use 'ulimit' to allow mysqld to use more memory or you can add more swap space ", MySQLVersion:"5.7"},
    1042: definition.ErrorDefinition{ErrorNumber:1042, ErrorType:"ServerError", Symbol:"ER_BAD_HOST_ERROR", SQLState:"08S01", Message:"Can't get hostname for your address ", Description:"", MySQLVersion:"5.7"},
    1043: definition.ErrorDefinition{ErrorNumber:1043, ErrorType:"ServerError", Symbol:"ER_HANDSHAKE_ERROR", SQLState:"08S01", Message:"Bad handshake ", Description:"", MySQLVersion:"5.7"},
    1044: definition.ErrorDefinition{ErrorNumber:1044, ErrorType:"ServerError", Symbol:"ER_DBACCESS_DENIED_ERROR", SQLState:"42000", Message:"Access denied for user '%s'@'%s' to database '%s' ", Description:"", MySQLVersion:"5.7"},
    1045: definition.ErrorDefinition{ErrorNumber:1045, ErrorType:"ServerError", Symbol:"ER_ACCESS_DENIED_ERROR", SQLState:"28000", Message:"Access denied for user '%s'@'%s' (using password: %s) ", Description:"", MySQLVersion:"5.7"},
    1046: definition.ErrorDefinition{ErrorNumber:1046, ErrorType:"ServerError", Symbol:"ER_NO_DB_ERROR", SQLState:"3D000", Message:"No database selected ", Description:"", MySQLVersion:"5.7"},
    1047: definition.ErrorDefinition{ErrorNumber:1047, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_COM_ERROR", SQLState:"08S01", Message:"Unknown command ", Description:"", MySQLVersion:"5.7"},
    1048: definition.ErrorDefinition{ErrorNumber:1048, ErrorType:"ServerError", Symbol:"ER_BAD_NULL_ERROR", SQLState:"23000", Message:"Column '%s' cannot be null ", Description:"", MySQLVersion:"5.7"},
    1049: definition.ErrorDefinition{ErrorNumber:1049, ErrorType:"ServerError", Symbol:"ER_BAD_DB_ERROR", SQLState:"42000", Message:"Unknown database '%s' ", Description:"", MySQLVersion:"5.7"},
    1050: definition.ErrorDefinition{ErrorNumber:1050, ErrorType:"ServerError", Symbol:"ER_TABLE_EXISTS_ERROR", SQLState:"42S01", Message:"Table '%s' already exists ", Description:"", MySQLVersion:"5.7"},
    1051: definition.ErrorDefinition{ErrorNumber:1051, ErrorType:"ServerError", Symbol:"ER_BAD_TABLE_ERROR", SQLState:"42S02", Message:"Unknown table '%s' ", Description:"", MySQLVersion:"5.7"},
    1052: definition.ErrorDefinition{ErrorNumber:1052, ErrorType:"ServerError", Symbol:"ER_NON_UNIQ_ERROR", SQLState:"23000", Message:"Column '%s' in %s is ambiguous %s = column name\n%s = location of column (for example, \"field list\") Likely cause: A column appears in a query without appropriate qualification, such as in a select list or ON clause.", Description:"\n", MySQLVersion:"5.7"},
    1053: definition.ErrorDefinition{ErrorNumber:1053, ErrorType:"ServerError", Symbol:"ER_SERVER_SHUTDOWN", SQLState:"08S01", Message:"Server shutdown in progress ", Description:"", MySQLVersion:"5.7"},
    1054: definition.ErrorDefinition{ErrorNumber:1054, ErrorType:"ServerError", Symbol:"ER_BAD_FIELD_ERROR", SQLState:"42S22", Message:"Unknown column '%s' in '%s' ", Description:"", MySQLVersion:"5.7"},
    1055: definition.ErrorDefinition{ErrorNumber:1055, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_WITH_GROUP", SQLState:"42000", Message:"'%s' isn't in GROUP BY ", Description:"", MySQLVersion:"5.7"},
    1056: definition.ErrorDefinition{ErrorNumber:1056, ErrorType:"ServerError", Symbol:"ER_WRONG_GROUP_FIELD", SQLState:"42000", Message:"Can't group on '%s' ", Description:"", MySQLVersion:"5.7"},
    1057: definition.ErrorDefinition{ErrorNumber:1057, ErrorType:"ServerError", Symbol:"ER_WRONG_SUM_SELECT", SQLState:"42000", Message:"Statement has sum functions and columns in same statement ", Description:"", MySQLVersion:"5.7"},
    1058: definition.ErrorDefinition{ErrorNumber:1058, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_COUNT", SQLState:"21S01", Message:"Column count doesn't match value count ", Description:"", MySQLVersion:"5.7"},
    1059: definition.ErrorDefinition{ErrorNumber:1059, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_IDENT", SQLState:"42000", Message:"Identifier name '%s' is too long ", Description:"", MySQLVersion:"5.7"},
    1060: definition.ErrorDefinition{ErrorNumber:1060, ErrorType:"ServerError", Symbol:"ER_DUP_FIELDNAME", SQLState:"42S21", Message:"Duplicate column name '%s' ", Description:"", MySQLVersion:"5.7"},
    1061: definition.ErrorDefinition{ErrorNumber:1061, ErrorType:"ServerError", Symbol:"ER_DUP_KEYNAME", SQLState:"42000", Message:"Duplicate key name '%s' ", Description:"", MySQLVersion:"5.7"},
    1062: definition.ErrorDefinition{ErrorNumber:1062, ErrorType:"ServerError", Symbol:"ER_DUP_ENTRY", SQLState:"23000", Message:"Duplicate entry '%s' for key %d", Description:"The message returned with this error uses the format string for ER_DUP_ENTRY_WITH_KEY_NAME. ", MySQLVersion:"5.7"},
    1063: definition.ErrorDefinition{ErrorNumber:1063, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_SPEC", SQLState:"42000", Message:"Incorrect column specifier for column '%s' ", Description:"", MySQLVersion:"5.7"},
    1064: definition.ErrorDefinition{ErrorNumber:1064, ErrorType:"ServerError", Symbol:"ER_PARSE_ERROR", SQLState:"42000", Message:"%s near '%s' at line %d ", Description:"", MySQLVersion:"5.7"},
    1065: definition.ErrorDefinition{ErrorNumber:1065, ErrorType:"ServerError", Symbol:"ER_EMPTY_QUERY", SQLState:"42000", Message:"Query was empty ", Description:"", MySQLVersion:"5.7"},
    1066: definition.ErrorDefinition{ErrorNumber:1066, ErrorType:"ServerError", Symbol:"ER_NONUNIQ_TABLE", SQLState:"42000", Message:"Not unique table/alias: '%s' ", Description:"", MySQLVersion:"5.7"},
    1067: definition.ErrorDefinition{ErrorNumber:1067, ErrorType:"ServerError", Symbol:"ER_INVALID_DEFAULT", SQLState:"42000", Message:"Invalid default value for '%s' ", Description:"", MySQLVersion:"5.7"},
    1068: definition.ErrorDefinition{ErrorNumber:1068, ErrorType:"ServerError", Symbol:"ER_MULTIPLE_PRI_KEY", SQLState:"42000", Message:"Multiple primary key defined ", Description:"", MySQLVersion:"5.7"},
    1069: definition.ErrorDefinition{ErrorNumber:1069, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_KEYS", SQLState:"42000", Message:"Too many keys specified", Description:"max %d keys allowed ", MySQLVersion:"5.7"},
    1070: definition.ErrorDefinition{ErrorNumber:1070, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_KEY_PARTS", SQLState:"42000", Message:"Too many key parts specified", Description:"max %d parts allowed ", MySQLVersion:"5.7"},
    1071: definition.ErrorDefinition{ErrorNumber:1071, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_KEY", SQLState:"42000", Message:"Specified key was too long", Description:"max key length is %d bytes ", MySQLVersion:"5.7"},
    1072: definition.ErrorDefinition{ErrorNumber:1072, ErrorType:"ServerError", Symbol:"ER_KEY_COLUMN_DOES_NOT_EXITS", SQLState:"42000", Message:"Key column '%s' doesn't exist in table ", Description:"", MySQLVersion:"5.7"},
    1073: definition.ErrorDefinition{ErrorNumber:1073, ErrorType:"ServerError", Symbol:"ER_BLOB_USED_AS_KEY", SQLState:"42000", Message:"BLOB column '%s' can't be used in key specification with the used table type ", Description:"", MySQLVersion:"5.7"},
    1074: definition.ErrorDefinition{ErrorNumber:1074, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_FIELDLENGTH", SQLState:"42000", Message:"Column length too big for column '%s' (max = %lu)", Description:"use BLOB or TEXT instead ", MySQLVersion:"5.7"},
    1075: definition.ErrorDefinition{ErrorNumber:1075, ErrorType:"ServerError", Symbol:"ER_WRONG_AUTO_KEY", SQLState:"42000", Message:"Incorrect table definition", Description:"there can be only one auto column and it must be defined as a key ", MySQLVersion:"5.7"},
    1076: definition.ErrorDefinition{ErrorNumber:1076, ErrorType:"ServerError", Symbol:"ER_READY", SQLState:"HY000", Message:"%s: ready for connections. Version: '%s' socket: '%s' port: %d ", Description:"", MySQLVersion:"5.7"},
    1077: definition.ErrorDefinition{ErrorNumber:1077, ErrorType:"ServerError", Symbol:"ER_NORMAL_SHUTDOWN", SQLState:"HY000", Message:"%s: Normal shutdown ", Description:"", MySQLVersion:"5.7"},
    1078: definition.ErrorDefinition{ErrorNumber:1078, ErrorType:"ServerError", Symbol:"ER_GOT_SIGNAL", SQLState:"HY000", Message:"%s: Got signal %d. Aborting! ", Description:"", MySQLVersion:"5.7"},
    1079: definition.ErrorDefinition{ErrorNumber:1079, ErrorType:"ServerError", Symbol:"ER_SHUTDOWN_COMPLETE", SQLState:"HY000", Message:"%s: Shutdown complete ", Description:"", MySQLVersion:"5.7"},
    1080: definition.ErrorDefinition{ErrorNumber:1080, ErrorType:"ServerError", Symbol:"ER_FORCING_CLOSE", SQLState:"08S01", Message:"%s: Forcing close of thread %ld user: '%s' ", Description:"", MySQLVersion:"5.7"},
    1081: definition.ErrorDefinition{ErrorNumber:1081, ErrorType:"ServerError", Symbol:"ER_IPSOCK_ERROR", SQLState:"08S01", Message:"Can't create IP socket ", Description:"", MySQLVersion:"5.7"},
    1082: definition.ErrorDefinition{ErrorNumber:1082, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_INDEX", SQLState:"42S12", Message:"Table '%s' has no index like the one used in CREATE INDEX", Description:"recreate the table ", MySQLVersion:"5.7"},
    1083: definition.ErrorDefinition{ErrorNumber:1083, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_TERMINATORS", SQLState:"42000", Message:"Field separator argument is not what is expected", Description:"check the manual ", MySQLVersion:"5.7"},
    1084: definition.ErrorDefinition{ErrorNumber:1084, ErrorType:"ServerError", Symbol:"ER_BLOBS_AND_NO_TERMINATED", SQLState:"42000", Message:"You can't use fixed rowlength with BLOBs", Description:"please use 'fields terminated by' ", MySQLVersion:"5.7"},
    1085: definition.ErrorDefinition{ErrorNumber:1085, ErrorType:"ServerError", Symbol:"ER_TEXTFILE_NOT_READABLE", SQLState:"HY000", Message:"The file '%s' must be in the database directory or be readable by all ", Description:"", MySQLVersion:"5.7"},
    1086: definition.ErrorDefinition{ErrorNumber:1086, ErrorType:"ServerError", Symbol:"ER_FILE_EXISTS_ERROR", SQLState:"HY000", Message:"File '%s' already exists ", Description:"", MySQLVersion:"5.7"},
    1087: definition.ErrorDefinition{ErrorNumber:1087, ErrorType:"ServerError", Symbol:"ER_LOAD_INFO", SQLState:"HY000", Message:"Records: %ld Deleted: %ld Skipped: %ld Warnings: %ld ", Description:"", MySQLVersion:"5.7"},
    1088: definition.ErrorDefinition{ErrorNumber:1088, ErrorType:"ServerError", Symbol:"ER_ALTER_INFO", SQLState:"HY000", Message:"Records: %ld Duplicates: %ld ", Description:"", MySQLVersion:"5.7"},
    1089: definition.ErrorDefinition{ErrorNumber:1089, ErrorType:"ServerError", Symbol:"ER_WRONG_SUB_KEY", SQLState:"HY000", Message:"Incorrect prefix key", Description:"the used key part isn't a string, the used length is longer than the key part, or the storage engine doesn't support unique prefix keys ", MySQLVersion:"5.7"},
    1090: definition.ErrorDefinition{ErrorNumber:1090, ErrorType:"ServerError", Symbol:"ER_CANT_REMOVE_ALL_FIELDS", SQLState:"42000", Message:"You can't delete all columns with ALTER TABLE", Description:"use DROP TABLE instead ", MySQLVersion:"5.7"},
    1091: definition.ErrorDefinition{ErrorNumber:1091, ErrorType:"ServerError", Symbol:"ER_CANT_DROP_FIELD_OR_KEY", SQLState:"42000", Message:"Can't DROP '%s'", Description:"check that column/key exists ", MySQLVersion:"5.7"},
    1092: definition.ErrorDefinition{ErrorNumber:1092, ErrorType:"ServerError", Symbol:"ER_INSERT_INFO", SQLState:"HY000", Message:"Records: %ld Duplicates: %ld Warnings: %ld ", Description:"", MySQLVersion:"5.7"},
    1093: definition.ErrorDefinition{ErrorNumber:1093, ErrorType:"ServerError", Symbol:"ER_UPDATE_TABLE_USED", SQLState:"HY000", Message:"You can't specify target table '%s' for update in FROM clause", Description:"This error occurs for attempts to select from and modify the same table within a single statement. If the select attempt occurs within a derived table, you can avoid this error by setting the derived_merge flag of the optimizer_switch system variable to force the subquery to be materialized into a temporary table, which effectively causes it to be a different table from the one modified. See Section\u00a08.2.2.4, “Optimizing Derived Tables and View References with Merging or Materialization”. ", MySQLVersion:"5.7"},
    1094: definition.ErrorDefinition{ErrorNumber:1094, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_THREAD", SQLState:"HY000", Message:"Unknown thread id: %lu ", Description:"", MySQLVersion:"5.7"},
    1095: definition.ErrorDefinition{ErrorNumber:1095, ErrorType:"ServerError", Symbol:"ER_KILL_DENIED_ERROR", SQLState:"HY000", Message:"You are not owner of thread %lu ", Description:"", MySQLVersion:"5.7"},
    1096: definition.ErrorDefinition{ErrorNumber:1096, ErrorType:"ServerError", Symbol:"ER_NO_TABLES_USED", SQLState:"HY000", Message:"No tables used ", Description:"", MySQLVersion:"5.7"},
    1097: definition.ErrorDefinition{ErrorNumber:1097, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_SET", SQLState:"HY000", Message:"Too many strings for column %s and SET ", Description:"", MySQLVersion:"5.7"},
    1098: definition.ErrorDefinition{ErrorNumber:1098, ErrorType:"ServerError", Symbol:"ER_NO_UNIQUE_LOGFILE", SQLState:"HY000", Message:"Can't generate a unique log-filename %s.(1-999) ", Description:"", MySQLVersion:"5.7"},
    1099: definition.ErrorDefinition{ErrorNumber:1099, ErrorType:"ServerError", Symbol:"ER_TABLE_NOT_LOCKED_FOR_WRITE", SQLState:"HY000", Message:"Table '%s' was locked with a READ lock and can't be updated ", Description:"", MySQLVersion:"5.7"},
    1100: definition.ErrorDefinition{ErrorNumber:1100, ErrorType:"ServerError", Symbol:"ER_TABLE_NOT_LOCKED", SQLState:"HY000", Message:"Table '%s' was not locked with LOCK TABLES ", Description:"", MySQLVersion:"5.7"},
    1101: definition.ErrorDefinition{ErrorNumber:1101, ErrorType:"ServerError", Symbol:"ER_BLOB_CANT_HAVE_DEFAULT", SQLState:"42000", Message:"BLOB, TEXT, GEOMETRY or JSON column '%s' can't have a default value ", Description:"", MySQLVersion:"5.7"},
    1102: definition.ErrorDefinition{ErrorNumber:1102, ErrorType:"ServerError", Symbol:"ER_WRONG_DB_NAME", SQLState:"42000", Message:"Incorrect database name '%s' ", Description:"", MySQLVersion:"5.7"},
    1103: definition.ErrorDefinition{ErrorNumber:1103, ErrorType:"ServerError", Symbol:"ER_WRONG_TABLE_NAME", SQLState:"42000", Message:"Incorrect table name '%s' ", Description:"", MySQLVersion:"5.7"},
    1104: definition.ErrorDefinition{ErrorNumber:1104, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_SELECT", SQLState:"42000", Message:"The SELECT would examine more than MAX_JOIN_SIZE rows", Description:"check your WHERE and use SET SQL_BIG_SELECTS=1 or SET MAX_JOIN_SIZE=# if the SELECT is okay ", MySQLVersion:"5.7"},
    1105: definition.ErrorDefinition{ErrorNumber:1105, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_ERROR", SQLState:"HY000", Message:"Unknown error ", Description:"", MySQLVersion:"5.7"},
    1106: definition.ErrorDefinition{ErrorNumber:1106, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_PROCEDURE", SQLState:"42000", Message:"Unknown procedure '%s' ", Description:"", MySQLVersion:"5.7"},
    1107: definition.ErrorDefinition{ErrorNumber:1107, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMCOUNT_TO_PROCEDURE", SQLState:"42000", Message:"Incorrect parameter count to procedure '%s' ", Description:"", MySQLVersion:"5.7"},
    1108: definition.ErrorDefinition{ErrorNumber:1108, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMETERS_TO_PROCEDURE", SQLState:"HY000", Message:"Incorrect parameters to procedure '%s' ", Description:"", MySQLVersion:"5.7"},
    1109: definition.ErrorDefinition{ErrorNumber:1109, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_TABLE", SQLState:"42S02", Message:"Unknown table '%s' in %s ", Description:"", MySQLVersion:"5.7"},
    1110: definition.ErrorDefinition{ErrorNumber:1110, ErrorType:"ServerError", Symbol:"ER_FIELD_SPECIFIED_TWICE", SQLState:"42000", Message:"Column '%s' specified twice ", Description:"", MySQLVersion:"5.7"},
    1111: definition.ErrorDefinition{ErrorNumber:1111, ErrorType:"ServerError", Symbol:"ER_INVALID_GROUP_FUNC_USE", SQLState:"HY000", Message:"Invalid use of group function ", Description:"", MySQLVersion:"5.7"},
    1112: definition.ErrorDefinition{ErrorNumber:1112, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_EXTENSION", SQLState:"42000", Message:"Table '%s' uses an extension that doesn't exist in this MySQL version ", Description:"", MySQLVersion:"5.7"},
    1113: definition.ErrorDefinition{ErrorNumber:1113, ErrorType:"ServerError", Symbol:"ER_TABLE_MUST_HAVE_COLUMNS", SQLState:"42000", Message:"A table must have at least 1 column ", Description:"", MySQLVersion:"5.7"},
    1114: definition.ErrorDefinition{ErrorNumber:1114, ErrorType:"ServerError", Symbol:"ER_RECORD_FILE_FULL", SQLState:"HY000", Message:"The table '%s' is full", Description:"InnoDB reports this error when the system tablespace runs out of free space. Reconfigure the system tablespace to add a new data file. ", MySQLVersion:"5.7"},
    1115: definition.ErrorDefinition{ErrorNumber:1115, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_CHARACTER_SET", SQLState:"42000", Message:"Unknown character set: '%s' ", Description:"", MySQLVersion:"5.7"},
    1116: definition.ErrorDefinition{ErrorNumber:1116, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_TABLES", SQLState:"HY000", Message:"Too many tables", Description:"MySQL can only use %d tables in a join ", MySQLVersion:"5.7"},
    1117: definition.ErrorDefinition{ErrorNumber:1117, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_FIELDS", SQLState:"HY000", Message:"Too many columns ", Description:"", MySQLVersion:"5.7"},
    1118: definition.ErrorDefinition{ErrorNumber:1118, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_ROWSIZE", SQLState:"42000", Message:"Row size too large. The maximum row size for the used table type, not counting BLOBs, is %ld. This includes storage overhead, check the manual. You have to change some columns to TEXT or BLOBs ", Description:"", MySQLVersion:"5.7"},
    1119: definition.ErrorDefinition{ErrorNumber:1119, ErrorType:"ServerError", Symbol:"ER_STACK_OVERRUN", SQLState:"HY000", Message:"Thread stack overrun: Used: %ld of a %ld stack. Use 'mysqld --thread_stack=#' to specify a bigger stack if needed ", Description:"", MySQLVersion:"5.7"},
    1120: definition.ErrorDefinition{ErrorNumber:1120, ErrorType:"ServerError", Symbol:"ER_WRONG_OUTER_JOIN", SQLState:"42000", Message:"Cross dependency found in OUTER JOIN", Description:"examine your ON conditions ", MySQLVersion:"5.7"},
    1121: definition.ErrorDefinition{ErrorNumber:1121, ErrorType:"ServerError", Symbol:"ER_NULL_COLUMN_IN_INDEX", SQLState:"42000", Message:"Table handler doesn't support NULL in given index. Please change column '%s' to be NOT NULL or use another handler ", Description:"", MySQLVersion:"5.7"},
    1122: definition.ErrorDefinition{ErrorNumber:1122, ErrorType:"ServerError", Symbol:"ER_CANT_FIND_UDF", SQLState:"HY000", Message:"Can't load function '%s' ", Description:"", MySQLVersion:"5.7"},
    1123: definition.ErrorDefinition{ErrorNumber:1123, ErrorType:"ServerError", Symbol:"ER_CANT_INITIALIZE_UDF", SQLState:"HY000", Message:"Can't initialize function '%s'", Description:"%s ", MySQLVersion:"5.7"},
    1124: definition.ErrorDefinition{ErrorNumber:1124, ErrorType:"ServerError", Symbol:"ER_UDF_NO_PATHS", SQLState:"HY000", Message:"No paths allowed for shared library ", Description:"", MySQLVersion:"5.7"},
    1125: definition.ErrorDefinition{ErrorNumber:1125, ErrorType:"ServerError", Symbol:"ER_UDF_EXISTS", SQLState:"HY000", Message:"Function '%s' already exists ", Description:"", MySQLVersion:"5.7"},
    1126: definition.ErrorDefinition{ErrorNumber:1126, ErrorType:"ServerError", Symbol:"ER_CANT_OPEN_LIBRARY", SQLState:"HY000", Message:"Can't open shared library '%s' (errno: %d %s) ", Description:"", MySQLVersion:"5.7"},
    1127: definition.ErrorDefinition{ErrorNumber:1127, ErrorType:"ServerError", Symbol:"ER_CANT_FIND_DL_ENTRY", SQLState:"HY000", Message:"Can't find symbol '%s' in library ", Description:"", MySQLVersion:"5.7"},
    1128: definition.ErrorDefinition{ErrorNumber:1128, ErrorType:"ServerError", Symbol:"ER_FUNCTION_NOT_DEFINED", SQLState:"HY000", Message:"Function '%s' is not defined ", Description:"", MySQLVersion:"5.7"},
    1129: definition.ErrorDefinition{ErrorNumber:1129, ErrorType:"ServerError", Symbol:"ER_HOST_IS_BLOCKED", SQLState:"HY000", Message:"Host '%s' is blocked because of many connection errors", Description:"unblock with 'mysqladmin flush-hosts' ", MySQLVersion:"5.7"},
    1130: definition.ErrorDefinition{ErrorNumber:1130, ErrorType:"ServerError", Symbol:"ER_HOST_NOT_PRIVILEGED", SQLState:"HY000", Message:"Host '%s' is not allowed to connect to this MySQL server ", Description:"", MySQLVersion:"5.7"},
    1131: definition.ErrorDefinition{ErrorNumber:1131, ErrorType:"ServerError", Symbol:"ER_PASSWORD_ANONYMOUS_USER", SQLState:"42000", Message:"You are using MySQL as an anonymous user and anonymous users are not allowed to change passwords ", Description:"", MySQLVersion:"5.7"},
    1132: definition.ErrorDefinition{ErrorNumber:1132, ErrorType:"ServerError", Symbol:"ER_PASSWORD_NOT_ALLOWED", SQLState:"42000", Message:"You must have privileges to update tables in the mysql database to be able to change passwords for others ", Description:"", MySQLVersion:"5.7"},
    1133: definition.ErrorDefinition{ErrorNumber:1133, ErrorType:"ServerError", Symbol:"ER_PASSWORD_NO_MATCH", SQLState:"42000", Message:"Can't find any matching row in the user table ", Description:"", MySQLVersion:"5.7"},
    1134: definition.ErrorDefinition{ErrorNumber:1134, ErrorType:"ServerError", Symbol:"ER_UPDATE_INFO", SQLState:"HY000", Message:"Rows matched: %ld Changed: %ld Warnings: %ld ", Description:"", MySQLVersion:"5.7"},
    1135: definition.ErrorDefinition{ErrorNumber:1135, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_THREAD", SQLState:"HY000", Message:"Can't create a new thread (errno %d)", Description:"if you are not out of available memory, you can consult the manual for a possible OS-dependent bug ", MySQLVersion:"5.7"},
    1136: definition.ErrorDefinition{ErrorNumber:1136, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_COUNT_ON_ROW", SQLState:"21S01", Message:"Column count doesn't match value count at row %ld ", Description:"", MySQLVersion:"5.7"},
    1137: definition.ErrorDefinition{ErrorNumber:1137, ErrorType:"ServerError", Symbol:"ER_CANT_REOPEN_TABLE", SQLState:"HY000", Message:"Can't reopen table: '%s' ", Description:"", MySQLVersion:"5.7"},
    1138: definition.ErrorDefinition{ErrorNumber:1138, ErrorType:"ServerError", Symbol:"ER_INVALID_USE_OF_NULL", SQLState:"22004", Message:"Invalid use of NULL value ", Description:"", MySQLVersion:"5.7"},
    1139: definition.ErrorDefinition{ErrorNumber:1139, ErrorType:"ServerError", Symbol:"ER_REGEXP_ERROR", SQLState:"42000", Message:"Got error '%s' from regexp ", Description:"", MySQLVersion:"5.7"},
    1140: definition.ErrorDefinition{ErrorNumber:1140, ErrorType:"ServerError", Symbol:"ER_MIX_OF_GROUP_FUNC_AND_FIELDS", SQLState:"42000", Message:"Mixing of GROUP columns (MIN(),MAX(),COUNT(),...) with no GROUP columns is illegal if there is no GROUP BY clause ", Description:"", MySQLVersion:"5.7"},
    1141: definition.ErrorDefinition{ErrorNumber:1141, ErrorType:"ServerError", Symbol:"ER_NONEXISTING_GRANT", SQLState:"42000", Message:"There is no such grant defined for user '%s' on host '%s' ", Description:"", MySQLVersion:"5.7"},
    1142: definition.ErrorDefinition{ErrorNumber:1142, ErrorType:"ServerError", Symbol:"ER_TABLEACCESS_DENIED_ERROR", SQLState:"42000", Message:"%s command denied to user '%s'@'%s' for table '%s' ", Description:"", MySQLVersion:"5.7"},
    1143: definition.ErrorDefinition{ErrorNumber:1143, ErrorType:"ServerError", Symbol:"ER_COLUMNACCESS_DENIED_ERROR", SQLState:"42000", Message:"%s command denied to user '%s'@'%s' for column '%s' in table '%s' ", Description:"", MySQLVersion:"5.7"},
    1144: definition.ErrorDefinition{ErrorNumber:1144, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_GRANT_FOR_TABLE", SQLState:"42000", Message:"Illegal GRANT/REVOKE command", Description:"please consult the manual to see which privileges can be used ", MySQLVersion:"5.7"},
    1145: definition.ErrorDefinition{ErrorNumber:1145, ErrorType:"ServerError", Symbol:"ER_GRANT_WRONG_HOST_OR_USER", SQLState:"42000", Message:"The host or user argument to GRANT is too long ", Description:"", MySQLVersion:"5.7"},
    1146: definition.ErrorDefinition{ErrorNumber:1146, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_TABLE", SQLState:"42S02", Message:"Table '%s.%s' doesn't exist ", Description:"", MySQLVersion:"5.7"},
    1147: definition.ErrorDefinition{ErrorNumber:1147, ErrorType:"ServerError", Symbol:"ER_NONEXISTING_TABLE_GRANT", SQLState:"42000", Message:"There is no such grant defined for user '%s' on host '%s' on table '%s' ", Description:"", MySQLVersion:"5.7"},
    1148: definition.ErrorDefinition{ErrorNumber:1148, ErrorType:"ServerError", Symbol:"ER_NOT_ALLOWED_COMMAND", SQLState:"42000", Message:"The used command is not allowed with this MySQL version ", Description:"", MySQLVersion:"5.7"},
    1149: definition.ErrorDefinition{ErrorNumber:1149, ErrorType:"ServerError", Symbol:"ER_SYNTAX_ERROR", SQLState:"42000", Message:"You have an error in your SQL syntax", Description:"check the manual that corresponds to your MySQL server version for the right syntax to use ", MySQLVersion:"5.7"},
    1150: definition.ErrorDefinition{ErrorNumber:1150, ErrorType:"ServerError", Symbol:"ER_UNUSED1", SQLState:"HY000", Message:"Delayed insert thread couldn't get requested lock for table %s ", Description:"", MySQLVersion:"5.7"},
    1151: definition.ErrorDefinition{ErrorNumber:1151, ErrorType:"ServerError", Symbol:"ER_UNUSED2", SQLState:"HY000", Message:"Too many delayed threads in use ", Description:"", MySQLVersion:"5.7"},
    1152: definition.ErrorDefinition{ErrorNumber:1152, ErrorType:"ServerError", Symbol:"ER_ABORTING_CONNECTION", SQLState:"08S01", Message:"Aborted connection %ld to db: '%s' user: '%s' (%s) ", Description:"", MySQLVersion:"5.7"},
    1153: definition.ErrorDefinition{ErrorNumber:1153, ErrorType:"ServerError", Symbol:"ER_NET_PACKET_TOO_LARGE", SQLState:"08S01", Message:"Got a packet bigger than 'max_allowed_packet' bytes ", Description:"", MySQLVersion:"5.7"},
    1154: definition.ErrorDefinition{ErrorNumber:1154, ErrorType:"ServerError", Symbol:"ER_NET_READ_ERROR_FROM_PIPE", SQLState:"08S01", Message:"Got a read error from the connection pipe ", Description:"", MySQLVersion:"5.7"},
    1155: definition.ErrorDefinition{ErrorNumber:1155, ErrorType:"ServerError", Symbol:"ER_NET_FCNTL_ERROR", SQLState:"08S01", Message:"Got an error from fcntl() ", Description:"", MySQLVersion:"5.7"},
    1156: definition.ErrorDefinition{ErrorNumber:1156, ErrorType:"ServerError", Symbol:"ER_NET_PACKETS_OUT_OF_ORDER", SQLState:"08S01", Message:"Got packets out of order ", Description:"", MySQLVersion:"5.7"},
    1157: definition.ErrorDefinition{ErrorNumber:1157, ErrorType:"ServerError", Symbol:"ER_NET_UNCOMPRESS_ERROR", SQLState:"08S01", Message:"Couldn't uncompress communication packet ", Description:"", MySQLVersion:"5.7"},
    1158: definition.ErrorDefinition{ErrorNumber:1158, ErrorType:"ServerError", Symbol:"ER_NET_READ_ERROR", SQLState:"08S01", Message:"Got an error reading communication packets ", Description:"", MySQLVersion:"5.7"},
    1159: definition.ErrorDefinition{ErrorNumber:1159, ErrorType:"ServerError", Symbol:"ER_NET_READ_INTERRUPTED", SQLState:"08S01", Message:"Got timeout reading communication packets ", Description:"", MySQLVersion:"5.7"},
    1160: definition.ErrorDefinition{ErrorNumber:1160, ErrorType:"ServerError", Symbol:"ER_NET_ERROR_ON_WRITE", SQLState:"08S01", Message:"Got an error writing communication packets ", Description:"", MySQLVersion:"5.7"},
    1161: definition.ErrorDefinition{ErrorNumber:1161, ErrorType:"ServerError", Symbol:"ER_NET_WRITE_INTERRUPTED", SQLState:"08S01", Message:"Got timeout writing communication packets ", Description:"", MySQLVersion:"5.7"},
    1162: definition.ErrorDefinition{ErrorNumber:1162, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_STRING", SQLState:"42000", Message:"Result string is longer than 'max_allowed_packet' bytes ", Description:"", MySQLVersion:"5.7"},
    1163: definition.ErrorDefinition{ErrorNumber:1163, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_BLOB", SQLState:"42000", Message:"The used table type doesn't support BLOB/TEXT columns ", Description:"", MySQLVersion:"5.7"},
    1164: definition.ErrorDefinition{ErrorNumber:1164, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_AUTO_INCREMENT", SQLState:"42000", Message:"The used table type doesn't support AUTO_INCREMENT columns ", Description:"", MySQLVersion:"5.7"},
    1165: definition.ErrorDefinition{ErrorNumber:1165, ErrorType:"ServerError", Symbol:"ER_UNUSED3", SQLState:"HY000", Message:"INSERT DELAYED can't be used with table '%s' because it is locked with LOCK TABLES ", Description:"", MySQLVersion:"5.7"},
    1166: definition.ErrorDefinition{ErrorNumber:1166, ErrorType:"ServerError", Symbol:"ER_WRONG_COLUMN_NAME", SQLState:"42000", Message:"Incorrect column name '%s' ", Description:"", MySQLVersion:"5.7"},
    1167: definition.ErrorDefinition{ErrorNumber:1167, ErrorType:"ServerError", Symbol:"ER_WRONG_KEY_COLUMN", SQLState:"42000", Message:"The used storage engine can't index column '%s' ", Description:"", MySQLVersion:"5.7"},
    1168: definition.ErrorDefinition{ErrorNumber:1168, ErrorType:"ServerError", Symbol:"ER_WRONG_MRG_TABLE", SQLState:"HY000", Message:"Unable to open underlying table which is differently defined or of non-MyISAM type or doesn't exist ", Description:"", MySQLVersion:"5.7"},
    1169: definition.ErrorDefinition{ErrorNumber:1169, ErrorType:"ServerError", Symbol:"ER_DUP_UNIQUE", SQLState:"23000", Message:"Can't write, because of unique constraint, to table '%s' ", Description:"", MySQLVersion:"5.7"},
    1170: definition.ErrorDefinition{ErrorNumber:1170, ErrorType:"ServerError", Symbol:"ER_BLOB_KEY_WITHOUT_LENGTH", SQLState:"42000", Message:"BLOB/TEXT column '%s' used in key specification without a key length ", Description:"", MySQLVersion:"5.7"},
    1171: definition.ErrorDefinition{ErrorNumber:1171, ErrorType:"ServerError", Symbol:"ER_PRIMARY_CANT_HAVE_NULL", SQLState:"42000", Message:"All parts of a PRIMARY KEY must be NOT NULL", Description:"if you need NULL in a key, use UNIQUE instead ", MySQLVersion:"5.7"},
    1172: definition.ErrorDefinition{ErrorNumber:1172, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_ROWS", SQLState:"42000", Message:"Result consisted of more than one row ", Description:"", MySQLVersion:"5.7"},
    1173: definition.ErrorDefinition{ErrorNumber:1173, ErrorType:"ServerError", Symbol:"ER_REQUIRES_PRIMARY_KEY", SQLState:"42000", Message:"This table type requires a primary key ", Description:"", MySQLVersion:"5.7"},
    1174: definition.ErrorDefinition{ErrorNumber:1174, ErrorType:"ServerError", Symbol:"ER_NO_RAID_COMPILED", SQLState:"HY000", Message:"This version of MySQL is not compiled with RAID support ", Description:"", MySQLVersion:"5.7"},
    1175: definition.ErrorDefinition{ErrorNumber:1175, ErrorType:"ServerError", Symbol:"ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE", SQLState:"HY000", Message:"You are using safe update mode and you tried to update a table without a WHERE that uses a KEY column. %s ", Description:"", MySQLVersion:"5.7"},
    1176: definition.ErrorDefinition{ErrorNumber:1176, ErrorType:"ServerError", Symbol:"ER_KEY_DOES_NOT_EXITS", SQLState:"42000", Message:"Key '%s' doesn't exist in table '%s' ", Description:"", MySQLVersion:"5.7"},
    1177: definition.ErrorDefinition{ErrorNumber:1177, ErrorType:"ServerError", Symbol:"ER_CHECK_NO_SUCH_TABLE", SQLState:"42000", Message:"Can't open table ", Description:"", MySQLVersion:"5.7"},
    1178: definition.ErrorDefinition{ErrorNumber:1178, ErrorType:"ServerError", Symbol:"ER_CHECK_NOT_IMPLEMENTED", SQLState:"42000", Message:"The storage engine for the table doesn't support %s ", Description:"", MySQLVersion:"5.7"},
    1179: definition.ErrorDefinition{ErrorNumber:1179, ErrorType:"ServerError", Symbol:"ER_CANT_DO_THIS_DURING_AN_TRANSACTION", SQLState:"25000", Message:"You are not allowed to execute this command in a transaction ", Description:"", MySQLVersion:"5.7"},
    1180: definition.ErrorDefinition{ErrorNumber:1180, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_COMMIT", SQLState:"HY000", Message:"Got error %d during COMMIT ", Description:"", MySQLVersion:"5.7"},
    1181: definition.ErrorDefinition{ErrorNumber:1181, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_ROLLBACK", SQLState:"HY000", Message:"Got error %d during ROLLBACK ", Description:"", MySQLVersion:"5.7"},
    1182: definition.ErrorDefinition{ErrorNumber:1182, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_FLUSH_LOGS", SQLState:"HY000", Message:"Got error %d during FLUSH_LOGS ", Description:"", MySQLVersion:"5.7"},
    1183: definition.ErrorDefinition{ErrorNumber:1183, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_CHECKPOINT", SQLState:"HY000", Message:"Got error %d during CHECKPOINT ", Description:"", MySQLVersion:"5.7"},
    1184: definition.ErrorDefinition{ErrorNumber:1184, ErrorType:"ServerError", Symbol:"ER_NEW_ABORTING_CONNECTION", SQLState:"08S01", Message:"Aborted connection %u to db: '%s' user: '%s' host: '%s' (%s) ", Description:"", MySQLVersion:"5.7"},
    1185: definition.ErrorDefinition{ErrorNumber:1185, ErrorType:"ServerError", Symbol:"ER_DUMP_NOT_IMPLEMENTED", SQLState:"HY000", Message:"The storage engine for the table does not support binary table dump ", Description:"", MySQLVersion:"5.7"},
    1186: definition.ErrorDefinition{ErrorNumber:1186, ErrorType:"ServerError", Symbol:"ER_FLUSH_MASTER_BINLOG_CLOSED", SQLState:"HY000", Message:"Binlog closed, cannot RESET MASTER ", Description:"", MySQLVersion:"5.7"},
    1187: definition.ErrorDefinition{ErrorNumber:1187, ErrorType:"ServerError", Symbol:"ER_INDEX_REBUILD", SQLState:"HY000", Message:"Failed rebuilding the index of dumped table '%s' ", Description:"", MySQLVersion:"5.7"},
    1188: definition.ErrorDefinition{ErrorNumber:1188, ErrorType:"ServerError", Symbol:"ER_MASTER", SQLState:"HY000", Message:"Error from master: '%s' ", Description:"", MySQLVersion:"5.7"},
    1189: definition.ErrorDefinition{ErrorNumber:1189, ErrorType:"ServerError", Symbol:"ER_MASTER_NET_READ", SQLState:"08S01", Message:"Net error reading from master ", Description:"", MySQLVersion:"5.7"},
    1190: definition.ErrorDefinition{ErrorNumber:1190, ErrorType:"ServerError", Symbol:"ER_MASTER_NET_WRITE", SQLState:"08S01", Message:"Net error writing to master ", Description:"", MySQLVersion:"5.7"},
    1191: definition.ErrorDefinition{ErrorNumber:1191, ErrorType:"ServerError", Symbol:"ER_FT_MATCHING_KEY_NOT_FOUND", SQLState:"HY000", Message:"Can't find FULLTEXT index matching the column list ", Description:"", MySQLVersion:"5.7"},
    1192: definition.ErrorDefinition{ErrorNumber:1192, ErrorType:"ServerError", Symbol:"ER_LOCK_OR_ACTIVE_TRANSACTION", SQLState:"HY000", Message:"Can't execute the given command because you have active locked tables or an active transaction ", Description:"", MySQLVersion:"5.7"},
    1193: definition.ErrorDefinition{ErrorNumber:1193, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_SYSTEM_VARIABLE", SQLState:"HY000", Message:"Unknown system variable '%s' ", Description:"", MySQLVersion:"5.7"},
    1194: definition.ErrorDefinition{ErrorNumber:1194, ErrorType:"ServerError", Symbol:"ER_CRASHED_ON_USAGE", SQLState:"HY000", Message:"Table '%s' is marked as crashed and should be repaired ", Description:"", MySQLVersion:"5.7"},
    1195: definition.ErrorDefinition{ErrorNumber:1195, ErrorType:"ServerError", Symbol:"ER_CRASHED_ON_REPAIR", SQLState:"HY000", Message:"Table '%s' is marked as crashed and last (automatic?) repair failed ", Description:"", MySQLVersion:"5.7"},
    1196: definition.ErrorDefinition{ErrorNumber:1196, ErrorType:"ServerError", Symbol:"ER_WARNING_NOT_COMPLETE_ROLLBACK", SQLState:"HY000", Message:"Some non-transactional changed tables couldn't be rolled back ", Description:"", MySQLVersion:"5.7"},
    1197: definition.ErrorDefinition{ErrorNumber:1197, ErrorType:"ServerError", Symbol:"ER_TRANS_CACHE_FULL", SQLState:"HY000", Message:"Multi-statement transaction required more than 'max_binlog_cache_size' bytes of storage", Description:"increase this mysqld variable and try again ", MySQLVersion:"5.7"},
    1198: definition.ErrorDefinition{ErrorNumber:1198, ErrorType:"ServerError", Symbol:"ER_SLAVE_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with a running slave", Description:"run STOP SLAVE first ", MySQLVersion:"5.7"},
    1199: definition.ErrorDefinition{ErrorNumber:1199, ErrorType:"ServerError", Symbol:"ER_SLAVE_NOT_RUNNING", SQLState:"HY000", Message:"This operation requires a running slave", Description:"configure slave and do START SLAVE ", MySQLVersion:"5.7"},
    1200: definition.ErrorDefinition{ErrorNumber:1200, ErrorType:"ServerError", Symbol:"ER_BAD_SLAVE", SQLState:"HY000", Message:"The server is not configured as slave", Description:"fix in config file or with CHANGE MASTER TO ", MySQLVersion:"5.7"},
    1201: definition.ErrorDefinition{ErrorNumber:1201, ErrorType:"ServerError", Symbol:"ER_MASTER_INFO", SQLState:"HY000", Message:"Could not initialize master info structure", Description:"more error messages can be found in the MySQL error log ", MySQLVersion:"5.7"},
    1202: definition.ErrorDefinition{ErrorNumber:1202, ErrorType:"ServerError", Symbol:"ER_SLAVE_THREAD", SQLState:"HY000", Message:"Could not create slave thread", Description:"check system resources ", MySQLVersion:"5.7"},
    1203: definition.ErrorDefinition{ErrorNumber:1203, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_USER_CONNECTIONS", SQLState:"42000", Message:"User %s already has more than 'max_user_connections' active connections ", Description:"", MySQLVersion:"5.7"},
    1204: definition.ErrorDefinition{ErrorNumber:1204, ErrorType:"ServerError", Symbol:"ER_SET_CONSTANTS_ONLY", SQLState:"HY000", Message:"You may only use constant expressions with SET ", Description:"", MySQLVersion:"5.7"},
    1205: definition.ErrorDefinition{ErrorNumber:1205, ErrorType:"ServerError", Symbol:"ER_LOCK_WAIT_TIMEOUT", SQLState:"HY000", Message:"Lock wait timeout exceeded", Description:"InnoDB reports this error when lock wait timeout expires. The statement that waited too long was rolled back (not the entire transaction). You can increase the value of the innodb_lock_wait_timeout configuration option if SQL statements should wait longer for other transactions to complete, or decrease it if too many long-running transactions are causing locking problems and reducing concurrency on a busy system. ", MySQLVersion:"5.7"},
    1206: definition.ErrorDefinition{ErrorNumber:1206, ErrorType:"ServerError", Symbol:"ER_LOCK_TABLE_FULL", SQLState:"HY000", Message:"The total number of locks exceeds the lock table size", Description:"InnoDB reports this error when the total number of locks exceeds the amount of memory devoted to managing locks. To avoid this error, increase the value of innodb_buffer_pool_size. Within an individual application, a workaround may be to break a large operation into smaller pieces. For example, if the error occurs for a large INSERT, perform several smaller INSERT operations. ", MySQLVersion:"5.7"},
    1207: definition.ErrorDefinition{ErrorNumber:1207, ErrorType:"ServerError", Symbol:"ER_READ_ONLY_TRANSACTION", SQLState:"25000", Message:"Update locks cannot be acquired during a READ UNCOMMITTED transaction ", Description:"", MySQLVersion:"5.7"},
    1208: definition.ErrorDefinition{ErrorNumber:1208, ErrorType:"ServerError", Symbol:"ER_DROP_DB_WITH_READ_LOCK", SQLState:"HY000", Message:"DROP DATABASE not allowed while thread is holding global read lock ", Description:"", MySQLVersion:"5.7"},
    1209: definition.ErrorDefinition{ErrorNumber:1209, ErrorType:"ServerError", Symbol:"ER_CREATE_DB_WITH_READ_LOCK", SQLState:"HY000", Message:"CREATE DATABASE not allowed while thread is holding global read lock ", Description:"", MySQLVersion:"5.7"},
    1210: definition.ErrorDefinition{ErrorNumber:1210, ErrorType:"ServerError", Symbol:"ER_WRONG_ARGUMENTS", SQLState:"HY000", Message:"Incorrect arguments to %s ", Description:"", MySQLVersion:"5.7"},
    1211: definition.ErrorDefinition{ErrorNumber:1211, ErrorType:"ServerError", Symbol:"ER_NO_PERMISSION_TO_CREATE_USER", SQLState:"42000", Message:"'%s'@'%s' is not allowed to create new users ", Description:"", MySQLVersion:"5.7"},
    1212: definition.ErrorDefinition{ErrorNumber:1212, ErrorType:"ServerError", Symbol:"ER_UNION_TABLES_IN_DIFFERENT_DIR", SQLState:"HY000", Message:"Incorrect table definition", Description:"all MERGE tables must be in the same database ", MySQLVersion:"5.7"},
    1213: definition.ErrorDefinition{ErrorNumber:1213, ErrorType:"ServerError", Symbol:"ER_LOCK_DEADLOCK", SQLState:"40001", Message:"Deadlock found when trying to get lock", Description:"InnoDB reports this error when a transaction encounters a deadlock and is automatically rolled back so that your application can take corrective action. To recover from this error, run all the operations in this transaction again. A deadlock occurs when requests for locks arrive in inconsistent order between transactions. The transaction that was rolled back released all its locks, and the other transaction can now get all the locks it requested. Thus, when you re-run the transaction that was rolled back, it might have to wait for other transactions to complete, but typically the deadlock does not recur. If you encounter frequent deadlocks, make the sequence of locking operations (LOCK TABLES, SELECT ... FOR UPDATE, and so on) consistent between the different transactions or applications that experience the issue. See Section\u00a014.7.5, “Deadlocks in InnoDB” for details. ", MySQLVersion:"5.7"},
    1214: definition.ErrorDefinition{ErrorNumber:1214, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_FT", SQLState:"HY000", Message:"The used table type doesn't support FULLTEXT indexes ", Description:"", MySQLVersion:"5.7"},
    1215: definition.ErrorDefinition{ErrorNumber:1215, ErrorType:"ServerError", Symbol:"ER_CANNOT_ADD_FOREIGN", SQLState:"HY000", Message:"Cannot add foreign key constraint ", Description:"", MySQLVersion:"5.7"},
    1216: definition.ErrorDefinition{ErrorNumber:1216, ErrorType:"ServerError", Symbol:"ER_NO_REFERENCED_ROW", SQLState:"23000", Message:"Cannot add or update a child row: a foreign key constraint fails", Description:"InnoDB reports this error when you try to add a row but there is no parent row, and a foreign key constraint fails. Add the parent row first. ", MySQLVersion:"5.7"},
    1217: definition.ErrorDefinition{ErrorNumber:1217, ErrorType:"ServerError", Symbol:"ER_ROW_IS_REFERENCED", SQLState:"23000", Message:"Cannot delete or update a parent row: a foreign key constraint fails", Description:"InnoDB reports this error when you try to delete a parent row that has children, and a foreign key constraint fails. Delete the children first. ", MySQLVersion:"5.7"},
    1218: definition.ErrorDefinition{ErrorNumber:1218, ErrorType:"ServerError", Symbol:"ER_CONNECT_TO_MASTER", SQLState:"08S01", Message:"Error connecting to master: %s ", Description:"", MySQLVersion:"5.7"},
    1219: definition.ErrorDefinition{ErrorNumber:1219, ErrorType:"ServerError", Symbol:"ER_QUERY_ON_MASTER", SQLState:"HY000", Message:"Error running query on master: %s ", Description:"", MySQLVersion:"5.7"},
    1220: definition.ErrorDefinition{ErrorNumber:1220, ErrorType:"ServerError", Symbol:"ER_ERROR_WHEN_EXECUTING_COMMAND", SQLState:"HY000", Message:"Error when executing command %s: %s ", Description:"", MySQLVersion:"5.7"},
    1221: definition.ErrorDefinition{ErrorNumber:1221, ErrorType:"ServerError", Symbol:"ER_WRONG_USAGE", SQLState:"HY000", Message:"Incorrect usage of %s and %s ", Description:"", MySQLVersion:"5.7"},
    1222: definition.ErrorDefinition{ErrorNumber:1222, ErrorType:"ServerError", Symbol:"ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT", SQLState:"21000", Message:"The used SELECT statements have a different number of columns ", Description:"", MySQLVersion:"5.7"},
    1223: definition.ErrorDefinition{ErrorNumber:1223, ErrorType:"ServerError", Symbol:"ER_CANT_UPDATE_WITH_READLOCK", SQLState:"HY000", Message:"Can't execute the query because you have a conflicting read lock ", Description:"", MySQLVersion:"5.7"},
    1224: definition.ErrorDefinition{ErrorNumber:1224, ErrorType:"ServerError", Symbol:"ER_MIXING_NOT_ALLOWED", SQLState:"HY000", Message:"Mixing of transactional and non-transactional tables is disabled ", Description:"", MySQLVersion:"5.7"},
    1225: definition.ErrorDefinition{ErrorNumber:1225, ErrorType:"ServerError", Symbol:"ER_DUP_ARGUMENT", SQLState:"HY000", Message:"Option '%s' used twice in statement ", Description:"", MySQLVersion:"5.7"},
    1226: definition.ErrorDefinition{ErrorNumber:1226, ErrorType:"ServerError", Symbol:"ER_USER_LIMIT_REACHED", SQLState:"42000", Message:"User '%s' has exceeded the '%s' resource (current value: %ld) ", Description:"", MySQLVersion:"5.7"},
    1227: definition.ErrorDefinition{ErrorNumber:1227, ErrorType:"ServerError", Symbol:"ER_SPECIFIC_ACCESS_DENIED_ERROR", SQLState:"42000", Message:"Access denied", Description:"you need (at least one of) the %s privilege(s) for this operation ", MySQLVersion:"5.7"},
    1228: definition.ErrorDefinition{ErrorNumber:1228, ErrorType:"ServerError", Symbol:"ER_LOCAL_VARIABLE", SQLState:"HY000", Message:"Variable '%s' is a SESSION variable and can't be used with SET GLOBAL ", Description:"", MySQLVersion:"5.7"},
    1229: definition.ErrorDefinition{ErrorNumber:1229, ErrorType:"ServerError", Symbol:"ER_GLOBAL_VARIABLE", SQLState:"HY000", Message:"Variable '%s' is a GLOBAL variable and should be set with SET GLOBAL ", Description:"", MySQLVersion:"5.7"},
    1230: definition.ErrorDefinition{ErrorNumber:1230, ErrorType:"ServerError", Symbol:"ER_NO_DEFAULT", SQLState:"42000", Message:"Variable '%s' doesn't have a default value ", Description:"", MySQLVersion:"5.7"},
    1231: definition.ErrorDefinition{ErrorNumber:1231, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_FOR_VAR", SQLState:"42000", Message:"Variable '%s' can't be set to the value of '%s' ", Description:"", MySQLVersion:"5.7"},
    1232: definition.ErrorDefinition{ErrorNumber:1232, ErrorType:"ServerError", Symbol:"ER_WRONG_TYPE_FOR_VAR", SQLState:"42000", Message:"Incorrect argument type to variable '%s' ", Description:"", MySQLVersion:"5.7"},
    1233: definition.ErrorDefinition{ErrorNumber:1233, ErrorType:"ServerError", Symbol:"ER_VAR_CANT_BE_READ", SQLState:"HY000", Message:"Variable '%s' can only be set, not read ", Description:"", MySQLVersion:"5.7"},
    1234: definition.ErrorDefinition{ErrorNumber:1234, ErrorType:"ServerError", Symbol:"ER_CANT_USE_OPTION_HERE", SQLState:"42000", Message:"Incorrect usage/placement of '%s' ", Description:"", MySQLVersion:"5.7"},
    1235: definition.ErrorDefinition{ErrorNumber:1235, ErrorType:"ServerError", Symbol:"ER_NOT_SUPPORTED_YET", SQLState:"42000", Message:"This version of MySQL doesn't yet support '%s' ", Description:"", MySQLVersion:"5.7"},
    1236: definition.ErrorDefinition{ErrorNumber:1236, ErrorType:"ServerError", Symbol:"ER_MASTER_FATAL_ERROR_READING_BINLOG", SQLState:"HY000", Message:"Got fatal error %d from master when reading data from binary log: '%s' ", Description:"", MySQLVersion:"5.7"},
    1237: definition.ErrorDefinition{ErrorNumber:1237, ErrorType:"ServerError", Symbol:"ER_SLAVE_IGNORED_TABLE", SQLState:"HY000", Message:"Slave SQL thread ignored the query because of replicate-*-table rules ", Description:"", MySQLVersion:"5.7"},
    1238: definition.ErrorDefinition{ErrorNumber:1238, ErrorType:"ServerError", Symbol:"ER_INCORRECT_GLOBAL_LOCAL_VAR", SQLState:"HY000", Message:"Variable '%s' is a %s variable ", Description:"", MySQLVersion:"5.7"},
    1239: definition.ErrorDefinition{ErrorNumber:1239, ErrorType:"ServerError", Symbol:"ER_WRONG_FK_DEF", SQLState:"42000", Message:"Incorrect foreign key definition for '%s': %s ", Description:"", MySQLVersion:"5.7"},
    1240: definition.ErrorDefinition{ErrorNumber:1240, ErrorType:"ServerError", Symbol:"ER_KEY_REF_DO_NOT_MATCH_TABLE_REF", SQLState:"HY000", Message:"Key reference and table reference don't match ", Description:"", MySQLVersion:"5.7"},
    1241: definition.ErrorDefinition{ErrorNumber:1241, ErrorType:"ServerError", Symbol:"ER_OPERAND_COLUMNS", SQLState:"21000", Message:"Operand should contain %d column(s) ", Description:"", MySQLVersion:"5.7"},
    1242: definition.ErrorDefinition{ErrorNumber:1242, ErrorType:"ServerError", Symbol:"ER_SUBQUERY_NO_1_ROW", SQLState:"21000", Message:"Subquery returns more than 1 row ", Description:"", MySQLVersion:"5.7"},
    1243: definition.ErrorDefinition{ErrorNumber:1243, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_STMT_HANDLER", SQLState:"HY000", Message:"Unknown prepared statement handler (%.*s) given to %s ", Description:"", MySQLVersion:"5.7"},
    1244: definition.ErrorDefinition{ErrorNumber:1244, ErrorType:"ServerError", Symbol:"ER_CORRUPT_HELP_DB", SQLState:"HY000", Message:"Help database is corrupt or does not exist ", Description:"", MySQLVersion:"5.7"},
    1245: definition.ErrorDefinition{ErrorNumber:1245, ErrorType:"ServerError", Symbol:"ER_CYCLIC_REFERENCE", SQLState:"HY000", Message:"Cyclic reference on subqueries ", Description:"", MySQLVersion:"5.7"},
    1246: definition.ErrorDefinition{ErrorNumber:1246, ErrorType:"ServerError", Symbol:"ER_AUTO_CONVERT", SQLState:"HY000", Message:"Converting column '%s' from %s to %s ", Description:"", MySQLVersion:"5.7"},
    1247: definition.ErrorDefinition{ErrorNumber:1247, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_REFERENCE", SQLState:"42S22", Message:"Reference '%s' not supported (%s) ", Description:"", MySQLVersion:"5.7"},
    1248: definition.ErrorDefinition{ErrorNumber:1248, ErrorType:"ServerError", Symbol:"ER_DERIVED_MUST_HAVE_ALIAS", SQLState:"42000", Message:"Every derived table must have its own alias ", Description:"", MySQLVersion:"5.7"},
    1249: definition.ErrorDefinition{ErrorNumber:1249, ErrorType:"ServerError", Symbol:"ER_SELECT_REDUCED", SQLState:"01000", Message:"Select %u was reduced during optimization ", Description:"", MySQLVersion:"5.7"},
    1250: definition.ErrorDefinition{ErrorNumber:1250, ErrorType:"ServerError", Symbol:"ER_TABLENAME_NOT_ALLOWED_HERE", SQLState:"42000", Message:"Table '%s' from one of the SELECTs cannot be used in %s ", Description:"", MySQLVersion:"5.7"},
    1251: definition.ErrorDefinition{ErrorNumber:1251, ErrorType:"ServerError", Symbol:"ER_NOT_SUPPORTED_AUTH_MODE", SQLState:"08004", Message:"Client does not support authentication protocol requested by server", Description:"consider upgrading MySQL client ", MySQLVersion:"5.7"},
    1252: definition.ErrorDefinition{ErrorNumber:1252, ErrorType:"ServerError", Symbol:"ER_SPATIAL_CANT_HAVE_NULL", SQLState:"42000", Message:"All parts of a SPATIAL index must be NOT NULL ", Description:"", MySQLVersion:"5.7"},
    1253: definition.ErrorDefinition{ErrorNumber:1253, ErrorType:"ServerError", Symbol:"ER_COLLATION_CHARSET_MISMATCH", SQLState:"42000", Message:"COLLATION '%s' is not valid for CHARACTER SET '%s' ", Description:"", MySQLVersion:"5.7"},
    1254: definition.ErrorDefinition{ErrorNumber:1254, ErrorType:"ServerError", Symbol:"ER_SLAVE_WAS_RUNNING", SQLState:"HY000", Message:"Slave is already running ", Description:"", MySQLVersion:"5.7"},
    1255: definition.ErrorDefinition{ErrorNumber:1255, ErrorType:"ServerError", Symbol:"ER_SLAVE_WAS_NOT_RUNNING", SQLState:"HY000", Message:"Slave already has been stopped ", Description:"", MySQLVersion:"5.7"},
    1256: definition.ErrorDefinition{ErrorNumber:1256, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_FOR_UNCOMPRESS", SQLState:"HY000", Message:"Uncompressed data size too large", Description:"the maximum size is %d (probably, length of uncompressed data was corrupted) ", MySQLVersion:"5.7"},
    1257: definition.ErrorDefinition{ErrorNumber:1257, ErrorType:"ServerError", Symbol:"ER_ZLIB_Z_MEM_ERROR", SQLState:"HY000", Message:"ZLIB: Not enough memory ", Description:"", MySQLVersion:"5.7"},
    1258: definition.ErrorDefinition{ErrorNumber:1258, ErrorType:"ServerError", Symbol:"ER_ZLIB_Z_BUF_ERROR", SQLState:"HY000", Message:"ZLIB: Not enough room in the output buffer (probably, length of uncompressed data was corrupted) ", Description:"", MySQLVersion:"5.7"},
    1259: definition.ErrorDefinition{ErrorNumber:1259, ErrorType:"ServerError", Symbol:"ER_ZLIB_Z_DATA_ERROR", SQLState:"HY000", Message:"ZLIB: Input data corrupted ", Description:"", MySQLVersion:"5.7"},
    1260: definition.ErrorDefinition{ErrorNumber:1260, ErrorType:"ServerError", Symbol:"ER_CUT_VALUE_GROUP_CONCAT", SQLState:"HY000", Message:"Row %u was cut by GROUP_CONCAT() ", Description:"", MySQLVersion:"5.7"},
    1261: definition.ErrorDefinition{ErrorNumber:1261, ErrorType:"ServerError", Symbol:"ER_WARN_TOO_FEW_RECORDS", SQLState:"01000", Message:"Row %ld doesn't contain data for all columns ", Description:"", MySQLVersion:"5.7"},
    1262: definition.ErrorDefinition{ErrorNumber:1262, ErrorType:"ServerError", Symbol:"ER_WARN_TOO_MANY_RECORDS", SQLState:"01000", Message:"Row %ld was truncated", Description:"it contained more data than there were input columns ", MySQLVersion:"5.7"},
    1263: definition.ErrorDefinition{ErrorNumber:1263, ErrorType:"ServerError", Symbol:"ER_WARN_NULL_TO_NOTNULL", SQLState:"22004", Message:"Column set to default value", Description:"NULL supplied to NOT NULL column '%s' at row %ld ", MySQLVersion:"5.7"},
    1264: definition.ErrorDefinition{ErrorNumber:1264, ErrorType:"ServerError", Symbol:"ER_WARN_DATA_OUT_OF_RANGE", SQLState:"22003", Message:"Out of range value for column '%s' at row %ld ", Description:"", MySQLVersion:"5.7"},
    1265: definition.ErrorDefinition{ErrorNumber:1265, ErrorType:"ServerError", Symbol:"WARN_DATA_TRUNCATED", SQLState:"01000", Message:"Data truncated for column '%s' at row %ld ", Description:"", MySQLVersion:"5.7"},
    1266: definition.ErrorDefinition{ErrorNumber:1266, ErrorType:"ServerError", Symbol:"ER_WARN_USING_OTHER_HANDLER", SQLState:"HY000", Message:"Using storage engine %s for table '%s' ", Description:"", MySQLVersion:"5.7"},
    1267: definition.ErrorDefinition{ErrorNumber:1267, ErrorType:"ServerError", Symbol:"ER_CANT_AGGREGATE_2COLLATIONS", SQLState:"HY000", Message:"Illegal mix of collations (%s,%s) and (%s,%s) for operation '%s' ", Description:"", MySQLVersion:"5.7"},
    1268: definition.ErrorDefinition{ErrorNumber:1268, ErrorType:"ServerError", Symbol:"ER_DROP_USER", SQLState:"HY000", Message:"Cannot drop one or more of the requested users ", Description:"", MySQLVersion:"5.7"},
    1269: definition.ErrorDefinition{ErrorNumber:1269, ErrorType:"ServerError", Symbol:"ER_REVOKE_GRANTS", SQLState:"HY000", Message:"Can't revoke all privileges for one or more of the requested users ", Description:"", MySQLVersion:"5.7"},
    1270: definition.ErrorDefinition{ErrorNumber:1270, ErrorType:"ServerError", Symbol:"ER_CANT_AGGREGATE_3COLLATIONS", SQLState:"HY000", Message:"Illegal mix of collations (%s,%s), (%s,%s), (%s,%s) for operation '%s' ", Description:"", MySQLVersion:"5.7"},
    1271: definition.ErrorDefinition{ErrorNumber:1271, ErrorType:"ServerError", Symbol:"ER_CANT_AGGREGATE_NCOLLATIONS", SQLState:"HY000", Message:"Illegal mix of collations for operation '%s' ", Description:"", MySQLVersion:"5.7"},
    1272: definition.ErrorDefinition{ErrorNumber:1272, ErrorType:"ServerError", Symbol:"ER_VARIABLE_IS_NOT_STRUCT", SQLState:"HY000", Message:"Variable '%s' is not a variable component (can't be used as XXXX.variable_name) ", Description:"", MySQLVersion:"5.7"},
    1273: definition.ErrorDefinition{ErrorNumber:1273, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_COLLATION", SQLState:"HY000", Message:"Unknown collation: '%s' ", Description:"", MySQLVersion:"5.7"},
    1274: definition.ErrorDefinition{ErrorNumber:1274, ErrorType:"ServerError", Symbol:"ER_SLAVE_IGNORED_SSL_PARAMS", SQLState:"HY000", Message:"SSL parameters in CHANGE MASTER are ignored because this MySQL slave was compiled without SSL support", Description:"they can be used later if MySQL slave with SSL is started ", MySQLVersion:"5.7"},
    1275: definition.ErrorDefinition{ErrorNumber:1275, ErrorType:"ServerError", Symbol:"ER_SERVER_IS_IN_SECURE_AUTH_MODE", SQLState:"HY000", Message:"Server is running in --secure-auth mode, but '%s'@'%s' has a password in the old format", Description:"please change the password to the new format ", MySQLVersion:"5.7"},
    1276: definition.ErrorDefinition{ErrorNumber:1276, ErrorType:"ServerError", Symbol:"ER_WARN_FIELD_RESOLVED", SQLState:"HY000", Message:"Field or reference '%s%s%s%s%s' of SELECT #%d was resolved in SELECT #%d ", Description:"", MySQLVersion:"5.7"},
    1277: definition.ErrorDefinition{ErrorNumber:1277, ErrorType:"ServerError", Symbol:"ER_BAD_SLAVE_UNTIL_COND", SQLState:"HY000", Message:"Incorrect parameter or combination of parameters for START SLAVE UNTIL ", Description:"", MySQLVersion:"5.7"},
    1278: definition.ErrorDefinition{ErrorNumber:1278, ErrorType:"ServerError", Symbol:"ER_MISSING_SKIP_SLAVE", SQLState:"HY000", Message:"It is recommended to use --skip-slave-start when doing step-by-step replication with START SLAVE UNTIL", Description:"otherwise, you will get problems if you get an unexpected slave's mysqld restart ", MySQLVersion:"5.7"},
    1279: definition.ErrorDefinition{ErrorNumber:1279, ErrorType:"ServerError", Symbol:"ER_UNTIL_COND_IGNORED", SQLState:"HY000", Message:"SQL thread is not to be started so UNTIL options are ignored ", Description:"", MySQLVersion:"5.7"},
    1280: definition.ErrorDefinition{ErrorNumber:1280, ErrorType:"ServerError", Symbol:"ER_WRONG_NAME_FOR_INDEX", SQLState:"42000", Message:"Incorrect index name '%s' ", Description:"", MySQLVersion:"5.7"},
    1281: definition.ErrorDefinition{ErrorNumber:1281, ErrorType:"ServerError", Symbol:"ER_WRONG_NAME_FOR_CATALOG", SQLState:"42000", Message:"Incorrect catalog name '%s' ", Description:"", MySQLVersion:"5.7"},
    1282: definition.ErrorDefinition{ErrorNumber:1282, ErrorType:"ServerError", Symbol:"ER_WARN_QC_RESIZE", SQLState:"HY000", Message:"Query cache failed to set size %lu", Description:"new query cache size is %lu ", MySQLVersion:"5.7"},
    1283: definition.ErrorDefinition{ErrorNumber:1283, ErrorType:"ServerError", Symbol:"ER_BAD_FT_COLUMN", SQLState:"HY000", Message:"Column '%s' cannot be part of FULLTEXT index ", Description:"", MySQLVersion:"5.7"},
    1284: definition.ErrorDefinition{ErrorNumber:1284, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_KEY_CACHE", SQLState:"HY000", Message:"Unknown key cache '%s' ", Description:"", MySQLVersion:"5.7"},
    1285: definition.ErrorDefinition{ErrorNumber:1285, ErrorType:"ServerError", Symbol:"ER_WARN_HOSTNAME_WONT_WORK", SQLState:"HY000", Message:"MySQL is started in --skip-name-resolve mode", Description:"you must restart it without this switch for this grant to work ", MySQLVersion:"5.7"},
    1286: definition.ErrorDefinition{ErrorNumber:1286, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_STORAGE_ENGINE", SQLState:"42000", Message:"Unknown storage engine '%s' ", Description:"", MySQLVersion:"5.7"},
    1287: definition.ErrorDefinition{ErrorNumber:1287, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYNTAX", SQLState:"HY000", Message:"'%s' is deprecated and will be removed in a future release. Please use %s instead ", Description:"", MySQLVersion:"5.7"},
    1288: definition.ErrorDefinition{ErrorNumber:1288, ErrorType:"ServerError", Symbol:"ER_NON_UPDATABLE_TABLE", SQLState:"HY000", Message:"The target table %s of the %s is not updatable ", Description:"", MySQLVersion:"5.7"},
    1289: definition.ErrorDefinition{ErrorNumber:1289, ErrorType:"ServerError", Symbol:"ER_FEATURE_DISABLED", SQLState:"HY000", Message:"The '%s' feature is disabled", Description:"you need MySQL built with '%s' to have it working ", MySQLVersion:"5.7"},
    1290: definition.ErrorDefinition{ErrorNumber:1290, ErrorType:"ServerError", Symbol:"ER_OPTION_PREVENTS_STATEMENT", SQLState:"HY000", Message:"The MySQL server is running with the %s option so it cannot execute this statement ", Description:"", MySQLVersion:"5.7"},
    1291: definition.ErrorDefinition{ErrorNumber:1291, ErrorType:"ServerError", Symbol:"ER_DUPLICATED_VALUE_IN_TYPE", SQLState:"HY000", Message:"Column '%s' has duplicated value '%s' in %s ", Description:"", MySQLVersion:"5.7"},
    1292: definition.ErrorDefinition{ErrorNumber:1292, ErrorType:"ServerError", Symbol:"ER_TRUNCATED_WRONG_VALUE", SQLState:"22007", Message:"Truncated incorrect %s value: '%s' ", Description:"", MySQLVersion:"5.7"},
    1293: definition.ErrorDefinition{ErrorNumber:1293, ErrorType:"ServerError", Symbol:"ER_TOO_MUCH_AUTO_TIMESTAMP_COLS", SQLState:"HY000", Message:"Incorrect table definition", Description:"there can be only one TIMESTAMP column with CURRENT_TIMESTAMP in DEFAULT or ON UPDATE clause ", MySQLVersion:"5.7"},
    1294: definition.ErrorDefinition{ErrorNumber:1294, ErrorType:"ServerError", Symbol:"ER_INVALID_ON_UPDATE", SQLState:"HY000", Message:"Invalid ON UPDATE clause for '%s' column ", Description:"", MySQLVersion:"5.7"},
    1295: definition.ErrorDefinition{ErrorNumber:1295, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_PS", SQLState:"HY000", Message:"This command is not supported in the prepared statement protocol yet ", Description:"", MySQLVersion:"5.7"},
    1296: definition.ErrorDefinition{ErrorNumber:1296, ErrorType:"ServerError", Symbol:"ER_GET_ERRMSG", SQLState:"HY000", Message:"Got error %d '%s' from %s ", Description:"", MySQLVersion:"5.7"},
    1297: definition.ErrorDefinition{ErrorNumber:1297, ErrorType:"ServerError", Symbol:"ER_GET_TEMPORARY_ERRMSG", SQLState:"HY000", Message:"Got temporary error %d '%s' from %s ", Description:"", MySQLVersion:"5.7"},
    1298: definition.ErrorDefinition{ErrorNumber:1298, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_TIME_ZONE", SQLState:"HY000", Message:"Unknown or incorrect time zone: '%s' ", Description:"", MySQLVersion:"5.7"},
    1299: definition.ErrorDefinition{ErrorNumber:1299, ErrorType:"ServerError", Symbol:"ER_WARN_INVALID_TIMESTAMP", SQLState:"HY000", Message:"Invalid TIMESTAMP value in column '%s' at row %ld ", Description:"", MySQLVersion:"5.7"},
    1300: definition.ErrorDefinition{ErrorNumber:1300, ErrorType:"ServerError", Symbol:"ER_INVALID_CHARACTER_STRING", SQLState:"HY000", Message:"Invalid %s character string: '%s' ", Description:"", MySQLVersion:"5.7"},
    1301: definition.ErrorDefinition{ErrorNumber:1301, ErrorType:"ServerError", Symbol:"ER_WARN_ALLOWED_PACKET_OVERFLOWED", SQLState:"HY000", Message:"Result of %s() was larger than max_allowed_packet (%ld) - truncated ", Description:"", MySQLVersion:"5.7"},
    1302: definition.ErrorDefinition{ErrorNumber:1302, ErrorType:"ServerError", Symbol:"ER_CONFLICTING_DECLARATIONS", SQLState:"HY000", Message:"Conflicting declarations: '%s%s' and '%s%s' ", Description:"", MySQLVersion:"5.7"},
    1303: definition.ErrorDefinition{ErrorNumber:1303, ErrorType:"ServerError", Symbol:"ER_SP_NO_RECURSIVE_CREATE", SQLState:"2F003", Message:"Can't create a %s from within another stored routine ", Description:"", MySQLVersion:"5.7"},
    1304: definition.ErrorDefinition{ErrorNumber:1304, ErrorType:"ServerError", Symbol:"ER_SP_ALREADY_EXISTS", SQLState:"42000", Message:"%s %s already exists ", Description:"", MySQLVersion:"5.7"},
    1305: definition.ErrorDefinition{ErrorNumber:1305, ErrorType:"ServerError", Symbol:"ER_SP_DOES_NOT_EXIST", SQLState:"42000", Message:"%s %s does not exist ", Description:"", MySQLVersion:"5.7"},
    1306: definition.ErrorDefinition{ErrorNumber:1306, ErrorType:"ServerError", Symbol:"ER_SP_DROP_FAILED", SQLState:"HY000", Message:"Failed to DROP %s %s ", Description:"", MySQLVersion:"5.7"},
    1307: definition.ErrorDefinition{ErrorNumber:1307, ErrorType:"ServerError", Symbol:"ER_SP_STORE_FAILED", SQLState:"HY000", Message:"Failed to CREATE %s %s ", Description:"", MySQLVersion:"5.7"},
    1308: definition.ErrorDefinition{ErrorNumber:1308, ErrorType:"ServerError", Symbol:"ER_SP_LILABEL_MISMATCH", SQLState:"42000", Message:"%s with no matching label: %s ", Description:"", MySQLVersion:"5.7"},
    1309: definition.ErrorDefinition{ErrorNumber:1309, ErrorType:"ServerError", Symbol:"ER_SP_LABEL_REDEFINE", SQLState:"42000", Message:"Redefining label %s ", Description:"", MySQLVersion:"5.7"},
    1310: definition.ErrorDefinition{ErrorNumber:1310, ErrorType:"ServerError", Symbol:"ER_SP_LABEL_MISMATCH", SQLState:"42000", Message:"End-label %s without match ", Description:"", MySQLVersion:"5.7"},
    1311: definition.ErrorDefinition{ErrorNumber:1311, ErrorType:"ServerError", Symbol:"ER_SP_UNINIT_VAR", SQLState:"01000", Message:"Referring to uninitialized variable %s ", Description:"", MySQLVersion:"5.7"},
    1312: definition.ErrorDefinition{ErrorNumber:1312, ErrorType:"ServerError", Symbol:"ER_SP_BADSELECT", SQLState:"0A000", Message:"PROCEDURE %s can't return a result set in the given context ", Description:"", MySQLVersion:"5.7"},
    1313: definition.ErrorDefinition{ErrorNumber:1313, ErrorType:"ServerError", Symbol:"ER_SP_BADRETURN", SQLState:"42000", Message:"RETURN is only allowed in a FUNCTION ", Description:"", MySQLVersion:"5.7"},
    1314: definition.ErrorDefinition{ErrorNumber:1314, ErrorType:"ServerError", Symbol:"ER_SP_BADSTATEMENT", SQLState:"0A000", Message:"%s is not allowed in stored procedures ", Description:"", MySQLVersion:"5.7"},
    1315: definition.ErrorDefinition{ErrorNumber:1315, ErrorType:"ServerError", Symbol:"ER_UPDATE_LOG_DEPRECATED_IGNORED", SQLState:"42000", Message:"The update log is deprecated and replaced by the binary log", Description:"SET SQL_LOG_UPDATE has been ignored. ", MySQLVersion:"5.7"},
    1316: definition.ErrorDefinition{ErrorNumber:1316, ErrorType:"ServerError", Symbol:"ER_UPDATE_LOG_DEPRECATED_TRANSLATED", SQLState:"42000", Message:"The update log is deprecated and replaced by the binary log", Description:"SET SQL_LOG_UPDATE has been translated to SET SQL_LOG_BIN. ", MySQLVersion:"5.7"},
    1317: definition.ErrorDefinition{ErrorNumber:1317, ErrorType:"ServerError", Symbol:"ER_QUERY_INTERRUPTED", SQLState:"70100", Message:"Query execution was interrupted ", Description:"", MySQLVersion:"5.7"},
    1318: definition.ErrorDefinition{ErrorNumber:1318, ErrorType:"ServerError", Symbol:"ER_SP_WRONG_NO_OF_ARGS", SQLState:"42000", Message:"Incorrect number of arguments for %s %s", Description:"expected %u, got %u ", MySQLVersion:"5.7"},
    1319: definition.ErrorDefinition{ErrorNumber:1319, ErrorType:"ServerError", Symbol:"ER_SP_COND_MISMATCH", SQLState:"42000", Message:"Undefined CONDITION: %s ", Description:"", MySQLVersion:"5.7"},
    1320: definition.ErrorDefinition{ErrorNumber:1320, ErrorType:"ServerError", Symbol:"ER_SP_NORETURN", SQLState:"42000", Message:"No RETURN found in FUNCTION %s ", Description:"", MySQLVersion:"5.7"},
    1321: definition.ErrorDefinition{ErrorNumber:1321, ErrorType:"ServerError", Symbol:"ER_SP_NORETURNEND", SQLState:"2F005", Message:"FUNCTION %s ended without RETURN ", Description:"", MySQLVersion:"5.7"},
    1322: definition.ErrorDefinition{ErrorNumber:1322, ErrorType:"ServerError", Symbol:"ER_SP_BAD_CURSOR_QUERY", SQLState:"42000", Message:"Cursor statement must be a SELECT ", Description:"", MySQLVersion:"5.7"},
    1323: definition.ErrorDefinition{ErrorNumber:1323, ErrorType:"ServerError", Symbol:"ER_SP_BAD_CURSOR_SELECT", SQLState:"42000", Message:"Cursor SELECT must not have INTO ", Description:"", MySQLVersion:"5.7"},
    1324: definition.ErrorDefinition{ErrorNumber:1324, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_MISMATCH", SQLState:"42000", Message:"Undefined CURSOR: %s ", Description:"", MySQLVersion:"5.7"},
    1325: definition.ErrorDefinition{ErrorNumber:1325, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_ALREADY_OPEN", SQLState:"24000", Message:"Cursor is already open ", Description:"", MySQLVersion:"5.7"},
    1326: definition.ErrorDefinition{ErrorNumber:1326, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_NOT_OPEN", SQLState:"24000", Message:"Cursor is not open ", Description:"", MySQLVersion:"5.7"},
    1327: definition.ErrorDefinition{ErrorNumber:1327, ErrorType:"ServerError", Symbol:"ER_SP_UNDECLARED_VAR", SQLState:"42000", Message:"Undeclared variable: %s ", Description:"", MySQLVersion:"5.7"},
    1328: definition.ErrorDefinition{ErrorNumber:1328, ErrorType:"ServerError", Symbol:"ER_SP_WRONG_NO_OF_FETCH_ARGS", SQLState:"HY000", Message:"Incorrect number of FETCH variables ", Description:"", MySQLVersion:"5.7"},
    1329: definition.ErrorDefinition{ErrorNumber:1329, ErrorType:"ServerError", Symbol:"ER_SP_FETCH_NO_DATA", SQLState:"02000", Message:"No data - zero rows fetched, selected, or processed ", Description:"", MySQLVersion:"5.7"},
    1330: definition.ErrorDefinition{ErrorNumber:1330, ErrorType:"ServerError", Symbol:"ER_SP_DUP_PARAM", SQLState:"42000", Message:"Duplicate parameter: %s ", Description:"", MySQLVersion:"5.7"},
    1331: definition.ErrorDefinition{ErrorNumber:1331, ErrorType:"ServerError", Symbol:"ER_SP_DUP_VAR", SQLState:"42000", Message:"Duplicate variable: %s ", Description:"", MySQLVersion:"5.7"},
    1332: definition.ErrorDefinition{ErrorNumber:1332, ErrorType:"ServerError", Symbol:"ER_SP_DUP_COND", SQLState:"42000", Message:"Duplicate condition: %s ", Description:"", MySQLVersion:"5.7"},
    1333: definition.ErrorDefinition{ErrorNumber:1333, ErrorType:"ServerError", Symbol:"ER_SP_DUP_CURS", SQLState:"42000", Message:"Duplicate cursor: %s ", Description:"", MySQLVersion:"5.7"},
    1334: definition.ErrorDefinition{ErrorNumber:1334, ErrorType:"ServerError", Symbol:"ER_SP_CANT_ALTER", SQLState:"HY000", Message:"Failed to ALTER %s %s ", Description:"", MySQLVersion:"5.7"},
    1335: definition.ErrorDefinition{ErrorNumber:1335, ErrorType:"ServerError", Symbol:"ER_SP_SUBSELECT_NYI", SQLState:"0A000", Message:"Subquery value not supported ", Description:"", MySQLVersion:"5.7"},
    1336: definition.ErrorDefinition{ErrorNumber:1336, ErrorType:"ServerError", Symbol:"ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG", SQLState:"0A000", Message:"%s is not allowed in stored function or trigger ", Description:"", MySQLVersion:"5.7"},
    1337: definition.ErrorDefinition{ErrorNumber:1337, ErrorType:"ServerError", Symbol:"ER_SP_VARCOND_AFTER_CURSHNDLR", SQLState:"42000", Message:"Variable or condition declaration after cursor or handler declaration ", Description:"", MySQLVersion:"5.7"},
    1338: definition.ErrorDefinition{ErrorNumber:1338, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_AFTER_HANDLER", SQLState:"42000", Message:"Cursor declaration after handler declaration ", Description:"", MySQLVersion:"5.7"},
    1339: definition.ErrorDefinition{ErrorNumber:1339, ErrorType:"ServerError", Symbol:"ER_SP_CASE_NOT_FOUND", SQLState:"20000", Message:"Case not found for CASE statement ", Description:"", MySQLVersion:"5.7"},
    1340: definition.ErrorDefinition{ErrorNumber:1340, ErrorType:"ServerError", Symbol:"ER_FPARSER_TOO_BIG_FILE", SQLState:"HY000", Message:"Configuration file '%s' is too big ", Description:"", MySQLVersion:"5.7"},
    1341: definition.ErrorDefinition{ErrorNumber:1341, ErrorType:"ServerError", Symbol:"ER_FPARSER_BAD_HEADER", SQLState:"HY000", Message:"Malformed file type header in file '%s' ", Description:"", MySQLVersion:"5.7"},
    1342: definition.ErrorDefinition{ErrorNumber:1342, ErrorType:"ServerError", Symbol:"ER_FPARSER_EOF_IN_COMMENT", SQLState:"HY000", Message:"Unexpected end of file while parsing comment '%s' ", Description:"", MySQLVersion:"5.7"},
    1343: definition.ErrorDefinition{ErrorNumber:1343, ErrorType:"ServerError", Symbol:"ER_FPARSER_ERROR_IN_PARAMETER", SQLState:"HY000", Message:"Error while parsing parameter '%s' (line: '%s') ", Description:"", MySQLVersion:"5.7"},
    1344: definition.ErrorDefinition{ErrorNumber:1344, ErrorType:"ServerError", Symbol:"ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER", SQLState:"HY000", Message:"Unexpected end of file while skipping unknown parameter '%s' ", Description:"", MySQLVersion:"5.7"},
    1345: definition.ErrorDefinition{ErrorNumber:1345, ErrorType:"ServerError", Symbol:"ER_VIEW_NO_EXPLAIN", SQLState:"HY000", Message:"EXPLAIN/SHOW can not be issued", Description:"lacking privileges for underlying table ", MySQLVersion:"5.7"},
    1346: definition.ErrorDefinition{ErrorNumber:1346, ErrorType:"ServerError", Symbol:"ER_FRM_UNKNOWN_TYPE", SQLState:"HY000", Message:"File '%s' has unknown type '%s' in its header ", Description:"", MySQLVersion:"5.7"},
    1347: definition.ErrorDefinition{ErrorNumber:1347, ErrorType:"ServerError", Symbol:"ER_WRONG_OBJECT", SQLState:"HY000", Message:"'%s.%s' is not %s", Description:"The named object is incorrect for the type of operation attempted on it. It must be an object of the named type. ", MySQLVersion:"5.7"},
    1348: definition.ErrorDefinition{ErrorNumber:1348, ErrorType:"ServerError", Symbol:"ER_NONUPDATEABLE_COLUMN", SQLState:"HY000", Message:"Column '%s' is not updatable ", Description:"", MySQLVersion:"5.7"},
    1349: definition.ErrorDefinition{ErrorNumber:1349, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_DERIVED", SQLState:"HY000", Message:"View's SELECT contains a subquery in the FROM clause", Description:"ER_VIEW_SELECT_DERIVED was removed after 5.7.6. ", MySQLVersion:"5.7"},
    1349: definition.ErrorDefinition{ErrorNumber:1349, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_DERIVED_UNUSED", SQLState:"HY000", Message:"View's SELECT contains a subquery in the FROM clause", Description:"ER_VIEW_SELECT_DERIVED_UNUSED was added in 5.7.7. ", MySQLVersion:"5.7"},
    1350: definition.ErrorDefinition{ErrorNumber:1350, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_CLAUSE", SQLState:"HY000", Message:"View's SELECT contains a '%s' clause ", Description:"", MySQLVersion:"5.7"},
    1351: definition.ErrorDefinition{ErrorNumber:1351, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_VARIABLE", SQLState:"HY000", Message:"View's SELECT contains a variable or parameter ", Description:"", MySQLVersion:"5.7"},
    1352: definition.ErrorDefinition{ErrorNumber:1352, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_TMPTABLE", SQLState:"HY000", Message:"View's SELECT refers to a temporary table '%s' ", Description:"", MySQLVersion:"5.7"},
    1353: definition.ErrorDefinition{ErrorNumber:1353, ErrorType:"ServerError", Symbol:"ER_VIEW_WRONG_LIST", SQLState:"HY000", Message:"View's SELECT and view's field list have different column counts ", Description:"", MySQLVersion:"5.7"},
    1354: definition.ErrorDefinition{ErrorNumber:1354, ErrorType:"ServerError", Symbol:"ER_WARN_VIEW_MERGE", SQLState:"HY000", Message:"View merge algorithm can't be used here for now (assumed undefined algorithm) ", Description:"", MySQLVersion:"5.7"},
    1355: definition.ErrorDefinition{ErrorNumber:1355, ErrorType:"ServerError", Symbol:"ER_WARN_VIEW_WITHOUT_KEY", SQLState:"HY000", Message:"View being updated does not have complete key of underlying table in it ", Description:"", MySQLVersion:"5.7"},
    1356: definition.ErrorDefinition{ErrorNumber:1356, ErrorType:"ServerError", Symbol:"ER_VIEW_INVALID", SQLState:"HY000", Message:"View '%s.%s' references invalid table(s) or column(s) or function(s) or definer/invoker of view lack rights to use them ", Description:"", MySQLVersion:"5.7"},
    1357: definition.ErrorDefinition{ErrorNumber:1357, ErrorType:"ServerError", Symbol:"ER_SP_NO_DROP_SP", SQLState:"HY000", Message:"Can't drop or alter a %s from within another stored routine ", Description:"", MySQLVersion:"5.7"},
    1358: definition.ErrorDefinition{ErrorNumber:1358, ErrorType:"ServerError", Symbol:"ER_SP_GOTO_IN_HNDLR", SQLState:"HY000", Message:"GOTO is not allowed in a stored procedure handler ", Description:"", MySQLVersion:"5.7"},
    1359: definition.ErrorDefinition{ErrorNumber:1359, ErrorType:"ServerError", Symbol:"ER_TRG_ALREADY_EXISTS", SQLState:"HY000", Message:"Trigger already exists ", Description:"", MySQLVersion:"5.7"},
    1360: definition.ErrorDefinition{ErrorNumber:1360, ErrorType:"ServerError", Symbol:"ER_TRG_DOES_NOT_EXIST", SQLState:"HY000", Message:"Trigger does not exist ", Description:"", MySQLVersion:"5.7"},
    1361: definition.ErrorDefinition{ErrorNumber:1361, ErrorType:"ServerError", Symbol:"ER_TRG_ON_VIEW_OR_TEMP_TABLE", SQLState:"HY000", Message:"Trigger's '%s' is view or temporary table ", Description:"", MySQLVersion:"5.7"},
    1362: definition.ErrorDefinition{ErrorNumber:1362, ErrorType:"ServerError", Symbol:"ER_TRG_CANT_CHANGE_ROW", SQLState:"HY000", Message:"Updating of %s row is not allowed in %strigger ", Description:"", MySQLVersion:"5.7"},
    1363: definition.ErrorDefinition{ErrorNumber:1363, ErrorType:"ServerError", Symbol:"ER_TRG_NO_SUCH_ROW_IN_TRG", SQLState:"HY000", Message:"There is no %s row in %s trigger ", Description:"", MySQLVersion:"5.7"},
    1364: definition.ErrorDefinition{ErrorNumber:1364, ErrorType:"ServerError", Symbol:"ER_NO_DEFAULT_FOR_FIELD", SQLState:"HY000", Message:"Field '%s' doesn't have a default value ", Description:"", MySQLVersion:"5.7"},
    1365: definition.ErrorDefinition{ErrorNumber:1365, ErrorType:"ServerError", Symbol:"ER_DIVISION_BY_ZERO", SQLState:"22012", Message:"Division by 0 ", Description:"", MySQLVersion:"5.7"},
    1366: definition.ErrorDefinition{ErrorNumber:1366, ErrorType:"ServerError", Symbol:"ER_TRUNCATED_WRONG_VALUE_FOR_FIELD", SQLState:"HY000", Message:"Incorrect %s value: '%s' for column '%s' at row %ld ", Description:"", MySQLVersion:"5.7"},
    1367: definition.ErrorDefinition{ErrorNumber:1367, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_VALUE_FOR_TYPE", SQLState:"22007", Message:"Illegal %s '%s' value found during parsing ", Description:"", MySQLVersion:"5.7"},
    1368: definition.ErrorDefinition{ErrorNumber:1368, ErrorType:"ServerError", Symbol:"ER_VIEW_NONUPD_CHECK", SQLState:"HY000", Message:"CHECK OPTION on non-updatable view '%s.%s' ", Description:"", MySQLVersion:"5.7"},
    1369: definition.ErrorDefinition{ErrorNumber:1369, ErrorType:"ServerError", Symbol:"ER_VIEW_CHECK_FAILED", SQLState:"HY000", Message:"CHECK OPTION failed '%s.%s' ", Description:"", MySQLVersion:"5.7"},
    1370: definition.ErrorDefinition{ErrorNumber:1370, ErrorType:"ServerError", Symbol:"ER_PROCACCESS_DENIED_ERROR", SQLState:"42000", Message:"%s command denied to user '%s'@'%s' for routine '%s' ", Description:"", MySQLVersion:"5.7"},
    1371: definition.ErrorDefinition{ErrorNumber:1371, ErrorType:"ServerError", Symbol:"ER_RELAY_LOG_FAIL", SQLState:"HY000", Message:"Failed purging old relay logs: %s ", Description:"", MySQLVersion:"5.7"},
    1372: definition.ErrorDefinition{ErrorNumber:1372, ErrorType:"ServerError", Symbol:"ER_PASSWD_LENGTH", SQLState:"HY000", Message:"Password hash should be a %d-digit hexadecimal number ", Description:"", MySQLVersion:"5.7"},
    1373: definition.ErrorDefinition{ErrorNumber:1373, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_TARGET_BINLOG", SQLState:"HY000", Message:"Target log not found in binlog index ", Description:"", MySQLVersion:"5.7"},
    1374: definition.ErrorDefinition{ErrorNumber:1374, ErrorType:"ServerError", Symbol:"ER_IO_ERR_LOG_INDEX_READ", SQLState:"HY000", Message:"I/O error reading log index file ", Description:"", MySQLVersion:"5.7"},
    1375: definition.ErrorDefinition{ErrorNumber:1375, ErrorType:"ServerError", Symbol:"ER_BINLOG_PURGE_PROHIBITED", SQLState:"HY000", Message:"Server configuration does not permit binlog purge ", Description:"", MySQLVersion:"5.7"},
    1376: definition.ErrorDefinition{ErrorNumber:1376, ErrorType:"ServerError", Symbol:"ER_FSEEK_FAIL", SQLState:"HY000", Message:"Failed on fseek() ", Description:"", MySQLVersion:"5.7"},
    1377: definition.ErrorDefinition{ErrorNumber:1377, ErrorType:"ServerError", Symbol:"ER_BINLOG_PURGE_FATAL_ERR", SQLState:"HY000", Message:"Fatal error during log purge ", Description:"", MySQLVersion:"5.7"},
    1378: definition.ErrorDefinition{ErrorNumber:1378, ErrorType:"ServerError", Symbol:"ER_LOG_IN_USE", SQLState:"HY000", Message:"A purgeable log is in use, will not purge ", Description:"", MySQLVersion:"5.7"},
    1379: definition.ErrorDefinition{ErrorNumber:1379, ErrorType:"ServerError", Symbol:"ER_LOG_PURGE_UNKNOWN_ERR", SQLState:"HY000", Message:"Unknown error during log purge ", Description:"", MySQLVersion:"5.7"},
    1380: definition.ErrorDefinition{ErrorNumber:1380, ErrorType:"ServerError", Symbol:"ER_RELAY_LOG_INIT", SQLState:"HY000", Message:"Failed initializing relay log position: %s ", Description:"", MySQLVersion:"5.7"},
    1381: definition.ErrorDefinition{ErrorNumber:1381, ErrorType:"ServerError", Symbol:"ER_NO_BINARY_LOGGING", SQLState:"HY000", Message:"You are not using binary logging ", Description:"", MySQLVersion:"5.7"},
    1382: definition.ErrorDefinition{ErrorNumber:1382, ErrorType:"ServerError", Symbol:"ER_RESERVED_SYNTAX", SQLState:"HY000", Message:"The '%s' syntax is reserved for purposes internal to the MySQL server ", Description:"", MySQLVersion:"5.7"},
    1383: definition.ErrorDefinition{ErrorNumber:1383, ErrorType:"ServerError", Symbol:"ER_WSAS_FAILED", SQLState:"HY000", Message:"WSAStartup Failed ", Description:"", MySQLVersion:"5.7"},
    1384: definition.ErrorDefinition{ErrorNumber:1384, ErrorType:"ServerError", Symbol:"ER_DIFF_GROUPS_PROC", SQLState:"HY000", Message:"Can't handle procedures with different groups yet ", Description:"", MySQLVersion:"5.7"},
    1385: definition.ErrorDefinition{ErrorNumber:1385, ErrorType:"ServerError", Symbol:"ER_NO_GROUP_FOR_PROC", SQLState:"HY000", Message:"Select must have a group with this procedure ", Description:"", MySQLVersion:"5.7"},
    1386: definition.ErrorDefinition{ErrorNumber:1386, ErrorType:"ServerError", Symbol:"ER_ORDER_WITH_PROC", SQLState:"HY000", Message:"Can't use ORDER clause with this procedure ", Description:"", MySQLVersion:"5.7"},
    1387: definition.ErrorDefinition{ErrorNumber:1387, ErrorType:"ServerError", Symbol:"ER_LOGGING_PROHIBIT_CHANGING_OF", SQLState:"HY000", Message:"Binary logging and replication forbid changing the global server %s ", Description:"", MySQLVersion:"5.7"},
    1388: definition.ErrorDefinition{ErrorNumber:1388, ErrorType:"ServerError", Symbol:"ER_NO_FILE_MAPPING", SQLState:"HY000", Message:"Can't map file: %s, errno: %d ", Description:"", MySQLVersion:"5.7"},
    1389: definition.ErrorDefinition{ErrorNumber:1389, ErrorType:"ServerError", Symbol:"ER_WRONG_MAGIC", SQLState:"HY000", Message:"Wrong magic in %s ", Description:"", MySQLVersion:"5.7"},
    1390: definition.ErrorDefinition{ErrorNumber:1390, ErrorType:"ServerError", Symbol:"ER_PS_MANY_PARAM", SQLState:"HY000", Message:"Prepared statement contains too many placeholders ", Description:"", MySQLVersion:"5.7"},
    1391: definition.ErrorDefinition{ErrorNumber:1391, ErrorType:"ServerError", Symbol:"ER_KEY_PART_0", SQLState:"HY000", Message:"Key part '%s' length cannot be 0 ", Description:"", MySQLVersion:"5.7"},
    1392: definition.ErrorDefinition{ErrorNumber:1392, ErrorType:"ServerError", Symbol:"ER_VIEW_CHECKSUM", SQLState:"HY000", Message:"View text checksum failed ", Description:"", MySQLVersion:"5.7"},
    1393: definition.ErrorDefinition{ErrorNumber:1393, ErrorType:"ServerError", Symbol:"ER_VIEW_MULTIUPDATE", SQLState:"HY000", Message:"Can not modify more than one base table through a join view '%s.%s' ", Description:"", MySQLVersion:"5.7"},
    1394: definition.ErrorDefinition{ErrorNumber:1394, ErrorType:"ServerError", Symbol:"ER_VIEW_NO_INSERT_FIELD_LIST", SQLState:"HY000", Message:"Can not insert into join view '%s.%s' without fields list ", Description:"", MySQLVersion:"5.7"},
    1395: definition.ErrorDefinition{ErrorNumber:1395, ErrorType:"ServerError", Symbol:"ER_VIEW_DELETE_MERGE_VIEW", SQLState:"HY000", Message:"Can not delete from join view '%s.%s' ", Description:"", MySQLVersion:"5.7"},
    1396: definition.ErrorDefinition{ErrorNumber:1396, ErrorType:"ServerError", Symbol:"ER_CANNOT_USER", SQLState:"HY000", Message:"Operation %s failed for %s ", Description:"", MySQLVersion:"5.7"},
    1397: definition.ErrorDefinition{ErrorNumber:1397, ErrorType:"ServerError", Symbol:"ER_XAER_NOTA", SQLState:"XAE04", Message:"XAER_NOTA: Unknown XID ", Description:"", MySQLVersion:"5.7"},
    1398: definition.ErrorDefinition{ErrorNumber:1398, ErrorType:"ServerError", Symbol:"ER_XAER_INVAL", SQLState:"XAE05", Message:"XAER_INVAL: Invalid arguments (or unsupported command) ", Description:"", MySQLVersion:"5.7"},
    1399: definition.ErrorDefinition{ErrorNumber:1399, ErrorType:"ServerError", Symbol:"ER_XAER_RMFAIL", SQLState:"XAE07", Message:"XAER_RMFAIL: The command cannot be executed when global transaction is in the %s state ", Description:"", MySQLVersion:"5.7"},
    1400: definition.ErrorDefinition{ErrorNumber:1400, ErrorType:"ServerError", Symbol:"ER_XAER_OUTSIDE", SQLState:"XAE09", Message:"XAER_OUTSIDE: Some work is done outside global transaction ", Description:"", MySQLVersion:"5.7"},
    1401: definition.ErrorDefinition{ErrorNumber:1401, ErrorType:"ServerError", Symbol:"ER_XAER_RMERR", SQLState:"XAE03", Message:"XAER_RMERR: Fatal error occurred in the transaction branch - check your data for consistency ", Description:"", MySQLVersion:"5.7"},
    1402: definition.ErrorDefinition{ErrorNumber:1402, ErrorType:"ServerError", Symbol:"ER_XA_RBROLLBACK", SQLState:"XA100", Message:"XA_RBROLLBACK: Transaction branch was rolled back ", Description:"", MySQLVersion:"5.7"},
    1403: definition.ErrorDefinition{ErrorNumber:1403, ErrorType:"ServerError", Symbol:"ER_NONEXISTING_PROC_GRANT", SQLState:"42000", Message:"There is no such grant defined for user '%s' on host '%s' on routine '%s' ", Description:"", MySQLVersion:"5.7"},
    1404: definition.ErrorDefinition{ErrorNumber:1404, ErrorType:"ServerError", Symbol:"ER_PROC_AUTO_GRANT_FAIL", SQLState:"HY000", Message:"Failed to grant EXECUTE and ALTER ROUTINE privileges ", Description:"", MySQLVersion:"5.7"},
    1405: definition.ErrorDefinition{ErrorNumber:1405, ErrorType:"ServerError", Symbol:"ER_PROC_AUTO_REVOKE_FAIL", SQLState:"HY000", Message:"Failed to revoke all privileges to dropped routine ", Description:"", MySQLVersion:"5.7"},
    1406: definition.ErrorDefinition{ErrorNumber:1406, ErrorType:"ServerError", Symbol:"ER_DATA_TOO_LONG", SQLState:"22001", Message:"Data too long for column '%s' at row %ld ", Description:"", MySQLVersion:"5.7"},
    1407: definition.ErrorDefinition{ErrorNumber:1407, ErrorType:"ServerError", Symbol:"ER_SP_BAD_SQLSTATE", SQLState:"42000", Message:"Bad SQLSTATE: '%s' ", Description:"", MySQLVersion:"5.7"},
    1408: definition.ErrorDefinition{ErrorNumber:1408, ErrorType:"ServerError", Symbol:"ER_STARTUP", SQLState:"HY000", Message:"%s: ready for connections. Version: '%s' socket: '%s' port: %d %s ", Description:"", MySQLVersion:"5.7"},
    1409: definition.ErrorDefinition{ErrorNumber:1409, ErrorType:"ServerError", Symbol:"ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR", SQLState:"HY000", Message:"Can't load value from file with fixed size rows to variable ", Description:"", MySQLVersion:"5.7"},
    1410: definition.ErrorDefinition{ErrorNumber:1410, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_USER_WITH_GRANT", SQLState:"42000", Message:"You are not allowed to create a user with GRANT ", Description:"", MySQLVersion:"5.7"},
    1411: definition.ErrorDefinition{ErrorNumber:1411, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_FOR_TYPE", SQLState:"HY000", Message:"Incorrect %s value: '%s' for function %s ", Description:"", MySQLVersion:"5.7"},
    1412: definition.ErrorDefinition{ErrorNumber:1412, ErrorType:"ServerError", Symbol:"ER_TABLE_DEF_CHANGED", SQLState:"HY000", Message:"Table definition has changed, please retry transaction ", Description:"", MySQLVersion:"5.7"},
    1413: definition.ErrorDefinition{ErrorNumber:1413, ErrorType:"ServerError", Symbol:"ER_SP_DUP_HANDLER", SQLState:"42000", Message:"Duplicate handler declared in the same block ", Description:"", MySQLVersion:"5.7"},
    1414: definition.ErrorDefinition{ErrorNumber:1414, ErrorType:"ServerError", Symbol:"ER_SP_NOT_VAR_ARG", SQLState:"42000", Message:"OUT or INOUT argument %d for routine %s is not a variable or NEW pseudo-variable in BEFORE trigger ", Description:"", MySQLVersion:"5.7"},
    1415: definition.ErrorDefinition{ErrorNumber:1415, ErrorType:"ServerError", Symbol:"ER_SP_NO_RETSET", SQLState:"0A000", Message:"Not allowed to return a result set from a %s ", Description:"", MySQLVersion:"5.7"},
    1416: definition.ErrorDefinition{ErrorNumber:1416, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_GEOMETRY_OBJECT", SQLState:"22003", Message:"Cannot get geometry object from data you send to the GEOMETRY field ", Description:"", MySQLVersion:"5.7"},
    1417: definition.ErrorDefinition{ErrorNumber:1417, ErrorType:"ServerError", Symbol:"ER_FAILED_ROUTINE_BREAK_BINLOG", SQLState:"HY000", Message:"A routine failed and has neither NO SQL nor READS SQL DATA in its declaration and binary logging is enabled", Description:"if non-transactional tables were updated, the binary log will miss their changes ", MySQLVersion:"5.7"},
    1418: definition.ErrorDefinition{ErrorNumber:1418, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_ROUTINE", SQLState:"HY000", Message:"This function has none of DETERMINISTIC, NO SQL, or READS SQL DATA in its declaration and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable) ", Description:"", MySQLVersion:"5.7"},
    1419: definition.ErrorDefinition{ErrorNumber:1419, ErrorType:"ServerError", Symbol:"ER_BINLOG_CREATE_ROUTINE_NEED_SUPER", SQLState:"HY000", Message:"You do not have the SUPER privilege and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable) ", Description:"", MySQLVersion:"5.7"},
    1420: definition.ErrorDefinition{ErrorNumber:1420, ErrorType:"ServerError", Symbol:"ER_EXEC_STMT_WITH_OPEN_CURSOR", SQLState:"HY000", Message:"You can't execute a prepared statement which has an open cursor associated with it. Reset the statement to re-execute it. ", Description:"", MySQLVersion:"5.7"},
    1421: definition.ErrorDefinition{ErrorNumber:1421, ErrorType:"ServerError", Symbol:"ER_STMT_HAS_NO_OPEN_CURSOR", SQLState:"HY000", Message:"The statement (%lu) has no open cursor. ", Description:"", MySQLVersion:"5.7"},
    1422: definition.ErrorDefinition{ErrorNumber:1422, ErrorType:"ServerError", Symbol:"ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG", SQLState:"HY000", Message:"Explicit or implicit commit is not allowed in stored function or trigger. ", Description:"", MySQLVersion:"5.7"},
    1423: definition.ErrorDefinition{ErrorNumber:1423, ErrorType:"ServerError", Symbol:"ER_NO_DEFAULT_FOR_VIEW_FIELD", SQLState:"HY000", Message:"Field of view '%s.%s' underlying table doesn't have a default value ", Description:"", MySQLVersion:"5.7"},
    1424: definition.ErrorDefinition{ErrorNumber:1424, ErrorType:"ServerError", Symbol:"ER_SP_NO_RECURSION", SQLState:"HY000", Message:"Recursive stored functions and triggers are not allowed. ", Description:"", MySQLVersion:"5.7"},
    1425: definition.ErrorDefinition{ErrorNumber:1425, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_SCALE", SQLState:"42000", Message:"Too big scale %d specified for column '%s'. Maximum is %lu. ", Description:"", MySQLVersion:"5.7"},
    1426: definition.ErrorDefinition{ErrorNumber:1426, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_PRECISION", SQLState:"42000", Message:"Too-big precision %d specified for '%s'. Maximum is %lu. ", Description:"", MySQLVersion:"5.7"},
    1427: definition.ErrorDefinition{ErrorNumber:1427, ErrorType:"ServerError", Symbol:"ER_M_BIGGER_THAN_D", SQLState:"42000", Message:"For float(M,D), double(M,D) or decimal(M,D), M must be >= D (column '%s'). ", Description:"", MySQLVersion:"5.7"},
    1428: definition.ErrorDefinition{ErrorNumber:1428, ErrorType:"ServerError", Symbol:"ER_WRONG_LOCK_OF_SYSTEM_TABLE", SQLState:"HY000", Message:"You can't combine write-locking of system tables with other tables or lock types ", Description:"", MySQLVersion:"5.7"},
    1429: definition.ErrorDefinition{ErrorNumber:1429, ErrorType:"ServerError", Symbol:"ER_CONNECT_TO_FOREIGN_DATA_SOURCE", SQLState:"HY000", Message:"Unable to connect to foreign data source: %s ", Description:"", MySQLVersion:"5.7"},
    1430: definition.ErrorDefinition{ErrorNumber:1430, ErrorType:"ServerError", Symbol:"ER_QUERY_ON_FOREIGN_DATA_SOURCE", SQLState:"HY000", Message:"There was a problem processing the query on the foreign data source. Data source error: %s ", Description:"", MySQLVersion:"5.7"},
    1431: definition.ErrorDefinition{ErrorNumber:1431, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST", SQLState:"HY000", Message:"The foreign data source you are trying to reference does not exist. Data source error: %s ", Description:"", MySQLVersion:"5.7"},
    1432: definition.ErrorDefinition{ErrorNumber:1432, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE", SQLState:"HY000", Message:"Can't create federated table. The data source connection string '%s' is not in the correct format ", Description:"", MySQLVersion:"5.7"},
    1433: definition.ErrorDefinition{ErrorNumber:1433, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DATA_STRING_INVALID", SQLState:"HY000", Message:"The data source connection string '%s' is not in the correct format ", Description:"", MySQLVersion:"5.7"},
    1434: definition.ErrorDefinition{ErrorNumber:1434, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_FEDERATED_TABLE", SQLState:"HY000", Message:"Can't create federated table. Foreign data src error: %s ", Description:"", MySQLVersion:"5.7"},
    1435: definition.ErrorDefinition{ErrorNumber:1435, ErrorType:"ServerError", Symbol:"ER_TRG_IN_WRONG_SCHEMA", SQLState:"HY000", Message:"Trigger in wrong schema ", Description:"", MySQLVersion:"5.7"},
    1436: definition.ErrorDefinition{ErrorNumber:1436, ErrorType:"ServerError", Symbol:"ER_STACK_OVERRUN_NEED_MORE", SQLState:"HY000", Message:"Thread stack overrun: %ld bytes used of a %ld byte stack, and %ld bytes needed. Use 'mysqld --thread_stack=#' to specify a bigger stack. ", Description:"", MySQLVersion:"5.7"},
    1437: definition.ErrorDefinition{ErrorNumber:1437, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_BODY", SQLState:"42000", Message:"Routine body for '%s' is too long ", Description:"", MySQLVersion:"5.7"},
    1438: definition.ErrorDefinition{ErrorNumber:1438, ErrorType:"ServerError", Symbol:"ER_WARN_CANT_DROP_DEFAULT_KEYCACHE", SQLState:"HY000", Message:"Cannot drop default keycache ", Description:"", MySQLVersion:"5.7"},
    1439: definition.ErrorDefinition{ErrorNumber:1439, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_DISPLAYWIDTH", SQLState:"42000", Message:"Display width out of range for column '%s' (max = %lu) ", Description:"", MySQLVersion:"5.7"},
    1440: definition.ErrorDefinition{ErrorNumber:1440, ErrorType:"ServerError", Symbol:"ER_XAER_DUPID", SQLState:"XAE08", Message:"XAER_DUPID: The XID already exists ", Description:"", MySQLVersion:"5.7"},
    1441: definition.ErrorDefinition{ErrorNumber:1441, ErrorType:"ServerError", Symbol:"ER_DATETIME_FUNCTION_OVERFLOW", SQLState:"22008", Message:"Datetime function: %s field overflow ", Description:"", MySQLVersion:"5.7"},
    1442: definition.ErrorDefinition{ErrorNumber:1442, ErrorType:"ServerError", Symbol:"ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG", SQLState:"HY000", Message:"Can't update table '%s' in stored function/trigger because it is already used by statement which invoked this stored function/trigger. ", Description:"", MySQLVersion:"5.7"},
    1443: definition.ErrorDefinition{ErrorNumber:1443, ErrorType:"ServerError", Symbol:"ER_VIEW_PREVENT_UPDATE", SQLState:"HY000", Message:"The definition of table '%s' prevents operation %s on table '%s'. ", Description:"", MySQLVersion:"5.7"},
    1444: definition.ErrorDefinition{ErrorNumber:1444, ErrorType:"ServerError", Symbol:"ER_PS_NO_RECURSION", SQLState:"HY000", Message:"The prepared statement contains a stored routine call that refers to that same statement. It's not allowed to execute a prepared statement in such a recursive manner ", Description:"", MySQLVersion:"5.7"},
    1445: definition.ErrorDefinition{ErrorNumber:1445, ErrorType:"ServerError", Symbol:"ER_SP_CANT_SET_AUTOCOMMIT", SQLState:"HY000", Message:"Not allowed to set autocommit from a stored function or trigger ", Description:"", MySQLVersion:"5.7"},
    1446: definition.ErrorDefinition{ErrorNumber:1446, ErrorType:"ServerError", Symbol:"ER_MALFORMED_DEFINER", SQLState:"HY000", Message:"Definer is not fully qualified ", Description:"", MySQLVersion:"5.7"},
    1447: definition.ErrorDefinition{ErrorNumber:1447, ErrorType:"ServerError", Symbol:"ER_VIEW_FRM_NO_USER", SQLState:"HY000", Message:"View '%s'.'%s' has no definer information (old table format). Current user is used as definer. Please recreate the view! ", Description:"", MySQLVersion:"5.7"},
    1448: definition.ErrorDefinition{ErrorNumber:1448, ErrorType:"ServerError", Symbol:"ER_VIEW_OTHER_USER", SQLState:"HY000", Message:"You need the SUPER privilege for creation view with '%s'@'%s' definer ", Description:"", MySQLVersion:"5.7"},
    1449: definition.ErrorDefinition{ErrorNumber:1449, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_USER", SQLState:"HY000", Message:"The user specified as a definer ('%s'@'%s') does not exist ", Description:"", MySQLVersion:"5.7"},
    1450: definition.ErrorDefinition{ErrorNumber:1450, ErrorType:"ServerError", Symbol:"ER_FORBID_SCHEMA_CHANGE", SQLState:"HY000", Message:"Changing schema from '%s' to '%s' is not allowed. ", Description:"", MySQLVersion:"5.7"},
    1451: definition.ErrorDefinition{ErrorNumber:1451, ErrorType:"ServerError", Symbol:"ER_ROW_IS_REFERENCED_2", SQLState:"23000", Message:"Cannot delete or update a parent row: a foreign key constraint fails (%s)", Description:"InnoDB reports this error when you try to delete a parent row that has children, and a foreign key constraint fails. Delete the children first. ", MySQLVersion:"5.7"},
    1452: definition.ErrorDefinition{ErrorNumber:1452, ErrorType:"ServerError", Symbol:"ER_NO_REFERENCED_ROW_2", SQLState:"23000", Message:"Cannot add or update a child row: a foreign key constraint fails (%s)", Description:"InnoDB reports this error when you try to add a row but there is no parent row, and a foreign key constraint fails. Add the parent row first. ", MySQLVersion:"5.7"},
    1453: definition.ErrorDefinition{ErrorNumber:1453, ErrorType:"ServerError", Symbol:"ER_SP_BAD_VAR_SHADOW", SQLState:"42000", Message:"Variable '%s' must be quoted with `...`, or renamed ", Description:"", MySQLVersion:"5.7"},
    1454: definition.ErrorDefinition{ErrorNumber:1454, ErrorType:"ServerError", Symbol:"ER_TRG_NO_DEFINER", SQLState:"HY000", Message:"No definer attribute for trigger '%s'.'%s'. The trigger will be activated under the authorization of the invoker, which may have insufficient privileges. Please recreate the trigger. ", Description:"", MySQLVersion:"5.7"},
    1455: definition.ErrorDefinition{ErrorNumber:1455, ErrorType:"ServerError", Symbol:"ER_OLD_FILE_FORMAT", SQLState:"HY000", Message:"'%s' has an old format, you should re-create the '%s' object(s) ", Description:"", MySQLVersion:"5.7"},
    1456: definition.ErrorDefinition{ErrorNumber:1456, ErrorType:"ServerError", Symbol:"ER_SP_RECURSION_LIMIT", SQLState:"HY000", Message:"Recursive limit %d (as set by the max_sp_recursion_depth variable) was exceeded for routine %s ", Description:"", MySQLVersion:"5.7"},
    1457: definition.ErrorDefinition{ErrorNumber:1457, ErrorType:"ServerError", Symbol:"ER_SP_PROC_TABLE_CORRUPT", SQLState:"HY000", Message:"Failed to load routine %s. The table mysql.proc is missing, corrupt, or contains bad data (internal code %d) ", Description:"", MySQLVersion:"5.7"},
    1458: definition.ErrorDefinition{ErrorNumber:1458, ErrorType:"ServerError", Symbol:"ER_SP_WRONG_NAME", SQLState:"42000", Message:"Incorrect routine name '%s' ", Description:"", MySQLVersion:"5.7"},
    1459: definition.ErrorDefinition{ErrorNumber:1459, ErrorType:"ServerError", Symbol:"ER_TABLE_NEEDS_UPGRADE", SQLState:"HY000", Message:"Table upgrade required. Please do \"REPAIR TABLE `%s`\" or dump/reload to fix it! ", Description:"", MySQLVersion:"5.7"},
    1460: definition.ErrorDefinition{ErrorNumber:1460, ErrorType:"ServerError", Symbol:"ER_SP_NO_AGGREGATE", SQLState:"42000", Message:"AGGREGATE is not supported for stored functions ", Description:"", MySQLVersion:"5.7"},
    1461: definition.ErrorDefinition{ErrorNumber:1461, ErrorType:"ServerError", Symbol:"ER_MAX_PREPARED_STMT_COUNT_REACHED", SQLState:"42000", Message:"Can't create more than max_prepared_stmt_count statements (current value: %lu) ", Description:"", MySQLVersion:"5.7"},
    1462: definition.ErrorDefinition{ErrorNumber:1462, ErrorType:"ServerError", Symbol:"ER_VIEW_RECURSIVE", SQLState:"HY000", Message:"`%s`.`%s` contains view recursion ", Description:"", MySQLVersion:"5.7"},
    1463: definition.ErrorDefinition{ErrorNumber:1463, ErrorType:"ServerError", Symbol:"ER_NON_GROUPING_FIELD_USED", SQLState:"42000", Message:"Non-grouping field '%s' is used in %s clause ", Description:"", MySQLVersion:"5.7"},
    1464: definition.ErrorDefinition{ErrorNumber:1464, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_SPKEYS", SQLState:"HY000", Message:"The used table type doesn't support SPATIAL indexes ", Description:"", MySQLVersion:"5.7"},
    1465: definition.ErrorDefinition{ErrorNumber:1465, ErrorType:"ServerError", Symbol:"ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA", SQLState:"HY000", Message:"Triggers can not be created on system tables ", Description:"", MySQLVersion:"5.7"},
    1466: definition.ErrorDefinition{ErrorNumber:1466, ErrorType:"ServerError", Symbol:"ER_REMOVED_SPACES", SQLState:"HY000", Message:"Leading spaces are removed from name '%s' ", Description:"", MySQLVersion:"5.7"},
    1467: definition.ErrorDefinition{ErrorNumber:1467, ErrorType:"ServerError", Symbol:"ER_AUTOINC_READ_FAILED", SQLState:"HY000", Message:"Failed to read auto-increment value from storage engine ", Description:"", MySQLVersion:"5.7"},
    1468: definition.ErrorDefinition{ErrorNumber:1468, ErrorType:"ServerError", Symbol:"ER_USERNAME", SQLState:"HY000", Message:"user name ", Description:"", MySQLVersion:"5.7"},
    1469: definition.ErrorDefinition{ErrorNumber:1469, ErrorType:"ServerError", Symbol:"ER_HOSTNAME", SQLState:"HY000", Message:"host name ", Description:"", MySQLVersion:"5.7"},
    1470: definition.ErrorDefinition{ErrorNumber:1470, ErrorType:"ServerError", Symbol:"ER_WRONG_STRING_LENGTH", SQLState:"HY000", Message:"String '%s' is too long for %s (should be no longer than %d) ", Description:"", MySQLVersion:"5.7"},
    1471: definition.ErrorDefinition{ErrorNumber:1471, ErrorType:"ServerError", Symbol:"ER_NON_INSERTABLE_TABLE", SQLState:"HY000", Message:"The target table %s of the %s is not insertable-into ", Description:"", MySQLVersion:"5.7"},
    1472: definition.ErrorDefinition{ErrorNumber:1472, ErrorType:"ServerError", Symbol:"ER_ADMIN_WRONG_MRG_TABLE", SQLState:"HY000", Message:"Table '%s' is differently defined or of non-MyISAM type or doesn't exist ", Description:"", MySQLVersion:"5.7"},
    1473: definition.ErrorDefinition{ErrorNumber:1473, ErrorType:"ServerError", Symbol:"ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT", SQLState:"HY000", Message:"Too high level of nesting for select ", Description:"", MySQLVersion:"5.7"},
    1474: definition.ErrorDefinition{ErrorNumber:1474, ErrorType:"ServerError", Symbol:"ER_NAME_BECOMES_EMPTY", SQLState:"HY000", Message:"Name '%s' has become '' ", Description:"", MySQLVersion:"5.7"},
    1475: definition.ErrorDefinition{ErrorNumber:1475, ErrorType:"ServerError", Symbol:"ER_AMBIGUOUS_FIELD_TERM", SQLState:"HY000", Message:"First character of the FIELDS TERMINATED string is ambiguous", Description:"please use non-optional and non-empty FIELDS ENCLOSED BY ", MySQLVersion:"5.7"},
    1476: definition.ErrorDefinition{ErrorNumber:1476, ErrorType:"ServerError", Symbol:"ER_FOREIGN_SERVER_EXISTS", SQLState:"HY000", Message:"The foreign server, %s, you are trying to create already exists. ", Description:"", MySQLVersion:"5.7"},
    1477: definition.ErrorDefinition{ErrorNumber:1477, ErrorType:"ServerError", Symbol:"ER_FOREIGN_SERVER_DOESNT_EXIST", SQLState:"HY000", Message:"The foreign server name you are trying to reference does not exist. Data source error: %s ", Description:"", MySQLVersion:"5.7"},
    1478: definition.ErrorDefinition{ErrorNumber:1478, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_HA_CREATE_OPTION", SQLState:"HY000", Message:"Table storage engine '%s' does not support the create option '%s' ", Description:"", MySQLVersion:"5.7"},
    1479: definition.ErrorDefinition{ErrorNumber:1479, ErrorType:"ServerError", Symbol:"ER_PARTITION_REQUIRES_VALUES_ERROR", SQLState:"HY000", Message:"Syntax error: %s PARTITIONING requires definition of VALUES %s for each partition ", Description:"", MySQLVersion:"5.7"},
    1480: definition.ErrorDefinition{ErrorNumber:1480, ErrorType:"ServerError", Symbol:"ER_PARTITION_WRONG_VALUES_ERROR", SQLState:"HY000", Message:"Only %s PARTITIONING can use VALUES %s in partition definition ", Description:"", MySQLVersion:"5.7"},
    1481: definition.ErrorDefinition{ErrorNumber:1481, ErrorType:"ServerError", Symbol:"ER_PARTITION_MAXVALUE_ERROR", SQLState:"HY000", Message:"MAXVALUE can only be used in last partition definition ", Description:"", MySQLVersion:"5.7"},
    1482: definition.ErrorDefinition{ErrorNumber:1482, ErrorType:"ServerError", Symbol:"ER_PARTITION_SUBPARTITION_ERROR", SQLState:"HY000", Message:"Subpartitions can only be hash partitions and by key ", Description:"", MySQLVersion:"5.7"},
    1483: definition.ErrorDefinition{ErrorNumber:1483, ErrorType:"ServerError", Symbol:"ER_PARTITION_SUBPART_MIX_ERROR", SQLState:"HY000", Message:"Must define subpartitions on all partitions if on one partition ", Description:"", MySQLVersion:"5.7"},
    1484: definition.ErrorDefinition{ErrorNumber:1484, ErrorType:"ServerError", Symbol:"ER_PARTITION_WRONG_NO_PART_ERROR", SQLState:"HY000", Message:"Wrong number of partitions defined, mismatch with previous setting ", Description:"", MySQLVersion:"5.7"},
    1485: definition.ErrorDefinition{ErrorNumber:1485, ErrorType:"ServerError", Symbol:"ER_PARTITION_WRONG_NO_SUBPART_ERROR", SQLState:"HY000", Message:"Wrong number of subpartitions defined, mismatch with previous setting ", Description:"", MySQLVersion:"5.7"},
    1486: definition.ErrorDefinition{ErrorNumber:1486, ErrorType:"ServerError", Symbol:"ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR", SQLState:"HY000", Message:"Constant, random or timezone-dependent expressions in (sub)partitioning function are not allowed ", Description:"", MySQLVersion:"5.7"},
    1487: definition.ErrorDefinition{ErrorNumber:1487, ErrorType:"ServerError", Symbol:"ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR", SQLState:"HY000", Message:"Expression in RANGE/LIST VALUES must be constant ", Description:"", MySQLVersion:"5.7"},
    1488: definition.ErrorDefinition{ErrorNumber:1488, ErrorType:"ServerError", Symbol:"ER_FIELD_NOT_FOUND_PART_ERROR", SQLState:"HY000", Message:"Field in list of fields for partition function not found in table ", Description:"", MySQLVersion:"5.7"},
    1489: definition.ErrorDefinition{ErrorNumber:1489, ErrorType:"ServerError", Symbol:"ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR", SQLState:"HY000", Message:"List of fields is only allowed in KEY partitions ", Description:"", MySQLVersion:"5.7"},
    1490: definition.ErrorDefinition{ErrorNumber:1490, ErrorType:"ServerError", Symbol:"ER_INCONSISTENT_PARTITION_INFO_ERROR", SQLState:"HY000", Message:"The partition info in the frm file is not consistent with what can be written into the frm file ", Description:"", MySQLVersion:"5.7"},
    1491: definition.ErrorDefinition{ErrorNumber:1491, ErrorType:"ServerError", Symbol:"ER_PARTITION_FUNC_NOT_ALLOWED_ERROR", SQLState:"HY000", Message:"The %s function returns the wrong type ", Description:"", MySQLVersion:"5.7"},
    1492: definition.ErrorDefinition{ErrorNumber:1492, ErrorType:"ServerError", Symbol:"ER_PARTITIONS_MUST_BE_DEFINED_ERROR", SQLState:"HY000", Message:"For %s partitions each partition must be defined ", Description:"", MySQLVersion:"5.7"},
    1493: definition.ErrorDefinition{ErrorNumber:1493, ErrorType:"ServerError", Symbol:"ER_RANGE_NOT_INCREASING_ERROR", SQLState:"HY000", Message:"VALUES LESS THAN value must be strictly increasing for each partition ", Description:"", MySQLVersion:"5.7"},
    1494: definition.ErrorDefinition{ErrorNumber:1494, ErrorType:"ServerError", Symbol:"ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR", SQLState:"HY000", Message:"VALUES value must be of same type as partition function ", Description:"", MySQLVersion:"5.7"},
    1495: definition.ErrorDefinition{ErrorNumber:1495, ErrorType:"ServerError", Symbol:"ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR", SQLState:"HY000", Message:"Multiple definition of same constant in list partitioning ", Description:"", MySQLVersion:"5.7"},
    1496: definition.ErrorDefinition{ErrorNumber:1496, ErrorType:"ServerError", Symbol:"ER_PARTITION_ENTRY_ERROR", SQLState:"HY000", Message:"Partitioning can not be used stand-alone in query ", Description:"", MySQLVersion:"5.7"},
    1497: definition.ErrorDefinition{ErrorNumber:1497, ErrorType:"ServerError", Symbol:"ER_MIX_HANDLER_ERROR", SQLState:"HY000", Message:"The mix of handlers in the partitions is not allowed in this version of MySQL ", Description:"", MySQLVersion:"5.7"},
    1498: definition.ErrorDefinition{ErrorNumber:1498, ErrorType:"ServerError", Symbol:"ER_PARTITION_NOT_DEFINED_ERROR", SQLState:"HY000", Message:"For the partitioned engine it is necessary to define all %s ", Description:"", MySQLVersion:"5.7"},
    1499: definition.ErrorDefinition{ErrorNumber:1499, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_PARTITIONS_ERROR", SQLState:"HY000", Message:"Too many partitions (including subpartitions) were defined ", Description:"", MySQLVersion:"5.7"},
    1500: definition.ErrorDefinition{ErrorNumber:1500, ErrorType:"ServerError", Symbol:"ER_SUBPARTITION_ERROR", SQLState:"HY000", Message:"It is only possible to mix RANGE/LIST partitioning with HASH/KEY partitioning for subpartitioning ", Description:"", MySQLVersion:"5.7"},
    1501: definition.ErrorDefinition{ErrorNumber:1501, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_HANDLER_FILE", SQLState:"HY000", Message:"Failed to create specific handler file ", Description:"", MySQLVersion:"5.7"},
    1502: definition.ErrorDefinition{ErrorNumber:1502, ErrorType:"ServerError", Symbol:"ER_BLOB_FIELD_IN_PART_FUNC_ERROR", SQLState:"HY000", Message:"A BLOB field is not allowed in partition function ", Description:"", MySQLVersion:"5.7"},
    1503: definition.ErrorDefinition{ErrorNumber:1503, ErrorType:"ServerError", Symbol:"ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF", SQLState:"HY000", Message:"A %s must include all columns in the table's partitioning function ", Description:"", MySQLVersion:"5.7"},
    1504: definition.ErrorDefinition{ErrorNumber:1504, ErrorType:"ServerError", Symbol:"ER_NO_PARTS_ERROR", SQLState:"HY000", Message:"Number of %s = 0 is not an allowed value ", Description:"", MySQLVersion:"5.7"},
    1505: definition.ErrorDefinition{ErrorNumber:1505, ErrorType:"ServerError", Symbol:"ER_PARTITION_MGMT_ON_NONPARTITIONED", SQLState:"HY000", Message:"Partition management on a not partitioned table is not possible ", Description:"", MySQLVersion:"5.7"},
    1506: definition.ErrorDefinition{ErrorNumber:1506, ErrorType:"ServerError", Symbol:"ER_FOREIGN_KEY_ON_PARTITIONED", SQLState:"HY000", Message:"Foreign keys are not yet supported in conjunction with partitioning ", Description:"", MySQLVersion:"5.7"},
    1507: definition.ErrorDefinition{ErrorNumber:1507, ErrorType:"ServerError", Symbol:"ER_DROP_PARTITION_NON_EXISTENT", SQLState:"HY000", Message:"Error in list of partitions to %s ", Description:"", MySQLVersion:"5.7"},
    1508: definition.ErrorDefinition{ErrorNumber:1508, ErrorType:"ServerError", Symbol:"ER_DROP_LAST_PARTITION", SQLState:"HY000", Message:"Cannot remove all partitions, use DROP TABLE instead ", Description:"", MySQLVersion:"5.7"},
    1509: definition.ErrorDefinition{ErrorNumber:1509, ErrorType:"ServerError", Symbol:"ER_COALESCE_ONLY_ON_HASH_PARTITION", SQLState:"HY000", Message:"COALESCE PARTITION can only be used on HASH/KEY partitions ", Description:"", MySQLVersion:"5.7"},
    1510: definition.ErrorDefinition{ErrorNumber:1510, ErrorType:"ServerError", Symbol:"ER_REORG_HASH_ONLY_ON_SAME_NO", SQLState:"HY000", Message:"REORGANIZE PARTITION can only be used to reorganize partitions not to change their numbers ", Description:"", MySQLVersion:"5.7"},
    1511: definition.ErrorDefinition{ErrorNumber:1511, ErrorType:"ServerError", Symbol:"ER_REORG_NO_PARAM_ERROR", SQLState:"HY000", Message:"REORGANIZE PARTITION without parameters can only be used on auto-partitioned tables using HASH PARTITIONs ", Description:"", MySQLVersion:"5.7"},
    1512: definition.ErrorDefinition{ErrorNumber:1512, ErrorType:"ServerError", Symbol:"ER_ONLY_ON_RANGE_LIST_PARTITION", SQLState:"HY000", Message:"%s PARTITION can only be used on RANGE/LIST partitions ", Description:"", MySQLVersion:"5.7"},
    1513: definition.ErrorDefinition{ErrorNumber:1513, ErrorType:"ServerError", Symbol:"ER_ADD_PARTITION_SUBPART_ERROR", SQLState:"HY000", Message:"Trying to Add partition(s) with wrong number of subpartitions ", Description:"", MySQLVersion:"5.7"},
    1514: definition.ErrorDefinition{ErrorNumber:1514, ErrorType:"ServerError", Symbol:"ER_ADD_PARTITION_NO_NEW_PARTITION", SQLState:"HY000", Message:"At least one partition must be added ", Description:"", MySQLVersion:"5.7"},
    1515: definition.ErrorDefinition{ErrorNumber:1515, ErrorType:"ServerError", Symbol:"ER_COALESCE_PARTITION_NO_PARTITION", SQLState:"HY000", Message:"At least one partition must be coalesced ", Description:"", MySQLVersion:"5.7"},
    1516: definition.ErrorDefinition{ErrorNumber:1516, ErrorType:"ServerError", Symbol:"ER_REORG_PARTITION_NOT_EXIST", SQLState:"HY000", Message:"More partitions to reorganize than there are partitions ", Description:"", MySQLVersion:"5.7"},
    1517: definition.ErrorDefinition{ErrorNumber:1517, ErrorType:"ServerError", Symbol:"ER_SAME_NAME_PARTITION", SQLState:"HY000", Message:"Duplicate partition name %s ", Description:"", MySQLVersion:"5.7"},
    1518: definition.ErrorDefinition{ErrorNumber:1518, ErrorType:"ServerError", Symbol:"ER_NO_BINLOG_ERROR", SQLState:"HY000", Message:"It is not allowed to shut off binlog on this command ", Description:"", MySQLVersion:"5.7"},
    1519: definition.ErrorDefinition{ErrorNumber:1519, ErrorType:"ServerError", Symbol:"ER_CONSECUTIVE_REORG_PARTITIONS", SQLState:"HY000", Message:"When reorganizing a set of partitions they must be in consecutive order ", Description:"", MySQLVersion:"5.7"},
    1520: definition.ErrorDefinition{ErrorNumber:1520, ErrorType:"ServerError", Symbol:"ER_REORG_OUTSIDE_RANGE", SQLState:"HY000", Message:"Reorganize of range partitions cannot change total ranges except for last partition where it can extend the range ", Description:"", MySQLVersion:"5.7"},
    1521: definition.ErrorDefinition{ErrorNumber:1521, ErrorType:"ServerError", Symbol:"ER_PARTITION_FUNCTION_FAILURE", SQLState:"HY000", Message:"Partition function not supported in this version for this handler ", Description:"", MySQLVersion:"5.7"},
    1522: definition.ErrorDefinition{ErrorNumber:1522, ErrorType:"ServerError", Symbol:"ER_PART_STATE_ERROR", SQLState:"HY000", Message:"Partition state cannot be defined from CREATE/ALTER TABLE ", Description:"", MySQLVersion:"5.7"},
    1523: definition.ErrorDefinition{ErrorNumber:1523, ErrorType:"ServerError", Symbol:"ER_LIMITED_PART_RANGE", SQLState:"HY000", Message:"The %s handler only supports 32 bit integers in VALUES ", Description:"", MySQLVersion:"5.7"},
    1524: definition.ErrorDefinition{ErrorNumber:1524, ErrorType:"ServerError", Symbol:"ER_PLUGIN_IS_NOT_LOADED", SQLState:"HY000", Message:"Plugin '%s' is not loaded ", Description:"", MySQLVersion:"5.7"},
    1525: definition.ErrorDefinition{ErrorNumber:1525, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE", SQLState:"HY000", Message:"Incorrect %s value: '%s' ", Description:"", MySQLVersion:"5.7"},
    1526: definition.ErrorDefinition{ErrorNumber:1526, ErrorType:"ServerError", Symbol:"ER_NO_PARTITION_FOR_GIVEN_VALUE", SQLState:"HY000", Message:"Table has no partition for value %s ", Description:"", MySQLVersion:"5.7"},
    1527: definition.ErrorDefinition{ErrorNumber:1527, ErrorType:"ServerError", Symbol:"ER_FILEGROUP_OPTION_ONLY_ONCE", SQLState:"HY000", Message:"It is not allowed to specify %s more than once ", Description:"", MySQLVersion:"5.7"},
    1528: definition.ErrorDefinition{ErrorNumber:1528, ErrorType:"ServerError", Symbol:"ER_CREATE_FILEGROUP_FAILED", SQLState:"HY000", Message:"Failed to create %s ", Description:"", MySQLVersion:"5.7"},
    1529: definition.ErrorDefinition{ErrorNumber:1529, ErrorType:"ServerError", Symbol:"ER_DROP_FILEGROUP_FAILED", SQLState:"HY000", Message:"Failed to drop %s ", Description:"", MySQLVersion:"5.7"},
    1530: definition.ErrorDefinition{ErrorNumber:1530, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_AUTO_EXTEND_ERROR", SQLState:"HY000", Message:"The handler doesn't support autoextend of tablespaces ", Description:"", MySQLVersion:"5.7"},
    1531: definition.ErrorDefinition{ErrorNumber:1531, ErrorType:"ServerError", Symbol:"ER_WRONG_SIZE_NUMBER", SQLState:"HY000", Message:"A size parameter was incorrectly specified, either number or on the form 10M ", Description:"", MySQLVersion:"5.7"},
    1532: definition.ErrorDefinition{ErrorNumber:1532, ErrorType:"ServerError", Symbol:"ER_SIZE_OVERFLOW_ERROR", SQLState:"HY000", Message:"The size number was correct but we don't allow the digit part to be more than 2 billion ", Description:"", MySQLVersion:"5.7"},
    1533: definition.ErrorDefinition{ErrorNumber:1533, ErrorType:"ServerError", Symbol:"ER_ALTER_FILEGROUP_FAILED", SQLState:"HY000", Message:"Failed to alter: %s ", Description:"", MySQLVersion:"5.7"},
    1534: definition.ErrorDefinition{ErrorNumber:1534, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_LOGGING_FAILED", SQLState:"HY000", Message:"Writing one row to the row-based binary log failed ", Description:"", MySQLVersion:"5.7"},
    1535: definition.ErrorDefinition{ErrorNumber:1535, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_WRONG_TABLE_DEF", SQLState:"HY000", Message:"Table definition on master and slave does not match: %s ", Description:"", MySQLVersion:"5.7"},
    1536: definition.ErrorDefinition{ErrorNumber:1536, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_RBR_TO_SBR", SQLState:"HY000", Message:"Slave running with --log-slave-updates must use row-based binary logging to be able to replicate row-based binary log events ", Description:"", MySQLVersion:"5.7"},
    1537: definition.ErrorDefinition{ErrorNumber:1537, ErrorType:"ServerError", Symbol:"ER_EVENT_ALREADY_EXISTS", SQLState:"HY000", Message:"Event '%s' already exists ", Description:"", MySQLVersion:"5.7"},
    1538: definition.ErrorDefinition{ErrorNumber:1538, ErrorType:"ServerError", Symbol:"ER_EVENT_STORE_FAILED", SQLState:"HY000", Message:"Failed to store event %s. Error code %d from storage engine. ", Description:"", MySQLVersion:"5.7"},
    1539: definition.ErrorDefinition{ErrorNumber:1539, ErrorType:"ServerError", Symbol:"ER_EVENT_DOES_NOT_EXIST", SQLState:"HY000", Message:"Unknown event '%s' ", Description:"", MySQLVersion:"5.7"},
    1540: definition.ErrorDefinition{ErrorNumber:1540, ErrorType:"ServerError", Symbol:"ER_EVENT_CANT_ALTER", SQLState:"HY000", Message:"Failed to alter event '%s' ", Description:"", MySQLVersion:"5.7"},
    1541: definition.ErrorDefinition{ErrorNumber:1541, ErrorType:"ServerError", Symbol:"ER_EVENT_DROP_FAILED", SQLState:"HY000", Message:"Failed to drop %s ", Description:"", MySQLVersion:"5.7"},
    1542: definition.ErrorDefinition{ErrorNumber:1542, ErrorType:"ServerError", Symbol:"ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG", SQLState:"HY000", Message:"INTERVAL is either not positive or too big ", Description:"", MySQLVersion:"5.7"},
    1543: definition.ErrorDefinition{ErrorNumber:1543, ErrorType:"ServerError", Symbol:"ER_EVENT_ENDS_BEFORE_STARTS", SQLState:"HY000", Message:"ENDS is either invalid or before STARTS ", Description:"", MySQLVersion:"5.7"},
    1544: definition.ErrorDefinition{ErrorNumber:1544, ErrorType:"ServerError", Symbol:"ER_EVENT_EXEC_TIME_IN_THE_PAST", SQLState:"HY000", Message:"Event execution time is in the past. Event has been disabled ", Description:"", MySQLVersion:"5.7"},
    1545: definition.ErrorDefinition{ErrorNumber:1545, ErrorType:"ServerError", Symbol:"ER_EVENT_OPEN_TABLE_FAILED", SQLState:"HY000", Message:"Failed to open mysql.event ", Description:"", MySQLVersion:"5.7"},
    1546: definition.ErrorDefinition{ErrorNumber:1546, ErrorType:"ServerError", Symbol:"ER_EVENT_NEITHER_M_EXPR_NOR_M_AT", SQLState:"HY000", Message:"No datetime expression provided ", Description:"", MySQLVersion:"5.7"},
    1547: definition.ErrorDefinition{ErrorNumber:1547, ErrorType:"ServerError", Symbol:"ER_OBSOLETE_COL_COUNT_DOESNT_MATCH_CORRUPTED", SQLState:"HY000", Message:"Column count of mysql.%s is wrong. Expected %d, found %d. The table is probably corrupted ", Description:"", MySQLVersion:"5.7"},
    1548: definition.ErrorDefinition{ErrorNumber:1548, ErrorType:"ServerError", Symbol:"ER_OBSOLETE_CANNOT_LOAD_FROM_TABLE", SQLState:"HY000", Message:"Cannot load from mysql.%s. The table is probably corrupted ", Description:"", MySQLVersion:"5.7"},
    1549: definition.ErrorDefinition{ErrorNumber:1549, ErrorType:"ServerError", Symbol:"ER_EVENT_CANNOT_DELETE", SQLState:"HY000", Message:"Failed to delete the event from mysql.event ", Description:"", MySQLVersion:"5.7"},
    1550: definition.ErrorDefinition{ErrorNumber:1550, ErrorType:"ServerError", Symbol:"ER_EVENT_COMPILE_ERROR", SQLState:"HY000", Message:"Error during compilation of event's body ", Description:"", MySQLVersion:"5.7"},
    1551: definition.ErrorDefinition{ErrorNumber:1551, ErrorType:"ServerError", Symbol:"ER_EVENT_SAME_NAME", SQLState:"HY000", Message:"Same old and new event name ", Description:"", MySQLVersion:"5.7"},
    1552: definition.ErrorDefinition{ErrorNumber:1552, ErrorType:"ServerError", Symbol:"ER_EVENT_DATA_TOO_LONG", SQLState:"HY000", Message:"Data for column '%s' too long ", Description:"", MySQLVersion:"5.7"},
    1553: definition.ErrorDefinition{ErrorNumber:1553, ErrorType:"ServerError", Symbol:"ER_DROP_INDEX_FK", SQLState:"HY000", Message:"Cannot drop index '%s': needed in a foreign key constraint", Description:"When you drop an index, InnoDB checks if the index is used for checking a foreign key constraint. It is still OK to drop the index if there is another index that can be used to enforce the same constraint. InnoDB prevents you from dropping the last index that can enforce a particular referential constraint. ", MySQLVersion:"5.7"},
    1554: definition.ErrorDefinition{ErrorNumber:1554, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYNTAX_WITH_VER", SQLState:"HY000", Message:"The syntax '%s' is deprecated and will be removed in MySQL %s. Please use %s instead ", Description:"", MySQLVersion:"5.7"},
    1555: definition.ErrorDefinition{ErrorNumber:1555, ErrorType:"ServerError", Symbol:"ER_CANT_WRITE_LOCK_LOG_TABLE", SQLState:"HY000", Message:"You can't write-lock a log table. Only read access is possible ", Description:"", MySQLVersion:"5.7"},
    1556: definition.ErrorDefinition{ErrorNumber:1556, ErrorType:"ServerError", Symbol:"ER_CANT_LOCK_LOG_TABLE", SQLState:"HY000", Message:"You can't use locks with log tables. ", Description:"", MySQLVersion:"5.7"},
    1557: definition.ErrorDefinition{ErrorNumber:1557, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED", SQLState:"23000", Message:"Upholding foreign key constraints for table '%s', entry '%s', key %d would lead to a duplicate entry ", Description:"", MySQLVersion:"5.7"},
    1558: definition.ErrorDefinition{ErrorNumber:1558, ErrorType:"ServerError", Symbol:"ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE", SQLState:"HY000", Message:"Column count of mysql.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please use mysql_upgrade to fix this error. ", Description:"", MySQLVersion:"5.7"},
    1559: definition.ErrorDefinition{ErrorNumber:1559, ErrorType:"ServerError", Symbol:"ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR", SQLState:"HY000", Message:"Cannot switch out of the row-based binary log format when the session has open temporary tables ", Description:"", MySQLVersion:"5.7"},
    1560: definition.ErrorDefinition{ErrorNumber:1560, ErrorType:"ServerError", Symbol:"ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT", SQLState:"HY000", Message:"Cannot change the binary logging format inside a stored function or trigger ", Description:"", MySQLVersion:"5.7"},
    1561: definition.ErrorDefinition{ErrorNumber:1561, ErrorType:"ServerError", Symbol:"ER_NDB_CANT_SWITCH_BINLOG_FORMAT", SQLState:"HY000", Message:"The NDB cluster engine does not support changing the binlog format on the fly yet ", Description:"", MySQLVersion:"5.7"},
    1562: definition.ErrorDefinition{ErrorNumber:1562, ErrorType:"ServerError", Symbol:"ER_PARTITION_NO_TEMPORARY", SQLState:"HY000", Message:"Cannot create temporary table with partitions ", Description:"", MySQLVersion:"5.7"},
    1563: definition.ErrorDefinition{ErrorNumber:1563, ErrorType:"ServerError", Symbol:"ER_PARTITION_CONST_DOMAIN_ERROR", SQLState:"HY000", Message:"Partition constant is out of partition function domain ", Description:"", MySQLVersion:"5.7"},
    1564: definition.ErrorDefinition{ErrorNumber:1564, ErrorType:"ServerError", Symbol:"ER_PARTITION_FUNCTION_IS_NOT_ALLOWED", SQLState:"HY000", Message:"This partition function is not allowed ", Description:"", MySQLVersion:"5.7"},
    1565: definition.ErrorDefinition{ErrorNumber:1565, ErrorType:"ServerError", Symbol:"ER_DDL_LOG_ERROR", SQLState:"HY000", Message:"Error in DDL log ", Description:"", MySQLVersion:"5.7"},
    1566: definition.ErrorDefinition{ErrorNumber:1566, ErrorType:"ServerError", Symbol:"ER_NULL_IN_VALUES_LESS_THAN", SQLState:"HY000", Message:"Not allowed to use NULL value in VALUES LESS THAN ", Description:"", MySQLVersion:"5.7"},
    1567: definition.ErrorDefinition{ErrorNumber:1567, ErrorType:"ServerError", Symbol:"ER_WRONG_PARTITION_NAME", SQLState:"HY000", Message:"Incorrect partition name ", Description:"", MySQLVersion:"5.7"},
    1568: definition.ErrorDefinition{ErrorNumber:1568, ErrorType:"ServerError", Symbol:"ER_CANT_CHANGE_TX_CHARACTERISTICS", SQLState:"25001", Message:"Transaction characteristics can't be changed while a transaction is in progress ", Description:"", MySQLVersion:"5.7"},
    1569: definition.ErrorDefinition{ErrorNumber:1569, ErrorType:"ServerError", Symbol:"ER_DUP_ENTRY_AUTOINCREMENT_CASE", SQLState:"HY000", Message:"ALTER TABLE causes auto_increment resequencing, resulting in duplicate entry '%s' for key '%s' ", Description:"", MySQLVersion:"5.7"},
    1570: definition.ErrorDefinition{ErrorNumber:1570, ErrorType:"ServerError", Symbol:"ER_EVENT_MODIFY_QUEUE_ERROR", SQLState:"HY000", Message:"Internal scheduler error %d ", Description:"", MySQLVersion:"5.7"},
    1571: definition.ErrorDefinition{ErrorNumber:1571, ErrorType:"ServerError", Symbol:"ER_EVENT_SET_VAR_ERROR", SQLState:"HY000", Message:"Error during starting/stopping of the scheduler. Error code %u ", Description:"", MySQLVersion:"5.7"},
    1572: definition.ErrorDefinition{ErrorNumber:1572, ErrorType:"ServerError", Symbol:"ER_PARTITION_MERGE_ERROR", SQLState:"HY000", Message:"Engine cannot be used in partitioned tables ", Description:"", MySQLVersion:"5.7"},
    1573: definition.ErrorDefinition{ErrorNumber:1573, ErrorType:"ServerError", Symbol:"ER_CANT_ACTIVATE_LOG", SQLState:"HY000", Message:"Cannot activate '%s' log ", Description:"", MySQLVersion:"5.7"},
    1574: definition.ErrorDefinition{ErrorNumber:1574, ErrorType:"ServerError", Symbol:"ER_RBR_NOT_AVAILABLE", SQLState:"HY000", Message:"The server was not built with row-based replication ", Description:"", MySQLVersion:"5.7"},
    1575: definition.ErrorDefinition{ErrorNumber:1575, ErrorType:"ServerError", Symbol:"ER_BASE64_DECODE_ERROR", SQLState:"HY000", Message:"Decoding of base64 string failed ", Description:"", MySQLVersion:"5.7"},
    1576: definition.ErrorDefinition{ErrorNumber:1576, ErrorType:"ServerError", Symbol:"ER_EVENT_RECURSION_FORBIDDEN", SQLState:"HY000", Message:"Recursion of EVENT DDL statements is forbidden when body is present ", Description:"", MySQLVersion:"5.7"},
    1577: definition.ErrorDefinition{ErrorNumber:1577, ErrorType:"ServerError", Symbol:"ER_EVENTS_DB_ERROR", SQLState:"HY000", Message:"Cannot proceed because system tables used by Event Scheduler were found damaged at server start", Description:"To address this issue, try running mysql_upgrade. ", MySQLVersion:"5.7"},
    1578: definition.ErrorDefinition{ErrorNumber:1578, ErrorType:"ServerError", Symbol:"ER_ONLY_INTEGERS_ALLOWED", SQLState:"HY000", Message:"Only integers allowed as number here ", Description:"", MySQLVersion:"5.7"},
    1579: definition.ErrorDefinition{ErrorNumber:1579, ErrorType:"ServerError", Symbol:"ER_UNSUPORTED_LOG_ENGINE", SQLState:"HY000", Message:"This storage engine cannot be used for log tables\" ", Description:"", MySQLVersion:"5.7"},
    1580: definition.ErrorDefinition{ErrorNumber:1580, ErrorType:"ServerError", Symbol:"ER_BAD_LOG_STATEMENT", SQLState:"HY000", Message:"You cannot '%s' a log table if logging is enabled ", Description:"", MySQLVersion:"5.7"},
    1581: definition.ErrorDefinition{ErrorNumber:1581, ErrorType:"ServerError", Symbol:"ER_CANT_RENAME_LOG_TABLE", SQLState:"HY000", Message:"Cannot rename '%s'. When logging enabled, rename to/from log table must rename two tables: the log table to an archive table and another table back to '%s' ", Description:"", MySQLVersion:"5.7"},
    1582: definition.ErrorDefinition{ErrorNumber:1582, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT", SQLState:"42000", Message:"Incorrect parameter count in the call to native function '%s' ", Description:"", MySQLVersion:"5.7"},
    1583: definition.ErrorDefinition{ErrorNumber:1583, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMETERS_TO_NATIVE_FCT", SQLState:"42000", Message:"Incorrect parameters in the call to native function '%s' ", Description:"", MySQLVersion:"5.7"},
    1584: definition.ErrorDefinition{ErrorNumber:1584, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMETERS_TO_STORED_FCT", SQLState:"42000", Message:"Incorrect parameters in the call to stored function %s ", Description:"", MySQLVersion:"5.7"},
    1585: definition.ErrorDefinition{ErrorNumber:1585, ErrorType:"ServerError", Symbol:"ER_NATIVE_FCT_NAME_COLLISION", SQLState:"HY000", Message:"This function '%s' has the same name as a native function ", Description:"", MySQLVersion:"5.7"},
    1586: definition.ErrorDefinition{ErrorNumber:1586, ErrorType:"ServerError", Symbol:"ER_DUP_ENTRY_WITH_KEY_NAME", SQLState:"23000", Message:"Duplicate entry '%s' for key '%s'", Description:"The format string for this error is also used with ER_DUP_ENTRY. ", MySQLVersion:"5.7"},
    1587: definition.ErrorDefinition{ErrorNumber:1587, ErrorType:"ServerError", Symbol:"ER_BINLOG_PURGE_EMFILE", SQLState:"HY000", Message:"Too many files opened, please execute the command again ", Description:"", MySQLVersion:"5.7"},
    1588: definition.ErrorDefinition{ErrorNumber:1588, ErrorType:"ServerError", Symbol:"ER_EVENT_CANNOT_CREATE_IN_THE_PAST", SQLState:"HY000", Message:"Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was dropped immediately after creation. ", Description:"", MySQLVersion:"5.7"},
    1589: definition.ErrorDefinition{ErrorNumber:1589, ErrorType:"ServerError", Symbol:"ER_EVENT_CANNOT_ALTER_IN_THE_PAST", SQLState:"HY000", Message:"Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was not changed. Specify a time in the future. ", Description:"", MySQLVersion:"5.7"},
    1590: definition.ErrorDefinition{ErrorNumber:1590, ErrorType:"ServerError", Symbol:"ER_SLAVE_INCIDENT", SQLState:"HY000", Message:"The incident %s occured on the master. Message: %s ", Description:"", MySQLVersion:"5.7"},
    1591: definition.ErrorDefinition{ErrorNumber:1591, ErrorType:"ServerError", Symbol:"ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT", SQLState:"HY000", Message:"Table has no partition for some existing values ", Description:"", MySQLVersion:"5.7"},
    1592: definition.ErrorDefinition{ErrorNumber:1592, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_STATEMENT", SQLState:"HY000", Message:"Unsafe statement written to the binary log using statement format since BINLOG_FORMAT = STATEMENT. %s ", Description:"", MySQLVersion:"5.7"},
    1593: definition.ErrorDefinition{ErrorNumber:1593, ErrorType:"ServerError", Symbol:"ER_SLAVE_FATAL_ERROR", SQLState:"HY000", Message:"Fatal error: %s ", Description:"", MySQLVersion:"5.7"},
    1594: definition.ErrorDefinition{ErrorNumber:1594, ErrorType:"ServerError", Symbol:"ER_SLAVE_RELAY_LOG_READ_FAILURE", SQLState:"HY000", Message:"Relay log read failure: %s ", Description:"", MySQLVersion:"5.7"},
    1595: definition.ErrorDefinition{ErrorNumber:1595, ErrorType:"ServerError", Symbol:"ER_SLAVE_RELAY_LOG_WRITE_FAILURE", SQLState:"HY000", Message:"Relay log write failure: %s ", Description:"", MySQLVersion:"5.7"},
    1596: definition.ErrorDefinition{ErrorNumber:1596, ErrorType:"ServerError", Symbol:"ER_SLAVE_CREATE_EVENT_FAILURE", SQLState:"HY000", Message:"Failed to create %s ", Description:"", MySQLVersion:"5.7"},
    1597: definition.ErrorDefinition{ErrorNumber:1597, ErrorType:"ServerError", Symbol:"ER_SLAVE_MASTER_COM_FAILURE", SQLState:"HY000", Message:"Master command %s failed: %s ", Description:"", MySQLVersion:"5.7"},
    1598: definition.ErrorDefinition{ErrorNumber:1598, ErrorType:"ServerError", Symbol:"ER_BINLOG_LOGGING_IMPOSSIBLE", SQLState:"HY000", Message:"Binary logging not possible. Message: %s ", Description:"", MySQLVersion:"5.7"},
    1599: definition.ErrorDefinition{ErrorNumber:1599, ErrorType:"ServerError", Symbol:"ER_VIEW_NO_CREATION_CTX", SQLState:"HY000", Message:"View `%s`.`%s` has no creation context ", Description:"", MySQLVersion:"5.7"},
    1600: definition.ErrorDefinition{ErrorNumber:1600, ErrorType:"ServerError", Symbol:"ER_VIEW_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Creation context of view `%s`.`%s' is invalid ", Description:"", MySQLVersion:"5.7"},
    1601: definition.ErrorDefinition{ErrorNumber:1601, ErrorType:"ServerError", Symbol:"ER_SR_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Creation context of stored routine `%s`.`%s` is invalid ", Description:"", MySQLVersion:"5.7"},
    1602: definition.ErrorDefinition{ErrorNumber:1602, ErrorType:"ServerError", Symbol:"ER_TRG_CORRUPTED_FILE", SQLState:"HY000", Message:"Corrupted TRG file for table `%s`.`%s` ", Description:"", MySQLVersion:"5.7"},
    1603: definition.ErrorDefinition{ErrorNumber:1603, ErrorType:"ServerError", Symbol:"ER_TRG_NO_CREATION_CTX", SQLState:"HY000", Message:"Triggers for table `%s`.`%s` have no creation context ", Description:"", MySQLVersion:"5.7"},
    1604: definition.ErrorDefinition{ErrorNumber:1604, ErrorType:"ServerError", Symbol:"ER_TRG_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Trigger creation context of table `%s`.`%s` is invalid ", Description:"", MySQLVersion:"5.7"},
    1605: definition.ErrorDefinition{ErrorNumber:1605, ErrorType:"ServerError", Symbol:"ER_EVENT_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Creation context of event `%s`.`%s` is invalid ", Description:"", MySQLVersion:"5.7"},
    1606: definition.ErrorDefinition{ErrorNumber:1606, ErrorType:"ServerError", Symbol:"ER_TRG_CANT_OPEN_TABLE", SQLState:"HY000", Message:"Cannot open table for trigger `%s`.`%s` ", Description:"", MySQLVersion:"5.7"},
    1607: definition.ErrorDefinition{ErrorNumber:1607, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_SROUTINE", SQLState:"HY000", Message:"Cannot create stored routine `%s`. Check warnings ", Description:"", MySQLVersion:"5.7"},
    1608: definition.ErrorDefinition{ErrorNumber:1608, ErrorType:"ServerError", Symbol:"ER_NEVER_USED", SQLState:"HY000", Message:"Ambiguous slave modes combination. %s ", Description:"", MySQLVersion:"5.7"},
    1609: definition.ErrorDefinition{ErrorNumber:1609, ErrorType:"ServerError", Symbol:"ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT", SQLState:"HY000", Message:"The BINLOG statement of type `%s` was not preceded by a format description BINLOG statement. ", Description:"", MySQLVersion:"5.7"},
    1610: definition.ErrorDefinition{ErrorNumber:1610, ErrorType:"ServerError", Symbol:"ER_SLAVE_CORRUPT_EVENT", SQLState:"HY000", Message:"Corrupted replication event was detected ", Description:"", MySQLVersion:"5.7"},
    1611: definition.ErrorDefinition{ErrorNumber:1611, ErrorType:"ServerError", Symbol:"ER_LOAD_DATA_INVALID_COLUMN", SQLState:"HY000", Message:"Invalid column reference (%s) in LOAD DATA", Description:"ER_LOAD_DATA_INVALID_COLUMN was removed after 5.7.7. ", MySQLVersion:"5.7"},
    1611: definition.ErrorDefinition{ErrorNumber:1611, ErrorType:"ServerError", Symbol:"ER_LOAD_DATA_INVALID_COLUMN_UNUSED", SQLState:"HY000", Message:"Invalid column reference (%s) in LOAD DATA", Description:"ER_LOAD_DATA_INVALID_COLUMN_UNUSED was added in 5.7.8. ", MySQLVersion:"5.7"},
    1612: definition.ErrorDefinition{ErrorNumber:1612, ErrorType:"ServerError", Symbol:"ER_LOG_PURGE_NO_FILE", SQLState:"HY000", Message:"Being purged log %s was not found ", Description:"", MySQLVersion:"5.7"},
    1613: definition.ErrorDefinition{ErrorNumber:1613, ErrorType:"ServerError", Symbol:"ER_XA_RBTIMEOUT", SQLState:"XA106", Message:"XA_RBTIMEOUT: Transaction branch was rolled back: took too long ", Description:"", MySQLVersion:"5.7"},
    1614: definition.ErrorDefinition{ErrorNumber:1614, ErrorType:"ServerError", Symbol:"ER_XA_RBDEADLOCK", SQLState:"XA102", Message:"XA_RBDEADLOCK: Transaction branch was rolled back: deadlock was detected ", Description:"", MySQLVersion:"5.7"},
    1615: definition.ErrorDefinition{ErrorNumber:1615, ErrorType:"ServerError", Symbol:"ER_NEED_REPREPARE", SQLState:"HY000", Message:"Prepared statement needs to be re-prepared ", Description:"", MySQLVersion:"5.7"},
    1616: definition.ErrorDefinition{ErrorNumber:1616, ErrorType:"ServerError", Symbol:"ER_DELAYED_NOT_SUPPORTED", SQLState:"HY000", Message:"DELAYED option not supported for table '%s' ", Description:"", MySQLVersion:"5.7"},
    1617: definition.ErrorDefinition{ErrorNumber:1617, ErrorType:"ServerError", Symbol:"WARN_NO_MASTER_INFO", SQLState:"HY000", Message:"The master info structure does not exist ", Description:"", MySQLVersion:"5.7"},
    1618: definition.ErrorDefinition{ErrorNumber:1618, ErrorType:"ServerError", Symbol:"WARN_OPTION_IGNORED", SQLState:"HY000", Message:"<%s> option ignored ", Description:"", MySQLVersion:"5.7"},
    1619: definition.ErrorDefinition{ErrorNumber:1619, ErrorType:"ServerError", Symbol:"WARN_PLUGIN_DELETE_BUILTIN", SQLState:"HY000", Message:"Built-in plugins cannot be deleted", Description:"WARN_PLUGIN_DELETE_BUILTIN was removed after 5.7.4. ", MySQLVersion:"5.7"},
    1619: definition.ErrorDefinition{ErrorNumber:1619, ErrorType:"ServerError", Symbol:"ER_PLUGIN_DELETE_BUILTIN", SQLState:"HY000", Message:"Built-in plugins cannot be deleted", Description:"ER_PLUGIN_DELETE_BUILTIN was added in 5.7.5. ", MySQLVersion:"5.7"},
    1620: definition.ErrorDefinition{ErrorNumber:1620, ErrorType:"ServerError", Symbol:"WARN_PLUGIN_BUSY", SQLState:"HY000", Message:"Plugin is busy and will be uninstalled on shutdown ", Description:"", MySQLVersion:"5.7"},
    1621: definition.ErrorDefinition{ErrorNumber:1621, ErrorType:"ServerError", Symbol:"ER_VARIABLE_IS_READONLY", SQLState:"HY000", Message:"%s variable '%s' is read-only. Use SET %s to assign the value ", Description:"", MySQLVersion:"5.7"},
    1622: definition.ErrorDefinition{ErrorNumber:1622, ErrorType:"ServerError", Symbol:"ER_WARN_ENGINE_TRANSACTION_ROLLBACK", SQLState:"HY000", Message:"Storage engine %s does not support rollback for this statement. Transaction rolled back and must be restarted ", Description:"", MySQLVersion:"5.7"},
    1623: definition.ErrorDefinition{ErrorNumber:1623, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_FAILURE", SQLState:"HY000", Message:"Unexpected master's heartbeat data: %s ", Description:"", MySQLVersion:"5.7"},
    1624: definition.ErrorDefinition{ErrorNumber:1624, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE", SQLState:"HY000", Message:"The requested value for the heartbeat period is either negative or exceeds the maximum allowed (%s seconds). ", Description:"", MySQLVersion:"5.7"},
    1625: definition.ErrorDefinition{ErrorNumber:1625, ErrorType:"ServerError", Symbol:"ER_NDB_REPLICATION_SCHEMA_ERROR", SQLState:"HY000", Message:"Bad schema for mysql.ndb_replication table. Message: %s ", Description:"", MySQLVersion:"5.7"},
    1626: definition.ErrorDefinition{ErrorNumber:1626, ErrorType:"ServerError", Symbol:"ER_CONFLICT_FN_PARSE_ERROR", SQLState:"HY000", Message:"Error in parsing conflict function. Message: %s ", Description:"", MySQLVersion:"5.7"},
    1627: definition.ErrorDefinition{ErrorNumber:1627, ErrorType:"ServerError", Symbol:"ER_EXCEPTIONS_WRITE_ERROR", SQLState:"HY000", Message:"Write to exceptions table failed. Message: %s\" ", Description:"", MySQLVersion:"5.7"},
    1628: definition.ErrorDefinition{ErrorNumber:1628, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_TABLE_COMMENT", SQLState:"HY000", Message:"Comment for table '%s' is too long (max = %lu) ", Description:"", MySQLVersion:"5.7"},
    1629: definition.ErrorDefinition{ErrorNumber:1629, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_FIELD_COMMENT", SQLState:"HY000", Message:"Comment for field '%s' is too long (max = %lu) ", Description:"", MySQLVersion:"5.7"},
    1630: definition.ErrorDefinition{ErrorNumber:1630, ErrorType:"ServerError", Symbol:"ER_FUNC_INEXISTENT_NAME_COLLISION", SQLState:"42000", Message:"FUNCTION %s does not exist. Check the 'Function Name Parsing and Resolution' section in the Reference Manual ", Description:"", MySQLVersion:"5.7"},
    1631: definition.ErrorDefinition{ErrorNumber:1631, ErrorType:"ServerError", Symbol:"ER_DATABASE_NAME", SQLState:"HY000", Message:"Database ", Description:"", MySQLVersion:"5.7"},
    1632: definition.ErrorDefinition{ErrorNumber:1632, ErrorType:"ServerError", Symbol:"ER_TABLE_NAME", SQLState:"HY000", Message:"Table ", Description:"", MySQLVersion:"5.7"},
    1633: definition.ErrorDefinition{ErrorNumber:1633, ErrorType:"ServerError", Symbol:"ER_PARTITION_NAME", SQLState:"HY000", Message:"Partition ", Description:"", MySQLVersion:"5.7"},
    1634: definition.ErrorDefinition{ErrorNumber:1634, ErrorType:"ServerError", Symbol:"ER_SUBPARTITION_NAME", SQLState:"HY000", Message:"Subpartition ", Description:"", MySQLVersion:"5.7"},
    1635: definition.ErrorDefinition{ErrorNumber:1635, ErrorType:"ServerError", Symbol:"ER_TEMPORARY_NAME", SQLState:"HY000", Message:"Temporary ", Description:"", MySQLVersion:"5.7"},
    1636: definition.ErrorDefinition{ErrorNumber:1636, ErrorType:"ServerError", Symbol:"ER_RENAMED_NAME", SQLState:"HY000", Message:"Renamed ", Description:"", MySQLVersion:"5.7"},
    1637: definition.ErrorDefinition{ErrorNumber:1637, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_CONCURRENT_TRXS", SQLState:"HY000", Message:"Too many active concurrent transactions ", Description:"", MySQLVersion:"5.7"},
    1638: definition.ErrorDefinition{ErrorNumber:1638, ErrorType:"ServerError", Symbol:"WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED", SQLState:"HY000", Message:"Non-ASCII separator arguments are not fully supported ", Description:"", MySQLVersion:"5.7"},
    1639: definition.ErrorDefinition{ErrorNumber:1639, ErrorType:"ServerError", Symbol:"ER_DEBUG_SYNC_TIMEOUT", SQLState:"HY000", Message:"debug sync point wait timed out ", Description:"", MySQLVersion:"5.7"},
    1640: definition.ErrorDefinition{ErrorNumber:1640, ErrorType:"ServerError", Symbol:"ER_DEBUG_SYNC_HIT_LIMIT", SQLState:"HY000", Message:"debug sync point hit limit reached ", Description:"", MySQLVersion:"5.7"},
    1641: definition.ErrorDefinition{ErrorNumber:1641, ErrorType:"ServerError", Symbol:"ER_DUP_SIGNAL_SET", SQLState:"42000", Message:"Duplicate condition information item '%s' ", Description:"", MySQLVersion:"5.7"},
    1642: definition.ErrorDefinition{ErrorNumber:1642, ErrorType:"ServerError", Symbol:"ER_SIGNAL_WARN", SQLState:"01000", Message:"Unhandled user-defined warning condition ", Description:"", MySQLVersion:"5.7"},
    1643: definition.ErrorDefinition{ErrorNumber:1643, ErrorType:"ServerError", Symbol:"ER_SIGNAL_NOT_FOUND", SQLState:"02000", Message:"Unhandled user-defined not found condition ", Description:"", MySQLVersion:"5.7"},
    1644: definition.ErrorDefinition{ErrorNumber:1644, ErrorType:"ServerError", Symbol:"ER_SIGNAL_EXCEPTION", SQLState:"HY000", Message:"Unhandled user-defined exception condition ", Description:"", MySQLVersion:"5.7"},
    1645: definition.ErrorDefinition{ErrorNumber:1645, ErrorType:"ServerError", Symbol:"ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER", SQLState:"0K000", Message:"RESIGNAL when handler not active ", Description:"", MySQLVersion:"5.7"},
    1646: definition.ErrorDefinition{ErrorNumber:1646, ErrorType:"ServerError", Symbol:"ER_SIGNAL_BAD_CONDITION_TYPE", SQLState:"HY000", Message:"SIGNAL/RESIGNAL can only use a CONDITION defined with SQLSTATE ", Description:"", MySQLVersion:"5.7"},
    1647: definition.ErrorDefinition{ErrorNumber:1647, ErrorType:"ServerError", Symbol:"WARN_COND_ITEM_TRUNCATED", SQLState:"HY000", Message:"Data truncated for condition item '%s' ", Description:"", MySQLVersion:"5.7"},
    1648: definition.ErrorDefinition{ErrorNumber:1648, ErrorType:"ServerError", Symbol:"ER_COND_ITEM_TOO_LONG", SQLState:"HY000", Message:"Data too long for condition item '%s' ", Description:"", MySQLVersion:"5.7"},
    1649: definition.ErrorDefinition{ErrorNumber:1649, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_LOCALE", SQLState:"HY000", Message:"Unknown locale: '%s' ", Description:"", MySQLVersion:"5.7"},
    1650: definition.ErrorDefinition{ErrorNumber:1650, ErrorType:"ServerError", Symbol:"ER_SLAVE_IGNORE_SERVER_IDS", SQLState:"HY000", Message:"The requested server id %d clashes with the slave startup option --replicate-same-server-id ", Description:"", MySQLVersion:"5.7"},
    1651: definition.ErrorDefinition{ErrorNumber:1651, ErrorType:"ServerError", Symbol:"ER_QUERY_CACHE_DISABLED", SQLState:"HY000", Message:"Query cache is disabled", Description:"restart the server with query_cache_type=1 to enable it ", MySQLVersion:"5.7"},
    1652: definition.ErrorDefinition{ErrorNumber:1652, ErrorType:"ServerError", Symbol:"ER_SAME_NAME_PARTITION_FIELD", SQLState:"HY000", Message:"Duplicate partition field name '%s' ", Description:"", MySQLVersion:"5.7"},
    1653: definition.ErrorDefinition{ErrorNumber:1653, ErrorType:"ServerError", Symbol:"ER_PARTITION_COLUMN_LIST_ERROR", SQLState:"HY000", Message:"Inconsistency in usage of column lists for partitioning ", Description:"", MySQLVersion:"5.7"},
    1654: definition.ErrorDefinition{ErrorNumber:1654, ErrorType:"ServerError", Symbol:"ER_WRONG_TYPE_COLUMN_VALUE_ERROR", SQLState:"HY000", Message:"Partition column values of incorrect type ", Description:"", MySQLVersion:"5.7"},
    1655: definition.ErrorDefinition{ErrorNumber:1655, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR", SQLState:"HY000", Message:"Too many fields in '%s' ", Description:"", MySQLVersion:"5.7"},
    1656: definition.ErrorDefinition{ErrorNumber:1656, ErrorType:"ServerError", Symbol:"ER_MAXVALUE_IN_VALUES_IN", SQLState:"HY000", Message:"Cannot use MAXVALUE as value in VALUES IN ", Description:"", MySQLVersion:"5.7"},
    1657: definition.ErrorDefinition{ErrorNumber:1657, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_VALUES_ERROR", SQLState:"HY000", Message:"Cannot have more than one value for this type of %s partitioning ", Description:"", MySQLVersion:"5.7"},
    1658: definition.ErrorDefinition{ErrorNumber:1658, ErrorType:"ServerError", Symbol:"ER_ROW_SINGLE_PARTITION_FIELD_ERROR", SQLState:"HY000", Message:"Row expressions in VALUES IN only allowed for multi-field column partitioning ", Description:"", MySQLVersion:"5.7"},
    1659: definition.ErrorDefinition{ErrorNumber:1659, ErrorType:"ServerError", Symbol:"ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD", SQLState:"HY000", Message:"Field '%s' is of a not allowed type for this type of partitioning ", Description:"", MySQLVersion:"5.7"},
    1660: definition.ErrorDefinition{ErrorNumber:1660, ErrorType:"ServerError", Symbol:"ER_PARTITION_FIELDS_TOO_LONG", SQLState:"HY000", Message:"The total length of the partitioning fields is too large ", Description:"", MySQLVersion:"5.7"},
    1661: definition.ErrorDefinition{ErrorNumber:1661, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since both row-incapable engines and statement-incapable engines are involved. ", Description:"", MySQLVersion:"5.7"},
    1662: definition.ErrorDefinition{ErrorNumber:1662, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_MODE_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = ROW and at least one table uses a storage engine limited to statement-based logging. ", Description:"", MySQLVersion:"5.7"},
    1663: definition.ErrorDefinition{ErrorNumber:1663, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since statement is unsafe, storage engine is limited to statement-based logging, and BINLOG_FORMAT = MIXED. %s ", Description:"", MySQLVersion:"5.7"},
    1664: definition.ErrorDefinition{ErrorNumber:1664, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since statement is in row format and at least one table uses a storage engine limited to statement-based logging. ", Description:"", MySQLVersion:"5.7"},
    1665: definition.ErrorDefinition{ErrorNumber:1665, ErrorType:"ServerError", Symbol:"ER_BINLOG_STMT_MODE_AND_ROW_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT and at least one table uses a storage engine limited to row-based logging.%s ", Description:"", MySQLVersion:"5.7"},
    1666: definition.ErrorDefinition{ErrorNumber:1666, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_INJECTION_AND_STMT_MODE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since statement is in row format and BINLOG_FORMAT = STATEMENT. ", Description:"", MySQLVersion:"5.7"},
    1667: definition.ErrorDefinition{ErrorNumber:1667, ErrorType:"ServerError", Symbol:"ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since more than one engine is involved and at least one engine is self-logging. ", Description:"", MySQLVersion:"5.7"},
    1668: definition.ErrorDefinition{ErrorNumber:1668, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_LIMIT", SQLState:"HY000", Message:"The statement is unsafe because it uses a LIMIT clause. This is unsafe because the set of rows included cannot be predicted. ", Description:"", MySQLVersion:"5.7"},
    1669: definition.ErrorDefinition{ErrorNumber:1669, ErrorType:"ServerError", Symbol:"ER_UNUSED4", SQLState:"HY000", Message:"The statement is unsafe because it uses INSERT DELAYED. This is unsafe because the times when rows are inserted cannot be predicted. ", Description:"", MySQLVersion:"5.7"},
    1670: definition.ErrorDefinition{ErrorNumber:1670, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_SYSTEM_TABLE", SQLState:"HY000", Message:"The statement is unsafe because it uses the general log, slow query log, or performance_schema table(s). This is unsafe because system tables may differ on slaves. ", Description:"", MySQLVersion:"5.7"},
    1671: definition.ErrorDefinition{ErrorNumber:1671, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_AUTOINC_COLUMNS", SQLState:"HY000", Message:"Statement is unsafe because it invokes a trigger or a stored function that inserts into an AUTO_INCREMENT column. Inserted values cannot be logged correctly. ", Description:"", MySQLVersion:"5.7"},
    1672: definition.ErrorDefinition{ErrorNumber:1672, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_UDF", SQLState:"HY000", Message:"Statement is unsafe because it uses a UDF which may not return the same value on the slave. ", Description:"", MySQLVersion:"5.7"},
    1673: definition.ErrorDefinition{ErrorNumber:1673, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_SYSTEM_VARIABLE", SQLState:"HY000", Message:"Statement is unsafe because it uses a system variable that may have a different value on the slave. ", Description:"", MySQLVersion:"5.7"},
    1674: definition.ErrorDefinition{ErrorNumber:1674, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_SYSTEM_FUNCTION", SQLState:"HY000", Message:"Statement is unsafe because it uses a system function that may return a different value on the slave. ", Description:"", MySQLVersion:"5.7"},
    1675: definition.ErrorDefinition{ErrorNumber:1675, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS", SQLState:"HY000", Message:"Statement is unsafe because it accesses a non-transactional table after accessing a transactional table within the same transaction. ", Description:"", MySQLVersion:"5.7"},
    1676: definition.ErrorDefinition{ErrorNumber:1676, ErrorType:"ServerError", Symbol:"ER_MESSAGE_AND_STATEMENT", SQLState:"HY000", Message:"%s Statement: %s ", Description:"", MySQLVersion:"5.7"},
    1677: definition.ErrorDefinition{ErrorNumber:1677, ErrorType:"ServerError", Symbol:"ER_SLAVE_CONVERSION_FAILED", SQLState:"HY000", Message:"Column %d of table '%s.%s' cannot be converted from type '%s' to type '%s' ", Description:"", MySQLVersion:"5.7"},
    1678: definition.ErrorDefinition{ErrorNumber:1678, ErrorType:"ServerError", Symbol:"ER_SLAVE_CANT_CREATE_CONVERSION", SQLState:"HY000", Message:"Can't create conversion table for table '%s.%s' ", Description:"", MySQLVersion:"5.7"},
    1679: definition.ErrorDefinition{ErrorNumber:1679, ErrorType:"ServerError", Symbol:"ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT", SQLState:"HY000", Message:"Cannot modify @@session.binlog_format inside a transaction ", Description:"", MySQLVersion:"5.7"},
    1680: definition.ErrorDefinition{ErrorNumber:1680, ErrorType:"ServerError", Symbol:"ER_PATH_LENGTH", SQLState:"HY000", Message:"The path specified for %s is too long. ", Description:"", MySQLVersion:"5.7"},
    1681: definition.ErrorDefinition{ErrorNumber:1681, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT", SQLState:"HY000", Message:"'%s' is deprecated and will be removed in a future release. ", Description:"", MySQLVersion:"5.7"},
    1682: definition.ErrorDefinition{ErrorNumber:1682, ErrorType:"ServerError", Symbol:"ER_WRONG_NATIVE_TABLE_STRUCTURE", SQLState:"HY000", Message:"Native table '%s'.'%s' has the wrong structure ", Description:"", MySQLVersion:"5.7"},
    1683: definition.ErrorDefinition{ErrorNumber:1683, ErrorType:"ServerError", Symbol:"ER_WRONG_PERFSCHEMA_USAGE", SQLState:"HY000", Message:"Invalid performance_schema usage. ", Description:"", MySQLVersion:"5.7"},
    1684: definition.ErrorDefinition{ErrorNumber:1684, ErrorType:"ServerError", Symbol:"ER_WARN_I_S_SKIPPED_TABLE", SQLState:"HY000", Message:"Table '%s'.'%s' was skipped since its definition is being modified by concurrent DDL statement ", Description:"", MySQLVersion:"5.7"},
    1685: definition.ErrorDefinition{ErrorNumber:1685, ErrorType:"ServerError", Symbol:"ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT", SQLState:"HY000", Message:"Cannot modify @@session.binlog_direct_non_transactional_updates inside a transaction ", Description:"", MySQLVersion:"5.7"},
    1686: definition.ErrorDefinition{ErrorNumber:1686, ErrorType:"ServerError", Symbol:"ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT", SQLState:"HY000", Message:"Cannot change the binlog direct flag inside a stored function or trigger ", Description:"", MySQLVersion:"5.7"},
    1687: definition.ErrorDefinition{ErrorNumber:1687, ErrorType:"ServerError", Symbol:"ER_SPATIAL_MUST_HAVE_GEOM_COL", SQLState:"42000", Message:"A SPATIAL index may only contain a geometrical type column ", Description:"", MySQLVersion:"5.7"},
    1688: definition.ErrorDefinition{ErrorNumber:1688, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_INDEX_COMMENT", SQLState:"HY000", Message:"Comment for index '%s' is too long (max = %lu) ", Description:"", MySQLVersion:"5.7"},
    1689: definition.ErrorDefinition{ErrorNumber:1689, ErrorType:"ServerError", Symbol:"ER_LOCK_ABORTED", SQLState:"HY000", Message:"Wait on a lock was aborted due to a pending exclusive lock ", Description:"", MySQLVersion:"5.7"},
    1690: definition.ErrorDefinition{ErrorNumber:1690, ErrorType:"ServerError", Symbol:"ER_DATA_OUT_OF_RANGE", SQLState:"22003", Message:"%s value is out of range in '%s' ", Description:"", MySQLVersion:"5.7"},
    1691: definition.ErrorDefinition{ErrorNumber:1691, ErrorType:"ServerError", Symbol:"ER_WRONG_SPVAR_TYPE_IN_LIMIT", SQLState:"HY000", Message:"A variable of a non-integer based type in LIMIT clause ", Description:"", MySQLVersion:"5.7"},
    1692: definition.ErrorDefinition{ErrorNumber:1692, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE", SQLState:"HY000", Message:"Mixing self-logging and non-self-logging engines in a statement is unsafe. ", Description:"", MySQLVersion:"5.7"},
    1693: definition.ErrorDefinition{ErrorNumber:1693, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_MIXED_STATEMENT", SQLState:"HY000", Message:"Statement accesses nontransactional table as well as transactional or temporary table, and writes to any of them. ", Description:"", MySQLVersion:"5.7"},
    1694: definition.ErrorDefinition{ErrorNumber:1694, ErrorType:"ServerError", Symbol:"ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN", SQLState:"HY000", Message:"Cannot modify @@session.sql_log_bin inside a transaction ", Description:"", MySQLVersion:"5.7"},
    1695: definition.ErrorDefinition{ErrorNumber:1695, ErrorType:"ServerError", Symbol:"ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN", SQLState:"HY000", Message:"Cannot change the sql_log_bin inside a stored function or trigger ", Description:"", MySQLVersion:"5.7"},
    1696: definition.ErrorDefinition{ErrorNumber:1696, ErrorType:"ServerError", Symbol:"ER_FAILED_READ_FROM_PAR_FILE", SQLState:"HY000", Message:"Failed to read from the .par file ", Description:"", MySQLVersion:"5.7"},
    1697: definition.ErrorDefinition{ErrorNumber:1697, ErrorType:"ServerError", Symbol:"ER_VALUES_IS_NOT_INT_TYPE_ERROR", SQLState:"HY000", Message:"VALUES value for partition '%s' must have type INT ", Description:"", MySQLVersion:"5.7"},
    1698: definition.ErrorDefinition{ErrorNumber:1698, ErrorType:"ServerError", Symbol:"ER_ACCESS_DENIED_NO_PASSWORD_ERROR", SQLState:"28000", Message:"Access denied for user '%s'@'%s' ", Description:"", MySQLVersion:"5.7"},
    1699: definition.ErrorDefinition{ErrorNumber:1699, ErrorType:"ServerError", Symbol:"ER_SET_PASSWORD_AUTH_PLUGIN", SQLState:"HY000", Message:"SET PASSWORD has no significance for users authenticating via plugins ", Description:"", MySQLVersion:"5.7"},
    1700: definition.ErrorDefinition{ErrorNumber:1700, ErrorType:"ServerError", Symbol:"ER_GRANT_PLUGIN_USER_EXISTS", SQLState:"HY000", Message:"GRANT with IDENTIFIED WITH is illegal because the user %-.*s already exists ", Description:"", MySQLVersion:"5.7"},
    1701: definition.ErrorDefinition{ErrorNumber:1701, ErrorType:"ServerError", Symbol:"ER_TRUNCATE_ILLEGAL_FK", SQLState:"42000", Message:"Cannot truncate a table referenced in a foreign key constraint (%s) ", Description:"", MySQLVersion:"5.7"},
    1702: definition.ErrorDefinition{ErrorNumber:1702, ErrorType:"ServerError", Symbol:"ER_PLUGIN_IS_PERMANENT", SQLState:"HY000", Message:"Plugin '%s' is force_plus_permanent and can not be unloaded ", Description:"", MySQLVersion:"5.7"},
    1703: definition.ErrorDefinition{ErrorNumber:1703, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN", SQLState:"HY000", Message:"The requested value for the heartbeat period is less than 1 millisecond. The value is reset to 0, meaning that heartbeating will effectively be disabled. ", Description:"", MySQLVersion:"5.7"},
    1704: definition.ErrorDefinition{ErrorNumber:1704, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX", SQLState:"HY000", Message:"The requested value for the heartbeat period exceeds the value of `slave_net_timeout' seconds. A sensible value for the period should be less than the timeout. ", Description:"", MySQLVersion:"5.7"},
    1705: definition.ErrorDefinition{ErrorNumber:1705, ErrorType:"ServerError", Symbol:"ER_STMT_CACHE_FULL", SQLState:"HY000", Message:"Multi-row statements required more than 'max_binlog_stmt_cache_size' bytes of storage", Description:"increase this mysqld variable and try again ", MySQLVersion:"5.7"},
    1706: definition.ErrorDefinition{ErrorNumber:1706, ErrorType:"ServerError", Symbol:"ER_MULTI_UPDATE_KEY_CONFLICT", SQLState:"HY000", Message:"Primary key/partition key update is not allowed since the table is updated both as '%s' and '%s'. ", Description:"", MySQLVersion:"5.7"},
    1707: definition.ErrorDefinition{ErrorNumber:1707, ErrorType:"ServerError", Symbol:"ER_TABLE_NEEDS_REBUILD", SQLState:"HY000", Message:"Table rebuild required. Please do \"ALTER TABLE `%s` FORCE\" or dump/reload to fix it! ", Description:"", MySQLVersion:"5.7"},
    1708: definition.ErrorDefinition{ErrorNumber:1708, ErrorType:"ServerError", Symbol:"WARN_OPTION_BELOW_LIMIT", SQLState:"HY000", Message:"The value of '%s' should be no less than the value of '%s' ", Description:"", MySQLVersion:"5.7"},
    1709: definition.ErrorDefinition{ErrorNumber:1709, ErrorType:"ServerError", Symbol:"ER_INDEX_COLUMN_TOO_LONG", SQLState:"HY000", Message:"Index column size too large. The maximum column size is %lu bytes. ", Description:"", MySQLVersion:"5.7"},
    1710: definition.ErrorDefinition{ErrorNumber:1710, ErrorType:"ServerError", Symbol:"ER_ERROR_IN_TRIGGER_BODY", SQLState:"HY000", Message:"Trigger '%s' has an error in its body: '%s' ", Description:"", MySQLVersion:"5.7"},
    1711: definition.ErrorDefinition{ErrorNumber:1711, ErrorType:"ServerError", Symbol:"ER_ERROR_IN_UNKNOWN_TRIGGER_BODY", SQLState:"HY000", Message:"Unknown trigger has an error in its body: '%s' ", Description:"", MySQLVersion:"5.7"},
    1712: definition.ErrorDefinition{ErrorNumber:1712, ErrorType:"ServerError", Symbol:"ER_INDEX_CORRUPT", SQLState:"HY000", Message:"Index %s is corrupted ", Description:"", MySQLVersion:"5.7"},
    1713: definition.ErrorDefinition{ErrorNumber:1713, ErrorType:"ServerError", Symbol:"ER_UNDO_RECORD_TOO_BIG", SQLState:"HY000", Message:"Undo log record is too big. ", Description:"", MySQLVersion:"5.7"},
    1714: definition.ErrorDefinition{ErrorNumber:1714, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT", SQLState:"HY000", Message:"INSERT IGNORE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1715: definition.ErrorDefinition{ErrorNumber:1715, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE", SQLState:"HY000", Message:"INSERT... SELECT... ON DUPLICATE KEY UPDATE is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are updated. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1716: definition.ErrorDefinition{ErrorNumber:1716, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_REPLACE_SELECT", SQLState:"HY000", Message:"REPLACE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1717: definition.ErrorDefinition{ErrorNumber:1717, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT", SQLState:"HY000", Message:"CREATE... IGNORE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1718: definition.ErrorDefinition{ErrorNumber:1718, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT", SQLState:"HY000", Message:"CREATE... REPLACE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1719: definition.ErrorDefinition{ErrorNumber:1719, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_UPDATE_IGNORE", SQLState:"HY000", Message:"UPDATE IGNORE is unsafe because the order in which rows are updated determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1720: definition.ErrorDefinition{ErrorNumber:1720, ErrorType:"ServerError", Symbol:"ER_PLUGIN_NO_UNINSTALL", SQLState:"HY000", Message:"Plugin '%s' is marked as not dynamically uninstallable. You have to stop the server to uninstall it. ", Description:"", MySQLVersion:"5.7"},
    1721: definition.ErrorDefinition{ErrorNumber:1721, ErrorType:"ServerError", Symbol:"ER_PLUGIN_NO_INSTALL", SQLState:"HY000", Message:"Plugin '%s' is marked as not dynamically installable. You have to stop the server to install it. ", Description:"", MySQLVersion:"5.7"},
    1722: definition.ErrorDefinition{ErrorNumber:1722, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT", SQLState:"HY000", Message:"Statements writing to a table with an auto-increment column after selecting from another table are unsafe because the order in which rows are retrieved determines what (if any) rows will be written. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1723: definition.ErrorDefinition{ErrorNumber:1723, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC", SQLState:"HY000", Message:"CREATE TABLE... SELECT... on a table with an auto-increment column is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are inserted. This order cannot be predicted and may differ on master and the slave. ", Description:"", MySQLVersion:"5.7"},
    1724: definition.ErrorDefinition{ErrorNumber:1724, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_TWO_KEYS", SQLState:"HY000", Message:"INSERT... ON DUPLICATE KEY UPDATE on a table with more than one UNIQUE KEY is unsafe ", Description:"", MySQLVersion:"5.7"},
    1725: definition.ErrorDefinition{ErrorNumber:1725, ErrorType:"ServerError", Symbol:"ER_TABLE_IN_FK_CHECK", SQLState:"HY000", Message:"Table is being used in foreign key check. ", Description:"", MySQLVersion:"5.7"},
    1726: definition.ErrorDefinition{ErrorNumber:1726, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_ENGINE", SQLState:"HY000", Message:"Storage engine '%s' does not support system tables. [%s.%s] ", Description:"", MySQLVersion:"5.7"},
    1727: definition.ErrorDefinition{ErrorNumber:1727, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST", SQLState:"HY000", Message:"INSERT into autoincrement field which is not the first part in the composed primary key is unsafe. ", Description:"", MySQLVersion:"5.7"},
    1728: definition.ErrorDefinition{ErrorNumber:1728, ErrorType:"ServerError", Symbol:"ER_CANNOT_LOAD_FROM_TABLE_V2", SQLState:"HY000", Message:"Cannot load from %s.%s. The table is probably corrupted ", Description:"", MySQLVersion:"5.7"},
    1729: definition.ErrorDefinition{ErrorNumber:1729, ErrorType:"ServerError", Symbol:"ER_MASTER_DELAY_VALUE_OUT_OF_RANGE", SQLState:"HY000", Message:"The requested value %s for the master delay exceeds the maximum %u ", Description:"", MySQLVersion:"5.7"},
    1730: definition.ErrorDefinition{ErrorNumber:1730, ErrorType:"ServerError", Symbol:"ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT", SQLState:"HY000", Message:"Only Format_description_log_event and row events are allowed in BINLOG statements (but %s was provided) ", Description:"", MySQLVersion:"5.7"},
    1731: definition.ErrorDefinition{ErrorNumber:1731, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_DIFFERENT_OPTION", SQLState:"HY000", Message:"Non matching attribute '%s' between partition and table ", Description:"", MySQLVersion:"5.7"},
    1732: definition.ErrorDefinition{ErrorNumber:1732, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_PART_TABLE", SQLState:"HY000", Message:"Table to exchange with partition is partitioned: '%s' ", Description:"", MySQLVersion:"5.7"},
    1733: definition.ErrorDefinition{ErrorNumber:1733, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_TEMP_TABLE", SQLState:"HY000", Message:"Table to exchange with partition is temporary: '%s' ", Description:"", MySQLVersion:"5.7"},
    1734: definition.ErrorDefinition{ErrorNumber:1734, ErrorType:"ServerError", Symbol:"ER_PARTITION_INSTEAD_OF_SUBPARTITION", SQLState:"HY000", Message:"Subpartitioned table, use subpartition instead of partition ", Description:"", MySQLVersion:"5.7"},
    1735: definition.ErrorDefinition{ErrorNumber:1735, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_PARTITION", SQLState:"HY000", Message:"Unknown partition '%s' in table '%s' ", Description:"", MySQLVersion:"5.7"},
    1736: definition.ErrorDefinition{ErrorNumber:1736, ErrorType:"ServerError", Symbol:"ER_TABLES_DIFFERENT_METADATA", SQLState:"HY000", Message:"Tables have different definitions ", Description:"", MySQLVersion:"5.7"},
    1737: definition.ErrorDefinition{ErrorNumber:1737, ErrorType:"ServerError", Symbol:"ER_ROW_DOES_NOT_MATCH_PARTITION", SQLState:"HY000", Message:"Found a row that does not match the partition ", Description:"", MySQLVersion:"5.7"},
    1738: definition.ErrorDefinition{ErrorNumber:1738, ErrorType:"ServerError", Symbol:"ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX", SQLState:"HY000", Message:"Option binlog_cache_size (%lu) is greater than max_binlog_cache_size (%lu)", Description:"setting binlog_cache_size equal to max_binlog_cache_size. ", MySQLVersion:"5.7"},
    1739: definition.ErrorDefinition{ErrorNumber:1739, ErrorType:"ServerError", Symbol:"ER_WARN_INDEX_NOT_APPLICABLE", SQLState:"HY000", Message:"Cannot use %s access on index '%s' due to type or collation conversion on field '%s' ", Description:"", MySQLVersion:"5.7"},
    1740: definition.ErrorDefinition{ErrorNumber:1740, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_FOREIGN_KEY", SQLState:"HY000", Message:"Table to exchange with partition has foreign key references: '%s' ", Description:"", MySQLVersion:"5.7"},
    1741: definition.ErrorDefinition{ErrorNumber:1741, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_KEY_VALUE", SQLState:"HY000", Message:"Key value '%s' was not found in table '%s.%s' ", Description:"", MySQLVersion:"5.7"},
    1742: definition.ErrorDefinition{ErrorNumber:1742, ErrorType:"ServerError", Symbol:"ER_RPL_INFO_DATA_TOO_LONG", SQLState:"HY000", Message:"Data for column '%s' too long ", Description:"", MySQLVersion:"5.7"},
    1743: definition.ErrorDefinition{ErrorNumber:1743, ErrorType:"ServerError", Symbol:"ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE", SQLState:"HY000", Message:"Replication event checksum verification failed while reading from network. ", Description:"", MySQLVersion:"5.7"},
    1744: definition.ErrorDefinition{ErrorNumber:1744, ErrorType:"ServerError", Symbol:"ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE", SQLState:"HY000", Message:"Replication event checksum verification failed while reading from a log file. ", Description:"", MySQLVersion:"5.7"},
    1745: definition.ErrorDefinition{ErrorNumber:1745, ErrorType:"ServerError", Symbol:"ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX", SQLState:"HY000", Message:"Option binlog_stmt_cache_size (%lu) is greater than max_binlog_stmt_cache_size (%lu)", Description:"setting binlog_stmt_cache_size equal to max_binlog_stmt_cache_size. ", MySQLVersion:"5.7"},
    1746: definition.ErrorDefinition{ErrorNumber:1746, ErrorType:"ServerError", Symbol:"ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT", SQLState:"HY000", Message:"Can't update table '%s' while '%s' is being created. ", Description:"", MySQLVersion:"5.7"},
    1747: definition.ErrorDefinition{ErrorNumber:1747, ErrorType:"ServerError", Symbol:"ER_PARTITION_CLAUSE_ON_NONPARTITIONED", SQLState:"HY000", Message:"PARTITION () clause on non partitioned table ", Description:"", MySQLVersion:"5.7"},
    1748: definition.ErrorDefinition{ErrorNumber:1748, ErrorType:"ServerError", Symbol:"ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET", SQLState:"HY000", Message:"Found a row not matching the given partition set ", Description:"", MySQLVersion:"5.7"},
    1749: definition.ErrorDefinition{ErrorNumber:1749, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_PARTITION__UNUSED", SQLState:"HY000", Message:"partition '%s' doesn't exist ", Description:"", MySQLVersion:"5.7"},
    1750: definition.ErrorDefinition{ErrorNumber:1750, ErrorType:"ServerError", Symbol:"ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE", SQLState:"HY000", Message:"Failure while changing the type of replication repository: %s. ", Description:"", MySQLVersion:"5.7"},
    1751: definition.ErrorDefinition{ErrorNumber:1751, ErrorType:"ServerError", Symbol:"ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE", SQLState:"HY000", Message:"The creation of some temporary tables could not be rolled back. ", Description:"", MySQLVersion:"5.7"},
    1752: definition.ErrorDefinition{ErrorNumber:1752, ErrorType:"ServerError", Symbol:"ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE", SQLState:"HY000", Message:"Some temporary tables were dropped, but these operations could not be rolled back. ", Description:"", MySQLVersion:"5.7"},
    1753: definition.ErrorDefinition{ErrorNumber:1753, ErrorType:"ServerError", Symbol:"ER_MTS_FEATURE_IS_NOT_SUPPORTED", SQLState:"HY000", Message:"%s is not supported in multi-threaded slave mode. %s ", Description:"", MySQLVersion:"5.7"},
    1754: definition.ErrorDefinition{ErrorNumber:1754, ErrorType:"ServerError", Symbol:"ER_MTS_UPDATED_DBS_GREATER_MAX", SQLState:"HY000", Message:"The number of modified databases exceeds the maximum %d", Description:"the database names will not be included in the replication event metadata. ", MySQLVersion:"5.7"},
    1755: definition.ErrorDefinition{ErrorNumber:1755, ErrorType:"ServerError", Symbol:"ER_MTS_CANT_PARALLEL", SQLState:"HY000", Message:"Cannot execute the current event group in the parallel mode. Encountered event %s, relay-log name %s, position %s which prevents execution of this event group in parallel mode. Reason: %s. ", Description:"", MySQLVersion:"5.7"},
    1756: definition.ErrorDefinition{ErrorNumber:1756, ErrorType:"ServerError", Symbol:"ER_MTS_INCONSISTENT_DATA", SQLState:"HY000", Message:"%s ", Description:"", MySQLVersion:"5.7"},
    1757: definition.ErrorDefinition{ErrorNumber:1757, ErrorType:"ServerError", Symbol:"ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING", SQLState:"HY000", Message:"FULLTEXT index is not supported for partitioned tables. ", Description:"", MySQLVersion:"5.7"},
    1758: definition.ErrorDefinition{ErrorNumber:1758, ErrorType:"ServerError", Symbol:"ER_DA_INVALID_CONDITION_NUMBER", SQLState:"35000", Message:"Invalid condition number ", Description:"", MySQLVersion:"5.7"},
    1759: definition.ErrorDefinition{ErrorNumber:1759, ErrorType:"ServerError", Symbol:"ER_INSECURE_PLAIN_TEXT", SQLState:"HY000", Message:"Sending passwords in plain text without SSL/TLS is extremely insecure. ", Description:"", MySQLVersion:"5.7"},
    1760: definition.ErrorDefinition{ErrorNumber:1760, ErrorType:"ServerError", Symbol:"ER_INSECURE_CHANGE_MASTER", SQLState:"HY000", Message:"Storing MySQL user name or password information in the master info repository is not secure and is therefore not recommended. Please consider using the USER and PASSWORD connection options for START SLAVE", Description:"see the 'START SLAVE Syntax' in the MySQL Manual for more information. ", MySQLVersion:"5.7"},
    1761: definition.ErrorDefinition{ErrorNumber:1761, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO", SQLState:"23000", Message:"Foreign key constraint for table '%s', record '%s' would lead to a duplicate entry in table '%s', key '%s' ", Description:"", MySQLVersion:"5.7"},
    1762: definition.ErrorDefinition{ErrorNumber:1762, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO", SQLState:"23000", Message:"Foreign key constraint for table '%s', record '%s' would lead to a duplicate entry in a child table ", Description:"", MySQLVersion:"5.7"},
    1763: definition.ErrorDefinition{ErrorNumber:1763, ErrorType:"ServerError", Symbol:"ER_SQLTHREAD_WITH_SECURE_SLAVE", SQLState:"HY000", Message:"Setting authentication options is not possible when only the Slave SQL Thread is being started. ", Description:"", MySQLVersion:"5.7"},
    1764: definition.ErrorDefinition{ErrorNumber:1764, ErrorType:"ServerError", Symbol:"ER_TABLE_HAS_NO_FT", SQLState:"HY000", Message:"The table does not have FULLTEXT index to support this query ", Description:"", MySQLVersion:"5.7"},
    1765: definition.ErrorDefinition{ErrorNumber:1765, ErrorType:"ServerError", Symbol:"ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER", SQLState:"HY000", Message:"The system variable %s cannot be set in stored functions or triggers. ", Description:"", MySQLVersion:"5.7"},
    1766: definition.ErrorDefinition{ErrorNumber:1766, ErrorType:"ServerError", Symbol:"ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION", SQLState:"HY000", Message:"The system variable %s cannot be set when there is an ongoing transaction. ", Description:"", MySQLVersion:"5.7"},
    1767: definition.ErrorDefinition{ErrorNumber:1767, ErrorType:"ServerError", Symbol:"ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT has the value %s, which is not listed in @@SESSION.GTID_NEXT_LIST. ", Description:"", MySQLVersion:"5.7"},
    1768: definition.ErrorDefinition{ErrorNumber:1768, ErrorType:"ServerError", Symbol:"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT cannot change inside a transaction.", Description:"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL was removed after 5.7.5. ", MySQLVersion:"5.7"},
    1768: definition.ErrorDefinition{ErrorNumber:1768, ErrorType:"ServerError", Symbol:"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT cannot change inside a transaction.", Description:"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION was added in 5.7.6. ", MySQLVersion:"5.7"},
    1769: definition.ErrorDefinition{ErrorNumber:1769, ErrorType:"ServerError", Symbol:"ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION", SQLState:"HY000", Message:"The statement 'SET %s' cannot invoke a stored function. ", Description:"", MySQLVersion:"5.7"},
    1770: definition.ErrorDefinition{ErrorNumber:1770, ErrorType:"ServerError", Symbol:"ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT cannot be 'AUTOMATIC' when @@SESSION.GTID_NEXT_LIST is non-NULL. ", Description:"", MySQLVersion:"5.7"},
    1771: definition.ErrorDefinition{ErrorNumber:1771, ErrorType:"ServerError", Symbol:"ER_SKIPPING_LOGGED_TRANSACTION", SQLState:"HY000", Message:"Skipping transaction %s because it has already been executed and logged. ", Description:"", MySQLVersion:"5.7"},
    1772: definition.ErrorDefinition{ErrorNumber:1772, ErrorType:"ServerError", Symbol:"ER_MALFORMED_GTID_SET_SPECIFICATION", SQLState:"HY000", Message:"Malformed GTID set specification '%s'. ", Description:"", MySQLVersion:"5.7"},
    1773: definition.ErrorDefinition{ErrorNumber:1773, ErrorType:"ServerError", Symbol:"ER_MALFORMED_GTID_SET_ENCODING", SQLState:"HY000", Message:"Malformed GTID set encoding. ", Description:"", MySQLVersion:"5.7"},
    1774: definition.ErrorDefinition{ErrorNumber:1774, ErrorType:"ServerError", Symbol:"ER_MALFORMED_GTID_SPECIFICATION", SQLState:"HY000", Message:"Malformed GTID specification '%s'. ", Description:"", MySQLVersion:"5.7"},
    1775: definition.ErrorDefinition{ErrorNumber:1775, ErrorType:"ServerError", Symbol:"ER_GNO_EXHAUSTED", SQLState:"HY000", Message:"Impossible to generate Global Transaction Identifier: the integer component reached the maximal value. Restart the server with a new server_uuid. ", Description:"", MySQLVersion:"5.7"},
    1776: definition.ErrorDefinition{ErrorNumber:1776, ErrorType:"ServerError", Symbol:"ER_BAD_SLAVE_AUTO_POSITION", SQLState:"HY000", Message:"Parameters MASTER_LOG_FILE, MASTER_LOG_POS, RELAY_LOG_FILE and RELAY_LOG_POS cannot be set when MASTER_AUTO_POSITION is active. ", Description:"", MySQLVersion:"5.7"},
    1777: definition.ErrorDefinition{ErrorNumber:1777, ErrorType:"ServerError", Symbol:"ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON", SQLState:"HY000", Message:"CHANGE MASTER TO MASTER_AUTO_POSITION = 1 can only be executed when @@GLOBAL.GTID_MODE = ON.", Description:"ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON was removed after 5.7.5. ", MySQLVersion:"5.7"},
    1777: definition.ErrorDefinition{ErrorNumber:1777, ErrorType:"ServerError", Symbol:"ER_AUTO_POSITION_REQUIRES_GTID_MODE_NOT_OFF", SQLState:"HY000", Message:"CHANGE MASTER TO MASTER_AUTO_POSITION = 1 cannot be executed because @@GLOBAL.GTID_MODE = OFF.", Description:"ER_AUTO_POSITION_REQUIRES_GTID_MODE_NOT_OFF was added in 5.7.6. ", MySQLVersion:"5.7"},
    1778: definition.ErrorDefinition{ErrorNumber:1778, ErrorType:"ServerError", Symbol:"ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET", SQLState:"HY000", Message:"Cannot execute statements with implicit commit inside a transaction when @@SESSION.GTID_NEXT == 'UUID:NUMBER'. ", Description:"", MySQLVersion:"5.7"},
    1779: definition.ErrorDefinition{ErrorNumber:1779, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON", SQLState:"HY000", Message:"@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_2 requires @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1.", Description:"ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON was removed after 5.7.5. ", MySQLVersion:"5.7"},
    1779: definition.ErrorDefinition{ErrorNumber:1779, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_ON_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON", SQLState:"HY000", Message:"GTID_MODE = ON requires ENFORCE_GTID_CONSISTENCY = ON.", Description:"ER_GTID_MODE_ON_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON was added in 5.7.6. ", MySQLVersion:"5.7"},
    1780: definition.ErrorDefinition{ErrorNumber:1780, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_REQUIRES_BINLOG", SQLState:"HY000", Message:"@@GLOBAL.GTID_MODE = ON or ON_PERMISSIVE or OFF_PERMISSIVE requires --log-bin and --log-slave-updates. ", Description:"", MySQLVersion:"5.7"},
    1781: definition.ErrorDefinition{ErrorNumber:1781, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT cannot be set to UUID:NUMBER when @@GLOBAL.GTID_MODE = OFF. ", Description:"", MySQLVersion:"5.7"},
    1782: definition.ErrorDefinition{ErrorNumber:1782, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT cannot be set to ANONYMOUS when @@GLOBAL.GTID_MODE = ON. ", Description:"", MySQLVersion:"5.7"},
    1783: definition.ErrorDefinition{ErrorNumber:1783, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT_LIST cannot be set to a non-NULL value when @@GLOBAL.GTID_MODE = OFF. ", Description:"", MySQLVersion:"5.7"},
    1784: definition.ErrorDefinition{ErrorNumber:1784, ErrorType:"ServerError", Symbol:"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"Found a Gtid_log_event or Previous_gtids_log_event when @@GLOBAL.GTID_MODE = OFF.", Description:"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF was removed after 5.7.5. ", MySQLVersion:"5.7"},
    1784: definition.ErrorDefinition{ErrorNumber:1784, ErrorType:"ServerError", Symbol:"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF__UNUSED", SQLState:"HY000", Message:"Found a Gtid_log_event when @@GLOBAL.GTID_MODE = OFF.", Description:"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF__UNUSED was added in 5.7.6. ", MySQLVersion:"5.7"},
    1785: definition.ErrorDefinition{ErrorNumber:1785, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE", SQLState:"HY000", Message:"Statement violates GTID consistency: Updates to non-transactional tables can only be done in either autocommitted statements or single-statement transactions, and never in the same statement as updates to transactional tables. ", Description:"", MySQLVersion:"5.7"},
    1786: definition.ErrorDefinition{ErrorNumber:1786, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_CREATE_SELECT", SQLState:"HY000", Message:"Statement violates GTID consistency: CREATE TABLE ... SELECT. ", Description:"", MySQLVersion:"5.7"},
    1787: definition.ErrorDefinition{ErrorNumber:1787, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTION", SQLState:"HY000", Message:"Statement violates GTID consistency: CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can only be executed outside transactional context. These statements are also not allowed in a function or trigger because functions and triggers are also considered to be multi-statement transactions. ", Description:"", MySQLVersion:"5.7"},
    1788: definition.ErrorDefinition{ErrorNumber:1788, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME", SQLState:"HY000", Message:"The value of @@GLOBAL.GTID_MODE can only be changed one step at a time: OFF <-> OFF_PERMISSIVE <-> ON_PERMISSIVE <-> ON. Also note that this value must be stepped up or down simultaneously on all servers. See the Manual for instructions. ", Description:"", MySQLVersion:"5.7"},
    1789: definition.ErrorDefinition{ErrorNumber:1789, ErrorType:"ServerError", Symbol:"ER_MASTER_HAS_PURGED_REQUIRED_GTIDS", SQLState:"HY000", Message:"The slave is connecting using CHANGE MASTER TO MASTER_AUTO_POSITION = 1, but the master has purged binary logs containing GTIDs that the slave requires. Replicate the missing transactions from elsewhere, or provision a new slave from backup. Consider increasing the master's binary log expiration period. %s. ", Description:"", MySQLVersion:"5.7"},
    1790: definition.ErrorDefinition{ErrorNumber:1790, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK. ", Description:"", MySQLVersion:"5.7"},
    1791: definition.ErrorDefinition{ErrorNumber:1791, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_EXPLAIN_FORMAT", SQLState:"HY000", Message:"Unknown EXPLAIN format name: '%s' ", Description:"", MySQLVersion:"5.7"},
    1792: definition.ErrorDefinition{ErrorNumber:1792, ErrorType:"ServerError", Symbol:"ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION", SQLState:"25006", Message:"Cannot execute statement in a READ ONLY transaction. ", Description:"", MySQLVersion:"5.7"},
    1793: definition.ErrorDefinition{ErrorNumber:1793, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_TABLE_PARTITION_COMMENT", SQLState:"HY000", Message:"Comment for table partition '%s' is too long (max = %lu) ", Description:"", MySQLVersion:"5.7"},
    1794: definition.ErrorDefinition{ErrorNumber:1794, ErrorType:"ServerError", Symbol:"ER_SLAVE_CONFIGURATION", SQLState:"HY000", Message:"Slave is not configured or failed to initialize properly. You must at least set --server-id to enable either a master or a slave. Additional error messages can be found in the MySQL error log. ", Description:"", MySQLVersion:"5.7"},
    1795: definition.ErrorDefinition{ErrorNumber:1795, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_LIMIT", SQLState:"HY000", Message:"InnoDB presently supports one FULLTEXT index creation at a time ", Description:"", MySQLVersion:"5.7"},
    1796: definition.ErrorDefinition{ErrorNumber:1796, ErrorType:"ServerError", Symbol:"ER_INNODB_NO_FT_TEMP_TABLE", SQLState:"HY000", Message:"Cannot create FULLTEXT index on temporary InnoDB table ", Description:"", MySQLVersion:"5.7"},
    1797: definition.ErrorDefinition{ErrorNumber:1797, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_WRONG_DOCID_COLUMN", SQLState:"HY000", Message:"Column '%s' is of wrong type for an InnoDB FULLTEXT index ", Description:"", MySQLVersion:"5.7"},
    1798: definition.ErrorDefinition{ErrorNumber:1798, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_WRONG_DOCID_INDEX", SQLState:"HY000", Message:"Index '%s' is of wrong type for an InnoDB FULLTEXT index ", Description:"", MySQLVersion:"5.7"},
    1799: definition.ErrorDefinition{ErrorNumber:1799, ErrorType:"ServerError", Symbol:"ER_INNODB_ONLINE_LOG_TOO_BIG", SQLState:"HY000", Message:"Creating index '%s' required more than 'innodb_online_alter_log_max_size' bytes of modification log. Please try again. ", Description:"", MySQLVersion:"5.7"},
    1800: definition.ErrorDefinition{ErrorNumber:1800, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_ALTER_ALGORITHM", SQLState:"HY000", Message:"Unknown ALGORITHM '%s' ", Description:"", MySQLVersion:"5.7"},
    1801: definition.ErrorDefinition{ErrorNumber:1801, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_ALTER_LOCK", SQLState:"HY000", Message:"Unknown LOCK type '%s' ", Description:"", MySQLVersion:"5.7"},
    1802: definition.ErrorDefinition{ErrorNumber:1802, ErrorType:"ServerError", Symbol:"ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS", SQLState:"HY000", Message:"CHANGE MASTER cannot be executed when the slave was stopped with an error or killed in MTS mode. Consider using RESET SLAVE or START SLAVE UNTIL. ", Description:"", MySQLVersion:"5.7"},
    1803: definition.ErrorDefinition{ErrorNumber:1803, ErrorType:"ServerError", Symbol:"ER_MTS_RECOVERY_FAILURE", SQLState:"HY000", Message:"Cannot recover after SLAVE errored out in parallel execution mode. Additional error messages can be found in the MySQL error log. ", Description:"", MySQLVersion:"5.7"},
    1804: definition.ErrorDefinition{ErrorNumber:1804, ErrorType:"ServerError", Symbol:"ER_MTS_RESET_WORKERS", SQLState:"HY000", Message:"Cannot clean up worker info tables. Additional error messages can be found in the MySQL error log. ", Description:"", MySQLVersion:"5.7"},
    1805: definition.ErrorDefinition{ErrorNumber:1805, ErrorType:"ServerError", Symbol:"ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2", SQLState:"HY000", Message:"Column count of %s.%s is wrong. Expected %d, found %d. The table is probably corrupted ", Description:"", MySQLVersion:"5.7"},
    1806: definition.ErrorDefinition{ErrorNumber:1806, ErrorType:"ServerError", Symbol:"ER_SLAVE_SILENT_RETRY_TRANSACTION", SQLState:"HY000", Message:"Slave must silently retry current transaction ", Description:"", MySQLVersion:"5.7"},
    1807: definition.ErrorDefinition{ErrorNumber:1807, ErrorType:"ServerError", Symbol:"ER_DISCARD_FK_CHECKS_RUNNING", SQLState:"HY000", Message:"There is a foreign key check running on table '%s'. Cannot discard the table. ", Description:"", MySQLVersion:"5.7"},
    1808: definition.ErrorDefinition{ErrorNumber:1808, ErrorType:"ServerError", Symbol:"ER_TABLE_SCHEMA_MISMATCH", SQLState:"HY000", Message:"Schema mismatch (%s) ", Description:"", MySQLVersion:"5.7"},
    1809: definition.ErrorDefinition{ErrorNumber:1809, ErrorType:"ServerError", Symbol:"ER_TABLE_IN_SYSTEM_TABLESPACE", SQLState:"HY000", Message:"Table '%s' in system tablespace ", Description:"", MySQLVersion:"5.7"},
    1810: definition.ErrorDefinition{ErrorNumber:1810, ErrorType:"ServerError", Symbol:"ER_IO_READ_ERROR", SQLState:"HY000", Message:"IO Read error: (%lu, %s) %s ", Description:"", MySQLVersion:"5.7"},
    1811: definition.ErrorDefinition{ErrorNumber:1811, ErrorType:"ServerError", Symbol:"ER_IO_WRITE_ERROR", SQLState:"HY000", Message:"IO Write error: (%lu, %s) %s ", Description:"", MySQLVersion:"5.7"},
    1812: definition.ErrorDefinition{ErrorNumber:1812, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_MISSING", SQLState:"HY000", Message:"Tablespace is missing for table %s. ", Description:"", MySQLVersion:"5.7"},
    1813: definition.ErrorDefinition{ErrorNumber:1813, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_EXISTS", SQLState:"HY000", Message:"Tablespace '%s' exists. ", Description:"", MySQLVersion:"5.7"},
    1814: definition.ErrorDefinition{ErrorNumber:1814, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_DISCARDED", SQLState:"HY000", Message:"Tablespace has been discarded for table '%s' ", Description:"", MySQLVersion:"5.7"},
    1815: definition.ErrorDefinition{ErrorNumber:1815, ErrorType:"ServerError", Symbol:"ER_INTERNAL_ERROR", SQLState:"HY000", Message:"Internal error: %s ", Description:"", MySQLVersion:"5.7"},
    1816: definition.ErrorDefinition{ErrorNumber:1816, ErrorType:"ServerError", Symbol:"ER_INNODB_IMPORT_ERROR", SQLState:"HY000", Message:"ALTER TABLE %s IMPORT TABLESPACE failed with error %lu : '%s' ", Description:"", MySQLVersion:"5.7"},
    1817: definition.ErrorDefinition{ErrorNumber:1817, ErrorType:"ServerError", Symbol:"ER_INNODB_INDEX_CORRUPT", SQLState:"HY000", Message:"Index corrupt: %s ", Description:"", MySQLVersion:"5.7"},
    1818: definition.ErrorDefinition{ErrorNumber:1818, ErrorType:"ServerError", Symbol:"ER_INVALID_YEAR_COLUMN_LENGTH", SQLState:"HY000", Message:"Supports only YEAR or YEAR(4) column. ", Description:"", MySQLVersion:"5.7"},
    1819: definition.ErrorDefinition{ErrorNumber:1819, ErrorType:"ServerError", Symbol:"ER_NOT_VALID_PASSWORD", SQLState:"HY000", Message:"Your password does not satisfy the current policy requirements ", Description:"", MySQLVersion:"5.7"},
    1820: definition.ErrorDefinition{ErrorNumber:1820, ErrorType:"ServerError", Symbol:"ER_MUST_CHANGE_PASSWORD", SQLState:"HY000", Message:"You must reset your password using ALTER USER statement before executing this statement. ", Description:"", MySQLVersion:"5.7"},
    1821: definition.ErrorDefinition{ErrorNumber:1821, ErrorType:"ServerError", Symbol:"ER_FK_NO_INDEX_CHILD", SQLState:"HY000", Message:"Failed to add the foreign key constaint. Missing index for constraint '%s' in the foreign table '%s' ", Description:"", MySQLVersion:"5.7"},
    1822: definition.ErrorDefinition{ErrorNumber:1822, ErrorType:"ServerError", Symbol:"ER_FK_NO_INDEX_PARENT", SQLState:"HY000", Message:"Failed to add the foreign key constaint. Missing index for constraint '%s' in the referenced table '%s' ", Description:"", MySQLVersion:"5.7"},
    1823: definition.ErrorDefinition{ErrorNumber:1823, ErrorType:"ServerError", Symbol:"ER_FK_FAIL_ADD_SYSTEM", SQLState:"HY000", Message:"Failed to add the foreign key constraint '%s' to system tables ", Description:"", MySQLVersion:"5.7"},
    1824: definition.ErrorDefinition{ErrorNumber:1824, ErrorType:"ServerError", Symbol:"ER_FK_CANNOT_OPEN_PARENT", SQLState:"HY000", Message:"Failed to open the referenced table '%s' ", Description:"", MySQLVersion:"5.7"},
    1825: definition.ErrorDefinition{ErrorNumber:1825, ErrorType:"ServerError", Symbol:"ER_FK_INCORRECT_OPTION", SQLState:"HY000", Message:"Failed to add the foreign key constraint on table '%s'. Incorrect options in FOREIGN KEY constraint '%s' ", Description:"", MySQLVersion:"5.7"},
    1826: definition.ErrorDefinition{ErrorNumber:1826, ErrorType:"ServerError", Symbol:"ER_FK_DUP_NAME", SQLState:"HY000", Message:"Duplicate foreign key constraint name '%s' ", Description:"", MySQLVersion:"5.7"},
    1827: definition.ErrorDefinition{ErrorNumber:1827, ErrorType:"ServerError", Symbol:"ER_PASSWORD_FORMAT", SQLState:"HY000", Message:"The password hash doesn't have the expected format. Check if the correct password algorithm is being used with the PASSWORD() function. ", Description:"", MySQLVersion:"5.7"},
    1828: definition.ErrorDefinition{ErrorNumber:1828, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_DROP", SQLState:"HY000", Message:"Cannot drop column '%s': needed in a foreign key constraint '%s' ", Description:"", MySQLVersion:"5.7"},
    1829: definition.ErrorDefinition{ErrorNumber:1829, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_DROP_CHILD", SQLState:"HY000", Message:"Cannot drop column '%s': needed in a foreign key constraint '%s' of table '%s' ", Description:"", MySQLVersion:"5.7"},
    1830: definition.ErrorDefinition{ErrorNumber:1830, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_NOT_NULL", SQLState:"HY000", Message:"Column '%s' cannot be NOT NULL: needed in a foreign key constraint '%s' SET NULL ", Description:"", MySQLVersion:"5.7"},
    1831: definition.ErrorDefinition{ErrorNumber:1831, ErrorType:"ServerError", Symbol:"ER_DUP_INDEX", SQLState:"HY000", Message:"Duplicate index '%s' defined on the table '%s.%s'. This is deprecated and will be disallowed in a future release. ", Description:"", MySQLVersion:"5.7"},
    1832: definition.ErrorDefinition{ErrorNumber:1832, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_CHANGE", SQLState:"HY000", Message:"Cannot change column '%s': used in a foreign key constraint '%s' ", Description:"", MySQLVersion:"5.7"},
    1833: definition.ErrorDefinition{ErrorNumber:1833, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_CHANGE_CHILD", SQLState:"HY000", Message:"Cannot change column '%s': used in a foreign key constraint '%s' of table '%s' ", Description:"", MySQLVersion:"5.7"},
    1834: definition.ErrorDefinition{ErrorNumber:1834, ErrorType:"ServerError", Symbol:"ER_FK_CANNOT_DELETE_PARENT", SQLState:"HY000", Message:"Cannot delete rows from table which is parent in a foreign key constraint '%s' of table '%s'", Description:"ER_FK_CANNOT_DELETE_PARENT was removed after 5.7.3. ", MySQLVersion:"5.7"},
    1834: definition.ErrorDefinition{ErrorNumber:1834, ErrorType:"ServerError", Symbol:"ER_UNUSED5", SQLState:"HY000", Message:"Cannot delete rows from table which is parent in a foreign key constraint '%s' of table '%s'", Description:"ER_UNUSED5 was added in 5.7.4. ", MySQLVersion:"5.7"},
    1835: definition.ErrorDefinition{ErrorNumber:1835, ErrorType:"ServerError", Symbol:"ER_MALFORMED_PACKET", SQLState:"HY000", Message:"Malformed communication packet. ", Description:"", MySQLVersion:"5.7"},
    1836: definition.ErrorDefinition{ErrorNumber:1836, ErrorType:"ServerError", Symbol:"ER_READ_ONLY_MODE", SQLState:"HY000", Message:"Running in read-only mode ", Description:"", MySQLVersion:"5.7"},
    1837: definition.ErrorDefinition{ErrorNumber:1837, ErrorType:"ServerError", Symbol:"ER_GTID_NEXT_TYPE_UNDEFINED_GROUP", SQLState:"HY000", Message:"When @@SESSION.GTID_NEXT is set to a GTID, you must explicitly set it to a different value after a COMMIT or ROLLBACK. Please check GTID_NEXT variable manual page for detailed explanation. Current @@SESSION.GTID_NEXT is '%s'. ", Description:"", MySQLVersion:"5.7"},
    1838: definition.ErrorDefinition{ErrorNumber:1838, ErrorType:"ServerError", Symbol:"ER_VARIABLE_NOT_SETTABLE_IN_SP", SQLState:"HY000", Message:"The system variable %s cannot be set in stored procedures. ", Description:"", MySQLVersion:"5.7"},
    1839: definition.ErrorDefinition{ErrorNumber:1839, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_MODE = ON. ", Description:"", MySQLVersion:"5.7"},
    1840: definition.ErrorDefinition{ErrorNumber:1840, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_EXECUTED is empty. ", Description:"", MySQLVersion:"5.7"},
    1841: definition.ErrorDefinition{ErrorNumber:1841, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED can only be set when there are no ongoing transactions (not even in other clients). ", Description:"", MySQLVersion:"5.7"},
    1842: definition.ErrorDefinition{ErrorNumber:1842, ErrorType:"ServerError", Symbol:"ER_GTID_PURGED_WAS_CHANGED", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'. ", Description:"", MySQLVersion:"5.7"},
    1843: definition.ErrorDefinition{ErrorNumber:1843, ErrorType:"ServerError", Symbol:"ER_GTID_EXECUTED_WAS_CHANGED", SQLState:"HY000", Message:"@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'. ", Description:"", MySQLVersion:"5.7"},
    1844: definition.ErrorDefinition{ErrorNumber:1844, ErrorType:"ServerError", Symbol:"ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT, and both replicated and non replicated tables are written to. ", Description:"", MySQLVersion:"5.7"},
    1845: definition.ErrorDefinition{ErrorNumber:1845, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED", SQLState:"0A000", Message:"%s is not supported for this operation. Try %s.", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED was added in 5.7.1. ", MySQLVersion:"5.7"},
    1846: definition.ErrorDefinition{ErrorNumber:1846, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON", SQLState:"0A000", Message:"%s is not supported. Reason: %s. Try %s.", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON was added in 5.7.1. ", MySQLVersion:"5.7"},
    1847: definition.ErrorDefinition{ErrorNumber:1847, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY", SQLState:"HY000", Message:"COPY algorithm requires a lock", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY was added in 5.7.1. ", MySQLVersion:"5.7"},
    1848: definition.ErrorDefinition{ErrorNumber:1848, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION", SQLState:"HY000", Message:"Partition specific operations do not yet support LOCK/ALGORITHM", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION was added in 5.7.1. ", MySQLVersion:"5.7"},
    1849: definition.ErrorDefinition{ErrorNumber:1849, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME", SQLState:"HY000", Message:"Columns participating in a foreign key are renamed", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME was added in 5.7.1. ", MySQLVersion:"5.7"},
    1850: definition.ErrorDefinition{ErrorNumber:1850, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE", SQLState:"HY000", Message:"Cannot change column type INPLACE", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE was added in 5.7.1. ", MySQLVersion:"5.7"},
    1851: definition.ErrorDefinition{ErrorNumber:1851, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK", SQLState:"HY000", Message:"Adding foreign keys needs foreign_key_checks=OFF", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK was added in 5.7.1. ", MySQLVersion:"5.7"},
    1852: definition.ErrorDefinition{ErrorNumber:1852, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE", SQLState:"HY000", Message:"Creating unique indexes with IGNORE requires COPY algorithm to remove duplicate rows", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE was added in 5.7.1, removed after 5.7.3. ", MySQLVersion:"5.7"},
    1852: definition.ErrorDefinition{ErrorNumber:1852, ErrorType:"ServerError", Symbol:"ER_UNUSED6", SQLState:"HY000", Message:"Creating unique indexes with IGNORE requires COPY algorithm to remove duplicate rows", Description:"ER_UNUSED6 was added in 5.7.4. ", MySQLVersion:"5.7"},
    1853: definition.ErrorDefinition{ErrorNumber:1853, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK", SQLState:"HY000", Message:"Dropping a primary key is not allowed without also adding a new primary key", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK was added in 5.7.1. ", MySQLVersion:"5.7"},
    1854: definition.ErrorDefinition{ErrorNumber:1854, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC", SQLState:"HY000", Message:"Adding an auto-increment column requires a lock", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC was added in 5.7.1. ", MySQLVersion:"5.7"},
    1855: definition.ErrorDefinition{ErrorNumber:1855, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS", SQLState:"HY000", Message:"Cannot replace hidden FTS_DOC_ID with a user-visible one", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS was added in 5.7.1. ", MySQLVersion:"5.7"},
    1856: definition.ErrorDefinition{ErrorNumber:1856, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS", SQLState:"HY000", Message:"Cannot drop or rename FTS_DOC_ID", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS was added in 5.7.1. ", MySQLVersion:"5.7"},
    1857: definition.ErrorDefinition{ErrorNumber:1857, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS", SQLState:"HY000", Message:"Fulltext index creation requires a lock", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS was added in 5.7.1. ", MySQLVersion:"5.7"},
    1858: definition.ErrorDefinition{ErrorNumber:1858, ErrorType:"ServerError", Symbol:"ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE", SQLState:"HY000", Message:"sql_slave_skip_counter can not be set when the server is running with @@GLOBAL.GTID_MODE = ON. Instead, for each transaction that you want to skip, generate an empty transaction with the same GTID as the transaction", Description:"ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE was added in 5.7.1. ", MySQLVersion:"5.7"},
    1859: definition.ErrorDefinition{ErrorNumber:1859, ErrorType:"ServerError", Symbol:"ER_DUP_UNKNOWN_IN_INDEX", SQLState:"23000", Message:"Duplicate entry for key '%s'", Description:"ER_DUP_UNKNOWN_IN_INDEX was added in 5.7.1. ", MySQLVersion:"5.7"},
    1860: definition.ErrorDefinition{ErrorNumber:1860, ErrorType:"ServerError", Symbol:"ER_IDENT_CAUSES_TOO_LONG_PATH", SQLState:"HY000", Message:"Long database name and identifier for object resulted in path length exceeding %d characters. Path: '%s'.", Description:"ER_IDENT_CAUSES_TOO_LONG_PATH was added in 5.7.1. ", MySQLVersion:"5.7"},
    1861: definition.ErrorDefinition{ErrorNumber:1861, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL", SQLState:"HY000", Message:"cannot silently convert NULL values, as required in this SQL_MODE", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL was added in 5.7.1. ", MySQLVersion:"5.7"},
    1862: definition.ErrorDefinition{ErrorNumber:1862, ErrorType:"ServerError", Symbol:"ER_MUST_CHANGE_PASSWORD_LOGIN", SQLState:"HY000", Message:"Your password has expired. To log in you must change it using a client that supports expired passwords.", Description:"ER_MUST_CHANGE_PASSWORD_LOGIN was added in 5.7.1. ", MySQLVersion:"5.7"},
    1863: definition.ErrorDefinition{ErrorNumber:1863, ErrorType:"ServerError", Symbol:"ER_ROW_IN_WRONG_PARTITION", SQLState:"HY000", Message:"Found a row in wrong partition %s", Description:"ER_ROW_IN_WRONG_PARTITION was added in 5.7.1. ", MySQLVersion:"5.7"},
    1864: definition.ErrorDefinition{ErrorNumber:1864, ErrorType:"ServerError", Symbol:"ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX", SQLState:"HY000", Message:"Cannot schedule event %s, relay-log name %s, position %s to Worker thread because its size %lu exceeds %lu of slave_pending_jobs_size_max.", Description:"ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX was added in 5.7.2. ", MySQLVersion:"5.7"},
    1865: definition.ErrorDefinition{ErrorNumber:1865, ErrorType:"ServerError", Symbol:"ER_INNODB_NO_FT_USES_PARSER", SQLState:"HY000", Message:"Cannot CREATE FULLTEXT INDEX WITH PARSER on InnoDB table", Description:"ER_INNODB_NO_FT_USES_PARSER was added in 5.7.2. ", MySQLVersion:"5.7"},
    1866: definition.ErrorDefinition{ErrorNumber:1866, ErrorType:"ServerError", Symbol:"ER_BINLOG_LOGICAL_CORRUPTION", SQLState:"HY000", Message:"The binary log file '%s' is logically corrupted: %s", Description:"ER_BINLOG_LOGICAL_CORRUPTION was added in 5.7.2. ", MySQLVersion:"5.7"},
    1867: definition.ErrorDefinition{ErrorNumber:1867, ErrorType:"ServerError", Symbol:"ER_WARN_PURGE_LOG_IN_USE", SQLState:"HY000", Message:"file %s was not purged because it was being read by %d thread(s), purged only %d out of %d files.", Description:"ER_WARN_PURGE_LOG_IN_USE was added in 5.7.2. ", MySQLVersion:"5.7"},
    1868: definition.ErrorDefinition{ErrorNumber:1868, ErrorType:"ServerError", Symbol:"ER_WARN_PURGE_LOG_IS_ACTIVE", SQLState:"HY000", Message:"file %s was not purged because it is the active log file.", Description:"ER_WARN_PURGE_LOG_IS_ACTIVE was added in 5.7.2. ", MySQLVersion:"5.7"},
    1869: definition.ErrorDefinition{ErrorNumber:1869, ErrorType:"ServerError", Symbol:"ER_AUTO_INCREMENT_CONFLICT", SQLState:"HY000", Message:"Auto-increment value in UPDATE conflicts with internally generated values", Description:"ER_AUTO_INCREMENT_CONFLICT was added in 5.7.2. ", MySQLVersion:"5.7"},
    1870: definition.ErrorDefinition{ErrorNumber:1870, ErrorType:"ServerError", Symbol:"WARN_ON_BLOCKHOLE_IN_RBR", SQLState:"HY000", Message:"Row events are not logged for %s statements that modify BLACKHOLE tables in row format. Table(s): '%s'", Description:"WARN_ON_BLOCKHOLE_IN_RBR was added in 5.7.2. ", MySQLVersion:"5.7"},
    1871: definition.ErrorDefinition{ErrorNumber:1871, ErrorType:"ServerError", Symbol:"ER_SLAVE_MI_INIT_REPOSITORY", SQLState:"HY000", Message:"Slave failed to initialize master info structure from the repository", Description:"ER_SLAVE_MI_INIT_REPOSITORY was added in 5.7.2. ", MySQLVersion:"5.7"},
    1872: definition.ErrorDefinition{ErrorNumber:1872, ErrorType:"ServerError", Symbol:"ER_SLAVE_RLI_INIT_REPOSITORY", SQLState:"HY000", Message:"Slave failed to initialize relay log info structure from the repository", Description:"ER_SLAVE_RLI_INIT_REPOSITORY was added in 5.7.2. ", MySQLVersion:"5.7"},
    1873: definition.ErrorDefinition{ErrorNumber:1873, ErrorType:"ServerError", Symbol:"ER_ACCESS_DENIED_CHANGE_USER_ERROR", SQLState:"28000", Message:"Access denied trying to change to user '%s'@'%s' (using password: %s). Disconnecting.", Description:"ER_ACCESS_DENIED_CHANGE_USER_ERROR was added in 5.7.2. ", MySQLVersion:"5.7"},
    1874: definition.ErrorDefinition{ErrorNumber:1874, ErrorType:"ServerError", Symbol:"ER_INNODB_READ_ONLY", SQLState:"HY000", Message:"InnoDB is in read only mode.", Description:"ER_INNODB_READ_ONLY was added in 5.7.2. ", MySQLVersion:"5.7"},
    1875: definition.ErrorDefinition{ErrorNumber:1875, ErrorType:"ServerError", Symbol:"ER_STOP_SLAVE_SQL_THREAD_TIMEOUT", SQLState:"HY000", Message:"STOP SLAVE command execution is incomplete: Slave SQL thread got the stop signal, thread is busy, SQL thread will stop once the current task is complete.", Description:"ER_STOP_SLAVE_SQL_THREAD_TIMEOUT was added in 5.7.2. ", MySQLVersion:"5.7"},
    1876: definition.ErrorDefinition{ErrorNumber:1876, ErrorType:"ServerError", Symbol:"ER_STOP_SLAVE_IO_THREAD_TIMEOUT", SQLState:"HY000", Message:"STOP SLAVE command execution is incomplete: Slave IO thread got the stop signal, thread is busy, IO thread will stop once the current task is complete.", Description:"ER_STOP_SLAVE_IO_THREAD_TIMEOUT was added in 5.7.2. ", MySQLVersion:"5.7"},
    1877: definition.ErrorDefinition{ErrorNumber:1877, ErrorType:"ServerError", Symbol:"ER_TABLE_CORRUPT", SQLState:"HY000", Message:"Operation cannot be performed. The table '%s.%s' is missing, corrupt or contains bad data.", Description:"ER_TABLE_CORRUPT was added in 5.7.2. ", MySQLVersion:"5.7"},
    1878: definition.ErrorDefinition{ErrorNumber:1878, ErrorType:"ServerError", Symbol:"ER_TEMP_FILE_WRITE_FAILURE", SQLState:"HY000", Message:"Temporary file write failure.", Description:"ER_TEMP_FILE_WRITE_FAILURE was added in 5.7.3. ", MySQLVersion:"5.7"},
    1879: definition.ErrorDefinition{ErrorNumber:1879, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_AUX_NOT_HEX_ID", SQLState:"HY000", Message:"Upgrade index name failed, please use create index(alter table) algorithm copy to rebuild index.", Description:"ER_INNODB_FT_AUX_NOT_HEX_ID was added in 5.7.4. ", MySQLVersion:"5.7"},
    1880: definition.ErrorDefinition{ErrorNumber:1880, ErrorType:"ServerError", Symbol:"ER_OLD_TEMPORALS_UPGRADED", SQLState:"HY000", Message:"TIME/TIMESTAMP/DATETIME columns of old format have been upgraded to the new format.", Description:"ER_OLD_TEMPORALS_UPGRADED was added in 5.7.4. ", MySQLVersion:"5.7"},
    1881: definition.ErrorDefinition{ErrorNumber:1881, ErrorType:"ServerError", Symbol:"ER_INNODB_FORCED_RECOVERY", SQLState:"HY000", Message:"Operation not allowed when innodb_forced_recovery > 0.", Description:"ER_INNODB_FORCED_RECOVERY was added in 5.7.4. ", MySQLVersion:"5.7"},
    1882: definition.ErrorDefinition{ErrorNumber:1882, ErrorType:"ServerError", Symbol:"ER_AES_INVALID_IV", SQLState:"HY000", Message:"The initialization vector supplied to %s is too short. Must be at least %d bytes long", Description:"ER_AES_INVALID_IV was added in 5.7.4. ", MySQLVersion:"5.7"},
    1883: definition.ErrorDefinition{ErrorNumber:1883, ErrorType:"ServerError", Symbol:"ER_PLUGIN_CANNOT_BE_UNINSTALLED", SQLState:"HY000", Message:"Plugin '%s' cannot be uninstalled now. %s", Description:"ER_PLUGIN_CANNOT_BE_UNINSTALLED was added in 5.7.5. ", MySQLVersion:"5.7"},
    1884: definition.ErrorDefinition{ErrorNumber:1884, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP", SQLState:"HY000", Message:"Cannot execute statement because it needs to be written to the binary log as multiple statements, and this is not allowed when @@SESSION.GTID_NEXT == 'UUID:NUMBER'.", Description:"ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP was added in 5.7.5. ", MySQLVersion:"5.7"},
    1885: definition.ErrorDefinition{ErrorNumber:1885, ErrorType:"ServerError", Symbol:"ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER", SQLState:"HY000", Message:"Slave has more GTIDs than the master has, using the master's SERVER_UUID. This may indicate that the end of the binary log was truncated or that the last binary log file was lost, e.g., after a power or disk failure when sync_binlog != 1. The master may or may not have rolled back transactions that were already replicated to the slave. Suggest to replicate any transactions that master has rolled back from slave to master, and/or commit empty transactions on master to account for transactions that have been committed on master but are not included in GTID_EXECUTED.", Description:"ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER was added in 5.7.6. ", MySQLVersion:"5.7"},
    1886: definition.ErrorDefinition{ErrorNumber:1886, ErrorType:"ServerError", Symbol:"ER_MISSING_KEY", SQLState:"HY000", Message:"The table '%s.%s' does not have the necessary key(s) defined on it. Please check the table definition and create index(s) accordingly.", Description:"ER_MISSING_KEY was added in 5.7.22. ", MySQLVersion:"5.7"},
    1887: definition.ErrorDefinition{ErrorNumber:1887, ErrorType:"ServerError", Symbol:"WARN_NAMED_PIPE_ACCESS_EVERYONE", SQLState:"HY000", Message:"Setting named_pipe_full_access_group='%s' is insecure. Consider using a Windows group with fewer members.", Description:"WARN_NAMED_PIPE_ACCESS_EVERYONE was added in 5.7.27. ", MySQLVersion:"5.7"},
    1888: definition.ErrorDefinition{ErrorNumber:1888, ErrorType:"ServerError", Symbol:"ER_FOUND_MISSING_GTIDS", SQLState:"HY000", Message:"Cannot replicate to server with server_uuid='%s' because the present server has purged required binary logs. The connecting server needs to replicate the missing transactions from elsewhere, or be replaced by a new server created from a more recent backup. To prevent this error in the future, consider increasing the binary log expiration period on the present server. %s.", Description:"ER_FOUND_MISSING_GTIDS was added in 5.7.29. ", MySQLVersion:"5.7"},
    1906: definition.ErrorDefinition{ErrorNumber:1906, ErrorType:"ServerError", Symbol:"ER_SLAVE_IO_THREAD_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with a running slave io thread", Description:"ER_SLAVE_IO_THREAD_MUST_STOP was added in 5.7.4, removed after 5.7.5. ", MySQLVersion:"5.7"},
    3000: definition.ErrorDefinition{ErrorNumber:3000, ErrorType:"ServerError", Symbol:"ER_FILE_CORRUPT", SQLState:"HY000", Message:"File %s is corrupted ", Description:"", MySQLVersion:"5.7"},
    3001: definition.ErrorDefinition{ErrorNumber:3001, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_MASTER", SQLState:"HY000", Message:"Query partially completed on the master (error on master: %d) and was aborted. There is a chance that your master is inconsistent at this point. If you are sure that your master is ok, run this query manually on the slave and then restart the slave with SET GLOBAL SQL_SLAVE_SKIP_COUNTER=1", Description:". Query:'%s' ", MySQLVersion:"5.7"},
    3002: definition.ErrorDefinition{ErrorNumber:3002, ErrorType:"ServerError", Symbol:"ER_INCONSISTENT_ERROR", SQLState:"HY000", Message:"Query caused different errors on master and slave. Error on master: message (format)='%s' error code=%d", Description:"Error on slave:actual message='%s', error code=%d. Default database:'%s'. Query:'%s' ", MySQLVersion:"5.7"},
    3003: definition.ErrorDefinition{ErrorNumber:3003, ErrorType:"ServerError", Symbol:"ER_STORAGE_ENGINE_NOT_LOADED", SQLState:"HY000", Message:"Storage engine for table '%s'.'%s' is not loaded. ", Description:"", MySQLVersion:"5.7"},
    3004: definition.ErrorDefinition{ErrorNumber:3004, ErrorType:"ServerError", Symbol:"ER_GET_STACKED_DA_WITHOUT_ACTIVE_HANDLER", SQLState:"0Z002", Message:"GET STACKED DIAGNOSTICS when handler not active ", Description:"", MySQLVersion:"5.7"},
    3005: definition.ErrorDefinition{ErrorNumber:3005, ErrorType:"ServerError", Symbol:"ER_WARN_LEGACY_SYNTAX_CONVERTED", SQLState:"HY000", Message:"%s is no longer supported. The statement was converted to %s. ", Description:"", MySQLVersion:"5.7"},
    3006: definition.ErrorDefinition{ErrorNumber:3006, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_FULLTEXT_PLUGIN", SQLState:"HY000", Message:"Statement is unsafe because it uses a fulltext parser plugin which may not return the same value on the slave.", Description:"ER_BINLOG_UNSAFE_FULLTEXT_PLUGIN was added in 5.7.1. ", MySQLVersion:"5.7"},
    3007: definition.ErrorDefinition{ErrorNumber:3007, ErrorType:"ServerError", Symbol:"ER_CANNOT_DISCARD_TEMPORARY_TABLE", SQLState:"HY000", Message:"Cannot DISCARD/IMPORT tablespace associated with temporary table", Description:"ER_CANNOT_DISCARD_TEMPORARY_TABLE was added in 5.7.1. ", MySQLVersion:"5.7"},
    3008: definition.ErrorDefinition{ErrorNumber:3008, ErrorType:"ServerError", Symbol:"ER_FK_DEPTH_EXCEEDED", SQLState:"HY000", Message:"Foreign key cascade delete/update exceeds max depth of %d.", Description:"ER_FK_DEPTH_EXCEEDED was added in 5.7.2. ", MySQLVersion:"5.7"},
    3009: definition.ErrorDefinition{ErrorNumber:3009, ErrorType:"ServerError", Symbol:"ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2", SQLState:"HY000", Message:"Column count of %s.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please use mysql_upgrade to fix this error.", Description:"ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2 was added in 5.7.2. ", MySQLVersion:"5.7"},
    3010: definition.ErrorDefinition{ErrorNumber:3010, ErrorType:"ServerError", Symbol:"ER_WARN_TRIGGER_DOESNT_HAVE_CREATED", SQLState:"HY000", Message:"Trigger %s.%s.%s does not have CREATED attribute.", Description:"ER_WARN_TRIGGER_DOESNT_HAVE_CREATED was added in 5.7.2. ", MySQLVersion:"5.7"},
    3011: definition.ErrorDefinition{ErrorNumber:3011, ErrorType:"ServerError", Symbol:"ER_REFERENCED_TRG_DOES_NOT_EXIST", SQLState:"HY000", Message:"Referenced trigger '%s' for the given action time and event type does not exist.", Description:"ER_REFERENCED_TRG_DOES_NOT_EXIST was added in 5.7.2. ", MySQLVersion:"5.7"},
    3012: definition.ErrorDefinition{ErrorNumber:3012, ErrorType:"ServerError", Symbol:"ER_EXPLAIN_NOT_SUPPORTED", SQLState:"HY000", Message:"EXPLAIN FOR CONNECTION command is supported only for SELECT/UPDATE/INSERT/DELETE/REPLACE", Description:"ER_EXPLAIN_NOT_SUPPORTED was added in 5.7.2. ", MySQLVersion:"5.7"},
    3013: definition.ErrorDefinition{ErrorNumber:3013, ErrorType:"ServerError", Symbol:"ER_INVALID_FIELD_SIZE", SQLState:"HY000", Message:"Invalid size for column '%s'.", Description:"ER_INVALID_FIELD_SIZE was added in 5.7.2. ", MySQLVersion:"5.7"},
    3014: definition.ErrorDefinition{ErrorNumber:3014, ErrorType:"ServerError", Symbol:"ER_MISSING_HA_CREATE_OPTION", SQLState:"HY000", Message:"Table storage engine '%s' found required create option missing", Description:"ER_MISSING_HA_CREATE_OPTION was added in 5.7.2. ", MySQLVersion:"5.7"},
    3015: definition.ErrorDefinition{ErrorNumber:3015, ErrorType:"ServerError", Symbol:"ER_ENGINE_OUT_OF_MEMORY", SQLState:"HY000", Message:"Out of memory in storage engine '%s'.", Description:"ER_ENGINE_OUT_OF_MEMORY was added in 5.7.3. ", MySQLVersion:"5.7"},
    3016: definition.ErrorDefinition{ErrorNumber:3016, ErrorType:"ServerError", Symbol:"ER_PASSWORD_EXPIRE_ANONYMOUS_USER", SQLState:"HY000", Message:"The password for anonymous user cannot be expired.", Description:"ER_PASSWORD_EXPIRE_ANONYMOUS_USER was added in 5.7.3. ", MySQLVersion:"5.7"},
    3017: definition.ErrorDefinition{ErrorNumber:3017, ErrorType:"ServerError", Symbol:"ER_SLAVE_SQL_THREAD_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with a running slave sql thread", Description:"ER_SLAVE_SQL_THREAD_MUST_STOP was added in 5.7.3. ", MySQLVersion:"5.7"},
    3018: definition.ErrorDefinition{ErrorNumber:3018, ErrorType:"ServerError", Symbol:"ER_NO_FT_MATERIALIZED_SUBQUERY", SQLState:"HY000", Message:"Cannot create FULLTEXT index on materialized subquery", Description:"ER_NO_FT_MATERIALIZED_SUBQUERY was added in 5.7.4. ", MySQLVersion:"5.7"},
    3019: definition.ErrorDefinition{ErrorNumber:3019, ErrorType:"ServerError", Symbol:"ER_INNODB_UNDO_LOG_FULL", SQLState:"HY000", Message:"Undo Log error: %s", Description:"ER_INNODB_UNDO_LOG_FULL was added in 5.7.4. ", MySQLVersion:"5.7"},
    3020: definition.ErrorDefinition{ErrorNumber:3020, ErrorType:"ServerError", Symbol:"ER_INVALID_ARGUMENT_FOR_LOGARITHM", SQLState:"2201E", Message:"Invalid argument for logarithm", Description:"ER_INVALID_ARGUMENT_FOR_LOGARITHM was added in 5.7.4. ", MySQLVersion:"5.7"},
    3021: definition.ErrorDefinition{ErrorNumber:3021, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_IO_THREAD_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with a running slave io thread", Description:"ER_SLAVE_CHANNEL_IO_THREAD_MUST_STOP was added in 5.7.6. ", MySQLVersion:"5.7"},
    3022: definition.ErrorDefinition{ErrorNumber:3022, ErrorType:"ServerError", Symbol:"ER_WARN_OPEN_TEMP_TABLES_MUST_BE_ZERO", SQLState:"HY000", Message:"This operation may not be safe when the slave has temporary tables. The tables will be kept open until the server restarts or until the tables are deleted by any replicated DROP statement. Suggest to wait until slave_open_temp_tables = 0.", Description:"ER_WARN_OPEN_TEMP_TABLES_MUST_BE_ZERO was added in 5.7.4. ", MySQLVersion:"5.7"},
    3023: definition.ErrorDefinition{ErrorNumber:3023, ErrorType:"ServerError", Symbol:"ER_WARN_ONLY_MASTER_LOG_FILE_NO_POS", SQLState:"HY000", Message:"CHANGE MASTER TO with a MASTER_LOG_FILE clause but no MASTER_LOG_POS clause may not be safe. The old position value may not be valid for the new binary log file.", Description:"ER_WARN_ONLY_MASTER_LOG_FILE_NO_POS was added in 5.7.4. ", MySQLVersion:"5.7"},
    3024: definition.ErrorDefinition{ErrorNumber:3024, ErrorType:"ServerError", Symbol:"ER_QUERY_TIMEOUT", SQLState:"HY000", Message:"Query execution was interrupted, maximum statement execution time exceeded", Description:"ER_QUERY_TIMEOUT was added in 5.7.4. ", MySQLVersion:"5.7"},
    3025: definition.ErrorDefinition{ErrorNumber:3025, ErrorType:"ServerError", Symbol:"ER_NON_RO_SELECT_DISABLE_TIMER", SQLState:"HY000", Message:"Select is not a read only statement, disabling timer", Description:"ER_NON_RO_SELECT_DISABLE_TIMER was added in 5.7.4. ", MySQLVersion:"5.7"},
    3026: definition.ErrorDefinition{ErrorNumber:3026, ErrorType:"ServerError", Symbol:"ER_DUP_LIST_ENTRY", SQLState:"HY000", Message:"Duplicate entry '%s'.", Description:"ER_DUP_LIST_ENTRY was added in 5.7.4. ", MySQLVersion:"5.7"},
    3027: definition.ErrorDefinition{ErrorNumber:3027, ErrorType:"ServerError", Symbol:"ER_SQL_MODE_NO_EFFECT", SQLState:"HY000", Message:"'%s' mode no longer has any effect. Use STRICT_ALL_TABLES or STRICT_TRANS_TABLES instead.", Description:"ER_SQL_MODE_NO_EFFECT was added in 5.7.4. ", MySQLVersion:"5.7"},
    3028: definition.ErrorDefinition{ErrorNumber:3028, ErrorType:"ServerError", Symbol:"ER_AGGREGATE_ORDER_FOR_UNION", SQLState:"HY000", Message:"Expression #%u of ORDER BY contains aggregate function and applies to a UNION", Description:"ER_AGGREGATE_ORDER_FOR_UNION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3029: definition.ErrorDefinition{ErrorNumber:3029, ErrorType:"ServerError", Symbol:"ER_AGGREGATE_ORDER_NON_AGG_QUERY", SQLState:"HY000", Message:"Expression #%u of ORDER BY contains aggregate function and applies to the result of a non-aggregated query", Description:"ER_AGGREGATE_ORDER_NON_AGG_QUERY was added in 5.7.5. ", MySQLVersion:"5.7"},
    3030: definition.ErrorDefinition{ErrorNumber:3030, ErrorType:"ServerError", Symbol:"ER_SLAVE_WORKER_STOPPED_PREVIOUS_THD_ERROR", SQLState:"HY000", Message:"Slave worker has stopped after at least one previous worker encountered an error when slave-preserve-commit-order was enabled. To preserve commit order, the last transaction executed by this thread has not been committed. When restarting the slave after fixing any failed threads, you should fix this worker as well.", Description:"ER_SLAVE_WORKER_STOPPED_PREVIOUS_THD_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3031: definition.ErrorDefinition{ErrorNumber:3031, ErrorType:"ServerError", Symbol:"ER_DONT_SUPPORT_SLAVE_PRESERVE_COMMIT_ORDER", SQLState:"HY000", Message:"slave_preserve_commit_order is not supported %s.", Description:"ER_DONT_SUPPORT_SLAVE_PRESERVE_COMMIT_ORDER was added in 5.7.5. ", MySQLVersion:"5.7"},
    3032: definition.ErrorDefinition{ErrorNumber:3032, ErrorType:"ServerError", Symbol:"ER_SERVER_OFFLINE_MODE", SQLState:"HY000", Message:"The server is currently in offline mode", Description:"ER_SERVER_OFFLINE_MODE was added in 5.7.5. ", MySQLVersion:"5.7"},
    3033: definition.ErrorDefinition{ErrorNumber:3033, ErrorType:"ServerError", Symbol:"ER_GIS_DIFFERENT_SRIDS", SQLState:"HY000", Message:"Binary geometry function %s given two geometries of different srids: %u and %u, which should have been identical.", Description:"ER_GIS_DIFFERENT_SRIDS was added in 5.7.5. ", MySQLVersion:"5.7"},
    3034: definition.ErrorDefinition{ErrorNumber:3034, ErrorType:"ServerError", Symbol:"ER_GIS_UNSUPPORTED_ARGUMENT", SQLState:"HY000", Message:"Calling geometry function %s with unsupported types of arguments.", Description:"ER_GIS_UNSUPPORTED_ARGUMENT was added in 5.7.5. ", MySQLVersion:"5.7"},
    3035: definition.ErrorDefinition{ErrorNumber:3035, ErrorType:"ServerError", Symbol:"ER_GIS_UNKNOWN_ERROR", SQLState:"HY000", Message:"Unknown GIS error occured in function %s.", Description:"ER_GIS_UNKNOWN_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3036: definition.ErrorDefinition{ErrorNumber:3036, ErrorType:"ServerError", Symbol:"ER_GIS_UNKNOWN_EXCEPTION", SQLState:"HY000", Message:"Unknown exception caught in GIS function %s.", Description:"ER_GIS_UNKNOWN_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3037: definition.ErrorDefinition{ErrorNumber:3037, ErrorType:"ServerError", Symbol:"ER_GIS_INVALID_DATA", SQLState:"22023", Message:"Invalid GIS data provided to function %s.", Description:"ER_GIS_INVALID_DATA was added in 5.7.5. ", MySQLVersion:"5.7"},
    3038: definition.ErrorDefinition{ErrorNumber:3038, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_EMPTY_INPUT_EXCEPTION", SQLState:"HY000", Message:"The geometry has no data in function %s.", Description:"ER_BOOST_GEOMETRY_EMPTY_INPUT_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3039: definition.ErrorDefinition{ErrorNumber:3039, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_CENTROID_EXCEPTION", SQLState:"HY000", Message:"Unable to calculate centroid because geometry is empty in function %s.", Description:"ER_BOOST_GEOMETRY_CENTROID_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3040: definition.ErrorDefinition{ErrorNumber:3040, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_OVERLAY_INVALID_INPUT_EXCEPTION", SQLState:"HY000", Message:"Geometry overlay calculation error: geometry data is invalid in function %s.", Description:"ER_BOOST_GEOMETRY_OVERLAY_INVALID_INPUT_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3041: definition.ErrorDefinition{ErrorNumber:3041, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_TURN_INFO_EXCEPTION", SQLState:"HY000", Message:"Geometry turn info calculation error: geometry data is invalid in function %s.", Description:"ER_BOOST_GEOMETRY_TURN_INFO_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3042: definition.ErrorDefinition{ErrorNumber:3042, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_SELF_INTERSECTION_POINT_EXCEPTION", SQLState:"HY000", Message:"Analysis procedures of intersection points interrupted unexpectedly in function %s.", Description:"ER_BOOST_GEOMETRY_SELF_INTERSECTION_POINT_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3043: definition.ErrorDefinition{ErrorNumber:3043, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_UNKNOWN_EXCEPTION", SQLState:"HY000", Message:"Unknown exception thrown in function %s.", Description:"ER_BOOST_GEOMETRY_UNKNOWN_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3044: definition.ErrorDefinition{ErrorNumber:3044, ErrorType:"ServerError", Symbol:"ER_STD_BAD_ALLOC_ERROR", SQLState:"HY000", Message:"Memory allocation error: %s in function %s.", Description:"ER_STD_BAD_ALLOC_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3045: definition.ErrorDefinition{ErrorNumber:3045, ErrorType:"ServerError", Symbol:"ER_STD_DOMAIN_ERROR", SQLState:"HY000", Message:"Domain error: %s in function %s.", Description:"ER_STD_DOMAIN_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3046: definition.ErrorDefinition{ErrorNumber:3046, ErrorType:"ServerError", Symbol:"ER_STD_LENGTH_ERROR", SQLState:"HY000", Message:"Length error: %s in function %s.", Description:"ER_STD_LENGTH_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3047: definition.ErrorDefinition{ErrorNumber:3047, ErrorType:"ServerError", Symbol:"ER_STD_INVALID_ARGUMENT", SQLState:"HY000", Message:"Invalid argument error: %s in function %s.", Description:"ER_STD_INVALID_ARGUMENT was added in 5.7.5. ", MySQLVersion:"5.7"},
    3048: definition.ErrorDefinition{ErrorNumber:3048, ErrorType:"ServerError", Symbol:"ER_STD_OUT_OF_RANGE_ERROR", SQLState:"HY000", Message:"Out of range error: %s in function %s.", Description:"ER_STD_OUT_OF_RANGE_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3049: definition.ErrorDefinition{ErrorNumber:3049, ErrorType:"ServerError", Symbol:"ER_STD_OVERFLOW_ERROR", SQLState:"HY000", Message:"Overflow error error: %s in function %s.", Description:"ER_STD_OVERFLOW_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3050: definition.ErrorDefinition{ErrorNumber:3050, ErrorType:"ServerError", Symbol:"ER_STD_RANGE_ERROR", SQLState:"HY000", Message:"Range error: %s in function %s.", Description:"ER_STD_RANGE_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3051: definition.ErrorDefinition{ErrorNumber:3051, ErrorType:"ServerError", Symbol:"ER_STD_UNDERFLOW_ERROR", SQLState:"HY000", Message:"Underflow error: %s in function %s.", Description:"ER_STD_UNDERFLOW_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3052: definition.ErrorDefinition{ErrorNumber:3052, ErrorType:"ServerError", Symbol:"ER_STD_LOGIC_ERROR", SQLState:"HY000", Message:"Logic error: %s in function %s.", Description:"ER_STD_LOGIC_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3053: definition.ErrorDefinition{ErrorNumber:3053, ErrorType:"ServerError", Symbol:"ER_STD_RUNTIME_ERROR", SQLState:"HY000", Message:"Runtime error: %s in function %s.", Description:"ER_STD_RUNTIME_ERROR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3054: definition.ErrorDefinition{ErrorNumber:3054, ErrorType:"ServerError", Symbol:"ER_STD_UNKNOWN_EXCEPTION", SQLState:"HY000", Message:"Unknown exception: %s in function %s.", Description:"ER_STD_UNKNOWN_EXCEPTION was added in 5.7.5. ", MySQLVersion:"5.7"},
    3055: definition.ErrorDefinition{ErrorNumber:3055, ErrorType:"ServerError", Symbol:"ER_GIS_DATA_WRONG_ENDIANESS", SQLState:"HY000", Message:"Geometry byte string must be little endian.", Description:"ER_GIS_DATA_WRONG_ENDIANESS was added in 5.7.5. ", MySQLVersion:"5.7"},
    3056: definition.ErrorDefinition{ErrorNumber:3056, ErrorType:"ServerError", Symbol:"ER_CHANGE_MASTER_PASSWORD_LENGTH", SQLState:"HY000", Message:"The password provided for the replication user exceeds the maximum length of 32 characters", Description:"ER_CHANGE_MASTER_PASSWORD_LENGTH was added in 5.7.5. ", MySQLVersion:"5.7"},
    3057: definition.ErrorDefinition{ErrorNumber:3057, ErrorType:"ServerError", Symbol:"ER_USER_LOCK_WRONG_NAME", SQLState:"42000", Message:"Incorrect user-level lock name '%s'.", Description:"ER_USER_LOCK_WRONG_NAME was added in 5.7.5. ", MySQLVersion:"5.7"},
    3058: definition.ErrorDefinition{ErrorNumber:3058, ErrorType:"ServerError", Symbol:"ER_USER_LOCK_DEADLOCK", SQLState:"HY000", Message:"Deadlock found when trying to get user-level lock", Description:"ER_USER_LOCK_DEADLOCK was added in 5.7.5. ", MySQLVersion:"5.7"},
    3059: definition.ErrorDefinition{ErrorNumber:3059, ErrorType:"ServerError", Symbol:"ER_REPLACE_INACCESSIBLE_ROWS", SQLState:"HY000", Message:"REPLACE cannot be executed as it requires deleting rows that are not in the view", Description:"ER_REPLACE_INACCESSIBLE_ROWS was added in 5.7.5. ", MySQLVersion:"5.7"},
    3060: definition.ErrorDefinition{ErrorNumber:3060, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_GIS", SQLState:"HY000", Message:"Do not support online operation on table with GIS index", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_GIS was added in 5.7.5. ", MySQLVersion:"5.7"},
    3061: definition.ErrorDefinition{ErrorNumber:3061, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_USER_VAR", SQLState:"42000", Message:"User variable name '%s' is illegal", Description:"ER_ILLEGAL_USER_VAR was added in 5.7.5. ", MySQLVersion:"5.7"},
    3062: definition.ErrorDefinition{ErrorNumber:3062, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_OFF", SQLState:"HY000", Message:"Cannot %s when GTID_MODE = OFF.", Description:"ER_GTID_MODE_OFF was added in 5.7.5. ", MySQLVersion:"5.7"},
    3063: definition.ErrorDefinition{ErrorNumber:3063, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_BY_REPLICATION_THREAD", SQLState:"HY000", Message:"Cannot %s from a replication slave thread.", Description:"ER_UNSUPPORTED_BY_REPLICATION_THREAD was added in 5.7.5. ", MySQLVersion:"5.7"},
    3064: definition.ErrorDefinition{ErrorNumber:3064, ErrorType:"ServerError", Symbol:"ER_INCORRECT_TYPE", SQLState:"HY000", Message:"Incorrect type for argument %s in function %s.", Description:"ER_INCORRECT_TYPE was added in 5.7.5. ", MySQLVersion:"5.7"},
    3065: definition.ErrorDefinition{ErrorNumber:3065, ErrorType:"ServerError", Symbol:"ER_FIELD_IN_ORDER_NOT_SELECT", SQLState:"HY000", Message:"Expression #%u of ORDER BY clause is not in SELECT list, references column '%s' which is not in SELECT list", Description:"ER_FIELD_IN_ORDER_NOT_SELECT was added in 5.7.5. ", MySQLVersion:"5.7"},
    3066: definition.ErrorDefinition{ErrorNumber:3066, ErrorType:"ServerError", Symbol:"ER_AGGREGATE_IN_ORDER_NOT_SELECT", SQLState:"HY000", Message:"Expression #%u of ORDER BY clause is not in SELECT list, contains aggregate function", Description:"ER_AGGREGATE_IN_ORDER_NOT_SELECT was added in 5.7.5. ", MySQLVersion:"5.7"},
    3067: definition.ErrorDefinition{ErrorNumber:3067, ErrorType:"ServerError", Symbol:"ER_INVALID_RPL_WILD_TABLE_FILTER_PATTERN", SQLState:"HY000", Message:"Supplied filter list contains a value which is not in the required format 'db_pattern.table_pattern'", Description:"ER_INVALID_RPL_WILD_TABLE_FILTER_PATTERN was added in 5.7.5. ", MySQLVersion:"5.7"},
    3068: definition.ErrorDefinition{ErrorNumber:3068, ErrorType:"ServerError", Symbol:"ER_NET_OK_PACKET_TOO_LARGE", SQLState:"08S01", Message:"OK packet too large", Description:"ER_NET_OK_PACKET_TOO_LARGE was added in 5.7.5. ", MySQLVersion:"5.7"},
    3069: definition.ErrorDefinition{ErrorNumber:3069, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_DATA", SQLState:"HY000", Message:"Invalid JSON data provided to function %s: %s", Description:"ER_INVALID_JSON_DATA was added in 5.7.5. ", MySQLVersion:"5.7"},
    3070: definition.ErrorDefinition{ErrorNumber:3070, ErrorType:"ServerError", Symbol:"ER_INVALID_GEOJSON_MISSING_MEMBER", SQLState:"HY000", Message:"Invalid GeoJSON data provided to function %s: Missing required member '%s'", Description:"ER_INVALID_GEOJSON_MISSING_MEMBER was added in 5.7.5. ", MySQLVersion:"5.7"},
    3071: definition.ErrorDefinition{ErrorNumber:3071, ErrorType:"ServerError", Symbol:"ER_INVALID_GEOJSON_WRONG_TYPE", SQLState:"HY000", Message:"Invalid GeoJSON data provided to function %s: Member '%s' must be of type '%s'", Description:"ER_INVALID_GEOJSON_WRONG_TYPE was added in 5.7.5. ", MySQLVersion:"5.7"},
    3072: definition.ErrorDefinition{ErrorNumber:3072, ErrorType:"ServerError", Symbol:"ER_INVALID_GEOJSON_UNSPECIFIED", SQLState:"HY000", Message:"Invalid GeoJSON data provided to function %s", Description:"ER_INVALID_GEOJSON_UNSPECIFIED was added in 5.7.5. ", MySQLVersion:"5.7"},
    3073: definition.ErrorDefinition{ErrorNumber:3073, ErrorType:"ServerError", Symbol:"ER_DIMENSION_UNSUPPORTED", SQLState:"HY000", Message:"Unsupported number of coordinate dimensions in function %s: Found %u, expected %u", Description:"ER_DIMENSION_UNSUPPORTED was added in 5.7.5. ", MySQLVersion:"5.7"},
    3074: definition.ErrorDefinition{ErrorNumber:3074, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_DOES_NOT_EXIST", SQLState:"HY000", Message:"Slave channel '%s' does not exist.", Description:"ER_SLAVE_CHANNEL_DOES_NOT_EXIST was added in 5.7.6. ", MySQLVersion:"5.7"},
    3075: definition.ErrorDefinition{ErrorNumber:3075, ErrorType:"ServerError", Symbol:"ER_SLAVE_MULTIPLE_CHANNELS_HOST_PORT", SQLState:"HY000", Message:"A slave channel '%s' already exists for the given host and port combination.", Description:"ER_SLAVE_MULTIPLE_CHANNELS_HOST_PORT was added in 5.7.6. ", MySQLVersion:"5.7"},
    3076: definition.ErrorDefinition{ErrorNumber:3076, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_NAME_INVALID_OR_TOO_LONG", SQLState:"HY000", Message:"Couldn't create channel: Channel name is either invalid or too long.", Description:"ER_SLAVE_CHANNEL_NAME_INVALID_OR_TOO_LONG was added in 5.7.6. ", MySQLVersion:"5.7"},
    3077: definition.ErrorDefinition{ErrorNumber:3077, ErrorType:"ServerError", Symbol:"ER_SLAVE_NEW_CHANNEL_WRONG_REPOSITORY", SQLState:"HY000", Message:"To have multiple channels, repository cannot be of type FILE", Description:"ER_SLAVE_NEW_CHANNEL_WRONG_REPOSITORY was added in 5.7.6. ", MySQLVersion:"5.7"},
    3078: definition.ErrorDefinition{ErrorNumber:3078, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_DELETE", SQLState:"HY000", Message:"Cannot delete slave info objects for channel '%s'.", Description:"ER_SLAVE_CHANNEL_DELETE was added in 5.7.6. ", MySQLVersion:"5.7"},
    3079: definition.ErrorDefinition{ErrorNumber:3079, ErrorType:"ServerError", Symbol:"ER_SLAVE_MULTIPLE_CHANNELS_CMD", SQLState:"HY000", Message:"Multiple channels exist on the slave. Please provide channel name as an argument.", Description:"ER_SLAVE_MULTIPLE_CHANNELS_CMD was added in 5.7.6. ", MySQLVersion:"5.7"},
    3080: definition.ErrorDefinition{ErrorNumber:3080, ErrorType:"ServerError", Symbol:"ER_SLAVE_MAX_CHANNELS_EXCEEDED", SQLState:"HY000", Message:"Maximum number of replication channels allowed exceeded.", Description:"ER_SLAVE_MAX_CHANNELS_EXCEEDED was added in 5.7.6. ", MySQLVersion:"5.7"},
    3081: definition.ErrorDefinition{ErrorNumber:3081, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with running replication threads", Description:"ER_SLAVE_CHANNEL_MUST_STOP was added in 5.7.6. ", MySQLVersion:"5.7"},
    3082: definition.ErrorDefinition{ErrorNumber:3082, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_NOT_RUNNING", SQLState:"HY000", Message:"This operation requires running replication threads", Description:"ER_SLAVE_CHANNEL_NOT_RUNNING was added in 5.7.6. ", MySQLVersion:"5.7"},
    3083: definition.ErrorDefinition{ErrorNumber:3083, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_WAS_RUNNING", SQLState:"HY000", Message:"Replication thread(s) for channel '%s' are already runnning.", Description:"ER_SLAVE_CHANNEL_WAS_RUNNING was added in 5.7.6. ", MySQLVersion:"5.7"},
    3084: definition.ErrorDefinition{ErrorNumber:3084, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_WAS_NOT_RUNNING", SQLState:"HY000", Message:"Replication thread(s) for channel '%s' are already stopped.", Description:"ER_SLAVE_CHANNEL_WAS_NOT_RUNNING was added in 5.7.6. ", MySQLVersion:"5.7"},
    3085: definition.ErrorDefinition{ErrorNumber:3085, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_SQL_THREAD_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with a running slave sql thread", Description:"ER_SLAVE_CHANNEL_SQL_THREAD_MUST_STOP was added in 5.7.6. ", MySQLVersion:"5.7"},
    3086: definition.ErrorDefinition{ErrorNumber:3086, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_SQL_SKIP_COUNTER", SQLState:"HY000", Message:"When sql_slave_skip_counter > 0, it is not allowed to start more than one SQL thread by using 'START SLAVE [SQL_THREAD]'. Value of sql_slave_skip_counter can only be used by one SQL thread at a time. Please use 'START SLAVE [SQL_THREAD] FOR CHANNEL' to start the SQL thread which will use the value of sql_slave_skip_counter.", Description:"ER_SLAVE_CHANNEL_SQL_SKIP_COUNTER was added in 5.7.6. ", MySQLVersion:"5.7"},
    3087: definition.ErrorDefinition{ErrorNumber:3087, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_WITH_GROUP_V2", SQLState:"HY000", Message:"Expression #%u of %s is not in GROUP BY clause and contains nonaggregated column '%s' which is not functionally dependent on columns in GROUP BY clause", Description:"ER_WRONG_FIELD_WITH_GROUP_V2 was added in 5.7.6. ", MySQLVersion:"5.7"},
    3088: definition.ErrorDefinition{ErrorNumber:3088, ErrorType:"ServerError", Symbol:"ER_MIX_OF_GROUP_FUNC_AND_FIELDS_V2", SQLState:"HY000", Message:"In aggregated query without GROUP BY, expression #%u of %s contains nonaggregated column '%s'", Description:"ER_MIX_OF_GROUP_FUNC_AND_FIELDS_V2 was added in 5.7.6. ", MySQLVersion:"5.7"},
    3089: definition.ErrorDefinition{ErrorNumber:3089, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYSVAR_UPDATE", SQLState:"HY000", Message:"Updating '%s' is deprecated. It will be made read-only in a future release.", Description:"ER_WARN_DEPRECATED_SYSVAR_UPDATE was added in 5.7.6. ", MySQLVersion:"5.7"},
    3090: definition.ErrorDefinition{ErrorNumber:3090, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SQLMODE", SQLState:"HY000", Message:"Changing sql mode '%s' is deprecated. It will be removed in a future release.", Description:"ER_WARN_DEPRECATED_SQLMODE was added in 5.7.6. ", MySQLVersion:"5.7"},
    3091: definition.ErrorDefinition{ErrorNumber:3091, ErrorType:"ServerError", Symbol:"ER_CANNOT_LOG_PARTIAL_DROP_DATABASE_WITH_GTID", SQLState:"HY000", Message:"DROP DATABASE failed", Description:"ER_CANNOT_LOG_PARTIAL_DROP_DATABASE_WITH_GTID was added in 5.7.6. ", MySQLVersion:"5.7"},
    3092: definition.ErrorDefinition{ErrorNumber:3092, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_CONFIGURATION", SQLState:"HY000", Message:"The server is not configured properly to be an active member of the group. Please see more details on error log.", Description:"ER_GROUP_REPLICATION_CONFIGURATION was added in 5.7.6. ", MySQLVersion:"5.7"},
    3093: definition.ErrorDefinition{ErrorNumber:3093, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_RUNNING", SQLState:"HY000", Message:"The START GROUP_REPLICATION command failed since the group is already running.", Description:"ER_GROUP_REPLICATION_RUNNING was added in 5.7.6. ", MySQLVersion:"5.7"},
    3094: definition.ErrorDefinition{ErrorNumber:3094, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_APPLIER_INIT_ERROR", SQLState:"HY000", Message:"The START GROUP_REPLICATION command failed as the applier module failed to start.", Description:"ER_GROUP_REPLICATION_APPLIER_INIT_ERROR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3095: definition.ErrorDefinition{ErrorNumber:3095, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_STOP_APPLIER_THREAD_TIMEOUT", SQLState:"HY000", Message:"The STOP GROUP_REPLICATION command execution is incomplete: The applier thread got the stop signal while it was busy. The applier thread will stop once the current task is complete.", Description:"ER_GROUP_REPLICATION_STOP_APPLIER_THREAD_TIMEOUT was added in 5.7.6. ", MySQLVersion:"5.7"},
    3096: definition.ErrorDefinition{ErrorNumber:3096, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_COMMUNICATION_LAYER_SESSION_ERROR", SQLState:"HY000", Message:"The START GROUP_REPLICATION command failed as there was an error when initializing the group communication layer.", Description:"ER_GROUP_REPLICATION_COMMUNICATION_LAYER_SESSION_ERROR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3097: definition.ErrorDefinition{ErrorNumber:3097, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_COMMUNICATION_LAYER_JOIN_ERROR", SQLState:"HY000", Message:"The START GROUP_REPLICATION command failed as there was an error when joining the communication group.", Description:"ER_GROUP_REPLICATION_COMMUNICATION_LAYER_JOIN_ERROR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3098: definition.ErrorDefinition{ErrorNumber:3098, ErrorType:"ServerError", Symbol:"ER_BEFORE_DML_VALIDATION_ERROR", SQLState:"HY000", Message:"The table does not comply with the requirements by an external plugin.", Description:"ER_BEFORE_DML_VALIDATION_ERROR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3099: definition.ErrorDefinition{ErrorNumber:3099, ErrorType:"ServerError", Symbol:"ER_PREVENTS_VARIABLE_WITHOUT_RBR", SQLState:"HY000", Message:"Cannot change the value of variable %s without binary log format as ROW.", Description:"ER_PREVENTS_VARIABLE_WITHOUT_RBR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3100: definition.ErrorDefinition{ErrorNumber:3100, ErrorType:"ServerError", Symbol:"ER_RUN_HOOK_ERROR", SQLState:"HY000", Message:"Error on observer while running replication hook '%s'.", Description:"ER_RUN_HOOK_ERROR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3101: definition.ErrorDefinition{ErrorNumber:3101, ErrorType:"ServerError", Symbol:"ER_TRANSACTION_ROLLBACK_DURING_COMMIT", SQLState:"HY000", Message:"Plugin instructed the server to rollback the current transaction.", Description:"ER_TRANSACTION_ROLLBACK_DURING_COMMIT was added in 5.7.6. ", MySQLVersion:"5.7"},
    3102: definition.ErrorDefinition{ErrorNumber:3102, ErrorType:"ServerError", Symbol:"ER_GENERATED_COLUMN_FUNCTION_IS_NOT_ALLOWED", SQLState:"HY000", Message:"Expression of generated column '%s' contains a disallowed function.", Description:"ER_GENERATED_COLUMN_FUNCTION_IS_NOT_ALLOWED was added in 5.7.6. ", MySQLVersion:"5.7"},
    3103: definition.ErrorDefinition{ErrorNumber:3103, ErrorType:"ServerError", Symbol:"ER_KEY_BASED_ON_GENERATED_COLUMN", SQLState:"HY000", Message:"Key/Index cannot be defined on a virtual generated column.", Description:"ER_KEY_BASED_ON_GENERATED_COLUMN was added in 5.7.6, removed after 5.7.7. ", MySQLVersion:"5.7"},
    3103: definition.ErrorDefinition{ErrorNumber:3103, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_ALTER_INPLACE_ON_VIRTUAL_COLUMN", SQLState:"HY000", Message:"INPLACE ADD or DROP of virtual columns cannot be combined with other ALTER TABLE actions", Description:"ER_UNSUPPORTED_ALTER_INPLACE_ON_VIRTUAL_COLUMN was added in 5.7.8. ", MySQLVersion:"5.7"},
    3104: definition.ErrorDefinition{ErrorNumber:3104, ErrorType:"ServerError", Symbol:"ER_WRONG_FK_OPTION_FOR_GENERATED_COLUMN", SQLState:"HY000", Message:"Cannot define foreign key with %s clause on a generated column.", Description:"ER_WRONG_FK_OPTION_FOR_GENERATED_COLUMN was added in 5.7.6. ", MySQLVersion:"5.7"},
    3105: definition.ErrorDefinition{ErrorNumber:3105, ErrorType:"ServerError", Symbol:"ER_NON_DEFAULT_VALUE_FOR_GENERATED_COLUMN", SQLState:"HY000", Message:"The value specified for generated column '%s' in table '%s' is not allowed.", Description:"ER_NON_DEFAULT_VALUE_FOR_GENERATED_COLUMN was added in 5.7.6. ", MySQLVersion:"5.7"},
    3106: definition.ErrorDefinition{ErrorNumber:3106, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_ACTION_ON_GENERATED_COLUMN", SQLState:"HY000", Message:"'%s' is not supported for generated columns.", Description:"ER_UNSUPPORTED_ACTION_ON_GENERATED_COLUMN was added in 5.7.6. ", MySQLVersion:"5.7"},
    3107: definition.ErrorDefinition{ErrorNumber:3107, ErrorType:"ServerError", Symbol:"ER_GENERATED_COLUMN_NON_PRIOR", SQLState:"HY000", Message:"Generated column can refer only to generated columns defined prior to it.", Description:"ER_GENERATED_COLUMN_NON_PRIOR was added in 5.7.6. ", MySQLVersion:"5.7"},
    3108: definition.ErrorDefinition{ErrorNumber:3108, ErrorType:"ServerError", Symbol:"ER_DEPENDENT_BY_GENERATED_COLUMN", SQLState:"HY000", Message:"Column '%s' has a generated column dependency.", Description:"ER_DEPENDENT_BY_GENERATED_COLUMN was added in 5.7.6. ", MySQLVersion:"5.7"},
    3109: definition.ErrorDefinition{ErrorNumber:3109, ErrorType:"ServerError", Symbol:"ER_GENERATED_COLUMN_REF_AUTO_INC", SQLState:"HY000", Message:"Generated column '%s' cannot refer to auto-increment column.", Description:"ER_GENERATED_COLUMN_REF_AUTO_INC was added in 5.7.6. ", MySQLVersion:"5.7"},
    3110: definition.ErrorDefinition{ErrorNumber:3110, ErrorType:"ServerError", Symbol:"ER_FEATURE_NOT_AVAILABLE", SQLState:"HY000", Message:"The '%s' feature is not available", Description:"ER_FEATURE_NOT_AVAILABLE was added in 5.7.6. ", MySQLVersion:"5.7"},
    3111: definition.ErrorDefinition{ErrorNumber:3111, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_MODE", SQLState:"HY000", Message:"SET @@GLOBAL.GTID_MODE = %s is not allowed because %s.", Description:"ER_CANT_SET_GTID_MODE was added in 5.7.6. ", MySQLVersion:"5.7"},
    3112: definition.ErrorDefinition{ErrorNumber:3112, ErrorType:"ServerError", Symbol:"ER_CANT_USE_AUTO_POSITION_WITH_GTID_MODE_OFF", SQLState:"HY000", Message:"The replication receiver thread%s cannot start in AUTO_POSITION mode: this server uses @@GLOBAL.GTID_MODE = OFF.", Description:"ER_CANT_USE_AUTO_POSITION_WITH_GTID_MODE_OFF was added in 5.7.6. ", MySQLVersion:"5.7"},
    3113: definition.ErrorDefinition{ErrorNumber:3113, ErrorType:"ServerError", Symbol:"ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION", SQLState:"HY000", Message:"Cannot replicate anonymous transaction when AUTO_POSITION = 1, at file %s, position %lld.", Description:"ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION was added in 5.7.6. ", MySQLVersion:"5.7"},
    3114: definition.ErrorDefinition{ErrorNumber:3114, ErrorType:"ServerError", Symbol:"ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON", SQLState:"HY000", Message:"Cannot replicate anonymous transaction when @@GLOBAL.GTID_MODE = ON, at file %s, position %lld.", Description:"ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON was added in 5.7.6. ", MySQLVersion:"5.7"},
    3115: definition.ErrorDefinition{ErrorNumber:3115, ErrorType:"ServerError", Symbol:"ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF", SQLState:"HY000", Message:"Cannot replicate GTID-transaction when @@GLOBAL.GTID_MODE = OFF, at file %s, position %lld.", Description:"ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF was added in 5.7.6. ", MySQLVersion:"5.7"},
    3116: definition.ErrorDefinition{ErrorNumber:3116, ErrorType:"ServerError", Symbol:"ER_CANT_SET_ENFORCE_GTID_CONSISTENCY_ON_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS", SQLState:"HY000", Message:"Cannot set ENFORCE_GTID_CONSISTENCY = ON because there are ongoing transactions that violate GTID consistency.", Description:"ER_CANT_SET_ENFORCE_GTID_CONSISTENCY_ON_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS was added in 5.7.6. ", MySQLVersion:"5.7"},
    3117: definition.ErrorDefinition{ErrorNumber:3117, ErrorType:"ServerError", Symbol:"ER_SET_ENFORCE_GTID_CONSISTENCY_WARN_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS", SQLState:"HY000", Message:"There are ongoing transactions that violate GTID consistency.", Description:"ER_SET_ENFORCE_GTID_CONSISTENCY_WARN_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS was added in 5.7.6. ", MySQLVersion:"5.7"},
    3118: definition.ErrorDefinition{ErrorNumber:3118, ErrorType:"ServerError", Symbol:"ER_ACCOUNT_HAS_BEEN_LOCKED", SQLState:"HY000", Message:"Access denied for user '%s'@'%s'. Account is locked.", Description:"ER_ACCOUNT_HAS_BEEN_LOCKED was added in 5.7.6. ", MySQLVersion:"5.7"},
    3119: definition.ErrorDefinition{ErrorNumber:3119, ErrorType:"ServerError", Symbol:"ER_WRONG_TABLESPACE_NAME", SQLState:"42000", Message:"Incorrect tablespace name `%s`", Description:"ER_WRONG_TABLESPACE_NAME was added in 5.7.6. ", MySQLVersion:"5.7"},
    3120: definition.ErrorDefinition{ErrorNumber:3120, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_IS_NOT_EMPTY", SQLState:"HY000", Message:"Tablespace `%s` is not empty.", Description:"ER_TABLESPACE_IS_NOT_EMPTY was added in 5.7.6. ", MySQLVersion:"5.7"},
    3121: definition.ErrorDefinition{ErrorNumber:3121, ErrorType:"ServerError", Symbol:"ER_WRONG_FILE_NAME", SQLState:"HY000", Message:"Incorrect File Name '%s'.", Description:"ER_WRONG_FILE_NAME was added in 5.7.6. ", MySQLVersion:"5.7"},
    3122: definition.ErrorDefinition{ErrorNumber:3122, ErrorType:"ServerError", Symbol:"ER_BOOST_GEOMETRY_INCONSISTENT_TURNS_EXCEPTION", SQLState:"HY000", Message:"Inconsistent intersection points.", Description:"ER_BOOST_GEOMETRY_INCONSISTENT_TURNS_EXCEPTION was added in 5.7.7. ", MySQLVersion:"5.7"},
    3123: definition.ErrorDefinition{ErrorNumber:3123, ErrorType:"ServerError", Symbol:"ER_WARN_OPTIMIZER_HINT_SYNTAX_ERROR", SQLState:"HY000", Message:"Optimizer hint syntax error", Description:"ER_WARN_OPTIMIZER_HINT_SYNTAX_ERROR was added in 5.7.7. ", MySQLVersion:"5.7"},
    3124: definition.ErrorDefinition{ErrorNumber:3124, ErrorType:"ServerError", Symbol:"ER_WARN_BAD_MAX_EXECUTION_TIME", SQLState:"HY000", Message:"Unsupported MAX_EXECUTION_TIME", Description:"ER_WARN_BAD_MAX_EXECUTION_TIME was added in 5.7.7. ", MySQLVersion:"5.7"},
    3125: definition.ErrorDefinition{ErrorNumber:3125, ErrorType:"ServerError", Symbol:"ER_WARN_UNSUPPORTED_MAX_EXECUTION_TIME", SQLState:"HY000", Message:"MAX_EXECUTION_TIME hint is supported by top-level standalone SELECT statements only", Description:"ER_WARN_UNSUPPORTED_MAX_EXECUTION_TIME was added in 5.7.7. ", MySQLVersion:"5.7"},
    3126: definition.ErrorDefinition{ErrorNumber:3126, ErrorType:"ServerError", Symbol:"ER_WARN_CONFLICTING_HINT", SQLState:"HY000", Message:"Hint %s is ignored as conflicting/duplicated", Description:"ER_WARN_CONFLICTING_HINT was added in 5.7.7. ", MySQLVersion:"5.7"},
    3127: definition.ErrorDefinition{ErrorNumber:3127, ErrorType:"ServerError", Symbol:"ER_WARN_UNKNOWN_QB_NAME", SQLState:"HY000", Message:"Query block name %s is not found for %s hint", Description:"ER_WARN_UNKNOWN_QB_NAME was added in 5.7.7. ", MySQLVersion:"5.7"},
    3128: definition.ErrorDefinition{ErrorNumber:3128, ErrorType:"ServerError", Symbol:"ER_UNRESOLVED_HINT_NAME", SQLState:"HY000", Message:"Unresolved name %s for %s hint", Description:"ER_UNRESOLVED_HINT_NAME was added in 5.7.7. ", MySQLVersion:"5.7"},
    3129: definition.ErrorDefinition{ErrorNumber:3129, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SQLMODE_UNSET", SQLState:"HY000", Message:"Unsetting sql mode '%s' is deprecated. It will be made read-only in a future release.", Description:"ER_WARN_DEPRECATED_SQLMODE_UNSET was added in 5.7.7, removed after 5.7.7. ", MySQLVersion:"5.7"},
    3129: definition.ErrorDefinition{ErrorNumber:3129, ErrorType:"ServerError", Symbol:"ER_WARN_ON_MODIFYING_GTID_EXECUTED_TABLE", SQLState:"HY000", Message:"Please do not modify the %s table. This is a mysql internal system table to store GTIDs for committed transactions. Modifying it can lead to an inconsistent GTID state.", Description:"ER_WARN_ON_MODIFYING_GTID_EXECUTED_TABLE was added in 5.7.8. ", MySQLVersion:"5.7"},
    3130: definition.ErrorDefinition{ErrorNumber:3130, ErrorType:"ServerError", Symbol:"ER_PLUGGABLE_PROTOCOL_COMMAND_NOT_SUPPORTED", SQLState:"HY000", Message:"Command not supported by pluggable protocols", Description:"ER_PLUGGABLE_PROTOCOL_COMMAND_NOT_SUPPORTED was added in 5.7.8. ", MySQLVersion:"5.7"},
    3131: definition.ErrorDefinition{ErrorNumber:3131, ErrorType:"ServerError", Symbol:"ER_LOCKING_SERVICE_WRONG_NAME", SQLState:"42000", Message:"Incorrect locking service lock name '%s'.", Description:"ER_LOCKING_SERVICE_WRONG_NAME was added in 5.7.8. ", MySQLVersion:"5.7"},
    3132: definition.ErrorDefinition{ErrorNumber:3132, ErrorType:"ServerError", Symbol:"ER_LOCKING_SERVICE_DEADLOCK", SQLState:"HY000", Message:"Deadlock found when trying to get locking service lock", Description:"ER_LOCKING_SERVICE_DEADLOCK was added in 5.7.8. ", MySQLVersion:"5.7"},
    3133: definition.ErrorDefinition{ErrorNumber:3133, ErrorType:"ServerError", Symbol:"ER_LOCKING_SERVICE_TIMEOUT", SQLState:"HY000", Message:"Service lock wait timeout exceeded.", Description:"ER_LOCKING_SERVICE_TIMEOUT was added in 5.7.8. ", MySQLVersion:"5.7"},
    3134: definition.ErrorDefinition{ErrorNumber:3134, ErrorType:"ServerError", Symbol:"ER_GIS_MAX_POINTS_IN_GEOMETRY_OVERFLOWED", SQLState:"HY000", Message:"Parameter %s exceeds the maximum number of points in a geometry (%lu) in function %s.", Description:"ER_GIS_MAX_POINTS_IN_GEOMETRY_OVERFLOWED was added in 5.7.8. ", MySQLVersion:"5.7"},
    3135: definition.ErrorDefinition{ErrorNumber:3135, ErrorType:"ServerError", Symbol:"ER_SQL_MODE_MERGED", SQLState:"HY000", Message:"'NO_ZERO_DATE', 'NO_ZERO_IN_DATE' and 'ERROR_FOR_DIVISION_BY_ZERO' sql modes should be used with strict mode. They will be merged with strict mode in a future release.", Description:"ER_SQL_MODE_MERGED was added in 5.7.8. ", MySQLVersion:"5.7"},
    3136: definition.ErrorDefinition{ErrorNumber:3136, ErrorType:"ServerError", Symbol:"ER_VTOKEN_PLUGIN_TOKEN_MISMATCH", SQLState:"HY000", Message:"Version token mismatch for %.*s. Correct value %.*s", Description:"ER_VTOKEN_PLUGIN_TOKEN_MISMATCH was added in 5.7.8. ", MySQLVersion:"5.7"},
    3137: definition.ErrorDefinition{ErrorNumber:3137, ErrorType:"ServerError", Symbol:"ER_VTOKEN_PLUGIN_TOKEN_NOT_FOUND", SQLState:"HY000", Message:"Version token %.*s not found.", Description:"ER_VTOKEN_PLUGIN_TOKEN_NOT_FOUND was added in 5.7.8. ", MySQLVersion:"5.7"},
    3138: definition.ErrorDefinition{ErrorNumber:3138, ErrorType:"ServerError", Symbol:"ER_CANT_SET_VARIABLE_WHEN_OWNING_GTID", SQLState:"HY000", Message:"Variable %s cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK.", Description:"ER_CANT_SET_VARIABLE_WHEN_OWNING_GTID was added in 5.7.8. ", MySQLVersion:"5.7"},
    3139: definition.ErrorDefinition{ErrorNumber:3139, ErrorType:"ServerError", Symbol:"ER_SLAVE_CHANNEL_OPERATION_NOT_ALLOWED", SQLState:"HY000", Message:"%s cannot be performed on channel '%s'.", Description:"ER_SLAVE_CHANNEL_OPERATION_NOT_ALLOWED was added in 5.7.8. ", MySQLVersion:"5.7"},
    3140: definition.ErrorDefinition{ErrorNumber:3140, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_TEXT", SQLState:"22032", Message:"Invalid JSON text: \"%s\" at position %u in value for column '%s'.", Description:"ER_INVALID_JSON_TEXT was added in 5.7.8. ", MySQLVersion:"5.7"},
    3141: definition.ErrorDefinition{ErrorNumber:3141, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_TEXT_IN_PARAM", SQLState:"22032", Message:"Invalid JSON text in argument %u to function %s: \"%s\" at position %u.%s", Description:"ER_INVALID_JSON_TEXT_IN_PARAM was added in 5.7.8. ", MySQLVersion:"5.7"},
    3142: definition.ErrorDefinition{ErrorNumber:3142, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_BINARY_DATA", SQLState:"HY000", Message:"The JSON binary value contains invalid data.", Description:"ER_INVALID_JSON_BINARY_DATA was added in 5.7.8. ", MySQLVersion:"5.7"},
    3143: definition.ErrorDefinition{ErrorNumber:3143, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_PATH", SQLState:"42000", Message:"Invalid JSON path expression. The error is around character position %u.%s", Description:"ER_INVALID_JSON_PATH was added in 5.7.8. ", MySQLVersion:"5.7"},
    3144: definition.ErrorDefinition{ErrorNumber:3144, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_CHARSET", SQLState:"22032", Message:"Cannot create a JSON value from a string with CHARACTER SET '%s'.", Description:"ER_INVALID_JSON_CHARSET was added in 5.7.8. ", MySQLVersion:"5.7"},
    3145: definition.ErrorDefinition{ErrorNumber:3145, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_CHARSET_IN_FUNCTION", SQLState:"22032", Message:"Invalid JSON character data provided to function %s: '%s'", Description:"ER_INVALID_JSON_CHARSET_IN_FUNCTION was added in 5.7.8. ", MySQLVersion:"5.7"},
    3146: definition.ErrorDefinition{ErrorNumber:3146, ErrorType:"ServerError", Symbol:"ER_INVALID_TYPE_FOR_JSON", SQLState:"22032", Message:"Invalid data type for JSON data in argument %u to function %s", Description:"ER_INVALID_TYPE_FOR_JSON was added in 5.7.8. ", MySQLVersion:"5.7"},
    3147: definition.ErrorDefinition{ErrorNumber:3147, ErrorType:"ServerError", Symbol:"ER_INVALID_CAST_TO_JSON", SQLState:"22032", Message:"Cannot CAST value to JSON.", Description:"ER_INVALID_CAST_TO_JSON was added in 5.7.8. ", MySQLVersion:"5.7"},
    3148: definition.ErrorDefinition{ErrorNumber:3148, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_PATH_CHARSET", SQLState:"42000", Message:"A path expression must be encoded in the utf8 character set. The path expression '%s' is encoded in character set '%s'.", Description:"ER_INVALID_JSON_PATH_CHARSET was added in 5.7.8. ", MySQLVersion:"5.7"},
    3149: definition.ErrorDefinition{ErrorNumber:3149, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_PATH_WILDCARD", SQLState:"42000", Message:"In this situation, path expressions may not contain the * and ** tokens.", Description:"ER_INVALID_JSON_PATH_WILDCARD was added in 5.7.8. ", MySQLVersion:"5.7"},
    3150: definition.ErrorDefinition{ErrorNumber:3150, ErrorType:"ServerError", Symbol:"ER_JSON_VALUE_TOO_BIG", SQLState:"22032", Message:"The JSON value is too big to be stored in a JSON column.", Description:"ER_JSON_VALUE_TOO_BIG was added in 5.7.8. ", MySQLVersion:"5.7"},
    3151: definition.ErrorDefinition{ErrorNumber:3151, ErrorType:"ServerError", Symbol:"ER_JSON_KEY_TOO_BIG", SQLState:"22032", Message:"The JSON object contains a key name that is too long.", Description:"ER_JSON_KEY_TOO_BIG was added in 5.7.8. ", MySQLVersion:"5.7"},
    3152: definition.ErrorDefinition{ErrorNumber:3152, ErrorType:"ServerError", Symbol:"ER_JSON_USED_AS_KEY", SQLState:"42000", Message:"JSON column '%s' cannot be used in key specification.", Description:"ER_JSON_USED_AS_KEY was added in 5.7.8. ", MySQLVersion:"5.7"},
    3153: definition.ErrorDefinition{ErrorNumber:3153, ErrorType:"ServerError", Symbol:"ER_JSON_VACUOUS_PATH", SQLState:"42000", Message:"The path expression '$' is not allowed in this context.", Description:"ER_JSON_VACUOUS_PATH was added in 5.7.8. ", MySQLVersion:"5.7"},
    3154: definition.ErrorDefinition{ErrorNumber:3154, ErrorType:"ServerError", Symbol:"ER_JSON_BAD_ONE_OR_ALL_ARG", SQLState:"42000", Message:"The oneOrAll argument to %s may take these values: 'one' or 'all'.", Description:"ER_JSON_BAD_ONE_OR_ALL_ARG was added in 5.7.8. ", MySQLVersion:"5.7"},
    3155: definition.ErrorDefinition{ErrorNumber:3155, ErrorType:"ServerError", Symbol:"ER_NUMERIC_JSON_VALUE_OUT_OF_RANGE", SQLState:"22003", Message:"Out of range JSON value for CAST to %s%s from column %s at row %ld", Description:"ER_NUMERIC_JSON_VALUE_OUT_OF_RANGE was added in 5.7.8. ", MySQLVersion:"5.7"},
    3156: definition.ErrorDefinition{ErrorNumber:3156, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_VALUE_FOR_CAST", SQLState:"22018", Message:"Invalid JSON value for CAST to %s%s from column %s at row %ld", Description:"ER_INVALID_JSON_VALUE_FOR_CAST was added in 5.7.8. ", MySQLVersion:"5.7"},
    3157: definition.ErrorDefinition{ErrorNumber:3157, ErrorType:"ServerError", Symbol:"ER_JSON_DOCUMENT_TOO_DEEP", SQLState:"22032", Message:"The JSON document exceeds the maximum depth.", Description:"ER_JSON_DOCUMENT_TOO_DEEP was added in 5.7.8. ", MySQLVersion:"5.7"},
    3158: definition.ErrorDefinition{ErrorNumber:3158, ErrorType:"ServerError", Symbol:"ER_JSON_DOCUMENT_NULL_KEY", SQLState:"22032", Message:"JSON documents may not contain NULL member names.", Description:"ER_JSON_DOCUMENT_NULL_KEY was added in 5.7.8. ", MySQLVersion:"5.7"},
    3159: definition.ErrorDefinition{ErrorNumber:3159, ErrorType:"ServerError", Symbol:"ER_SECURE_TRANSPORT_REQUIRED", SQLState:"HY000", Message:"Connections using insecure transport are prohibited while --require_secure_transport=ON.", Description:"ER_SECURE_TRANSPORT_REQUIRED was added in 5.7.8. ", MySQLVersion:"5.7"},
    3160: definition.ErrorDefinition{ErrorNumber:3160, ErrorType:"ServerError", Symbol:"ER_NO_SECURE_TRANSPORTS_CONFIGURED", SQLState:"HY000", Message:"No secure transports (SSL or Shared Memory) are configured, unable to set --require_secure_transport=ON.", Description:"ER_NO_SECURE_TRANSPORTS_CONFIGURED was added in 5.7.8. ", MySQLVersion:"5.7"},
    3161: definition.ErrorDefinition{ErrorNumber:3161, ErrorType:"ServerError", Symbol:"ER_DISABLED_STORAGE_ENGINE", SQLState:"HY000", Message:"Storage engine %s is disabled (Table creation is disallowed).", Description:"ER_DISABLED_STORAGE_ENGINE was added in 5.7.8. ", MySQLVersion:"5.7"},
    3162: definition.ErrorDefinition{ErrorNumber:3162, ErrorType:"ServerError", Symbol:"ER_USER_DOES_NOT_EXIST", SQLState:"HY000", Message:"User %s does not exist.", Description:"ER_USER_DOES_NOT_EXIST was added in 5.7.8. ", MySQLVersion:"5.7"},
    3163: definition.ErrorDefinition{ErrorNumber:3163, ErrorType:"ServerError", Symbol:"ER_USER_ALREADY_EXISTS", SQLState:"HY000", Message:"User %s already exists.", Description:"ER_USER_ALREADY_EXISTS was added in 5.7.8. ", MySQLVersion:"5.7"},
    3164: definition.ErrorDefinition{ErrorNumber:3164, ErrorType:"ServerError", Symbol:"ER_AUDIT_API_ABORT", SQLState:"HY000", Message:"Aborted by Audit API ('%s'", Description:"ER_AUDIT_API_ABORT was added in 5.7.8. ", MySQLVersion:"5.7"},
    3165: definition.ErrorDefinition{ErrorNumber:3165, ErrorType:"ServerError", Symbol:"ER_INVALID_JSON_PATH_ARRAY_CELL", SQLState:"42000", Message:"A path expression is not a path to a cell in an array.", Description:"ER_INVALID_JSON_PATH_ARRAY_CELL was added in 5.7.8. ", MySQLVersion:"5.7"},
    3166: definition.ErrorDefinition{ErrorNumber:3166, ErrorType:"ServerError", Symbol:"ER_BUFPOOL_RESIZE_INPROGRESS", SQLState:"HY000", Message:"Another buffer pool resize is already in progress.", Description:"ER_BUFPOOL_RESIZE_INPROGRESS was added in 5.7.9. ", MySQLVersion:"5.7"},
    3167: definition.ErrorDefinition{ErrorNumber:3167, ErrorType:"ServerError", Symbol:"ER_FEATURE_DISABLED_SEE_DOC", SQLState:"HY000", Message:"The '%s' feature is disabled", Description:"ER_FEATURE_DISABLED_SEE_DOC was added in 5.7.9. ", MySQLVersion:"5.7"},
    3168: definition.ErrorDefinition{ErrorNumber:3168, ErrorType:"ServerError", Symbol:"ER_SERVER_ISNT_AVAILABLE", SQLState:"HY000", Message:"Server isn't available", Description:"ER_SERVER_ISNT_AVAILABLE was added in 5.7.9. ", MySQLVersion:"5.7"},
    3169: definition.ErrorDefinition{ErrorNumber:3169, ErrorType:"ServerError", Symbol:"ER_SESSION_WAS_KILLED", SQLState:"HY000", Message:"Session was killed", Description:"ER_SESSION_WAS_KILLED was added in 5.7.9. ", MySQLVersion:"5.7"},
    3170: definition.ErrorDefinition{ErrorNumber:3170, ErrorType:"ServerError", Symbol:"ER_CAPACITY_EXCEEDED", SQLState:"HY000", Message:"Memory capacity of %llu bytes for '%s' exceeded. %s", Description:"ER_CAPACITY_EXCEEDED was added in 5.7.9. ", MySQLVersion:"5.7"},
    3171: definition.ErrorDefinition{ErrorNumber:3171, ErrorType:"ServerError", Symbol:"ER_CAPACITY_EXCEEDED_IN_RANGE_OPTIMIZER", SQLState:"HY000", Message:"Range optimization was not done for this query.", Description:"ER_CAPACITY_EXCEEDED_IN_RANGE_OPTIMIZER was added in 5.7.9. ", MySQLVersion:"5.7"},
    3172: definition.ErrorDefinition{ErrorNumber:3172, ErrorType:"ServerError", Symbol:"ER_TABLE_NEEDS_UPG_PART", SQLState:"HY000", Message:"Partitioning upgrade required. Please dump/reload to fix it or do: ALTER TABLE `%s`.`%s` UPGRADE PARTITIONING", Description:"ER_TABLE_NEEDS_UPG_PART was added in 5.7.9. ", MySQLVersion:"5.7"},
    3173: definition.ErrorDefinition{ErrorNumber:3173, ErrorType:"ServerError", Symbol:"ER_CANT_WAIT_FOR_EXECUTED_GTID_SET_WHILE_OWNING_A_GTID", SQLState:"HY000", Message:"The client holds ownership of the GTID %s. Therefore, WAIT_FOR_EXECUTED_GTID_SET cannot wait for this GTID.", Description:"ER_CANT_WAIT_FOR_EXECUTED_GTID_SET_WHILE_OWNING_A_GTID was added in 5.7.9. ", MySQLVersion:"5.7"},
    3174: definition.ErrorDefinition{ErrorNumber:3174, ErrorType:"ServerError", Symbol:"ER_CANNOT_ADD_FOREIGN_BASE_COL_VIRTUAL", SQLState:"HY000", Message:"Cannot add foreign key on the base column of indexed virtual column.", Description:"ER_CANNOT_ADD_FOREIGN_BASE_COL_VIRTUAL was added in 5.7.10. ", MySQLVersion:"5.7"},
    3175: definition.ErrorDefinition{ErrorNumber:3175, ErrorType:"ServerError", Symbol:"ER_CANNOT_CREATE_VIRTUAL_INDEX_CONSTRAINT", SQLState:"HY000", Message:"Cannot create index on virtual column whose base column has foreign constraint.", Description:"ER_CANNOT_CREATE_VIRTUAL_INDEX_CONSTRAINT was added in 5.7.10. ", MySQLVersion:"5.7"},
    3176: definition.ErrorDefinition{ErrorNumber:3176, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_MODIFYING_GTID_EXECUTED_TABLE", SQLState:"HY000", Message:"Please do not modify the %s table with an XA transaction. This is an internal system table used to store GTIDs for committed transactions. Although modifying it can lead to an inconsistent GTID state, if neccessary you can modify it with a non-XA transaction.", Description:"ER_ERROR_ON_MODIFYING_GTID_EXECUTED_TABLE was added in 5.7.11. ", MySQLVersion:"5.7"},
    3177: definition.ErrorDefinition{ErrorNumber:3177, ErrorType:"ServerError", Symbol:"ER_LOCK_REFUSED_BY_ENGINE", SQLState:"HY000", Message:"Lock acquisition refused by storage engine.", Description:"ER_LOCK_REFUSED_BY_ENGINE was added in 5.7.11. ", MySQLVersion:"5.7"},
    3178: definition.ErrorDefinition{ErrorNumber:3178, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_ALTER_ONLINE_ON_VIRTUAL_COLUMN", SQLState:"HY000", Message:"ADD COLUMN col...VIRTUAL, ADD INDEX(col)", Description:"ER_UNSUPPORTED_ALTER_ONLINE_ON_VIRTUAL_COLUMN was added in 5.7.11. ", MySQLVersion:"5.7"},
    3179: definition.ErrorDefinition{ErrorNumber:3179, ErrorType:"ServerError", Symbol:"ER_MASTER_KEY_ROTATION_NOT_SUPPORTED_BY_SE", SQLState:"HY000", Message:"Master key rotation is not supported by storage engine.", Description:"ER_MASTER_KEY_ROTATION_NOT_SUPPORTED_BY_SE was added in 5.7.11. ", MySQLVersion:"5.7"},
    3180: definition.ErrorDefinition{ErrorNumber:3180, ErrorType:"ServerError", Symbol:"ER_MASTER_KEY_ROTATION_ERROR_BY_SE", SQLState:"HY000", Message:"Encryption key rotation error reported by SE: %s", Description:"ER_MASTER_KEY_ROTATION_ERROR_BY_SE was added in 5.7.11. ", MySQLVersion:"5.7"},
    3181: definition.ErrorDefinition{ErrorNumber:3181, ErrorType:"ServerError", Symbol:"ER_MASTER_KEY_ROTATION_BINLOG_FAILED", SQLState:"HY000", Message:"Write to binlog failed. However, master key rotation has been completed successfully.", Description:"ER_MASTER_KEY_ROTATION_BINLOG_FAILED was added in 5.7.11. ", MySQLVersion:"5.7"},
    3182: definition.ErrorDefinition{ErrorNumber:3182, ErrorType:"ServerError", Symbol:"ER_MASTER_KEY_ROTATION_SE_UNAVAILABLE", SQLState:"HY000", Message:"Storage engine is not available.", Description:"ER_MASTER_KEY_ROTATION_SE_UNAVAILABLE was added in 5.7.11. ", MySQLVersion:"5.7"},
    3183: definition.ErrorDefinition{ErrorNumber:3183, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_CANNOT_ENCRYPT", SQLState:"HY000", Message:"This tablespace can't be encrypted.", Description:"ER_TABLESPACE_CANNOT_ENCRYPT was added in 5.7.11. ", MySQLVersion:"5.7"},
    3184: definition.ErrorDefinition{ErrorNumber:3184, ErrorType:"ServerError", Symbol:"ER_INVALID_ENCRYPTION_OPTION", SQLState:"HY000", Message:"Invalid encryption option.", Description:"ER_INVALID_ENCRYPTION_OPTION was added in 5.7.11. ", MySQLVersion:"5.7"},
    3185: definition.ErrorDefinition{ErrorNumber:3185, ErrorType:"ServerError", Symbol:"ER_CANNOT_FIND_KEY_IN_KEYRING", SQLState:"HY000", Message:"Can't find master key from keyring, please check in the server log if a keyring plugin is loaded and initialized successfully.", Description:"ER_CANNOT_FIND_KEY_IN_KEYRING was added in 5.7.11. ", MySQLVersion:"5.7"},
    3186: definition.ErrorDefinition{ErrorNumber:3186, ErrorType:"ServerError", Symbol:"ER_CAPACITY_EXCEEDED_IN_PARSER", SQLState:"HY000", Message:"Parser bailed out for this query.", Description:"ER_CAPACITY_EXCEEDED_IN_PARSER was added in 5.7.12. ", MySQLVersion:"5.7"},
    3187: definition.ErrorDefinition{ErrorNumber:3187, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_ALTER_ENCRYPTION_INPLACE", SQLState:"HY000", Message:"Cannot alter encryption attribute by inplace algorithm.", Description:"ER_UNSUPPORTED_ALTER_ENCRYPTION_INPLACE was added in 5.7.13. ", MySQLVersion:"5.7"},
    3188: definition.ErrorDefinition{ErrorNumber:3188, ErrorType:"ServerError", Symbol:"ER_KEYRING_UDF_KEYRING_SERVICE_ERROR", SQLState:"HY000", Message:"Function '%s' failed because underlying keyring service returned an error. Please check if a keyring plugin is installed and that provided arguments are valid for the keyring you are using.", Description:"ER_KEYRING_UDF_KEYRING_SERVICE_ERROR was added in 5.7.13. ", MySQLVersion:"5.7"},
    3189: definition.ErrorDefinition{ErrorNumber:3189, ErrorType:"ServerError", Symbol:"ER_USER_COLUMN_OLD_LENGTH", SQLState:"HY000", Message:"It seems that your db schema is old. The %s column is 77 characters long and should be 93 characters long. Please run mysql_upgrade.", Description:"ER_USER_COLUMN_OLD_LENGTH was added in 5.7.13. ", MySQLVersion:"5.7"},
    3190: definition.ErrorDefinition{ErrorNumber:3190, ErrorType:"ServerError", Symbol:"ER_CANT_RESET_MASTER", SQLState:"HY000", Message:"RESET MASTER is not allowed because %s.", Description:"ER_CANT_RESET_MASTER was added in 5.7.14. ", MySQLVersion:"5.7"},
    3191: definition.ErrorDefinition{ErrorNumber:3191, ErrorType:"ServerError", Symbol:"ER_GROUP_REPLICATION_MAX_GROUP_SIZE", SQLState:"HY000", Message:"The START GROUP_REPLICATION command failed since the group already has 9 members.", Description:"ER_GROUP_REPLICATION_MAX_GROUP_SIZE was added in 5.7.14. ", MySQLVersion:"5.7"},
    3192: definition.ErrorDefinition{ErrorNumber:3192, ErrorType:"ServerError", Symbol:"ER_CANNOT_ADD_FOREIGN_BASE_COL_STORED", SQLState:"HY000", Message:"Cannot add foreign key on the base column of stored column.", Description:"ER_CANNOT_ADD_FOREIGN_BASE_COL_STORED was added in 5.7.14. ", MySQLVersion:"5.7"},
    3193: definition.ErrorDefinition{ErrorNumber:3193, ErrorType:"ServerError", Symbol:"ER_TABLE_REFERENCED", SQLState:"HY000", Message:"Cannot complete the operation because table is referenced by another connection.", Description:"ER_TABLE_REFERENCED was added in 5.7.14. ", MySQLVersion:"5.7"},
    3194: definition.ErrorDefinition{ErrorNumber:3194, ErrorType:"ServerError", Symbol:"ER_PARTITION_ENGINE_DEPRECATED_FOR_TABLE", SQLState:"HY000", Message:"The partition engine, used by table '%s.%s', is deprecated and will be removed in a future release. Please use native partitioning instead.", Description:"ER_PARTITION_ENGINE_DEPRECATED_FOR_TABLE was added in 5.7.17. ", MySQLVersion:"5.7"},
    3195: definition.ErrorDefinition{ErrorNumber:3195, ErrorType:"ServerError", Symbol:"ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID_ZERO", SQLState:"01000", Message:"%s(geometry) is deprecated and will be replaced by st_srid(geometry, 0) in a future version. Use %s(st_aswkb(geometry), 0) instead.", Description:"ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID_ZERO was added in 5.7.19. ", MySQLVersion:"5.7"},
    3196: definition.ErrorDefinition{ErrorNumber:3196, ErrorType:"ServerError", Symbol:"ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID", SQLState:"01000", Message:"%s(geometry, srid) is deprecated and will be replaced by st_srid(geometry, srid) in a future version. Use %s(st_aswkb(geometry), srid) instead.", Description:"ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID was added in 5.7.19. ", MySQLVersion:"5.7"},
    3197: definition.ErrorDefinition{ErrorNumber:3197, ErrorType:"ServerError", Symbol:"ER_XA_RETRY", SQLState:"HY000", Message:"The resource manager is not able to commit the transaction branch at this time. Please retry later.", Description:"ER_XA_RETRY was added in 5.7.19. ", MySQLVersion:"5.7"},
    3198: definition.ErrorDefinition{ErrorNumber:3198, ErrorType:"ServerError", Symbol:"ER_KEYRING_AWS_UDF_AWS_KMS_ERROR", SQLState:"HY000", Message:"Function %s failed due to: %s.", Description:"ER_KEYRING_AWS_UDF_AWS_KMS_ERROR was added in 5.7.19. ", MySQLVersion:"5.7"},
    3199: definition.ErrorDefinition{ErrorNumber:3199, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_XA", SQLState:"HY000", Message:"Statement is unsafe because it is being used inside a XA transaction. Concurrent XA transactions may deadlock on slaves when replicated using statements.", Description:"ER_BINLOG_UNSAFE_XA was added in 5.7.20. ", MySQLVersion:"5.7"},
    3200: definition.ErrorDefinition{ErrorNumber:3200, ErrorType:"ServerError", Symbol:"ER_UDF_ERROR", SQLState:"HY000", Message:"%s UDF failed", Description:"ER_UDF_ERROR was added in 5.7.21. ", MySQLVersion:"5.7"},
    3201: definition.ErrorDefinition{ErrorNumber:3201, ErrorType:"ServerError", Symbol:"ER_KEYRING_MIGRATION_FAILURE", SQLState:"HY000", Message:"Can not perform keyring migration : %s", Description:"ER_KEYRING_MIGRATION_FAILURE was added in 5.7.21. ", MySQLVersion:"5.7"},
    3202: definition.ErrorDefinition{ErrorNumber:3202, ErrorType:"ServerError", Symbol:"ER_KEYRING_ACCESS_DENIED_ERROR", SQLState:"42000", Message:"Access denied", Description:"ER_KEYRING_ACCESS_DENIED_ERROR was added in 5.7.21. ", MySQLVersion:"5.7"},
    3203: definition.ErrorDefinition{ErrorNumber:3203, ErrorType:"ServerError", Symbol:"ER_KEYRING_MIGRATION_STATUS", SQLState:"HY000", Message:"Keyring migration %s.", Description:"ER_KEYRING_MIGRATION_STATUS was added in 5.7.21. ", MySQLVersion:"5.7"},
    3204: definition.ErrorDefinition{ErrorNumber:3204, ErrorType:"ServerError", Symbol:"ER_PLUGIN_FAILED_TO_OPEN_TABLES", SQLState:"HY000", Message:"Failed to open the %s filter tables.", Description:"ER_PLUGIN_FAILED_TO_OPEN_TABLES was added in 5.7.22. ", MySQLVersion:"5.7"},
    3205: definition.ErrorDefinition{ErrorNumber:3205, ErrorType:"ServerError", Symbol:"ER_PLUGIN_FAILED_TO_OPEN_TABLE", SQLState:"HY000", Message:"Failed to open '%s.%s' %s table.", Description:"ER_PLUGIN_FAILED_TO_OPEN_TABLE was added in 5.7.22. ", MySQLVersion:"5.7"},
    3206: definition.ErrorDefinition{ErrorNumber:3206, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED", SQLState:"HY000", Message:"No keyring plugin installed.", Description:"ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED was added in 5.7.22. ", MySQLVersion:"5.7"},
    3207: definition.ErrorDefinition{ErrorNumber:3207, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET", SQLState:"HY000", Message:"Audit log encryption password has not been set", Description:"ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET was added in 5.7.22. ", MySQLVersion:"5.7"},
    3208: definition.ErrorDefinition{ErrorNumber:3208, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY", SQLState:"HY000", Message:"Could not create AES key. OpenSSL's EVP_BytesToKey function failed.", Description:"ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY was added in 5.7.22. ", MySQLVersion:"5.7"},
    3209: definition.ErrorDefinition{ErrorNumber:3209, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED", SQLState:"HY000", Message:"Audit log encryption password cannot be fetched from the keyring. Password used so far is used for encryption.", Description:"ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED was added in 5.7.22. ", MySQLVersion:"5.7"},
    3210: definition.ErrorDefinition{ErrorNumber:3210, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED", SQLState:"HY000", Message:"Audit Log filtering has not been installed.", Description:"ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED was added in 5.7.22. ", MySQLVersion:"5.7"},
    3211: definition.ErrorDefinition{ErrorNumber:3211, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE", SQLState:"HY000", Message:"Request ignored for '%s'@'%s'. SUPER_ACL needed to perform operation", Description:"ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE was added in 5.7.22. ", MySQLVersion:"5.7"},
    3212: definition.ErrorDefinition{ErrorNumber:3212, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED", SQLState:"HY000", Message:"SUPER privilege required for '%s'@'%s' user.", Description:"ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED was added in 5.7.22. ", MySQLVersion:"5.7"},
    3213: definition.ErrorDefinition{ErrorNumber:3213, ErrorType:"ServerError", Symbol:"ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS", SQLState:"HY000", Message:"Could not reinitialize audit log filters.", Description:"ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS was added in 5.7.22. ", MySQLVersion:"5.7"},
    3214: definition.ErrorDefinition{ErrorNumber:3214, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE", SQLState:"HY000", Message:"Invalid argument type", Description:"ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE was added in 5.7.22. ", MySQLVersion:"5.7"},
    3215: definition.ErrorDefinition{ErrorNumber:3215, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT", SQLState:"HY000", Message:"Invalid argument count", Description:"ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT was added in 5.7.22. ", MySQLVersion:"5.7"},
    3216: definition.ErrorDefinition{ErrorNumber:3216, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED", SQLState:"HY000", Message:"audit_log plugin has not been installed using INSTALL PLUGIN syntax.", Description:"ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED was added in 5.7.22. ", MySQLVersion:"5.7"},
    3217: definition.ErrorDefinition{ErrorNumber:3217, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE", SQLState:"HY000", Message:"Invalid \"max_array_length\" argument type.", Description:"ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE was added in 5.7.22. ", MySQLVersion:"5.7"},
    3218: definition.ErrorDefinition{ErrorNumber:3218, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_VALUE", SQLState:"HY000", Message:"Invalid \"max_array_length\" argument value.", Description:"ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_VALUE was added in 5.7.22. ", MySQLVersion:"5.7"},
    3219: definition.ErrorDefinition{ErrorNumber:3219, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR", SQLState:"HY000", Message:"%s", Description:"ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR was added in 5.7.22. ", MySQLVersion:"5.7"},
    3220: definition.ErrorDefinition{ErrorNumber:3220, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY", SQLState:"HY000", Message:"Filter name cannot be empty.", Description:"ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY was added in 5.7.22. ", MySQLVersion:"5.7"},
    3221: definition.ErrorDefinition{ErrorNumber:3221, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY", SQLState:"HY000", Message:"User cannot be empty.", Description:"ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY was added in 5.7.22. ", MySQLVersion:"5.7"},
    3222: definition.ErrorDefinition{ErrorNumber:3222, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXISTS", SQLState:"HY000", Message:"Specified filter has not been found.", Description:"ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXISTS was added in 5.7.22. ", MySQLVersion:"5.7"},
    3223: definition.ErrorDefinition{ErrorNumber:3223, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC", SQLState:"HY000", Message:"First character of the user name must be alphanumeric.", Description:"ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC was added in 5.7.22. ", MySQLVersion:"5.7"},
    3224: definition.ErrorDefinition{ErrorNumber:3224, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER", SQLState:"HY000", Message:"Invalid character in the user name.", Description:"ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER was added in 5.7.22. ", MySQLVersion:"5.7"},
    3225: definition.ErrorDefinition{ErrorNumber:3225, ErrorType:"ServerError", Symbol:"ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER", SQLState:"HY000", Message:"Invalid character in the host name.", Description:"ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER was added in 5.7.22. ", MySQLVersion:"5.7"},
    3226: definition.ErrorDefinition{ErrorNumber:3226, ErrorType:"ServerError", Symbol:"WARN_DEPRECATED_MAXDB_SQL_MODE_FOR_TIMESTAMP", SQLState:"HY000", Message:"With the MAXDB SQL mode enabled, TIMESTAMP is identical with DATETIME. The MAXDB SQL mode is deprecated and will be removed in a future release. Please disable the MAXDB SQL mode and use DATETIME instead.", Description:"WARN_DEPRECATED_MAXDB_SQL_MODE_FOR_TIMESTAMP was added in 5.7.22. ", MySQLVersion:"5.7"},
    3227: definition.ErrorDefinition{ErrorNumber:3227, ErrorType:"ServerError", Symbol:"ER_XA_REPLICATION_FILTERS", SQLState:"HY000", Message:"The use of replication filters with XA transactions is not supported, and can lead to an undefined state in the replication slave.", Description:"ER_XA_REPLICATION_FILTERS was added in 5.7.23. ", MySQLVersion:"5.7"},
    3228: definition.ErrorDefinition{ErrorNumber:3228, ErrorType:"ServerError", Symbol:"ER_CANT_OPEN_ERROR_LOG", SQLState:"HY000", Message:"Could not open file '%s' for error logging%s%s", Description:"ER_CANT_OPEN_ERROR_LOG was added in 5.7.24. ", MySQLVersion:"5.7"},
    3229: definition.ErrorDefinition{ErrorNumber:3229, ErrorType:"ServerError", Symbol:"ER_GROUPING_ON_TIMESTAMP_IN_DST", SQLState:"HY000", Message:"Grouping on temporal is non-deterministic for timezones having DST. Please consider switching to UTC for this query.", Description:"ER_GROUPING_ON_TIMESTAMP_IN_DST was added in 5.7.27. ", MySQLVersion:"5.7"},
    3230: definition.ErrorDefinition{ErrorNumber:3230, ErrorType:"ServerError", Symbol:"ER_CANT_START_SERVER_NAMED_PIPE", SQLState:"HY000", Message:"Can't start server : Named Pipe \"%s\" already in use.", Description:"ER_CANT_START_SERVER_NAMED_PIPE was added in 5.7.27.", MySQLVersion:"5.7"},
    2000: definition.ErrorDefinition{ErrorNumber:2000, ErrorType:"ClientError", Symbol:"CR_UNKNOWN_ERROR", SQLState:"", Message:"Unknown MySQL error ", Description:"", MySQLVersion:"5.7"},
    2001: definition.ErrorDefinition{ErrorNumber:2001, ErrorType:"ClientError", Symbol:"CR_SOCKET_CREATE_ERROR", SQLState:"", Message:"Can't create UNIX socket (%d) ", Description:"", MySQLVersion:"5.7"},
    2002: definition.ErrorDefinition{ErrorNumber:2002, ErrorType:"ClientError", Symbol:"CR_CONNECTION_ERROR", SQLState:"", Message:"Can't connect to local MySQL server through socket '%s' (%d) ", Description:"", MySQLVersion:"5.7"},
    2003: definition.ErrorDefinition{ErrorNumber:2003, ErrorType:"ClientError", Symbol:"CR_CONN_HOST_ERROR", SQLState:"", Message:"Can't connect to MySQL server on '%s' (%d) ", Description:"", MySQLVersion:"5.7"},
    2004: definition.ErrorDefinition{ErrorNumber:2004, ErrorType:"ClientError", Symbol:"CR_IPSOCK_ERROR", SQLState:"", Message:"Can't create TCP/IP socket (%d) ", Description:"", MySQLVersion:"5.7"},
    2005: definition.ErrorDefinition{ErrorNumber:2005, ErrorType:"ClientError", Symbol:"CR_UNKNOWN_HOST", SQLState:"", Message:"Unknown MySQL server host '%s' (%d) ", Description:"", MySQLVersion:"5.7"},
    2006: definition.ErrorDefinition{ErrorNumber:2006, ErrorType:"ClientError", Symbol:"CR_SERVER_GONE_ERROR", SQLState:"", Message:"MySQL server has gone away ", Description:"", MySQLVersion:"5.7"},
    2007: definition.ErrorDefinition{ErrorNumber:2007, ErrorType:"ClientError", Symbol:"CR_VERSION_ERROR", SQLState:"", Message:"Protocol mismatch", Description:"server version = %d, client version = %d ", MySQLVersion:"5.7"},
    2008: definition.ErrorDefinition{ErrorNumber:2008, ErrorType:"ClientError", Symbol:"CR_OUT_OF_MEMORY", SQLState:"", Message:"MySQL client ran out of memory ", Description:"", MySQLVersion:"5.7"},
    2009: definition.ErrorDefinition{ErrorNumber:2009, ErrorType:"ClientError", Symbol:"CR_WRONG_HOST_INFO", SQLState:"", Message:"Wrong host info ", Description:"", MySQLVersion:"5.7"},
    2010: definition.ErrorDefinition{ErrorNumber:2010, ErrorType:"ClientError", Symbol:"CR_LOCALHOST_CONNECTION", SQLState:"", Message:"Localhost via UNIX socket ", Description:"", MySQLVersion:"5.7"},
    2011: definition.ErrorDefinition{ErrorNumber:2011, ErrorType:"ClientError", Symbol:"CR_TCP_CONNECTION", SQLState:"", Message:"%s via TCP/IP ", Description:"", MySQLVersion:"5.7"},
    2012: definition.ErrorDefinition{ErrorNumber:2012, ErrorType:"ClientError", Symbol:"CR_SERVER_HANDSHAKE_ERR", SQLState:"", Message:"Error in server handshake ", Description:"", MySQLVersion:"5.7"},
    2013: definition.ErrorDefinition{ErrorNumber:2013, ErrorType:"ClientError", Symbol:"CR_SERVER_LOST", SQLState:"", Message:"Lost connection to MySQL server during query ", Description:"", MySQLVersion:"5.7"},
    2014: definition.ErrorDefinition{ErrorNumber:2014, ErrorType:"ClientError", Symbol:"CR_COMMANDS_OUT_OF_SYNC", SQLState:"", Message:"Commands out of sync", Description:"Commands were executed in an improper order. This error occurs when a function is called that is not appropriate for the current state of the connection. For example, if mysql_stmt_fetch() is not called enough times to read an entire result set (that is, enough times to return MYSQL_NO_DATA), this error may occur for the following C API call. ", MySQLVersion:"5.7"},
    2015: definition.ErrorDefinition{ErrorNumber:2015, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPE_CONNECTION", SQLState:"", Message:"Named pipe: %s ", Description:"", MySQLVersion:"5.7"},
    2016: definition.ErrorDefinition{ErrorNumber:2016, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPEWAIT_ERROR", SQLState:"", Message:"Can't wait for named pipe to host: %s pipe: %s (%lu) ", Description:"", MySQLVersion:"5.7"},
    2017: definition.ErrorDefinition{ErrorNumber:2017, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPEOPEN_ERROR", SQLState:"", Message:"Can't open named pipe to host: %s pipe: %s (%lu) ", Description:"", MySQLVersion:"5.7"},
    2018: definition.ErrorDefinition{ErrorNumber:2018, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPESETSTATE_ERROR", SQLState:"", Message:"Can't set state of named pipe to host: %s pipe: %s (%lu) ", Description:"", MySQLVersion:"5.7"},
    2019: definition.ErrorDefinition{ErrorNumber:2019, ErrorType:"ClientError", Symbol:"CR_CANT_READ_CHARSET", SQLState:"", Message:"Can't initialize character set %s (path: %s) ", Description:"", MySQLVersion:"5.7"},
    2020: definition.ErrorDefinition{ErrorNumber:2020, ErrorType:"ClientError", Symbol:"CR_NET_PACKET_TOO_LARGE", SQLState:"", Message:"Got packet bigger than 'max_allowed_packet' bytes ", Description:"", MySQLVersion:"5.7"},
    2021: definition.ErrorDefinition{ErrorNumber:2021, ErrorType:"ClientError", Symbol:"CR_EMBEDDED_CONNECTION", SQLState:"", Message:"Embedded server ", Description:"", MySQLVersion:"5.7"},
    2022: definition.ErrorDefinition{ErrorNumber:2022, ErrorType:"ClientError", Symbol:"CR_PROBE_SLAVE_STATUS", SQLState:"", Message:"Error on SHOW SLAVE STATUS: ", Description:"", MySQLVersion:"5.7"},
    2023: definition.ErrorDefinition{ErrorNumber:2023, ErrorType:"ClientError", Symbol:"CR_PROBE_SLAVE_HOSTS", SQLState:"", Message:"Error on SHOW SLAVE HOSTS: ", Description:"", MySQLVersion:"5.7"},
    2024: definition.ErrorDefinition{ErrorNumber:2024, ErrorType:"ClientError", Symbol:"CR_PROBE_SLAVE_CONNECT", SQLState:"", Message:"Error connecting to slave: ", Description:"", MySQLVersion:"5.7"},
    2025: definition.ErrorDefinition{ErrorNumber:2025, ErrorType:"ClientError", Symbol:"CR_PROBE_MASTER_CONNECT", SQLState:"", Message:"Error connecting to master: ", Description:"", MySQLVersion:"5.7"},
    2026: definition.ErrorDefinition{ErrorNumber:2026, ErrorType:"ClientError", Symbol:"CR_SSL_CONNECTION_ERROR", SQLState:"", Message:"SSL connection error: %s ", Description:"", MySQLVersion:"5.7"},
    2027: definition.ErrorDefinition{ErrorNumber:2027, ErrorType:"ClientError", Symbol:"CR_MALFORMED_PACKET", SQLState:"", Message:"Malformed packet ", Description:"", MySQLVersion:"5.7"},
    2028: definition.ErrorDefinition{ErrorNumber:2028, ErrorType:"ClientError", Symbol:"CR_WRONG_LICENSE", SQLState:"", Message:"This client library is licensed only for use with MySQL servers having '%s' license ", Description:"", MySQLVersion:"5.7"},
    2029: definition.ErrorDefinition{ErrorNumber:2029, ErrorType:"ClientError", Symbol:"CR_NULL_POINTER", SQLState:"", Message:"Invalid use of null pointer ", Description:"", MySQLVersion:"5.7"},
    2030: definition.ErrorDefinition{ErrorNumber:2030, ErrorType:"ClientError", Symbol:"CR_NO_PREPARE_STMT", SQLState:"", Message:"Statement not prepared ", Description:"", MySQLVersion:"5.7"},
    2031: definition.ErrorDefinition{ErrorNumber:2031, ErrorType:"ClientError", Symbol:"CR_PARAMS_NOT_BOUND", SQLState:"", Message:"No data supplied for parameters in prepared statement ", Description:"", MySQLVersion:"5.7"},
    2032: definition.ErrorDefinition{ErrorNumber:2032, ErrorType:"ClientError", Symbol:"CR_DATA_TRUNCATED", SQLState:"", Message:"Data truncated ", Description:"", MySQLVersion:"5.7"},
    2033: definition.ErrorDefinition{ErrorNumber:2033, ErrorType:"ClientError", Symbol:"CR_NO_PARAMETERS_EXISTS", SQLState:"", Message:"No parameters exist in the statement ", Description:"", MySQLVersion:"5.7"},
    2034: definition.ErrorDefinition{ErrorNumber:2034, ErrorType:"ClientError", Symbol:"CR_INVALID_PARAMETER_NO", SQLState:"", Message:"Invalid parameter number", Description:"A key name was empty or the amount of connection attribute data for mysql_options4() exceeds the 64KB limit. ", MySQLVersion:"5.7"},
    2035: definition.ErrorDefinition{ErrorNumber:2035, ErrorType:"ClientError", Symbol:"CR_INVALID_BUFFER_USE", SQLState:"", Message:"Can't send long data for non-string/non-binary data types (parameter: %d) ", Description:"", MySQLVersion:"5.7"},
    2036: definition.ErrorDefinition{ErrorNumber:2036, ErrorType:"ClientError", Symbol:"CR_UNSUPPORTED_PARAM_TYPE", SQLState:"", Message:"Using unsupported buffer type: %d (parameter: %d) ", Description:"", MySQLVersion:"5.7"},
    2037: definition.ErrorDefinition{ErrorNumber:2037, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECTION", SQLState:"", Message:"Shared memory: %s ", Description:"", MySQLVersion:"5.7"},
    2038: definition.ErrorDefinition{ErrorNumber:2038, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_REQUEST_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not create request event (%lu) ", MySQLVersion:"5.7"},
    2039: definition.ErrorDefinition{ErrorNumber:2039, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_ANSWER_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"no answer event received from server (%lu) ", MySQLVersion:"5.7"},
    2040: definition.ErrorDefinition{ErrorNumber:2040, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_FILE_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"server could not allocate file mapping (%lu) ", MySQLVersion:"5.7"},
    2041: definition.ErrorDefinition{ErrorNumber:2041, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"server could not get pointer to file mapping (%lu) ", MySQLVersion:"5.7"},
    2042: definition.ErrorDefinition{ErrorNumber:2042, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_FILE_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not allocate file mapping (%lu) ", MySQLVersion:"5.7"},
    2043: definition.ErrorDefinition{ErrorNumber:2043, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not get pointer to file mapping (%lu) ", MySQLVersion:"5.7"},
    2044: definition.ErrorDefinition{ErrorNumber:2044, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_EVENT_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not create %s event (%lu) ", MySQLVersion:"5.7"},
    2045: definition.ErrorDefinition{ErrorNumber:2045, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_ABANDONED_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"no answer from server (%lu) ", MySQLVersion:"5.7"},
    2046: definition.ErrorDefinition{ErrorNumber:2046, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_SET_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"cannot send request event to server (%lu) ", MySQLVersion:"5.7"},
    2047: definition.ErrorDefinition{ErrorNumber:2047, ErrorType:"ClientError", Symbol:"CR_CONN_UNKNOW_PROTOCOL", SQLState:"", Message:"Wrong or unknown protocol ", Description:"", MySQLVersion:"5.7"},
    2048: definition.ErrorDefinition{ErrorNumber:2048, ErrorType:"ClientError", Symbol:"CR_INVALID_CONN_HANDLE", SQLState:"", Message:"Invalid connection handle ", Description:"", MySQLVersion:"5.7"},
    2049: definition.ErrorDefinition{ErrorNumber:2049, ErrorType:"ClientError", Symbol:"CR_SECURE_AUTH", SQLState:"", Message:"Connection using old (pre-4.1.1) authentication protocol refused (client option 'secure_auth' enabled)", Description:"CR_SECURE_AUTH was removed after 5.7.4. ", MySQLVersion:"5.7"},
    2049: definition.ErrorDefinition{ErrorNumber:2049, ErrorType:"ClientError", Symbol:"CR_UNUSED_1", SQLState:"", Message:"Connection using old (pre-4.1.1) authentication protocol refused (client option 'secure_auth' enabled)", Description:"CR_UNUSED_1 was added in 5.7.5. ", MySQLVersion:"5.7"},
    2050: definition.ErrorDefinition{ErrorNumber:2050, ErrorType:"ClientError", Symbol:"CR_FETCH_CANCELED", SQLState:"", Message:"Row retrieval was canceled by mysql_stmt_close() call ", Description:"", MySQLVersion:"5.7"},
    2051: definition.ErrorDefinition{ErrorNumber:2051, ErrorType:"ClientError", Symbol:"CR_NO_DATA", SQLState:"", Message:"Attempt to read column without prior row fetch ", Description:"", MySQLVersion:"5.7"},
    2052: definition.ErrorDefinition{ErrorNumber:2052, ErrorType:"ClientError", Symbol:"CR_NO_STMT_METADATA", SQLState:"", Message:"Prepared statement contains no metadata ", Description:"", MySQLVersion:"5.7"},
    2053: definition.ErrorDefinition{ErrorNumber:2053, ErrorType:"ClientError", Symbol:"CR_NO_RESULT_SET", SQLState:"", Message:"Attempt to read a row while there is no result set associated with the statement ", Description:"", MySQLVersion:"5.7"},
    2054: definition.ErrorDefinition{ErrorNumber:2054, ErrorType:"ClientError", Symbol:"CR_NOT_IMPLEMENTED", SQLState:"", Message:"This feature is not implemented yet ", Description:"", MySQLVersion:"5.7"},
    2055: definition.ErrorDefinition{ErrorNumber:2055, ErrorType:"ClientError", Symbol:"CR_SERVER_LOST_EXTENDED", SQLState:"", Message:"Lost connection to MySQL server at '%s', system error: %d ", Description:"", MySQLVersion:"5.7"},
    2056: definition.ErrorDefinition{ErrorNumber:2056, ErrorType:"ClientError", Symbol:"CR_STMT_CLOSED", SQLState:"", Message:"Statement closed indirectly because of a preceding %s() call ", Description:"", MySQLVersion:"5.7"},
    2057: definition.ErrorDefinition{ErrorNumber:2057, ErrorType:"ClientError", Symbol:"CR_NEW_STMT_METADATA", SQLState:"", Message:"The number of columns in the result set differs from the number of bound buffers. You must reset the statement, rebind the result set columns, and execute the statement again ", Description:"", MySQLVersion:"5.7"},
    2058: definition.ErrorDefinition{ErrorNumber:2058, ErrorType:"ClientError", Symbol:"CR_ALREADY_CONNECTED", SQLState:"", Message:"This handle is already connected. Use a separate handle for each connection. ", Description:"", MySQLVersion:"5.7"},
    2059: definition.ErrorDefinition{ErrorNumber:2059, ErrorType:"ClientError", Symbol:"CR_AUTH_PLUGIN_CANNOT_LOAD", SQLState:"", Message:"Authentication plugin '%s' cannot be loaded: %s ", Description:"", MySQLVersion:"5.7"},
    2060: definition.ErrorDefinition{ErrorNumber:2060, ErrorType:"ClientError", Symbol:"CR_DUPLICATE_CONNECTION_ATTR", SQLState:"", Message:"There is an attribute with the same name already", Description:"A duplicate connection attribute name was specified for mysql_options4(). ", MySQLVersion:"5.7"},
    2061: definition.ErrorDefinition{ErrorNumber:2061, ErrorType:"ClientError", Symbol:"CR_AUTH_PLUGIN_ERR", SQLState:"", Message:"Authentication plugin '%s' reported error: %s", Description:"CR_AUTH_PLUGIN_ERR was added in 5.7.1. ", MySQLVersion:"5.7"},
    2062: definition.ErrorDefinition{ErrorNumber:2062, ErrorType:"ClientError", Symbol:"CR_INSECURE_API_ERR", SQLState:"", Message:"Insecure API function call: '%s' Use instead: '%s'", Description:"CR_INSECURE_API_ERR was added in 5.7.6.", MySQLVersion:"5.7"},
    1: definition.ErrorDefinition{ErrorNumber:1, ErrorType:"GlobalError", Symbol:"EE_CANTCREATEFILE", SQLState:"", Message:"Can't create/write to file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    2: definition.ErrorDefinition{ErrorNumber:2, ErrorType:"GlobalError", Symbol:"EE_READ", SQLState:"", Message:"Error reading file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    3: definition.ErrorDefinition{ErrorNumber:3, ErrorType:"GlobalError", Symbol:"EE_WRITE", SQLState:"", Message:"Error writing file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    4: definition.ErrorDefinition{ErrorNumber:4, ErrorType:"GlobalError", Symbol:"EE_BADCLOSE", SQLState:"", Message:"Error on close of '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    5: definition.ErrorDefinition{ErrorNumber:5, ErrorType:"GlobalError", Symbol:"EE_OUTOFMEMORY", SQLState:"", Message:"Out of memory (Needed %u bytes) ", Description:"", MySQLVersion:"5.7"},
    6: definition.ErrorDefinition{ErrorNumber:6, ErrorType:"GlobalError", Symbol:"EE_DELETE", SQLState:"", Message:"Error on delete of '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    7: definition.ErrorDefinition{ErrorNumber:7, ErrorType:"GlobalError", Symbol:"EE_LINK", SQLState:"", Message:"Error on rename of '%s' to '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    9: definition.ErrorDefinition{ErrorNumber:9, ErrorType:"GlobalError", Symbol:"EE_EOFERR", SQLState:"", Message:"Unexpected EOF found when reading file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    10: definition.ErrorDefinition{ErrorNumber:10, ErrorType:"GlobalError", Symbol:"EE_CANTLOCK", SQLState:"", Message:"Can't lock file (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    11: definition.ErrorDefinition{ErrorNumber:11, ErrorType:"GlobalError", Symbol:"EE_CANTUNLOCK", SQLState:"", Message:"Can't unlock file (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    12: definition.ErrorDefinition{ErrorNumber:12, ErrorType:"GlobalError", Symbol:"EE_DIR", SQLState:"", Message:"Can't read dir of '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    13: definition.ErrorDefinition{ErrorNumber:13, ErrorType:"GlobalError", Symbol:"EE_STAT", SQLState:"", Message:"Can't get stat of '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    14: definition.ErrorDefinition{ErrorNumber:14, ErrorType:"GlobalError", Symbol:"EE_CANT_CHSIZE", SQLState:"", Message:"Can't change size of file (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    15: definition.ErrorDefinition{ErrorNumber:15, ErrorType:"GlobalError", Symbol:"EE_CANT_OPEN_STREAM", SQLState:"", Message:"Can't open stream from handle (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    16: definition.ErrorDefinition{ErrorNumber:16, ErrorType:"GlobalError", Symbol:"EE_GETWD", SQLState:"", Message:"Can't get working directory (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    17: definition.ErrorDefinition{ErrorNumber:17, ErrorType:"GlobalError", Symbol:"EE_SETWD", SQLState:"", Message:"Can't change dir to '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    18: definition.ErrorDefinition{ErrorNumber:18, ErrorType:"GlobalError", Symbol:"EE_LINK_WARNING", SQLState:"", Message:"Warning: '%s' had %d links ", Description:"", MySQLVersion:"5.7"},
    19: definition.ErrorDefinition{ErrorNumber:19, ErrorType:"GlobalError", Symbol:"EE_OPEN_WARNING", SQLState:"", Message:"Warning: %d files and %d streams is left open ", Description:"", MySQLVersion:"5.7"},
    20: definition.ErrorDefinition{ErrorNumber:20, ErrorType:"GlobalError", Symbol:"EE_DISK_FULL", SQLState:"", Message:"Disk is full writing '%s' (Errcode: %d - %s). Waiting for someone to free space... ", Description:"", MySQLVersion:"5.7"},
    21: definition.ErrorDefinition{ErrorNumber:21, ErrorType:"GlobalError", Symbol:"EE_CANT_MKDIR", SQLState:"", Message:"Can't create directory '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    22: definition.ErrorDefinition{ErrorNumber:22, ErrorType:"GlobalError", Symbol:"EE_UNKNOWN_CHARSET", SQLState:"", Message:"Character set '%s' is not a compiled character set and is not specified in the '%s' file ", Description:"", MySQLVersion:"5.7"},
    23: definition.ErrorDefinition{ErrorNumber:23, ErrorType:"GlobalError", Symbol:"EE_OUT_OF_FILERESOURCES", SQLState:"", Message:"Out of resources when opening file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    24: definition.ErrorDefinition{ErrorNumber:24, ErrorType:"GlobalError", Symbol:"EE_CANT_READLINK", SQLState:"", Message:"Can't read value for symlink '%s' (Error %d - %s) ", Description:"", MySQLVersion:"5.7"},
    25: definition.ErrorDefinition{ErrorNumber:25, ErrorType:"GlobalError", Symbol:"EE_CANT_SYMLINK", SQLState:"", Message:"Can't create symlink '%s' pointing at '%s' (Error %d - %s) ", Description:"", MySQLVersion:"5.7"},
    26: definition.ErrorDefinition{ErrorNumber:26, ErrorType:"GlobalError", Symbol:"EE_REALPATH", SQLState:"", Message:"Error on realpath() on '%s' (Error %d - %s) ", Description:"", MySQLVersion:"5.7"},
    27: definition.ErrorDefinition{ErrorNumber:27, ErrorType:"GlobalError", Symbol:"EE_SYNC", SQLState:"", Message:"Can't sync file '%s' to disk (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    28: definition.ErrorDefinition{ErrorNumber:28, ErrorType:"GlobalError", Symbol:"EE_UNKNOWN_COLLATION", SQLState:"", Message:"Collation '%s' is not a compiled collation and is not specified in the '%s' file ", Description:"", MySQLVersion:"5.7"},
    29: definition.ErrorDefinition{ErrorNumber:29, ErrorType:"GlobalError", Symbol:"EE_FILENOTFOUND", SQLState:"", Message:"File '%s' not found (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    30: definition.ErrorDefinition{ErrorNumber:30, ErrorType:"GlobalError", Symbol:"EE_FILE_NOT_CLOSED", SQLState:"", Message:"File '%s' (fileno: %d) was not closed ", Description:"", MySQLVersion:"5.7"},
    31: definition.ErrorDefinition{ErrorNumber:31, ErrorType:"GlobalError", Symbol:"EE_CHANGE_OWNERSHIP", SQLState:"", Message:"Can't change ownership of the file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    32: definition.ErrorDefinition{ErrorNumber:32, ErrorType:"GlobalError", Symbol:"EE_CHANGE_PERMISSIONS", SQLState:"", Message:"Can't change permissions of the file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    33: definition.ErrorDefinition{ErrorNumber:33, ErrorType:"GlobalError", Symbol:"EE_CANT_SEEK", SQLState:"", Message:"Can't seek in file '%s' (Errcode: %d - %s) ", Description:"", MySQLVersion:"5.7"},
    34: definition.ErrorDefinition{ErrorNumber:34, ErrorType:"GlobalError", Symbol:"EE_CAPACITY_EXCEEDED", SQLState:"", Message:"Memory capacity exceeded (capacity %llu bytes)", Description:"EE_CAPACITY_EXCEEDED was added in 5.7.9.", MySQLVersion:"5.7"},
}

const (
ER_HASHCHK int = 1000
    ER_NISAMCHK int = 1001
    ER_NO int = 1002
    ER_YES int = 1003
    ER_CANT_CREATE_FILE int = 1004
    ER_CANT_CREATE_TABLE int = 1005
    ER_CANT_CREATE_DB int = 1006
    ER_DB_CREATE_EXISTS int = 1007
    ER_DB_DROP_EXISTS int = 1008
    ER_DB_DROP_DELETE int = 1009
    ER_DB_DROP_RMDIR int = 1010
    ER_CANT_DELETE_FILE int = 1011
    ER_CANT_FIND_SYSTEM_REC int = 1012
    ER_CANT_GET_STAT int = 1013
    ER_CANT_GET_WD int = 1014
    ER_CANT_LOCK int = 1015
    ER_CANT_OPEN_FILE int = 1016
    ER_FILE_NOT_FOUND int = 1017
    ER_CANT_READ_DIR int = 1018
    ER_CANT_SET_WD int = 1019
    ER_CHECKREAD int = 1020
    ER_DISK_FULL int = 1021
    ER_DUP_KEY int = 1022
    ER_ERROR_ON_CLOSE int = 1023
    ER_ERROR_ON_READ int = 1024
    ER_ERROR_ON_RENAME int = 1025
    ER_ERROR_ON_WRITE int = 1026
    ER_FILE_USED int = 1027
    ER_FILSORT_ABORT int = 1028
    ER_FORM_NOT_FOUND int = 1029
    ER_GET_ERRNO int = 1030
    ER_ILLEGAL_HA int = 1031
    ER_KEY_NOT_FOUND int = 1032
    ER_NOT_FORM_FILE int = 1033
    ER_NOT_KEYFILE int = 1034
    ER_OLD_KEYFILE int = 1035
    ER_OPEN_AS_READONLY int = 1036
    ER_OUTOFMEMORY int = 1037
    ER_OUT_OF_SORTMEMORY int = 1038
    ER_UNEXPECTED_EOF int = 1039
    ER_CON_COUNT_ERROR int = 1040
    ER_OUT_OF_RESOURCES int = 1041
    ER_BAD_HOST_ERROR int = 1042
    ER_HANDSHAKE_ERROR int = 1043
    ER_DBACCESS_DENIED_ERROR int = 1044
    ER_ACCESS_DENIED_ERROR int = 1045
    ER_NO_DB_ERROR int = 1046
    ER_UNKNOWN_COM_ERROR int = 1047
    ER_BAD_NULL_ERROR int = 1048
    ER_BAD_DB_ERROR int = 1049
    ER_TABLE_EXISTS_ERROR int = 1050
    ER_BAD_TABLE_ERROR int = 1051
    ER_NON_UNIQ_ERROR int = 1052
    ER_SERVER_SHUTDOWN int = 1053
    ER_BAD_FIELD_ERROR int = 1054
    ER_WRONG_FIELD_WITH_GROUP int = 1055
    ER_WRONG_GROUP_FIELD int = 1056
    ER_WRONG_SUM_SELECT int = 1057
    ER_WRONG_VALUE_COUNT int = 1058
    ER_TOO_LONG_IDENT int = 1059
    ER_DUP_FIELDNAME int = 1060
    ER_DUP_KEYNAME int = 1061
    ER_DUP_ENTRY int = 1062
    ER_WRONG_FIELD_SPEC int = 1063
    ER_PARSE_ERROR int = 1064
    ER_EMPTY_QUERY int = 1065
    ER_NONUNIQ_TABLE int = 1066
    ER_INVALID_DEFAULT int = 1067
    ER_MULTIPLE_PRI_KEY int = 1068
    ER_TOO_MANY_KEYS int = 1069
    ER_TOO_MANY_KEY_PARTS int = 1070
    ER_TOO_LONG_KEY int = 1071
    ER_KEY_COLUMN_DOES_NOT_EXITS int = 1072
    ER_BLOB_USED_AS_KEY int = 1073
    ER_TOO_BIG_FIELDLENGTH int = 1074
    ER_WRONG_AUTO_KEY int = 1075
    ER_READY int = 1076
    ER_NORMAL_SHUTDOWN int = 1077
    ER_GOT_SIGNAL int = 1078
    ER_SHUTDOWN_COMPLETE int = 1079
    ER_FORCING_CLOSE int = 1080
    ER_IPSOCK_ERROR int = 1081
    ER_NO_SUCH_INDEX int = 1082
    ER_WRONG_FIELD_TERMINATORS int = 1083
    ER_BLOBS_AND_NO_TERMINATED int = 1084
    ER_TEXTFILE_NOT_READABLE int = 1085
    ER_FILE_EXISTS_ERROR int = 1086
    ER_LOAD_INFO int = 1087
    ER_ALTER_INFO int = 1088
    ER_WRONG_SUB_KEY int = 1089
    ER_CANT_REMOVE_ALL_FIELDS int = 1090
    ER_CANT_DROP_FIELD_OR_KEY int = 1091
    ER_INSERT_INFO int = 1092
    ER_UPDATE_TABLE_USED int = 1093
    ER_NO_SUCH_THREAD int = 1094
    ER_KILL_DENIED_ERROR int = 1095
    ER_NO_TABLES_USED int = 1096
    ER_TOO_BIG_SET int = 1097
    ER_NO_UNIQUE_LOGFILE int = 1098
    ER_TABLE_NOT_LOCKED_FOR_WRITE int = 1099
    ER_TABLE_NOT_LOCKED int = 1100
    ER_BLOB_CANT_HAVE_DEFAULT int = 1101
    ER_WRONG_DB_NAME int = 1102
    ER_WRONG_TABLE_NAME int = 1103
    ER_TOO_BIG_SELECT int = 1104
    ER_UNKNOWN_ERROR int = 1105
    ER_UNKNOWN_PROCEDURE int = 1106
    ER_WRONG_PARAMCOUNT_TO_PROCEDURE int = 1107
    ER_WRONG_PARAMETERS_TO_PROCEDURE int = 1108
    ER_UNKNOWN_TABLE int = 1109
    ER_FIELD_SPECIFIED_TWICE int = 1110
    ER_INVALID_GROUP_FUNC_USE int = 1111
    ER_UNSUPPORTED_EXTENSION int = 1112
    ER_TABLE_MUST_HAVE_COLUMNS int = 1113
    ER_RECORD_FILE_FULL int = 1114
    ER_UNKNOWN_CHARACTER_SET int = 1115
    ER_TOO_MANY_TABLES int = 1116
    ER_TOO_MANY_FIELDS int = 1117
    ER_TOO_BIG_ROWSIZE int = 1118
    ER_STACK_OVERRUN int = 1119
    ER_WRONG_OUTER_JOIN int = 1120
    ER_NULL_COLUMN_IN_INDEX int = 1121
    ER_CANT_FIND_UDF int = 1122
    ER_CANT_INITIALIZE_UDF int = 1123
    ER_UDF_NO_PATHS int = 1124
    ER_UDF_EXISTS int = 1125
    ER_CANT_OPEN_LIBRARY int = 1126
    ER_CANT_FIND_DL_ENTRY int = 1127
    ER_FUNCTION_NOT_DEFINED int = 1128
    ER_HOST_IS_BLOCKED int = 1129
    ER_HOST_NOT_PRIVILEGED int = 1130
    ER_PASSWORD_ANONYMOUS_USER int = 1131
    ER_PASSWORD_NOT_ALLOWED int = 1132
    ER_PASSWORD_NO_MATCH int = 1133
    ER_UPDATE_INFO int = 1134
    ER_CANT_CREATE_THREAD int = 1135
    ER_WRONG_VALUE_COUNT_ON_ROW int = 1136
    ER_CANT_REOPEN_TABLE int = 1137
    ER_INVALID_USE_OF_NULL int = 1138
    ER_REGEXP_ERROR int = 1139
    ER_MIX_OF_GROUP_FUNC_AND_FIELDS int = 1140
    ER_NONEXISTING_GRANT int = 1141
    ER_TABLEACCESS_DENIED_ERROR int = 1142
    ER_COLUMNACCESS_DENIED_ERROR int = 1143
    ER_ILLEGAL_GRANT_FOR_TABLE int = 1144
    ER_GRANT_WRONG_HOST_OR_USER int = 1145
    ER_NO_SUCH_TABLE int = 1146
    ER_NONEXISTING_TABLE_GRANT int = 1147
    ER_NOT_ALLOWED_COMMAND int = 1148
    ER_SYNTAX_ERROR int = 1149
    ER_UNUSED1 int = 1150
    ER_UNUSED2 int = 1151
    ER_ABORTING_CONNECTION int = 1152
    ER_NET_PACKET_TOO_LARGE int = 1153
    ER_NET_READ_ERROR_FROM_PIPE int = 1154
    ER_NET_FCNTL_ERROR int = 1155
    ER_NET_PACKETS_OUT_OF_ORDER int = 1156
    ER_NET_UNCOMPRESS_ERROR int = 1157
    ER_NET_READ_ERROR int = 1158
    ER_NET_READ_INTERRUPTED int = 1159
    ER_NET_ERROR_ON_WRITE int = 1160
    ER_NET_WRITE_INTERRUPTED int = 1161
    ER_TOO_LONG_STRING int = 1162
    ER_TABLE_CANT_HANDLE_BLOB int = 1163
    ER_TABLE_CANT_HANDLE_AUTO_INCREMENT int = 1164
    ER_UNUSED3 int = 1165
    ER_WRONG_COLUMN_NAME int = 1166
    ER_WRONG_KEY_COLUMN int = 1167
    ER_WRONG_MRG_TABLE int = 1168
    ER_DUP_UNIQUE int = 1169
    ER_BLOB_KEY_WITHOUT_LENGTH int = 1170
    ER_PRIMARY_CANT_HAVE_NULL int = 1171
    ER_TOO_MANY_ROWS int = 1172
    ER_REQUIRES_PRIMARY_KEY int = 1173
    ER_NO_RAID_COMPILED int = 1174
    ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE int = 1175
    ER_KEY_DOES_NOT_EXITS int = 1176
    ER_CHECK_NO_SUCH_TABLE int = 1177
    ER_CHECK_NOT_IMPLEMENTED int = 1178
    ER_CANT_DO_THIS_DURING_AN_TRANSACTION int = 1179
    ER_ERROR_DURING_COMMIT int = 1180
    ER_ERROR_DURING_ROLLBACK int = 1181
    ER_ERROR_DURING_FLUSH_LOGS int = 1182
    ER_ERROR_DURING_CHECKPOINT int = 1183
    ER_NEW_ABORTING_CONNECTION int = 1184
    ER_DUMP_NOT_IMPLEMENTED int = 1185
    ER_FLUSH_MASTER_BINLOG_CLOSED int = 1186
    ER_INDEX_REBUILD int = 1187
    ER_MASTER int = 1188
    ER_MASTER_NET_READ int = 1189
    ER_MASTER_NET_WRITE int = 1190
    ER_FT_MATCHING_KEY_NOT_FOUND int = 1191
    ER_LOCK_OR_ACTIVE_TRANSACTION int = 1192
    ER_UNKNOWN_SYSTEM_VARIABLE int = 1193
    ER_CRASHED_ON_USAGE int = 1194
    ER_CRASHED_ON_REPAIR int = 1195
    ER_WARNING_NOT_COMPLETE_ROLLBACK int = 1196
    ER_TRANS_CACHE_FULL int = 1197
    ER_SLAVE_MUST_STOP int = 1198
    ER_SLAVE_NOT_RUNNING int = 1199
    ER_BAD_SLAVE int = 1200
    ER_MASTER_INFO int = 1201
    ER_SLAVE_THREAD int = 1202
    ER_TOO_MANY_USER_CONNECTIONS int = 1203
    ER_SET_CONSTANTS_ONLY int = 1204
    ER_LOCK_WAIT_TIMEOUT int = 1205
    ER_LOCK_TABLE_FULL int = 1206
    ER_READ_ONLY_TRANSACTION int = 1207
    ER_DROP_DB_WITH_READ_LOCK int = 1208
    ER_CREATE_DB_WITH_READ_LOCK int = 1209
    ER_WRONG_ARGUMENTS int = 1210
    ER_NO_PERMISSION_TO_CREATE_USER int = 1211
    ER_UNION_TABLES_IN_DIFFERENT_DIR int = 1212
    ER_LOCK_DEADLOCK int = 1213
    ER_TABLE_CANT_HANDLE_FT int = 1214
    ER_CANNOT_ADD_FOREIGN int = 1215
    ER_NO_REFERENCED_ROW int = 1216
    ER_ROW_IS_REFERENCED int = 1217
    ER_CONNECT_TO_MASTER int = 1218
    ER_QUERY_ON_MASTER int = 1219
    ER_ERROR_WHEN_EXECUTING_COMMAND int = 1220
    ER_WRONG_USAGE int = 1221
    ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT int = 1222
    ER_CANT_UPDATE_WITH_READLOCK int = 1223
    ER_MIXING_NOT_ALLOWED int = 1224
    ER_DUP_ARGUMENT int = 1225
    ER_USER_LIMIT_REACHED int = 1226
    ER_SPECIFIC_ACCESS_DENIED_ERROR int = 1227
    ER_LOCAL_VARIABLE int = 1228
    ER_GLOBAL_VARIABLE int = 1229
    ER_NO_DEFAULT int = 1230
    ER_WRONG_VALUE_FOR_VAR int = 1231
    ER_WRONG_TYPE_FOR_VAR int = 1232
    ER_VAR_CANT_BE_READ int = 1233
    ER_CANT_USE_OPTION_HERE int = 1234
    ER_NOT_SUPPORTED_YET int = 1235
    ER_MASTER_FATAL_ERROR_READING_BINLOG int = 1236
    ER_SLAVE_IGNORED_TABLE int = 1237
    ER_INCORRECT_GLOBAL_LOCAL_VAR int = 1238
    ER_WRONG_FK_DEF int = 1239
    ER_KEY_REF_DO_NOT_MATCH_TABLE_REF int = 1240
    ER_OPERAND_COLUMNS int = 1241
    ER_SUBQUERY_NO_1_ROW int = 1242
    ER_UNKNOWN_STMT_HANDLER int = 1243
    ER_CORRUPT_HELP_DB int = 1244
    ER_CYCLIC_REFERENCE int = 1245
    ER_AUTO_CONVERT int = 1246
    ER_ILLEGAL_REFERENCE int = 1247
    ER_DERIVED_MUST_HAVE_ALIAS int = 1248
    ER_SELECT_REDUCED int = 1249
    ER_TABLENAME_NOT_ALLOWED_HERE int = 1250
    ER_NOT_SUPPORTED_AUTH_MODE int = 1251
    ER_SPATIAL_CANT_HAVE_NULL int = 1252
    ER_COLLATION_CHARSET_MISMATCH int = 1253
    ER_SLAVE_WAS_RUNNING int = 1254
    ER_SLAVE_WAS_NOT_RUNNING int = 1255
    ER_TOO_BIG_FOR_UNCOMPRESS int = 1256
    ER_ZLIB_Z_MEM_ERROR int = 1257
    ER_ZLIB_Z_BUF_ERROR int = 1258
    ER_ZLIB_Z_DATA_ERROR int = 1259
    ER_CUT_VALUE_GROUP_CONCAT int = 1260
    ER_WARN_TOO_FEW_RECORDS int = 1261
    ER_WARN_TOO_MANY_RECORDS int = 1262
    ER_WARN_NULL_TO_NOTNULL int = 1263
    ER_WARN_DATA_OUT_OF_RANGE int = 1264
    WARN_DATA_TRUNCATED int = 1265
    ER_WARN_USING_OTHER_HANDLER int = 1266
    ER_CANT_AGGREGATE_2COLLATIONS int = 1267
    ER_DROP_USER int = 1268
    ER_REVOKE_GRANTS int = 1269
    ER_CANT_AGGREGATE_3COLLATIONS int = 1270
    ER_CANT_AGGREGATE_NCOLLATIONS int = 1271
    ER_VARIABLE_IS_NOT_STRUCT int = 1272
    ER_UNKNOWN_COLLATION int = 1273
    ER_SLAVE_IGNORED_SSL_PARAMS int = 1274
    ER_SERVER_IS_IN_SECURE_AUTH_MODE int = 1275
    ER_WARN_FIELD_RESOLVED int = 1276
    ER_BAD_SLAVE_UNTIL_COND int = 1277
    ER_MISSING_SKIP_SLAVE int = 1278
    ER_UNTIL_COND_IGNORED int = 1279
    ER_WRONG_NAME_FOR_INDEX int = 1280
    ER_WRONG_NAME_FOR_CATALOG int = 1281
    ER_WARN_QC_RESIZE int = 1282
    ER_BAD_FT_COLUMN int = 1283
    ER_UNKNOWN_KEY_CACHE int = 1284
    ER_WARN_HOSTNAME_WONT_WORK int = 1285
    ER_UNKNOWN_STORAGE_ENGINE int = 1286
    ER_WARN_DEPRECATED_SYNTAX int = 1287
    ER_NON_UPDATABLE_TABLE int = 1288
    ER_FEATURE_DISABLED int = 1289
    ER_OPTION_PREVENTS_STATEMENT int = 1290
    ER_DUPLICATED_VALUE_IN_TYPE int = 1291
    ER_TRUNCATED_WRONG_VALUE int = 1292
    ER_TOO_MUCH_AUTO_TIMESTAMP_COLS int = 1293
    ER_INVALID_ON_UPDATE int = 1294
    ER_UNSUPPORTED_PS int = 1295
    ER_GET_ERRMSG int = 1296
    ER_GET_TEMPORARY_ERRMSG int = 1297
    ER_UNKNOWN_TIME_ZONE int = 1298
    ER_WARN_INVALID_TIMESTAMP int = 1299
    ER_INVALID_CHARACTER_STRING int = 1300
    ER_WARN_ALLOWED_PACKET_OVERFLOWED int = 1301
    ER_CONFLICTING_DECLARATIONS int = 1302
    ER_SP_NO_RECURSIVE_CREATE int = 1303
    ER_SP_ALREADY_EXISTS int = 1304
    ER_SP_DOES_NOT_EXIST int = 1305
    ER_SP_DROP_FAILED int = 1306
    ER_SP_STORE_FAILED int = 1307
    ER_SP_LILABEL_MISMATCH int = 1308
    ER_SP_LABEL_REDEFINE int = 1309
    ER_SP_LABEL_MISMATCH int = 1310
    ER_SP_UNINIT_VAR int = 1311
    ER_SP_BADSELECT int = 1312
    ER_SP_BADRETURN int = 1313
    ER_SP_BADSTATEMENT int = 1314
    ER_UPDATE_LOG_DEPRECATED_IGNORED int = 1315
    ER_UPDATE_LOG_DEPRECATED_TRANSLATED int = 1316
    ER_QUERY_INTERRUPTED int = 1317
    ER_SP_WRONG_NO_OF_ARGS int = 1318
    ER_SP_COND_MISMATCH int = 1319
    ER_SP_NORETURN int = 1320
    ER_SP_NORETURNEND int = 1321
    ER_SP_BAD_CURSOR_QUERY int = 1322
    ER_SP_BAD_CURSOR_SELECT int = 1323
    ER_SP_CURSOR_MISMATCH int = 1324
    ER_SP_CURSOR_ALREADY_OPEN int = 1325
    ER_SP_CURSOR_NOT_OPEN int = 1326
    ER_SP_UNDECLARED_VAR int = 1327
    ER_SP_WRONG_NO_OF_FETCH_ARGS int = 1328
    ER_SP_FETCH_NO_DATA int = 1329
    ER_SP_DUP_PARAM int = 1330
    ER_SP_DUP_VAR int = 1331
    ER_SP_DUP_COND int = 1332
    ER_SP_DUP_CURS int = 1333
    ER_SP_CANT_ALTER int = 1334
    ER_SP_SUBSELECT_NYI int = 1335
    ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG int = 1336
    ER_SP_VARCOND_AFTER_CURSHNDLR int = 1337
    ER_SP_CURSOR_AFTER_HANDLER int = 1338
    ER_SP_CASE_NOT_FOUND int = 1339
    ER_FPARSER_TOO_BIG_FILE int = 1340
    ER_FPARSER_BAD_HEADER int = 1341
    ER_FPARSER_EOF_IN_COMMENT int = 1342
    ER_FPARSER_ERROR_IN_PARAMETER int = 1343
    ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER int = 1344
    ER_VIEW_NO_EXPLAIN int = 1345
    ER_FRM_UNKNOWN_TYPE int = 1346
    ER_WRONG_OBJECT int = 1347
    ER_NONUPDATEABLE_COLUMN int = 1348
    ER_VIEW_SELECT_DERIVED int = 1349
    ER_VIEW_SELECT_DERIVED_UNUSED int = 1349
    ER_VIEW_SELECT_CLAUSE int = 1350
    ER_VIEW_SELECT_VARIABLE int = 1351
    ER_VIEW_SELECT_TMPTABLE int = 1352
    ER_VIEW_WRONG_LIST int = 1353
    ER_WARN_VIEW_MERGE int = 1354
    ER_WARN_VIEW_WITHOUT_KEY int = 1355
    ER_VIEW_INVALID int = 1356
    ER_SP_NO_DROP_SP int = 1357
    ER_SP_GOTO_IN_HNDLR int = 1358
    ER_TRG_ALREADY_EXISTS int = 1359
    ER_TRG_DOES_NOT_EXIST int = 1360
    ER_TRG_ON_VIEW_OR_TEMP_TABLE int = 1361
    ER_TRG_CANT_CHANGE_ROW int = 1362
    ER_TRG_NO_SUCH_ROW_IN_TRG int = 1363
    ER_NO_DEFAULT_FOR_FIELD int = 1364
    ER_DIVISION_BY_ZERO int = 1365
    ER_TRUNCATED_WRONG_VALUE_FOR_FIELD int = 1366
    ER_ILLEGAL_VALUE_FOR_TYPE int = 1367
    ER_VIEW_NONUPD_CHECK int = 1368
    ER_VIEW_CHECK_FAILED int = 1369
    ER_PROCACCESS_DENIED_ERROR int = 1370
    ER_RELAY_LOG_FAIL int = 1371
    ER_PASSWD_LENGTH int = 1372
    ER_UNKNOWN_TARGET_BINLOG int = 1373
    ER_IO_ERR_LOG_INDEX_READ int = 1374
    ER_BINLOG_PURGE_PROHIBITED int = 1375
    ER_FSEEK_FAIL int = 1376
    ER_BINLOG_PURGE_FATAL_ERR int = 1377
    ER_LOG_IN_USE int = 1378
    ER_LOG_PURGE_UNKNOWN_ERR int = 1379
    ER_RELAY_LOG_INIT int = 1380
    ER_NO_BINARY_LOGGING int = 1381
    ER_RESERVED_SYNTAX int = 1382
    ER_WSAS_FAILED int = 1383
    ER_DIFF_GROUPS_PROC int = 1384
    ER_NO_GROUP_FOR_PROC int = 1385
    ER_ORDER_WITH_PROC int = 1386
    ER_LOGGING_PROHIBIT_CHANGING_OF int = 1387
    ER_NO_FILE_MAPPING int = 1388
    ER_WRONG_MAGIC int = 1389
    ER_PS_MANY_PARAM int = 1390
    ER_KEY_PART_0 int = 1391
    ER_VIEW_CHECKSUM int = 1392
    ER_VIEW_MULTIUPDATE int = 1393
    ER_VIEW_NO_INSERT_FIELD_LIST int = 1394
    ER_VIEW_DELETE_MERGE_VIEW int = 1395
    ER_CANNOT_USER int = 1396
    ER_XAER_NOTA int = 1397
    ER_XAER_INVAL int = 1398
    ER_XAER_RMFAIL int = 1399
    ER_XAER_OUTSIDE int = 1400
    ER_XAER_RMERR int = 1401
    ER_XA_RBROLLBACK int = 1402
    ER_NONEXISTING_PROC_GRANT int = 1403
    ER_PROC_AUTO_GRANT_FAIL int = 1404
    ER_PROC_AUTO_REVOKE_FAIL int = 1405
    ER_DATA_TOO_LONG int = 1406
    ER_SP_BAD_SQLSTATE int = 1407
    ER_STARTUP int = 1408
    ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR int = 1409
    ER_CANT_CREATE_USER_WITH_GRANT int = 1410
    ER_WRONG_VALUE_FOR_TYPE int = 1411
    ER_TABLE_DEF_CHANGED int = 1412
    ER_SP_DUP_HANDLER int = 1413
    ER_SP_NOT_VAR_ARG int = 1414
    ER_SP_NO_RETSET int = 1415
    ER_CANT_CREATE_GEOMETRY_OBJECT int = 1416
    ER_FAILED_ROUTINE_BREAK_BINLOG int = 1417
    ER_BINLOG_UNSAFE_ROUTINE int = 1418
    ER_BINLOG_CREATE_ROUTINE_NEED_SUPER int = 1419
    ER_EXEC_STMT_WITH_OPEN_CURSOR int = 1420
    ER_STMT_HAS_NO_OPEN_CURSOR int = 1421
    ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG int = 1422
    ER_NO_DEFAULT_FOR_VIEW_FIELD int = 1423
    ER_SP_NO_RECURSION int = 1424
    ER_TOO_BIG_SCALE int = 1425
    ER_TOO_BIG_PRECISION int = 1426
    ER_M_BIGGER_THAN_D int = 1427
    ER_WRONG_LOCK_OF_SYSTEM_TABLE int = 1428
    ER_CONNECT_TO_FOREIGN_DATA_SOURCE int = 1429
    ER_QUERY_ON_FOREIGN_DATA_SOURCE int = 1430
    ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST int = 1431
    ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE int = 1432
    ER_FOREIGN_DATA_STRING_INVALID int = 1433
    ER_CANT_CREATE_FEDERATED_TABLE int = 1434
    ER_TRG_IN_WRONG_SCHEMA int = 1435
    ER_STACK_OVERRUN_NEED_MORE int = 1436
    ER_TOO_LONG_BODY int = 1437
    ER_WARN_CANT_DROP_DEFAULT_KEYCACHE int = 1438
    ER_TOO_BIG_DISPLAYWIDTH int = 1439
    ER_XAER_DUPID int = 1440
    ER_DATETIME_FUNCTION_OVERFLOW int = 1441
    ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG int = 1442
    ER_VIEW_PREVENT_UPDATE int = 1443
    ER_PS_NO_RECURSION int = 1444
    ER_SP_CANT_SET_AUTOCOMMIT int = 1445
    ER_MALFORMED_DEFINER int = 1446
    ER_VIEW_FRM_NO_USER int = 1447
    ER_VIEW_OTHER_USER int = 1448
    ER_NO_SUCH_USER int = 1449
    ER_FORBID_SCHEMA_CHANGE int = 1450
    ER_ROW_IS_REFERENCED_2 int = 1451
    ER_NO_REFERENCED_ROW_2 int = 1452
    ER_SP_BAD_VAR_SHADOW int = 1453
    ER_TRG_NO_DEFINER int = 1454
    ER_OLD_FILE_FORMAT int = 1455
    ER_SP_RECURSION_LIMIT int = 1456
    ER_SP_PROC_TABLE_CORRUPT int = 1457
    ER_SP_WRONG_NAME int = 1458
    ER_TABLE_NEEDS_UPGRADE int = 1459
    ER_SP_NO_AGGREGATE int = 1460
    ER_MAX_PREPARED_STMT_COUNT_REACHED int = 1461
    ER_VIEW_RECURSIVE int = 1462
    ER_NON_GROUPING_FIELD_USED int = 1463
    ER_TABLE_CANT_HANDLE_SPKEYS int = 1464
    ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA int = 1465
    ER_REMOVED_SPACES int = 1466
    ER_AUTOINC_READ_FAILED int = 1467
    ER_USERNAME int = 1468
    ER_HOSTNAME int = 1469
    ER_WRONG_STRING_LENGTH int = 1470
    ER_NON_INSERTABLE_TABLE int = 1471
    ER_ADMIN_WRONG_MRG_TABLE int = 1472
    ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT int = 1473
    ER_NAME_BECOMES_EMPTY int = 1474
    ER_AMBIGUOUS_FIELD_TERM int = 1475
    ER_FOREIGN_SERVER_EXISTS int = 1476
    ER_FOREIGN_SERVER_DOESNT_EXIST int = 1477
    ER_ILLEGAL_HA_CREATE_OPTION int = 1478
    ER_PARTITION_REQUIRES_VALUES_ERROR int = 1479
    ER_PARTITION_WRONG_VALUES_ERROR int = 1480
    ER_PARTITION_MAXVALUE_ERROR int = 1481
    ER_PARTITION_SUBPARTITION_ERROR int = 1482
    ER_PARTITION_SUBPART_MIX_ERROR int = 1483
    ER_PARTITION_WRONG_NO_PART_ERROR int = 1484
    ER_PARTITION_WRONG_NO_SUBPART_ERROR int = 1485
    ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR int = 1486
    ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR int = 1487
    ER_FIELD_NOT_FOUND_PART_ERROR int = 1488
    ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR int = 1489
    ER_INCONSISTENT_PARTITION_INFO_ERROR int = 1490
    ER_PARTITION_FUNC_NOT_ALLOWED_ERROR int = 1491
    ER_PARTITIONS_MUST_BE_DEFINED_ERROR int = 1492
    ER_RANGE_NOT_INCREASING_ERROR int = 1493
    ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR int = 1494
    ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR int = 1495
    ER_PARTITION_ENTRY_ERROR int = 1496
    ER_MIX_HANDLER_ERROR int = 1497
    ER_PARTITION_NOT_DEFINED_ERROR int = 1498
    ER_TOO_MANY_PARTITIONS_ERROR int = 1499
    ER_SUBPARTITION_ERROR int = 1500
    ER_CANT_CREATE_HANDLER_FILE int = 1501
    ER_BLOB_FIELD_IN_PART_FUNC_ERROR int = 1502
    ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF int = 1503
    ER_NO_PARTS_ERROR int = 1504
    ER_PARTITION_MGMT_ON_NONPARTITIONED int = 1505
    ER_FOREIGN_KEY_ON_PARTITIONED int = 1506
    ER_DROP_PARTITION_NON_EXISTENT int = 1507
    ER_DROP_LAST_PARTITION int = 1508
    ER_COALESCE_ONLY_ON_HASH_PARTITION int = 1509
    ER_REORG_HASH_ONLY_ON_SAME_NO int = 1510
    ER_REORG_NO_PARAM_ERROR int = 1511
    ER_ONLY_ON_RANGE_LIST_PARTITION int = 1512
    ER_ADD_PARTITION_SUBPART_ERROR int = 1513
    ER_ADD_PARTITION_NO_NEW_PARTITION int = 1514
    ER_COALESCE_PARTITION_NO_PARTITION int = 1515
    ER_REORG_PARTITION_NOT_EXIST int = 1516
    ER_SAME_NAME_PARTITION int = 1517
    ER_NO_BINLOG_ERROR int = 1518
    ER_CONSECUTIVE_REORG_PARTITIONS int = 1519
    ER_REORG_OUTSIDE_RANGE int = 1520
    ER_PARTITION_FUNCTION_FAILURE int = 1521
    ER_PART_STATE_ERROR int = 1522
    ER_LIMITED_PART_RANGE int = 1523
    ER_PLUGIN_IS_NOT_LOADED int = 1524
    ER_WRONG_VALUE int = 1525
    ER_NO_PARTITION_FOR_GIVEN_VALUE int = 1526
    ER_FILEGROUP_OPTION_ONLY_ONCE int = 1527
    ER_CREATE_FILEGROUP_FAILED int = 1528
    ER_DROP_FILEGROUP_FAILED int = 1529
    ER_TABLESPACE_AUTO_EXTEND_ERROR int = 1530
    ER_WRONG_SIZE_NUMBER int = 1531
    ER_SIZE_OVERFLOW_ERROR int = 1532
    ER_ALTER_FILEGROUP_FAILED int = 1533
    ER_BINLOG_ROW_LOGGING_FAILED int = 1534
    ER_BINLOG_ROW_WRONG_TABLE_DEF int = 1535
    ER_BINLOG_ROW_RBR_TO_SBR int = 1536
    ER_EVENT_ALREADY_EXISTS int = 1537
    ER_EVENT_STORE_FAILED int = 1538
    ER_EVENT_DOES_NOT_EXIST int = 1539
    ER_EVENT_CANT_ALTER int = 1540
    ER_EVENT_DROP_FAILED int = 1541
    ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG int = 1542
    ER_EVENT_ENDS_BEFORE_STARTS int = 1543
    ER_EVENT_EXEC_TIME_IN_THE_PAST int = 1544
    ER_EVENT_OPEN_TABLE_FAILED int = 1545
    ER_EVENT_NEITHER_M_EXPR_NOR_M_AT int = 1546
    ER_OBSOLETE_COL_COUNT_DOESNT_MATCH_CORRUPTED int = 1547
    ER_OBSOLETE_CANNOT_LOAD_FROM_TABLE int = 1548
    ER_EVENT_CANNOT_DELETE int = 1549
    ER_EVENT_COMPILE_ERROR int = 1550
    ER_EVENT_SAME_NAME int = 1551
    ER_EVENT_DATA_TOO_LONG int = 1552
    ER_DROP_INDEX_FK int = 1553
    ER_WARN_DEPRECATED_SYNTAX_WITH_VER int = 1554
    ER_CANT_WRITE_LOCK_LOG_TABLE int = 1555
    ER_CANT_LOCK_LOG_TABLE int = 1556
    ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED int = 1557
    ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE int = 1558
    ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR int = 1559
    ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT int = 1560
    ER_NDB_CANT_SWITCH_BINLOG_FORMAT int = 1561
    ER_PARTITION_NO_TEMPORARY int = 1562
    ER_PARTITION_CONST_DOMAIN_ERROR int = 1563
    ER_PARTITION_FUNCTION_IS_NOT_ALLOWED int = 1564
    ER_DDL_LOG_ERROR int = 1565
    ER_NULL_IN_VALUES_LESS_THAN int = 1566
    ER_WRONG_PARTITION_NAME int = 1567
    ER_CANT_CHANGE_TX_CHARACTERISTICS int = 1568
    ER_DUP_ENTRY_AUTOINCREMENT_CASE int = 1569
    ER_EVENT_MODIFY_QUEUE_ERROR int = 1570
    ER_EVENT_SET_VAR_ERROR int = 1571
    ER_PARTITION_MERGE_ERROR int = 1572
    ER_CANT_ACTIVATE_LOG int = 1573
    ER_RBR_NOT_AVAILABLE int = 1574
    ER_BASE64_DECODE_ERROR int = 1575
    ER_EVENT_RECURSION_FORBIDDEN int = 1576
    ER_EVENTS_DB_ERROR int = 1577
    ER_ONLY_INTEGERS_ALLOWED int = 1578
    ER_UNSUPORTED_LOG_ENGINE int = 1579
    ER_BAD_LOG_STATEMENT int = 1580
    ER_CANT_RENAME_LOG_TABLE int = 1581
    ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT int = 1582
    ER_WRONG_PARAMETERS_TO_NATIVE_FCT int = 1583
    ER_WRONG_PARAMETERS_TO_STORED_FCT int = 1584
    ER_NATIVE_FCT_NAME_COLLISION int = 1585
    ER_DUP_ENTRY_WITH_KEY_NAME int = 1586
    ER_BINLOG_PURGE_EMFILE int = 1587
    ER_EVENT_CANNOT_CREATE_IN_THE_PAST int = 1588
    ER_EVENT_CANNOT_ALTER_IN_THE_PAST int = 1589
    ER_SLAVE_INCIDENT int = 1590
    ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT int = 1591
    ER_BINLOG_UNSAFE_STATEMENT int = 1592
    ER_SLAVE_FATAL_ERROR int = 1593
    ER_SLAVE_RELAY_LOG_READ_FAILURE int = 1594
    ER_SLAVE_RELAY_LOG_WRITE_FAILURE int = 1595
    ER_SLAVE_CREATE_EVENT_FAILURE int = 1596
    ER_SLAVE_MASTER_COM_FAILURE int = 1597
    ER_BINLOG_LOGGING_IMPOSSIBLE int = 1598
    ER_VIEW_NO_CREATION_CTX int = 1599
    ER_VIEW_INVALID_CREATION_CTX int = 1600
    ER_SR_INVALID_CREATION_CTX int = 1601
    ER_TRG_CORRUPTED_FILE int = 1602
    ER_TRG_NO_CREATION_CTX int = 1603
    ER_TRG_INVALID_CREATION_CTX int = 1604
    ER_EVENT_INVALID_CREATION_CTX int = 1605
    ER_TRG_CANT_OPEN_TABLE int = 1606
    ER_CANT_CREATE_SROUTINE int = 1607
    ER_NEVER_USED int = 1608
    ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT int = 1609
    ER_SLAVE_CORRUPT_EVENT int = 1610
    ER_LOAD_DATA_INVALID_COLUMN int = 1611
    ER_LOAD_DATA_INVALID_COLUMN_UNUSED int = 1611
    ER_LOG_PURGE_NO_FILE int = 1612
    ER_XA_RBTIMEOUT int = 1613
    ER_XA_RBDEADLOCK int = 1614
    ER_NEED_REPREPARE int = 1615
    ER_DELAYED_NOT_SUPPORTED int = 1616
    WARN_NO_MASTER_INFO int = 1617
    WARN_OPTION_IGNORED int = 1618
    WARN_PLUGIN_DELETE_BUILTIN int = 1619
    ER_PLUGIN_DELETE_BUILTIN int = 1619
    WARN_PLUGIN_BUSY int = 1620
    ER_VARIABLE_IS_READONLY int = 1621
    ER_WARN_ENGINE_TRANSACTION_ROLLBACK int = 1622
    ER_SLAVE_HEARTBEAT_FAILURE int = 1623
    ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE int = 1624
    ER_NDB_REPLICATION_SCHEMA_ERROR int = 1625
    ER_CONFLICT_FN_PARSE_ERROR int = 1626
    ER_EXCEPTIONS_WRITE_ERROR int = 1627
    ER_TOO_LONG_TABLE_COMMENT int = 1628
    ER_TOO_LONG_FIELD_COMMENT int = 1629
    ER_FUNC_INEXISTENT_NAME_COLLISION int = 1630
    ER_DATABASE_NAME int = 1631
    ER_TABLE_NAME int = 1632
    ER_PARTITION_NAME int = 1633
    ER_SUBPARTITION_NAME int = 1634
    ER_TEMPORARY_NAME int = 1635
    ER_RENAMED_NAME int = 1636
    ER_TOO_MANY_CONCURRENT_TRXS int = 1637
    WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED int = 1638
    ER_DEBUG_SYNC_TIMEOUT int = 1639
    ER_DEBUG_SYNC_HIT_LIMIT int = 1640
    ER_DUP_SIGNAL_SET int = 1641
    ER_SIGNAL_WARN int = 1642
    ER_SIGNAL_NOT_FOUND int = 1643
    ER_SIGNAL_EXCEPTION int = 1644
    ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER int = 1645
    ER_SIGNAL_BAD_CONDITION_TYPE int = 1646
    WARN_COND_ITEM_TRUNCATED int = 1647
    ER_COND_ITEM_TOO_LONG int = 1648
    ER_UNKNOWN_LOCALE int = 1649
    ER_SLAVE_IGNORE_SERVER_IDS int = 1650
    ER_QUERY_CACHE_DISABLED int = 1651
    ER_SAME_NAME_PARTITION_FIELD int = 1652
    ER_PARTITION_COLUMN_LIST_ERROR int = 1653
    ER_WRONG_TYPE_COLUMN_VALUE_ERROR int = 1654
    ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR int = 1655
    ER_MAXVALUE_IN_VALUES_IN int = 1656
    ER_TOO_MANY_VALUES_ERROR int = 1657
    ER_ROW_SINGLE_PARTITION_FIELD_ERROR int = 1658
    ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD int = 1659
    ER_PARTITION_FIELDS_TOO_LONG int = 1660
    ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE int = 1661
    ER_BINLOG_ROW_MODE_AND_STMT_ENGINE int = 1662
    ER_BINLOG_UNSAFE_AND_STMT_ENGINE int = 1663
    ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE int = 1664
    ER_BINLOG_STMT_MODE_AND_ROW_ENGINE int = 1665
    ER_BINLOG_ROW_INJECTION_AND_STMT_MODE int = 1666
    ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE int = 1667
    ER_BINLOG_UNSAFE_LIMIT int = 1668
    ER_UNUSED4 int = 1669
    ER_BINLOG_UNSAFE_SYSTEM_TABLE int = 1670
    ER_BINLOG_UNSAFE_AUTOINC_COLUMNS int = 1671
    ER_BINLOG_UNSAFE_UDF int = 1672
    ER_BINLOG_UNSAFE_SYSTEM_VARIABLE int = 1673
    ER_BINLOG_UNSAFE_SYSTEM_FUNCTION int = 1674
    ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS int = 1675
    ER_MESSAGE_AND_STATEMENT int = 1676
    ER_SLAVE_CONVERSION_FAILED int = 1677
    ER_SLAVE_CANT_CREATE_CONVERSION int = 1678
    ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT int = 1679
    ER_PATH_LENGTH int = 1680
    ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT int = 1681
    ER_WRONG_NATIVE_TABLE_STRUCTURE int = 1682
    ER_WRONG_PERFSCHEMA_USAGE int = 1683
    ER_WARN_I_S_SKIPPED_TABLE int = 1684
    ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT int = 1685
    ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT int = 1686
    ER_SPATIAL_MUST_HAVE_GEOM_COL int = 1687
    ER_TOO_LONG_INDEX_COMMENT int = 1688
    ER_LOCK_ABORTED int = 1689
    ER_DATA_OUT_OF_RANGE int = 1690
    ER_WRONG_SPVAR_TYPE_IN_LIMIT int = 1691
    ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE int = 1692
    ER_BINLOG_UNSAFE_MIXED_STATEMENT int = 1693
    ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN int = 1694
    ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN int = 1695
    ER_FAILED_READ_FROM_PAR_FILE int = 1696
    ER_VALUES_IS_NOT_INT_TYPE_ERROR int = 1697
    ER_ACCESS_DENIED_NO_PASSWORD_ERROR int = 1698
    ER_SET_PASSWORD_AUTH_PLUGIN int = 1699
    ER_GRANT_PLUGIN_USER_EXISTS int = 1700
    ER_TRUNCATE_ILLEGAL_FK int = 1701
    ER_PLUGIN_IS_PERMANENT int = 1702
    ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN int = 1703
    ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX int = 1704
    ER_STMT_CACHE_FULL int = 1705
    ER_MULTI_UPDATE_KEY_CONFLICT int = 1706
    ER_TABLE_NEEDS_REBUILD int = 1707
    WARN_OPTION_BELOW_LIMIT int = 1708
    ER_INDEX_COLUMN_TOO_LONG int = 1709
    ER_ERROR_IN_TRIGGER_BODY int = 1710
    ER_ERROR_IN_UNKNOWN_TRIGGER_BODY int = 1711
    ER_INDEX_CORRUPT int = 1712
    ER_UNDO_RECORD_TOO_BIG int = 1713
    ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT int = 1714
    ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE int = 1715
    ER_BINLOG_UNSAFE_REPLACE_SELECT int = 1716
    ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT int = 1717
    ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT int = 1718
    ER_BINLOG_UNSAFE_UPDATE_IGNORE int = 1719
    ER_PLUGIN_NO_UNINSTALL int = 1720
    ER_PLUGIN_NO_INSTALL int = 1721
    ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT int = 1722
    ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC int = 1723
    ER_BINLOG_UNSAFE_INSERT_TWO_KEYS int = 1724
    ER_TABLE_IN_FK_CHECK int = 1725
    ER_UNSUPPORTED_ENGINE int = 1726
    ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST int = 1727
    ER_CANNOT_LOAD_FROM_TABLE_V2 int = 1728
    ER_MASTER_DELAY_VALUE_OUT_OF_RANGE int = 1729
    ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT int = 1730
    ER_PARTITION_EXCHANGE_DIFFERENT_OPTION int = 1731
    ER_PARTITION_EXCHANGE_PART_TABLE int = 1732
    ER_PARTITION_EXCHANGE_TEMP_TABLE int = 1733
    ER_PARTITION_INSTEAD_OF_SUBPARTITION int = 1734
    ER_UNKNOWN_PARTITION int = 1735
    ER_TABLES_DIFFERENT_METADATA int = 1736
    ER_ROW_DOES_NOT_MATCH_PARTITION int = 1737
    ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX int = 1738
    ER_WARN_INDEX_NOT_APPLICABLE int = 1739
    ER_PARTITION_EXCHANGE_FOREIGN_KEY int = 1740
    ER_NO_SUCH_KEY_VALUE int = 1741
    ER_RPL_INFO_DATA_TOO_LONG int = 1742
    ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE int = 1743
    ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE int = 1744
    ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX int = 1745
    ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT int = 1746
    ER_PARTITION_CLAUSE_ON_NONPARTITIONED int = 1747
    ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET int = 1748
    ER_NO_SUCH_PARTITION__UNUSED int = 1749
    ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE int = 1750
    ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE int = 1751
    ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE int = 1752
    ER_MTS_FEATURE_IS_NOT_SUPPORTED int = 1753
    ER_MTS_UPDATED_DBS_GREATER_MAX int = 1754
    ER_MTS_CANT_PARALLEL int = 1755
    ER_MTS_INCONSISTENT_DATA int = 1756
    ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING int = 1757
    ER_DA_INVALID_CONDITION_NUMBER int = 1758
    ER_INSECURE_PLAIN_TEXT int = 1759
    ER_INSECURE_CHANGE_MASTER int = 1760
    ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO int = 1761
    ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO int = 1762
    ER_SQLTHREAD_WITH_SECURE_SLAVE int = 1763
    ER_TABLE_HAS_NO_FT int = 1764
    ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER int = 1765
    ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION int = 1766
    ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST int = 1767
    ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL int = 1768
    ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION int = 1768
    ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION int = 1769
    ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL int = 1770
    ER_SKIPPING_LOGGED_TRANSACTION int = 1771
    ER_MALFORMED_GTID_SET_SPECIFICATION int = 1772
    ER_MALFORMED_GTID_SET_ENCODING int = 1773
    ER_MALFORMED_GTID_SPECIFICATION int = 1774
    ER_GNO_EXHAUSTED int = 1775
    ER_BAD_SLAVE_AUTO_POSITION int = 1776
    ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON int = 1777
    ER_AUTO_POSITION_REQUIRES_GTID_MODE_NOT_OFF int = 1777
    ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET int = 1778
    ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON int = 1779
    ER_GTID_MODE_ON_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON int = 1779
    ER_GTID_MODE_REQUIRES_BINLOG int = 1780
    ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF int = 1781
    ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON int = 1782
    ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF int = 1783
    ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF int = 1784
    ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF__UNUSED int = 1784
    ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE int = 1785
    ER_GTID_UNSAFE_CREATE_SELECT int = 1786
    ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTION int = 1787
    ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME int = 1788
    ER_MASTER_HAS_PURGED_REQUIRED_GTIDS int = 1789
    ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID int = 1790
    ER_UNKNOWN_EXPLAIN_FORMAT int = 1791
    ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION int = 1792
    ER_TOO_LONG_TABLE_PARTITION_COMMENT int = 1793
    ER_SLAVE_CONFIGURATION int = 1794
    ER_INNODB_FT_LIMIT int = 1795
    ER_INNODB_NO_FT_TEMP_TABLE int = 1796
    ER_INNODB_FT_WRONG_DOCID_COLUMN int = 1797
    ER_INNODB_FT_WRONG_DOCID_INDEX int = 1798
    ER_INNODB_ONLINE_LOG_TOO_BIG int = 1799
    ER_UNKNOWN_ALTER_ALGORITHM int = 1800
    ER_UNKNOWN_ALTER_LOCK int = 1801
    ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS int = 1802
    ER_MTS_RECOVERY_FAILURE int = 1803
    ER_MTS_RESET_WORKERS int = 1804
    ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2 int = 1805
    ER_SLAVE_SILENT_RETRY_TRANSACTION int = 1806
    ER_DISCARD_FK_CHECKS_RUNNING int = 1807
    ER_TABLE_SCHEMA_MISMATCH int = 1808
    ER_TABLE_IN_SYSTEM_TABLESPACE int = 1809
    ER_IO_READ_ERROR int = 1810
    ER_IO_WRITE_ERROR int = 1811
    ER_TABLESPACE_MISSING int = 1812
    ER_TABLESPACE_EXISTS int = 1813
    ER_TABLESPACE_DISCARDED int = 1814
    ER_INTERNAL_ERROR int = 1815
    ER_INNODB_IMPORT_ERROR int = 1816
    ER_INNODB_INDEX_CORRUPT int = 1817
    ER_INVALID_YEAR_COLUMN_LENGTH int = 1818
    ER_NOT_VALID_PASSWORD int = 1819
    ER_MUST_CHANGE_PASSWORD int = 1820
    ER_FK_NO_INDEX_CHILD int = 1821
    ER_FK_NO_INDEX_PARENT int = 1822
    ER_FK_FAIL_ADD_SYSTEM int = 1823
    ER_FK_CANNOT_OPEN_PARENT int = 1824
    ER_FK_INCORRECT_OPTION int = 1825
    ER_FK_DUP_NAME int = 1826
    ER_PASSWORD_FORMAT int = 1827
    ER_FK_COLUMN_CANNOT_DROP int = 1828
    ER_FK_COLUMN_CANNOT_DROP_CHILD int = 1829
    ER_FK_COLUMN_NOT_NULL int = 1830
    ER_DUP_INDEX int = 1831
    ER_FK_COLUMN_CANNOT_CHANGE int = 1832
    ER_FK_COLUMN_CANNOT_CHANGE_CHILD int = 1833
    ER_FK_CANNOT_DELETE_PARENT int = 1834
    ER_UNUSED5 int = 1834
    ER_MALFORMED_PACKET int = 1835
    ER_READ_ONLY_MODE int = 1836
    ER_GTID_NEXT_TYPE_UNDEFINED_GROUP int = 1837
    ER_VARIABLE_NOT_SETTABLE_IN_SP int = 1838
    ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF int = 1839
    ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY int = 1840
    ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY int = 1841
    ER_GTID_PURGED_WAS_CHANGED int = 1842
    ER_GTID_EXECUTED_WAS_CHANGED int = 1843
    ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES int = 1844
    ER_ALTER_OPERATION_NOT_SUPPORTED int = 1845
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON int = 1846
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY int = 1847
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION int = 1848
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME int = 1849
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE int = 1850
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK int = 1851
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE int = 1852
    ER_UNUSED6 int = 1852
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK int = 1853
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC int = 1854
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS int = 1855
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS int = 1856
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS int = 1857
    ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE int = 1858
    ER_DUP_UNKNOWN_IN_INDEX int = 1859
    ER_IDENT_CAUSES_TOO_LONG_PATH int = 1860
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL int = 1861
    ER_MUST_CHANGE_PASSWORD_LOGIN int = 1862
    ER_ROW_IN_WRONG_PARTITION int = 1863
    ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX int = 1864
    ER_INNODB_NO_FT_USES_PARSER int = 1865
    ER_BINLOG_LOGICAL_CORRUPTION int = 1866
    ER_WARN_PURGE_LOG_IN_USE int = 1867
    ER_WARN_PURGE_LOG_IS_ACTIVE int = 1868
    ER_AUTO_INCREMENT_CONFLICT int = 1869
    WARN_ON_BLOCKHOLE_IN_RBR int = 1870
    ER_SLAVE_MI_INIT_REPOSITORY int = 1871
    ER_SLAVE_RLI_INIT_REPOSITORY int = 1872
    ER_ACCESS_DENIED_CHANGE_USER_ERROR int = 1873
    ER_INNODB_READ_ONLY int = 1874
    ER_STOP_SLAVE_SQL_THREAD_TIMEOUT int = 1875
    ER_STOP_SLAVE_IO_THREAD_TIMEOUT int = 1876
    ER_TABLE_CORRUPT int = 1877
    ER_TEMP_FILE_WRITE_FAILURE int = 1878
    ER_INNODB_FT_AUX_NOT_HEX_ID int = 1879
    ER_OLD_TEMPORALS_UPGRADED int = 1880
    ER_INNODB_FORCED_RECOVERY int = 1881
    ER_AES_INVALID_IV int = 1882
    ER_PLUGIN_CANNOT_BE_UNINSTALLED int = 1883
    ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP int = 1884
    ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER int = 1885
    ER_MISSING_KEY int = 1886
    WARN_NAMED_PIPE_ACCESS_EVERYONE int = 1887
    ER_FOUND_MISSING_GTIDS int = 1888
    ER_SLAVE_IO_THREAD_MUST_STOP int = 1906
    ER_FILE_CORRUPT int = 3000
    ER_ERROR_ON_MASTER int = 3001
    ER_INCONSISTENT_ERROR int = 3002
    ER_STORAGE_ENGINE_NOT_LOADED int = 3003
    ER_GET_STACKED_DA_WITHOUT_ACTIVE_HANDLER int = 3004
    ER_WARN_LEGACY_SYNTAX_CONVERTED int = 3005
    ER_BINLOG_UNSAFE_FULLTEXT_PLUGIN int = 3006
    ER_CANNOT_DISCARD_TEMPORARY_TABLE int = 3007
    ER_FK_DEPTH_EXCEEDED int = 3008
    ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2 int = 3009
    ER_WARN_TRIGGER_DOESNT_HAVE_CREATED int = 3010
    ER_REFERENCED_TRG_DOES_NOT_EXIST int = 3011
    ER_EXPLAIN_NOT_SUPPORTED int = 3012
    ER_INVALID_FIELD_SIZE int = 3013
    ER_MISSING_HA_CREATE_OPTION int = 3014
    ER_ENGINE_OUT_OF_MEMORY int = 3015
    ER_PASSWORD_EXPIRE_ANONYMOUS_USER int = 3016
    ER_SLAVE_SQL_THREAD_MUST_STOP int = 3017
    ER_NO_FT_MATERIALIZED_SUBQUERY int = 3018
    ER_INNODB_UNDO_LOG_FULL int = 3019
    ER_INVALID_ARGUMENT_FOR_LOGARITHM int = 3020
    ER_SLAVE_CHANNEL_IO_THREAD_MUST_STOP int = 3021
    ER_WARN_OPEN_TEMP_TABLES_MUST_BE_ZERO int = 3022
    ER_WARN_ONLY_MASTER_LOG_FILE_NO_POS int = 3023
    ER_QUERY_TIMEOUT int = 3024
    ER_NON_RO_SELECT_DISABLE_TIMER int = 3025
    ER_DUP_LIST_ENTRY int = 3026
    ER_SQL_MODE_NO_EFFECT int = 3027
    ER_AGGREGATE_ORDER_FOR_UNION int = 3028
    ER_AGGREGATE_ORDER_NON_AGG_QUERY int = 3029
    ER_SLAVE_WORKER_STOPPED_PREVIOUS_THD_ERROR int = 3030
    ER_DONT_SUPPORT_SLAVE_PRESERVE_COMMIT_ORDER int = 3031
    ER_SERVER_OFFLINE_MODE int = 3032
    ER_GIS_DIFFERENT_SRIDS int = 3033
    ER_GIS_UNSUPPORTED_ARGUMENT int = 3034
    ER_GIS_UNKNOWN_ERROR int = 3035
    ER_GIS_UNKNOWN_EXCEPTION int = 3036
    ER_GIS_INVALID_DATA int = 3037
    ER_BOOST_GEOMETRY_EMPTY_INPUT_EXCEPTION int = 3038
    ER_BOOST_GEOMETRY_CENTROID_EXCEPTION int = 3039
    ER_BOOST_GEOMETRY_OVERLAY_INVALID_INPUT_EXCEPTION int = 3040
    ER_BOOST_GEOMETRY_TURN_INFO_EXCEPTION int = 3041
    ER_BOOST_GEOMETRY_SELF_INTERSECTION_POINT_EXCEPTION int = 3042
    ER_BOOST_GEOMETRY_UNKNOWN_EXCEPTION int = 3043
    ER_STD_BAD_ALLOC_ERROR int = 3044
    ER_STD_DOMAIN_ERROR int = 3045
    ER_STD_LENGTH_ERROR int = 3046
    ER_STD_INVALID_ARGUMENT int = 3047
    ER_STD_OUT_OF_RANGE_ERROR int = 3048
    ER_STD_OVERFLOW_ERROR int = 3049
    ER_STD_RANGE_ERROR int = 3050
    ER_STD_UNDERFLOW_ERROR int = 3051
    ER_STD_LOGIC_ERROR int = 3052
    ER_STD_RUNTIME_ERROR int = 3053
    ER_STD_UNKNOWN_EXCEPTION int = 3054
    ER_GIS_DATA_WRONG_ENDIANESS int = 3055
    ER_CHANGE_MASTER_PASSWORD_LENGTH int = 3056
    ER_USER_LOCK_WRONG_NAME int = 3057
    ER_USER_LOCK_DEADLOCK int = 3058
    ER_REPLACE_INACCESSIBLE_ROWS int = 3059
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_GIS int = 3060
    ER_ILLEGAL_USER_VAR int = 3061
    ER_GTID_MODE_OFF int = 3062
    ER_UNSUPPORTED_BY_REPLICATION_THREAD int = 3063
    ER_INCORRECT_TYPE int = 3064
    ER_FIELD_IN_ORDER_NOT_SELECT int = 3065
    ER_AGGREGATE_IN_ORDER_NOT_SELECT int = 3066
    ER_INVALID_RPL_WILD_TABLE_FILTER_PATTERN int = 3067
    ER_NET_OK_PACKET_TOO_LARGE int = 3068
    ER_INVALID_JSON_DATA int = 3069
    ER_INVALID_GEOJSON_MISSING_MEMBER int = 3070
    ER_INVALID_GEOJSON_WRONG_TYPE int = 3071
    ER_INVALID_GEOJSON_UNSPECIFIED int = 3072
    ER_DIMENSION_UNSUPPORTED int = 3073
    ER_SLAVE_CHANNEL_DOES_NOT_EXIST int = 3074
    ER_SLAVE_MULTIPLE_CHANNELS_HOST_PORT int = 3075
    ER_SLAVE_CHANNEL_NAME_INVALID_OR_TOO_LONG int = 3076
    ER_SLAVE_NEW_CHANNEL_WRONG_REPOSITORY int = 3077
    ER_SLAVE_CHANNEL_DELETE int = 3078
    ER_SLAVE_MULTIPLE_CHANNELS_CMD int = 3079
    ER_SLAVE_MAX_CHANNELS_EXCEEDED int = 3080
    ER_SLAVE_CHANNEL_MUST_STOP int = 3081
    ER_SLAVE_CHANNEL_NOT_RUNNING int = 3082
    ER_SLAVE_CHANNEL_WAS_RUNNING int = 3083
    ER_SLAVE_CHANNEL_WAS_NOT_RUNNING int = 3084
    ER_SLAVE_CHANNEL_SQL_THREAD_MUST_STOP int = 3085
    ER_SLAVE_CHANNEL_SQL_SKIP_COUNTER int = 3086
    ER_WRONG_FIELD_WITH_GROUP_V2 int = 3087
    ER_MIX_OF_GROUP_FUNC_AND_FIELDS_V2 int = 3088
    ER_WARN_DEPRECATED_SYSVAR_UPDATE int = 3089
    ER_WARN_DEPRECATED_SQLMODE int = 3090
    ER_CANNOT_LOG_PARTIAL_DROP_DATABASE_WITH_GTID int = 3091
    ER_GROUP_REPLICATION_CONFIGURATION int = 3092
    ER_GROUP_REPLICATION_RUNNING int = 3093
    ER_GROUP_REPLICATION_APPLIER_INIT_ERROR int = 3094
    ER_GROUP_REPLICATION_STOP_APPLIER_THREAD_TIMEOUT int = 3095
    ER_GROUP_REPLICATION_COMMUNICATION_LAYER_SESSION_ERROR int = 3096
    ER_GROUP_REPLICATION_COMMUNICATION_LAYER_JOIN_ERROR int = 3097
    ER_BEFORE_DML_VALIDATION_ERROR int = 3098
    ER_PREVENTS_VARIABLE_WITHOUT_RBR int = 3099
    ER_RUN_HOOK_ERROR int = 3100
    ER_TRANSACTION_ROLLBACK_DURING_COMMIT int = 3101
    ER_GENERATED_COLUMN_FUNCTION_IS_NOT_ALLOWED int = 3102
    ER_KEY_BASED_ON_GENERATED_COLUMN int = 3103
    ER_UNSUPPORTED_ALTER_INPLACE_ON_VIRTUAL_COLUMN int = 3103
    ER_WRONG_FK_OPTION_FOR_GENERATED_COLUMN int = 3104
    ER_NON_DEFAULT_VALUE_FOR_GENERATED_COLUMN int = 3105
    ER_UNSUPPORTED_ACTION_ON_GENERATED_COLUMN int = 3106
    ER_GENERATED_COLUMN_NON_PRIOR int = 3107
    ER_DEPENDENT_BY_GENERATED_COLUMN int = 3108
    ER_GENERATED_COLUMN_REF_AUTO_INC int = 3109
    ER_FEATURE_NOT_AVAILABLE int = 3110
    ER_CANT_SET_GTID_MODE int = 3111
    ER_CANT_USE_AUTO_POSITION_WITH_GTID_MODE_OFF int = 3112
    ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION int = 3113
    ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON int = 3114
    ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF int = 3115
    ER_CANT_SET_ENFORCE_GTID_CONSISTENCY_ON_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS int = 3116
    ER_SET_ENFORCE_GTID_CONSISTENCY_WARN_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS int = 3117
    ER_ACCOUNT_HAS_BEEN_LOCKED int = 3118
    ER_WRONG_TABLESPACE_NAME int = 3119
    ER_TABLESPACE_IS_NOT_EMPTY int = 3120
    ER_WRONG_FILE_NAME int = 3121
    ER_BOOST_GEOMETRY_INCONSISTENT_TURNS_EXCEPTION int = 3122
    ER_WARN_OPTIMIZER_HINT_SYNTAX_ERROR int = 3123
    ER_WARN_BAD_MAX_EXECUTION_TIME int = 3124
    ER_WARN_UNSUPPORTED_MAX_EXECUTION_TIME int = 3125
    ER_WARN_CONFLICTING_HINT int = 3126
    ER_WARN_UNKNOWN_QB_NAME int = 3127
    ER_UNRESOLVED_HINT_NAME int = 3128
    ER_WARN_DEPRECATED_SQLMODE_UNSET int = 3129
    ER_WARN_ON_MODIFYING_GTID_EXECUTED_TABLE int = 3129
    ER_PLUGGABLE_PROTOCOL_COMMAND_NOT_SUPPORTED int = 3130
    ER_LOCKING_SERVICE_WRONG_NAME int = 3131
    ER_LOCKING_SERVICE_DEADLOCK int = 3132
    ER_LOCKING_SERVICE_TIMEOUT int = 3133
    ER_GIS_MAX_POINTS_IN_GEOMETRY_OVERFLOWED int = 3134
    ER_SQL_MODE_MERGED int = 3135
    ER_VTOKEN_PLUGIN_TOKEN_MISMATCH int = 3136
    ER_VTOKEN_PLUGIN_TOKEN_NOT_FOUND int = 3137
    ER_CANT_SET_VARIABLE_WHEN_OWNING_GTID int = 3138
    ER_SLAVE_CHANNEL_OPERATION_NOT_ALLOWED int = 3139
    ER_INVALID_JSON_TEXT int = 3140
    ER_INVALID_JSON_TEXT_IN_PARAM int = 3141
    ER_INVALID_JSON_BINARY_DATA int = 3142
    ER_INVALID_JSON_PATH int = 3143
    ER_INVALID_JSON_CHARSET int = 3144
    ER_INVALID_JSON_CHARSET_IN_FUNCTION int = 3145
    ER_INVALID_TYPE_FOR_JSON int = 3146
    ER_INVALID_CAST_TO_JSON int = 3147
    ER_INVALID_JSON_PATH_CHARSET int = 3148
    ER_INVALID_JSON_PATH_WILDCARD int = 3149
    ER_JSON_VALUE_TOO_BIG int = 3150
    ER_JSON_KEY_TOO_BIG int = 3151
    ER_JSON_USED_AS_KEY int = 3152
    ER_JSON_VACUOUS_PATH int = 3153
    ER_JSON_BAD_ONE_OR_ALL_ARG int = 3154
    ER_NUMERIC_JSON_VALUE_OUT_OF_RANGE int = 3155
    ER_INVALID_JSON_VALUE_FOR_CAST int = 3156
    ER_JSON_DOCUMENT_TOO_DEEP int = 3157
    ER_JSON_DOCUMENT_NULL_KEY int = 3158
    ER_SECURE_TRANSPORT_REQUIRED int = 3159
    ER_NO_SECURE_TRANSPORTS_CONFIGURED int = 3160
    ER_DISABLED_STORAGE_ENGINE int = 3161
    ER_USER_DOES_NOT_EXIST int = 3162
    ER_USER_ALREADY_EXISTS int = 3163
    ER_AUDIT_API_ABORT int = 3164
    ER_INVALID_JSON_PATH_ARRAY_CELL int = 3165
    ER_BUFPOOL_RESIZE_INPROGRESS int = 3166
    ER_FEATURE_DISABLED_SEE_DOC int = 3167
    ER_SERVER_ISNT_AVAILABLE int = 3168
    ER_SESSION_WAS_KILLED int = 3169
    ER_CAPACITY_EXCEEDED int = 3170
    ER_CAPACITY_EXCEEDED_IN_RANGE_OPTIMIZER int = 3171
    ER_TABLE_NEEDS_UPG_PART int = 3172
    ER_CANT_WAIT_FOR_EXECUTED_GTID_SET_WHILE_OWNING_A_GTID int = 3173
    ER_CANNOT_ADD_FOREIGN_BASE_COL_VIRTUAL int = 3174
    ER_CANNOT_CREATE_VIRTUAL_INDEX_CONSTRAINT int = 3175
    ER_ERROR_ON_MODIFYING_GTID_EXECUTED_TABLE int = 3176
    ER_LOCK_REFUSED_BY_ENGINE int = 3177
    ER_UNSUPPORTED_ALTER_ONLINE_ON_VIRTUAL_COLUMN int = 3178
    ER_MASTER_KEY_ROTATION_NOT_SUPPORTED_BY_SE int = 3179
    ER_MASTER_KEY_ROTATION_ERROR_BY_SE int = 3180
    ER_MASTER_KEY_ROTATION_BINLOG_FAILED int = 3181
    ER_MASTER_KEY_ROTATION_SE_UNAVAILABLE int = 3182
    ER_TABLESPACE_CANNOT_ENCRYPT int = 3183
    ER_INVALID_ENCRYPTION_OPTION int = 3184
    ER_CANNOT_FIND_KEY_IN_KEYRING int = 3185
    ER_CAPACITY_EXCEEDED_IN_PARSER int = 3186
    ER_UNSUPPORTED_ALTER_ENCRYPTION_INPLACE int = 3187
    ER_KEYRING_UDF_KEYRING_SERVICE_ERROR int = 3188
    ER_USER_COLUMN_OLD_LENGTH int = 3189
    ER_CANT_RESET_MASTER int = 3190
    ER_GROUP_REPLICATION_MAX_GROUP_SIZE int = 3191
    ER_CANNOT_ADD_FOREIGN_BASE_COL_STORED int = 3192
    ER_TABLE_REFERENCED int = 3193
    ER_PARTITION_ENGINE_DEPRECATED_FOR_TABLE int = 3194
    ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID_ZERO int = 3195
    ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID int = 3196
    ER_XA_RETRY int = 3197
    ER_KEYRING_AWS_UDF_AWS_KMS_ERROR int = 3198
    ER_BINLOG_UNSAFE_XA int = 3199
    ER_UDF_ERROR int = 3200
    ER_KEYRING_MIGRATION_FAILURE int = 3201
    ER_KEYRING_ACCESS_DENIED_ERROR int = 3202
    ER_KEYRING_MIGRATION_STATUS int = 3203
    ER_PLUGIN_FAILED_TO_OPEN_TABLES int = 3204
    ER_PLUGIN_FAILED_TO_OPEN_TABLE int = 3205
    ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED int = 3206
    ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET int = 3207
    ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY int = 3208
    ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED int = 3209
    ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED int = 3210
    ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE int = 3211
    ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED int = 3212
    ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS int = 3213
    ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE int = 3214
    ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT int = 3215
    ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED int = 3216
    ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE int = 3217
    ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_VALUE int = 3218
    ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR int = 3219
    ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY int = 3220
    ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY int = 3221
    ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXISTS int = 3222
    ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC int = 3223
    ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER int = 3224
    ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER int = 3225
    WARN_DEPRECATED_MAXDB_SQL_MODE_FOR_TIMESTAMP int = 3226
    ER_XA_REPLICATION_FILTERS int = 3227
    ER_CANT_OPEN_ERROR_LOG int = 3228
    ER_GROUPING_ON_TIMESTAMP_IN_DST int = 3229
    ER_CANT_START_SERVER_NAMED_PIPE int = 3230
    CR_UNKNOWN_ERROR int = 2000
    CR_SOCKET_CREATE_ERROR int = 2001
    CR_CONNECTION_ERROR int = 2002
    CR_CONN_HOST_ERROR int = 2003
    CR_IPSOCK_ERROR int = 2004
    CR_UNKNOWN_HOST int = 2005
    CR_SERVER_GONE_ERROR int = 2006
    CR_VERSION_ERROR int = 2007
    CR_OUT_OF_MEMORY int = 2008
    CR_WRONG_HOST_INFO int = 2009
    CR_LOCALHOST_CONNECTION int = 2010
    CR_TCP_CONNECTION int = 2011
    CR_SERVER_HANDSHAKE_ERR int = 2012
    CR_SERVER_LOST int = 2013
    CR_COMMANDS_OUT_OF_SYNC int = 2014
    CR_NAMEDPIPE_CONNECTION int = 2015
    CR_NAMEDPIPEWAIT_ERROR int = 2016
    CR_NAMEDPIPEOPEN_ERROR int = 2017
    CR_NAMEDPIPESETSTATE_ERROR int = 2018
    CR_CANT_READ_CHARSET int = 2019
    CR_NET_PACKET_TOO_LARGE int = 2020
    CR_EMBEDDED_CONNECTION int = 2021
    CR_PROBE_SLAVE_STATUS int = 2022
    CR_PROBE_SLAVE_HOSTS int = 2023
    CR_PROBE_SLAVE_CONNECT int = 2024
    CR_PROBE_MASTER_CONNECT int = 2025
    CR_SSL_CONNECTION_ERROR int = 2026
    CR_MALFORMED_PACKET int = 2027
    CR_WRONG_LICENSE int = 2028
    CR_NULL_POINTER int = 2029
    CR_NO_PREPARE_STMT int = 2030
    CR_PARAMS_NOT_BOUND int = 2031
    CR_DATA_TRUNCATED int = 2032
    CR_NO_PARAMETERS_EXISTS int = 2033
    CR_INVALID_PARAMETER_NO int = 2034
    CR_INVALID_BUFFER_USE int = 2035
    CR_UNSUPPORTED_PARAM_TYPE int = 2036
    CR_SHARED_MEMORY_CONNECTION int = 2037
    CR_SHARED_MEMORY_CONNECT_REQUEST_ERROR int = 2038
    CR_SHARED_MEMORY_CONNECT_ANSWER_ERROR int = 2039
    CR_SHARED_MEMORY_CONNECT_FILE_MAP_ERROR int = 2040
    CR_SHARED_MEMORY_CONNECT_MAP_ERROR int = 2041
    CR_SHARED_MEMORY_FILE_MAP_ERROR int = 2042
    CR_SHARED_MEMORY_MAP_ERROR int = 2043
    CR_SHARED_MEMORY_EVENT_ERROR int = 2044
    CR_SHARED_MEMORY_CONNECT_ABANDONED_ERROR int = 2045
    CR_SHARED_MEMORY_CONNECT_SET_ERROR int = 2046
    CR_CONN_UNKNOW_PROTOCOL int = 2047
    CR_INVALID_CONN_HANDLE int = 2048
    CR_SECURE_AUTH int = 2049
    CR_UNUSED_1 int = 2049
    CR_FETCH_CANCELED int = 2050
    CR_NO_DATA int = 2051
    CR_NO_STMT_METADATA int = 2052
    CR_NO_RESULT_SET int = 2053
    CR_NOT_IMPLEMENTED int = 2054
    CR_SERVER_LOST_EXTENDED int = 2055
    CR_STMT_CLOSED int = 2056
    CR_NEW_STMT_METADATA int = 2057
    CR_ALREADY_CONNECTED int = 2058
    CR_AUTH_PLUGIN_CANNOT_LOAD int = 2059
    CR_DUPLICATE_CONNECTION_ATTR int = 2060
    CR_AUTH_PLUGIN_ERR int = 2061
    CR_INSECURE_API_ERR int = 2062
    EE_CANTCREATEFILE int = 1
    EE_READ int = 2
    EE_WRITE int = 3
    EE_BADCLOSE int = 4
    EE_OUTOFMEMORY int = 5
    EE_DELETE int = 6
    EE_LINK int = 7
    EE_EOFERR int = 9
    EE_CANTLOCK int = 10
    EE_CANTUNLOCK int = 11
    EE_DIR int = 12
    EE_STAT int = 13
    EE_CANT_CHSIZE int = 14
    EE_CANT_OPEN_STREAM int = 15
    EE_GETWD int = 16
    EE_SETWD int = 17
    EE_LINK_WARNING int = 18
    EE_OPEN_WARNING int = 19
    EE_DISK_FULL int = 20
    EE_CANT_MKDIR int = 21
    EE_UNKNOWN_CHARSET int = 22
    EE_OUT_OF_FILERESOURCES int = 23
    EE_CANT_READLINK int = 24
    EE_CANT_SYMLINK int = 25
    EE_REALPATH int = 26
    EE_SYNC int = 27
    EE_UNKNOWN_COLLATION int = 28
    EE_FILENOTFOUND int = 29
    EE_FILE_NOT_CLOSED int = 30
    EE_CHANGE_OWNERSHIP int = 31
    EE_CHANGE_PERMISSIONS int = 32
    EE_CANT_SEEK int = 33
    EE_CAPACITY_EXCEEDED int = 34
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


// IsServerErrorHashchk check mysql error is "hashchk" 
func IsServerErrorHashchk(err error) bool {
    result := Isa(err, ER_HASHCHK)
    return result
}

    
// IsServerErrorNisamchk check mysql error is "isamchk" 
func IsServerErrorNisamchk(err error) bool {
    result := Isa(err, ER_NISAMCHK)
    return result
}

    
// IsServerErrorNo check mysql error is "NO" 
func IsServerErrorNo(err error) bool {
    result := Isa(err, ER_NO)
    return result
}

    
// IsServerErrorYes check mysql error is "YES" 
func IsServerErrorYes(err error) bool {
    result := Isa(err, ER_YES)
    return result
}

    
// IsServerErrorCantCreateFile check mysql error is "Can't create file '%s' (errno: %d - %s)" 
func IsServerErrorCantCreateFile(err error) bool {
    result := Isa(err, ER_CANT_CREATE_FILE)
    return result
}

    
// IsServerErrorCantCreateTable check mysql error is "Can't create table '%s' (errno: %d)" 
func IsServerErrorCantCreateTable(err error) bool {
    result := Isa(err, ER_CANT_CREATE_TABLE)
    return result
}

    
// IsServerErrorCantCreateDb check mysql error is "Can't create database '%s' (errno: %d) " 
func IsServerErrorCantCreateDb(err error) bool {
    result := Isa(err, ER_CANT_CREATE_DB)
    return result
}

    
// IsServerErrorDbCreateExists check mysql error is "Can't create database '%s'" 
func IsServerErrorDbCreateExists(err error) bool {
    result := Isa(err, ER_DB_CREATE_EXISTS)
    return result
}

    
// IsServerErrorDbDropExists check mysql error is "Can't drop database '%s'" 
func IsServerErrorDbDropExists(err error) bool {
    result := Isa(err, ER_DB_DROP_EXISTS)
    return result
}

    
// IsServerErrorDbDropDelete check mysql error is "Error dropping database (can't delete '%s', errno: %d) " 
func IsServerErrorDbDropDelete(err error) bool {
    result := Isa(err, ER_DB_DROP_DELETE)
    return result
}

    
// IsServerErrorDbDropRmdir check mysql error is "Error dropping database (can't rmdir '%s', errno: %d) " 
func IsServerErrorDbDropRmdir(err error) bool {
    result := Isa(err, ER_DB_DROP_RMDIR)
    return result
}

    
// IsServerErrorCantDeleteFile check mysql error is "Error on delete of '%s' (errno: %d - %s) " 
func IsServerErrorCantDeleteFile(err error) bool {
    result := Isa(err, ER_CANT_DELETE_FILE)
    return result
}

    
// IsServerErrorCantFindSystemRec check mysql error is "Can't read record in system table" 
func IsServerErrorCantFindSystemRec(err error) bool {
    result := Isa(err, ER_CANT_FIND_SYSTEM_REC)
    return result
}

    
// IsServerErrorCantGetStat check mysql error is "Can't get status of '%s' (errno: %d - %s) " 
func IsServerErrorCantGetStat(err error) bool {
    result := Isa(err, ER_CANT_GET_STAT)
    return result
}

    
// IsServerErrorCantGetWd check mysql error is "Can't get working directory (errno: %d - %s) " 
func IsServerErrorCantGetWd(err error) bool {
    result := Isa(err, ER_CANT_GET_WD)
    return result
}

    
// IsServerErrorCantLock check mysql error is "Can't lock file (errno: %d - %s) " 
func IsServerErrorCantLock(err error) bool {
    result := Isa(err, ER_CANT_LOCK)
    return result
}

    
// IsServerErrorCantOpenFile check mysql error is "Can't open file: '%s' (errno: %d - %s)" 
func IsServerErrorCantOpenFile(err error) bool {
    result := Isa(err, ER_CANT_OPEN_FILE)
    return result
}

    
// IsServerErrorFileNotFound check mysql error is "Can't find file: '%s' (errno: %d - %s) " 
func IsServerErrorFileNotFound(err error) bool {
    result := Isa(err, ER_FILE_NOT_FOUND)
    return result
}

    
// IsServerErrorCantReadDir check mysql error is "Can't read dir of '%s' (errno: %d - %s) " 
func IsServerErrorCantReadDir(err error) bool {
    result := Isa(err, ER_CANT_READ_DIR)
    return result
}

    
// IsServerErrorCantSetWd check mysql error is "Can't change dir to '%s' (errno: %d - %s) " 
func IsServerErrorCantSetWd(err error) bool {
    result := Isa(err, ER_CANT_SET_WD)
    return result
}

    
// IsServerErrorCheckread check mysql error is "Record has changed since last read in table '%s' " 
func IsServerErrorCheckread(err error) bool {
    result := Isa(err, ER_CHECKREAD)
    return result
}

    
// IsServerErrorDiskFull check mysql error is "Disk full (%s)" 
func IsServerErrorDiskFull(err error) bool {
    result := Isa(err, ER_DISK_FULL)
    return result
}

    
// IsServerErrorDupKey check mysql error is "Can't write" 
func IsServerErrorDupKey(err error) bool {
    result := Isa(err, ER_DUP_KEY)
    return result
}

    
// IsServerErrorErrorOnClose check mysql error is "Error on close of '%s' (errno: %d - %s) " 
func IsServerErrorErrorOnClose(err error) bool {
    result := Isa(err, ER_ERROR_ON_CLOSE)
    return result
}

    
// IsServerErrorErrorOnRead check mysql error is "Error reading file '%s' (errno: %d - %s) " 
func IsServerErrorErrorOnRead(err error) bool {
    result := Isa(err, ER_ERROR_ON_READ)
    return result
}

    
// IsServerErrorErrorOnRename check mysql error is "Error on rename of '%s' to '%s' (errno: %d - %s) " 
func IsServerErrorErrorOnRename(err error) bool {
    result := Isa(err, ER_ERROR_ON_RENAME)
    return result
}

    
// IsServerErrorErrorOnWrite check mysql error is "Error writing file '%s' (errno: %d - %s) " 
func IsServerErrorErrorOnWrite(err error) bool {
    result := Isa(err, ER_ERROR_ON_WRITE)
    return result
}

    
// IsServerErrorFileUsed check mysql error is "'%s' is locked against change " 
func IsServerErrorFileUsed(err error) bool {
    result := Isa(err, ER_FILE_USED)
    return result
}

    
// IsServerErrorFilsortAbort check mysql error is "Sort aborted " 
func IsServerErrorFilsortAbort(err error) bool {
    result := Isa(err, ER_FILSORT_ABORT)
    return result
}

    
// IsServerErrorFormNotFound check mysql error is "View '%s' doesn't exist for '%s' " 
func IsServerErrorFormNotFound(err error) bool {
    result := Isa(err, ER_FORM_NOT_FOUND)
    return result
}

    
// IsServerErrorGetErrno check mysql error is "Got error %d from storage engine" 
func IsServerErrorGetErrno(err error) bool {
    result := Isa(err, ER_GET_ERRNO)
    return result
}

    
// IsServerErrorIllegalHa check mysql error is "Table storage engine for '%s' doesn't have this option " 
func IsServerErrorIllegalHa(err error) bool {
    result := Isa(err, ER_ILLEGAL_HA)
    return result
}

    
// IsServerErrorKeyNotFound check mysql error is "Can't find record in '%s' " 
func IsServerErrorKeyNotFound(err error) bool {
    result := Isa(err, ER_KEY_NOT_FOUND)
    return result
}

    
// IsServerErrorNotFormFile check mysql error is "Incorrect information in file: '%s' " 
func IsServerErrorNotFormFile(err error) bool {
    result := Isa(err, ER_NOT_FORM_FILE)
    return result
}

    
// IsServerErrorNotKeyfile check mysql error is "Incorrect key file for table '%s'" 
func IsServerErrorNotKeyfile(err error) bool {
    result := Isa(err, ER_NOT_KEYFILE)
    return result
}

    
// IsServerErrorOldKeyfile check mysql error is "Old key file for table '%s'" 
func IsServerErrorOldKeyfile(err error) bool {
    result := Isa(err, ER_OLD_KEYFILE)
    return result
}

    
// IsServerErrorOpenAsReadonly check mysql error is "Table '%s' is read only " 
func IsServerErrorOpenAsReadonly(err error) bool {
    result := Isa(err, ER_OPEN_AS_READONLY)
    return result
}

    
// IsServerErrorOutofmemory check mysql error is "Out of memory" 
func IsServerErrorOutofmemory(err error) bool {
    result := Isa(err, ER_OUTOFMEMORY)
    return result
}

    
// IsServerErrorOutOfSortmemory check mysql error is "Out of sort memory, consider increasing server sort buffer size " 
func IsServerErrorOutOfSortmemory(err error) bool {
    result := Isa(err, ER_OUT_OF_SORTMEMORY)
    return result
}

    
// IsServerErrorUnexpectedEof check mysql error is "Unexpected EOF found when reading file '%s' (errno: %d - %s) " 
func IsServerErrorUnexpectedEof(err error) bool {
    result := Isa(err, ER_UNEXPECTED_EOF)
    return result
}

    
// IsServerErrorConCountError check mysql error is "Too many connections " 
func IsServerErrorConCountError(err error) bool {
    result := Isa(err, ER_CON_COUNT_ERROR)
    return result
}

    
// IsServerErrorOutOfResources check mysql error is "Out of memory" 
func IsServerErrorOutOfResources(err error) bool {
    result := Isa(err, ER_OUT_OF_RESOURCES)
    return result
}

    
// IsServerErrorBadHostError check mysql error is "Can't get hostname for your address " 
func IsServerErrorBadHostError(err error) bool {
    result := Isa(err, ER_BAD_HOST_ERROR)
    return result
}

    
// IsServerErrorHandshakeError check mysql error is "Bad handshake " 
func IsServerErrorHandshakeError(err error) bool {
    result := Isa(err, ER_HANDSHAKE_ERROR)
    return result
}

    
// IsServerErrorDbaccessDeniedError check mysql error is "Access denied for user '%s'@'%s' to database '%s' " 
func IsServerErrorDbaccessDeniedError(err error) bool {
    result := Isa(err, ER_DBACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorAccessDeniedError check mysql error is "Access denied for user '%s'@'%s' (using password: %s) " 
func IsServerErrorAccessDeniedError(err error) bool {
    result := Isa(err, ER_ACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorNoDbError check mysql error is "No database selected " 
func IsServerErrorNoDbError(err error) bool {
    result := Isa(err, ER_NO_DB_ERROR)
    return result
}

    
// IsServerErrorUnknownComError check mysql error is "Unknown command " 
func IsServerErrorUnknownComError(err error) bool {
    result := Isa(err, ER_UNKNOWN_COM_ERROR)
    return result
}

    
// IsServerErrorBadNullError check mysql error is "Column '%s' cannot be null " 
func IsServerErrorBadNullError(err error) bool {
    result := Isa(err, ER_BAD_NULL_ERROR)
    return result
}

    
// IsServerErrorBadDbError check mysql error is "Unknown database '%s' " 
func IsServerErrorBadDbError(err error) bool {
    result := Isa(err, ER_BAD_DB_ERROR)
    return result
}

    
// IsServerErrorTableExistsError check mysql error is "Table '%s' already exists " 
func IsServerErrorTableExistsError(err error) bool {
    result := Isa(err, ER_TABLE_EXISTS_ERROR)
    return result
}

    
// IsServerErrorBadTableError check mysql error is "Unknown table '%s' " 
func IsServerErrorBadTableError(err error) bool {
    result := Isa(err, ER_BAD_TABLE_ERROR)
    return result
}

    
// IsServerErrorNonUniqError check mysql error is "Column '%s' in %s is ambiguous %s = column name %s = location of column (for example, "field list") Likely cause: A column appears in a query without appropriate qualification, such as in a select list or ON clause." 
func IsServerErrorNonUniqError(err error) bool {
    result := Isa(err, ER_NON_UNIQ_ERROR)
    return result
}

    
// IsServerErrorServerShutdown check mysql error is "Server shutdown in progress " 
func IsServerErrorServerShutdown(err error) bool {
    result := Isa(err, ER_SERVER_SHUTDOWN)
    return result
}

    
// IsServerErrorBadFieldError check mysql error is "Unknown column '%s' in '%s' " 
func IsServerErrorBadFieldError(err error) bool {
    result := Isa(err, ER_BAD_FIELD_ERROR)
    return result
}

    
// IsServerErrorWrongFieldWithGroup check mysql error is "'%s' isn't in GROUP BY " 
func IsServerErrorWrongFieldWithGroup(err error) bool {
    result := Isa(err, ER_WRONG_FIELD_WITH_GROUP)
    return result
}

    
// IsServerErrorWrongGroupField check mysql error is "Can't group on '%s' " 
func IsServerErrorWrongGroupField(err error) bool {
    result := Isa(err, ER_WRONG_GROUP_FIELD)
    return result
}

    
// IsServerErrorWrongSumSelect check mysql error is "Statement has sum functions and columns in same statement " 
func IsServerErrorWrongSumSelect(err error) bool {
    result := Isa(err, ER_WRONG_SUM_SELECT)
    return result
}

    
// IsServerErrorWrongValueCount check mysql error is "Column count doesn't match value count " 
func IsServerErrorWrongValueCount(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_COUNT)
    return result
}

    
// IsServerErrorTooLongIdent check mysql error is "Identifier name '%s' is too long " 
func IsServerErrorTooLongIdent(err error) bool {
    result := Isa(err, ER_TOO_LONG_IDENT)
    return result
}

    
// IsServerErrorDupFieldname check mysql error is "Duplicate column name '%s' " 
func IsServerErrorDupFieldname(err error) bool {
    result := Isa(err, ER_DUP_FIELDNAME)
    return result
}

    
// IsServerErrorDupKeyname check mysql error is "Duplicate key name '%s' " 
func IsServerErrorDupKeyname(err error) bool {
    result := Isa(err, ER_DUP_KEYNAME)
    return result
}

    
// IsServerErrorDupEntry check mysql error is "Duplicate entry '%s' for key %d" 
func IsServerErrorDupEntry(err error) bool {
    result := Isa(err, ER_DUP_ENTRY)
    return result
}

    
// IsServerErrorWrongFieldSpec check mysql error is "Incorrect column specifier for column '%s' " 
func IsServerErrorWrongFieldSpec(err error) bool {
    result := Isa(err, ER_WRONG_FIELD_SPEC)
    return result
}

    
// IsServerErrorParseError check mysql error is "%s near '%s' at line %d " 
func IsServerErrorParseError(err error) bool {
    result := Isa(err, ER_PARSE_ERROR)
    return result
}

    
// IsServerErrorEmptyQuery check mysql error is "Query was empty " 
func IsServerErrorEmptyQuery(err error) bool {
    result := Isa(err, ER_EMPTY_QUERY)
    return result
}

    
// IsServerErrorNonuniqTable check mysql error is "Not unique table/alias: '%s' " 
func IsServerErrorNonuniqTable(err error) bool {
    result := Isa(err, ER_NONUNIQ_TABLE)
    return result
}

    
// IsServerErrorInvalidDefault check mysql error is "Invalid default value for '%s' " 
func IsServerErrorInvalidDefault(err error) bool {
    result := Isa(err, ER_INVALID_DEFAULT)
    return result
}

    
// IsServerErrorMultiplePriKey check mysql error is "Multiple primary key defined " 
func IsServerErrorMultiplePriKey(err error) bool {
    result := Isa(err, ER_MULTIPLE_PRI_KEY)
    return result
}

    
// IsServerErrorTooManyKeys check mysql error is "Too many keys specified" 
func IsServerErrorTooManyKeys(err error) bool {
    result := Isa(err, ER_TOO_MANY_KEYS)
    return result
}

    
// IsServerErrorTooManyKeyParts check mysql error is "Too many key parts specified" 
func IsServerErrorTooManyKeyParts(err error) bool {
    result := Isa(err, ER_TOO_MANY_KEY_PARTS)
    return result
}

    
// IsServerErrorTooLongKey check mysql error is "Specified key was too long" 
func IsServerErrorTooLongKey(err error) bool {
    result := Isa(err, ER_TOO_LONG_KEY)
    return result
}

    
// IsServerErrorKeyColumnDoesNotExits check mysql error is "Key column '%s' doesn't exist in table " 
func IsServerErrorKeyColumnDoesNotExits(err error) bool {
    result := Isa(err, ER_KEY_COLUMN_DOES_NOT_EXITS)
    return result
}

    
// IsServerErrorBlobUsedAsKey check mysql error is "BLOB column '%s' can't be used in key specification with the used table type " 
func IsServerErrorBlobUsedAsKey(err error) bool {
    result := Isa(err, ER_BLOB_USED_AS_KEY)
    return result
}

    
// IsServerErrorTooBigFieldlength check mysql error is "Column length too big for column '%s' (max = %lu)" 
func IsServerErrorTooBigFieldlength(err error) bool {
    result := Isa(err, ER_TOO_BIG_FIELDLENGTH)
    return result
}

    
// IsServerErrorWrongAutoKey check mysql error is "Incorrect table definition" 
func IsServerErrorWrongAutoKey(err error) bool {
    result := Isa(err, ER_WRONG_AUTO_KEY)
    return result
}

    
// IsServerErrorReady check mysql error is "%s: ready for connections. Version: '%s' socket: '%s' port: %d " 
func IsServerErrorReady(err error) bool {
    result := Isa(err, ER_READY)
    return result
}

    
// IsServerErrorNormalShutdown check mysql error is "%s: Normal shutdown " 
func IsServerErrorNormalShutdown(err error) bool {
    result := Isa(err, ER_NORMAL_SHUTDOWN)
    return result
}

    
// IsServerErrorGotSignal check mysql error is "%s: Got signal %d. Aborting! " 
func IsServerErrorGotSignal(err error) bool {
    result := Isa(err, ER_GOT_SIGNAL)
    return result
}

    
// IsServerErrorShutdownComplete check mysql error is "%s: Shutdown complete " 
func IsServerErrorShutdownComplete(err error) bool {
    result := Isa(err, ER_SHUTDOWN_COMPLETE)
    return result
}

    
// IsServerErrorForcingClose check mysql error is "%s: Forcing close of thread %ld user: '%s' " 
func IsServerErrorForcingClose(err error) bool {
    result := Isa(err, ER_FORCING_CLOSE)
    return result
}

    
// IsServerErrorIpsockError check mysql error is "Can't create IP socket " 
func IsServerErrorIpsockError(err error) bool {
    result := Isa(err, ER_IPSOCK_ERROR)
    return result
}

    
// IsServerErrorNoSuchIndex check mysql error is "Table '%s' has no index like the one used in CREATE INDEX" 
func IsServerErrorNoSuchIndex(err error) bool {
    result := Isa(err, ER_NO_SUCH_INDEX)
    return result
}

    
// IsServerErrorWrongFieldTerminators check mysql error is "Field separator argument is not what is expected" 
func IsServerErrorWrongFieldTerminators(err error) bool {
    result := Isa(err, ER_WRONG_FIELD_TERMINATORS)
    return result
}

    
// IsServerErrorBlobsAndNoTerminated check mysql error is "You can't use fixed rowlength with BLOBs" 
func IsServerErrorBlobsAndNoTerminated(err error) bool {
    result := Isa(err, ER_BLOBS_AND_NO_TERMINATED)
    return result
}

    
// IsServerErrorTextfileNotReadable check mysql error is "The file '%s' must be in the database directory or be readable by all " 
func IsServerErrorTextfileNotReadable(err error) bool {
    result := Isa(err, ER_TEXTFILE_NOT_READABLE)
    return result
}

    
// IsServerErrorFileExistsError check mysql error is "File '%s' already exists " 
func IsServerErrorFileExistsError(err error) bool {
    result := Isa(err, ER_FILE_EXISTS_ERROR)
    return result
}

    
// IsServerErrorLoadInfo check mysql error is "Records: %ld Deleted: %ld Skipped: %ld Warnings: %ld " 
func IsServerErrorLoadInfo(err error) bool {
    result := Isa(err, ER_LOAD_INFO)
    return result
}

    
// IsServerErrorAlterInfo check mysql error is "Records: %ld Duplicates: %ld " 
func IsServerErrorAlterInfo(err error) bool {
    result := Isa(err, ER_ALTER_INFO)
    return result
}

    
// IsServerErrorWrongSubKey check mysql error is "Incorrect prefix key" 
func IsServerErrorWrongSubKey(err error) bool {
    result := Isa(err, ER_WRONG_SUB_KEY)
    return result
}

    
// IsServerErrorCantRemoveAllFields check mysql error is "You can't delete all columns with ALTER TABLE" 
func IsServerErrorCantRemoveAllFields(err error) bool {
    result := Isa(err, ER_CANT_REMOVE_ALL_FIELDS)
    return result
}

    
// IsServerErrorCantDropFieldOrKey check mysql error is "Can't DROP '%s'" 
func IsServerErrorCantDropFieldOrKey(err error) bool {
    result := Isa(err, ER_CANT_DROP_FIELD_OR_KEY)
    return result
}

    
// IsServerErrorInsertInfo check mysql error is "Records: %ld Duplicates: %ld Warnings: %ld " 
func IsServerErrorInsertInfo(err error) bool {
    result := Isa(err, ER_INSERT_INFO)
    return result
}

    
// IsServerErrorUpdateTableUsed check mysql error is "You can't specify target table '%s' for update in FROM clause" 
func IsServerErrorUpdateTableUsed(err error) bool {
    result := Isa(err, ER_UPDATE_TABLE_USED)
    return result
}

    
// IsServerErrorNoSuchThread check mysql error is "Unknown thread id: %lu " 
func IsServerErrorNoSuchThread(err error) bool {
    result := Isa(err, ER_NO_SUCH_THREAD)
    return result
}

    
// IsServerErrorKillDeniedError check mysql error is "You are not owner of thread %lu " 
func IsServerErrorKillDeniedError(err error) bool {
    result := Isa(err, ER_KILL_DENIED_ERROR)
    return result
}

    
// IsServerErrorNoTablesUsed check mysql error is "No tables used " 
func IsServerErrorNoTablesUsed(err error) bool {
    result := Isa(err, ER_NO_TABLES_USED)
    return result
}

    
// IsServerErrorTooBigSet check mysql error is "Too many strings for column %s and SET " 
func IsServerErrorTooBigSet(err error) bool {
    result := Isa(err, ER_TOO_BIG_SET)
    return result
}

    
// IsServerErrorNoUniqueLogfile check mysql error is "Can't generate a unique log-filename %s.(1-999) " 
func IsServerErrorNoUniqueLogfile(err error) bool {
    result := Isa(err, ER_NO_UNIQUE_LOGFILE)
    return result
}

    
// IsServerErrorTableNotLockedForWrite check mysql error is "Table '%s' was locked with a READ lock and can't be updated " 
func IsServerErrorTableNotLockedForWrite(err error) bool {
    result := Isa(err, ER_TABLE_NOT_LOCKED_FOR_WRITE)
    return result
}

    
// IsServerErrorTableNotLocked check mysql error is "Table '%s' was not locked with LOCK TABLES " 
func IsServerErrorTableNotLocked(err error) bool {
    result := Isa(err, ER_TABLE_NOT_LOCKED)
    return result
}

    
// IsServerErrorBlobCantHaveDefault check mysql error is "BLOB, TEXT, GEOMETRY or JSON column '%s' can't have a default value " 
func IsServerErrorBlobCantHaveDefault(err error) bool {
    result := Isa(err, ER_BLOB_CANT_HAVE_DEFAULT)
    return result
}

    
// IsServerErrorWrongDbName check mysql error is "Incorrect database name '%s' " 
func IsServerErrorWrongDbName(err error) bool {
    result := Isa(err, ER_WRONG_DB_NAME)
    return result
}

    
// IsServerErrorWrongTableName check mysql error is "Incorrect table name '%s' " 
func IsServerErrorWrongTableName(err error) bool {
    result := Isa(err, ER_WRONG_TABLE_NAME)
    return result
}

    
// IsServerErrorTooBigSelect check mysql error is "The SELECT would examine more than MAX_JOIN_SIZE rows" 
func IsServerErrorTooBigSelect(err error) bool {
    result := Isa(err, ER_TOO_BIG_SELECT)
    return result
}

    
// IsServerErrorUnknownError check mysql error is "Unknown error " 
func IsServerErrorUnknownError(err error) bool {
    result := Isa(err, ER_UNKNOWN_ERROR)
    return result
}

    
// IsServerErrorUnknownProcedure check mysql error is "Unknown procedure '%s' " 
func IsServerErrorUnknownProcedure(err error) bool {
    result := Isa(err, ER_UNKNOWN_PROCEDURE)
    return result
}

    
// IsServerErrorWrongParamcountToProcedure check mysql error is "Incorrect parameter count to procedure '%s' " 
func IsServerErrorWrongParamcountToProcedure(err error) bool {
    result := Isa(err, ER_WRONG_PARAMCOUNT_TO_PROCEDURE)
    return result
}

    
// IsServerErrorWrongParametersToProcedure check mysql error is "Incorrect parameters to procedure '%s' " 
func IsServerErrorWrongParametersToProcedure(err error) bool {
    result := Isa(err, ER_WRONG_PARAMETERS_TO_PROCEDURE)
    return result
}

    
// IsServerErrorUnknownTable check mysql error is "Unknown table '%s' in %s " 
func IsServerErrorUnknownTable(err error) bool {
    result := Isa(err, ER_UNKNOWN_TABLE)
    return result
}

    
// IsServerErrorFieldSpecifiedTwice check mysql error is "Column '%s' specified twice " 
func IsServerErrorFieldSpecifiedTwice(err error) bool {
    result := Isa(err, ER_FIELD_SPECIFIED_TWICE)
    return result
}

    
// IsServerErrorInvalidGroupFuncUse check mysql error is "Invalid use of group function " 
func IsServerErrorInvalidGroupFuncUse(err error) bool {
    result := Isa(err, ER_INVALID_GROUP_FUNC_USE)
    return result
}

    
// IsServerErrorUnsupportedExtension check mysql error is "Table '%s' uses an extension that doesn't exist in this MySQL version " 
func IsServerErrorUnsupportedExtension(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_EXTENSION)
    return result
}

    
// IsServerErrorTableMustHaveColumns check mysql error is "A table must have at least 1 column " 
func IsServerErrorTableMustHaveColumns(err error) bool {
    result := Isa(err, ER_TABLE_MUST_HAVE_COLUMNS)
    return result
}

    
// IsServerErrorRecordFileFull check mysql error is "The table '%s' is full" 
func IsServerErrorRecordFileFull(err error) bool {
    result := Isa(err, ER_RECORD_FILE_FULL)
    return result
}

    
// IsServerErrorUnknownCharacterSet check mysql error is "Unknown character set: '%s' " 
func IsServerErrorUnknownCharacterSet(err error) bool {
    result := Isa(err, ER_UNKNOWN_CHARACTER_SET)
    return result
}

    
// IsServerErrorTooManyTables check mysql error is "Too many tables" 
func IsServerErrorTooManyTables(err error) bool {
    result := Isa(err, ER_TOO_MANY_TABLES)
    return result
}

    
// IsServerErrorTooManyFields check mysql error is "Too many columns " 
func IsServerErrorTooManyFields(err error) bool {
    result := Isa(err, ER_TOO_MANY_FIELDS)
    return result
}

    
// IsServerErrorTooBigRowsize check mysql error is "Row size too large. The maximum row size for the used table type, not counting BLOBs, is %ld. This includes storage overhead, check the manual. You have to change some columns to TEXT or BLOBs " 
func IsServerErrorTooBigRowsize(err error) bool {
    result := Isa(err, ER_TOO_BIG_ROWSIZE)
    return result
}

    
// IsServerErrorStackOverrun check mysql error is "Thread stack overrun: Used: %ld of a %ld stack. Use 'mysqld --thread_stack=#' to specify a bigger stack if needed " 
func IsServerErrorStackOverrun(err error) bool {
    result := Isa(err, ER_STACK_OVERRUN)
    return result
}

    
// IsServerErrorWrongOuterJoin check mysql error is "Cross dependency found in OUTER JOIN" 
func IsServerErrorWrongOuterJoin(err error) bool {
    result := Isa(err, ER_WRONG_OUTER_JOIN)
    return result
}

    
// IsServerErrorNullColumnInIndex check mysql error is "Table handler doesn't support NULL in given index. Please change column '%s' to be NOT NULL or use another handler " 
func IsServerErrorNullColumnInIndex(err error) bool {
    result := Isa(err, ER_NULL_COLUMN_IN_INDEX)
    return result
}

    
// IsServerErrorCantFindUdf check mysql error is "Can't load function '%s' " 
func IsServerErrorCantFindUdf(err error) bool {
    result := Isa(err, ER_CANT_FIND_UDF)
    return result
}

    
// IsServerErrorCantInitializeUdf check mysql error is "Can't initialize function '%s'" 
func IsServerErrorCantInitializeUdf(err error) bool {
    result := Isa(err, ER_CANT_INITIALIZE_UDF)
    return result
}

    
// IsServerErrorUdfNoPaths check mysql error is "No paths allowed for shared library " 
func IsServerErrorUdfNoPaths(err error) bool {
    result := Isa(err, ER_UDF_NO_PATHS)
    return result
}

    
// IsServerErrorUdfExists check mysql error is "Function '%s' already exists " 
func IsServerErrorUdfExists(err error) bool {
    result := Isa(err, ER_UDF_EXISTS)
    return result
}

    
// IsServerErrorCantOpenLibrary check mysql error is "Can't open shared library '%s' (errno: %d %s) " 
func IsServerErrorCantOpenLibrary(err error) bool {
    result := Isa(err, ER_CANT_OPEN_LIBRARY)
    return result
}

    
// IsServerErrorCantFindDlEntry check mysql error is "Can't find symbol '%s' in library " 
func IsServerErrorCantFindDlEntry(err error) bool {
    result := Isa(err, ER_CANT_FIND_DL_ENTRY)
    return result
}

    
// IsServerErrorFunctionNotDefined check mysql error is "Function '%s' is not defined " 
func IsServerErrorFunctionNotDefined(err error) bool {
    result := Isa(err, ER_FUNCTION_NOT_DEFINED)
    return result
}

    
// IsServerErrorHostIsBlocked check mysql error is "Host '%s' is blocked because of many connection errors" 
func IsServerErrorHostIsBlocked(err error) bool {
    result := Isa(err, ER_HOST_IS_BLOCKED)
    return result
}

    
// IsServerErrorHostNotPrivileged check mysql error is "Host '%s' is not allowed to connect to this MySQL server " 
func IsServerErrorHostNotPrivileged(err error) bool {
    result := Isa(err, ER_HOST_NOT_PRIVILEGED)
    return result
}

    
// IsServerErrorPasswordAnonymousUser check mysql error is "You are using MySQL as an anonymous user and anonymous users are not allowed to change passwords " 
func IsServerErrorPasswordAnonymousUser(err error) bool {
    result := Isa(err, ER_PASSWORD_ANONYMOUS_USER)
    return result
}

    
// IsServerErrorPasswordNotAllowed check mysql error is "You must have privileges to update tables in the mysql database to be able to change passwords for others " 
func IsServerErrorPasswordNotAllowed(err error) bool {
    result := Isa(err, ER_PASSWORD_NOT_ALLOWED)
    return result
}

    
// IsServerErrorPasswordNoMatch check mysql error is "Can't find any matching row in the user table " 
func IsServerErrorPasswordNoMatch(err error) bool {
    result := Isa(err, ER_PASSWORD_NO_MATCH)
    return result
}

    
// IsServerErrorUpdateInfo check mysql error is "Rows matched: %ld Changed: %ld Warnings: %ld " 
func IsServerErrorUpdateInfo(err error) bool {
    result := Isa(err, ER_UPDATE_INFO)
    return result
}

    
// IsServerErrorCantCreateThread check mysql error is "Can't create a new thread (errno %d)" 
func IsServerErrorCantCreateThread(err error) bool {
    result := Isa(err, ER_CANT_CREATE_THREAD)
    return result
}

    
// IsServerErrorWrongValueCountOnRow check mysql error is "Column count doesn't match value count at row %ld " 
func IsServerErrorWrongValueCountOnRow(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_COUNT_ON_ROW)
    return result
}

    
// IsServerErrorCantReopenTable check mysql error is "Can't reopen table: '%s' " 
func IsServerErrorCantReopenTable(err error) bool {
    result := Isa(err, ER_CANT_REOPEN_TABLE)
    return result
}

    
// IsServerErrorInvalidUseOfNull check mysql error is "Invalid use of NULL value " 
func IsServerErrorInvalidUseOfNull(err error) bool {
    result := Isa(err, ER_INVALID_USE_OF_NULL)
    return result
}

    
// IsServerErrorRegexpError check mysql error is "Got error '%s' from regexp " 
func IsServerErrorRegexpError(err error) bool {
    result := Isa(err, ER_REGEXP_ERROR)
    return result
}

    
// IsServerErrorMixOfGroupFuncAndFields check mysql error is "Mixing of GROUP columns (MIN(),MAX(),COUNT(),...) with no GROUP columns is illegal if there is no GROUP BY clause " 
func IsServerErrorMixOfGroupFuncAndFields(err error) bool {
    result := Isa(err, ER_MIX_OF_GROUP_FUNC_AND_FIELDS)
    return result
}

    
// IsServerErrorNonexistingGrant check mysql error is "There is no such grant defined for user '%s' on host '%s' " 
func IsServerErrorNonexistingGrant(err error) bool {
    result := Isa(err, ER_NONEXISTING_GRANT)
    return result
}

    
// IsServerErrorTableaccessDeniedError check mysql error is "%s command denied to user '%s'@'%s' for table '%s' " 
func IsServerErrorTableaccessDeniedError(err error) bool {
    result := Isa(err, ER_TABLEACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorColumnaccessDeniedError check mysql error is "%s command denied to user '%s'@'%s' for column '%s' in table '%s' " 
func IsServerErrorColumnaccessDeniedError(err error) bool {
    result := Isa(err, ER_COLUMNACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorIllegalGrantForTable check mysql error is "Illegal GRANT/REVOKE command" 
func IsServerErrorIllegalGrantForTable(err error) bool {
    result := Isa(err, ER_ILLEGAL_GRANT_FOR_TABLE)
    return result
}

    
// IsServerErrorGrantWrongHostOrUser check mysql error is "The host or user argument to GRANT is too long " 
func IsServerErrorGrantWrongHostOrUser(err error) bool {
    result := Isa(err, ER_GRANT_WRONG_HOST_OR_USER)
    return result
}

    
// IsServerErrorNoSuchTable check mysql error is "Table '%s.%s' doesn't exist " 
func IsServerErrorNoSuchTable(err error) bool {
    result := Isa(err, ER_NO_SUCH_TABLE)
    return result
}

    
// IsServerErrorNonexistingTableGrant check mysql error is "There is no such grant defined for user '%s' on host '%s' on table '%s' " 
func IsServerErrorNonexistingTableGrant(err error) bool {
    result := Isa(err, ER_NONEXISTING_TABLE_GRANT)
    return result
}

    
// IsServerErrorNotAllowedCommand check mysql error is "The used command is not allowed with this MySQL version " 
func IsServerErrorNotAllowedCommand(err error) bool {
    result := Isa(err, ER_NOT_ALLOWED_COMMAND)
    return result
}

    
// IsServerErrorSyntaxError check mysql error is "You have an error in your SQL syntax" 
func IsServerErrorSyntaxError(err error) bool {
    result := Isa(err, ER_SYNTAX_ERROR)
    return result
}

    
// IsServerErrorUnused1 check mysql error is "Delayed insert thread couldn't get requested lock for table %s " 
func IsServerErrorUnused1(err error) bool {
    result := Isa(err, ER_UNUSED1)
    return result
}

    
// IsServerErrorUnused2 check mysql error is "Too many delayed threads in use " 
func IsServerErrorUnused2(err error) bool {
    result := Isa(err, ER_UNUSED2)
    return result
}

    
// IsServerErrorAbortingConnection check mysql error is "Aborted connection %ld to db: '%s' user: '%s' (%s) " 
func IsServerErrorAbortingConnection(err error) bool {
    result := Isa(err, ER_ABORTING_CONNECTION)
    return result
}

    
// IsServerErrorNetPacketTooLarge check mysql error is "Got a packet bigger than 'max_allowed_packet' bytes " 
func IsServerErrorNetPacketTooLarge(err error) bool {
    result := Isa(err, ER_NET_PACKET_TOO_LARGE)
    return result
}

    
// IsServerErrorNetReadErrorFromPipe check mysql error is "Got a read error from the connection pipe " 
func IsServerErrorNetReadErrorFromPipe(err error) bool {
    result := Isa(err, ER_NET_READ_ERROR_FROM_PIPE)
    return result
}

    
// IsServerErrorNetFcntlError check mysql error is "Got an error from fcntl() " 
func IsServerErrorNetFcntlError(err error) bool {
    result := Isa(err, ER_NET_FCNTL_ERROR)
    return result
}

    
// IsServerErrorNetPacketsOutOfOrder check mysql error is "Got packets out of order " 
func IsServerErrorNetPacketsOutOfOrder(err error) bool {
    result := Isa(err, ER_NET_PACKETS_OUT_OF_ORDER)
    return result
}

    
// IsServerErrorNetUncompressError check mysql error is "Couldn't uncompress communication packet " 
func IsServerErrorNetUncompressError(err error) bool {
    result := Isa(err, ER_NET_UNCOMPRESS_ERROR)
    return result
}

    
// IsServerErrorNetReadError check mysql error is "Got an error reading communication packets " 
func IsServerErrorNetReadError(err error) bool {
    result := Isa(err, ER_NET_READ_ERROR)
    return result
}

    
// IsServerErrorNetReadInterrupted check mysql error is "Got timeout reading communication packets " 
func IsServerErrorNetReadInterrupted(err error) bool {
    result := Isa(err, ER_NET_READ_INTERRUPTED)
    return result
}

    
// IsServerErrorNetErrorOnWrite check mysql error is "Got an error writing communication packets " 
func IsServerErrorNetErrorOnWrite(err error) bool {
    result := Isa(err, ER_NET_ERROR_ON_WRITE)
    return result
}

    
// IsServerErrorNetWriteInterrupted check mysql error is "Got timeout writing communication packets " 
func IsServerErrorNetWriteInterrupted(err error) bool {
    result := Isa(err, ER_NET_WRITE_INTERRUPTED)
    return result
}

    
// IsServerErrorTooLongString check mysql error is "Result string is longer than 'max_allowed_packet' bytes " 
func IsServerErrorTooLongString(err error) bool {
    result := Isa(err, ER_TOO_LONG_STRING)
    return result
}

    
// IsServerErrorTableCantHandleBlob check mysql error is "The used table type doesn't support BLOB/TEXT columns " 
func IsServerErrorTableCantHandleBlob(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_BLOB)
    return result
}

    
// IsServerErrorTableCantHandleAutoIncrement check mysql error is "The used table type doesn't support AUTO_INCREMENT columns " 
func IsServerErrorTableCantHandleAutoIncrement(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_AUTO_INCREMENT)
    return result
}

    
// IsServerErrorUnused3 check mysql error is "INSERT DELAYED can't be used with table '%s' because it is locked with LOCK TABLES " 
func IsServerErrorUnused3(err error) bool {
    result := Isa(err, ER_UNUSED3)
    return result
}

    
// IsServerErrorWrongColumnName check mysql error is "Incorrect column name '%s' " 
func IsServerErrorWrongColumnName(err error) bool {
    result := Isa(err, ER_WRONG_COLUMN_NAME)
    return result
}

    
// IsServerErrorWrongKeyColumn check mysql error is "The used storage engine can't index column '%s' " 
func IsServerErrorWrongKeyColumn(err error) bool {
    result := Isa(err, ER_WRONG_KEY_COLUMN)
    return result
}

    
// IsServerErrorWrongMrgTable check mysql error is "Unable to open underlying table which is differently defined or of non-MyISAM type or doesn't exist " 
func IsServerErrorWrongMrgTable(err error) bool {
    result := Isa(err, ER_WRONG_MRG_TABLE)
    return result
}

    
// IsServerErrorDupUnique check mysql error is "Can't write, because of unique constraint, to table '%s' " 
func IsServerErrorDupUnique(err error) bool {
    result := Isa(err, ER_DUP_UNIQUE)
    return result
}

    
// IsServerErrorBlobKeyWithoutLength check mysql error is "BLOB/TEXT column '%s' used in key specification without a key length " 
func IsServerErrorBlobKeyWithoutLength(err error) bool {
    result := Isa(err, ER_BLOB_KEY_WITHOUT_LENGTH)
    return result
}

    
// IsServerErrorPrimaryCantHaveNull check mysql error is "All parts of a PRIMARY KEY must be NOT NULL" 
func IsServerErrorPrimaryCantHaveNull(err error) bool {
    result := Isa(err, ER_PRIMARY_CANT_HAVE_NULL)
    return result
}

    
// IsServerErrorTooManyRows check mysql error is "Result consisted of more than one row " 
func IsServerErrorTooManyRows(err error) bool {
    result := Isa(err, ER_TOO_MANY_ROWS)
    return result
}

    
// IsServerErrorRequiresPrimaryKey check mysql error is "This table type requires a primary key " 
func IsServerErrorRequiresPrimaryKey(err error) bool {
    result := Isa(err, ER_REQUIRES_PRIMARY_KEY)
    return result
}

    
// IsServerErrorNoRaidCompiled check mysql error is "This version of MySQL is not compiled with RAID support " 
func IsServerErrorNoRaidCompiled(err error) bool {
    result := Isa(err, ER_NO_RAID_COMPILED)
    return result
}

    
// IsServerErrorUpdateWithoutKeyInSafeMode check mysql error is "You are using safe update mode and you tried to update a table without a WHERE that uses a KEY column. %s " 
func IsServerErrorUpdateWithoutKeyInSafeMode(err error) bool {
    result := Isa(err, ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE)
    return result
}

    
// IsServerErrorKeyDoesNotExits check mysql error is "Key '%s' doesn't exist in table '%s' " 
func IsServerErrorKeyDoesNotExits(err error) bool {
    result := Isa(err, ER_KEY_DOES_NOT_EXITS)
    return result
}

    
// IsServerErrorCheckNoSuchTable check mysql error is "Can't open table " 
func IsServerErrorCheckNoSuchTable(err error) bool {
    result := Isa(err, ER_CHECK_NO_SUCH_TABLE)
    return result
}

    
// IsServerErrorCheckNotImplemented check mysql error is "The storage engine for the table doesn't support %s " 
func IsServerErrorCheckNotImplemented(err error) bool {
    result := Isa(err, ER_CHECK_NOT_IMPLEMENTED)
    return result
}

    
// IsServerErrorCantDoThisDuringAnTransaction check mysql error is "You are not allowed to execute this command in a transaction " 
func IsServerErrorCantDoThisDuringAnTransaction(err error) bool {
    result := Isa(err, ER_CANT_DO_THIS_DURING_AN_TRANSACTION)
    return result
}

    
// IsServerErrorErrorDuringCommit check mysql error is "Got error %d during COMMIT " 
func IsServerErrorErrorDuringCommit(err error) bool {
    result := Isa(err, ER_ERROR_DURING_COMMIT)
    return result
}

    
// IsServerErrorErrorDuringRollback check mysql error is "Got error %d during ROLLBACK " 
func IsServerErrorErrorDuringRollback(err error) bool {
    result := Isa(err, ER_ERROR_DURING_ROLLBACK)
    return result
}

    
// IsServerErrorErrorDuringFlushLogs check mysql error is "Got error %d during FLUSH_LOGS " 
func IsServerErrorErrorDuringFlushLogs(err error) bool {
    result := Isa(err, ER_ERROR_DURING_FLUSH_LOGS)
    return result
}

    
// IsServerErrorErrorDuringCheckpoint check mysql error is "Got error %d during CHECKPOINT " 
func IsServerErrorErrorDuringCheckpoint(err error) bool {
    result := Isa(err, ER_ERROR_DURING_CHECKPOINT)
    return result
}

    
// IsServerErrorNewAbortingConnection check mysql error is "Aborted connection %u to db: '%s' user: '%s' host: '%s' (%s) " 
func IsServerErrorNewAbortingConnection(err error) bool {
    result := Isa(err, ER_NEW_ABORTING_CONNECTION)
    return result
}

    
// IsServerErrorDumpNotImplemented check mysql error is "The storage engine for the table does not support binary table dump " 
func IsServerErrorDumpNotImplemented(err error) bool {
    result := Isa(err, ER_DUMP_NOT_IMPLEMENTED)
    return result
}

    
// IsServerErrorFlushMasterBinlogClosed check mysql error is "Binlog closed, cannot RESET MASTER " 
func IsServerErrorFlushMasterBinlogClosed(err error) bool {
    result := Isa(err, ER_FLUSH_MASTER_BINLOG_CLOSED)
    return result
}

    
// IsServerErrorIndexRebuild check mysql error is "Failed rebuilding the index of dumped table '%s' " 
func IsServerErrorIndexRebuild(err error) bool {
    result := Isa(err, ER_INDEX_REBUILD)
    return result
}

    
// IsServerErrorMaster check mysql error is "Error from master: '%s' " 
func IsServerErrorMaster(err error) bool {
    result := Isa(err, ER_MASTER)
    return result
}

    
// IsServerErrorMasterNetRead check mysql error is "Net error reading from master " 
func IsServerErrorMasterNetRead(err error) bool {
    result := Isa(err, ER_MASTER_NET_READ)
    return result
}

    
// IsServerErrorMasterNetWrite check mysql error is "Net error writing to master " 
func IsServerErrorMasterNetWrite(err error) bool {
    result := Isa(err, ER_MASTER_NET_WRITE)
    return result
}

    
// IsServerErrorFtMatchingKeyNotFound check mysql error is "Can't find FULLTEXT index matching the column list " 
func IsServerErrorFtMatchingKeyNotFound(err error) bool {
    result := Isa(err, ER_FT_MATCHING_KEY_NOT_FOUND)
    return result
}

    
// IsServerErrorLockOrActiveTransaction check mysql error is "Can't execute the given command because you have active locked tables or an active transaction " 
func IsServerErrorLockOrActiveTransaction(err error) bool {
    result := Isa(err, ER_LOCK_OR_ACTIVE_TRANSACTION)
    return result
}

    
// IsServerErrorUnknownSystemVariable check mysql error is "Unknown system variable '%s' " 
func IsServerErrorUnknownSystemVariable(err error) bool {
    result := Isa(err, ER_UNKNOWN_SYSTEM_VARIABLE)
    return result
}

    
// IsServerErrorCrashedOnUsage check mysql error is "Table '%s' is marked as crashed and should be repaired " 
func IsServerErrorCrashedOnUsage(err error) bool {
    result := Isa(err, ER_CRASHED_ON_USAGE)
    return result
}

    
// IsServerErrorCrashedOnRepair check mysql error is "Table '%s' is marked as crashed and last (automatic?) repair failed " 
func IsServerErrorCrashedOnRepair(err error) bool {
    result := Isa(err, ER_CRASHED_ON_REPAIR)
    return result
}

    
// IsServerErrorWarningNotCompleteRollback check mysql error is "Some non-transactional changed tables couldn't be rolled back " 
func IsServerErrorWarningNotCompleteRollback(err error) bool {
    result := Isa(err, ER_WARNING_NOT_COMPLETE_ROLLBACK)
    return result
}

    
// IsServerErrorTransCacheFull check mysql error is "Multi-statement transaction required more than 'max_binlog_cache_size' bytes of storage" 
func IsServerErrorTransCacheFull(err error) bool {
    result := Isa(err, ER_TRANS_CACHE_FULL)
    return result
}

    
// IsServerErrorSlaveMustStop check mysql error is "This operation cannot be performed with a running slave" 
func IsServerErrorSlaveMustStop(err error) bool {
    result := Isa(err, ER_SLAVE_MUST_STOP)
    return result
}

    
// IsServerErrorSlaveNotRunning check mysql error is "This operation requires a running slave" 
func IsServerErrorSlaveNotRunning(err error) bool {
    result := Isa(err, ER_SLAVE_NOT_RUNNING)
    return result
}

    
// IsServerErrorBadSlave check mysql error is "The server is not configured as slave" 
func IsServerErrorBadSlave(err error) bool {
    result := Isa(err, ER_BAD_SLAVE)
    return result
}

    
// IsServerErrorMasterInfo check mysql error is "Could not initialize master info structure" 
func IsServerErrorMasterInfo(err error) bool {
    result := Isa(err, ER_MASTER_INFO)
    return result
}

    
// IsServerErrorSlaveThread check mysql error is "Could not create slave thread" 
func IsServerErrorSlaveThread(err error) bool {
    result := Isa(err, ER_SLAVE_THREAD)
    return result
}

    
// IsServerErrorTooManyUserConnections check mysql error is "User %s already has more than 'max_user_connections' active connections " 
func IsServerErrorTooManyUserConnections(err error) bool {
    result := Isa(err, ER_TOO_MANY_USER_CONNECTIONS)
    return result
}

    
// IsServerErrorSetConstantsOnly check mysql error is "You may only use constant expressions with SET " 
func IsServerErrorSetConstantsOnly(err error) bool {
    result := Isa(err, ER_SET_CONSTANTS_ONLY)
    return result
}

    
// IsServerErrorLockWaitTimeout check mysql error is "Lock wait timeout exceeded" 
func IsServerErrorLockWaitTimeout(err error) bool {
    result := Isa(err, ER_LOCK_WAIT_TIMEOUT)
    return result
}

    
// IsServerErrorLockTableFull check mysql error is "The total number of locks exceeds the lock table size" 
func IsServerErrorLockTableFull(err error) bool {
    result := Isa(err, ER_LOCK_TABLE_FULL)
    return result
}

    
// IsServerErrorReadOnlyTransaction check mysql error is "Update locks cannot be acquired during a READ UNCOMMITTED transaction " 
func IsServerErrorReadOnlyTransaction(err error) bool {
    result := Isa(err, ER_READ_ONLY_TRANSACTION)
    return result
}

    
// IsServerErrorDropDbWithReadLock check mysql error is "DROP DATABASE not allowed while thread is holding global read lock " 
func IsServerErrorDropDbWithReadLock(err error) bool {
    result := Isa(err, ER_DROP_DB_WITH_READ_LOCK)
    return result
}

    
// IsServerErrorCreateDbWithReadLock check mysql error is "CREATE DATABASE not allowed while thread is holding global read lock " 
func IsServerErrorCreateDbWithReadLock(err error) bool {
    result := Isa(err, ER_CREATE_DB_WITH_READ_LOCK)
    return result
}

    
// IsServerErrorWrongArguments check mysql error is "Incorrect arguments to %s " 
func IsServerErrorWrongArguments(err error) bool {
    result := Isa(err, ER_WRONG_ARGUMENTS)
    return result
}

    
// IsServerErrorNoPermissionToCreateUser check mysql error is "'%s'@'%s' is not allowed to create new users " 
func IsServerErrorNoPermissionToCreateUser(err error) bool {
    result := Isa(err, ER_NO_PERMISSION_TO_CREATE_USER)
    return result
}

    
// IsServerErrorUnionTablesInDifferentDir check mysql error is "Incorrect table definition" 
func IsServerErrorUnionTablesInDifferentDir(err error) bool {
    result := Isa(err, ER_UNION_TABLES_IN_DIFFERENT_DIR)
    return result
}

    
// IsServerErrorLockDeadlock check mysql error is "Deadlock found when trying to get lock" 
func IsServerErrorLockDeadlock(err error) bool {
    result := Isa(err, ER_LOCK_DEADLOCK)
    return result
}

    
// IsServerErrorTableCantHandleFt check mysql error is "The used table type doesn't support FULLTEXT indexes " 
func IsServerErrorTableCantHandleFt(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_FT)
    return result
}

    
// IsServerErrorCannotAddForeign check mysql error is "Cannot add foreign key constraint " 
func IsServerErrorCannotAddForeign(err error) bool {
    result := Isa(err, ER_CANNOT_ADD_FOREIGN)
    return result
}

    
// IsServerErrorNoReferencedRow check mysql error is "Cannot add or update a child row: a foreign key constraint fails" 
func IsServerErrorNoReferencedRow(err error) bool {
    result := Isa(err, ER_NO_REFERENCED_ROW)
    return result
}

    
// IsServerErrorRowIsReferenced check mysql error is "Cannot delete or update a parent row: a foreign key constraint fails" 
func IsServerErrorRowIsReferenced(err error) bool {
    result := Isa(err, ER_ROW_IS_REFERENCED)
    return result
}

    
// IsServerErrorConnectToMaster check mysql error is "Error connecting to master: %s " 
func IsServerErrorConnectToMaster(err error) bool {
    result := Isa(err, ER_CONNECT_TO_MASTER)
    return result
}

    
// IsServerErrorQueryOnMaster check mysql error is "Error running query on master: %s " 
func IsServerErrorQueryOnMaster(err error) bool {
    result := Isa(err, ER_QUERY_ON_MASTER)
    return result
}

    
// IsServerErrorErrorWhenExecutingCommand check mysql error is "Error when executing command %s: %s " 
func IsServerErrorErrorWhenExecutingCommand(err error) bool {
    result := Isa(err, ER_ERROR_WHEN_EXECUTING_COMMAND)
    return result
}

    
// IsServerErrorWrongUsage check mysql error is "Incorrect usage of %s and %s " 
func IsServerErrorWrongUsage(err error) bool {
    result := Isa(err, ER_WRONG_USAGE)
    return result
}

    
// IsServerErrorWrongNumberOfColumnsInSelect check mysql error is "The used SELECT statements have a different number of columns " 
func IsServerErrorWrongNumberOfColumnsInSelect(err error) bool {
    result := Isa(err, ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT)
    return result
}

    
// IsServerErrorCantUpdateWithReadlock check mysql error is "Can't execute the query because you have a conflicting read lock " 
func IsServerErrorCantUpdateWithReadlock(err error) bool {
    result := Isa(err, ER_CANT_UPDATE_WITH_READLOCK)
    return result
}

    
// IsServerErrorMixingNotAllowed check mysql error is "Mixing of transactional and non-transactional tables is disabled " 
func IsServerErrorMixingNotAllowed(err error) bool {
    result := Isa(err, ER_MIXING_NOT_ALLOWED)
    return result
}

    
// IsServerErrorDupArgument check mysql error is "Option '%s' used twice in statement " 
func IsServerErrorDupArgument(err error) bool {
    result := Isa(err, ER_DUP_ARGUMENT)
    return result
}

    
// IsServerErrorUserLimitReached check mysql error is "User '%s' has exceeded the '%s' resource (current value: %ld) " 
func IsServerErrorUserLimitReached(err error) bool {
    result := Isa(err, ER_USER_LIMIT_REACHED)
    return result
}

    
// IsServerErrorSpecificAccessDeniedError check mysql error is "Access denied" 
func IsServerErrorSpecificAccessDeniedError(err error) bool {
    result := Isa(err, ER_SPECIFIC_ACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorLocalVariable check mysql error is "Variable '%s' is a SESSION variable and can't be used with SET GLOBAL " 
func IsServerErrorLocalVariable(err error) bool {
    result := Isa(err, ER_LOCAL_VARIABLE)
    return result
}

    
// IsServerErrorGlobalVariable check mysql error is "Variable '%s' is a GLOBAL variable and should be set with SET GLOBAL " 
func IsServerErrorGlobalVariable(err error) bool {
    result := Isa(err, ER_GLOBAL_VARIABLE)
    return result
}

    
// IsServerErrorNoDefault check mysql error is "Variable '%s' doesn't have a default value " 
func IsServerErrorNoDefault(err error) bool {
    result := Isa(err, ER_NO_DEFAULT)
    return result
}

    
// IsServerErrorWrongValueForVar check mysql error is "Variable '%s' can't be set to the value of '%s' " 
func IsServerErrorWrongValueForVar(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_FOR_VAR)
    return result
}

    
// IsServerErrorWrongTypeForVar check mysql error is "Incorrect argument type to variable '%s' " 
func IsServerErrorWrongTypeForVar(err error) bool {
    result := Isa(err, ER_WRONG_TYPE_FOR_VAR)
    return result
}

    
// IsServerErrorVarCantBeRead check mysql error is "Variable '%s' can only be set, not read " 
func IsServerErrorVarCantBeRead(err error) bool {
    result := Isa(err, ER_VAR_CANT_BE_READ)
    return result
}

    
// IsServerErrorCantUseOptionHere check mysql error is "Incorrect usage/placement of '%s' " 
func IsServerErrorCantUseOptionHere(err error) bool {
    result := Isa(err, ER_CANT_USE_OPTION_HERE)
    return result
}

    
// IsServerErrorNotSupportedYet check mysql error is "This version of MySQL doesn't yet support '%s' " 
func IsServerErrorNotSupportedYet(err error) bool {
    result := Isa(err, ER_NOT_SUPPORTED_YET)
    return result
}

    
// IsServerErrorMasterFatalErrorReadingBinlog check mysql error is "Got fatal error %d from master when reading data from binary log: '%s' " 
func IsServerErrorMasterFatalErrorReadingBinlog(err error) bool {
    result := Isa(err, ER_MASTER_FATAL_ERROR_READING_BINLOG)
    return result
}

    
// IsServerErrorSlaveIgnoredTable check mysql error is "Slave SQL thread ignored the query because of replicate-*-table rules " 
func IsServerErrorSlaveIgnoredTable(err error) bool {
    result := Isa(err, ER_SLAVE_IGNORED_TABLE)
    return result
}

    
// IsServerErrorIncorrectGlobalLocalVar check mysql error is "Variable '%s' is a %s variable " 
func IsServerErrorIncorrectGlobalLocalVar(err error) bool {
    result := Isa(err, ER_INCORRECT_GLOBAL_LOCAL_VAR)
    return result
}

    
// IsServerErrorWrongFkDef check mysql error is "Incorrect foreign key definition for '%s': %s " 
func IsServerErrorWrongFkDef(err error) bool {
    result := Isa(err, ER_WRONG_FK_DEF)
    return result
}

    
// IsServerErrorKeyRefDoNotMatchTableRef check mysql error is "Key reference and table reference don't match " 
func IsServerErrorKeyRefDoNotMatchTableRef(err error) bool {
    result := Isa(err, ER_KEY_REF_DO_NOT_MATCH_TABLE_REF)
    return result
}

    
// IsServerErrorOperandColumns check mysql error is "Operand should contain %d column(s) " 
func IsServerErrorOperandColumns(err error) bool {
    result := Isa(err, ER_OPERAND_COLUMNS)
    return result
}

    
// IsServerErrorSubqueryNo1Row check mysql error is "Subquery returns more than 1 row " 
func IsServerErrorSubqueryNo1Row(err error) bool {
    result := Isa(err, ER_SUBQUERY_NO_1_ROW)
    return result
}

    
// IsServerErrorUnknownStmtHandler check mysql error is "Unknown prepared statement handler (%.*s) given to %s " 
func IsServerErrorUnknownStmtHandler(err error) bool {
    result := Isa(err, ER_UNKNOWN_STMT_HANDLER)
    return result
}

    
// IsServerErrorCorruptHelpDb check mysql error is "Help database is corrupt or does not exist " 
func IsServerErrorCorruptHelpDb(err error) bool {
    result := Isa(err, ER_CORRUPT_HELP_DB)
    return result
}

    
// IsServerErrorCyclicReference check mysql error is "Cyclic reference on subqueries " 
func IsServerErrorCyclicReference(err error) bool {
    result := Isa(err, ER_CYCLIC_REFERENCE)
    return result
}

    
// IsServerErrorAutoConvert check mysql error is "Converting column '%s' from %s to %s " 
func IsServerErrorAutoConvert(err error) bool {
    result := Isa(err, ER_AUTO_CONVERT)
    return result
}

    
// IsServerErrorIllegalReference check mysql error is "Reference '%s' not supported (%s) " 
func IsServerErrorIllegalReference(err error) bool {
    result := Isa(err, ER_ILLEGAL_REFERENCE)
    return result
}

    
// IsServerErrorDerivedMustHaveAlias check mysql error is "Every derived table must have its own alias " 
func IsServerErrorDerivedMustHaveAlias(err error) bool {
    result := Isa(err, ER_DERIVED_MUST_HAVE_ALIAS)
    return result
}

    
// IsServerErrorSelectReduced check mysql error is "Select %u was reduced during optimization " 
func IsServerErrorSelectReduced(err error) bool {
    result := Isa(err, ER_SELECT_REDUCED)
    return result
}

    
// IsServerErrorTablenameNotAllowedHere check mysql error is "Table '%s' from one of the SELECTs cannot be used in %s " 
func IsServerErrorTablenameNotAllowedHere(err error) bool {
    result := Isa(err, ER_TABLENAME_NOT_ALLOWED_HERE)
    return result
}

    
// IsServerErrorNotSupportedAuthMode check mysql error is "Client does not support authentication protocol requested by server" 
func IsServerErrorNotSupportedAuthMode(err error) bool {
    result := Isa(err, ER_NOT_SUPPORTED_AUTH_MODE)
    return result
}

    
// IsServerErrorSpatialCantHaveNull check mysql error is "All parts of a SPATIAL index must be NOT NULL " 
func IsServerErrorSpatialCantHaveNull(err error) bool {
    result := Isa(err, ER_SPATIAL_CANT_HAVE_NULL)
    return result
}

    
// IsServerErrorCollationCharsetMismatch check mysql error is "COLLATION '%s' is not valid for CHARACTER SET '%s' " 
func IsServerErrorCollationCharsetMismatch(err error) bool {
    result := Isa(err, ER_COLLATION_CHARSET_MISMATCH)
    return result
}

    
// IsServerErrorSlaveWasRunning check mysql error is "Slave is already running " 
func IsServerErrorSlaveWasRunning(err error) bool {
    result := Isa(err, ER_SLAVE_WAS_RUNNING)
    return result
}

    
// IsServerErrorSlaveWasNotRunning check mysql error is "Slave already has been stopped " 
func IsServerErrorSlaveWasNotRunning(err error) bool {
    result := Isa(err, ER_SLAVE_WAS_NOT_RUNNING)
    return result
}

    
// IsServerErrorTooBigForUncompress check mysql error is "Uncompressed data size too large" 
func IsServerErrorTooBigForUncompress(err error) bool {
    result := Isa(err, ER_TOO_BIG_FOR_UNCOMPRESS)
    return result
}

    
// IsServerErrorZlibZMemError check mysql error is "ZLIB: Not enough memory " 
func IsServerErrorZlibZMemError(err error) bool {
    result := Isa(err, ER_ZLIB_Z_MEM_ERROR)
    return result
}

    
// IsServerErrorZlibZBufError check mysql error is "ZLIB: Not enough room in the output buffer (probably, length of uncompressed data was corrupted) " 
func IsServerErrorZlibZBufError(err error) bool {
    result := Isa(err, ER_ZLIB_Z_BUF_ERROR)
    return result
}

    
// IsServerErrorZlibZDataError check mysql error is "ZLIB: Input data corrupted " 
func IsServerErrorZlibZDataError(err error) bool {
    result := Isa(err, ER_ZLIB_Z_DATA_ERROR)
    return result
}

    
// IsServerErrorCutValueGroupConcat check mysql error is "Row %u was cut by GROUP_CONCAT() " 
func IsServerErrorCutValueGroupConcat(err error) bool {
    result := Isa(err, ER_CUT_VALUE_GROUP_CONCAT)
    return result
}

    
// IsServerErrorWarnTooFewRecords check mysql error is "Row %ld doesn't contain data for all columns " 
func IsServerErrorWarnTooFewRecords(err error) bool {
    result := Isa(err, ER_WARN_TOO_FEW_RECORDS)
    return result
}

    
// IsServerErrorWarnTooManyRecords check mysql error is "Row %ld was truncated" 
func IsServerErrorWarnTooManyRecords(err error) bool {
    result := Isa(err, ER_WARN_TOO_MANY_RECORDS)
    return result
}

    
// IsServerErrorWarnNullToNotnull check mysql error is "Column set to default value" 
func IsServerErrorWarnNullToNotnull(err error) bool {
    result := Isa(err, ER_WARN_NULL_TO_NOTNULL)
    return result
}

    
// IsServerErrorWarnDataOutOfRange check mysql error is "Out of range value for column '%s' at row %ld " 
func IsServerErrorWarnDataOutOfRange(err error) bool {
    result := Isa(err, ER_WARN_DATA_OUT_OF_RANGE)
    return result
}

    
// IsWarnDataTruncated check mysql error is "Data truncated for column '%s' at row %ld " 
func IsWarnDataTruncated(err error) bool {
    result := Isa(err, WARN_DATA_TRUNCATED)
    return result
}

    
// IsServerErrorWarnUsingOtherHandler check mysql error is "Using storage engine %s for table '%s' " 
func IsServerErrorWarnUsingOtherHandler(err error) bool {
    result := Isa(err, ER_WARN_USING_OTHER_HANDLER)
    return result
}

    
// IsServerErrorCantAggregate2collations check mysql error is "Illegal mix of collations (%s,%s) and (%s,%s) for operation '%s' " 
func IsServerErrorCantAggregate2collations(err error) bool {
    result := Isa(err, ER_CANT_AGGREGATE_2COLLATIONS)
    return result
}

    
// IsServerErrorDropUser check mysql error is "Cannot drop one or more of the requested users " 
func IsServerErrorDropUser(err error) bool {
    result := Isa(err, ER_DROP_USER)
    return result
}

    
// IsServerErrorRevokeGrants check mysql error is "Can't revoke all privileges for one or more of the requested users " 
func IsServerErrorRevokeGrants(err error) bool {
    result := Isa(err, ER_REVOKE_GRANTS)
    return result
}

    
// IsServerErrorCantAggregate3collations check mysql error is "Illegal mix of collations (%s,%s), (%s,%s), (%s,%s) for operation '%s' " 
func IsServerErrorCantAggregate3collations(err error) bool {
    result := Isa(err, ER_CANT_AGGREGATE_3COLLATIONS)
    return result
}

    
// IsServerErrorCantAggregateNcollations check mysql error is "Illegal mix of collations for operation '%s' " 
func IsServerErrorCantAggregateNcollations(err error) bool {
    result := Isa(err, ER_CANT_AGGREGATE_NCOLLATIONS)
    return result
}

    
// IsServerErrorVariableIsNotStruct check mysql error is "Variable '%s' is not a variable component (can't be used as XXXX.variable_name) " 
func IsServerErrorVariableIsNotStruct(err error) bool {
    result := Isa(err, ER_VARIABLE_IS_NOT_STRUCT)
    return result
}

    
// IsServerErrorUnknownCollation check mysql error is "Unknown collation: '%s' " 
func IsServerErrorUnknownCollation(err error) bool {
    result := Isa(err, ER_UNKNOWN_COLLATION)
    return result
}

    
// IsServerErrorSlaveIgnoredSslParams check mysql error is "SSL parameters in CHANGE MASTER are ignored because this MySQL slave was compiled without SSL support" 
func IsServerErrorSlaveIgnoredSslParams(err error) bool {
    result := Isa(err, ER_SLAVE_IGNORED_SSL_PARAMS)
    return result
}

    
// IsServerErrorServerIsInSecureAuthMode check mysql error is "Server is running in --secure-auth mode, but '%s'@'%s' has a password in the old format" 
func IsServerErrorServerIsInSecureAuthMode(err error) bool {
    result := Isa(err, ER_SERVER_IS_IN_SECURE_AUTH_MODE)
    return result
}

    
// IsServerErrorWarnFieldResolved check mysql error is "Field or reference '%s%s%s%s%s' of SELECT #%d was resolved in SELECT #%d " 
func IsServerErrorWarnFieldResolved(err error) bool {
    result := Isa(err, ER_WARN_FIELD_RESOLVED)
    return result
}

    
// IsServerErrorBadSlaveUntilCond check mysql error is "Incorrect parameter or combination of parameters for START SLAVE UNTIL " 
func IsServerErrorBadSlaveUntilCond(err error) bool {
    result := Isa(err, ER_BAD_SLAVE_UNTIL_COND)
    return result
}

    
// IsServerErrorMissingSkipSlave check mysql error is "It is recommended to use --skip-slave-start when doing step-by-step replication with START SLAVE UNTIL" 
func IsServerErrorMissingSkipSlave(err error) bool {
    result := Isa(err, ER_MISSING_SKIP_SLAVE)
    return result
}

    
// IsServerErrorUntilCondIgnored check mysql error is "SQL thread is not to be started so UNTIL options are ignored " 
func IsServerErrorUntilCondIgnored(err error) bool {
    result := Isa(err, ER_UNTIL_COND_IGNORED)
    return result
}

    
// IsServerErrorWrongNameForIndex check mysql error is "Incorrect index name '%s' " 
func IsServerErrorWrongNameForIndex(err error) bool {
    result := Isa(err, ER_WRONG_NAME_FOR_INDEX)
    return result
}

    
// IsServerErrorWrongNameForCatalog check mysql error is "Incorrect catalog name '%s' " 
func IsServerErrorWrongNameForCatalog(err error) bool {
    result := Isa(err, ER_WRONG_NAME_FOR_CATALOG)
    return result
}

    
// IsServerErrorWarnQcResize check mysql error is "Query cache failed to set size %lu" 
func IsServerErrorWarnQcResize(err error) bool {
    result := Isa(err, ER_WARN_QC_RESIZE)
    return result
}

    
// IsServerErrorBadFtColumn check mysql error is "Column '%s' cannot be part of FULLTEXT index " 
func IsServerErrorBadFtColumn(err error) bool {
    result := Isa(err, ER_BAD_FT_COLUMN)
    return result
}

    
// IsServerErrorUnknownKeyCache check mysql error is "Unknown key cache '%s' " 
func IsServerErrorUnknownKeyCache(err error) bool {
    result := Isa(err, ER_UNKNOWN_KEY_CACHE)
    return result
}

    
// IsServerErrorWarnHostnameWontWork check mysql error is "MySQL is started in --skip-name-resolve mode" 
func IsServerErrorWarnHostnameWontWork(err error) bool {
    result := Isa(err, ER_WARN_HOSTNAME_WONT_WORK)
    return result
}

    
// IsServerErrorUnknownStorageEngine check mysql error is "Unknown storage engine '%s' " 
func IsServerErrorUnknownStorageEngine(err error) bool {
    result := Isa(err, ER_UNKNOWN_STORAGE_ENGINE)
    return result
}

    
// IsServerErrorWarnDeprecatedSyntax check mysql error is "'%s' is deprecated and will be removed in a future release. Please use %s instead " 
func IsServerErrorWarnDeprecatedSyntax(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYNTAX)
    return result
}

    
// IsServerErrorNonUpdatableTable check mysql error is "The target table %s of the %s is not updatable " 
func IsServerErrorNonUpdatableTable(err error) bool {
    result := Isa(err, ER_NON_UPDATABLE_TABLE)
    return result
}

    
// IsServerErrorFeatureDisabled check mysql error is "The '%s' feature is disabled" 
func IsServerErrorFeatureDisabled(err error) bool {
    result := Isa(err, ER_FEATURE_DISABLED)
    return result
}

    
// IsServerErrorOptionPreventsStatement check mysql error is "The MySQL server is running with the %s option so it cannot execute this statement " 
func IsServerErrorOptionPreventsStatement(err error) bool {
    result := Isa(err, ER_OPTION_PREVENTS_STATEMENT)
    return result
}

    
// IsServerErrorDuplicatedValueInType check mysql error is "Column '%s' has duplicated value '%s' in %s " 
func IsServerErrorDuplicatedValueInType(err error) bool {
    result := Isa(err, ER_DUPLICATED_VALUE_IN_TYPE)
    return result
}

    
// IsServerErrorTruncatedWrongValue check mysql error is "Truncated incorrect %s value: '%s' " 
func IsServerErrorTruncatedWrongValue(err error) bool {
    result := Isa(err, ER_TRUNCATED_WRONG_VALUE)
    return result
}

    
// IsServerErrorTooMuchAutoTimestampCols check mysql error is "Incorrect table definition" 
func IsServerErrorTooMuchAutoTimestampCols(err error) bool {
    result := Isa(err, ER_TOO_MUCH_AUTO_TIMESTAMP_COLS)
    return result
}

    
// IsServerErrorInvalidOnUpdate check mysql error is "Invalid ON UPDATE clause for '%s' column " 
func IsServerErrorInvalidOnUpdate(err error) bool {
    result := Isa(err, ER_INVALID_ON_UPDATE)
    return result
}

    
// IsServerErrorUnsupportedPs check mysql error is "This command is not supported in the prepared statement protocol yet " 
func IsServerErrorUnsupportedPs(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_PS)
    return result
}

    
// IsServerErrorGetErrmsg check mysql error is "Got error %d '%s' from %s " 
func IsServerErrorGetErrmsg(err error) bool {
    result := Isa(err, ER_GET_ERRMSG)
    return result
}

    
// IsServerErrorGetTemporaryErrmsg check mysql error is "Got temporary error %d '%s' from %s " 
func IsServerErrorGetTemporaryErrmsg(err error) bool {
    result := Isa(err, ER_GET_TEMPORARY_ERRMSG)
    return result
}

    
// IsServerErrorUnknownTimeZone check mysql error is "Unknown or incorrect time zone: '%s' " 
func IsServerErrorUnknownTimeZone(err error) bool {
    result := Isa(err, ER_UNKNOWN_TIME_ZONE)
    return result
}

    
// IsServerErrorWarnInvalidTimestamp check mysql error is "Invalid TIMESTAMP value in column '%s' at row %ld " 
func IsServerErrorWarnInvalidTimestamp(err error) bool {
    result := Isa(err, ER_WARN_INVALID_TIMESTAMP)
    return result
}

    
// IsServerErrorInvalidCharacterString check mysql error is "Invalid %s character string: '%s' " 
func IsServerErrorInvalidCharacterString(err error) bool {
    result := Isa(err, ER_INVALID_CHARACTER_STRING)
    return result
}

    
// IsServerErrorWarnAllowedPacketOverflowed check mysql error is "Result of %s() was larger than max_allowed_packet (%ld) - truncated " 
func IsServerErrorWarnAllowedPacketOverflowed(err error) bool {
    result := Isa(err, ER_WARN_ALLOWED_PACKET_OVERFLOWED)
    return result
}

    
// IsServerErrorConflictingDeclarations check mysql error is "Conflicting declarations: '%s%s' and '%s%s' " 
func IsServerErrorConflictingDeclarations(err error) bool {
    result := Isa(err, ER_CONFLICTING_DECLARATIONS)
    return result
}

    
// IsServerErrorSpNoRecursiveCreate check mysql error is "Can't create a %s from within another stored routine " 
func IsServerErrorSpNoRecursiveCreate(err error) bool {
    result := Isa(err, ER_SP_NO_RECURSIVE_CREATE)
    return result
}

    
// IsServerErrorSpAlreadyExists check mysql error is "%s %s already exists " 
func IsServerErrorSpAlreadyExists(err error) bool {
    result := Isa(err, ER_SP_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorSpDoesNotExist check mysql error is "%s %s does not exist " 
func IsServerErrorSpDoesNotExist(err error) bool {
    result := Isa(err, ER_SP_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorSpDropFailed check mysql error is "Failed to DROP %s %s " 
func IsServerErrorSpDropFailed(err error) bool {
    result := Isa(err, ER_SP_DROP_FAILED)
    return result
}

    
// IsServerErrorSpStoreFailed check mysql error is "Failed to CREATE %s %s " 
func IsServerErrorSpStoreFailed(err error) bool {
    result := Isa(err, ER_SP_STORE_FAILED)
    return result
}

    
// IsServerErrorSpLilabelMismatch check mysql error is "%s with no matching label: %s " 
func IsServerErrorSpLilabelMismatch(err error) bool {
    result := Isa(err, ER_SP_LILABEL_MISMATCH)
    return result
}

    
// IsServerErrorSpLabelRedefine check mysql error is "Redefining label %s " 
func IsServerErrorSpLabelRedefine(err error) bool {
    result := Isa(err, ER_SP_LABEL_REDEFINE)
    return result
}

    
// IsServerErrorSpLabelMismatch check mysql error is "End-label %s without match " 
func IsServerErrorSpLabelMismatch(err error) bool {
    result := Isa(err, ER_SP_LABEL_MISMATCH)
    return result
}

    
// IsServerErrorSpUninitVar check mysql error is "Referring to uninitialized variable %s " 
func IsServerErrorSpUninitVar(err error) bool {
    result := Isa(err, ER_SP_UNINIT_VAR)
    return result
}

    
// IsServerErrorSpBadselect check mysql error is "PROCEDURE %s can't return a result set in the given context " 
func IsServerErrorSpBadselect(err error) bool {
    result := Isa(err, ER_SP_BADSELECT)
    return result
}

    
// IsServerErrorSpBadreturn check mysql error is "RETURN is only allowed in a FUNCTION " 
func IsServerErrorSpBadreturn(err error) bool {
    result := Isa(err, ER_SP_BADRETURN)
    return result
}

    
// IsServerErrorSpBadstatement check mysql error is "%s is not allowed in stored procedures " 
func IsServerErrorSpBadstatement(err error) bool {
    result := Isa(err, ER_SP_BADSTATEMENT)
    return result
}

    
// IsServerErrorUpdateLogDeprecatedIgnored check mysql error is "The update log is deprecated and replaced by the binary log" 
func IsServerErrorUpdateLogDeprecatedIgnored(err error) bool {
    result := Isa(err, ER_UPDATE_LOG_DEPRECATED_IGNORED)
    return result
}

    
// IsServerErrorUpdateLogDeprecatedTranslated check mysql error is "The update log is deprecated and replaced by the binary log" 
func IsServerErrorUpdateLogDeprecatedTranslated(err error) bool {
    result := Isa(err, ER_UPDATE_LOG_DEPRECATED_TRANSLATED)
    return result
}

    
// IsServerErrorQueryInterrupted check mysql error is "Query execution was interrupted " 
func IsServerErrorQueryInterrupted(err error) bool {
    result := Isa(err, ER_QUERY_INTERRUPTED)
    return result
}

    
// IsServerErrorSpWrongNoOfArgs check mysql error is "Incorrect number of arguments for %s %s" 
func IsServerErrorSpWrongNoOfArgs(err error) bool {
    result := Isa(err, ER_SP_WRONG_NO_OF_ARGS)
    return result
}

    
// IsServerErrorSpCondMismatch check mysql error is "Undefined CONDITION: %s " 
func IsServerErrorSpCondMismatch(err error) bool {
    result := Isa(err, ER_SP_COND_MISMATCH)
    return result
}

    
// IsServerErrorSpNoreturn check mysql error is "No RETURN found in FUNCTION %s " 
func IsServerErrorSpNoreturn(err error) bool {
    result := Isa(err, ER_SP_NORETURN)
    return result
}

    
// IsServerErrorSpNoreturnend check mysql error is "FUNCTION %s ended without RETURN " 
func IsServerErrorSpNoreturnend(err error) bool {
    result := Isa(err, ER_SP_NORETURNEND)
    return result
}

    
// IsServerErrorSpBadCursorQuery check mysql error is "Cursor statement must be a SELECT " 
func IsServerErrorSpBadCursorQuery(err error) bool {
    result := Isa(err, ER_SP_BAD_CURSOR_QUERY)
    return result
}

    
// IsServerErrorSpBadCursorSelect check mysql error is "Cursor SELECT must not have INTO " 
func IsServerErrorSpBadCursorSelect(err error) bool {
    result := Isa(err, ER_SP_BAD_CURSOR_SELECT)
    return result
}

    
// IsServerErrorSpCursorMismatch check mysql error is "Undefined CURSOR: %s " 
func IsServerErrorSpCursorMismatch(err error) bool {
    result := Isa(err, ER_SP_CURSOR_MISMATCH)
    return result
}

    
// IsServerErrorSpCursorAlreadyOpen check mysql error is "Cursor is already open " 
func IsServerErrorSpCursorAlreadyOpen(err error) bool {
    result := Isa(err, ER_SP_CURSOR_ALREADY_OPEN)
    return result
}

    
// IsServerErrorSpCursorNotOpen check mysql error is "Cursor is not open " 
func IsServerErrorSpCursorNotOpen(err error) bool {
    result := Isa(err, ER_SP_CURSOR_NOT_OPEN)
    return result
}

    
// IsServerErrorSpUndeclaredVar check mysql error is "Undeclared variable: %s " 
func IsServerErrorSpUndeclaredVar(err error) bool {
    result := Isa(err, ER_SP_UNDECLARED_VAR)
    return result
}

    
// IsServerErrorSpWrongNoOfFetchArgs check mysql error is "Incorrect number of FETCH variables " 
func IsServerErrorSpWrongNoOfFetchArgs(err error) bool {
    result := Isa(err, ER_SP_WRONG_NO_OF_FETCH_ARGS)
    return result
}

    
// IsServerErrorSpFetchNoData check mysql error is "No data - zero rows fetched, selected, or processed " 
func IsServerErrorSpFetchNoData(err error) bool {
    result := Isa(err, ER_SP_FETCH_NO_DATA)
    return result
}

    
// IsServerErrorSpDupParam check mysql error is "Duplicate parameter: %s " 
func IsServerErrorSpDupParam(err error) bool {
    result := Isa(err, ER_SP_DUP_PARAM)
    return result
}

    
// IsServerErrorSpDupVar check mysql error is "Duplicate variable: %s " 
func IsServerErrorSpDupVar(err error) bool {
    result := Isa(err, ER_SP_DUP_VAR)
    return result
}

    
// IsServerErrorSpDupCond check mysql error is "Duplicate condition: %s " 
func IsServerErrorSpDupCond(err error) bool {
    result := Isa(err, ER_SP_DUP_COND)
    return result
}

    
// IsServerErrorSpDupCurs check mysql error is "Duplicate cursor: %s " 
func IsServerErrorSpDupCurs(err error) bool {
    result := Isa(err, ER_SP_DUP_CURS)
    return result
}

    
// IsServerErrorSpCantAlter check mysql error is "Failed to ALTER %s %s " 
func IsServerErrorSpCantAlter(err error) bool {
    result := Isa(err, ER_SP_CANT_ALTER)
    return result
}

    
// IsServerErrorSpSubselectNyi check mysql error is "Subquery value not supported " 
func IsServerErrorSpSubselectNyi(err error) bool {
    result := Isa(err, ER_SP_SUBSELECT_NYI)
    return result
}

    
// IsServerErrorStmtNotAllowedInSfOrTrg check mysql error is "%s is not allowed in stored function or trigger " 
func IsServerErrorStmtNotAllowedInSfOrTrg(err error) bool {
    result := Isa(err, ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG)
    return result
}

    
// IsServerErrorSpVarcondAfterCurshndlr check mysql error is "Variable or condition declaration after cursor or handler declaration " 
func IsServerErrorSpVarcondAfterCurshndlr(err error) bool {
    result := Isa(err, ER_SP_VARCOND_AFTER_CURSHNDLR)
    return result
}

    
// IsServerErrorSpCursorAfterHandler check mysql error is "Cursor declaration after handler declaration " 
func IsServerErrorSpCursorAfterHandler(err error) bool {
    result := Isa(err, ER_SP_CURSOR_AFTER_HANDLER)
    return result
}

    
// IsServerErrorSpCaseNotFound check mysql error is "Case not found for CASE statement " 
func IsServerErrorSpCaseNotFound(err error) bool {
    result := Isa(err, ER_SP_CASE_NOT_FOUND)
    return result
}

    
// IsServerErrorFparserTooBigFile check mysql error is "Configuration file '%s' is too big " 
func IsServerErrorFparserTooBigFile(err error) bool {
    result := Isa(err, ER_FPARSER_TOO_BIG_FILE)
    return result
}

    
// IsServerErrorFparserBadHeader check mysql error is "Malformed file type header in file '%s' " 
func IsServerErrorFparserBadHeader(err error) bool {
    result := Isa(err, ER_FPARSER_BAD_HEADER)
    return result
}

    
// IsServerErrorFparserEofInComment check mysql error is "Unexpected end of file while parsing comment '%s' " 
func IsServerErrorFparserEofInComment(err error) bool {
    result := Isa(err, ER_FPARSER_EOF_IN_COMMENT)
    return result
}

    
// IsServerErrorFparserErrorInParameter check mysql error is "Error while parsing parameter '%s' (line: '%s') " 
func IsServerErrorFparserErrorInParameter(err error) bool {
    result := Isa(err, ER_FPARSER_ERROR_IN_PARAMETER)
    return result
}

    
// IsServerErrorFparserEofInUnknownParameter check mysql error is "Unexpected end of file while skipping unknown parameter '%s' " 
func IsServerErrorFparserEofInUnknownParameter(err error) bool {
    result := Isa(err, ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER)
    return result
}

    
// IsServerErrorViewNoExplain check mysql error is "EXPLAIN/SHOW can not be issued" 
func IsServerErrorViewNoExplain(err error) bool {
    result := Isa(err, ER_VIEW_NO_EXPLAIN)
    return result
}

    
// IsServerErrorFrmUnknownType check mysql error is "File '%s' has unknown type '%s' in its header " 
func IsServerErrorFrmUnknownType(err error) bool {
    result := Isa(err, ER_FRM_UNKNOWN_TYPE)
    return result
}

    
// IsServerErrorWrongObject check mysql error is "'%s.%s' is not %s" 
func IsServerErrorWrongObject(err error) bool {
    result := Isa(err, ER_WRONG_OBJECT)
    return result
}

    
// IsServerErrorNonupdateableColumn check mysql error is "Column '%s' is not updatable " 
func IsServerErrorNonupdateableColumn(err error) bool {
    result := Isa(err, ER_NONUPDATEABLE_COLUMN)
    return result
}

    
// IsServerErrorViewSelectDerived check mysql error is "View's SELECT contains a subquery in the FROM clause" 
func IsServerErrorViewSelectDerived(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_DERIVED)
    return result
}

    
// IsServerErrorViewSelectDerivedUnused check mysql error is "View's SELECT contains a subquery in the FROM clause" 
func IsServerErrorViewSelectDerivedUnused(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_DERIVED_UNUSED)
    return result
}

    
// IsServerErrorViewSelectClause check mysql error is "View's SELECT contains a '%s' clause " 
func IsServerErrorViewSelectClause(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_CLAUSE)
    return result
}

    
// IsServerErrorViewSelectVariable check mysql error is "View's SELECT contains a variable or parameter " 
func IsServerErrorViewSelectVariable(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_VARIABLE)
    return result
}

    
// IsServerErrorViewSelectTmptable check mysql error is "View's SELECT refers to a temporary table '%s' " 
func IsServerErrorViewSelectTmptable(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_TMPTABLE)
    return result
}

    
// IsServerErrorViewWrongList check mysql error is "View's SELECT and view's field list have different column counts " 
func IsServerErrorViewWrongList(err error) bool {
    result := Isa(err, ER_VIEW_WRONG_LIST)
    return result
}

    
// IsServerErrorWarnViewMerge check mysql error is "View merge algorithm can't be used here for now (assumed undefined algorithm) " 
func IsServerErrorWarnViewMerge(err error) bool {
    result := Isa(err, ER_WARN_VIEW_MERGE)
    return result
}

    
// IsServerErrorWarnViewWithoutKey check mysql error is "View being updated does not have complete key of underlying table in it " 
func IsServerErrorWarnViewWithoutKey(err error) bool {
    result := Isa(err, ER_WARN_VIEW_WITHOUT_KEY)
    return result
}

    
// IsServerErrorViewInvalid check mysql error is "View '%s.%s' references invalid table(s) or column(s) or function(s) or definer/invoker of view lack rights to use them " 
func IsServerErrorViewInvalid(err error) bool {
    result := Isa(err, ER_VIEW_INVALID)
    return result
}

    
// IsServerErrorSpNoDropSp check mysql error is "Can't drop or alter a %s from within another stored routine " 
func IsServerErrorSpNoDropSp(err error) bool {
    result := Isa(err, ER_SP_NO_DROP_SP)
    return result
}

    
// IsServerErrorSpGotoInHndlr check mysql error is "GOTO is not allowed in a stored procedure handler " 
func IsServerErrorSpGotoInHndlr(err error) bool {
    result := Isa(err, ER_SP_GOTO_IN_HNDLR)
    return result
}

    
// IsServerErrorTrgAlreadyExists check mysql error is "Trigger already exists " 
func IsServerErrorTrgAlreadyExists(err error) bool {
    result := Isa(err, ER_TRG_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorTrgDoesNotExist check mysql error is "Trigger does not exist " 
func IsServerErrorTrgDoesNotExist(err error) bool {
    result := Isa(err, ER_TRG_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorTrgOnViewOrTempTable check mysql error is "Trigger's '%s' is view or temporary table " 
func IsServerErrorTrgOnViewOrTempTable(err error) bool {
    result := Isa(err, ER_TRG_ON_VIEW_OR_TEMP_TABLE)
    return result
}

    
// IsServerErrorTrgCantChangeRow check mysql error is "Updating of %s row is not allowed in %strigger " 
func IsServerErrorTrgCantChangeRow(err error) bool {
    result := Isa(err, ER_TRG_CANT_CHANGE_ROW)
    return result
}

    
// IsServerErrorTrgNoSuchRowInTrg check mysql error is "There is no %s row in %s trigger " 
func IsServerErrorTrgNoSuchRowInTrg(err error) bool {
    result := Isa(err, ER_TRG_NO_SUCH_ROW_IN_TRG)
    return result
}

    
// IsServerErrorNoDefaultForField check mysql error is "Field '%s' doesn't have a default value " 
func IsServerErrorNoDefaultForField(err error) bool {
    result := Isa(err, ER_NO_DEFAULT_FOR_FIELD)
    return result
}

    
// IsServerErrorDivisionByZero check mysql error is "Division by 0 " 
func IsServerErrorDivisionByZero(err error) bool {
    result := Isa(err, ER_DIVISION_BY_ZERO)
    return result
}

    
// IsServerErrorTruncatedWrongValueForField check mysql error is "Incorrect %s value: '%s' for column '%s' at row %ld " 
func IsServerErrorTruncatedWrongValueForField(err error) bool {
    result := Isa(err, ER_TRUNCATED_WRONG_VALUE_FOR_FIELD)
    return result
}

    
// IsServerErrorIllegalValueForType check mysql error is "Illegal %s '%s' value found during parsing " 
func IsServerErrorIllegalValueForType(err error) bool {
    result := Isa(err, ER_ILLEGAL_VALUE_FOR_TYPE)
    return result
}

    
// IsServerErrorViewNonupdCheck check mysql error is "CHECK OPTION on non-updatable view '%s.%s' " 
func IsServerErrorViewNonupdCheck(err error) bool {
    result := Isa(err, ER_VIEW_NONUPD_CHECK)
    return result
}

    
// IsServerErrorViewCheckFailed check mysql error is "CHECK OPTION failed '%s.%s' " 
func IsServerErrorViewCheckFailed(err error) bool {
    result := Isa(err, ER_VIEW_CHECK_FAILED)
    return result
}

    
// IsServerErrorProcaccessDeniedError check mysql error is "%s command denied to user '%s'@'%s' for routine '%s' " 
func IsServerErrorProcaccessDeniedError(err error) bool {
    result := Isa(err, ER_PROCACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorRelayLogFail check mysql error is "Failed purging old relay logs: %s " 
func IsServerErrorRelayLogFail(err error) bool {
    result := Isa(err, ER_RELAY_LOG_FAIL)
    return result
}

    
// IsServerErrorPasswdLength check mysql error is "Password hash should be a %d-digit hexadecimal number " 
func IsServerErrorPasswdLength(err error) bool {
    result := Isa(err, ER_PASSWD_LENGTH)
    return result
}

    
// IsServerErrorUnknownTargetBinlog check mysql error is "Target log not found in binlog index " 
func IsServerErrorUnknownTargetBinlog(err error) bool {
    result := Isa(err, ER_UNKNOWN_TARGET_BINLOG)
    return result
}

    
// IsServerErrorIoErrLogIndexRead check mysql error is "I/O error reading log index file " 
func IsServerErrorIoErrLogIndexRead(err error) bool {
    result := Isa(err, ER_IO_ERR_LOG_INDEX_READ)
    return result
}

    
// IsServerErrorBinlogPurgeProhibited check mysql error is "Server configuration does not permit binlog purge " 
func IsServerErrorBinlogPurgeProhibited(err error) bool {
    result := Isa(err, ER_BINLOG_PURGE_PROHIBITED)
    return result
}

    
// IsServerErrorFseekFail check mysql error is "Failed on fseek() " 
func IsServerErrorFseekFail(err error) bool {
    result := Isa(err, ER_FSEEK_FAIL)
    return result
}

    
// IsServerErrorBinlogPurgeFatalErr check mysql error is "Fatal error during log purge " 
func IsServerErrorBinlogPurgeFatalErr(err error) bool {
    result := Isa(err, ER_BINLOG_PURGE_FATAL_ERR)
    return result
}

    
// IsServerErrorLogInUse check mysql error is "A purgeable log is in use, will not purge " 
func IsServerErrorLogInUse(err error) bool {
    result := Isa(err, ER_LOG_IN_USE)
    return result
}

    
// IsServerErrorLogPurgeUnknownErr check mysql error is "Unknown error during log purge " 
func IsServerErrorLogPurgeUnknownErr(err error) bool {
    result := Isa(err, ER_LOG_PURGE_UNKNOWN_ERR)
    return result
}

    
// IsServerErrorRelayLogInit check mysql error is "Failed initializing relay log position: %s " 
func IsServerErrorRelayLogInit(err error) bool {
    result := Isa(err, ER_RELAY_LOG_INIT)
    return result
}

    
// IsServerErrorNoBinaryLogging check mysql error is "You are not using binary logging " 
func IsServerErrorNoBinaryLogging(err error) bool {
    result := Isa(err, ER_NO_BINARY_LOGGING)
    return result
}

    
// IsServerErrorReservedSyntax check mysql error is "The '%s' syntax is reserved for purposes internal to the MySQL server " 
func IsServerErrorReservedSyntax(err error) bool {
    result := Isa(err, ER_RESERVED_SYNTAX)
    return result
}

    
// IsServerErrorWsasFailed check mysql error is "WSAStartup Failed " 
func IsServerErrorWsasFailed(err error) bool {
    result := Isa(err, ER_WSAS_FAILED)
    return result
}

    
// IsServerErrorDiffGroupsProc check mysql error is "Can't handle procedures with different groups yet " 
func IsServerErrorDiffGroupsProc(err error) bool {
    result := Isa(err, ER_DIFF_GROUPS_PROC)
    return result
}

    
// IsServerErrorNoGroupForProc check mysql error is "Select must have a group with this procedure " 
func IsServerErrorNoGroupForProc(err error) bool {
    result := Isa(err, ER_NO_GROUP_FOR_PROC)
    return result
}

    
// IsServerErrorOrderWithProc check mysql error is "Can't use ORDER clause with this procedure " 
func IsServerErrorOrderWithProc(err error) bool {
    result := Isa(err, ER_ORDER_WITH_PROC)
    return result
}

    
// IsServerErrorLoggingProhibitChangingOf check mysql error is "Binary logging and replication forbid changing the global server %s " 
func IsServerErrorLoggingProhibitChangingOf(err error) bool {
    result := Isa(err, ER_LOGGING_PROHIBIT_CHANGING_OF)
    return result
}

    
// IsServerErrorNoFileMapping check mysql error is "Can't map file: %s, errno: %d " 
func IsServerErrorNoFileMapping(err error) bool {
    result := Isa(err, ER_NO_FILE_MAPPING)
    return result
}

    
// IsServerErrorWrongMagic check mysql error is "Wrong magic in %s " 
func IsServerErrorWrongMagic(err error) bool {
    result := Isa(err, ER_WRONG_MAGIC)
    return result
}

    
// IsServerErrorPsManyParam check mysql error is "Prepared statement contains too many placeholders " 
func IsServerErrorPsManyParam(err error) bool {
    result := Isa(err, ER_PS_MANY_PARAM)
    return result
}

    
// IsServerErrorKeyPart0 check mysql error is "Key part '%s' length cannot be 0 " 
func IsServerErrorKeyPart0(err error) bool {
    result := Isa(err, ER_KEY_PART_0)
    return result
}

    
// IsServerErrorViewChecksum check mysql error is "View text checksum failed " 
func IsServerErrorViewChecksum(err error) bool {
    result := Isa(err, ER_VIEW_CHECKSUM)
    return result
}

    
// IsServerErrorViewMultiupdate check mysql error is "Can not modify more than one base table through a join view '%s.%s' " 
func IsServerErrorViewMultiupdate(err error) bool {
    result := Isa(err, ER_VIEW_MULTIUPDATE)
    return result
}

    
// IsServerErrorViewNoInsertFieldList check mysql error is "Can not insert into join view '%s.%s' without fields list " 
func IsServerErrorViewNoInsertFieldList(err error) bool {
    result := Isa(err, ER_VIEW_NO_INSERT_FIELD_LIST)
    return result
}

    
// IsServerErrorViewDeleteMergeView check mysql error is "Can not delete from join view '%s.%s' " 
func IsServerErrorViewDeleteMergeView(err error) bool {
    result := Isa(err, ER_VIEW_DELETE_MERGE_VIEW)
    return result
}

    
// IsServerErrorCannotUser check mysql error is "Operation %s failed for %s " 
func IsServerErrorCannotUser(err error) bool {
    result := Isa(err, ER_CANNOT_USER)
    return result
}

    
// IsServerErrorXaerNota check mysql error is "XAER_NOTA: Unknown XID " 
func IsServerErrorXaerNota(err error) bool {
    result := Isa(err, ER_XAER_NOTA)
    return result
}

    
// IsServerErrorXaerInval check mysql error is "XAER_INVAL: Invalid arguments (or unsupported command) " 
func IsServerErrorXaerInval(err error) bool {
    result := Isa(err, ER_XAER_INVAL)
    return result
}

    
// IsServerErrorXaerRmfail check mysql error is "XAER_RMFAIL: The command cannot be executed when global transaction is in the %s state " 
func IsServerErrorXaerRmfail(err error) bool {
    result := Isa(err, ER_XAER_RMFAIL)
    return result
}

    
// IsServerErrorXaerOutside check mysql error is "XAER_OUTSIDE: Some work is done outside global transaction " 
func IsServerErrorXaerOutside(err error) bool {
    result := Isa(err, ER_XAER_OUTSIDE)
    return result
}

    
// IsServerErrorXaerRmerr check mysql error is "XAER_RMERR: Fatal error occurred in the transaction branch - check your data for consistency " 
func IsServerErrorXaerRmerr(err error) bool {
    result := Isa(err, ER_XAER_RMERR)
    return result
}

    
// IsServerErrorXaRbrollback check mysql error is "XA_RBROLLBACK: Transaction branch was rolled back " 
func IsServerErrorXaRbrollback(err error) bool {
    result := Isa(err, ER_XA_RBROLLBACK)
    return result
}

    
// IsServerErrorNonexistingProcGrant check mysql error is "There is no such grant defined for user '%s' on host '%s' on routine '%s' " 
func IsServerErrorNonexistingProcGrant(err error) bool {
    result := Isa(err, ER_NONEXISTING_PROC_GRANT)
    return result
}

    
// IsServerErrorProcAutoGrantFail check mysql error is "Failed to grant EXECUTE and ALTER ROUTINE privileges " 
func IsServerErrorProcAutoGrantFail(err error) bool {
    result := Isa(err, ER_PROC_AUTO_GRANT_FAIL)
    return result
}

    
// IsServerErrorProcAutoRevokeFail check mysql error is "Failed to revoke all privileges to dropped routine " 
func IsServerErrorProcAutoRevokeFail(err error) bool {
    result := Isa(err, ER_PROC_AUTO_REVOKE_FAIL)
    return result
}

    
// IsServerErrorDataTooLong check mysql error is "Data too long for column '%s' at row %ld " 
func IsServerErrorDataTooLong(err error) bool {
    result := Isa(err, ER_DATA_TOO_LONG)
    return result
}

    
// IsServerErrorSpBadSqlstate check mysql error is "Bad SQLSTATE: '%s' " 
func IsServerErrorSpBadSqlstate(err error) bool {
    result := Isa(err, ER_SP_BAD_SQLSTATE)
    return result
}

    
// IsServerErrorStartup check mysql error is "%s: ready for connections. Version: '%s' socket: '%s' port: %d %s " 
func IsServerErrorStartup(err error) bool {
    result := Isa(err, ER_STARTUP)
    return result
}

    
// IsServerErrorLoadFromFixedSizeRowsToVar check mysql error is "Can't load value from file with fixed size rows to variable " 
func IsServerErrorLoadFromFixedSizeRowsToVar(err error) bool {
    result := Isa(err, ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR)
    return result
}

    
// IsServerErrorCantCreateUserWithGrant check mysql error is "You are not allowed to create a user with GRANT " 
func IsServerErrorCantCreateUserWithGrant(err error) bool {
    result := Isa(err, ER_CANT_CREATE_USER_WITH_GRANT)
    return result
}

    
// IsServerErrorWrongValueForType check mysql error is "Incorrect %s value: '%s' for function %s " 
func IsServerErrorWrongValueForType(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_FOR_TYPE)
    return result
}

    
// IsServerErrorTableDefChanged check mysql error is "Table definition has changed, please retry transaction " 
func IsServerErrorTableDefChanged(err error) bool {
    result := Isa(err, ER_TABLE_DEF_CHANGED)
    return result
}

    
// IsServerErrorSpDupHandler check mysql error is "Duplicate handler declared in the same block " 
func IsServerErrorSpDupHandler(err error) bool {
    result := Isa(err, ER_SP_DUP_HANDLER)
    return result
}

    
// IsServerErrorSpNotVarArg check mysql error is "OUT or INOUT argument %d for routine %s is not a variable or NEW pseudo-variable in BEFORE trigger " 
func IsServerErrorSpNotVarArg(err error) bool {
    result := Isa(err, ER_SP_NOT_VAR_ARG)
    return result
}

    
// IsServerErrorSpNoRetset check mysql error is "Not allowed to return a result set from a %s " 
func IsServerErrorSpNoRetset(err error) bool {
    result := Isa(err, ER_SP_NO_RETSET)
    return result
}

    
// IsServerErrorCantCreateGeometryObject check mysql error is "Cannot get geometry object from data you send to the GEOMETRY field " 
func IsServerErrorCantCreateGeometryObject(err error) bool {
    result := Isa(err, ER_CANT_CREATE_GEOMETRY_OBJECT)
    return result
}

    
// IsServerErrorFailedRoutineBreakBinlog check mysql error is "A routine failed and has neither NO SQL nor READS SQL DATA in its declaration and binary logging is enabled" 
func IsServerErrorFailedRoutineBreakBinlog(err error) bool {
    result := Isa(err, ER_FAILED_ROUTINE_BREAK_BINLOG)
    return result
}

    
// IsServerErrorBinlogUnsafeRoutine check mysql error is "This function has none of DETERMINISTIC, NO SQL, or READS SQL DATA in its declaration and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable) " 
func IsServerErrorBinlogUnsafeRoutine(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_ROUTINE)
    return result
}

    
// IsServerErrorBinlogCreateRoutineNeedSuper check mysql error is "You do not have the SUPER privilege and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable) " 
func IsServerErrorBinlogCreateRoutineNeedSuper(err error) bool {
    result := Isa(err, ER_BINLOG_CREATE_ROUTINE_NEED_SUPER)
    return result
}

    
// IsServerErrorExecStmtWithOpenCursor check mysql error is "You can't execute a prepared statement which has an open cursor associated with it. Reset the statement to re-execute it. " 
func IsServerErrorExecStmtWithOpenCursor(err error) bool {
    result := Isa(err, ER_EXEC_STMT_WITH_OPEN_CURSOR)
    return result
}

    
// IsServerErrorStmtHasNoOpenCursor check mysql error is "The statement (%lu) has no open cursor. " 
func IsServerErrorStmtHasNoOpenCursor(err error) bool {
    result := Isa(err, ER_STMT_HAS_NO_OPEN_CURSOR)
    return result
}

    
// IsServerErrorCommitNotAllowedInSfOrTrg check mysql error is "Explicit or implicit commit is not allowed in stored function or trigger. " 
func IsServerErrorCommitNotAllowedInSfOrTrg(err error) bool {
    result := Isa(err, ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG)
    return result
}

    
// IsServerErrorNoDefaultForViewField check mysql error is "Field of view '%s.%s' underlying table doesn't have a default value " 
func IsServerErrorNoDefaultForViewField(err error) bool {
    result := Isa(err, ER_NO_DEFAULT_FOR_VIEW_FIELD)
    return result
}

    
// IsServerErrorSpNoRecursion check mysql error is "Recursive stored functions and triggers are not allowed. " 
func IsServerErrorSpNoRecursion(err error) bool {
    result := Isa(err, ER_SP_NO_RECURSION)
    return result
}

    
// IsServerErrorTooBigScale check mysql error is "Too big scale %d specified for column '%s'. Maximum is %lu. " 
func IsServerErrorTooBigScale(err error) bool {
    result := Isa(err, ER_TOO_BIG_SCALE)
    return result
}

    
// IsServerErrorTooBigPrecision check mysql error is "Too-big precision %d specified for '%s'. Maximum is %lu. " 
func IsServerErrorTooBigPrecision(err error) bool {
    result := Isa(err, ER_TOO_BIG_PRECISION)
    return result
}

    
// IsServerErrorMBiggerThanD check mysql error is "For float(M,D), double(M,D) or decimal(M,D), M must be >= D (column '%s'). " 
func IsServerErrorMBiggerThanD(err error) bool {
    result := Isa(err, ER_M_BIGGER_THAN_D)
    return result
}

    
// IsServerErrorWrongLockOfSystemTable check mysql error is "You can't combine write-locking of system tables with other tables or lock types " 
func IsServerErrorWrongLockOfSystemTable(err error) bool {
    result := Isa(err, ER_WRONG_LOCK_OF_SYSTEM_TABLE)
    return result
}

    
// IsServerErrorConnectToForeignDataSource check mysql error is "Unable to connect to foreign data source: %s " 
func IsServerErrorConnectToForeignDataSource(err error) bool {
    result := Isa(err, ER_CONNECT_TO_FOREIGN_DATA_SOURCE)
    return result
}

    
// IsServerErrorQueryOnForeignDataSource check mysql error is "There was a problem processing the query on the foreign data source. Data source error: %s " 
func IsServerErrorQueryOnForeignDataSource(err error) bool {
    result := Isa(err, ER_QUERY_ON_FOREIGN_DATA_SOURCE)
    return result
}

    
// IsServerErrorForeignDataSourceDoesntExist check mysql error is "The foreign data source you are trying to reference does not exist. Data source error: %s " 
func IsServerErrorForeignDataSourceDoesntExist(err error) bool {
    result := Isa(err, ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST)
    return result
}

    
// IsServerErrorForeignDataStringInvalidCantCreate check mysql error is "Can't create federated table. The data source connection string '%s' is not in the correct format " 
func IsServerErrorForeignDataStringInvalidCantCreate(err error) bool {
    result := Isa(err, ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE)
    return result
}

    
// IsServerErrorForeignDataStringInvalid check mysql error is "The data source connection string '%s' is not in the correct format " 
func IsServerErrorForeignDataStringInvalid(err error) bool {
    result := Isa(err, ER_FOREIGN_DATA_STRING_INVALID)
    return result
}

    
// IsServerErrorCantCreateFederatedTable check mysql error is "Can't create federated table. Foreign data src error: %s " 
func IsServerErrorCantCreateFederatedTable(err error) bool {
    result := Isa(err, ER_CANT_CREATE_FEDERATED_TABLE)
    return result
}

    
// IsServerErrorTrgInWrongSchema check mysql error is "Trigger in wrong schema " 
func IsServerErrorTrgInWrongSchema(err error) bool {
    result := Isa(err, ER_TRG_IN_WRONG_SCHEMA)
    return result
}

    
// IsServerErrorStackOverrunNeedMore check mysql error is "Thread stack overrun: %ld bytes used of a %ld byte stack, and %ld bytes needed. Use 'mysqld --thread_stack=#' to specify a bigger stack. " 
func IsServerErrorStackOverrunNeedMore(err error) bool {
    result := Isa(err, ER_STACK_OVERRUN_NEED_MORE)
    return result
}

    
// IsServerErrorTooLongBody check mysql error is "Routine body for '%s' is too long " 
func IsServerErrorTooLongBody(err error) bool {
    result := Isa(err, ER_TOO_LONG_BODY)
    return result
}

    
// IsServerErrorWarnCantDropDefaultKeycache check mysql error is "Cannot drop default keycache " 
func IsServerErrorWarnCantDropDefaultKeycache(err error) bool {
    result := Isa(err, ER_WARN_CANT_DROP_DEFAULT_KEYCACHE)
    return result
}

    
// IsServerErrorTooBigDisplaywidth check mysql error is "Display width out of range for column '%s' (max = %lu) " 
func IsServerErrorTooBigDisplaywidth(err error) bool {
    result := Isa(err, ER_TOO_BIG_DISPLAYWIDTH)
    return result
}

    
// IsServerErrorXaerDupid check mysql error is "XAER_DUPID: The XID already exists " 
func IsServerErrorXaerDupid(err error) bool {
    result := Isa(err, ER_XAER_DUPID)
    return result
}

    
// IsServerErrorDatetimeFunctionOverflow check mysql error is "Datetime function: %s field overflow " 
func IsServerErrorDatetimeFunctionOverflow(err error) bool {
    result := Isa(err, ER_DATETIME_FUNCTION_OVERFLOW)
    return result
}

    
// IsServerErrorCantUpdateUsedTableInSfOrTrg check mysql error is "Can't update table '%s' in stored function/trigger because it is already used by statement which invoked this stored function/trigger. " 
func IsServerErrorCantUpdateUsedTableInSfOrTrg(err error) bool {
    result := Isa(err, ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG)
    return result
}

    
// IsServerErrorViewPreventUpdate check mysql error is "The definition of table '%s' prevents operation %s on table '%s'. " 
func IsServerErrorViewPreventUpdate(err error) bool {
    result := Isa(err, ER_VIEW_PREVENT_UPDATE)
    return result
}

    
// IsServerErrorPsNoRecursion check mysql error is "The prepared statement contains a stored routine call that refers to that same statement. It's not allowed to execute a prepared statement in such a recursive manner " 
func IsServerErrorPsNoRecursion(err error) bool {
    result := Isa(err, ER_PS_NO_RECURSION)
    return result
}

    
// IsServerErrorSpCantSetAutocommit check mysql error is "Not allowed to set autocommit from a stored function or trigger " 
func IsServerErrorSpCantSetAutocommit(err error) bool {
    result := Isa(err, ER_SP_CANT_SET_AUTOCOMMIT)
    return result
}

    
// IsServerErrorMalformedDefiner check mysql error is "Definer is not fully qualified " 
func IsServerErrorMalformedDefiner(err error) bool {
    result := Isa(err, ER_MALFORMED_DEFINER)
    return result
}

    
// IsServerErrorViewFrmNoUser check mysql error is "View '%s'.'%s' has no definer information (old table format). Current user is used as definer. Please recreate the view! " 
func IsServerErrorViewFrmNoUser(err error) bool {
    result := Isa(err, ER_VIEW_FRM_NO_USER)
    return result
}

    
// IsServerErrorViewOtherUser check mysql error is "You need the SUPER privilege for creation view with '%s'@'%s' definer " 
func IsServerErrorViewOtherUser(err error) bool {
    result := Isa(err, ER_VIEW_OTHER_USER)
    return result
}

    
// IsServerErrorNoSuchUser check mysql error is "The user specified as a definer ('%s'@'%s') does not exist " 
func IsServerErrorNoSuchUser(err error) bool {
    result := Isa(err, ER_NO_SUCH_USER)
    return result
}

    
// IsServerErrorForbidSchemaChange check mysql error is "Changing schema from '%s' to '%s' is not allowed. " 
func IsServerErrorForbidSchemaChange(err error) bool {
    result := Isa(err, ER_FORBID_SCHEMA_CHANGE)
    return result
}

    
// IsServerErrorRowIsReferenced2 check mysql error is "Cannot delete or update a parent row: a foreign key constraint fails (%s)" 
func IsServerErrorRowIsReferenced2(err error) bool {
    result := Isa(err, ER_ROW_IS_REFERENCED_2)
    return result
}

    
// IsServerErrorNoReferencedRow2 check mysql error is "Cannot add or update a child row: a foreign key constraint fails (%s)" 
func IsServerErrorNoReferencedRow2(err error) bool {
    result := Isa(err, ER_NO_REFERENCED_ROW_2)
    return result
}

    
// IsServerErrorSpBadVarShadow check mysql error is "Variable '%s' must be quoted with `...`, or renamed " 
func IsServerErrorSpBadVarShadow(err error) bool {
    result := Isa(err, ER_SP_BAD_VAR_SHADOW)
    return result
}

    
// IsServerErrorTrgNoDefiner check mysql error is "No definer attribute for trigger '%s'.'%s'. The trigger will be activated under the authorization of the invoker, which may have insufficient privileges. Please recreate the trigger. " 
func IsServerErrorTrgNoDefiner(err error) bool {
    result := Isa(err, ER_TRG_NO_DEFINER)
    return result
}

    
// IsServerErrorOldFileFormat check mysql error is "'%s' has an old format, you should re-create the '%s' object(s) " 
func IsServerErrorOldFileFormat(err error) bool {
    result := Isa(err, ER_OLD_FILE_FORMAT)
    return result
}

    
// IsServerErrorSpRecursionLimit check mysql error is "Recursive limit %d (as set by the max_sp_recursion_depth variable) was exceeded for routine %s " 
func IsServerErrorSpRecursionLimit(err error) bool {
    result := Isa(err, ER_SP_RECURSION_LIMIT)
    return result
}

    
// IsServerErrorSpProcTableCorrupt check mysql error is "Failed to load routine %s. The table mysql.proc is missing, corrupt, or contains bad data (internal code %d) " 
func IsServerErrorSpProcTableCorrupt(err error) bool {
    result := Isa(err, ER_SP_PROC_TABLE_CORRUPT)
    return result
}

    
// IsServerErrorSpWrongName check mysql error is "Incorrect routine name '%s' " 
func IsServerErrorSpWrongName(err error) bool {
    result := Isa(err, ER_SP_WRONG_NAME)
    return result
}

    
// IsServerErrorTableNeedsUpgrade check mysql error is "Table upgrade required. Please do "REPAIR TABLE `%s`" or dump/reload to fix it! " 
func IsServerErrorTableNeedsUpgrade(err error) bool {
    result := Isa(err, ER_TABLE_NEEDS_UPGRADE)
    return result
}

    
// IsServerErrorSpNoAggregate check mysql error is "AGGREGATE is not supported for stored functions " 
func IsServerErrorSpNoAggregate(err error) bool {
    result := Isa(err, ER_SP_NO_AGGREGATE)
    return result
}

    
// IsServerErrorMaxPreparedStmtCountReached check mysql error is "Can't create more than max_prepared_stmt_count statements (current value: %lu) " 
func IsServerErrorMaxPreparedStmtCountReached(err error) bool {
    result := Isa(err, ER_MAX_PREPARED_STMT_COUNT_REACHED)
    return result
}

    
// IsServerErrorViewRecursive check mysql error is "`%s`.`%s` contains view recursion " 
func IsServerErrorViewRecursive(err error) bool {
    result := Isa(err, ER_VIEW_RECURSIVE)
    return result
}

    
// IsServerErrorNonGroupingFieldUsed check mysql error is "Non-grouping field '%s' is used in %s clause " 
func IsServerErrorNonGroupingFieldUsed(err error) bool {
    result := Isa(err, ER_NON_GROUPING_FIELD_USED)
    return result
}

    
// IsServerErrorTableCantHandleSpkeys check mysql error is "The used table type doesn't support SPATIAL indexes " 
func IsServerErrorTableCantHandleSpkeys(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_SPKEYS)
    return result
}

    
// IsServerErrorNoTriggersOnSystemSchema check mysql error is "Triggers can not be created on system tables " 
func IsServerErrorNoTriggersOnSystemSchema(err error) bool {
    result := Isa(err, ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA)
    return result
}

    
// IsServerErrorRemovedSpaces check mysql error is "Leading spaces are removed from name '%s' " 
func IsServerErrorRemovedSpaces(err error) bool {
    result := Isa(err, ER_REMOVED_SPACES)
    return result
}

    
// IsServerErrorAutoincReadFailed check mysql error is "Failed to read auto-increment value from storage engine " 
func IsServerErrorAutoincReadFailed(err error) bool {
    result := Isa(err, ER_AUTOINC_READ_FAILED)
    return result
}

    
// IsServerErrorUsername check mysql error is "user name " 
func IsServerErrorUsername(err error) bool {
    result := Isa(err, ER_USERNAME)
    return result
}

    
// IsServerErrorHostname check mysql error is "host name " 
func IsServerErrorHostname(err error) bool {
    result := Isa(err, ER_HOSTNAME)
    return result
}

    
// IsServerErrorWrongStringLength check mysql error is "String '%s' is too long for %s (should be no longer than %d) " 
func IsServerErrorWrongStringLength(err error) bool {
    result := Isa(err, ER_WRONG_STRING_LENGTH)
    return result
}

    
// IsServerErrorNonInsertableTable check mysql error is "The target table %s of the %s is not insertable-into " 
func IsServerErrorNonInsertableTable(err error) bool {
    result := Isa(err, ER_NON_INSERTABLE_TABLE)
    return result
}

    
// IsServerErrorAdminWrongMrgTable check mysql error is "Table '%s' is differently defined or of non-MyISAM type or doesn't exist " 
func IsServerErrorAdminWrongMrgTable(err error) bool {
    result := Isa(err, ER_ADMIN_WRONG_MRG_TABLE)
    return result
}

    
// IsServerErrorTooHighLevelOfNestingForSelect check mysql error is "Too high level of nesting for select " 
func IsServerErrorTooHighLevelOfNestingForSelect(err error) bool {
    result := Isa(err, ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT)
    return result
}

    
// IsServerErrorNameBecomesEmpty check mysql error is "Name '%s' has become '' " 
func IsServerErrorNameBecomesEmpty(err error) bool {
    result := Isa(err, ER_NAME_BECOMES_EMPTY)
    return result
}

    
// IsServerErrorAmbiguousFieldTerm check mysql error is "First character of the FIELDS TERMINATED string is ambiguous" 
func IsServerErrorAmbiguousFieldTerm(err error) bool {
    result := Isa(err, ER_AMBIGUOUS_FIELD_TERM)
    return result
}

    
// IsServerErrorForeignServerExists check mysql error is "The foreign server, %s, you are trying to create already exists. " 
func IsServerErrorForeignServerExists(err error) bool {
    result := Isa(err, ER_FOREIGN_SERVER_EXISTS)
    return result
}

    
// IsServerErrorForeignServerDoesntExist check mysql error is "The foreign server name you are trying to reference does not exist. Data source error: %s " 
func IsServerErrorForeignServerDoesntExist(err error) bool {
    result := Isa(err, ER_FOREIGN_SERVER_DOESNT_EXIST)
    return result
}

    
// IsServerErrorIllegalHaCreateOption check mysql error is "Table storage engine '%s' does not support the create option '%s' " 
func IsServerErrorIllegalHaCreateOption(err error) bool {
    result := Isa(err, ER_ILLEGAL_HA_CREATE_OPTION)
    return result
}

    
// IsServerErrorPartitionRequiresValuesError check mysql error is "Syntax error: %s PARTITIONING requires definition of VALUES %s for each partition " 
func IsServerErrorPartitionRequiresValuesError(err error) bool {
    result := Isa(err, ER_PARTITION_REQUIRES_VALUES_ERROR)
    return result
}

    
// IsServerErrorPartitionWrongValuesError check mysql error is "Only %s PARTITIONING can use VALUES %s in partition definition " 
func IsServerErrorPartitionWrongValuesError(err error) bool {
    result := Isa(err, ER_PARTITION_WRONG_VALUES_ERROR)
    return result
}

    
// IsServerErrorPartitionMaxvalueError check mysql error is "MAXVALUE can only be used in last partition definition " 
func IsServerErrorPartitionMaxvalueError(err error) bool {
    result := Isa(err, ER_PARTITION_MAXVALUE_ERROR)
    return result
}

    
// IsServerErrorPartitionSubpartitionError check mysql error is "Subpartitions can only be hash partitions and by key " 
func IsServerErrorPartitionSubpartitionError(err error) bool {
    result := Isa(err, ER_PARTITION_SUBPARTITION_ERROR)
    return result
}

    
// IsServerErrorPartitionSubpartMixError check mysql error is "Must define subpartitions on all partitions if on one partition " 
func IsServerErrorPartitionSubpartMixError(err error) bool {
    result := Isa(err, ER_PARTITION_SUBPART_MIX_ERROR)
    return result
}

    
// IsServerErrorPartitionWrongNoPartError check mysql error is "Wrong number of partitions defined, mismatch with previous setting " 
func IsServerErrorPartitionWrongNoPartError(err error) bool {
    result := Isa(err, ER_PARTITION_WRONG_NO_PART_ERROR)
    return result
}

    
// IsServerErrorPartitionWrongNoSubpartError check mysql error is "Wrong number of subpartitions defined, mismatch with previous setting " 
func IsServerErrorPartitionWrongNoSubpartError(err error) bool {
    result := Isa(err, ER_PARTITION_WRONG_NO_SUBPART_ERROR)
    return result
}

    
// IsServerErrorWrongExprInPartitionFuncError check mysql error is "Constant, random or timezone-dependent expressions in (sub)partitioning function are not allowed " 
func IsServerErrorWrongExprInPartitionFuncError(err error) bool {
    result := Isa(err, ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR)
    return result
}

    
// IsServerErrorNoConstExprInRangeOrListError check mysql error is "Expression in RANGE/LIST VALUES must be constant " 
func IsServerErrorNoConstExprInRangeOrListError(err error) bool {
    result := Isa(err, ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR)
    return result
}

    
// IsServerErrorFieldNotFoundPartError check mysql error is "Field in list of fields for partition function not found in table " 
func IsServerErrorFieldNotFoundPartError(err error) bool {
    result := Isa(err, ER_FIELD_NOT_FOUND_PART_ERROR)
    return result
}

    
// IsServerErrorListOfFieldsOnlyInHashError check mysql error is "List of fields is only allowed in KEY partitions " 
func IsServerErrorListOfFieldsOnlyInHashError(err error) bool {
    result := Isa(err, ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR)
    return result
}

    
// IsServerErrorInconsistentPartitionInfoError check mysql error is "The partition info in the frm file is not consistent with what can be written into the frm file " 
func IsServerErrorInconsistentPartitionInfoError(err error) bool {
    result := Isa(err, ER_INCONSISTENT_PARTITION_INFO_ERROR)
    return result
}

    
// IsServerErrorPartitionFuncNotAllowedError check mysql error is "The %s function returns the wrong type " 
func IsServerErrorPartitionFuncNotAllowedError(err error) bool {
    result := Isa(err, ER_PARTITION_FUNC_NOT_ALLOWED_ERROR)
    return result
}

    
// IsServerErrorPartitionsMustBeDefinedError check mysql error is "For %s partitions each partition must be defined " 
func IsServerErrorPartitionsMustBeDefinedError(err error) bool {
    result := Isa(err, ER_PARTITIONS_MUST_BE_DEFINED_ERROR)
    return result
}

    
// IsServerErrorRangeNotIncreasingError check mysql error is "VALUES LESS THAN value must be strictly increasing for each partition " 
func IsServerErrorRangeNotIncreasingError(err error) bool {
    result := Isa(err, ER_RANGE_NOT_INCREASING_ERROR)
    return result
}

    
// IsServerErrorInconsistentTypeOfFunctionsError check mysql error is "VALUES value must be of same type as partition function " 
func IsServerErrorInconsistentTypeOfFunctionsError(err error) bool {
    result := Isa(err, ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR)
    return result
}

    
// IsServerErrorMultipleDefConstInListPartError check mysql error is "Multiple definition of same constant in list partitioning " 
func IsServerErrorMultipleDefConstInListPartError(err error) bool {
    result := Isa(err, ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR)
    return result
}

    
// IsServerErrorPartitionEntryError check mysql error is "Partitioning can not be used stand-alone in query " 
func IsServerErrorPartitionEntryError(err error) bool {
    result := Isa(err, ER_PARTITION_ENTRY_ERROR)
    return result
}

    
// IsServerErrorMixHandlerError check mysql error is "The mix of handlers in the partitions is not allowed in this version of MySQL " 
func IsServerErrorMixHandlerError(err error) bool {
    result := Isa(err, ER_MIX_HANDLER_ERROR)
    return result
}

    
// IsServerErrorPartitionNotDefinedError check mysql error is "For the partitioned engine it is necessary to define all %s " 
func IsServerErrorPartitionNotDefinedError(err error) bool {
    result := Isa(err, ER_PARTITION_NOT_DEFINED_ERROR)
    return result
}

    
// IsServerErrorTooManyPartitionsError check mysql error is "Too many partitions (including subpartitions) were defined " 
func IsServerErrorTooManyPartitionsError(err error) bool {
    result := Isa(err, ER_TOO_MANY_PARTITIONS_ERROR)
    return result
}

    
// IsServerErrorSubpartitionError check mysql error is "It is only possible to mix RANGE/LIST partitioning with HASH/KEY partitioning for subpartitioning " 
func IsServerErrorSubpartitionError(err error) bool {
    result := Isa(err, ER_SUBPARTITION_ERROR)
    return result
}

    
// IsServerErrorCantCreateHandlerFile check mysql error is "Failed to create specific handler file " 
func IsServerErrorCantCreateHandlerFile(err error) bool {
    result := Isa(err, ER_CANT_CREATE_HANDLER_FILE)
    return result
}

    
// IsServerErrorBlobFieldInPartFuncError check mysql error is "A BLOB field is not allowed in partition function " 
func IsServerErrorBlobFieldInPartFuncError(err error) bool {
    result := Isa(err, ER_BLOB_FIELD_IN_PART_FUNC_ERROR)
    return result
}

    
// IsServerErrorUniqueKeyNeedAllFieldsInPf check mysql error is "A %s must include all columns in the table's partitioning function " 
func IsServerErrorUniqueKeyNeedAllFieldsInPf(err error) bool {
    result := Isa(err, ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF)
    return result
}

    
// IsServerErrorNoPartsError check mysql error is "Number of %s = 0 is not an allowed value " 
func IsServerErrorNoPartsError(err error) bool {
    result := Isa(err, ER_NO_PARTS_ERROR)
    return result
}

    
// IsServerErrorPartitionMgmtOnNonpartitioned check mysql error is "Partition management on a not partitioned table is not possible " 
func IsServerErrorPartitionMgmtOnNonpartitioned(err error) bool {
    result := Isa(err, ER_PARTITION_MGMT_ON_NONPARTITIONED)
    return result
}

    
// IsServerErrorForeignKeyOnPartitioned check mysql error is "Foreign keys are not yet supported in conjunction with partitioning " 
func IsServerErrorForeignKeyOnPartitioned(err error) bool {
    result := Isa(err, ER_FOREIGN_KEY_ON_PARTITIONED)
    return result
}

    
// IsServerErrorDropPartitionNonExistent check mysql error is "Error in list of partitions to %s " 
func IsServerErrorDropPartitionNonExistent(err error) bool {
    result := Isa(err, ER_DROP_PARTITION_NON_EXISTENT)
    return result
}

    
// IsServerErrorDropLastPartition check mysql error is "Cannot remove all partitions, use DROP TABLE instead " 
func IsServerErrorDropLastPartition(err error) bool {
    result := Isa(err, ER_DROP_LAST_PARTITION)
    return result
}

    
// IsServerErrorCoalesceOnlyOnHashPartition check mysql error is "COALESCE PARTITION can only be used on HASH/KEY partitions " 
func IsServerErrorCoalesceOnlyOnHashPartition(err error) bool {
    result := Isa(err, ER_COALESCE_ONLY_ON_HASH_PARTITION)
    return result
}

    
// IsServerErrorReorgHashOnlyOnSameNo check mysql error is "REORGANIZE PARTITION can only be used to reorganize partitions not to change their numbers " 
func IsServerErrorReorgHashOnlyOnSameNo(err error) bool {
    result := Isa(err, ER_REORG_HASH_ONLY_ON_SAME_NO)
    return result
}

    
// IsServerErrorReorgNoParamError check mysql error is "REORGANIZE PARTITION without parameters can only be used on auto-partitioned tables using HASH PARTITIONs " 
func IsServerErrorReorgNoParamError(err error) bool {
    result := Isa(err, ER_REORG_NO_PARAM_ERROR)
    return result
}

    
// IsServerErrorOnlyOnRangeListPartition check mysql error is "%s PARTITION can only be used on RANGE/LIST partitions " 
func IsServerErrorOnlyOnRangeListPartition(err error) bool {
    result := Isa(err, ER_ONLY_ON_RANGE_LIST_PARTITION)
    return result
}

    
// IsServerErrorAddPartitionSubpartError check mysql error is "Trying to Add partition(s) with wrong number of subpartitions " 
func IsServerErrorAddPartitionSubpartError(err error) bool {
    result := Isa(err, ER_ADD_PARTITION_SUBPART_ERROR)
    return result
}

    
// IsServerErrorAddPartitionNoNewPartition check mysql error is "At least one partition must be added " 
func IsServerErrorAddPartitionNoNewPartition(err error) bool {
    result := Isa(err, ER_ADD_PARTITION_NO_NEW_PARTITION)
    return result
}

    
// IsServerErrorCoalescePartitionNoPartition check mysql error is "At least one partition must be coalesced " 
func IsServerErrorCoalescePartitionNoPartition(err error) bool {
    result := Isa(err, ER_COALESCE_PARTITION_NO_PARTITION)
    return result
}

    
// IsServerErrorReorgPartitionNotExist check mysql error is "More partitions to reorganize than there are partitions " 
func IsServerErrorReorgPartitionNotExist(err error) bool {
    result := Isa(err, ER_REORG_PARTITION_NOT_EXIST)
    return result
}

    
// IsServerErrorSameNamePartition check mysql error is "Duplicate partition name %s " 
func IsServerErrorSameNamePartition(err error) bool {
    result := Isa(err, ER_SAME_NAME_PARTITION)
    return result
}

    
// IsServerErrorNoBinlogError check mysql error is "It is not allowed to shut off binlog on this command " 
func IsServerErrorNoBinlogError(err error) bool {
    result := Isa(err, ER_NO_BINLOG_ERROR)
    return result
}

    
// IsServerErrorConsecutiveReorgPartitions check mysql error is "When reorganizing a set of partitions they must be in consecutive order " 
func IsServerErrorConsecutiveReorgPartitions(err error) bool {
    result := Isa(err, ER_CONSECUTIVE_REORG_PARTITIONS)
    return result
}

    
// IsServerErrorReorgOutsideRange check mysql error is "Reorganize of range partitions cannot change total ranges except for last partition where it can extend the range " 
func IsServerErrorReorgOutsideRange(err error) bool {
    result := Isa(err, ER_REORG_OUTSIDE_RANGE)
    return result
}

    
// IsServerErrorPartitionFunctionFailure check mysql error is "Partition function not supported in this version for this handler " 
func IsServerErrorPartitionFunctionFailure(err error) bool {
    result := Isa(err, ER_PARTITION_FUNCTION_FAILURE)
    return result
}

    
// IsServerErrorPartStateError check mysql error is "Partition state cannot be defined from CREATE/ALTER TABLE " 
func IsServerErrorPartStateError(err error) bool {
    result := Isa(err, ER_PART_STATE_ERROR)
    return result
}

    
// IsServerErrorLimitedPartRange check mysql error is "The %s handler only supports 32 bit integers in VALUES " 
func IsServerErrorLimitedPartRange(err error) bool {
    result := Isa(err, ER_LIMITED_PART_RANGE)
    return result
}

    
// IsServerErrorPluginIsNotLoaded check mysql error is "Plugin '%s' is not loaded " 
func IsServerErrorPluginIsNotLoaded(err error) bool {
    result := Isa(err, ER_PLUGIN_IS_NOT_LOADED)
    return result
}

    
// IsServerErrorWrongValue check mysql error is "Incorrect %s value: '%s' " 
func IsServerErrorWrongValue(err error) bool {
    result := Isa(err, ER_WRONG_VALUE)
    return result
}

    
// IsServerErrorNoPartitionForGivenValue check mysql error is "Table has no partition for value %s " 
func IsServerErrorNoPartitionForGivenValue(err error) bool {
    result := Isa(err, ER_NO_PARTITION_FOR_GIVEN_VALUE)
    return result
}

    
// IsServerErrorFilegroupOptionOnlyOnce check mysql error is "It is not allowed to specify %s more than once " 
func IsServerErrorFilegroupOptionOnlyOnce(err error) bool {
    result := Isa(err, ER_FILEGROUP_OPTION_ONLY_ONCE)
    return result
}

    
// IsServerErrorCreateFilegroupFailed check mysql error is "Failed to create %s " 
func IsServerErrorCreateFilegroupFailed(err error) bool {
    result := Isa(err, ER_CREATE_FILEGROUP_FAILED)
    return result
}

    
// IsServerErrorDropFilegroupFailed check mysql error is "Failed to drop %s " 
func IsServerErrorDropFilegroupFailed(err error) bool {
    result := Isa(err, ER_DROP_FILEGROUP_FAILED)
    return result
}

    
// IsServerErrorTablespaceAutoExtendError check mysql error is "The handler doesn't support autoextend of tablespaces " 
func IsServerErrorTablespaceAutoExtendError(err error) bool {
    result := Isa(err, ER_TABLESPACE_AUTO_EXTEND_ERROR)
    return result
}

    
// IsServerErrorWrongSizeNumber check mysql error is "A size parameter was incorrectly specified, either number or on the form 10M " 
func IsServerErrorWrongSizeNumber(err error) bool {
    result := Isa(err, ER_WRONG_SIZE_NUMBER)
    return result
}

    
// IsServerErrorSizeOverflowError check mysql error is "The size number was correct but we don't allow the digit part to be more than 2 billion " 
func IsServerErrorSizeOverflowError(err error) bool {
    result := Isa(err, ER_SIZE_OVERFLOW_ERROR)
    return result
}

    
// IsServerErrorAlterFilegroupFailed check mysql error is "Failed to alter: %s " 
func IsServerErrorAlterFilegroupFailed(err error) bool {
    result := Isa(err, ER_ALTER_FILEGROUP_FAILED)
    return result
}

    
// IsServerErrorBinlogRowLoggingFailed check mysql error is "Writing one row to the row-based binary log failed " 
func IsServerErrorBinlogRowLoggingFailed(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_LOGGING_FAILED)
    return result
}

    
// IsServerErrorBinlogRowWrongTableDef check mysql error is "Table definition on master and slave does not match: %s " 
func IsServerErrorBinlogRowWrongTableDef(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_WRONG_TABLE_DEF)
    return result
}

    
// IsServerErrorBinlogRowRbrToSbr check mysql error is "Slave running with --log-slave-updates must use row-based binary logging to be able to replicate row-based binary log events " 
func IsServerErrorBinlogRowRbrToSbr(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_RBR_TO_SBR)
    return result
}

    
// IsServerErrorEventAlreadyExists check mysql error is "Event '%s' already exists " 
func IsServerErrorEventAlreadyExists(err error) bool {
    result := Isa(err, ER_EVENT_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorEventStoreFailed check mysql error is "Failed to store event %s. Error code %d from storage engine. " 
func IsServerErrorEventStoreFailed(err error) bool {
    result := Isa(err, ER_EVENT_STORE_FAILED)
    return result
}

    
// IsServerErrorEventDoesNotExist check mysql error is "Unknown event '%s' " 
func IsServerErrorEventDoesNotExist(err error) bool {
    result := Isa(err, ER_EVENT_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorEventCantAlter check mysql error is "Failed to alter event '%s' " 
func IsServerErrorEventCantAlter(err error) bool {
    result := Isa(err, ER_EVENT_CANT_ALTER)
    return result
}

    
// IsServerErrorEventDropFailed check mysql error is "Failed to drop %s " 
func IsServerErrorEventDropFailed(err error) bool {
    result := Isa(err, ER_EVENT_DROP_FAILED)
    return result
}

    
// IsServerErrorEventIntervalNotPositiveOrTooBig check mysql error is "INTERVAL is either not positive or too big " 
func IsServerErrorEventIntervalNotPositiveOrTooBig(err error) bool {
    result := Isa(err, ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG)
    return result
}

    
// IsServerErrorEventEndsBeforeStarts check mysql error is "ENDS is either invalid or before STARTS " 
func IsServerErrorEventEndsBeforeStarts(err error) bool {
    result := Isa(err, ER_EVENT_ENDS_BEFORE_STARTS)
    return result
}

    
// IsServerErrorEventExecTimeInThePast check mysql error is "Event execution time is in the past. Event has been disabled " 
func IsServerErrorEventExecTimeInThePast(err error) bool {
    result := Isa(err, ER_EVENT_EXEC_TIME_IN_THE_PAST)
    return result
}

    
// IsServerErrorEventOpenTableFailed check mysql error is "Failed to open mysql.event " 
func IsServerErrorEventOpenTableFailed(err error) bool {
    result := Isa(err, ER_EVENT_OPEN_TABLE_FAILED)
    return result
}

    
// IsServerErrorEventNeitherMExprNorMAt check mysql error is "No datetime expression provided " 
func IsServerErrorEventNeitherMExprNorMAt(err error) bool {
    result := Isa(err, ER_EVENT_NEITHER_M_EXPR_NOR_M_AT)
    return result
}

    
// IsServerErrorObsoleteColCountDoesntMatchCorrupted check mysql error is "Column count of mysql.%s is wrong. Expected %d, found %d. The table is probably corrupted " 
func IsServerErrorObsoleteColCountDoesntMatchCorrupted(err error) bool {
    result := Isa(err, ER_OBSOLETE_COL_COUNT_DOESNT_MATCH_CORRUPTED)
    return result
}

    
// IsServerErrorObsoleteCannotLoadFromTable check mysql error is "Cannot load from mysql.%s. The table is probably corrupted " 
func IsServerErrorObsoleteCannotLoadFromTable(err error) bool {
    result := Isa(err, ER_OBSOLETE_CANNOT_LOAD_FROM_TABLE)
    return result
}

    
// IsServerErrorEventCannotDelete check mysql error is "Failed to delete the event from mysql.event " 
func IsServerErrorEventCannotDelete(err error) bool {
    result := Isa(err, ER_EVENT_CANNOT_DELETE)
    return result
}

    
// IsServerErrorEventCompileError check mysql error is "Error during compilation of event's body " 
func IsServerErrorEventCompileError(err error) bool {
    result := Isa(err, ER_EVENT_COMPILE_ERROR)
    return result
}

    
// IsServerErrorEventSameName check mysql error is "Same old and new event name " 
func IsServerErrorEventSameName(err error) bool {
    result := Isa(err, ER_EVENT_SAME_NAME)
    return result
}

    
// IsServerErrorEventDataTooLong check mysql error is "Data for column '%s' too long " 
func IsServerErrorEventDataTooLong(err error) bool {
    result := Isa(err, ER_EVENT_DATA_TOO_LONG)
    return result
}

    
// IsServerErrorDropIndexFk check mysql error is "Cannot drop index '%s': needed in a foreign key constraint" 
func IsServerErrorDropIndexFk(err error) bool {
    result := Isa(err, ER_DROP_INDEX_FK)
    return result
}

    
// IsServerErrorWarnDeprecatedSyntaxWithVer check mysql error is "The syntax '%s' is deprecated and will be removed in MySQL %s. Please use %s instead " 
func IsServerErrorWarnDeprecatedSyntaxWithVer(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYNTAX_WITH_VER)
    return result
}

    
// IsServerErrorCantWriteLockLogTable check mysql error is "You can't write-lock a log table. Only read access is possible " 
func IsServerErrorCantWriteLockLogTable(err error) bool {
    result := Isa(err, ER_CANT_WRITE_LOCK_LOG_TABLE)
    return result
}

    
// IsServerErrorCantLockLogTable check mysql error is "You can't use locks with log tables. " 
func IsServerErrorCantLockLogTable(err error) bool {
    result := Isa(err, ER_CANT_LOCK_LOG_TABLE)
    return result
}

    
// IsServerErrorForeignDuplicateKeyOldUnused check mysql error is "Upholding foreign key constraints for table '%s', entry '%s', key %d would lead to a duplicate entry " 
func IsServerErrorForeignDuplicateKeyOldUnused(err error) bool {
    result := Isa(err, ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED)
    return result
}

    
// IsServerErrorColCountDoesntMatchPleaseUpdate check mysql error is "Column count of mysql.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please use mysql_upgrade to fix this error. " 
func IsServerErrorColCountDoesntMatchPleaseUpdate(err error) bool {
    result := Isa(err, ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE)
    return result
}

    
// IsServerErrorTempTablePreventsSwitchOutOfRbr check mysql error is "Cannot switch out of the row-based binary log format when the session has open temporary tables " 
func IsServerErrorTempTablePreventsSwitchOutOfRbr(err error) bool {
    result := Isa(err, ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR)
    return result
}

    
// IsServerErrorStoredFunctionPreventsSwitchBinlogFormat check mysql error is "Cannot change the binary logging format inside a stored function or trigger " 
func IsServerErrorStoredFunctionPreventsSwitchBinlogFormat(err error) bool {
    result := Isa(err, ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT)
    return result
}

    
// IsServerErrorNdbCantSwitchBinlogFormat check mysql error is "The NDB cluster engine does not support changing the binlog format on the fly yet " 
func IsServerErrorNdbCantSwitchBinlogFormat(err error) bool {
    result := Isa(err, ER_NDB_CANT_SWITCH_BINLOG_FORMAT)
    return result
}

    
// IsServerErrorPartitionNoTemporary check mysql error is "Cannot create temporary table with partitions " 
func IsServerErrorPartitionNoTemporary(err error) bool {
    result := Isa(err, ER_PARTITION_NO_TEMPORARY)
    return result
}

    
// IsServerErrorPartitionConstDomainError check mysql error is "Partition constant is out of partition function domain " 
func IsServerErrorPartitionConstDomainError(err error) bool {
    result := Isa(err, ER_PARTITION_CONST_DOMAIN_ERROR)
    return result
}

    
// IsServerErrorPartitionFunctionIsNotAllowed check mysql error is "This partition function is not allowed " 
func IsServerErrorPartitionFunctionIsNotAllowed(err error) bool {
    result := Isa(err, ER_PARTITION_FUNCTION_IS_NOT_ALLOWED)
    return result
}

    
// IsServerErrorDdlLogError check mysql error is "Error in DDL log " 
func IsServerErrorDdlLogError(err error) bool {
    result := Isa(err, ER_DDL_LOG_ERROR)
    return result
}

    
// IsServerErrorNullInValuesLessThan check mysql error is "Not allowed to use NULL value in VALUES LESS THAN " 
func IsServerErrorNullInValuesLessThan(err error) bool {
    result := Isa(err, ER_NULL_IN_VALUES_LESS_THAN)
    return result
}

    
// IsServerErrorWrongPartitionName check mysql error is "Incorrect partition name " 
func IsServerErrorWrongPartitionName(err error) bool {
    result := Isa(err, ER_WRONG_PARTITION_NAME)
    return result
}

    
// IsServerErrorCantChangeTxCharacteristics check mysql error is "Transaction characteristics can't be changed while a transaction is in progress " 
func IsServerErrorCantChangeTxCharacteristics(err error) bool {
    result := Isa(err, ER_CANT_CHANGE_TX_CHARACTERISTICS)
    return result
}

    
// IsServerErrorDupEntryAutoincrementCase check mysql error is "ALTER TABLE causes auto_increment resequencing, resulting in duplicate entry '%s' for key '%s' " 
func IsServerErrorDupEntryAutoincrementCase(err error) bool {
    result := Isa(err, ER_DUP_ENTRY_AUTOINCREMENT_CASE)
    return result
}

    
// IsServerErrorEventModifyQueueError check mysql error is "Internal scheduler error %d " 
func IsServerErrorEventModifyQueueError(err error) bool {
    result := Isa(err, ER_EVENT_MODIFY_QUEUE_ERROR)
    return result
}

    
// IsServerErrorEventSetVarError check mysql error is "Error during starting/stopping of the scheduler. Error code %u " 
func IsServerErrorEventSetVarError(err error) bool {
    result := Isa(err, ER_EVENT_SET_VAR_ERROR)
    return result
}

    
// IsServerErrorPartitionMergeError check mysql error is "Engine cannot be used in partitioned tables " 
func IsServerErrorPartitionMergeError(err error) bool {
    result := Isa(err, ER_PARTITION_MERGE_ERROR)
    return result
}

    
// IsServerErrorCantActivateLog check mysql error is "Cannot activate '%s' log " 
func IsServerErrorCantActivateLog(err error) bool {
    result := Isa(err, ER_CANT_ACTIVATE_LOG)
    return result
}

    
// IsServerErrorRbrNotAvailable check mysql error is "The server was not built with row-based replication " 
func IsServerErrorRbrNotAvailable(err error) bool {
    result := Isa(err, ER_RBR_NOT_AVAILABLE)
    return result
}

    
// IsServerErrorBase64DecodeError check mysql error is "Decoding of base64 string failed " 
func IsServerErrorBase64DecodeError(err error) bool {
    result := Isa(err, ER_BASE64_DECODE_ERROR)
    return result
}

    
// IsServerErrorEventRecursionForbidden check mysql error is "Recursion of EVENT DDL statements is forbidden when body is present " 
func IsServerErrorEventRecursionForbidden(err error) bool {
    result := Isa(err, ER_EVENT_RECURSION_FORBIDDEN)
    return result
}

    
// IsServerErrorEventsDbError check mysql error is "Cannot proceed because system tables used by Event Scheduler were found damaged at server start" 
func IsServerErrorEventsDbError(err error) bool {
    result := Isa(err, ER_EVENTS_DB_ERROR)
    return result
}

    
// IsServerErrorOnlyIntegersAllowed check mysql error is "Only integers allowed as number here " 
func IsServerErrorOnlyIntegersAllowed(err error) bool {
    result := Isa(err, ER_ONLY_INTEGERS_ALLOWED)
    return result
}

    
// IsServerErrorUnsuportedLogEngine check mysql error is "This storage engine cannot be used for log tables" " 
func IsServerErrorUnsuportedLogEngine(err error) bool {
    result := Isa(err, ER_UNSUPORTED_LOG_ENGINE)
    return result
}

    
// IsServerErrorBadLogStatement check mysql error is "You cannot '%s' a log table if logging is enabled " 
func IsServerErrorBadLogStatement(err error) bool {
    result := Isa(err, ER_BAD_LOG_STATEMENT)
    return result
}

    
// IsServerErrorCantRenameLogTable check mysql error is "Cannot rename '%s'. When logging enabled, rename to/from log table must rename two tables: the log table to an archive table and another table back to '%s' " 
func IsServerErrorCantRenameLogTable(err error) bool {
    result := Isa(err, ER_CANT_RENAME_LOG_TABLE)
    return result
}

    
// IsServerErrorWrongParamcountToNativeFct check mysql error is "Incorrect parameter count in the call to native function '%s' " 
func IsServerErrorWrongParamcountToNativeFct(err error) bool {
    result := Isa(err, ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT)
    return result
}

    
// IsServerErrorWrongParametersToNativeFct check mysql error is "Incorrect parameters in the call to native function '%s' " 
func IsServerErrorWrongParametersToNativeFct(err error) bool {
    result := Isa(err, ER_WRONG_PARAMETERS_TO_NATIVE_FCT)
    return result
}

    
// IsServerErrorWrongParametersToStoredFct check mysql error is "Incorrect parameters in the call to stored function %s " 
func IsServerErrorWrongParametersToStoredFct(err error) bool {
    result := Isa(err, ER_WRONG_PARAMETERS_TO_STORED_FCT)
    return result
}

    
// IsServerErrorNativeFctNameCollision check mysql error is "This function '%s' has the same name as a native function " 
func IsServerErrorNativeFctNameCollision(err error) bool {
    result := Isa(err, ER_NATIVE_FCT_NAME_COLLISION)
    return result
}

    
// IsServerErrorDupEntryWithKeyName check mysql error is "Duplicate entry '%s' for key '%s'" 
func IsServerErrorDupEntryWithKeyName(err error) bool {
    result := Isa(err, ER_DUP_ENTRY_WITH_KEY_NAME)
    return result
}

    
// IsServerErrorBinlogPurgeEmfile check mysql error is "Too many files opened, please execute the command again " 
func IsServerErrorBinlogPurgeEmfile(err error) bool {
    result := Isa(err, ER_BINLOG_PURGE_EMFILE)
    return result
}

    
// IsServerErrorEventCannotCreateInThePast check mysql error is "Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was dropped immediately after creation. " 
func IsServerErrorEventCannotCreateInThePast(err error) bool {
    result := Isa(err, ER_EVENT_CANNOT_CREATE_IN_THE_PAST)
    return result
}

    
// IsServerErrorEventCannotAlterInThePast check mysql error is "Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was not changed. Specify a time in the future. " 
func IsServerErrorEventCannotAlterInThePast(err error) bool {
    result := Isa(err, ER_EVENT_CANNOT_ALTER_IN_THE_PAST)
    return result
}

    
// IsServerErrorSlaveIncident check mysql error is "The incident %s occured on the master. Message: %s " 
func IsServerErrorSlaveIncident(err error) bool {
    result := Isa(err, ER_SLAVE_INCIDENT)
    return result
}

    
// IsServerErrorNoPartitionForGivenValueSilent check mysql error is "Table has no partition for some existing values " 
func IsServerErrorNoPartitionForGivenValueSilent(err error) bool {
    result := Isa(err, ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT)
    return result
}

    
// IsServerErrorBinlogUnsafeStatement check mysql error is "Unsafe statement written to the binary log using statement format since BINLOG_FORMAT = STATEMENT. %s " 
func IsServerErrorBinlogUnsafeStatement(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_STATEMENT)
    return result
}

    
// IsServerErrorSlaveFatalError check mysql error is "Fatal error: %s " 
func IsServerErrorSlaveFatalError(err error) bool {
    result := Isa(err, ER_SLAVE_FATAL_ERROR)
    return result
}

    
// IsServerErrorSlaveRelayLogReadFailure check mysql error is "Relay log read failure: %s " 
func IsServerErrorSlaveRelayLogReadFailure(err error) bool {
    result := Isa(err, ER_SLAVE_RELAY_LOG_READ_FAILURE)
    return result
}

    
// IsServerErrorSlaveRelayLogWriteFailure check mysql error is "Relay log write failure: %s " 
func IsServerErrorSlaveRelayLogWriteFailure(err error) bool {
    result := Isa(err, ER_SLAVE_RELAY_LOG_WRITE_FAILURE)
    return result
}

    
// IsServerErrorSlaveCreateEventFailure check mysql error is "Failed to create %s " 
func IsServerErrorSlaveCreateEventFailure(err error) bool {
    result := Isa(err, ER_SLAVE_CREATE_EVENT_FAILURE)
    return result
}

    
// IsServerErrorSlaveMasterComFailure check mysql error is "Master command %s failed: %s " 
func IsServerErrorSlaveMasterComFailure(err error) bool {
    result := Isa(err, ER_SLAVE_MASTER_COM_FAILURE)
    return result
}

    
// IsServerErrorBinlogLoggingImpossible check mysql error is "Binary logging not possible. Message: %s " 
func IsServerErrorBinlogLoggingImpossible(err error) bool {
    result := Isa(err, ER_BINLOG_LOGGING_IMPOSSIBLE)
    return result
}

    
// IsServerErrorViewNoCreationCtx check mysql error is "View `%s`.`%s` has no creation context " 
func IsServerErrorViewNoCreationCtx(err error) bool {
    result := Isa(err, ER_VIEW_NO_CREATION_CTX)
    return result
}

    
// IsServerErrorViewInvalidCreationCtx check mysql error is "Creation context of view `%s`.`%s' is invalid " 
func IsServerErrorViewInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_VIEW_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorSrInvalidCreationCtx check mysql error is "Creation context of stored routine `%s`.`%s` is invalid " 
func IsServerErrorSrInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_SR_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorTrgCorruptedFile check mysql error is "Corrupted TRG file for table `%s`.`%s` " 
func IsServerErrorTrgCorruptedFile(err error) bool {
    result := Isa(err, ER_TRG_CORRUPTED_FILE)
    return result
}

    
// IsServerErrorTrgNoCreationCtx check mysql error is "Triggers for table `%s`.`%s` have no creation context " 
func IsServerErrorTrgNoCreationCtx(err error) bool {
    result := Isa(err, ER_TRG_NO_CREATION_CTX)
    return result
}

    
// IsServerErrorTrgInvalidCreationCtx check mysql error is "Trigger creation context of table `%s`.`%s` is invalid " 
func IsServerErrorTrgInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_TRG_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorEventInvalidCreationCtx check mysql error is "Creation context of event `%s`.`%s` is invalid " 
func IsServerErrorEventInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_EVENT_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorTrgCantOpenTable check mysql error is "Cannot open table for trigger `%s`.`%s` " 
func IsServerErrorTrgCantOpenTable(err error) bool {
    result := Isa(err, ER_TRG_CANT_OPEN_TABLE)
    return result
}

    
// IsServerErrorCantCreateSroutine check mysql error is "Cannot create stored routine `%s`. Check warnings " 
func IsServerErrorCantCreateSroutine(err error) bool {
    result := Isa(err, ER_CANT_CREATE_SROUTINE)
    return result
}

    
// IsServerErrorNeverUsed check mysql error is "Ambiguous slave modes combination. %s " 
func IsServerErrorNeverUsed(err error) bool {
    result := Isa(err, ER_NEVER_USED)
    return result
}

    
// IsServerErrorNoFormatDescriptionEventBeforeBinlogStatement check mysql error is "The BINLOG statement of type `%s` was not preceded by a format description BINLOG statement. " 
func IsServerErrorNoFormatDescriptionEventBeforeBinlogStatement(err error) bool {
    result := Isa(err, ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT)
    return result
}

    
// IsServerErrorSlaveCorruptEvent check mysql error is "Corrupted replication event was detected " 
func IsServerErrorSlaveCorruptEvent(err error) bool {
    result := Isa(err, ER_SLAVE_CORRUPT_EVENT)
    return result
}

    
// IsServerErrorLoadDataInvalidColumn check mysql error is "Invalid column reference (%s) in LOAD DATA" 
func IsServerErrorLoadDataInvalidColumn(err error) bool {
    result := Isa(err, ER_LOAD_DATA_INVALID_COLUMN)
    return result
}

    
// IsServerErrorLoadDataInvalidColumnUnused check mysql error is "Invalid column reference (%s) in LOAD DATA" 
func IsServerErrorLoadDataInvalidColumnUnused(err error) bool {
    result := Isa(err, ER_LOAD_DATA_INVALID_COLUMN_UNUSED)
    return result
}

    
// IsServerErrorLogPurgeNoFile check mysql error is "Being purged log %s was not found " 
func IsServerErrorLogPurgeNoFile(err error) bool {
    result := Isa(err, ER_LOG_PURGE_NO_FILE)
    return result
}

    
// IsServerErrorXaRbtimeout check mysql error is "XA_RBTIMEOUT: Transaction branch was rolled back: took too long " 
func IsServerErrorXaRbtimeout(err error) bool {
    result := Isa(err, ER_XA_RBTIMEOUT)
    return result
}

    
// IsServerErrorXaRbdeadlock check mysql error is "XA_RBDEADLOCK: Transaction branch was rolled back: deadlock was detected " 
func IsServerErrorXaRbdeadlock(err error) bool {
    result := Isa(err, ER_XA_RBDEADLOCK)
    return result
}

    
// IsServerErrorNeedReprepare check mysql error is "Prepared statement needs to be re-prepared " 
func IsServerErrorNeedReprepare(err error) bool {
    result := Isa(err, ER_NEED_REPREPARE)
    return result
}

    
// IsServerErrorDelayedNotSupported check mysql error is "DELAYED option not supported for table '%s' " 
func IsServerErrorDelayedNotSupported(err error) bool {
    result := Isa(err, ER_DELAYED_NOT_SUPPORTED)
    return result
}

    
// IsWarnNoMasterInfo check mysql error is "The master info structure does not exist " 
func IsWarnNoMasterInfo(err error) bool {
    result := Isa(err, WARN_NO_MASTER_INFO)
    return result
}

    
// IsWarnOptionIgnored check mysql error is "<%s> option ignored " 
func IsWarnOptionIgnored(err error) bool {
    result := Isa(err, WARN_OPTION_IGNORED)
    return result
}

    
// IsWarnPluginDeleteBuiltin check mysql error is "Built-in plugins cannot be deleted" 
func IsWarnPluginDeleteBuiltin(err error) bool {
    result := Isa(err, WARN_PLUGIN_DELETE_BUILTIN)
    return result
}

    
// IsServerErrorPluginDeleteBuiltin check mysql error is "Built-in plugins cannot be deleted" 
func IsServerErrorPluginDeleteBuiltin(err error) bool {
    result := Isa(err, ER_PLUGIN_DELETE_BUILTIN)
    return result
}

    
// IsWarnPluginBusy check mysql error is "Plugin is busy and will be uninstalled on shutdown " 
func IsWarnPluginBusy(err error) bool {
    result := Isa(err, WARN_PLUGIN_BUSY)
    return result
}

    
// IsServerErrorVariableIsReadonly check mysql error is "%s variable '%s' is read-only. Use SET %s to assign the value " 
func IsServerErrorVariableIsReadonly(err error) bool {
    result := Isa(err, ER_VARIABLE_IS_READONLY)
    return result
}

    
// IsServerErrorWarnEngineTransactionRollback check mysql error is "Storage engine %s does not support rollback for this statement. Transaction rolled back and must be restarted " 
func IsServerErrorWarnEngineTransactionRollback(err error) bool {
    result := Isa(err, ER_WARN_ENGINE_TRANSACTION_ROLLBACK)
    return result
}

    
// IsServerErrorSlaveHeartbeatFailure check mysql error is "Unexpected master's heartbeat data: %s " 
func IsServerErrorSlaveHeartbeatFailure(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_FAILURE)
    return result
}

    
// IsServerErrorSlaveHeartbeatValueOutOfRange check mysql error is "The requested value for the heartbeat period is either negative or exceeds the maximum allowed (%s seconds). " 
func IsServerErrorSlaveHeartbeatValueOutOfRange(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorNdbReplicationSchemaError check mysql error is "Bad schema for mysql.ndb_replication table. Message: %s " 
func IsServerErrorNdbReplicationSchemaError(err error) bool {
    result := Isa(err, ER_NDB_REPLICATION_SCHEMA_ERROR)
    return result
}

    
// IsServerErrorConflictFnParseError check mysql error is "Error in parsing conflict function. Message: %s " 
func IsServerErrorConflictFnParseError(err error) bool {
    result := Isa(err, ER_CONFLICT_FN_PARSE_ERROR)
    return result
}

    
// IsServerErrorExceptionsWriteError check mysql error is "Write to exceptions table failed. Message: %s" " 
func IsServerErrorExceptionsWriteError(err error) bool {
    result := Isa(err, ER_EXCEPTIONS_WRITE_ERROR)
    return result
}

    
// IsServerErrorTooLongTableComment check mysql error is "Comment for table '%s' is too long (max = %lu) " 
func IsServerErrorTooLongTableComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_TABLE_COMMENT)
    return result
}

    
// IsServerErrorTooLongFieldComment check mysql error is "Comment for field '%s' is too long (max = %lu) " 
func IsServerErrorTooLongFieldComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_FIELD_COMMENT)
    return result
}

    
// IsServerErrorFuncInexistentNameCollision check mysql error is "FUNCTION %s does not exist. Check the 'Function Name Parsing and Resolution' section in the Reference Manual " 
func IsServerErrorFuncInexistentNameCollision(err error) bool {
    result := Isa(err, ER_FUNC_INEXISTENT_NAME_COLLISION)
    return result
}

    
// IsServerErrorDatabaseName check mysql error is "Database " 
func IsServerErrorDatabaseName(err error) bool {
    result := Isa(err, ER_DATABASE_NAME)
    return result
}

    
// IsServerErrorTableName check mysql error is "Table " 
func IsServerErrorTableName(err error) bool {
    result := Isa(err, ER_TABLE_NAME)
    return result
}

    
// IsServerErrorPartitionName check mysql error is "Partition " 
func IsServerErrorPartitionName(err error) bool {
    result := Isa(err, ER_PARTITION_NAME)
    return result
}

    
// IsServerErrorSubpartitionName check mysql error is "Subpartition " 
func IsServerErrorSubpartitionName(err error) bool {
    result := Isa(err, ER_SUBPARTITION_NAME)
    return result
}

    
// IsServerErrorTemporaryName check mysql error is "Temporary " 
func IsServerErrorTemporaryName(err error) bool {
    result := Isa(err, ER_TEMPORARY_NAME)
    return result
}

    
// IsServerErrorRenamedName check mysql error is "Renamed " 
func IsServerErrorRenamedName(err error) bool {
    result := Isa(err, ER_RENAMED_NAME)
    return result
}

    
// IsServerErrorTooManyConcurrentTrxs check mysql error is "Too many active concurrent transactions " 
func IsServerErrorTooManyConcurrentTrxs(err error) bool {
    result := Isa(err, ER_TOO_MANY_CONCURRENT_TRXS)
    return result
}

    
// IsWarnNonAsciiSeparatorNotImplemented check mysql error is "Non-ASCII separator arguments are not fully supported " 
func IsWarnNonAsciiSeparatorNotImplemented(err error) bool {
    result := Isa(err, WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED)
    return result
}

    
// IsServerErrorDebugSyncTimeout check mysql error is "debug sync point wait timed out " 
func IsServerErrorDebugSyncTimeout(err error) bool {
    result := Isa(err, ER_DEBUG_SYNC_TIMEOUT)
    return result
}

    
// IsServerErrorDebugSyncHitLimit check mysql error is "debug sync point hit limit reached " 
func IsServerErrorDebugSyncHitLimit(err error) bool {
    result := Isa(err, ER_DEBUG_SYNC_HIT_LIMIT)
    return result
}

    
// IsServerErrorDupSignalSet check mysql error is "Duplicate condition information item '%s' " 
func IsServerErrorDupSignalSet(err error) bool {
    result := Isa(err, ER_DUP_SIGNAL_SET)
    return result
}

    
// IsServerErrorSignalWarn check mysql error is "Unhandled user-defined warning condition " 
func IsServerErrorSignalWarn(err error) bool {
    result := Isa(err, ER_SIGNAL_WARN)
    return result
}

    
// IsServerErrorSignalNotFound check mysql error is "Unhandled user-defined not found condition " 
func IsServerErrorSignalNotFound(err error) bool {
    result := Isa(err, ER_SIGNAL_NOT_FOUND)
    return result
}

    
// IsServerErrorSignalException check mysql error is "Unhandled user-defined exception condition " 
func IsServerErrorSignalException(err error) bool {
    result := Isa(err, ER_SIGNAL_EXCEPTION)
    return result
}

    
// IsServerErrorResignalWithoutActiveHandler check mysql error is "RESIGNAL when handler not active " 
func IsServerErrorResignalWithoutActiveHandler(err error) bool {
    result := Isa(err, ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER)
    return result
}

    
// IsServerErrorSignalBadConditionType check mysql error is "SIGNAL/RESIGNAL can only use a CONDITION defined with SQLSTATE " 
func IsServerErrorSignalBadConditionType(err error) bool {
    result := Isa(err, ER_SIGNAL_BAD_CONDITION_TYPE)
    return result
}

    
// IsWarnCondItemTruncated check mysql error is "Data truncated for condition item '%s' " 
func IsWarnCondItemTruncated(err error) bool {
    result := Isa(err, WARN_COND_ITEM_TRUNCATED)
    return result
}

    
// IsServerErrorCondItemTooLong check mysql error is "Data too long for condition item '%s' " 
func IsServerErrorCondItemTooLong(err error) bool {
    result := Isa(err, ER_COND_ITEM_TOO_LONG)
    return result
}

    
// IsServerErrorUnknownLocale check mysql error is "Unknown locale: '%s' " 
func IsServerErrorUnknownLocale(err error) bool {
    result := Isa(err, ER_UNKNOWN_LOCALE)
    return result
}

    
// IsServerErrorSlaveIgnoreServerIds check mysql error is "The requested server id %d clashes with the slave startup option --replicate-same-server-id " 
func IsServerErrorSlaveIgnoreServerIds(err error) bool {
    result := Isa(err, ER_SLAVE_IGNORE_SERVER_IDS)
    return result
}

    
// IsServerErrorQueryCacheDisabled check mysql error is "Query cache is disabled" 
func IsServerErrorQueryCacheDisabled(err error) bool {
    result := Isa(err, ER_QUERY_CACHE_DISABLED)
    return result
}

    
// IsServerErrorSameNamePartitionField check mysql error is "Duplicate partition field name '%s' " 
func IsServerErrorSameNamePartitionField(err error) bool {
    result := Isa(err, ER_SAME_NAME_PARTITION_FIELD)
    return result
}

    
// IsServerErrorPartitionColumnListError check mysql error is "Inconsistency in usage of column lists for partitioning " 
func IsServerErrorPartitionColumnListError(err error) bool {
    result := Isa(err, ER_PARTITION_COLUMN_LIST_ERROR)
    return result
}

    
// IsServerErrorWrongTypeColumnValueError check mysql error is "Partition column values of incorrect type " 
func IsServerErrorWrongTypeColumnValueError(err error) bool {
    result := Isa(err, ER_WRONG_TYPE_COLUMN_VALUE_ERROR)
    return result
}

    
// IsServerErrorTooManyPartitionFuncFieldsError check mysql error is "Too many fields in '%s' " 
func IsServerErrorTooManyPartitionFuncFieldsError(err error) bool {
    result := Isa(err, ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR)
    return result
}

    
// IsServerErrorMaxvalueInValuesIn check mysql error is "Cannot use MAXVALUE as value in VALUES IN " 
func IsServerErrorMaxvalueInValuesIn(err error) bool {
    result := Isa(err, ER_MAXVALUE_IN_VALUES_IN)
    return result
}

    
// IsServerErrorTooManyValuesError check mysql error is "Cannot have more than one value for this type of %s partitioning " 
func IsServerErrorTooManyValuesError(err error) bool {
    result := Isa(err, ER_TOO_MANY_VALUES_ERROR)
    return result
}

    
// IsServerErrorRowSinglePartitionFieldError check mysql error is "Row expressions in VALUES IN only allowed for multi-field column partitioning " 
func IsServerErrorRowSinglePartitionFieldError(err error) bool {
    result := Isa(err, ER_ROW_SINGLE_PARTITION_FIELD_ERROR)
    return result
}

    
// IsServerErrorFieldTypeNotAllowedAsPartitionField check mysql error is "Field '%s' is of a not allowed type for this type of partitioning " 
func IsServerErrorFieldTypeNotAllowedAsPartitionField(err error) bool {
    result := Isa(err, ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD)
    return result
}

    
// IsServerErrorPartitionFieldsTooLong check mysql error is "The total length of the partitioning fields is too large " 
func IsServerErrorPartitionFieldsTooLong(err error) bool {
    result := Isa(err, ER_PARTITION_FIELDS_TOO_LONG)
    return result
}

    
// IsServerErrorBinlogRowEngineAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binary log since both row-incapable engines and statement-incapable engines are involved. " 
func IsServerErrorBinlogRowEngineAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogRowModeAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = ROW and at least one table uses a storage engine limited to statement-based logging. " 
func IsServerErrorBinlogRowModeAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_MODE_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binary log since statement is unsafe, storage engine is limited to statement-based logging, and BINLOG_FORMAT = MIXED. %s " 
func IsServerErrorBinlogUnsafeAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogRowInjectionAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binary log since statement is in row format and at least one table uses a storage engine limited to statement-based logging. " 
func IsServerErrorBinlogRowInjectionAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogStmtModeAndRowEngine check mysql error is "Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT and at least one table uses a storage engine limited to row-based logging.%s " 
func IsServerErrorBinlogStmtModeAndRowEngine(err error) bool {
    result := Isa(err, ER_BINLOG_STMT_MODE_AND_ROW_ENGINE)
    return result
}

    
// IsServerErrorBinlogRowInjectionAndStmtMode check mysql error is "Cannot execute statement: impossible to write to binary log since statement is in row format and BINLOG_FORMAT = STATEMENT. " 
func IsServerErrorBinlogRowInjectionAndStmtMode(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_INJECTION_AND_STMT_MODE)
    return result
}

    
// IsServerErrorBinlogMultipleEnginesAndSelfLoggingEngine check mysql error is "Cannot execute statement: impossible to write to binary log since more than one engine is involved and at least one engine is self-logging. " 
func IsServerErrorBinlogMultipleEnginesAndSelfLoggingEngine(err error) bool {
    result := Isa(err, ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeLimit check mysql error is "The statement is unsafe because it uses a LIMIT clause. This is unsafe because the set of rows included cannot be predicted. " 
func IsServerErrorBinlogUnsafeLimit(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_LIMIT)
    return result
}

    
// IsServerErrorUnused4 check mysql error is "The statement is unsafe because it uses INSERT DELAYED. This is unsafe because the times when rows are inserted cannot be predicted. " 
func IsServerErrorUnused4(err error) bool {
    result := Isa(err, ER_UNUSED4)
    return result
}

    
// IsServerErrorBinlogUnsafeSystemTable check mysql error is "The statement is unsafe because it uses the general log, slow query log, or performance_schema table(s). This is unsafe because system tables may differ on slaves. " 
func IsServerErrorBinlogUnsafeSystemTable(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_SYSTEM_TABLE)
    return result
}

    
// IsServerErrorBinlogUnsafeAutoincColumns check mysql error is "Statement is unsafe because it invokes a trigger or a stored function that inserts into an AUTO_INCREMENT column. Inserted values cannot be logged correctly. " 
func IsServerErrorBinlogUnsafeAutoincColumns(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_AUTOINC_COLUMNS)
    return result
}

    
// IsServerErrorBinlogUnsafeUdf check mysql error is "Statement is unsafe because it uses a UDF which may not return the same value on the slave. " 
func IsServerErrorBinlogUnsafeUdf(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_UDF)
    return result
}

    
// IsServerErrorBinlogUnsafeSystemVariable check mysql error is "Statement is unsafe because it uses a system variable that may have a different value on the slave. " 
func IsServerErrorBinlogUnsafeSystemVariable(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_SYSTEM_VARIABLE)
    return result
}

    
// IsServerErrorBinlogUnsafeSystemFunction check mysql error is "Statement is unsafe because it uses a system function that may return a different value on the slave. " 
func IsServerErrorBinlogUnsafeSystemFunction(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_SYSTEM_FUNCTION)
    return result
}

    
// IsServerErrorBinlogUnsafeNontransAfterTrans check mysql error is "Statement is unsafe because it accesses a non-transactional table after accessing a transactional table within the same transaction. " 
func IsServerErrorBinlogUnsafeNontransAfterTrans(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS)
    return result
}

    
// IsServerErrorMessageAndStatement check mysql error is "%s Statement: %s " 
func IsServerErrorMessageAndStatement(err error) bool {
    result := Isa(err, ER_MESSAGE_AND_STATEMENT)
    return result
}

    
// IsServerErrorSlaveConversionFailed check mysql error is "Column %d of table '%s.%s' cannot be converted from type '%s' to type '%s' " 
func IsServerErrorSlaveConversionFailed(err error) bool {
    result := Isa(err, ER_SLAVE_CONVERSION_FAILED)
    return result
}

    
// IsServerErrorSlaveCantCreateConversion check mysql error is "Can't create conversion table for table '%s.%s' " 
func IsServerErrorSlaveCantCreateConversion(err error) bool {
    result := Isa(err, ER_SLAVE_CANT_CREATE_CONVERSION)
    return result
}

    
// IsServerErrorInsideTransactionPreventsSwitchBinlogFormat check mysql error is "Cannot modify @@session.binlog_format inside a transaction " 
func IsServerErrorInsideTransactionPreventsSwitchBinlogFormat(err error) bool {
    result := Isa(err, ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT)
    return result
}

    
// IsServerErrorPathLength check mysql error is "The path specified for %s is too long. " 
func IsServerErrorPathLength(err error) bool {
    result := Isa(err, ER_PATH_LENGTH)
    return result
}

    
// IsServerErrorWarnDeprecatedSyntaxNoReplacement check mysql error is "'%s' is deprecated and will be removed in a future release. " 
func IsServerErrorWarnDeprecatedSyntaxNoReplacement(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT)
    return result
}

    
// IsServerErrorWrongNativeTableStructure check mysql error is "Native table '%s'.'%s' has the wrong structure " 
func IsServerErrorWrongNativeTableStructure(err error) bool {
    result := Isa(err, ER_WRONG_NATIVE_TABLE_STRUCTURE)
    return result
}

    
// IsServerErrorWrongPerfschemaUsage check mysql error is "Invalid performance_schema usage. " 
func IsServerErrorWrongPerfschemaUsage(err error) bool {
    result := Isa(err, ER_WRONG_PERFSCHEMA_USAGE)
    return result
}

    
// IsServerErrorWarnISSkippedTable check mysql error is "Table '%s'.'%s' was skipped since its definition is being modified by concurrent DDL statement " 
func IsServerErrorWarnISSkippedTable(err error) bool {
    result := Isa(err, ER_WARN_I_S_SKIPPED_TABLE)
    return result
}

    
// IsServerErrorInsideTransactionPreventsSwitchBinlogDirect check mysql error is "Cannot modify @@session.binlog_direct_non_transactional_updates inside a transaction " 
func IsServerErrorInsideTransactionPreventsSwitchBinlogDirect(err error) bool {
    result := Isa(err, ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT)
    return result
}

    
// IsServerErrorStoredFunctionPreventsSwitchBinlogDirect check mysql error is "Cannot change the binlog direct flag inside a stored function or trigger " 
func IsServerErrorStoredFunctionPreventsSwitchBinlogDirect(err error) bool {
    result := Isa(err, ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT)
    return result
}

    
// IsServerErrorSpatialMustHaveGeomCol check mysql error is "A SPATIAL index may only contain a geometrical type column " 
func IsServerErrorSpatialMustHaveGeomCol(err error) bool {
    result := Isa(err, ER_SPATIAL_MUST_HAVE_GEOM_COL)
    return result
}

    
// IsServerErrorTooLongIndexComment check mysql error is "Comment for index '%s' is too long (max = %lu) " 
func IsServerErrorTooLongIndexComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_INDEX_COMMENT)
    return result
}

    
// IsServerErrorLockAborted check mysql error is "Wait on a lock was aborted due to a pending exclusive lock " 
func IsServerErrorLockAborted(err error) bool {
    result := Isa(err, ER_LOCK_ABORTED)
    return result
}

    
// IsServerErrorDataOutOfRange check mysql error is "%s value is out of range in '%s' " 
func IsServerErrorDataOutOfRange(err error) bool {
    result := Isa(err, ER_DATA_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorWrongSpvarTypeInLimit check mysql error is "A variable of a non-integer based type in LIMIT clause " 
func IsServerErrorWrongSpvarTypeInLimit(err error) bool {
    result := Isa(err, ER_WRONG_SPVAR_TYPE_IN_LIMIT)
    return result
}

    
// IsServerErrorBinlogUnsafeMultipleEnginesAndSelfLoggingEngine check mysql error is "Mixing self-logging and non-self-logging engines in a statement is unsafe. " 
func IsServerErrorBinlogUnsafeMultipleEnginesAndSelfLoggingEngine(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeMixedStatement check mysql error is "Statement accesses nontransactional table as well as transactional or temporary table, and writes to any of them. " 
func IsServerErrorBinlogUnsafeMixedStatement(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_MIXED_STATEMENT)
    return result
}

    
// IsServerErrorInsideTransactionPreventsSwitchSqlLogBin check mysql error is "Cannot modify @@session.sql_log_bin inside a transaction " 
func IsServerErrorInsideTransactionPreventsSwitchSqlLogBin(err error) bool {
    result := Isa(err, ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN)
    return result
}

    
// IsServerErrorStoredFunctionPreventsSwitchSqlLogBin check mysql error is "Cannot change the sql_log_bin inside a stored function or trigger " 
func IsServerErrorStoredFunctionPreventsSwitchSqlLogBin(err error) bool {
    result := Isa(err, ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN)
    return result
}

    
// IsServerErrorFailedReadFromParFile check mysql error is "Failed to read from the .par file " 
func IsServerErrorFailedReadFromParFile(err error) bool {
    result := Isa(err, ER_FAILED_READ_FROM_PAR_FILE)
    return result
}

    
// IsServerErrorValuesIsNotIntTypeError check mysql error is "VALUES value for partition '%s' must have type INT " 
func IsServerErrorValuesIsNotIntTypeError(err error) bool {
    result := Isa(err, ER_VALUES_IS_NOT_INT_TYPE_ERROR)
    return result
}

    
// IsServerErrorAccessDeniedNoPasswordError check mysql error is "Access denied for user '%s'@'%s' " 
func IsServerErrorAccessDeniedNoPasswordError(err error) bool {
    result := Isa(err, ER_ACCESS_DENIED_NO_PASSWORD_ERROR)
    return result
}

    
// IsServerErrorSetPasswordAuthPlugin check mysql error is "SET PASSWORD has no significance for users authenticating via plugins " 
func IsServerErrorSetPasswordAuthPlugin(err error) bool {
    result := Isa(err, ER_SET_PASSWORD_AUTH_PLUGIN)
    return result
}

    
// IsServerErrorGrantPluginUserExists check mysql error is "GRANT with IDENTIFIED WITH is illegal because the user %-.*s already exists " 
func IsServerErrorGrantPluginUserExists(err error) bool {
    result := Isa(err, ER_GRANT_PLUGIN_USER_EXISTS)
    return result
}

    
// IsServerErrorTruncateIllegalFk check mysql error is "Cannot truncate a table referenced in a foreign key constraint (%s) " 
func IsServerErrorTruncateIllegalFk(err error) bool {
    result := Isa(err, ER_TRUNCATE_ILLEGAL_FK)
    return result
}

    
// IsServerErrorPluginIsPermanent check mysql error is "Plugin '%s' is force_plus_permanent and can not be unloaded " 
func IsServerErrorPluginIsPermanent(err error) bool {
    result := Isa(err, ER_PLUGIN_IS_PERMANENT)
    return result
}

    
// IsServerErrorSlaveHeartbeatValueOutOfRangeMin check mysql error is "The requested value for the heartbeat period is less than 1 millisecond. The value is reset to 0, meaning that heartbeating will effectively be disabled. " 
func IsServerErrorSlaveHeartbeatValueOutOfRangeMin(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN)
    return result
}

    
// IsServerErrorSlaveHeartbeatValueOutOfRangeMax check mysql error is "The requested value for the heartbeat period exceeds the value of `slave_net_timeout' seconds. A sensible value for the period should be less than the timeout. " 
func IsServerErrorSlaveHeartbeatValueOutOfRangeMax(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX)
    return result
}

    
// IsServerErrorStmtCacheFull check mysql error is "Multi-row statements required more than 'max_binlog_stmt_cache_size' bytes of storage" 
func IsServerErrorStmtCacheFull(err error) bool {
    result := Isa(err, ER_STMT_CACHE_FULL)
    return result
}

    
// IsServerErrorMultiUpdateKeyConflict check mysql error is "Primary key/partition key update is not allowed since the table is updated both as '%s' and '%s'. " 
func IsServerErrorMultiUpdateKeyConflict(err error) bool {
    result := Isa(err, ER_MULTI_UPDATE_KEY_CONFLICT)
    return result
}

    
// IsServerErrorTableNeedsRebuild check mysql error is "Table rebuild required. Please do "ALTER TABLE `%s` FORCE" or dump/reload to fix it! " 
func IsServerErrorTableNeedsRebuild(err error) bool {
    result := Isa(err, ER_TABLE_NEEDS_REBUILD)
    return result
}

    
// IsWarnOptionBelowLimit check mysql error is "The value of '%s' should be no less than the value of '%s' " 
func IsWarnOptionBelowLimit(err error) bool {
    result := Isa(err, WARN_OPTION_BELOW_LIMIT)
    return result
}

    
// IsServerErrorIndexColumnTooLong check mysql error is "Index column size too large. The maximum column size is %lu bytes. " 
func IsServerErrorIndexColumnTooLong(err error) bool {
    result := Isa(err, ER_INDEX_COLUMN_TOO_LONG)
    return result
}

    
// IsServerErrorErrorInTriggerBody check mysql error is "Trigger '%s' has an error in its body: '%s' " 
func IsServerErrorErrorInTriggerBody(err error) bool {
    result := Isa(err, ER_ERROR_IN_TRIGGER_BODY)
    return result
}

    
// IsServerErrorErrorInUnknownTriggerBody check mysql error is "Unknown trigger has an error in its body: '%s' " 
func IsServerErrorErrorInUnknownTriggerBody(err error) bool {
    result := Isa(err, ER_ERROR_IN_UNKNOWN_TRIGGER_BODY)
    return result
}

    
// IsServerErrorIndexCorrupt check mysql error is "Index %s is corrupted " 
func IsServerErrorIndexCorrupt(err error) bool {
    result := Isa(err, ER_INDEX_CORRUPT)
    return result
}

    
// IsServerErrorUndoRecordTooBig check mysql error is "Undo log record is too big. " 
func IsServerErrorUndoRecordTooBig(err error) bool {
    result := Isa(err, ER_UNDO_RECORD_TOO_BIG)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertIgnoreSelect check mysql error is "INSERT IGNORE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeInsertIgnoreSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertSelectUpdate check mysql error is "INSERT... SELECT... ON DUPLICATE KEY UPDATE is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are updated. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeInsertSelectUpdate(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE)
    return result
}

    
// IsServerErrorBinlogUnsafeReplaceSelect check mysql error is "REPLACE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeReplaceSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_REPLACE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeCreateIgnoreSelect check mysql error is "CREATE... IGNORE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeCreateIgnoreSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeCreateReplaceSelect check mysql error is "CREATE... REPLACE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeCreateReplaceSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeUpdateIgnore check mysql error is "UPDATE IGNORE is unsafe because the order in which rows are updated determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeUpdateIgnore(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_UPDATE_IGNORE)
    return result
}

    
// IsServerErrorPluginNoUninstall check mysql error is "Plugin '%s' is marked as not dynamically uninstallable. You have to stop the server to uninstall it. " 
func IsServerErrorPluginNoUninstall(err error) bool {
    result := Isa(err, ER_PLUGIN_NO_UNINSTALL)
    return result
}

    
// IsServerErrorPluginNoInstall check mysql error is "Plugin '%s' is marked as not dynamically installable. You have to stop the server to install it. " 
func IsServerErrorPluginNoInstall(err error) bool {
    result := Isa(err, ER_PLUGIN_NO_INSTALL)
    return result
}

    
// IsServerErrorBinlogUnsafeWriteAutoincSelect check mysql error is "Statements writing to a table with an auto-increment column after selecting from another table are unsafe because the order in which rows are retrieved determines what (if any) rows will be written. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeWriteAutoincSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeCreateSelectAutoinc check mysql error is "CREATE TABLE... SELECT... on a table with an auto-increment column is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are inserted. This order cannot be predicted and may differ on master and the slave. " 
func IsServerErrorBinlogUnsafeCreateSelectAutoinc(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertTwoKeys check mysql error is "INSERT... ON DUPLICATE KEY UPDATE on a table with more than one UNIQUE KEY is unsafe " 
func IsServerErrorBinlogUnsafeInsertTwoKeys(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_TWO_KEYS)
    return result
}

    
// IsServerErrorTableInFkCheck check mysql error is "Table is being used in foreign key check. " 
func IsServerErrorTableInFkCheck(err error) bool {
    result := Isa(err, ER_TABLE_IN_FK_CHECK)
    return result
}

    
// IsServerErrorUnsupportedEngine check mysql error is "Storage engine '%s' does not support system tables. [%s.%s] " 
func IsServerErrorUnsupportedEngine(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeAutoincNotFirst check mysql error is "INSERT into autoincrement field which is not the first part in the composed primary key is unsafe. " 
func IsServerErrorBinlogUnsafeAutoincNotFirst(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST)
    return result
}

    
// IsServerErrorCannotLoadFromTableV2 check mysql error is "Cannot load from %s.%s. The table is probably corrupted " 
func IsServerErrorCannotLoadFromTableV2(err error) bool {
    result := Isa(err, ER_CANNOT_LOAD_FROM_TABLE_V2)
    return result
}

    
// IsServerErrorMasterDelayValueOutOfRange check mysql error is "The requested value %s for the master delay exceeds the maximum %u " 
func IsServerErrorMasterDelayValueOutOfRange(err error) bool {
    result := Isa(err, ER_MASTER_DELAY_VALUE_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorOnlyFdAndRbrEventsAllowedInBinlogStatement check mysql error is "Only Format_description_log_event and row events are allowed in BINLOG statements (but %s was provided) " 
func IsServerErrorOnlyFdAndRbrEventsAllowedInBinlogStatement(err error) bool {
    result := Isa(err, ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT)
    return result
}

    
// IsServerErrorPartitionExchangeDifferentOption check mysql error is "Non matching attribute '%s' between partition and table " 
func IsServerErrorPartitionExchangeDifferentOption(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_DIFFERENT_OPTION)
    return result
}

    
// IsServerErrorPartitionExchangePartTable check mysql error is "Table to exchange with partition is partitioned: '%s' " 
func IsServerErrorPartitionExchangePartTable(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_PART_TABLE)
    return result
}

    
// IsServerErrorPartitionExchangeTempTable check mysql error is "Table to exchange with partition is temporary: '%s' " 
func IsServerErrorPartitionExchangeTempTable(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_TEMP_TABLE)
    return result
}

    
// IsServerErrorPartitionInsteadOfSubpartition check mysql error is "Subpartitioned table, use subpartition instead of partition " 
func IsServerErrorPartitionInsteadOfSubpartition(err error) bool {
    result := Isa(err, ER_PARTITION_INSTEAD_OF_SUBPARTITION)
    return result
}

    
// IsServerErrorUnknownPartition check mysql error is "Unknown partition '%s' in table '%s' " 
func IsServerErrorUnknownPartition(err error) bool {
    result := Isa(err, ER_UNKNOWN_PARTITION)
    return result
}

    
// IsServerErrorTablesDifferentMetadata check mysql error is "Tables have different definitions " 
func IsServerErrorTablesDifferentMetadata(err error) bool {
    result := Isa(err, ER_TABLES_DIFFERENT_METADATA)
    return result
}

    
// IsServerErrorRowDoesNotMatchPartition check mysql error is "Found a row that does not match the partition " 
func IsServerErrorRowDoesNotMatchPartition(err error) bool {
    result := Isa(err, ER_ROW_DOES_NOT_MATCH_PARTITION)
    return result
}

    
// IsServerErrorBinlogCacheSizeGreaterThanMax check mysql error is "Option binlog_cache_size (%lu) is greater than max_binlog_cache_size (%lu)" 
func IsServerErrorBinlogCacheSizeGreaterThanMax(err error) bool {
    result := Isa(err, ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX)
    return result
}

    
// IsServerErrorWarnIndexNotApplicable check mysql error is "Cannot use %s access on index '%s' due to type or collation conversion on field '%s' " 
func IsServerErrorWarnIndexNotApplicable(err error) bool {
    result := Isa(err, ER_WARN_INDEX_NOT_APPLICABLE)
    return result
}

    
// IsServerErrorPartitionExchangeForeignKey check mysql error is "Table to exchange with partition has foreign key references: '%s' " 
func IsServerErrorPartitionExchangeForeignKey(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_FOREIGN_KEY)
    return result
}

    
// IsServerErrorNoSuchKeyValue check mysql error is "Key value '%s' was not found in table '%s.%s' " 
func IsServerErrorNoSuchKeyValue(err error) bool {
    result := Isa(err, ER_NO_SUCH_KEY_VALUE)
    return result
}

    
// IsServerErrorRplInfoDataTooLong check mysql error is "Data for column '%s' too long " 
func IsServerErrorRplInfoDataTooLong(err error) bool {
    result := Isa(err, ER_RPL_INFO_DATA_TOO_LONG)
    return result
}

    
// IsServerErrorNetworkReadEventChecksumFailure check mysql error is "Replication event checksum verification failed while reading from network. " 
func IsServerErrorNetworkReadEventChecksumFailure(err error) bool {
    result := Isa(err, ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE)
    return result
}

    
// IsServerErrorBinlogReadEventChecksumFailure check mysql error is "Replication event checksum verification failed while reading from a log file. " 
func IsServerErrorBinlogReadEventChecksumFailure(err error) bool {
    result := Isa(err, ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE)
    return result
}

    
// IsServerErrorBinlogStmtCacheSizeGreaterThanMax check mysql error is "Option binlog_stmt_cache_size (%lu) is greater than max_binlog_stmt_cache_size (%lu)" 
func IsServerErrorBinlogStmtCacheSizeGreaterThanMax(err error) bool {
    result := Isa(err, ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX)
    return result
}

    
// IsServerErrorCantUpdateTableInCreateTableSelect check mysql error is "Can't update table '%s' while '%s' is being created. " 
func IsServerErrorCantUpdateTableInCreateTableSelect(err error) bool {
    result := Isa(err, ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT)
    return result
}

    
// IsServerErrorPartitionClauseOnNonpartitioned check mysql error is "PARTITION () clause on non partitioned table " 
func IsServerErrorPartitionClauseOnNonpartitioned(err error) bool {
    result := Isa(err, ER_PARTITION_CLAUSE_ON_NONPARTITIONED)
    return result
}

    
// IsServerErrorRowDoesNotMatchGivenPartitionSet check mysql error is "Found a row not matching the given partition set " 
func IsServerErrorRowDoesNotMatchGivenPartitionSet(err error) bool {
    result := Isa(err, ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET)
    return result
}

    
// IsServerErrorNoSuchPartitionUnused check mysql error is "partition '%s' doesn't exist " 
func IsServerErrorNoSuchPartitionUnused(err error) bool {
    result := Isa(err, ER_NO_SUCH_PARTITION__UNUSED)
    return result
}

    
// IsServerErrorChangeRplInfoRepositoryFailure check mysql error is "Failure while changing the type of replication repository: %s. " 
func IsServerErrorChangeRplInfoRepositoryFailure(err error) bool {
    result := Isa(err, ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE)
    return result
}

    
// IsServerErrorWarningNotCompleteRollbackWithCreatedTempTable check mysql error is "The creation of some temporary tables could not be rolled back. " 
func IsServerErrorWarningNotCompleteRollbackWithCreatedTempTable(err error) bool {
    result := Isa(err, ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE)
    return result
}

    
// IsServerErrorWarningNotCompleteRollbackWithDroppedTempTable check mysql error is "Some temporary tables were dropped, but these operations could not be rolled back. " 
func IsServerErrorWarningNotCompleteRollbackWithDroppedTempTable(err error) bool {
    result := Isa(err, ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE)
    return result
}

    
// IsServerErrorMtsFeatureIsNotSupported check mysql error is "%s is not supported in multi-threaded slave mode. %s " 
func IsServerErrorMtsFeatureIsNotSupported(err error) bool {
    result := Isa(err, ER_MTS_FEATURE_IS_NOT_SUPPORTED)
    return result
}

    
// IsServerErrorMtsUpdatedDbsGreaterMax check mysql error is "The number of modified databases exceeds the maximum %d" 
func IsServerErrorMtsUpdatedDbsGreaterMax(err error) bool {
    result := Isa(err, ER_MTS_UPDATED_DBS_GREATER_MAX)
    return result
}

    
// IsServerErrorMtsCantParallel check mysql error is "Cannot execute the current event group in the parallel mode. Encountered event %s, relay-log name %s, position %s which prevents execution of this event group in parallel mode. Reason: %s. " 
func IsServerErrorMtsCantParallel(err error) bool {
    result := Isa(err, ER_MTS_CANT_PARALLEL)
    return result
}

    
// IsServerErrorMtsInconsistentData check mysql error is "%s " 
func IsServerErrorMtsInconsistentData(err error) bool {
    result := Isa(err, ER_MTS_INCONSISTENT_DATA)
    return result
}

    
// IsServerErrorFulltextNotSupportedWithPartitioning check mysql error is "FULLTEXT index is not supported for partitioned tables. " 
func IsServerErrorFulltextNotSupportedWithPartitioning(err error) bool {
    result := Isa(err, ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING)
    return result
}

    
// IsServerErrorDaInvalidConditionNumber check mysql error is "Invalid condition number " 
func IsServerErrorDaInvalidConditionNumber(err error) bool {
    result := Isa(err, ER_DA_INVALID_CONDITION_NUMBER)
    return result
}

    
// IsServerErrorInsecurePlainText check mysql error is "Sending passwords in plain text without SSL/TLS is extremely insecure. " 
func IsServerErrorInsecurePlainText(err error) bool {
    result := Isa(err, ER_INSECURE_PLAIN_TEXT)
    return result
}

    
// IsServerErrorInsecureChangeMaster check mysql error is "Storing MySQL user name or password information in the master info repository is not secure and is therefore not recommended. Please consider using the USER and PASSWORD connection options for START SLAVE" 
func IsServerErrorInsecureChangeMaster(err error) bool {
    result := Isa(err, ER_INSECURE_CHANGE_MASTER)
    return result
}

    
// IsServerErrorForeignDuplicateKeyWithChildInfo check mysql error is "Foreign key constraint for table '%s', record '%s' would lead to a duplicate entry in table '%s', key '%s' " 
func IsServerErrorForeignDuplicateKeyWithChildInfo(err error) bool {
    result := Isa(err, ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO)
    return result
}

    
// IsServerErrorForeignDuplicateKeyWithoutChildInfo check mysql error is "Foreign key constraint for table '%s', record '%s' would lead to a duplicate entry in a child table " 
func IsServerErrorForeignDuplicateKeyWithoutChildInfo(err error) bool {
    result := Isa(err, ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO)
    return result
}

    
// IsServerErrorSqlthreadWithSecureSlave check mysql error is "Setting authentication options is not possible when only the Slave SQL Thread is being started. " 
func IsServerErrorSqlthreadWithSecureSlave(err error) bool {
    result := Isa(err, ER_SQLTHREAD_WITH_SECURE_SLAVE)
    return result
}

    
// IsServerErrorTableHasNoFt check mysql error is "The table does not have FULLTEXT index to support this query " 
func IsServerErrorTableHasNoFt(err error) bool {
    result := Isa(err, ER_TABLE_HAS_NO_FT)
    return result
}

    
// IsServerErrorVariableNotSettableInSfOrTrigger check mysql error is "The system variable %s cannot be set in stored functions or triggers. " 
func IsServerErrorVariableNotSettableInSfOrTrigger(err error) bool {
    result := Isa(err, ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER)
    return result
}

    
// IsServerErrorVariableNotSettableInTransaction check mysql error is "The system variable %s cannot be set when there is an ongoing transaction. " 
func IsServerErrorVariableNotSettableInTransaction(err error) bool {
    result := Isa(err, ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION)
    return result
}

    
// IsServerErrorGtidNextIsNotInGtidNextList check mysql error is "The system variable @@SESSION.GTID_NEXT has the value %s, which is not listed in @@SESSION.GTID_NEXT_LIST. " 
func IsServerErrorGtidNextIsNotInGtidNextList(err error) bool {
    result := Isa(err, ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST)
    return result
}

    
// IsServerErrorCantChangeGtidNextInTransactionWhenGtidNextListIsNull check mysql error is "The system variable @@SESSION.GTID_NEXT cannot change inside a transaction." 
func IsServerErrorCantChangeGtidNextInTransactionWhenGtidNextListIsNull(err error) bool {
    result := Isa(err, ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL)
    return result
}

    
// IsServerErrorCantChangeGtidNextInTransaction check mysql error is "The system variable @@SESSION.GTID_NEXT cannot change inside a transaction." 
func IsServerErrorCantChangeGtidNextInTransaction(err error) bool {
    result := Isa(err, ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION)
    return result
}

    
// IsServerErrorSetStatementCannotInvokeFunction check mysql error is "The statement 'SET %s' cannot invoke a stored function. " 
func IsServerErrorSetStatementCannotInvokeFunction(err error) bool {
    result := Isa(err, ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION)
    return result
}

    
// IsServerErrorGtidNextCantBeAutomaticIfGtidNextListIsNonNull check mysql error is "The system variable @@SESSION.GTID_NEXT cannot be 'AUTOMATIC' when @@SESSION.GTID_NEXT_LIST is non-NULL. " 
func IsServerErrorGtidNextCantBeAutomaticIfGtidNextListIsNonNull(err error) bool {
    result := Isa(err, ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL)
    return result
}

    
// IsServerErrorSkippingLoggedTransaction check mysql error is "Skipping transaction %s because it has already been executed and logged. " 
func IsServerErrorSkippingLoggedTransaction(err error) bool {
    result := Isa(err, ER_SKIPPING_LOGGED_TRANSACTION)
    return result
}

    
// IsServerErrorMalformedGtidSetSpecification check mysql error is "Malformed GTID set specification '%s'. " 
func IsServerErrorMalformedGtidSetSpecification(err error) bool {
    result := Isa(err, ER_MALFORMED_GTID_SET_SPECIFICATION)
    return result
}

    
// IsServerErrorMalformedGtidSetEncoding check mysql error is "Malformed GTID set encoding. " 
func IsServerErrorMalformedGtidSetEncoding(err error) bool {
    result := Isa(err, ER_MALFORMED_GTID_SET_ENCODING)
    return result
}

    
// IsServerErrorMalformedGtidSpecification check mysql error is "Malformed GTID specification '%s'. " 
func IsServerErrorMalformedGtidSpecification(err error) bool {
    result := Isa(err, ER_MALFORMED_GTID_SPECIFICATION)
    return result
}

    
// IsServerErrorGnoExhausted check mysql error is "Impossible to generate Global Transaction Identifier: the integer component reached the maximal value. Restart the server with a new server_uuid. " 
func IsServerErrorGnoExhausted(err error) bool {
    result := Isa(err, ER_GNO_EXHAUSTED)
    return result
}

    
// IsServerErrorBadSlaveAutoPosition check mysql error is "Parameters MASTER_LOG_FILE, MASTER_LOG_POS, RELAY_LOG_FILE and RELAY_LOG_POS cannot be set when MASTER_AUTO_POSITION is active. " 
func IsServerErrorBadSlaveAutoPosition(err error) bool {
    result := Isa(err, ER_BAD_SLAVE_AUTO_POSITION)
    return result
}

    
// IsServerErrorAutoPositionRequiresGtidModeOn check mysql error is "CHANGE MASTER TO MASTER_AUTO_POSITION = 1 can only be executed when @@GLOBAL.GTID_MODE = ON." 
func IsServerErrorAutoPositionRequiresGtidModeOn(err error) bool {
    result := Isa(err, ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON)
    return result
}

    
// IsServerErrorAutoPositionRequiresGtidModeNotOff check mysql error is "CHANGE MASTER TO MASTER_AUTO_POSITION = 1 cannot be executed because @@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorAutoPositionRequiresGtidModeNotOff(err error) bool {
    result := Isa(err, ER_AUTO_POSITION_REQUIRES_GTID_MODE_NOT_OFF)
    return result
}

    
// IsServerErrorCantDoImplicitCommitInTrxWhenGtidNextIsSet check mysql error is "Cannot execute statements with implicit commit inside a transaction when @@SESSION.GTID_NEXT == 'UUID:NUMBER'. " 
func IsServerErrorCantDoImplicitCommitInTrxWhenGtidNextIsSet(err error) bool {
    result := Isa(err, ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET)
    return result
}

    
// IsServerErrorGtidMode2Or3RequiresEnforceGtidConsistencyOn check mysql error is "@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_2 requires @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1." 
func IsServerErrorGtidMode2Or3RequiresEnforceGtidConsistencyOn(err error) bool {
    result := Isa(err, ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON)
    return result
}

    
// IsServerErrorGtidModeOnRequiresEnforceGtidConsistencyOn check mysql error is "GTID_MODE = ON requires ENFORCE_GTID_CONSISTENCY = ON." 
func IsServerErrorGtidModeOnRequiresEnforceGtidConsistencyOn(err error) bool {
    result := Isa(err, ER_GTID_MODE_ON_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON)
    return result
}

    
// IsServerErrorGtidModeRequiresBinlog check mysql error is "@@GLOBAL.GTID_MODE = ON or ON_PERMISSIVE or OFF_PERMISSIVE requires --log-bin and --log-slave-updates. " 
func IsServerErrorGtidModeRequiresBinlog(err error) bool {
    result := Isa(err, ER_GTID_MODE_REQUIRES_BINLOG)
    return result
}

    
// IsServerErrorCantSetGtidNextToGtidWhenGtidModeIsOff check mysql error is "@@SESSION.GTID_NEXT cannot be set to UUID:NUMBER when @@GLOBAL.GTID_MODE = OFF. " 
func IsServerErrorCantSetGtidNextToGtidWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorCantSetGtidNextToAnonymousWhenGtidModeIsOn check mysql error is "@@SESSION.GTID_NEXT cannot be set to ANONYMOUS when @@GLOBAL.GTID_MODE = ON. " 
func IsServerErrorCantSetGtidNextToAnonymousWhenGtidModeIsOn(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON)
    return result
}

    
// IsServerErrorCantSetGtidNextListToNonNullWhenGtidModeIsOff check mysql error is "@@SESSION.GTID_NEXT_LIST cannot be set to a non-NULL value when @@GLOBAL.GTID_MODE = OFF. " 
func IsServerErrorCantSetGtidNextListToNonNullWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorFoundGtidEventWhenGtidModeIsOff check mysql error is "Found a Gtid_log_event or Previous_gtids_log_event when @@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorFoundGtidEventWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorFoundGtidEventWhenGtidModeIsOffUnused check mysql error is "Found a Gtid_log_event when @@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorFoundGtidEventWhenGtidModeIsOffUnused(err error) bool {
    result := Isa(err, ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF__UNUSED)
    return result
}

    
// IsServerErrorGtidUnsafeNonTransactionalTable check mysql error is "Statement violates GTID consistency: Updates to non-transactional tables can only be done in either autocommitted statements or single-statement transactions, and never in the same statement as updates to transactional tables. " 
func IsServerErrorGtidUnsafeNonTransactionalTable(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE)
    return result
}

    
// IsServerErrorGtidUnsafeCreateSelect check mysql error is "Statement violates GTID consistency: CREATE TABLE ... SELECT. " 
func IsServerErrorGtidUnsafeCreateSelect(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_CREATE_SELECT)
    return result
}

    
// IsServerErrorGtidUnsafeCreateDropTemporaryTableInTransaction check mysql error is "Statement violates GTID consistency: CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can only be executed outside transactional context. These statements are also not allowed in a function or trigger because functions and triggers are also considered to be multi-statement transactions. " 
func IsServerErrorGtidUnsafeCreateDropTemporaryTableInTransaction(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTION)
    return result
}

    
// IsServerErrorGtidModeCanOnlyChangeOneStepAtATime check mysql error is "The value of @@GLOBAL.GTID_MODE can only be changed one step at a time: OFF <-> OFF_PERMISSIVE <-> ON_PERMISSIVE <-> ON. Also note that this value must be stepped up or down simultaneously on all servers. See the Manual for instructions. " 
func IsServerErrorGtidModeCanOnlyChangeOneStepAtATime(err error) bool {
    result := Isa(err, ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME)
    return result
}

    
// IsServerErrorMasterHasPurgedRequiredGtids check mysql error is "The slave is connecting using CHANGE MASTER TO MASTER_AUTO_POSITION = 1, but the master has purged binary logs containing GTIDs that the slave requires. Replicate the missing transactions from elsewhere, or provision a new slave from backup. Consider increasing the master's binary log expiration period. %s. " 
func IsServerErrorMasterHasPurgedRequiredGtids(err error) bool {
    result := Isa(err, ER_MASTER_HAS_PURGED_REQUIRED_GTIDS)
    return result
}

    
// IsServerErrorCantSetGtidNextWhenOwningGtid check mysql error is "@@SESSION.GTID_NEXT cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK. " 
func IsServerErrorCantSetGtidNextWhenOwningGtid(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID)
    return result
}

    
// IsServerErrorUnknownExplainFormat check mysql error is "Unknown EXPLAIN format name: '%s' " 
func IsServerErrorUnknownExplainFormat(err error) bool {
    result := Isa(err, ER_UNKNOWN_EXPLAIN_FORMAT)
    return result
}

    
// IsServerErrorCantExecuteInReadOnlyTransaction check mysql error is "Cannot execute statement in a READ ONLY transaction. " 
func IsServerErrorCantExecuteInReadOnlyTransaction(err error) bool {
    result := Isa(err, ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION)
    return result
}

    
// IsServerErrorTooLongTablePartitionComment check mysql error is "Comment for table partition '%s' is too long (max = %lu) " 
func IsServerErrorTooLongTablePartitionComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_TABLE_PARTITION_COMMENT)
    return result
}

    
// IsServerErrorSlaveConfiguration check mysql error is "Slave is not configured or failed to initialize properly. You must at least set --server-id to enable either a master or a slave. Additional error messages can be found in the MySQL error log. " 
func IsServerErrorSlaveConfiguration(err error) bool {
    result := Isa(err, ER_SLAVE_CONFIGURATION)
    return result
}

    
// IsServerErrorInnodbFtLimit check mysql error is "InnoDB presently supports one FULLTEXT index creation at a time " 
func IsServerErrorInnodbFtLimit(err error) bool {
    result := Isa(err, ER_INNODB_FT_LIMIT)
    return result
}

    
// IsServerErrorInnodbNoFtTempTable check mysql error is "Cannot create FULLTEXT index on temporary InnoDB table " 
func IsServerErrorInnodbNoFtTempTable(err error) bool {
    result := Isa(err, ER_INNODB_NO_FT_TEMP_TABLE)
    return result
}

    
// IsServerErrorInnodbFtWrongDocidColumn check mysql error is "Column '%s' is of wrong type for an InnoDB FULLTEXT index " 
func IsServerErrorInnodbFtWrongDocidColumn(err error) bool {
    result := Isa(err, ER_INNODB_FT_WRONG_DOCID_COLUMN)
    return result
}

    
// IsServerErrorInnodbFtWrongDocidIndex check mysql error is "Index '%s' is of wrong type for an InnoDB FULLTEXT index " 
func IsServerErrorInnodbFtWrongDocidIndex(err error) bool {
    result := Isa(err, ER_INNODB_FT_WRONG_DOCID_INDEX)
    return result
}

    
// IsServerErrorInnodbOnlineLogTooBig check mysql error is "Creating index '%s' required more than 'innodb_online_alter_log_max_size' bytes of modification log. Please try again. " 
func IsServerErrorInnodbOnlineLogTooBig(err error) bool {
    result := Isa(err, ER_INNODB_ONLINE_LOG_TOO_BIG)
    return result
}

    
// IsServerErrorUnknownAlterAlgorithm check mysql error is "Unknown ALGORITHM '%s' " 
func IsServerErrorUnknownAlterAlgorithm(err error) bool {
    result := Isa(err, ER_UNKNOWN_ALTER_ALGORITHM)
    return result
}

    
// IsServerErrorUnknownAlterLock check mysql error is "Unknown LOCK type '%s' " 
func IsServerErrorUnknownAlterLock(err error) bool {
    result := Isa(err, ER_UNKNOWN_ALTER_LOCK)
    return result
}

    
// IsServerErrorMtsChangeMasterCantRunWithGaps check mysql error is "CHANGE MASTER cannot be executed when the slave was stopped with an error or killed in MTS mode. Consider using RESET SLAVE or START SLAVE UNTIL. " 
func IsServerErrorMtsChangeMasterCantRunWithGaps(err error) bool {
    result := Isa(err, ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS)
    return result
}

    
// IsServerErrorMtsRecoveryFailure check mysql error is "Cannot recover after SLAVE errored out in parallel execution mode. Additional error messages can be found in the MySQL error log. " 
func IsServerErrorMtsRecoveryFailure(err error) bool {
    result := Isa(err, ER_MTS_RECOVERY_FAILURE)
    return result
}

    
// IsServerErrorMtsResetWorkers check mysql error is "Cannot clean up worker info tables. Additional error messages can be found in the MySQL error log. " 
func IsServerErrorMtsResetWorkers(err error) bool {
    result := Isa(err, ER_MTS_RESET_WORKERS)
    return result
}

    
// IsServerErrorColCountDoesntMatchCorruptedV2 check mysql error is "Column count of %s.%s is wrong. Expected %d, found %d. The table is probably corrupted " 
func IsServerErrorColCountDoesntMatchCorruptedV2(err error) bool {
    result := Isa(err, ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2)
    return result
}

    
// IsServerErrorSlaveSilentRetryTransaction check mysql error is "Slave must silently retry current transaction " 
func IsServerErrorSlaveSilentRetryTransaction(err error) bool {
    result := Isa(err, ER_SLAVE_SILENT_RETRY_TRANSACTION)
    return result
}

    
// IsServerErrorDiscardFkChecksRunning check mysql error is "There is a foreign key check running on table '%s'. Cannot discard the table. " 
func IsServerErrorDiscardFkChecksRunning(err error) bool {
    result := Isa(err, ER_DISCARD_FK_CHECKS_RUNNING)
    return result
}

    
// IsServerErrorTableSchemaMismatch check mysql error is "Schema mismatch (%s) " 
func IsServerErrorTableSchemaMismatch(err error) bool {
    result := Isa(err, ER_TABLE_SCHEMA_MISMATCH)
    return result
}

    
// IsServerErrorTableInSystemTablespace check mysql error is "Table '%s' in system tablespace " 
func IsServerErrorTableInSystemTablespace(err error) bool {
    result := Isa(err, ER_TABLE_IN_SYSTEM_TABLESPACE)
    return result
}

    
// IsServerErrorIoReadError check mysql error is "IO Read error: (%lu, %s) %s " 
func IsServerErrorIoReadError(err error) bool {
    result := Isa(err, ER_IO_READ_ERROR)
    return result
}

    
// IsServerErrorIoWriteError check mysql error is "IO Write error: (%lu, %s) %s " 
func IsServerErrorIoWriteError(err error) bool {
    result := Isa(err, ER_IO_WRITE_ERROR)
    return result
}

    
// IsServerErrorTablespaceMissing check mysql error is "Tablespace is missing for table %s. " 
func IsServerErrorTablespaceMissing(err error) bool {
    result := Isa(err, ER_TABLESPACE_MISSING)
    return result
}

    
// IsServerErrorTablespaceExists check mysql error is "Tablespace '%s' exists. " 
func IsServerErrorTablespaceExists(err error) bool {
    result := Isa(err, ER_TABLESPACE_EXISTS)
    return result
}

    
// IsServerErrorTablespaceDiscarded check mysql error is "Tablespace has been discarded for table '%s' " 
func IsServerErrorTablespaceDiscarded(err error) bool {
    result := Isa(err, ER_TABLESPACE_DISCARDED)
    return result
}

    
// IsServerErrorInternalError check mysql error is "Internal error: %s " 
func IsServerErrorInternalError(err error) bool {
    result := Isa(err, ER_INTERNAL_ERROR)
    return result
}

    
// IsServerErrorInnodbImportError check mysql error is "ALTER TABLE %s IMPORT TABLESPACE failed with error %lu : '%s' " 
func IsServerErrorInnodbImportError(err error) bool {
    result := Isa(err, ER_INNODB_IMPORT_ERROR)
    return result
}

    
// IsServerErrorInnodbIndexCorrupt check mysql error is "Index corrupt: %s " 
func IsServerErrorInnodbIndexCorrupt(err error) bool {
    result := Isa(err, ER_INNODB_INDEX_CORRUPT)
    return result
}

    
// IsServerErrorInvalidYearColumnLength check mysql error is "Supports only YEAR or YEAR(4) column. " 
func IsServerErrorInvalidYearColumnLength(err error) bool {
    result := Isa(err, ER_INVALID_YEAR_COLUMN_LENGTH)
    return result
}

    
// IsServerErrorNotValidPassword check mysql error is "Your password does not satisfy the current policy requirements " 
func IsServerErrorNotValidPassword(err error) bool {
    result := Isa(err, ER_NOT_VALID_PASSWORD)
    return result
}

    
// IsServerErrorMustChangePassword check mysql error is "You must reset your password using ALTER USER statement before executing this statement. " 
func IsServerErrorMustChangePassword(err error) bool {
    result := Isa(err, ER_MUST_CHANGE_PASSWORD)
    return result
}

    
// IsServerErrorFkNoIndexChild check mysql error is "Failed to add the foreign key constaint. Missing index for constraint '%s' in the foreign table '%s' " 
func IsServerErrorFkNoIndexChild(err error) bool {
    result := Isa(err, ER_FK_NO_INDEX_CHILD)
    return result
}

    
// IsServerErrorFkNoIndexParent check mysql error is "Failed to add the foreign key constaint. Missing index for constraint '%s' in the referenced table '%s' " 
func IsServerErrorFkNoIndexParent(err error) bool {
    result := Isa(err, ER_FK_NO_INDEX_PARENT)
    return result
}

    
// IsServerErrorFkFailAddSystem check mysql error is "Failed to add the foreign key constraint '%s' to system tables " 
func IsServerErrorFkFailAddSystem(err error) bool {
    result := Isa(err, ER_FK_FAIL_ADD_SYSTEM)
    return result
}

    
// IsServerErrorFkCannotOpenParent check mysql error is "Failed to open the referenced table '%s' " 
func IsServerErrorFkCannotOpenParent(err error) bool {
    result := Isa(err, ER_FK_CANNOT_OPEN_PARENT)
    return result
}

    
// IsServerErrorFkIncorrectOption check mysql error is "Failed to add the foreign key constraint on table '%s'. Incorrect options in FOREIGN KEY constraint '%s' " 
func IsServerErrorFkIncorrectOption(err error) bool {
    result := Isa(err, ER_FK_INCORRECT_OPTION)
    return result
}

    
// IsServerErrorFkDupName check mysql error is "Duplicate foreign key constraint name '%s' " 
func IsServerErrorFkDupName(err error) bool {
    result := Isa(err, ER_FK_DUP_NAME)
    return result
}

    
// IsServerErrorPasswordFormat check mysql error is "The password hash doesn't have the expected format. Check if the correct password algorithm is being used with the PASSWORD() function. " 
func IsServerErrorPasswordFormat(err error) bool {
    result := Isa(err, ER_PASSWORD_FORMAT)
    return result
}

    
// IsServerErrorFkColumnCannotDrop check mysql error is "Cannot drop column '%s': needed in a foreign key constraint '%s' " 
func IsServerErrorFkColumnCannotDrop(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_DROP)
    return result
}

    
// IsServerErrorFkColumnCannotDropChild check mysql error is "Cannot drop column '%s': needed in a foreign key constraint '%s' of table '%s' " 
func IsServerErrorFkColumnCannotDropChild(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_DROP_CHILD)
    return result
}

    
// IsServerErrorFkColumnNotNull check mysql error is "Column '%s' cannot be NOT NULL: needed in a foreign key constraint '%s' SET NULL " 
func IsServerErrorFkColumnNotNull(err error) bool {
    result := Isa(err, ER_FK_COLUMN_NOT_NULL)
    return result
}

    
// IsServerErrorDupIndex check mysql error is "Duplicate index '%s' defined on the table '%s.%s'. This is deprecated and will be disallowed in a future release. " 
func IsServerErrorDupIndex(err error) bool {
    result := Isa(err, ER_DUP_INDEX)
    return result
}

    
// IsServerErrorFkColumnCannotChange check mysql error is "Cannot change column '%s': used in a foreign key constraint '%s' " 
func IsServerErrorFkColumnCannotChange(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_CHANGE)
    return result
}

    
// IsServerErrorFkColumnCannotChangeChild check mysql error is "Cannot change column '%s': used in a foreign key constraint '%s' of table '%s' " 
func IsServerErrorFkColumnCannotChangeChild(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_CHANGE_CHILD)
    return result
}

    
// IsServerErrorFkCannotDeleteParent check mysql error is "Cannot delete rows from table which is parent in a foreign key constraint '%s' of table '%s'" 
func IsServerErrorFkCannotDeleteParent(err error) bool {
    result := Isa(err, ER_FK_CANNOT_DELETE_PARENT)
    return result
}

    
// IsServerErrorUnused5 check mysql error is "Cannot delete rows from table which is parent in a foreign key constraint '%s' of table '%s'" 
func IsServerErrorUnused5(err error) bool {
    result := Isa(err, ER_UNUSED5)
    return result
}

    
// IsServerErrorMalformedPacket check mysql error is "Malformed communication packet. " 
func IsServerErrorMalformedPacket(err error) bool {
    result := Isa(err, ER_MALFORMED_PACKET)
    return result
}

    
// IsServerErrorReadOnlyMode check mysql error is "Running in read-only mode " 
func IsServerErrorReadOnlyMode(err error) bool {
    result := Isa(err, ER_READ_ONLY_MODE)
    return result
}

    
// IsServerErrorGtidNextTypeUndefinedGroup check mysql error is "When @@SESSION.GTID_NEXT is set to a GTID, you must explicitly set it to a different value after a COMMIT or ROLLBACK. Please check GTID_NEXT variable manual page for detailed explanation. Current @@SESSION.GTID_NEXT is '%s'. " 
func IsServerErrorGtidNextTypeUndefinedGroup(err error) bool {
    result := Isa(err, ER_GTID_NEXT_TYPE_UNDEFINED_GROUP)
    return result
}

    
// IsServerErrorVariableNotSettableInSp check mysql error is "The system variable %s cannot be set in stored procedures. " 
func IsServerErrorVariableNotSettableInSp(err error) bool {
    result := Isa(err, ER_VARIABLE_NOT_SETTABLE_IN_SP)
    return result
}

    
// IsServerErrorCantSetGtidPurgedWhenGtidModeIsOff check mysql error is "@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_MODE = ON. " 
func IsServerErrorCantSetGtidPurgedWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorCantSetGtidPurgedWhenGtidExecutedIsNotEmpty check mysql error is "@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_EXECUTED is empty. " 
func IsServerErrorCantSetGtidPurgedWhenGtidExecutedIsNotEmpty(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY)
    return result
}

    
// IsServerErrorCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty check mysql error is "@@GLOBAL.GTID_PURGED can only be set when there are no ongoing transactions (not even in other clients). " 
func IsServerErrorCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY)
    return result
}

    
// IsServerErrorGtidPurgedWasChanged check mysql error is "@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'. " 
func IsServerErrorGtidPurgedWasChanged(err error) bool {
    result := Isa(err, ER_GTID_PURGED_WAS_CHANGED)
    return result
}

    
// IsServerErrorGtidExecutedWasChanged check mysql error is "@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'. " 
func IsServerErrorGtidExecutedWasChanged(err error) bool {
    result := Isa(err, ER_GTID_EXECUTED_WAS_CHANGED)
    return result
}

    
// IsServerErrorBinlogStmtModeAndNoReplTables check mysql error is "Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT, and both replicated and non replicated tables are written to. " 
func IsServerErrorBinlogStmtModeAndNoReplTables(err error) bool {
    result := Isa(err, ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES)
    return result
}

    
// IsServerErrorAlterOperationNotSupported check mysql error is "%s is not supported for this operation. Try %s." 
func IsServerErrorAlterOperationNotSupported(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReason check mysql error is "%s is not supported. Reason: %s. Try %s." 
func IsServerErrorAlterOperationNotSupportedReason(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonCopy check mysql error is "COPY algorithm requires a lock" 
func IsServerErrorAlterOperationNotSupportedReasonCopy(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonPartition check mysql error is "Partition specific operations do not yet support LOCK/ALGORITHM" 
func IsServerErrorAlterOperationNotSupportedReasonPartition(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonFkRename check mysql error is "Columns participating in a foreign key are renamed" 
func IsServerErrorAlterOperationNotSupportedReasonFkRename(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonColumnType check mysql error is "Cannot change column type INPLACE" 
func IsServerErrorAlterOperationNotSupportedReasonColumnType(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonFkCheck check mysql error is "Adding foreign keys needs foreign_key_checks=OFF" 
func IsServerErrorAlterOperationNotSupportedReasonFkCheck(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonIgnore check mysql error is "Creating unique indexes with IGNORE requires COPY algorithm to remove duplicate rows" 
func IsServerErrorAlterOperationNotSupportedReasonIgnore(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE)
    return result
}

    
// IsServerErrorUnused6 check mysql error is "Creating unique indexes with IGNORE requires COPY algorithm to remove duplicate rows" 
func IsServerErrorUnused6(err error) bool {
    result := Isa(err, ER_UNUSED6)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonNopk check mysql error is "Dropping a primary key is not allowed without also adding a new primary key" 
func IsServerErrorAlterOperationNotSupportedReasonNopk(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonAutoinc check mysql error is "Adding an auto-increment column requires a lock" 
func IsServerErrorAlterOperationNotSupportedReasonAutoinc(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonHiddenFts check mysql error is "Cannot replace hidden FTS_DOC_ID with a user-visible one" 
func IsServerErrorAlterOperationNotSupportedReasonHiddenFts(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonChangeFts check mysql error is "Cannot drop or rename FTS_DOC_ID" 
func IsServerErrorAlterOperationNotSupportedReasonChangeFts(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonFts check mysql error is "Fulltext index creation requires a lock" 
func IsServerErrorAlterOperationNotSupportedReasonFts(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS)
    return result
}

    
// IsServerErrorSqlSlaveSkipCounterNotSettableInGtidMode check mysql error is "sql_slave_skip_counter can not be set when the server is running with @@GLOBAL.GTID_MODE = ON. Instead, for each transaction that you want to skip, generate an empty transaction with the same GTID as the transaction" 
func IsServerErrorSqlSlaveSkipCounterNotSettableInGtidMode(err error) bool {
    result := Isa(err, ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE)
    return result
}

    
// IsServerErrorDupUnknownInIndex check mysql error is "Duplicate entry for key '%s'" 
func IsServerErrorDupUnknownInIndex(err error) bool {
    result := Isa(err, ER_DUP_UNKNOWN_IN_INDEX)
    return result
}

    
// IsServerErrorIdentCausesTooLongPath check mysql error is "Long database name and identifier for object resulted in path length exceeding %d characters. Path: '%s'." 
func IsServerErrorIdentCausesTooLongPath(err error) bool {
    result := Isa(err, ER_IDENT_CAUSES_TOO_LONG_PATH)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonNotNull check mysql error is "cannot silently convert NULL values, as required in this SQL_MODE" 
func IsServerErrorAlterOperationNotSupportedReasonNotNull(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL)
    return result
}

    
// IsServerErrorMustChangePasswordLogin check mysql error is "Your password has expired. To log in you must change it using a client that supports expired passwords." 
func IsServerErrorMustChangePasswordLogin(err error) bool {
    result := Isa(err, ER_MUST_CHANGE_PASSWORD_LOGIN)
    return result
}

    
// IsServerErrorRowInWrongPartition check mysql error is "Found a row in wrong partition %s" 
func IsServerErrorRowInWrongPartition(err error) bool {
    result := Isa(err, ER_ROW_IN_WRONG_PARTITION)
    return result
}

    
// IsServerErrorMtsEventBiggerPendingJobsSizeMax check mysql error is "Cannot schedule event %s, relay-log name %s, position %s to Worker thread because its size %lu exceeds %lu of slave_pending_jobs_size_max." 
func IsServerErrorMtsEventBiggerPendingJobsSizeMax(err error) bool {
    result := Isa(err, ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX)
    return result
}

    
// IsServerErrorInnodbNoFtUsesParser check mysql error is "Cannot CREATE FULLTEXT INDEX WITH PARSER on InnoDB table" 
func IsServerErrorInnodbNoFtUsesParser(err error) bool {
    result := Isa(err, ER_INNODB_NO_FT_USES_PARSER)
    return result
}

    
// IsServerErrorBinlogLogicalCorruption check mysql error is "The binary log file '%s' is logically corrupted: %s" 
func IsServerErrorBinlogLogicalCorruption(err error) bool {
    result := Isa(err, ER_BINLOG_LOGICAL_CORRUPTION)
    return result
}

    
// IsServerErrorWarnPurgeLogInUse check mysql error is "file %s was not purged because it was being read by %d thread(s), purged only %d out of %d files." 
func IsServerErrorWarnPurgeLogInUse(err error) bool {
    result := Isa(err, ER_WARN_PURGE_LOG_IN_USE)
    return result
}

    
// IsServerErrorWarnPurgeLogIsActive check mysql error is "file %s was not purged because it is the active log file." 
func IsServerErrorWarnPurgeLogIsActive(err error) bool {
    result := Isa(err, ER_WARN_PURGE_LOG_IS_ACTIVE)
    return result
}

    
// IsServerErrorAutoIncrementConflict check mysql error is "Auto-increment value in UPDATE conflicts with internally generated values" 
func IsServerErrorAutoIncrementConflict(err error) bool {
    result := Isa(err, ER_AUTO_INCREMENT_CONFLICT)
    return result
}

    
// IsWarnOnBlockholeInRbr check mysql error is "Row events are not logged for %s statements that modify BLACKHOLE tables in row format. Table(s): '%s'" 
func IsWarnOnBlockholeInRbr(err error) bool {
    result := Isa(err, WARN_ON_BLOCKHOLE_IN_RBR)
    return result
}

    
// IsServerErrorSlaveMiInitRepository check mysql error is "Slave failed to initialize master info structure from the repository" 
func IsServerErrorSlaveMiInitRepository(err error) bool {
    result := Isa(err, ER_SLAVE_MI_INIT_REPOSITORY)
    return result
}

    
// IsServerErrorSlaveRliInitRepository check mysql error is "Slave failed to initialize relay log info structure from the repository" 
func IsServerErrorSlaveRliInitRepository(err error) bool {
    result := Isa(err, ER_SLAVE_RLI_INIT_REPOSITORY)
    return result
}

    
// IsServerErrorAccessDeniedChangeUserError check mysql error is "Access denied trying to change to user '%s'@'%s' (using password: %s). Disconnecting." 
func IsServerErrorAccessDeniedChangeUserError(err error) bool {
    result := Isa(err, ER_ACCESS_DENIED_CHANGE_USER_ERROR)
    return result
}

    
// IsServerErrorInnodbReadOnly check mysql error is "InnoDB is in read only mode." 
func IsServerErrorInnodbReadOnly(err error) bool {
    result := Isa(err, ER_INNODB_READ_ONLY)
    return result
}

    
// IsServerErrorStopSlaveSqlThreadTimeout check mysql error is "STOP SLAVE command execution is incomplete: Slave SQL thread got the stop signal, thread is busy, SQL thread will stop once the current task is complete." 
func IsServerErrorStopSlaveSqlThreadTimeout(err error) bool {
    result := Isa(err, ER_STOP_SLAVE_SQL_THREAD_TIMEOUT)
    return result
}

    
// IsServerErrorStopSlaveIoThreadTimeout check mysql error is "STOP SLAVE command execution is incomplete: Slave IO thread got the stop signal, thread is busy, IO thread will stop once the current task is complete." 
func IsServerErrorStopSlaveIoThreadTimeout(err error) bool {
    result := Isa(err, ER_STOP_SLAVE_IO_THREAD_TIMEOUT)
    return result
}

    
// IsServerErrorTableCorrupt check mysql error is "Operation cannot be performed. The table '%s.%s' is missing, corrupt or contains bad data." 
func IsServerErrorTableCorrupt(err error) bool {
    result := Isa(err, ER_TABLE_CORRUPT)
    return result
}

    
// IsServerErrorTempFileWriteFailure check mysql error is "Temporary file write failure." 
func IsServerErrorTempFileWriteFailure(err error) bool {
    result := Isa(err, ER_TEMP_FILE_WRITE_FAILURE)
    return result
}

    
// IsServerErrorInnodbFtAuxNotHexId check mysql error is "Upgrade index name failed, please use create index(alter table) algorithm copy to rebuild index." 
func IsServerErrorInnodbFtAuxNotHexId(err error) bool {
    result := Isa(err, ER_INNODB_FT_AUX_NOT_HEX_ID)
    return result
}

    
// IsServerErrorOldTemporalsUpgraded check mysql error is "TIME/TIMESTAMP/DATETIME columns of old format have been upgraded to the new format." 
func IsServerErrorOldTemporalsUpgraded(err error) bool {
    result := Isa(err, ER_OLD_TEMPORALS_UPGRADED)
    return result
}

    
// IsServerErrorInnodbForcedRecovery check mysql error is "Operation not allowed when innodb_forced_recovery > 0." 
func IsServerErrorInnodbForcedRecovery(err error) bool {
    result := Isa(err, ER_INNODB_FORCED_RECOVERY)
    return result
}

    
// IsServerErrorAesInvalidIv check mysql error is "The initialization vector supplied to %s is too short. Must be at least %d bytes long" 
func IsServerErrorAesInvalidIv(err error) bool {
    result := Isa(err, ER_AES_INVALID_IV)
    return result
}

    
// IsServerErrorPluginCannotBeUninstalled check mysql error is "Plugin '%s' cannot be uninstalled now. %s" 
func IsServerErrorPluginCannotBeUninstalled(err error) bool {
    result := Isa(err, ER_PLUGIN_CANNOT_BE_UNINSTALLED)
    return result
}

    
// IsServerErrorGtidUnsafeBinlogSplittableStatementAndGtidGroup check mysql error is "Cannot execute statement because it needs to be written to the binary log as multiple statements, and this is not allowed when @@SESSION.GTID_NEXT == 'UUID:NUMBER'." 
func IsServerErrorGtidUnsafeBinlogSplittableStatementAndGtidGroup(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP)
    return result
}

    
// IsServerErrorSlaveHasMoreGtidsThanMaster check mysql error is "Slave has more GTIDs than the master has, using the master's SERVER_UUID. This may indicate that the end of the binary log was truncated or that the last binary log file was lost, e.g., after a power or disk failure when sync_binlog != 1. The master may or may not have rolled back transactions that were already replicated to the slave. Suggest to replicate any transactions that master has rolled back from slave to master, and/or commit empty transactions on master to account for transactions that have been committed on master but are not included in GTID_EXECUTED." 
func IsServerErrorSlaveHasMoreGtidsThanMaster(err error) bool {
    result := Isa(err, ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER)
    return result
}

    
// IsServerErrorMissingKey check mysql error is "The table '%s.%s' does not have the necessary key(s) defined on it. Please check the table definition and create index(s) accordingly." 
func IsServerErrorMissingKey(err error) bool {
    result := Isa(err, ER_MISSING_KEY)
    return result
}

    
// IsWarnNamedPipeAccessEveryone check mysql error is "Setting named_pipe_full_access_group='%s' is insecure. Consider using a Windows group with fewer members." 
func IsWarnNamedPipeAccessEveryone(err error) bool {
    result := Isa(err, WARN_NAMED_PIPE_ACCESS_EVERYONE)
    return result
}

    
// IsServerErrorFoundMissingGtids check mysql error is "Cannot replicate to server with server_uuid='%s' because the present server has purged required binary logs. The connecting server needs to replicate the missing transactions from elsewhere, or be replaced by a new server created from a more recent backup. To prevent this error in the future, consider increasing the binary log expiration period on the present server. %s." 
func IsServerErrorFoundMissingGtids(err error) bool {
    result := Isa(err, ER_FOUND_MISSING_GTIDS)
    return result
}

    
// IsServerErrorSlaveIoThreadMustStop check mysql error is "This operation cannot be performed with a running slave io thread" 
func IsServerErrorSlaveIoThreadMustStop(err error) bool {
    result := Isa(err, ER_SLAVE_IO_THREAD_MUST_STOP)
    return result
}

    
// IsServerErrorFileCorrupt check mysql error is "File %s is corrupted " 
func IsServerErrorFileCorrupt(err error) bool {
    result := Isa(err, ER_FILE_CORRUPT)
    return result
}

    
// IsServerErrorErrorOnMaster check mysql error is "Query partially completed on the master (error on master: %d) and was aborted. There is a chance that your master is inconsistent at this point. If you are sure that your master is ok, run this query manually on the slave and then restart the slave with SET GLOBAL SQL_SLAVE_SKIP_COUNTER=1" 
func IsServerErrorErrorOnMaster(err error) bool {
    result := Isa(err, ER_ERROR_ON_MASTER)
    return result
}

    
// IsServerErrorInconsistentError check mysql error is "Query caused different errors on master and slave. Error on master: message (format)='%s' error code=%d" 
func IsServerErrorInconsistentError(err error) bool {
    result := Isa(err, ER_INCONSISTENT_ERROR)
    return result
}

    
// IsServerErrorStorageEngineNotLoaded check mysql error is "Storage engine for table '%s'.'%s' is not loaded. " 
func IsServerErrorStorageEngineNotLoaded(err error) bool {
    result := Isa(err, ER_STORAGE_ENGINE_NOT_LOADED)
    return result
}

    
// IsServerErrorGetStackedDaWithoutActiveHandler check mysql error is "GET STACKED DIAGNOSTICS when handler not active " 
func IsServerErrorGetStackedDaWithoutActiveHandler(err error) bool {
    result := Isa(err, ER_GET_STACKED_DA_WITHOUT_ACTIVE_HANDLER)
    return result
}

    
// IsServerErrorWarnLegacySyntaxConverted check mysql error is "%s is no longer supported. The statement was converted to %s. " 
func IsServerErrorWarnLegacySyntaxConverted(err error) bool {
    result := Isa(err, ER_WARN_LEGACY_SYNTAX_CONVERTED)
    return result
}

    
// IsServerErrorBinlogUnsafeFulltextPlugin check mysql error is "Statement is unsafe because it uses a fulltext parser plugin which may not return the same value on the slave." 
func IsServerErrorBinlogUnsafeFulltextPlugin(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_FULLTEXT_PLUGIN)
    return result
}

    
// IsServerErrorCannotDiscardTemporaryTable check mysql error is "Cannot DISCARD/IMPORT tablespace associated with temporary table" 
func IsServerErrorCannotDiscardTemporaryTable(err error) bool {
    result := Isa(err, ER_CANNOT_DISCARD_TEMPORARY_TABLE)
    return result
}

    
// IsServerErrorFkDepthExceeded check mysql error is "Foreign key cascade delete/update exceeds max depth of %d." 
func IsServerErrorFkDepthExceeded(err error) bool {
    result := Isa(err, ER_FK_DEPTH_EXCEEDED)
    return result
}

    
// IsServerErrorColCountDoesntMatchPleaseUpdateV2 check mysql error is "Column count of %s.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please use mysql_upgrade to fix this error." 
func IsServerErrorColCountDoesntMatchPleaseUpdateV2(err error) bool {
    result := Isa(err, ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2)
    return result
}

    
// IsServerErrorWarnTriggerDoesntHaveCreated check mysql error is "Trigger %s.%s.%s does not have CREATED attribute." 
func IsServerErrorWarnTriggerDoesntHaveCreated(err error) bool {
    result := Isa(err, ER_WARN_TRIGGER_DOESNT_HAVE_CREATED)
    return result
}

    
// IsServerErrorReferencedTrgDoesNotExist check mysql error is "Referenced trigger '%s' for the given action time and event type does not exist." 
func IsServerErrorReferencedTrgDoesNotExist(err error) bool {
    result := Isa(err, ER_REFERENCED_TRG_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorExplainNotSupported check mysql error is "EXPLAIN FOR CONNECTION command is supported only for SELECT/UPDATE/INSERT/DELETE/REPLACE" 
func IsServerErrorExplainNotSupported(err error) bool {
    result := Isa(err, ER_EXPLAIN_NOT_SUPPORTED)
    return result
}

    
// IsServerErrorInvalidFieldSize check mysql error is "Invalid size for column '%s'." 
func IsServerErrorInvalidFieldSize(err error) bool {
    result := Isa(err, ER_INVALID_FIELD_SIZE)
    return result
}

    
// IsServerErrorMissingHaCreateOption check mysql error is "Table storage engine '%s' found required create option missing" 
func IsServerErrorMissingHaCreateOption(err error) bool {
    result := Isa(err, ER_MISSING_HA_CREATE_OPTION)
    return result
}

    
// IsServerErrorEngineOutOfMemory check mysql error is "Out of memory in storage engine '%s'." 
func IsServerErrorEngineOutOfMemory(err error) bool {
    result := Isa(err, ER_ENGINE_OUT_OF_MEMORY)
    return result
}

    
// IsServerErrorPasswordExpireAnonymousUser check mysql error is "The password for anonymous user cannot be expired." 
func IsServerErrorPasswordExpireAnonymousUser(err error) bool {
    result := Isa(err, ER_PASSWORD_EXPIRE_ANONYMOUS_USER)
    return result
}

    
// IsServerErrorSlaveSqlThreadMustStop check mysql error is "This operation cannot be performed with a running slave sql thread" 
func IsServerErrorSlaveSqlThreadMustStop(err error) bool {
    result := Isa(err, ER_SLAVE_SQL_THREAD_MUST_STOP)
    return result
}

    
// IsServerErrorNoFtMaterializedSubquery check mysql error is "Cannot create FULLTEXT index on materialized subquery" 
func IsServerErrorNoFtMaterializedSubquery(err error) bool {
    result := Isa(err, ER_NO_FT_MATERIALIZED_SUBQUERY)
    return result
}

    
// IsServerErrorInnodbUndoLogFull check mysql error is "Undo Log error: %s" 
func IsServerErrorInnodbUndoLogFull(err error) bool {
    result := Isa(err, ER_INNODB_UNDO_LOG_FULL)
    return result
}

    
// IsServerErrorInvalidArgumentForLogarithm check mysql error is "Invalid argument for logarithm" 
func IsServerErrorInvalidArgumentForLogarithm(err error) bool {
    result := Isa(err, ER_INVALID_ARGUMENT_FOR_LOGARITHM)
    return result
}

    
// IsServerErrorSlaveChannelIoThreadMustStop check mysql error is "This operation cannot be performed with a running slave io thread" 
func IsServerErrorSlaveChannelIoThreadMustStop(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_IO_THREAD_MUST_STOP)
    return result
}

    
// IsServerErrorWarnOpenTempTablesMustBeZero check mysql error is "This operation may not be safe when the slave has temporary tables. The tables will be kept open until the server restarts or until the tables are deleted by any replicated DROP statement. Suggest to wait until slave_open_temp_tables = 0." 
func IsServerErrorWarnOpenTempTablesMustBeZero(err error) bool {
    result := Isa(err, ER_WARN_OPEN_TEMP_TABLES_MUST_BE_ZERO)
    return result
}

    
// IsServerErrorWarnOnlyMasterLogFileNoPos check mysql error is "CHANGE MASTER TO with a MASTER_LOG_FILE clause but no MASTER_LOG_POS clause may not be safe. The old position value may not be valid for the new binary log file." 
func IsServerErrorWarnOnlyMasterLogFileNoPos(err error) bool {
    result := Isa(err, ER_WARN_ONLY_MASTER_LOG_FILE_NO_POS)
    return result
}

    
// IsServerErrorQueryTimeout check mysql error is "Query execution was interrupted, maximum statement execution time exceeded" 
func IsServerErrorQueryTimeout(err error) bool {
    result := Isa(err, ER_QUERY_TIMEOUT)
    return result
}

    
// IsServerErrorNonRoSelectDisableTimer check mysql error is "Select is not a read only statement, disabling timer" 
func IsServerErrorNonRoSelectDisableTimer(err error) bool {
    result := Isa(err, ER_NON_RO_SELECT_DISABLE_TIMER)
    return result
}

    
// IsServerErrorDupListEntry check mysql error is "Duplicate entry '%s'." 
func IsServerErrorDupListEntry(err error) bool {
    result := Isa(err, ER_DUP_LIST_ENTRY)
    return result
}

    
// IsServerErrorSqlModeNoEffect check mysql error is "'%s' mode no longer has any effect. Use STRICT_ALL_TABLES or STRICT_TRANS_TABLES instead." 
func IsServerErrorSqlModeNoEffect(err error) bool {
    result := Isa(err, ER_SQL_MODE_NO_EFFECT)
    return result
}

    
// IsServerErrorAggregateOrderForUnion check mysql error is "Expression #%u of ORDER BY contains aggregate function and applies to a UNION" 
func IsServerErrorAggregateOrderForUnion(err error) bool {
    result := Isa(err, ER_AGGREGATE_ORDER_FOR_UNION)
    return result
}

    
// IsServerErrorAggregateOrderNonAggQuery check mysql error is "Expression #%u of ORDER BY contains aggregate function and applies to the result of a non-aggregated query" 
func IsServerErrorAggregateOrderNonAggQuery(err error) bool {
    result := Isa(err, ER_AGGREGATE_ORDER_NON_AGG_QUERY)
    return result
}

    
// IsServerErrorSlaveWorkerStoppedPreviousThdError check mysql error is "Slave worker has stopped after at least one previous worker encountered an error when slave-preserve-commit-order was enabled. To preserve commit order, the last transaction executed by this thread has not been committed. When restarting the slave after fixing any failed threads, you should fix this worker as well." 
func IsServerErrorSlaveWorkerStoppedPreviousThdError(err error) bool {
    result := Isa(err, ER_SLAVE_WORKER_STOPPED_PREVIOUS_THD_ERROR)
    return result
}

    
// IsServerErrorDontSupportSlavePreserveCommitOrder check mysql error is "slave_preserve_commit_order is not supported %s." 
func IsServerErrorDontSupportSlavePreserveCommitOrder(err error) bool {
    result := Isa(err, ER_DONT_SUPPORT_SLAVE_PRESERVE_COMMIT_ORDER)
    return result
}

    
// IsServerErrorServerOfflineMode check mysql error is "The server is currently in offline mode" 
func IsServerErrorServerOfflineMode(err error) bool {
    result := Isa(err, ER_SERVER_OFFLINE_MODE)
    return result
}

    
// IsServerErrorGisDifferentSrids check mysql error is "Binary geometry function %s given two geometries of different srids: %u and %u, which should have been identical." 
func IsServerErrorGisDifferentSrids(err error) bool {
    result := Isa(err, ER_GIS_DIFFERENT_SRIDS)
    return result
}

    
// IsServerErrorGisUnsupportedArgument check mysql error is "Calling geometry function %s with unsupported types of arguments." 
func IsServerErrorGisUnsupportedArgument(err error) bool {
    result := Isa(err, ER_GIS_UNSUPPORTED_ARGUMENT)
    return result
}

    
// IsServerErrorGisUnknownError check mysql error is "Unknown GIS error occured in function %s." 
func IsServerErrorGisUnknownError(err error) bool {
    result := Isa(err, ER_GIS_UNKNOWN_ERROR)
    return result
}

    
// IsServerErrorGisUnknownException check mysql error is "Unknown exception caught in GIS function %s." 
func IsServerErrorGisUnknownException(err error) bool {
    result := Isa(err, ER_GIS_UNKNOWN_EXCEPTION)
    return result
}

    
// IsServerErrorGisInvalidData check mysql error is "Invalid GIS data provided to function %s." 
func IsServerErrorGisInvalidData(err error) bool {
    result := Isa(err, ER_GIS_INVALID_DATA)
    return result
}

    
// IsServerErrorBoostGeometryEmptyInputException check mysql error is "The geometry has no data in function %s." 
func IsServerErrorBoostGeometryEmptyInputException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_EMPTY_INPUT_EXCEPTION)
    return result
}

    
// IsServerErrorBoostGeometryCentroidException check mysql error is "Unable to calculate centroid because geometry is empty in function %s." 
func IsServerErrorBoostGeometryCentroidException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_CENTROID_EXCEPTION)
    return result
}

    
// IsServerErrorBoostGeometryOverlayInvalidInputException check mysql error is "Geometry overlay calculation error: geometry data is invalid in function %s." 
func IsServerErrorBoostGeometryOverlayInvalidInputException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_OVERLAY_INVALID_INPUT_EXCEPTION)
    return result
}

    
// IsServerErrorBoostGeometryTurnInfoException check mysql error is "Geometry turn info calculation error: geometry data is invalid in function %s." 
func IsServerErrorBoostGeometryTurnInfoException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_TURN_INFO_EXCEPTION)
    return result
}

    
// IsServerErrorBoostGeometrySelfIntersectionPointException check mysql error is "Analysis procedures of intersection points interrupted unexpectedly in function %s." 
func IsServerErrorBoostGeometrySelfIntersectionPointException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_SELF_INTERSECTION_POINT_EXCEPTION)
    return result
}

    
// IsServerErrorBoostGeometryUnknownException check mysql error is "Unknown exception thrown in function %s." 
func IsServerErrorBoostGeometryUnknownException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_UNKNOWN_EXCEPTION)
    return result
}

    
// IsServerErrorStdBadAllocError check mysql error is "Memory allocation error: %s in function %s." 
func IsServerErrorStdBadAllocError(err error) bool {
    result := Isa(err, ER_STD_BAD_ALLOC_ERROR)
    return result
}

    
// IsServerErrorStdDomainError check mysql error is "Domain error: %s in function %s." 
func IsServerErrorStdDomainError(err error) bool {
    result := Isa(err, ER_STD_DOMAIN_ERROR)
    return result
}

    
// IsServerErrorStdLengthError check mysql error is "Length error: %s in function %s." 
func IsServerErrorStdLengthError(err error) bool {
    result := Isa(err, ER_STD_LENGTH_ERROR)
    return result
}

    
// IsServerErrorStdInvalidArgument check mysql error is "Invalid argument error: %s in function %s." 
func IsServerErrorStdInvalidArgument(err error) bool {
    result := Isa(err, ER_STD_INVALID_ARGUMENT)
    return result
}

    
// IsServerErrorStdOutOfRangeError check mysql error is "Out of range error: %s in function %s." 
func IsServerErrorStdOutOfRangeError(err error) bool {
    result := Isa(err, ER_STD_OUT_OF_RANGE_ERROR)
    return result
}

    
// IsServerErrorStdOverflowError check mysql error is "Overflow error error: %s in function %s." 
func IsServerErrorStdOverflowError(err error) bool {
    result := Isa(err, ER_STD_OVERFLOW_ERROR)
    return result
}

    
// IsServerErrorStdRangeError check mysql error is "Range error: %s in function %s." 
func IsServerErrorStdRangeError(err error) bool {
    result := Isa(err, ER_STD_RANGE_ERROR)
    return result
}

    
// IsServerErrorStdUnderflowError check mysql error is "Underflow error: %s in function %s." 
func IsServerErrorStdUnderflowError(err error) bool {
    result := Isa(err, ER_STD_UNDERFLOW_ERROR)
    return result
}

    
// IsServerErrorStdLogicError check mysql error is "Logic error: %s in function %s." 
func IsServerErrorStdLogicError(err error) bool {
    result := Isa(err, ER_STD_LOGIC_ERROR)
    return result
}

    
// IsServerErrorStdRuntimeError check mysql error is "Runtime error: %s in function %s." 
func IsServerErrorStdRuntimeError(err error) bool {
    result := Isa(err, ER_STD_RUNTIME_ERROR)
    return result
}

    
// IsServerErrorStdUnknownException check mysql error is "Unknown exception: %s in function %s." 
func IsServerErrorStdUnknownException(err error) bool {
    result := Isa(err, ER_STD_UNKNOWN_EXCEPTION)
    return result
}

    
// IsServerErrorGisDataWrongEndianess check mysql error is "Geometry byte string must be little endian." 
func IsServerErrorGisDataWrongEndianess(err error) bool {
    result := Isa(err, ER_GIS_DATA_WRONG_ENDIANESS)
    return result
}

    
// IsServerErrorChangeMasterPasswordLength check mysql error is "The password provided for the replication user exceeds the maximum length of 32 characters" 
func IsServerErrorChangeMasterPasswordLength(err error) bool {
    result := Isa(err, ER_CHANGE_MASTER_PASSWORD_LENGTH)
    return result
}

    
// IsServerErrorUserLockWrongName check mysql error is "Incorrect user-level lock name '%s'." 
func IsServerErrorUserLockWrongName(err error) bool {
    result := Isa(err, ER_USER_LOCK_WRONG_NAME)
    return result
}

    
// IsServerErrorUserLockDeadlock check mysql error is "Deadlock found when trying to get user-level lock" 
func IsServerErrorUserLockDeadlock(err error) bool {
    result := Isa(err, ER_USER_LOCK_DEADLOCK)
    return result
}

    
// IsServerErrorReplaceInaccessibleRows check mysql error is "REPLACE cannot be executed as it requires deleting rows that are not in the view" 
func IsServerErrorReplaceInaccessibleRows(err error) bool {
    result := Isa(err, ER_REPLACE_INACCESSIBLE_ROWS)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonGis check mysql error is "Do not support online operation on table with GIS index" 
func IsServerErrorAlterOperationNotSupportedReasonGis(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_GIS)
    return result
}

    
// IsServerErrorIllegalUserVar check mysql error is "User variable name '%s' is illegal" 
func IsServerErrorIllegalUserVar(err error) bool {
    result := Isa(err, ER_ILLEGAL_USER_VAR)
    return result
}

    
// IsServerErrorGtidModeOff check mysql error is "Cannot %s when GTID_MODE = OFF." 
func IsServerErrorGtidModeOff(err error) bool {
    result := Isa(err, ER_GTID_MODE_OFF)
    return result
}

    
// IsServerErrorUnsupportedByReplicationThread check mysql error is "Cannot %s from a replication slave thread." 
func IsServerErrorUnsupportedByReplicationThread(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_BY_REPLICATION_THREAD)
    return result
}

    
// IsServerErrorIncorrectType check mysql error is "Incorrect type for argument %s in function %s." 
func IsServerErrorIncorrectType(err error) bool {
    result := Isa(err, ER_INCORRECT_TYPE)
    return result
}

    
// IsServerErrorFieldInOrderNotSelect check mysql error is "Expression #%u of ORDER BY clause is not in SELECT list, references column '%s' which is not in SELECT list" 
func IsServerErrorFieldInOrderNotSelect(err error) bool {
    result := Isa(err, ER_FIELD_IN_ORDER_NOT_SELECT)
    return result
}

    
// IsServerErrorAggregateInOrderNotSelect check mysql error is "Expression #%u of ORDER BY clause is not in SELECT list, contains aggregate function" 
func IsServerErrorAggregateInOrderNotSelect(err error) bool {
    result := Isa(err, ER_AGGREGATE_IN_ORDER_NOT_SELECT)
    return result
}

    
// IsServerErrorInvalidRplWildTableFilterPattern check mysql error is "Supplied filter list contains a value which is not in the required format 'db_pattern.table_pattern'" 
func IsServerErrorInvalidRplWildTableFilterPattern(err error) bool {
    result := Isa(err, ER_INVALID_RPL_WILD_TABLE_FILTER_PATTERN)
    return result
}

    
// IsServerErrorNetOkPacketTooLarge check mysql error is "OK packet too large" 
func IsServerErrorNetOkPacketTooLarge(err error) bool {
    result := Isa(err, ER_NET_OK_PACKET_TOO_LARGE)
    return result
}

    
// IsServerErrorInvalidJsonData check mysql error is "Invalid JSON data provided to function %s: %s" 
func IsServerErrorInvalidJsonData(err error) bool {
    result := Isa(err, ER_INVALID_JSON_DATA)
    return result
}

    
// IsServerErrorInvalidGeojsonMissingMember check mysql error is "Invalid GeoJSON data provided to function %s: Missing required member '%s'" 
func IsServerErrorInvalidGeojsonMissingMember(err error) bool {
    result := Isa(err, ER_INVALID_GEOJSON_MISSING_MEMBER)
    return result
}

    
// IsServerErrorInvalidGeojsonWrongType check mysql error is "Invalid GeoJSON data provided to function %s: Member '%s' must be of type '%s'" 
func IsServerErrorInvalidGeojsonWrongType(err error) bool {
    result := Isa(err, ER_INVALID_GEOJSON_WRONG_TYPE)
    return result
}

    
// IsServerErrorInvalidGeojsonUnspecified check mysql error is "Invalid GeoJSON data provided to function %s" 
func IsServerErrorInvalidGeojsonUnspecified(err error) bool {
    result := Isa(err, ER_INVALID_GEOJSON_UNSPECIFIED)
    return result
}

    
// IsServerErrorDimensionUnsupported check mysql error is "Unsupported number of coordinate dimensions in function %s: Found %u, expected %u" 
func IsServerErrorDimensionUnsupported(err error) bool {
    result := Isa(err, ER_DIMENSION_UNSUPPORTED)
    return result
}

    
// IsServerErrorSlaveChannelDoesNotExist check mysql error is "Slave channel '%s' does not exist." 
func IsServerErrorSlaveChannelDoesNotExist(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorSlaveMultipleChannelsHostPort check mysql error is "A slave channel '%s' already exists for the given host and port combination." 
func IsServerErrorSlaveMultipleChannelsHostPort(err error) bool {
    result := Isa(err, ER_SLAVE_MULTIPLE_CHANNELS_HOST_PORT)
    return result
}

    
// IsServerErrorSlaveChannelNameInvalidOrTooLong check mysql error is "Couldn't create channel: Channel name is either invalid or too long." 
func IsServerErrorSlaveChannelNameInvalidOrTooLong(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_NAME_INVALID_OR_TOO_LONG)
    return result
}

    
// IsServerErrorSlaveNewChannelWrongRepository check mysql error is "To have multiple channels, repository cannot be of type FILE" 
func IsServerErrorSlaveNewChannelWrongRepository(err error) bool {
    result := Isa(err, ER_SLAVE_NEW_CHANNEL_WRONG_REPOSITORY)
    return result
}

    
// IsServerErrorSlaveChannelDelete check mysql error is "Cannot delete slave info objects for channel '%s'." 
func IsServerErrorSlaveChannelDelete(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_DELETE)
    return result
}

    
// IsServerErrorSlaveMultipleChannelsCmd check mysql error is "Multiple channels exist on the slave. Please provide channel name as an argument." 
func IsServerErrorSlaveMultipleChannelsCmd(err error) bool {
    result := Isa(err, ER_SLAVE_MULTIPLE_CHANNELS_CMD)
    return result
}

    
// IsServerErrorSlaveMaxChannelsExceeded check mysql error is "Maximum number of replication channels allowed exceeded." 
func IsServerErrorSlaveMaxChannelsExceeded(err error) bool {
    result := Isa(err, ER_SLAVE_MAX_CHANNELS_EXCEEDED)
    return result
}

    
// IsServerErrorSlaveChannelMustStop check mysql error is "This operation cannot be performed with running replication threads" 
func IsServerErrorSlaveChannelMustStop(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_MUST_STOP)
    return result
}

    
// IsServerErrorSlaveChannelNotRunning check mysql error is "This operation requires running replication threads" 
func IsServerErrorSlaveChannelNotRunning(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_NOT_RUNNING)
    return result
}

    
// IsServerErrorSlaveChannelWasRunning check mysql error is "Replication thread(s) for channel '%s' are already runnning." 
func IsServerErrorSlaveChannelWasRunning(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_WAS_RUNNING)
    return result
}

    
// IsServerErrorSlaveChannelWasNotRunning check mysql error is "Replication thread(s) for channel '%s' are already stopped." 
func IsServerErrorSlaveChannelWasNotRunning(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_WAS_NOT_RUNNING)
    return result
}

    
// IsServerErrorSlaveChannelSqlThreadMustStop check mysql error is "This operation cannot be performed with a running slave sql thread" 
func IsServerErrorSlaveChannelSqlThreadMustStop(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_SQL_THREAD_MUST_STOP)
    return result
}

    
// IsServerErrorSlaveChannelSqlSkipCounter check mysql error is "When sql_slave_skip_counter > 0, it is not allowed to start more than one SQL thread by using 'START SLAVE [SQL_THREAD]'. Value of sql_slave_skip_counter can only be used by one SQL thread at a time. Please use 'START SLAVE [SQL_THREAD] FOR CHANNEL' to start the SQL thread which will use the value of sql_slave_skip_counter." 
func IsServerErrorSlaveChannelSqlSkipCounter(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_SQL_SKIP_COUNTER)
    return result
}

    
// IsServerErrorWrongFieldWithGroupV2 check mysql error is "Expression #%u of %s is not in GROUP BY clause and contains nonaggregated column '%s' which is not functionally dependent on columns in GROUP BY clause" 
func IsServerErrorWrongFieldWithGroupV2(err error) bool {
    result := Isa(err, ER_WRONG_FIELD_WITH_GROUP_V2)
    return result
}

    
// IsServerErrorMixOfGroupFuncAndFieldsV2 check mysql error is "In aggregated query without GROUP BY, expression #%u of %s contains nonaggregated column '%s'" 
func IsServerErrorMixOfGroupFuncAndFieldsV2(err error) bool {
    result := Isa(err, ER_MIX_OF_GROUP_FUNC_AND_FIELDS_V2)
    return result
}

    
// IsServerErrorWarnDeprecatedSysvarUpdate check mysql error is "Updating '%s' is deprecated. It will be made read-only in a future release." 
func IsServerErrorWarnDeprecatedSysvarUpdate(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYSVAR_UPDATE)
    return result
}

    
// IsServerErrorWarnDeprecatedSqlmode check mysql error is "Changing sql mode '%s' is deprecated. It will be removed in a future release." 
func IsServerErrorWarnDeprecatedSqlmode(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SQLMODE)
    return result
}

    
// IsServerErrorCannotLogPartialDropDatabaseWithGtid check mysql error is "DROP DATABASE failed" 
func IsServerErrorCannotLogPartialDropDatabaseWithGtid(err error) bool {
    result := Isa(err, ER_CANNOT_LOG_PARTIAL_DROP_DATABASE_WITH_GTID)
    return result
}

    
// IsServerErrorGroupReplicationConfiguration check mysql error is "The server is not configured properly to be an active member of the group. Please see more details on error log." 
func IsServerErrorGroupReplicationConfiguration(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_CONFIGURATION)
    return result
}

    
// IsServerErrorGroupReplicationRunning check mysql error is "The START GROUP_REPLICATION command failed since the group is already running." 
func IsServerErrorGroupReplicationRunning(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_RUNNING)
    return result
}

    
// IsServerErrorGroupReplicationApplierInitError check mysql error is "The START GROUP_REPLICATION command failed as the applier module failed to start." 
func IsServerErrorGroupReplicationApplierInitError(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_APPLIER_INIT_ERROR)
    return result
}

    
// IsServerErrorGroupReplicationStopApplierThreadTimeout check mysql error is "The STOP GROUP_REPLICATION command execution is incomplete: The applier thread got the stop signal while it was busy. The applier thread will stop once the current task is complete." 
func IsServerErrorGroupReplicationStopApplierThreadTimeout(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_STOP_APPLIER_THREAD_TIMEOUT)
    return result
}

    
// IsServerErrorGroupReplicationCommunicationLayerSessionError check mysql error is "The START GROUP_REPLICATION command failed as there was an error when initializing the group communication layer." 
func IsServerErrorGroupReplicationCommunicationLayerSessionError(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_COMMUNICATION_LAYER_SESSION_ERROR)
    return result
}

    
// IsServerErrorGroupReplicationCommunicationLayerJoinError check mysql error is "The START GROUP_REPLICATION command failed as there was an error when joining the communication group." 
func IsServerErrorGroupReplicationCommunicationLayerJoinError(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_COMMUNICATION_LAYER_JOIN_ERROR)
    return result
}

    
// IsServerErrorBeforeDmlValidationError check mysql error is "The table does not comply with the requirements by an external plugin." 
func IsServerErrorBeforeDmlValidationError(err error) bool {
    result := Isa(err, ER_BEFORE_DML_VALIDATION_ERROR)
    return result
}

    
// IsServerErrorPreventsVariableWithoutRbr check mysql error is "Cannot change the value of variable %s without binary log format as ROW." 
func IsServerErrorPreventsVariableWithoutRbr(err error) bool {
    result := Isa(err, ER_PREVENTS_VARIABLE_WITHOUT_RBR)
    return result
}

    
// IsServerErrorRunHookError check mysql error is "Error on observer while running replication hook '%s'." 
func IsServerErrorRunHookError(err error) bool {
    result := Isa(err, ER_RUN_HOOK_ERROR)
    return result
}

    
// IsServerErrorTransactionRollbackDuringCommit check mysql error is "Plugin instructed the server to rollback the current transaction." 
func IsServerErrorTransactionRollbackDuringCommit(err error) bool {
    result := Isa(err, ER_TRANSACTION_ROLLBACK_DURING_COMMIT)
    return result
}

    
// IsServerErrorGeneratedColumnFunctionIsNotAllowed check mysql error is "Expression of generated column '%s' contains a disallowed function." 
func IsServerErrorGeneratedColumnFunctionIsNotAllowed(err error) bool {
    result := Isa(err, ER_GENERATED_COLUMN_FUNCTION_IS_NOT_ALLOWED)
    return result
}

    
// IsServerErrorKeyBasedOnGeneratedColumn check mysql error is "Key/Index cannot be defined on a virtual generated column." 
func IsServerErrorKeyBasedOnGeneratedColumn(err error) bool {
    result := Isa(err, ER_KEY_BASED_ON_GENERATED_COLUMN)
    return result
}

    
// IsServerErrorUnsupportedAlterInplaceOnVirtualColumn check mysql error is "INPLACE ADD or DROP of virtual columns cannot be combined with other ALTER TABLE actions" 
func IsServerErrorUnsupportedAlterInplaceOnVirtualColumn(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_ALTER_INPLACE_ON_VIRTUAL_COLUMN)
    return result
}

    
// IsServerErrorWrongFkOptionForGeneratedColumn check mysql error is "Cannot define foreign key with %s clause on a generated column." 
func IsServerErrorWrongFkOptionForGeneratedColumn(err error) bool {
    result := Isa(err, ER_WRONG_FK_OPTION_FOR_GENERATED_COLUMN)
    return result
}

    
// IsServerErrorNonDefaultValueForGeneratedColumn check mysql error is "The value specified for generated column '%s' in table '%s' is not allowed." 
func IsServerErrorNonDefaultValueForGeneratedColumn(err error) bool {
    result := Isa(err, ER_NON_DEFAULT_VALUE_FOR_GENERATED_COLUMN)
    return result
}

    
// IsServerErrorUnsupportedActionOnGeneratedColumn check mysql error is "'%s' is not supported for generated columns." 
func IsServerErrorUnsupportedActionOnGeneratedColumn(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_ACTION_ON_GENERATED_COLUMN)
    return result
}

    
// IsServerErrorGeneratedColumnNonPrior check mysql error is "Generated column can refer only to generated columns defined prior to it." 
func IsServerErrorGeneratedColumnNonPrior(err error) bool {
    result := Isa(err, ER_GENERATED_COLUMN_NON_PRIOR)
    return result
}

    
// IsServerErrorDependentByGeneratedColumn check mysql error is "Column '%s' has a generated column dependency." 
func IsServerErrorDependentByGeneratedColumn(err error) bool {
    result := Isa(err, ER_DEPENDENT_BY_GENERATED_COLUMN)
    return result
}

    
// IsServerErrorGeneratedColumnRefAutoInc check mysql error is "Generated column '%s' cannot refer to auto-increment column." 
func IsServerErrorGeneratedColumnRefAutoInc(err error) bool {
    result := Isa(err, ER_GENERATED_COLUMN_REF_AUTO_INC)
    return result
}

    
// IsServerErrorFeatureNotAvailable check mysql error is "The '%s' feature is not available" 
func IsServerErrorFeatureNotAvailable(err error) bool {
    result := Isa(err, ER_FEATURE_NOT_AVAILABLE)
    return result
}

    
// IsServerErrorCantSetGtidMode check mysql error is "SET @@GLOBAL.GTID_MODE = %s is not allowed because %s." 
func IsServerErrorCantSetGtidMode(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_MODE)
    return result
}

    
// IsServerErrorCantUseAutoPositionWithGtidModeOff check mysql error is "The replication receiver thread%s cannot start in AUTO_POSITION mode: this server uses @@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorCantUseAutoPositionWithGtidModeOff(err error) bool {
    result := Isa(err, ER_CANT_USE_AUTO_POSITION_WITH_GTID_MODE_OFF)
    return result
}

    
// IsServerErrorCantReplicateAnonymousWithAutoPosition check mysql error is "Cannot replicate anonymous transaction when AUTO_POSITION = 1, at file %s, position %lld." 
func IsServerErrorCantReplicateAnonymousWithAutoPosition(err error) bool {
    result := Isa(err, ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION)
    return result
}

    
// IsServerErrorCantReplicateAnonymousWithGtidModeOn check mysql error is "Cannot replicate anonymous transaction when @@GLOBAL.GTID_MODE = ON, at file %s, position %lld." 
func IsServerErrorCantReplicateAnonymousWithGtidModeOn(err error) bool {
    result := Isa(err, ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON)
    return result
}

    
// IsServerErrorCantReplicateGtidWithGtidModeOff check mysql error is "Cannot replicate GTID-transaction when @@GLOBAL.GTID_MODE = OFF, at file %s, position %lld." 
func IsServerErrorCantReplicateGtidWithGtidModeOff(err error) bool {
    result := Isa(err, ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF)
    return result
}

    
// IsServerErrorCantSetEnforceGtidConsistencyOnWithOngoingGtidViolatingTransactions check mysql error is "Cannot set ENFORCE_GTID_CONSISTENCY = ON because there are ongoing transactions that violate GTID consistency." 
func IsServerErrorCantSetEnforceGtidConsistencyOnWithOngoingGtidViolatingTransactions(err error) bool {
    result := Isa(err, ER_CANT_SET_ENFORCE_GTID_CONSISTENCY_ON_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS)
    return result
}

    
// IsServerErrorSetEnforceGtidConsistencyWarnWithOngoingGtidViolatingTransactions check mysql error is "There are ongoing transactions that violate GTID consistency." 
func IsServerErrorSetEnforceGtidConsistencyWarnWithOngoingGtidViolatingTransactions(err error) bool {
    result := Isa(err, ER_SET_ENFORCE_GTID_CONSISTENCY_WARN_WITH_ONGOING_GTID_VIOLATING_TRANSACTIONS)
    return result
}

    
// IsServerErrorAccountHasBeenLocked check mysql error is "Access denied for user '%s'@'%s'. Account is locked." 
func IsServerErrorAccountHasBeenLocked(err error) bool {
    result := Isa(err, ER_ACCOUNT_HAS_BEEN_LOCKED)
    return result
}

    
// IsServerErrorWrongTablespaceName check mysql error is "Incorrect tablespace name `%s`" 
func IsServerErrorWrongTablespaceName(err error) bool {
    result := Isa(err, ER_WRONG_TABLESPACE_NAME)
    return result
}

    
// IsServerErrorTablespaceIsNotEmpty check mysql error is "Tablespace `%s` is not empty." 
func IsServerErrorTablespaceIsNotEmpty(err error) bool {
    result := Isa(err, ER_TABLESPACE_IS_NOT_EMPTY)
    return result
}

    
// IsServerErrorWrongFileName check mysql error is "Incorrect File Name '%s'." 
func IsServerErrorWrongFileName(err error) bool {
    result := Isa(err, ER_WRONG_FILE_NAME)
    return result
}

    
// IsServerErrorBoostGeometryInconsistentTurnsException check mysql error is "Inconsistent intersection points." 
func IsServerErrorBoostGeometryInconsistentTurnsException(err error) bool {
    result := Isa(err, ER_BOOST_GEOMETRY_INCONSISTENT_TURNS_EXCEPTION)
    return result
}

    
// IsServerErrorWarnOptimizerHintSyntaxError check mysql error is "Optimizer hint syntax error" 
func IsServerErrorWarnOptimizerHintSyntaxError(err error) bool {
    result := Isa(err, ER_WARN_OPTIMIZER_HINT_SYNTAX_ERROR)
    return result
}

    
// IsServerErrorWarnBadMaxExecutionTime check mysql error is "Unsupported MAX_EXECUTION_TIME" 
func IsServerErrorWarnBadMaxExecutionTime(err error) bool {
    result := Isa(err, ER_WARN_BAD_MAX_EXECUTION_TIME)
    return result
}

    
// IsServerErrorWarnUnsupportedMaxExecutionTime check mysql error is "MAX_EXECUTION_TIME hint is supported by top-level standalone SELECT statements only" 
func IsServerErrorWarnUnsupportedMaxExecutionTime(err error) bool {
    result := Isa(err, ER_WARN_UNSUPPORTED_MAX_EXECUTION_TIME)
    return result
}

    
// IsServerErrorWarnConflictingHint check mysql error is "Hint %s is ignored as conflicting/duplicated" 
func IsServerErrorWarnConflictingHint(err error) bool {
    result := Isa(err, ER_WARN_CONFLICTING_HINT)
    return result
}

    
// IsServerErrorWarnUnknownQbName check mysql error is "Query block name %s is not found for %s hint" 
func IsServerErrorWarnUnknownQbName(err error) bool {
    result := Isa(err, ER_WARN_UNKNOWN_QB_NAME)
    return result
}

    
// IsServerErrorUnresolvedHintName check mysql error is "Unresolved name %s for %s hint" 
func IsServerErrorUnresolvedHintName(err error) bool {
    result := Isa(err, ER_UNRESOLVED_HINT_NAME)
    return result
}

    
// IsServerErrorWarnDeprecatedSqlmodeUnset check mysql error is "Unsetting sql mode '%s' is deprecated. It will be made read-only in a future release." 
func IsServerErrorWarnDeprecatedSqlmodeUnset(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SQLMODE_UNSET)
    return result
}

    
// IsServerErrorWarnOnModifyingGtidExecutedTable check mysql error is "Please do not modify the %s table. This is a mysql internal system table to store GTIDs for committed transactions. Modifying it can lead to an inconsistent GTID state." 
func IsServerErrorWarnOnModifyingGtidExecutedTable(err error) bool {
    result := Isa(err, ER_WARN_ON_MODIFYING_GTID_EXECUTED_TABLE)
    return result
}

    
// IsServerErrorPluggableProtocolCommandNotSupported check mysql error is "Command not supported by pluggable protocols" 
func IsServerErrorPluggableProtocolCommandNotSupported(err error) bool {
    result := Isa(err, ER_PLUGGABLE_PROTOCOL_COMMAND_NOT_SUPPORTED)
    return result
}

    
// IsServerErrorLockingServiceWrongName check mysql error is "Incorrect locking service lock name '%s'." 
func IsServerErrorLockingServiceWrongName(err error) bool {
    result := Isa(err, ER_LOCKING_SERVICE_WRONG_NAME)
    return result
}

    
// IsServerErrorLockingServiceDeadlock check mysql error is "Deadlock found when trying to get locking service lock" 
func IsServerErrorLockingServiceDeadlock(err error) bool {
    result := Isa(err, ER_LOCKING_SERVICE_DEADLOCK)
    return result
}

    
// IsServerErrorLockingServiceTimeout check mysql error is "Service lock wait timeout exceeded." 
func IsServerErrorLockingServiceTimeout(err error) bool {
    result := Isa(err, ER_LOCKING_SERVICE_TIMEOUT)
    return result
}

    
// IsServerErrorGisMaxPointsInGeometryOverflowed check mysql error is "Parameter %s exceeds the maximum number of points in a geometry (%lu) in function %s." 
func IsServerErrorGisMaxPointsInGeometryOverflowed(err error) bool {
    result := Isa(err, ER_GIS_MAX_POINTS_IN_GEOMETRY_OVERFLOWED)
    return result
}

    
// IsServerErrorSqlModeMerged check mysql error is "'NO_ZERO_DATE', 'NO_ZERO_IN_DATE' and 'ERROR_FOR_DIVISION_BY_ZERO' sql modes should be used with strict mode. They will be merged with strict mode in a future release." 
func IsServerErrorSqlModeMerged(err error) bool {
    result := Isa(err, ER_SQL_MODE_MERGED)
    return result
}

    
// IsServerErrorVtokenPluginTokenMismatch check mysql error is "Version token mismatch for %.*s. Correct value %.*s" 
func IsServerErrorVtokenPluginTokenMismatch(err error) bool {
    result := Isa(err, ER_VTOKEN_PLUGIN_TOKEN_MISMATCH)
    return result
}

    
// IsServerErrorVtokenPluginTokenNotFound check mysql error is "Version token %.*s not found." 
func IsServerErrorVtokenPluginTokenNotFound(err error) bool {
    result := Isa(err, ER_VTOKEN_PLUGIN_TOKEN_NOT_FOUND)
    return result
}

    
// IsServerErrorCantSetVariableWhenOwningGtid check mysql error is "Variable %s cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK." 
func IsServerErrorCantSetVariableWhenOwningGtid(err error) bool {
    result := Isa(err, ER_CANT_SET_VARIABLE_WHEN_OWNING_GTID)
    return result
}

    
// IsServerErrorSlaveChannelOperationNotAllowed check mysql error is "%s cannot be performed on channel '%s'." 
func IsServerErrorSlaveChannelOperationNotAllowed(err error) bool {
    result := Isa(err, ER_SLAVE_CHANNEL_OPERATION_NOT_ALLOWED)
    return result
}

    
// IsServerErrorInvalidJsonText check mysql error is "Invalid JSON text: "%s" at position %u in value for column '%s'." 
func IsServerErrorInvalidJsonText(err error) bool {
    result := Isa(err, ER_INVALID_JSON_TEXT)
    return result
}

    
// IsServerErrorInvalidJsonTextInParam check mysql error is "Invalid JSON text in argument %u to function %s: "%s" at position %u.%s" 
func IsServerErrorInvalidJsonTextInParam(err error) bool {
    result := Isa(err, ER_INVALID_JSON_TEXT_IN_PARAM)
    return result
}

    
// IsServerErrorInvalidJsonBinaryData check mysql error is "The JSON binary value contains invalid data." 
func IsServerErrorInvalidJsonBinaryData(err error) bool {
    result := Isa(err, ER_INVALID_JSON_BINARY_DATA)
    return result
}

    
// IsServerErrorInvalidJsonPath check mysql error is "Invalid JSON path expression. The error is around character position %u.%s" 
func IsServerErrorInvalidJsonPath(err error) bool {
    result := Isa(err, ER_INVALID_JSON_PATH)
    return result
}

    
// IsServerErrorInvalidJsonCharset check mysql error is "Cannot create a JSON value from a string with CHARACTER SET '%s'." 
func IsServerErrorInvalidJsonCharset(err error) bool {
    result := Isa(err, ER_INVALID_JSON_CHARSET)
    return result
}

    
// IsServerErrorInvalidJsonCharsetInFunction check mysql error is "Invalid JSON character data provided to function %s: '%s'" 
func IsServerErrorInvalidJsonCharsetInFunction(err error) bool {
    result := Isa(err, ER_INVALID_JSON_CHARSET_IN_FUNCTION)
    return result
}

    
// IsServerErrorInvalidTypeForJson check mysql error is "Invalid data type for JSON data in argument %u to function %s" 
func IsServerErrorInvalidTypeForJson(err error) bool {
    result := Isa(err, ER_INVALID_TYPE_FOR_JSON)
    return result
}

    
// IsServerErrorInvalidCastToJson check mysql error is "Cannot CAST value to JSON." 
func IsServerErrorInvalidCastToJson(err error) bool {
    result := Isa(err, ER_INVALID_CAST_TO_JSON)
    return result
}

    
// IsServerErrorInvalidJsonPathCharset check mysql error is "A path expression must be encoded in the utf8 character set. The path expression '%s' is encoded in character set '%s'." 
func IsServerErrorInvalidJsonPathCharset(err error) bool {
    result := Isa(err, ER_INVALID_JSON_PATH_CHARSET)
    return result
}

    
// IsServerErrorInvalidJsonPathWildcard check mysql error is "In this situation, path expressions may not contain the * and ** tokens." 
func IsServerErrorInvalidJsonPathWildcard(err error) bool {
    result := Isa(err, ER_INVALID_JSON_PATH_WILDCARD)
    return result
}

    
// IsServerErrorJsonValueTooBig check mysql error is "The JSON value is too big to be stored in a JSON column." 
func IsServerErrorJsonValueTooBig(err error) bool {
    result := Isa(err, ER_JSON_VALUE_TOO_BIG)
    return result
}

    
// IsServerErrorJsonKeyTooBig check mysql error is "The JSON object contains a key name that is too long." 
func IsServerErrorJsonKeyTooBig(err error) bool {
    result := Isa(err, ER_JSON_KEY_TOO_BIG)
    return result
}

    
// IsServerErrorJsonUsedAsKey check mysql error is "JSON column '%s' cannot be used in key specification." 
func IsServerErrorJsonUsedAsKey(err error) bool {
    result := Isa(err, ER_JSON_USED_AS_KEY)
    return result
}

    
// IsServerErrorJsonVacuousPath check mysql error is "The path expression '$' is not allowed in this context." 
func IsServerErrorJsonVacuousPath(err error) bool {
    result := Isa(err, ER_JSON_VACUOUS_PATH)
    return result
}

    
// IsServerErrorJsonBadOneOrAllArg check mysql error is "The oneOrAll argument to %s may take these values: 'one' or 'all'." 
func IsServerErrorJsonBadOneOrAllArg(err error) bool {
    result := Isa(err, ER_JSON_BAD_ONE_OR_ALL_ARG)
    return result
}

    
// IsServerErrorNumericJsonValueOutOfRange check mysql error is "Out of range JSON value for CAST to %s%s from column %s at row %ld" 
func IsServerErrorNumericJsonValueOutOfRange(err error) bool {
    result := Isa(err, ER_NUMERIC_JSON_VALUE_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorInvalidJsonValueForCast check mysql error is "Invalid JSON value for CAST to %s%s from column %s at row %ld" 
func IsServerErrorInvalidJsonValueForCast(err error) bool {
    result := Isa(err, ER_INVALID_JSON_VALUE_FOR_CAST)
    return result
}

    
// IsServerErrorJsonDocumentTooDeep check mysql error is "The JSON document exceeds the maximum depth." 
func IsServerErrorJsonDocumentTooDeep(err error) bool {
    result := Isa(err, ER_JSON_DOCUMENT_TOO_DEEP)
    return result
}

    
// IsServerErrorJsonDocumentNullKey check mysql error is "JSON documents may not contain NULL member names." 
func IsServerErrorJsonDocumentNullKey(err error) bool {
    result := Isa(err, ER_JSON_DOCUMENT_NULL_KEY)
    return result
}

    
// IsServerErrorSecureTransportRequired check mysql error is "Connections using insecure transport are prohibited while --require_secure_transport=ON." 
func IsServerErrorSecureTransportRequired(err error) bool {
    result := Isa(err, ER_SECURE_TRANSPORT_REQUIRED)
    return result
}

    
// IsServerErrorNoSecureTransportsConfigured check mysql error is "No secure transports (SSL or Shared Memory) are configured, unable to set --require_secure_transport=ON." 
func IsServerErrorNoSecureTransportsConfigured(err error) bool {
    result := Isa(err, ER_NO_SECURE_TRANSPORTS_CONFIGURED)
    return result
}

    
// IsServerErrorDisabledStorageEngine check mysql error is "Storage engine %s is disabled (Table creation is disallowed)." 
func IsServerErrorDisabledStorageEngine(err error) bool {
    result := Isa(err, ER_DISABLED_STORAGE_ENGINE)
    return result
}

    
// IsServerErrorUserDoesNotExist check mysql error is "User %s does not exist." 
func IsServerErrorUserDoesNotExist(err error) bool {
    result := Isa(err, ER_USER_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorUserAlreadyExists check mysql error is "User %s already exists." 
func IsServerErrorUserAlreadyExists(err error) bool {
    result := Isa(err, ER_USER_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorAuditApiAbort check mysql error is "Aborted by Audit API ('%s'" 
func IsServerErrorAuditApiAbort(err error) bool {
    result := Isa(err, ER_AUDIT_API_ABORT)
    return result
}

    
// IsServerErrorInvalidJsonPathArrayCell check mysql error is "A path expression is not a path to a cell in an array." 
func IsServerErrorInvalidJsonPathArrayCell(err error) bool {
    result := Isa(err, ER_INVALID_JSON_PATH_ARRAY_CELL)
    return result
}

    
// IsServerErrorBufpoolResizeInprogress check mysql error is "Another buffer pool resize is already in progress." 
func IsServerErrorBufpoolResizeInprogress(err error) bool {
    result := Isa(err, ER_BUFPOOL_RESIZE_INPROGRESS)
    return result
}

    
// IsServerErrorFeatureDisabledSeeDoc check mysql error is "The '%s' feature is disabled" 
func IsServerErrorFeatureDisabledSeeDoc(err error) bool {
    result := Isa(err, ER_FEATURE_DISABLED_SEE_DOC)
    return result
}

    
// IsServerErrorServerIsntAvailable check mysql error is "Server isn't available" 
func IsServerErrorServerIsntAvailable(err error) bool {
    result := Isa(err, ER_SERVER_ISNT_AVAILABLE)
    return result
}

    
// IsServerErrorSessionWasKilled check mysql error is "Session was killed" 
func IsServerErrorSessionWasKilled(err error) bool {
    result := Isa(err, ER_SESSION_WAS_KILLED)
    return result
}

    
// IsServerErrorCapacityExceeded check mysql error is "Memory capacity of %llu bytes for '%s' exceeded. %s" 
func IsServerErrorCapacityExceeded(err error) bool {
    result := Isa(err, ER_CAPACITY_EXCEEDED)
    return result
}

    
// IsServerErrorCapacityExceededInRangeOptimizer check mysql error is "Range optimization was not done for this query." 
func IsServerErrorCapacityExceededInRangeOptimizer(err error) bool {
    result := Isa(err, ER_CAPACITY_EXCEEDED_IN_RANGE_OPTIMIZER)
    return result
}

    
// IsServerErrorTableNeedsUpgPart check mysql error is "Partitioning upgrade required. Please dump/reload to fix it or do: ALTER TABLE `%s`.`%s` UPGRADE PARTITIONING" 
func IsServerErrorTableNeedsUpgPart(err error) bool {
    result := Isa(err, ER_TABLE_NEEDS_UPG_PART)
    return result
}

    
// IsServerErrorCantWaitForExecutedGtidSetWhileOwningAGtid check mysql error is "The client holds ownership of the GTID %s. Therefore, WAIT_FOR_EXECUTED_GTID_SET cannot wait for this GTID." 
func IsServerErrorCantWaitForExecutedGtidSetWhileOwningAGtid(err error) bool {
    result := Isa(err, ER_CANT_WAIT_FOR_EXECUTED_GTID_SET_WHILE_OWNING_A_GTID)
    return result
}

    
// IsServerErrorCannotAddForeignBaseColVirtual check mysql error is "Cannot add foreign key on the base column of indexed virtual column." 
func IsServerErrorCannotAddForeignBaseColVirtual(err error) bool {
    result := Isa(err, ER_CANNOT_ADD_FOREIGN_BASE_COL_VIRTUAL)
    return result
}

    
// IsServerErrorCannotCreateVirtualIndexConstraint check mysql error is "Cannot create index on virtual column whose base column has foreign constraint." 
func IsServerErrorCannotCreateVirtualIndexConstraint(err error) bool {
    result := Isa(err, ER_CANNOT_CREATE_VIRTUAL_INDEX_CONSTRAINT)
    return result
}

    
// IsServerErrorErrorOnModifyingGtidExecutedTable check mysql error is "Please do not modify the %s table with an XA transaction. This is an internal system table used to store GTIDs for committed transactions. Although modifying it can lead to an inconsistent GTID state, if neccessary you can modify it with a non-XA transaction." 
func IsServerErrorErrorOnModifyingGtidExecutedTable(err error) bool {
    result := Isa(err, ER_ERROR_ON_MODIFYING_GTID_EXECUTED_TABLE)
    return result
}

    
// IsServerErrorLockRefusedByEngine check mysql error is "Lock acquisition refused by storage engine." 
func IsServerErrorLockRefusedByEngine(err error) bool {
    result := Isa(err, ER_LOCK_REFUSED_BY_ENGINE)
    return result
}

    
// IsServerErrorUnsupportedAlterOnlineOnVirtualColumn check mysql error is "ADD COLUMN col...VIRTUAL, ADD INDEX(col)" 
func IsServerErrorUnsupportedAlterOnlineOnVirtualColumn(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_ALTER_ONLINE_ON_VIRTUAL_COLUMN)
    return result
}

    
// IsServerErrorMasterKeyRotationNotSupportedBySe check mysql error is "Master key rotation is not supported by storage engine." 
func IsServerErrorMasterKeyRotationNotSupportedBySe(err error) bool {
    result := Isa(err, ER_MASTER_KEY_ROTATION_NOT_SUPPORTED_BY_SE)
    return result
}

    
// IsServerErrorMasterKeyRotationErrorBySe check mysql error is "Encryption key rotation error reported by SE: %s" 
func IsServerErrorMasterKeyRotationErrorBySe(err error) bool {
    result := Isa(err, ER_MASTER_KEY_ROTATION_ERROR_BY_SE)
    return result
}

    
// IsServerErrorMasterKeyRotationBinlogFailed check mysql error is "Write to binlog failed. However, master key rotation has been completed successfully." 
func IsServerErrorMasterKeyRotationBinlogFailed(err error) bool {
    result := Isa(err, ER_MASTER_KEY_ROTATION_BINLOG_FAILED)
    return result
}

    
// IsServerErrorMasterKeyRotationSeUnavailable check mysql error is "Storage engine is not available." 
func IsServerErrorMasterKeyRotationSeUnavailable(err error) bool {
    result := Isa(err, ER_MASTER_KEY_ROTATION_SE_UNAVAILABLE)
    return result
}

    
// IsServerErrorTablespaceCannotEncrypt check mysql error is "This tablespace can't be encrypted." 
func IsServerErrorTablespaceCannotEncrypt(err error) bool {
    result := Isa(err, ER_TABLESPACE_CANNOT_ENCRYPT)
    return result
}

    
// IsServerErrorInvalidEncryptionOption check mysql error is "Invalid encryption option." 
func IsServerErrorInvalidEncryptionOption(err error) bool {
    result := Isa(err, ER_INVALID_ENCRYPTION_OPTION)
    return result
}

    
// IsServerErrorCannotFindKeyInKeyring check mysql error is "Can't find master key from keyring, please check in the server log if a keyring plugin is loaded and initialized successfully." 
func IsServerErrorCannotFindKeyInKeyring(err error) bool {
    result := Isa(err, ER_CANNOT_FIND_KEY_IN_KEYRING)
    return result
}

    
// IsServerErrorCapacityExceededInParser check mysql error is "Parser bailed out for this query." 
func IsServerErrorCapacityExceededInParser(err error) bool {
    result := Isa(err, ER_CAPACITY_EXCEEDED_IN_PARSER)
    return result
}

    
// IsServerErrorUnsupportedAlterEncryptionInplace check mysql error is "Cannot alter encryption attribute by inplace algorithm." 
func IsServerErrorUnsupportedAlterEncryptionInplace(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_ALTER_ENCRYPTION_INPLACE)
    return result
}

    
// IsServerErrorKeyringUdfKeyringServiceError check mysql error is "Function '%s' failed because underlying keyring service returned an error. Please check if a keyring plugin is installed and that provided arguments are valid for the keyring you are using." 
func IsServerErrorKeyringUdfKeyringServiceError(err error) bool {
    result := Isa(err, ER_KEYRING_UDF_KEYRING_SERVICE_ERROR)
    return result
}

    
// IsServerErrorUserColumnOldLength check mysql error is "It seems that your db schema is old. The %s column is 77 characters long and should be 93 characters long. Please run mysql_upgrade." 
func IsServerErrorUserColumnOldLength(err error) bool {
    result := Isa(err, ER_USER_COLUMN_OLD_LENGTH)
    return result
}

    
// IsServerErrorCantResetMaster check mysql error is "RESET MASTER is not allowed because %s." 
func IsServerErrorCantResetMaster(err error) bool {
    result := Isa(err, ER_CANT_RESET_MASTER)
    return result
}

    
// IsServerErrorGroupReplicationMaxGroupSize check mysql error is "The START GROUP_REPLICATION command failed since the group already has 9 members." 
func IsServerErrorGroupReplicationMaxGroupSize(err error) bool {
    result := Isa(err, ER_GROUP_REPLICATION_MAX_GROUP_SIZE)
    return result
}

    
// IsServerErrorCannotAddForeignBaseColStored check mysql error is "Cannot add foreign key on the base column of stored column." 
func IsServerErrorCannotAddForeignBaseColStored(err error) bool {
    result := Isa(err, ER_CANNOT_ADD_FOREIGN_BASE_COL_STORED)
    return result
}

    
// IsServerErrorTableReferenced check mysql error is "Cannot complete the operation because table is referenced by another connection." 
func IsServerErrorTableReferenced(err error) bool {
    result := Isa(err, ER_TABLE_REFERENCED)
    return result
}

    
// IsServerErrorPartitionEngineDeprecatedForTable check mysql error is "The partition engine, used by table '%s.%s', is deprecated and will be removed in a future release. Please use native partitioning instead." 
func IsServerErrorPartitionEngineDeprecatedForTable(err error) bool {
    result := Isa(err, ER_PARTITION_ENGINE_DEPRECATED_FOR_TABLE)
    return result
}

    
// IsServerErrorWarnUsingGeomfromwkbToSetSridZero check mysql error is "%s(geometry) is deprecated and will be replaced by st_srid(geometry, 0) in a future version. Use %s(st_aswkb(geometry), 0) instead." 
func IsServerErrorWarnUsingGeomfromwkbToSetSridZero(err error) bool {
    result := Isa(err, ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID_ZERO)
    return result
}

    
// IsServerErrorWarnUsingGeomfromwkbToSetSrid check mysql error is "%s(geometry, srid) is deprecated and will be replaced by st_srid(geometry, srid) in a future version. Use %s(st_aswkb(geometry), srid) instead." 
func IsServerErrorWarnUsingGeomfromwkbToSetSrid(err error) bool {
    result := Isa(err, ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID)
    return result
}

    
// IsServerErrorXaRetry check mysql error is "The resource manager is not able to commit the transaction branch at this time. Please retry later." 
func IsServerErrorXaRetry(err error) bool {
    result := Isa(err, ER_XA_RETRY)
    return result
}

    
// IsServerErrorKeyringAwsUdfAwsKmsError check mysql error is "Function %s failed due to: %s." 
func IsServerErrorKeyringAwsUdfAwsKmsError(err error) bool {
    result := Isa(err, ER_KEYRING_AWS_UDF_AWS_KMS_ERROR)
    return result
}

    
// IsServerErrorBinlogUnsafeXa check mysql error is "Statement is unsafe because it is being used inside a XA transaction. Concurrent XA transactions may deadlock on slaves when replicated using statements." 
func IsServerErrorBinlogUnsafeXa(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_XA)
    return result
}

    
// IsServerErrorUdfError check mysql error is "%s UDF failed" 
func IsServerErrorUdfError(err error) bool {
    result := Isa(err, ER_UDF_ERROR)
    return result
}

    
// IsServerErrorKeyringMigrationFailure check mysql error is "Can not perform keyring migration : %s" 
func IsServerErrorKeyringMigrationFailure(err error) bool {
    result := Isa(err, ER_KEYRING_MIGRATION_FAILURE)
    return result
}

    
// IsServerErrorKeyringAccessDeniedError check mysql error is "Access denied" 
func IsServerErrorKeyringAccessDeniedError(err error) bool {
    result := Isa(err, ER_KEYRING_ACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorKeyringMigrationStatus check mysql error is "Keyring migration %s." 
func IsServerErrorKeyringMigrationStatus(err error) bool {
    result := Isa(err, ER_KEYRING_MIGRATION_STATUS)
    return result
}

    
// IsServerErrorPluginFailedToOpenTables check mysql error is "Failed to open the %s filter tables." 
func IsServerErrorPluginFailedToOpenTables(err error) bool {
    result := Isa(err, ER_PLUGIN_FAILED_TO_OPEN_TABLES)
    return result
}

    
// IsServerErrorPluginFailedToOpenTable check mysql error is "Failed to open '%s.%s' %s table." 
func IsServerErrorPluginFailedToOpenTable(err error) bool {
    result := Isa(err, ER_PLUGIN_FAILED_TO_OPEN_TABLE)
    return result
}

    
// IsServerErrorAuditLogNoKeyringPluginInstalled check mysql error is "No keyring plugin installed." 
func IsServerErrorAuditLogNoKeyringPluginInstalled(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED)
    return result
}

    
// IsServerErrorAuditLogEncryptionPasswordHasNotBeenSet check mysql error is "Audit log encryption password has not been set" 
func IsServerErrorAuditLogEncryptionPasswordHasNotBeenSet(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET)
    return result
}

    
// IsServerErrorAuditLogCouldNotCreateAesKey check mysql error is "Could not create AES key. OpenSSL's EVP_BytesToKey function failed." 
func IsServerErrorAuditLogCouldNotCreateAesKey(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY)
    return result
}

    
// IsServerErrorAuditLogEncryptionPasswordCannotBeFetched check mysql error is "Audit log encryption password cannot be fetched from the keyring. Password used so far is used for encryption." 
func IsServerErrorAuditLogEncryptionPasswordCannotBeFetched(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED)
    return result
}

    
// IsServerErrorAuditLogJsonFilteringNotEnabled check mysql error is "Audit Log filtering has not been installed." 
func IsServerErrorAuditLogJsonFilteringNotEnabled(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED)
    return result
}

    
// IsServerErrorAuditLogUdfInsufficientPrivilege check mysql error is "Request ignored for '%s'@'%s'. SUPER_ACL needed to perform operation" 
func IsServerErrorAuditLogUdfInsufficientPrivilege(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE)
    return result
}

    
// IsServerErrorAuditLogSuperPrivilegeRequired check mysql error is "SUPER privilege required for '%s'@'%s' user." 
func IsServerErrorAuditLogSuperPrivilegeRequired(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED)
    return result
}

    
// IsServerErrorCouldNotReinitializeAuditLogFilters check mysql error is "Could not reinitialize audit log filters." 
func IsServerErrorCouldNotReinitializeAuditLogFilters(err error) bool {
    result := Isa(err, ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS)
    return result
}

    
// IsServerErrorAuditLogUdfInvalidArgumentType check mysql error is "Invalid argument type" 
func IsServerErrorAuditLogUdfInvalidArgumentType(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE)
    return result
}

    
// IsServerErrorAuditLogUdfInvalidArgumentCount check mysql error is "Invalid argument count" 
func IsServerErrorAuditLogUdfInvalidArgumentCount(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT)
    return result
}

    
// IsServerErrorAuditLogHasNotBeenInstalled check mysql error is "audit_log plugin has not been installed using INSTALL PLUGIN syntax." 
func IsServerErrorAuditLogHasNotBeenInstalled(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED)
    return result
}

    
// IsServerErrorAuditLogUdfReadInvalidMaxArrayLengthArgType check mysql error is "Invalid "max_array_length" argument type." 
func IsServerErrorAuditLogUdfReadInvalidMaxArrayLengthArgType(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE)
    return result
}

    
// IsServerErrorAuditLogUdfReadInvalidMaxArrayLengthArgValue check mysql error is "Invalid "max_array_length" argument value." 
func IsServerErrorAuditLogUdfReadInvalidMaxArrayLengthArgValue(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_VALUE)
    return result
}

    
// IsServerErrorAuditLogJsonFilterParsingError check mysql error is "%s" 
func IsServerErrorAuditLogJsonFilterParsingError(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR)
    return result
}

    
// IsServerErrorAuditLogJsonFilterNameCannotBeEmpty check mysql error is "Filter name cannot be empty." 
func IsServerErrorAuditLogJsonFilterNameCannotBeEmpty(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY)
    return result
}

    
// IsServerErrorAuditLogJsonUserNameCannotBeEmpty check mysql error is "User cannot be empty." 
func IsServerErrorAuditLogJsonUserNameCannotBeEmpty(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY)
    return result
}

    
// IsServerErrorAuditLogJsonFilterDoesNotExists check mysql error is "Specified filter has not been found." 
func IsServerErrorAuditLogJsonFilterDoesNotExists(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXISTS)
    return result
}

    
// IsServerErrorAuditLogUserFirstCharacterMustBeAlphanumeric check mysql error is "First character of the user name must be alphanumeric." 
func IsServerErrorAuditLogUserFirstCharacterMustBeAlphanumeric(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC)
    return result
}

    
// IsServerErrorAuditLogUserNameInvalidCharacter check mysql error is "Invalid character in the user name." 
func IsServerErrorAuditLogUserNameInvalidCharacter(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER)
    return result
}

    
// IsServerErrorAuditLogHostNameInvalidCharacter check mysql error is "Invalid character in the host name." 
func IsServerErrorAuditLogHostNameInvalidCharacter(err error) bool {
    result := Isa(err, ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER)
    return result
}

    
// IsWarnDeprecatedMaxdbSqlModeForTimestamp check mysql error is "With the MAXDB SQL mode enabled, TIMESTAMP is identical with DATETIME. The MAXDB SQL mode is deprecated and will be removed in a future release. Please disable the MAXDB SQL mode and use DATETIME instead." 
func IsWarnDeprecatedMaxdbSqlModeForTimestamp(err error) bool {
    result := Isa(err, WARN_DEPRECATED_MAXDB_SQL_MODE_FOR_TIMESTAMP)
    return result
}

    
// IsServerErrorXaReplicationFilters check mysql error is "The use of replication filters with XA transactions is not supported, and can lead to an undefined state in the replication slave." 
func IsServerErrorXaReplicationFilters(err error) bool {
    result := Isa(err, ER_XA_REPLICATION_FILTERS)
    return result
}

    
// IsServerErrorCantOpenErrorLog check mysql error is "Could not open file '%s' for error logging%s%s" 
func IsServerErrorCantOpenErrorLog(err error) bool {
    result := Isa(err, ER_CANT_OPEN_ERROR_LOG)
    return result
}

    
// IsServerErrorGroupingOnTimestampInDst check mysql error is "Grouping on temporal is non-deterministic for timezones having DST. Please consider switching to UTC for this query." 
func IsServerErrorGroupingOnTimestampInDst(err error) bool {
    result := Isa(err, ER_GROUPING_ON_TIMESTAMP_IN_DST)
    return result
}

    
// IsServerErrorCantStartServerNamedPipe check mysql error is "Can't start server : Named Pipe "%s" already in use." 
func IsServerErrorCantStartServerNamedPipe(err error) bool {
    result := Isa(err, ER_CANT_START_SERVER_NAMED_PIPE)
    return result
}

    
// IsClientErrorUnknownError check mysql error is "Unknown MySQL error " 
func IsClientErrorUnknownError(err error) bool {
    result := Isa(err, CR_UNKNOWN_ERROR)
    return result
}

    
// IsClientErrorSocketCreateError check mysql error is "Can't create UNIX socket (%d) " 
func IsClientErrorSocketCreateError(err error) bool {
    result := Isa(err, CR_SOCKET_CREATE_ERROR)
    return result
}

    
// IsClientErrorConnectionError check mysql error is "Can't connect to local MySQL server through socket '%s' (%d) " 
func IsClientErrorConnectionError(err error) bool {
    result := Isa(err, CR_CONNECTION_ERROR)
    return result
}

    
// IsClientErrorConnHostError check mysql error is "Can't connect to MySQL server on '%s' (%d) " 
func IsClientErrorConnHostError(err error) bool {
    result := Isa(err, CR_CONN_HOST_ERROR)
    return result
}

    
// IsClientErrorIpsockError check mysql error is "Can't create TCP/IP socket (%d) " 
func IsClientErrorIpsockError(err error) bool {
    result := Isa(err, CR_IPSOCK_ERROR)
    return result
}

    
// IsClientErrorUnknownHost check mysql error is "Unknown MySQL server host '%s' (%d) " 
func IsClientErrorUnknownHost(err error) bool {
    result := Isa(err, CR_UNKNOWN_HOST)
    return result
}

    
// IsClientErrorServerGoneError check mysql error is "MySQL server has gone away " 
func IsClientErrorServerGoneError(err error) bool {
    result := Isa(err, CR_SERVER_GONE_ERROR)
    return result
}

    
// IsClientErrorVersionError check mysql error is "Protocol mismatch" 
func IsClientErrorVersionError(err error) bool {
    result := Isa(err, CR_VERSION_ERROR)
    return result
}

    
// IsClientErrorOutOfMemory check mysql error is "MySQL client ran out of memory " 
func IsClientErrorOutOfMemory(err error) bool {
    result := Isa(err, CR_OUT_OF_MEMORY)
    return result
}

    
// IsClientErrorWrongHostInfo check mysql error is "Wrong host info " 
func IsClientErrorWrongHostInfo(err error) bool {
    result := Isa(err, CR_WRONG_HOST_INFO)
    return result
}

    
// IsClientErrorLocalhostConnection check mysql error is "Localhost via UNIX socket " 
func IsClientErrorLocalhostConnection(err error) bool {
    result := Isa(err, CR_LOCALHOST_CONNECTION)
    return result
}

    
// IsClientErrorTcpConnection check mysql error is "%s via TCP/IP " 
func IsClientErrorTcpConnection(err error) bool {
    result := Isa(err, CR_TCP_CONNECTION)
    return result
}

    
// IsClientErrorServerHandshakeErr check mysql error is "Error in server handshake " 
func IsClientErrorServerHandshakeErr(err error) bool {
    result := Isa(err, CR_SERVER_HANDSHAKE_ERR)
    return result
}

    
// IsClientErrorServerLost check mysql error is "Lost connection to MySQL server during query " 
func IsClientErrorServerLost(err error) bool {
    result := Isa(err, CR_SERVER_LOST)
    return result
}

    
// IsClientErrorCommandsOutOfSync check mysql error is "Commands out of sync" 
func IsClientErrorCommandsOutOfSync(err error) bool {
    result := Isa(err, CR_COMMANDS_OUT_OF_SYNC)
    return result
}

    
// IsClientErrorNamedpipeConnection check mysql error is "Named pipe: %s " 
func IsClientErrorNamedpipeConnection(err error) bool {
    result := Isa(err, CR_NAMEDPIPE_CONNECTION)
    return result
}

    
// IsClientErrorNamedpipewaitError check mysql error is "Can't wait for named pipe to host: %s pipe: %s (%lu) " 
func IsClientErrorNamedpipewaitError(err error) bool {
    result := Isa(err, CR_NAMEDPIPEWAIT_ERROR)
    return result
}

    
// IsClientErrorNamedpipeopenError check mysql error is "Can't open named pipe to host: %s pipe: %s (%lu) " 
func IsClientErrorNamedpipeopenError(err error) bool {
    result := Isa(err, CR_NAMEDPIPEOPEN_ERROR)
    return result
}

    
// IsClientErrorNamedpipesetstateError check mysql error is "Can't set state of named pipe to host: %s pipe: %s (%lu) " 
func IsClientErrorNamedpipesetstateError(err error) bool {
    result := Isa(err, CR_NAMEDPIPESETSTATE_ERROR)
    return result
}

    
// IsClientErrorCantReadCharset check mysql error is "Can't initialize character set %s (path: %s) " 
func IsClientErrorCantReadCharset(err error) bool {
    result := Isa(err, CR_CANT_READ_CHARSET)
    return result
}

    
// IsClientErrorNetPacketTooLarge check mysql error is "Got packet bigger than 'max_allowed_packet' bytes " 
func IsClientErrorNetPacketTooLarge(err error) bool {
    result := Isa(err, CR_NET_PACKET_TOO_LARGE)
    return result
}

    
// IsClientErrorEmbeddedConnection check mysql error is "Embedded server " 
func IsClientErrorEmbeddedConnection(err error) bool {
    result := Isa(err, CR_EMBEDDED_CONNECTION)
    return result
}

    
// IsClientErrorProbeSlaveStatus check mysql error is "Error on SHOW SLAVE STATUS: " 
func IsClientErrorProbeSlaveStatus(err error) bool {
    result := Isa(err, CR_PROBE_SLAVE_STATUS)
    return result
}

    
// IsClientErrorProbeSlaveHosts check mysql error is "Error on SHOW SLAVE HOSTS: " 
func IsClientErrorProbeSlaveHosts(err error) bool {
    result := Isa(err, CR_PROBE_SLAVE_HOSTS)
    return result
}

    
// IsClientErrorProbeSlaveConnect check mysql error is "Error connecting to slave: " 
func IsClientErrorProbeSlaveConnect(err error) bool {
    result := Isa(err, CR_PROBE_SLAVE_CONNECT)
    return result
}

    
// IsClientErrorProbeMasterConnect check mysql error is "Error connecting to master: " 
func IsClientErrorProbeMasterConnect(err error) bool {
    result := Isa(err, CR_PROBE_MASTER_CONNECT)
    return result
}

    
// IsClientErrorSslConnectionError check mysql error is "SSL connection error: %s " 
func IsClientErrorSslConnectionError(err error) bool {
    result := Isa(err, CR_SSL_CONNECTION_ERROR)
    return result
}

    
// IsClientErrorMalformedPacket check mysql error is "Malformed packet " 
func IsClientErrorMalformedPacket(err error) bool {
    result := Isa(err, CR_MALFORMED_PACKET)
    return result
}

    
// IsClientErrorWrongLicense check mysql error is "This client library is licensed only for use with MySQL servers having '%s' license " 
func IsClientErrorWrongLicense(err error) bool {
    result := Isa(err, CR_WRONG_LICENSE)
    return result
}

    
// IsClientErrorNullPointer check mysql error is "Invalid use of null pointer " 
func IsClientErrorNullPointer(err error) bool {
    result := Isa(err, CR_NULL_POINTER)
    return result
}

    
// IsClientErrorNoPrepareStmt check mysql error is "Statement not prepared " 
func IsClientErrorNoPrepareStmt(err error) bool {
    result := Isa(err, CR_NO_PREPARE_STMT)
    return result
}

    
// IsClientErrorParamsNotBound check mysql error is "No data supplied for parameters in prepared statement " 
func IsClientErrorParamsNotBound(err error) bool {
    result := Isa(err, CR_PARAMS_NOT_BOUND)
    return result
}

    
// IsClientErrorDataTruncated check mysql error is "Data truncated " 
func IsClientErrorDataTruncated(err error) bool {
    result := Isa(err, CR_DATA_TRUNCATED)
    return result
}

    
// IsClientErrorNoParametersExists check mysql error is "No parameters exist in the statement " 
func IsClientErrorNoParametersExists(err error) bool {
    result := Isa(err, CR_NO_PARAMETERS_EXISTS)
    return result
}

    
// IsClientErrorInvalidParameterNo check mysql error is "Invalid parameter number" 
func IsClientErrorInvalidParameterNo(err error) bool {
    result := Isa(err, CR_INVALID_PARAMETER_NO)
    return result
}

    
// IsClientErrorInvalidBufferUse check mysql error is "Can't send long data for non-string/non-binary data types (parameter: %d) " 
func IsClientErrorInvalidBufferUse(err error) bool {
    result := Isa(err, CR_INVALID_BUFFER_USE)
    return result
}

    
// IsClientErrorUnsupportedParamType check mysql error is "Using unsupported buffer type: %d (parameter: %d) " 
func IsClientErrorUnsupportedParamType(err error) bool {
    result := Isa(err, CR_UNSUPPORTED_PARAM_TYPE)
    return result
}

    
// IsClientErrorSharedMemoryConnection check mysql error is "Shared memory: %s " 
func IsClientErrorSharedMemoryConnection(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECTION)
    return result
}

    
// IsClientErrorSharedMemoryConnectRequestError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryConnectRequestError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECT_REQUEST_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryConnectAnswerError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryConnectAnswerError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECT_ANSWER_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryConnectFileMapError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryConnectFileMapError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECT_FILE_MAP_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryConnectMapError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryConnectMapError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECT_MAP_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryFileMapError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryFileMapError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_FILE_MAP_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryMapError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryMapError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_MAP_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryEventError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryEventError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_EVENT_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryConnectAbandonedError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryConnectAbandonedError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECT_ABANDONED_ERROR)
    return result
}

    
// IsClientErrorSharedMemoryConnectSetError check mysql error is "Can't open shared memory" 
func IsClientErrorSharedMemoryConnectSetError(err error) bool {
    result := Isa(err, CR_SHARED_MEMORY_CONNECT_SET_ERROR)
    return result
}

    
// IsClientErrorConnUnknowProtocol check mysql error is "Wrong or unknown protocol " 
func IsClientErrorConnUnknowProtocol(err error) bool {
    result := Isa(err, CR_CONN_UNKNOW_PROTOCOL)
    return result
}

    
// IsClientErrorInvalidConnHandle check mysql error is "Invalid connection handle " 
func IsClientErrorInvalidConnHandle(err error) bool {
    result := Isa(err, CR_INVALID_CONN_HANDLE)
    return result
}

    
// IsClientErrorSecureAuth check mysql error is "Connection using old (pre-4.1.1) authentication protocol refused (client option 'secure_auth' enabled)" 
func IsClientErrorSecureAuth(err error) bool {
    result := Isa(err, CR_SECURE_AUTH)
    return result
}

    
// IsClientErrorUnused1 check mysql error is "Connection using old (pre-4.1.1) authentication protocol refused (client option 'secure_auth' enabled)" 
func IsClientErrorUnused1(err error) bool {
    result := Isa(err, CR_UNUSED_1)
    return result
}

    
// IsClientErrorFetchCanceled check mysql error is "Row retrieval was canceled by mysql_stmt_close() call " 
func IsClientErrorFetchCanceled(err error) bool {
    result := Isa(err, CR_FETCH_CANCELED)
    return result
}

    
// IsClientErrorNoData check mysql error is "Attempt to read column without prior row fetch " 
func IsClientErrorNoData(err error) bool {
    result := Isa(err, CR_NO_DATA)
    return result
}

    
// IsClientErrorNoStmtMetadata check mysql error is "Prepared statement contains no metadata " 
func IsClientErrorNoStmtMetadata(err error) bool {
    result := Isa(err, CR_NO_STMT_METADATA)
    return result
}

    
// IsClientErrorNoResultSet check mysql error is "Attempt to read a row while there is no result set associated with the statement " 
func IsClientErrorNoResultSet(err error) bool {
    result := Isa(err, CR_NO_RESULT_SET)
    return result
}

    
// IsClientErrorNotImplemented check mysql error is "This feature is not implemented yet " 
func IsClientErrorNotImplemented(err error) bool {
    result := Isa(err, CR_NOT_IMPLEMENTED)
    return result
}

    
// IsClientErrorServerLostExtended check mysql error is "Lost connection to MySQL server at '%s', system error: %d " 
func IsClientErrorServerLostExtended(err error) bool {
    result := Isa(err, CR_SERVER_LOST_EXTENDED)
    return result
}

    
// IsClientErrorStmtClosed check mysql error is "Statement closed indirectly because of a preceding %s() call " 
func IsClientErrorStmtClosed(err error) bool {
    result := Isa(err, CR_STMT_CLOSED)
    return result
}

    
// IsClientErrorNewStmtMetadata check mysql error is "The number of columns in the result set differs from the number of bound buffers. You must reset the statement, rebind the result set columns, and execute the statement again " 
func IsClientErrorNewStmtMetadata(err error) bool {
    result := Isa(err, CR_NEW_STMT_METADATA)
    return result
}

    
// IsClientErrorAlreadyConnected check mysql error is "This handle is already connected. Use a separate handle for each connection. " 
func IsClientErrorAlreadyConnected(err error) bool {
    result := Isa(err, CR_ALREADY_CONNECTED)
    return result
}

    
// IsClientErrorAuthPluginCannotLoad check mysql error is "Authentication plugin '%s' cannot be loaded: %s " 
func IsClientErrorAuthPluginCannotLoad(err error) bool {
    result := Isa(err, CR_AUTH_PLUGIN_CANNOT_LOAD)
    return result
}

    
// IsClientErrorDuplicateConnectionAttr check mysql error is "There is an attribute with the same name already" 
func IsClientErrorDuplicateConnectionAttr(err error) bool {
    result := Isa(err, CR_DUPLICATE_CONNECTION_ATTR)
    return result
}

    
// IsClientErrorAuthPluginErr check mysql error is "Authentication plugin '%s' reported error: %s" 
func IsClientErrorAuthPluginErr(err error) bool {
    result := Isa(err, CR_AUTH_PLUGIN_ERR)
    return result
}

    
// IsClientErrorInsecureApiErr check mysql error is "Insecure API function call: '%s' Use instead: '%s'" 
func IsClientErrorInsecureApiErr(err error) bool {
    result := Isa(err, CR_INSECURE_API_ERR)
    return result
}

    
// IsGlobalErrorCantcreatefile check mysql error is "Can't create/write to file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorCantcreatefile(err error) bool {
    result := Isa(err, EE_CANTCREATEFILE)
    return result
}

    
// IsGlobalErrorRead check mysql error is "Error reading file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorRead(err error) bool {
    result := Isa(err, EE_READ)
    return result
}

    
// IsGlobalErrorWrite check mysql error is "Error writing file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorWrite(err error) bool {
    result := Isa(err, EE_WRITE)
    return result
}

    
// IsGlobalErrorBadclose check mysql error is "Error on close of '%s' (Errcode: %d - %s) " 
func IsGlobalErrorBadclose(err error) bool {
    result := Isa(err, EE_BADCLOSE)
    return result
}

    
// IsGlobalErrorOutofmemory check mysql error is "Out of memory (Needed %u bytes) " 
func IsGlobalErrorOutofmemory(err error) bool {
    result := Isa(err, EE_OUTOFMEMORY)
    return result
}

    
// IsGlobalErrorDelete check mysql error is "Error on delete of '%s' (Errcode: %d - %s) " 
func IsGlobalErrorDelete(err error) bool {
    result := Isa(err, EE_DELETE)
    return result
}

    
// IsGlobalErrorLink check mysql error is "Error on rename of '%s' to '%s' (Errcode: %d - %s) " 
func IsGlobalErrorLink(err error) bool {
    result := Isa(err, EE_LINK)
    return result
}

    
// IsGlobalErrorEoferr check mysql error is "Unexpected EOF found when reading file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorEoferr(err error) bool {
    result := Isa(err, EE_EOFERR)
    return result
}

    
// IsGlobalErrorCantlock check mysql error is "Can't lock file (Errcode: %d - %s) " 
func IsGlobalErrorCantlock(err error) bool {
    result := Isa(err, EE_CANTLOCK)
    return result
}

    
// IsGlobalErrorCantunlock check mysql error is "Can't unlock file (Errcode: %d - %s) " 
func IsGlobalErrorCantunlock(err error) bool {
    result := Isa(err, EE_CANTUNLOCK)
    return result
}

    
// IsGlobalErrorDir check mysql error is "Can't read dir of '%s' (Errcode: %d - %s) " 
func IsGlobalErrorDir(err error) bool {
    result := Isa(err, EE_DIR)
    return result
}

    
// IsGlobalErrorStat check mysql error is "Can't get stat of '%s' (Errcode: %d - %s) " 
func IsGlobalErrorStat(err error) bool {
    result := Isa(err, EE_STAT)
    return result
}

    
// IsGlobalErrorCantChsize check mysql error is "Can't change size of file (Errcode: %d - %s) " 
func IsGlobalErrorCantChsize(err error) bool {
    result := Isa(err, EE_CANT_CHSIZE)
    return result
}

    
// IsGlobalErrorCantOpenStream check mysql error is "Can't open stream from handle (Errcode: %d - %s) " 
func IsGlobalErrorCantOpenStream(err error) bool {
    result := Isa(err, EE_CANT_OPEN_STREAM)
    return result
}

    
// IsGlobalErrorGetwd check mysql error is "Can't get working directory (Errcode: %d - %s) " 
func IsGlobalErrorGetwd(err error) bool {
    result := Isa(err, EE_GETWD)
    return result
}

    
// IsGlobalErrorSetwd check mysql error is "Can't change dir to '%s' (Errcode: %d - %s) " 
func IsGlobalErrorSetwd(err error) bool {
    result := Isa(err, EE_SETWD)
    return result
}

    
// IsGlobalErrorLinkWarning check mysql error is "Warning: '%s' had %d links " 
func IsGlobalErrorLinkWarning(err error) bool {
    result := Isa(err, EE_LINK_WARNING)
    return result
}

    
// IsGlobalErrorOpenWarning check mysql error is "Warning: %d files and %d streams is left open " 
func IsGlobalErrorOpenWarning(err error) bool {
    result := Isa(err, EE_OPEN_WARNING)
    return result
}

    
// IsGlobalErrorDiskFull check mysql error is "Disk is full writing '%s' (Errcode: %d - %s). Waiting for someone to free space... " 
func IsGlobalErrorDiskFull(err error) bool {
    result := Isa(err, EE_DISK_FULL)
    return result
}

    
// IsGlobalErrorCantMkdir check mysql error is "Can't create directory '%s' (Errcode: %d - %s) " 
func IsGlobalErrorCantMkdir(err error) bool {
    result := Isa(err, EE_CANT_MKDIR)
    return result
}

    
// IsGlobalErrorUnknownCharset check mysql error is "Character set '%s' is not a compiled character set and is not specified in the '%s' file " 
func IsGlobalErrorUnknownCharset(err error) bool {
    result := Isa(err, EE_UNKNOWN_CHARSET)
    return result
}

    
// IsGlobalErrorOutOfFileresources check mysql error is "Out of resources when opening file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorOutOfFileresources(err error) bool {
    result := Isa(err, EE_OUT_OF_FILERESOURCES)
    return result
}

    
// IsGlobalErrorCantReadlink check mysql error is "Can't read value for symlink '%s' (Error %d - %s) " 
func IsGlobalErrorCantReadlink(err error) bool {
    result := Isa(err, EE_CANT_READLINK)
    return result
}

    
// IsGlobalErrorCantSymlink check mysql error is "Can't create symlink '%s' pointing at '%s' (Error %d - %s) " 
func IsGlobalErrorCantSymlink(err error) bool {
    result := Isa(err, EE_CANT_SYMLINK)
    return result
}

    
// IsGlobalErrorRealpath check mysql error is "Error on realpath() on '%s' (Error %d - %s) " 
func IsGlobalErrorRealpath(err error) bool {
    result := Isa(err, EE_REALPATH)
    return result
}

    
// IsGlobalErrorSync check mysql error is "Can't sync file '%s' to disk (Errcode: %d - %s) " 
func IsGlobalErrorSync(err error) bool {
    result := Isa(err, EE_SYNC)
    return result
}

    
// IsGlobalErrorUnknownCollation check mysql error is "Collation '%s' is not a compiled collation and is not specified in the '%s' file " 
func IsGlobalErrorUnknownCollation(err error) bool {
    result := Isa(err, EE_UNKNOWN_COLLATION)
    return result
}

    
// IsGlobalErrorFilenotfound check mysql error is "File '%s' not found (Errcode: %d - %s) " 
func IsGlobalErrorFilenotfound(err error) bool {
    result := Isa(err, EE_FILENOTFOUND)
    return result
}

    
// IsGlobalErrorFileNotClosed check mysql error is "File '%s' (fileno: %d) was not closed " 
func IsGlobalErrorFileNotClosed(err error) bool {
    result := Isa(err, EE_FILE_NOT_CLOSED)
    return result
}

    
// IsGlobalErrorChangeOwnership check mysql error is "Can't change ownership of the file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorChangeOwnership(err error) bool {
    result := Isa(err, EE_CHANGE_OWNERSHIP)
    return result
}

    
// IsGlobalErrorChangePermissions check mysql error is "Can't change permissions of the file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorChangePermissions(err error) bool {
    result := Isa(err, EE_CHANGE_PERMISSIONS)
    return result
}

    
// IsGlobalErrorCantSeek check mysql error is "Can't seek in file '%s' (Errcode: %d - %s) " 
func IsGlobalErrorCantSeek(err error) bool {
    result := Isa(err, EE_CANT_SEEK)
    return result
}

    
// IsGlobalErrorCapacityExceeded check mysql error is "Memory capacity exceeded (capacity %llu bytes)" 
func IsGlobalErrorCapacityExceeded(err error) bool {
    result := Isa(err, EE_CAPACITY_EXCEEDED)
    return result
}




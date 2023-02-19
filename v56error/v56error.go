package v56error

import (
    "github.com/ytnobody/gomysqlerror/definition"
    "github.com/imdario/mergo"
)

var ErrorMap = map[int64]definition.ErrorDefinition{
    1000: {ErrorNumber:1000, ErrorType:"ServerError", Symbol:"ER_HASHCHK", SQLState:"HY000", Message:"hashchk", Description:"Unused.", MySQLVersion:"5.6"},
    1001: {ErrorNumber:1001, ErrorType:"ServerError", Symbol:"ER_NISAMCHK", SQLState:"HY000", Message:"isamchk", Description:"Unused.", MySQLVersion:"5.6"},
    1002: {ErrorNumber:1002, ErrorType:"ServerError", Symbol:"ER_NO", SQLState:"HY000", Message:"NO", Description:"Used in the construction of other messages.", MySQLVersion:"5.6"},
    1003: {ErrorNumber:1003, ErrorType:"ServerError", Symbol:"ER_YES", SQLState:"HY000", Message:"YES", Description:"Extended EXPLAIN format generatesNote messages. ER_YES is used inthe Code column for these messages insubsequent SHOW WARNINGS output.", MySQLVersion:"5.6"},
    1004: {ErrorNumber:1004, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_FILE", SQLState:"HY000", Message:"Can't create file '%s' (errno: %d - %s)", Description:"destinationfile already exists but is not writeable.", MySQLVersion:"5.6"},
    1005: {ErrorNumber:1005, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_TABLE", SQLState:"HY000", Message:"Can't create table '%s' (errno: %d)", Description:"InnoDB reports this error when a table cannotbe created. If the error message refers to error 150, tablecreation failed because aforeign keyconstraint was not correctly formed. If the error messagerefers to error −1, table creation probably failed becausethe table includes a column name that matched the name of aninternal InnoDB table.", MySQLVersion:"5.6"},
    1006: {ErrorNumber:1006, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_DB", SQLState:"HY000", Message:"Can't create database '%s' (errno: %d)", Description:"", MySQLVersion:"5.6"},
    1007: {ErrorNumber:1007, ErrorType:"ServerError", Symbol:"ER_DB_CREATE_EXISTS", SQLState:"HY000", Message:"Can't create database '%s'", Description:"Drop the database first if you really want to replace an existingdatabase, or add an IF NOT EXISTS clause to theCREATE DATABASE statement if toretain an existing database without having the statement producean error.", MySQLVersion:"5.6"},
    1008: {ErrorNumber:1008, ErrorType:"ServerError", Symbol:"ER_DB_DROP_EXISTS", SQLState:"HY000", Message:"Can't drop database '%s'", Description:"database doesn't exist", MySQLVersion:"5.6"},
    1009: {ErrorNumber:1009, ErrorType:"ServerError", Symbol:"ER_DB_DROP_DELETE", SQLState:"HY000", Message:"Error dropping database (can't delete '%s', errno: %d)", Description:"", MySQLVersion:"5.6"},
    1010: {ErrorNumber:1010, ErrorType:"ServerError", Symbol:"ER_DB_DROP_RMDIR", SQLState:"HY000", Message:"Error dropping database (can't rmdir '%s', errno: %d)", Description:"", MySQLVersion:"5.6"},
    1011: {ErrorNumber:1011, ErrorType:"ServerError", Symbol:"ER_CANT_DELETE_FILE", SQLState:"HY000", Message:"Error on delete of '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1012: {ErrorNumber:1012, ErrorType:"ServerError", Symbol:"ER_CANT_FIND_SYSTEM_REC", SQLState:"HY000", Message:"Can't read record in system table", Description:"Returned by InnoDB for attempts to accessInnoDB INFORMATION_SCHEMAtables when InnoDB is unavailable.", MySQLVersion:"5.6"},
    1013: {ErrorNumber:1013, ErrorType:"ServerError", Symbol:"ER_CANT_GET_STAT", SQLState:"HY000", Message:"Can't get status of '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1014: {ErrorNumber:1014, ErrorType:"ServerError", Symbol:"ER_CANT_GET_WD", SQLState:"HY000", Message:"Can't get working directory (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1015: {ErrorNumber:1015, ErrorType:"ServerError", Symbol:"ER_CANT_LOCK", SQLState:"HY000", Message:"Can't lock file (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1016: {ErrorNumber:1016, ErrorType:"ServerError", Symbol:"ER_CANT_OPEN_FILE", SQLState:"HY000", Message:"Can't open file: '%s' (errno: %d - %s)", Description:"InnoDB reports this error when the table fromthe InnoDB datafiles cannot be found, even though the.frm file for the table exists. SeeTroubleshooting InnoDB Data Dictionary Operations.", MySQLVersion:"5.6"},
    1017: {ErrorNumber:1017, ErrorType:"ServerError", Symbol:"ER_FILE_NOT_FOUND", SQLState:"HY000", Message:"Can't find file: '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1018: {ErrorNumber:1018, ErrorType:"ServerError", Symbol:"ER_CANT_READ_DIR", SQLState:"HY000", Message:"Can't read dir of '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1019: {ErrorNumber:1019, ErrorType:"ServerError", Symbol:"ER_CANT_SET_WD", SQLState:"HY000", Message:"Can't change dir to '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1020: {ErrorNumber:1020, ErrorType:"ServerError", Symbol:"ER_CHECKREAD", SQLState:"HY000", Message:"Record has changed since last read in table '%s'", Description:"", MySQLVersion:"5.6"},
    1021: {ErrorNumber:1021, ErrorType:"ServerError", Symbol:"ER_DISK_FULL", SQLState:"HY000", Message:"Disk full (%s)", Description:"waiting for someone to free some space...(errno: %d - %s)", MySQLVersion:"5.6"},
    1022: {ErrorNumber:1022, ErrorType:"ServerError", Symbol:"ER_DUP_KEY", SQLState:"23000", Message:"Can't write", Description:"duplicate key in table '%s'", MySQLVersion:"5.6"},
    1023: {ErrorNumber:1023, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_CLOSE", SQLState:"HY000", Message:"Error on close of '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1024: {ErrorNumber:1024, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_READ", SQLState:"HY000", Message:"Error reading file '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1025: {ErrorNumber:1025, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_RENAME", SQLState:"HY000", Message:"Error on rename of '%s' to '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1026: {ErrorNumber:1026, ErrorType:"ServerError", Symbol:"ER_ERROR_ON_WRITE", SQLState:"HY000", Message:"Error writing file '%s' (errno: %d - %s)", Description:"", MySQLVersion:"5.6"},
    1027: {ErrorNumber:1027, ErrorType:"ServerError", Symbol:"ER_FILE_USED", SQLState:"HY000", Message:"'%s' is locked against change", Description:"", MySQLVersion:"5.6"},
    1028: {ErrorNumber:1028, ErrorType:"ServerError", Symbol:"ER_FILSORT_ABORT", SQLState:"HY000", Message:"Sort aborted", Description:"", MySQLVersion:"5.6"},
    1029: {ErrorNumber:1029, ErrorType:"ServerError", Symbol:"ER_FORM_NOT_FOUND", SQLState:"HY000", Message:"View '%s' doesn't exist for '%s'", Description:"", MySQLVersion:"5.6"},
    1030: {ErrorNumber:1030, ErrorType:"ServerError", Symbol:"ER_GET_ERRNO", SQLState:"HY000", Message:"Got error %d from storage engine", Description:"Check the %d value to see what the OS errormeans. For example, 28 indicates that you have run out of diskspace.", MySQLVersion:"5.6"},
    1031: {ErrorNumber:1031, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_HA", SQLState:"HY000", Message:"Table storage engine for '%s' doesn't have this option", Description:"", MySQLVersion:"5.6"},
    1032: {ErrorNumber:1032, ErrorType:"ServerError", Symbol:"ER_KEY_NOT_FOUND", SQLState:"HY000", Message:"Can't find record in '%s'", Description:"", MySQLVersion:"5.6"},
    1033: {ErrorNumber:1033, ErrorType:"ServerError", Symbol:"ER_NOT_FORM_FILE", SQLState:"HY000", Message:"Incorrect information in file: '%s'", Description:"", MySQLVersion:"5.6"},
    1034: {ErrorNumber:1034, ErrorType:"ServerError", Symbol:"ER_NOT_KEYFILE", SQLState:"HY000", Message:"Incorrect key file for table '%s'", Description:"try to repair it", MySQLVersion:"5.6"},
    1035: {ErrorNumber:1035, ErrorType:"ServerError", Symbol:"ER_OLD_KEYFILE", SQLState:"HY000", Message:"Old key file for table '%s'", Description:"repair it!", MySQLVersion:"5.6"},
    1036: {ErrorNumber:1036, ErrorType:"ServerError", Symbol:"ER_OPEN_AS_READONLY", SQLState:"HY000", Message:"Table '%s' is read only", Description:"", MySQLVersion:"5.6"},
    1037: {ErrorNumber:1037, ErrorType:"ServerError", Symbol:"ER_OUTOFMEMORY", SQLState:"HY001", Message:"Out of memory", Description:"restart server and try again (needed %dbytes)", MySQLVersion:"5.6"},
    1038: {ErrorNumber:1038, ErrorType:"ServerError", Symbol:"ER_OUT_OF_SORTMEMORY", SQLState:"HY001", Message:"Out of sort memory, consider increasing server sortbuffer size", Description:"", MySQLVersion:"5.6"},
    1039: {ErrorNumber:1039, ErrorType:"ServerError", Symbol:"ER_UNEXPECTED_EOF", SQLState:"HY000", Message:"Unexpected EOF found when reading file '%s' (errno: %d -%s)", Description:"", MySQLVersion:"5.6"},
    1040: {ErrorNumber:1040, ErrorType:"ServerError", Symbol:"ER_CON_COUNT_ERROR", SQLState:"08004", Message:"Too many connections", Description:"", MySQLVersion:"5.6"},
    1041: {ErrorNumber:1041, ErrorType:"ServerError", Symbol:"ER_OUT_OF_RESOURCES", SQLState:"HY000", Message:"Out of memory", Description:"if not, you may have to use 'ulimit' toallow mysqld to use more memory or you can add more swap space", MySQLVersion:"5.6"},
    1042: {ErrorNumber:1042, ErrorType:"ServerError", Symbol:"ER_BAD_HOST_ERROR", SQLState:"08S01", Message:"Can't get hostname for your address", Description:"", MySQLVersion:"5.6"},
    1043: {ErrorNumber:1043, ErrorType:"ServerError", Symbol:"ER_HANDSHAKE_ERROR", SQLState:"08S01", Message:"Bad handshake", Description:"", MySQLVersion:"5.6"},
    1044: {ErrorNumber:1044, ErrorType:"ServerError", Symbol:"ER_DBACCESS_DENIED_ERROR", SQLState:"42000", Message:"Access denied for user '%s'@'%s' to database '%s'", Description:"", MySQLVersion:"5.6"},
    1045: {ErrorNumber:1045, ErrorType:"ServerError", Symbol:"ER_ACCESS_DENIED_ERROR", SQLState:"28000", Message:"Access denied for user '%s'@'%s' (using password: %s)", Description:"", MySQLVersion:"5.6"},
    1046: {ErrorNumber:1046, ErrorType:"ServerError", Symbol:"ER_NO_DB_ERROR", SQLState:"3D000", Message:"No database selected", Description:"", MySQLVersion:"5.6"},
    1047: {ErrorNumber:1047, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_COM_ERROR", SQLState:"08S01", Message:"Unknown command", Description:"", MySQLVersion:"5.6"},
    1048: {ErrorNumber:1048, ErrorType:"ServerError", Symbol:"ER_BAD_NULL_ERROR", SQLState:"23000", Message:"Column '%s' cannot be null", Description:"", MySQLVersion:"5.6"},
    1049: {ErrorNumber:1049, ErrorType:"ServerError", Symbol:"ER_BAD_DB_ERROR", SQLState:"42000", Message:"Unknown database '%s'", Description:"", MySQLVersion:"5.6"},
    1050: {ErrorNumber:1050, ErrorType:"ServerError", Symbol:"ER_TABLE_EXISTS_ERROR", SQLState:"42S01", Message:"Table '%s' already exists", Description:"", MySQLVersion:"5.6"},
    1051: {ErrorNumber:1051, ErrorType:"ServerError", Symbol:"ER_BAD_TABLE_ERROR", SQLState:"42S02", Message:"Unknown table '%s'", Description:"", MySQLVersion:"5.6"},
    1052: {ErrorNumber:1052, ErrorType:"ServerError", Symbol:"ER_NON_UNIQ_ERROR", SQLState:"23000", Message:"Column '%s' in %s is ambiguous", Description:"Resolution:", MySQLVersion:"5.6"},
    1053: {ErrorNumber:1053, ErrorType:"ServerError", Symbol:"ER_SERVER_SHUTDOWN", SQLState:"08S01", Message:"Server shutdown in progress", Description:"", MySQLVersion:"5.6"},
    1054: {ErrorNumber:1054, ErrorType:"ServerError", Symbol:"ER_BAD_FIELD_ERROR", SQLState:"42S22", Message:"Unknown column '%s' in '%s'", Description:"", MySQLVersion:"5.6"},
    1055: {ErrorNumber:1055, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_WITH_GROUP", SQLState:"42000", Message:"'%s' isn't in GROUP BY", Description:"", MySQLVersion:"5.6"},
    1056: {ErrorNumber:1056, ErrorType:"ServerError", Symbol:"ER_WRONG_GROUP_FIELD", SQLState:"42000", Message:"Can't group on '%s'", Description:"", MySQLVersion:"5.6"},
    1057: {ErrorNumber:1057, ErrorType:"ServerError", Symbol:"ER_WRONG_SUM_SELECT", SQLState:"42000", Message:"Statement has sum functions and columns in same statement", Description:"", MySQLVersion:"5.6"},
    1058: {ErrorNumber:1058, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_COUNT", SQLState:"21S01", Message:"Column count doesn't match value count", Description:"", MySQLVersion:"5.6"},
    1059: {ErrorNumber:1059, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_IDENT", SQLState:"42000", Message:"Identifier name '%s' is too long", Description:"", MySQLVersion:"5.6"},
    1060: {ErrorNumber:1060, ErrorType:"ServerError", Symbol:"ER_DUP_FIELDNAME", SQLState:"42S21", Message:"Duplicate column name '%s'", Description:"", MySQLVersion:"5.6"},
    1061: {ErrorNumber:1061, ErrorType:"ServerError", Symbol:"ER_DUP_KEYNAME", SQLState:"42000", Message:"Duplicate key name '%s'", Description:"", MySQLVersion:"5.6"},
    1062: {ErrorNumber:1062, ErrorType:"ServerError", Symbol:"ER_DUP_ENTRY", SQLState:"23000", Message:"Duplicate entry '%s' for key %d", Description:"The message returned with this error uses the format string forER_DUP_ENTRY_WITH_KEY_NAME.", MySQLVersion:"5.6"},
    1063: {ErrorNumber:1063, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_SPEC", SQLState:"42000", Message:"Incorrect column specifier for column '%s'", Description:"", MySQLVersion:"5.6"},
    1064: {ErrorNumber:1064, ErrorType:"ServerError", Symbol:"ER_PARSE_ERROR", SQLState:"42000", Message:"%s near '%s' at line %d", Description:"", MySQLVersion:"5.6"},
    1065: {ErrorNumber:1065, ErrorType:"ServerError", Symbol:"ER_EMPTY_QUERY", SQLState:"42000", Message:"Query was empty", Description:"", MySQLVersion:"5.6"},
    1066: {ErrorNumber:1066, ErrorType:"ServerError", Symbol:"ER_NONUNIQ_TABLE", SQLState:"42000", Message:"Not unique table/alias: '%s'", Description:"", MySQLVersion:"5.6"},
    1067: {ErrorNumber:1067, ErrorType:"ServerError", Symbol:"ER_INVALID_DEFAULT", SQLState:"42000", Message:"Invalid default value for '%s'", Description:"", MySQLVersion:"5.6"},
    1068: {ErrorNumber:1068, ErrorType:"ServerError", Symbol:"ER_MULTIPLE_PRI_KEY", SQLState:"42000", Message:"Multiple primary key defined", Description:"", MySQLVersion:"5.6"},
    1069: {ErrorNumber:1069, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_KEYS", SQLState:"42000", Message:"Too many keys specified", Description:"max %d keys allowed", MySQLVersion:"5.6"},
    1070: {ErrorNumber:1070, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_KEY_PARTS", SQLState:"42000", Message:"Too many key parts specified", Description:"max %d parts allowed", MySQLVersion:"5.6"},
    1071: {ErrorNumber:1071, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_KEY", SQLState:"42000", Message:"Specified key was too long", Description:"max key length is %d bytes", MySQLVersion:"5.6"},
    1072: {ErrorNumber:1072, ErrorType:"ServerError", Symbol:"ER_KEY_COLUMN_DOES_NOT_EXITS", SQLState:"42000", Message:"Key column '%s' doesn't exist in table", Description:"", MySQLVersion:"5.6"},
    1073: {ErrorNumber:1073, ErrorType:"ServerError", Symbol:"ER_BLOB_USED_AS_KEY", SQLState:"42000", Message:"BLOB column '%s' can't be used in key specification withthe used table type", Description:"", MySQLVersion:"5.6"},
    1074: {ErrorNumber:1074, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_FIELDLENGTH", SQLState:"42000", Message:"Column length too big for column '%s' (max = %lu)", Description:"useBLOB or TEXT instead", MySQLVersion:"5.6"},
    1075: {ErrorNumber:1075, ErrorType:"ServerError", Symbol:"ER_WRONG_AUTO_KEY", SQLState:"42000", Message:"Incorrect table definition", Description:"there can be only one autocolumn and it must be defined as a key", MySQLVersion:"5.6"},
    1076: {ErrorNumber:1076, ErrorType:"ServerError", Symbol:"ER_READY", SQLState:"HY000", Message:"%s: ready for connections. Version: '%s' socket: '%s'port: %d", Description:"", MySQLVersion:"5.6"},
    1077: {ErrorNumber:1077, ErrorType:"ServerError", Symbol:"ER_NORMAL_SHUTDOWN", SQLState:"HY000", Message:"%s: Normal shutdown", Description:"", MySQLVersion:"5.6"},
    1078: {ErrorNumber:1078, ErrorType:"ServerError", Symbol:"ER_GOT_SIGNAL", SQLState:"HY000", Message:"%s: Got signal %d. Aborting!", Description:"", MySQLVersion:"5.6"},
    1079: {ErrorNumber:1079, ErrorType:"ServerError", Symbol:"ER_SHUTDOWN_COMPLETE", SQLState:"HY000", Message:"%s: Shutdown complete", Description:"", MySQLVersion:"5.6"},
    1080: {ErrorNumber:1080, ErrorType:"ServerError", Symbol:"ER_FORCING_CLOSE", SQLState:"08S01", Message:"%s: Forcing close of thread %ld user: '%s'", Description:"", MySQLVersion:"5.6"},
    1081: {ErrorNumber:1081, ErrorType:"ServerError", Symbol:"ER_IPSOCK_ERROR", SQLState:"08S01", Message:"Can't create IP socket", Description:"", MySQLVersion:"5.6"},
    1082: {ErrorNumber:1082, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_INDEX", SQLState:"42S12", Message:"Table '%s' has no index like the one used in CREATEINDEX", Description:"recreate the table", MySQLVersion:"5.6"},
    1083: {ErrorNumber:1083, ErrorType:"ServerError", Symbol:"ER_WRONG_FIELD_TERMINATORS", SQLState:"42000", Message:"Field separator argument is not what is expected", Description:"checkthe manual", MySQLVersion:"5.6"},
    1084: {ErrorNumber:1084, ErrorType:"ServerError", Symbol:"ER_BLOBS_AND_NO_TERMINATED", SQLState:"42000", Message:"You can't use fixed rowlength with BLOBs", Description:"please use'fields terminated by'", MySQLVersion:"5.6"},
    1085: {ErrorNumber:1085, ErrorType:"ServerError", Symbol:"ER_TEXTFILE_NOT_READABLE", SQLState:"HY000", Message:"The file '%s' must be in the database directory or bereadable by all", Description:"", MySQLVersion:"5.6"},
    1086: {ErrorNumber:1086, ErrorType:"ServerError", Symbol:"ER_FILE_EXISTS_ERROR", SQLState:"HY000", Message:"File '%s' already exists", Description:"", MySQLVersion:"5.6"},
    1087: {ErrorNumber:1087, ErrorType:"ServerError", Symbol:"ER_LOAD_INFO", SQLState:"HY000", Message:"Records: %ld Deleted: %ld Skipped: %ld Warnings: %ld", Description:"", MySQLVersion:"5.6"},
    1088: {ErrorNumber:1088, ErrorType:"ServerError", Symbol:"ER_ALTER_INFO", SQLState:"HY000", Message:"Records: %ld Duplicates: %ld", Description:"", MySQLVersion:"5.6"},
    1089: {ErrorNumber:1089, ErrorType:"ServerError", Symbol:"ER_WRONG_SUB_KEY", SQLState:"HY000", Message:"Incorrect prefix key", Description:"the used key part isn't a string,the used length is longer than the key part, or the storage enginedoesn't support unique prefix keys", MySQLVersion:"5.6"},
    1090: {ErrorNumber:1090, ErrorType:"ServerError", Symbol:"ER_CANT_REMOVE_ALL_FIELDS", SQLState:"42000", Message:"You can't delete all columns with ALTER TABLE", Description:"use DROPTABLE instead", MySQLVersion:"5.6"},
    1091: {ErrorNumber:1091, ErrorType:"ServerError", Symbol:"ER_CANT_DROP_FIELD_OR_KEY", SQLState:"42000", Message:"Can't DROP '%s'", Description:"check that column/key exists", MySQLVersion:"5.6"},
    1092: {ErrorNumber:1092, ErrorType:"ServerError", Symbol:"ER_INSERT_INFO", SQLState:"HY000", Message:"Records: %ld Duplicates: %ld Warnings: %ld", Description:"", MySQLVersion:"5.6"},
    1093: {ErrorNumber:1093, ErrorType:"ServerError", Symbol:"ER_UPDATE_TABLE_USED", SQLState:"HY000", Message:"You can't specify target table '%s' for update in FROMclause", Description:"This error occurs for attempts to select from and modify the sametable within a single statement. SeeRestrictions on Subqueries.", MySQLVersion:"5.6"},
    1094: {ErrorNumber:1094, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_THREAD", SQLState:"HY000", Message:"Unknown thread id: %lu", Description:"", MySQLVersion:"5.6"},
    1095: {ErrorNumber:1095, ErrorType:"ServerError", Symbol:"ER_KILL_DENIED_ERROR", SQLState:"HY000", Message:"You are not owner of thread %lu", Description:"", MySQLVersion:"5.6"},
    1096: {ErrorNumber:1096, ErrorType:"ServerError", Symbol:"ER_NO_TABLES_USED", SQLState:"HY000", Message:"No tables used", Description:"", MySQLVersion:"5.6"},
    1097: {ErrorNumber:1097, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_SET", SQLState:"HY000", Message:"Too many strings for column %s and SET", Description:"", MySQLVersion:"5.6"},
    1098: {ErrorNumber:1098, ErrorType:"ServerError", Symbol:"ER_NO_UNIQUE_LOGFILE", SQLState:"HY000", Message:"Can't generate a unique log-filename %s.(1-999)", Description:"", MySQLVersion:"5.6"},
    1099: {ErrorNumber:1099, ErrorType:"ServerError", Symbol:"ER_TABLE_NOT_LOCKED_FOR_WRITE", SQLState:"HY000", Message:"Table '%s' was locked with a READ lock and can't beupdated", Description:"", MySQLVersion:"5.6"},
    1100: {ErrorNumber:1100, ErrorType:"ServerError", Symbol:"ER_TABLE_NOT_LOCKED", SQLState:"HY000", Message:"Table '%s' was not locked with LOCK TABLES", Description:"", MySQLVersion:"5.6"},
    1101: {ErrorNumber:1101, ErrorType:"ServerError", Symbol:"ER_BLOB_CANT_HAVE_DEFAULT", SQLState:"42000", Message:"BLOB/TEXT column '%s' can't have a default value", Description:"", MySQLVersion:"5.6"},
    1102: {ErrorNumber:1102, ErrorType:"ServerError", Symbol:"ER_WRONG_DB_NAME", SQLState:"42000", Message:"Incorrect database name '%s'", Description:"", MySQLVersion:"5.6"},
    1103: {ErrorNumber:1103, ErrorType:"ServerError", Symbol:"ER_WRONG_TABLE_NAME", SQLState:"42000", Message:"Incorrect table name '%s'", Description:"", MySQLVersion:"5.6"},
    1104: {ErrorNumber:1104, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_SELECT", SQLState:"42000", Message:"The SELECT would examine more than MAX_JOIN_SIZE rows", Description:"check your WHERE and use SET SQL_BIG_SELECTS=1 or SETMAX_JOIN_SIZE=# if the SELECT is okay", MySQLVersion:"5.6"},
    1105: {ErrorNumber:1105, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_ERROR", SQLState:"HY000", Message:"Unknown error", Description:"", MySQLVersion:"5.6"},
    1106: {ErrorNumber:1106, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_PROCEDURE", SQLState:"42000", Message:"Unknown procedure '%s'", Description:"", MySQLVersion:"5.6"},
    1107: {ErrorNumber:1107, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMCOUNT_TO_PROCEDURE", SQLState:"42000", Message:"Incorrect parameter count to procedure '%s'", Description:"", MySQLVersion:"5.6"},
    1108: {ErrorNumber:1108, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMETERS_TO_PROCEDURE", SQLState:"HY000", Message:"Incorrect parameters to procedure '%s'", Description:"", MySQLVersion:"5.6"},
    1109: {ErrorNumber:1109, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_TABLE", SQLState:"42S02", Message:"Unknown table '%s' in %s", Description:"", MySQLVersion:"5.6"},
    1110: {ErrorNumber:1110, ErrorType:"ServerError", Symbol:"ER_FIELD_SPECIFIED_TWICE", SQLState:"42000", Message:"Column '%s' specified twice", Description:"", MySQLVersion:"5.6"},
    1111: {ErrorNumber:1111, ErrorType:"ServerError", Symbol:"ER_INVALID_GROUP_FUNC_USE", SQLState:"HY000", Message:"Invalid use of group function", Description:"", MySQLVersion:"5.6"},
    1112: {ErrorNumber:1112, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_EXTENSION", SQLState:"42000", Message:"Table '%s' uses an extension that doesn't exist in thisMySQL version", Description:"", MySQLVersion:"5.6"},
    1113: {ErrorNumber:1113, ErrorType:"ServerError", Symbol:"ER_TABLE_MUST_HAVE_COLUMNS", SQLState:"42000", Message:"A table must have at least 1 column", Description:"", MySQLVersion:"5.6"},
    1114: {ErrorNumber:1114, ErrorType:"ServerError", Symbol:"ER_RECORD_FILE_FULL", SQLState:"HY000", Message:"The table '%s' is full", Description:"InnoDB reports this error when the systemtablespace runs out of free space. Reconfigure the systemtablespace to add a new data file.", MySQLVersion:"5.6"},
    1115: {ErrorNumber:1115, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_CHARACTER_SET", SQLState:"42000", Message:"Unknown character set: '%s'", Description:"", MySQLVersion:"5.6"},
    1116: {ErrorNumber:1116, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_TABLES", SQLState:"HY000", Message:"Too many tables", Description:"MySQL can only use %d tables in a join", MySQLVersion:"5.6"},
    1117: {ErrorNumber:1117, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_FIELDS", SQLState:"HY000", Message:"Too many columns", Description:"", MySQLVersion:"5.6"},
    1118: {ErrorNumber:1118, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_ROWSIZE", SQLState:"42000", Message:"Row size too large. The maximum row size for the usedtable type, not counting BLOBs, is %ld. This includes storageoverhead, check the manual. You have to change some columns toTEXT or BLOBs", Description:"", MySQLVersion:"5.6"},
    1119: {ErrorNumber:1119, ErrorType:"ServerError", Symbol:"ER_STACK_OVERRUN", SQLState:"HY000", Message:"Thread stack overrun: Used: %ld of a %ld stack. Use'mysqld --thread_stack=#' to specify a bigger stack if needed", Description:"", MySQLVersion:"5.6"},
    1120: {ErrorNumber:1120, ErrorType:"ServerError", Symbol:"ER_WRONG_OUTER_JOIN", SQLState:"42000", Message:"Cross dependency found in OUTER JOIN", Description:"examine your ONconditions", MySQLVersion:"5.6"},
    1121: {ErrorNumber:1121, ErrorType:"ServerError", Symbol:"ER_NULL_COLUMN_IN_INDEX", SQLState:"42000", Message:"Table handler doesn't support NULL in given index. Pleasechange column '%s' to be NOT NULL or use another handler", Description:"", MySQLVersion:"5.6"},
    1122: {ErrorNumber:1122, ErrorType:"ServerError", Symbol:"ER_CANT_FIND_UDF", SQLState:"HY000", Message:"Can't load function '%s'", Description:"", MySQLVersion:"5.6"},
    1123: {ErrorNumber:1123, ErrorType:"ServerError", Symbol:"ER_CANT_INITIALIZE_UDF", SQLState:"HY000", Message:"Can't initialize function '%s'", Description:"%s", MySQLVersion:"5.6"},
    1124: {ErrorNumber:1124, ErrorType:"ServerError", Symbol:"ER_UDF_NO_PATHS", SQLState:"HY000", Message:"No paths allowed for shared library", Description:"", MySQLVersion:"5.6"},
    1125: {ErrorNumber:1125, ErrorType:"ServerError", Symbol:"ER_UDF_EXISTS", SQLState:"HY000", Message:"Function '%s' already exists", Description:"", MySQLVersion:"5.6"},
    1126: {ErrorNumber:1126, ErrorType:"ServerError", Symbol:"ER_CANT_OPEN_LIBRARY", SQLState:"HY000", Message:"Can't open shared library '%s' (errno: %d %s)", Description:"", MySQLVersion:"5.6"},
    1127: {ErrorNumber:1127, ErrorType:"ServerError", Symbol:"ER_CANT_FIND_DL_ENTRY", SQLState:"HY000", Message:"Can't find symbol '%s' in library", Description:"", MySQLVersion:"5.6"},
    1128: {ErrorNumber:1128, ErrorType:"ServerError", Symbol:"ER_FUNCTION_NOT_DEFINED", SQLState:"HY000", Message:"Function '%s' is not defined", Description:"", MySQLVersion:"5.6"},
    1129: {ErrorNumber:1129, ErrorType:"ServerError", Symbol:"ER_HOST_IS_BLOCKED", SQLState:"HY000", Message:"Host '%s' is blocked because of many connection errors", Description:"unblock with 'mysqladmin flush-hosts'", MySQLVersion:"5.6"},
    1130: {ErrorNumber:1130, ErrorType:"ServerError", Symbol:"ER_HOST_NOT_PRIVILEGED", SQLState:"HY000", Message:"Host '%s' is not allowed to connect to this MySQL server", Description:"", MySQLVersion:"5.6"},
    1131: {ErrorNumber:1131, ErrorType:"ServerError", Symbol:"ER_PASSWORD_ANONYMOUS_USER", SQLState:"42000", Message:"You are using MySQL as an anonymous user and anonymoususers are not allowed to change passwords", Description:"", MySQLVersion:"5.6"},
    1132: {ErrorNumber:1132, ErrorType:"ServerError", Symbol:"ER_PASSWORD_NOT_ALLOWED", SQLState:"42000", Message:"You must have privileges to update tables in the mysqldatabase to be able to change passwords for others", Description:"", MySQLVersion:"5.6"},
    1133: {ErrorNumber:1133, ErrorType:"ServerError", Symbol:"ER_PASSWORD_NO_MATCH", SQLState:"42000", Message:"Can't find any matching row in the user table", Description:"", MySQLVersion:"5.6"},
    1134: {ErrorNumber:1134, ErrorType:"ServerError", Symbol:"ER_UPDATE_INFO", SQLState:"HY000", Message:"Rows matched: %ld Changed: %ld Warnings: %ld", Description:"", MySQLVersion:"5.6"},
    1135: {ErrorNumber:1135, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_THREAD", SQLState:"HY000", Message:"Can't create a new thread (errno %d)", Description:"if you are not outof available memory, you can consult the manual for a possibleOS-dependent bug", MySQLVersion:"5.6"},
    1136: {ErrorNumber:1136, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_COUNT_ON_ROW", SQLState:"21S01", Message:"Column count doesn't match value count at row %ld", Description:"", MySQLVersion:"5.6"},
    1137: {ErrorNumber:1137, ErrorType:"ServerError", Symbol:"ER_CANT_REOPEN_TABLE", SQLState:"HY000", Message:"Can't reopen table: '%s'", Description:"", MySQLVersion:"5.6"},
    1138: {ErrorNumber:1138, ErrorType:"ServerError", Symbol:"ER_INVALID_USE_OF_NULL", SQLState:"22004", Message:"Invalid use of NULL value", Description:"", MySQLVersion:"5.6"},
    1139: {ErrorNumber:1139, ErrorType:"ServerError", Symbol:"ER_REGEXP_ERROR", SQLState:"42000", Message:"Got error '%s' from regexp", Description:"", MySQLVersion:"5.6"},
    1140: {ErrorNumber:1140, ErrorType:"ServerError", Symbol:"ER_MIX_OF_GROUP_FUNC_AND_FIELDS", SQLState:"42000", Message:"Mixing of GROUP columns (MIN(),MAX(),COUNT(),...) with noGROUP columns is illegal if there is no GROUP BY clause", Description:"", MySQLVersion:"5.6"},
    1141: {ErrorNumber:1141, ErrorType:"ServerError", Symbol:"ER_NONEXISTING_GRANT", SQLState:"42000", Message:"There is no such grant defined for user '%s' on host '%s'", Description:"", MySQLVersion:"5.6"},
    1142: {ErrorNumber:1142, ErrorType:"ServerError", Symbol:"ER_TABLEACCESS_DENIED_ERROR", SQLState:"42000", Message:"%s command denied to user '%s'@'%s' for table '%s'", Description:"", MySQLVersion:"5.6"},
    1143: {ErrorNumber:1143, ErrorType:"ServerError", Symbol:"ER_COLUMNACCESS_DENIED_ERROR", SQLState:"42000", Message:"%s command denied to user '%s'@'%s' for column '%s' intable '%s'", Description:"", MySQLVersion:"5.6"},
    1144: {ErrorNumber:1144, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_GRANT_FOR_TABLE", SQLState:"42000", Message:"Illegal GRANT/REVOKE command", Description:"please consult the manualto see which privileges can be used", MySQLVersion:"5.6"},
    1145: {ErrorNumber:1145, ErrorType:"ServerError", Symbol:"ER_GRANT_WRONG_HOST_OR_USER", SQLState:"42000", Message:"The host or user argument to GRANT is too long", Description:"", MySQLVersion:"5.6"},
    1146: {ErrorNumber:1146, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_TABLE", SQLState:"42S02", Message:"Table '%s.%s' doesn't exist", Description:"", MySQLVersion:"5.6"},
    1147: {ErrorNumber:1147, ErrorType:"ServerError", Symbol:"ER_NONEXISTING_TABLE_GRANT", SQLState:"42000", Message:"There is no such grant defined for user '%s' on host '%s'on table '%s'", Description:"", MySQLVersion:"5.6"},
    1148: {ErrorNumber:1148, ErrorType:"ServerError", Symbol:"ER_NOT_ALLOWED_COMMAND", SQLState:"42000", Message:"The used command is not allowed with this MySQL version", Description:"", MySQLVersion:"5.6"},
    1149: {ErrorNumber:1149, ErrorType:"ServerError", Symbol:"ER_SYNTAX_ERROR", SQLState:"42000", Message:"You have an error in your SQL syntax", Description:"check the manualthat corresponds to your MySQL server version for the right syntaxto use", MySQLVersion:"5.6"},
    1150: {ErrorNumber:1150, ErrorType:"ServerError", Symbol:"ER_DELAYED_CANT_CHANGE_LOCK", SQLState:"HY000", Message:"Delayed insert thread couldn't get requested lock fortable %s", Description:"", MySQLVersion:"5.6"},
    1151: {ErrorNumber:1151, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_DELAYED_THREADS", SQLState:"HY000", Message:"Too many delayed threads in use", Description:"", MySQLVersion:"5.6"},
    1152: {ErrorNumber:1152, ErrorType:"ServerError", Symbol:"ER_ABORTING_CONNECTION", SQLState:"08S01", Message:"Aborted connection %ld to db: '%s' user: '%s' (%s)", Description:"", MySQLVersion:"5.6"},
    1153: {ErrorNumber:1153, ErrorType:"ServerError", Symbol:"ER_NET_PACKET_TOO_LARGE", SQLState:"08S01", Message:"Got a packet bigger than 'max_allowed_packet' bytes", Description:"", MySQLVersion:"5.6"},
    1154: {ErrorNumber:1154, ErrorType:"ServerError", Symbol:"ER_NET_READ_ERROR_FROM_PIPE", SQLState:"08S01", Message:"Got a read error from the connection pipe", Description:"", MySQLVersion:"5.6"},
    1155: {ErrorNumber:1155, ErrorType:"ServerError", Symbol:"ER_NET_FCNTL_ERROR", SQLState:"08S01", Message:"Got an error from fcntl()", Description:"", MySQLVersion:"5.6"},
    1156: {ErrorNumber:1156, ErrorType:"ServerError", Symbol:"ER_NET_PACKETS_OUT_OF_ORDER", SQLState:"08S01", Message:"Got packets out of order", Description:"", MySQLVersion:"5.6"},
    1157: {ErrorNumber:1157, ErrorType:"ServerError", Symbol:"ER_NET_UNCOMPRESS_ERROR", SQLState:"08S01", Message:"Couldn't uncompress communication packet", Description:"", MySQLVersion:"5.6"},
    1158: {ErrorNumber:1158, ErrorType:"ServerError", Symbol:"ER_NET_READ_ERROR", SQLState:"08S01", Message:"Got an error reading communication packets", Description:"", MySQLVersion:"5.6"},
    1159: {ErrorNumber:1159, ErrorType:"ServerError", Symbol:"ER_NET_READ_INTERRUPTED", SQLState:"08S01", Message:"Got timeout reading communication packets", Description:"", MySQLVersion:"5.6"},
    1160: {ErrorNumber:1160, ErrorType:"ServerError", Symbol:"ER_NET_ERROR_ON_WRITE", SQLState:"08S01", Message:"Got an error writing communication packets", Description:"", MySQLVersion:"5.6"},
    1161: {ErrorNumber:1161, ErrorType:"ServerError", Symbol:"ER_NET_WRITE_INTERRUPTED", SQLState:"08S01", Message:"Got timeout writing communication packets", Description:"", MySQLVersion:"5.6"},
    1162: {ErrorNumber:1162, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_STRING", SQLState:"42000", Message:"Result string is longer than 'max_allowed_packet' bytes", Description:"", MySQLVersion:"5.6"},
    1163: {ErrorNumber:1163, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_BLOB", SQLState:"42000", Message:"The used table type doesn't support BLOB/TEXT columns", Description:"", MySQLVersion:"5.6"},
    1164: {ErrorNumber:1164, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_AUTO_INCREMENT", SQLState:"42000", Message:"The used table type doesn't support AUTO_INCREMENTcolumns", Description:"", MySQLVersion:"5.6"},
    1165: {ErrorNumber:1165, ErrorType:"ServerError", Symbol:"ER_DELAYED_INSERT_TABLE_LOCKED", SQLState:"HY000", Message:"INSERT DELAYED can't be used with table '%s' because itis locked with LOCK TABLES", Description:"", MySQLVersion:"5.6"},
    1166: {ErrorNumber:1166, ErrorType:"ServerError", Symbol:"ER_WRONG_COLUMN_NAME", SQLState:"42000", Message:"Incorrect column name '%s'", Description:"", MySQLVersion:"5.6"},
    1167: {ErrorNumber:1167, ErrorType:"ServerError", Symbol:"ER_WRONG_KEY_COLUMN", SQLState:"42000", Message:"The used storage engine can't index column '%s'", Description:"", MySQLVersion:"5.6"},
    1168: {ErrorNumber:1168, ErrorType:"ServerError", Symbol:"ER_WRONG_MRG_TABLE", SQLState:"HY000", Message:"Unable to open underlying table which is differentlydefined or of non-MyISAM type or doesn't exist", Description:"", MySQLVersion:"5.6"},
    1169: {ErrorNumber:1169, ErrorType:"ServerError", Symbol:"ER_DUP_UNIQUE", SQLState:"23000", Message:"Can't write, because of unique constraint, to table '%s'", Description:"", MySQLVersion:"5.6"},
    1170: {ErrorNumber:1170, ErrorType:"ServerError", Symbol:"ER_BLOB_KEY_WITHOUT_LENGTH", SQLState:"42000", Message:"BLOB/TEXT column '%s' used in key specification without akey length", Description:"", MySQLVersion:"5.6"},
    1171: {ErrorNumber:1171, ErrorType:"ServerError", Symbol:"ER_PRIMARY_CANT_HAVE_NULL", SQLState:"42000", Message:"All parts of a PRIMARY KEY must be NOT NULL", Description:"if you needNULL in a key, use UNIQUE instead", MySQLVersion:"5.6"},
    1172: {ErrorNumber:1172, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_ROWS", SQLState:"42000", Message:"Result consisted of more than one row", Description:"", MySQLVersion:"5.6"},
    1173: {ErrorNumber:1173, ErrorType:"ServerError", Symbol:"ER_REQUIRES_PRIMARY_KEY", SQLState:"42000", Message:"This table type requires a primary key", Description:"", MySQLVersion:"5.6"},
    1174: {ErrorNumber:1174, ErrorType:"ServerError", Symbol:"ER_NO_RAID_COMPILED", SQLState:"HY000", Message:"This version of MySQL is not compiled with RAID support", Description:"", MySQLVersion:"5.6"},
    1175: {ErrorNumber:1175, ErrorType:"ServerError", Symbol:"ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE", SQLState:"HY000", Message:"You are using safe update mode and you tried to update atable without a WHERE that uses a KEY column", Description:"", MySQLVersion:"5.6"},
    1176: {ErrorNumber:1176, ErrorType:"ServerError", Symbol:"ER_KEY_DOES_NOT_EXITS", SQLState:"42000", Message:"Key '%s' doesn't exist in table '%s'", Description:"", MySQLVersion:"5.6"},
    1177: {ErrorNumber:1177, ErrorType:"ServerError", Symbol:"ER_CHECK_NO_SUCH_TABLE", SQLState:"42000", Message:"Can't open table", Description:"", MySQLVersion:"5.6"},
    1178: {ErrorNumber:1178, ErrorType:"ServerError", Symbol:"ER_CHECK_NOT_IMPLEMENTED", SQLState:"42000", Message:"The storage engine for the table doesn't support %s", Description:"", MySQLVersion:"5.6"},
    1179: {ErrorNumber:1179, ErrorType:"ServerError", Symbol:"ER_CANT_DO_THIS_DURING_AN_TRANSACTION", SQLState:"25000", Message:"You are not allowed to execute this command in atransaction", Description:"", MySQLVersion:"5.6"},
    1180: {ErrorNumber:1180, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_COMMIT", SQLState:"HY000", Message:"Got error %d during COMMIT", Description:"", MySQLVersion:"5.6"},
    1181: {ErrorNumber:1181, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_ROLLBACK", SQLState:"HY000", Message:"Got error %d during ROLLBACK", Description:"", MySQLVersion:"5.6"},
    1182: {ErrorNumber:1182, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_FLUSH_LOGS", SQLState:"HY000", Message:"Got error %d during FLUSH_LOGS", Description:"", MySQLVersion:"5.6"},
    1183: {ErrorNumber:1183, ErrorType:"ServerError", Symbol:"ER_ERROR_DURING_CHECKPOINT", SQLState:"HY000", Message:"Got error %d during CHECKPOINT", Description:"", MySQLVersion:"5.6"},
    1184: {ErrorNumber:1184, ErrorType:"ServerError", Symbol:"ER_NEW_ABORTING_CONNECTION", SQLState:"08S01", Message:"Aborted connection %ld to db: '%s' user: '%s' host: '%s'(%s)", Description:"", MySQLVersion:"5.6"},
    1185: {ErrorNumber:1185, ErrorType:"ServerError", Symbol:"ER_DUMP_NOT_IMPLEMENTED", SQLState:"HY000", Message:"The storage engine for the table does not support binarytable dump", Description:"", MySQLVersion:"5.6"},
    1186: {ErrorNumber:1186, ErrorType:"ServerError", Symbol:"ER_FLUSH_MASTER_BINLOG_CLOSED", SQLState:"HY000", Message:"Binlog closed, cannot RESET MASTER", Description:"", MySQLVersion:"5.6"},
    1187: {ErrorNumber:1187, ErrorType:"ServerError", Symbol:"ER_INDEX_REBUILD", SQLState:"HY000", Message:"Failed rebuilding the index of dumped table '%s'", Description:"", MySQLVersion:"5.6"},
    1188: {ErrorNumber:1188, ErrorType:"ServerError", Symbol:"ER_MASTER", SQLState:"HY000", Message:"Error from master: '%s'", Description:"", MySQLVersion:"5.6"},
    1189: {ErrorNumber:1189, ErrorType:"ServerError", Symbol:"ER_MASTER_NET_READ", SQLState:"08S01", Message:"Net error reading from master", Description:"", MySQLVersion:"5.6"},
    1190: {ErrorNumber:1190, ErrorType:"ServerError", Symbol:"ER_MASTER_NET_WRITE", SQLState:"08S01", Message:"Net error writing to master", Description:"", MySQLVersion:"5.6"},
    1191: {ErrorNumber:1191, ErrorType:"ServerError", Symbol:"ER_FT_MATCHING_KEY_NOT_FOUND", SQLState:"HY000", Message:"Can't find FULLTEXT index matching the column list", Description:"", MySQLVersion:"5.6"},
    1192: {ErrorNumber:1192, ErrorType:"ServerError", Symbol:"ER_LOCK_OR_ACTIVE_TRANSACTION", SQLState:"HY000", Message:"Can't execute the given command because you have activelocked tables or an active transaction", Description:"", MySQLVersion:"5.6"},
    1193: {ErrorNumber:1193, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_SYSTEM_VARIABLE", SQLState:"HY000", Message:"Unknown system variable '%s'", Description:"", MySQLVersion:"5.6"},
    1194: {ErrorNumber:1194, ErrorType:"ServerError", Symbol:"ER_CRASHED_ON_USAGE", SQLState:"HY000", Message:"Table '%s' is marked as crashed and should be repaired", Description:"", MySQLVersion:"5.6"},
    1195: {ErrorNumber:1195, ErrorType:"ServerError", Symbol:"ER_CRASHED_ON_REPAIR", SQLState:"HY000", Message:"Table '%s' is marked as crashed and last (automatic?)repair failed", Description:"", MySQLVersion:"5.6"},
    1196: {ErrorNumber:1196, ErrorType:"ServerError", Symbol:"ER_WARNING_NOT_COMPLETE_ROLLBACK", SQLState:"HY000", Message:"Some non-transactional changed tables couldn't be rolledback", Description:"", MySQLVersion:"5.6"},
    1197: {ErrorNumber:1197, ErrorType:"ServerError", Symbol:"ER_TRANS_CACHE_FULL", SQLState:"HY000", Message:"Multi-statement transaction required more than'max_binlog_cache_size' bytes of storage", Description:"increase this mysqldvariable and try again", MySQLVersion:"5.6"},
    1198: {ErrorNumber:1198, ErrorType:"ServerError", Symbol:"ER_SLAVE_MUST_STOP", SQLState:"HY000", Message:"This operation cannot be performed with a running slave", Description:"run STOP SLAVE first", MySQLVersion:"5.6"},
    1199: {ErrorNumber:1199, ErrorType:"ServerError", Symbol:"ER_SLAVE_NOT_RUNNING", SQLState:"HY000", Message:"This operation requires a running slave", Description:"configure slaveand do START SLAVE", MySQLVersion:"5.6"},
    1200: {ErrorNumber:1200, ErrorType:"ServerError", Symbol:"ER_BAD_SLAVE", SQLState:"HY000", Message:"The server is not configured as slave", Description:"fix in config fileor with CHANGE MASTER TO", MySQLVersion:"5.6"},
    1201: {ErrorNumber:1201, ErrorType:"ServerError", Symbol:"ER_MASTER_INFO", SQLState:"HY000", Message:"Could not initialize master info structure", Description:"more errormessages can be found in the MySQL error log", MySQLVersion:"5.6"},
    1202: {ErrorNumber:1202, ErrorType:"ServerError", Symbol:"ER_SLAVE_THREAD", SQLState:"HY000", Message:"Could not create slave thread", Description:"check system resources", MySQLVersion:"5.6"},
    1203: {ErrorNumber:1203, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_USER_CONNECTIONS", SQLState:"42000", Message:"User %s already has more than 'max_user_connections'active connections", Description:"", MySQLVersion:"5.6"},
    1204: {ErrorNumber:1204, ErrorType:"ServerError", Symbol:"ER_SET_CONSTANTS_ONLY", SQLState:"HY000", Message:"You may only use constant expressions with SET", Description:"", MySQLVersion:"5.6"},
    1205: {ErrorNumber:1205, ErrorType:"ServerError", Symbol:"ER_LOCK_WAIT_TIMEOUT", SQLState:"HY000", Message:"Lock wait timeout exceeded", Description:"InnoDB reports this error when lock waittimeout expires. The statement that waited too long wasrolled back (not the entiretransaction). You canincrease the value of theinnodb_lock_wait_timeoutconfiguration option if SQL statements should wait longer forother transactions to complete, or decrease it if too manylong-running transactions are causinglocking problems and reducingconcurrency on a busysystem.", MySQLVersion:"5.6"},
    1206: {ErrorNumber:1206, ErrorType:"ServerError", Symbol:"ER_LOCK_TABLE_FULL", SQLState:"HY000", Message:"The total number of locks exceeds the lock table size", Description:"InnoDB reports this error when the total numberof locks exceeds the amount of memory devoted to managing locks.To avoid this error, increase the value ofinnodb_buffer_pool_size. Withinan individual application, a workaround may be to break a largeoperation into smaller pieces. For example, if the error occursfor a large INSERT, perform severalsmaller INSERT operations.", MySQLVersion:"5.6"},
    1207: {ErrorNumber:1207, ErrorType:"ServerError", Symbol:"ER_READ_ONLY_TRANSACTION", SQLState:"25000", Message:"Update locks cannot be acquired during a READ UNCOMMITTEDtransaction", Description:"", MySQLVersion:"5.6"},
    1208: {ErrorNumber:1208, ErrorType:"ServerError", Symbol:"ER_DROP_DB_WITH_READ_LOCK", SQLState:"HY000", Message:"DROP DATABASE not allowed while thread is holding globalread lock", Description:"", MySQLVersion:"5.6"},
    1209: {ErrorNumber:1209, ErrorType:"ServerError", Symbol:"ER_CREATE_DB_WITH_READ_LOCK", SQLState:"HY000", Message:"CREATE DATABASE not allowed while thread is holdingglobal read lock", Description:"", MySQLVersion:"5.6"},
    1210: {ErrorNumber:1210, ErrorType:"ServerError", Symbol:"ER_WRONG_ARGUMENTS", SQLState:"HY000", Message:"Incorrect arguments to %s", Description:"", MySQLVersion:"5.6"},
    1211: {ErrorNumber:1211, ErrorType:"ServerError", Symbol:"ER_NO_PERMISSION_TO_CREATE_USER", SQLState:"42000", Message:"'%s'@'%s' is not allowed to create new users", Description:"", MySQLVersion:"5.6"},
    1212: {ErrorNumber:1212, ErrorType:"ServerError", Symbol:"ER_UNION_TABLES_IN_DIFFERENT_DIR", SQLState:"HY000", Message:"Incorrect table definition", Description:"all MERGE tables must be inthe same database", MySQLVersion:"5.6"},
    1213: {ErrorNumber:1213, ErrorType:"ServerError", Symbol:"ER_LOCK_DEADLOCK", SQLState:"40001", Message:"Deadlock found when trying to get lock", Description:"InnoDB reports this error when atransaction encounters adeadlock and is automaticallyrolled back so that yourapplication can take corrective action. To recover from thiserror, run all the operations in this transaction again. Adeadlock occurs when requests for locks arrive in inconsistentorder between transactions. The transaction that was rolled backreleased all its locks, and the other transaction can now get allthe locks it requested. Thus, when you re-run the transaction thatwas rolled back, it might have to wait for other transactions tocomplete, but typically the deadlock does not recur. If youencounter frequent deadlocks, make the sequence of lockingoperations (LOCK TABLES, SELECT ...FOR UPDATE, and so on) consistent between the differenttransactions or applications that experience the issue. SeeDeadlocks in InnoDB for details.", MySQLVersion:"5.6"},
    1214: {ErrorNumber:1214, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_FT", SQLState:"HY000", Message:"The used table type doesn't support FULLTEXT indexes", Description:"", MySQLVersion:"5.6"},
    1215: {ErrorNumber:1215, ErrorType:"ServerError", Symbol:"ER_CANNOT_ADD_FOREIGN", SQLState:"HY000", Message:"Cannot add foreign key constraint", Description:"", MySQLVersion:"5.6"},
    1216: {ErrorNumber:1216, ErrorType:"ServerError", Symbol:"ER_NO_REFERENCED_ROW", SQLState:"23000", Message:"Cannot add or update a child row: a foreign keyconstraint fails", Description:"InnoDB reports this error when you try to add arow but there is no parent row, and aforeign keyconstraint fails. Add the parent row first.", MySQLVersion:"5.6"},
    1217: {ErrorNumber:1217, ErrorType:"ServerError", Symbol:"ER_ROW_IS_REFERENCED", SQLState:"23000", Message:"Cannot delete or update a parent row: a foreign keyconstraint fails", Description:"InnoDB reports this error when you try todelete a parent row that has children, and aforeign keyconstraint fails. Delete the children first.", MySQLVersion:"5.6"},
    1218: {ErrorNumber:1218, ErrorType:"ServerError", Symbol:"ER_CONNECT_TO_MASTER", SQLState:"08S01", Message:"Error connecting to master: %s", Description:"", MySQLVersion:"5.6"},
    1219: {ErrorNumber:1219, ErrorType:"ServerError", Symbol:"ER_QUERY_ON_MASTER", SQLState:"HY000", Message:"Error running query on master: %s", Description:"", MySQLVersion:"5.6"},
    1220: {ErrorNumber:1220, ErrorType:"ServerError", Symbol:"ER_ERROR_WHEN_EXECUTING_COMMAND", SQLState:"HY000", Message:"Error when executing command %s: %s", Description:"", MySQLVersion:"5.6"},
    1221: {ErrorNumber:1221, ErrorType:"ServerError", Symbol:"ER_WRONG_USAGE", SQLState:"HY000", Message:"Incorrect usage of %s and %s", Description:"", MySQLVersion:"5.6"},
    1222: {ErrorNumber:1222, ErrorType:"ServerError", Symbol:"ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT", SQLState:"21000", Message:"The used SELECT statements have a different number ofcolumns", Description:"", MySQLVersion:"5.6"},
    1223: {ErrorNumber:1223, ErrorType:"ServerError", Symbol:"ER_CANT_UPDATE_WITH_READLOCK", SQLState:"HY000", Message:"Can't execute the query because you have a conflictingread lock", Description:"", MySQLVersion:"5.6"},
    1224: {ErrorNumber:1224, ErrorType:"ServerError", Symbol:"ER_MIXING_NOT_ALLOWED", SQLState:"HY000", Message:"Mixing of transactional and non-transactional tables isdisabled", Description:"", MySQLVersion:"5.6"},
    1225: {ErrorNumber:1225, ErrorType:"ServerError", Symbol:"ER_DUP_ARGUMENT", SQLState:"HY000", Message:"Option '%s' used twice in statement", Description:"", MySQLVersion:"5.6"},
    1226: {ErrorNumber:1226, ErrorType:"ServerError", Symbol:"ER_USER_LIMIT_REACHED", SQLState:"42000", Message:"User '%s' has exceeded the '%s' resource (current value:%ld)", Description:"", MySQLVersion:"5.6"},
    1227: {ErrorNumber:1227, ErrorType:"ServerError", Symbol:"ER_SPECIFIC_ACCESS_DENIED_ERROR", SQLState:"42000", Message:"Access denied", Description:"you need (at least one of) the %sprivilege(s) for this operation", MySQLVersion:"5.6"},
    1228: {ErrorNumber:1228, ErrorType:"ServerError", Symbol:"ER_LOCAL_VARIABLE", SQLState:"HY000", Message:"Variable '%s' is a SESSION variable and can't be usedwith SET GLOBAL", Description:"", MySQLVersion:"5.6"},
    1229: {ErrorNumber:1229, ErrorType:"ServerError", Symbol:"ER_GLOBAL_VARIABLE", SQLState:"HY000", Message:"Variable '%s' is a GLOBAL variable and should be set withSET GLOBAL", Description:"", MySQLVersion:"5.6"},
    1230: {ErrorNumber:1230, ErrorType:"ServerError", Symbol:"ER_NO_DEFAULT", SQLState:"42000", Message:"Variable '%s' doesn't have a default value", Description:"", MySQLVersion:"5.6"},
    1231: {ErrorNumber:1231, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_FOR_VAR", SQLState:"42000", Message:"Variable '%s' can't be set to the value of '%s'", Description:"", MySQLVersion:"5.6"},
    1232: {ErrorNumber:1232, ErrorType:"ServerError", Symbol:"ER_WRONG_TYPE_FOR_VAR", SQLState:"42000", Message:"Incorrect argument type to variable '%s'", Description:"", MySQLVersion:"5.6"},
    1233: {ErrorNumber:1233, ErrorType:"ServerError", Symbol:"ER_VAR_CANT_BE_READ", SQLState:"HY000", Message:"Variable '%s' can only be set, not read", Description:"", MySQLVersion:"5.6"},
    1234: {ErrorNumber:1234, ErrorType:"ServerError", Symbol:"ER_CANT_USE_OPTION_HERE", SQLState:"42000", Message:"Incorrect usage/placement of '%s'", Description:"", MySQLVersion:"5.6"},
    1235: {ErrorNumber:1235, ErrorType:"ServerError", Symbol:"ER_NOT_SUPPORTED_YET", SQLState:"42000", Message:"This version of MySQL doesn't yet support '%s'", Description:"", MySQLVersion:"5.6"},
    1236: {ErrorNumber:1236, ErrorType:"ServerError", Symbol:"ER_MASTER_FATAL_ERROR_READING_BINLOG", SQLState:"HY000", Message:"Got fatal error %d from master when reading data frombinary log: '%s'", Description:"", MySQLVersion:"5.6"},
    1237: {ErrorNumber:1237, ErrorType:"ServerError", Symbol:"ER_SLAVE_IGNORED_TABLE", SQLState:"HY000", Message:"Slave SQL thread ignored the query because ofreplicate-*-table rules", Description:"", MySQLVersion:"5.6"},
    1238: {ErrorNumber:1238, ErrorType:"ServerError", Symbol:"ER_INCORRECT_GLOBAL_LOCAL_VAR", SQLState:"HY000", Message:"Variable '%s' is a %s variable", Description:"", MySQLVersion:"5.6"},
    1239: {ErrorNumber:1239, ErrorType:"ServerError", Symbol:"ER_WRONG_FK_DEF", SQLState:"42000", Message:"Incorrect foreign key definition for '%s': %s", Description:"", MySQLVersion:"5.6"},
    1240: {ErrorNumber:1240, ErrorType:"ServerError", Symbol:"ER_KEY_REF_DO_NOT_MATCH_TABLE_REF", SQLState:"HY000", Message:"Key reference and table reference don't match", Description:"", MySQLVersion:"5.6"},
    1241: {ErrorNumber:1241, ErrorType:"ServerError", Symbol:"ER_OPERAND_COLUMNS", SQLState:"21000", Message:"Operand should contain %d column(s)", Description:"", MySQLVersion:"5.6"},
    1242: {ErrorNumber:1242, ErrorType:"ServerError", Symbol:"ER_SUBQUERY_NO_1_ROW", SQLState:"21000", Message:"Subquery returns more than 1 row", Description:"", MySQLVersion:"5.6"},
    1243: {ErrorNumber:1243, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_STMT_HANDLER", SQLState:"HY000", Message:"Unknown prepared statement handler (%.*s) given to %s", Description:"", MySQLVersion:"5.6"},
    1244: {ErrorNumber:1244, ErrorType:"ServerError", Symbol:"ER_CORRUPT_HELP_DB", SQLState:"HY000", Message:"Help database is corrupt or does not exist", Description:"", MySQLVersion:"5.6"},
    1245: {ErrorNumber:1245, ErrorType:"ServerError", Symbol:"ER_CYCLIC_REFERENCE", SQLState:"HY000", Message:"Cyclic reference on subqueries", Description:"", MySQLVersion:"5.6"},
    1246: {ErrorNumber:1246, ErrorType:"ServerError", Symbol:"ER_AUTO_CONVERT", SQLState:"HY000", Message:"Converting column '%s' from %s to %s", Description:"", MySQLVersion:"5.6"},
    1247: {ErrorNumber:1247, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_REFERENCE", SQLState:"42S22", Message:"Reference '%s' not supported (%s)", Description:"", MySQLVersion:"5.6"},
    1248: {ErrorNumber:1248, ErrorType:"ServerError", Symbol:"ER_DERIVED_MUST_HAVE_ALIAS", SQLState:"42000", Message:"Every derived table must have its own alias", Description:"", MySQLVersion:"5.6"},
    1249: {ErrorNumber:1249, ErrorType:"ServerError", Symbol:"ER_SELECT_REDUCED", SQLState:"01000", Message:"Select %u was reduced during optimization", Description:"", MySQLVersion:"5.6"},
    1250: {ErrorNumber:1250, ErrorType:"ServerError", Symbol:"ER_TABLENAME_NOT_ALLOWED_HERE", SQLState:"42000", Message:"Table '%s' from one of the SELECTs cannot be used in %s", Description:"", MySQLVersion:"5.6"},
    1251: {ErrorNumber:1251, ErrorType:"ServerError", Symbol:"ER_NOT_SUPPORTED_AUTH_MODE", SQLState:"08004", Message:"Client does not support authentication protocol requestedby server", Description:"consider upgrading MySQL client", MySQLVersion:"5.6"},
    1252: {ErrorNumber:1252, ErrorType:"ServerError", Symbol:"ER_SPATIAL_CANT_HAVE_NULL", SQLState:"42000", Message:"All parts of a SPATIAL index must be NOT NULL", Description:"", MySQLVersion:"5.6"},
    1253: {ErrorNumber:1253, ErrorType:"ServerError", Symbol:"ER_COLLATION_CHARSET_MISMATCH", SQLState:"42000", Message:"COLLATION '%s' is not valid for CHARACTER SET '%s'", Description:"", MySQLVersion:"5.6"},
    1254: {ErrorNumber:1254, ErrorType:"ServerError", Symbol:"ER_SLAVE_WAS_RUNNING", SQLState:"HY000", Message:"Slave is already running", Description:"", MySQLVersion:"5.6"},
    1255: {ErrorNumber:1255, ErrorType:"ServerError", Symbol:"ER_SLAVE_WAS_NOT_RUNNING", SQLState:"HY000", Message:"Slave already has been stopped", Description:"", MySQLVersion:"5.6"},
    1256: {ErrorNumber:1256, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_FOR_UNCOMPRESS", SQLState:"HY000", Message:"Uncompressed data size too large", Description:"the maximum size is %d(probably, length of uncompressed data was corrupted)", MySQLVersion:"5.6"},
    1257: {ErrorNumber:1257, ErrorType:"ServerError", Symbol:"ER_ZLIB_Z_MEM_ERROR", SQLState:"HY000", Message:"ZLIB: Not enough memory", Description:"", MySQLVersion:"5.6"},
    1258: {ErrorNumber:1258, ErrorType:"ServerError", Symbol:"ER_ZLIB_Z_BUF_ERROR", SQLState:"HY000", Message:"ZLIB: Not enough room in the output buffer (probably,length of uncompressed data was corrupted)", Description:"", MySQLVersion:"5.6"},
    1259: {ErrorNumber:1259, ErrorType:"ServerError", Symbol:"ER_ZLIB_Z_DATA_ERROR", SQLState:"HY000", Message:"ZLIB: Input data corrupted", Description:"", MySQLVersion:"5.6"},
    1260: {ErrorNumber:1260, ErrorType:"ServerError", Symbol:"ER_CUT_VALUE_GROUP_CONCAT", SQLState:"HY000", Message:"Row %u was cut by GROUP_CONCAT()", Description:"", MySQLVersion:"5.6"},
    1261: {ErrorNumber:1261, ErrorType:"ServerError", Symbol:"ER_WARN_TOO_FEW_RECORDS", SQLState:"01000", Message:"Row %ld doesn't contain data for all columns", Description:"", MySQLVersion:"5.6"},
    1262: {ErrorNumber:1262, ErrorType:"ServerError", Symbol:"ER_WARN_TOO_MANY_RECORDS", SQLState:"01000", Message:"Row %ld was truncated", Description:"it contained more data than therewere input columns", MySQLVersion:"5.6"},
    1263: {ErrorNumber:1263, ErrorType:"ServerError", Symbol:"ER_WARN_NULL_TO_NOTNULL", SQLState:"22004", Message:"Column set to default value", Description:"NULL supplied to NOT NULLcolumn '%s' at row %ld", MySQLVersion:"5.6"},
    1264: {ErrorNumber:1264, ErrorType:"ServerError", Symbol:"ER_WARN_DATA_OUT_OF_RANGE", SQLState:"22003", Message:"Out of range value for column '%s' at row %ld", Description:"", MySQLVersion:"5.6"},
    1265: {ErrorNumber:1265, ErrorType:"ServerError", Symbol:"WARN_DATA_TRUNCATED", SQLState:"01000", Message:"Data truncated for column '%s' at row %ld", Description:"", MySQLVersion:"5.6"},
    1266: {ErrorNumber:1266, ErrorType:"ServerError", Symbol:"ER_WARN_USING_OTHER_HANDLER", SQLState:"HY000", Message:"Using storage engine %s for table '%s'", Description:"", MySQLVersion:"5.6"},
    1267: {ErrorNumber:1267, ErrorType:"ServerError", Symbol:"ER_CANT_AGGREGATE_2COLLATIONS", SQLState:"HY000", Message:"Illegal mix of collations (%s,%s) and (%s,%s) foroperation '%s'", Description:"", MySQLVersion:"5.6"},
    1268: {ErrorNumber:1268, ErrorType:"ServerError", Symbol:"ER_DROP_USER", SQLState:"HY000", Message:"Cannot drop one or more of the requested users", Description:"", MySQLVersion:"5.6"},
    1269: {ErrorNumber:1269, ErrorType:"ServerError", Symbol:"ER_REVOKE_GRANTS", SQLState:"HY000", Message:"Can't revoke all privileges for one or more of therequested users", Description:"", MySQLVersion:"5.6"},
    1270: {ErrorNumber:1270, ErrorType:"ServerError", Symbol:"ER_CANT_AGGREGATE_3COLLATIONS", SQLState:"HY000", Message:"Illegal mix of collations (%s,%s), (%s,%s), (%s,%s) foroperation '%s'", Description:"", MySQLVersion:"5.6"},
    1271: {ErrorNumber:1271, ErrorType:"ServerError", Symbol:"ER_CANT_AGGREGATE_NCOLLATIONS", SQLState:"HY000", Message:"Illegal mix of collations for operation '%s'", Description:"", MySQLVersion:"5.6"},
    1272: {ErrorNumber:1272, ErrorType:"ServerError", Symbol:"ER_VARIABLE_IS_NOT_STRUCT", SQLState:"HY000", Message:"Variable '%s' is not a variable component (can't be usedas XXXX.variable_name)", Description:"", MySQLVersion:"5.6"},
    1273: {ErrorNumber:1273, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_COLLATION", SQLState:"HY000", Message:"Unknown collation: '%s'", Description:"", MySQLVersion:"5.6"},
    1274: {ErrorNumber:1274, ErrorType:"ServerError", Symbol:"ER_SLAVE_IGNORED_SSL_PARAMS", SQLState:"HY000", Message:"SSL parameters in CHANGE MASTER are ignored because thisMySQL slave was compiled without SSL support", Description:"they can be usedlater if MySQL slave with SSL is started", MySQLVersion:"5.6"},
    1275: {ErrorNumber:1275, ErrorType:"ServerError", Symbol:"ER_SERVER_IS_IN_SECURE_AUTH_MODE", SQLState:"HY000", Message:"Server is running in --secure-auth mode, but '%s'@'%s'has a password in the old format", Description:"please change the password tothe new format", MySQLVersion:"5.6"},
    1276: {ErrorNumber:1276, ErrorType:"ServerError", Symbol:"ER_WARN_FIELD_RESOLVED", SQLState:"HY000", Message:"Field or reference '%s%s%s%s%s' of SELECT #%d wasresolved in SELECT #%d", Description:"", MySQLVersion:"5.6"},
    1277: {ErrorNumber:1277, ErrorType:"ServerError", Symbol:"ER_BAD_SLAVE_UNTIL_COND", SQLState:"HY000", Message:"Incorrect parameter or combination of parameters forSTART SLAVE UNTIL", Description:"", MySQLVersion:"5.6"},
    1278: {ErrorNumber:1278, ErrorType:"ServerError", Symbol:"ER_MISSING_SKIP_SLAVE", SQLState:"HY000", Message:"It is recommended to use --skip-slave-start when doingstep-by-step replication with START SLAVE UNTIL", Description:"otherwise, youwill get problems if you get an unexpected slave's mysqld restart", MySQLVersion:"5.6"},
    1279: {ErrorNumber:1279, ErrorType:"ServerError", Symbol:"ER_UNTIL_COND_IGNORED", SQLState:"HY000", Message:"SQL thread is not to be started so UNTIL options areignored", Description:"", MySQLVersion:"5.6"},
    1280: {ErrorNumber:1280, ErrorType:"ServerError", Symbol:"ER_WRONG_NAME_FOR_INDEX", SQLState:"42000", Message:"Incorrect index name '%s'", Description:"", MySQLVersion:"5.6"},
    1281: {ErrorNumber:1281, ErrorType:"ServerError", Symbol:"ER_WRONG_NAME_FOR_CATALOG", SQLState:"42000", Message:"Incorrect catalog name '%s'", Description:"", MySQLVersion:"5.6"},
    1282: {ErrorNumber:1282, ErrorType:"ServerError", Symbol:"ER_WARN_QC_RESIZE", SQLState:"HY000", Message:"Query cache failed to set size %lu", Description:"new query cache sizeis %lu", MySQLVersion:"5.6"},
    1283: {ErrorNumber:1283, ErrorType:"ServerError", Symbol:"ER_BAD_FT_COLUMN", SQLState:"HY000", Message:"Column '%s' cannot be part of FULLTEXT index", Description:"", MySQLVersion:"5.6"},
    1284: {ErrorNumber:1284, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_KEY_CACHE", SQLState:"HY000", Message:"Unknown key cache '%s'", Description:"", MySQLVersion:"5.6"},
    1285: {ErrorNumber:1285, ErrorType:"ServerError", Symbol:"ER_WARN_HOSTNAME_WONT_WORK", SQLState:"HY000", Message:"MySQL is started in --skip-name-resolve mode", Description:"you mustrestart it without this switch for this grant to work", MySQLVersion:"5.6"},
    1286: {ErrorNumber:1286, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_STORAGE_ENGINE", SQLState:"42000", Message:"Unknown storage engine '%s'", Description:"", MySQLVersion:"5.6"},
    1287: {ErrorNumber:1287, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYNTAX", SQLState:"HY000", Message:"'%s' is deprecated and will be removed in a futurerelease. Please use %s instead", Description:"", MySQLVersion:"5.6"},
    1288: {ErrorNumber:1288, ErrorType:"ServerError", Symbol:"ER_NON_UPDATABLE_TABLE", SQLState:"HY000", Message:"The target table %s of the %s is not updatable", Description:"", MySQLVersion:"5.6"},
    1289: {ErrorNumber:1289, ErrorType:"ServerError", Symbol:"ER_FEATURE_DISABLED", SQLState:"HY000", Message:"The '%s' feature is disabled", Description:"you need MySQL built with'%s' to have it working", MySQLVersion:"5.6"},
    1290: {ErrorNumber:1290, ErrorType:"ServerError", Symbol:"ER_OPTION_PREVENTS_STATEMENT", SQLState:"HY000", Message:"The MySQL server is running with the %s option so itcannot execute this statement", Description:"", MySQLVersion:"5.6"},
    1291: {ErrorNumber:1291, ErrorType:"ServerError", Symbol:"ER_DUPLICATED_VALUE_IN_TYPE", SQLState:"HY000", Message:"Column '%s' has duplicated value '%s' in %s", Description:"", MySQLVersion:"5.6"},
    1292: {ErrorNumber:1292, ErrorType:"ServerError", Symbol:"ER_TRUNCATED_WRONG_VALUE", SQLState:"22007", Message:"Truncated incorrect %s value: '%s'", Description:"", MySQLVersion:"5.6"},
    1293: {ErrorNumber:1293, ErrorType:"ServerError", Symbol:"ER_TOO_MUCH_AUTO_TIMESTAMP_COLS", SQLState:"HY000", Message:"Incorrect table definition", Description:"there can be only oneTIMESTAMP column with CURRENT_TIMESTAMP in DEFAULT or ON UPDATEclause", MySQLVersion:"5.6"},
    1294: {ErrorNumber:1294, ErrorType:"ServerError", Symbol:"ER_INVALID_ON_UPDATE", SQLState:"HY000", Message:"Invalid ON UPDATE clause for '%s' column", Description:"", MySQLVersion:"5.6"},
    1295: {ErrorNumber:1295, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_PS", SQLState:"HY000", Message:"This command is not supported in the prepared statementprotocol yet", Description:"", MySQLVersion:"5.6"},
    1296: {ErrorNumber:1296, ErrorType:"ServerError", Symbol:"ER_GET_ERRMSG", SQLState:"HY000", Message:"Got error %d '%s' from %s", Description:"", MySQLVersion:"5.6"},
    1297: {ErrorNumber:1297, ErrorType:"ServerError", Symbol:"ER_GET_TEMPORARY_ERRMSG", SQLState:"HY000", Message:"Got temporary error %d '%s' from %s", Description:"", MySQLVersion:"5.6"},
    1298: {ErrorNumber:1298, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_TIME_ZONE", SQLState:"HY000", Message:"Unknown or incorrect time zone: '%s'", Description:"", MySQLVersion:"5.6"},
    1299: {ErrorNumber:1299, ErrorType:"ServerError", Symbol:"ER_WARN_INVALID_TIMESTAMP", SQLState:"HY000", Message:"Invalid TIMESTAMP value in column '%s' at row %ld", Description:"", MySQLVersion:"5.6"},
    1300: {ErrorNumber:1300, ErrorType:"ServerError", Symbol:"ER_INVALID_CHARACTER_STRING", SQLState:"HY000", Message:"Invalid %s character string: '%s'", Description:"", MySQLVersion:"5.6"},
    1301: {ErrorNumber:1301, ErrorType:"ServerError", Symbol:"ER_WARN_ALLOWED_PACKET_OVERFLOWED", SQLState:"HY000", Message:"Result of %s() was larger than max_allowed_packet (%ld) -truncated", Description:"", MySQLVersion:"5.6"},
    1302: {ErrorNumber:1302, ErrorType:"ServerError", Symbol:"ER_CONFLICTING_DECLARATIONS", SQLState:"HY000", Message:"Conflicting declarations: '%s%s' and '%s%s'", Description:"", MySQLVersion:"5.6"},
    1303: {ErrorNumber:1303, ErrorType:"ServerError", Symbol:"ER_SP_NO_RECURSIVE_CREATE", SQLState:"2F003", Message:"Can't create a %s from within another stored routine", Description:"", MySQLVersion:"5.6"},
    1304: {ErrorNumber:1304, ErrorType:"ServerError", Symbol:"ER_SP_ALREADY_EXISTS", SQLState:"42000", Message:"%s %s already exists", Description:"", MySQLVersion:"5.6"},
    1305: {ErrorNumber:1305, ErrorType:"ServerError", Symbol:"ER_SP_DOES_NOT_EXIST", SQLState:"42000", Message:"%s %s does not exist", Description:"", MySQLVersion:"5.6"},
    1306: {ErrorNumber:1306, ErrorType:"ServerError", Symbol:"ER_SP_DROP_FAILED", SQLState:"HY000", Message:"Failed to DROP %s %s", Description:"", MySQLVersion:"5.6"},
    1307: {ErrorNumber:1307, ErrorType:"ServerError", Symbol:"ER_SP_STORE_FAILED", SQLState:"HY000", Message:"Failed to CREATE %s %s", Description:"", MySQLVersion:"5.6"},
    1308: {ErrorNumber:1308, ErrorType:"ServerError", Symbol:"ER_SP_LILABEL_MISMATCH", SQLState:"42000", Message:"%s with no matching label: %s", Description:"", MySQLVersion:"5.6"},
    1309: {ErrorNumber:1309, ErrorType:"ServerError", Symbol:"ER_SP_LABEL_REDEFINE", SQLState:"42000", Message:"Redefining label %s", Description:"", MySQLVersion:"5.6"},
    1310: {ErrorNumber:1310, ErrorType:"ServerError", Symbol:"ER_SP_LABEL_MISMATCH", SQLState:"42000", Message:"End-label %s without match", Description:"", MySQLVersion:"5.6"},
    1311: {ErrorNumber:1311, ErrorType:"ServerError", Symbol:"ER_SP_UNINIT_VAR", SQLState:"01000", Message:"Referring to uninitialized variable %s", Description:"", MySQLVersion:"5.6"},
    1312: {ErrorNumber:1312, ErrorType:"ServerError", Symbol:"ER_SP_BADSELECT", SQLState:"0A000", Message:"PROCEDURE %s can't return a result set in the givencontext", Description:"", MySQLVersion:"5.6"},
    1313: {ErrorNumber:1313, ErrorType:"ServerError", Symbol:"ER_SP_BADRETURN", SQLState:"42000", Message:"RETURN is only allowed in a FUNCTION", Description:"", MySQLVersion:"5.6"},
    1314: {ErrorNumber:1314, ErrorType:"ServerError", Symbol:"ER_SP_BADSTATEMENT", SQLState:"0A000", Message:"%s is not allowed in stored procedures", Description:"", MySQLVersion:"5.6"},
    1315: {ErrorNumber:1315, ErrorType:"ServerError", Symbol:"ER_UPDATE_LOG_DEPRECATED_IGNORED", SQLState:"42000", Message:"The update log is deprecated and replaced by the binarylog", Description:"SET SQL_LOG_UPDATE has been ignored.", MySQLVersion:"5.6"},
    1316: {ErrorNumber:1316, ErrorType:"ServerError", Symbol:"ER_UPDATE_LOG_DEPRECATED_TRANSLATED", SQLState:"42000", Message:"The update log is deprecated and replaced by the binarylog", Description:"SET SQL_LOG_UPDATE has been translated to SET SQL_LOG_BIN.", MySQLVersion:"5.6"},
    1317: {ErrorNumber:1317, ErrorType:"ServerError", Symbol:"ER_QUERY_INTERRUPTED", SQLState:"70100", Message:"Query execution was interrupted", Description:"", MySQLVersion:"5.6"},
    1318: {ErrorNumber:1318, ErrorType:"ServerError", Symbol:"ER_SP_WRONG_NO_OF_ARGS", SQLState:"42000", Message:"Incorrect number of arguments for %s %s", Description:"expected %u, got%u", MySQLVersion:"5.6"},
    1319: {ErrorNumber:1319, ErrorType:"ServerError", Symbol:"ER_SP_COND_MISMATCH", SQLState:"42000", Message:"Undefined CONDITION: %s", Description:"", MySQLVersion:"5.6"},
    1320: {ErrorNumber:1320, ErrorType:"ServerError", Symbol:"ER_SP_NORETURN", SQLState:"42000", Message:"No RETURN found in FUNCTION %s", Description:"", MySQLVersion:"5.6"},
    1321: {ErrorNumber:1321, ErrorType:"ServerError", Symbol:"ER_SP_NORETURNEND", SQLState:"2F005", Message:"FUNCTION %s ended without RETURN", Description:"", MySQLVersion:"5.6"},
    1322: {ErrorNumber:1322, ErrorType:"ServerError", Symbol:"ER_SP_BAD_CURSOR_QUERY", SQLState:"42000", Message:"Cursor statement must be a SELECT", Description:"", MySQLVersion:"5.6"},
    1323: {ErrorNumber:1323, ErrorType:"ServerError", Symbol:"ER_SP_BAD_CURSOR_SELECT", SQLState:"42000", Message:"Cursor SELECT must not have INTO", Description:"", MySQLVersion:"5.6"},
    1324: {ErrorNumber:1324, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_MISMATCH", SQLState:"42000", Message:"Undefined CURSOR: %s", Description:"", MySQLVersion:"5.6"},
    1325: {ErrorNumber:1325, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_ALREADY_OPEN", SQLState:"24000", Message:"Cursor is already open", Description:"", MySQLVersion:"5.6"},
    1326: {ErrorNumber:1326, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_NOT_OPEN", SQLState:"24000", Message:"Cursor is not open", Description:"", MySQLVersion:"5.6"},
    1327: {ErrorNumber:1327, ErrorType:"ServerError", Symbol:"ER_SP_UNDECLARED_VAR", SQLState:"42000", Message:"Undeclared variable: %s", Description:"", MySQLVersion:"5.6"},
    1328: {ErrorNumber:1328, ErrorType:"ServerError", Symbol:"ER_SP_WRONG_NO_OF_FETCH_ARGS", SQLState:"HY000", Message:"Incorrect number of FETCH variables", Description:"", MySQLVersion:"5.6"},
    1329: {ErrorNumber:1329, ErrorType:"ServerError", Symbol:"ER_SP_FETCH_NO_DATA", SQLState:"02000", Message:"No data - zero rows fetched, selected, or processed", Description:"", MySQLVersion:"5.6"},
    1330: {ErrorNumber:1330, ErrorType:"ServerError", Symbol:"ER_SP_DUP_PARAM", SQLState:"42000", Message:"Duplicate parameter: %s", Description:"", MySQLVersion:"5.6"},
    1331: {ErrorNumber:1331, ErrorType:"ServerError", Symbol:"ER_SP_DUP_VAR", SQLState:"42000", Message:"Duplicate variable: %s", Description:"", MySQLVersion:"5.6"},
    1332: {ErrorNumber:1332, ErrorType:"ServerError", Symbol:"ER_SP_DUP_COND", SQLState:"42000", Message:"Duplicate condition: %s", Description:"", MySQLVersion:"5.6"},
    1333: {ErrorNumber:1333, ErrorType:"ServerError", Symbol:"ER_SP_DUP_CURS", SQLState:"42000", Message:"Duplicate cursor: %s", Description:"", MySQLVersion:"5.6"},
    1334: {ErrorNumber:1334, ErrorType:"ServerError", Symbol:"ER_SP_CANT_ALTER", SQLState:"HY000", Message:"Failed to ALTER %s %s", Description:"", MySQLVersion:"5.6"},
    1335: {ErrorNumber:1335, ErrorType:"ServerError", Symbol:"ER_SP_SUBSELECT_NYI", SQLState:"0A000", Message:"Subquery value not supported", Description:"", MySQLVersion:"5.6"},
    1336: {ErrorNumber:1336, ErrorType:"ServerError", Symbol:"ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG", SQLState:"0A000", Message:"%s is not allowed in stored function or trigger", Description:"", MySQLVersion:"5.6"},
    1337: {ErrorNumber:1337, ErrorType:"ServerError", Symbol:"ER_SP_VARCOND_AFTER_CURSHNDLR", SQLState:"42000", Message:"Variable or condition declaration after cursor or handlerdeclaration", Description:"", MySQLVersion:"5.6"},
    1338: {ErrorNumber:1338, ErrorType:"ServerError", Symbol:"ER_SP_CURSOR_AFTER_HANDLER", SQLState:"42000", Message:"Cursor declaration after handler declaration", Description:"", MySQLVersion:"5.6"},
    1339: {ErrorNumber:1339, ErrorType:"ServerError", Symbol:"ER_SP_CASE_NOT_FOUND", SQLState:"20000", Message:"Case not found for CASE statement", Description:"", MySQLVersion:"5.6"},
    1340: {ErrorNumber:1340, ErrorType:"ServerError", Symbol:"ER_FPARSER_TOO_BIG_FILE", SQLState:"HY000", Message:"Configuration file '%s' is too big", Description:"", MySQLVersion:"5.6"},
    1341: {ErrorNumber:1341, ErrorType:"ServerError", Symbol:"ER_FPARSER_BAD_HEADER", SQLState:"HY000", Message:"Malformed file type header in file '%s'", Description:"", MySQLVersion:"5.6"},
    1342: {ErrorNumber:1342, ErrorType:"ServerError", Symbol:"ER_FPARSER_EOF_IN_COMMENT", SQLState:"HY000", Message:"Unexpected end of file while parsing comment '%s'", Description:"", MySQLVersion:"5.6"},
    1343: {ErrorNumber:1343, ErrorType:"ServerError", Symbol:"ER_FPARSER_ERROR_IN_PARAMETER", SQLState:"HY000", Message:"Error while parsing parameter '%s' (line: '%s')", Description:"", MySQLVersion:"5.6"},
    1344: {ErrorNumber:1344, ErrorType:"ServerError", Symbol:"ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER", SQLState:"HY000", Message:"Unexpected end of file while skipping unknown parameter'%s'", Description:"", MySQLVersion:"5.6"},
    1345: {ErrorNumber:1345, ErrorType:"ServerError", Symbol:"ER_VIEW_NO_EXPLAIN", SQLState:"HY000", Message:"EXPLAIN/SHOW can not be issued", Description:"lacking privileges forunderlying table", MySQLVersion:"5.6"},
    1346: {ErrorNumber:1346, ErrorType:"ServerError", Symbol:"ER_FRM_UNKNOWN_TYPE", SQLState:"HY000", Message:"File '%s' has unknown type '%s' in its header", Description:"", MySQLVersion:"5.6"},
    1347: {ErrorNumber:1347, ErrorType:"ServerError", Symbol:"ER_WRONG_OBJECT", SQLState:"HY000", Message:"'%s.%s' is not %s", Description:"The named object is incorrect for the type of operation attemptedon it. It must be an object of the named type.", MySQLVersion:"5.6"},
    1348: {ErrorNumber:1348, ErrorType:"ServerError", Symbol:"ER_NONUPDATEABLE_COLUMN", SQLState:"HY000", Message:"Column '%s' is not updatable", Description:"", MySQLVersion:"5.6"},
    1349: {ErrorNumber:1349, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_DERIVED", SQLState:"HY000", Message:"View's SELECT contains a subquery in the FROM clause", Description:"", MySQLVersion:"5.6"},
    1350: {ErrorNumber:1350, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_CLAUSE", SQLState:"HY000", Message:"View's SELECT contains a '%s' clause", Description:"", MySQLVersion:"5.6"},
    1351: {ErrorNumber:1351, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_VARIABLE", SQLState:"HY000", Message:"View's SELECT contains a variable or parameter", Description:"", MySQLVersion:"5.6"},
    1352: {ErrorNumber:1352, ErrorType:"ServerError", Symbol:"ER_VIEW_SELECT_TMPTABLE", SQLState:"HY000", Message:"View's SELECT refers to a temporary table '%s'", Description:"", MySQLVersion:"5.6"},
    1353: {ErrorNumber:1353, ErrorType:"ServerError", Symbol:"ER_VIEW_WRONG_LIST", SQLState:"HY000", Message:"View's SELECT and view's field list have different columncounts", Description:"", MySQLVersion:"5.6"},
    1354: {ErrorNumber:1354, ErrorType:"ServerError", Symbol:"ER_WARN_VIEW_MERGE", SQLState:"HY000", Message:"View merge algorithm can't be used here for now (assumedundefined algorithm)", Description:"", MySQLVersion:"5.6"},
    1355: {ErrorNumber:1355, ErrorType:"ServerError", Symbol:"ER_WARN_VIEW_WITHOUT_KEY", SQLState:"HY000", Message:"View being updated does not have complete key ofunderlying table in it", Description:"", MySQLVersion:"5.6"},
    1356: {ErrorNumber:1356, ErrorType:"ServerError", Symbol:"ER_VIEW_INVALID", SQLState:"HY000", Message:"View '%s.%s' references invalid table(s) or column(s) orfunction(s) or definer/invoker of view lack rights to use them", Description:"", MySQLVersion:"5.6"},
    1357: {ErrorNumber:1357, ErrorType:"ServerError", Symbol:"ER_SP_NO_DROP_SP", SQLState:"HY000", Message:"Can't drop or alter a %s from within another storedroutine", Description:"", MySQLVersion:"5.6"},
    1358: {ErrorNumber:1358, ErrorType:"ServerError", Symbol:"ER_SP_GOTO_IN_HNDLR", SQLState:"HY000", Message:"GOTO is not allowed in a stored procedure handler", Description:"", MySQLVersion:"5.6"},
    1359: {ErrorNumber:1359, ErrorType:"ServerError", Symbol:"ER_TRG_ALREADY_EXISTS", SQLState:"HY000", Message:"Trigger already exists", Description:"", MySQLVersion:"5.6"},
    1360: {ErrorNumber:1360, ErrorType:"ServerError", Symbol:"ER_TRG_DOES_NOT_EXIST", SQLState:"HY000", Message:"Trigger does not exist", Description:"", MySQLVersion:"5.6"},
    1361: {ErrorNumber:1361, ErrorType:"ServerError", Symbol:"ER_TRG_ON_VIEW_OR_TEMP_TABLE", SQLState:"HY000", Message:"Trigger's '%s' is view or temporary table", Description:"", MySQLVersion:"5.6"},
    1362: {ErrorNumber:1362, ErrorType:"ServerError", Symbol:"ER_TRG_CANT_CHANGE_ROW", SQLState:"HY000", Message:"Updating of %s row is not allowed in %strigger", Description:"", MySQLVersion:"5.6"},
    1363: {ErrorNumber:1363, ErrorType:"ServerError", Symbol:"ER_TRG_NO_SUCH_ROW_IN_TRG", SQLState:"HY000", Message:"There is no %s row in %s trigger", Description:"", MySQLVersion:"5.6"},
    1364: {ErrorNumber:1364, ErrorType:"ServerError", Symbol:"ER_NO_DEFAULT_FOR_FIELD", SQLState:"HY000", Message:"Field '%s' doesn't have a default value", Description:"", MySQLVersion:"5.6"},
    1365: {ErrorNumber:1365, ErrorType:"ServerError", Symbol:"ER_DIVISION_BY_ZERO", SQLState:"22012", Message:"Division by 0", Description:"", MySQLVersion:"5.6"},
    1366: {ErrorNumber:1366, ErrorType:"ServerError", Symbol:"ER_TRUNCATED_WRONG_VALUE_FOR_FIELD", SQLState:"HY000", Message:"Incorrect %s value: '%s' for column '%s' at row %ld", Description:"", MySQLVersion:"5.6"},
    1367: {ErrorNumber:1367, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_VALUE_FOR_TYPE", SQLState:"22007", Message:"Illegal %s '%s' value found during parsing", Description:"", MySQLVersion:"5.6"},
    1368: {ErrorNumber:1368, ErrorType:"ServerError", Symbol:"ER_VIEW_NONUPD_CHECK", SQLState:"HY000", Message:"CHECK OPTION on non-updatable view '%s.%s'", Description:"", MySQLVersion:"5.6"},
    1369: {ErrorNumber:1369, ErrorType:"ServerError", Symbol:"ER_VIEW_CHECK_FAILED", SQLState:"HY000", Message:"CHECK OPTION failed '%s.%s'", Description:"", MySQLVersion:"5.6"},
    1370: {ErrorNumber:1370, ErrorType:"ServerError", Symbol:"ER_PROCACCESS_DENIED_ERROR", SQLState:"42000", Message:"%s command denied to user '%s'@'%s' for routine '%s'", Description:"", MySQLVersion:"5.6"},
    1371: {ErrorNumber:1371, ErrorType:"ServerError", Symbol:"ER_RELAY_LOG_FAIL", SQLState:"HY000", Message:"Failed purging old relay logs: %s", Description:"", MySQLVersion:"5.6"},
    1372: {ErrorNumber:1372, ErrorType:"ServerError", Symbol:"ER_PASSWD_LENGTH", SQLState:"HY000", Message:"Password hash should be a %d-digit hexadecimal number", Description:"", MySQLVersion:"5.6"},
    1373: {ErrorNumber:1373, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_TARGET_BINLOG", SQLState:"HY000", Message:"Target log not found in binlog index", Description:"", MySQLVersion:"5.6"},
    1374: {ErrorNumber:1374, ErrorType:"ServerError", Symbol:"ER_IO_ERR_LOG_INDEX_READ", SQLState:"HY000", Message:"I/O error reading log index file", Description:"", MySQLVersion:"5.6"},
    1375: {ErrorNumber:1375, ErrorType:"ServerError", Symbol:"ER_BINLOG_PURGE_PROHIBITED", SQLState:"HY000", Message:"Server configuration does not permit binlog purge", Description:"", MySQLVersion:"5.6"},
    1376: {ErrorNumber:1376, ErrorType:"ServerError", Symbol:"ER_FSEEK_FAIL", SQLState:"HY000", Message:"Failed on fseek()", Description:"", MySQLVersion:"5.6"},
    1377: {ErrorNumber:1377, ErrorType:"ServerError", Symbol:"ER_BINLOG_PURGE_FATAL_ERR", SQLState:"HY000", Message:"Fatal error during log purge", Description:"", MySQLVersion:"5.6"},
    1378: {ErrorNumber:1378, ErrorType:"ServerError", Symbol:"ER_LOG_IN_USE", SQLState:"HY000", Message:"A purgeable log is in use, will not purge", Description:"", MySQLVersion:"5.6"},
    1379: {ErrorNumber:1379, ErrorType:"ServerError", Symbol:"ER_LOG_PURGE_UNKNOWN_ERR", SQLState:"HY000", Message:"Unknown error during log purge", Description:"", MySQLVersion:"5.6"},
    1380: {ErrorNumber:1380, ErrorType:"ServerError", Symbol:"ER_RELAY_LOG_INIT", SQLState:"HY000", Message:"Failed initializing relay log position: %s", Description:"", MySQLVersion:"5.6"},
    1381: {ErrorNumber:1381, ErrorType:"ServerError", Symbol:"ER_NO_BINARY_LOGGING", SQLState:"HY000", Message:"You are not using binary logging", Description:"", MySQLVersion:"5.6"},
    1382: {ErrorNumber:1382, ErrorType:"ServerError", Symbol:"ER_RESERVED_SYNTAX", SQLState:"HY000", Message:"The '%s' syntax is reserved for purposes internal to theMySQL server", Description:"", MySQLVersion:"5.6"},
    1383: {ErrorNumber:1383, ErrorType:"ServerError", Symbol:"ER_WSAS_FAILED", SQLState:"HY000", Message:"WSAStartup Failed", Description:"", MySQLVersion:"5.6"},
    1384: {ErrorNumber:1384, ErrorType:"ServerError", Symbol:"ER_DIFF_GROUPS_PROC", SQLState:"HY000", Message:"Can't handle procedures with different groups yet", Description:"", MySQLVersion:"5.6"},
    1385: {ErrorNumber:1385, ErrorType:"ServerError", Symbol:"ER_NO_GROUP_FOR_PROC", SQLState:"HY000", Message:"Select must have a group with this procedure", Description:"", MySQLVersion:"5.6"},
    1386: {ErrorNumber:1386, ErrorType:"ServerError", Symbol:"ER_ORDER_WITH_PROC", SQLState:"HY000", Message:"Can't use ORDER clause with this procedure", Description:"", MySQLVersion:"5.6"},
    1387: {ErrorNumber:1387, ErrorType:"ServerError", Symbol:"ER_LOGGING_PROHIBIT_CHANGING_OF", SQLState:"HY000", Message:"Binary logging and replication forbid changing the globalserver %s", Description:"", MySQLVersion:"5.6"},
    1388: {ErrorNumber:1388, ErrorType:"ServerError", Symbol:"ER_NO_FILE_MAPPING", SQLState:"HY000", Message:"Can't map file: %s, errno: %d", Description:"", MySQLVersion:"5.6"},
    1389: {ErrorNumber:1389, ErrorType:"ServerError", Symbol:"ER_WRONG_MAGIC", SQLState:"HY000", Message:"Wrong magic in %s", Description:"", MySQLVersion:"5.6"},
    1390: {ErrorNumber:1390, ErrorType:"ServerError", Symbol:"ER_PS_MANY_PARAM", SQLState:"HY000", Message:"Prepared statement contains too many placeholders", Description:"", MySQLVersion:"5.6"},
    1391: {ErrorNumber:1391, ErrorType:"ServerError", Symbol:"ER_KEY_PART_0", SQLState:"HY000", Message:"Key part '%s' length cannot be 0", Description:"", MySQLVersion:"5.6"},
    1392: {ErrorNumber:1392, ErrorType:"ServerError", Symbol:"ER_VIEW_CHECKSUM", SQLState:"HY000", Message:"View text checksum failed", Description:"", MySQLVersion:"5.6"},
    1393: {ErrorNumber:1393, ErrorType:"ServerError", Symbol:"ER_VIEW_MULTIUPDATE", SQLState:"HY000", Message:"Can not modify more than one base table through a joinview '%s.%s'", Description:"", MySQLVersion:"5.6"},
    1394: {ErrorNumber:1394, ErrorType:"ServerError", Symbol:"ER_VIEW_NO_INSERT_FIELD_LIST", SQLState:"HY000", Message:"Can not insert into join view '%s.%s' without fields list", Description:"", MySQLVersion:"5.6"},
    1395: {ErrorNumber:1395, ErrorType:"ServerError", Symbol:"ER_VIEW_DELETE_MERGE_VIEW", SQLState:"HY000", Message:"Can not delete from join view '%s.%s'", Description:"", MySQLVersion:"5.6"},
    1396: {ErrorNumber:1396, ErrorType:"ServerError", Symbol:"ER_CANNOT_USER", SQLState:"HY000", Message:"Operation %s failed for %s", Description:"", MySQLVersion:"5.6"},
    1397: {ErrorNumber:1397, ErrorType:"ServerError", Symbol:"ER_XAER_NOTA", SQLState:"XAE04", Message:"XAER_NOTA: Unknown XID", Description:"", MySQLVersion:"5.6"},
    1398: {ErrorNumber:1398, ErrorType:"ServerError", Symbol:"ER_XAER_INVAL", SQLState:"XAE05", Message:"XAER_INVAL: Invalid arguments (or unsupported command)", Description:"", MySQLVersion:"5.6"},
    1399: {ErrorNumber:1399, ErrorType:"ServerError", Symbol:"ER_XAER_RMFAIL", SQLState:"XAE07", Message:"XAER_RMFAIL: The command cannot be executed when globaltransaction is in the %s state", Description:"", MySQLVersion:"5.6"},
    1400: {ErrorNumber:1400, ErrorType:"ServerError", Symbol:"ER_XAER_OUTSIDE", SQLState:"XAE09", Message:"XAER_OUTSIDE: Some work is done outside globaltransaction", Description:"", MySQLVersion:"5.6"},
    1401: {ErrorNumber:1401, ErrorType:"ServerError", Symbol:"ER_XAER_RMERR", SQLState:"XAE03", Message:"XAER_RMERR: Fatal error occurred in the transactionbranch - check your data for consistency", Description:"", MySQLVersion:"5.6"},
    1402: {ErrorNumber:1402, ErrorType:"ServerError", Symbol:"ER_XA_RBROLLBACK", SQLState:"XA100", Message:"XA_RBROLLBACK: Transaction branch was rolled back", Description:"", MySQLVersion:"5.6"},
    1403: {ErrorNumber:1403, ErrorType:"ServerError", Symbol:"ER_NONEXISTING_PROC_GRANT", SQLState:"42000", Message:"There is no such grant defined for user '%s' on host '%s'on routine '%s'", Description:"", MySQLVersion:"5.6"},
    1404: {ErrorNumber:1404, ErrorType:"ServerError", Symbol:"ER_PROC_AUTO_GRANT_FAIL", SQLState:"HY000", Message:"Failed to grant EXECUTE and ALTER ROUTINE privileges", Description:"", MySQLVersion:"5.6"},
    1405: {ErrorNumber:1405, ErrorType:"ServerError", Symbol:"ER_PROC_AUTO_REVOKE_FAIL", SQLState:"HY000", Message:"Failed to revoke all privileges to dropped routine", Description:"", MySQLVersion:"5.6"},
    1406: {ErrorNumber:1406, ErrorType:"ServerError", Symbol:"ER_DATA_TOO_LONG", SQLState:"22001", Message:"Data too long for column '%s' at row %ld", Description:"", MySQLVersion:"5.6"},
    1407: {ErrorNumber:1407, ErrorType:"ServerError", Symbol:"ER_SP_BAD_SQLSTATE", SQLState:"42000", Message:"Bad SQLSTATE: '%s'", Description:"", MySQLVersion:"5.6"},
    1408: {ErrorNumber:1408, ErrorType:"ServerError", Symbol:"ER_STARTUP", SQLState:"HY000", Message:"%s: ready for connections. Version: '%s' socket: '%s'port: %d %s", Description:"", MySQLVersion:"5.6"},
    1409: {ErrorNumber:1409, ErrorType:"ServerError", Symbol:"ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR", SQLState:"HY000", Message:"Can't load value from file with fixed size rows tovariable", Description:"", MySQLVersion:"5.6"},
    1410: {ErrorNumber:1410, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_USER_WITH_GRANT", SQLState:"42000", Message:"You are not allowed to create a user with GRANT", Description:"", MySQLVersion:"5.6"},
    1411: {ErrorNumber:1411, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE_FOR_TYPE", SQLState:"HY000", Message:"Incorrect %s value: '%s' for function %s", Description:"", MySQLVersion:"5.6"},
    1412: {ErrorNumber:1412, ErrorType:"ServerError", Symbol:"ER_TABLE_DEF_CHANGED", SQLState:"HY000", Message:"Table definition has changed, please retry transaction", Description:"", MySQLVersion:"5.6"},
    1413: {ErrorNumber:1413, ErrorType:"ServerError", Symbol:"ER_SP_DUP_HANDLER", SQLState:"42000", Message:"Duplicate handler declared in the same block", Description:"", MySQLVersion:"5.6"},
    1414: {ErrorNumber:1414, ErrorType:"ServerError", Symbol:"ER_SP_NOT_VAR_ARG", SQLState:"42000", Message:"OUT or INOUT argument %d for routine %s is not a variableor NEW pseudo-variable in BEFORE trigger", Description:"", MySQLVersion:"5.6"},
    1415: {ErrorNumber:1415, ErrorType:"ServerError", Symbol:"ER_SP_NO_RETSET", SQLState:"0A000", Message:"Not allowed to return a result set from a %s", Description:"", MySQLVersion:"5.6"},
    1416: {ErrorNumber:1416, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_GEOMETRY_OBJECT", SQLState:"22003", Message:"Cannot get geometry object from data you send to theGEOMETRY field", Description:"", MySQLVersion:"5.6"},
    1417: {ErrorNumber:1417, ErrorType:"ServerError", Symbol:"ER_FAILED_ROUTINE_BREAK_BINLOG", SQLState:"HY000", Message:"A routine failed and has neither NO SQL nor READS SQLDATA in its declaration and binary logging is enabled", Description:"ifnon-transactional tables were updated, the binary log will misstheir changes", MySQLVersion:"5.6"},
    1418: {ErrorNumber:1418, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_ROUTINE", SQLState:"HY000", Message:"This function has none of DETERMINISTIC, NO SQL, or READSSQL DATA in its declaration and binary logging is enabled (you*might* want to use the less safe log_bin_trust_function_creatorsvariable)", Description:"", MySQLVersion:"5.6"},
    1419: {ErrorNumber:1419, ErrorType:"ServerError", Symbol:"ER_BINLOG_CREATE_ROUTINE_NEED_SUPER", SQLState:"HY000", Message:"You do not have the SUPER privilege and binary logging isenabled (you *might* want to use the less safelog_bin_trust_function_creators variable)", Description:"", MySQLVersion:"5.6"},
    1420: {ErrorNumber:1420, ErrorType:"ServerError", Symbol:"ER_EXEC_STMT_WITH_OPEN_CURSOR", SQLState:"HY000", Message:"You can't execute a prepared statement which has an opencursor associated with it. Reset the statement to re-execute it.", Description:"", MySQLVersion:"5.6"},
    1421: {ErrorNumber:1421, ErrorType:"ServerError", Symbol:"ER_STMT_HAS_NO_OPEN_CURSOR", SQLState:"HY000", Message:"The statement (%lu) has no open cursor.", Description:"", MySQLVersion:"5.6"},
    1422: {ErrorNumber:1422, ErrorType:"ServerError", Symbol:"ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG", SQLState:"HY000", Message:"Explicit or implicit commit is not allowed in storedfunction or trigger.", Description:"", MySQLVersion:"5.6"},
    1423: {ErrorNumber:1423, ErrorType:"ServerError", Symbol:"ER_NO_DEFAULT_FOR_VIEW_FIELD", SQLState:"HY000", Message:"Field of view '%s.%s' underlying table doesn't have adefault value", Description:"", MySQLVersion:"5.6"},
    1424: {ErrorNumber:1424, ErrorType:"ServerError", Symbol:"ER_SP_NO_RECURSION", SQLState:"HY000", Message:"Recursive stored functions and triggers are not allowed.", Description:"", MySQLVersion:"5.6"},
    1425: {ErrorNumber:1425, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_SCALE", SQLState:"42000", Message:"Too big scale %d specified for column '%s'. Maximum is%lu.", Description:"", MySQLVersion:"5.6"},
    1426: {ErrorNumber:1426, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_PRECISION", SQLState:"42000", Message:"Too big precision %d specified for column '%s'. Maximumis %lu.", Description:"", MySQLVersion:"5.6"},
    1427: {ErrorNumber:1427, ErrorType:"ServerError", Symbol:"ER_M_BIGGER_THAN_D", SQLState:"42000", Message:"For float(M,D), double(M,D) or decimal(M,D), M must be>= D (column '%s').", Description:"", MySQLVersion:"5.6"},
    1428: {ErrorNumber:1428, ErrorType:"ServerError", Symbol:"ER_WRONG_LOCK_OF_SYSTEM_TABLE", SQLState:"HY000", Message:"You can't combine write-locking of system tables withother tables or lock types", Description:"", MySQLVersion:"5.6"},
    1429: {ErrorNumber:1429, ErrorType:"ServerError", Symbol:"ER_CONNECT_TO_FOREIGN_DATA_SOURCE", SQLState:"HY000", Message:"Unable to connect to foreign data source: %s", Description:"", MySQLVersion:"5.6"},
    1430: {ErrorNumber:1430, ErrorType:"ServerError", Symbol:"ER_QUERY_ON_FOREIGN_DATA_SOURCE", SQLState:"HY000", Message:"There was a problem processing the query on the foreigndata source. Data source error: %s", Description:"", MySQLVersion:"5.6"},
    1431: {ErrorNumber:1431, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST", SQLState:"HY000", Message:"The foreign data source you are trying to reference doesnot exist. Data source error: %s", Description:"", MySQLVersion:"5.6"},
    1432: {ErrorNumber:1432, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE", SQLState:"HY000", Message:"Can't create federated table. The data source connectionstring '%s' is not in the correct format", Description:"", MySQLVersion:"5.6"},
    1433: {ErrorNumber:1433, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DATA_STRING_INVALID", SQLState:"HY000", Message:"The data source connection string '%s' is not in thecorrect format", Description:"", MySQLVersion:"5.6"},
    1434: {ErrorNumber:1434, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_FEDERATED_TABLE", SQLState:"HY000", Message:"Can't create federated table. Foreign data src error: %s", Description:"", MySQLVersion:"5.6"},
    1435: {ErrorNumber:1435, ErrorType:"ServerError", Symbol:"ER_TRG_IN_WRONG_SCHEMA", SQLState:"HY000", Message:"Trigger in wrong schema", Description:"", MySQLVersion:"5.6"},
    1436: {ErrorNumber:1436, ErrorType:"ServerError", Symbol:"ER_STACK_OVERRUN_NEED_MORE", SQLState:"HY000", Message:"Thread stack overrun: %ld bytes used of a %ld byte stack,and %ld bytes needed. Use 'mysqld --thread_stack=#' to specify abigger stack.", Description:"", MySQLVersion:"5.6"},
    1437: {ErrorNumber:1437, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_BODY", SQLState:"42000", Message:"Routine body for '%s' is too long", Description:"", MySQLVersion:"5.6"},
    1438: {ErrorNumber:1438, ErrorType:"ServerError", Symbol:"ER_WARN_CANT_DROP_DEFAULT_KEYCACHE", SQLState:"HY000", Message:"Cannot drop default keycache", Description:"", MySQLVersion:"5.6"},
    1439: {ErrorNumber:1439, ErrorType:"ServerError", Symbol:"ER_TOO_BIG_DISPLAYWIDTH", SQLState:"42000", Message:"Display width out of range for column '%s' (max = %lu)", Description:"", MySQLVersion:"5.6"},
    1440: {ErrorNumber:1440, ErrorType:"ServerError", Symbol:"ER_XAER_DUPID", SQLState:"XAE08", Message:"XAER_DUPID: The XID already exists", Description:"", MySQLVersion:"5.6"},
    1441: {ErrorNumber:1441, ErrorType:"ServerError", Symbol:"ER_DATETIME_FUNCTION_OVERFLOW", SQLState:"22008", Message:"Datetime function: %s field overflow", Description:"", MySQLVersion:"5.6"},
    1442: {ErrorNumber:1442, ErrorType:"ServerError", Symbol:"ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG", SQLState:"HY000", Message:"Can't update table '%s' in stored function/triggerbecause it is already used by statement which invoked this storedfunction/trigger.", Description:"", MySQLVersion:"5.6"},
    1443: {ErrorNumber:1443, ErrorType:"ServerError", Symbol:"ER_VIEW_PREVENT_UPDATE", SQLState:"HY000", Message:"The definition of table '%s' prevents operation %s ontable '%s'.", Description:"", MySQLVersion:"5.6"},
    1444: {ErrorNumber:1444, ErrorType:"ServerError", Symbol:"ER_PS_NO_RECURSION", SQLState:"HY000", Message:"The prepared statement contains a stored routine callthat refers to that same statement. It's not allowed to execute aprepared statement in such a recursive manner", Description:"", MySQLVersion:"5.6"},
    1445: {ErrorNumber:1445, ErrorType:"ServerError", Symbol:"ER_SP_CANT_SET_AUTOCOMMIT", SQLState:"HY000", Message:"Not allowed to set autocommit from a stored function ortrigger", Description:"", MySQLVersion:"5.6"},
    1446: {ErrorNumber:1446, ErrorType:"ServerError", Symbol:"ER_MALFORMED_DEFINER", SQLState:"HY000", Message:"Definer is not fully qualified", Description:"", MySQLVersion:"5.6"},
    1447: {ErrorNumber:1447, ErrorType:"ServerError", Symbol:"ER_VIEW_FRM_NO_USER", SQLState:"HY000", Message:"View '%s'.'%s' has no definer information (old tableformat). Current user is used as definer. Please recreate theview!", Description:"", MySQLVersion:"5.6"},
    1448: {ErrorNumber:1448, ErrorType:"ServerError", Symbol:"ER_VIEW_OTHER_USER", SQLState:"HY000", Message:"You need the SUPER privilege for creation view with'%s'@'%s' definer", Description:"", MySQLVersion:"5.6"},
    1449: {ErrorNumber:1449, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_USER", SQLState:"HY000", Message:"The user specified as a definer ('%s'@'%s') does notexist", Description:"", MySQLVersion:"5.6"},
    1450: {ErrorNumber:1450, ErrorType:"ServerError", Symbol:"ER_FORBID_SCHEMA_CHANGE", SQLState:"HY000", Message:"Changing schema from '%s' to '%s' is not allowed.", Description:"", MySQLVersion:"5.6"},
    1451: {ErrorNumber:1451, ErrorType:"ServerError", Symbol:"ER_ROW_IS_REFERENCED_2", SQLState:"23000", Message:"Cannot delete or update a parent row: a foreign keyconstraint fails (%s)", Description:"InnoDB reports this error when you try todelete a parent row that has children, and aforeign keyconstraint fails. Delete the children first.", MySQLVersion:"5.6"},
    1452: {ErrorNumber:1452, ErrorType:"ServerError", Symbol:"ER_NO_REFERENCED_ROW_2", SQLState:"23000", Message:"Cannot add or update a child row: a foreign keyconstraint fails (%s)", Description:"InnoDB reports this error when you try to add arow but there is no parent row, and aforeign keyconstraint fails. Add the parent row first.", MySQLVersion:"5.6"},
    1453: {ErrorNumber:1453, ErrorType:"ServerError", Symbol:"ER_SP_BAD_VAR_SHADOW", SQLState:"42000", Message:"Variable '%s' must be quoted with `...`, or renamed", Description:"", MySQLVersion:"5.6"},
    1454: {ErrorNumber:1454, ErrorType:"ServerError", Symbol:"ER_TRG_NO_DEFINER", SQLState:"HY000", Message:"No definer attribute for trigger '%s'.'%s'. The triggerwill be activated under the authorization of the invoker, whichmay have insufficient privileges. Please recreate the trigger.", Description:"", MySQLVersion:"5.6"},
    1455: {ErrorNumber:1455, ErrorType:"ServerError", Symbol:"ER_OLD_FILE_FORMAT", SQLState:"HY000", Message:"'%s' has an old format, you should re-create the '%s'object(s)", Description:"", MySQLVersion:"5.6"},
    1456: {ErrorNumber:1456, ErrorType:"ServerError", Symbol:"ER_SP_RECURSION_LIMIT", SQLState:"HY000", Message:"Recursive limit %d (as set by the max_sp_recursion_depthvariable) was exceeded for routine %s", Description:"", MySQLVersion:"5.6"},
    1457: {ErrorNumber:1457, ErrorType:"ServerError", Symbol:"ER_SP_PROC_TABLE_CORRUPT", SQLState:"HY000", Message:"Failed to load routine %s. The table mysql.proc ismissing, corrupt, or contains bad data (internal code %d)", Description:"", MySQLVersion:"5.6"},
    1458: {ErrorNumber:1458, ErrorType:"ServerError", Symbol:"ER_SP_WRONG_NAME", SQLState:"42000", Message:"Incorrect routine name '%s'", Description:"", MySQLVersion:"5.6"},
    1459: {ErrorNumber:1459, ErrorType:"ServerError", Symbol:"ER_TABLE_NEEDS_UPGRADE", SQLState:"HY000", Message:"Table upgrade required. Please do \"REPAIR TABLE `%s`\" ordump/reload to fix it!", Description:"", MySQLVersion:"5.6"},
    1460: {ErrorNumber:1460, ErrorType:"ServerError", Symbol:"ER_SP_NO_AGGREGATE", SQLState:"42000", Message:"AGGREGATE is not supported for stored functions", Description:"", MySQLVersion:"5.6"},
    1461: {ErrorNumber:1461, ErrorType:"ServerError", Symbol:"ER_MAX_PREPARED_STMT_COUNT_REACHED", SQLState:"42000", Message:"Can't create more than max_prepared_stmt_count statements(current value: %lu)", Description:"", MySQLVersion:"5.6"},
    1462: {ErrorNumber:1462, ErrorType:"ServerError", Symbol:"ER_VIEW_RECURSIVE", SQLState:"HY000", Message:"`%s`.`%s` contains view recursion", Description:"", MySQLVersion:"5.6"},
    1463: {ErrorNumber:1463, ErrorType:"ServerError", Symbol:"ER_NON_GROUPING_FIELD_USED", SQLState:"42000", Message:"Non-grouping field '%s' is used in %s clause", Description:"", MySQLVersion:"5.6"},
    1464: {ErrorNumber:1464, ErrorType:"ServerError", Symbol:"ER_TABLE_CANT_HANDLE_SPKEYS", SQLState:"HY000", Message:"The used table type doesn't support SPATIAL indexes", Description:"", MySQLVersion:"5.6"},
    1465: {ErrorNumber:1465, ErrorType:"ServerError", Symbol:"ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA", SQLState:"HY000", Message:"Triggers can not be created on system tables", Description:"", MySQLVersion:"5.6"},
    1466: {ErrorNumber:1466, ErrorType:"ServerError", Symbol:"ER_REMOVED_SPACES", SQLState:"HY000", Message:"Leading spaces are removed from name '%s'", Description:"", MySQLVersion:"5.6"},
    1467: {ErrorNumber:1467, ErrorType:"ServerError", Symbol:"ER_AUTOINC_READ_FAILED", SQLState:"HY000", Message:"Failed to read auto-increment value from storage engine", Description:"", MySQLVersion:"5.6"},
    1468: {ErrorNumber:1468, ErrorType:"ServerError", Symbol:"ER_USERNAME", SQLState:"HY000", Message:"user name", Description:"", MySQLVersion:"5.6"},
    1469: {ErrorNumber:1469, ErrorType:"ServerError", Symbol:"ER_HOSTNAME", SQLState:"HY000", Message:"host name", Description:"", MySQLVersion:"5.6"},
    1470: {ErrorNumber:1470, ErrorType:"ServerError", Symbol:"ER_WRONG_STRING_LENGTH", SQLState:"HY000", Message:"String '%s' is too long for %s (should be no longer than%d)", Description:"", MySQLVersion:"5.6"},
    1471: {ErrorNumber:1471, ErrorType:"ServerError", Symbol:"ER_NON_INSERTABLE_TABLE", SQLState:"HY000", Message:"The target table %s of the %s is not insertable-into", Description:"", MySQLVersion:"5.6"},
    1472: {ErrorNumber:1472, ErrorType:"ServerError", Symbol:"ER_ADMIN_WRONG_MRG_TABLE", SQLState:"HY000", Message:"Table '%s' is differently defined or of non-MyISAM typeor doesn't exist", Description:"", MySQLVersion:"5.6"},
    1473: {ErrorNumber:1473, ErrorType:"ServerError", Symbol:"ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT", SQLState:"HY000", Message:"Too high level of nesting for select", Description:"", MySQLVersion:"5.6"},
    1474: {ErrorNumber:1474, ErrorType:"ServerError", Symbol:"ER_NAME_BECOMES_EMPTY", SQLState:"HY000", Message:"Name '%s' has become ''", Description:"", MySQLVersion:"5.6"},
    1475: {ErrorNumber:1475, ErrorType:"ServerError", Symbol:"ER_AMBIGUOUS_FIELD_TERM", SQLState:"HY000", Message:"First character of the FIELDS TERMINATED string isambiguous", Description:"please use non-optional and non-empty FIELDS ENCLOSEDBY", MySQLVersion:"5.6"},
    1476: {ErrorNumber:1476, ErrorType:"ServerError", Symbol:"ER_FOREIGN_SERVER_EXISTS", SQLState:"HY000", Message:"The foreign server, %s, you are trying to create alreadyexists.", Description:"", MySQLVersion:"5.6"},
    1477: {ErrorNumber:1477, ErrorType:"ServerError", Symbol:"ER_FOREIGN_SERVER_DOESNT_EXIST", SQLState:"HY000", Message:"The foreign server name you are trying to reference doesnot exist. Data source error: %s", Description:"", MySQLVersion:"5.6"},
    1478: {ErrorNumber:1478, ErrorType:"ServerError", Symbol:"ER_ILLEGAL_HA_CREATE_OPTION", SQLState:"HY000", Message:"Table storage engine '%s' does not support the createoption '%s'", Description:"", MySQLVersion:"5.6"},
    1479: {ErrorNumber:1479, ErrorType:"ServerError", Symbol:"ER_PARTITION_REQUIRES_VALUES_ERROR", SQLState:"HY000", Message:"Syntax error: %s PARTITIONING requires definition ofVALUES %s for each partition", Description:"", MySQLVersion:"5.6"},
    1480: {ErrorNumber:1480, ErrorType:"ServerError", Symbol:"ER_PARTITION_WRONG_VALUES_ERROR", SQLState:"HY000", Message:"Only %s PARTITIONING can use VALUES %s in partitiondefinition", Description:"", MySQLVersion:"5.6"},
    1481: {ErrorNumber:1481, ErrorType:"ServerError", Symbol:"ER_PARTITION_MAXVALUE_ERROR", SQLState:"HY000", Message:"MAXVALUE can only be used in last partition definition", Description:"", MySQLVersion:"5.6"},
    1482: {ErrorNumber:1482, ErrorType:"ServerError", Symbol:"ER_PARTITION_SUBPARTITION_ERROR", SQLState:"HY000", Message:"Subpartitions can only be hash partitions and by key", Description:"", MySQLVersion:"5.6"},
    1483: {ErrorNumber:1483, ErrorType:"ServerError", Symbol:"ER_PARTITION_SUBPART_MIX_ERROR", SQLState:"HY000", Message:"Must define subpartitions on all partitions if on onepartition", Description:"", MySQLVersion:"5.6"},
    1484: {ErrorNumber:1484, ErrorType:"ServerError", Symbol:"ER_PARTITION_WRONG_NO_PART_ERROR", SQLState:"HY000", Message:"Wrong number of partitions defined, mismatch withprevious setting", Description:"", MySQLVersion:"5.6"},
    1485: {ErrorNumber:1485, ErrorType:"ServerError", Symbol:"ER_PARTITION_WRONG_NO_SUBPART_ERROR", SQLState:"HY000", Message:"Wrong number of subpartitions defined, mismatch withprevious setting", Description:"", MySQLVersion:"5.6"},
    1486: {ErrorNumber:1486, ErrorType:"ServerError", Symbol:"ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR", SQLState:"HY000", Message:"Constant, random or timezone-dependent expressions in(sub)partitioning function are not allowed", Description:"", MySQLVersion:"5.6"},
    1487: {ErrorNumber:1487, ErrorType:"ServerError", Symbol:"ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR", SQLState:"HY000", Message:"Expression in RANGE/LIST VALUES must be constant", Description:"", MySQLVersion:"5.6"},
    1488: {ErrorNumber:1488, ErrorType:"ServerError", Symbol:"ER_FIELD_NOT_FOUND_PART_ERROR", SQLState:"HY000", Message:"Field in list of fields for partition function not foundin table", Description:"", MySQLVersion:"5.6"},
    1489: {ErrorNumber:1489, ErrorType:"ServerError", Symbol:"ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR", SQLState:"HY000", Message:"List of fields is only allowed in KEY partitions", Description:"", MySQLVersion:"5.6"},
    1490: {ErrorNumber:1490, ErrorType:"ServerError", Symbol:"ER_INCONSISTENT_PARTITION_INFO_ERROR", SQLState:"HY000", Message:"The partition info in the frm file is not consistent withwhat can be written into the frm file", Description:"", MySQLVersion:"5.6"},
    1491: {ErrorNumber:1491, ErrorType:"ServerError", Symbol:"ER_PARTITION_FUNC_NOT_ALLOWED_ERROR", SQLState:"HY000", Message:"The %s function returns the wrong type", Description:"", MySQLVersion:"5.6"},
    1492: {ErrorNumber:1492, ErrorType:"ServerError", Symbol:"ER_PARTITIONS_MUST_BE_DEFINED_ERROR", SQLState:"HY000", Message:"For %s partitions each partition must be defined", Description:"", MySQLVersion:"5.6"},
    1493: {ErrorNumber:1493, ErrorType:"ServerError", Symbol:"ER_RANGE_NOT_INCREASING_ERROR", SQLState:"HY000", Message:"VALUES LESS THAN value must be strictly increasing foreach partition", Description:"", MySQLVersion:"5.6"},
    1494: {ErrorNumber:1494, ErrorType:"ServerError", Symbol:"ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR", SQLState:"HY000", Message:"VALUES value must be of same type as partition function", Description:"", MySQLVersion:"5.6"},
    1495: {ErrorNumber:1495, ErrorType:"ServerError", Symbol:"ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR", SQLState:"HY000", Message:"Multiple definition of same constant in list partitioning", Description:"", MySQLVersion:"5.6"},
    1496: {ErrorNumber:1496, ErrorType:"ServerError", Symbol:"ER_PARTITION_ENTRY_ERROR", SQLState:"HY000", Message:"Partitioning can not be used stand-alone in query", Description:"", MySQLVersion:"5.6"},
    1497: {ErrorNumber:1497, ErrorType:"ServerError", Symbol:"ER_MIX_HANDLER_ERROR", SQLState:"HY000", Message:"The mix of handlers in the partitions is not allowed inthis version of MySQL", Description:"", MySQLVersion:"5.6"},
    1498: {ErrorNumber:1498, ErrorType:"ServerError", Symbol:"ER_PARTITION_NOT_DEFINED_ERROR", SQLState:"HY000", Message:"For the partitioned engine it is necessary to define all%s", Description:"", MySQLVersion:"5.6"},
    1499: {ErrorNumber:1499, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_PARTITIONS_ERROR", SQLState:"HY000", Message:"Too many partitions (including subpartitions) weredefined", Description:"", MySQLVersion:"5.6"},
    1500: {ErrorNumber:1500, ErrorType:"ServerError", Symbol:"ER_SUBPARTITION_ERROR", SQLState:"HY000", Message:"It is only possible to mix RANGE/LIST partitioning withHASH/KEY partitioning for subpartitioning", Description:"", MySQLVersion:"5.6"},
    1501: {ErrorNumber:1501, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_HANDLER_FILE", SQLState:"HY000", Message:"Failed to create specific handler file", Description:"", MySQLVersion:"5.6"},
    1502: {ErrorNumber:1502, ErrorType:"ServerError", Symbol:"ER_BLOB_FIELD_IN_PART_FUNC_ERROR", SQLState:"HY000", Message:"A BLOB field is not allowed in partition function", Description:"", MySQLVersion:"5.6"},
    1503: {ErrorNumber:1503, ErrorType:"ServerError", Symbol:"ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF", SQLState:"HY000", Message:"A %s must include all columns in the table's partitioningfunction", Description:"", MySQLVersion:"5.6"},
    1504: {ErrorNumber:1504, ErrorType:"ServerError", Symbol:"ER_NO_PARTS_ERROR", SQLState:"HY000", Message:"Number of %s = 0 is not an allowed value", Description:"", MySQLVersion:"5.6"},
    1505: {ErrorNumber:1505, ErrorType:"ServerError", Symbol:"ER_PARTITION_MGMT_ON_NONPARTITIONED", SQLState:"HY000", Message:"Partition management on a not partitioned table is notpossible", Description:"", MySQLVersion:"5.6"},
    1506: {ErrorNumber:1506, ErrorType:"ServerError", Symbol:"ER_FOREIGN_KEY_ON_PARTITIONED", SQLState:"HY000", Message:"Foreign key clause is not yet supported in conjunctionwith partitioning", Description:"", MySQLVersion:"5.6"},
    1507: {ErrorNumber:1507, ErrorType:"ServerError", Symbol:"ER_DROP_PARTITION_NON_EXISTENT", SQLState:"HY000", Message:"Error in list of partitions to %s", Description:"", MySQLVersion:"5.6"},
    1508: {ErrorNumber:1508, ErrorType:"ServerError", Symbol:"ER_DROP_LAST_PARTITION", SQLState:"HY000", Message:"Cannot remove all partitions, use DROP TABLE instead", Description:"", MySQLVersion:"5.6"},
    1509: {ErrorNumber:1509, ErrorType:"ServerError", Symbol:"ER_COALESCE_ONLY_ON_HASH_PARTITION", SQLState:"HY000", Message:"COALESCE PARTITION can only be used on HASH/KEYpartitions", Description:"", MySQLVersion:"5.6"},
    1510: {ErrorNumber:1510, ErrorType:"ServerError", Symbol:"ER_REORG_HASH_ONLY_ON_SAME_NO", SQLState:"HY000", Message:"REORGANIZE PARTITION can only be used to reorganizepartitions not to change their numbers", Description:"", MySQLVersion:"5.6"},
    1511: {ErrorNumber:1511, ErrorType:"ServerError", Symbol:"ER_REORG_NO_PARAM_ERROR", SQLState:"HY000", Message:"REORGANIZE PARTITION without parameters can only be usedon auto-partitioned tables using HASH PARTITIONs", Description:"", MySQLVersion:"5.6"},
    1512: {ErrorNumber:1512, ErrorType:"ServerError", Symbol:"ER_ONLY_ON_RANGE_LIST_PARTITION", SQLState:"HY000", Message:"%s PARTITION can only be used on RANGE/LIST partitions", Description:"", MySQLVersion:"5.6"},
    1513: {ErrorNumber:1513, ErrorType:"ServerError", Symbol:"ER_ADD_PARTITION_SUBPART_ERROR", SQLState:"HY000", Message:"Trying to Add partition(s) with wrong number ofsubpartitions", Description:"", MySQLVersion:"5.6"},
    1514: {ErrorNumber:1514, ErrorType:"ServerError", Symbol:"ER_ADD_PARTITION_NO_NEW_PARTITION", SQLState:"HY000", Message:"At least one partition must be added", Description:"", MySQLVersion:"5.6"},
    1515: {ErrorNumber:1515, ErrorType:"ServerError", Symbol:"ER_COALESCE_PARTITION_NO_PARTITION", SQLState:"HY000", Message:"At least one partition must be coalesced", Description:"", MySQLVersion:"5.6"},
    1516: {ErrorNumber:1516, ErrorType:"ServerError", Symbol:"ER_REORG_PARTITION_NOT_EXIST", SQLState:"HY000", Message:"More partitions to reorganize than there are partitions", Description:"", MySQLVersion:"5.6"},
    1517: {ErrorNumber:1517, ErrorType:"ServerError", Symbol:"ER_SAME_NAME_PARTITION", SQLState:"HY000", Message:"Duplicate partition name %s", Description:"", MySQLVersion:"5.6"},
    1518: {ErrorNumber:1518, ErrorType:"ServerError", Symbol:"ER_NO_BINLOG_ERROR", SQLState:"HY000", Message:"It is not allowed to shut off binlog on this command", Description:"", MySQLVersion:"5.6"},
    1519: {ErrorNumber:1519, ErrorType:"ServerError", Symbol:"ER_CONSECUTIVE_REORG_PARTITIONS", SQLState:"HY000", Message:"When reorganizing a set of partitions they must be inconsecutive order", Description:"", MySQLVersion:"5.6"},
    1520: {ErrorNumber:1520, ErrorType:"ServerError", Symbol:"ER_REORG_OUTSIDE_RANGE", SQLState:"HY000", Message:"Reorganize of range partitions cannot change total rangesexcept for last partition where it can extend the range", Description:"", MySQLVersion:"5.6"},
    1521: {ErrorNumber:1521, ErrorType:"ServerError", Symbol:"ER_PARTITION_FUNCTION_FAILURE", SQLState:"HY000", Message:"Partition function not supported in this version for thishandler", Description:"", MySQLVersion:"5.6"},
    1522: {ErrorNumber:1522, ErrorType:"ServerError", Symbol:"ER_PART_STATE_ERROR", SQLState:"HY000", Message:"Partition state cannot be defined from CREATE/ALTER TABLE", Description:"", MySQLVersion:"5.6"},
    1523: {ErrorNumber:1523, ErrorType:"ServerError", Symbol:"ER_LIMITED_PART_RANGE", SQLState:"HY000", Message:"The %s handler only supports 32 bit integers in VALUES", Description:"", MySQLVersion:"5.6"},
    1524: {ErrorNumber:1524, ErrorType:"ServerError", Symbol:"ER_PLUGIN_IS_NOT_LOADED", SQLState:"HY000", Message:"Plugin '%s' is not loaded", Description:"", MySQLVersion:"5.6"},
    1525: {ErrorNumber:1525, ErrorType:"ServerError", Symbol:"ER_WRONG_VALUE", SQLState:"HY000", Message:"Incorrect %s value: '%s'", Description:"", MySQLVersion:"5.6"},
    1526: {ErrorNumber:1526, ErrorType:"ServerError", Symbol:"ER_NO_PARTITION_FOR_GIVEN_VALUE", SQLState:"HY000", Message:"Table has no partition for value %s", Description:"", MySQLVersion:"5.6"},
    1527: {ErrorNumber:1527, ErrorType:"ServerError", Symbol:"ER_FILEGROUP_OPTION_ONLY_ONCE", SQLState:"HY000", Message:"It is not allowed to specify %s more than once", Description:"", MySQLVersion:"5.6"},
    1528: {ErrorNumber:1528, ErrorType:"ServerError", Symbol:"ER_CREATE_FILEGROUP_FAILED", SQLState:"HY000", Message:"Failed to create %s", Description:"", MySQLVersion:"5.6"},
    1529: {ErrorNumber:1529, ErrorType:"ServerError", Symbol:"ER_DROP_FILEGROUP_FAILED", SQLState:"HY000", Message:"Failed to drop %s", Description:"", MySQLVersion:"5.6"},
    1530: {ErrorNumber:1530, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_AUTO_EXTEND_ERROR", SQLState:"HY000", Message:"The handler doesn't support autoextend of tablespaces", Description:"", MySQLVersion:"5.6"},
    1531: {ErrorNumber:1531, ErrorType:"ServerError", Symbol:"ER_WRONG_SIZE_NUMBER", SQLState:"HY000", Message:"A size parameter was incorrectly specified, either numberor on the form 10M", Description:"", MySQLVersion:"5.6"},
    1532: {ErrorNumber:1532, ErrorType:"ServerError", Symbol:"ER_SIZE_OVERFLOW_ERROR", SQLState:"HY000", Message:"The size number was correct but we don't allow the digitpart to be more than 2 billion", Description:"", MySQLVersion:"5.6"},
    1533: {ErrorNumber:1533, ErrorType:"ServerError", Symbol:"ER_ALTER_FILEGROUP_FAILED", SQLState:"HY000", Message:"Failed to alter: %s", Description:"", MySQLVersion:"5.6"},
    1534: {ErrorNumber:1534, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_LOGGING_FAILED", SQLState:"HY000", Message:"Writing one row to the row-based binary log failed", Description:"", MySQLVersion:"5.6"},
    1535: {ErrorNumber:1535, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_WRONG_TABLE_DEF", SQLState:"HY000", Message:"Table definition on master and slave does not match: %s", Description:"", MySQLVersion:"5.6"},
    1536: {ErrorNumber:1536, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_RBR_TO_SBR", SQLState:"HY000", Message:"Slave running with --log-slave-updates must use row-basedbinary logging to be able to replicate row-based binary log events", Description:"", MySQLVersion:"5.6"},
    1537: {ErrorNumber:1537, ErrorType:"ServerError", Symbol:"ER_EVENT_ALREADY_EXISTS", SQLState:"HY000", Message:"Event '%s' already exists", Description:"", MySQLVersion:"5.6"},
    1538: {ErrorNumber:1538, ErrorType:"ServerError", Symbol:"ER_EVENT_STORE_FAILED", SQLState:"HY000", Message:"Failed to store event %s. Error code %d from storageengine.", Description:"", MySQLVersion:"5.6"},
    1539: {ErrorNumber:1539, ErrorType:"ServerError", Symbol:"ER_EVENT_DOES_NOT_EXIST", SQLState:"HY000", Message:"Unknown event '%s'", Description:"", MySQLVersion:"5.6"},
    1540: {ErrorNumber:1540, ErrorType:"ServerError", Symbol:"ER_EVENT_CANT_ALTER", SQLState:"HY000", Message:"Failed to alter event '%s'", Description:"", MySQLVersion:"5.6"},
    1541: {ErrorNumber:1541, ErrorType:"ServerError", Symbol:"ER_EVENT_DROP_FAILED", SQLState:"HY000", Message:"Failed to drop %s", Description:"", MySQLVersion:"5.6"},
    1542: {ErrorNumber:1542, ErrorType:"ServerError", Symbol:"ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG", SQLState:"HY000", Message:"INTERVAL is either not positive or too big", Description:"", MySQLVersion:"5.6"},
    1543: {ErrorNumber:1543, ErrorType:"ServerError", Symbol:"ER_EVENT_ENDS_BEFORE_STARTS", SQLState:"HY000", Message:"ENDS is either invalid or before STARTS", Description:"", MySQLVersion:"5.6"},
    1544: {ErrorNumber:1544, ErrorType:"ServerError", Symbol:"ER_EVENT_EXEC_TIME_IN_THE_PAST", SQLState:"HY000", Message:"Event execution time is in the past. Event has beendisabled", Description:"", MySQLVersion:"5.6"},
    1545: {ErrorNumber:1545, ErrorType:"ServerError", Symbol:"ER_EVENT_OPEN_TABLE_FAILED", SQLState:"HY000", Message:"Failed to open mysql.event", Description:"", MySQLVersion:"5.6"},
    1546: {ErrorNumber:1546, ErrorType:"ServerError", Symbol:"ER_EVENT_NEITHER_M_EXPR_NOR_M_AT", SQLState:"HY000", Message:"No datetime expression provided", Description:"", MySQLVersion:"5.6"},
    1547: {ErrorNumber:1547, ErrorType:"ServerError", Symbol:"ER_OBSOLETE_COL_COUNT_DOESNT_MATCH_CORRUPTED", SQLState:"HY000", Message:"Column count of mysql.%s is wrong. Expected %d, found %d.The table is probably corrupted", Description:"", MySQLVersion:"5.6"},
    1548: {ErrorNumber:1548, ErrorType:"ServerError", Symbol:"ER_OBSOLETE_CANNOT_LOAD_FROM_TABLE", SQLState:"HY000", Message:"Cannot load from mysql.%s. The table is probablycorrupted", Description:"", MySQLVersion:"5.6"},
    1549: {ErrorNumber:1549, ErrorType:"ServerError", Symbol:"ER_EVENT_CANNOT_DELETE", SQLState:"HY000", Message:"Failed to delete the event from mysql.event", Description:"", MySQLVersion:"5.6"},
    1550: {ErrorNumber:1550, ErrorType:"ServerError", Symbol:"ER_EVENT_COMPILE_ERROR", SQLState:"HY000", Message:"Error during compilation of event's body", Description:"", MySQLVersion:"5.6"},
    1551: {ErrorNumber:1551, ErrorType:"ServerError", Symbol:"ER_EVENT_SAME_NAME", SQLState:"HY000", Message:"Same old and new event name", Description:"", MySQLVersion:"5.6"},
    1552: {ErrorNumber:1552, ErrorType:"ServerError", Symbol:"ER_EVENT_DATA_TOO_LONG", SQLState:"HY000", Message:"Data for column '%s' too long", Description:"", MySQLVersion:"5.6"},
    1553: {ErrorNumber:1553, ErrorType:"ServerError", Symbol:"ER_DROP_INDEX_FK", SQLState:"HY000", Message:"Cannot drop index '%s': needed in a foreign keyconstraint", Description:"When you drop an index, InnoDB checks if theindex is used for checking a foreign key constraint. It is stillOK to drop the index if there is another index that can be used toenforce the same constraint. InnoDB preventsyou from dropping the last index that can enforce a particularreferential constraint.", MySQLVersion:"5.6"},
    1554: {ErrorNumber:1554, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYNTAX_WITH_VER", SQLState:"HY000", Message:"The syntax '%s' is deprecated and will be removed inMySQL %s. Please use %s instead", Description:"", MySQLVersion:"5.6"},
    1555: {ErrorNumber:1555, ErrorType:"ServerError", Symbol:"ER_CANT_WRITE_LOCK_LOG_TABLE", SQLState:"HY000", Message:"You can't write-lock a log table. Only read access ispossible", Description:"", MySQLVersion:"5.6"},
    1556: {ErrorNumber:1556, ErrorType:"ServerError", Symbol:"ER_CANT_LOCK_LOG_TABLE", SQLState:"HY000", Message:"You can't use locks with log tables.", Description:"", MySQLVersion:"5.6"},
    1557: {ErrorNumber:1557, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED", SQLState:"23000", Message:"Upholding foreign key constraints for table '%s', entry'%s', key %d would lead to a duplicate entry", Description:"ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSEDwas added in 5.6.4.", MySQLVersion:"5.6"},
    1558: {ErrorNumber:1558, ErrorType:"ServerError", Symbol:"ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE", SQLState:"HY000", Message:"Column count of mysql.%s is wrong. Expected %d, found %d.Created with MySQL %d, now running %d. Please use mysql_upgrade tofix this error.", Description:"", MySQLVersion:"5.6"},
    1559: {ErrorNumber:1559, ErrorType:"ServerError", Symbol:"ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR", SQLState:"HY000", Message:"Cannot switch out of the row-based binary log format whenthe session has open temporary tables", Description:"", MySQLVersion:"5.6"},
    1560: {ErrorNumber:1560, ErrorType:"ServerError", Symbol:"ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT", SQLState:"HY000", Message:"Cannot change the binary logging format inside a storedfunction or trigger", Description:"", MySQLVersion:"5.6"},
    1561: {ErrorNumber:1561, ErrorType:"ServerError", Symbol:"ER_NDB_CANT_SWITCH_BINLOG_FORMAT", SQLState:"HY000", Message:"The NDB cluster engine does not support changing thebinlog format on the fly yet", Description:"", MySQLVersion:"5.6"},
    1562: {ErrorNumber:1562, ErrorType:"ServerError", Symbol:"ER_PARTITION_NO_TEMPORARY", SQLState:"HY000", Message:"Cannot create temporary table with partitions", Description:"", MySQLVersion:"5.6"},
    1563: {ErrorNumber:1563, ErrorType:"ServerError", Symbol:"ER_PARTITION_CONST_DOMAIN_ERROR", SQLState:"HY000", Message:"Partition constant is out of partition function domain", Description:"", MySQLVersion:"5.6"},
    1564: {ErrorNumber:1564, ErrorType:"ServerError", Symbol:"ER_PARTITION_FUNCTION_IS_NOT_ALLOWED", SQLState:"HY000", Message:"This partition function is not allowed", Description:"", MySQLVersion:"5.6"},
    1565: {ErrorNumber:1565, ErrorType:"ServerError", Symbol:"ER_DDL_LOG_ERROR", SQLState:"HY000", Message:"Error in DDL log", Description:"", MySQLVersion:"5.6"},
    1566: {ErrorNumber:1566, ErrorType:"ServerError", Symbol:"ER_NULL_IN_VALUES_LESS_THAN", SQLState:"HY000", Message:"Not allowed to use NULL value in VALUES LESS THAN", Description:"", MySQLVersion:"5.6"},
    1567: {ErrorNumber:1567, ErrorType:"ServerError", Symbol:"ER_WRONG_PARTITION_NAME", SQLState:"HY000", Message:"Incorrect partition name", Description:"", MySQLVersion:"5.6"},
    1568: {ErrorNumber:1568, ErrorType:"ServerError", Symbol:"ER_CANT_CHANGE_TX_CHARACTERISTICS", SQLState:"25001", Message:"Transaction characteristics can't be changed while atransaction is in progress", Description:"ER_CANT_CHANGE_TX_CHARACTERISTICSwas added in 5.6.5.", MySQLVersion:"5.6"},
    1569: {ErrorNumber:1569, ErrorType:"ServerError", Symbol:"ER_DUP_ENTRY_AUTOINCREMENT_CASE", SQLState:"HY000", Message:"ALTER TABLE causes auto_increment resequencing, resultingin duplicate entry '%s' for key '%s'", Description:"", MySQLVersion:"5.6"},
    1570: {ErrorNumber:1570, ErrorType:"ServerError", Symbol:"ER_EVENT_MODIFY_QUEUE_ERROR", SQLState:"HY000", Message:"Internal scheduler error %d", Description:"", MySQLVersion:"5.6"},
    1571: {ErrorNumber:1571, ErrorType:"ServerError", Symbol:"ER_EVENT_SET_VAR_ERROR", SQLState:"HY000", Message:"Error during starting/stopping of the scheduler. Errorcode %u", Description:"", MySQLVersion:"5.6"},
    1572: {ErrorNumber:1572, ErrorType:"ServerError", Symbol:"ER_PARTITION_MERGE_ERROR", SQLState:"HY000", Message:"Engine cannot be used in partitioned tables", Description:"", MySQLVersion:"5.6"},
    1573: {ErrorNumber:1573, ErrorType:"ServerError", Symbol:"ER_CANT_ACTIVATE_LOG", SQLState:"HY000", Message:"Cannot activate '%s' log", Description:"", MySQLVersion:"5.6"},
    1574: {ErrorNumber:1574, ErrorType:"ServerError", Symbol:"ER_RBR_NOT_AVAILABLE", SQLState:"HY000", Message:"The server was not built with row-based replication", Description:"", MySQLVersion:"5.6"},
    1575: {ErrorNumber:1575, ErrorType:"ServerError", Symbol:"ER_BASE64_DECODE_ERROR", SQLState:"HY000", Message:"Decoding of base64 string failed", Description:"", MySQLVersion:"5.6"},
    1576: {ErrorNumber:1576, ErrorType:"ServerError", Symbol:"ER_EVENT_RECURSION_FORBIDDEN", SQLState:"HY000", Message:"Recursion of EVENT DDL statements is forbidden when bodyis present", Description:"", MySQLVersion:"5.6"},
    1577: {ErrorNumber:1577, ErrorType:"ServerError", Symbol:"ER_EVENTS_DB_ERROR", SQLState:"HY000", Message:"Cannot proceed because system tables used by EventScheduler were found damaged at server start", Description:"To address this issue, try runningmysql_upgrade.", MySQLVersion:"5.6"},
    1578: {ErrorNumber:1578, ErrorType:"ServerError", Symbol:"ER_ONLY_INTEGERS_ALLOWED", SQLState:"HY000", Message:"Only integers allowed as number here", Description:"", MySQLVersion:"5.6"},
    1579: {ErrorNumber:1579, ErrorType:"ServerError", Symbol:"ER_UNSUPORTED_LOG_ENGINE", SQLState:"HY000", Message:"This storage engine cannot be used for log tables\"", Description:"", MySQLVersion:"5.6"},
    1580: {ErrorNumber:1580, ErrorType:"ServerError", Symbol:"ER_BAD_LOG_STATEMENT", SQLState:"HY000", Message:"You cannot '%s' a log table if logging is enabled", Description:"", MySQLVersion:"5.6"},
    1581: {ErrorNumber:1581, ErrorType:"ServerError", Symbol:"ER_CANT_RENAME_LOG_TABLE", SQLState:"HY000", Message:"Cannot rename '%s'. When logging enabled, rename to/fromlog table must rename two tables: the log table to an archivetable and another table back to '%s'", Description:"", MySQLVersion:"5.6"},
    1582: {ErrorNumber:1582, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT", SQLState:"42000", Message:"Incorrect parameter count in the call to native function'%s'", Description:"", MySQLVersion:"5.6"},
    1583: {ErrorNumber:1583, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMETERS_TO_NATIVE_FCT", SQLState:"42000", Message:"Incorrect parameters in the call to native function '%s'", Description:"", MySQLVersion:"5.6"},
    1584: {ErrorNumber:1584, ErrorType:"ServerError", Symbol:"ER_WRONG_PARAMETERS_TO_STORED_FCT", SQLState:"42000", Message:"Incorrect parameters in the call to stored function '%s'", Description:"", MySQLVersion:"5.6"},
    1585: {ErrorNumber:1585, ErrorType:"ServerError", Symbol:"ER_NATIVE_FCT_NAME_COLLISION", SQLState:"HY000", Message:"This function '%s' has the same name as a native function", Description:"", MySQLVersion:"5.6"},
    1586: {ErrorNumber:1586, ErrorType:"ServerError", Symbol:"ER_DUP_ENTRY_WITH_KEY_NAME", SQLState:"23000", Message:"Duplicate entry '%s' for key '%s'", Description:"The format string for this error is also used withER_DUP_ENTRY.", MySQLVersion:"5.6"},
    1587: {ErrorNumber:1587, ErrorType:"ServerError", Symbol:"ER_BINLOG_PURGE_EMFILE", SQLState:"HY000", Message:"Too many files opened, please execute the command again", Description:"", MySQLVersion:"5.6"},
    1588: {ErrorNumber:1588, ErrorType:"ServerError", Symbol:"ER_EVENT_CANNOT_CREATE_IN_THE_PAST", SQLState:"HY000", Message:"Event execution time is in the past and ON COMPLETION NOTPRESERVE is set. The event was dropped immediately after creation.", Description:"", MySQLVersion:"5.6"},
    1589: {ErrorNumber:1589, ErrorType:"ServerError", Symbol:"ER_EVENT_CANNOT_ALTER_IN_THE_PAST", SQLState:"HY000", Message:"Event execution time is in the past and ON COMPLETION NOTPRESERVE is set. The event was not changed. Specify a time in thefuture.", Description:"", MySQLVersion:"5.6"},
    1590: {ErrorNumber:1590, ErrorType:"ServerError", Symbol:"ER_SLAVE_INCIDENT", SQLState:"HY000", Message:"The incident %s occured on the master.Message: %s", Description:"", MySQLVersion:"5.6"},
    1591: {ErrorNumber:1591, ErrorType:"ServerError", Symbol:"ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT", SQLState:"HY000", Message:"Table has no partition for some existing values", Description:"", MySQLVersion:"5.6"},
    1592: {ErrorNumber:1592, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_STATEMENT", SQLState:"HY000", Message:"Unsafe statement written to the binary log usingstatement format since BINLOG_FORMAT = STATEMENT. %s", Description:"", MySQLVersion:"5.6"},
    1593: {ErrorNumber:1593, ErrorType:"ServerError", Symbol:"ER_SLAVE_FATAL_ERROR", SQLState:"HY000", Message:"Fatal error: %s", Description:"", MySQLVersion:"5.6"},
    1594: {ErrorNumber:1594, ErrorType:"ServerError", Symbol:"ER_SLAVE_RELAY_LOG_READ_FAILURE", SQLState:"HY000", Message:"Relay log read failure: %s", Description:"", MySQLVersion:"5.6"},
    1595: {ErrorNumber:1595, ErrorType:"ServerError", Symbol:"ER_SLAVE_RELAY_LOG_WRITE_FAILURE", SQLState:"HY000", Message:"Relay log write failure: %s", Description:"", MySQLVersion:"5.6"},
    1596: {ErrorNumber:1596, ErrorType:"ServerError", Symbol:"ER_SLAVE_CREATE_EVENT_FAILURE", SQLState:"HY000", Message:"Failed to create %s", Description:"", MySQLVersion:"5.6"},
    1597: {ErrorNumber:1597, ErrorType:"ServerError", Symbol:"ER_SLAVE_MASTER_COM_FAILURE", SQLState:"HY000", Message:"Master command %s failed: %s", Description:"", MySQLVersion:"5.6"},
    1598: {ErrorNumber:1598, ErrorType:"ServerError", Symbol:"ER_BINLOG_LOGGING_IMPOSSIBLE", SQLState:"HY000", Message:"Binary logging not possible.Message: %s", Description:"", MySQLVersion:"5.6"},
    1599: {ErrorNumber:1599, ErrorType:"ServerError", Symbol:"ER_VIEW_NO_CREATION_CTX", SQLState:"HY000", Message:"View `%s`.`%s` has no creation context", Description:"", MySQLVersion:"5.6"},
    1600: {ErrorNumber:1600, ErrorType:"ServerError", Symbol:"ER_VIEW_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Creation context of view `%s`.`%s' is invalid", Description:"", MySQLVersion:"5.6"},
    1601: {ErrorNumber:1601, ErrorType:"ServerError", Symbol:"ER_SR_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Creation context of stored routine `%s`.`%s` is invalid", Description:"", MySQLVersion:"5.6"},
    1602: {ErrorNumber:1602, ErrorType:"ServerError", Symbol:"ER_TRG_CORRUPTED_FILE", SQLState:"HY000", Message:"Corrupted TRG file for table `%s`.`%s`", Description:"", MySQLVersion:"5.6"},
    1603: {ErrorNumber:1603, ErrorType:"ServerError", Symbol:"ER_TRG_NO_CREATION_CTX", SQLState:"HY000", Message:"Triggers for table `%s`.`%s` have no creation context", Description:"", MySQLVersion:"5.6"},
    1604: {ErrorNumber:1604, ErrorType:"ServerError", Symbol:"ER_TRG_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Trigger creation context of table `%s`.`%s` is invalid", Description:"", MySQLVersion:"5.6"},
    1605: {ErrorNumber:1605, ErrorType:"ServerError", Symbol:"ER_EVENT_INVALID_CREATION_CTX", SQLState:"HY000", Message:"Creation context of event `%s`.`%s` is invalid", Description:"", MySQLVersion:"5.6"},
    1606: {ErrorNumber:1606, ErrorType:"ServerError", Symbol:"ER_TRG_CANT_OPEN_TABLE", SQLState:"HY000", Message:"Cannot open table for trigger `%s`.`%s`", Description:"", MySQLVersion:"5.6"},
    1607: {ErrorNumber:1607, ErrorType:"ServerError", Symbol:"ER_CANT_CREATE_SROUTINE", SQLState:"HY000", Message:"Cannot create stored routine `%s`. Check warnings", Description:"", MySQLVersion:"5.6"},
    1608: {ErrorNumber:1608, ErrorType:"ServerError", Symbol:"ER_NEVER_USED", SQLState:"HY000", Message:"Ambiguous slave modes combination. %s", Description:"", MySQLVersion:"5.6"},
    1609: {ErrorNumber:1609, ErrorType:"ServerError", Symbol:"ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT", SQLState:"HY000", Message:"The BINLOG statement of type `%s` was not preceded by aformat description BINLOG statement.", Description:"", MySQLVersion:"5.6"},
    1610: {ErrorNumber:1610, ErrorType:"ServerError", Symbol:"ER_SLAVE_CORRUPT_EVENT", SQLState:"HY000", Message:"Corrupted replication event was detected", Description:"", MySQLVersion:"5.6"},
    1611: {ErrorNumber:1611, ErrorType:"ServerError", Symbol:"ER_LOAD_DATA_INVALID_COLUMN", SQLState:"HY000", Message:"Invalid column reference (%s) in LOAD DATA", Description:"", MySQLVersion:"5.6"},
    1612: {ErrorNumber:1612, ErrorType:"ServerError", Symbol:"ER_LOG_PURGE_NO_FILE", SQLState:"HY000", Message:"Being purged log %s was not found", Description:"", MySQLVersion:"5.6"},
    1613: {ErrorNumber:1613, ErrorType:"ServerError", Symbol:"ER_XA_RBTIMEOUT", SQLState:"XA106", Message:"XA_RBTIMEOUT: Transaction branch was rolled back: tooktoo long", Description:"", MySQLVersion:"5.6"},
    1614: {ErrorNumber:1614, ErrorType:"ServerError", Symbol:"ER_XA_RBDEADLOCK", SQLState:"XA102", Message:"XA_RBDEADLOCK: Transaction branch was rolled back:deadlock was detected", Description:"", MySQLVersion:"5.6"},
    1615: {ErrorNumber:1615, ErrorType:"ServerError", Symbol:"ER_NEED_REPREPARE", SQLState:"HY000", Message:"Prepared statement needs to be re-prepared", Description:"", MySQLVersion:"5.6"},
    1616: {ErrorNumber:1616, ErrorType:"ServerError", Symbol:"ER_DELAYED_NOT_SUPPORTED", SQLState:"HY000", Message:"DELAYED option not supported for table '%s'", Description:"", MySQLVersion:"5.6"},
    1617: {ErrorNumber:1617, ErrorType:"ServerError", Symbol:"WARN_NO_MASTER_INFO", SQLState:"HY000", Message:"The master info structure does not exist", Description:"", MySQLVersion:"5.6"},
    1618: {ErrorNumber:1618, ErrorType:"ServerError", Symbol:"WARN_OPTION_IGNORED", SQLState:"HY000", Message:"<%s> option ignored", Description:"", MySQLVersion:"5.6"},
    1619: {ErrorNumber:1619, ErrorType:"ServerError", Symbol:"WARN_PLUGIN_DELETE_BUILTIN", SQLState:"HY000", Message:"Built-in plugins cannot be deleted", Description:"", MySQLVersion:"5.6"},
    1620: {ErrorNumber:1620, ErrorType:"ServerError", Symbol:"WARN_PLUGIN_BUSY", SQLState:"HY000", Message:"Plugin is busy and will be uninstalled on shutdown", Description:"", MySQLVersion:"5.6"},
    1621: {ErrorNumber:1621, ErrorType:"ServerError", Symbol:"ER_VARIABLE_IS_READONLY", SQLState:"HY000", Message:"%s variable '%s' is read-only. Use SET %s to assign thevalue", Description:"", MySQLVersion:"5.6"},
    1622: {ErrorNumber:1622, ErrorType:"ServerError", Symbol:"ER_WARN_ENGINE_TRANSACTION_ROLLBACK", SQLState:"HY000", Message:"Storage engine %s does not support rollback for thisstatement. Transaction rolled back and must be restarted", Description:"", MySQLVersion:"5.6"},
    1623: {ErrorNumber:1623, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_FAILURE", SQLState:"HY000", Message:"Unexpected master's heartbeat data: %s", Description:"", MySQLVersion:"5.6"},
    1624: {ErrorNumber:1624, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE", SQLState:"HY000", Message:"The requested value for the heartbeat period is eithernegative or exceeds the maximum allowed (%s seconds).", Description:"", MySQLVersion:"5.6"},
    1625: {ErrorNumber:1625, ErrorType:"ServerError", Symbol:"ER_NDB_REPLICATION_SCHEMA_ERROR", SQLState:"HY000", Message:"Bad schema for mysql.ndb_replication table.Message: %s", Description:"", MySQLVersion:"5.6"},
    1626: {ErrorNumber:1626, ErrorType:"ServerError", Symbol:"ER_CONFLICT_FN_PARSE_ERROR", SQLState:"HY000", Message:"Error in parsing conflict function.Message: %s", Description:"", MySQLVersion:"5.6"},
    1627: {ErrorNumber:1627, ErrorType:"ServerError", Symbol:"ER_EXCEPTIONS_WRITE_ERROR", SQLState:"HY000", Message:"Write to exceptions table failed.Message: %s\"", Description:"", MySQLVersion:"5.6"},
    1628: {ErrorNumber:1628, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_TABLE_COMMENT", SQLState:"HY000", Message:"Comment for table '%s' is too long (max = %lu)", Description:"", MySQLVersion:"5.6"},
    1629: {ErrorNumber:1629, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_FIELD_COMMENT", SQLState:"HY000", Message:"Comment for field '%s' is too long (max = %lu)", Description:"", MySQLVersion:"5.6"},
    1630: {ErrorNumber:1630, ErrorType:"ServerError", Symbol:"ER_FUNC_INEXISTENT_NAME_COLLISION", SQLState:"42000", Message:"FUNCTION %s does not exist. Check the 'Function NameParsing and Resolution' section in the Reference Manual", Description:"", MySQLVersion:"5.6"},
    1631: {ErrorNumber:1631, ErrorType:"ServerError", Symbol:"ER_DATABASE_NAME", SQLState:"HY000", Message:"Database", Description:"", MySQLVersion:"5.6"},
    1632: {ErrorNumber:1632, ErrorType:"ServerError", Symbol:"ER_TABLE_NAME", SQLState:"HY000", Message:"Table", Description:"", MySQLVersion:"5.6"},
    1633: {ErrorNumber:1633, ErrorType:"ServerError", Symbol:"ER_PARTITION_NAME", SQLState:"HY000", Message:"Partition", Description:"", MySQLVersion:"5.6"},
    1634: {ErrorNumber:1634, ErrorType:"ServerError", Symbol:"ER_SUBPARTITION_NAME", SQLState:"HY000", Message:"Subpartition", Description:"", MySQLVersion:"5.6"},
    1635: {ErrorNumber:1635, ErrorType:"ServerError", Symbol:"ER_TEMPORARY_NAME", SQLState:"HY000", Message:"Temporary", Description:"", MySQLVersion:"5.6"},
    1636: {ErrorNumber:1636, ErrorType:"ServerError", Symbol:"ER_RENAMED_NAME", SQLState:"HY000", Message:"Renamed", Description:"", MySQLVersion:"5.6"},
    1637: {ErrorNumber:1637, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_CONCURRENT_TRXS", SQLState:"HY000", Message:"Too many active concurrent transactions", Description:"", MySQLVersion:"5.6"},
    1638: {ErrorNumber:1638, ErrorType:"ServerError", Symbol:"WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED", SQLState:"HY000", Message:"Non-ASCII separator arguments are not fully supported", Description:"", MySQLVersion:"5.6"},
    1639: {ErrorNumber:1639, ErrorType:"ServerError", Symbol:"ER_DEBUG_SYNC_TIMEOUT", SQLState:"HY000", Message:"debug sync point wait timed out", Description:"", MySQLVersion:"5.6"},
    1640: {ErrorNumber:1640, ErrorType:"ServerError", Symbol:"ER_DEBUG_SYNC_HIT_LIMIT", SQLState:"HY000", Message:"debug sync point hit limit reached", Description:"", MySQLVersion:"5.6"},
    1641: {ErrorNumber:1641, ErrorType:"ServerError", Symbol:"ER_DUP_SIGNAL_SET", SQLState:"42000", Message:"Duplicate condition information item '%s'", Description:"", MySQLVersion:"5.6"},
    1642: {ErrorNumber:1642, ErrorType:"ServerError", Symbol:"ER_SIGNAL_WARN", SQLState:"01000", Message:"Unhandled user-defined warning condition", Description:"", MySQLVersion:"5.6"},
    1643: {ErrorNumber:1643, ErrorType:"ServerError", Symbol:"ER_SIGNAL_NOT_FOUND", SQLState:"02000", Message:"Unhandled user-defined not found condition", Description:"", MySQLVersion:"5.6"},
    1644: {ErrorNumber:1644, ErrorType:"ServerError", Symbol:"ER_SIGNAL_EXCEPTION", SQLState:"HY000", Message:"Unhandled user-defined exception condition", Description:"", MySQLVersion:"5.6"},
    1645: {ErrorNumber:1645, ErrorType:"ServerError", Symbol:"ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER", SQLState:"0K000", Message:"RESIGNAL when handler not active", Description:"", MySQLVersion:"5.6"},
    1646: {ErrorNumber:1646, ErrorType:"ServerError", Symbol:"ER_SIGNAL_BAD_CONDITION_TYPE", SQLState:"HY000", Message:"SIGNAL/RESIGNAL can only use a CONDITION defined withSQLSTATE", Description:"", MySQLVersion:"5.6"},
    1647: {ErrorNumber:1647, ErrorType:"ServerError", Symbol:"WARN_COND_ITEM_TRUNCATED", SQLState:"HY000", Message:"Data truncated for condition item '%s'", Description:"", MySQLVersion:"5.6"},
    1648: {ErrorNumber:1648, ErrorType:"ServerError", Symbol:"ER_COND_ITEM_TOO_LONG", SQLState:"HY000", Message:"Data too long for condition item '%s'", Description:"", MySQLVersion:"5.6"},
    1649: {ErrorNumber:1649, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_LOCALE", SQLState:"HY000", Message:"Unknown locale: '%s'", Description:"", MySQLVersion:"5.6"},
    1650: {ErrorNumber:1650, ErrorType:"ServerError", Symbol:"ER_SLAVE_IGNORE_SERVER_IDS", SQLState:"HY000", Message:"The requested server id %d clashes with the slave startupoption --replicate-same-server-id", Description:"", MySQLVersion:"5.6"},
    1651: {ErrorNumber:1651, ErrorType:"ServerError", Symbol:"ER_QUERY_CACHE_DISABLED", SQLState:"HY000", Message:"Query cache is disabled", Description:"restart the server withquery_cache_type=1 to enable it", MySQLVersion:"5.6"},
    1652: {ErrorNumber:1652, ErrorType:"ServerError", Symbol:"ER_SAME_NAME_PARTITION_FIELD", SQLState:"HY000", Message:"Duplicate partition field name '%s'", Description:"", MySQLVersion:"5.6"},
    1653: {ErrorNumber:1653, ErrorType:"ServerError", Symbol:"ER_PARTITION_COLUMN_LIST_ERROR", SQLState:"HY000", Message:"Inconsistency in usage of column lists for partitioning", Description:"", MySQLVersion:"5.6"},
    1654: {ErrorNumber:1654, ErrorType:"ServerError", Symbol:"ER_WRONG_TYPE_COLUMN_VALUE_ERROR", SQLState:"HY000", Message:"Partition column values of incorrect type", Description:"", MySQLVersion:"5.6"},
    1655: {ErrorNumber:1655, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR", SQLState:"HY000", Message:"Too many fields in '%s'", Description:"", MySQLVersion:"5.6"},
    1656: {ErrorNumber:1656, ErrorType:"ServerError", Symbol:"ER_MAXVALUE_IN_VALUES_IN", SQLState:"HY000", Message:"Cannot use MAXVALUE as value in VALUES IN", Description:"", MySQLVersion:"5.6"},
    1657: {ErrorNumber:1657, ErrorType:"ServerError", Symbol:"ER_TOO_MANY_VALUES_ERROR", SQLState:"HY000", Message:"Cannot have more than one value for this type of %spartitioning", Description:"", MySQLVersion:"5.6"},
    1658: {ErrorNumber:1658, ErrorType:"ServerError", Symbol:"ER_ROW_SINGLE_PARTITION_FIELD_ERROR", SQLState:"HY000", Message:"Row expressions in VALUES IN only allowed for multi-fieldcolumn partitioning", Description:"", MySQLVersion:"5.6"},
    1659: {ErrorNumber:1659, ErrorType:"ServerError", Symbol:"ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD", SQLState:"HY000", Message:"Field '%s' is of a not allowed type for this type ofpartitioning", Description:"", MySQLVersion:"5.6"},
    1660: {ErrorNumber:1660, ErrorType:"ServerError", Symbol:"ER_PARTITION_FIELDS_TOO_LONG", SQLState:"HY000", Message:"The total length of the partitioning fields is too large", Description:"", MySQLVersion:"5.6"},
    1661: {ErrorNumber:1661, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since both row-incapable engines and statement-incapableengines are involved.", Description:"", MySQLVersion:"5.6"},
    1662: {ErrorNumber:1662, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_MODE_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since BINLOG_FORMAT = ROW and at least one table uses astorage engine limited to statement-based logging.", Description:"", MySQLVersion:"5.6"},
    1663: {ErrorNumber:1663, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since statement is unsafe, storage engine is limited tostatement-based logging, and BINLOG_FORMAT = MIXED. %s", Description:"", MySQLVersion:"5.6"},
    1664: {ErrorNumber:1664, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since statement is in row format and at least one table uses astorage engine limited to statement-based logging.", Description:"", MySQLVersion:"5.6"},
    1665: {ErrorNumber:1665, ErrorType:"ServerError", Symbol:"ER_BINLOG_STMT_MODE_AND_ROW_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since BINLOG_FORMAT = STATEMENT and at least one table uses astorage engine limited to row-based logging.%s", Description:"", MySQLVersion:"5.6"},
    1666: {ErrorNumber:1666, ErrorType:"ServerError", Symbol:"ER_BINLOG_ROW_INJECTION_AND_STMT_MODE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since statement is in row format and BINLOG_FORMAT =STATEMENT.", Description:"", MySQLVersion:"5.6"},
    1667: {ErrorNumber:1667, ErrorType:"ServerError", Symbol:"ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since more than one engine is involved and at least one engineis self-logging.", Description:"", MySQLVersion:"5.6"},
    1668: {ErrorNumber:1668, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_LIMIT", SQLState:"HY000", Message:"The statement is unsafe because it uses a LIMIT clause.This is unsafe because the set of rows included cannot bepredicted.", Description:"", MySQLVersion:"5.6"},
    1669: {ErrorNumber:1669, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_DELAYED", SQLState:"HY000", Message:"The statement is unsafe because it uses INSERT DELAYED.This is unsafe because the times when rows are inserted cannot bepredicted.", Description:"", MySQLVersion:"5.6"},
    1670: {ErrorNumber:1670, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_SYSTEM_TABLE", SQLState:"HY000", Message:"The statement is unsafe because it uses the general log,slow query log, or performance_schema table(s). This is unsafebecause system tables may differ on slaves.", Description:"", MySQLVersion:"5.6"},
    1671: {ErrorNumber:1671, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_AUTOINC_COLUMNS", SQLState:"HY000", Message:"Statement is unsafe because it invokes a trigger or astored function that inserts into an AUTO_INCREMENT column.Inserted values cannot be logged correctly.", Description:"", MySQLVersion:"5.6"},
    1672: {ErrorNumber:1672, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_UDF", SQLState:"HY000", Message:"Statement is unsafe because it uses a UDF which may notreturn the same value on the slave.", Description:"", MySQLVersion:"5.6"},
    1673: {ErrorNumber:1673, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_SYSTEM_VARIABLE", SQLState:"HY000", Message:"Statement is unsafe because it uses a system variablethat may have a different value on the slave.", Description:"", MySQLVersion:"5.6"},
    1674: {ErrorNumber:1674, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_SYSTEM_FUNCTION", SQLState:"HY000", Message:"Statement is unsafe because it uses a system functionthat may return a different value on the slave.", Description:"", MySQLVersion:"5.6"},
    1675: {ErrorNumber:1675, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS", SQLState:"HY000", Message:"Statement is unsafe because it accesses anon-transactional table after accessing a transactional tablewithin the same transaction.", Description:"", MySQLVersion:"5.6"},
    1676: {ErrorNumber:1676, ErrorType:"ServerError", Symbol:"ER_MESSAGE_AND_STATEMENT", SQLState:"HY000", Message:"%s Statement: %s", Description:"", MySQLVersion:"5.6"},
    1677: {ErrorNumber:1677, ErrorType:"ServerError", Symbol:"ER_SLAVE_CONVERSION_FAILED", SQLState:"HY000", Message:"Column %d of table '%s.%s' cannot be converted from type'%s' to type '%s'", Description:"", MySQLVersion:"5.6"},
    1678: {ErrorNumber:1678, ErrorType:"ServerError", Symbol:"ER_SLAVE_CANT_CREATE_CONVERSION", SQLState:"HY000", Message:"Can't create conversion table for table '%s.%s'", Description:"", MySQLVersion:"5.6"},
    1679: {ErrorNumber:1679, ErrorType:"ServerError", Symbol:"ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT", SQLState:"HY000", Message:"Cannot modify @@session.binlog_format inside atransaction", Description:"", MySQLVersion:"5.6"},
    1680: {ErrorNumber:1680, ErrorType:"ServerError", Symbol:"ER_PATH_LENGTH", SQLState:"HY000", Message:"The path specified for %s is too long.", Description:"", MySQLVersion:"5.6"},
    1681: {ErrorNumber:1681, ErrorType:"ServerError", Symbol:"ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT", SQLState:"HY000", Message:"'%s' is deprecated and will be removed in a futurerelease.", Description:"", MySQLVersion:"5.6"},
    1682: {ErrorNumber:1682, ErrorType:"ServerError", Symbol:"ER_WRONG_NATIVE_TABLE_STRUCTURE", SQLState:"HY000", Message:"Native table '%s'.'%s' has the wrong structure", Description:"", MySQLVersion:"5.6"},
    1683: {ErrorNumber:1683, ErrorType:"ServerError", Symbol:"ER_WRONG_PERFSCHEMA_USAGE", SQLState:"HY000", Message:"Invalid performance_schema usage.", Description:"", MySQLVersion:"5.6"},
    1684: {ErrorNumber:1684, ErrorType:"ServerError", Symbol:"ER_WARN_I_S_SKIPPED_TABLE", SQLState:"HY000", Message:"Table '%s'.'%s' was skipped since its definition is beingmodified by concurrent DDL statement", Description:"", MySQLVersion:"5.6"},
    1685: {ErrorNumber:1685, ErrorType:"ServerError", Symbol:"ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT", SQLState:"HY000", Message:"Cannot modify@@session.binlog_direct_non_transactional_updates inside atransaction", Description:"", MySQLVersion:"5.6"},
    1686: {ErrorNumber:1686, ErrorType:"ServerError", Symbol:"ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT", SQLState:"HY000", Message:"Cannot change the binlog direct flag inside a storedfunction or trigger", Description:"", MySQLVersion:"5.6"},
    1687: {ErrorNumber:1687, ErrorType:"ServerError", Symbol:"ER_SPATIAL_MUST_HAVE_GEOM_COL", SQLState:"42000", Message:"A SPATIAL index may only contain a geometrical typecolumn", Description:"", MySQLVersion:"5.6"},
    1688: {ErrorNumber:1688, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_INDEX_COMMENT", SQLState:"HY000", Message:"Comment for index '%s' is too long (max = %lu)", Description:"", MySQLVersion:"5.6"},
    1689: {ErrorNumber:1689, ErrorType:"ServerError", Symbol:"ER_LOCK_ABORTED", SQLState:"HY000", Message:"Wait on a lock was aborted due to a pending exclusivelock", Description:"", MySQLVersion:"5.6"},
    1690: {ErrorNumber:1690, ErrorType:"ServerError", Symbol:"ER_DATA_OUT_OF_RANGE", SQLState:"22003", Message:"%s value is out of range in '%s'", Description:"", MySQLVersion:"5.6"},
    1691: {ErrorNumber:1691, ErrorType:"ServerError", Symbol:"ER_WRONG_SPVAR_TYPE_IN_LIMIT", SQLState:"HY000", Message:"A variable of a non-integer based type in LIMIT clause", Description:"", MySQLVersion:"5.6"},
    1692: {ErrorNumber:1692, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE", SQLState:"HY000", Message:"Mixing self-logging and non-self-logging engines in astatement is unsafe.", Description:"", MySQLVersion:"5.6"},
    1693: {ErrorNumber:1693, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_MIXED_STATEMENT", SQLState:"HY000", Message:"Statement accesses nontransactional table as well astransactional or temporary table, and writes to any of them.", Description:"", MySQLVersion:"5.6"},
    1694: {ErrorNumber:1694, ErrorType:"ServerError", Symbol:"ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN", SQLState:"HY000", Message:"Cannot modify @@session.sql_log_bin inside a transaction", Description:"", MySQLVersion:"5.6"},
    1695: {ErrorNumber:1695, ErrorType:"ServerError", Symbol:"ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN", SQLState:"HY000", Message:"Cannot change the sql_log_bin inside a stored function ortrigger", Description:"", MySQLVersion:"5.6"},
    1696: {ErrorNumber:1696, ErrorType:"ServerError", Symbol:"ER_FAILED_READ_FROM_PAR_FILE", SQLState:"HY000", Message:"Failed to read from the .par file", Description:"", MySQLVersion:"5.6"},
    1697: {ErrorNumber:1697, ErrorType:"ServerError", Symbol:"ER_VALUES_IS_NOT_INT_TYPE_ERROR", SQLState:"HY000", Message:"VALUES value for partition '%s' must have type INT", Description:"ER_VALUES_IS_NOT_INT_TYPE_ERRORwas added in 5.6.1.", MySQLVersion:"5.6"},
    1698: {ErrorNumber:1698, ErrorType:"ServerError", Symbol:"ER_ACCESS_DENIED_NO_PASSWORD_ERROR", SQLState:"28000", Message:"Access denied for user '%s'@'%s'", Description:"ER_ACCESS_DENIED_NO_PASSWORD_ERRORwas added in 5.6.1.", MySQLVersion:"5.6"},
    1699: {ErrorNumber:1699, ErrorType:"ServerError", Symbol:"ER_SET_PASSWORD_AUTH_PLUGIN", SQLState:"HY000", Message:"SET PASSWORD has no significance for users authenticatingvia plugins", Description:"ER_SET_PASSWORD_AUTH_PLUGIN wasadded in 5.6.1.", MySQLVersion:"5.6"},
    1700: {ErrorNumber:1700, ErrorType:"ServerError", Symbol:"ER_GRANT_PLUGIN_USER_EXISTS", SQLState:"HY000", Message:"GRANT with IDENTIFIED WITH is illegal because the user%-.*s already exists", Description:"ER_GRANT_PLUGIN_USER_EXISTS wasadded in 5.6.1.", MySQLVersion:"5.6"},
    1701: {ErrorNumber:1701, ErrorType:"ServerError", Symbol:"ER_TRUNCATE_ILLEGAL_FK", SQLState:"42000", Message:"Cannot truncate a table referenced in a foreign keyconstraint (%s)", Description:"ER_TRUNCATE_ILLEGAL_FK was addedin 5.6.1.", MySQLVersion:"5.6"},
    1702: {ErrorNumber:1702, ErrorType:"ServerError", Symbol:"ER_PLUGIN_IS_PERMANENT", SQLState:"HY000", Message:"Plugin '%s' is force_plus_permanent and can not beunloaded", Description:"ER_PLUGIN_IS_PERMANENT was addedin 5.6.1.", MySQLVersion:"5.6"},
    1703: {ErrorNumber:1703, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN", SQLState:"HY000", Message:"The requested value for the heartbeat period is less than1 millisecond. The value is reset to 0, meaning that heartbeatingwill effectively be disabled.", Description:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MINwas added in 5.6.1.", MySQLVersion:"5.6"},
    1704: {ErrorNumber:1704, ErrorType:"ServerError", Symbol:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX", SQLState:"HY000", Message:"The requested value for the heartbeat period exceeds thevalue of `slave_net_timeout' seconds. A sensible value for theperiod should be less than the timeout.", Description:"ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAXwas added in 5.6.1.", MySQLVersion:"5.6"},
    1705: {ErrorNumber:1705, ErrorType:"ServerError", Symbol:"ER_STMT_CACHE_FULL", SQLState:"HY000", Message:"Multi-row statements required more than'max_binlog_stmt_cache_size' bytes of storage", Description:"ER_STMT_CACHE_FULL was added in5.6.1.", MySQLVersion:"5.6"},
    1706: {ErrorNumber:1706, ErrorType:"ServerError", Symbol:"ER_MULTI_UPDATE_KEY_CONFLICT", SQLState:"HY000", Message:"Primary key/partition key update is not allowed since thetable is updated both as '%s' and '%s'.", Description:"ER_MULTI_UPDATE_KEY_CONFLICT wasadded in 5.6.2.", MySQLVersion:"5.6"},
    1707: {ErrorNumber:1707, ErrorType:"ServerError", Symbol:"ER_TABLE_NEEDS_REBUILD", SQLState:"HY000", Message:"Table rebuild required. Please do \"ALTER TABLE `%s`FORCE\" or dump/reload to fix it!", Description:"ER_TABLE_NEEDS_REBUILD was addedin 5.6.3.", MySQLVersion:"5.6"},
    1708: {ErrorNumber:1708, ErrorType:"ServerError", Symbol:"WARN_OPTION_BELOW_LIMIT", SQLState:"HY000", Message:"The value of '%s' should be no less than the value of'%s'", Description:"WARN_OPTION_BELOW_LIMIT was addedin 5.6.3.", MySQLVersion:"5.6"},
    1709: {ErrorNumber:1709, ErrorType:"ServerError", Symbol:"ER_INDEX_COLUMN_TOO_LONG", SQLState:"HY000", Message:"Index column size too large. The maximum column size is%lu bytes.", Description:"ER_INDEX_COLUMN_TOO_LONG was addedin 5.6.3.", MySQLVersion:"5.6"},
    1710: {ErrorNumber:1710, ErrorType:"ServerError", Symbol:"ER_ERROR_IN_TRIGGER_BODY", SQLState:"HY000", Message:"Trigger '%s' has an error in its body: '%s'", Description:"ER_ERROR_IN_TRIGGER_BODY was addedin 5.6.3.", MySQLVersion:"5.6"},
    1711: {ErrorNumber:1711, ErrorType:"ServerError", Symbol:"ER_ERROR_IN_UNKNOWN_TRIGGER_BODY", SQLState:"HY000", Message:"Unknown trigger has an error in its body: '%s'", Description:"ER_ERROR_IN_UNKNOWN_TRIGGER_BODYwas added in 5.6.3.", MySQLVersion:"5.6"},
    1712: {ErrorNumber:1712, ErrorType:"ServerError", Symbol:"ER_INDEX_CORRUPT", SQLState:"HY000", Message:"Index %s is corrupted", Description:"ER_INDEX_CORRUPT was added in5.6.3.", MySQLVersion:"5.6"},
    1713: {ErrorNumber:1713, ErrorType:"ServerError", Symbol:"ER_UNDO_RECORD_TOO_BIG", SQLState:"HY000", Message:"Undo log record is too big.", Description:"ER_UNDO_RECORD_TOO_BIG was addedin 5.6.4.", MySQLVersion:"5.6"},
    1714: {ErrorNumber:1714, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT", SQLState:"HY000", Message:"INSERT IGNORE... SELECT is unsafe because the order inwhich rows are retrieved by the SELECT determines which (if any)rows are ignored. This order cannot be predicted and may differ onmaster and the slave.", Description:"ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECTwas added in 5.6.4.", MySQLVersion:"5.6"},
    1715: {ErrorNumber:1715, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE", SQLState:"HY000", Message:"INSERT... SELECT... ON DUPLICATE KEY UPDATE is unsafebecause the order in which rows are retrieved by the SELECTdetermines which (if any) rows are updated. This order cannot bepredicted and may differ on master and the slave.", Description:"ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATEwas added in 5.6.4.", MySQLVersion:"5.6"},
    1716: {ErrorNumber:1716, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_REPLACE_SELECT", SQLState:"HY000", Message:"REPLACE... SELECT is unsafe because the order in whichrows are retrieved by the SELECT determines which (if any) rowsare replaced. This order cannot be predicted and may differ onmaster and the slave.", Description:"ER_BINLOG_UNSAFE_REPLACE_SELECTwas added in 5.6.4.", MySQLVersion:"5.6"},
    1717: {ErrorNumber:1717, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT", SQLState:"HY000", Message:"CREATE... IGNORE SELECT is unsafe because the order inwhich rows are retrieved by the SELECT determines which (if any)rows are ignored. This order cannot be predicted and may differ onmaster and the slave.", Description:"ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECTwas added in 5.6.4.", MySQLVersion:"5.6"},
    1718: {ErrorNumber:1718, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT", SQLState:"HY000", Message:"CREATE... REPLACE SELECT is unsafe because the order inwhich rows are retrieved by the SELECT determines which (if any)rows are replaced. This order cannot be predicted and may differon master and the slave.", Description:"ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECTwas added in 5.6.4.", MySQLVersion:"5.6"},
    1719: {ErrorNumber:1719, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_UPDATE_IGNORE", SQLState:"HY000", Message:"UPDATE IGNORE is unsafe because the order in which rowsare updated determines which (if any) rows are ignored. This ordercannot be predicted and may differ on master and the slave.", Description:"ER_BINLOG_UNSAFE_UPDATE_IGNORE wasadded in 5.6.4.", MySQLVersion:"5.6"},
    1720: {ErrorNumber:1720, ErrorType:"ServerError", Symbol:"ER_PLUGIN_NO_UNINSTALL", SQLState:"HY000", Message:"Plugin '%s' is marked as not dynamically uninstallable.You have to stop the server to uninstall it.", Description:"ER_PLUGIN_NO_UNINSTALL was addedin 5.6.4.", MySQLVersion:"5.6"},
    1721: {ErrorNumber:1721, ErrorType:"ServerError", Symbol:"ER_PLUGIN_NO_INSTALL", SQLState:"HY000", Message:"Plugin '%s' is marked as not dynamically installable. Youhave to stop the server to install it.", Description:"ER_PLUGIN_NO_INSTALL was added in5.6.4.", MySQLVersion:"5.6"},
    1722: {ErrorNumber:1722, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT", SQLState:"HY000", Message:"Statements writing to a table with an auto-incrementcolumn after selecting from another table are unsafe because theorder in which rows are retrieved determines what (if any) rowswill be written. This order cannot be predicted and may differ onmaster and the slave.", Description:"ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECTwas added in 5.6.5.", MySQLVersion:"5.6"},
    1723: {ErrorNumber:1723, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC", SQLState:"HY000", Message:"CREATE TABLE... SELECT... on a table with anauto-increment column is unsafe because the order in which rowsare retrieved by the SELECT determines which (if any) rows areinserted. This order cannot be predicted and may differ on masterand the slave.", Description:"ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINCwas added in 5.6.5.", MySQLVersion:"5.6"},
    1724: {ErrorNumber:1724, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_INSERT_TWO_KEYS", SQLState:"HY000", Message:"INSERT... ON DUPLICATE KEY UPDATE on a table with morethan one UNIQUE KEY is unsafe", Description:"ER_BINLOG_UNSAFE_INSERT_TWO_KEYSwas added in 5.6.6.", MySQLVersion:"5.6"},
    1725: {ErrorNumber:1725, ErrorType:"ServerError", Symbol:"ER_TABLE_IN_FK_CHECK", SQLState:"HY000", Message:"Table is being used in foreign key check.", Description:"ER_TABLE_IN_FK_CHECK was added in5.6.6.", MySQLVersion:"5.6"},
    1726: {ErrorNumber:1726, ErrorType:"ServerError", Symbol:"ER_UNSUPPORTED_ENGINE", SQLState:"HY000", Message:"Storage engine '%s' does not support system tables.[%s.%s]", Description:"ER_UNSUPPORTED_ENGINE was added in5.6.6.", MySQLVersion:"5.6"},
    1727: {ErrorNumber:1727, ErrorType:"ServerError", Symbol:"ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST", SQLState:"HY000", Message:"INSERT into autoincrement field which is not the firstpart in the composed primary key is unsafe.", Description:"ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRSTwas added in 5.6.6.", MySQLVersion:"5.6"},
    1728: {ErrorNumber:1728, ErrorType:"ServerError", Symbol:"ER_CANNOT_LOAD_FROM_TABLE_V2", SQLState:"HY000", Message:"Cannot load from %s.%s. The table is probably corrupted", Description:"", MySQLVersion:"5.6"},
    1729: {ErrorNumber:1729, ErrorType:"ServerError", Symbol:"ER_MASTER_DELAY_VALUE_OUT_OF_RANGE", SQLState:"HY000", Message:"The requested value %s for the master delay exceeds themaximum %u", Description:"", MySQLVersion:"5.6"},
    1730: {ErrorNumber:1730, ErrorType:"ServerError", Symbol:"ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT", SQLState:"HY000", Message:"Only Format_description_log_event and row events areallowed in BINLOG statements (but %s was provided)", Description:"", MySQLVersion:"5.6"},
    1731: {ErrorNumber:1731, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_DIFFERENT_OPTION", SQLState:"HY000", Message:"Non matching attribute '%s' between partition and table", Description:"", MySQLVersion:"5.6"},
    1732: {ErrorNumber:1732, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_PART_TABLE", SQLState:"HY000", Message:"Table to exchange with partition is partitioned: '%s'", Description:"", MySQLVersion:"5.6"},
    1733: {ErrorNumber:1733, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_TEMP_TABLE", SQLState:"HY000", Message:"Table to exchange with partition is temporary: '%s'", Description:"", MySQLVersion:"5.6"},
    1734: {ErrorNumber:1734, ErrorType:"ServerError", Symbol:"ER_PARTITION_INSTEAD_OF_SUBPARTITION", SQLState:"HY000", Message:"Subpartitioned table, use subpartition instead ofpartition", Description:"", MySQLVersion:"5.6"},
    1735: {ErrorNumber:1735, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_PARTITION", SQLState:"HY000", Message:"Unknown partition '%s' in table '%s'", Description:"", MySQLVersion:"5.6"},
    1736: {ErrorNumber:1736, ErrorType:"ServerError", Symbol:"ER_TABLES_DIFFERENT_METADATA", SQLState:"HY000", Message:"Tables have different definitions", Description:"", MySQLVersion:"5.6"},
    1737: {ErrorNumber:1737, ErrorType:"ServerError", Symbol:"ER_ROW_DOES_NOT_MATCH_PARTITION", SQLState:"HY000", Message:"Found a row that does not match the partition", Description:"", MySQLVersion:"5.6"},
    1738: {ErrorNumber:1738, ErrorType:"ServerError", Symbol:"ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX", SQLState:"HY000", Message:"Option binlog_cache_size (%lu) is greater thanmax_binlog_cache_size (%lu)", Description:"ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAXwas added in 5.6.1.", MySQLVersion:"5.6"},
    1739: {ErrorNumber:1739, ErrorType:"ServerError", Symbol:"ER_WARN_INDEX_NOT_APPLICABLE", SQLState:"HY000", Message:"Cannot use %s access on index '%s' due to type orcollation conversion on field '%s'", Description:"ER_WARN_INDEX_NOT_APPLICABLE wasadded in 5.6.1.", MySQLVersion:"5.6"},
    1740: {ErrorNumber:1740, ErrorType:"ServerError", Symbol:"ER_PARTITION_EXCHANGE_FOREIGN_KEY", SQLState:"HY000", Message:"Table to exchange with partition has foreign keyreferences: '%s'", Description:"ER_PARTITION_EXCHANGE_FOREIGN_KEYwas added in 5.6.1.", MySQLVersion:"5.6"},
    1741: {ErrorNumber:1741, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_KEY_VALUE", SQLState:"HY000", Message:"Key value '%s' was not found in table '%s.%s'", Description:"ER_NO_SUCH_KEY_VALUE was added in5.6.1.", MySQLVersion:"5.6"},
    1742: {ErrorNumber:1742, ErrorType:"ServerError", Symbol:"ER_RPL_INFO_DATA_TOO_LONG", SQLState:"HY000", Message:"Data for column '%s' too long", Description:"ER_RPL_INFO_DATA_TOO_LONG wasadded in 5.6.1.", MySQLVersion:"5.6"},
    1743: {ErrorNumber:1743, ErrorType:"ServerError", Symbol:"ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE", SQLState:"HY000", Message:"Replication event checksum verification failed whilereading from network.", Description:"ER_NETWORK_READ_EVENT_CHECKSUM_FAILUREwas added in 5.6.1.", MySQLVersion:"5.6"},
    1744: {ErrorNumber:1744, ErrorType:"ServerError", Symbol:"ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE", SQLState:"HY000", Message:"Replication event checksum verification failed whilereading from a log file.", Description:"ER_BINLOG_READ_EVENT_CHECKSUM_FAILUREwas added in 5.6.1.", MySQLVersion:"5.6"},
    1745: {ErrorNumber:1745, ErrorType:"ServerError", Symbol:"ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX", SQLState:"HY000", Message:"Option binlog_stmt_cache_size (%lu) is greater thanmax_binlog_stmt_cache_size (%lu)", Description:"ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAXwas added in 5.6.1.", MySQLVersion:"5.6"},
    1746: {ErrorNumber:1746, ErrorType:"ServerError", Symbol:"ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT", SQLState:"HY000", Message:"Can't update table '%s' while '%s' is being created.", Description:"ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECTwas added in 5.6.2.", MySQLVersion:"5.6"},
    1747: {ErrorNumber:1747, ErrorType:"ServerError", Symbol:"ER_PARTITION_CLAUSE_ON_NONPARTITIONED", SQLState:"HY000", Message:"PARTITION () clause on non partitioned table", Description:"ER_PARTITION_CLAUSE_ON_NONPARTITIONEDwas added in 5.6.2.", MySQLVersion:"5.6"},
    1748: {ErrorNumber:1748, ErrorType:"ServerError", Symbol:"ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET", SQLState:"HY000", Message:"Found a row not matching the given partition set", Description:"ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SETwas added in 5.6.2.", MySQLVersion:"5.6"},
    1749: {ErrorNumber:1749, ErrorType:"ServerError", Symbol:"ER_NO_SUCH_PARTITION__UNUSED", SQLState:"HY000", Message:"partition '%s' doesn't exist", Description:"ER_NO_SUCH_PARTITION__UNUSED wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1750: {ErrorNumber:1750, ErrorType:"ServerError", Symbol:"ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE", SQLState:"HY000", Message:"Failure while changing the type of replicationrepository: %s.", Description:"ER_CHANGE_RPL_INFO_REPOSITORY_FAILUREwas added in 5.6.3.", MySQLVersion:"5.6"},
    1751: {ErrorNumber:1751, ErrorType:"ServerError", Symbol:"ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE", SQLState:"HY000", Message:"The creation of some temporary tables could not be rolledback.", Description:"ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLEwas added in 5.6.3.", MySQLVersion:"5.6"},
    1752: {ErrorNumber:1752, ErrorType:"ServerError", Symbol:"ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE", SQLState:"HY000", Message:"Some temporary tables were dropped, but these operationscould not be rolled back.", Description:"ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLEwas added in 5.6.3.", MySQLVersion:"5.6"},
    1753: {ErrorNumber:1753, ErrorType:"ServerError", Symbol:"ER_MTS_FEATURE_IS_NOT_SUPPORTED", SQLState:"HY000", Message:"%s is not supported in multi-threaded slave mode. %s", Description:"ER_MTS_FEATURE_IS_NOT_SUPPORTEDwas added in 5.6.3.", MySQLVersion:"5.6"},
    1754: {ErrorNumber:1754, ErrorType:"ServerError", Symbol:"ER_MTS_UPDATED_DBS_GREATER_MAX", SQLState:"HY000", Message:"The number of modified databases exceeds the maximum %d", Description:"ER_MTS_UPDATED_DBS_GREATER_MAX wasadded in 5.6.3.", MySQLVersion:"5.6"},
    1755: {ErrorNumber:1755, ErrorType:"ServerError", Symbol:"ER_MTS_CANT_PARALLEL", SQLState:"HY000", Message:"Cannot execute the current event group in the parallelmode. Encountered event %s, relay-log name %s, position %s whichprevents execution of this event group in parallel mode. Reason:%s.", Description:"ER_MTS_CANT_PARALLEL was added in5.6.3.", MySQLVersion:"5.6"},
    1756: {ErrorNumber:1756, ErrorType:"ServerError", Symbol:"ER_MTS_INCONSISTENT_DATA", SQLState:"HY000", Message:"%s", Description:"ER_MTS_INCONSISTENT_DATA was addedin 5.6.3.", MySQLVersion:"5.6"},
    1757: {ErrorNumber:1757, ErrorType:"ServerError", Symbol:"ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING", SQLState:"HY000", Message:"FULLTEXT index is not supported for partitioned tables.", Description:"ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONINGwas added in 5.6.4.", MySQLVersion:"5.6"},
    1758: {ErrorNumber:1758, ErrorType:"ServerError", Symbol:"ER_DA_INVALID_CONDITION_NUMBER", SQLState:"35000", Message:"Invalid condition number", Description:"ER_DA_INVALID_CONDITION_NUMBER wasadded in 5.6.4.", MySQLVersion:"5.6"},
    1759: {ErrorNumber:1759, ErrorType:"ServerError", Symbol:"ER_INSECURE_PLAIN_TEXT", SQLState:"HY000", Message:"Sending passwords in plain text without SSL/TLS isextremely insecure.", Description:"ER_INSECURE_PLAIN_TEXT was addedin 5.6.4.", MySQLVersion:"5.6"},
    1760: {ErrorNumber:1760, ErrorType:"ServerError", Symbol:"ER_INSECURE_CHANGE_MASTER", SQLState:"HY000", Message:"Storing MySQL user name or password information in themaster info repository is not secure and is therefore notrecommended. Please consider using the USER and PASSWORDconnection options for START SLAVE", Description:"ER_INSECURE_CHANGE_MASTER wasadded in 5.6.4.", MySQLVersion:"5.6"},
    1761: {ErrorNumber:1761, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO", SQLState:"23000", Message:"Foreign key constraint for table '%s', record '%s' wouldlead to a duplicate entry in table '%s', key '%s'", Description:"ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFOwas added in 5.6.4.", MySQLVersion:"5.6"},
    1762: {ErrorNumber:1762, ErrorType:"ServerError", Symbol:"ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO", SQLState:"23000", Message:"Foreign key constraint for table '%s', record '%s' wouldlead to a duplicate entry in a child table", Description:"ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFOwas added in 5.6.4.", MySQLVersion:"5.6"},
    1763: {ErrorNumber:1763, ErrorType:"ServerError", Symbol:"ER_SQLTHREAD_WITH_SECURE_SLAVE", SQLState:"HY000", Message:"Setting authentication options is not possible when onlythe Slave SQL Thread is being started.", Description:"ER_SQLTHREAD_WITH_SECURE_SLAVE wasadded in 5.6.4.", MySQLVersion:"5.6"},
    1764: {ErrorNumber:1764, ErrorType:"ServerError", Symbol:"ER_TABLE_HAS_NO_FT", SQLState:"HY000", Message:"The table does not have FULLTEXT index to support thisquery", Description:"ER_TABLE_HAS_NO_FT was added in5.6.4.", MySQLVersion:"5.6"},
    1765: {ErrorNumber:1765, ErrorType:"ServerError", Symbol:"ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER", SQLState:"HY000", Message:"The system variable %s cannot be set in stored functionsor triggers.", Description:"ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGERwas added in 5.6.5.", MySQLVersion:"5.6"},
    1766: {ErrorNumber:1766, ErrorType:"ServerError", Symbol:"ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION", SQLState:"HY000", Message:"The system variable %s cannot be set when there is anongoing transaction.", Description:"ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTIONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1767: {ErrorNumber:1767, ErrorType:"ServerError", Symbol:"ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT has the value %s,which is not listed in @@SESSION.GTID_NEXT_LIST.", Description:"ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LISTwas added in 5.6.5.", MySQLVersion:"5.6"},
    1768: {ErrorNumber:1768, ErrorType:"ServerError", Symbol:"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT cannot changeinside a transaction.", Description:"ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULLwas added in 5.6.5.", MySQLVersion:"5.6"},
    1769: {ErrorNumber:1769, ErrorType:"ServerError", Symbol:"ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION", SQLState:"HY000", Message:"The statement 'SET %s' cannot invoke a stored function.", Description:"ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTIONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1770: {ErrorNumber:1770, ErrorType:"ServerError", Symbol:"ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL", SQLState:"HY000", Message:"The system variable @@SESSION.GTID_NEXT cannot be'AUTOMATIC' when @@SESSION.GTID_NEXT_LIST is non-NULL.", Description:"ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULLwas added in 5.6.5.", MySQLVersion:"5.6"},
    1771: {ErrorNumber:1771, ErrorType:"ServerError", Symbol:"ER_SKIPPING_LOGGED_TRANSACTION", SQLState:"HY000", Message:"Skipping transaction %s because it has already beenexecuted and logged.", Description:"ER_SKIPPING_LOGGED_TRANSACTION wasadded in 5.6.5.", MySQLVersion:"5.6"},
    1772: {ErrorNumber:1772, ErrorType:"ServerError", Symbol:"ER_MALFORMED_GTID_SET_SPECIFICATION", SQLState:"HY000", Message:"Malformed GTID set specification '%s'.", Description:"ER_MALFORMED_GTID_SET_SPECIFICATIONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1773: {ErrorNumber:1773, ErrorType:"ServerError", Symbol:"ER_MALFORMED_GTID_SET_ENCODING", SQLState:"HY000", Message:"Malformed GTID set encoding.", Description:"ER_MALFORMED_GTID_SET_ENCODING wasadded in 5.6.5.", MySQLVersion:"5.6"},
    1774: {ErrorNumber:1774, ErrorType:"ServerError", Symbol:"ER_MALFORMED_GTID_SPECIFICATION", SQLState:"HY000", Message:"Malformed GTID specification '%s'.", Description:"ER_MALFORMED_GTID_SPECIFICATIONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1775: {ErrorNumber:1775, ErrorType:"ServerError", Symbol:"ER_GNO_EXHAUSTED", SQLState:"HY000", Message:"Impossible to generate Global Transaction Identifier: theinteger component reached the maximal value. Restart the serverwith a new server_uuid.", Description:"ER_GNO_EXHAUSTED was added in5.6.5.", MySQLVersion:"5.6"},
    1776: {ErrorNumber:1776, ErrorType:"ServerError", Symbol:"ER_BAD_SLAVE_AUTO_POSITION", SQLState:"HY000", Message:"Parameters MASTER_LOG_FILE, MASTER_LOG_POS,RELAY_LOG_FILE and RELAY_LOG_POS cannot be set whenMASTER_AUTO_POSITION is active.", Description:"ER_BAD_SLAVE_AUTO_POSITION wasadded in 5.6.5.", MySQLVersion:"5.6"},
    1777: {ErrorNumber:1777, ErrorType:"ServerError", Symbol:"ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON", SQLState:"HY000", Message:"CHANGE MASTER TO MASTER_AUTO_POSITION = 1 can only beexecuted when @@GLOBAL.GTID_MODE = ON.", Description:"ER_AUTO_POSITION_REQUIRES_GTID_MODE_ONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1778: {ErrorNumber:1778, ErrorType:"ServerError", Symbol:"ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET", SQLState:"HY000", Message:"Cannot execute statements with implicit commit inside atransaction when @@SESSION.GTID_NEXT != AUTOMATIC.", Description:"ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SETwas added in 5.6.5.", MySQLVersion:"5.6"},
    1779: {ErrorNumber:1779, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON", SQLState:"HY000", Message:"@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_2 requires@@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1.", Description:"ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ONwas added in 5.6.9.", MySQLVersion:"5.6"},
    1780: {ErrorNumber:1780, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_REQUIRES_BINLOG", SQLState:"HY000", Message:"@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_1 orUPGRADE_STEP_2 requires --log-bin and --log-slave-updates.", Description:"ER_GTID_MODE_REQUIRES_BINLOG wasadded in 5.6.5.", MySQLVersion:"5.6"},
    1781: {ErrorNumber:1781, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT cannot be set to UUID:NUMBER when@@GLOBAL.GTID_MODE = OFF.", Description:"ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFFwas added in 5.6.5.", MySQLVersion:"5.6"},
    1782: {ErrorNumber:1782, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT cannot be set to ANONYMOUS when@@GLOBAL.GTID_MODE = ON.", Description:"ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1783: {ErrorNumber:1783, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT_LIST cannot be set to a non-NULLvalue when @@GLOBAL.GTID_MODE = OFF.", Description:"ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFFwas added in 5.6.5.", MySQLVersion:"5.6"},
    1784: {ErrorNumber:1784, ErrorType:"ServerError", Symbol:"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"Found a Gtid_log_event or Previous_gtids_log_event when@@GLOBAL.GTID_MODE = OFF.", Description:"ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFFwas added in 5.6.5.", MySQLVersion:"5.6"},
    1785: {ErrorNumber:1785, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE", SQLState:"HY000", Message:"When @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1, updates tonon-transactional tables can only be done in either autocommittedstatements or single-statement transactions, and never in the samestatement as updates to transactional tables.", Description:"ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLEwas added in 5.6.5.", MySQLVersion:"5.6"},
    1786: {ErrorNumber:1786, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_CREATE_SELECT", SQLState:"HY000", Message:"CREATE TABLE ... SELECT is forbidden when@@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1.", Description:"ER_GTID_UNSAFE_CREATE_SELECT wasadded in 5.6.5.", MySQLVersion:"5.6"},
    1787: {ErrorNumber:1787, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTION", SQLState:"HY000", Message:"When @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1, thestatements CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can beexecuted in a non-transactional context only, and require thatAUTOCOMMIT = 1. These statements are also not allowed in afunction or trigger because functions and triggers are alsoconsidered to be multi-statement transactions.", Description:"ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTIONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1788: {ErrorNumber:1788, ErrorType:"ServerError", Symbol:"ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME", SQLState:"HY000", Message:"The value of @@GLOBAL.GTID_MODE can only change one stepat a time: OFF <-> UPGRADE_STEP_1 <-> UPGRADE_STEP_2<-> ON. Also note that this value must be stepped up or downsimultaneously on all servers", Description:"ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIMEwas added in 5.6.5.", MySQLVersion:"5.6"},
    1789: {ErrorNumber:1789, ErrorType:"ServerError", Symbol:"ER_MASTER_HAS_PURGED_REQUIRED_GTIDS", SQLState:"HY000", Message:"The slave is connecting using CHANGE MASTER TOMASTER_AUTO_POSITION = 1, but the master has purged binary logscontaining GTIDs that the slave requires. Replicate the missingtransactions from elsewhere, or provision a new slave from backup.Consider increasing the master's binary log expiration period. %s.", Description:"ER_MASTER_HAS_PURGED_REQUIRED_GTIDSwas added in 5.6.5.", MySQLVersion:"5.6"},
    1790: {ErrorNumber:1790, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID", SQLState:"HY000", Message:"@@SESSION.GTID_NEXT cannot be changed by a client thatowns a GTID. The client owns %s. Ownership is released on COMMITor ROLLBACK.", Description:"ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTIDwas added in 5.6.5.", MySQLVersion:"5.6"},
    1791: {ErrorNumber:1791, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_EXPLAIN_FORMAT", SQLState:"HY000", Message:"Unknown EXPLAIN format name: '%s'", Description:"ER_UNKNOWN_EXPLAIN_FORMAT wasadded in 5.6.5.", MySQLVersion:"5.6"},
    1792: {ErrorNumber:1792, ErrorType:"ServerError", Symbol:"ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION", SQLState:"25006", Message:"Cannot execute statement in a READ ONLY transaction.", Description:"ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTIONwas added in 5.6.5.", MySQLVersion:"5.6"},
    1793: {ErrorNumber:1793, ErrorType:"ServerError", Symbol:"ER_TOO_LONG_TABLE_PARTITION_COMMENT", SQLState:"HY000", Message:"Comment for table partition '%s' is too long (max = %lu)", Description:"ER_TOO_LONG_TABLE_PARTITION_COMMENTwas added in 5.6.6.", MySQLVersion:"5.6"},
    1794: {ErrorNumber:1794, ErrorType:"ServerError", Symbol:"ER_SLAVE_CONFIGURATION", SQLState:"HY000", Message:"Slave is not configured or failed to initialize properly.You must at least set --server-id to enable either a master or aslave. Additional error messages can be found in the MySQL errorlog.", Description:"ER_SLAVE_CONFIGURATION was addedin 5.6.6.", MySQLVersion:"5.6"},
    1795: {ErrorNumber:1795, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_LIMIT", SQLState:"HY000", Message:"InnoDB presently supports one FULLTEXT index creation ata time", Description:"ER_INNODB_FT_LIMIT was added in5.6.4.", MySQLVersion:"5.6"},
    1796: {ErrorNumber:1796, ErrorType:"ServerError", Symbol:"ER_INNODB_NO_FT_TEMP_TABLE", SQLState:"HY000", Message:"Cannot create FULLTEXT index on temporary InnoDB table", Description:"ER_INNODB_NO_FT_TEMP_TABLE wasadded in 5.6.4.", MySQLVersion:"5.6"},
    1797: {ErrorNumber:1797, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_WRONG_DOCID_COLUMN", SQLState:"HY000", Message:"Column '%s' is of wrong type for an InnoDB FULLTEXT index", Description:"ER_INNODB_FT_WRONG_DOCID_COLUMNwas added in 5.6.6.", MySQLVersion:"5.6"},
    1798: {ErrorNumber:1798, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_WRONG_DOCID_INDEX", SQLState:"HY000", Message:"Index '%s' is of wrong type for an InnoDB FULLTEXT index", Description:"ER_INNODB_FT_WRONG_DOCID_INDEX wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1799: {ErrorNumber:1799, ErrorType:"ServerError", Symbol:"ER_INNODB_ONLINE_LOG_TOO_BIG", SQLState:"HY000", Message:"Creating index '%s' required more than'innodb_online_alter_log_max_size' bytes of modification log.Please try again.", Description:"ER_INNODB_ONLINE_LOG_TOO_BIG wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1800: {ErrorNumber:1800, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_ALTER_ALGORITHM", SQLState:"HY000", Message:"Unknown ALGORITHM '%s'", Description:"ER_UNKNOWN_ALTER_ALGORITHM wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1801: {ErrorNumber:1801, ErrorType:"ServerError", Symbol:"ER_UNKNOWN_ALTER_LOCK", SQLState:"HY000", Message:"Unknown LOCK type '%s'", Description:"ER_UNKNOWN_ALTER_LOCK was added in5.6.6.", MySQLVersion:"5.6"},
    1802: {ErrorNumber:1802, ErrorType:"ServerError", Symbol:"ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS", SQLState:"HY000", Message:"CHANGE MASTER cannot be executed when the slave wasstopped with an error or killed in MTS mode. Consider using RESETSLAVE or START SLAVE UNTIL.", Description:"ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPSwas added in 5.6.6.", MySQLVersion:"5.6"},
    1803: {ErrorNumber:1803, ErrorType:"ServerError", Symbol:"ER_MTS_RECOVERY_FAILURE", SQLState:"HY000", Message:"Cannot recover after SLAVE errored out in parallelexecution mode. Additional error messages can be found in theMySQL error log.", Description:"ER_MTS_RECOVERY_FAILURE was addedin 5.6.6.", MySQLVersion:"5.6"},
    1804: {ErrorNumber:1804, ErrorType:"ServerError", Symbol:"ER_MTS_RESET_WORKERS", SQLState:"HY000", Message:"Cannot clean up worker info tables. Additional errormessages can be found in the MySQL error log.", Description:"ER_MTS_RESET_WORKERS was added in5.6.6.", MySQLVersion:"5.6"},
    1805: {ErrorNumber:1805, ErrorType:"ServerError", Symbol:"ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2", SQLState:"HY000", Message:"Column count of %s.%s is wrong. Expected %d, found %d.The table is probably corrupted", Description:"", MySQLVersion:"5.6"},
    1806: {ErrorNumber:1806, ErrorType:"ServerError", Symbol:"ER_SLAVE_SILENT_RETRY_TRANSACTION", SQLState:"HY000", Message:"Slave must silently retry current transaction", Description:"ER_SLAVE_SILENT_RETRY_TRANSACTIONwas added in 5.6.6.", MySQLVersion:"5.6"},
    1807: {ErrorNumber:1807, ErrorType:"ServerError", Symbol:"ER_DISCARD_FK_CHECKS_RUNNING", SQLState:"HY000", Message:"There is a foreign key check running on table '%s'.Cannot discard the table.", Description:"ER_DISCARD_FK_CHECKS_RUNNING wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1808: {ErrorNumber:1808, ErrorType:"ServerError", Symbol:"ER_TABLE_SCHEMA_MISMATCH", SQLState:"HY000", Message:"Schema mismatch (%s)", Description:"ER_TABLE_SCHEMA_MISMATCH was addedin 5.6.6.", MySQLVersion:"5.6"},
    1809: {ErrorNumber:1809, ErrorType:"ServerError", Symbol:"ER_TABLE_IN_SYSTEM_TABLESPACE", SQLState:"HY000", Message:"Table '%s' in system tablespace", Description:"ER_TABLE_IN_SYSTEM_TABLESPACE wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1810: {ErrorNumber:1810, ErrorType:"ServerError", Symbol:"ER_IO_READ_ERROR", SQLState:"HY000", Message:"IO Read error: (%lu, %s) %s", Description:"ER_IO_READ_ERROR was added in5.6.6.", MySQLVersion:"5.6"},
    1811: {ErrorNumber:1811, ErrorType:"ServerError", Symbol:"ER_IO_WRITE_ERROR", SQLState:"HY000", Message:"IO Write error: (%lu, %s) %s", Description:"ER_IO_WRITE_ERROR was added in5.6.6.", MySQLVersion:"5.6"},
    1812: {ErrorNumber:1812, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_MISSING", SQLState:"HY000", Message:"Tablespace is missing for table '%s'", Description:"ER_TABLESPACE_MISSING was added in5.6.6.", MySQLVersion:"5.6"},
    1813: {ErrorNumber:1813, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_EXISTS", SQLState:"HY000", Message:"Tablespace for table '%s' exists. Please DISCARD thetablespace before IMPORT.", Description:"ER_TABLESPACE_EXISTS was added in5.6.6.", MySQLVersion:"5.6"},
    1814: {ErrorNumber:1814, ErrorType:"ServerError", Symbol:"ER_TABLESPACE_DISCARDED", SQLState:"HY000", Message:"Tablespace has been discarded for table '%s'", Description:"ER_TABLESPACE_DISCARDED was addedin 5.6.6.", MySQLVersion:"5.6"},
    1815: {ErrorNumber:1815, ErrorType:"ServerError", Symbol:"ER_INTERNAL_ERROR", SQLState:"HY000", Message:"Internal error: %s", Description:"ER_INTERNAL_ERROR was added in5.6.6.", MySQLVersion:"5.6"},
    1816: {ErrorNumber:1816, ErrorType:"ServerError", Symbol:"ER_INNODB_IMPORT_ERROR", SQLState:"HY000", Message:"ALTER TABLE '%s' IMPORT TABLESPACE failed with error %lu: '%s'", Description:"ER_INNODB_IMPORT_ERROR was addedin 5.6.6.", MySQLVersion:"5.6"},
    1817: {ErrorNumber:1817, ErrorType:"ServerError", Symbol:"ER_INNODB_INDEX_CORRUPT", SQLState:"HY000", Message:"Index corrupt: %s", Description:"ER_INNODB_INDEX_CORRUPT was addedin 5.6.6.", MySQLVersion:"5.6"},
    1818: {ErrorNumber:1818, ErrorType:"ServerError", Symbol:"ER_INVALID_YEAR_COLUMN_LENGTH", SQLState:"HY000", Message:"YEAR(%lu) column type is deprecated. Creating YEAR(4)column instead.", Description:"ER_INVALID_YEAR_COLUMN_LENGTH wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1819: {ErrorNumber:1819, ErrorType:"ServerError", Symbol:"ER_NOT_VALID_PASSWORD", SQLState:"HY000", Message:"Your password does not satisfy the current policyrequirements", Description:"ER_NOT_VALID_PASSWORD was added in5.6.6.", MySQLVersion:"5.6"},
    1820: {ErrorNumber:1820, ErrorType:"ServerError", Symbol:"ER_MUST_CHANGE_PASSWORD", SQLState:"HY000", Message:"You must SET PASSWORD before executing this statement", Description:"ER_MUST_CHANGE_PASSWORD was addedin 5.6.6.", MySQLVersion:"5.6"},
    1821: {ErrorNumber:1821, ErrorType:"ServerError", Symbol:"ER_FK_NO_INDEX_CHILD", SQLState:"HY000", Message:"Failed to add the foreign key constaint. Missing indexfor constraint '%s' in the foreign table '%s'", Description:"ER_FK_NO_INDEX_CHILD was added in5.6.6.", MySQLVersion:"5.6"},
    1822: {ErrorNumber:1822, ErrorType:"ServerError", Symbol:"ER_FK_NO_INDEX_PARENT", SQLState:"HY000", Message:"Failed to add the foreign key constaint. Missing indexfor constraint '%s' in the referenced table '%s'", Description:"ER_FK_NO_INDEX_PARENT was added in5.6.6.", MySQLVersion:"5.6"},
    1823: {ErrorNumber:1823, ErrorType:"ServerError", Symbol:"ER_FK_FAIL_ADD_SYSTEM", SQLState:"HY000", Message:"Failed to add the foreign key constraint '%s' to systemtables", Description:"ER_FK_FAIL_ADD_SYSTEM was added in5.6.6.", MySQLVersion:"5.6"},
    1824: {ErrorNumber:1824, ErrorType:"ServerError", Symbol:"ER_FK_CANNOT_OPEN_PARENT", SQLState:"HY000", Message:"Failed to open the referenced table '%s'", Description:"ER_FK_CANNOT_OPEN_PARENT was addedin 5.6.6.", MySQLVersion:"5.6"},
    1825: {ErrorNumber:1825, ErrorType:"ServerError", Symbol:"ER_FK_INCORRECT_OPTION", SQLState:"HY000", Message:"Failed to add the foreign key constraint on table '%s'.Incorrect options in FOREIGN KEY constraint '%s'", Description:"ER_FK_INCORRECT_OPTION was addedin 5.6.6.", MySQLVersion:"5.6"},
    1826: {ErrorNumber:1826, ErrorType:"ServerError", Symbol:"ER_FK_DUP_NAME", SQLState:"HY000", Message:"Duplicate foreign key constraint name '%s'", Description:"ER_FK_DUP_NAME was added in 5.6.6.", MySQLVersion:"5.6"},
    1827: {ErrorNumber:1827, ErrorType:"ServerError", Symbol:"ER_PASSWORD_FORMAT", SQLState:"HY000", Message:"The password hash doesn't have the expected format. Checkif the correct password algorithm is being used with thePASSWORD() function.", Description:"ER_PASSWORD_FORMAT was added in5.6.6.", MySQLVersion:"5.6"},
    1828: {ErrorNumber:1828, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_DROP", SQLState:"HY000", Message:"Cannot drop column '%s': needed in a foreign keyconstraint '%s'", Description:"ER_FK_COLUMN_CANNOT_DROP was addedin 5.6.6.", MySQLVersion:"5.6"},
    1829: {ErrorNumber:1829, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_DROP_CHILD", SQLState:"HY000", Message:"Cannot drop column '%s': needed in a foreign keyconstraint '%s' of table '%s'", Description:"ER_FK_COLUMN_CANNOT_DROP_CHILD wasadded in 5.6.6.", MySQLVersion:"5.6"},
    1830: {ErrorNumber:1830, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_NOT_NULL", SQLState:"HY000", Message:"Column '%s' cannot be NOT NULL: needed in a foreign keyconstraint '%s' SET NULL", Description:"ER_FK_COLUMN_NOT_NULL was added in5.6.6.", MySQLVersion:"5.6"},
    1831: {ErrorNumber:1831, ErrorType:"ServerError", Symbol:"ER_DUP_INDEX", SQLState:"HY000", Message:"Duplicate index '%s' defined on the table '%s.%s'. Thisis deprecated and will be disallowed in a future release.", Description:"ER_DUP_INDEX was added in 5.6.7.", MySQLVersion:"5.6"},
    1832: {ErrorNumber:1832, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_CHANGE", SQLState:"HY000", Message:"Cannot change column '%s': used in a foreign keyconstraint '%s'", Description:"ER_FK_COLUMN_CANNOT_CHANGE wasadded in 5.6.7.", MySQLVersion:"5.6"},
    1833: {ErrorNumber:1833, ErrorType:"ServerError", Symbol:"ER_FK_COLUMN_CANNOT_CHANGE_CHILD", SQLState:"HY000", Message:"Cannot change column '%s': used in a foreign keyconstraint '%s' of table '%s'", Description:"ER_FK_COLUMN_CANNOT_CHANGE_CHILDwas added in 5.6.7.", MySQLVersion:"5.6"},
    1834: {ErrorNumber:1834, ErrorType:"ServerError", Symbol:"ER_FK_CANNOT_DELETE_PARENT", SQLState:"HY000", Message:"Cannot delete rows from table which is parent in aforeign key constraint '%s' of table '%s'", Description:"ER_FK_CANNOT_DELETE_PARENT wasadded in 5.6.7.", MySQLVersion:"5.6"},
    1835: {ErrorNumber:1835, ErrorType:"ServerError", Symbol:"ER_MALFORMED_PACKET", SQLState:"HY000", Message:"Malformed communication packet.", Description:"ER_MALFORMED_PACKET was added in5.6.7.", MySQLVersion:"5.6"},
    1836: {ErrorNumber:1836, ErrorType:"ServerError", Symbol:"ER_READ_ONLY_MODE", SQLState:"HY000", Message:"Running in read-only mode", Description:"ER_READ_ONLY_MODE was added in5.6.8.", MySQLVersion:"5.6"},
    1837: {ErrorNumber:1837, ErrorType:"ServerError", Symbol:"ER_GTID_NEXT_TYPE_UNDEFINED_GROUP", SQLState:"HY000", Message:"When @@SESSION.GTID_NEXT is set to a GTID, you mustexplicitly set it to a different value after a COMMIT or ROLLBACK.Please check GTID_NEXT variable manual page for detailedexplanation. Current @@SESSION.GTID_NEXT is '%s'.", Description:"ER_GTID_NEXT_TYPE_UNDEFINED_GROUPwas added in 5.6.9.", MySQLVersion:"5.6"},
    1838: {ErrorNumber:1838, ErrorType:"ServerError", Symbol:"ER_VARIABLE_NOT_SETTABLE_IN_SP", SQLState:"HY000", Message:"The system variable %s cannot be set in storedprocedures.", Description:"ER_VARIABLE_NOT_SETTABLE_IN_SP wasadded in 5.6.9.", MySQLVersion:"5.6"},
    1839: {ErrorNumber:1839, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED can only be set when@@GLOBAL.GTID_MODE = ON.", Description:"ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFFwas added in 5.6.9.", MySQLVersion:"5.6"},
    1840: {ErrorNumber:1840, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED can only be set when@@GLOBAL.GTID_EXECUTED is empty.", Description:"ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTYwas added in 5.6.9.", MySQLVersion:"5.6"},
    1841: {ErrorNumber:1841, ErrorType:"ServerError", Symbol:"ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED can only be set when there are noongoing transactions (not even in other clients).", Description:"ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTYwas added in 5.6.9.", MySQLVersion:"5.6"},
    1842: {ErrorNumber:1842, ErrorType:"ServerError", Symbol:"ER_GTID_PURGED_WAS_CHANGED", SQLState:"HY000", Message:"@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'.", Description:"ER_GTID_PURGED_WAS_CHANGED wasadded in 5.6.9.", MySQLVersion:"5.6"},
    1843: {ErrorNumber:1843, ErrorType:"ServerError", Symbol:"ER_GTID_EXECUTED_WAS_CHANGED", SQLState:"HY000", Message:"@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'.", Description:"ER_GTID_EXECUTED_WAS_CHANGED wasadded in 5.6.9.", MySQLVersion:"5.6"},
    1844: {ErrorNumber:1844, ErrorType:"ServerError", Symbol:"ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES", SQLState:"HY000", Message:"Cannot execute statement: impossible to write to binarylog since BINLOG_FORMAT = STATEMENT, and both replicated and nonreplicated tables are written to.", Description:"ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLESwas added in 5.6.9.", MySQLVersion:"5.6"},
    1845: {ErrorNumber:1845, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED", SQLState:"0A000", Message:"%s is not supported for this operation. Try %s.", Description:"ER_ALTER_OPERATION_NOT_SUPPORTEDwas added in 5.6.10.", MySQLVersion:"5.6"},
    1846: {ErrorNumber:1846, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON", SQLState:"0A000", Message:"%s is not supported. Reason: %s. Try %s.", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASONwas added in 5.6.10.", MySQLVersion:"5.6"},
    1847: {ErrorNumber:1847, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY", SQLState:"HY000", Message:"COPY algorithm requires a lock", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPYwas added in 5.6.10.", MySQLVersion:"5.6"},
    1848: {ErrorNumber:1848, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION", SQLState:"HY000", Message:"Partition specific operations do not yet supportLOCK/ALGORITHM", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITIONwas added in 5.6.10.", MySQLVersion:"5.6"},
    1849: {ErrorNumber:1849, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME", SQLState:"HY000", Message:"Columns participating in a foreign key are renamed", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAMEwas added in 5.6.10.", MySQLVersion:"5.6"},
    1850: {ErrorNumber:1850, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE", SQLState:"HY000", Message:"Cannot change column type INPLACE", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPEwas added in 5.6.10.", MySQLVersion:"5.6"},
    1851: {ErrorNumber:1851, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK", SQLState:"HY000", Message:"Adding foreign keys needs foreign_key_checks=OFF", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECKwas added in 5.6.10.", MySQLVersion:"5.6"},
    1852: {ErrorNumber:1852, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE", SQLState:"HY000", Message:"Creating unique indexes with IGNORE requires COPYalgorithm to remove duplicate rows", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNOREwas added in 5.6.10.", MySQLVersion:"5.6"},
    1853: {ErrorNumber:1853, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK", SQLState:"HY000", Message:"Dropping a primary key is not allowed without also addinga new primary key", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPKwas added in 5.6.10.", MySQLVersion:"5.6"},
    1854: {ErrorNumber:1854, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC", SQLState:"HY000", Message:"Adding an auto-increment column requires a lock", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINCwas added in 5.6.10.", MySQLVersion:"5.6"},
    1855: {ErrorNumber:1855, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS", SQLState:"HY000", Message:"Cannot replace hidden FTS_DOC_ID with a user-visible one", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTSwas added in 5.6.10.", MySQLVersion:"5.6"},
    1856: {ErrorNumber:1856, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS", SQLState:"HY000", Message:"Cannot drop or rename FTS_DOC_ID", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTSwas added in 5.6.10.", MySQLVersion:"5.6"},
    1857: {ErrorNumber:1857, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS", SQLState:"HY000", Message:"Fulltext index creation requires a lock", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTSwas added in 5.6.10.", MySQLVersion:"5.6"},
    1858: {ErrorNumber:1858, ErrorType:"ServerError", Symbol:"ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE", SQLState:"HY000", Message:"sql_slave_skip_counter can not be set when the server isrunning with @@GLOBAL.GTID_MODE = ON. Instead, for eachtransaction that you want to skip, generate an empty transactionwith the same GTID as the transaction", Description:"ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODEwas added in 5.6.10.", MySQLVersion:"5.6"},
    1859: {ErrorNumber:1859, ErrorType:"ServerError", Symbol:"ER_DUP_UNKNOWN_IN_INDEX", SQLState:"23000", Message:"Duplicate entry for key '%s'", Description:"ER_DUP_UNKNOWN_IN_INDEX was addedin 5.6.10.", MySQLVersion:"5.6"},
    1860: {ErrorNumber:1860, ErrorType:"ServerError", Symbol:"ER_IDENT_CAUSES_TOO_LONG_PATH", SQLState:"HY000", Message:"Long database name and identifier for object resulted inpath length exceeding %d characters. Path: '%s'.", Description:"ER_IDENT_CAUSES_TOO_LONG_PATH wasadded in 5.6.10.", MySQLVersion:"5.6"},
    1861: {ErrorNumber:1861, ErrorType:"ServerError", Symbol:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL", SQLState:"HY000", Message:"cannot silently convert NULL values, as required in thisSQL_MODE", Description:"ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULLwas added in 5.6.10.", MySQLVersion:"5.6"},
    1862: {ErrorNumber:1862, ErrorType:"ServerError", Symbol:"ER_MUST_CHANGE_PASSWORD_LOGIN", SQLState:"HY000", Message:"Your password has expired. To log in you must change itusing a client that supports expired passwords.", Description:"ER_MUST_CHANGE_PASSWORD_LOGIN wasadded in 5.6.11.", MySQLVersion:"5.6"},
    1863: {ErrorNumber:1863, ErrorType:"ServerError", Symbol:"ER_ROW_IN_WRONG_PARTITION", SQLState:"HY000", Message:"Found a row in wrong partition %s", Description:"ER_ROW_IN_WRONG_PARTITION wasadded in 5.6.11.", MySQLVersion:"5.6"},
    1864: {ErrorNumber:1864, ErrorType:"ServerError", Symbol:"ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX", SQLState:"HY000", Message:"Cannot schedule event %s, relay-log name %s, position %sto Worker thread because its size %lu exceeds %lu ofslave_pending_jobs_size_max.", Description:"ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAXwas added in 5.6.12.", MySQLVersion:"5.6"},
    1865: {ErrorNumber:1865, ErrorType:"ServerError", Symbol:"ER_INNODB_NO_FT_USES_PARSER", SQLState:"HY000", Message:"Cannot CREATE FULLTEXT INDEX WITH PARSER on InnoDB table", Description:"ER_INNODB_NO_FT_USES_PARSER wasadded in 5.6.12.", MySQLVersion:"5.6"},
    1866: {ErrorNumber:1866, ErrorType:"ServerError", Symbol:"ER_BINLOG_LOGICAL_CORRUPTION", SQLState:"HY000", Message:"The binary log file '%s' is logically corrupted: %s", Description:"ER_BINLOG_LOGICAL_CORRUPTION wasadded in 5.6.12.", MySQLVersion:"5.6"},
    1867: {ErrorNumber:1867, ErrorType:"ServerError", Symbol:"ER_WARN_PURGE_LOG_IN_USE", SQLState:"HY000", Message:"file %s was not purged because it was being read by %dthread(s), purged only %d out of %d files.", Description:"ER_WARN_PURGE_LOG_IN_USE was addedin 5.6.12.", MySQLVersion:"5.6"},
    1868: {ErrorNumber:1868, ErrorType:"ServerError", Symbol:"ER_WARN_PURGE_LOG_IS_ACTIVE", SQLState:"HY000", Message:"file %s was not purged because it is the active log file.", Description:"ER_WARN_PURGE_LOG_IS_ACTIVE wasadded in 5.6.12.", MySQLVersion:"5.6"},
    1869: {ErrorNumber:1869, ErrorType:"ServerError", Symbol:"ER_AUTO_INCREMENT_CONFLICT", SQLState:"HY000", Message:"Auto-increment value in UPDATE conflicts with internallygenerated values", Description:"ER_AUTO_INCREMENT_CONFLICT wasadded in 5.6.12.", MySQLVersion:"5.6"},
    1870: {ErrorNumber:1870, ErrorType:"ServerError", Symbol:"WARN_ON_BLOCKHOLE_IN_RBR", SQLState:"HY000", Message:"Row events are not logged for %s statements that modifyBLACKHOLE tables in row format. Table(s): '%s'", Description:"WARN_ON_BLOCKHOLE_IN_RBR was addedin 5.6.12.", MySQLVersion:"5.6"},
    1871: {ErrorNumber:1871, ErrorType:"ServerError", Symbol:"ER_SLAVE_MI_INIT_REPOSITORY", SQLState:"HY000", Message:"Slave failed to initialize master info structure from therepository", Description:"ER_SLAVE_MI_INIT_REPOSITORY wasadded in 5.6.12.", MySQLVersion:"5.6"},
    1872: {ErrorNumber:1872, ErrorType:"ServerError", Symbol:"ER_SLAVE_RLI_INIT_REPOSITORY", SQLState:"HY000", Message:"Slave failed to initialize relay log info structure fromthe repository", Description:"ER_SLAVE_RLI_INIT_REPOSITORY wasadded in 5.6.12.", MySQLVersion:"5.6"},
    1873: {ErrorNumber:1873, ErrorType:"ServerError", Symbol:"ER_ACCESS_DENIED_CHANGE_USER_ERROR", SQLState:"28000", Message:"Access denied trying to change to user '%s'@'%s' (usingpassword: %s). Disconnecting.", Description:"ER_ACCESS_DENIED_CHANGE_USER_ERRORwas added in 5.6.13.", MySQLVersion:"5.6"},
    1874: {ErrorNumber:1874, ErrorType:"ServerError", Symbol:"ER_INNODB_READ_ONLY", SQLState:"HY000", Message:"InnoDB is in read only mode.", Description:"ER_INNODB_READ_ONLY was added in5.6.13.", MySQLVersion:"5.6"},
    1875: {ErrorNumber:1875, ErrorType:"ServerError", Symbol:"ER_STOP_SLAVE_SQL_THREAD_TIMEOUT", SQLState:"HY000", Message:"STOP SLAVE command execution is incomplete: Slave SQLthread got the stop signal, thread is busy, SQL thread will stoponce the current task is complete.", Description:"ER_STOP_SLAVE_SQL_THREAD_TIMEOUTwas added in 5.6.13.", MySQLVersion:"5.6"},
    1876: {ErrorNumber:1876, ErrorType:"ServerError", Symbol:"ER_STOP_SLAVE_IO_THREAD_TIMEOUT", SQLState:"HY000", Message:"STOP SLAVE command execution is incomplete: Slave IOthread got the stop signal, thread is busy, IO thread will stoponce the current task is complete.", Description:"ER_STOP_SLAVE_IO_THREAD_TIMEOUTwas added in 5.6.13.", MySQLVersion:"5.6"},
    1877: {ErrorNumber:1877, ErrorType:"ServerError", Symbol:"ER_TABLE_CORRUPT", SQLState:"HY000", Message:"Operation cannot be performed. The table '%s.%s' ismissing, corrupt or contains bad data.", Description:"ER_TABLE_CORRUPT was added in5.6.14.", MySQLVersion:"5.6"},
    1878: {ErrorNumber:1878, ErrorType:"ServerError", Symbol:"ER_TEMP_FILE_WRITE_FAILURE", SQLState:"HY000", Message:"Temporary file write failure.", Description:"ER_TEMP_FILE_WRITE_FAILURE wasadded in 5.6.15.", MySQLVersion:"5.6"},
    1879: {ErrorNumber:1879, ErrorType:"ServerError", Symbol:"ER_INNODB_FT_AUX_NOT_HEX_ID", SQLState:"HY000", Message:"Upgrade index name failed, please use create index(altertable) algorithm copy to rebuild index.", Description:"ER_INNODB_FT_AUX_NOT_HEX_ID wasadded in 5.6.16.", MySQLVersion:"5.6"},
    1880: {ErrorNumber:1880, ErrorType:"ServerError", Symbol:"ER_OLD_TEMPORALS_UPGRADED", SQLState:"HY000", Message:"TIME/TIMESTAMP/DATETIME columns of old format have beenupgraded to the new format.", Description:"ER_OLD_TEMPORALS_UPGRADED wasadded in 5.6.16.", MySQLVersion:"5.6"},
    1881: {ErrorNumber:1881, ErrorType:"ServerError", Symbol:"ER_INNODB_FORCED_RECOVERY", SQLState:"HY000", Message:"Operation not allowed when innodb_forced_recovery > 0.", Description:"ER_INNODB_FORCED_RECOVERY wasadded in 5.6.16.", MySQLVersion:"5.6"},
    1882: {ErrorNumber:1882, ErrorType:"ServerError", Symbol:"ER_AES_INVALID_IV", SQLState:"HY000", Message:"The initialization vector supplied to %s is too short.Must be at least %d bytes long", Description:"ER_AES_INVALID_IV was added in5.6.17.", MySQLVersion:"5.6"},
    1883: {ErrorNumber:1883, ErrorType:"ServerError", Symbol:"ER_PLUGIN_CANNOT_BE_UNINSTALLED", SQLState:"HY000", Message:"Plugin '%s' cannot be uninstalled now. %s", Description:"ER_PLUGIN_CANNOT_BE_UNINSTALLEDwas added in 5.6.20.", MySQLVersion:"5.6"},
    1884: {ErrorNumber:1884, ErrorType:"ServerError", Symbol:"ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP", SQLState:"HY000", Message:"Cannot execute statement because it needs to be writtento the binary log as multiple statements, and this is not allowedwhen @@SESSION.GTID_NEXT == 'UUID:NUMBER'.", Description:"ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUPwas added in 5.6.20.", MySQLVersion:"5.6"},
    1885: {ErrorNumber:1885, ErrorType:"ServerError", Symbol:"ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER", SQLState:"HY000", Message:"Slave has more GTIDs than the master has, using themaster's SERVER_UUID. This may indicate that the end of the binarylog was truncated or that the last binary log file was lost, e.g.,after a power or disk failure when sync_binlog != 1. The mastermay or may not have rolled back transactions that were alreadyreplicated to the slave. Suggest to replicate any transactionsthat master has rolled back from slave to master, and/or commitempty transactions on master to account for transactions that havebeen committed on master but are not included in GTID_EXECUTED.", Description:"ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTERwas added in 5.6.23.", MySQLVersion:"5.6"},
    1886: {ErrorNumber:1886, ErrorType:"ServerError", Symbol:"ER_MISSING_KEY", SQLState:"HY000", Message:"The table '%s.%s' does not have the necessary key(s)defined on it. Please check the table definition and createindex(s) accordingly.", Description:"ER_MISSING_KEY was added in5.6.40.", MySQLVersion:"5.6"},
    1887: {ErrorNumber:1887, ErrorType:"ServerError", Symbol:"WARN_NAMED_PIPE_ACCESS_EVERYONE", SQLState:"HY000", Message:"Setting named_pipe_full_access_group='%s' is insecure.Consider using a Windows group with fewer members.", Description:"WARN_NAMED_PIPE_ACCESS_EVERYONEwas added in 5.6.45.", MySQLVersion:"5.6"},
    1888: {ErrorNumber:1888, ErrorType:"ServerError", Symbol:"ER_FOUND_MISSING_GTIDS", SQLState:"HY000", Message:"Cannot replicate to server with server_uuid='%s' becausethe present server has purged required binary logs. The connectingserver needs to replicate the missing transactions from elsewhere,or be replaced by a new server created from a more recent backup.To prevent this error in the future, consider increasing thebinary log expiration period on the present server. %s.", Description:"ER_FOUND_MISSING_GTIDS was addedin 5.6.47.", MySQLVersion:"5.6"},
    2000: {ErrorNumber:2000, ErrorType:"ClientError", Symbol:"CR_UNKNOWN_ERROR", SQLState:"", Message:"Unknown MySQL error", Description:"", MySQLVersion:"5.6"},
    2001: {ErrorNumber:2001, ErrorType:"ClientError", Symbol:"CR_SOCKET_CREATE_ERROR", SQLState:"", Message:"Can't create UNIX socket (%d)", Description:"", MySQLVersion:"5.6"},
    2002: {ErrorNumber:2002, ErrorType:"ClientError", Symbol:"CR_CONNECTION_ERROR", SQLState:"", Message:"Can't connect to local MySQL server through socket '%s'(%d)", Description:"", MySQLVersion:"5.6"},
    2003: {ErrorNumber:2003, ErrorType:"ClientError", Symbol:"CR_CONN_HOST_ERROR", SQLState:"", Message:"Can't connect to MySQL server on '%s' (%d)", Description:"", MySQLVersion:"5.6"},
    2004: {ErrorNumber:2004, ErrorType:"ClientError", Symbol:"CR_IPSOCK_ERROR", SQLState:"", Message:"Can't create TCP/IP socket (%d)", Description:"", MySQLVersion:"5.6"},
    2005: {ErrorNumber:2005, ErrorType:"ClientError", Symbol:"CR_UNKNOWN_HOST", SQLState:"", Message:"Unknown MySQL server host '%s' (%d)", Description:"", MySQLVersion:"5.6"},
    2006: {ErrorNumber:2006, ErrorType:"ClientError", Symbol:"CR_SERVER_GONE_ERROR", SQLState:"", Message:"MySQL server has gone away", Description:"", MySQLVersion:"5.6"},
    2007: {ErrorNumber:2007, ErrorType:"ClientError", Symbol:"CR_VERSION_ERROR", SQLState:"", Message:"Protocol mismatch", Description:"server version = %d, client version =%d", MySQLVersion:"5.6"},
    2008: {ErrorNumber:2008, ErrorType:"ClientError", Symbol:"CR_OUT_OF_MEMORY", SQLState:"", Message:"MySQL client ran out of memory", Description:"", MySQLVersion:"5.6"},
    2009: {ErrorNumber:2009, ErrorType:"ClientError", Symbol:"CR_WRONG_HOST_INFO", SQLState:"", Message:"Wrong host info", Description:"", MySQLVersion:"5.6"},
    2010: {ErrorNumber:2010, ErrorType:"ClientError", Symbol:"CR_LOCALHOST_CONNECTION", SQLState:"", Message:"Localhost via UNIX socket", Description:"", MySQLVersion:"5.6"},
    2011: {ErrorNumber:2011, ErrorType:"ClientError", Symbol:"CR_TCP_CONNECTION", SQLState:"", Message:"%s via TCP/IP", Description:"", MySQLVersion:"5.6"},
    2012: {ErrorNumber:2012, ErrorType:"ClientError", Symbol:"CR_SERVER_HANDSHAKE_ERR", SQLState:"", Message:"Error in server handshake", Description:"", MySQLVersion:"5.6"},
    2013: {ErrorNumber:2013, ErrorType:"ClientError", Symbol:"CR_SERVER_LOST", SQLState:"", Message:"Lost connection to MySQL server during query", Description:"", MySQLVersion:"5.6"},
    2014: {ErrorNumber:2014, ErrorType:"ClientError", Symbol:"CR_COMMANDS_OUT_OF_SYNC", SQLState:"", Message:"Commands out of sync", Description:"Commands were executed in an improper order. This error occurswhen a function is called that is not appropriate for the currentstate of the connection. For example, ifmysql_stmt_fetch() is not calledenough times to read an entire result set (that is, enough timesto return MYSQL_NO_DATA), this error may occurfor the following C API call.", MySQLVersion:"5.6"},
    2015: {ErrorNumber:2015, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPE_CONNECTION", SQLState:"", Message:"Named pipe: %s", Description:"", MySQLVersion:"5.6"},
    2016: {ErrorNumber:2016, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPEWAIT_ERROR", SQLState:"", Message:"Can't wait for named pipe to host: %s pipe: %s (%lu)", Description:"", MySQLVersion:"5.6"},
    2017: {ErrorNumber:2017, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPEOPEN_ERROR", SQLState:"", Message:"Can't open named pipe to host: %s pipe: %s (%lu)", Description:"", MySQLVersion:"5.6"},
    2018: {ErrorNumber:2018, ErrorType:"ClientError", Symbol:"CR_NAMEDPIPESETSTATE_ERROR", SQLState:"", Message:"Can't set state of named pipe to host: %s pipe: %s (%lu)", Description:"", MySQLVersion:"5.6"},
    2019: {ErrorNumber:2019, ErrorType:"ClientError", Symbol:"CR_CANT_READ_CHARSET", SQLState:"", Message:"Can't initialize character set %s (path: %s)", Description:"", MySQLVersion:"5.6"},
    2020: {ErrorNumber:2020, ErrorType:"ClientError", Symbol:"CR_NET_PACKET_TOO_LARGE", SQLState:"", Message:"Got packet bigger than 'max_allowed_packet' bytes", Description:"", MySQLVersion:"5.6"},
    2021: {ErrorNumber:2021, ErrorType:"ClientError", Symbol:"CR_EMBEDDED_CONNECTION", SQLState:"", Message:"Embedded server", Description:"", MySQLVersion:"5.6"},
    2022: {ErrorNumber:2022, ErrorType:"ClientError", Symbol:"CR_PROBE_SLAVE_STATUS", SQLState:"", Message:"Error on SHOW SLAVE STATUS:", Description:"", MySQLVersion:"5.6"},
    2023: {ErrorNumber:2023, ErrorType:"ClientError", Symbol:"CR_PROBE_SLAVE_HOSTS", SQLState:"", Message:"Error on SHOW SLAVE HOSTS:", Description:"", MySQLVersion:"5.6"},
    2024: {ErrorNumber:2024, ErrorType:"ClientError", Symbol:"CR_PROBE_SLAVE_CONNECT", SQLState:"", Message:"Error connecting to slave:", Description:"", MySQLVersion:"5.6"},
    2025: {ErrorNumber:2025, ErrorType:"ClientError", Symbol:"CR_PROBE_MASTER_CONNECT", SQLState:"", Message:"Error connecting to master:", Description:"", MySQLVersion:"5.6"},
    2026: {ErrorNumber:2026, ErrorType:"ClientError", Symbol:"CR_SSL_CONNECTION_ERROR", SQLState:"", Message:"SSL connection error: %s", Description:"", MySQLVersion:"5.6"},
    2027: {ErrorNumber:2027, ErrorType:"ClientError", Symbol:"CR_MALFORMED_PACKET", SQLState:"", Message:"Malformed packet", Description:"", MySQLVersion:"5.6"},
    2028: {ErrorNumber:2028, ErrorType:"ClientError", Symbol:"CR_WRONG_LICENSE", SQLState:"", Message:"This client library is licensed only for use with MySQLservers having '%s' license", Description:"", MySQLVersion:"5.6"},
    2029: {ErrorNumber:2029, ErrorType:"ClientError", Symbol:"CR_NULL_POINTER", SQLState:"", Message:"Invalid use of null pointer", Description:"", MySQLVersion:"5.6"},
    2030: {ErrorNumber:2030, ErrorType:"ClientError", Symbol:"CR_NO_PREPARE_STMT", SQLState:"", Message:"Statement not prepared", Description:"", MySQLVersion:"5.6"},
    2031: {ErrorNumber:2031, ErrorType:"ClientError", Symbol:"CR_PARAMS_NOT_BOUND", SQLState:"", Message:"No data supplied for parameters in prepared statement", Description:"", MySQLVersion:"5.6"},
    2032: {ErrorNumber:2032, ErrorType:"ClientError", Symbol:"CR_DATA_TRUNCATED", SQLState:"", Message:"Data truncated", Description:"", MySQLVersion:"5.6"},
    2033: {ErrorNumber:2033, ErrorType:"ClientError", Symbol:"CR_NO_PARAMETERS_EXISTS", SQLState:"", Message:"No parameters exist in the statement", Description:"", MySQLVersion:"5.6"},
    2034: {ErrorNumber:2034, ErrorType:"ClientError", Symbol:"CR_INVALID_PARAMETER_NO", SQLState:"", Message:"Invalid parameter number", Description:"From 5.6.6: A key name was empty or the amount of connectionattribute data formysql_options4() exceeds the 64KBlimit.", MySQLVersion:"5.6"},
    2035: {ErrorNumber:2035, ErrorType:"ClientError", Symbol:"CR_INVALID_BUFFER_USE", SQLState:"", Message:"Can't send long data for non-string/non-binary data types(parameter: %d)", Description:"", MySQLVersion:"5.6"},
    2036: {ErrorNumber:2036, ErrorType:"ClientError", Symbol:"CR_UNSUPPORTED_PARAM_TYPE", SQLState:"", Message:"Using unsupported buffer type: %d (parameter: %d)", Description:"", MySQLVersion:"5.6"},
    2037: {ErrorNumber:2037, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECTION", SQLState:"", Message:"Shared memory: %s", Description:"", MySQLVersion:"5.6"},
    2038: {ErrorNumber:2038, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_REQUEST_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not create requestevent (%lu)", MySQLVersion:"5.6"},
    2039: {ErrorNumber:2039, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_ANSWER_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"no answer event received fromserver (%lu)", MySQLVersion:"5.6"},
    2040: {ErrorNumber:2040, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_FILE_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"server could not allocate filemapping (%lu)", MySQLVersion:"5.6"},
    2041: {ErrorNumber:2041, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"server could not get pointer tofile mapping (%lu)", MySQLVersion:"5.6"},
    2042: {ErrorNumber:2042, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_FILE_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not allocate filemapping (%lu)", MySQLVersion:"5.6"},
    2043: {ErrorNumber:2043, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_MAP_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not get pointer tofile mapping (%lu)", MySQLVersion:"5.6"},
    2044: {ErrorNumber:2044, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_EVENT_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"client could not create %sevent (%lu)", MySQLVersion:"5.6"},
    2045: {ErrorNumber:2045, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_ABANDONED_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"no answer from server (%lu)", MySQLVersion:"5.6"},
    2046: {ErrorNumber:2046, ErrorType:"ClientError", Symbol:"CR_SHARED_MEMORY_CONNECT_SET_ERROR", SQLState:"", Message:"Can't open shared memory", Description:"cannot send request event toserver (%lu)", MySQLVersion:"5.6"},
    2047: {ErrorNumber:2047, ErrorType:"ClientError", Symbol:"CR_CONN_UNKNOW_PROTOCOL", SQLState:"", Message:"Wrong or unknown protocol", Description:"", MySQLVersion:"5.6"},
    2048: {ErrorNumber:2048, ErrorType:"ClientError", Symbol:"CR_INVALID_CONN_HANDLE", SQLState:"", Message:"Invalid connection handle", Description:"", MySQLVersion:"5.6"},
    2049: {ErrorNumber:2049, ErrorType:"ClientError", Symbol:"CR_SECURE_AUTH", SQLState:"", Message:"Connection using old (pre-4.1.1) authentication protocolrefused (client option 'secure_auth' enabled)", Description:"", MySQLVersion:"5.6"},
    2050: {ErrorNumber:2050, ErrorType:"ClientError", Symbol:"CR_FETCH_CANCELED", SQLState:"", Message:"Row retrieval was canceled by mysql_stmt_close() call", Description:"", MySQLVersion:"5.6"},
    2051: {ErrorNumber:2051, ErrorType:"ClientError", Symbol:"CR_NO_DATA", SQLState:"", Message:"Attempt to read column without prior row fetch", Description:"", MySQLVersion:"5.6"},
    2052: {ErrorNumber:2052, ErrorType:"ClientError", Symbol:"CR_NO_STMT_METADATA", SQLState:"", Message:"Prepared statement contains no metadata", Description:"", MySQLVersion:"5.6"},
    2053: {ErrorNumber:2053, ErrorType:"ClientError", Symbol:"CR_NO_RESULT_SET", SQLState:"", Message:"Attempt to read a row while there is no result setassociated with the statement", Description:"", MySQLVersion:"5.6"},
    2054: {ErrorNumber:2054, ErrorType:"ClientError", Symbol:"CR_NOT_IMPLEMENTED", SQLState:"", Message:"This feature is not implemented yet", Description:"", MySQLVersion:"5.6"},
    2055: {ErrorNumber:2055, ErrorType:"ClientError", Symbol:"CR_SERVER_LOST_EXTENDED", SQLState:"", Message:"Lost connection to MySQL server at '%s', system error: %d", Description:"", MySQLVersion:"5.6"},
    2056: {ErrorNumber:2056, ErrorType:"ClientError", Symbol:"CR_STMT_CLOSED", SQLState:"", Message:"Statement closed indirectly because of a preceeding %s()call", Description:"", MySQLVersion:"5.6"},
    2057: {ErrorNumber:2057, ErrorType:"ClientError", Symbol:"CR_NEW_STMT_METADATA", SQLState:"", Message:"The number of columns in the result set differs from thenumber of bound buffers. You must reset the statement, rebind theresult set columns, and execute the statement again", Description:"", MySQLVersion:"5.6"},
    2058: {ErrorNumber:2058, ErrorType:"ClientError", Symbol:"CR_ALREADY_CONNECTED", SQLState:"", Message:"This handle is already connected. Use a separate handlefor each connection.", Description:"", MySQLVersion:"5.6"},
    2059: {ErrorNumber:2059, ErrorType:"ClientError", Symbol:"CR_AUTH_PLUGIN_CANNOT_LOAD", SQLState:"", Message:"Authentication plugin '%s' cannot be loaded: %s", Description:"CR_AUTH_PLUGIN_CANNOT_LOAD wasadded in 5.6.1.", MySQLVersion:"5.6"},
    2060: {ErrorNumber:2060, ErrorType:"ClientError", Symbol:"CR_DUPLICATE_CONNECTION_ATTR", SQLState:"", Message:"There is an attribute with the same name already", Description:"CR_DUPLICATE_CONNECTION_ATTR wasadded in 5.6.6.", MySQLVersion:"5.6"},
    2061: {ErrorNumber:2061, ErrorType:"ClientError", Symbol:"CR_AUTH_PLUGIN_ERR", SQLState:"", Message:"Authentication plugin '%s' reported error: %s", Description:"CR_AUTH_PLUGIN_ERR was added in5.6.10.", MySQLVersion:"5.6"},
    1: {ErrorNumber:1, ErrorType:"GlobalError", Symbol:"EE_CANTCREATEFILE", SQLState:"", Message:"Can't create/write to file '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    2: {ErrorNumber:2, ErrorType:"GlobalError", Symbol:"EE_READ", SQLState:"", Message:"Error reading file '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    3: {ErrorNumber:3, ErrorType:"GlobalError", Symbol:"EE_WRITE", SQLState:"", Message:"Error writing file '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    4: {ErrorNumber:4, ErrorType:"GlobalError", Symbol:"EE_BADCLOSE", SQLState:"", Message:"Error on close of '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    5: {ErrorNumber:5, ErrorType:"GlobalError", Symbol:"EE_OUTOFMEMORY", SQLState:"", Message:"Out of memory (Needed %u bytes)", Description:"", MySQLVersion:"5.6"},
    6: {ErrorNumber:6, ErrorType:"GlobalError", Symbol:"EE_DELETE", SQLState:"", Message:"Error on delete of '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    7: {ErrorNumber:7, ErrorType:"GlobalError", Symbol:"EE_LINK", SQLState:"", Message:"Error on rename of '%s' to '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    9: {ErrorNumber:9, ErrorType:"GlobalError", Symbol:"EE_EOFERR", SQLState:"", Message:"Unexpected EOF found when reading file '%s' (Errcode: %d- %s)", Description:"", MySQLVersion:"5.6"},
    10: {ErrorNumber:10, ErrorType:"GlobalError", Symbol:"EE_CANTLOCK", SQLState:"", Message:"Can't lock file (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    11: {ErrorNumber:11, ErrorType:"GlobalError", Symbol:"EE_CANTUNLOCK", SQLState:"", Message:"Can't unlock file (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    12: {ErrorNumber:12, ErrorType:"GlobalError", Symbol:"EE_DIR", SQLState:"", Message:"Can't read dir of '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    13: {ErrorNumber:13, ErrorType:"GlobalError", Symbol:"EE_STAT", SQLState:"", Message:"Can't get stat of '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    14: {ErrorNumber:14, ErrorType:"GlobalError", Symbol:"EE_CANT_CHSIZE", SQLState:"", Message:"Can't change size of file (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    15: {ErrorNumber:15, ErrorType:"GlobalError", Symbol:"EE_CANT_OPEN_STREAM", SQLState:"", Message:"Can't open stream from handle (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    16: {ErrorNumber:16, ErrorType:"GlobalError", Symbol:"EE_GETWD", SQLState:"", Message:"Can't get working directory (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    17: {ErrorNumber:17, ErrorType:"GlobalError", Symbol:"EE_SETWD", SQLState:"", Message:"Can't change dir to '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    18: {ErrorNumber:18, ErrorType:"GlobalError", Symbol:"EE_LINK_WARNING", SQLState:"", Message:"Warning: '%s' had %d links", Description:"", MySQLVersion:"5.6"},
    19: {ErrorNumber:19, ErrorType:"GlobalError", Symbol:"EE_OPEN_WARNING", SQLState:"", Message:"Warning: %d files and %d streams is left open", Description:"", MySQLVersion:"5.6"},
    20: {ErrorNumber:20, ErrorType:"GlobalError", Symbol:"EE_DISK_FULL", SQLState:"", Message:"Disk is full writing '%s' (Errcode: %d - %s). Waiting forsomeone to free space...", Description:"", MySQLVersion:"5.6"},
    21: {ErrorNumber:21, ErrorType:"GlobalError", Symbol:"EE_CANT_MKDIR", SQLState:"", Message:"Can't create directory '%s' (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    22: {ErrorNumber:22, ErrorType:"GlobalError", Symbol:"EE_UNKNOWN_CHARSET", SQLState:"", Message:"Character set '%s' is not a compiled character set and isnot specified in the '%s' file", Description:"", MySQLVersion:"5.6"},
    23: {ErrorNumber:23, ErrorType:"GlobalError", Symbol:"EE_OUT_OF_FILERESOURCES", SQLState:"", Message:"Out of resources when opening file '%s' (Errcode: %d -%s)", Description:"", MySQLVersion:"5.6"},
    24: {ErrorNumber:24, ErrorType:"GlobalError", Symbol:"EE_CANT_READLINK", SQLState:"", Message:"Can't read value for symlink '%s' (Error %d - %s)", Description:"", MySQLVersion:"5.6"},
    25: {ErrorNumber:25, ErrorType:"GlobalError", Symbol:"EE_CANT_SYMLINK", SQLState:"", Message:"Can't create symlink '%s' pointing at '%s' (Error %d -%s)", Description:"", MySQLVersion:"5.6"},
    26: {ErrorNumber:26, ErrorType:"GlobalError", Symbol:"EE_REALPATH", SQLState:"", Message:"Error on realpath() on '%s' (Error %d - %s)", Description:"", MySQLVersion:"5.6"},
    27: {ErrorNumber:27, ErrorType:"GlobalError", Symbol:"EE_SYNC", SQLState:"", Message:"Can't sync file '%s' to disk (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    28: {ErrorNumber:28, ErrorType:"GlobalError", Symbol:"EE_UNKNOWN_COLLATION", SQLState:"", Message:"Collation '%s' is not a compiled collation and is notspecified in the '%s' file", Description:"", MySQLVersion:"5.6"},
    29: {ErrorNumber:29, ErrorType:"GlobalError", Symbol:"EE_FILENOTFOUND", SQLState:"", Message:"File '%s' not found (Errcode: %d - %s)", Description:"", MySQLVersion:"5.6"},
    30: {ErrorNumber:30, ErrorType:"GlobalError", Symbol:"EE_FILE_NOT_CLOSED", SQLState:"", Message:"File '%s' (fileno: %d) was not closed", Description:"", MySQLVersion:"5.6"},
    31: {ErrorNumber:31, ErrorType:"GlobalError", Symbol:"EE_CHANGE_OWNERSHIP", SQLState:"", Message:"Can't change ownership of the file '%s' (Errcode: %d -%s)", Description:"", MySQLVersion:"5.6"},
    32: {ErrorNumber:32, ErrorType:"GlobalError", Symbol:"EE_CHANGE_PERMISSIONS", SQLState:"", Message:"Can't change permissions of the file '%s' (Errcode: %d -%s)", Description:"", MySQLVersion:"5.6"},
    33: {ErrorNumber:33, ErrorType:"GlobalError", Symbol:"EE_CANT_SEEK", SQLState:"", Message:"Can't seek in file '%s' (Errcode: %d - %s)", Description:"EE_CANT_SEEK was added in 5.6.1.", MySQLVersion:"5.6"},
}

const (
ER_HASHCHK int64 = 1000
    ER_NISAMCHK int64 = 1001
    ER_NO int64 = 1002
    ER_YES int64 = 1003
    ER_CANT_CREATE_FILE int64 = 1004
    ER_CANT_CREATE_TABLE int64 = 1005
    ER_CANT_CREATE_DB int64 = 1006
    ER_DB_CREATE_EXISTS int64 = 1007
    ER_DB_DROP_EXISTS int64 = 1008
    ER_DB_DROP_DELETE int64 = 1009
    ER_DB_DROP_RMDIR int64 = 1010
    ER_CANT_DELETE_FILE int64 = 1011
    ER_CANT_FIND_SYSTEM_REC int64 = 1012
    ER_CANT_GET_STAT int64 = 1013
    ER_CANT_GET_WD int64 = 1014
    ER_CANT_LOCK int64 = 1015
    ER_CANT_OPEN_FILE int64 = 1016
    ER_FILE_NOT_FOUND int64 = 1017
    ER_CANT_READ_DIR int64 = 1018
    ER_CANT_SET_WD int64 = 1019
    ER_CHECKREAD int64 = 1020
    ER_DISK_FULL int64 = 1021
    ER_DUP_KEY int64 = 1022
    ER_ERROR_ON_CLOSE int64 = 1023
    ER_ERROR_ON_READ int64 = 1024
    ER_ERROR_ON_RENAME int64 = 1025
    ER_ERROR_ON_WRITE int64 = 1026
    ER_FILE_USED int64 = 1027
    ER_FILSORT_ABORT int64 = 1028
    ER_FORM_NOT_FOUND int64 = 1029
    ER_GET_ERRNO int64 = 1030
    ER_ILLEGAL_HA int64 = 1031
    ER_KEY_NOT_FOUND int64 = 1032
    ER_NOT_FORM_FILE int64 = 1033
    ER_NOT_KEYFILE int64 = 1034
    ER_OLD_KEYFILE int64 = 1035
    ER_OPEN_AS_READONLY int64 = 1036
    ER_OUTOFMEMORY int64 = 1037
    ER_OUT_OF_SORTMEMORY int64 = 1038
    ER_UNEXPECTED_EOF int64 = 1039
    ER_CON_COUNT_ERROR int64 = 1040
    ER_OUT_OF_RESOURCES int64 = 1041
    ER_BAD_HOST_ERROR int64 = 1042
    ER_HANDSHAKE_ERROR int64 = 1043
    ER_DBACCESS_DENIED_ERROR int64 = 1044
    ER_ACCESS_DENIED_ERROR int64 = 1045
    ER_NO_DB_ERROR int64 = 1046
    ER_UNKNOWN_COM_ERROR int64 = 1047
    ER_BAD_NULL_ERROR int64 = 1048
    ER_BAD_DB_ERROR int64 = 1049
    ER_TABLE_EXISTS_ERROR int64 = 1050
    ER_BAD_TABLE_ERROR int64 = 1051
    ER_NON_UNIQ_ERROR int64 = 1052
    ER_SERVER_SHUTDOWN int64 = 1053
    ER_BAD_FIELD_ERROR int64 = 1054
    ER_WRONG_FIELD_WITH_GROUP int64 = 1055
    ER_WRONG_GROUP_FIELD int64 = 1056
    ER_WRONG_SUM_SELECT int64 = 1057
    ER_WRONG_VALUE_COUNT int64 = 1058
    ER_TOO_LONG_IDENT int64 = 1059
    ER_DUP_FIELDNAME int64 = 1060
    ER_DUP_KEYNAME int64 = 1061
    ER_DUP_ENTRY int64 = 1062
    ER_WRONG_FIELD_SPEC int64 = 1063
    ER_PARSE_ERROR int64 = 1064
    ER_EMPTY_QUERY int64 = 1065
    ER_NONUNIQ_TABLE int64 = 1066
    ER_INVALID_DEFAULT int64 = 1067
    ER_MULTIPLE_PRI_KEY int64 = 1068
    ER_TOO_MANY_KEYS int64 = 1069
    ER_TOO_MANY_KEY_PARTS int64 = 1070
    ER_TOO_LONG_KEY int64 = 1071
    ER_KEY_COLUMN_DOES_NOT_EXITS int64 = 1072
    ER_BLOB_USED_AS_KEY int64 = 1073
    ER_TOO_BIG_FIELDLENGTH int64 = 1074
    ER_WRONG_AUTO_KEY int64 = 1075
    ER_READY int64 = 1076
    ER_NORMAL_SHUTDOWN int64 = 1077
    ER_GOT_SIGNAL int64 = 1078
    ER_SHUTDOWN_COMPLETE int64 = 1079
    ER_FORCING_CLOSE int64 = 1080
    ER_IPSOCK_ERROR int64 = 1081
    ER_NO_SUCH_INDEX int64 = 1082
    ER_WRONG_FIELD_TERMINATORS int64 = 1083
    ER_BLOBS_AND_NO_TERMINATED int64 = 1084
    ER_TEXTFILE_NOT_READABLE int64 = 1085
    ER_FILE_EXISTS_ERROR int64 = 1086
    ER_LOAD_INFO int64 = 1087
    ER_ALTER_INFO int64 = 1088
    ER_WRONG_SUB_KEY int64 = 1089
    ER_CANT_REMOVE_ALL_FIELDS int64 = 1090
    ER_CANT_DROP_FIELD_OR_KEY int64 = 1091
    ER_INSERT_INFO int64 = 1092
    ER_UPDATE_TABLE_USED int64 = 1093
    ER_NO_SUCH_THREAD int64 = 1094
    ER_KILL_DENIED_ERROR int64 = 1095
    ER_NO_TABLES_USED int64 = 1096
    ER_TOO_BIG_SET int64 = 1097
    ER_NO_UNIQUE_LOGFILE int64 = 1098
    ER_TABLE_NOT_LOCKED_FOR_WRITE int64 = 1099
    ER_TABLE_NOT_LOCKED int64 = 1100
    ER_BLOB_CANT_HAVE_DEFAULT int64 = 1101
    ER_WRONG_DB_NAME int64 = 1102
    ER_WRONG_TABLE_NAME int64 = 1103
    ER_TOO_BIG_SELECT int64 = 1104
    ER_UNKNOWN_ERROR int64 = 1105
    ER_UNKNOWN_PROCEDURE int64 = 1106
    ER_WRONG_PARAMCOUNT_TO_PROCEDURE int64 = 1107
    ER_WRONG_PARAMETERS_TO_PROCEDURE int64 = 1108
    ER_UNKNOWN_TABLE int64 = 1109
    ER_FIELD_SPECIFIED_TWICE int64 = 1110
    ER_INVALID_GROUP_FUNC_USE int64 = 1111
    ER_UNSUPPORTED_EXTENSION int64 = 1112
    ER_TABLE_MUST_HAVE_COLUMNS int64 = 1113
    ER_RECORD_FILE_FULL int64 = 1114
    ER_UNKNOWN_CHARACTER_SET int64 = 1115
    ER_TOO_MANY_TABLES int64 = 1116
    ER_TOO_MANY_FIELDS int64 = 1117
    ER_TOO_BIG_ROWSIZE int64 = 1118
    ER_STACK_OVERRUN int64 = 1119
    ER_WRONG_OUTER_JOIN int64 = 1120
    ER_NULL_COLUMN_IN_INDEX int64 = 1121
    ER_CANT_FIND_UDF int64 = 1122
    ER_CANT_INITIALIZE_UDF int64 = 1123
    ER_UDF_NO_PATHS int64 = 1124
    ER_UDF_EXISTS int64 = 1125
    ER_CANT_OPEN_LIBRARY int64 = 1126
    ER_CANT_FIND_DL_ENTRY int64 = 1127
    ER_FUNCTION_NOT_DEFINED int64 = 1128
    ER_HOST_IS_BLOCKED int64 = 1129
    ER_HOST_NOT_PRIVILEGED int64 = 1130
    ER_PASSWORD_ANONYMOUS_USER int64 = 1131
    ER_PASSWORD_NOT_ALLOWED int64 = 1132
    ER_PASSWORD_NO_MATCH int64 = 1133
    ER_UPDATE_INFO int64 = 1134
    ER_CANT_CREATE_THREAD int64 = 1135
    ER_WRONG_VALUE_COUNT_ON_ROW int64 = 1136
    ER_CANT_REOPEN_TABLE int64 = 1137
    ER_INVALID_USE_OF_NULL int64 = 1138
    ER_REGEXP_ERROR int64 = 1139
    ER_MIX_OF_GROUP_FUNC_AND_FIELDS int64 = 1140
    ER_NONEXISTING_GRANT int64 = 1141
    ER_TABLEACCESS_DENIED_ERROR int64 = 1142
    ER_COLUMNACCESS_DENIED_ERROR int64 = 1143
    ER_ILLEGAL_GRANT_FOR_TABLE int64 = 1144
    ER_GRANT_WRONG_HOST_OR_USER int64 = 1145
    ER_NO_SUCH_TABLE int64 = 1146
    ER_NONEXISTING_TABLE_GRANT int64 = 1147
    ER_NOT_ALLOWED_COMMAND int64 = 1148
    ER_SYNTAX_ERROR int64 = 1149
    ER_DELAYED_CANT_CHANGE_LOCK int64 = 1150
    ER_TOO_MANY_DELAYED_THREADS int64 = 1151
    ER_ABORTING_CONNECTION int64 = 1152
    ER_NET_PACKET_TOO_LARGE int64 = 1153
    ER_NET_READ_ERROR_FROM_PIPE int64 = 1154
    ER_NET_FCNTL_ERROR int64 = 1155
    ER_NET_PACKETS_OUT_OF_ORDER int64 = 1156
    ER_NET_UNCOMPRESS_ERROR int64 = 1157
    ER_NET_READ_ERROR int64 = 1158
    ER_NET_READ_INTERRUPTED int64 = 1159
    ER_NET_ERROR_ON_WRITE int64 = 1160
    ER_NET_WRITE_INTERRUPTED int64 = 1161
    ER_TOO_LONG_STRING int64 = 1162
    ER_TABLE_CANT_HANDLE_BLOB int64 = 1163
    ER_TABLE_CANT_HANDLE_AUTO_INCREMENT int64 = 1164
    ER_DELAYED_INSERT_TABLE_LOCKED int64 = 1165
    ER_WRONG_COLUMN_NAME int64 = 1166
    ER_WRONG_KEY_COLUMN int64 = 1167
    ER_WRONG_MRG_TABLE int64 = 1168
    ER_DUP_UNIQUE int64 = 1169
    ER_BLOB_KEY_WITHOUT_LENGTH int64 = 1170
    ER_PRIMARY_CANT_HAVE_NULL int64 = 1171
    ER_TOO_MANY_ROWS int64 = 1172
    ER_REQUIRES_PRIMARY_KEY int64 = 1173
    ER_NO_RAID_COMPILED int64 = 1174
    ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE int64 = 1175
    ER_KEY_DOES_NOT_EXITS int64 = 1176
    ER_CHECK_NO_SUCH_TABLE int64 = 1177
    ER_CHECK_NOT_IMPLEMENTED int64 = 1178
    ER_CANT_DO_THIS_DURING_AN_TRANSACTION int64 = 1179
    ER_ERROR_DURING_COMMIT int64 = 1180
    ER_ERROR_DURING_ROLLBACK int64 = 1181
    ER_ERROR_DURING_FLUSH_LOGS int64 = 1182
    ER_ERROR_DURING_CHECKPOINT int64 = 1183
    ER_NEW_ABORTING_CONNECTION int64 = 1184
    ER_DUMP_NOT_IMPLEMENTED int64 = 1185
    ER_FLUSH_MASTER_BINLOG_CLOSED int64 = 1186
    ER_INDEX_REBUILD int64 = 1187
    ER_MASTER int64 = 1188
    ER_MASTER_NET_READ int64 = 1189
    ER_MASTER_NET_WRITE int64 = 1190
    ER_FT_MATCHING_KEY_NOT_FOUND int64 = 1191
    ER_LOCK_OR_ACTIVE_TRANSACTION int64 = 1192
    ER_UNKNOWN_SYSTEM_VARIABLE int64 = 1193
    ER_CRASHED_ON_USAGE int64 = 1194
    ER_CRASHED_ON_REPAIR int64 = 1195
    ER_WARNING_NOT_COMPLETE_ROLLBACK int64 = 1196
    ER_TRANS_CACHE_FULL int64 = 1197
    ER_SLAVE_MUST_STOP int64 = 1198
    ER_SLAVE_NOT_RUNNING int64 = 1199
    ER_BAD_SLAVE int64 = 1200
    ER_MASTER_INFO int64 = 1201
    ER_SLAVE_THREAD int64 = 1202
    ER_TOO_MANY_USER_CONNECTIONS int64 = 1203
    ER_SET_CONSTANTS_ONLY int64 = 1204
    ER_LOCK_WAIT_TIMEOUT int64 = 1205
    ER_LOCK_TABLE_FULL int64 = 1206
    ER_READ_ONLY_TRANSACTION int64 = 1207
    ER_DROP_DB_WITH_READ_LOCK int64 = 1208
    ER_CREATE_DB_WITH_READ_LOCK int64 = 1209
    ER_WRONG_ARGUMENTS int64 = 1210
    ER_NO_PERMISSION_TO_CREATE_USER int64 = 1211
    ER_UNION_TABLES_IN_DIFFERENT_DIR int64 = 1212
    ER_LOCK_DEADLOCK int64 = 1213
    ER_TABLE_CANT_HANDLE_FT int64 = 1214
    ER_CANNOT_ADD_FOREIGN int64 = 1215
    ER_NO_REFERENCED_ROW int64 = 1216
    ER_ROW_IS_REFERENCED int64 = 1217
    ER_CONNECT_TO_MASTER int64 = 1218
    ER_QUERY_ON_MASTER int64 = 1219
    ER_ERROR_WHEN_EXECUTING_COMMAND int64 = 1220
    ER_WRONG_USAGE int64 = 1221
    ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT int64 = 1222
    ER_CANT_UPDATE_WITH_READLOCK int64 = 1223
    ER_MIXING_NOT_ALLOWED int64 = 1224
    ER_DUP_ARGUMENT int64 = 1225
    ER_USER_LIMIT_REACHED int64 = 1226
    ER_SPECIFIC_ACCESS_DENIED_ERROR int64 = 1227
    ER_LOCAL_VARIABLE int64 = 1228
    ER_GLOBAL_VARIABLE int64 = 1229
    ER_NO_DEFAULT int64 = 1230
    ER_WRONG_VALUE_FOR_VAR int64 = 1231
    ER_WRONG_TYPE_FOR_VAR int64 = 1232
    ER_VAR_CANT_BE_READ int64 = 1233
    ER_CANT_USE_OPTION_HERE int64 = 1234
    ER_NOT_SUPPORTED_YET int64 = 1235
    ER_MASTER_FATAL_ERROR_READING_BINLOG int64 = 1236
    ER_SLAVE_IGNORED_TABLE int64 = 1237
    ER_INCORRECT_GLOBAL_LOCAL_VAR int64 = 1238
    ER_WRONG_FK_DEF int64 = 1239
    ER_KEY_REF_DO_NOT_MATCH_TABLE_REF int64 = 1240
    ER_OPERAND_COLUMNS int64 = 1241
    ER_SUBQUERY_NO_1_ROW int64 = 1242
    ER_UNKNOWN_STMT_HANDLER int64 = 1243
    ER_CORRUPT_HELP_DB int64 = 1244
    ER_CYCLIC_REFERENCE int64 = 1245
    ER_AUTO_CONVERT int64 = 1246
    ER_ILLEGAL_REFERENCE int64 = 1247
    ER_DERIVED_MUST_HAVE_ALIAS int64 = 1248
    ER_SELECT_REDUCED int64 = 1249
    ER_TABLENAME_NOT_ALLOWED_HERE int64 = 1250
    ER_NOT_SUPPORTED_AUTH_MODE int64 = 1251
    ER_SPATIAL_CANT_HAVE_NULL int64 = 1252
    ER_COLLATION_CHARSET_MISMATCH int64 = 1253
    ER_SLAVE_WAS_RUNNING int64 = 1254
    ER_SLAVE_WAS_NOT_RUNNING int64 = 1255
    ER_TOO_BIG_FOR_UNCOMPRESS int64 = 1256
    ER_ZLIB_Z_MEM_ERROR int64 = 1257
    ER_ZLIB_Z_BUF_ERROR int64 = 1258
    ER_ZLIB_Z_DATA_ERROR int64 = 1259
    ER_CUT_VALUE_GROUP_CONCAT int64 = 1260
    ER_WARN_TOO_FEW_RECORDS int64 = 1261
    ER_WARN_TOO_MANY_RECORDS int64 = 1262
    ER_WARN_NULL_TO_NOTNULL int64 = 1263
    ER_WARN_DATA_OUT_OF_RANGE int64 = 1264
    WARN_DATA_TRUNCATED int64 = 1265
    ER_WARN_USING_OTHER_HANDLER int64 = 1266
    ER_CANT_AGGREGATE_2COLLATIONS int64 = 1267
    ER_DROP_USER int64 = 1268
    ER_REVOKE_GRANTS int64 = 1269
    ER_CANT_AGGREGATE_3COLLATIONS int64 = 1270
    ER_CANT_AGGREGATE_NCOLLATIONS int64 = 1271
    ER_VARIABLE_IS_NOT_STRUCT int64 = 1272
    ER_UNKNOWN_COLLATION int64 = 1273
    ER_SLAVE_IGNORED_SSL_PARAMS int64 = 1274
    ER_SERVER_IS_IN_SECURE_AUTH_MODE int64 = 1275
    ER_WARN_FIELD_RESOLVED int64 = 1276
    ER_BAD_SLAVE_UNTIL_COND int64 = 1277
    ER_MISSING_SKIP_SLAVE int64 = 1278
    ER_UNTIL_COND_IGNORED int64 = 1279
    ER_WRONG_NAME_FOR_INDEX int64 = 1280
    ER_WRONG_NAME_FOR_CATALOG int64 = 1281
    ER_WARN_QC_RESIZE int64 = 1282
    ER_BAD_FT_COLUMN int64 = 1283
    ER_UNKNOWN_KEY_CACHE int64 = 1284
    ER_WARN_HOSTNAME_WONT_WORK int64 = 1285
    ER_UNKNOWN_STORAGE_ENGINE int64 = 1286
    ER_WARN_DEPRECATED_SYNTAX int64 = 1287
    ER_NON_UPDATABLE_TABLE int64 = 1288
    ER_FEATURE_DISABLED int64 = 1289
    ER_OPTION_PREVENTS_STATEMENT int64 = 1290
    ER_DUPLICATED_VALUE_IN_TYPE int64 = 1291
    ER_TRUNCATED_WRONG_VALUE int64 = 1292
    ER_TOO_MUCH_AUTO_TIMESTAMP_COLS int64 = 1293
    ER_INVALID_ON_UPDATE int64 = 1294
    ER_UNSUPPORTED_PS int64 = 1295
    ER_GET_ERRMSG int64 = 1296
    ER_GET_TEMPORARY_ERRMSG int64 = 1297
    ER_UNKNOWN_TIME_ZONE int64 = 1298
    ER_WARN_INVALID_TIMESTAMP int64 = 1299
    ER_INVALID_CHARACTER_STRING int64 = 1300
    ER_WARN_ALLOWED_PACKET_OVERFLOWED int64 = 1301
    ER_CONFLICTING_DECLARATIONS int64 = 1302
    ER_SP_NO_RECURSIVE_CREATE int64 = 1303
    ER_SP_ALREADY_EXISTS int64 = 1304
    ER_SP_DOES_NOT_EXIST int64 = 1305
    ER_SP_DROP_FAILED int64 = 1306
    ER_SP_STORE_FAILED int64 = 1307
    ER_SP_LILABEL_MISMATCH int64 = 1308
    ER_SP_LABEL_REDEFINE int64 = 1309
    ER_SP_LABEL_MISMATCH int64 = 1310
    ER_SP_UNINIT_VAR int64 = 1311
    ER_SP_BADSELECT int64 = 1312
    ER_SP_BADRETURN int64 = 1313
    ER_SP_BADSTATEMENT int64 = 1314
    ER_UPDATE_LOG_DEPRECATED_IGNORED int64 = 1315
    ER_UPDATE_LOG_DEPRECATED_TRANSLATED int64 = 1316
    ER_QUERY_INTERRUPTED int64 = 1317
    ER_SP_WRONG_NO_OF_ARGS int64 = 1318
    ER_SP_COND_MISMATCH int64 = 1319
    ER_SP_NORETURN int64 = 1320
    ER_SP_NORETURNEND int64 = 1321
    ER_SP_BAD_CURSOR_QUERY int64 = 1322
    ER_SP_BAD_CURSOR_SELECT int64 = 1323
    ER_SP_CURSOR_MISMATCH int64 = 1324
    ER_SP_CURSOR_ALREADY_OPEN int64 = 1325
    ER_SP_CURSOR_NOT_OPEN int64 = 1326
    ER_SP_UNDECLARED_VAR int64 = 1327
    ER_SP_WRONG_NO_OF_FETCH_ARGS int64 = 1328
    ER_SP_FETCH_NO_DATA int64 = 1329
    ER_SP_DUP_PARAM int64 = 1330
    ER_SP_DUP_VAR int64 = 1331
    ER_SP_DUP_COND int64 = 1332
    ER_SP_DUP_CURS int64 = 1333
    ER_SP_CANT_ALTER int64 = 1334
    ER_SP_SUBSELECT_NYI int64 = 1335
    ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG int64 = 1336
    ER_SP_VARCOND_AFTER_CURSHNDLR int64 = 1337
    ER_SP_CURSOR_AFTER_HANDLER int64 = 1338
    ER_SP_CASE_NOT_FOUND int64 = 1339
    ER_FPARSER_TOO_BIG_FILE int64 = 1340
    ER_FPARSER_BAD_HEADER int64 = 1341
    ER_FPARSER_EOF_IN_COMMENT int64 = 1342
    ER_FPARSER_ERROR_IN_PARAMETER int64 = 1343
    ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER int64 = 1344
    ER_VIEW_NO_EXPLAIN int64 = 1345
    ER_FRM_UNKNOWN_TYPE int64 = 1346
    ER_WRONG_OBJECT int64 = 1347
    ER_NONUPDATEABLE_COLUMN int64 = 1348
    ER_VIEW_SELECT_DERIVED int64 = 1349
    ER_VIEW_SELECT_CLAUSE int64 = 1350
    ER_VIEW_SELECT_VARIABLE int64 = 1351
    ER_VIEW_SELECT_TMPTABLE int64 = 1352
    ER_VIEW_WRONG_LIST int64 = 1353
    ER_WARN_VIEW_MERGE int64 = 1354
    ER_WARN_VIEW_WITHOUT_KEY int64 = 1355
    ER_VIEW_INVALID int64 = 1356
    ER_SP_NO_DROP_SP int64 = 1357
    ER_SP_GOTO_IN_HNDLR int64 = 1358
    ER_TRG_ALREADY_EXISTS int64 = 1359
    ER_TRG_DOES_NOT_EXIST int64 = 1360
    ER_TRG_ON_VIEW_OR_TEMP_TABLE int64 = 1361
    ER_TRG_CANT_CHANGE_ROW int64 = 1362
    ER_TRG_NO_SUCH_ROW_IN_TRG int64 = 1363
    ER_NO_DEFAULT_FOR_FIELD int64 = 1364
    ER_DIVISION_BY_ZERO int64 = 1365
    ER_TRUNCATED_WRONG_VALUE_FOR_FIELD int64 = 1366
    ER_ILLEGAL_VALUE_FOR_TYPE int64 = 1367
    ER_VIEW_NONUPD_CHECK int64 = 1368
    ER_VIEW_CHECK_FAILED int64 = 1369
    ER_PROCACCESS_DENIED_ERROR int64 = 1370
    ER_RELAY_LOG_FAIL int64 = 1371
    ER_PASSWD_LENGTH int64 = 1372
    ER_UNKNOWN_TARGET_BINLOG int64 = 1373
    ER_IO_ERR_LOG_INDEX_READ int64 = 1374
    ER_BINLOG_PURGE_PROHIBITED int64 = 1375
    ER_FSEEK_FAIL int64 = 1376
    ER_BINLOG_PURGE_FATAL_ERR int64 = 1377
    ER_LOG_IN_USE int64 = 1378
    ER_LOG_PURGE_UNKNOWN_ERR int64 = 1379
    ER_RELAY_LOG_INIT int64 = 1380
    ER_NO_BINARY_LOGGING int64 = 1381
    ER_RESERVED_SYNTAX int64 = 1382
    ER_WSAS_FAILED int64 = 1383
    ER_DIFF_GROUPS_PROC int64 = 1384
    ER_NO_GROUP_FOR_PROC int64 = 1385
    ER_ORDER_WITH_PROC int64 = 1386
    ER_LOGGING_PROHIBIT_CHANGING_OF int64 = 1387
    ER_NO_FILE_MAPPING int64 = 1388
    ER_WRONG_MAGIC int64 = 1389
    ER_PS_MANY_PARAM int64 = 1390
    ER_KEY_PART_0 int64 = 1391
    ER_VIEW_CHECKSUM int64 = 1392
    ER_VIEW_MULTIUPDATE int64 = 1393
    ER_VIEW_NO_INSERT_FIELD_LIST int64 = 1394
    ER_VIEW_DELETE_MERGE_VIEW int64 = 1395
    ER_CANNOT_USER int64 = 1396
    ER_XAER_NOTA int64 = 1397
    ER_XAER_INVAL int64 = 1398
    ER_XAER_RMFAIL int64 = 1399
    ER_XAER_OUTSIDE int64 = 1400
    ER_XAER_RMERR int64 = 1401
    ER_XA_RBROLLBACK int64 = 1402
    ER_NONEXISTING_PROC_GRANT int64 = 1403
    ER_PROC_AUTO_GRANT_FAIL int64 = 1404
    ER_PROC_AUTO_REVOKE_FAIL int64 = 1405
    ER_DATA_TOO_LONG int64 = 1406
    ER_SP_BAD_SQLSTATE int64 = 1407
    ER_STARTUP int64 = 1408
    ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR int64 = 1409
    ER_CANT_CREATE_USER_WITH_GRANT int64 = 1410
    ER_WRONG_VALUE_FOR_TYPE int64 = 1411
    ER_TABLE_DEF_CHANGED int64 = 1412
    ER_SP_DUP_HANDLER int64 = 1413
    ER_SP_NOT_VAR_ARG int64 = 1414
    ER_SP_NO_RETSET int64 = 1415
    ER_CANT_CREATE_GEOMETRY_OBJECT int64 = 1416
    ER_FAILED_ROUTINE_BREAK_BINLOG int64 = 1417
    ER_BINLOG_UNSAFE_ROUTINE int64 = 1418
    ER_BINLOG_CREATE_ROUTINE_NEED_SUPER int64 = 1419
    ER_EXEC_STMT_WITH_OPEN_CURSOR int64 = 1420
    ER_STMT_HAS_NO_OPEN_CURSOR int64 = 1421
    ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG int64 = 1422
    ER_NO_DEFAULT_FOR_VIEW_FIELD int64 = 1423
    ER_SP_NO_RECURSION int64 = 1424
    ER_TOO_BIG_SCALE int64 = 1425
    ER_TOO_BIG_PRECISION int64 = 1426
    ER_M_BIGGER_THAN_D int64 = 1427
    ER_WRONG_LOCK_OF_SYSTEM_TABLE int64 = 1428
    ER_CONNECT_TO_FOREIGN_DATA_SOURCE int64 = 1429
    ER_QUERY_ON_FOREIGN_DATA_SOURCE int64 = 1430
    ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST int64 = 1431
    ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE int64 = 1432
    ER_FOREIGN_DATA_STRING_INVALID int64 = 1433
    ER_CANT_CREATE_FEDERATED_TABLE int64 = 1434
    ER_TRG_IN_WRONG_SCHEMA int64 = 1435
    ER_STACK_OVERRUN_NEED_MORE int64 = 1436
    ER_TOO_LONG_BODY int64 = 1437
    ER_WARN_CANT_DROP_DEFAULT_KEYCACHE int64 = 1438
    ER_TOO_BIG_DISPLAYWIDTH int64 = 1439
    ER_XAER_DUPID int64 = 1440
    ER_DATETIME_FUNCTION_OVERFLOW int64 = 1441
    ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG int64 = 1442
    ER_VIEW_PREVENT_UPDATE int64 = 1443
    ER_PS_NO_RECURSION int64 = 1444
    ER_SP_CANT_SET_AUTOCOMMIT int64 = 1445
    ER_MALFORMED_DEFINER int64 = 1446
    ER_VIEW_FRM_NO_USER int64 = 1447
    ER_VIEW_OTHER_USER int64 = 1448
    ER_NO_SUCH_USER int64 = 1449
    ER_FORBID_SCHEMA_CHANGE int64 = 1450
    ER_ROW_IS_REFERENCED_2 int64 = 1451
    ER_NO_REFERENCED_ROW_2 int64 = 1452
    ER_SP_BAD_VAR_SHADOW int64 = 1453
    ER_TRG_NO_DEFINER int64 = 1454
    ER_OLD_FILE_FORMAT int64 = 1455
    ER_SP_RECURSION_LIMIT int64 = 1456
    ER_SP_PROC_TABLE_CORRUPT int64 = 1457
    ER_SP_WRONG_NAME int64 = 1458
    ER_TABLE_NEEDS_UPGRADE int64 = 1459
    ER_SP_NO_AGGREGATE int64 = 1460
    ER_MAX_PREPARED_STMT_COUNT_REACHED int64 = 1461
    ER_VIEW_RECURSIVE int64 = 1462
    ER_NON_GROUPING_FIELD_USED int64 = 1463
    ER_TABLE_CANT_HANDLE_SPKEYS int64 = 1464
    ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA int64 = 1465
    ER_REMOVED_SPACES int64 = 1466
    ER_AUTOINC_READ_FAILED int64 = 1467
    ER_USERNAME int64 = 1468
    ER_HOSTNAME int64 = 1469
    ER_WRONG_STRING_LENGTH int64 = 1470
    ER_NON_INSERTABLE_TABLE int64 = 1471
    ER_ADMIN_WRONG_MRG_TABLE int64 = 1472
    ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT int64 = 1473
    ER_NAME_BECOMES_EMPTY int64 = 1474
    ER_AMBIGUOUS_FIELD_TERM int64 = 1475
    ER_FOREIGN_SERVER_EXISTS int64 = 1476
    ER_FOREIGN_SERVER_DOESNT_EXIST int64 = 1477
    ER_ILLEGAL_HA_CREATE_OPTION int64 = 1478
    ER_PARTITION_REQUIRES_VALUES_ERROR int64 = 1479
    ER_PARTITION_WRONG_VALUES_ERROR int64 = 1480
    ER_PARTITION_MAXVALUE_ERROR int64 = 1481
    ER_PARTITION_SUBPARTITION_ERROR int64 = 1482
    ER_PARTITION_SUBPART_MIX_ERROR int64 = 1483
    ER_PARTITION_WRONG_NO_PART_ERROR int64 = 1484
    ER_PARTITION_WRONG_NO_SUBPART_ERROR int64 = 1485
    ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR int64 = 1486
    ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR int64 = 1487
    ER_FIELD_NOT_FOUND_PART_ERROR int64 = 1488
    ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR int64 = 1489
    ER_INCONSISTENT_PARTITION_INFO_ERROR int64 = 1490
    ER_PARTITION_FUNC_NOT_ALLOWED_ERROR int64 = 1491
    ER_PARTITIONS_MUST_BE_DEFINED_ERROR int64 = 1492
    ER_RANGE_NOT_INCREASING_ERROR int64 = 1493
    ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR int64 = 1494
    ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR int64 = 1495
    ER_PARTITION_ENTRY_ERROR int64 = 1496
    ER_MIX_HANDLER_ERROR int64 = 1497
    ER_PARTITION_NOT_DEFINED_ERROR int64 = 1498
    ER_TOO_MANY_PARTITIONS_ERROR int64 = 1499
    ER_SUBPARTITION_ERROR int64 = 1500
    ER_CANT_CREATE_HANDLER_FILE int64 = 1501
    ER_BLOB_FIELD_IN_PART_FUNC_ERROR int64 = 1502
    ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF int64 = 1503
    ER_NO_PARTS_ERROR int64 = 1504
    ER_PARTITION_MGMT_ON_NONPARTITIONED int64 = 1505
    ER_FOREIGN_KEY_ON_PARTITIONED int64 = 1506
    ER_DROP_PARTITION_NON_EXISTENT int64 = 1507
    ER_DROP_LAST_PARTITION int64 = 1508
    ER_COALESCE_ONLY_ON_HASH_PARTITION int64 = 1509
    ER_REORG_HASH_ONLY_ON_SAME_NO int64 = 1510
    ER_REORG_NO_PARAM_ERROR int64 = 1511
    ER_ONLY_ON_RANGE_LIST_PARTITION int64 = 1512
    ER_ADD_PARTITION_SUBPART_ERROR int64 = 1513
    ER_ADD_PARTITION_NO_NEW_PARTITION int64 = 1514
    ER_COALESCE_PARTITION_NO_PARTITION int64 = 1515
    ER_REORG_PARTITION_NOT_EXIST int64 = 1516
    ER_SAME_NAME_PARTITION int64 = 1517
    ER_NO_BINLOG_ERROR int64 = 1518
    ER_CONSECUTIVE_REORG_PARTITIONS int64 = 1519
    ER_REORG_OUTSIDE_RANGE int64 = 1520
    ER_PARTITION_FUNCTION_FAILURE int64 = 1521
    ER_PART_STATE_ERROR int64 = 1522
    ER_LIMITED_PART_RANGE int64 = 1523
    ER_PLUGIN_IS_NOT_LOADED int64 = 1524
    ER_WRONG_VALUE int64 = 1525
    ER_NO_PARTITION_FOR_GIVEN_VALUE int64 = 1526
    ER_FILEGROUP_OPTION_ONLY_ONCE int64 = 1527
    ER_CREATE_FILEGROUP_FAILED int64 = 1528
    ER_DROP_FILEGROUP_FAILED int64 = 1529
    ER_TABLESPACE_AUTO_EXTEND_ERROR int64 = 1530
    ER_WRONG_SIZE_NUMBER int64 = 1531
    ER_SIZE_OVERFLOW_ERROR int64 = 1532
    ER_ALTER_FILEGROUP_FAILED int64 = 1533
    ER_BINLOG_ROW_LOGGING_FAILED int64 = 1534
    ER_BINLOG_ROW_WRONG_TABLE_DEF int64 = 1535
    ER_BINLOG_ROW_RBR_TO_SBR int64 = 1536
    ER_EVENT_ALREADY_EXISTS int64 = 1537
    ER_EVENT_STORE_FAILED int64 = 1538
    ER_EVENT_DOES_NOT_EXIST int64 = 1539
    ER_EVENT_CANT_ALTER int64 = 1540
    ER_EVENT_DROP_FAILED int64 = 1541
    ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG int64 = 1542
    ER_EVENT_ENDS_BEFORE_STARTS int64 = 1543
    ER_EVENT_EXEC_TIME_IN_THE_PAST int64 = 1544
    ER_EVENT_OPEN_TABLE_FAILED int64 = 1545
    ER_EVENT_NEITHER_M_EXPR_NOR_M_AT int64 = 1546
    ER_OBSOLETE_COL_COUNT_DOESNT_MATCH_CORRUPTED int64 = 1547
    ER_OBSOLETE_CANNOT_LOAD_FROM_TABLE int64 = 1548
    ER_EVENT_CANNOT_DELETE int64 = 1549
    ER_EVENT_COMPILE_ERROR int64 = 1550
    ER_EVENT_SAME_NAME int64 = 1551
    ER_EVENT_DATA_TOO_LONG int64 = 1552
    ER_DROP_INDEX_FK int64 = 1553
    ER_WARN_DEPRECATED_SYNTAX_WITH_VER int64 = 1554
    ER_CANT_WRITE_LOCK_LOG_TABLE int64 = 1555
    ER_CANT_LOCK_LOG_TABLE int64 = 1556
    ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED int64 = 1557
    ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE int64 = 1558
    ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR int64 = 1559
    ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT int64 = 1560
    ER_NDB_CANT_SWITCH_BINLOG_FORMAT int64 = 1561
    ER_PARTITION_NO_TEMPORARY int64 = 1562
    ER_PARTITION_CONST_DOMAIN_ERROR int64 = 1563
    ER_PARTITION_FUNCTION_IS_NOT_ALLOWED int64 = 1564
    ER_DDL_LOG_ERROR int64 = 1565
    ER_NULL_IN_VALUES_LESS_THAN int64 = 1566
    ER_WRONG_PARTITION_NAME int64 = 1567
    ER_CANT_CHANGE_TX_CHARACTERISTICS int64 = 1568
    ER_DUP_ENTRY_AUTOINCREMENT_CASE int64 = 1569
    ER_EVENT_MODIFY_QUEUE_ERROR int64 = 1570
    ER_EVENT_SET_VAR_ERROR int64 = 1571
    ER_PARTITION_MERGE_ERROR int64 = 1572
    ER_CANT_ACTIVATE_LOG int64 = 1573
    ER_RBR_NOT_AVAILABLE int64 = 1574
    ER_BASE64_DECODE_ERROR int64 = 1575
    ER_EVENT_RECURSION_FORBIDDEN int64 = 1576
    ER_EVENTS_DB_ERROR int64 = 1577
    ER_ONLY_INTEGERS_ALLOWED int64 = 1578
    ER_UNSUPORTED_LOG_ENGINE int64 = 1579
    ER_BAD_LOG_STATEMENT int64 = 1580
    ER_CANT_RENAME_LOG_TABLE int64 = 1581
    ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT int64 = 1582
    ER_WRONG_PARAMETERS_TO_NATIVE_FCT int64 = 1583
    ER_WRONG_PARAMETERS_TO_STORED_FCT int64 = 1584
    ER_NATIVE_FCT_NAME_COLLISION int64 = 1585
    ER_DUP_ENTRY_WITH_KEY_NAME int64 = 1586
    ER_BINLOG_PURGE_EMFILE int64 = 1587
    ER_EVENT_CANNOT_CREATE_IN_THE_PAST int64 = 1588
    ER_EVENT_CANNOT_ALTER_IN_THE_PAST int64 = 1589
    ER_SLAVE_INCIDENT int64 = 1590
    ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT int64 = 1591
    ER_BINLOG_UNSAFE_STATEMENT int64 = 1592
    ER_SLAVE_FATAL_ERROR int64 = 1593
    ER_SLAVE_RELAY_LOG_READ_FAILURE int64 = 1594
    ER_SLAVE_RELAY_LOG_WRITE_FAILURE int64 = 1595
    ER_SLAVE_CREATE_EVENT_FAILURE int64 = 1596
    ER_SLAVE_MASTER_COM_FAILURE int64 = 1597
    ER_BINLOG_LOGGING_IMPOSSIBLE int64 = 1598
    ER_VIEW_NO_CREATION_CTX int64 = 1599
    ER_VIEW_INVALID_CREATION_CTX int64 = 1600
    ER_SR_INVALID_CREATION_CTX int64 = 1601
    ER_TRG_CORRUPTED_FILE int64 = 1602
    ER_TRG_NO_CREATION_CTX int64 = 1603
    ER_TRG_INVALID_CREATION_CTX int64 = 1604
    ER_EVENT_INVALID_CREATION_CTX int64 = 1605
    ER_TRG_CANT_OPEN_TABLE int64 = 1606
    ER_CANT_CREATE_SROUTINE int64 = 1607
    ER_NEVER_USED int64 = 1608
    ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT int64 = 1609
    ER_SLAVE_CORRUPT_EVENT int64 = 1610
    ER_LOAD_DATA_INVALID_COLUMN int64 = 1611
    ER_LOG_PURGE_NO_FILE int64 = 1612
    ER_XA_RBTIMEOUT int64 = 1613
    ER_XA_RBDEADLOCK int64 = 1614
    ER_NEED_REPREPARE int64 = 1615
    ER_DELAYED_NOT_SUPPORTED int64 = 1616
    WARN_NO_MASTER_INFO int64 = 1617
    WARN_OPTION_IGNORED int64 = 1618
    WARN_PLUGIN_DELETE_BUILTIN int64 = 1619
    WARN_PLUGIN_BUSY int64 = 1620
    ER_VARIABLE_IS_READONLY int64 = 1621
    ER_WARN_ENGINE_TRANSACTION_ROLLBACK int64 = 1622
    ER_SLAVE_HEARTBEAT_FAILURE int64 = 1623
    ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE int64 = 1624
    ER_NDB_REPLICATION_SCHEMA_ERROR int64 = 1625
    ER_CONFLICT_FN_PARSE_ERROR int64 = 1626
    ER_EXCEPTIONS_WRITE_ERROR int64 = 1627
    ER_TOO_LONG_TABLE_COMMENT int64 = 1628
    ER_TOO_LONG_FIELD_COMMENT int64 = 1629
    ER_FUNC_INEXISTENT_NAME_COLLISION int64 = 1630
    ER_DATABASE_NAME int64 = 1631
    ER_TABLE_NAME int64 = 1632
    ER_PARTITION_NAME int64 = 1633
    ER_SUBPARTITION_NAME int64 = 1634
    ER_TEMPORARY_NAME int64 = 1635
    ER_RENAMED_NAME int64 = 1636
    ER_TOO_MANY_CONCURRENT_TRXS int64 = 1637
    WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED int64 = 1638
    ER_DEBUG_SYNC_TIMEOUT int64 = 1639
    ER_DEBUG_SYNC_HIT_LIMIT int64 = 1640
    ER_DUP_SIGNAL_SET int64 = 1641
    ER_SIGNAL_WARN int64 = 1642
    ER_SIGNAL_NOT_FOUND int64 = 1643
    ER_SIGNAL_EXCEPTION int64 = 1644
    ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER int64 = 1645
    ER_SIGNAL_BAD_CONDITION_TYPE int64 = 1646
    WARN_COND_ITEM_TRUNCATED int64 = 1647
    ER_COND_ITEM_TOO_LONG int64 = 1648
    ER_UNKNOWN_LOCALE int64 = 1649
    ER_SLAVE_IGNORE_SERVER_IDS int64 = 1650
    ER_QUERY_CACHE_DISABLED int64 = 1651
    ER_SAME_NAME_PARTITION_FIELD int64 = 1652
    ER_PARTITION_COLUMN_LIST_ERROR int64 = 1653
    ER_WRONG_TYPE_COLUMN_VALUE_ERROR int64 = 1654
    ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR int64 = 1655
    ER_MAXVALUE_IN_VALUES_IN int64 = 1656
    ER_TOO_MANY_VALUES_ERROR int64 = 1657
    ER_ROW_SINGLE_PARTITION_FIELD_ERROR int64 = 1658
    ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD int64 = 1659
    ER_PARTITION_FIELDS_TOO_LONG int64 = 1660
    ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE int64 = 1661
    ER_BINLOG_ROW_MODE_AND_STMT_ENGINE int64 = 1662
    ER_BINLOG_UNSAFE_AND_STMT_ENGINE int64 = 1663
    ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE int64 = 1664
    ER_BINLOG_STMT_MODE_AND_ROW_ENGINE int64 = 1665
    ER_BINLOG_ROW_INJECTION_AND_STMT_MODE int64 = 1666
    ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE int64 = 1667
    ER_BINLOG_UNSAFE_LIMIT int64 = 1668
    ER_BINLOG_UNSAFE_INSERT_DELAYED int64 = 1669
    ER_BINLOG_UNSAFE_SYSTEM_TABLE int64 = 1670
    ER_BINLOG_UNSAFE_AUTOINC_COLUMNS int64 = 1671
    ER_BINLOG_UNSAFE_UDF int64 = 1672
    ER_BINLOG_UNSAFE_SYSTEM_VARIABLE int64 = 1673
    ER_BINLOG_UNSAFE_SYSTEM_FUNCTION int64 = 1674
    ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS int64 = 1675
    ER_MESSAGE_AND_STATEMENT int64 = 1676
    ER_SLAVE_CONVERSION_FAILED int64 = 1677
    ER_SLAVE_CANT_CREATE_CONVERSION int64 = 1678
    ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT int64 = 1679
    ER_PATH_LENGTH int64 = 1680
    ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT int64 = 1681
    ER_WRONG_NATIVE_TABLE_STRUCTURE int64 = 1682
    ER_WRONG_PERFSCHEMA_USAGE int64 = 1683
    ER_WARN_I_S_SKIPPED_TABLE int64 = 1684
    ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT int64 = 1685
    ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT int64 = 1686
    ER_SPATIAL_MUST_HAVE_GEOM_COL int64 = 1687
    ER_TOO_LONG_INDEX_COMMENT int64 = 1688
    ER_LOCK_ABORTED int64 = 1689
    ER_DATA_OUT_OF_RANGE int64 = 1690
    ER_WRONG_SPVAR_TYPE_IN_LIMIT int64 = 1691
    ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE int64 = 1692
    ER_BINLOG_UNSAFE_MIXED_STATEMENT int64 = 1693
    ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN int64 = 1694
    ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN int64 = 1695
    ER_FAILED_READ_FROM_PAR_FILE int64 = 1696
    ER_VALUES_IS_NOT_INT_TYPE_ERROR int64 = 1697
    ER_ACCESS_DENIED_NO_PASSWORD_ERROR int64 = 1698
    ER_SET_PASSWORD_AUTH_PLUGIN int64 = 1699
    ER_GRANT_PLUGIN_USER_EXISTS int64 = 1700
    ER_TRUNCATE_ILLEGAL_FK int64 = 1701
    ER_PLUGIN_IS_PERMANENT int64 = 1702
    ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN int64 = 1703
    ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX int64 = 1704
    ER_STMT_CACHE_FULL int64 = 1705
    ER_MULTI_UPDATE_KEY_CONFLICT int64 = 1706
    ER_TABLE_NEEDS_REBUILD int64 = 1707
    WARN_OPTION_BELOW_LIMIT int64 = 1708
    ER_INDEX_COLUMN_TOO_LONG int64 = 1709
    ER_ERROR_IN_TRIGGER_BODY int64 = 1710
    ER_ERROR_IN_UNKNOWN_TRIGGER_BODY int64 = 1711
    ER_INDEX_CORRUPT int64 = 1712
    ER_UNDO_RECORD_TOO_BIG int64 = 1713
    ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT int64 = 1714
    ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE int64 = 1715
    ER_BINLOG_UNSAFE_REPLACE_SELECT int64 = 1716
    ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT int64 = 1717
    ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT int64 = 1718
    ER_BINLOG_UNSAFE_UPDATE_IGNORE int64 = 1719
    ER_PLUGIN_NO_UNINSTALL int64 = 1720
    ER_PLUGIN_NO_INSTALL int64 = 1721
    ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT int64 = 1722
    ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC int64 = 1723
    ER_BINLOG_UNSAFE_INSERT_TWO_KEYS int64 = 1724
    ER_TABLE_IN_FK_CHECK int64 = 1725
    ER_UNSUPPORTED_ENGINE int64 = 1726
    ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST int64 = 1727
    ER_CANNOT_LOAD_FROM_TABLE_V2 int64 = 1728
    ER_MASTER_DELAY_VALUE_OUT_OF_RANGE int64 = 1729
    ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT int64 = 1730
    ER_PARTITION_EXCHANGE_DIFFERENT_OPTION int64 = 1731
    ER_PARTITION_EXCHANGE_PART_TABLE int64 = 1732
    ER_PARTITION_EXCHANGE_TEMP_TABLE int64 = 1733
    ER_PARTITION_INSTEAD_OF_SUBPARTITION int64 = 1734
    ER_UNKNOWN_PARTITION int64 = 1735
    ER_TABLES_DIFFERENT_METADATA int64 = 1736
    ER_ROW_DOES_NOT_MATCH_PARTITION int64 = 1737
    ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX int64 = 1738
    ER_WARN_INDEX_NOT_APPLICABLE int64 = 1739
    ER_PARTITION_EXCHANGE_FOREIGN_KEY int64 = 1740
    ER_NO_SUCH_KEY_VALUE int64 = 1741
    ER_RPL_INFO_DATA_TOO_LONG int64 = 1742
    ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE int64 = 1743
    ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE int64 = 1744
    ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX int64 = 1745
    ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT int64 = 1746
    ER_PARTITION_CLAUSE_ON_NONPARTITIONED int64 = 1747
    ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET int64 = 1748
    ER_NO_SUCH_PARTITION__UNUSED int64 = 1749
    ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE int64 = 1750
    ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE int64 = 1751
    ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE int64 = 1752
    ER_MTS_FEATURE_IS_NOT_SUPPORTED int64 = 1753
    ER_MTS_UPDATED_DBS_GREATER_MAX int64 = 1754
    ER_MTS_CANT_PARALLEL int64 = 1755
    ER_MTS_INCONSISTENT_DATA int64 = 1756
    ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING int64 = 1757
    ER_DA_INVALID_CONDITION_NUMBER int64 = 1758
    ER_INSECURE_PLAIN_TEXT int64 = 1759
    ER_INSECURE_CHANGE_MASTER int64 = 1760
    ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO int64 = 1761
    ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO int64 = 1762
    ER_SQLTHREAD_WITH_SECURE_SLAVE int64 = 1763
    ER_TABLE_HAS_NO_FT int64 = 1764
    ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER int64 = 1765
    ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION int64 = 1766
    ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST int64 = 1767
    ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL int64 = 1768
    ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION int64 = 1769
    ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL int64 = 1770
    ER_SKIPPING_LOGGED_TRANSACTION int64 = 1771
    ER_MALFORMED_GTID_SET_SPECIFICATION int64 = 1772
    ER_MALFORMED_GTID_SET_ENCODING int64 = 1773
    ER_MALFORMED_GTID_SPECIFICATION int64 = 1774
    ER_GNO_EXHAUSTED int64 = 1775
    ER_BAD_SLAVE_AUTO_POSITION int64 = 1776
    ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON int64 = 1777
    ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET int64 = 1778
    ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON int64 = 1779
    ER_GTID_MODE_REQUIRES_BINLOG int64 = 1780
    ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF int64 = 1781
    ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON int64 = 1782
    ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF int64 = 1783
    ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF int64 = 1784
    ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE int64 = 1785
    ER_GTID_UNSAFE_CREATE_SELECT int64 = 1786
    ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTION int64 = 1787
    ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME int64 = 1788
    ER_MASTER_HAS_PURGED_REQUIRED_GTIDS int64 = 1789
    ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID int64 = 1790
    ER_UNKNOWN_EXPLAIN_FORMAT int64 = 1791
    ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION int64 = 1792
    ER_TOO_LONG_TABLE_PARTITION_COMMENT int64 = 1793
    ER_SLAVE_CONFIGURATION int64 = 1794
    ER_INNODB_FT_LIMIT int64 = 1795
    ER_INNODB_NO_FT_TEMP_TABLE int64 = 1796
    ER_INNODB_FT_WRONG_DOCID_COLUMN int64 = 1797
    ER_INNODB_FT_WRONG_DOCID_INDEX int64 = 1798
    ER_INNODB_ONLINE_LOG_TOO_BIG int64 = 1799
    ER_UNKNOWN_ALTER_ALGORITHM int64 = 1800
    ER_UNKNOWN_ALTER_LOCK int64 = 1801
    ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS int64 = 1802
    ER_MTS_RECOVERY_FAILURE int64 = 1803
    ER_MTS_RESET_WORKERS int64 = 1804
    ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2 int64 = 1805
    ER_SLAVE_SILENT_RETRY_TRANSACTION int64 = 1806
    ER_DISCARD_FK_CHECKS_RUNNING int64 = 1807
    ER_TABLE_SCHEMA_MISMATCH int64 = 1808
    ER_TABLE_IN_SYSTEM_TABLESPACE int64 = 1809
    ER_IO_READ_ERROR int64 = 1810
    ER_IO_WRITE_ERROR int64 = 1811
    ER_TABLESPACE_MISSING int64 = 1812
    ER_TABLESPACE_EXISTS int64 = 1813
    ER_TABLESPACE_DISCARDED int64 = 1814
    ER_INTERNAL_ERROR int64 = 1815
    ER_INNODB_IMPORT_ERROR int64 = 1816
    ER_INNODB_INDEX_CORRUPT int64 = 1817
    ER_INVALID_YEAR_COLUMN_LENGTH int64 = 1818
    ER_NOT_VALID_PASSWORD int64 = 1819
    ER_MUST_CHANGE_PASSWORD int64 = 1820
    ER_FK_NO_INDEX_CHILD int64 = 1821
    ER_FK_NO_INDEX_PARENT int64 = 1822
    ER_FK_FAIL_ADD_SYSTEM int64 = 1823
    ER_FK_CANNOT_OPEN_PARENT int64 = 1824
    ER_FK_INCORRECT_OPTION int64 = 1825
    ER_FK_DUP_NAME int64 = 1826
    ER_PASSWORD_FORMAT int64 = 1827
    ER_FK_COLUMN_CANNOT_DROP int64 = 1828
    ER_FK_COLUMN_CANNOT_DROP_CHILD int64 = 1829
    ER_FK_COLUMN_NOT_NULL int64 = 1830
    ER_DUP_INDEX int64 = 1831
    ER_FK_COLUMN_CANNOT_CHANGE int64 = 1832
    ER_FK_COLUMN_CANNOT_CHANGE_CHILD int64 = 1833
    ER_FK_CANNOT_DELETE_PARENT int64 = 1834
    ER_MALFORMED_PACKET int64 = 1835
    ER_READ_ONLY_MODE int64 = 1836
    ER_GTID_NEXT_TYPE_UNDEFINED_GROUP int64 = 1837
    ER_VARIABLE_NOT_SETTABLE_IN_SP int64 = 1838
    ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF int64 = 1839
    ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY int64 = 1840
    ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY int64 = 1841
    ER_GTID_PURGED_WAS_CHANGED int64 = 1842
    ER_GTID_EXECUTED_WAS_CHANGED int64 = 1843
    ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES int64 = 1844
    ER_ALTER_OPERATION_NOT_SUPPORTED int64 = 1845
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON int64 = 1846
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY int64 = 1847
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION int64 = 1848
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME int64 = 1849
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE int64 = 1850
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK int64 = 1851
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE int64 = 1852
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK int64 = 1853
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC int64 = 1854
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS int64 = 1855
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS int64 = 1856
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS int64 = 1857
    ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE int64 = 1858
    ER_DUP_UNKNOWN_IN_INDEX int64 = 1859
    ER_IDENT_CAUSES_TOO_LONG_PATH int64 = 1860
    ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL int64 = 1861
    ER_MUST_CHANGE_PASSWORD_LOGIN int64 = 1862
    ER_ROW_IN_WRONG_PARTITION int64 = 1863
    ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX int64 = 1864
    ER_INNODB_NO_FT_USES_PARSER int64 = 1865
    ER_BINLOG_LOGICAL_CORRUPTION int64 = 1866
    ER_WARN_PURGE_LOG_IN_USE int64 = 1867
    ER_WARN_PURGE_LOG_IS_ACTIVE int64 = 1868
    ER_AUTO_INCREMENT_CONFLICT int64 = 1869
    WARN_ON_BLOCKHOLE_IN_RBR int64 = 1870
    ER_SLAVE_MI_INIT_REPOSITORY int64 = 1871
    ER_SLAVE_RLI_INIT_REPOSITORY int64 = 1872
    ER_ACCESS_DENIED_CHANGE_USER_ERROR int64 = 1873
    ER_INNODB_READ_ONLY int64 = 1874
    ER_STOP_SLAVE_SQL_THREAD_TIMEOUT int64 = 1875
    ER_STOP_SLAVE_IO_THREAD_TIMEOUT int64 = 1876
    ER_TABLE_CORRUPT int64 = 1877
    ER_TEMP_FILE_WRITE_FAILURE int64 = 1878
    ER_INNODB_FT_AUX_NOT_HEX_ID int64 = 1879
    ER_OLD_TEMPORALS_UPGRADED int64 = 1880
    ER_INNODB_FORCED_RECOVERY int64 = 1881
    ER_AES_INVALID_IV int64 = 1882
    ER_PLUGIN_CANNOT_BE_UNINSTALLED int64 = 1883
    ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP int64 = 1884
    ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER int64 = 1885
    ER_MISSING_KEY int64 = 1886
    WARN_NAMED_PIPE_ACCESS_EVERYONE int64 = 1887
    ER_FOUND_MISSING_GTIDS int64 = 1888
    CR_UNKNOWN_ERROR int64 = 2000
    CR_SOCKET_CREATE_ERROR int64 = 2001
    CR_CONNECTION_ERROR int64 = 2002
    CR_CONN_HOST_ERROR int64 = 2003
    CR_IPSOCK_ERROR int64 = 2004
    CR_UNKNOWN_HOST int64 = 2005
    CR_SERVER_GONE_ERROR int64 = 2006
    CR_VERSION_ERROR int64 = 2007
    CR_OUT_OF_MEMORY int64 = 2008
    CR_WRONG_HOST_INFO int64 = 2009
    CR_LOCALHOST_CONNECTION int64 = 2010
    CR_TCP_CONNECTION int64 = 2011
    CR_SERVER_HANDSHAKE_ERR int64 = 2012
    CR_SERVER_LOST int64 = 2013
    CR_COMMANDS_OUT_OF_SYNC int64 = 2014
    CR_NAMEDPIPE_CONNECTION int64 = 2015
    CR_NAMEDPIPEWAIT_ERROR int64 = 2016
    CR_NAMEDPIPEOPEN_ERROR int64 = 2017
    CR_NAMEDPIPESETSTATE_ERROR int64 = 2018
    CR_CANT_READ_CHARSET int64 = 2019
    CR_NET_PACKET_TOO_LARGE int64 = 2020
    CR_EMBEDDED_CONNECTION int64 = 2021
    CR_PROBE_SLAVE_STATUS int64 = 2022
    CR_PROBE_SLAVE_HOSTS int64 = 2023
    CR_PROBE_SLAVE_CONNECT int64 = 2024
    CR_PROBE_MASTER_CONNECT int64 = 2025
    CR_SSL_CONNECTION_ERROR int64 = 2026
    CR_MALFORMED_PACKET int64 = 2027
    CR_WRONG_LICENSE int64 = 2028
    CR_NULL_POINTER int64 = 2029
    CR_NO_PREPARE_STMT int64 = 2030
    CR_PARAMS_NOT_BOUND int64 = 2031
    CR_DATA_TRUNCATED int64 = 2032
    CR_NO_PARAMETERS_EXISTS int64 = 2033
    CR_INVALID_PARAMETER_NO int64 = 2034
    CR_INVALID_BUFFER_USE int64 = 2035
    CR_UNSUPPORTED_PARAM_TYPE int64 = 2036
    CR_SHARED_MEMORY_CONNECTION int64 = 2037
    CR_SHARED_MEMORY_CONNECT_REQUEST_ERROR int64 = 2038
    CR_SHARED_MEMORY_CONNECT_ANSWER_ERROR int64 = 2039
    CR_SHARED_MEMORY_CONNECT_FILE_MAP_ERROR int64 = 2040
    CR_SHARED_MEMORY_CONNECT_MAP_ERROR int64 = 2041
    CR_SHARED_MEMORY_FILE_MAP_ERROR int64 = 2042
    CR_SHARED_MEMORY_MAP_ERROR int64 = 2043
    CR_SHARED_MEMORY_EVENT_ERROR int64 = 2044
    CR_SHARED_MEMORY_CONNECT_ABANDONED_ERROR int64 = 2045
    CR_SHARED_MEMORY_CONNECT_SET_ERROR int64 = 2046
    CR_CONN_UNKNOW_PROTOCOL int64 = 2047
    CR_INVALID_CONN_HANDLE int64 = 2048
    CR_SECURE_AUTH int64 = 2049
    CR_FETCH_CANCELED int64 = 2050
    CR_NO_DATA int64 = 2051
    CR_NO_STMT_METADATA int64 = 2052
    CR_NO_RESULT_SET int64 = 2053
    CR_NOT_IMPLEMENTED int64 = 2054
    CR_SERVER_LOST_EXTENDED int64 = 2055
    CR_STMT_CLOSED int64 = 2056
    CR_NEW_STMT_METADATA int64 = 2057
    CR_ALREADY_CONNECTED int64 = 2058
    CR_AUTH_PLUGIN_CANNOT_LOAD int64 = 2059
    CR_DUPLICATE_CONNECTION_ATTR int64 = 2060
    CR_AUTH_PLUGIN_ERR int64 = 2061
    EE_CANTCREATEFILE int64 = 1
    EE_READ int64 = 2
    EE_WRITE int64 = 3
    EE_BADCLOSE int64 = 4
    EE_OUTOFMEMORY int64 = 5
    EE_DELETE int64 = 6
    EE_LINK int64 = 7
    EE_EOFERR int64 = 9
    EE_CANTLOCK int64 = 10
    EE_CANTUNLOCK int64 = 11
    EE_DIR int64 = 12
    EE_STAT int64 = 13
    EE_CANT_CHSIZE int64 = 14
    EE_CANT_OPEN_STREAM int64 = 15
    EE_GETWD int64 = 16
    EE_SETWD int64 = 17
    EE_LINK_WARNING int64 = 18
    EE_OPEN_WARNING int64 = 19
    EE_DISK_FULL int64 = 20
    EE_CANT_MKDIR int64 = 21
    EE_UNKNOWN_CHARSET int64 = 22
    EE_OUT_OF_FILERESOURCES int64 = 23
    EE_CANT_READLINK int64 = 24
    EE_CANT_SYMLINK int64 = 25
    EE_REALPATH int64 = 26
    EE_SYNC int64 = 27
    EE_UNKNOWN_COLLATION int64 = 28
    EE_FILENOTFOUND int64 = 29
    EE_FILE_NOT_CLOSED int64 = 30
    EE_CHANGE_OWNERSHIP int64 = 31
    EE_CHANGE_PERMISSIONS int64 = 32
    EE_CANT_SEEK int64 = 33
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

    
// IsServerErrorCantCreateDb check mysql error is "Can't create database '%s' (errno: %d)" 
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

    
// IsServerErrorDbDropDelete check mysql error is "Error dropping database (can't delete '%s', errno: %d)" 
func IsServerErrorDbDropDelete(err error) bool {
    result := Isa(err, ER_DB_DROP_DELETE)
    return result
}

    
// IsServerErrorDbDropRmdir check mysql error is "Error dropping database (can't rmdir '%s', errno: %d)" 
func IsServerErrorDbDropRmdir(err error) bool {
    result := Isa(err, ER_DB_DROP_RMDIR)
    return result
}

    
// IsServerErrorCantDeleteFile check mysql error is "Error on delete of '%s' (errno: %d - %s)" 
func IsServerErrorCantDeleteFile(err error) bool {
    result := Isa(err, ER_CANT_DELETE_FILE)
    return result
}

    
// IsServerErrorCantFindSystemRec check mysql error is "Can't read record in system table" 
func IsServerErrorCantFindSystemRec(err error) bool {
    result := Isa(err, ER_CANT_FIND_SYSTEM_REC)
    return result
}

    
// IsServerErrorCantGetStat check mysql error is "Can't get status of '%s' (errno: %d - %s)" 
func IsServerErrorCantGetStat(err error) bool {
    result := Isa(err, ER_CANT_GET_STAT)
    return result
}

    
// IsServerErrorCantGetWd check mysql error is "Can't get working directory (errno: %d - %s)" 
func IsServerErrorCantGetWd(err error) bool {
    result := Isa(err, ER_CANT_GET_WD)
    return result
}

    
// IsServerErrorCantLock check mysql error is "Can't lock file (errno: %d - %s)" 
func IsServerErrorCantLock(err error) bool {
    result := Isa(err, ER_CANT_LOCK)
    return result
}

    
// IsServerErrorCantOpenFile check mysql error is "Can't open file: '%s' (errno: %d - %s)" 
func IsServerErrorCantOpenFile(err error) bool {
    result := Isa(err, ER_CANT_OPEN_FILE)
    return result
}

    
// IsServerErrorFileNotFound check mysql error is "Can't find file: '%s' (errno: %d - %s)" 
func IsServerErrorFileNotFound(err error) bool {
    result := Isa(err, ER_FILE_NOT_FOUND)
    return result
}

    
// IsServerErrorCantReadDir check mysql error is "Can't read dir of '%s' (errno: %d - %s)" 
func IsServerErrorCantReadDir(err error) bool {
    result := Isa(err, ER_CANT_READ_DIR)
    return result
}

    
// IsServerErrorCantSetWd check mysql error is "Can't change dir to '%s' (errno: %d - %s)" 
func IsServerErrorCantSetWd(err error) bool {
    result := Isa(err, ER_CANT_SET_WD)
    return result
}

    
// IsServerErrorCheckread check mysql error is "Record has changed since last read in table '%s'" 
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

    
// IsServerErrorErrorOnClose check mysql error is "Error on close of '%s' (errno: %d - %s)" 
func IsServerErrorErrorOnClose(err error) bool {
    result := Isa(err, ER_ERROR_ON_CLOSE)
    return result
}

    
// IsServerErrorErrorOnRead check mysql error is "Error reading file '%s' (errno: %d - %s)" 
func IsServerErrorErrorOnRead(err error) bool {
    result := Isa(err, ER_ERROR_ON_READ)
    return result
}

    
// IsServerErrorErrorOnRename check mysql error is "Error on rename of '%s' to '%s' (errno: %d - %s)" 
func IsServerErrorErrorOnRename(err error) bool {
    result := Isa(err, ER_ERROR_ON_RENAME)
    return result
}

    
// IsServerErrorErrorOnWrite check mysql error is "Error writing file '%s' (errno: %d - %s)" 
func IsServerErrorErrorOnWrite(err error) bool {
    result := Isa(err, ER_ERROR_ON_WRITE)
    return result
}

    
// IsServerErrorFileUsed check mysql error is "'%s' is locked against change" 
func IsServerErrorFileUsed(err error) bool {
    result := Isa(err, ER_FILE_USED)
    return result
}

    
// IsServerErrorFilsortAbort check mysql error is "Sort aborted" 
func IsServerErrorFilsortAbort(err error) bool {
    result := Isa(err, ER_FILSORT_ABORT)
    return result
}

    
// IsServerErrorFormNotFound check mysql error is "View '%s' doesn't exist for '%s'" 
func IsServerErrorFormNotFound(err error) bool {
    result := Isa(err, ER_FORM_NOT_FOUND)
    return result
}

    
// IsServerErrorGetErrno check mysql error is "Got error %d from storage engine" 
func IsServerErrorGetErrno(err error) bool {
    result := Isa(err, ER_GET_ERRNO)
    return result
}

    
// IsServerErrorIllegalHa check mysql error is "Table storage engine for '%s' doesn't have this option" 
func IsServerErrorIllegalHa(err error) bool {
    result := Isa(err, ER_ILLEGAL_HA)
    return result
}

    
// IsServerErrorKeyNotFound check mysql error is "Can't find record in '%s'" 
func IsServerErrorKeyNotFound(err error) bool {
    result := Isa(err, ER_KEY_NOT_FOUND)
    return result
}

    
// IsServerErrorNotFormFile check mysql error is "Incorrect information in file: '%s'" 
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

    
// IsServerErrorOpenAsReadonly check mysql error is "Table '%s' is read only" 
func IsServerErrorOpenAsReadonly(err error) bool {
    result := Isa(err, ER_OPEN_AS_READONLY)
    return result
}

    
// IsServerErrorOutofmemory check mysql error is "Out of memory" 
func IsServerErrorOutofmemory(err error) bool {
    result := Isa(err, ER_OUTOFMEMORY)
    return result
}

    
// IsServerErrorOutOfSortmemory check mysql error is "Out of sort memory, consider increasing server sortbuffer size" 
func IsServerErrorOutOfSortmemory(err error) bool {
    result := Isa(err, ER_OUT_OF_SORTMEMORY)
    return result
}

    
// IsServerErrorUnexpectedEof check mysql error is "Unexpected EOF found when reading file '%s' (errno: %d -%s)" 
func IsServerErrorUnexpectedEof(err error) bool {
    result := Isa(err, ER_UNEXPECTED_EOF)
    return result
}

    
// IsServerErrorConCountError check mysql error is "Too many connections" 
func IsServerErrorConCountError(err error) bool {
    result := Isa(err, ER_CON_COUNT_ERROR)
    return result
}

    
// IsServerErrorOutOfResources check mysql error is "Out of memory" 
func IsServerErrorOutOfResources(err error) bool {
    result := Isa(err, ER_OUT_OF_RESOURCES)
    return result
}

    
// IsServerErrorBadHostError check mysql error is "Can't get hostname for your address" 
func IsServerErrorBadHostError(err error) bool {
    result := Isa(err, ER_BAD_HOST_ERROR)
    return result
}

    
// IsServerErrorHandshakeError check mysql error is "Bad handshake" 
func IsServerErrorHandshakeError(err error) bool {
    result := Isa(err, ER_HANDSHAKE_ERROR)
    return result
}

    
// IsServerErrorDbaccessDeniedError check mysql error is "Access denied for user '%s'@'%s' to database '%s'" 
func IsServerErrorDbaccessDeniedError(err error) bool {
    result := Isa(err, ER_DBACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorAccessDeniedError check mysql error is "Access denied for user '%s'@'%s' (using password: %s)" 
func IsServerErrorAccessDeniedError(err error) bool {
    result := Isa(err, ER_ACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorNoDbError check mysql error is "No database selected" 
func IsServerErrorNoDbError(err error) bool {
    result := Isa(err, ER_NO_DB_ERROR)
    return result
}

    
// IsServerErrorUnknownComError check mysql error is "Unknown command" 
func IsServerErrorUnknownComError(err error) bool {
    result := Isa(err, ER_UNKNOWN_COM_ERROR)
    return result
}

    
// IsServerErrorBadNullError check mysql error is "Column '%s' cannot be null" 
func IsServerErrorBadNullError(err error) bool {
    result := Isa(err, ER_BAD_NULL_ERROR)
    return result
}

    
// IsServerErrorBadDbError check mysql error is "Unknown database '%s'" 
func IsServerErrorBadDbError(err error) bool {
    result := Isa(err, ER_BAD_DB_ERROR)
    return result
}

    
// IsServerErrorTableExistsError check mysql error is "Table '%s' already exists" 
func IsServerErrorTableExistsError(err error) bool {
    result := Isa(err, ER_TABLE_EXISTS_ERROR)
    return result
}

    
// IsServerErrorBadTableError check mysql error is "Unknown table '%s'" 
func IsServerErrorBadTableError(err error) bool {
    result := Isa(err, ER_BAD_TABLE_ERROR)
    return result
}

    
// IsServerErrorNonUniqError check mysql error is "Column '%s' in %s is ambiguous" 
func IsServerErrorNonUniqError(err error) bool {
    result := Isa(err, ER_NON_UNIQ_ERROR)
    return result
}

    
// IsServerErrorServerShutdown check mysql error is "Server shutdown in progress" 
func IsServerErrorServerShutdown(err error) bool {
    result := Isa(err, ER_SERVER_SHUTDOWN)
    return result
}

    
// IsServerErrorBadFieldError check mysql error is "Unknown column '%s' in '%s'" 
func IsServerErrorBadFieldError(err error) bool {
    result := Isa(err, ER_BAD_FIELD_ERROR)
    return result
}

    
// IsServerErrorWrongFieldWithGroup check mysql error is "'%s' isn't in GROUP BY" 
func IsServerErrorWrongFieldWithGroup(err error) bool {
    result := Isa(err, ER_WRONG_FIELD_WITH_GROUP)
    return result
}

    
// IsServerErrorWrongGroupField check mysql error is "Can't group on '%s'" 
func IsServerErrorWrongGroupField(err error) bool {
    result := Isa(err, ER_WRONG_GROUP_FIELD)
    return result
}

    
// IsServerErrorWrongSumSelect check mysql error is "Statement has sum functions and columns in same statement" 
func IsServerErrorWrongSumSelect(err error) bool {
    result := Isa(err, ER_WRONG_SUM_SELECT)
    return result
}

    
// IsServerErrorWrongValueCount check mysql error is "Column count doesn't match value count" 
func IsServerErrorWrongValueCount(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_COUNT)
    return result
}

    
// IsServerErrorTooLongIdent check mysql error is "Identifier name '%s' is too long" 
func IsServerErrorTooLongIdent(err error) bool {
    result := Isa(err, ER_TOO_LONG_IDENT)
    return result
}

    
// IsServerErrorDupFieldname check mysql error is "Duplicate column name '%s'" 
func IsServerErrorDupFieldname(err error) bool {
    result := Isa(err, ER_DUP_FIELDNAME)
    return result
}

    
// IsServerErrorDupKeyname check mysql error is "Duplicate key name '%s'" 
func IsServerErrorDupKeyname(err error) bool {
    result := Isa(err, ER_DUP_KEYNAME)
    return result
}

    
// IsServerErrorDupEntry check mysql error is "Duplicate entry '%s' for key %d" 
func IsServerErrorDupEntry(err error) bool {
    result := Isa(err, ER_DUP_ENTRY)
    return result
}

    
// IsServerErrorWrongFieldSpec check mysql error is "Incorrect column specifier for column '%s'" 
func IsServerErrorWrongFieldSpec(err error) bool {
    result := Isa(err, ER_WRONG_FIELD_SPEC)
    return result
}

    
// IsServerErrorParseError check mysql error is "%s near '%s' at line %d" 
func IsServerErrorParseError(err error) bool {
    result := Isa(err, ER_PARSE_ERROR)
    return result
}

    
// IsServerErrorEmptyQuery check mysql error is "Query was empty" 
func IsServerErrorEmptyQuery(err error) bool {
    result := Isa(err, ER_EMPTY_QUERY)
    return result
}

    
// IsServerErrorNonuniqTable check mysql error is "Not unique table/alias: '%s'" 
func IsServerErrorNonuniqTable(err error) bool {
    result := Isa(err, ER_NONUNIQ_TABLE)
    return result
}

    
// IsServerErrorInvalidDefault check mysql error is "Invalid default value for '%s'" 
func IsServerErrorInvalidDefault(err error) bool {
    result := Isa(err, ER_INVALID_DEFAULT)
    return result
}

    
// IsServerErrorMultiplePriKey check mysql error is "Multiple primary key defined" 
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

    
// IsServerErrorKeyColumnDoesNotExits check mysql error is "Key column '%s' doesn't exist in table" 
func IsServerErrorKeyColumnDoesNotExits(err error) bool {
    result := Isa(err, ER_KEY_COLUMN_DOES_NOT_EXITS)
    return result
}

    
// IsServerErrorBlobUsedAsKey check mysql error is "BLOB column '%s' can't be used in key specification withthe used table type" 
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

    
// IsServerErrorReady check mysql error is "%s: ready for connections. Version: '%s' socket: '%s'port: %d" 
func IsServerErrorReady(err error) bool {
    result := Isa(err, ER_READY)
    return result
}

    
// IsServerErrorNormalShutdown check mysql error is "%s: Normal shutdown" 
func IsServerErrorNormalShutdown(err error) bool {
    result := Isa(err, ER_NORMAL_SHUTDOWN)
    return result
}

    
// IsServerErrorGotSignal check mysql error is "%s: Got signal %d. Aborting!" 
func IsServerErrorGotSignal(err error) bool {
    result := Isa(err, ER_GOT_SIGNAL)
    return result
}

    
// IsServerErrorShutdownComplete check mysql error is "%s: Shutdown complete" 
func IsServerErrorShutdownComplete(err error) bool {
    result := Isa(err, ER_SHUTDOWN_COMPLETE)
    return result
}

    
// IsServerErrorForcingClose check mysql error is "%s: Forcing close of thread %ld user: '%s'" 
func IsServerErrorForcingClose(err error) bool {
    result := Isa(err, ER_FORCING_CLOSE)
    return result
}

    
// IsServerErrorIpsockError check mysql error is "Can't create IP socket" 
func IsServerErrorIpsockError(err error) bool {
    result := Isa(err, ER_IPSOCK_ERROR)
    return result
}

    
// IsServerErrorNoSuchIndex check mysql error is "Table '%s' has no index like the one used in CREATEINDEX" 
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

    
// IsServerErrorTextfileNotReadable check mysql error is "The file '%s' must be in the database directory or bereadable by all" 
func IsServerErrorTextfileNotReadable(err error) bool {
    result := Isa(err, ER_TEXTFILE_NOT_READABLE)
    return result
}

    
// IsServerErrorFileExistsError check mysql error is "File '%s' already exists" 
func IsServerErrorFileExistsError(err error) bool {
    result := Isa(err, ER_FILE_EXISTS_ERROR)
    return result
}

    
// IsServerErrorLoadInfo check mysql error is "Records: %ld Deleted: %ld Skipped: %ld Warnings: %ld" 
func IsServerErrorLoadInfo(err error) bool {
    result := Isa(err, ER_LOAD_INFO)
    return result
}

    
// IsServerErrorAlterInfo check mysql error is "Records: %ld Duplicates: %ld" 
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

    
// IsServerErrorInsertInfo check mysql error is "Records: %ld Duplicates: %ld Warnings: %ld" 
func IsServerErrorInsertInfo(err error) bool {
    result := Isa(err, ER_INSERT_INFO)
    return result
}

    
// IsServerErrorUpdateTableUsed check mysql error is "You can't specify target table '%s' for update in FROMclause" 
func IsServerErrorUpdateTableUsed(err error) bool {
    result := Isa(err, ER_UPDATE_TABLE_USED)
    return result
}

    
// IsServerErrorNoSuchThread check mysql error is "Unknown thread id: %lu" 
func IsServerErrorNoSuchThread(err error) bool {
    result := Isa(err, ER_NO_SUCH_THREAD)
    return result
}

    
// IsServerErrorKillDeniedError check mysql error is "You are not owner of thread %lu" 
func IsServerErrorKillDeniedError(err error) bool {
    result := Isa(err, ER_KILL_DENIED_ERROR)
    return result
}

    
// IsServerErrorNoTablesUsed check mysql error is "No tables used" 
func IsServerErrorNoTablesUsed(err error) bool {
    result := Isa(err, ER_NO_TABLES_USED)
    return result
}

    
// IsServerErrorTooBigSet check mysql error is "Too many strings for column %s and SET" 
func IsServerErrorTooBigSet(err error) bool {
    result := Isa(err, ER_TOO_BIG_SET)
    return result
}

    
// IsServerErrorNoUniqueLogfile check mysql error is "Can't generate a unique log-filename %s.(1-999)" 
func IsServerErrorNoUniqueLogfile(err error) bool {
    result := Isa(err, ER_NO_UNIQUE_LOGFILE)
    return result
}

    
// IsServerErrorTableNotLockedForWrite check mysql error is "Table '%s' was locked with a READ lock and can't beupdated" 
func IsServerErrorTableNotLockedForWrite(err error) bool {
    result := Isa(err, ER_TABLE_NOT_LOCKED_FOR_WRITE)
    return result
}

    
// IsServerErrorTableNotLocked check mysql error is "Table '%s' was not locked with LOCK TABLES" 
func IsServerErrorTableNotLocked(err error) bool {
    result := Isa(err, ER_TABLE_NOT_LOCKED)
    return result
}

    
// IsServerErrorBlobCantHaveDefault check mysql error is "BLOB/TEXT column '%s' can't have a default value" 
func IsServerErrorBlobCantHaveDefault(err error) bool {
    result := Isa(err, ER_BLOB_CANT_HAVE_DEFAULT)
    return result
}

    
// IsServerErrorWrongDbName check mysql error is "Incorrect database name '%s'" 
func IsServerErrorWrongDbName(err error) bool {
    result := Isa(err, ER_WRONG_DB_NAME)
    return result
}

    
// IsServerErrorWrongTableName check mysql error is "Incorrect table name '%s'" 
func IsServerErrorWrongTableName(err error) bool {
    result := Isa(err, ER_WRONG_TABLE_NAME)
    return result
}

    
// IsServerErrorTooBigSelect check mysql error is "The SELECT would examine more than MAX_JOIN_SIZE rows" 
func IsServerErrorTooBigSelect(err error) bool {
    result := Isa(err, ER_TOO_BIG_SELECT)
    return result
}

    
// IsServerErrorUnknownError check mysql error is "Unknown error" 
func IsServerErrorUnknownError(err error) bool {
    result := Isa(err, ER_UNKNOWN_ERROR)
    return result
}

    
// IsServerErrorUnknownProcedure check mysql error is "Unknown procedure '%s'" 
func IsServerErrorUnknownProcedure(err error) bool {
    result := Isa(err, ER_UNKNOWN_PROCEDURE)
    return result
}

    
// IsServerErrorWrongParamcountToProcedure check mysql error is "Incorrect parameter count to procedure '%s'" 
func IsServerErrorWrongParamcountToProcedure(err error) bool {
    result := Isa(err, ER_WRONG_PARAMCOUNT_TO_PROCEDURE)
    return result
}

    
// IsServerErrorWrongParametersToProcedure check mysql error is "Incorrect parameters to procedure '%s'" 
func IsServerErrorWrongParametersToProcedure(err error) bool {
    result := Isa(err, ER_WRONG_PARAMETERS_TO_PROCEDURE)
    return result
}

    
// IsServerErrorUnknownTable check mysql error is "Unknown table '%s' in %s" 
func IsServerErrorUnknownTable(err error) bool {
    result := Isa(err, ER_UNKNOWN_TABLE)
    return result
}

    
// IsServerErrorFieldSpecifiedTwice check mysql error is "Column '%s' specified twice" 
func IsServerErrorFieldSpecifiedTwice(err error) bool {
    result := Isa(err, ER_FIELD_SPECIFIED_TWICE)
    return result
}

    
// IsServerErrorInvalidGroupFuncUse check mysql error is "Invalid use of group function" 
func IsServerErrorInvalidGroupFuncUse(err error) bool {
    result := Isa(err, ER_INVALID_GROUP_FUNC_USE)
    return result
}

    
// IsServerErrorUnsupportedExtension check mysql error is "Table '%s' uses an extension that doesn't exist in thisMySQL version" 
func IsServerErrorUnsupportedExtension(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_EXTENSION)
    return result
}

    
// IsServerErrorTableMustHaveColumns check mysql error is "A table must have at least 1 column" 
func IsServerErrorTableMustHaveColumns(err error) bool {
    result := Isa(err, ER_TABLE_MUST_HAVE_COLUMNS)
    return result
}

    
// IsServerErrorRecordFileFull check mysql error is "The table '%s' is full" 
func IsServerErrorRecordFileFull(err error) bool {
    result := Isa(err, ER_RECORD_FILE_FULL)
    return result
}

    
// IsServerErrorUnknownCharacterSet check mysql error is "Unknown character set: '%s'" 
func IsServerErrorUnknownCharacterSet(err error) bool {
    result := Isa(err, ER_UNKNOWN_CHARACTER_SET)
    return result
}

    
// IsServerErrorTooManyTables check mysql error is "Too many tables" 
func IsServerErrorTooManyTables(err error) bool {
    result := Isa(err, ER_TOO_MANY_TABLES)
    return result
}

    
// IsServerErrorTooManyFields check mysql error is "Too many columns" 
func IsServerErrorTooManyFields(err error) bool {
    result := Isa(err, ER_TOO_MANY_FIELDS)
    return result
}

    
// IsServerErrorTooBigRowsize check mysql error is "Row size too large. The maximum row size for the usedtable type, not counting BLOBs, is %ld. This includes storageoverhead, check the manual. You have to change some columns toTEXT or BLOBs" 
func IsServerErrorTooBigRowsize(err error) bool {
    result := Isa(err, ER_TOO_BIG_ROWSIZE)
    return result
}

    
// IsServerErrorStackOverrun check mysql error is "Thread stack overrun: Used: %ld of a %ld stack. Use'mysqld --thread_stack=#' to specify a bigger stack if needed" 
func IsServerErrorStackOverrun(err error) bool {
    result := Isa(err, ER_STACK_OVERRUN)
    return result
}

    
// IsServerErrorWrongOuterJoin check mysql error is "Cross dependency found in OUTER JOIN" 
func IsServerErrorWrongOuterJoin(err error) bool {
    result := Isa(err, ER_WRONG_OUTER_JOIN)
    return result
}

    
// IsServerErrorNullColumnInIndex check mysql error is "Table handler doesn't support NULL in given index. Pleasechange column '%s' to be NOT NULL or use another handler" 
func IsServerErrorNullColumnInIndex(err error) bool {
    result := Isa(err, ER_NULL_COLUMN_IN_INDEX)
    return result
}

    
// IsServerErrorCantFindUdf check mysql error is "Can't load function '%s'" 
func IsServerErrorCantFindUdf(err error) bool {
    result := Isa(err, ER_CANT_FIND_UDF)
    return result
}

    
// IsServerErrorCantInitializeUdf check mysql error is "Can't initialize function '%s'" 
func IsServerErrorCantInitializeUdf(err error) bool {
    result := Isa(err, ER_CANT_INITIALIZE_UDF)
    return result
}

    
// IsServerErrorUdfNoPaths check mysql error is "No paths allowed for shared library" 
func IsServerErrorUdfNoPaths(err error) bool {
    result := Isa(err, ER_UDF_NO_PATHS)
    return result
}

    
// IsServerErrorUdfExists check mysql error is "Function '%s' already exists" 
func IsServerErrorUdfExists(err error) bool {
    result := Isa(err, ER_UDF_EXISTS)
    return result
}

    
// IsServerErrorCantOpenLibrary check mysql error is "Can't open shared library '%s' (errno: %d %s)" 
func IsServerErrorCantOpenLibrary(err error) bool {
    result := Isa(err, ER_CANT_OPEN_LIBRARY)
    return result
}

    
// IsServerErrorCantFindDlEntry check mysql error is "Can't find symbol '%s' in library" 
func IsServerErrorCantFindDlEntry(err error) bool {
    result := Isa(err, ER_CANT_FIND_DL_ENTRY)
    return result
}

    
// IsServerErrorFunctionNotDefined check mysql error is "Function '%s' is not defined" 
func IsServerErrorFunctionNotDefined(err error) bool {
    result := Isa(err, ER_FUNCTION_NOT_DEFINED)
    return result
}

    
// IsServerErrorHostIsBlocked check mysql error is "Host '%s' is blocked because of many connection errors" 
func IsServerErrorHostIsBlocked(err error) bool {
    result := Isa(err, ER_HOST_IS_BLOCKED)
    return result
}

    
// IsServerErrorHostNotPrivileged check mysql error is "Host '%s' is not allowed to connect to this MySQL server" 
func IsServerErrorHostNotPrivileged(err error) bool {
    result := Isa(err, ER_HOST_NOT_PRIVILEGED)
    return result
}

    
// IsServerErrorPasswordAnonymousUser check mysql error is "You are using MySQL as an anonymous user and anonymoususers are not allowed to change passwords" 
func IsServerErrorPasswordAnonymousUser(err error) bool {
    result := Isa(err, ER_PASSWORD_ANONYMOUS_USER)
    return result
}

    
// IsServerErrorPasswordNotAllowed check mysql error is "You must have privileges to update tables in the mysqldatabase to be able to change passwords for others" 
func IsServerErrorPasswordNotAllowed(err error) bool {
    result := Isa(err, ER_PASSWORD_NOT_ALLOWED)
    return result
}

    
// IsServerErrorPasswordNoMatch check mysql error is "Can't find any matching row in the user table" 
func IsServerErrorPasswordNoMatch(err error) bool {
    result := Isa(err, ER_PASSWORD_NO_MATCH)
    return result
}

    
// IsServerErrorUpdateInfo check mysql error is "Rows matched: %ld Changed: %ld Warnings: %ld" 
func IsServerErrorUpdateInfo(err error) bool {
    result := Isa(err, ER_UPDATE_INFO)
    return result
}

    
// IsServerErrorCantCreateThread check mysql error is "Can't create a new thread (errno %d)" 
func IsServerErrorCantCreateThread(err error) bool {
    result := Isa(err, ER_CANT_CREATE_THREAD)
    return result
}

    
// IsServerErrorWrongValueCountOnRow check mysql error is "Column count doesn't match value count at row %ld" 
func IsServerErrorWrongValueCountOnRow(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_COUNT_ON_ROW)
    return result
}

    
// IsServerErrorCantReopenTable check mysql error is "Can't reopen table: '%s'" 
func IsServerErrorCantReopenTable(err error) bool {
    result := Isa(err, ER_CANT_REOPEN_TABLE)
    return result
}

    
// IsServerErrorInvalidUseOfNull check mysql error is "Invalid use of NULL value" 
func IsServerErrorInvalidUseOfNull(err error) bool {
    result := Isa(err, ER_INVALID_USE_OF_NULL)
    return result
}

    
// IsServerErrorRegexpError check mysql error is "Got error '%s' from regexp" 
func IsServerErrorRegexpError(err error) bool {
    result := Isa(err, ER_REGEXP_ERROR)
    return result
}

    
// IsServerErrorMixOfGroupFuncAndFields check mysql error is "Mixing of GROUP columns (MIN(),MAX(),COUNT(),...) with noGROUP columns is illegal if there is no GROUP BY clause" 
func IsServerErrorMixOfGroupFuncAndFields(err error) bool {
    result := Isa(err, ER_MIX_OF_GROUP_FUNC_AND_FIELDS)
    return result
}

    
// IsServerErrorNonexistingGrant check mysql error is "There is no such grant defined for user '%s' on host '%s'" 
func IsServerErrorNonexistingGrant(err error) bool {
    result := Isa(err, ER_NONEXISTING_GRANT)
    return result
}

    
// IsServerErrorTableaccessDeniedError check mysql error is "%s command denied to user '%s'@'%s' for table '%s'" 
func IsServerErrorTableaccessDeniedError(err error) bool {
    result := Isa(err, ER_TABLEACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorColumnaccessDeniedError check mysql error is "%s command denied to user '%s'@'%s' for column '%s' intable '%s'" 
func IsServerErrorColumnaccessDeniedError(err error) bool {
    result := Isa(err, ER_COLUMNACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorIllegalGrantForTable check mysql error is "Illegal GRANT/REVOKE command" 
func IsServerErrorIllegalGrantForTable(err error) bool {
    result := Isa(err, ER_ILLEGAL_GRANT_FOR_TABLE)
    return result
}

    
// IsServerErrorGrantWrongHostOrUser check mysql error is "The host or user argument to GRANT is too long" 
func IsServerErrorGrantWrongHostOrUser(err error) bool {
    result := Isa(err, ER_GRANT_WRONG_HOST_OR_USER)
    return result
}

    
// IsServerErrorNoSuchTable check mysql error is "Table '%s.%s' doesn't exist" 
func IsServerErrorNoSuchTable(err error) bool {
    result := Isa(err, ER_NO_SUCH_TABLE)
    return result
}

    
// IsServerErrorNonexistingTableGrant check mysql error is "There is no such grant defined for user '%s' on host '%s'on table '%s'" 
func IsServerErrorNonexistingTableGrant(err error) bool {
    result := Isa(err, ER_NONEXISTING_TABLE_GRANT)
    return result
}

    
// IsServerErrorNotAllowedCommand check mysql error is "The used command is not allowed with this MySQL version" 
func IsServerErrorNotAllowedCommand(err error) bool {
    result := Isa(err, ER_NOT_ALLOWED_COMMAND)
    return result
}

    
// IsServerErrorSyntaxError check mysql error is "You have an error in your SQL syntax" 
func IsServerErrorSyntaxError(err error) bool {
    result := Isa(err, ER_SYNTAX_ERROR)
    return result
}

    
// IsServerErrorDelayedCantChangeLock check mysql error is "Delayed insert thread couldn't get requested lock fortable %s" 
func IsServerErrorDelayedCantChangeLock(err error) bool {
    result := Isa(err, ER_DELAYED_CANT_CHANGE_LOCK)
    return result
}

    
// IsServerErrorTooManyDelayedThreads check mysql error is "Too many delayed threads in use" 
func IsServerErrorTooManyDelayedThreads(err error) bool {
    result := Isa(err, ER_TOO_MANY_DELAYED_THREADS)
    return result
}

    
// IsServerErrorAbortingConnection check mysql error is "Aborted connection %ld to db: '%s' user: '%s' (%s)" 
func IsServerErrorAbortingConnection(err error) bool {
    result := Isa(err, ER_ABORTING_CONNECTION)
    return result
}

    
// IsServerErrorNetPacketTooLarge check mysql error is "Got a packet bigger than 'max_allowed_packet' bytes" 
func IsServerErrorNetPacketTooLarge(err error) bool {
    result := Isa(err, ER_NET_PACKET_TOO_LARGE)
    return result
}

    
// IsServerErrorNetReadErrorFromPipe check mysql error is "Got a read error from the connection pipe" 
func IsServerErrorNetReadErrorFromPipe(err error) bool {
    result := Isa(err, ER_NET_READ_ERROR_FROM_PIPE)
    return result
}

    
// IsServerErrorNetFcntlError check mysql error is "Got an error from fcntl()" 
func IsServerErrorNetFcntlError(err error) bool {
    result := Isa(err, ER_NET_FCNTL_ERROR)
    return result
}

    
// IsServerErrorNetPacketsOutOfOrder check mysql error is "Got packets out of order" 
func IsServerErrorNetPacketsOutOfOrder(err error) bool {
    result := Isa(err, ER_NET_PACKETS_OUT_OF_ORDER)
    return result
}

    
// IsServerErrorNetUncompressError check mysql error is "Couldn't uncompress communication packet" 
func IsServerErrorNetUncompressError(err error) bool {
    result := Isa(err, ER_NET_UNCOMPRESS_ERROR)
    return result
}

    
// IsServerErrorNetReadError check mysql error is "Got an error reading communication packets" 
func IsServerErrorNetReadError(err error) bool {
    result := Isa(err, ER_NET_READ_ERROR)
    return result
}

    
// IsServerErrorNetReadInterrupted check mysql error is "Got timeout reading communication packets" 
func IsServerErrorNetReadInterrupted(err error) bool {
    result := Isa(err, ER_NET_READ_INTERRUPTED)
    return result
}

    
// IsServerErrorNetErrorOnWrite check mysql error is "Got an error writing communication packets" 
func IsServerErrorNetErrorOnWrite(err error) bool {
    result := Isa(err, ER_NET_ERROR_ON_WRITE)
    return result
}

    
// IsServerErrorNetWriteInterrupted check mysql error is "Got timeout writing communication packets" 
func IsServerErrorNetWriteInterrupted(err error) bool {
    result := Isa(err, ER_NET_WRITE_INTERRUPTED)
    return result
}

    
// IsServerErrorTooLongString check mysql error is "Result string is longer than 'max_allowed_packet' bytes" 
func IsServerErrorTooLongString(err error) bool {
    result := Isa(err, ER_TOO_LONG_STRING)
    return result
}

    
// IsServerErrorTableCantHandleBlob check mysql error is "The used table type doesn't support BLOB/TEXT columns" 
func IsServerErrorTableCantHandleBlob(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_BLOB)
    return result
}

    
// IsServerErrorTableCantHandleAutoIncrement check mysql error is "The used table type doesn't support AUTO_INCREMENTcolumns" 
func IsServerErrorTableCantHandleAutoIncrement(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_AUTO_INCREMENT)
    return result
}

    
// IsServerErrorDelayedInsertTableLocked check mysql error is "INSERT DELAYED can't be used with table '%s' because itis locked with LOCK TABLES" 
func IsServerErrorDelayedInsertTableLocked(err error) bool {
    result := Isa(err, ER_DELAYED_INSERT_TABLE_LOCKED)
    return result
}

    
// IsServerErrorWrongColumnName check mysql error is "Incorrect column name '%s'" 
func IsServerErrorWrongColumnName(err error) bool {
    result := Isa(err, ER_WRONG_COLUMN_NAME)
    return result
}

    
// IsServerErrorWrongKeyColumn check mysql error is "The used storage engine can't index column '%s'" 
func IsServerErrorWrongKeyColumn(err error) bool {
    result := Isa(err, ER_WRONG_KEY_COLUMN)
    return result
}

    
// IsServerErrorWrongMrgTable check mysql error is "Unable to open underlying table which is differentlydefined or of non-MyISAM type or doesn't exist" 
func IsServerErrorWrongMrgTable(err error) bool {
    result := Isa(err, ER_WRONG_MRG_TABLE)
    return result
}

    
// IsServerErrorDupUnique check mysql error is "Can't write, because of unique constraint, to table '%s'" 
func IsServerErrorDupUnique(err error) bool {
    result := Isa(err, ER_DUP_UNIQUE)
    return result
}

    
// IsServerErrorBlobKeyWithoutLength check mysql error is "BLOB/TEXT column '%s' used in key specification without akey length" 
func IsServerErrorBlobKeyWithoutLength(err error) bool {
    result := Isa(err, ER_BLOB_KEY_WITHOUT_LENGTH)
    return result
}

    
// IsServerErrorPrimaryCantHaveNull check mysql error is "All parts of a PRIMARY KEY must be NOT NULL" 
func IsServerErrorPrimaryCantHaveNull(err error) bool {
    result := Isa(err, ER_PRIMARY_CANT_HAVE_NULL)
    return result
}

    
// IsServerErrorTooManyRows check mysql error is "Result consisted of more than one row" 
func IsServerErrorTooManyRows(err error) bool {
    result := Isa(err, ER_TOO_MANY_ROWS)
    return result
}

    
// IsServerErrorRequiresPrimaryKey check mysql error is "This table type requires a primary key" 
func IsServerErrorRequiresPrimaryKey(err error) bool {
    result := Isa(err, ER_REQUIRES_PRIMARY_KEY)
    return result
}

    
// IsServerErrorNoRaidCompiled check mysql error is "This version of MySQL is not compiled with RAID support" 
func IsServerErrorNoRaidCompiled(err error) bool {
    result := Isa(err, ER_NO_RAID_COMPILED)
    return result
}

    
// IsServerErrorUpdateWithoutKeyInSafeMode check mysql error is "You are using safe update mode and you tried to update atable without a WHERE that uses a KEY column" 
func IsServerErrorUpdateWithoutKeyInSafeMode(err error) bool {
    result := Isa(err, ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE)
    return result
}

    
// IsServerErrorKeyDoesNotExits check mysql error is "Key '%s' doesn't exist in table '%s'" 
func IsServerErrorKeyDoesNotExits(err error) bool {
    result := Isa(err, ER_KEY_DOES_NOT_EXITS)
    return result
}

    
// IsServerErrorCheckNoSuchTable check mysql error is "Can't open table" 
func IsServerErrorCheckNoSuchTable(err error) bool {
    result := Isa(err, ER_CHECK_NO_SUCH_TABLE)
    return result
}

    
// IsServerErrorCheckNotImplemented check mysql error is "The storage engine for the table doesn't support %s" 
func IsServerErrorCheckNotImplemented(err error) bool {
    result := Isa(err, ER_CHECK_NOT_IMPLEMENTED)
    return result
}

    
// IsServerErrorCantDoThisDuringAnTransaction check mysql error is "You are not allowed to execute this command in atransaction" 
func IsServerErrorCantDoThisDuringAnTransaction(err error) bool {
    result := Isa(err, ER_CANT_DO_THIS_DURING_AN_TRANSACTION)
    return result
}

    
// IsServerErrorErrorDuringCommit check mysql error is "Got error %d during COMMIT" 
func IsServerErrorErrorDuringCommit(err error) bool {
    result := Isa(err, ER_ERROR_DURING_COMMIT)
    return result
}

    
// IsServerErrorErrorDuringRollback check mysql error is "Got error %d during ROLLBACK" 
func IsServerErrorErrorDuringRollback(err error) bool {
    result := Isa(err, ER_ERROR_DURING_ROLLBACK)
    return result
}

    
// IsServerErrorErrorDuringFlushLogs check mysql error is "Got error %d during FLUSH_LOGS" 
func IsServerErrorErrorDuringFlushLogs(err error) bool {
    result := Isa(err, ER_ERROR_DURING_FLUSH_LOGS)
    return result
}

    
// IsServerErrorErrorDuringCheckpoint check mysql error is "Got error %d during CHECKPOINT" 
func IsServerErrorErrorDuringCheckpoint(err error) bool {
    result := Isa(err, ER_ERROR_DURING_CHECKPOINT)
    return result
}

    
// IsServerErrorNewAbortingConnection check mysql error is "Aborted connection %ld to db: '%s' user: '%s' host: '%s'(%s)" 
func IsServerErrorNewAbortingConnection(err error) bool {
    result := Isa(err, ER_NEW_ABORTING_CONNECTION)
    return result
}

    
// IsServerErrorDumpNotImplemented check mysql error is "The storage engine for the table does not support binarytable dump" 
func IsServerErrorDumpNotImplemented(err error) bool {
    result := Isa(err, ER_DUMP_NOT_IMPLEMENTED)
    return result
}

    
// IsServerErrorFlushMasterBinlogClosed check mysql error is "Binlog closed, cannot RESET MASTER" 
func IsServerErrorFlushMasterBinlogClosed(err error) bool {
    result := Isa(err, ER_FLUSH_MASTER_BINLOG_CLOSED)
    return result
}

    
// IsServerErrorIndexRebuild check mysql error is "Failed rebuilding the index of dumped table '%s'" 
func IsServerErrorIndexRebuild(err error) bool {
    result := Isa(err, ER_INDEX_REBUILD)
    return result
}

    
// IsServerErrorMaster check mysql error is "Error from master: '%s'" 
func IsServerErrorMaster(err error) bool {
    result := Isa(err, ER_MASTER)
    return result
}

    
// IsServerErrorMasterNetRead check mysql error is "Net error reading from master" 
func IsServerErrorMasterNetRead(err error) bool {
    result := Isa(err, ER_MASTER_NET_READ)
    return result
}

    
// IsServerErrorMasterNetWrite check mysql error is "Net error writing to master" 
func IsServerErrorMasterNetWrite(err error) bool {
    result := Isa(err, ER_MASTER_NET_WRITE)
    return result
}

    
// IsServerErrorFtMatchingKeyNotFound check mysql error is "Can't find FULLTEXT index matching the column list" 
func IsServerErrorFtMatchingKeyNotFound(err error) bool {
    result := Isa(err, ER_FT_MATCHING_KEY_NOT_FOUND)
    return result
}

    
// IsServerErrorLockOrActiveTransaction check mysql error is "Can't execute the given command because you have activelocked tables or an active transaction" 
func IsServerErrorLockOrActiveTransaction(err error) bool {
    result := Isa(err, ER_LOCK_OR_ACTIVE_TRANSACTION)
    return result
}

    
// IsServerErrorUnknownSystemVariable check mysql error is "Unknown system variable '%s'" 
func IsServerErrorUnknownSystemVariable(err error) bool {
    result := Isa(err, ER_UNKNOWN_SYSTEM_VARIABLE)
    return result
}

    
// IsServerErrorCrashedOnUsage check mysql error is "Table '%s' is marked as crashed and should be repaired" 
func IsServerErrorCrashedOnUsage(err error) bool {
    result := Isa(err, ER_CRASHED_ON_USAGE)
    return result
}

    
// IsServerErrorCrashedOnRepair check mysql error is "Table '%s' is marked as crashed and last (automatic?)repair failed" 
func IsServerErrorCrashedOnRepair(err error) bool {
    result := Isa(err, ER_CRASHED_ON_REPAIR)
    return result
}

    
// IsServerErrorWarningNotCompleteRollback check mysql error is "Some non-transactional changed tables couldn't be rolledback" 
func IsServerErrorWarningNotCompleteRollback(err error) bool {
    result := Isa(err, ER_WARNING_NOT_COMPLETE_ROLLBACK)
    return result
}

    
// IsServerErrorTransCacheFull check mysql error is "Multi-statement transaction required more than'max_binlog_cache_size' bytes of storage" 
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

    
// IsServerErrorTooManyUserConnections check mysql error is "User %s already has more than 'max_user_connections'active connections" 
func IsServerErrorTooManyUserConnections(err error) bool {
    result := Isa(err, ER_TOO_MANY_USER_CONNECTIONS)
    return result
}

    
// IsServerErrorSetConstantsOnly check mysql error is "You may only use constant expressions with SET" 
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

    
// IsServerErrorReadOnlyTransaction check mysql error is "Update locks cannot be acquired during a READ UNCOMMITTEDtransaction" 
func IsServerErrorReadOnlyTransaction(err error) bool {
    result := Isa(err, ER_READ_ONLY_TRANSACTION)
    return result
}

    
// IsServerErrorDropDbWithReadLock check mysql error is "DROP DATABASE not allowed while thread is holding globalread lock" 
func IsServerErrorDropDbWithReadLock(err error) bool {
    result := Isa(err, ER_DROP_DB_WITH_READ_LOCK)
    return result
}

    
// IsServerErrorCreateDbWithReadLock check mysql error is "CREATE DATABASE not allowed while thread is holdingglobal read lock" 
func IsServerErrorCreateDbWithReadLock(err error) bool {
    result := Isa(err, ER_CREATE_DB_WITH_READ_LOCK)
    return result
}

    
// IsServerErrorWrongArguments check mysql error is "Incorrect arguments to %s" 
func IsServerErrorWrongArguments(err error) bool {
    result := Isa(err, ER_WRONG_ARGUMENTS)
    return result
}

    
// IsServerErrorNoPermissionToCreateUser check mysql error is "'%s'@'%s' is not allowed to create new users" 
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

    
// IsServerErrorTableCantHandleFt check mysql error is "The used table type doesn't support FULLTEXT indexes" 
func IsServerErrorTableCantHandleFt(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_FT)
    return result
}

    
// IsServerErrorCannotAddForeign check mysql error is "Cannot add foreign key constraint" 
func IsServerErrorCannotAddForeign(err error) bool {
    result := Isa(err, ER_CANNOT_ADD_FOREIGN)
    return result
}

    
// IsServerErrorNoReferencedRow check mysql error is "Cannot add or update a child row: a foreign keyconstraint fails" 
func IsServerErrorNoReferencedRow(err error) bool {
    result := Isa(err, ER_NO_REFERENCED_ROW)
    return result
}

    
// IsServerErrorRowIsReferenced check mysql error is "Cannot delete or update a parent row: a foreign keyconstraint fails" 
func IsServerErrorRowIsReferenced(err error) bool {
    result := Isa(err, ER_ROW_IS_REFERENCED)
    return result
}

    
// IsServerErrorConnectToMaster check mysql error is "Error connecting to master: %s" 
func IsServerErrorConnectToMaster(err error) bool {
    result := Isa(err, ER_CONNECT_TO_MASTER)
    return result
}

    
// IsServerErrorQueryOnMaster check mysql error is "Error running query on master: %s" 
func IsServerErrorQueryOnMaster(err error) bool {
    result := Isa(err, ER_QUERY_ON_MASTER)
    return result
}

    
// IsServerErrorErrorWhenExecutingCommand check mysql error is "Error when executing command %s: %s" 
func IsServerErrorErrorWhenExecutingCommand(err error) bool {
    result := Isa(err, ER_ERROR_WHEN_EXECUTING_COMMAND)
    return result
}

    
// IsServerErrorWrongUsage check mysql error is "Incorrect usage of %s and %s" 
func IsServerErrorWrongUsage(err error) bool {
    result := Isa(err, ER_WRONG_USAGE)
    return result
}

    
// IsServerErrorWrongNumberOfColumnsInSelect check mysql error is "The used SELECT statements have a different number ofcolumns" 
func IsServerErrorWrongNumberOfColumnsInSelect(err error) bool {
    result := Isa(err, ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT)
    return result
}

    
// IsServerErrorCantUpdateWithReadlock check mysql error is "Can't execute the query because you have a conflictingread lock" 
func IsServerErrorCantUpdateWithReadlock(err error) bool {
    result := Isa(err, ER_CANT_UPDATE_WITH_READLOCK)
    return result
}

    
// IsServerErrorMixingNotAllowed check mysql error is "Mixing of transactional and non-transactional tables isdisabled" 
func IsServerErrorMixingNotAllowed(err error) bool {
    result := Isa(err, ER_MIXING_NOT_ALLOWED)
    return result
}

    
// IsServerErrorDupArgument check mysql error is "Option '%s' used twice in statement" 
func IsServerErrorDupArgument(err error) bool {
    result := Isa(err, ER_DUP_ARGUMENT)
    return result
}

    
// IsServerErrorUserLimitReached check mysql error is "User '%s' has exceeded the '%s' resource (current value:%ld)" 
func IsServerErrorUserLimitReached(err error) bool {
    result := Isa(err, ER_USER_LIMIT_REACHED)
    return result
}

    
// IsServerErrorSpecificAccessDeniedError check mysql error is "Access denied" 
func IsServerErrorSpecificAccessDeniedError(err error) bool {
    result := Isa(err, ER_SPECIFIC_ACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorLocalVariable check mysql error is "Variable '%s' is a SESSION variable and can't be usedwith SET GLOBAL" 
func IsServerErrorLocalVariable(err error) bool {
    result := Isa(err, ER_LOCAL_VARIABLE)
    return result
}

    
// IsServerErrorGlobalVariable check mysql error is "Variable '%s' is a GLOBAL variable and should be set withSET GLOBAL" 
func IsServerErrorGlobalVariable(err error) bool {
    result := Isa(err, ER_GLOBAL_VARIABLE)
    return result
}

    
// IsServerErrorNoDefault check mysql error is "Variable '%s' doesn't have a default value" 
func IsServerErrorNoDefault(err error) bool {
    result := Isa(err, ER_NO_DEFAULT)
    return result
}

    
// IsServerErrorWrongValueForVar check mysql error is "Variable '%s' can't be set to the value of '%s'" 
func IsServerErrorWrongValueForVar(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_FOR_VAR)
    return result
}

    
// IsServerErrorWrongTypeForVar check mysql error is "Incorrect argument type to variable '%s'" 
func IsServerErrorWrongTypeForVar(err error) bool {
    result := Isa(err, ER_WRONG_TYPE_FOR_VAR)
    return result
}

    
// IsServerErrorVarCantBeRead check mysql error is "Variable '%s' can only be set, not read" 
func IsServerErrorVarCantBeRead(err error) bool {
    result := Isa(err, ER_VAR_CANT_BE_READ)
    return result
}

    
// IsServerErrorCantUseOptionHere check mysql error is "Incorrect usage/placement of '%s'" 
func IsServerErrorCantUseOptionHere(err error) bool {
    result := Isa(err, ER_CANT_USE_OPTION_HERE)
    return result
}

    
// IsServerErrorNotSupportedYet check mysql error is "This version of MySQL doesn't yet support '%s'" 
func IsServerErrorNotSupportedYet(err error) bool {
    result := Isa(err, ER_NOT_SUPPORTED_YET)
    return result
}

    
// IsServerErrorMasterFatalErrorReadingBinlog check mysql error is "Got fatal error %d from master when reading data frombinary log: '%s'" 
func IsServerErrorMasterFatalErrorReadingBinlog(err error) bool {
    result := Isa(err, ER_MASTER_FATAL_ERROR_READING_BINLOG)
    return result
}

    
// IsServerErrorSlaveIgnoredTable check mysql error is "Slave SQL thread ignored the query because ofreplicate-*-table rules" 
func IsServerErrorSlaveIgnoredTable(err error) bool {
    result := Isa(err, ER_SLAVE_IGNORED_TABLE)
    return result
}

    
// IsServerErrorIncorrectGlobalLocalVar check mysql error is "Variable '%s' is a %s variable" 
func IsServerErrorIncorrectGlobalLocalVar(err error) bool {
    result := Isa(err, ER_INCORRECT_GLOBAL_LOCAL_VAR)
    return result
}

    
// IsServerErrorWrongFkDef check mysql error is "Incorrect foreign key definition for '%s': %s" 
func IsServerErrorWrongFkDef(err error) bool {
    result := Isa(err, ER_WRONG_FK_DEF)
    return result
}

    
// IsServerErrorKeyRefDoNotMatchTableRef check mysql error is "Key reference and table reference don't match" 
func IsServerErrorKeyRefDoNotMatchTableRef(err error) bool {
    result := Isa(err, ER_KEY_REF_DO_NOT_MATCH_TABLE_REF)
    return result
}

    
// IsServerErrorOperandColumns check mysql error is "Operand should contain %d column(s)" 
func IsServerErrorOperandColumns(err error) bool {
    result := Isa(err, ER_OPERAND_COLUMNS)
    return result
}

    
// IsServerErrorSubqueryNo1Row check mysql error is "Subquery returns more than 1 row" 
func IsServerErrorSubqueryNo1Row(err error) bool {
    result := Isa(err, ER_SUBQUERY_NO_1_ROW)
    return result
}

    
// IsServerErrorUnknownStmtHandler check mysql error is "Unknown prepared statement handler (%.*s) given to %s" 
func IsServerErrorUnknownStmtHandler(err error) bool {
    result := Isa(err, ER_UNKNOWN_STMT_HANDLER)
    return result
}

    
// IsServerErrorCorruptHelpDb check mysql error is "Help database is corrupt or does not exist" 
func IsServerErrorCorruptHelpDb(err error) bool {
    result := Isa(err, ER_CORRUPT_HELP_DB)
    return result
}

    
// IsServerErrorCyclicReference check mysql error is "Cyclic reference on subqueries" 
func IsServerErrorCyclicReference(err error) bool {
    result := Isa(err, ER_CYCLIC_REFERENCE)
    return result
}

    
// IsServerErrorAutoConvert check mysql error is "Converting column '%s' from %s to %s" 
func IsServerErrorAutoConvert(err error) bool {
    result := Isa(err, ER_AUTO_CONVERT)
    return result
}

    
// IsServerErrorIllegalReference check mysql error is "Reference '%s' not supported (%s)" 
func IsServerErrorIllegalReference(err error) bool {
    result := Isa(err, ER_ILLEGAL_REFERENCE)
    return result
}

    
// IsServerErrorDerivedMustHaveAlias check mysql error is "Every derived table must have its own alias" 
func IsServerErrorDerivedMustHaveAlias(err error) bool {
    result := Isa(err, ER_DERIVED_MUST_HAVE_ALIAS)
    return result
}

    
// IsServerErrorSelectReduced check mysql error is "Select %u was reduced during optimization" 
func IsServerErrorSelectReduced(err error) bool {
    result := Isa(err, ER_SELECT_REDUCED)
    return result
}

    
// IsServerErrorTablenameNotAllowedHere check mysql error is "Table '%s' from one of the SELECTs cannot be used in %s" 
func IsServerErrorTablenameNotAllowedHere(err error) bool {
    result := Isa(err, ER_TABLENAME_NOT_ALLOWED_HERE)
    return result
}

    
// IsServerErrorNotSupportedAuthMode check mysql error is "Client does not support authentication protocol requestedby server" 
func IsServerErrorNotSupportedAuthMode(err error) bool {
    result := Isa(err, ER_NOT_SUPPORTED_AUTH_MODE)
    return result
}

    
// IsServerErrorSpatialCantHaveNull check mysql error is "All parts of a SPATIAL index must be NOT NULL" 
func IsServerErrorSpatialCantHaveNull(err error) bool {
    result := Isa(err, ER_SPATIAL_CANT_HAVE_NULL)
    return result
}

    
// IsServerErrorCollationCharsetMismatch check mysql error is "COLLATION '%s' is not valid for CHARACTER SET '%s'" 
func IsServerErrorCollationCharsetMismatch(err error) bool {
    result := Isa(err, ER_COLLATION_CHARSET_MISMATCH)
    return result
}

    
// IsServerErrorSlaveWasRunning check mysql error is "Slave is already running" 
func IsServerErrorSlaveWasRunning(err error) bool {
    result := Isa(err, ER_SLAVE_WAS_RUNNING)
    return result
}

    
// IsServerErrorSlaveWasNotRunning check mysql error is "Slave already has been stopped" 
func IsServerErrorSlaveWasNotRunning(err error) bool {
    result := Isa(err, ER_SLAVE_WAS_NOT_RUNNING)
    return result
}

    
// IsServerErrorTooBigForUncompress check mysql error is "Uncompressed data size too large" 
func IsServerErrorTooBigForUncompress(err error) bool {
    result := Isa(err, ER_TOO_BIG_FOR_UNCOMPRESS)
    return result
}

    
// IsServerErrorZlibZMemError check mysql error is "ZLIB: Not enough memory" 
func IsServerErrorZlibZMemError(err error) bool {
    result := Isa(err, ER_ZLIB_Z_MEM_ERROR)
    return result
}

    
// IsServerErrorZlibZBufError check mysql error is "ZLIB: Not enough room in the output buffer (probably,length of uncompressed data was corrupted)" 
func IsServerErrorZlibZBufError(err error) bool {
    result := Isa(err, ER_ZLIB_Z_BUF_ERROR)
    return result
}

    
// IsServerErrorZlibZDataError check mysql error is "ZLIB: Input data corrupted" 
func IsServerErrorZlibZDataError(err error) bool {
    result := Isa(err, ER_ZLIB_Z_DATA_ERROR)
    return result
}

    
// IsServerErrorCutValueGroupConcat check mysql error is "Row %u was cut by GROUP_CONCAT()" 
func IsServerErrorCutValueGroupConcat(err error) bool {
    result := Isa(err, ER_CUT_VALUE_GROUP_CONCAT)
    return result
}

    
// IsServerErrorWarnTooFewRecords check mysql error is "Row %ld doesn't contain data for all columns" 
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

    
// IsServerErrorWarnDataOutOfRange check mysql error is "Out of range value for column '%s' at row %ld" 
func IsServerErrorWarnDataOutOfRange(err error) bool {
    result := Isa(err, ER_WARN_DATA_OUT_OF_RANGE)
    return result
}

    
// IsWarnDataTruncated check mysql error is "Data truncated for column '%s' at row %ld" 
func IsWarnDataTruncated(err error) bool {
    result := Isa(err, WARN_DATA_TRUNCATED)
    return result
}

    
// IsServerErrorWarnUsingOtherHandler check mysql error is "Using storage engine %s for table '%s'" 
func IsServerErrorWarnUsingOtherHandler(err error) bool {
    result := Isa(err, ER_WARN_USING_OTHER_HANDLER)
    return result
}

    
// IsServerErrorCantAggregate2collations check mysql error is "Illegal mix of collations (%s,%s) and (%s,%s) foroperation '%s'" 
func IsServerErrorCantAggregate2collations(err error) bool {
    result := Isa(err, ER_CANT_AGGREGATE_2COLLATIONS)
    return result
}

    
// IsServerErrorDropUser check mysql error is "Cannot drop one or more of the requested users" 
func IsServerErrorDropUser(err error) bool {
    result := Isa(err, ER_DROP_USER)
    return result
}

    
// IsServerErrorRevokeGrants check mysql error is "Can't revoke all privileges for one or more of therequested users" 
func IsServerErrorRevokeGrants(err error) bool {
    result := Isa(err, ER_REVOKE_GRANTS)
    return result
}

    
// IsServerErrorCantAggregate3collations check mysql error is "Illegal mix of collations (%s,%s), (%s,%s), (%s,%s) foroperation '%s'" 
func IsServerErrorCantAggregate3collations(err error) bool {
    result := Isa(err, ER_CANT_AGGREGATE_3COLLATIONS)
    return result
}

    
// IsServerErrorCantAggregateNcollations check mysql error is "Illegal mix of collations for operation '%s'" 
func IsServerErrorCantAggregateNcollations(err error) bool {
    result := Isa(err, ER_CANT_AGGREGATE_NCOLLATIONS)
    return result
}

    
// IsServerErrorVariableIsNotStruct check mysql error is "Variable '%s' is not a variable component (can't be usedas XXXX.variable_name)" 
func IsServerErrorVariableIsNotStruct(err error) bool {
    result := Isa(err, ER_VARIABLE_IS_NOT_STRUCT)
    return result
}

    
// IsServerErrorUnknownCollation check mysql error is "Unknown collation: '%s'" 
func IsServerErrorUnknownCollation(err error) bool {
    result := Isa(err, ER_UNKNOWN_COLLATION)
    return result
}

    
// IsServerErrorSlaveIgnoredSslParams check mysql error is "SSL parameters in CHANGE MASTER are ignored because thisMySQL slave was compiled without SSL support" 
func IsServerErrorSlaveIgnoredSslParams(err error) bool {
    result := Isa(err, ER_SLAVE_IGNORED_SSL_PARAMS)
    return result
}

    
// IsServerErrorServerIsInSecureAuthMode check mysql error is "Server is running in --secure-auth mode, but '%s'@'%s'has a password in the old format" 
func IsServerErrorServerIsInSecureAuthMode(err error) bool {
    result := Isa(err, ER_SERVER_IS_IN_SECURE_AUTH_MODE)
    return result
}

    
// IsServerErrorWarnFieldResolved check mysql error is "Field or reference '%s%s%s%s%s' of SELECT #%d wasresolved in SELECT #%d" 
func IsServerErrorWarnFieldResolved(err error) bool {
    result := Isa(err, ER_WARN_FIELD_RESOLVED)
    return result
}

    
// IsServerErrorBadSlaveUntilCond check mysql error is "Incorrect parameter or combination of parameters forSTART SLAVE UNTIL" 
func IsServerErrorBadSlaveUntilCond(err error) bool {
    result := Isa(err, ER_BAD_SLAVE_UNTIL_COND)
    return result
}

    
// IsServerErrorMissingSkipSlave check mysql error is "It is recommended to use --skip-slave-start when doingstep-by-step replication with START SLAVE UNTIL" 
func IsServerErrorMissingSkipSlave(err error) bool {
    result := Isa(err, ER_MISSING_SKIP_SLAVE)
    return result
}

    
// IsServerErrorUntilCondIgnored check mysql error is "SQL thread is not to be started so UNTIL options areignored" 
func IsServerErrorUntilCondIgnored(err error) bool {
    result := Isa(err, ER_UNTIL_COND_IGNORED)
    return result
}

    
// IsServerErrorWrongNameForIndex check mysql error is "Incorrect index name '%s'" 
func IsServerErrorWrongNameForIndex(err error) bool {
    result := Isa(err, ER_WRONG_NAME_FOR_INDEX)
    return result
}

    
// IsServerErrorWrongNameForCatalog check mysql error is "Incorrect catalog name '%s'" 
func IsServerErrorWrongNameForCatalog(err error) bool {
    result := Isa(err, ER_WRONG_NAME_FOR_CATALOG)
    return result
}

    
// IsServerErrorWarnQcResize check mysql error is "Query cache failed to set size %lu" 
func IsServerErrorWarnQcResize(err error) bool {
    result := Isa(err, ER_WARN_QC_RESIZE)
    return result
}

    
// IsServerErrorBadFtColumn check mysql error is "Column '%s' cannot be part of FULLTEXT index" 
func IsServerErrorBadFtColumn(err error) bool {
    result := Isa(err, ER_BAD_FT_COLUMN)
    return result
}

    
// IsServerErrorUnknownKeyCache check mysql error is "Unknown key cache '%s'" 
func IsServerErrorUnknownKeyCache(err error) bool {
    result := Isa(err, ER_UNKNOWN_KEY_CACHE)
    return result
}

    
// IsServerErrorWarnHostnameWontWork check mysql error is "MySQL is started in --skip-name-resolve mode" 
func IsServerErrorWarnHostnameWontWork(err error) bool {
    result := Isa(err, ER_WARN_HOSTNAME_WONT_WORK)
    return result
}

    
// IsServerErrorUnknownStorageEngine check mysql error is "Unknown storage engine '%s'" 
func IsServerErrorUnknownStorageEngine(err error) bool {
    result := Isa(err, ER_UNKNOWN_STORAGE_ENGINE)
    return result
}

    
// IsServerErrorWarnDeprecatedSyntax check mysql error is "'%s' is deprecated and will be removed in a futurerelease. Please use %s instead" 
func IsServerErrorWarnDeprecatedSyntax(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYNTAX)
    return result
}

    
// IsServerErrorNonUpdatableTable check mysql error is "The target table %s of the %s is not updatable" 
func IsServerErrorNonUpdatableTable(err error) bool {
    result := Isa(err, ER_NON_UPDATABLE_TABLE)
    return result
}

    
// IsServerErrorFeatureDisabled check mysql error is "The '%s' feature is disabled" 
func IsServerErrorFeatureDisabled(err error) bool {
    result := Isa(err, ER_FEATURE_DISABLED)
    return result
}

    
// IsServerErrorOptionPreventsStatement check mysql error is "The MySQL server is running with the %s option so itcannot execute this statement" 
func IsServerErrorOptionPreventsStatement(err error) bool {
    result := Isa(err, ER_OPTION_PREVENTS_STATEMENT)
    return result
}

    
// IsServerErrorDuplicatedValueInType check mysql error is "Column '%s' has duplicated value '%s' in %s" 
func IsServerErrorDuplicatedValueInType(err error) bool {
    result := Isa(err, ER_DUPLICATED_VALUE_IN_TYPE)
    return result
}

    
// IsServerErrorTruncatedWrongValue check mysql error is "Truncated incorrect %s value: '%s'" 
func IsServerErrorTruncatedWrongValue(err error) bool {
    result := Isa(err, ER_TRUNCATED_WRONG_VALUE)
    return result
}

    
// IsServerErrorTooMuchAutoTimestampCols check mysql error is "Incorrect table definition" 
func IsServerErrorTooMuchAutoTimestampCols(err error) bool {
    result := Isa(err, ER_TOO_MUCH_AUTO_TIMESTAMP_COLS)
    return result
}

    
// IsServerErrorInvalidOnUpdate check mysql error is "Invalid ON UPDATE clause for '%s' column" 
func IsServerErrorInvalidOnUpdate(err error) bool {
    result := Isa(err, ER_INVALID_ON_UPDATE)
    return result
}

    
// IsServerErrorUnsupportedPs check mysql error is "This command is not supported in the prepared statementprotocol yet" 
func IsServerErrorUnsupportedPs(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_PS)
    return result
}

    
// IsServerErrorGetErrmsg check mysql error is "Got error %d '%s' from %s" 
func IsServerErrorGetErrmsg(err error) bool {
    result := Isa(err, ER_GET_ERRMSG)
    return result
}

    
// IsServerErrorGetTemporaryErrmsg check mysql error is "Got temporary error %d '%s' from %s" 
func IsServerErrorGetTemporaryErrmsg(err error) bool {
    result := Isa(err, ER_GET_TEMPORARY_ERRMSG)
    return result
}

    
// IsServerErrorUnknownTimeZone check mysql error is "Unknown or incorrect time zone: '%s'" 
func IsServerErrorUnknownTimeZone(err error) bool {
    result := Isa(err, ER_UNKNOWN_TIME_ZONE)
    return result
}

    
// IsServerErrorWarnInvalidTimestamp check mysql error is "Invalid TIMESTAMP value in column '%s' at row %ld" 
func IsServerErrorWarnInvalidTimestamp(err error) bool {
    result := Isa(err, ER_WARN_INVALID_TIMESTAMP)
    return result
}

    
// IsServerErrorInvalidCharacterString check mysql error is "Invalid %s character string: '%s'" 
func IsServerErrorInvalidCharacterString(err error) bool {
    result := Isa(err, ER_INVALID_CHARACTER_STRING)
    return result
}

    
// IsServerErrorWarnAllowedPacketOverflowed check mysql error is "Result of %s() was larger than max_allowed_packet (%ld) -truncated" 
func IsServerErrorWarnAllowedPacketOverflowed(err error) bool {
    result := Isa(err, ER_WARN_ALLOWED_PACKET_OVERFLOWED)
    return result
}

    
// IsServerErrorConflictingDeclarations check mysql error is "Conflicting declarations: '%s%s' and '%s%s'" 
func IsServerErrorConflictingDeclarations(err error) bool {
    result := Isa(err, ER_CONFLICTING_DECLARATIONS)
    return result
}

    
// IsServerErrorSpNoRecursiveCreate check mysql error is "Can't create a %s from within another stored routine" 
func IsServerErrorSpNoRecursiveCreate(err error) bool {
    result := Isa(err, ER_SP_NO_RECURSIVE_CREATE)
    return result
}

    
// IsServerErrorSpAlreadyExists check mysql error is "%s %s already exists" 
func IsServerErrorSpAlreadyExists(err error) bool {
    result := Isa(err, ER_SP_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorSpDoesNotExist check mysql error is "%s %s does not exist" 
func IsServerErrorSpDoesNotExist(err error) bool {
    result := Isa(err, ER_SP_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorSpDropFailed check mysql error is "Failed to DROP %s %s" 
func IsServerErrorSpDropFailed(err error) bool {
    result := Isa(err, ER_SP_DROP_FAILED)
    return result
}

    
// IsServerErrorSpStoreFailed check mysql error is "Failed to CREATE %s %s" 
func IsServerErrorSpStoreFailed(err error) bool {
    result := Isa(err, ER_SP_STORE_FAILED)
    return result
}

    
// IsServerErrorSpLilabelMismatch check mysql error is "%s with no matching label: %s" 
func IsServerErrorSpLilabelMismatch(err error) bool {
    result := Isa(err, ER_SP_LILABEL_MISMATCH)
    return result
}

    
// IsServerErrorSpLabelRedefine check mysql error is "Redefining label %s" 
func IsServerErrorSpLabelRedefine(err error) bool {
    result := Isa(err, ER_SP_LABEL_REDEFINE)
    return result
}

    
// IsServerErrorSpLabelMismatch check mysql error is "End-label %s without match" 
func IsServerErrorSpLabelMismatch(err error) bool {
    result := Isa(err, ER_SP_LABEL_MISMATCH)
    return result
}

    
// IsServerErrorSpUninitVar check mysql error is "Referring to uninitialized variable %s" 
func IsServerErrorSpUninitVar(err error) bool {
    result := Isa(err, ER_SP_UNINIT_VAR)
    return result
}

    
// IsServerErrorSpBadselect check mysql error is "PROCEDURE %s can't return a result set in the givencontext" 
func IsServerErrorSpBadselect(err error) bool {
    result := Isa(err, ER_SP_BADSELECT)
    return result
}

    
// IsServerErrorSpBadreturn check mysql error is "RETURN is only allowed in a FUNCTION" 
func IsServerErrorSpBadreturn(err error) bool {
    result := Isa(err, ER_SP_BADRETURN)
    return result
}

    
// IsServerErrorSpBadstatement check mysql error is "%s is not allowed in stored procedures" 
func IsServerErrorSpBadstatement(err error) bool {
    result := Isa(err, ER_SP_BADSTATEMENT)
    return result
}

    
// IsServerErrorUpdateLogDeprecatedIgnored check mysql error is "The update log is deprecated and replaced by the binarylog" 
func IsServerErrorUpdateLogDeprecatedIgnored(err error) bool {
    result := Isa(err, ER_UPDATE_LOG_DEPRECATED_IGNORED)
    return result
}

    
// IsServerErrorUpdateLogDeprecatedTranslated check mysql error is "The update log is deprecated and replaced by the binarylog" 
func IsServerErrorUpdateLogDeprecatedTranslated(err error) bool {
    result := Isa(err, ER_UPDATE_LOG_DEPRECATED_TRANSLATED)
    return result
}

    
// IsServerErrorQueryInterrupted check mysql error is "Query execution was interrupted" 
func IsServerErrorQueryInterrupted(err error) bool {
    result := Isa(err, ER_QUERY_INTERRUPTED)
    return result
}

    
// IsServerErrorSpWrongNoOfArgs check mysql error is "Incorrect number of arguments for %s %s" 
func IsServerErrorSpWrongNoOfArgs(err error) bool {
    result := Isa(err, ER_SP_WRONG_NO_OF_ARGS)
    return result
}

    
// IsServerErrorSpCondMismatch check mysql error is "Undefined CONDITION: %s" 
func IsServerErrorSpCondMismatch(err error) bool {
    result := Isa(err, ER_SP_COND_MISMATCH)
    return result
}

    
// IsServerErrorSpNoreturn check mysql error is "No RETURN found in FUNCTION %s" 
func IsServerErrorSpNoreturn(err error) bool {
    result := Isa(err, ER_SP_NORETURN)
    return result
}

    
// IsServerErrorSpNoreturnend check mysql error is "FUNCTION %s ended without RETURN" 
func IsServerErrorSpNoreturnend(err error) bool {
    result := Isa(err, ER_SP_NORETURNEND)
    return result
}

    
// IsServerErrorSpBadCursorQuery check mysql error is "Cursor statement must be a SELECT" 
func IsServerErrorSpBadCursorQuery(err error) bool {
    result := Isa(err, ER_SP_BAD_CURSOR_QUERY)
    return result
}

    
// IsServerErrorSpBadCursorSelect check mysql error is "Cursor SELECT must not have INTO" 
func IsServerErrorSpBadCursorSelect(err error) bool {
    result := Isa(err, ER_SP_BAD_CURSOR_SELECT)
    return result
}

    
// IsServerErrorSpCursorMismatch check mysql error is "Undefined CURSOR: %s" 
func IsServerErrorSpCursorMismatch(err error) bool {
    result := Isa(err, ER_SP_CURSOR_MISMATCH)
    return result
}

    
// IsServerErrorSpCursorAlreadyOpen check mysql error is "Cursor is already open" 
func IsServerErrorSpCursorAlreadyOpen(err error) bool {
    result := Isa(err, ER_SP_CURSOR_ALREADY_OPEN)
    return result
}

    
// IsServerErrorSpCursorNotOpen check mysql error is "Cursor is not open" 
func IsServerErrorSpCursorNotOpen(err error) bool {
    result := Isa(err, ER_SP_CURSOR_NOT_OPEN)
    return result
}

    
// IsServerErrorSpUndeclaredVar check mysql error is "Undeclared variable: %s" 
func IsServerErrorSpUndeclaredVar(err error) bool {
    result := Isa(err, ER_SP_UNDECLARED_VAR)
    return result
}

    
// IsServerErrorSpWrongNoOfFetchArgs check mysql error is "Incorrect number of FETCH variables" 
func IsServerErrorSpWrongNoOfFetchArgs(err error) bool {
    result := Isa(err, ER_SP_WRONG_NO_OF_FETCH_ARGS)
    return result
}

    
// IsServerErrorSpFetchNoData check mysql error is "No data - zero rows fetched, selected, or processed" 
func IsServerErrorSpFetchNoData(err error) bool {
    result := Isa(err, ER_SP_FETCH_NO_DATA)
    return result
}

    
// IsServerErrorSpDupParam check mysql error is "Duplicate parameter: %s" 
func IsServerErrorSpDupParam(err error) bool {
    result := Isa(err, ER_SP_DUP_PARAM)
    return result
}

    
// IsServerErrorSpDupVar check mysql error is "Duplicate variable: %s" 
func IsServerErrorSpDupVar(err error) bool {
    result := Isa(err, ER_SP_DUP_VAR)
    return result
}

    
// IsServerErrorSpDupCond check mysql error is "Duplicate condition: %s" 
func IsServerErrorSpDupCond(err error) bool {
    result := Isa(err, ER_SP_DUP_COND)
    return result
}

    
// IsServerErrorSpDupCurs check mysql error is "Duplicate cursor: %s" 
func IsServerErrorSpDupCurs(err error) bool {
    result := Isa(err, ER_SP_DUP_CURS)
    return result
}

    
// IsServerErrorSpCantAlter check mysql error is "Failed to ALTER %s %s" 
func IsServerErrorSpCantAlter(err error) bool {
    result := Isa(err, ER_SP_CANT_ALTER)
    return result
}

    
// IsServerErrorSpSubselectNyi check mysql error is "Subquery value not supported" 
func IsServerErrorSpSubselectNyi(err error) bool {
    result := Isa(err, ER_SP_SUBSELECT_NYI)
    return result
}

    
// IsServerErrorStmtNotAllowedInSfOrTrg check mysql error is "%s is not allowed in stored function or trigger" 
func IsServerErrorStmtNotAllowedInSfOrTrg(err error) bool {
    result := Isa(err, ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG)
    return result
}

    
// IsServerErrorSpVarcondAfterCurshndlr check mysql error is "Variable or condition declaration after cursor or handlerdeclaration" 
func IsServerErrorSpVarcondAfterCurshndlr(err error) bool {
    result := Isa(err, ER_SP_VARCOND_AFTER_CURSHNDLR)
    return result
}

    
// IsServerErrorSpCursorAfterHandler check mysql error is "Cursor declaration after handler declaration" 
func IsServerErrorSpCursorAfterHandler(err error) bool {
    result := Isa(err, ER_SP_CURSOR_AFTER_HANDLER)
    return result
}

    
// IsServerErrorSpCaseNotFound check mysql error is "Case not found for CASE statement" 
func IsServerErrorSpCaseNotFound(err error) bool {
    result := Isa(err, ER_SP_CASE_NOT_FOUND)
    return result
}

    
// IsServerErrorFparserTooBigFile check mysql error is "Configuration file '%s' is too big" 
func IsServerErrorFparserTooBigFile(err error) bool {
    result := Isa(err, ER_FPARSER_TOO_BIG_FILE)
    return result
}

    
// IsServerErrorFparserBadHeader check mysql error is "Malformed file type header in file '%s'" 
func IsServerErrorFparserBadHeader(err error) bool {
    result := Isa(err, ER_FPARSER_BAD_HEADER)
    return result
}

    
// IsServerErrorFparserEofInComment check mysql error is "Unexpected end of file while parsing comment '%s'" 
func IsServerErrorFparserEofInComment(err error) bool {
    result := Isa(err, ER_FPARSER_EOF_IN_COMMENT)
    return result
}

    
// IsServerErrorFparserErrorInParameter check mysql error is "Error while parsing parameter '%s' (line: '%s')" 
func IsServerErrorFparserErrorInParameter(err error) bool {
    result := Isa(err, ER_FPARSER_ERROR_IN_PARAMETER)
    return result
}

    
// IsServerErrorFparserEofInUnknownParameter check mysql error is "Unexpected end of file while skipping unknown parameter'%s'" 
func IsServerErrorFparserEofInUnknownParameter(err error) bool {
    result := Isa(err, ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER)
    return result
}

    
// IsServerErrorViewNoExplain check mysql error is "EXPLAIN/SHOW can not be issued" 
func IsServerErrorViewNoExplain(err error) bool {
    result := Isa(err, ER_VIEW_NO_EXPLAIN)
    return result
}

    
// IsServerErrorFrmUnknownType check mysql error is "File '%s' has unknown type '%s' in its header" 
func IsServerErrorFrmUnknownType(err error) bool {
    result := Isa(err, ER_FRM_UNKNOWN_TYPE)
    return result
}

    
// IsServerErrorWrongObject check mysql error is "'%s.%s' is not %s" 
func IsServerErrorWrongObject(err error) bool {
    result := Isa(err, ER_WRONG_OBJECT)
    return result
}

    
// IsServerErrorNonupdateableColumn check mysql error is "Column '%s' is not updatable" 
func IsServerErrorNonupdateableColumn(err error) bool {
    result := Isa(err, ER_NONUPDATEABLE_COLUMN)
    return result
}

    
// IsServerErrorViewSelectDerived check mysql error is "View's SELECT contains a subquery in the FROM clause" 
func IsServerErrorViewSelectDerived(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_DERIVED)
    return result
}

    
// IsServerErrorViewSelectClause check mysql error is "View's SELECT contains a '%s' clause" 
func IsServerErrorViewSelectClause(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_CLAUSE)
    return result
}

    
// IsServerErrorViewSelectVariable check mysql error is "View's SELECT contains a variable or parameter" 
func IsServerErrorViewSelectVariable(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_VARIABLE)
    return result
}

    
// IsServerErrorViewSelectTmptable check mysql error is "View's SELECT refers to a temporary table '%s'" 
func IsServerErrorViewSelectTmptable(err error) bool {
    result := Isa(err, ER_VIEW_SELECT_TMPTABLE)
    return result
}

    
// IsServerErrorViewWrongList check mysql error is "View's SELECT and view's field list have different columncounts" 
func IsServerErrorViewWrongList(err error) bool {
    result := Isa(err, ER_VIEW_WRONG_LIST)
    return result
}

    
// IsServerErrorWarnViewMerge check mysql error is "View merge algorithm can't be used here for now (assumedundefined algorithm)" 
func IsServerErrorWarnViewMerge(err error) bool {
    result := Isa(err, ER_WARN_VIEW_MERGE)
    return result
}

    
// IsServerErrorWarnViewWithoutKey check mysql error is "View being updated does not have complete key ofunderlying table in it" 
func IsServerErrorWarnViewWithoutKey(err error) bool {
    result := Isa(err, ER_WARN_VIEW_WITHOUT_KEY)
    return result
}

    
// IsServerErrorViewInvalid check mysql error is "View '%s.%s' references invalid table(s) or column(s) orfunction(s) or definer/invoker of view lack rights to use them" 
func IsServerErrorViewInvalid(err error) bool {
    result := Isa(err, ER_VIEW_INVALID)
    return result
}

    
// IsServerErrorSpNoDropSp check mysql error is "Can't drop or alter a %s from within another storedroutine" 
func IsServerErrorSpNoDropSp(err error) bool {
    result := Isa(err, ER_SP_NO_DROP_SP)
    return result
}

    
// IsServerErrorSpGotoInHndlr check mysql error is "GOTO is not allowed in a stored procedure handler" 
func IsServerErrorSpGotoInHndlr(err error) bool {
    result := Isa(err, ER_SP_GOTO_IN_HNDLR)
    return result
}

    
// IsServerErrorTrgAlreadyExists check mysql error is "Trigger already exists" 
func IsServerErrorTrgAlreadyExists(err error) bool {
    result := Isa(err, ER_TRG_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorTrgDoesNotExist check mysql error is "Trigger does not exist" 
func IsServerErrorTrgDoesNotExist(err error) bool {
    result := Isa(err, ER_TRG_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorTrgOnViewOrTempTable check mysql error is "Trigger's '%s' is view or temporary table" 
func IsServerErrorTrgOnViewOrTempTable(err error) bool {
    result := Isa(err, ER_TRG_ON_VIEW_OR_TEMP_TABLE)
    return result
}

    
// IsServerErrorTrgCantChangeRow check mysql error is "Updating of %s row is not allowed in %strigger" 
func IsServerErrorTrgCantChangeRow(err error) bool {
    result := Isa(err, ER_TRG_CANT_CHANGE_ROW)
    return result
}

    
// IsServerErrorTrgNoSuchRowInTrg check mysql error is "There is no %s row in %s trigger" 
func IsServerErrorTrgNoSuchRowInTrg(err error) bool {
    result := Isa(err, ER_TRG_NO_SUCH_ROW_IN_TRG)
    return result
}

    
// IsServerErrorNoDefaultForField check mysql error is "Field '%s' doesn't have a default value" 
func IsServerErrorNoDefaultForField(err error) bool {
    result := Isa(err, ER_NO_DEFAULT_FOR_FIELD)
    return result
}

    
// IsServerErrorDivisionByZero check mysql error is "Division by 0" 
func IsServerErrorDivisionByZero(err error) bool {
    result := Isa(err, ER_DIVISION_BY_ZERO)
    return result
}

    
// IsServerErrorTruncatedWrongValueForField check mysql error is "Incorrect %s value: '%s' for column '%s' at row %ld" 
func IsServerErrorTruncatedWrongValueForField(err error) bool {
    result := Isa(err, ER_TRUNCATED_WRONG_VALUE_FOR_FIELD)
    return result
}

    
// IsServerErrorIllegalValueForType check mysql error is "Illegal %s '%s' value found during parsing" 
func IsServerErrorIllegalValueForType(err error) bool {
    result := Isa(err, ER_ILLEGAL_VALUE_FOR_TYPE)
    return result
}

    
// IsServerErrorViewNonupdCheck check mysql error is "CHECK OPTION on non-updatable view '%s.%s'" 
func IsServerErrorViewNonupdCheck(err error) bool {
    result := Isa(err, ER_VIEW_NONUPD_CHECK)
    return result
}

    
// IsServerErrorViewCheckFailed check mysql error is "CHECK OPTION failed '%s.%s'" 
func IsServerErrorViewCheckFailed(err error) bool {
    result := Isa(err, ER_VIEW_CHECK_FAILED)
    return result
}

    
// IsServerErrorProcaccessDeniedError check mysql error is "%s command denied to user '%s'@'%s' for routine '%s'" 
func IsServerErrorProcaccessDeniedError(err error) bool {
    result := Isa(err, ER_PROCACCESS_DENIED_ERROR)
    return result
}

    
// IsServerErrorRelayLogFail check mysql error is "Failed purging old relay logs: %s" 
func IsServerErrorRelayLogFail(err error) bool {
    result := Isa(err, ER_RELAY_LOG_FAIL)
    return result
}

    
// IsServerErrorPasswdLength check mysql error is "Password hash should be a %d-digit hexadecimal number" 
func IsServerErrorPasswdLength(err error) bool {
    result := Isa(err, ER_PASSWD_LENGTH)
    return result
}

    
// IsServerErrorUnknownTargetBinlog check mysql error is "Target log not found in binlog index" 
func IsServerErrorUnknownTargetBinlog(err error) bool {
    result := Isa(err, ER_UNKNOWN_TARGET_BINLOG)
    return result
}

    
// IsServerErrorIoErrLogIndexRead check mysql error is "I/O error reading log index file" 
func IsServerErrorIoErrLogIndexRead(err error) bool {
    result := Isa(err, ER_IO_ERR_LOG_INDEX_READ)
    return result
}

    
// IsServerErrorBinlogPurgeProhibited check mysql error is "Server configuration does not permit binlog purge" 
func IsServerErrorBinlogPurgeProhibited(err error) bool {
    result := Isa(err, ER_BINLOG_PURGE_PROHIBITED)
    return result
}

    
// IsServerErrorFseekFail check mysql error is "Failed on fseek()" 
func IsServerErrorFseekFail(err error) bool {
    result := Isa(err, ER_FSEEK_FAIL)
    return result
}

    
// IsServerErrorBinlogPurgeFatalErr check mysql error is "Fatal error during log purge" 
func IsServerErrorBinlogPurgeFatalErr(err error) bool {
    result := Isa(err, ER_BINLOG_PURGE_FATAL_ERR)
    return result
}

    
// IsServerErrorLogInUse check mysql error is "A purgeable log is in use, will not purge" 
func IsServerErrorLogInUse(err error) bool {
    result := Isa(err, ER_LOG_IN_USE)
    return result
}

    
// IsServerErrorLogPurgeUnknownErr check mysql error is "Unknown error during log purge" 
func IsServerErrorLogPurgeUnknownErr(err error) bool {
    result := Isa(err, ER_LOG_PURGE_UNKNOWN_ERR)
    return result
}

    
// IsServerErrorRelayLogInit check mysql error is "Failed initializing relay log position: %s" 
func IsServerErrorRelayLogInit(err error) bool {
    result := Isa(err, ER_RELAY_LOG_INIT)
    return result
}

    
// IsServerErrorNoBinaryLogging check mysql error is "You are not using binary logging" 
func IsServerErrorNoBinaryLogging(err error) bool {
    result := Isa(err, ER_NO_BINARY_LOGGING)
    return result
}

    
// IsServerErrorReservedSyntax check mysql error is "The '%s' syntax is reserved for purposes internal to theMySQL server" 
func IsServerErrorReservedSyntax(err error) bool {
    result := Isa(err, ER_RESERVED_SYNTAX)
    return result
}

    
// IsServerErrorWsasFailed check mysql error is "WSAStartup Failed" 
func IsServerErrorWsasFailed(err error) bool {
    result := Isa(err, ER_WSAS_FAILED)
    return result
}

    
// IsServerErrorDiffGroupsProc check mysql error is "Can't handle procedures with different groups yet" 
func IsServerErrorDiffGroupsProc(err error) bool {
    result := Isa(err, ER_DIFF_GROUPS_PROC)
    return result
}

    
// IsServerErrorNoGroupForProc check mysql error is "Select must have a group with this procedure" 
func IsServerErrorNoGroupForProc(err error) bool {
    result := Isa(err, ER_NO_GROUP_FOR_PROC)
    return result
}

    
// IsServerErrorOrderWithProc check mysql error is "Can't use ORDER clause with this procedure" 
func IsServerErrorOrderWithProc(err error) bool {
    result := Isa(err, ER_ORDER_WITH_PROC)
    return result
}

    
// IsServerErrorLoggingProhibitChangingOf check mysql error is "Binary logging and replication forbid changing the globalserver %s" 
func IsServerErrorLoggingProhibitChangingOf(err error) bool {
    result := Isa(err, ER_LOGGING_PROHIBIT_CHANGING_OF)
    return result
}

    
// IsServerErrorNoFileMapping check mysql error is "Can't map file: %s, errno: %d" 
func IsServerErrorNoFileMapping(err error) bool {
    result := Isa(err, ER_NO_FILE_MAPPING)
    return result
}

    
// IsServerErrorWrongMagic check mysql error is "Wrong magic in %s" 
func IsServerErrorWrongMagic(err error) bool {
    result := Isa(err, ER_WRONG_MAGIC)
    return result
}

    
// IsServerErrorPsManyParam check mysql error is "Prepared statement contains too many placeholders" 
func IsServerErrorPsManyParam(err error) bool {
    result := Isa(err, ER_PS_MANY_PARAM)
    return result
}

    
// IsServerErrorKeyPart0 check mysql error is "Key part '%s' length cannot be 0" 
func IsServerErrorKeyPart0(err error) bool {
    result := Isa(err, ER_KEY_PART_0)
    return result
}

    
// IsServerErrorViewChecksum check mysql error is "View text checksum failed" 
func IsServerErrorViewChecksum(err error) bool {
    result := Isa(err, ER_VIEW_CHECKSUM)
    return result
}

    
// IsServerErrorViewMultiupdate check mysql error is "Can not modify more than one base table through a joinview '%s.%s'" 
func IsServerErrorViewMultiupdate(err error) bool {
    result := Isa(err, ER_VIEW_MULTIUPDATE)
    return result
}

    
// IsServerErrorViewNoInsertFieldList check mysql error is "Can not insert into join view '%s.%s' without fields list" 
func IsServerErrorViewNoInsertFieldList(err error) bool {
    result := Isa(err, ER_VIEW_NO_INSERT_FIELD_LIST)
    return result
}

    
// IsServerErrorViewDeleteMergeView check mysql error is "Can not delete from join view '%s.%s'" 
func IsServerErrorViewDeleteMergeView(err error) bool {
    result := Isa(err, ER_VIEW_DELETE_MERGE_VIEW)
    return result
}

    
// IsServerErrorCannotUser check mysql error is "Operation %s failed for %s" 
func IsServerErrorCannotUser(err error) bool {
    result := Isa(err, ER_CANNOT_USER)
    return result
}

    
// IsServerErrorXaerNota check mysql error is "XAER_NOTA: Unknown XID" 
func IsServerErrorXaerNota(err error) bool {
    result := Isa(err, ER_XAER_NOTA)
    return result
}

    
// IsServerErrorXaerInval check mysql error is "XAER_INVAL: Invalid arguments (or unsupported command)" 
func IsServerErrorXaerInval(err error) bool {
    result := Isa(err, ER_XAER_INVAL)
    return result
}

    
// IsServerErrorXaerRmfail check mysql error is "XAER_RMFAIL: The command cannot be executed when globaltransaction is in the %s state" 
func IsServerErrorXaerRmfail(err error) bool {
    result := Isa(err, ER_XAER_RMFAIL)
    return result
}

    
// IsServerErrorXaerOutside check mysql error is "XAER_OUTSIDE: Some work is done outside globaltransaction" 
func IsServerErrorXaerOutside(err error) bool {
    result := Isa(err, ER_XAER_OUTSIDE)
    return result
}

    
// IsServerErrorXaerRmerr check mysql error is "XAER_RMERR: Fatal error occurred in the transactionbranch - check your data for consistency" 
func IsServerErrorXaerRmerr(err error) bool {
    result := Isa(err, ER_XAER_RMERR)
    return result
}

    
// IsServerErrorXaRbrollback check mysql error is "XA_RBROLLBACK: Transaction branch was rolled back" 
func IsServerErrorXaRbrollback(err error) bool {
    result := Isa(err, ER_XA_RBROLLBACK)
    return result
}

    
// IsServerErrorNonexistingProcGrant check mysql error is "There is no such grant defined for user '%s' on host '%s'on routine '%s'" 
func IsServerErrorNonexistingProcGrant(err error) bool {
    result := Isa(err, ER_NONEXISTING_PROC_GRANT)
    return result
}

    
// IsServerErrorProcAutoGrantFail check mysql error is "Failed to grant EXECUTE and ALTER ROUTINE privileges" 
func IsServerErrorProcAutoGrantFail(err error) bool {
    result := Isa(err, ER_PROC_AUTO_GRANT_FAIL)
    return result
}

    
// IsServerErrorProcAutoRevokeFail check mysql error is "Failed to revoke all privileges to dropped routine" 
func IsServerErrorProcAutoRevokeFail(err error) bool {
    result := Isa(err, ER_PROC_AUTO_REVOKE_FAIL)
    return result
}

    
// IsServerErrorDataTooLong check mysql error is "Data too long for column '%s' at row %ld" 
func IsServerErrorDataTooLong(err error) bool {
    result := Isa(err, ER_DATA_TOO_LONG)
    return result
}

    
// IsServerErrorSpBadSqlstate check mysql error is "Bad SQLSTATE: '%s'" 
func IsServerErrorSpBadSqlstate(err error) bool {
    result := Isa(err, ER_SP_BAD_SQLSTATE)
    return result
}

    
// IsServerErrorStartup check mysql error is "%s: ready for connections. Version: '%s' socket: '%s'port: %d %s" 
func IsServerErrorStartup(err error) bool {
    result := Isa(err, ER_STARTUP)
    return result
}

    
// IsServerErrorLoadFromFixedSizeRowsToVar check mysql error is "Can't load value from file with fixed size rows tovariable" 
func IsServerErrorLoadFromFixedSizeRowsToVar(err error) bool {
    result := Isa(err, ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR)
    return result
}

    
// IsServerErrorCantCreateUserWithGrant check mysql error is "You are not allowed to create a user with GRANT" 
func IsServerErrorCantCreateUserWithGrant(err error) bool {
    result := Isa(err, ER_CANT_CREATE_USER_WITH_GRANT)
    return result
}

    
// IsServerErrorWrongValueForType check mysql error is "Incorrect %s value: '%s' for function %s" 
func IsServerErrorWrongValueForType(err error) bool {
    result := Isa(err, ER_WRONG_VALUE_FOR_TYPE)
    return result
}

    
// IsServerErrorTableDefChanged check mysql error is "Table definition has changed, please retry transaction" 
func IsServerErrorTableDefChanged(err error) bool {
    result := Isa(err, ER_TABLE_DEF_CHANGED)
    return result
}

    
// IsServerErrorSpDupHandler check mysql error is "Duplicate handler declared in the same block" 
func IsServerErrorSpDupHandler(err error) bool {
    result := Isa(err, ER_SP_DUP_HANDLER)
    return result
}

    
// IsServerErrorSpNotVarArg check mysql error is "OUT or INOUT argument %d for routine %s is not a variableor NEW pseudo-variable in BEFORE trigger" 
func IsServerErrorSpNotVarArg(err error) bool {
    result := Isa(err, ER_SP_NOT_VAR_ARG)
    return result
}

    
// IsServerErrorSpNoRetset check mysql error is "Not allowed to return a result set from a %s" 
func IsServerErrorSpNoRetset(err error) bool {
    result := Isa(err, ER_SP_NO_RETSET)
    return result
}

    
// IsServerErrorCantCreateGeometryObject check mysql error is "Cannot get geometry object from data you send to theGEOMETRY field" 
func IsServerErrorCantCreateGeometryObject(err error) bool {
    result := Isa(err, ER_CANT_CREATE_GEOMETRY_OBJECT)
    return result
}

    
// IsServerErrorFailedRoutineBreakBinlog check mysql error is "A routine failed and has neither NO SQL nor READS SQLDATA in its declaration and binary logging is enabled" 
func IsServerErrorFailedRoutineBreakBinlog(err error) bool {
    result := Isa(err, ER_FAILED_ROUTINE_BREAK_BINLOG)
    return result
}

    
// IsServerErrorBinlogUnsafeRoutine check mysql error is "This function has none of DETERMINISTIC, NO SQL, or READSSQL DATA in its declaration and binary logging is enabled (you*might* want to use the less safe log_bin_trust_function_creatorsvariable)" 
func IsServerErrorBinlogUnsafeRoutine(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_ROUTINE)
    return result
}

    
// IsServerErrorBinlogCreateRoutineNeedSuper check mysql error is "You do not have the SUPER privilege and binary logging isenabled (you *might* want to use the less safelog_bin_trust_function_creators variable)" 
func IsServerErrorBinlogCreateRoutineNeedSuper(err error) bool {
    result := Isa(err, ER_BINLOG_CREATE_ROUTINE_NEED_SUPER)
    return result
}

    
// IsServerErrorExecStmtWithOpenCursor check mysql error is "You can't execute a prepared statement which has an opencursor associated with it. Reset the statement to re-execute it." 
func IsServerErrorExecStmtWithOpenCursor(err error) bool {
    result := Isa(err, ER_EXEC_STMT_WITH_OPEN_CURSOR)
    return result
}

    
// IsServerErrorStmtHasNoOpenCursor check mysql error is "The statement (%lu) has no open cursor." 
func IsServerErrorStmtHasNoOpenCursor(err error) bool {
    result := Isa(err, ER_STMT_HAS_NO_OPEN_CURSOR)
    return result
}

    
// IsServerErrorCommitNotAllowedInSfOrTrg check mysql error is "Explicit or implicit commit is not allowed in storedfunction or trigger." 
func IsServerErrorCommitNotAllowedInSfOrTrg(err error) bool {
    result := Isa(err, ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG)
    return result
}

    
// IsServerErrorNoDefaultForViewField check mysql error is "Field of view '%s.%s' underlying table doesn't have adefault value" 
func IsServerErrorNoDefaultForViewField(err error) bool {
    result := Isa(err, ER_NO_DEFAULT_FOR_VIEW_FIELD)
    return result
}

    
// IsServerErrorSpNoRecursion check mysql error is "Recursive stored functions and triggers are not allowed." 
func IsServerErrorSpNoRecursion(err error) bool {
    result := Isa(err, ER_SP_NO_RECURSION)
    return result
}

    
// IsServerErrorTooBigScale check mysql error is "Too big scale %d specified for column '%s'. Maximum is%lu." 
func IsServerErrorTooBigScale(err error) bool {
    result := Isa(err, ER_TOO_BIG_SCALE)
    return result
}

    
// IsServerErrorTooBigPrecision check mysql error is "Too big precision %d specified for column '%s'. Maximumis %lu." 
func IsServerErrorTooBigPrecision(err error) bool {
    result := Isa(err, ER_TOO_BIG_PRECISION)
    return result
}

    
// IsServerErrorMBiggerThanD check mysql error is "For float(M,D), double(M,D) or decimal(M,D), M must be>= D (column '%s')." 
func IsServerErrorMBiggerThanD(err error) bool {
    result := Isa(err, ER_M_BIGGER_THAN_D)
    return result
}

    
// IsServerErrorWrongLockOfSystemTable check mysql error is "You can't combine write-locking of system tables withother tables or lock types" 
func IsServerErrorWrongLockOfSystemTable(err error) bool {
    result := Isa(err, ER_WRONG_LOCK_OF_SYSTEM_TABLE)
    return result
}

    
// IsServerErrorConnectToForeignDataSource check mysql error is "Unable to connect to foreign data source: %s" 
func IsServerErrorConnectToForeignDataSource(err error) bool {
    result := Isa(err, ER_CONNECT_TO_FOREIGN_DATA_SOURCE)
    return result
}

    
// IsServerErrorQueryOnForeignDataSource check mysql error is "There was a problem processing the query on the foreigndata source. Data source error: %s" 
func IsServerErrorQueryOnForeignDataSource(err error) bool {
    result := Isa(err, ER_QUERY_ON_FOREIGN_DATA_SOURCE)
    return result
}

    
// IsServerErrorForeignDataSourceDoesntExist check mysql error is "The foreign data source you are trying to reference doesnot exist. Data source error: %s" 
func IsServerErrorForeignDataSourceDoesntExist(err error) bool {
    result := Isa(err, ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST)
    return result
}

    
// IsServerErrorForeignDataStringInvalidCantCreate check mysql error is "Can't create federated table. The data source connectionstring '%s' is not in the correct format" 
func IsServerErrorForeignDataStringInvalidCantCreate(err error) bool {
    result := Isa(err, ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE)
    return result
}

    
// IsServerErrorForeignDataStringInvalid check mysql error is "The data source connection string '%s' is not in thecorrect format" 
func IsServerErrorForeignDataStringInvalid(err error) bool {
    result := Isa(err, ER_FOREIGN_DATA_STRING_INVALID)
    return result
}

    
// IsServerErrorCantCreateFederatedTable check mysql error is "Can't create federated table. Foreign data src error: %s" 
func IsServerErrorCantCreateFederatedTable(err error) bool {
    result := Isa(err, ER_CANT_CREATE_FEDERATED_TABLE)
    return result
}

    
// IsServerErrorTrgInWrongSchema check mysql error is "Trigger in wrong schema" 
func IsServerErrorTrgInWrongSchema(err error) bool {
    result := Isa(err, ER_TRG_IN_WRONG_SCHEMA)
    return result
}

    
// IsServerErrorStackOverrunNeedMore check mysql error is "Thread stack overrun: %ld bytes used of a %ld byte stack,and %ld bytes needed. Use 'mysqld --thread_stack=#' to specify abigger stack." 
func IsServerErrorStackOverrunNeedMore(err error) bool {
    result := Isa(err, ER_STACK_OVERRUN_NEED_MORE)
    return result
}

    
// IsServerErrorTooLongBody check mysql error is "Routine body for '%s' is too long" 
func IsServerErrorTooLongBody(err error) bool {
    result := Isa(err, ER_TOO_LONG_BODY)
    return result
}

    
// IsServerErrorWarnCantDropDefaultKeycache check mysql error is "Cannot drop default keycache" 
func IsServerErrorWarnCantDropDefaultKeycache(err error) bool {
    result := Isa(err, ER_WARN_CANT_DROP_DEFAULT_KEYCACHE)
    return result
}

    
// IsServerErrorTooBigDisplaywidth check mysql error is "Display width out of range for column '%s' (max = %lu)" 
func IsServerErrorTooBigDisplaywidth(err error) bool {
    result := Isa(err, ER_TOO_BIG_DISPLAYWIDTH)
    return result
}

    
// IsServerErrorXaerDupid check mysql error is "XAER_DUPID: The XID already exists" 
func IsServerErrorXaerDupid(err error) bool {
    result := Isa(err, ER_XAER_DUPID)
    return result
}

    
// IsServerErrorDatetimeFunctionOverflow check mysql error is "Datetime function: %s field overflow" 
func IsServerErrorDatetimeFunctionOverflow(err error) bool {
    result := Isa(err, ER_DATETIME_FUNCTION_OVERFLOW)
    return result
}

    
// IsServerErrorCantUpdateUsedTableInSfOrTrg check mysql error is "Can't update table '%s' in stored function/triggerbecause it is already used by statement which invoked this storedfunction/trigger." 
func IsServerErrorCantUpdateUsedTableInSfOrTrg(err error) bool {
    result := Isa(err, ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG)
    return result
}

    
// IsServerErrorViewPreventUpdate check mysql error is "The definition of table '%s' prevents operation %s ontable '%s'." 
func IsServerErrorViewPreventUpdate(err error) bool {
    result := Isa(err, ER_VIEW_PREVENT_UPDATE)
    return result
}

    
// IsServerErrorPsNoRecursion check mysql error is "The prepared statement contains a stored routine callthat refers to that same statement. It's not allowed to execute aprepared statement in such a recursive manner" 
func IsServerErrorPsNoRecursion(err error) bool {
    result := Isa(err, ER_PS_NO_RECURSION)
    return result
}

    
// IsServerErrorSpCantSetAutocommit check mysql error is "Not allowed to set autocommit from a stored function ortrigger" 
func IsServerErrorSpCantSetAutocommit(err error) bool {
    result := Isa(err, ER_SP_CANT_SET_AUTOCOMMIT)
    return result
}

    
// IsServerErrorMalformedDefiner check mysql error is "Definer is not fully qualified" 
func IsServerErrorMalformedDefiner(err error) bool {
    result := Isa(err, ER_MALFORMED_DEFINER)
    return result
}

    
// IsServerErrorViewFrmNoUser check mysql error is "View '%s'.'%s' has no definer information (old tableformat). Current user is used as definer. Please recreate theview!" 
func IsServerErrorViewFrmNoUser(err error) bool {
    result := Isa(err, ER_VIEW_FRM_NO_USER)
    return result
}

    
// IsServerErrorViewOtherUser check mysql error is "You need the SUPER privilege for creation view with'%s'@'%s' definer" 
func IsServerErrorViewOtherUser(err error) bool {
    result := Isa(err, ER_VIEW_OTHER_USER)
    return result
}

    
// IsServerErrorNoSuchUser check mysql error is "The user specified as a definer ('%s'@'%s') does notexist" 
func IsServerErrorNoSuchUser(err error) bool {
    result := Isa(err, ER_NO_SUCH_USER)
    return result
}

    
// IsServerErrorForbidSchemaChange check mysql error is "Changing schema from '%s' to '%s' is not allowed." 
func IsServerErrorForbidSchemaChange(err error) bool {
    result := Isa(err, ER_FORBID_SCHEMA_CHANGE)
    return result
}

    
// IsServerErrorRowIsReferenced2 check mysql error is "Cannot delete or update a parent row: a foreign keyconstraint fails (%s)" 
func IsServerErrorRowIsReferenced2(err error) bool {
    result := Isa(err, ER_ROW_IS_REFERENCED_2)
    return result
}

    
// IsServerErrorNoReferencedRow2 check mysql error is "Cannot add or update a child row: a foreign keyconstraint fails (%s)" 
func IsServerErrorNoReferencedRow2(err error) bool {
    result := Isa(err, ER_NO_REFERENCED_ROW_2)
    return result
}

    
// IsServerErrorSpBadVarShadow check mysql error is "Variable '%s' must be quoted with `...`, or renamed" 
func IsServerErrorSpBadVarShadow(err error) bool {
    result := Isa(err, ER_SP_BAD_VAR_SHADOW)
    return result
}

    
// IsServerErrorTrgNoDefiner check mysql error is "No definer attribute for trigger '%s'.'%s'. The triggerwill be activated under the authorization of the invoker, whichmay have insufficient privileges. Please recreate the trigger." 
func IsServerErrorTrgNoDefiner(err error) bool {
    result := Isa(err, ER_TRG_NO_DEFINER)
    return result
}

    
// IsServerErrorOldFileFormat check mysql error is "'%s' has an old format, you should re-create the '%s'object(s)" 
func IsServerErrorOldFileFormat(err error) bool {
    result := Isa(err, ER_OLD_FILE_FORMAT)
    return result
}

    
// IsServerErrorSpRecursionLimit check mysql error is "Recursive limit %d (as set by the max_sp_recursion_depthvariable) was exceeded for routine %s" 
func IsServerErrorSpRecursionLimit(err error) bool {
    result := Isa(err, ER_SP_RECURSION_LIMIT)
    return result
}

    
// IsServerErrorSpProcTableCorrupt check mysql error is "Failed to load routine %s. The table mysql.proc ismissing, corrupt, or contains bad data (internal code %d)" 
func IsServerErrorSpProcTableCorrupt(err error) bool {
    result := Isa(err, ER_SP_PROC_TABLE_CORRUPT)
    return result
}

    
// IsServerErrorSpWrongName check mysql error is "Incorrect routine name '%s'" 
func IsServerErrorSpWrongName(err error) bool {
    result := Isa(err, ER_SP_WRONG_NAME)
    return result
}

    
// IsServerErrorTableNeedsUpgrade check mysql error is "Table upgrade required. Please do "REPAIR TABLE `%s`" ordump/reload to fix it!" 
func IsServerErrorTableNeedsUpgrade(err error) bool {
    result := Isa(err, ER_TABLE_NEEDS_UPGRADE)
    return result
}

    
// IsServerErrorSpNoAggregate check mysql error is "AGGREGATE is not supported for stored functions" 
func IsServerErrorSpNoAggregate(err error) bool {
    result := Isa(err, ER_SP_NO_AGGREGATE)
    return result
}

    
// IsServerErrorMaxPreparedStmtCountReached check mysql error is "Can't create more than max_prepared_stmt_count statements(current value: %lu)" 
func IsServerErrorMaxPreparedStmtCountReached(err error) bool {
    result := Isa(err, ER_MAX_PREPARED_STMT_COUNT_REACHED)
    return result
}

    
// IsServerErrorViewRecursive check mysql error is "`%s`.`%s` contains view recursion" 
func IsServerErrorViewRecursive(err error) bool {
    result := Isa(err, ER_VIEW_RECURSIVE)
    return result
}

    
// IsServerErrorNonGroupingFieldUsed check mysql error is "Non-grouping field '%s' is used in %s clause" 
func IsServerErrorNonGroupingFieldUsed(err error) bool {
    result := Isa(err, ER_NON_GROUPING_FIELD_USED)
    return result
}

    
// IsServerErrorTableCantHandleSpkeys check mysql error is "The used table type doesn't support SPATIAL indexes" 
func IsServerErrorTableCantHandleSpkeys(err error) bool {
    result := Isa(err, ER_TABLE_CANT_HANDLE_SPKEYS)
    return result
}

    
// IsServerErrorNoTriggersOnSystemSchema check mysql error is "Triggers can not be created on system tables" 
func IsServerErrorNoTriggersOnSystemSchema(err error) bool {
    result := Isa(err, ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA)
    return result
}

    
// IsServerErrorRemovedSpaces check mysql error is "Leading spaces are removed from name '%s'" 
func IsServerErrorRemovedSpaces(err error) bool {
    result := Isa(err, ER_REMOVED_SPACES)
    return result
}

    
// IsServerErrorAutoincReadFailed check mysql error is "Failed to read auto-increment value from storage engine" 
func IsServerErrorAutoincReadFailed(err error) bool {
    result := Isa(err, ER_AUTOINC_READ_FAILED)
    return result
}

    
// IsServerErrorUsername check mysql error is "user name" 
func IsServerErrorUsername(err error) bool {
    result := Isa(err, ER_USERNAME)
    return result
}

    
// IsServerErrorHostname check mysql error is "host name" 
func IsServerErrorHostname(err error) bool {
    result := Isa(err, ER_HOSTNAME)
    return result
}

    
// IsServerErrorWrongStringLength check mysql error is "String '%s' is too long for %s (should be no longer than%d)" 
func IsServerErrorWrongStringLength(err error) bool {
    result := Isa(err, ER_WRONG_STRING_LENGTH)
    return result
}

    
// IsServerErrorNonInsertableTable check mysql error is "The target table %s of the %s is not insertable-into" 
func IsServerErrorNonInsertableTable(err error) bool {
    result := Isa(err, ER_NON_INSERTABLE_TABLE)
    return result
}

    
// IsServerErrorAdminWrongMrgTable check mysql error is "Table '%s' is differently defined or of non-MyISAM typeor doesn't exist" 
func IsServerErrorAdminWrongMrgTable(err error) bool {
    result := Isa(err, ER_ADMIN_WRONG_MRG_TABLE)
    return result
}

    
// IsServerErrorTooHighLevelOfNestingForSelect check mysql error is "Too high level of nesting for select" 
func IsServerErrorTooHighLevelOfNestingForSelect(err error) bool {
    result := Isa(err, ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT)
    return result
}

    
// IsServerErrorNameBecomesEmpty check mysql error is "Name '%s' has become ''" 
func IsServerErrorNameBecomesEmpty(err error) bool {
    result := Isa(err, ER_NAME_BECOMES_EMPTY)
    return result
}

    
// IsServerErrorAmbiguousFieldTerm check mysql error is "First character of the FIELDS TERMINATED string isambiguous" 
func IsServerErrorAmbiguousFieldTerm(err error) bool {
    result := Isa(err, ER_AMBIGUOUS_FIELD_TERM)
    return result
}

    
// IsServerErrorForeignServerExists check mysql error is "The foreign server, %s, you are trying to create alreadyexists." 
func IsServerErrorForeignServerExists(err error) bool {
    result := Isa(err, ER_FOREIGN_SERVER_EXISTS)
    return result
}

    
// IsServerErrorForeignServerDoesntExist check mysql error is "The foreign server name you are trying to reference doesnot exist. Data source error: %s" 
func IsServerErrorForeignServerDoesntExist(err error) bool {
    result := Isa(err, ER_FOREIGN_SERVER_DOESNT_EXIST)
    return result
}

    
// IsServerErrorIllegalHaCreateOption check mysql error is "Table storage engine '%s' does not support the createoption '%s'" 
func IsServerErrorIllegalHaCreateOption(err error) bool {
    result := Isa(err, ER_ILLEGAL_HA_CREATE_OPTION)
    return result
}

    
// IsServerErrorPartitionRequiresValuesError check mysql error is "Syntax error: %s PARTITIONING requires definition ofVALUES %s for each partition" 
func IsServerErrorPartitionRequiresValuesError(err error) bool {
    result := Isa(err, ER_PARTITION_REQUIRES_VALUES_ERROR)
    return result
}

    
// IsServerErrorPartitionWrongValuesError check mysql error is "Only %s PARTITIONING can use VALUES %s in partitiondefinition" 
func IsServerErrorPartitionWrongValuesError(err error) bool {
    result := Isa(err, ER_PARTITION_WRONG_VALUES_ERROR)
    return result
}

    
// IsServerErrorPartitionMaxvalueError check mysql error is "MAXVALUE can only be used in last partition definition" 
func IsServerErrorPartitionMaxvalueError(err error) bool {
    result := Isa(err, ER_PARTITION_MAXVALUE_ERROR)
    return result
}

    
// IsServerErrorPartitionSubpartitionError check mysql error is "Subpartitions can only be hash partitions and by key" 
func IsServerErrorPartitionSubpartitionError(err error) bool {
    result := Isa(err, ER_PARTITION_SUBPARTITION_ERROR)
    return result
}

    
// IsServerErrorPartitionSubpartMixError check mysql error is "Must define subpartitions on all partitions if on onepartition" 
func IsServerErrorPartitionSubpartMixError(err error) bool {
    result := Isa(err, ER_PARTITION_SUBPART_MIX_ERROR)
    return result
}

    
// IsServerErrorPartitionWrongNoPartError check mysql error is "Wrong number of partitions defined, mismatch withprevious setting" 
func IsServerErrorPartitionWrongNoPartError(err error) bool {
    result := Isa(err, ER_PARTITION_WRONG_NO_PART_ERROR)
    return result
}

    
// IsServerErrorPartitionWrongNoSubpartError check mysql error is "Wrong number of subpartitions defined, mismatch withprevious setting" 
func IsServerErrorPartitionWrongNoSubpartError(err error) bool {
    result := Isa(err, ER_PARTITION_WRONG_NO_SUBPART_ERROR)
    return result
}

    
// IsServerErrorWrongExprInPartitionFuncError check mysql error is "Constant, random or timezone-dependent expressions in(sub)partitioning function are not allowed" 
func IsServerErrorWrongExprInPartitionFuncError(err error) bool {
    result := Isa(err, ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR)
    return result
}

    
// IsServerErrorNoConstExprInRangeOrListError check mysql error is "Expression in RANGE/LIST VALUES must be constant" 
func IsServerErrorNoConstExprInRangeOrListError(err error) bool {
    result := Isa(err, ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR)
    return result
}

    
// IsServerErrorFieldNotFoundPartError check mysql error is "Field in list of fields for partition function not foundin table" 
func IsServerErrorFieldNotFoundPartError(err error) bool {
    result := Isa(err, ER_FIELD_NOT_FOUND_PART_ERROR)
    return result
}

    
// IsServerErrorListOfFieldsOnlyInHashError check mysql error is "List of fields is only allowed in KEY partitions" 
func IsServerErrorListOfFieldsOnlyInHashError(err error) bool {
    result := Isa(err, ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR)
    return result
}

    
// IsServerErrorInconsistentPartitionInfoError check mysql error is "The partition info in the frm file is not consistent withwhat can be written into the frm file" 
func IsServerErrorInconsistentPartitionInfoError(err error) bool {
    result := Isa(err, ER_INCONSISTENT_PARTITION_INFO_ERROR)
    return result
}

    
// IsServerErrorPartitionFuncNotAllowedError check mysql error is "The %s function returns the wrong type" 
func IsServerErrorPartitionFuncNotAllowedError(err error) bool {
    result := Isa(err, ER_PARTITION_FUNC_NOT_ALLOWED_ERROR)
    return result
}

    
// IsServerErrorPartitionsMustBeDefinedError check mysql error is "For %s partitions each partition must be defined" 
func IsServerErrorPartitionsMustBeDefinedError(err error) bool {
    result := Isa(err, ER_PARTITIONS_MUST_BE_DEFINED_ERROR)
    return result
}

    
// IsServerErrorRangeNotIncreasingError check mysql error is "VALUES LESS THAN value must be strictly increasing foreach partition" 
func IsServerErrorRangeNotIncreasingError(err error) bool {
    result := Isa(err, ER_RANGE_NOT_INCREASING_ERROR)
    return result
}

    
// IsServerErrorInconsistentTypeOfFunctionsError check mysql error is "VALUES value must be of same type as partition function" 
func IsServerErrorInconsistentTypeOfFunctionsError(err error) bool {
    result := Isa(err, ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR)
    return result
}

    
// IsServerErrorMultipleDefConstInListPartError check mysql error is "Multiple definition of same constant in list partitioning" 
func IsServerErrorMultipleDefConstInListPartError(err error) bool {
    result := Isa(err, ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR)
    return result
}

    
// IsServerErrorPartitionEntryError check mysql error is "Partitioning can not be used stand-alone in query" 
func IsServerErrorPartitionEntryError(err error) bool {
    result := Isa(err, ER_PARTITION_ENTRY_ERROR)
    return result
}

    
// IsServerErrorMixHandlerError check mysql error is "The mix of handlers in the partitions is not allowed inthis version of MySQL" 
func IsServerErrorMixHandlerError(err error) bool {
    result := Isa(err, ER_MIX_HANDLER_ERROR)
    return result
}

    
// IsServerErrorPartitionNotDefinedError check mysql error is "For the partitioned engine it is necessary to define all%s" 
func IsServerErrorPartitionNotDefinedError(err error) bool {
    result := Isa(err, ER_PARTITION_NOT_DEFINED_ERROR)
    return result
}

    
// IsServerErrorTooManyPartitionsError check mysql error is "Too many partitions (including subpartitions) weredefined" 
func IsServerErrorTooManyPartitionsError(err error) bool {
    result := Isa(err, ER_TOO_MANY_PARTITIONS_ERROR)
    return result
}

    
// IsServerErrorSubpartitionError check mysql error is "It is only possible to mix RANGE/LIST partitioning withHASH/KEY partitioning for subpartitioning" 
func IsServerErrorSubpartitionError(err error) bool {
    result := Isa(err, ER_SUBPARTITION_ERROR)
    return result
}

    
// IsServerErrorCantCreateHandlerFile check mysql error is "Failed to create specific handler file" 
func IsServerErrorCantCreateHandlerFile(err error) bool {
    result := Isa(err, ER_CANT_CREATE_HANDLER_FILE)
    return result
}

    
// IsServerErrorBlobFieldInPartFuncError check mysql error is "A BLOB field is not allowed in partition function" 
func IsServerErrorBlobFieldInPartFuncError(err error) bool {
    result := Isa(err, ER_BLOB_FIELD_IN_PART_FUNC_ERROR)
    return result
}

    
// IsServerErrorUniqueKeyNeedAllFieldsInPf check mysql error is "A %s must include all columns in the table's partitioningfunction" 
func IsServerErrorUniqueKeyNeedAllFieldsInPf(err error) bool {
    result := Isa(err, ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF)
    return result
}

    
// IsServerErrorNoPartsError check mysql error is "Number of %s = 0 is not an allowed value" 
func IsServerErrorNoPartsError(err error) bool {
    result := Isa(err, ER_NO_PARTS_ERROR)
    return result
}

    
// IsServerErrorPartitionMgmtOnNonpartitioned check mysql error is "Partition management on a not partitioned table is notpossible" 
func IsServerErrorPartitionMgmtOnNonpartitioned(err error) bool {
    result := Isa(err, ER_PARTITION_MGMT_ON_NONPARTITIONED)
    return result
}

    
// IsServerErrorForeignKeyOnPartitioned check mysql error is "Foreign key clause is not yet supported in conjunctionwith partitioning" 
func IsServerErrorForeignKeyOnPartitioned(err error) bool {
    result := Isa(err, ER_FOREIGN_KEY_ON_PARTITIONED)
    return result
}

    
// IsServerErrorDropPartitionNonExistent check mysql error is "Error in list of partitions to %s" 
func IsServerErrorDropPartitionNonExistent(err error) bool {
    result := Isa(err, ER_DROP_PARTITION_NON_EXISTENT)
    return result
}

    
// IsServerErrorDropLastPartition check mysql error is "Cannot remove all partitions, use DROP TABLE instead" 
func IsServerErrorDropLastPartition(err error) bool {
    result := Isa(err, ER_DROP_LAST_PARTITION)
    return result
}

    
// IsServerErrorCoalesceOnlyOnHashPartition check mysql error is "COALESCE PARTITION can only be used on HASH/KEYpartitions" 
func IsServerErrorCoalesceOnlyOnHashPartition(err error) bool {
    result := Isa(err, ER_COALESCE_ONLY_ON_HASH_PARTITION)
    return result
}

    
// IsServerErrorReorgHashOnlyOnSameNo check mysql error is "REORGANIZE PARTITION can only be used to reorganizepartitions not to change their numbers" 
func IsServerErrorReorgHashOnlyOnSameNo(err error) bool {
    result := Isa(err, ER_REORG_HASH_ONLY_ON_SAME_NO)
    return result
}

    
// IsServerErrorReorgNoParamError check mysql error is "REORGANIZE PARTITION without parameters can only be usedon auto-partitioned tables using HASH PARTITIONs" 
func IsServerErrorReorgNoParamError(err error) bool {
    result := Isa(err, ER_REORG_NO_PARAM_ERROR)
    return result
}

    
// IsServerErrorOnlyOnRangeListPartition check mysql error is "%s PARTITION can only be used on RANGE/LIST partitions" 
func IsServerErrorOnlyOnRangeListPartition(err error) bool {
    result := Isa(err, ER_ONLY_ON_RANGE_LIST_PARTITION)
    return result
}

    
// IsServerErrorAddPartitionSubpartError check mysql error is "Trying to Add partition(s) with wrong number ofsubpartitions" 
func IsServerErrorAddPartitionSubpartError(err error) bool {
    result := Isa(err, ER_ADD_PARTITION_SUBPART_ERROR)
    return result
}

    
// IsServerErrorAddPartitionNoNewPartition check mysql error is "At least one partition must be added" 
func IsServerErrorAddPartitionNoNewPartition(err error) bool {
    result := Isa(err, ER_ADD_PARTITION_NO_NEW_PARTITION)
    return result
}

    
// IsServerErrorCoalescePartitionNoPartition check mysql error is "At least one partition must be coalesced" 
func IsServerErrorCoalescePartitionNoPartition(err error) bool {
    result := Isa(err, ER_COALESCE_PARTITION_NO_PARTITION)
    return result
}

    
// IsServerErrorReorgPartitionNotExist check mysql error is "More partitions to reorganize than there are partitions" 
func IsServerErrorReorgPartitionNotExist(err error) bool {
    result := Isa(err, ER_REORG_PARTITION_NOT_EXIST)
    return result
}

    
// IsServerErrorSameNamePartition check mysql error is "Duplicate partition name %s" 
func IsServerErrorSameNamePartition(err error) bool {
    result := Isa(err, ER_SAME_NAME_PARTITION)
    return result
}

    
// IsServerErrorNoBinlogError check mysql error is "It is not allowed to shut off binlog on this command" 
func IsServerErrorNoBinlogError(err error) bool {
    result := Isa(err, ER_NO_BINLOG_ERROR)
    return result
}

    
// IsServerErrorConsecutiveReorgPartitions check mysql error is "When reorganizing a set of partitions they must be inconsecutive order" 
func IsServerErrorConsecutiveReorgPartitions(err error) bool {
    result := Isa(err, ER_CONSECUTIVE_REORG_PARTITIONS)
    return result
}

    
// IsServerErrorReorgOutsideRange check mysql error is "Reorganize of range partitions cannot change total rangesexcept for last partition where it can extend the range" 
func IsServerErrorReorgOutsideRange(err error) bool {
    result := Isa(err, ER_REORG_OUTSIDE_RANGE)
    return result
}

    
// IsServerErrorPartitionFunctionFailure check mysql error is "Partition function not supported in this version for thishandler" 
func IsServerErrorPartitionFunctionFailure(err error) bool {
    result := Isa(err, ER_PARTITION_FUNCTION_FAILURE)
    return result
}

    
// IsServerErrorPartStateError check mysql error is "Partition state cannot be defined from CREATE/ALTER TABLE" 
func IsServerErrorPartStateError(err error) bool {
    result := Isa(err, ER_PART_STATE_ERROR)
    return result
}

    
// IsServerErrorLimitedPartRange check mysql error is "The %s handler only supports 32 bit integers in VALUES" 
func IsServerErrorLimitedPartRange(err error) bool {
    result := Isa(err, ER_LIMITED_PART_RANGE)
    return result
}

    
// IsServerErrorPluginIsNotLoaded check mysql error is "Plugin '%s' is not loaded" 
func IsServerErrorPluginIsNotLoaded(err error) bool {
    result := Isa(err, ER_PLUGIN_IS_NOT_LOADED)
    return result
}

    
// IsServerErrorWrongValue check mysql error is "Incorrect %s value: '%s'" 
func IsServerErrorWrongValue(err error) bool {
    result := Isa(err, ER_WRONG_VALUE)
    return result
}

    
// IsServerErrorNoPartitionForGivenValue check mysql error is "Table has no partition for value %s" 
func IsServerErrorNoPartitionForGivenValue(err error) bool {
    result := Isa(err, ER_NO_PARTITION_FOR_GIVEN_VALUE)
    return result
}

    
// IsServerErrorFilegroupOptionOnlyOnce check mysql error is "It is not allowed to specify %s more than once" 
func IsServerErrorFilegroupOptionOnlyOnce(err error) bool {
    result := Isa(err, ER_FILEGROUP_OPTION_ONLY_ONCE)
    return result
}

    
// IsServerErrorCreateFilegroupFailed check mysql error is "Failed to create %s" 
func IsServerErrorCreateFilegroupFailed(err error) bool {
    result := Isa(err, ER_CREATE_FILEGROUP_FAILED)
    return result
}

    
// IsServerErrorDropFilegroupFailed check mysql error is "Failed to drop %s" 
func IsServerErrorDropFilegroupFailed(err error) bool {
    result := Isa(err, ER_DROP_FILEGROUP_FAILED)
    return result
}

    
// IsServerErrorTablespaceAutoExtendError check mysql error is "The handler doesn't support autoextend of tablespaces" 
func IsServerErrorTablespaceAutoExtendError(err error) bool {
    result := Isa(err, ER_TABLESPACE_AUTO_EXTEND_ERROR)
    return result
}

    
// IsServerErrorWrongSizeNumber check mysql error is "A size parameter was incorrectly specified, either numberor on the form 10M" 
func IsServerErrorWrongSizeNumber(err error) bool {
    result := Isa(err, ER_WRONG_SIZE_NUMBER)
    return result
}

    
// IsServerErrorSizeOverflowError check mysql error is "The size number was correct but we don't allow the digitpart to be more than 2 billion" 
func IsServerErrorSizeOverflowError(err error) bool {
    result := Isa(err, ER_SIZE_OVERFLOW_ERROR)
    return result
}

    
// IsServerErrorAlterFilegroupFailed check mysql error is "Failed to alter: %s" 
func IsServerErrorAlterFilegroupFailed(err error) bool {
    result := Isa(err, ER_ALTER_FILEGROUP_FAILED)
    return result
}

    
// IsServerErrorBinlogRowLoggingFailed check mysql error is "Writing one row to the row-based binary log failed" 
func IsServerErrorBinlogRowLoggingFailed(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_LOGGING_FAILED)
    return result
}

    
// IsServerErrorBinlogRowWrongTableDef check mysql error is "Table definition on master and slave does not match: %s" 
func IsServerErrorBinlogRowWrongTableDef(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_WRONG_TABLE_DEF)
    return result
}

    
// IsServerErrorBinlogRowRbrToSbr check mysql error is "Slave running with --log-slave-updates must use row-basedbinary logging to be able to replicate row-based binary log events" 
func IsServerErrorBinlogRowRbrToSbr(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_RBR_TO_SBR)
    return result
}

    
// IsServerErrorEventAlreadyExists check mysql error is "Event '%s' already exists" 
func IsServerErrorEventAlreadyExists(err error) bool {
    result := Isa(err, ER_EVENT_ALREADY_EXISTS)
    return result
}

    
// IsServerErrorEventStoreFailed check mysql error is "Failed to store event %s. Error code %d from storageengine." 
func IsServerErrorEventStoreFailed(err error) bool {
    result := Isa(err, ER_EVENT_STORE_FAILED)
    return result
}

    
// IsServerErrorEventDoesNotExist check mysql error is "Unknown event '%s'" 
func IsServerErrorEventDoesNotExist(err error) bool {
    result := Isa(err, ER_EVENT_DOES_NOT_EXIST)
    return result
}

    
// IsServerErrorEventCantAlter check mysql error is "Failed to alter event '%s'" 
func IsServerErrorEventCantAlter(err error) bool {
    result := Isa(err, ER_EVENT_CANT_ALTER)
    return result
}

    
// IsServerErrorEventDropFailed check mysql error is "Failed to drop %s" 
func IsServerErrorEventDropFailed(err error) bool {
    result := Isa(err, ER_EVENT_DROP_FAILED)
    return result
}

    
// IsServerErrorEventIntervalNotPositiveOrTooBig check mysql error is "INTERVAL is either not positive or too big" 
func IsServerErrorEventIntervalNotPositiveOrTooBig(err error) bool {
    result := Isa(err, ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG)
    return result
}

    
// IsServerErrorEventEndsBeforeStarts check mysql error is "ENDS is either invalid or before STARTS" 
func IsServerErrorEventEndsBeforeStarts(err error) bool {
    result := Isa(err, ER_EVENT_ENDS_BEFORE_STARTS)
    return result
}

    
// IsServerErrorEventExecTimeInThePast check mysql error is "Event execution time is in the past. Event has beendisabled" 
func IsServerErrorEventExecTimeInThePast(err error) bool {
    result := Isa(err, ER_EVENT_EXEC_TIME_IN_THE_PAST)
    return result
}

    
// IsServerErrorEventOpenTableFailed check mysql error is "Failed to open mysql.event" 
func IsServerErrorEventOpenTableFailed(err error) bool {
    result := Isa(err, ER_EVENT_OPEN_TABLE_FAILED)
    return result
}

    
// IsServerErrorEventNeitherMExprNorMAt check mysql error is "No datetime expression provided" 
func IsServerErrorEventNeitherMExprNorMAt(err error) bool {
    result := Isa(err, ER_EVENT_NEITHER_M_EXPR_NOR_M_AT)
    return result
}

    
// IsServerErrorObsoleteColCountDoesntMatchCorrupted check mysql error is "Column count of mysql.%s is wrong. Expected %d, found %d.The table is probably corrupted" 
func IsServerErrorObsoleteColCountDoesntMatchCorrupted(err error) bool {
    result := Isa(err, ER_OBSOLETE_COL_COUNT_DOESNT_MATCH_CORRUPTED)
    return result
}

    
// IsServerErrorObsoleteCannotLoadFromTable check mysql error is "Cannot load from mysql.%s. The table is probablycorrupted" 
func IsServerErrorObsoleteCannotLoadFromTable(err error) bool {
    result := Isa(err, ER_OBSOLETE_CANNOT_LOAD_FROM_TABLE)
    return result
}

    
// IsServerErrorEventCannotDelete check mysql error is "Failed to delete the event from mysql.event" 
func IsServerErrorEventCannotDelete(err error) bool {
    result := Isa(err, ER_EVENT_CANNOT_DELETE)
    return result
}

    
// IsServerErrorEventCompileError check mysql error is "Error during compilation of event's body" 
func IsServerErrorEventCompileError(err error) bool {
    result := Isa(err, ER_EVENT_COMPILE_ERROR)
    return result
}

    
// IsServerErrorEventSameName check mysql error is "Same old and new event name" 
func IsServerErrorEventSameName(err error) bool {
    result := Isa(err, ER_EVENT_SAME_NAME)
    return result
}

    
// IsServerErrorEventDataTooLong check mysql error is "Data for column '%s' too long" 
func IsServerErrorEventDataTooLong(err error) bool {
    result := Isa(err, ER_EVENT_DATA_TOO_LONG)
    return result
}

    
// IsServerErrorDropIndexFk check mysql error is "Cannot drop index '%s': needed in a foreign keyconstraint" 
func IsServerErrorDropIndexFk(err error) bool {
    result := Isa(err, ER_DROP_INDEX_FK)
    return result
}

    
// IsServerErrorWarnDeprecatedSyntaxWithVer check mysql error is "The syntax '%s' is deprecated and will be removed inMySQL %s. Please use %s instead" 
func IsServerErrorWarnDeprecatedSyntaxWithVer(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYNTAX_WITH_VER)
    return result
}

    
// IsServerErrorCantWriteLockLogTable check mysql error is "You can't write-lock a log table. Only read access ispossible" 
func IsServerErrorCantWriteLockLogTable(err error) bool {
    result := Isa(err, ER_CANT_WRITE_LOCK_LOG_TABLE)
    return result
}

    
// IsServerErrorCantLockLogTable check mysql error is "You can't use locks with log tables." 
func IsServerErrorCantLockLogTable(err error) bool {
    result := Isa(err, ER_CANT_LOCK_LOG_TABLE)
    return result
}

    
// IsServerErrorForeignDuplicateKeyOldUnused check mysql error is "Upholding foreign key constraints for table '%s', entry'%s', key %d would lead to a duplicate entry" 
func IsServerErrorForeignDuplicateKeyOldUnused(err error) bool {
    result := Isa(err, ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED)
    return result
}

    
// IsServerErrorColCountDoesntMatchPleaseUpdate check mysql error is "Column count of mysql.%s is wrong. Expected %d, found %d.Created with MySQL %d, now running %d. Please use mysql_upgrade tofix this error." 
func IsServerErrorColCountDoesntMatchPleaseUpdate(err error) bool {
    result := Isa(err, ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE)
    return result
}

    
// IsServerErrorTempTablePreventsSwitchOutOfRbr check mysql error is "Cannot switch out of the row-based binary log format whenthe session has open temporary tables" 
func IsServerErrorTempTablePreventsSwitchOutOfRbr(err error) bool {
    result := Isa(err, ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR)
    return result
}

    
// IsServerErrorStoredFunctionPreventsSwitchBinlogFormat check mysql error is "Cannot change the binary logging format inside a storedfunction or trigger" 
func IsServerErrorStoredFunctionPreventsSwitchBinlogFormat(err error) bool {
    result := Isa(err, ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT)
    return result
}

    
// IsServerErrorNdbCantSwitchBinlogFormat check mysql error is "The NDB cluster engine does not support changing thebinlog format on the fly yet" 
func IsServerErrorNdbCantSwitchBinlogFormat(err error) bool {
    result := Isa(err, ER_NDB_CANT_SWITCH_BINLOG_FORMAT)
    return result
}

    
// IsServerErrorPartitionNoTemporary check mysql error is "Cannot create temporary table with partitions" 
func IsServerErrorPartitionNoTemporary(err error) bool {
    result := Isa(err, ER_PARTITION_NO_TEMPORARY)
    return result
}

    
// IsServerErrorPartitionConstDomainError check mysql error is "Partition constant is out of partition function domain" 
func IsServerErrorPartitionConstDomainError(err error) bool {
    result := Isa(err, ER_PARTITION_CONST_DOMAIN_ERROR)
    return result
}

    
// IsServerErrorPartitionFunctionIsNotAllowed check mysql error is "This partition function is not allowed" 
func IsServerErrorPartitionFunctionIsNotAllowed(err error) bool {
    result := Isa(err, ER_PARTITION_FUNCTION_IS_NOT_ALLOWED)
    return result
}

    
// IsServerErrorDdlLogError check mysql error is "Error in DDL log" 
func IsServerErrorDdlLogError(err error) bool {
    result := Isa(err, ER_DDL_LOG_ERROR)
    return result
}

    
// IsServerErrorNullInValuesLessThan check mysql error is "Not allowed to use NULL value in VALUES LESS THAN" 
func IsServerErrorNullInValuesLessThan(err error) bool {
    result := Isa(err, ER_NULL_IN_VALUES_LESS_THAN)
    return result
}

    
// IsServerErrorWrongPartitionName check mysql error is "Incorrect partition name" 
func IsServerErrorWrongPartitionName(err error) bool {
    result := Isa(err, ER_WRONG_PARTITION_NAME)
    return result
}

    
// IsServerErrorCantChangeTxCharacteristics check mysql error is "Transaction characteristics can't be changed while atransaction is in progress" 
func IsServerErrorCantChangeTxCharacteristics(err error) bool {
    result := Isa(err, ER_CANT_CHANGE_TX_CHARACTERISTICS)
    return result
}

    
// IsServerErrorDupEntryAutoincrementCase check mysql error is "ALTER TABLE causes auto_increment resequencing, resultingin duplicate entry '%s' for key '%s'" 
func IsServerErrorDupEntryAutoincrementCase(err error) bool {
    result := Isa(err, ER_DUP_ENTRY_AUTOINCREMENT_CASE)
    return result
}

    
// IsServerErrorEventModifyQueueError check mysql error is "Internal scheduler error %d" 
func IsServerErrorEventModifyQueueError(err error) bool {
    result := Isa(err, ER_EVENT_MODIFY_QUEUE_ERROR)
    return result
}

    
// IsServerErrorEventSetVarError check mysql error is "Error during starting/stopping of the scheduler. Errorcode %u" 
func IsServerErrorEventSetVarError(err error) bool {
    result := Isa(err, ER_EVENT_SET_VAR_ERROR)
    return result
}

    
// IsServerErrorPartitionMergeError check mysql error is "Engine cannot be used in partitioned tables" 
func IsServerErrorPartitionMergeError(err error) bool {
    result := Isa(err, ER_PARTITION_MERGE_ERROR)
    return result
}

    
// IsServerErrorCantActivateLog check mysql error is "Cannot activate '%s' log" 
func IsServerErrorCantActivateLog(err error) bool {
    result := Isa(err, ER_CANT_ACTIVATE_LOG)
    return result
}

    
// IsServerErrorRbrNotAvailable check mysql error is "The server was not built with row-based replication" 
func IsServerErrorRbrNotAvailable(err error) bool {
    result := Isa(err, ER_RBR_NOT_AVAILABLE)
    return result
}

    
// IsServerErrorBase64DecodeError check mysql error is "Decoding of base64 string failed" 
func IsServerErrorBase64DecodeError(err error) bool {
    result := Isa(err, ER_BASE64_DECODE_ERROR)
    return result
}

    
// IsServerErrorEventRecursionForbidden check mysql error is "Recursion of EVENT DDL statements is forbidden when bodyis present" 
func IsServerErrorEventRecursionForbidden(err error) bool {
    result := Isa(err, ER_EVENT_RECURSION_FORBIDDEN)
    return result
}

    
// IsServerErrorEventsDbError check mysql error is "Cannot proceed because system tables used by EventScheduler were found damaged at server start" 
func IsServerErrorEventsDbError(err error) bool {
    result := Isa(err, ER_EVENTS_DB_ERROR)
    return result
}

    
// IsServerErrorOnlyIntegersAllowed check mysql error is "Only integers allowed as number here" 
func IsServerErrorOnlyIntegersAllowed(err error) bool {
    result := Isa(err, ER_ONLY_INTEGERS_ALLOWED)
    return result
}

    
// IsServerErrorUnsuportedLogEngine check mysql error is "This storage engine cannot be used for log tables"" 
func IsServerErrorUnsuportedLogEngine(err error) bool {
    result := Isa(err, ER_UNSUPORTED_LOG_ENGINE)
    return result
}

    
// IsServerErrorBadLogStatement check mysql error is "You cannot '%s' a log table if logging is enabled" 
func IsServerErrorBadLogStatement(err error) bool {
    result := Isa(err, ER_BAD_LOG_STATEMENT)
    return result
}

    
// IsServerErrorCantRenameLogTable check mysql error is "Cannot rename '%s'. When logging enabled, rename to/fromlog table must rename two tables: the log table to an archivetable and another table back to '%s'" 
func IsServerErrorCantRenameLogTable(err error) bool {
    result := Isa(err, ER_CANT_RENAME_LOG_TABLE)
    return result
}

    
// IsServerErrorWrongParamcountToNativeFct check mysql error is "Incorrect parameter count in the call to native function'%s'" 
func IsServerErrorWrongParamcountToNativeFct(err error) bool {
    result := Isa(err, ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT)
    return result
}

    
// IsServerErrorWrongParametersToNativeFct check mysql error is "Incorrect parameters in the call to native function '%s'" 
func IsServerErrorWrongParametersToNativeFct(err error) bool {
    result := Isa(err, ER_WRONG_PARAMETERS_TO_NATIVE_FCT)
    return result
}

    
// IsServerErrorWrongParametersToStoredFct check mysql error is "Incorrect parameters in the call to stored function '%s'" 
func IsServerErrorWrongParametersToStoredFct(err error) bool {
    result := Isa(err, ER_WRONG_PARAMETERS_TO_STORED_FCT)
    return result
}

    
// IsServerErrorNativeFctNameCollision check mysql error is "This function '%s' has the same name as a native function" 
func IsServerErrorNativeFctNameCollision(err error) bool {
    result := Isa(err, ER_NATIVE_FCT_NAME_COLLISION)
    return result
}

    
// IsServerErrorDupEntryWithKeyName check mysql error is "Duplicate entry '%s' for key '%s'" 
func IsServerErrorDupEntryWithKeyName(err error) bool {
    result := Isa(err, ER_DUP_ENTRY_WITH_KEY_NAME)
    return result
}

    
// IsServerErrorBinlogPurgeEmfile check mysql error is "Too many files opened, please execute the command again" 
func IsServerErrorBinlogPurgeEmfile(err error) bool {
    result := Isa(err, ER_BINLOG_PURGE_EMFILE)
    return result
}

    
// IsServerErrorEventCannotCreateInThePast check mysql error is "Event execution time is in the past and ON COMPLETION NOTPRESERVE is set. The event was dropped immediately after creation." 
func IsServerErrorEventCannotCreateInThePast(err error) bool {
    result := Isa(err, ER_EVENT_CANNOT_CREATE_IN_THE_PAST)
    return result
}

    
// IsServerErrorEventCannotAlterInThePast check mysql error is "Event execution time is in the past and ON COMPLETION NOTPRESERVE is set. The event was not changed. Specify a time in thefuture." 
func IsServerErrorEventCannotAlterInThePast(err error) bool {
    result := Isa(err, ER_EVENT_CANNOT_ALTER_IN_THE_PAST)
    return result
}

    
// IsServerErrorSlaveIncident check mysql error is "The incident %s occured on the master.Message: %s" 
func IsServerErrorSlaveIncident(err error) bool {
    result := Isa(err, ER_SLAVE_INCIDENT)
    return result
}

    
// IsServerErrorNoPartitionForGivenValueSilent check mysql error is "Table has no partition for some existing values" 
func IsServerErrorNoPartitionForGivenValueSilent(err error) bool {
    result := Isa(err, ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT)
    return result
}

    
// IsServerErrorBinlogUnsafeStatement check mysql error is "Unsafe statement written to the binary log usingstatement format since BINLOG_FORMAT = STATEMENT. %s" 
func IsServerErrorBinlogUnsafeStatement(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_STATEMENT)
    return result
}

    
// IsServerErrorSlaveFatalError check mysql error is "Fatal error: %s" 
func IsServerErrorSlaveFatalError(err error) bool {
    result := Isa(err, ER_SLAVE_FATAL_ERROR)
    return result
}

    
// IsServerErrorSlaveRelayLogReadFailure check mysql error is "Relay log read failure: %s" 
func IsServerErrorSlaveRelayLogReadFailure(err error) bool {
    result := Isa(err, ER_SLAVE_RELAY_LOG_READ_FAILURE)
    return result
}

    
// IsServerErrorSlaveRelayLogWriteFailure check mysql error is "Relay log write failure: %s" 
func IsServerErrorSlaveRelayLogWriteFailure(err error) bool {
    result := Isa(err, ER_SLAVE_RELAY_LOG_WRITE_FAILURE)
    return result
}

    
// IsServerErrorSlaveCreateEventFailure check mysql error is "Failed to create %s" 
func IsServerErrorSlaveCreateEventFailure(err error) bool {
    result := Isa(err, ER_SLAVE_CREATE_EVENT_FAILURE)
    return result
}

    
// IsServerErrorSlaveMasterComFailure check mysql error is "Master command %s failed: %s" 
func IsServerErrorSlaveMasterComFailure(err error) bool {
    result := Isa(err, ER_SLAVE_MASTER_COM_FAILURE)
    return result
}

    
// IsServerErrorBinlogLoggingImpossible check mysql error is "Binary logging not possible.Message: %s" 
func IsServerErrorBinlogLoggingImpossible(err error) bool {
    result := Isa(err, ER_BINLOG_LOGGING_IMPOSSIBLE)
    return result
}

    
// IsServerErrorViewNoCreationCtx check mysql error is "View `%s`.`%s` has no creation context" 
func IsServerErrorViewNoCreationCtx(err error) bool {
    result := Isa(err, ER_VIEW_NO_CREATION_CTX)
    return result
}

    
// IsServerErrorViewInvalidCreationCtx check mysql error is "Creation context of view `%s`.`%s' is invalid" 
func IsServerErrorViewInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_VIEW_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorSrInvalidCreationCtx check mysql error is "Creation context of stored routine `%s`.`%s` is invalid" 
func IsServerErrorSrInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_SR_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorTrgCorruptedFile check mysql error is "Corrupted TRG file for table `%s`.`%s`" 
func IsServerErrorTrgCorruptedFile(err error) bool {
    result := Isa(err, ER_TRG_CORRUPTED_FILE)
    return result
}

    
// IsServerErrorTrgNoCreationCtx check mysql error is "Triggers for table `%s`.`%s` have no creation context" 
func IsServerErrorTrgNoCreationCtx(err error) bool {
    result := Isa(err, ER_TRG_NO_CREATION_CTX)
    return result
}

    
// IsServerErrorTrgInvalidCreationCtx check mysql error is "Trigger creation context of table `%s`.`%s` is invalid" 
func IsServerErrorTrgInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_TRG_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorEventInvalidCreationCtx check mysql error is "Creation context of event `%s`.`%s` is invalid" 
func IsServerErrorEventInvalidCreationCtx(err error) bool {
    result := Isa(err, ER_EVENT_INVALID_CREATION_CTX)
    return result
}

    
// IsServerErrorTrgCantOpenTable check mysql error is "Cannot open table for trigger `%s`.`%s`" 
func IsServerErrorTrgCantOpenTable(err error) bool {
    result := Isa(err, ER_TRG_CANT_OPEN_TABLE)
    return result
}

    
// IsServerErrorCantCreateSroutine check mysql error is "Cannot create stored routine `%s`. Check warnings" 
func IsServerErrorCantCreateSroutine(err error) bool {
    result := Isa(err, ER_CANT_CREATE_SROUTINE)
    return result
}

    
// IsServerErrorNeverUsed check mysql error is "Ambiguous slave modes combination. %s" 
func IsServerErrorNeverUsed(err error) bool {
    result := Isa(err, ER_NEVER_USED)
    return result
}

    
// IsServerErrorNoFormatDescriptionEventBeforeBinlogStatement check mysql error is "The BINLOG statement of type `%s` was not preceded by aformat description BINLOG statement." 
func IsServerErrorNoFormatDescriptionEventBeforeBinlogStatement(err error) bool {
    result := Isa(err, ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT)
    return result
}

    
// IsServerErrorSlaveCorruptEvent check mysql error is "Corrupted replication event was detected" 
func IsServerErrorSlaveCorruptEvent(err error) bool {
    result := Isa(err, ER_SLAVE_CORRUPT_EVENT)
    return result
}

    
// IsServerErrorLoadDataInvalidColumn check mysql error is "Invalid column reference (%s) in LOAD DATA" 
func IsServerErrorLoadDataInvalidColumn(err error) bool {
    result := Isa(err, ER_LOAD_DATA_INVALID_COLUMN)
    return result
}

    
// IsServerErrorLogPurgeNoFile check mysql error is "Being purged log %s was not found" 
func IsServerErrorLogPurgeNoFile(err error) bool {
    result := Isa(err, ER_LOG_PURGE_NO_FILE)
    return result
}

    
// IsServerErrorXaRbtimeout check mysql error is "XA_RBTIMEOUT: Transaction branch was rolled back: tooktoo long" 
func IsServerErrorXaRbtimeout(err error) bool {
    result := Isa(err, ER_XA_RBTIMEOUT)
    return result
}

    
// IsServerErrorXaRbdeadlock check mysql error is "XA_RBDEADLOCK: Transaction branch was rolled back:deadlock was detected" 
func IsServerErrorXaRbdeadlock(err error) bool {
    result := Isa(err, ER_XA_RBDEADLOCK)
    return result
}

    
// IsServerErrorNeedReprepare check mysql error is "Prepared statement needs to be re-prepared" 
func IsServerErrorNeedReprepare(err error) bool {
    result := Isa(err, ER_NEED_REPREPARE)
    return result
}

    
// IsServerErrorDelayedNotSupported check mysql error is "DELAYED option not supported for table '%s'" 
func IsServerErrorDelayedNotSupported(err error) bool {
    result := Isa(err, ER_DELAYED_NOT_SUPPORTED)
    return result
}

    
// IsWarnNoMasterInfo check mysql error is "The master info structure does not exist" 
func IsWarnNoMasterInfo(err error) bool {
    result := Isa(err, WARN_NO_MASTER_INFO)
    return result
}

    
// IsWarnOptionIgnored check mysql error is "<%s> option ignored" 
func IsWarnOptionIgnored(err error) bool {
    result := Isa(err, WARN_OPTION_IGNORED)
    return result
}

    
// IsWarnPluginDeleteBuiltin check mysql error is "Built-in plugins cannot be deleted" 
func IsWarnPluginDeleteBuiltin(err error) bool {
    result := Isa(err, WARN_PLUGIN_DELETE_BUILTIN)
    return result
}

    
// IsWarnPluginBusy check mysql error is "Plugin is busy and will be uninstalled on shutdown" 
func IsWarnPluginBusy(err error) bool {
    result := Isa(err, WARN_PLUGIN_BUSY)
    return result
}

    
// IsServerErrorVariableIsReadonly check mysql error is "%s variable '%s' is read-only. Use SET %s to assign thevalue" 
func IsServerErrorVariableIsReadonly(err error) bool {
    result := Isa(err, ER_VARIABLE_IS_READONLY)
    return result
}

    
// IsServerErrorWarnEngineTransactionRollback check mysql error is "Storage engine %s does not support rollback for thisstatement. Transaction rolled back and must be restarted" 
func IsServerErrorWarnEngineTransactionRollback(err error) bool {
    result := Isa(err, ER_WARN_ENGINE_TRANSACTION_ROLLBACK)
    return result
}

    
// IsServerErrorSlaveHeartbeatFailure check mysql error is "Unexpected master's heartbeat data: %s" 
func IsServerErrorSlaveHeartbeatFailure(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_FAILURE)
    return result
}

    
// IsServerErrorSlaveHeartbeatValueOutOfRange check mysql error is "The requested value for the heartbeat period is eithernegative or exceeds the maximum allowed (%s seconds)." 
func IsServerErrorSlaveHeartbeatValueOutOfRange(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorNdbReplicationSchemaError check mysql error is "Bad schema for mysql.ndb_replication table.Message: %s" 
func IsServerErrorNdbReplicationSchemaError(err error) bool {
    result := Isa(err, ER_NDB_REPLICATION_SCHEMA_ERROR)
    return result
}

    
// IsServerErrorConflictFnParseError check mysql error is "Error in parsing conflict function.Message: %s" 
func IsServerErrorConflictFnParseError(err error) bool {
    result := Isa(err, ER_CONFLICT_FN_PARSE_ERROR)
    return result
}

    
// IsServerErrorExceptionsWriteError check mysql error is "Write to exceptions table failed.Message: %s"" 
func IsServerErrorExceptionsWriteError(err error) bool {
    result := Isa(err, ER_EXCEPTIONS_WRITE_ERROR)
    return result
}

    
// IsServerErrorTooLongTableComment check mysql error is "Comment for table '%s' is too long (max = %lu)" 
func IsServerErrorTooLongTableComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_TABLE_COMMENT)
    return result
}

    
// IsServerErrorTooLongFieldComment check mysql error is "Comment for field '%s' is too long (max = %lu)" 
func IsServerErrorTooLongFieldComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_FIELD_COMMENT)
    return result
}

    
// IsServerErrorFuncInexistentNameCollision check mysql error is "FUNCTION %s does not exist. Check the 'Function NameParsing and Resolution' section in the Reference Manual" 
func IsServerErrorFuncInexistentNameCollision(err error) bool {
    result := Isa(err, ER_FUNC_INEXISTENT_NAME_COLLISION)
    return result
}

    
// IsServerErrorDatabaseName check mysql error is "Database" 
func IsServerErrorDatabaseName(err error) bool {
    result := Isa(err, ER_DATABASE_NAME)
    return result
}

    
// IsServerErrorTableName check mysql error is "Table" 
func IsServerErrorTableName(err error) bool {
    result := Isa(err, ER_TABLE_NAME)
    return result
}

    
// IsServerErrorPartitionName check mysql error is "Partition" 
func IsServerErrorPartitionName(err error) bool {
    result := Isa(err, ER_PARTITION_NAME)
    return result
}

    
// IsServerErrorSubpartitionName check mysql error is "Subpartition" 
func IsServerErrorSubpartitionName(err error) bool {
    result := Isa(err, ER_SUBPARTITION_NAME)
    return result
}

    
// IsServerErrorTemporaryName check mysql error is "Temporary" 
func IsServerErrorTemporaryName(err error) bool {
    result := Isa(err, ER_TEMPORARY_NAME)
    return result
}

    
// IsServerErrorRenamedName check mysql error is "Renamed" 
func IsServerErrorRenamedName(err error) bool {
    result := Isa(err, ER_RENAMED_NAME)
    return result
}

    
// IsServerErrorTooManyConcurrentTrxs check mysql error is "Too many active concurrent transactions" 
func IsServerErrorTooManyConcurrentTrxs(err error) bool {
    result := Isa(err, ER_TOO_MANY_CONCURRENT_TRXS)
    return result
}

    
// IsWarnNonAsciiSeparatorNotImplemented check mysql error is "Non-ASCII separator arguments are not fully supported" 
func IsWarnNonAsciiSeparatorNotImplemented(err error) bool {
    result := Isa(err, WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED)
    return result
}

    
// IsServerErrorDebugSyncTimeout check mysql error is "debug sync point wait timed out" 
func IsServerErrorDebugSyncTimeout(err error) bool {
    result := Isa(err, ER_DEBUG_SYNC_TIMEOUT)
    return result
}

    
// IsServerErrorDebugSyncHitLimit check mysql error is "debug sync point hit limit reached" 
func IsServerErrorDebugSyncHitLimit(err error) bool {
    result := Isa(err, ER_DEBUG_SYNC_HIT_LIMIT)
    return result
}

    
// IsServerErrorDupSignalSet check mysql error is "Duplicate condition information item '%s'" 
func IsServerErrorDupSignalSet(err error) bool {
    result := Isa(err, ER_DUP_SIGNAL_SET)
    return result
}

    
// IsServerErrorSignalWarn check mysql error is "Unhandled user-defined warning condition" 
func IsServerErrorSignalWarn(err error) bool {
    result := Isa(err, ER_SIGNAL_WARN)
    return result
}

    
// IsServerErrorSignalNotFound check mysql error is "Unhandled user-defined not found condition" 
func IsServerErrorSignalNotFound(err error) bool {
    result := Isa(err, ER_SIGNAL_NOT_FOUND)
    return result
}

    
// IsServerErrorSignalException check mysql error is "Unhandled user-defined exception condition" 
func IsServerErrorSignalException(err error) bool {
    result := Isa(err, ER_SIGNAL_EXCEPTION)
    return result
}

    
// IsServerErrorResignalWithoutActiveHandler check mysql error is "RESIGNAL when handler not active" 
func IsServerErrorResignalWithoutActiveHandler(err error) bool {
    result := Isa(err, ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER)
    return result
}

    
// IsServerErrorSignalBadConditionType check mysql error is "SIGNAL/RESIGNAL can only use a CONDITION defined withSQLSTATE" 
func IsServerErrorSignalBadConditionType(err error) bool {
    result := Isa(err, ER_SIGNAL_BAD_CONDITION_TYPE)
    return result
}

    
// IsWarnCondItemTruncated check mysql error is "Data truncated for condition item '%s'" 
func IsWarnCondItemTruncated(err error) bool {
    result := Isa(err, WARN_COND_ITEM_TRUNCATED)
    return result
}

    
// IsServerErrorCondItemTooLong check mysql error is "Data too long for condition item '%s'" 
func IsServerErrorCondItemTooLong(err error) bool {
    result := Isa(err, ER_COND_ITEM_TOO_LONG)
    return result
}

    
// IsServerErrorUnknownLocale check mysql error is "Unknown locale: '%s'" 
func IsServerErrorUnknownLocale(err error) bool {
    result := Isa(err, ER_UNKNOWN_LOCALE)
    return result
}

    
// IsServerErrorSlaveIgnoreServerIds check mysql error is "The requested server id %d clashes with the slave startupoption --replicate-same-server-id" 
func IsServerErrorSlaveIgnoreServerIds(err error) bool {
    result := Isa(err, ER_SLAVE_IGNORE_SERVER_IDS)
    return result
}

    
// IsServerErrorQueryCacheDisabled check mysql error is "Query cache is disabled" 
func IsServerErrorQueryCacheDisabled(err error) bool {
    result := Isa(err, ER_QUERY_CACHE_DISABLED)
    return result
}

    
// IsServerErrorSameNamePartitionField check mysql error is "Duplicate partition field name '%s'" 
func IsServerErrorSameNamePartitionField(err error) bool {
    result := Isa(err, ER_SAME_NAME_PARTITION_FIELD)
    return result
}

    
// IsServerErrorPartitionColumnListError check mysql error is "Inconsistency in usage of column lists for partitioning" 
func IsServerErrorPartitionColumnListError(err error) bool {
    result := Isa(err, ER_PARTITION_COLUMN_LIST_ERROR)
    return result
}

    
// IsServerErrorWrongTypeColumnValueError check mysql error is "Partition column values of incorrect type" 
func IsServerErrorWrongTypeColumnValueError(err error) bool {
    result := Isa(err, ER_WRONG_TYPE_COLUMN_VALUE_ERROR)
    return result
}

    
// IsServerErrorTooManyPartitionFuncFieldsError check mysql error is "Too many fields in '%s'" 
func IsServerErrorTooManyPartitionFuncFieldsError(err error) bool {
    result := Isa(err, ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR)
    return result
}

    
// IsServerErrorMaxvalueInValuesIn check mysql error is "Cannot use MAXVALUE as value in VALUES IN" 
func IsServerErrorMaxvalueInValuesIn(err error) bool {
    result := Isa(err, ER_MAXVALUE_IN_VALUES_IN)
    return result
}

    
// IsServerErrorTooManyValuesError check mysql error is "Cannot have more than one value for this type of %spartitioning" 
func IsServerErrorTooManyValuesError(err error) bool {
    result := Isa(err, ER_TOO_MANY_VALUES_ERROR)
    return result
}

    
// IsServerErrorRowSinglePartitionFieldError check mysql error is "Row expressions in VALUES IN only allowed for multi-fieldcolumn partitioning" 
func IsServerErrorRowSinglePartitionFieldError(err error) bool {
    result := Isa(err, ER_ROW_SINGLE_PARTITION_FIELD_ERROR)
    return result
}

    
// IsServerErrorFieldTypeNotAllowedAsPartitionField check mysql error is "Field '%s' is of a not allowed type for this type ofpartitioning" 
func IsServerErrorFieldTypeNotAllowedAsPartitionField(err error) bool {
    result := Isa(err, ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD)
    return result
}

    
// IsServerErrorPartitionFieldsTooLong check mysql error is "The total length of the partitioning fields is too large" 
func IsServerErrorPartitionFieldsTooLong(err error) bool {
    result := Isa(err, ER_PARTITION_FIELDS_TOO_LONG)
    return result
}

    
// IsServerErrorBinlogRowEngineAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binarylog since both row-incapable engines and statement-incapableengines are involved." 
func IsServerErrorBinlogRowEngineAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogRowModeAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binarylog since BINLOG_FORMAT = ROW and at least one table uses astorage engine limited to statement-based logging." 
func IsServerErrorBinlogRowModeAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_MODE_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binarylog since statement is unsafe, storage engine is limited tostatement-based logging, and BINLOG_FORMAT = MIXED. %s" 
func IsServerErrorBinlogUnsafeAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogRowInjectionAndStmtEngine check mysql error is "Cannot execute statement: impossible to write to binarylog since statement is in row format and at least one table uses astorage engine limited to statement-based logging." 
func IsServerErrorBinlogRowInjectionAndStmtEngine(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE)
    return result
}

    
// IsServerErrorBinlogStmtModeAndRowEngine check mysql error is "Cannot execute statement: impossible to write to binarylog since BINLOG_FORMAT = STATEMENT and at least one table uses astorage engine limited to row-based logging.%s" 
func IsServerErrorBinlogStmtModeAndRowEngine(err error) bool {
    result := Isa(err, ER_BINLOG_STMT_MODE_AND_ROW_ENGINE)
    return result
}

    
// IsServerErrorBinlogRowInjectionAndStmtMode check mysql error is "Cannot execute statement: impossible to write to binarylog since statement is in row format and BINLOG_FORMAT =STATEMENT." 
func IsServerErrorBinlogRowInjectionAndStmtMode(err error) bool {
    result := Isa(err, ER_BINLOG_ROW_INJECTION_AND_STMT_MODE)
    return result
}

    
// IsServerErrorBinlogMultipleEnginesAndSelfLoggingEngine check mysql error is "Cannot execute statement: impossible to write to binarylog since more than one engine is involved and at least one engineis self-logging." 
func IsServerErrorBinlogMultipleEnginesAndSelfLoggingEngine(err error) bool {
    result := Isa(err, ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeLimit check mysql error is "The statement is unsafe because it uses a LIMIT clause.This is unsafe because the set of rows included cannot bepredicted." 
func IsServerErrorBinlogUnsafeLimit(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_LIMIT)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertDelayed check mysql error is "The statement is unsafe because it uses INSERT DELAYED.This is unsafe because the times when rows are inserted cannot bepredicted." 
func IsServerErrorBinlogUnsafeInsertDelayed(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_DELAYED)
    return result
}

    
// IsServerErrorBinlogUnsafeSystemTable check mysql error is "The statement is unsafe because it uses the general log,slow query log, or performance_schema table(s). This is unsafebecause system tables may differ on slaves." 
func IsServerErrorBinlogUnsafeSystemTable(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_SYSTEM_TABLE)
    return result
}

    
// IsServerErrorBinlogUnsafeAutoincColumns check mysql error is "Statement is unsafe because it invokes a trigger or astored function that inserts into an AUTO_INCREMENT column.Inserted values cannot be logged correctly." 
func IsServerErrorBinlogUnsafeAutoincColumns(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_AUTOINC_COLUMNS)
    return result
}

    
// IsServerErrorBinlogUnsafeUdf check mysql error is "Statement is unsafe because it uses a UDF which may notreturn the same value on the slave." 
func IsServerErrorBinlogUnsafeUdf(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_UDF)
    return result
}

    
// IsServerErrorBinlogUnsafeSystemVariable check mysql error is "Statement is unsafe because it uses a system variablethat may have a different value on the slave." 
func IsServerErrorBinlogUnsafeSystemVariable(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_SYSTEM_VARIABLE)
    return result
}

    
// IsServerErrorBinlogUnsafeSystemFunction check mysql error is "Statement is unsafe because it uses a system functionthat may return a different value on the slave." 
func IsServerErrorBinlogUnsafeSystemFunction(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_SYSTEM_FUNCTION)
    return result
}

    
// IsServerErrorBinlogUnsafeNontransAfterTrans check mysql error is "Statement is unsafe because it accesses anon-transactional table after accessing a transactional tablewithin the same transaction." 
func IsServerErrorBinlogUnsafeNontransAfterTrans(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS)
    return result
}

    
// IsServerErrorMessageAndStatement check mysql error is "%s Statement: %s" 
func IsServerErrorMessageAndStatement(err error) bool {
    result := Isa(err, ER_MESSAGE_AND_STATEMENT)
    return result
}

    
// IsServerErrorSlaveConversionFailed check mysql error is "Column %d of table '%s.%s' cannot be converted from type'%s' to type '%s'" 
func IsServerErrorSlaveConversionFailed(err error) bool {
    result := Isa(err, ER_SLAVE_CONVERSION_FAILED)
    return result
}

    
// IsServerErrorSlaveCantCreateConversion check mysql error is "Can't create conversion table for table '%s.%s'" 
func IsServerErrorSlaveCantCreateConversion(err error) bool {
    result := Isa(err, ER_SLAVE_CANT_CREATE_CONVERSION)
    return result
}

    
// IsServerErrorInsideTransactionPreventsSwitchBinlogFormat check mysql error is "Cannot modify @@session.binlog_format inside atransaction" 
func IsServerErrorInsideTransactionPreventsSwitchBinlogFormat(err error) bool {
    result := Isa(err, ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT)
    return result
}

    
// IsServerErrorPathLength check mysql error is "The path specified for %s is too long." 
func IsServerErrorPathLength(err error) bool {
    result := Isa(err, ER_PATH_LENGTH)
    return result
}

    
// IsServerErrorWarnDeprecatedSyntaxNoReplacement check mysql error is "'%s' is deprecated and will be removed in a futurerelease." 
func IsServerErrorWarnDeprecatedSyntaxNoReplacement(err error) bool {
    result := Isa(err, ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT)
    return result
}

    
// IsServerErrorWrongNativeTableStructure check mysql error is "Native table '%s'.'%s' has the wrong structure" 
func IsServerErrorWrongNativeTableStructure(err error) bool {
    result := Isa(err, ER_WRONG_NATIVE_TABLE_STRUCTURE)
    return result
}

    
// IsServerErrorWrongPerfschemaUsage check mysql error is "Invalid performance_schema usage." 
func IsServerErrorWrongPerfschemaUsage(err error) bool {
    result := Isa(err, ER_WRONG_PERFSCHEMA_USAGE)
    return result
}

    
// IsServerErrorWarnISSkippedTable check mysql error is "Table '%s'.'%s' was skipped since its definition is beingmodified by concurrent DDL statement" 
func IsServerErrorWarnISSkippedTable(err error) bool {
    result := Isa(err, ER_WARN_I_S_SKIPPED_TABLE)
    return result
}

    
// IsServerErrorInsideTransactionPreventsSwitchBinlogDirect check mysql error is "Cannot modify@@session.binlog_direct_non_transactional_updates inside atransaction" 
func IsServerErrorInsideTransactionPreventsSwitchBinlogDirect(err error) bool {
    result := Isa(err, ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT)
    return result
}

    
// IsServerErrorStoredFunctionPreventsSwitchBinlogDirect check mysql error is "Cannot change the binlog direct flag inside a storedfunction or trigger" 
func IsServerErrorStoredFunctionPreventsSwitchBinlogDirect(err error) bool {
    result := Isa(err, ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT)
    return result
}

    
// IsServerErrorSpatialMustHaveGeomCol check mysql error is "A SPATIAL index may only contain a geometrical typecolumn" 
func IsServerErrorSpatialMustHaveGeomCol(err error) bool {
    result := Isa(err, ER_SPATIAL_MUST_HAVE_GEOM_COL)
    return result
}

    
// IsServerErrorTooLongIndexComment check mysql error is "Comment for index '%s' is too long (max = %lu)" 
func IsServerErrorTooLongIndexComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_INDEX_COMMENT)
    return result
}

    
// IsServerErrorLockAborted check mysql error is "Wait on a lock was aborted due to a pending exclusivelock" 
func IsServerErrorLockAborted(err error) bool {
    result := Isa(err, ER_LOCK_ABORTED)
    return result
}

    
// IsServerErrorDataOutOfRange check mysql error is "%s value is out of range in '%s'" 
func IsServerErrorDataOutOfRange(err error) bool {
    result := Isa(err, ER_DATA_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorWrongSpvarTypeInLimit check mysql error is "A variable of a non-integer based type in LIMIT clause" 
func IsServerErrorWrongSpvarTypeInLimit(err error) bool {
    result := Isa(err, ER_WRONG_SPVAR_TYPE_IN_LIMIT)
    return result
}

    
// IsServerErrorBinlogUnsafeMultipleEnginesAndSelfLoggingEngine check mysql error is "Mixing self-logging and non-self-logging engines in astatement is unsafe." 
func IsServerErrorBinlogUnsafeMultipleEnginesAndSelfLoggingEngine(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeMixedStatement check mysql error is "Statement accesses nontransactional table as well astransactional or temporary table, and writes to any of them." 
func IsServerErrorBinlogUnsafeMixedStatement(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_MIXED_STATEMENT)
    return result
}

    
// IsServerErrorInsideTransactionPreventsSwitchSqlLogBin check mysql error is "Cannot modify @@session.sql_log_bin inside a transaction" 
func IsServerErrorInsideTransactionPreventsSwitchSqlLogBin(err error) bool {
    result := Isa(err, ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN)
    return result
}

    
// IsServerErrorStoredFunctionPreventsSwitchSqlLogBin check mysql error is "Cannot change the sql_log_bin inside a stored function ortrigger" 
func IsServerErrorStoredFunctionPreventsSwitchSqlLogBin(err error) bool {
    result := Isa(err, ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN)
    return result
}

    
// IsServerErrorFailedReadFromParFile check mysql error is "Failed to read from the .par file" 
func IsServerErrorFailedReadFromParFile(err error) bool {
    result := Isa(err, ER_FAILED_READ_FROM_PAR_FILE)
    return result
}

    
// IsServerErrorValuesIsNotIntTypeError check mysql error is "VALUES value for partition '%s' must have type INT" 
func IsServerErrorValuesIsNotIntTypeError(err error) bool {
    result := Isa(err, ER_VALUES_IS_NOT_INT_TYPE_ERROR)
    return result
}

    
// IsServerErrorAccessDeniedNoPasswordError check mysql error is "Access denied for user '%s'@'%s'" 
func IsServerErrorAccessDeniedNoPasswordError(err error) bool {
    result := Isa(err, ER_ACCESS_DENIED_NO_PASSWORD_ERROR)
    return result
}

    
// IsServerErrorSetPasswordAuthPlugin check mysql error is "SET PASSWORD has no significance for users authenticatingvia plugins" 
func IsServerErrorSetPasswordAuthPlugin(err error) bool {
    result := Isa(err, ER_SET_PASSWORD_AUTH_PLUGIN)
    return result
}

    
// IsServerErrorGrantPluginUserExists check mysql error is "GRANT with IDENTIFIED WITH is illegal because the user%-.*s already exists" 
func IsServerErrorGrantPluginUserExists(err error) bool {
    result := Isa(err, ER_GRANT_PLUGIN_USER_EXISTS)
    return result
}

    
// IsServerErrorTruncateIllegalFk check mysql error is "Cannot truncate a table referenced in a foreign keyconstraint (%s)" 
func IsServerErrorTruncateIllegalFk(err error) bool {
    result := Isa(err, ER_TRUNCATE_ILLEGAL_FK)
    return result
}

    
// IsServerErrorPluginIsPermanent check mysql error is "Plugin '%s' is force_plus_permanent and can not beunloaded" 
func IsServerErrorPluginIsPermanent(err error) bool {
    result := Isa(err, ER_PLUGIN_IS_PERMANENT)
    return result
}

    
// IsServerErrorSlaveHeartbeatValueOutOfRangeMin check mysql error is "The requested value for the heartbeat period is less than1 millisecond. The value is reset to 0, meaning that heartbeatingwill effectively be disabled." 
func IsServerErrorSlaveHeartbeatValueOutOfRangeMin(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN)
    return result
}

    
// IsServerErrorSlaveHeartbeatValueOutOfRangeMax check mysql error is "The requested value for the heartbeat period exceeds thevalue of `slave_net_timeout' seconds. A sensible value for theperiod should be less than the timeout." 
func IsServerErrorSlaveHeartbeatValueOutOfRangeMax(err error) bool {
    result := Isa(err, ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX)
    return result
}

    
// IsServerErrorStmtCacheFull check mysql error is "Multi-row statements required more than'max_binlog_stmt_cache_size' bytes of storage" 
func IsServerErrorStmtCacheFull(err error) bool {
    result := Isa(err, ER_STMT_CACHE_FULL)
    return result
}

    
// IsServerErrorMultiUpdateKeyConflict check mysql error is "Primary key/partition key update is not allowed since thetable is updated both as '%s' and '%s'." 
func IsServerErrorMultiUpdateKeyConflict(err error) bool {
    result := Isa(err, ER_MULTI_UPDATE_KEY_CONFLICT)
    return result
}

    
// IsServerErrorTableNeedsRebuild check mysql error is "Table rebuild required. Please do "ALTER TABLE `%s`FORCE" or dump/reload to fix it!" 
func IsServerErrorTableNeedsRebuild(err error) bool {
    result := Isa(err, ER_TABLE_NEEDS_REBUILD)
    return result
}

    
// IsWarnOptionBelowLimit check mysql error is "The value of '%s' should be no less than the value of'%s'" 
func IsWarnOptionBelowLimit(err error) bool {
    result := Isa(err, WARN_OPTION_BELOW_LIMIT)
    return result
}

    
// IsServerErrorIndexColumnTooLong check mysql error is "Index column size too large. The maximum column size is%lu bytes." 
func IsServerErrorIndexColumnTooLong(err error) bool {
    result := Isa(err, ER_INDEX_COLUMN_TOO_LONG)
    return result
}

    
// IsServerErrorErrorInTriggerBody check mysql error is "Trigger '%s' has an error in its body: '%s'" 
func IsServerErrorErrorInTriggerBody(err error) bool {
    result := Isa(err, ER_ERROR_IN_TRIGGER_BODY)
    return result
}

    
// IsServerErrorErrorInUnknownTriggerBody check mysql error is "Unknown trigger has an error in its body: '%s'" 
func IsServerErrorErrorInUnknownTriggerBody(err error) bool {
    result := Isa(err, ER_ERROR_IN_UNKNOWN_TRIGGER_BODY)
    return result
}

    
// IsServerErrorIndexCorrupt check mysql error is "Index %s is corrupted" 
func IsServerErrorIndexCorrupt(err error) bool {
    result := Isa(err, ER_INDEX_CORRUPT)
    return result
}

    
// IsServerErrorUndoRecordTooBig check mysql error is "Undo log record is too big." 
func IsServerErrorUndoRecordTooBig(err error) bool {
    result := Isa(err, ER_UNDO_RECORD_TOO_BIG)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertIgnoreSelect check mysql error is "INSERT IGNORE... SELECT is unsafe because the order inwhich rows are retrieved by the SELECT determines which (if any)rows are ignored. This order cannot be predicted and may differ onmaster and the slave." 
func IsServerErrorBinlogUnsafeInsertIgnoreSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertSelectUpdate check mysql error is "INSERT... SELECT... ON DUPLICATE KEY UPDATE is unsafebecause the order in which rows are retrieved by the SELECTdetermines which (if any) rows are updated. This order cannot bepredicted and may differ on master and the slave." 
func IsServerErrorBinlogUnsafeInsertSelectUpdate(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE)
    return result
}

    
// IsServerErrorBinlogUnsafeReplaceSelect check mysql error is "REPLACE... SELECT is unsafe because the order in whichrows are retrieved by the SELECT determines which (if any) rowsare replaced. This order cannot be predicted and may differ onmaster and the slave." 
func IsServerErrorBinlogUnsafeReplaceSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_REPLACE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeCreateIgnoreSelect check mysql error is "CREATE... IGNORE SELECT is unsafe because the order inwhich rows are retrieved by the SELECT determines which (if any)rows are ignored. This order cannot be predicted and may differ onmaster and the slave." 
func IsServerErrorBinlogUnsafeCreateIgnoreSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeCreateReplaceSelect check mysql error is "CREATE... REPLACE SELECT is unsafe because the order inwhich rows are retrieved by the SELECT determines which (if any)rows are replaced. This order cannot be predicted and may differon master and the slave." 
func IsServerErrorBinlogUnsafeCreateReplaceSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeUpdateIgnore check mysql error is "UPDATE IGNORE is unsafe because the order in which rowsare updated determines which (if any) rows are ignored. This ordercannot be predicted and may differ on master and the slave." 
func IsServerErrorBinlogUnsafeUpdateIgnore(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_UPDATE_IGNORE)
    return result
}

    
// IsServerErrorPluginNoUninstall check mysql error is "Plugin '%s' is marked as not dynamically uninstallable.You have to stop the server to uninstall it." 
func IsServerErrorPluginNoUninstall(err error) bool {
    result := Isa(err, ER_PLUGIN_NO_UNINSTALL)
    return result
}

    
// IsServerErrorPluginNoInstall check mysql error is "Plugin '%s' is marked as not dynamically installable. Youhave to stop the server to install it." 
func IsServerErrorPluginNoInstall(err error) bool {
    result := Isa(err, ER_PLUGIN_NO_INSTALL)
    return result
}

    
// IsServerErrorBinlogUnsafeWriteAutoincSelect check mysql error is "Statements writing to a table with an auto-incrementcolumn after selecting from another table are unsafe because theorder in which rows are retrieved determines what (if any) rowswill be written. This order cannot be predicted and may differ onmaster and the slave." 
func IsServerErrorBinlogUnsafeWriteAutoincSelect(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT)
    return result
}

    
// IsServerErrorBinlogUnsafeCreateSelectAutoinc check mysql error is "CREATE TABLE... SELECT... on a table with anauto-increment column is unsafe because the order in which rowsare retrieved by the SELECT determines which (if any) rows areinserted. This order cannot be predicted and may differ on masterand the slave." 
func IsServerErrorBinlogUnsafeCreateSelectAutoinc(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC)
    return result
}

    
// IsServerErrorBinlogUnsafeInsertTwoKeys check mysql error is "INSERT... ON DUPLICATE KEY UPDATE on a table with morethan one UNIQUE KEY is unsafe" 
func IsServerErrorBinlogUnsafeInsertTwoKeys(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_INSERT_TWO_KEYS)
    return result
}

    
// IsServerErrorTableInFkCheck check mysql error is "Table is being used in foreign key check." 
func IsServerErrorTableInFkCheck(err error) bool {
    result := Isa(err, ER_TABLE_IN_FK_CHECK)
    return result
}

    
// IsServerErrorUnsupportedEngine check mysql error is "Storage engine '%s' does not support system tables.[%s.%s]" 
func IsServerErrorUnsupportedEngine(err error) bool {
    result := Isa(err, ER_UNSUPPORTED_ENGINE)
    return result
}

    
// IsServerErrorBinlogUnsafeAutoincNotFirst check mysql error is "INSERT into autoincrement field which is not the firstpart in the composed primary key is unsafe." 
func IsServerErrorBinlogUnsafeAutoincNotFirst(err error) bool {
    result := Isa(err, ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST)
    return result
}

    
// IsServerErrorCannotLoadFromTableV2 check mysql error is "Cannot load from %s.%s. The table is probably corrupted" 
func IsServerErrorCannotLoadFromTableV2(err error) bool {
    result := Isa(err, ER_CANNOT_LOAD_FROM_TABLE_V2)
    return result
}

    
// IsServerErrorMasterDelayValueOutOfRange check mysql error is "The requested value %s for the master delay exceeds themaximum %u" 
func IsServerErrorMasterDelayValueOutOfRange(err error) bool {
    result := Isa(err, ER_MASTER_DELAY_VALUE_OUT_OF_RANGE)
    return result
}

    
// IsServerErrorOnlyFdAndRbrEventsAllowedInBinlogStatement check mysql error is "Only Format_description_log_event and row events areallowed in BINLOG statements (but %s was provided)" 
func IsServerErrorOnlyFdAndRbrEventsAllowedInBinlogStatement(err error) bool {
    result := Isa(err, ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT)
    return result
}

    
// IsServerErrorPartitionExchangeDifferentOption check mysql error is "Non matching attribute '%s' between partition and table" 
func IsServerErrorPartitionExchangeDifferentOption(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_DIFFERENT_OPTION)
    return result
}

    
// IsServerErrorPartitionExchangePartTable check mysql error is "Table to exchange with partition is partitioned: '%s'" 
func IsServerErrorPartitionExchangePartTable(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_PART_TABLE)
    return result
}

    
// IsServerErrorPartitionExchangeTempTable check mysql error is "Table to exchange with partition is temporary: '%s'" 
func IsServerErrorPartitionExchangeTempTable(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_TEMP_TABLE)
    return result
}

    
// IsServerErrorPartitionInsteadOfSubpartition check mysql error is "Subpartitioned table, use subpartition instead ofpartition" 
func IsServerErrorPartitionInsteadOfSubpartition(err error) bool {
    result := Isa(err, ER_PARTITION_INSTEAD_OF_SUBPARTITION)
    return result
}

    
// IsServerErrorUnknownPartition check mysql error is "Unknown partition '%s' in table '%s'" 
func IsServerErrorUnknownPartition(err error) bool {
    result := Isa(err, ER_UNKNOWN_PARTITION)
    return result
}

    
// IsServerErrorTablesDifferentMetadata check mysql error is "Tables have different definitions" 
func IsServerErrorTablesDifferentMetadata(err error) bool {
    result := Isa(err, ER_TABLES_DIFFERENT_METADATA)
    return result
}

    
// IsServerErrorRowDoesNotMatchPartition check mysql error is "Found a row that does not match the partition" 
func IsServerErrorRowDoesNotMatchPartition(err error) bool {
    result := Isa(err, ER_ROW_DOES_NOT_MATCH_PARTITION)
    return result
}

    
// IsServerErrorBinlogCacheSizeGreaterThanMax check mysql error is "Option binlog_cache_size (%lu) is greater thanmax_binlog_cache_size (%lu)" 
func IsServerErrorBinlogCacheSizeGreaterThanMax(err error) bool {
    result := Isa(err, ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX)
    return result
}

    
// IsServerErrorWarnIndexNotApplicable check mysql error is "Cannot use %s access on index '%s' due to type orcollation conversion on field '%s'" 
func IsServerErrorWarnIndexNotApplicable(err error) bool {
    result := Isa(err, ER_WARN_INDEX_NOT_APPLICABLE)
    return result
}

    
// IsServerErrorPartitionExchangeForeignKey check mysql error is "Table to exchange with partition has foreign keyreferences: '%s'" 
func IsServerErrorPartitionExchangeForeignKey(err error) bool {
    result := Isa(err, ER_PARTITION_EXCHANGE_FOREIGN_KEY)
    return result
}

    
// IsServerErrorNoSuchKeyValue check mysql error is "Key value '%s' was not found in table '%s.%s'" 
func IsServerErrorNoSuchKeyValue(err error) bool {
    result := Isa(err, ER_NO_SUCH_KEY_VALUE)
    return result
}

    
// IsServerErrorRplInfoDataTooLong check mysql error is "Data for column '%s' too long" 
func IsServerErrorRplInfoDataTooLong(err error) bool {
    result := Isa(err, ER_RPL_INFO_DATA_TOO_LONG)
    return result
}

    
// IsServerErrorNetworkReadEventChecksumFailure check mysql error is "Replication event checksum verification failed whilereading from network." 
func IsServerErrorNetworkReadEventChecksumFailure(err error) bool {
    result := Isa(err, ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE)
    return result
}

    
// IsServerErrorBinlogReadEventChecksumFailure check mysql error is "Replication event checksum verification failed whilereading from a log file." 
func IsServerErrorBinlogReadEventChecksumFailure(err error) bool {
    result := Isa(err, ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE)
    return result
}

    
// IsServerErrorBinlogStmtCacheSizeGreaterThanMax check mysql error is "Option binlog_stmt_cache_size (%lu) is greater thanmax_binlog_stmt_cache_size (%lu)" 
func IsServerErrorBinlogStmtCacheSizeGreaterThanMax(err error) bool {
    result := Isa(err, ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX)
    return result
}

    
// IsServerErrorCantUpdateTableInCreateTableSelect check mysql error is "Can't update table '%s' while '%s' is being created." 
func IsServerErrorCantUpdateTableInCreateTableSelect(err error) bool {
    result := Isa(err, ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT)
    return result
}

    
// IsServerErrorPartitionClauseOnNonpartitioned check mysql error is "PARTITION () clause on non partitioned table" 
func IsServerErrorPartitionClauseOnNonpartitioned(err error) bool {
    result := Isa(err, ER_PARTITION_CLAUSE_ON_NONPARTITIONED)
    return result
}

    
// IsServerErrorRowDoesNotMatchGivenPartitionSet check mysql error is "Found a row not matching the given partition set" 
func IsServerErrorRowDoesNotMatchGivenPartitionSet(err error) bool {
    result := Isa(err, ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET)
    return result
}

    
// IsServerErrorNoSuchPartitionUnused check mysql error is "partition '%s' doesn't exist" 
func IsServerErrorNoSuchPartitionUnused(err error) bool {
    result := Isa(err, ER_NO_SUCH_PARTITION__UNUSED)
    return result
}

    
// IsServerErrorChangeRplInfoRepositoryFailure check mysql error is "Failure while changing the type of replicationrepository: %s." 
func IsServerErrorChangeRplInfoRepositoryFailure(err error) bool {
    result := Isa(err, ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE)
    return result
}

    
// IsServerErrorWarningNotCompleteRollbackWithCreatedTempTable check mysql error is "The creation of some temporary tables could not be rolledback." 
func IsServerErrorWarningNotCompleteRollbackWithCreatedTempTable(err error) bool {
    result := Isa(err, ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE)
    return result
}

    
// IsServerErrorWarningNotCompleteRollbackWithDroppedTempTable check mysql error is "Some temporary tables were dropped, but these operationscould not be rolled back." 
func IsServerErrorWarningNotCompleteRollbackWithDroppedTempTable(err error) bool {
    result := Isa(err, ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE)
    return result
}

    
// IsServerErrorMtsFeatureIsNotSupported check mysql error is "%s is not supported in multi-threaded slave mode. %s" 
func IsServerErrorMtsFeatureIsNotSupported(err error) bool {
    result := Isa(err, ER_MTS_FEATURE_IS_NOT_SUPPORTED)
    return result
}

    
// IsServerErrorMtsUpdatedDbsGreaterMax check mysql error is "The number of modified databases exceeds the maximum %d" 
func IsServerErrorMtsUpdatedDbsGreaterMax(err error) bool {
    result := Isa(err, ER_MTS_UPDATED_DBS_GREATER_MAX)
    return result
}

    
// IsServerErrorMtsCantParallel check mysql error is "Cannot execute the current event group in the parallelmode. Encountered event %s, relay-log name %s, position %s whichprevents execution of this event group in parallel mode. Reason:%s." 
func IsServerErrorMtsCantParallel(err error) bool {
    result := Isa(err, ER_MTS_CANT_PARALLEL)
    return result
}

    
// IsServerErrorMtsInconsistentData check mysql error is "%s" 
func IsServerErrorMtsInconsistentData(err error) bool {
    result := Isa(err, ER_MTS_INCONSISTENT_DATA)
    return result
}

    
// IsServerErrorFulltextNotSupportedWithPartitioning check mysql error is "FULLTEXT index is not supported for partitioned tables." 
func IsServerErrorFulltextNotSupportedWithPartitioning(err error) bool {
    result := Isa(err, ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING)
    return result
}

    
// IsServerErrorDaInvalidConditionNumber check mysql error is "Invalid condition number" 
func IsServerErrorDaInvalidConditionNumber(err error) bool {
    result := Isa(err, ER_DA_INVALID_CONDITION_NUMBER)
    return result
}

    
// IsServerErrorInsecurePlainText check mysql error is "Sending passwords in plain text without SSL/TLS isextremely insecure." 
func IsServerErrorInsecurePlainText(err error) bool {
    result := Isa(err, ER_INSECURE_PLAIN_TEXT)
    return result
}

    
// IsServerErrorInsecureChangeMaster check mysql error is "Storing MySQL user name or password information in themaster info repository is not secure and is therefore notrecommended. Please consider using the USER and PASSWORDconnection options for START SLAVE" 
func IsServerErrorInsecureChangeMaster(err error) bool {
    result := Isa(err, ER_INSECURE_CHANGE_MASTER)
    return result
}

    
// IsServerErrorForeignDuplicateKeyWithChildInfo check mysql error is "Foreign key constraint for table '%s', record '%s' wouldlead to a duplicate entry in table '%s', key '%s'" 
func IsServerErrorForeignDuplicateKeyWithChildInfo(err error) bool {
    result := Isa(err, ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO)
    return result
}

    
// IsServerErrorForeignDuplicateKeyWithoutChildInfo check mysql error is "Foreign key constraint for table '%s', record '%s' wouldlead to a duplicate entry in a child table" 
func IsServerErrorForeignDuplicateKeyWithoutChildInfo(err error) bool {
    result := Isa(err, ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO)
    return result
}

    
// IsServerErrorSqlthreadWithSecureSlave check mysql error is "Setting authentication options is not possible when onlythe Slave SQL Thread is being started." 
func IsServerErrorSqlthreadWithSecureSlave(err error) bool {
    result := Isa(err, ER_SQLTHREAD_WITH_SECURE_SLAVE)
    return result
}

    
// IsServerErrorTableHasNoFt check mysql error is "The table does not have FULLTEXT index to support thisquery" 
func IsServerErrorTableHasNoFt(err error) bool {
    result := Isa(err, ER_TABLE_HAS_NO_FT)
    return result
}

    
// IsServerErrorVariableNotSettableInSfOrTrigger check mysql error is "The system variable %s cannot be set in stored functionsor triggers." 
func IsServerErrorVariableNotSettableInSfOrTrigger(err error) bool {
    result := Isa(err, ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER)
    return result
}

    
// IsServerErrorVariableNotSettableInTransaction check mysql error is "The system variable %s cannot be set when there is anongoing transaction." 
func IsServerErrorVariableNotSettableInTransaction(err error) bool {
    result := Isa(err, ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION)
    return result
}

    
// IsServerErrorGtidNextIsNotInGtidNextList check mysql error is "The system variable @@SESSION.GTID_NEXT has the value %s,which is not listed in @@SESSION.GTID_NEXT_LIST." 
func IsServerErrorGtidNextIsNotInGtidNextList(err error) bool {
    result := Isa(err, ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST)
    return result
}

    
// IsServerErrorCantChangeGtidNextInTransactionWhenGtidNextListIsNull check mysql error is "The system variable @@SESSION.GTID_NEXT cannot changeinside a transaction." 
func IsServerErrorCantChangeGtidNextInTransactionWhenGtidNextListIsNull(err error) bool {
    result := Isa(err, ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION_WHEN_GTID_NEXT_LIST_IS_NULL)
    return result
}

    
// IsServerErrorSetStatementCannotInvokeFunction check mysql error is "The statement 'SET %s' cannot invoke a stored function." 
func IsServerErrorSetStatementCannotInvokeFunction(err error) bool {
    result := Isa(err, ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION)
    return result
}

    
// IsServerErrorGtidNextCantBeAutomaticIfGtidNextListIsNonNull check mysql error is "The system variable @@SESSION.GTID_NEXT cannot be'AUTOMATIC' when @@SESSION.GTID_NEXT_LIST is non-NULL." 
func IsServerErrorGtidNextCantBeAutomaticIfGtidNextListIsNonNull(err error) bool {
    result := Isa(err, ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL)
    return result
}

    
// IsServerErrorSkippingLoggedTransaction check mysql error is "Skipping transaction %s because it has already beenexecuted and logged." 
func IsServerErrorSkippingLoggedTransaction(err error) bool {
    result := Isa(err, ER_SKIPPING_LOGGED_TRANSACTION)
    return result
}

    
// IsServerErrorMalformedGtidSetSpecification check mysql error is "Malformed GTID set specification '%s'." 
func IsServerErrorMalformedGtidSetSpecification(err error) bool {
    result := Isa(err, ER_MALFORMED_GTID_SET_SPECIFICATION)
    return result
}

    
// IsServerErrorMalformedGtidSetEncoding check mysql error is "Malformed GTID set encoding." 
func IsServerErrorMalformedGtidSetEncoding(err error) bool {
    result := Isa(err, ER_MALFORMED_GTID_SET_ENCODING)
    return result
}

    
// IsServerErrorMalformedGtidSpecification check mysql error is "Malformed GTID specification '%s'." 
func IsServerErrorMalformedGtidSpecification(err error) bool {
    result := Isa(err, ER_MALFORMED_GTID_SPECIFICATION)
    return result
}

    
// IsServerErrorGnoExhausted check mysql error is "Impossible to generate Global Transaction Identifier: theinteger component reached the maximal value. Restart the serverwith a new server_uuid." 
func IsServerErrorGnoExhausted(err error) bool {
    result := Isa(err, ER_GNO_EXHAUSTED)
    return result
}

    
// IsServerErrorBadSlaveAutoPosition check mysql error is "Parameters MASTER_LOG_FILE, MASTER_LOG_POS,RELAY_LOG_FILE and RELAY_LOG_POS cannot be set whenMASTER_AUTO_POSITION is active." 
func IsServerErrorBadSlaveAutoPosition(err error) bool {
    result := Isa(err, ER_BAD_SLAVE_AUTO_POSITION)
    return result
}

    
// IsServerErrorAutoPositionRequiresGtidModeOn check mysql error is "CHANGE MASTER TO MASTER_AUTO_POSITION = 1 can only beexecuted when @@GLOBAL.GTID_MODE = ON." 
func IsServerErrorAutoPositionRequiresGtidModeOn(err error) bool {
    result := Isa(err, ER_AUTO_POSITION_REQUIRES_GTID_MODE_ON)
    return result
}

    
// IsServerErrorCantDoImplicitCommitInTrxWhenGtidNextIsSet check mysql error is "Cannot execute statements with implicit commit inside atransaction when @@SESSION.GTID_NEXT != AUTOMATIC." 
func IsServerErrorCantDoImplicitCommitInTrxWhenGtidNextIsSet(err error) bool {
    result := Isa(err, ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET)
    return result
}

    
// IsServerErrorGtidMode2Or3RequiresEnforceGtidConsistencyOn check mysql error is "@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_2 requires@@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1." 
func IsServerErrorGtidMode2Or3RequiresEnforceGtidConsistencyOn(err error) bool {
    result := Isa(err, ER_GTID_MODE_2_OR_3_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON)
    return result
}

    
// IsServerErrorGtidModeRequiresBinlog check mysql error is "@@GLOBAL.GTID_MODE = ON or UPGRADE_STEP_1 orUPGRADE_STEP_2 requires --log-bin and --log-slave-updates." 
func IsServerErrorGtidModeRequiresBinlog(err error) bool {
    result := Isa(err, ER_GTID_MODE_REQUIRES_BINLOG)
    return result
}

    
// IsServerErrorCantSetGtidNextToGtidWhenGtidModeIsOff check mysql error is "@@SESSION.GTID_NEXT cannot be set to UUID:NUMBER when@@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorCantSetGtidNextToGtidWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorCantSetGtidNextToAnonymousWhenGtidModeIsOn check mysql error is "@@SESSION.GTID_NEXT cannot be set to ANONYMOUS when@@GLOBAL.GTID_MODE = ON." 
func IsServerErrorCantSetGtidNextToAnonymousWhenGtidModeIsOn(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON)
    return result
}

    
// IsServerErrorCantSetGtidNextListToNonNullWhenGtidModeIsOff check mysql error is "@@SESSION.GTID_NEXT_LIST cannot be set to a non-NULLvalue when @@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorCantSetGtidNextListToNonNullWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorFoundGtidEventWhenGtidModeIsOff check mysql error is "Found a Gtid_log_event or Previous_gtids_log_event when@@GLOBAL.GTID_MODE = OFF." 
func IsServerErrorFoundGtidEventWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorGtidUnsafeNonTransactionalTable check mysql error is "When @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1, updates tonon-transactional tables can only be done in either autocommittedstatements or single-statement transactions, and never in the samestatement as updates to transactional tables." 
func IsServerErrorGtidUnsafeNonTransactionalTable(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE)
    return result
}

    
// IsServerErrorGtidUnsafeCreateSelect check mysql error is "CREATE TABLE ... SELECT is forbidden when@@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1." 
func IsServerErrorGtidUnsafeCreateSelect(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_CREATE_SELECT)
    return result
}

    
// IsServerErrorGtidUnsafeCreateDropTemporaryTableInTransaction check mysql error is "When @@GLOBAL.ENFORCE_GTID_CONSISTENCY = 1, thestatements CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can beexecuted in a non-transactional context only, and require thatAUTOCOMMIT = 1. These statements are also not allowed in afunction or trigger because functions and triggers are alsoconsidered to be multi-statement transactions." 
func IsServerErrorGtidUnsafeCreateDropTemporaryTableInTransaction(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_CREATE_DROP_TEMPORARY_TABLE_IN_TRANSACTION)
    return result
}

    
// IsServerErrorGtidModeCanOnlyChangeOneStepAtATime check mysql error is "The value of @@GLOBAL.GTID_MODE can only change one stepat a time: OFF <-> UPGRADE_STEP_1 <-> UPGRADE_STEP_2<-> ON. Also note that this value must be stepped up or downsimultaneously on all servers" 
func IsServerErrorGtidModeCanOnlyChangeOneStepAtATime(err error) bool {
    result := Isa(err, ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME)
    return result
}

    
// IsServerErrorMasterHasPurgedRequiredGtids check mysql error is "The slave is connecting using CHANGE MASTER TOMASTER_AUTO_POSITION = 1, but the master has purged binary logscontaining GTIDs that the slave requires. Replicate the missingtransactions from elsewhere, or provision a new slave from backup.Consider increasing the master's binary log expiration period. %s." 
func IsServerErrorMasterHasPurgedRequiredGtids(err error) bool {
    result := Isa(err, ER_MASTER_HAS_PURGED_REQUIRED_GTIDS)
    return result
}

    
// IsServerErrorCantSetGtidNextWhenOwningGtid check mysql error is "@@SESSION.GTID_NEXT cannot be changed by a client thatowns a GTID. The client owns %s. Ownership is released on COMMITor ROLLBACK." 
func IsServerErrorCantSetGtidNextWhenOwningGtid(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID)
    return result
}

    
// IsServerErrorUnknownExplainFormat check mysql error is "Unknown EXPLAIN format name: '%s'" 
func IsServerErrorUnknownExplainFormat(err error) bool {
    result := Isa(err, ER_UNKNOWN_EXPLAIN_FORMAT)
    return result
}

    
// IsServerErrorCantExecuteInReadOnlyTransaction check mysql error is "Cannot execute statement in a READ ONLY transaction." 
func IsServerErrorCantExecuteInReadOnlyTransaction(err error) bool {
    result := Isa(err, ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION)
    return result
}

    
// IsServerErrorTooLongTablePartitionComment check mysql error is "Comment for table partition '%s' is too long (max = %lu)" 
func IsServerErrorTooLongTablePartitionComment(err error) bool {
    result := Isa(err, ER_TOO_LONG_TABLE_PARTITION_COMMENT)
    return result
}

    
// IsServerErrorSlaveConfiguration check mysql error is "Slave is not configured or failed to initialize properly.You must at least set --server-id to enable either a master or aslave. Additional error messages can be found in the MySQL errorlog." 
func IsServerErrorSlaveConfiguration(err error) bool {
    result := Isa(err, ER_SLAVE_CONFIGURATION)
    return result
}

    
// IsServerErrorInnodbFtLimit check mysql error is "InnoDB presently supports one FULLTEXT index creation ata time" 
func IsServerErrorInnodbFtLimit(err error) bool {
    result := Isa(err, ER_INNODB_FT_LIMIT)
    return result
}

    
// IsServerErrorInnodbNoFtTempTable check mysql error is "Cannot create FULLTEXT index on temporary InnoDB table" 
func IsServerErrorInnodbNoFtTempTable(err error) bool {
    result := Isa(err, ER_INNODB_NO_FT_TEMP_TABLE)
    return result
}

    
// IsServerErrorInnodbFtWrongDocidColumn check mysql error is "Column '%s' is of wrong type for an InnoDB FULLTEXT index" 
func IsServerErrorInnodbFtWrongDocidColumn(err error) bool {
    result := Isa(err, ER_INNODB_FT_WRONG_DOCID_COLUMN)
    return result
}

    
// IsServerErrorInnodbFtWrongDocidIndex check mysql error is "Index '%s' is of wrong type for an InnoDB FULLTEXT index" 
func IsServerErrorInnodbFtWrongDocidIndex(err error) bool {
    result := Isa(err, ER_INNODB_FT_WRONG_DOCID_INDEX)
    return result
}

    
// IsServerErrorInnodbOnlineLogTooBig check mysql error is "Creating index '%s' required more than'innodb_online_alter_log_max_size' bytes of modification log.Please try again." 
func IsServerErrorInnodbOnlineLogTooBig(err error) bool {
    result := Isa(err, ER_INNODB_ONLINE_LOG_TOO_BIG)
    return result
}

    
// IsServerErrorUnknownAlterAlgorithm check mysql error is "Unknown ALGORITHM '%s'" 
func IsServerErrorUnknownAlterAlgorithm(err error) bool {
    result := Isa(err, ER_UNKNOWN_ALTER_ALGORITHM)
    return result
}

    
// IsServerErrorUnknownAlterLock check mysql error is "Unknown LOCK type '%s'" 
func IsServerErrorUnknownAlterLock(err error) bool {
    result := Isa(err, ER_UNKNOWN_ALTER_LOCK)
    return result
}

    
// IsServerErrorMtsChangeMasterCantRunWithGaps check mysql error is "CHANGE MASTER cannot be executed when the slave wasstopped with an error or killed in MTS mode. Consider using RESETSLAVE or START SLAVE UNTIL." 
func IsServerErrorMtsChangeMasterCantRunWithGaps(err error) bool {
    result := Isa(err, ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS)
    return result
}

    
// IsServerErrorMtsRecoveryFailure check mysql error is "Cannot recover after SLAVE errored out in parallelexecution mode. Additional error messages can be found in theMySQL error log." 
func IsServerErrorMtsRecoveryFailure(err error) bool {
    result := Isa(err, ER_MTS_RECOVERY_FAILURE)
    return result
}

    
// IsServerErrorMtsResetWorkers check mysql error is "Cannot clean up worker info tables. Additional errormessages can be found in the MySQL error log." 
func IsServerErrorMtsResetWorkers(err error) bool {
    result := Isa(err, ER_MTS_RESET_WORKERS)
    return result
}

    
// IsServerErrorColCountDoesntMatchCorruptedV2 check mysql error is "Column count of %s.%s is wrong. Expected %d, found %d.The table is probably corrupted" 
func IsServerErrorColCountDoesntMatchCorruptedV2(err error) bool {
    result := Isa(err, ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2)
    return result
}

    
// IsServerErrorSlaveSilentRetryTransaction check mysql error is "Slave must silently retry current transaction" 
func IsServerErrorSlaveSilentRetryTransaction(err error) bool {
    result := Isa(err, ER_SLAVE_SILENT_RETRY_TRANSACTION)
    return result
}

    
// IsServerErrorDiscardFkChecksRunning check mysql error is "There is a foreign key check running on table '%s'.Cannot discard the table." 
func IsServerErrorDiscardFkChecksRunning(err error) bool {
    result := Isa(err, ER_DISCARD_FK_CHECKS_RUNNING)
    return result
}

    
// IsServerErrorTableSchemaMismatch check mysql error is "Schema mismatch (%s)" 
func IsServerErrorTableSchemaMismatch(err error) bool {
    result := Isa(err, ER_TABLE_SCHEMA_MISMATCH)
    return result
}

    
// IsServerErrorTableInSystemTablespace check mysql error is "Table '%s' in system tablespace" 
func IsServerErrorTableInSystemTablespace(err error) bool {
    result := Isa(err, ER_TABLE_IN_SYSTEM_TABLESPACE)
    return result
}

    
// IsServerErrorIoReadError check mysql error is "IO Read error: (%lu, %s) %s" 
func IsServerErrorIoReadError(err error) bool {
    result := Isa(err, ER_IO_READ_ERROR)
    return result
}

    
// IsServerErrorIoWriteError check mysql error is "IO Write error: (%lu, %s) %s" 
func IsServerErrorIoWriteError(err error) bool {
    result := Isa(err, ER_IO_WRITE_ERROR)
    return result
}

    
// IsServerErrorTablespaceMissing check mysql error is "Tablespace is missing for table '%s'" 
func IsServerErrorTablespaceMissing(err error) bool {
    result := Isa(err, ER_TABLESPACE_MISSING)
    return result
}

    
// IsServerErrorTablespaceExists check mysql error is "Tablespace for table '%s' exists. Please DISCARD thetablespace before IMPORT." 
func IsServerErrorTablespaceExists(err error) bool {
    result := Isa(err, ER_TABLESPACE_EXISTS)
    return result
}

    
// IsServerErrorTablespaceDiscarded check mysql error is "Tablespace has been discarded for table '%s'" 
func IsServerErrorTablespaceDiscarded(err error) bool {
    result := Isa(err, ER_TABLESPACE_DISCARDED)
    return result
}

    
// IsServerErrorInternalError check mysql error is "Internal error: %s" 
func IsServerErrorInternalError(err error) bool {
    result := Isa(err, ER_INTERNAL_ERROR)
    return result
}

    
// IsServerErrorInnodbImportError check mysql error is "ALTER TABLE '%s' IMPORT TABLESPACE failed with error %lu: '%s'" 
func IsServerErrorInnodbImportError(err error) bool {
    result := Isa(err, ER_INNODB_IMPORT_ERROR)
    return result
}

    
// IsServerErrorInnodbIndexCorrupt check mysql error is "Index corrupt: %s" 
func IsServerErrorInnodbIndexCorrupt(err error) bool {
    result := Isa(err, ER_INNODB_INDEX_CORRUPT)
    return result
}

    
// IsServerErrorInvalidYearColumnLength check mysql error is "YEAR(%lu) column type is deprecated. Creating YEAR(4)column instead." 
func IsServerErrorInvalidYearColumnLength(err error) bool {
    result := Isa(err, ER_INVALID_YEAR_COLUMN_LENGTH)
    return result
}

    
// IsServerErrorNotValidPassword check mysql error is "Your password does not satisfy the current policyrequirements" 
func IsServerErrorNotValidPassword(err error) bool {
    result := Isa(err, ER_NOT_VALID_PASSWORD)
    return result
}

    
// IsServerErrorMustChangePassword check mysql error is "You must SET PASSWORD before executing this statement" 
func IsServerErrorMustChangePassword(err error) bool {
    result := Isa(err, ER_MUST_CHANGE_PASSWORD)
    return result
}

    
// IsServerErrorFkNoIndexChild check mysql error is "Failed to add the foreign key constaint. Missing indexfor constraint '%s' in the foreign table '%s'" 
func IsServerErrorFkNoIndexChild(err error) bool {
    result := Isa(err, ER_FK_NO_INDEX_CHILD)
    return result
}

    
// IsServerErrorFkNoIndexParent check mysql error is "Failed to add the foreign key constaint. Missing indexfor constraint '%s' in the referenced table '%s'" 
func IsServerErrorFkNoIndexParent(err error) bool {
    result := Isa(err, ER_FK_NO_INDEX_PARENT)
    return result
}

    
// IsServerErrorFkFailAddSystem check mysql error is "Failed to add the foreign key constraint '%s' to systemtables" 
func IsServerErrorFkFailAddSystem(err error) bool {
    result := Isa(err, ER_FK_FAIL_ADD_SYSTEM)
    return result
}

    
// IsServerErrorFkCannotOpenParent check mysql error is "Failed to open the referenced table '%s'" 
func IsServerErrorFkCannotOpenParent(err error) bool {
    result := Isa(err, ER_FK_CANNOT_OPEN_PARENT)
    return result
}

    
// IsServerErrorFkIncorrectOption check mysql error is "Failed to add the foreign key constraint on table '%s'.Incorrect options in FOREIGN KEY constraint '%s'" 
func IsServerErrorFkIncorrectOption(err error) bool {
    result := Isa(err, ER_FK_INCORRECT_OPTION)
    return result
}

    
// IsServerErrorFkDupName check mysql error is "Duplicate foreign key constraint name '%s'" 
func IsServerErrorFkDupName(err error) bool {
    result := Isa(err, ER_FK_DUP_NAME)
    return result
}

    
// IsServerErrorPasswordFormat check mysql error is "The password hash doesn't have the expected format. Checkif the correct password algorithm is being used with thePASSWORD() function." 
func IsServerErrorPasswordFormat(err error) bool {
    result := Isa(err, ER_PASSWORD_FORMAT)
    return result
}

    
// IsServerErrorFkColumnCannotDrop check mysql error is "Cannot drop column '%s': needed in a foreign keyconstraint '%s'" 
func IsServerErrorFkColumnCannotDrop(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_DROP)
    return result
}

    
// IsServerErrorFkColumnCannotDropChild check mysql error is "Cannot drop column '%s': needed in a foreign keyconstraint '%s' of table '%s'" 
func IsServerErrorFkColumnCannotDropChild(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_DROP_CHILD)
    return result
}

    
// IsServerErrorFkColumnNotNull check mysql error is "Column '%s' cannot be NOT NULL: needed in a foreign keyconstraint '%s' SET NULL" 
func IsServerErrorFkColumnNotNull(err error) bool {
    result := Isa(err, ER_FK_COLUMN_NOT_NULL)
    return result
}

    
// IsServerErrorDupIndex check mysql error is "Duplicate index '%s' defined on the table '%s.%s'. Thisis deprecated and will be disallowed in a future release." 
func IsServerErrorDupIndex(err error) bool {
    result := Isa(err, ER_DUP_INDEX)
    return result
}

    
// IsServerErrorFkColumnCannotChange check mysql error is "Cannot change column '%s': used in a foreign keyconstraint '%s'" 
func IsServerErrorFkColumnCannotChange(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_CHANGE)
    return result
}

    
// IsServerErrorFkColumnCannotChangeChild check mysql error is "Cannot change column '%s': used in a foreign keyconstraint '%s' of table '%s'" 
func IsServerErrorFkColumnCannotChangeChild(err error) bool {
    result := Isa(err, ER_FK_COLUMN_CANNOT_CHANGE_CHILD)
    return result
}

    
// IsServerErrorFkCannotDeleteParent check mysql error is "Cannot delete rows from table which is parent in aforeign key constraint '%s' of table '%s'" 
func IsServerErrorFkCannotDeleteParent(err error) bool {
    result := Isa(err, ER_FK_CANNOT_DELETE_PARENT)
    return result
}

    
// IsServerErrorMalformedPacket check mysql error is "Malformed communication packet." 
func IsServerErrorMalformedPacket(err error) bool {
    result := Isa(err, ER_MALFORMED_PACKET)
    return result
}

    
// IsServerErrorReadOnlyMode check mysql error is "Running in read-only mode" 
func IsServerErrorReadOnlyMode(err error) bool {
    result := Isa(err, ER_READ_ONLY_MODE)
    return result
}

    
// IsServerErrorGtidNextTypeUndefinedGroup check mysql error is "When @@SESSION.GTID_NEXT is set to a GTID, you mustexplicitly set it to a different value after a COMMIT or ROLLBACK.Please check GTID_NEXT variable manual page for detailedexplanation. Current @@SESSION.GTID_NEXT is '%s'." 
func IsServerErrorGtidNextTypeUndefinedGroup(err error) bool {
    result := Isa(err, ER_GTID_NEXT_TYPE_UNDEFINED_GROUP)
    return result
}

    
// IsServerErrorVariableNotSettableInSp check mysql error is "The system variable %s cannot be set in storedprocedures." 
func IsServerErrorVariableNotSettableInSp(err error) bool {
    result := Isa(err, ER_VARIABLE_NOT_SETTABLE_IN_SP)
    return result
}

    
// IsServerErrorCantSetGtidPurgedWhenGtidModeIsOff check mysql error is "@@GLOBAL.GTID_PURGED can only be set when@@GLOBAL.GTID_MODE = ON." 
func IsServerErrorCantSetGtidPurgedWhenGtidModeIsOff(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF)
    return result
}

    
// IsServerErrorCantSetGtidPurgedWhenGtidExecutedIsNotEmpty check mysql error is "@@GLOBAL.GTID_PURGED can only be set when@@GLOBAL.GTID_EXECUTED is empty." 
func IsServerErrorCantSetGtidPurgedWhenGtidExecutedIsNotEmpty(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY)
    return result
}

    
// IsServerErrorCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty check mysql error is "@@GLOBAL.GTID_PURGED can only be set when there are noongoing transactions (not even in other clients)." 
func IsServerErrorCantSetGtidPurgedWhenOwnedGtidsIsNotEmpty(err error) bool {
    result := Isa(err, ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY)
    return result
}

    
// IsServerErrorGtidPurgedWasChanged check mysql error is "@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'." 
func IsServerErrorGtidPurgedWasChanged(err error) bool {
    result := Isa(err, ER_GTID_PURGED_WAS_CHANGED)
    return result
}

    
// IsServerErrorGtidExecutedWasChanged check mysql error is "@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'." 
func IsServerErrorGtidExecutedWasChanged(err error) bool {
    result := Isa(err, ER_GTID_EXECUTED_WAS_CHANGED)
    return result
}

    
// IsServerErrorBinlogStmtModeAndNoReplTables check mysql error is "Cannot execute statement: impossible to write to binarylog since BINLOG_FORMAT = STATEMENT, and both replicated and nonreplicated tables are written to." 
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

    
// IsServerErrorAlterOperationNotSupportedReasonPartition check mysql error is "Partition specific operations do not yet supportLOCK/ALGORITHM" 
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

    
// IsServerErrorAlterOperationNotSupportedReasonIgnore check mysql error is "Creating unique indexes with IGNORE requires COPYalgorithm to remove duplicate rows" 
func IsServerErrorAlterOperationNotSupportedReasonIgnore(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_IGNORE)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonNopk check mysql error is "Dropping a primary key is not allowed without also addinga new primary key" 
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

    
// IsServerErrorSqlSlaveSkipCounterNotSettableInGtidMode check mysql error is "sql_slave_skip_counter can not be set when the server isrunning with @@GLOBAL.GTID_MODE = ON. Instead, for eachtransaction that you want to skip, generate an empty transactionwith the same GTID as the transaction" 
func IsServerErrorSqlSlaveSkipCounterNotSettableInGtidMode(err error) bool {
    result := Isa(err, ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE)
    return result
}

    
// IsServerErrorDupUnknownInIndex check mysql error is "Duplicate entry for key '%s'" 
func IsServerErrorDupUnknownInIndex(err error) bool {
    result := Isa(err, ER_DUP_UNKNOWN_IN_INDEX)
    return result
}

    
// IsServerErrorIdentCausesTooLongPath check mysql error is "Long database name and identifier for object resulted inpath length exceeding %d characters. Path: '%s'." 
func IsServerErrorIdentCausesTooLongPath(err error) bool {
    result := Isa(err, ER_IDENT_CAUSES_TOO_LONG_PATH)
    return result
}

    
// IsServerErrorAlterOperationNotSupportedReasonNotNull check mysql error is "cannot silently convert NULL values, as required in thisSQL_MODE" 
func IsServerErrorAlterOperationNotSupportedReasonNotNull(err error) bool {
    result := Isa(err, ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL)
    return result
}

    
// IsServerErrorMustChangePasswordLogin check mysql error is "Your password has expired. To log in you must change itusing a client that supports expired passwords." 
func IsServerErrorMustChangePasswordLogin(err error) bool {
    result := Isa(err, ER_MUST_CHANGE_PASSWORD_LOGIN)
    return result
}

    
// IsServerErrorRowInWrongPartition check mysql error is "Found a row in wrong partition %s" 
func IsServerErrorRowInWrongPartition(err error) bool {
    result := Isa(err, ER_ROW_IN_WRONG_PARTITION)
    return result
}

    
// IsServerErrorMtsEventBiggerPendingJobsSizeMax check mysql error is "Cannot schedule event %s, relay-log name %s, position %sto Worker thread because its size %lu exceeds %lu ofslave_pending_jobs_size_max." 
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

    
// IsServerErrorWarnPurgeLogInUse check mysql error is "file %s was not purged because it was being read by %dthread(s), purged only %d out of %d files." 
func IsServerErrorWarnPurgeLogInUse(err error) bool {
    result := Isa(err, ER_WARN_PURGE_LOG_IN_USE)
    return result
}

    
// IsServerErrorWarnPurgeLogIsActive check mysql error is "file %s was not purged because it is the active log file." 
func IsServerErrorWarnPurgeLogIsActive(err error) bool {
    result := Isa(err, ER_WARN_PURGE_LOG_IS_ACTIVE)
    return result
}

    
// IsServerErrorAutoIncrementConflict check mysql error is "Auto-increment value in UPDATE conflicts with internallygenerated values" 
func IsServerErrorAutoIncrementConflict(err error) bool {
    result := Isa(err, ER_AUTO_INCREMENT_CONFLICT)
    return result
}

    
// IsWarnOnBlockholeInRbr check mysql error is "Row events are not logged for %s statements that modifyBLACKHOLE tables in row format. Table(s): '%s'" 
func IsWarnOnBlockholeInRbr(err error) bool {
    result := Isa(err, WARN_ON_BLOCKHOLE_IN_RBR)
    return result
}

    
// IsServerErrorSlaveMiInitRepository check mysql error is "Slave failed to initialize master info structure from therepository" 
func IsServerErrorSlaveMiInitRepository(err error) bool {
    result := Isa(err, ER_SLAVE_MI_INIT_REPOSITORY)
    return result
}

    
// IsServerErrorSlaveRliInitRepository check mysql error is "Slave failed to initialize relay log info structure fromthe repository" 
func IsServerErrorSlaveRliInitRepository(err error) bool {
    result := Isa(err, ER_SLAVE_RLI_INIT_REPOSITORY)
    return result
}

    
// IsServerErrorAccessDeniedChangeUserError check mysql error is "Access denied trying to change to user '%s'@'%s' (usingpassword: %s). Disconnecting." 
func IsServerErrorAccessDeniedChangeUserError(err error) bool {
    result := Isa(err, ER_ACCESS_DENIED_CHANGE_USER_ERROR)
    return result
}

    
// IsServerErrorInnodbReadOnly check mysql error is "InnoDB is in read only mode." 
func IsServerErrorInnodbReadOnly(err error) bool {
    result := Isa(err, ER_INNODB_READ_ONLY)
    return result
}

    
// IsServerErrorStopSlaveSqlThreadTimeout check mysql error is "STOP SLAVE command execution is incomplete: Slave SQLthread got the stop signal, thread is busy, SQL thread will stoponce the current task is complete." 
func IsServerErrorStopSlaveSqlThreadTimeout(err error) bool {
    result := Isa(err, ER_STOP_SLAVE_SQL_THREAD_TIMEOUT)
    return result
}

    
// IsServerErrorStopSlaveIoThreadTimeout check mysql error is "STOP SLAVE command execution is incomplete: Slave IOthread got the stop signal, thread is busy, IO thread will stoponce the current task is complete." 
func IsServerErrorStopSlaveIoThreadTimeout(err error) bool {
    result := Isa(err, ER_STOP_SLAVE_IO_THREAD_TIMEOUT)
    return result
}

    
// IsServerErrorTableCorrupt check mysql error is "Operation cannot be performed. The table '%s.%s' ismissing, corrupt or contains bad data." 
func IsServerErrorTableCorrupt(err error) bool {
    result := Isa(err, ER_TABLE_CORRUPT)
    return result
}

    
// IsServerErrorTempFileWriteFailure check mysql error is "Temporary file write failure." 
func IsServerErrorTempFileWriteFailure(err error) bool {
    result := Isa(err, ER_TEMP_FILE_WRITE_FAILURE)
    return result
}

    
// IsServerErrorInnodbFtAuxNotHexId check mysql error is "Upgrade index name failed, please use create index(altertable) algorithm copy to rebuild index." 
func IsServerErrorInnodbFtAuxNotHexId(err error) bool {
    result := Isa(err, ER_INNODB_FT_AUX_NOT_HEX_ID)
    return result
}

    
// IsServerErrorOldTemporalsUpgraded check mysql error is "TIME/TIMESTAMP/DATETIME columns of old format have beenupgraded to the new format." 
func IsServerErrorOldTemporalsUpgraded(err error) bool {
    result := Isa(err, ER_OLD_TEMPORALS_UPGRADED)
    return result
}

    
// IsServerErrorInnodbForcedRecovery check mysql error is "Operation not allowed when innodb_forced_recovery > 0." 
func IsServerErrorInnodbForcedRecovery(err error) bool {
    result := Isa(err, ER_INNODB_FORCED_RECOVERY)
    return result
}

    
// IsServerErrorAesInvalidIv check mysql error is "The initialization vector supplied to %s is too short.Must be at least %d bytes long" 
func IsServerErrorAesInvalidIv(err error) bool {
    result := Isa(err, ER_AES_INVALID_IV)
    return result
}

    
// IsServerErrorPluginCannotBeUninstalled check mysql error is "Plugin '%s' cannot be uninstalled now. %s" 
func IsServerErrorPluginCannotBeUninstalled(err error) bool {
    result := Isa(err, ER_PLUGIN_CANNOT_BE_UNINSTALLED)
    return result
}

    
// IsServerErrorGtidUnsafeBinlogSplittableStatementAndGtidGroup check mysql error is "Cannot execute statement because it needs to be writtento the binary log as multiple statements, and this is not allowedwhen @@SESSION.GTID_NEXT == 'UUID:NUMBER'." 
func IsServerErrorGtidUnsafeBinlogSplittableStatementAndGtidGroup(err error) bool {
    result := Isa(err, ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_GTID_GROUP)
    return result
}

    
// IsServerErrorSlaveHasMoreGtidsThanMaster check mysql error is "Slave has more GTIDs than the master has, using themaster's SERVER_UUID. This may indicate that the end of the binarylog was truncated or that the last binary log file was lost, e.g.,after a power or disk failure when sync_binlog != 1. The mastermay or may not have rolled back transactions that were alreadyreplicated to the slave. Suggest to replicate any transactionsthat master has rolled back from slave to master, and/or commitempty transactions on master to account for transactions that havebeen committed on master but are not included in GTID_EXECUTED." 
func IsServerErrorSlaveHasMoreGtidsThanMaster(err error) bool {
    result := Isa(err, ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER)
    return result
}

    
// IsServerErrorMissingKey check mysql error is "The table '%s.%s' does not have the necessary key(s)defined on it. Please check the table definition and createindex(s) accordingly." 
func IsServerErrorMissingKey(err error) bool {
    result := Isa(err, ER_MISSING_KEY)
    return result
}

    
// IsWarnNamedPipeAccessEveryone check mysql error is "Setting named_pipe_full_access_group='%s' is insecure.Consider using a Windows group with fewer members." 
func IsWarnNamedPipeAccessEveryone(err error) bool {
    result := Isa(err, WARN_NAMED_PIPE_ACCESS_EVERYONE)
    return result
}

    
// IsServerErrorFoundMissingGtids check mysql error is "Cannot replicate to server with server_uuid='%s' becausethe present server has purged required binary logs. The connectingserver needs to replicate the missing transactions from elsewhere,or be replaced by a new server created from a more recent backup.To prevent this error in the future, consider increasing thebinary log expiration period on the present server. %s." 
func IsServerErrorFoundMissingGtids(err error) bool {
    result := Isa(err, ER_FOUND_MISSING_GTIDS)
    return result
}

    
// IsClientErrorUnknownError check mysql error is "Unknown MySQL error" 
func IsClientErrorUnknownError(err error) bool {
    result := Isa(err, CR_UNKNOWN_ERROR)
    return result
}

    
// IsClientErrorSocketCreateError check mysql error is "Can't create UNIX socket (%d)" 
func IsClientErrorSocketCreateError(err error) bool {
    result := Isa(err, CR_SOCKET_CREATE_ERROR)
    return result
}

    
// IsClientErrorConnectionError check mysql error is "Can't connect to local MySQL server through socket '%s'(%d)" 
func IsClientErrorConnectionError(err error) bool {
    result := Isa(err, CR_CONNECTION_ERROR)
    return result
}

    
// IsClientErrorConnHostError check mysql error is "Can't connect to MySQL server on '%s' (%d)" 
func IsClientErrorConnHostError(err error) bool {
    result := Isa(err, CR_CONN_HOST_ERROR)
    return result
}

    
// IsClientErrorIpsockError check mysql error is "Can't create TCP/IP socket (%d)" 
func IsClientErrorIpsockError(err error) bool {
    result := Isa(err, CR_IPSOCK_ERROR)
    return result
}

    
// IsClientErrorUnknownHost check mysql error is "Unknown MySQL server host '%s' (%d)" 
func IsClientErrorUnknownHost(err error) bool {
    result := Isa(err, CR_UNKNOWN_HOST)
    return result
}

    
// IsClientErrorServerGoneError check mysql error is "MySQL server has gone away" 
func IsClientErrorServerGoneError(err error) bool {
    result := Isa(err, CR_SERVER_GONE_ERROR)
    return result
}

    
// IsClientErrorVersionError check mysql error is "Protocol mismatch" 
func IsClientErrorVersionError(err error) bool {
    result := Isa(err, CR_VERSION_ERROR)
    return result
}

    
// IsClientErrorOutOfMemory check mysql error is "MySQL client ran out of memory" 
func IsClientErrorOutOfMemory(err error) bool {
    result := Isa(err, CR_OUT_OF_MEMORY)
    return result
}

    
// IsClientErrorWrongHostInfo check mysql error is "Wrong host info" 
func IsClientErrorWrongHostInfo(err error) bool {
    result := Isa(err, CR_WRONG_HOST_INFO)
    return result
}

    
// IsClientErrorLocalhostConnection check mysql error is "Localhost via UNIX socket" 
func IsClientErrorLocalhostConnection(err error) bool {
    result := Isa(err, CR_LOCALHOST_CONNECTION)
    return result
}

    
// IsClientErrorTcpConnection check mysql error is "%s via TCP/IP" 
func IsClientErrorTcpConnection(err error) bool {
    result := Isa(err, CR_TCP_CONNECTION)
    return result
}

    
// IsClientErrorServerHandshakeErr check mysql error is "Error in server handshake" 
func IsClientErrorServerHandshakeErr(err error) bool {
    result := Isa(err, CR_SERVER_HANDSHAKE_ERR)
    return result
}

    
// IsClientErrorServerLost check mysql error is "Lost connection to MySQL server during query" 
func IsClientErrorServerLost(err error) bool {
    result := Isa(err, CR_SERVER_LOST)
    return result
}

    
// IsClientErrorCommandsOutOfSync check mysql error is "Commands out of sync" 
func IsClientErrorCommandsOutOfSync(err error) bool {
    result := Isa(err, CR_COMMANDS_OUT_OF_SYNC)
    return result
}

    
// IsClientErrorNamedpipeConnection check mysql error is "Named pipe: %s" 
func IsClientErrorNamedpipeConnection(err error) bool {
    result := Isa(err, CR_NAMEDPIPE_CONNECTION)
    return result
}

    
// IsClientErrorNamedpipewaitError check mysql error is "Can't wait for named pipe to host: %s pipe: %s (%lu)" 
func IsClientErrorNamedpipewaitError(err error) bool {
    result := Isa(err, CR_NAMEDPIPEWAIT_ERROR)
    return result
}

    
// IsClientErrorNamedpipeopenError check mysql error is "Can't open named pipe to host: %s pipe: %s (%lu)" 
func IsClientErrorNamedpipeopenError(err error) bool {
    result := Isa(err, CR_NAMEDPIPEOPEN_ERROR)
    return result
}

    
// IsClientErrorNamedpipesetstateError check mysql error is "Can't set state of named pipe to host: %s pipe: %s (%lu)" 
func IsClientErrorNamedpipesetstateError(err error) bool {
    result := Isa(err, CR_NAMEDPIPESETSTATE_ERROR)
    return result
}

    
// IsClientErrorCantReadCharset check mysql error is "Can't initialize character set %s (path: %s)" 
func IsClientErrorCantReadCharset(err error) bool {
    result := Isa(err, CR_CANT_READ_CHARSET)
    return result
}

    
// IsClientErrorNetPacketTooLarge check mysql error is "Got packet bigger than 'max_allowed_packet' bytes" 
func IsClientErrorNetPacketTooLarge(err error) bool {
    result := Isa(err, CR_NET_PACKET_TOO_LARGE)
    return result
}

    
// IsClientErrorEmbeddedConnection check mysql error is "Embedded server" 
func IsClientErrorEmbeddedConnection(err error) bool {
    result := Isa(err, CR_EMBEDDED_CONNECTION)
    return result
}

    
// IsClientErrorProbeSlaveStatus check mysql error is "Error on SHOW SLAVE STATUS:" 
func IsClientErrorProbeSlaveStatus(err error) bool {
    result := Isa(err, CR_PROBE_SLAVE_STATUS)
    return result
}

    
// IsClientErrorProbeSlaveHosts check mysql error is "Error on SHOW SLAVE HOSTS:" 
func IsClientErrorProbeSlaveHosts(err error) bool {
    result := Isa(err, CR_PROBE_SLAVE_HOSTS)
    return result
}

    
// IsClientErrorProbeSlaveConnect check mysql error is "Error connecting to slave:" 
func IsClientErrorProbeSlaveConnect(err error) bool {
    result := Isa(err, CR_PROBE_SLAVE_CONNECT)
    return result
}

    
// IsClientErrorProbeMasterConnect check mysql error is "Error connecting to master:" 
func IsClientErrorProbeMasterConnect(err error) bool {
    result := Isa(err, CR_PROBE_MASTER_CONNECT)
    return result
}

    
// IsClientErrorSslConnectionError check mysql error is "SSL connection error: %s" 
func IsClientErrorSslConnectionError(err error) bool {
    result := Isa(err, CR_SSL_CONNECTION_ERROR)
    return result
}

    
// IsClientErrorMalformedPacket check mysql error is "Malformed packet" 
func IsClientErrorMalformedPacket(err error) bool {
    result := Isa(err, CR_MALFORMED_PACKET)
    return result
}

    
// IsClientErrorWrongLicense check mysql error is "This client library is licensed only for use with MySQLservers having '%s' license" 
func IsClientErrorWrongLicense(err error) bool {
    result := Isa(err, CR_WRONG_LICENSE)
    return result
}

    
// IsClientErrorNullPointer check mysql error is "Invalid use of null pointer" 
func IsClientErrorNullPointer(err error) bool {
    result := Isa(err, CR_NULL_POINTER)
    return result
}

    
// IsClientErrorNoPrepareStmt check mysql error is "Statement not prepared" 
func IsClientErrorNoPrepareStmt(err error) bool {
    result := Isa(err, CR_NO_PREPARE_STMT)
    return result
}

    
// IsClientErrorParamsNotBound check mysql error is "No data supplied for parameters in prepared statement" 
func IsClientErrorParamsNotBound(err error) bool {
    result := Isa(err, CR_PARAMS_NOT_BOUND)
    return result
}

    
// IsClientErrorDataTruncated check mysql error is "Data truncated" 
func IsClientErrorDataTruncated(err error) bool {
    result := Isa(err, CR_DATA_TRUNCATED)
    return result
}

    
// IsClientErrorNoParametersExists check mysql error is "No parameters exist in the statement" 
func IsClientErrorNoParametersExists(err error) bool {
    result := Isa(err, CR_NO_PARAMETERS_EXISTS)
    return result
}

    
// IsClientErrorInvalidParameterNo check mysql error is "Invalid parameter number" 
func IsClientErrorInvalidParameterNo(err error) bool {
    result := Isa(err, CR_INVALID_PARAMETER_NO)
    return result
}

    
// IsClientErrorInvalidBufferUse check mysql error is "Can't send long data for non-string/non-binary data types(parameter: %d)" 
func IsClientErrorInvalidBufferUse(err error) bool {
    result := Isa(err, CR_INVALID_BUFFER_USE)
    return result
}

    
// IsClientErrorUnsupportedParamType check mysql error is "Using unsupported buffer type: %d (parameter: %d)" 
func IsClientErrorUnsupportedParamType(err error) bool {
    result := Isa(err, CR_UNSUPPORTED_PARAM_TYPE)
    return result
}

    
// IsClientErrorSharedMemoryConnection check mysql error is "Shared memory: %s" 
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

    
// IsClientErrorConnUnknowProtocol check mysql error is "Wrong or unknown protocol" 
func IsClientErrorConnUnknowProtocol(err error) bool {
    result := Isa(err, CR_CONN_UNKNOW_PROTOCOL)
    return result
}

    
// IsClientErrorInvalidConnHandle check mysql error is "Invalid connection handle" 
func IsClientErrorInvalidConnHandle(err error) bool {
    result := Isa(err, CR_INVALID_CONN_HANDLE)
    return result
}

    
// IsClientErrorSecureAuth check mysql error is "Connection using old (pre-4.1.1) authentication protocolrefused (client option 'secure_auth' enabled)" 
func IsClientErrorSecureAuth(err error) bool {
    result := Isa(err, CR_SECURE_AUTH)
    return result
}

    
// IsClientErrorFetchCanceled check mysql error is "Row retrieval was canceled by mysql_stmt_close() call" 
func IsClientErrorFetchCanceled(err error) bool {
    result := Isa(err, CR_FETCH_CANCELED)
    return result
}

    
// IsClientErrorNoData check mysql error is "Attempt to read column without prior row fetch" 
func IsClientErrorNoData(err error) bool {
    result := Isa(err, CR_NO_DATA)
    return result
}

    
// IsClientErrorNoStmtMetadata check mysql error is "Prepared statement contains no metadata" 
func IsClientErrorNoStmtMetadata(err error) bool {
    result := Isa(err, CR_NO_STMT_METADATA)
    return result
}

    
// IsClientErrorNoResultSet check mysql error is "Attempt to read a row while there is no result setassociated with the statement" 
func IsClientErrorNoResultSet(err error) bool {
    result := Isa(err, CR_NO_RESULT_SET)
    return result
}

    
// IsClientErrorNotImplemented check mysql error is "This feature is not implemented yet" 
func IsClientErrorNotImplemented(err error) bool {
    result := Isa(err, CR_NOT_IMPLEMENTED)
    return result
}

    
// IsClientErrorServerLostExtended check mysql error is "Lost connection to MySQL server at '%s', system error: %d" 
func IsClientErrorServerLostExtended(err error) bool {
    result := Isa(err, CR_SERVER_LOST_EXTENDED)
    return result
}

    
// IsClientErrorStmtClosed check mysql error is "Statement closed indirectly because of a preceeding %s()call" 
func IsClientErrorStmtClosed(err error) bool {
    result := Isa(err, CR_STMT_CLOSED)
    return result
}

    
// IsClientErrorNewStmtMetadata check mysql error is "The number of columns in the result set differs from thenumber of bound buffers. You must reset the statement, rebind theresult set columns, and execute the statement again" 
func IsClientErrorNewStmtMetadata(err error) bool {
    result := Isa(err, CR_NEW_STMT_METADATA)
    return result
}

    
// IsClientErrorAlreadyConnected check mysql error is "This handle is already connected. Use a separate handlefor each connection." 
func IsClientErrorAlreadyConnected(err error) bool {
    result := Isa(err, CR_ALREADY_CONNECTED)
    return result
}

    
// IsClientErrorAuthPluginCannotLoad check mysql error is "Authentication plugin '%s' cannot be loaded: %s" 
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

    
// IsGlobalErrorCantcreatefile check mysql error is "Can't create/write to file '%s' (Errcode: %d - %s)" 
func IsGlobalErrorCantcreatefile(err error) bool {
    result := Isa(err, EE_CANTCREATEFILE)
    return result
}

    
// IsGlobalErrorRead check mysql error is "Error reading file '%s' (Errcode: %d - %s)" 
func IsGlobalErrorRead(err error) bool {
    result := Isa(err, EE_READ)
    return result
}

    
// IsGlobalErrorWrite check mysql error is "Error writing file '%s' (Errcode: %d - %s)" 
func IsGlobalErrorWrite(err error) bool {
    result := Isa(err, EE_WRITE)
    return result
}

    
// IsGlobalErrorBadclose check mysql error is "Error on close of '%s' (Errcode: %d - %s)" 
func IsGlobalErrorBadclose(err error) bool {
    result := Isa(err, EE_BADCLOSE)
    return result
}

    
// IsGlobalErrorOutofmemory check mysql error is "Out of memory (Needed %u bytes)" 
func IsGlobalErrorOutofmemory(err error) bool {
    result := Isa(err, EE_OUTOFMEMORY)
    return result
}

    
// IsGlobalErrorDelete check mysql error is "Error on delete of '%s' (Errcode: %d - %s)" 
func IsGlobalErrorDelete(err error) bool {
    result := Isa(err, EE_DELETE)
    return result
}

    
// IsGlobalErrorLink check mysql error is "Error on rename of '%s' to '%s' (Errcode: %d - %s)" 
func IsGlobalErrorLink(err error) bool {
    result := Isa(err, EE_LINK)
    return result
}

    
// IsGlobalErrorEoferr check mysql error is "Unexpected EOF found when reading file '%s' (Errcode: %d- %s)" 
func IsGlobalErrorEoferr(err error) bool {
    result := Isa(err, EE_EOFERR)
    return result
}

    
// IsGlobalErrorCantlock check mysql error is "Can't lock file (Errcode: %d - %s)" 
func IsGlobalErrorCantlock(err error) bool {
    result := Isa(err, EE_CANTLOCK)
    return result
}

    
// IsGlobalErrorCantunlock check mysql error is "Can't unlock file (Errcode: %d - %s)" 
func IsGlobalErrorCantunlock(err error) bool {
    result := Isa(err, EE_CANTUNLOCK)
    return result
}

    
// IsGlobalErrorDir check mysql error is "Can't read dir of '%s' (Errcode: %d - %s)" 
func IsGlobalErrorDir(err error) bool {
    result := Isa(err, EE_DIR)
    return result
}

    
// IsGlobalErrorStat check mysql error is "Can't get stat of '%s' (Errcode: %d - %s)" 
func IsGlobalErrorStat(err error) bool {
    result := Isa(err, EE_STAT)
    return result
}

    
// IsGlobalErrorCantChsize check mysql error is "Can't change size of file (Errcode: %d - %s)" 
func IsGlobalErrorCantChsize(err error) bool {
    result := Isa(err, EE_CANT_CHSIZE)
    return result
}

    
// IsGlobalErrorCantOpenStream check mysql error is "Can't open stream from handle (Errcode: %d - %s)" 
func IsGlobalErrorCantOpenStream(err error) bool {
    result := Isa(err, EE_CANT_OPEN_STREAM)
    return result
}

    
// IsGlobalErrorGetwd check mysql error is "Can't get working directory (Errcode: %d - %s)" 
func IsGlobalErrorGetwd(err error) bool {
    result := Isa(err, EE_GETWD)
    return result
}

    
// IsGlobalErrorSetwd check mysql error is "Can't change dir to '%s' (Errcode: %d - %s)" 
func IsGlobalErrorSetwd(err error) bool {
    result := Isa(err, EE_SETWD)
    return result
}

    
// IsGlobalErrorLinkWarning check mysql error is "Warning: '%s' had %d links" 
func IsGlobalErrorLinkWarning(err error) bool {
    result := Isa(err, EE_LINK_WARNING)
    return result
}

    
// IsGlobalErrorOpenWarning check mysql error is "Warning: %d files and %d streams is left open" 
func IsGlobalErrorOpenWarning(err error) bool {
    result := Isa(err, EE_OPEN_WARNING)
    return result
}

    
// IsGlobalErrorDiskFull check mysql error is "Disk is full writing '%s' (Errcode: %d - %s). Waiting forsomeone to free space..." 
func IsGlobalErrorDiskFull(err error) bool {
    result := Isa(err, EE_DISK_FULL)
    return result
}

    
// IsGlobalErrorCantMkdir check mysql error is "Can't create directory '%s' (Errcode: %d - %s)" 
func IsGlobalErrorCantMkdir(err error) bool {
    result := Isa(err, EE_CANT_MKDIR)
    return result
}

    
// IsGlobalErrorUnknownCharset check mysql error is "Character set '%s' is not a compiled character set and isnot specified in the '%s' file" 
func IsGlobalErrorUnknownCharset(err error) bool {
    result := Isa(err, EE_UNKNOWN_CHARSET)
    return result
}

    
// IsGlobalErrorOutOfFileresources check mysql error is "Out of resources when opening file '%s' (Errcode: %d -%s)" 
func IsGlobalErrorOutOfFileresources(err error) bool {
    result := Isa(err, EE_OUT_OF_FILERESOURCES)
    return result
}

    
// IsGlobalErrorCantReadlink check mysql error is "Can't read value for symlink '%s' (Error %d - %s)" 
func IsGlobalErrorCantReadlink(err error) bool {
    result := Isa(err, EE_CANT_READLINK)
    return result
}

    
// IsGlobalErrorCantSymlink check mysql error is "Can't create symlink '%s' pointing at '%s' (Error %d -%s)" 
func IsGlobalErrorCantSymlink(err error) bool {
    result := Isa(err, EE_CANT_SYMLINK)
    return result
}

    
// IsGlobalErrorRealpath check mysql error is "Error on realpath() on '%s' (Error %d - %s)" 
func IsGlobalErrorRealpath(err error) bool {
    result := Isa(err, EE_REALPATH)
    return result
}

    
// IsGlobalErrorSync check mysql error is "Can't sync file '%s' to disk (Errcode: %d - %s)" 
func IsGlobalErrorSync(err error) bool {
    result := Isa(err, EE_SYNC)
    return result
}

    
// IsGlobalErrorUnknownCollation check mysql error is "Collation '%s' is not a compiled collation and is notspecified in the '%s' file" 
func IsGlobalErrorUnknownCollation(err error) bool {
    result := Isa(err, EE_UNKNOWN_COLLATION)
    return result
}

    
// IsGlobalErrorFilenotfound check mysql error is "File '%s' not found (Errcode: %d - %s)" 
func IsGlobalErrorFilenotfound(err error) bool {
    result := Isa(err, EE_FILENOTFOUND)
    return result
}

    
// IsGlobalErrorFileNotClosed check mysql error is "File '%s' (fileno: %d) was not closed" 
func IsGlobalErrorFileNotClosed(err error) bool {
    result := Isa(err, EE_FILE_NOT_CLOSED)
    return result
}

    
// IsGlobalErrorChangeOwnership check mysql error is "Can't change ownership of the file '%s' (Errcode: %d -%s)" 
func IsGlobalErrorChangeOwnership(err error) bool {
    result := Isa(err, EE_CHANGE_OWNERSHIP)
    return result
}

    
// IsGlobalErrorChangePermissions check mysql error is "Can't change permissions of the file '%s' (Errcode: %d -%s)" 
func IsGlobalErrorChangePermissions(err error) bool {
    result := Isa(err, EE_CHANGE_PERMISSIONS)
    return result
}

    
// IsGlobalErrorCantSeek check mysql error is "Can't seek in file '%s' (Errcode: %d - %s)" 
func IsGlobalErrorCantSeek(err error) bool {
    result := Isa(err, EE_CANT_SEEK)
    return result
}




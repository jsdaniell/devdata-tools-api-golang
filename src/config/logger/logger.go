package logger
//
//import (
//	"flag"
//	"log"
//	"os"
//)
//
//var (
//	LogUser      *log.Logger
//	LogTestCase      *log.Logger
//)
//
//
//func init() {
//	// Log for User Operations
//	var logPathUser = "github.com/jsdaniell/devdata-tools-api-golang/config/logger/logs/user.log"
//
//	flag.Parse()
//	var fileUser, err1 = os.Create(logPathUser)
//	if err1 != nil {
//		panic(err1)
//	}
//
//	LogUser = log.New(fileUser, "", log.LstdFlags|log.Lshortfile)
//	LogUser.Println("LogFile : " + logPathUser)
//
//
//
//	// Log for TestCase Operations
//	var logPathTestCase = "github.com/jsdaniell/devdata-tools-api-golang/config/logger/logs/test_case.log"
//
//
//	var fileTestCase, err2 = os.Create(logPathTestCase)
//	if err2 != nil {
//		panic(err1)
//	}
//
//	LogTestCase = log.New(fileTestCase, "", log.LstdFlags|log.Lshortfile)
//	LogTestCase.Println("LogFile : " + logPathTestCase)
//
//}

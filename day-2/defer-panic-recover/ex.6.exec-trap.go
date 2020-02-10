package execTrapDefault

import (
	"runtime/debug"
	"gitbut.com/sameeroak11/logger"
)


type ExecTrapDefault struct {
	Exec    func()
	Trap    func(Exception)
	Default func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (etd ExecTrapDefault) Execute() {
	if etd.Default != nil {
		defer etd.Default()
	}
	if etd.Trap != nil {
		defer func() {
			if r := recover(); r != nil {
				logger.Log(PKGNAME, logger.ERROR, "Recovered from panic: %v", r)
				logger.Log(PKGNAME, logger.ERROR, "Dumping stack:\n==============")
				debug.PrintStack()
				etd.Trap(r)
			}
		}()
	}
	etd.Exec()
}


//

/*
func APIHandler(responseWriter http.ResponseWriter, pReq *http.Request) {
	execTrapDefault.ExecTrapDefault {
		Exec: func() {
			// your code here
		},

		Trap: func(e trycatch.Exception) {
			if e != nil {
				// your code here
			} else {
				// your code here
			}
		},

		Default: func() {
			// your code here
		},
	}.Execute()
}

*/

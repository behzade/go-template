package service

var Errors = struct {
	InitConfigError   error
	DBConnectionError error
}{}

func init() {
	err := initConfig()
	if err != nil {
		Errors.InitConfigError = err
		return
	}

	Errors.DBConnectionError = initDBConnection()
}

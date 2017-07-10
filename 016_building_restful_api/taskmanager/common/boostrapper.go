package common

func Startup() {
	initConfig()
	initKeys()
	createDbSession()
	AddIndexes()
}

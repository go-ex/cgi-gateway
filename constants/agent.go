package constants

var appAgent map[int]string

func init() {
	appAgent = make(map[int]string)
}

func AddAppAgent(app int, url string) {
	appAgent[app] = url
}

func GetAppAgent(app int) string {
	return appAgent[app]
}

func GetAll() map[int]string {
	return appAgent
}

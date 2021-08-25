package settings

type SettingsDB struct {
	Username string  `json:"username"`
	Password string `json:"password"`
	DatabaseName string `json:"database_name"`
}
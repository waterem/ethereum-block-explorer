package config

var config map[string]string

func init() {
	config = make(map[string]string)
	config["lockFile"] = "/var/tmp/explorer-cli.lock"
	config["lastBlockFile"] = "/var/tmp/explorer-cli.last-block"
	config["errorFile"] = "/var/tmp/explorer-cli.error"
}
func Get(key string) string {
	return config[key]
}

package templates

import "fmt"

// GenerateCFG ...
func GenerateCFG(cfgName string) (string, string) {
	filename := fmt.Sprintf("config/%s.toml", cfgName)
	body := `base_add=":8080"
db_user=""
db_password=""
db_host="localhost"
db_port="5432"
db_name=""
`
	return filename, body
}

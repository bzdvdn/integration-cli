package templates

// ResponseGO ...
func ResponseGO() (string, string) {
	filename := "server/response/response.go"
	body := `package api`
	return filename, body
}

// ValidationsGO ...
func ValidationsGO() (string, string) {
	filename := "server/response/validations.go"
	body := `package api`
	return filename, body
}

// ValidationResponseGO ...
func ValidationResponseGO() (string, string) {
	filename := "server/response/validation_response.go"
	body := `package api`
	return filename, body
}

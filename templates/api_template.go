package templates

// DocGO ...
func DocGO() (string, string) {
	filename := "api/doc.go"
	body := `package api

/*
The user module containing the user CRU operation.
model.go: definition of orm based data model
router.go: router binding and core logic
serializer.go: definition the schema of return data
validator.go: definition the validator of form data
*/
`
	return filename, body
}

// ModelGO ...
func ModelGO() (string, string) {
	filename := "api/model.go"
	body := `
package api
`
	return filename, body
}

// SerializerGO ...
func SerializerGO() (string, string) {
	filename := "api/serializer.go"
	body :=
		`
package api
`
	return filename, body
}

// ValidatorGO ....
func ValidatorGO() (string, string) {
	filename := "api/validator.go"
	body := `
package api
`
	return filename, body
}

// APIRouter ....
func APIRouter() (string, string) {
	filename := "api/router.go"
	body := `
package api

import "github.com/gin-gonic/gin"

// Health ...
func Health(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "working!",
		})
	})
}
`
	return filename, body
}

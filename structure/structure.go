package structure

import (
	"integration-cli/templates"
	"log"
	"os"
	"strings"
	"text/template"
)

func generateDirs(projectName string) {
	if _, err := os.Stat(projectName); err != nil {
		os.Mkdir(projectName, 0700)
	}
	dirs := []string{
		"api",
		"server",
		"middleware",
		"database",
		"config",
		"utils",
	}
	for _, dir := range dirs {
		if _, err := os.Stat(projectName + dir); err != nil {
			os.Mkdir(projectName+dir, 0700)
		}
	}
	respDir := "server/response"
	if _, err := os.Stat(projectName + respDir); err != nil {
		os.Mkdir(projectName+respDir, 0700)
	}
}

func generateFile(projectName string, filename string, body string) {

	body = strings.Replace(body, "'", "`", -1)
	t, err := template.New("gen").Parse(body)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, projectName)
	if err != nil {
		panic(err)
	}
}

//GenerateStucts ...
func GenerateStucts(projectName string) {
	generateDirs(projectName + "/")
	for _, function := range []func() (string, string){
		templates.ServerGO,
		templates.ConfigGO,
		templates.Gomod,
		templates.SerializerGO,
		templates.ModelGO,
		templates.ResponseGO,
		templates.ValidationResponseGO,
		templates.ValidationsGO,
		templates.ValidatorGO,
		templates.DatabaseGO,
		templates.APIRouter,
		templates.MainGO,
		templates.RouterGO,
		templates.MakefileGO,
	} {
		f, b := function()
		generateFile(projectName, projectName+"/"+f, b)
	}
	for _, cfgName := range []string{
		"dev-config",
		"prod-config",
		"stage-config",
	} {
		f, b := templates.GenerateCFG(cfgName)
		generateFile(projectName, projectName+"/"+f, b)
	}
}

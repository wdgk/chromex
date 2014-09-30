package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"text/template"
)

var manifestJsonTemplate = template.Must(ParseAsset("manifest", "templates/manifest.tmpl"))
var jsTemplate = template.Must(ParseAsset("js", "templates/mainJS.tmpl"))
var htmlTemplate = template.Must(ParseAsset("html", "templates/mainHTML.tmpl"))

func ParseAsset(name string, path string) (*template.Template, error) {
	src, err := Asset(path)
	if err != nil {
		return nil, err
	}

	return template.New(name).Parse(string(src))
}

var manifestJson = Source{
	Name:     "manifest.json",
	Template: *manifestJsonTemplate,
}

var Commands = []cli.Command{
	commandInit,
}

var commandInit = cli.Command{
	Name:        "init",
	Usage:       "Create template files.",
	Description: ``,
	Action:      doInit,
}

type Application struct {
	Name, Author, Email string
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doInit(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("Usage:chromex init [Project Name]")
	}
	name := c.Args()[0]
	
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatal(err)
	}

	application := Application{
		Name:   name,
		Author: GitConfig("user.name"),
		Email:  GitConfig("user.email"),
	}

	err = manifestJson.generate(name, application)
	assert(err)
	mainJs := Source{
		Name:     name + ".js",
		Template: *jsTemplate,
	}
	mainHtml := Source{
		Name:     name + ".html",
		Template: *htmlTemplate,
	}

	err = mainJs.generate(name, application)
	assert(err)
	err = mainHtml.generate(name, application)
	assert(err)

}

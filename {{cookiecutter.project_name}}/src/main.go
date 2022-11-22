package main

import (
	"github.com/{{ cookiecutter.github_org }}/{{ cookiecutter.project_slug }}/cmd"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	cmd.Execute(version, commit)
}

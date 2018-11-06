package main

import (
	"context"
	"flag"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"os"
	"text/template"
)

var (
	repository   = flag.String("repo", "", "Name of repository to read")
	organization = flag.String("org", "", "Organization to use")
)

func main() {

	// Parse flag arguments
	flag.Parse()

	// Check if required variables are set
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Error: Environment variable GITHUB_TOKEN not set")
	}
	if *repository == "" {
		log.Fatal("Error: Name of repository not given")
	}
	if *organization == "" {
		log.Fatal("Error: Name of organization not given")
	}

	// Initialize oauth2 token and http client
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Initialize GitHub client
	client := github.NewClient(httpClient)

	// Get the desired repository
	r, _, _ := client.Repositories.Get(context.Background(), *organization, *repository)

	// Parse the template file
	tmpl, err := template.ParseFiles("templates/github_repository.tmpl")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	// Print the rendered template to STDOUT
	tmpl_err := tmpl.Execute(os.Stdout, *r)
	if tmpl_err != nil {
		log.Fatal("Execute: ", tmpl_err)
		return
	}
}

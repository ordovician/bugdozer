// Copyright 2013 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Reads JSON data on STDIN and shows it in a space separated table.
// The JSON is in the format used by JIRA for issues
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"bugdozer/jira"
	"flag"
)

var usageString = `bug ls is a tool for listing projects and issues.

Usage:

	bug-ls item 

The items are:
	project		list projects
	issue		list issues
`

type Issues struct {
	Issues []jira.Issue
}

func lsIssues(input []byte)  {
	var issues Issues
	
	err := json.Unmarshal(input, &issues)
	
	if err != nil {
		fmt.Errorf("Error parsing JSON input: %v", err)
	}
	
	for _, issue := range issues.Issues {
		fmt.Printf("%v %v %v\n", issue.Id, issue.Key, issue.Fields.Summary)
	}
}

func lsProjects(input []byte)  {
	var projects []jira.Project
	
	err := json.Unmarshal(input, &projects)
	
	if err != nil {
		fmt.Errorf("Error parsing JSON input: %v", err)
	}
	
	for _, proj := range projects {
		fmt.Printf("%v %v %v %v\n", proj.Id, proj.Key, proj.Name, proj.Description)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	
	input, _ := ioutil.ReadAll(os.Stdin)

	switch {
	case args[0] == "issue":
		lsIssues(input)
	case args[0] == "project":
		lsProjects(input)
	default:
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, usageString)
	os.Exit(2)
}
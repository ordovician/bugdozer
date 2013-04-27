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
//	"net/url"
)

type URL string
type ID string

type Person struct {
	DisplayName, EmailAddress, Name string
	Self URL	
}

type Comment struct {
	Author Person
	Body string
	Id ID
	Self URL
	UpdateAuthor Person
}

type CommentsSection struct {
	Comments []Comment
	MaxResults int
	StartAt int
	Total int
}

// Person assigned to a issue
type Assignee struct {
	Person
}

// A task, bug etc
type Issue struct {
	Fields struct {
		Description, Summary string
		Labels []string
		Comment CommentsSection
		Assignee Assignee
	}
	Id ID		// unique number identfier
	Key string		// a unique string identifier
	Self URL
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	var issue Issue
	
	err := json.Unmarshal(input, &issue)
	
	if err != nil {
		fmt.Errorf("Error parsing JSON input: %v", err)
	}

	fmt.Printf("Summary: %v\n\n", issue.Fields.Summary)
	fmt.Printf("Description: %v\n\n", issue.Fields.Description)
	fmt.Printf("Comments:\n")
	for _, comment := range issue.Fields.Comment.Comments {
		fmt.Printf("  %v: %v\n", comment.Author.DisplayName, comment.Body)
	}
}

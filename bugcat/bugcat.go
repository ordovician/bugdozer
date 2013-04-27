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
//	"net/url"
)



func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	var issue jira.Issue
	
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

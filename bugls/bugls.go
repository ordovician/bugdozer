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
	"net/url"
)

type Issues struct {
	Issues []struct {
		Fields struct {
			Description, Summary string
		}
		Id string
		Key string
		Self url.URL
	}
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	var issues Issues
	
	json.Unmarshal(input, &issues)
	
	for _, issue := range issues.Issues {
		fmt.Printf("%v %v %v\n", issue.Id, issue.Key, issue.Fields.Summary)
	}
	
}

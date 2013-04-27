// Copyright 2013 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// Common data types used when accessing JIRA REST API.
// JIRA is a bug and issue tracking system
package jira

type URL string
type ID string

// A user in JIRA. Person reporting bug, fixing it or commenting on it
type Person struct {
	DisplayName, EmailAddress, Name string
	Self URL	
}

// A comment to an issue
type Comment struct {
	Author Person
	Body string
	Id ID
	Self URL
	UpdateAuthor Person
}

// Contains all comments to an issue
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

// Usually used to track a software project. The tasks needed to be done etc.
type Project struct {
	Id ID
	Key string
	Name string
	Description string
	Self URL
}
package main

// Session : User session info
type Session struct {
	username   string
	backupName string
}

var sessions = make(map[string]Session)

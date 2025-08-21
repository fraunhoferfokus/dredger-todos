package usecases

import (
	"maps"
	"slices"
)

// Set of Sessions
type sessionSet map[string]struct{}

// Adds a session to the set
func (s sessionSet) add(session string) {
	s[session] = struct{}{}
}

// Removes a session from the set
func (s sessionSet) remove(session string) {
	delete(s, session)
}

// Returns a boolean value describing if the session exists in the set
func (s sessionSet) has(session string) bool {
	_, ok := s[session]
	return ok
}

// Returns a boolean value describing if the session exists in the set
func (s sessionSet) empty() bool {
	return len(s) == 0
}

// Sessions for each label
type LabelsMap map[string]sessionSet

var CurrentLabels LabelsMap = LabelsMap{}

// Returns a boolean value describing if the label exists in the set
func (s LabelsMap) has(label string) bool {
	_, ok := s[label]
	return ok
}

// Adds a label to the set
func (s LabelsMap) add(label string) {
	s[label] = sessionSet{}
}

// Removes a label from the set
func (s LabelsMap) remove(label string) {
	delete(s, label)
}

// Assign session to label
func (l LabelsMap) Set(label string, session string) {
	// remove session from all labels
	for x := range l {
		if l[x].has(session) {
			l[x].remove(session)
			if l[x].empty() {
				l.remove(x)
			}
		}
	}
	// set session for label
	if label != "" {
		if !l.has(label) {
			l.add(label)
		}
		l[label].add(session)
	}
}

// Find the label of a session
func (l LabelsMap) GetLabel(session string) (string, bool) {
	for x := range l {
		if l[x].has(session) {
			return x, true
		}
	}
	return "", false
}

// Get all session for a label
func (l LabelsMap) GetSessions(label string) []string {
	if l.has(label) {
		return slices.Sorted(maps.Keys(l[label]))
	}
	return []string{}
}

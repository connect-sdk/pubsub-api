package pubsubv1

// Key returns the key path.
func (x ParsedProjectName) Key() string {
	return x.ProjectID
}

// String returns the string representation.
func (x ParsedProjectName) String() string {
	return x.Name()
}

// Key returns the key path.
func (x ParsedTopicName) Key() string {
	return x.TopicID
}

// Parent returns the parent name.
func (x ParsedTopicName) Parent() ParsedProjectName {
	return ParsedProjectName{
		ProjectID: x.ProjectID,
	}
}

// String returns the string representation.
func (x ParsedTopicName) String() string {
	return x.Name()
}

package zip

// Extract opens the supplied zip archive file and reads the contents to the given destination.
// If a map of file extension is supplied only those that match will be extracted.
// The returned slice holds the paths to all the extracted files. If there is an error returned the
// slice will be nil.
func Extract(archive, dest string, formats map[string]struct{}) ([]string, error) {
	return nil, nil
}

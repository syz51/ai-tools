package handlers

import (
	"net/url"
	"path"
)

// ParseAndJoinURL parses the base URL and joins it with additional path segments.
func JoinUrl(base string, parts ...string) (string, error) {
	parsedURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	// Join the additional parts with the base URL's path
	parsedURL.Path = path.Join(parsedURL.Path, path.Join(parts...))

	return parsedURL.String(), nil
}

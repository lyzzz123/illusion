package filter

import "regexp"

type filterWrapper struct {
	innerFilter Filter

	pathRegex *regexp.Regexp
}

func (object *filterWrapper) Match(path string) bool {
	return object.pathRegex.MatchString(path)
}

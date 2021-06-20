package utils

import "regexp"

const replaceForStar = "[a-zA-Z\\d]+"
const replaceForDoubleStar = "[a-zA-Z\\d/]+"
const star = "\\*"
const doubleStar = "\\*\\*"

const pathValue = "{([a-zA-Z\\d]+)}"
const replaceForPathValue = "([a-zA-Z\\d]+)"

var starPattern = regexp.MustCompile(star)
var doubleStarPattern = regexp.MustCompile(doubleStar)

var handlerPattern = regexp.MustCompile(pathValue)

func ParseFilterUrlPattern(urlPattern string) *regexp.Regexp {
	urlPattern = doubleStarPattern.ReplaceAllString(urlPattern, replaceForDoubleStar)
	urlPattern = starPattern.ReplaceAllString(urlPattern, replaceForStar)
	return regexp.MustCompile("^" + urlPattern + "$")
}

func ParseHandlerUrlPattern(urlPattern string) *regexp.Regexp {
	urlPattern = handlerPattern.ReplaceAllString(urlPattern, replaceForStar)
	return regexp.MustCompile("^" + urlPattern + "$")
}

func ParseHandlerUrlPatternForPathValue(urlPattern string) *regexp.Regexp {
	urlPattern = handlerPattern.ReplaceAllString(urlPattern, replaceForPathValue)
	return regexp.MustCompile("^" + urlPattern + "$")
}
func ParseHandlerUrlPatternForPath(urlPattern string) *regexp.Regexp {
	urlPattern = handlerPattern.ReplaceAllString(urlPattern, pathValue)
	return regexp.MustCompile("^" + urlPattern + "$")
}

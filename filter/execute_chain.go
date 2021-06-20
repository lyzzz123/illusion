package filter

import (
	"github.com/lyzzz123/illusion/utils"
	"net/http"
	"sort"
)

var FilterChain = make([]*filterWrapper, 0)

func RegisterFilter(filter Filter) {
	if filter == nil {
		panic("RegisterFilter filter must not be nil")
	}

	if filter.GetPathPattern() == "" {
		panic("RegisterFilter PathPattern must not be blank")
	}

	filterWrapper := &filterWrapper{}
	filterWrapper.innerFilter = filter
	filterWrapper.pathRegex = utils.ParseFilterUrlPattern(filter.GetPathPattern())
	FilterChain = append(FilterChain, filterWrapper)
	sort.SliceStable(FilterChain, func(i, j int) bool {
		return FilterChain[i].innerFilter.GetPriority() > FilterChain[j].innerFilter.GetPriority()
	})
}

func ExecutePreHandle(writer http.ResponseWriter, request *http.Request) error {

	url := request.URL.Path
	for i := 0; i < len(FilterChain); i++ {
		if FilterChain[i].Match(url) {
			err := FilterChain[i].innerFilter.PreHandle(writer, request)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ExecutePostHandle(writer http.ResponseWriter, request *http.Request) error {
	url := request.URL.Path
	for i := 0; i < len(FilterChain); i++ {
		if FilterChain[i].Match(url) {
			err := FilterChain[i].innerFilter.PostHandle(writer)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

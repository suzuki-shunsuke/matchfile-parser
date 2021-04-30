package matchfile

import (
	"fmt"
)

type Parser struct {
	matcherFuncs map[string]NewMatcherFunc
}

func NewParser() *Parser {
	return &Parser{
		matcherFuncs: map[string]NewMatcherFunc{
			"dir":    newDirMatcher,
			"regexp": newRegexpMatcher,
			"glob":   newGlobMatcher,
			"equal":  newEqualMatcher,
		},
	}
}

func (parser *Parser) AddMatcherFunc(kind string, fn NewMatcherFunc) {
	if fn == nil {
		delete(parser.matcherFuncs, kind)
		return
	}
	parser.matcherFuncs[kind] = fn
}

func (parser *Parser) Match(checkedFiles []string, conditions []Condition) (bool, error) {
	for _, checkedFile := range checkedFiles {
		f, err := parser.MatchFile(checkedFile, conditions)
		if err != nil {
			return false, err
		}
		if f {
			return true, nil
		}
	}
	return false, nil
}

func (parser *Parser) MatchFile(checkedFile string, conditions []Condition) (bool, error) {
	matched := false
	for _, condition := range conditions {
		if matched && !condition.Exclude {
			continue
		}
		if !matched && condition.Exclude {
			continue
		}
		b, err := condition.Match(checkedFile)
		if err != nil {
			return false, fmt.Errorf("condition matching error: %w", err)
		}
		if b {
			matched = !condition.Exclude
			continue
		}
	}
	return matched, nil
}

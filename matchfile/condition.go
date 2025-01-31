package matchfile

import "fmt"

type Condition struct {
	Exclude bool
	matcher Matcher
}

func (cond *Condition) Match(p string) (bool, error) {
	s, err := cond.matcher.Match(p)
	if err != nil {
		return false, fmt.Errorf("check if the string '%s' matches: %w", p, err)
	}
	return s, nil
}

func (parser *Parser) ParseConditions(conditionLines []string) ([]Condition, error) {
	conditions := make([]Condition, 0, len(conditionLines))
	for _, conditionLine := range conditionLines {
		matchParam := parseLine(conditionLine)
		if matchParam.Comment {
			continue
		}
		matchers := make([]Matcher, len(matchParam.Kinds))
		for j, kind := range matchParam.Kinds {
			matcher, err := parser.NewMatcher(matchParam.Path, kind)
			if err != nil {
				return nil, fmt.Errorf("initialize a matcher: %w", err)
			}
			matchers[j] = matcher
		}
		conditions = append(conditions, Condition{
			Exclude: matchParam.Exclude,
			matcher: &combinedMatcher{matchers: matchers},
		})
	}
	return conditions, nil
}

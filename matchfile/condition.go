package matchfile

import "fmt"

type Condition struct {
	Exclude bool
	matcher Matcher
}

func (cond *Condition) Match(p string) (bool, error) {
	return cond.matcher.Match(p)
}

func (parser *Parser) GetConditions(conditionLines []string) ([]Condition, error) {
	conditions := make([]Condition, 0, len(conditionLines))
	for _, conditionLine := range conditionLines {
		matchParam := parseLine(conditionLine)
		if matchParam.Comment {
			continue
		}
		matchers := make([]Matcher, len(matchParam.Kinds))
		for j, kind := range matchParam.Kinds {
			matcher, err := NewMatcher(matchParam.Path, kind)
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

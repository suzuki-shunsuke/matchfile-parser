package matchfile

func (parser *Parser) List(checkedFiles []string, conditions []Condition) ([]string, error) {
	matchedFiles := make([]string, 0, len(checkedFiles))
	for _, checkedFile := range checkedFiles {
		f, err := parser.MatchFile(checkedFile, conditions)
		if err != nil {
			return nil, err
		}
		if f {
			matchedFiles = append(matchedFiles, checkedFile)
			continue
		}
	}
	return matchedFiles, nil
}

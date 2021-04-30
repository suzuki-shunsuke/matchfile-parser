# matchfile-parser

[![Build Status](https://github.com/suzuki-shunsuke/matchfile-parser/workflows/test/badge.svg)](https://github.com/suzuki-shunsuke/matchfile-parser/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/suzuki-shunsuke/matchfile-parser)](https://goreportcard.com/report/github.com/suzuki-shunsuke/matchfile-parser)
[![GitHub last commit](https://img.shields.io/github/last-commit/suzuki-shunsuke/matchfile-parser.svg)](https://github.com/suzuki-shunsuke/matchfile-parser)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/suzuki-shunsuke/matchfile-parser/main/LICENSE)

Parser of [matchfile](https://github.com/suzuki-shunsuke/matchfile) file format

## Format

The format is inspired by [gitignore](https://git-scm.com/docs/gitignore).

```
[#][!][<kind>,...] <path>
...
```

When the multiple kinds are specified, the condition matches when either of them matches.

The line starts with "#" is ignored as code comment.
Note that the comment in the middle of the line isn't supported.

`[<kind>,...]` is optional, and the default value is `equal,dir,glob`.

### kind

* equal: check the equality
* dir: [strings.HasPrefix](https://golang.org/pkg/strings/#HasPrefix)
* regexp: [regexp.MatchString](https://golang.org/pkg/regexp/#Regexp.MatchString)
* glob: [filepath.Match](https://golang.org/pkg/path/filepath/#Match)

## LICENSE

[MIT](LICENSE)

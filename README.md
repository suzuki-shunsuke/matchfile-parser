# matchfile-parser

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

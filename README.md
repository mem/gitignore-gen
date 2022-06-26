gitignore-gen
=============

gitignore-gen generates .gitignore files from snippets, like those
provided by [Github's gitignore repository](https://github.com/github/gitignore).

The configuration file lists all the inputs, in the order they should be
used. They can be local files (relative to the directory specified using
`-local-data-dir`) or URLs.

```yaml
inputs:
        - Global/Vim.gitignore
        - Global/Backup.gitignore
        - Global/Linux.gitignore
        - https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore
```

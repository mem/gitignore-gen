gitignore-gen
=============

gitignore-gen generates .gitignore files from snippets, like those
provided by [Github's gitignore repository](https://github.com/github/gitignore).

The configuration file lists all the inputs, in the order they should be
used:

```yaml
inputs:
        - Global/Vim.gitignore
        - Global/Backup.gitignore
        - Global/Linux.gitignore
        - Go.gitignore
```

they refer to the root of the data directory specified using
`-local-data-dir`.

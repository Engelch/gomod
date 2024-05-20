# CHANGELOG

LIFO documentation style

## V2.3

- new IsDirectory(filename string) bool
  
## V2.2

- ExecCmd is DEPRECATED, to be replaced by ExecNoOutputCmd
- New functions
    - ExecNoOutputCmd(cmdAndArgs string) error
    - ExecOutputCmd(cmdAndArgs string) ([]byte, error)
    - IsExistingFile(filename string) bool
    - IsExistingPlainFile(filename string) bool
## V2.1

- ExecutableReachableByPath(cmd...) error
  - added to package
  - test routine implemented and passing

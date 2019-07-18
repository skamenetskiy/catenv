# catenv
[![Build Status](https://travis-ci.com/skamenetskiy/catenv.svg?branch=master)](https://travis-ci.com/skamenetskiy/catenv)

### Description
A simple tool to append environment variables to a file contents. It supports both `$` and `${}` syntax.

### Usage
```
Usage: catenv <filename>
Usage: cat <filename> | catenv -in
```

### Example
```yaml
test_file:
  go_root: ${GOROOT}
  go_root_2: $GOROOT
```
Will output something like:
```yaml
test_file:
  go_root: C:\Go
  go_root_2: C:\Go
```

# tos
[Mackerel-Agent](https://mackerel.io/ja/docs/entry/custom-checks)のaction指定時に、条件分岐の記述を簡潔に記述出来ます。

## Usage
```
$ tos -h
Usage of tos:
  -c string
        CRITICAL(Short)
  -critical string
        CRITICAL
  -o string
        OK(Short)
  -ok string
        OK
  -u string
        Unknown(Short)
  -unknown string
        Unknown
  -version
        Print version information and quit.
  -w string
        WARNING(Short)
  -warning string
        WARNING
```

```bash
[plugin.checks.log]
command = "check-log -f /path/to/file -p PATTERN"
# action = { command = "bash -c '[ \"$MACKEREL_STATUS\" != \"OK\" ]' && ruby /path/to/something.rb", user = "someone" }
action = { command = "tos -no 'ruby /path/to/something.rb'", user = "someone" }
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/pyama86/tos
```

## Contribution

1. Fork ([https://github.com/pyama86/tos/fork](https://github.com/pyama86/tos/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[pyama86](https://github.com/pyama86)

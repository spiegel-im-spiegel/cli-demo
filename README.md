# [cli-demo] - Demonstration for Command Line Interface

## Install

```
$ go get -v github.com/spiegel-im-spiegel/cli-demo
```

## Usage

```
$ cli-demo -h
Long comment

Usage:
  cli-demo [flags]
  cli-demo [command]

Available Commands:
  help        Help about any command
  show        Short comment for show sub-command

Flags:
  -h, --help   help for cli-demo

Use "cli-demo [command] --help" for more information about a command.

$ cli-demo show -h
Long comment for show sub-command

Usage:
  cli-demo show [flags]

Flags:
  -b, --boolean         boolean option
  -h, --help            help for show
  -i, --integer int     integer option
  -s, --string string   string option

$ cli-demo show -i 123 -s 日本語 -b
Integer option value: 123
 String option value: 日本語
Boolean option value: true
```

## Dependencies

```
dep status -dot | dot -Tpng -o dependency.png
```

[![Dependencies](dependency.png)](dependency.png)

[cli-demo]: https://github.com/spiegel-im-spiegel/cli-demo "spiegel-im-spiegel/cli-demo: Demonstration for Command Line Interface"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"

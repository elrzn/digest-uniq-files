# digest-uniq-files

[![Go Report Card](https://goreportcard.com/badge/github.com/elrzn/digest-uniq-files)](https://goreportcard.com/report/github.com/elrzn/digest-uniq-files)
[![License](https://img.shields.io/github/license/elrzn/digest-uniq-files.svg)](https://github.com/elrzn/digest-uniq-files/blob/master/LICENSE)

## Synopsis

    $ go get github.com/elrzn/digest-uniq-files
    $ ls | grep -E '(jpg|png)$' | wc -l
        37
    $ digest-uniq-files -ext 'jpg,png'
        37  35

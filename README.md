# goblin

goblin is an embeddable scripting engine written in Go.

## Notes

To regenerate `parser.go`, you need to have `goyacc` installed:

`go get -u golang.org/x/tools/cmd/goyacc/...`

## To do

- Support more slice notations:
  - s[:3]
  - s[3:]

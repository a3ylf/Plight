package flags

import "flag"

var (
    Reset bool
    Dev bool
)

func ParseFlags() {
    flag.BoolVar(&Dev,"dev", false,"Use file debug.json")
    flag.BoolVar(&Reset,"reset", false,"reset file")
    flag.Parse()
}
func ParseArgs() []string {
    return flag.Args()
}

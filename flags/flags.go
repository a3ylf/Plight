package flags

import "flag"

var (
    Reset bool
    Dev bool
    Raw bool
)

func ParseFlags() {
    flag.BoolVar(&Dev,"dev", false,"Use file debug.json")
    flag.BoolVar(&Reset,"reset", false,"reset file")
    flag.BoolVar(&Raw,"raw", false,"receive raw data")
    flag.Parse()
}
func ParseArgs() []string {
    return flag.Args()
}

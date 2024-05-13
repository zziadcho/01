package variables

import "regexp"

var (
    // targeting flags with regexp
    HexFlag = regexp.MustCompile(`\(hex\)`)
    BinFlag = regexp.MustCompile(`\(bin\)`)
    UpFlag = regexp.MustCompile(`\(up\)`)
    LowFlag = regexp.MustCompile(`\(low\)`)
    CapFlag = regexp.MustCompile(`\(cap\)`)
    // multi flags
    UpFlagMulti  = regexp.MustCompile(`\(up,`)
    LowFlagMulti = regexp.MustCompile(`\(low,`)
    CapFlagMulti = regexp.MustCompile(`\(cap,`)
)

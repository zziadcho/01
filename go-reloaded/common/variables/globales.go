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
    UpFlagMulti  = regexp.MustCompile(`\(up,\s+(\d+)\)`)
    LowFlagMulti = regexp.MustCompile(`\(low,\s+(\d+)\)`)
    CapFlagMulti = regexp.MustCompile(`\(cap,\s+(\d+)\)`)
)

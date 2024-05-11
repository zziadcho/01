package variables

import "regexp"

var (
    // targeting flags with regexp
    HexFlag = regexp.MustCompile(`^\\(hex\\)$`)
    BinFlag = regexp.MustCompile(`^\\(bin\\)$`)
    UpFlag = regexp.MustCompile(`^\\(up\\)$`)
    LowFlag = regexp.MustCompile(`^\\(low\\)$`)
    CapFlag = regexp.MustCompile(`^\\(cap\\)$`)
    // multi flags
    HexFlagMulti = regexp.MustCompile(`^\(hex,\s+[0-9]+\)$`)
    BinFlagMulti = regexp.MustCompile(`^\(bin,\s+[0-9]+\)$`)
    UpFlagMulti  = regexp.MustCompile(`^\(up,\s+[0-9]+\)$`)
    LowFlagMulti = regexp.MustCompile(`^\(low,\s+[0-9]+\)$`)
    CapFlagMulti = regexp.MustCompile(`^\(cap,\s+[0-9]+\)$`)
)

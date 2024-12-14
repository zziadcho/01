package functions

import (
	"fmt"
	"os"
)

func FormatPermissions(mode os.FileMode) string {
    typeChar := '-'
    switch {
    case mode&os.ModeSymlink != 0:
        typeChar = 'l'
    case mode&os.ModeDir != 0:
        typeChar = 'd'
    case mode&os.ModeCharDevice != 0:
        typeChar = 'c'
    case mode&os.ModeDevice != 0:
        typeChar = 'b'
    case mode&os.ModeNamedPipe != 0:
        typeChar = 'p'
    case mode&os.ModeSocket != 0:
        typeChar = 's'
    }

    ownerR := '-'
    if mode&0400 != 0 {
        ownerR = 'r'
    }
    ownerW := '-'
    if mode&0200 != 0 {
        ownerW = 'w'
    }
    ownerX := '-'
    if mode&0100 != 0 {
        ownerX = 'x'
    }

    groupR := '-'
    if mode&0040 != 0 {
        groupR = 'r'
    }
    groupW := '-'
    if mode&0020 != 0 {
        groupW = 'w'
    }
    groupX := '-'
    if mode&0010 != 0 {
        groupX = 'x'
    }

    othersR := '-'
    if mode&0004 != 0 {
        othersR = 'r'
    }
    othersW := '-'
    if mode&0002 != 0 {
        othersW = 'w'
    }
    othersX := '-'
    if mode&0001 != 0 {
        othersX = 'x'
    }

    return fmt.Sprintf("%c%c%c%c%c%c%c%c%c%c", 
        typeChar, 
        ownerR, ownerW, ownerX, 
        groupR, groupW, groupX, 
        othersR, othersW, othersX)
}
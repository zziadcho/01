package functions

func Major(dev uint64) uint64 {
    return (dev >> 8) & 0xfff
}

func Minor(dev uint64) uint64 {
    return (dev & 0xff) | ((dev >> 12) & 0xfff00)
}
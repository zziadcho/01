package common

import "regexp"

	// targeting flags with regexp
	var (
		//uni flags
		hexFlag = regexp.MustCompile(`^\\(hex\\)$`)
		binFlag = regexp.MustCompile(`^\\(bin\\)$`)
		upFlag = regexp.MustCompile(`^\\(up\\)$`)
		lowFlag = regexp.MustCompile(`^\\(low\\)$`)
		capFlag = regexp.MustCompile(`^\\(cap\\)$`)
		// multi flags
		hexFlagMulti = regexp.MustCompile(`^\(hex,\s+[0-9]+\)$`)
		binFlagMulti = regexp.MustCompile(`^\(bin,\s+[0-9]+\)$`)
		upFlagMulti  = regexp.MustCompile(`^\(up,\s+[0-9]+\)$`)
		lowFlagMulti = regexp.MustCompile(`^\(low,\s+[0-9]+\)$`)
		capFlagMulti = regexp.MustCompile(`^\(cap,\s+[0-9]+\)$`)
	)
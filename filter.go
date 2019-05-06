package main

import (
	"sort"
)

func filter(l *logicals) string {
	fls := iterateAndFilter(l.LogicalServers, checkCountry)
	fls = iterateAndFilter(fls, checkTier)
	sortScoreAsc(fls)
	for _, ls := range fls {
		debug(ls.Name, "\t", ls.Tier, "\t", ls.Load, "\t", ls.Score)
	}

	if len(fls) > 0 {
		return fls[0].Name
	}
	return ""
}

func iterateAndFilter(lss []logicalServer, check func(logicalServer) bool) []logicalServer {
	ret := []logicalServer{}
	for _, ls := range lss {
		if check(ls) {
			ret = append(ret, ls)
		}
	}
	return ret
}

func checkCountry(ls logicalServer) bool {
	return ls.EntryCountry == "NL"
}

func checkTier(ls logicalServer) bool {
	return ls.Tier == 2
}

func sortScoreAsc(lss []logicalServer) {
	sort.Slice(lss, func(i, j int) bool { return lss[i].Score < lss[j].Score })
}

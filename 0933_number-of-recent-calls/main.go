package main

type RecentCounter struct {
	ts []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.ts = append(rc.ts, t)
	for len(rc.ts) > 0 && rc.ts[0] < t-3000 {
		rc.ts = rc.ts[1:]
	}
	return len(rc.ts)
}

package example1604

import "sort"

type worker struct {
	periods []int
}

func alertNames(keyName, keyTime []string) []string {
	var alert []string
	workers := make(map[string]*worker)
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		if w, ok := workers[name]; ok {
			w.periods = append(w.periods, hour*60+minute)
		} else {
			wo := &worker{}
			wo.periods = append(wo.periods, hour*60+minute)
			workers[name] = wo
		}

	}
	for name, w := range workers {
		sort.Ints(w.periods)
		for i, t := range w.periods[2:] {
			if t-w.periods[i] <= 60 {
				alert = append(alert, name)
				break
			}
		}
	}
	sort.Strings(alert)
	return alert
}

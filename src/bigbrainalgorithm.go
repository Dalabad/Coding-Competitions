package src

func (d *Dataset) Solve() {
	for timeStamp := 0; timeStamp < d.Time; timeStamp++ {
		for _, intersection := range d.Intersections {
			if CheckCycles(intersection.Schedule.Sched, intersection.Schedule.Streets) {
				continue
			}
			streets := intersection.Schedule.Streets
			BubbleSort(streets)
			intersection.Schedule.Sched = append(intersection.Schedule.Sched, streets[0].Name)
		}
	}
	for _, intersection := range d.Intersections {
		intersection.setSched(d)
	}
}

func BubbleSort(data []*Street) {
	for i := 0; i < len(data); i++ {
		for _, car := range data[i].Cars {
			for _, street := range car.Path {
				data[i].SumOfPaths += street.Length
			}
		}
	}
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j].SumOfPaths > data[j-1].SumOfPaths {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

func CheckCycles(sched []string, streets []*Street) bool {
	if len(sched) == 0 {
		return false
	}
	last := sched[0]
	step := 0
	for i := 1; i < len(sched); i++ {
		if sched[i] != last {
			step++
			if step == len(streets)-1 {
				return true
			}
		}
		last = sched[i]
	}
	return false
}

func (in *Intersection) setSched(d *Dataset) {
	last := in.Schedule.Sched[0]
	step := 0
	in.Schedule.Streets = make([]*Street, len(in.Schedule.Streets))
	in.Schedule.Duration = make([]int, len(in.Schedule.Streets))
	//insert first street already
	for _, street := range d.Streets {
		if street.Name == in.Schedule.Sched[0] {
			in.Schedule.Streets[0] = &street
			in.Schedule.Duration[0]++

		}
	}

	for i := 1; i < len(in.Schedule.Sched); i++ {
		if in.Schedule.Sched[i] != last {
			for _, street := range d.Streets {
				if street.Name == in.Schedule.Sched[i] {
					in.Schedule.Streets[i] = &street
				}
			}
		}
		step++
		if step == len(in.Schedule.Streets)-1 {
			return
		}
	}
	in.Schedule.Duration[step]++
}

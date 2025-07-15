//go:build !solution

package hogwarts

func GetCourseList(prereqs map[string][]string) []string {
	num := make(map[string]bool)
	result := []string{}

	for key, prereq := range prereqs {
		num[key] = true
		for _, j := range prereq {
			num[j] = true

		}
	}

	tempmap := make(map[string]bool)

	for len(result) != len(num) {
		progress := false
	outer:
		for i := range num {
			if tempmap[i] {
				continue
			}
			for _, j := range prereqs[i] {
				if !tempmap[j] {
					continue outer
				}
			}

			result = append(result, i)
			tempmap[i] = true
			progress = true

		}
		if !progress {
			panic("Ошибка")
		}

	}

	return result
}

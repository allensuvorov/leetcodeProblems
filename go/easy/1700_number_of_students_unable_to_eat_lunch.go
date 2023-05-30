// my bulky solution
func myCountStudents(students []int, sandwiches []int) int {
    i, j := 0, 0
    lineStart := j
    for i < len(sandwiches) {
        if students[j] == sandwiches[i] {
            i++
            if j == len(students) - 1 {
                students = students[:j]
                j = 0
            } else {
                students = append(students[:j], students[j+1:]...)
            }
            lineStart = j
        } else {
            if j == len(students) - 1 {
                j = 0
            } else {
                j++
            } 
            if j == lineStart {
                return len(sandwiches) - i
            }
        }
    }
    return 0
}

// someone elses cool short solution 
func countStudents(students []int, sandwiches []int) int {
	missed := 0

	for missed != len(students) {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			missed = 0
		} else {
			students = append(students[1:], students[0])
			missed++
		}
	}

	return missed
}

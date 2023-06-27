package lib

func SemesterNumber(semester string) int {
	if semester == "even" {
		return 2
	} else {
		return 1
	}
}

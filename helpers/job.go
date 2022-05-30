package helpers

func IsMeetRequirement(age int, gender string) bool {
	if gender == "Pria" && (age > 17 && age < 60) {
		return true
	} else if gender == "Wanita" && (age > 19 && age < 60) {
		return true
	} else {
		return false
	}
}

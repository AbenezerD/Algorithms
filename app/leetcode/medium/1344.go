package medium

func angleClock(hour int, minutes int) float64 {
	hourDegree := 30.0*float64(hour) + float64(minutes)/2
	minutesDegree := float64(minutes) * 6

	angle := hourDegree - minutesDegree
	if angle < 0 {
		angle = -1 * float64(angle)
	}

	if angle > 180 {
		angle = 360 - angle
	}

	return angle
}

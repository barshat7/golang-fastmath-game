package scorer

// Score the user based on previous score and correct or incorrect
func Score(isCorrect bool, previousScore float64, timeTaken int64) float64 {
	score := previousScore
	floatTimeTaken := float64(timeTaken)
	if isCorrect {
		score = (previousScore + 1.0) * (1000 / floatTimeTaken)
	}
	return score
}
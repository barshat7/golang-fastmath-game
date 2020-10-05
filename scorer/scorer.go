package scorer

// Score the user based on previous score and correct or incorrect
func Score(isCorrect bool, previousScore float64) float64 {
	score := previousScore
	if isCorrect {
		score = previousScore + 1.0
	}
	return score
}
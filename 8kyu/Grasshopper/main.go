package kata_grasshopper

func combat(health, damage float64) float64 {
	health -= damage

	if health < 0 {
		return 0
	}

	return health
}

package robot

// Model type is used to define a robots possible behaviors.
type Model uint

const (
	MobileBot Model = iota
	MobileDiagonalBot
	MobileCraneBot
	MobileDiagonalCraneBot
)

// getRobotCommands returns the valid commands according to the provided Model.
func getRobotCommands(m Model) []string {
	switch m {
	case MobileBot:
		return []string{"N", "E", "S", "W"}
	case MobileDiagonalBot:
		return []string{"NE", "EN", "NW", "WN", "SE", "ES", "SW", "WS", "N", "E", "S", "W"}
	case MobileCraneBot:
		return []string{"N", "E", "S", "W", "G", "D"}
	case MobileDiagonalCraneBot:
		return []string{"NE", "EN", "NW", "WN", "SE", "ES", "SW", "WS", "N", "E", "S", "W", "G", "D"}
	default:
		return []string{}
	}
}

// ALL DONE.

// Package space computes one's age on Solar System planets' given their lifetime in seconds
package space

type Planet string


const earthPeriod = 31557600
// maps the planet name to seconds per year on it
var periods = map[Planet]float64{
	"Earth":   earthPeriod,
	"Mercury": 0.2408467 * earthPeriod,
	"Venus":   0.61519726 * earthPeriod,
	"Mars":    1.8808158 * earthPeriod,
	"Jupiter": 11.862615 * earthPeriod,
	"Saturn":  29.447498 * earthPeriod,
	"Uranus":  84.016846 * earthPeriod,
	"Neptune": 164.79132 * earthPeriod,
}

// takes one's float64 lifetime in seconds and Solar System planet string
// returns one's age on that planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / periods[planet]
}

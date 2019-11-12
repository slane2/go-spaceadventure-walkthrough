package main

import "github.com/slane2/go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(spaceadventure.PlanetarySystem{Name: "Solar System", Planets: []spaceadventure.Planet{}})
}

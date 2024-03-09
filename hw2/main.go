package main

import "fmt"

const (
	HighWalls string = "high walls"
	Water     string = "clean water"
	Lianas    string = "natural lianas"
)

type AnimalName string

type Animal struct {
	name        AnimalName
	animType    string
	catched     bool
	cageRequire string
}

var animals = []Animal{
	{
		"Alex",
		"lion",
		false,
		HighWalls,
	}, {
		"Marty",
		"zebra",
		false,
		HighWalls,
	}, {
		"Melman",
		"giraffe",
		false,
		HighWalls,
	}, {
		"Gloria",
		"hippo",
		true,
		Water,
	}, {
		"Skipper",
		"penguin",
		false,
		Water,
	}, {
		"Kowalski",
		"penguin",
		false,
		Water,
	}, {
		"Rico",
		"penguin",
		false,
		Water,
	}, {
		"King Julien XIII",
		"lemur",
		false,
		Lianas,
	}, {
		"Mason",
		"monkey",
		false,
		Lianas,
	}, {
		"Phil",
		"monkey",
		false,
		Lianas,
	},
}

type Cage struct {
	Feature      string
	Number       int
	AnimalPlaced *Animal
}

var cages = []Cage{
	{
		Water,
		1,
		nil,
	},
	{
		Water,
		2,
		&animals[3], // Gloria did not run away from the cage
	},
	{
		HighWalls,
		3,
		nil,
	},
	{
		HighWalls,
		4,
		nil,
	},
	{
		HighWalls,
		5,
		nil,
	},
	{
		Lianas,
		6,
		nil,
	},
	{
		HighWalls,
		7,
		nil,
	},
	{
		Lianas,
		8,
		nil,
	},
	{
		Lianas,
		9,
		nil,
	},
	{
		Water,
		10,
		nil,
	},
	{
		Lianas,
		11,
		nil,
	},
	{
		HighWalls,
		12,
		nil,
	},
	{
		Water,
		13,
		nil,
	},
	{
		Water,
		14,
		nil,
	},
	{
		Lianas,
		15,
		nil,
	},
}

type Zookeeper string

func (z Zookeeper) lookingForAnimals(anims []Animal) []Animal {
	for i := 0; i < len(anims); i++ {
		anims[i].catched = true
	}

	return anims
}

func (z Zookeeper) isAnimalPlacedInCages(animal Animal, cags []Cage) (status bool) {
	for i := 0; i < len(cags); i++ {
		if cags[i].AnimalPlaced != nil && cags[i].AnimalPlaced.name == animal.name {
			return true
		}
	}

	return false
}

func (z Zookeeper) isSuitableCage(animal Animal, cage Cage) (isSuitable bool) {
	if cage.Feature == animal.cageRequire && cage.AnimalPlaced == nil {
		cage.AnimalPlaced = &animal
		return true
	}

	return false
}

func (z Zookeeper) returnAnimalsToZoo(anims []Animal, cags []Cage) []Cage {
	for i := 0; i < len(anims); i++ {
		isInCage := z.isAnimalPlacedInCages(anims[i], cags)
		if isInCage {
			continue
		}

		for k := 0; k < len(cags); k++ {
			isSuitable := z.isSuitableCage(anims[i], cags[k])
			if isSuitable {
				cags[k].AnimalPlaced = &anims[i]
				break
			}
		}
	}

	return cags
}

func (c Cage) cageInfo() {
	if c.AnimalPlaced == nil {
		fmt.Printf("cage #%d (%s) is free\n", c.Number, c.Feature)

		return
	}
	fmt.Printf("%s %s in cage #%d (%s)\n", c.AnimalPlaced.animType, c.AnimalPlaced.name, c.Number, c.Feature)
}

func (z Zookeeper) showZoo(cags []Cage) {
	fmt.Printf("\nWelcome to %s's Zoo!\n\n~ ~ Roadmap ~ ~\n", z)
	for i := 0; i < len(cags); i++ {
		cags[i].cageInfo()
	}
}

func main() {
	const keeper Zookeeper = "Felix"
	var catchedAnimals = keeper.lookingForAnimals(animals)
	var cages = keeper.returnAnimalsToZoo(catchedAnimals, cages)
	keeper.showZoo(cages)
}

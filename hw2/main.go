package main

import "fmt"

const (
	HighWalls string = "high walls"
	Water     string = "clean water"
	Lianas    string = "natural lianas"
)

type Cage struct {
	Feature      string
	Number       int
	AnimalPlaced *Animal
}

var freeCages = [...]Cage{
	{
		Water,
		1,
		nil,
	},
	{
		Water,
		2,
		nil,
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

type Animal struct {
	name        string
	animType    string
	catched     bool
	cagePlaced  bool
	cageRequire string
}

var runawayAnimals = [...]Animal{
	{
		"Alex",
		"lion",
		false,
		false,
		HighWalls,
	}, {
		"Marty",
		"zebra",
		false,
		false,
		HighWalls,
	}, {
		"Melman",
		"giraffe",
		false,
		false,
		HighWalls,
	}, {
		"Gloria",
		"hippo",
		false,
		false,
		Water,
	}, {
		"Skipper",
		"penguin",
		false,
		false,
		Water,
	}, {
		"Kowalski",
		"penguin",
		false,
		false,
		Water,
	}, {
		"Rico",
		"penguin",
		false,
		false,
		Water,
	}, {
		"King Julien XIII",
		"lemur",
		false,
		false,
		Lianas,
	}, {
		"Mason",
		"monkey",
		false,
		false,
		Lianas,
	}, {
		"Phil",
		"monkey",
		false,
		false,
		Lianas,
	},
}

type Zookeeper string

func (z Zookeeper) lookingForAnimals(anims []Animal) []Animal {
	for i := 0; i < len(anims); i++ {
		anims[i].catched = true
	}

	return anims
}

func (z Zookeeper) returnAnimalsToZoo(anims []Animal, cags []Cage) []Cage {
	for i := 0; i < len(anims); i++ {
		if !anims[i].catched {
			continue
		}
		for k := 0; k < len(cags); k++ {
			if !anims[i].cagePlaced && cags[k].Feature == anims[i].cageRequire && cags[k].AnimalPlaced == nil {
				anims[i].cagePlaced = true
				cags[k].AnimalPlaced = &anims[i]
			}
		}
	}
	return cags
}

func (z Zookeeper) cageInfo(c Cage) {

	if c.AnimalPlaced == nil {
		fmt.Printf("cage #%d (%s) is free\n", c.Number, c.Feature)
		return
	}
	fmt.Printf("%s %s in cage #%d (%s)\n", c.AnimalPlaced.animType, c.AnimalPlaced.name, c.Number, c.Feature)
}

func (z Zookeeper) showZoo(k Zookeeper, cage []Cage) {
	fmt.Printf("\nWelcome to %s's Zoo!\n\n~ ~ Roadmap ~ ~\n", k)
	for i := 0; i < len(cage); i++ {
		z.cageInfo(cage[i])
	}
}

func main() {
	const keeper Zookeeper = "Felix"
	var catchedAnimals = keeper.lookingForAnimals(runawayAnimals[:])
	var zoo = keeper.returnAnimalsToZoo(catchedAnimals[:], freeCages[:])
	keeper.showZoo(keeper, zoo)
}

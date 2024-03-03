package main

import "fmt"

type DateYear int

const (
	part1 string = "народився"
	part2 string = "року Одесі в єврейській родині. Активний учасник російського народницького руху. Через переслідування російською владою за політичні погляди переїхав до Лейпцига. Працював там столяром, приєднався до соціал-демократичного руху. Згодом увійшов до студентського гуртка, де познайомився з Кларою Айснер, яка пізніше стала відомою як"
	part3 string = "року під час зустрічі з соціалістами Августом Бебелем та Вільгельмом Лібкнехтом відповідно до антисоціалістичного закону Осип Центкін був заарештований та висланий з Лейпцига, як «проблемний іноземець»"
	part4 string = "У Парижі молода родина часто змінювала орендовані квартири, бо не мали постійної роботи: Осип публікував статті та нотатки в газетах лівої спрямованості, заробляв перекладами та викладанням іноземних мов, а Клара давала приватні уроки та прала білизну в багатих сім'ях. Але більшість часу сім'я Цеткін жили в крихітній квартирі на Монмартрі. Там же народилися двоє їхніх синів:"
	part5 string = "Осип Цеткін зіграв важливу роль у підготовці першого конгресу Робітничого Інтернаціоналу в"
	part6 string = "році. Того ж року він помер від туберкульозу"

	heroName        string   = "Осип Цеткін"
	heroBirthDate   DateYear = 1850
	heroDeathDate   DateYear = 1889
	wifeName        string   = "Клара Цеткін"
	extraditionDate DateYear = 1880
	child1Name      string   = "Максим"
	child1BirthDate DateYear = 1883
	child1DeathDate DateYear = 1965
	child2Name      string   = "Костянтин"
	child2BirthDate DateYear = 1885
	child2DeathDate DateYear = 1980
)

func getBio() {
	fmt.Printf("%s\n%s\n%s\n%s\n", sentence1(), sentence2(), sentence3(), sentence4())
}

func sentence1() string {
	return fmt.Sprintf("%s %s %d %s %s.", heroName, part1, heroBirthDate, part2, wifeName)
}

func sentence2() string {
	return fmt.Sprintf("%d %s.", extraditionDate, part3)
}

func sentence3() string {
	return fmt.Sprintf("%s %s (%d-%d) %s (%d-%d).", part4, child1Name, child1BirthDate, child1DeathDate, child2Name, child2BirthDate, child2DeathDate)
}

func sentence4() string {
	return fmt.Sprintf("%s %s %d %s.", heroName, part5, heroDeathDate, part6)
}

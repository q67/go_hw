package main

type SceneName string
type Item string
type SceneDesc struct {
	desc       string
	actions    []Action
	actionList []SceneName
	items      []Item
}

var Scenes = make(map[SceneName]*SceneDesc)

var scene1wake SceneName = "scene1wake"
var key1 Item = "key1"
var code1 Item = "7190"

var scene2 SceneName = "scene2"

type Action struct {
	whatToDo       string
	resultHappened string
	modify         func()
	nextScene      SceneName
}

var Pocket = []Item{}

func addKeyToPocket() {
	Pocket = append(Pocket, key1)
}

func writeCode1() {
	Pocket = append(Pocket, code1)
}

func InitScenes() {

	Scenes[scene1wake] = &SceneDesc{
		"Ви проснулись в кімнаті на ліжку. Поряд з ліжком стоїть тумбочка, яка має одну шуфлядку. " +
			"На тумбочці лежить книжка-оповідання відомого письменника",
		[]Action{
			{
				"Подивитись що в тумбочці",
				"В шухляді лежить ключ",
				nil,
				scene1wake,
			},
			{
				"Погортати книжку",
				"Всередині обгортки написані цифри 7190",
				nil,
				scene1wake,
			},
			{
				"Покласти ключ до карману",
				"Ключ перекладено до карману",
				addKeyToPocket,
				scene1wake,
			},
			{
				"Десь записати ці цифри",
				"був знайдений клаптик паперу та записані цифри",
				writeCode1,
				scene1wake,
			},
		},
		[]SceneName{scene1wake, scene2},
		[]Item{key1, book1},
	}

}

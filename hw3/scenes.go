package main

type sceneId string

type scene struct {
	desc       string
	nextScenes []sceneId
}

const Start sceneId = "Почати гру заново"
const NearCave sceneId = "До печери"
const Forest sceneId = "Йти ліс"
const Cave sceneId = "Зайти в печеру"
const Camp sceneId = "До табіру"
const Tent sceneId = "Зайти в намет"
const Exit sceneId = "Припинити гру"
const Safe sceneId = "Сейф"
const Death sceneId = "Непритомність"
const Flashlight sceneId = "Ліхтарик"
const Matches sceneId = "Сірники"

func InitScenes() map[sceneId]scene {
	scenes := map[sceneId]scene{
		Start: {
			"Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.",
			[]sceneId{Forest, Cave},
		},

		NearCave: {
			"Стівен біля печери. З ним рюкзак, в якому є сірники, ліхтарик і ніж.",
			[]sceneId{Forest, Cave},
		},

		Cave: {
			"Стівен в печері але в ній темно, треба чимось підсвітити",
			[]sceneId{Flashlight, Matches},
		},

		Forest: {
			"У лісі Стівен натикається на мертве тіло дивної тварини.",
			[]sceneId{NearCave, Camp},
		},

		Camp: {
			"Через деякий час Стівен приходить до безлюдного табору. Він вже втомлений і вирішує відпочити, а не йти далі. Заходить в найближчий намет",
			[]sceneId{Forest, Tent},
		},

		Tent: {
			"У найближчому наметі він знаходить сейф з кодовим замком з двох чисел. ",
			[]sceneId{Safe, Camp},
		},

		Safe: {
			"Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.",
			[]sceneId{Death},
		},

		Death: {
			"Стівен непритомніє. А все могло бути зовсім інакше.",
			[]sceneId{Start},
		},
	}

	return scenes
}

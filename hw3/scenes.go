package main

type sceneId string

type scene struct {
	desc       string
	nextScenes []sceneId
}

const Start sceneId = "start"
const NearCave sceneId = "Біля печери"
const Forest sceneId = "Ліс"
const Cave sceneId = "Печера"
const Camp sceneId = "Табір"
const Tent sceneId = "Намет"
const Exit sceneId = "Припинити гру"

func InitScenes() map[sceneId]scene {
	return map[sceneId]scene{
		Start: {
			"Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.",
			[]sceneId{Forest, Cave},
		},

		NearCave: {
			"Стівен біля печери. З ним рюкзак, в якому є сірники, ліхтарик і ніж.",
			[]sceneId{Forest, Cave},
		},

		Forest: {
			"У лісі Стівен натикається на мертве тіло дивної тварини.",
			[]sceneId{NearCave, Camp},
		},

		Camp: {
			"Через деякий час Стівен приходить до безлюдного табору. Він вже втомлений і вирішує відпочити, а не йти далі. Заходить в найближчий намет",
			[]sceneId{Forest, Tent},
		},
	}
}

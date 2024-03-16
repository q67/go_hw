package main

type sceneId string

// type scenesList []map[sceneId]string

type nextScene struct {
	nextSceneId sceneId
	action      string
}
type scene struct {
	desc       string
	nextScenes *[]nextScene
}

type scenesIn map[sceneId]scene

const Start sceneId = "Заново"
const NearCave sceneId = "Біля печери"
const Forest sceneId = "Ліс"
const Cave sceneId = "Печера"
const Camp sceneId = "Табір"
const Tent sceneId = "Намет"
const Exit sceneId = "Припинити гру"
const Strongbox sceneId = "Сейф"
const StrongBoxOpen sceneId = "Сейф-відчинений"
const Death sceneId = "Непритомність"
const Flashlight sceneId = "Ліхтарик"
const Matches sceneId = "Сірники"

func InitScenes() scenesIn {
	return scenesIn{
		Start: {
			"Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.",
			&[]nextScene{
				{Forest, "Йти до лісу"},
				{Cave, "Оглянути печеру"},
			},
		},

		NearCave: {
			"Стівен біля печери. З ним рюкзак, в якому є сірники, ліхтарик і ніж.",
			&[]nextScene{
				{Forest, "Йти до лісу"},
				{Cave, "Зайти печеру"},
			},
		},

		Cave: {
			"В печері але в ній темно, треба чимось підсвітити",
			&[]nextScene{
				{Matches, "Спробувати сірники"},
				{Flashlight, "Здається був ліхтарик"},
				{NearCave, "Вийти з печери"},
			},
		},

		Matches: {
			"Сірники запалюються і швидко згорають, але все одно Стівен замітив надпис на стіні",
			&[]nextScene{
				{Flashlight, "Дістати ліхтарик"},
				{NearCave, "Вийти з печери"},
			},
		},

		Flashlight: {
			"На стіні хтось написав 3489, цікаво, що це може бути",
			&[]nextScene{
				{NearCave, "Запам'ятати ці цифри та вийти"},
			},
		},

		Forest: {
			"У лісі Стівен натикається на мертве тіло дивної тварини.",
			&[]nextScene{
				{NearCave, "Повернутись до печери"},
				{Camp, "Нічого не робити та йти до табору"},
			},
		},

		Camp: {
			"Стівен приходить до безлюдного табору. Він вже втомлений і вирішує відпочити, а не йти далі.",
			&[]nextScene{
				{Forest, "Повернутись до лісу"},
				{Tent, "Зайти в найближчий намет"},
			},
		},

		Tent: {
			"У найближчому наметі він знаходить сейф з кодовим замком",
			&[]nextScene{
				{Strongbox, "Спробувати ввести якийсь код"},
				{Camp, "Вийти з намету"},
			},
		},

		Strongbox: {
			"Сейф пише введений код не підходить",
			&[]nextScene{
				{Strongbox, "Ще раз спробувати ввести інший код"},
				{Camp, "Вийти з намету"},
			},
		},

		StrongBoxOpen: {
			"Cейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.",
			&[]nextScene{
				{Start, "Стівен непритомніє. А все могло бути зовсім інакше. Почати гру заново"},
			},
		},
	}
}

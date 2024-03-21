package main

import (
	"hw6/primitives"
	"hw6/routing"
	"hw6/transport"
)

func main() {
	start := primitives.Point{
		Address: "вул. Тичини",
		City:    "Київ",
		Country: "Україна",
	}

	transfer1 := primitives.Point{
		Address: "Дарниця",
		City:    "Київ",
		Country: "Україна",
	}

	transfer2 := primitives.Point{
		Address: "Аеропорт Жуляни",
		City:    "Київ",
		Country: "Україна",
	}

	transfer3 := primitives.Point{
		Address: "Аеропорт Дортмунд",
		City:    "Дортмунд",
		Country: "Німеччина",
	}

	finish := primitives.Point{
		Address: "Вокзал",
		City:    "Кельн",
		Country: "Німеччина",
	}

	route := routing.NewRoute()
	route.AddTransport("Маршрутка 575", &transport.Bus{}, start, transfer1)
	route.AddTransport("Автобус 44", &transport.Bus{}, transfer1, transfer2)
	route.AddTransport("Рейс 8943", &transport.Plane{}, transfer2, transfer3)
	route.AddTransport("Rhein Bahn 03", &transport.Train{}, transfer3, finish)

	route.ShowTransport()
}

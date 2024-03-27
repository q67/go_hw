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
	bus1 := transport.NewBus("Маршрутка 575", start, transfer1)
	route.AddTransport(bus1)
	bus2 := transport.NewBus("Автобус 44", transfer1, transfer2)
	route.AddTransport(bus2)
	plane := transport.NewPlane("Рейс 8943", transfer2, transfer3)
	route.AddTransport(plane)
	train := transport.NewTrain("Rhein Bahn 03", transfer3, finish)
	route.AddTransport(train)

	route.ShowRoute()
}

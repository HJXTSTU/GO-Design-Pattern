package observer

type(
	Event struct{
		Data int64
	}

	Observer interface{
		OnNotify(Event)
	}

	Notifier interface{
		Register(Observer)

		Degister(Observer)

		Notify(Event)
	}
)



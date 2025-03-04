package status

type Status string

const (
	NEW        Status = "НОВАЯ"
	INPROGRESS Status = "В РАБОТЕ"
	PENDING    Status = "ОТЛОЖЕНА"
	CANCELED   Status = "ОТМЕНЕНА"
	DONE       Status = "ВЫПОЛНЕНА"
)

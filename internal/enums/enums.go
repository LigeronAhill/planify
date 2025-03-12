package enums

type Priority string

const (
	HIGH   Priority = "ВЫСОКИЙ"
	NORMAL Priority = "ОБЫЧНЫЙ"
	LOW    Priority = "НИЗКИЙ"
)

type Status string

const (
	NEW        Status = "НОВАЯ"
	INPROGRESS Status = "В РАБОТЕ"
	PENDING    Status = "ОТЛОЖЕНА"
	CANCELED   Status = "ОТМЕНЕНА"
	DONE       Status = "ВЫПОЛНЕНА"
)

type Category string

const (
	TASK Category = "ДЕЛО"
	MEET Category = "ВСТРЕЧА"
	CALL Category = "ЗВОНОК/СООБЩЕНИЕ"
)

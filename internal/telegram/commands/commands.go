package commands

const (
	Start      = "/start"
	Help       = "/help"
	Cancel     = "/cancel"
	AddTask    = "Создать задачу"
	MyTasks    = "Мои задачи"
	DoneTasks  = "Выполненные"
	Statistics = "Статистика"
)

func List() []string {
	return []string{Start, Help, Cancel, AddTask, MyTasks, DoneTasks, Statistics}
}

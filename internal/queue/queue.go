package queue

func NewQueue() chan string {
	return make(chan string, 100)
}

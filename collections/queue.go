package collections

/*
	Queue (Ring Buffer)

	Search O(n)
	Insert O(1)
	Delete O(1)
*/

type QueueParams interface {
	string | int | bool
}

type Queue[T QueueParams] struct {
	queue []T
	start int
	end int
	capacity int
}

func NewQueue[T QueueParams](capacity int) *Queue[T]{
	return &Queue[T]{
		queue : []T{},
		start : 0,
		end : 0,
		capacity : capacity,
	}
}

func (q *Queue[QueueParams]) AddFront(data QueueParams){
	
}

func (q *Queue[QueueParams]) AddBack(data QueueParams){

}

func (q *Queue[QueueParams]) PopFront(){

}

func (q *Queue[QueueParams]) PopBack(){

}

func (q *Queue[QueueParams]) Clear(){

}

func (q *Queue[QueueParams]) GetLength(){

}

func (q *Queue[QueueParams]) isOverCapacity(){

}
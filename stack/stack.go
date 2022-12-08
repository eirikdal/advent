package stack

type Stack[T any] []T

func (stack *Stack[T]) Multipush(c []T) {
	arr := make([]T, 0)
	arr = append(arr, c...)

	*stack = append(arr, *stack...)
}

func (stack *Stack[T]) Push(c T) {
	arr := make([]T, 0)
	arr = append(arr, c)

	*stack = append(arr, *stack...)
}

func (stack *Stack[T]) Pop(number int) Stack[T] {
	el := (*stack)[0:number]
	*stack = (*stack)[number:]
	return el
}

func (stack *Stack[T]) PopOne() T {
	el := (*stack)[0]
	*stack = (*stack)[1:]
	return el
}

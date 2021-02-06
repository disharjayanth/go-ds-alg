package main

import "fmt"

// Queue is slice of pointer to Order
type Queue []*Order

// Order type has priority field, bigger priority come first
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func New(priority int, quantity int, product string, customerName string) *Order {
	order := &Order{}
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
	return order
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func (queue *Queue) Remove() *Order {
	if len(*queue) == 0 {
		return nil
	}

	order := (*queue)[0]
	*queue = (*queue)[1:]
	return order
}

func main() {
	queue := &Queue{}

	order1 := New(2, 20, "Computer", "John")
	order2 := New(1, 10, "Monitor", "Smith")
	order3 := New(4, 22, "Mobile", "Rames")
	order4 := New(3, 42, "Brush", "Lee")
	order5 := New(10, 2, "Plane Ticket", "Ar")
	order6 := New(6, 12, "Wire", "Y")

	queue.Add(order1)
	queue.Add(order2)
	queue.Add(order3)
	queue.Add(order4)
	queue.Add(order5)
	queue.Add(order6)

	for i, order := range *queue {
		fmt.Println("Order no:", i, order)
	}

	fmt.Println("Removing Item from Queue")
	removedItem := queue.Remove()
	fmt.Println("Removed item:", *removedItem)

	fmt.Println("Queue after removing item")
	for i, order := range *queue {
		fmt.Println("Order no:", i, order)
	}
}

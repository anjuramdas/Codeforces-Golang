package main

import "fmt"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.elements = make([]interface{}, 0)
	return q
}
func (q *Queue) Enqueue(e interface{}) *Queue {
	q.elements = append(q.elements, e)
	return q
}
func (q *Queue) Dequeue() (e interface{}) {
	e, q.elements = q.elements[0], q.elements[1:len(q.elements)]
	return
}
func (q *Queue) Size() int {
	return len(q.elements)
}
func Min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}
func main() {
	var floors, start, goal, up, down, up_floor, down_floor int64
	fmt.Scan(&floors, &start, &goal, &up, &down)
	var vis [1000001]int64
	var dist [1000001]int64
	if start == goal {
		fmt.Println(0)
	} else {
		if 1 <= start && goal <= floors && floors <= 1000000 && 0 <= up && down <= 1000000 {
			vis[start] = 1
			dist[start] = 0
			dist[goal] = -1
			q := NewQueue()
			q.Enqueue(start)
			r := q.Size()
			for r > 0 {
				val := q.Dequeue().(int64)
				up_floor = val + up
				down_floor = val - down
				if up_floor <= floors && vis[up_floor] != 1 {
					vis[up_floor] = 1
					dist[up_floor] = dist[val] + 1
					q.Enqueue(up_floor)
				}
				if down_floor >= 1 && vis[down_floor] != 1 {
					vis[down_floor] = 1
					dist[down_floor] = dist[val] + 1
					q.Enqueue(down_floor)
				}
				r = q.Size()
			}
			if dist[goal] == -1 {
				fmt.Println("use the stairs")
			} else {
				fmt.Println(dist[goal])
			}
		}
	}
}

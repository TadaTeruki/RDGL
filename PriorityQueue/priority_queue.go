/*
github.com/TadaTeruki/RDGL/PriorityQueue/priority_queue.go
Copyright (C) 2021 Tada Teruki

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation; either version 3 of the License, or 
(at your option) any later version.

This program is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the 
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/
package priority_queue

type Object struct {
	Value	 interface{}
	Priority float64
}

type PriorityQueue struct {
	Tree []Object
	
}

func MakeObject(value interface{}, priority float64) Object{
	var obj Object
	obj.Value = value
	obj.Priority = priority
	return obj
}

func (queue *PriorityQueue) GetSize() int{
	return len(queue.Tree)
}

func (queue *PriorityQueue) GetFront() *Object{
	if queue.GetSize() == 0 {
		return nil
	}
	return &queue.Tree[0]
}

func (queue *PriorityQueue) Pop(){

	if queue.GetSize() == 0 {
		return
	}

	obj_id := 0
	queue.Tree[obj_id] = queue.Tree[queue.GetSize()-1]

	for ;obj_id < queue.GetSize();{
		a_id := obj_id*2+1
		b_id := a_id+1

		a_ok := queue.GetSize() > a_id && queue.Tree[a_id].Priority > queue.Tree[obj_id].Priority
		b_ok := queue.GetSize() > b_id && queue.Tree[b_id].Priority > queue.Tree[obj_id].Priority

		if b_ok && (a_ok == false || queue.Tree[b_id].Priority > queue.Tree[a_id].Priority){
			swap := queue.Tree[obj_id]
			queue.Tree[obj_id] = queue.Tree[b_id]
			queue.Tree[b_id] = swap
			obj_id = b_id
		} else if a_ok {
			swap := queue.Tree[obj_id]
			queue.Tree[obj_id] = queue.Tree[a_id]
			queue.Tree[a_id] = swap
			obj_id = a_id
		} else {
			break
		}
	}

	queue.Tree = queue.Tree[0:queue.GetSize()-1]

}

func (queue *PriorityQueue) Push(obj Object){
	queue.Tree = append(queue.Tree, obj)

	obj_id := queue.GetSize()-1

	for ;obj_id != 0;{
		parent_id := (obj_id-1)/2

		if queue.Tree[obj_id].Priority > queue.Tree[parent_id].Priority{
			swap := queue.Tree[obj_id]
			queue.Tree[obj_id] = queue.Tree[parent_id]
			queue.Tree[parent_id] = swap
			obj_id = parent_id
		} else {
			break
		}
	}
}

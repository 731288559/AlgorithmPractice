package leetcode

import "container/heap"

// 3408. 设计任务管理器

type TaskManager struct {
	taskIdToTaskMap map[int]*Task
	taskHeap        taskHeap
}

type Task struct {
	UserId   int
	TaskId   int
	Priority int
}

func Constructor3408(tasks [][]int) TaskManager {
	mgr := TaskManager{
		taskIdToTaskMap: make(map[int]*Task, len(tasks)),
	}

	for _, task := range tasks {
		mgr.Add(task[0], task[1], task[2])
	}
	return mgr
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	t := &Task{
		UserId:   userId,
		TaskId:   taskId,
		Priority: priority,
	}
	this.taskIdToTaskMap[taskId] = t
	heap.Push(&this.taskHeap, t)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	this.Add(this.taskIdToTaskMap[taskId].UserId, taskId, newPriority)
}

func (this *TaskManager) Rmv(taskId int) {
	this.taskIdToTaskMap[taskId] = nil
}

func (this *TaskManager) ExecTop() int {
	// 按注释这么写会异常
	//h := this.taskHeap
	//for h.Len() > 0 {
	//	top := heap.Pop(&h).(*Task)
	//	task := this.taskIdToTaskMap[top.TaskId]
	//	if task != nil && task.Priority == top.Priority && task.UserId == top.UserId {
	//		this.Rmv(task.TaskId)
	//		return task.UserId
	//	}
	//}
	for this.taskHeap.Len() > 0 {
		top := heap.Pop(&this.taskHeap).(*Task)
		task := this.taskIdToTaskMap[top.TaskId]
		if task != nil && task.Priority == top.Priority && task.UserId == top.UserId {
			this.Rmv(task.TaskId)
			return task.UserId
		}
	}
	return -1
}

type taskHeap struct {
	list []*Task
}

func (t taskHeap) Len() int {
	return len(t.list)
}

func (t taskHeap) Less(i, j int) bool {
	if t.list[i].Priority != t.list[j].Priority {
		return t.list[i].Priority > t.list[j].Priority
	}
	return t.list[i].TaskId > t.list[j].TaskId
}

func (t taskHeap) Swap(i, j int) {
	t.list[i], t.list[j] = t.list[j], t.list[i]
}

func (t *taskHeap) Push(x any) {
	t.list = append(t.list, x.(*Task))
}

func (t *taskHeap) Pop() any {
	//if len(t.list) == 0 {
	//	return nil
	//}
	//top := t.list[len(t.list)-1]
	//t.list = t.list[:len(t.list)-1]
	//return top
	a := t.list
	top := a[len(a)-1]
	t.list = a[:len(a)-1]
	return top
}

var _ heap.Interface = &taskHeap{}

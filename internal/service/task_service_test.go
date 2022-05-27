package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/JimySheepman/to-do-api/internal/domain/task"
)

type postgresqlTaskMockRepository struct {
	create map[int]error
	list   []task.Task
	update map[int]error
	delete map[int]error
}

func NewMockTaskRepository() *postgresqlTaskMockRepository {
	return &postgresqlTaskMockRepository{
		create: map[int]error{},
		list:   []task.Task{},
		update: map[int]error{},
		delete: map[int]error{},
	}
}

func (r *postgresqlTaskMockRepository) OnCreate(id int, err error) {
	r.create[id] = err
}

func (r *postgresqlTaskMockRepository) OnList(c int) {
	r.list = make([]task.Task, c, c)
}

func (r *postgresqlTaskMockRepository) OnUpdate(id int, err error) {
	r.update[id] = err
}

func (r *postgresqlTaskMockRepository) OnDelete(id int, err error) {
	r.delete[id] = err
}

func (r *postgresqlTaskMockRepository) CreateTask(ctx context.Context, task *task.Task) error {
	return r.create[task.Id]
}

func (r *postgresqlTaskMockRepository) ListTask(ctx context.Context) (*[]task.Task, error) {
	return &r.list, nil
}

func (r *postgresqlTaskMockRepository) UpdateTask(ctx context.Context, id int, task *task.Task) error {
	return r.update[id]
}

func (r *postgresqlTaskMockRepository) DeleteTask(ctx context.Context, id int) error {
	return r.delete[id]
}

func TestCreateTask(t *testing.T) {
	r := NewMockTaskRepository()
	s := NewTaskService(r)

	var tests = []struct {
		num    int
		tsk    *task.Task
		expect error
	}{
		{1, &task.Task{Id: 1}, nil},
		{3, &task.Task{Id: 3}, nil},
		{5, &task.Task{Id: 5}, nil},
		{10, &task.Task{Id: 10}, errors.New("10")},
		{100, &task.Task{Id: 100}, errors.New("100")},
		{1000, &task.Task{Id: 1000}, errors.New("1000")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnCreate(test.num, test.expect)
			actual := s.CreateTask(context.Background(), test.tsk)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func TestListTask(t *testing.T) {
	r := NewMockTaskRepository()
	s := NewTaskService(r)

	var tests = []struct {
		num    int
		expect error
	}{
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, nil},
		{6, nil},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnList(test.num)
			actual, _ := s.ListTask(context.Background())
			if !reflect.DeepEqual(len(*actual), test.num) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func TestUpdateTask(t *testing.T) {
	r := NewMockTaskRepository()
	s := NewTaskService(r)

	var tests = []struct {
		num    int
		tsk    *task.Task
		expect error
	}{
		{1, &task.Task{}, nil},
		{3, &task.Task{}, nil},
		{5, &task.Task{}, nil},
		{10, &task.Task{}, errors.New("10")},
		{100, &task.Task{}, errors.New("100")},
		{1000, &task.Task{}, errors.New("1000")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnUpdate(test.num, test.expect)
			actual := s.UpdateTask(context.Background(), test.num, test.tsk)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	r := NewMockTaskRepository()
	s := NewTaskService(r)

	var tests = []struct {
		num    int
		tsk    *task.Task
		expect error
	}{
		{1, &task.Task{}, nil},
		{3, &task.Task{}, nil},
		{5, &task.Task{}, nil},
		{10, &task.Task{}, errors.New("10")},
		{100, &task.Task{}, errors.New("100")},
		{1000, &task.Task{}, errors.New("1000")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("actual:%d expected error:%v", test.num, test.expect)
		t.Run(testName, func(t *testing.T) {
			r.OnDelete(test.num, test.expect)
			actual := s.DeleteTask(context.Background(), test.num)
			if !reflect.DeepEqual(actual, test.expect) {
				t.Errorf("%d value error", test.num)
			}
		})
	}
}

func ExampleCreateComment(t *testing.T) {
	r := NewMockTaskRepository()
	s := NewTaskService(r)

	r.OnCreate(1000, nil)
	r.OnCreate(2000, errors.New("test"))

	tsk1 := &task.Task{
		Id: 1000,
	}

	tsk2 := &task.Task{
		Id: 2000,
	}

	actual1 := s.CreateTask(context.Background(), tsk1)
	var expected1 error

	actual2 := s.CreateTask(context.Background(), tsk2)
	expected2 := errors.New("test")

	if actual1 != expected1 {
		t.Error("test1")
	}

	if !reflect.DeepEqual(actual2, expected2) {
		t.Error("test2")
	}
}

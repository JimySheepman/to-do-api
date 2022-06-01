package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/JimySheepman/to-do-api/internal/domain/task"
	"github.com/gofiber/fiber/v2"
)

type taskMockService struct {
	create map[int]error
	list   []task.Task
	update map[int]error
	delete map[int]error
}

func NewTaskMockService() *taskMockService {
	return &taskMockService{
		create: map[int]error{},
		list:   []task.Task{},
		update: map[int]error{},
		delete: map[int]error{},
	}
}

func (s *taskMockService) OnCreate(id int, err error) {
	s.create[id] = err
}

func (s *taskMockService) OnList(c int) {
	s.list = make([]task.Task, c, c)
}

func (s *taskMockService) OnUpdate(id int, err error) {
	s.update[id] = err
}

func (s *taskMockService) OnDelete(id int, err error) {
	s.delete[id] = err
}

func (s *taskMockService) CreateTask(ctx context.Context, task *task.Task) error {
	return s.create[task.Id]
}

func (s *taskMockService) ListTask(ctx context.Context) (*[]task.Task, error) {
	if len(s.list) == 0 {
		return nil, errors.New("service unavailable")
	}
	return &s.list, nil
}

func (s *taskMockService) UpdateTask(ctx context.Context, id int, task *task.Task) error {
	return s.update[id]
}

func (s *taskMockService) DeleteTask(ctx context.Context, id int) error {
	return s.delete[id]
}

func Test_createTask(t *testing.T) {

	//setup
	app := fiber.New()
	s := NewTaskMockService()
	NewTaskRouter(app.Group("/"), s)

	var tests = []struct {
		testName         string
		method           string
		path             string
		contentTypeKey   string
		contentTypeValue string
		body             task.Task
		expectedStatus   int
		expectedErrorId  int
		expectedError    error
	}{
		{
			testName:         "bad request error",
			method:           "POST",
			path:             "/",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "text/plain;charset=utf-8",
			body: task.Task{
				Id:       1,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  400,
			expectedErrorId: 1,
			expectedError:   nil,
		},
		{
			testName:         "service unavailable error",
			method:           "POST",
			path:             "/",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       2,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  500,
			expectedErrorId: 2,
			expectedError:   errors.New("service unavailable"),
		},
		{
			testName:         "wrong path request",
			method:           "POST",
			path:             "/sele/sa",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       3,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  404,
			expectedErrorId: 3,
			expectedError:   nil,
		},
		{
			testName:         "successful request",
			method:           "POST",
			path:             "/",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       4,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  201,
			expectedErrorId: 4,
			expectedError:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			body, _ := json.Marshal(test.body)

			s.OnCreate(test.expectedErrorId, test.expectedError)
			req := httptest.NewRequest(test.method, test.path, bytes.NewReader(body))
			req.Header.Add(test.contentTypeKey, test.contentTypeValue)
			resp, _ := app.Test(req)

			actual := resp.StatusCode
			want := test.expectedStatus
			if !reflect.DeepEqual(actual, want) {
				t.Errorf("actual:%v want:%v", actual, want)
			}
		})
	}
}

func Test_listTask(t *testing.T) {

	//setup
	app := fiber.New()
	s := NewTaskMockService()
	NewTaskRouter(app.Group("/"), s)

	var tests = []struct {
		testName         string
		method           string
		path             string
		contentTypeKey   string
		contentTypeValue string
		body             *task.Task
		expectedStatus   int
		expectedErrorId  int
		expectedError    error
	}{
		{
			testName:         "service unavailable error",
			method:           "GET",
			path:             "/",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: &task.Task{
				Id:       0,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  500,
			expectedErrorId: 0,
			expectedError:   errors.New("service unavailable"),
		},
		{
			testName:         "successful request",
			method:           "GET",
			path:             "/",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "text/plain;charset=utf-8",
			body: &task.Task{
				Id:       1,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  200,
			expectedErrorId: 1,
			expectedError:   nil,
		},
		{
			testName:         "wrong path request",
			method:           "GET",
			path:             "/sele/sa",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: &task.Task{
				Id:       2,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  404,
			expectedErrorId: 2,
			expectedError:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			body, _ := json.Marshal(test.body)

			s.OnList(test.expectedErrorId)

			req := httptest.NewRequest(test.method, test.path, bytes.NewReader(body))
			req.Header.Add(test.contentTypeKey, test.contentTypeValue)
			resp, _ := app.Test(req)

			actual := resp.StatusCode
			want := test.expectedStatus
			if !reflect.DeepEqual(actual, want) {
				t.Errorf("actual:%v want:%v", actual, want)
			}
		})
	}
}

func Test_updateTask(t *testing.T) {
	//setup
	app := fiber.New()
	s := NewTaskMockService()
	NewTaskRouter(app.Group("/"), s)

	var tests = []struct {
		testName         string
		method           string
		path             string
		contentTypeKey   string
		contentTypeValue string
		body             task.Task
		expectedStatus   int
		expectedErrorId  int
		expectedError    error
	}{
		{
			testName:         "wrong path request",
			method:           "PUT",
			path:             "/sele/sa",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       1,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  404,
			expectedErrorId: 1,
			expectedError:   nil,
		},
		{
			testName:         "wrong path param type error",
			method:           "PUT",
			path:             "/sele",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       2,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  400,
			expectedErrorId: 2,
			expectedError:   errors.New("param type error"),
		},
		{
			testName:         "bad request error",
			method:           "PUT",
			path:             "/3",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "text/plain;charset=utf-8",
			body: task.Task{
				Id:       3,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  400,
			expectedErrorId: 3,
			expectedError:   nil,
		},
		{
			testName:         "service unavailable error",
			method:           "PUT",
			path:             "/4",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       4,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  500,
			expectedErrorId: 4,
			expectedError:   errors.New("service unavailable"),
		},
		{
			testName:         "successful request",
			method:           "PUT",
			path:             "/5",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       5,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  200,
			expectedErrorId: 5,
			expectedError:   nil,
		},
		{
			testName:         "miss id request",
			method:           "PUT",
			path:             "/6",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       6,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  500,
			expectedErrorId: 6,
			expectedError:   errors.New("id is not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			body, _ := json.Marshal(test.body)

			s.OnUpdate(test.expectedErrorId, test.expectedError)
			req := httptest.NewRequest(test.method, test.path, bytes.NewReader(body))
			req.Header.Add(test.contentTypeKey, test.contentTypeValue)
			resp, _ := app.Test(req)

			actual := resp.StatusCode
			want := test.expectedStatus
			if !reflect.DeepEqual(actual, want) {
				t.Errorf("actual:%v want:%v", actual, want)
			}
		})
	}
}

func Test_deleteTask(t *testing.T) {
	//setup
	app := fiber.New()
	s := NewTaskMockService()
	NewTaskRouter(app.Group("/"), s)

	var tests = []struct {
		testName         string
		method           string
		path             string
		contentTypeKey   string
		contentTypeValue string
		body             task.Task
		expectedStatus   int
		expectedErrorId  int
		expectedError    error
	}{
		{
			testName:         "wrong path request",
			method:           "DELETE",
			path:             "/sele/sa",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       1,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  404,
			expectedErrorId: 1,
			expectedError:   nil,
		},
		{
			testName:         "wrong path param type error",
			method:           "DELETE",
			path:             "/sele",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       2,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  400,
			expectedErrorId: 2,
			expectedError:   errors.New("param type error"),
		},
		{
			testName:         "miss id request",
			method:           "DELETE",
			path:             "/3",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       3,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  500,
			expectedErrorId: 3,
			expectedError:   errors.New("id is not found"),
		},
		{
			testName:         "service unavailable error",
			method:           "DELETE",
			path:             "/4",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       4,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  500,
			expectedErrorId: 4,
			expectedError:   errors.New("service unavailable"),
		},
		{
			testName:         "successful request",
			method:           "DELETE",
			path:             "/5",
			contentTypeKey:   "Content-Type",
			contentTypeValue: "application/json;charset=utf-8",
			body: task.Task{
				Id:       5,
				Title:    "test",
				Content:  "test",
				Category: "test",
				Statu:    "test",
			},
			expectedStatus:  204,
			expectedErrorId: 5,
			expectedError:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {

			body, _ := json.Marshal(test.body)

			s.OnDelete(test.expectedErrorId, test.expectedError)
			req := httptest.NewRequest(test.method, test.path, bytes.NewReader(body))
			req.Header.Add(test.contentTypeKey, test.contentTypeValue)
			resp, _ := app.Test(req)

			actual := resp.StatusCode
			want := test.expectedStatus
			if !reflect.DeepEqual(actual, want) {
				t.Errorf("actual:%v want:%v", actual, want)
			}
		})
	}
}

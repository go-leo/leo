package cqrs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/leo/v3/metadatax"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type notMetadata struct {
}

func (h notMetadata) Handle(ctx context.Context, in *struct{}) (http.Header, error) {
	return nil, nil
}

type noHandleIn struct{}

type noHandle struct {
}

func (h noHandle) Invoke(ctx context.Context, cmd *noHandleIn) error {
	return nil
}

func (h noHandle) Handle(ctx context.Context, in1, in2 *noHandleIn) error {
	return nil
}

type unknown struct{}

type unknownQuery struct{}

var _ CommandHandler[*studyCmd] = new(study)

type studyCmd struct{}

type study struct{}

func (s study) Handle(ctx context.Context, args *studyCmd) error {
	fmt.Println("studying...")
	<-time.After(1 + time.Second)
	return nil
}

var _ CommandHandler[examCmd] = new(exam)

var errNotPassed = errors.New("not passed")

type examCmd struct{}

type exam struct{}

func (e exam) Handle(ctx context.Context, cmd examCmd) error {
	fmt.Println("taking an exam ...")
	<-time.After(1 + time.Second)
	return errNotPassed
}

var _ QueryHandler[*studentQuery, *studentResult] = new(student)

type studentQuery struct {
	Id int
}

type studentResult struct {
	Name string
}

type student struct{}

func (s student) Handle(ctx context.Context, q *studentQuery) (*studentResult, error) {
	fmt.Printf("finding %d student\n", q.Id)
	<-time.After(1 + time.Second)
	return &studentResult{Name: "jax"}, nil
}

var _ QueryHandler[teacherQuery, teacherResult] = new(teacher)

var errNotFoundTeacher = errors.New("not found teacher")

type teacherQuery struct {
	Id int
}

type teacherResult struct {
	Name string
}

type teacher struct{}

func (s teacher) Handle(ctx context.Context, q teacherQuery) (teacherResult, error) {
	fmt.Printf("finding %d teacher\n", q.Id)
	<-time.After(1 + time.Second)
	return teacherResult{}, errNotFoundTeacher
}

func TestBus(t *testing.T) {
	var bus SampleBus
	var err error
	var metadata metadatax.Metadata
	_ = metadata
	err = bus.RegisterCommand(&study{})
	assert.NoError(t, err)

	err = bus.RegisterCommand(study{})
	assert.Error(t, err)

	err = bus.RegisterCommand(new(exam))
	assert.NoError(t, err)

	err = bus.RegisterCommand(new(exam))
	assert.Error(t, err)

	err = bus.RegisterCommand(exam{})
	assert.Error(t, err)

	err = bus.RegisterCommand(nil)
	assert.Error(t, err)

	err = bus.RegisterCommand(new(noHandle))
	assert.Error(t, err)

	err = bus.RegisterCommand(new(notMetadata))
	assert.Error(t, err)

	err = bus.Exec(context.Background(), unknown{})
	assert.Error(t, err)

	err = bus.Exec(context.Background(), studyCmd{})
	assert.Error(t, err)

	err = bus.Exec(context.Background(), &studyCmd{})
	assert.NoError(t, err)

	err = bus.Exec(context.Background(), examCmd{})
	assert.ErrorIs(t, err, errNotPassed)

	err = bus.Exec(context.Background(), &examCmd{})
	assert.Error(t, err)

	err = bus.RegisterQuery(&student{})
	assert.NoError(t, err)

	err = bus.RegisterQuery(student{})
	assert.Error(t, err)

	err = bus.RegisterQuery(&teacher{})
	assert.NoError(t, err)

	err = bus.RegisterQuery(teacher{})
	assert.Error(t, err)

	err = bus.RegisterQuery(new(noHandle))
	assert.Error(t, err)

	_, err = bus.Query(context.Background(), unknown{})
	assert.Error(t, err)

	_, err = bus.Query(context.Background(), studentQuery{})
	assert.Error(t, err)

	r, err := bus.Query(context.Background(), &studentQuery{Id: 10})
	assert.NoError(t, err)
	assert.EqualValues(t, &studentResult{Name: "jax"}, r)

	r, err = bus.Query(context.Background(), teacherQuery{Id: 1})
	assert.ErrorIs(t, err, errNotFoundTeacher)

}

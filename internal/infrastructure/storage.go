package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/WanderningMaster/tcli/internal/encoding"
	"github.com/WanderningMaster/tcli/internal/logger"
	"github.com/WanderningMaster/tcli/internal/model"
)

type StorageState uint8

const (
	INITIALIZED StorageState = 0
	ERRORED     StorageState = 1
	EMPTY       StorageState = 2
	IN_MEMORY   StorageState = 3
)

var (
	StorageNotReady = errors.New("Storage not ready yet")
	StorageEmpty    = errors.New("Storage is empty")
	TaskNotExists   = errors.New("Task with this ID not exists")
)

type Storage struct {
	dir    string
	tasks  []*model.Task
	parser encoding.Parser
	state  StorageState
}

func NewStorage(ctx context.Context, dir string, p encoding.Parser) *Storage {
	return &Storage{
		dir:    dir,
		parser: p,
		state:  INITIALIZED,
	}
}

func (s *Storage) LoadTasks(ctx context.Context) error {
	logger := logger.FromContext(ctx)

	fPath := path.Join(s.dir, fmt.Sprintf("tcli.%s", s.parser.Extension()))
	buff, err := os.ReadFile(fPath)

	if err != nil {
		_, err = os.Create(fPath)
		if err != nil {
			logger.Error(err.Error())
			s.state = ERRORED

			return err
		}
	}

	if len(buff) == 0 {
		s.state = EMPTY

		return nil
	}

	err = s.parser.Unmarshal(ctx, buff, &s.tasks)
	if err != nil {
		s.state = ERRORED
		return err
	}

	s.state = IN_MEMORY
	return nil
}

func (s *Storage) Tasks(ctx context.Context) ([]*model.Task, error) {
	if s.state == EMPTY {
		return nil, StorageEmpty
	}
	if s.state != IN_MEMORY {
		return nil, StorageNotReady
	}

	return s.tasks, nil
}

func (s *Storage) TasksByTag(ctx context.Context, tag string) ([]*model.Task, error) {
	if s.state == EMPTY {
		return nil, StorageEmpty
	}
	if s.state != IN_MEMORY {
		return nil, StorageNotReady
	}

	out := []*model.Task{}
	for _, t := range s.tasks {
		if t.Tag != tag {
			continue
		}

		out = append(out, t)
	}
	if len(out) == 0 {
		return nil, StorageEmpty
	}

	return out, nil
}

func (s *Storage) Add(ctx context.Context, tag string, content string) error {
	if s.state == ERRORED || s.state == INITIALIZED {
		return StorageNotReady
	}

	task := &model.Task{
		Content: content,
		Tag:     tag,
	}
	s.tasks = append(s.tasks, task)
	b, err := s.parser.Marshal(ctx, s.tasks)
	if err != nil {
		return err
	}

	fPath := path.Join(s.dir, fmt.Sprintf("tcli.%s", s.parser.Extension()))

	fd, err := os.Create(fPath)
	if err != nil {
		return err
	}

	_, err = fd.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Remove(ctx context.Context, id int) error {
	if s.state == ERRORED || s.state == INITIALIZED {
		return StorageNotReady
	}
	if s.state == EMPTY {
		return StorageEmpty
	}

	idx := id - 1
	if idx+1 > len(s.tasks) {
		return TaskNotExists
	}

	s.tasks = append(s.tasks[:idx], s.tasks[idx+1:]...)
	b, err := s.parser.Marshal(ctx, s.tasks)

	fPath := path.Join(s.dir, fmt.Sprintf("tcli.%s", s.parser.Extension()))

	fd, err := os.Create(fPath)
	if err != nil {
		return err
	}

	_, err = fd.Write(b)
	if err != nil {
		return err
	}

	if len(s.tasks) == 0 {
		err := os.Remove(fPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) Reset(ctx context.Context) error {
	fPath := path.Join(s.dir, fmt.Sprintf("tcli.%s", s.parser.Extension()))

	err := os.Remove(fPath)
	if err != nil {
		return err
	}

	return nil
}

package list

import (
	"dev-oleksandrv/todogo-cli/internal/datetime"
	"errors"
	"strconv"
	"time"
)

type List struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func NewList(id int, name string) *List {
	return &List{
		ID: id,
		Name: name,
		CreatedAt: time.Now(),
	}
}

func (l *List) ToRecord() []string {
	return []string{strconv.Itoa(l.ID), l.Name, datetime.GetFormattedTime(l.CreatedAt)}
}

func FromRecord(record []string) (*List, error) {
	parsedID, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, errors.New("invalid id")
	}
	parsedTime, err := datetime.ParseFormattedTime(record[2])
	if err != nil {
		return nil, errors.New("invalid time")
	}
	return &List{
		ID: parsedID,
		Name: record[1],
		CreatedAt: *parsedTime,
	}, nil
}
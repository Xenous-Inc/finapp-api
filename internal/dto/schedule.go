package dto

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Date string

func (d *Date) UnmarshalJSON(bytes []byte) error {
	date, err := time.Parse(`"2006-01-02T15:04:05.000-0700"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(date.Format("2024.12.01"))

	return nil
}

func (d *Date) String() string {
	if d == nil {
		return ""
	}

	return string(*d)
}

type GetScheduleRequest struct {
	EntityId  string `json:"id" validate:"required"`
	StartDate Date   `json:"startDate"`
	EndDate   Date   `json:"endDate"`
}

type LessonType uint8

const (
	LESSON_TYPE_VALUE_UNKNOWN      string = "Неизвестно"
	LESSON_TYPE_VALUE_LECTURE             = "Лекция"
	LESSON_TYPE_VALUE_SEMINAR             = "Семинар"
	LESSON_TYPE_VALUE_TEST                = "Зачёт"
	LESSON_TYPE_VALUE_CONSULTATION        = "Консультация"
	LESSON_TYPE_VALUE_EXAM                = "Экзамен"
)

var (
	LESSON_TYPE_KEYS = map[uint8]string{
		0: LESSON_TYPE_VALUE_UNKNOWN,
		1: LESSON_TYPE_VALUE_LECTURE,
		2: LESSON_TYPE_VALUE_SEMINAR,
		3: LESSON_TYPE_VALUE_CONSULTATION,
		4: LESSON_TYPE_VALUE_TEST,
		5: LESSON_TYPE_VALUE_EXAM,
	}

	LESSON_TYPE_VALUES = map[string]uint8{
		LESSON_TYPE_VALUE_UNKNOWN:      0,
		LESSON_TYPE_VALUE_LECTURE:      1,
		LESSON_TYPE_VALUE_SEMINAR:      2,
		LESSON_TYPE_VALUE_CONSULTATION: 3,
		LESSON_TYPE_VALUE_TEST:         4,
		LESSON_TYPE_VALUE_EXAM:         5,
	}
)

func (l LessonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

func (l *LessonType) UnmarshalJSON(data []byte) (err error) {
	var suits string
	if err := json.Unmarshal(data, &suits); err != nil {
		return err
	}

	if *l, err = ParseLessonType(suits); err != nil {
		return err
	}

	return nil
}

func (l LessonType) String() string {
	return LESSON_TYPE_KEYS[uint8(l)]
}

func ParseLessonType(s string) (LessonType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := LESSON_TYPE_VALUES[s]
	if !ok {
		return LessonType(0), fmt.Errorf("%q is not a valid LessonType", s)
	}
	return LessonType(value), nil
}

type WeekDay uint8

const (
	WEEKDAY_VALUE_UNKNOWN   string = "Неизвестно"
	WEEKDAY_VALUE_MONDAY           = "Понедельник"
	WEEKDAY_VALUE_TUESDAY          = "Вторник"
	WEEKDAY_VALUE_WEDNESDAY        = "Среда"
	WEEKDAY_VALUE_THURDSAY         = "Четверг"
	WEEKDAY_VALUE_FRIDAY           = "Пятница"
	WEEKDAY_VALUE_SATURDAY         = "Суббота"
	WEEKDAY_VALUE_SUNDAY           = "Воскресенье"
)

var (
	WEEKDAY_KEYS = map[uint8]string{
		0: WEEKDAY_VALUE_UNKNOWN,
		1: WEEKDAY_VALUE_MONDAY,
		2: WEEKDAY_VALUE_THURDSAY,
		3: WEEKDAY_VALUE_WEDNESDAY,
		4: WEEKDAY_VALUE_THURDSAY,
		5: WEEKDAY_VALUE_FRIDAY,
		6: WEEKDAY_VALUE_SATURDAY,
		7: WEEKDAY_VALUE_SUNDAY,
	}

	WEEKDAY_VALUES = map[string]uint8{
		WEEKDAY_VALUE_UNKNOWN:   0,
		WEEKDAY_VALUE_MONDAY:    1,
		WEEKDAY_VALUE_TUESDAY:   2,
		WEEKDAY_VALUE_WEDNESDAY: 3,
		WEEKDAY_VALUE_THURDSAY:  4,
		WEEKDAY_VALUE_FRIDAY:    5,
		WEEKDAY_VALUE_SATURDAY:  6,
		WEEKDAY_VALUE_SUNDAY:    7,
	}
)

func (w WeekDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String())
}

func (w *WeekDay) UnmarshalJSON(data []byte) (err error) {
	var weekday string
	if err := json.Unmarshal(data, &weekday); err != nil {
		return err
	}

	if *w, err = ParseWeekDay(weekday); err != nil {
		return err
	}

	return nil
}

func (w WeekDay) String() string {
	return WEEKDAY_KEYS[uint8(w)]
}

func ParseWeekDay(s string) (WeekDay, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := LESSON_TYPE_VALUES[s]
	if !ok {
		return WeekDay(0), fmt.Errorf("%q is not a valid WeekDay", s)
	}

	return WeekDay(value), nil
}

type ScheduleItem struct {
	ClassroomNumber string     `json:"classroom"`
	StartsAt        string     `json:"startsAt"`
	EndsAt          string     `json:"endsAt"`
	Building        string     `json:"building"`
	Lesson          string     `json:"lesson"`
	LessonType      LessonType `json:"lessonType"`
}

package dto

import (
	"fmt"
	"time"

	"github.com/Xenous-Inc/finapp-api/internal/clients/ruzfaclient/models"
)

type Date string

func (d Date) Validate() error {
	_, err := time.Parse("2006.01.02", string(d))
	if err != nil {
		return fmt.Errorf("Date Format is invalid")
	}

	return nil
}

type GetScheduleRequest struct {
	EntityId  string `json:"id" validate:"required"`
	StartDate Date   `json:"startDate" example:"2023.09.02"`
	EndDate   Date   `json:"endDate" example:"2023.09.11"`
} //@name GetScheduleRequest

type ScheduleItem struct {
	ClassroomNumber   string `json:"classroom" example:"ОД/341"`
	StartsAt          string `json:"startsAt" example:"10:10"`
	EndsAt            string `json:"endsAt" example:"11:40"`
	Address           string `json:"address" example:"ул. Олеко Дундича, 23"`
	Lesson            string `json:"lesson" example:"Иностранный язык"`
	LessonType        string `json:"lessonType" example:"Лекции"`
	LessonNumberStart uint8  `json:"lessonNumberStart" example:"1"`
	LessonNumberEnd   uint8  `json:"lessonNumberEnd" example:"1"`
	Lecturer          string `json:"lecturer" example:"Бердышев Александр Валентинович"`

	Date    string `json:"date" example:"2023.11.27"`
	WeekDay uint8  `json:"weekDay" example:"0"`
} //@name ScheduleItem

func ScheduleItemFromClientModel(m *models.Schedule) ScheduleItem {
	return ScheduleItem{
		ClassroomNumber:   m.Auditorium,
		StartsAt:          m.BeginLesson,
		EndsAt:            m.EndLesson,
		Address:           m.Building,
		Lesson:            m.Discipline,
		LessonType:        m.KindOfWork,
		LessonNumberStart: m.LessonNumberStart,
		LessonNumberEnd:   m.LessonNumberEnd,
		Lecturer:          m.Lecturer,
		Date:              m.Date,
		WeekDay:           m.DayOfWeek,
	}
}

type MiniScheduleItem struct {
	Date       string `json:"date" example:"2023.11.27"`
	LessonType string `json:"lessonType" example:"Лекции"`
} //@name MiniScheduleItem

func MiniScheduleItemFromClientModel(m *models.Schedule) MiniScheduleItem {
	return MiniScheduleItem{
		Date:       m.Date,
		LessonType: m.KindOfWork,
	}
}

// * Keep it for best times when SWAGGO begins enum support * //

/*
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
*/

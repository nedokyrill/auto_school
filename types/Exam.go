package types

import "time"

type ExamSubType int
type ExamType int

type ExamStore interface {
}

const (
	INNER = iota + 1
	GAI
)

const (
	THEORY ExamSubType = iota + 1
	SQUARE_DRIVING
	CITY_DRIVING
)

type Exam struct {
	ID       int         `json:"id"`
	Type     ExamType    `json:"type"`
	SubType  ExamSubType `json:"sub_type"`
	Date     time.Time   `json:"date"`
	ExamsCar Car         `json:"exams_car"`
}

type PassedExams struct {
	InnerTheoryPassed  bool `json:"inner_theory_passed"`
	InnerSquaredPassed bool `json:"inner_squared_passed"`
	InnerCityPassed    bool `json:"inner_city_passed"`
	GAITheoryPassed    bool `json:"gai_theory_passed"`
	GAICityPassed      bool `json:"gai_city_passed"`
}

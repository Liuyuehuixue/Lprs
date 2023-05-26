package models

type PlotInfo struct {
	TestN    int `json:"test_n" db:"test_n"`
	Age      int `json:"age" db:"age"`
	AllScore [8]Score
}

type Score struct {
	SurveyId    int `json:"survey_id"`
	PN          int `json:"p_n" db:"p_n"`
	AgeEvaluate int `json:"age_evaluate"`
}

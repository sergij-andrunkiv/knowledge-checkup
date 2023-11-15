package model

// структура для збереження питання і відповідей
type JSONPayload struct {
	TestTitle        string `json:"testTitle"`
	CountOfQuestions int    `json:"countOfQuestions"`
	MaxMark          int    `json:"maxMark"`
	Tags             string `json:"tags"`
	Question         string `json:"question"`
	QuestionType     string `json:"questionType"`
	Answers          []struct {
		Answer    string `json:"answer"`
		IsCorrect int    `json:"isCorrect"`
	} `json:"answers"`
}

type EditJSONPayload struct {
	QuestionsToDelete []int
	AnswersToDelete   []int
	Test              TestEntity
}

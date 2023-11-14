package model

import "database/sql"

type ANSWER_TYPE string

const (
	SINGLE_ANSWER    ANSWER_TYPE = "single"
	MULTIPLE_ANSWERS ANSWER_TYPE = "multiple"
)

// Ініціалізація питань
type QuestionEntity struct {
	ID            int
	Label         string
	Type          ANSWER_TYPE
	AnswerOptions []AnswerEntity
}

// Ініціалізація питання
func (q *QuestionEntity) Create(id int, label string, questionType ANSWER_TYPE) {
	q.ID = id
	q.Label = label
	q.Type = questionType
}

// Додати відповідь
func (q *QuestionEntity) AddAnswer(answer AnswerEntity) {
	q.AnswerOptions = append(q.AnswerOptions, answer)
}

// Зберегти питання (включно з відповідями)
func (q *QuestionEntity) Save(tx *sql.Tx, creatorId int, testId int64) error {
	// Якщо коректного ід немає, значить це нове питання
	if q.ID == -1 {
		return q.createNew(tx, creatorId, testId)
	}

	return q.update()
}

// Створення нового питання
func (q *QuestionEntity) createNew(tx *sql.Tx, creatorId int, testId int64) error {
	insert, err := tx.Exec("INSERT INTO questions(text, id_creator, type) VALUES(?, ?, ?)", q.Label, creatorId, q.Type)

	if err != nil {
		return err
	}

	questionId, err := insert.LastInsertId()

	if err != nil {
		return err
	}

	for _, answer := range q.AnswerOptions {
		answer.Save(tx, questionId, testId, creatorId)
	}

	return nil
}

func (q *QuestionEntity) update() error {
	return nil
}

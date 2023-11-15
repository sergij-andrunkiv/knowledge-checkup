package model

import (
	"database/sql"
	"knowledge_checkup/backend/dataStorage"
)

type ANSWER_TYPE string

const (
	SINGLE_ANSWER    ANSWER_TYPE = "single"
	MULTIPLE_ANSWERS ANSWER_TYPE = "multiple"
)

// Структура питання
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

	return q.update(tx, testId, creatorId)
}

// Отримати всі питання до відповідного тесту
func (q QuestionEntity) GetByTestID(testID int) ([]QuestionEntity, error) {
	var resultQuestions []QuestionEntity

	db := dataStorage.GetDB()
	defer db.Close()

	// Отримання зведеної інформації про питання та варіанти відповедей
	rows, err := db.Query("SELECT questions.id_q, questions.text, questions.id_creator, questions.type, answers.id_a, answers.text, answers.is_correct FROM `questions` LEFT JOIN tests_structure ON questions.id_q = tests_structure.id_q LEFT JOIN answers ON tests_structure.id_a = answers.id_a WHERE tests_structure.id_t = ?", testID)

	if err != nil {
		return resultQuestions, err
	}

	// Хештаблиця для тимчасового зберігання даних
	questionsSummary := make(map[int]QuestionEntity)

	for rows.Next() {
		var questionId int
		var questionText string
		var questionCreatorId int
		var questionType ANSWER_TYPE
		var answerId int
		var answerText string
		var answerIsCorrect bool

		rows.Scan(&questionId, &questionText, &questionCreatorId, &questionType, &answerId, &answerText, &answerIsCorrect)

		// В зв'язку з тим, що запит з JOIN'ами виконаний вище містить дублікати даних, які стосується питання, використовуємо хештаблицю для того, щоб не створити зайвих питань
		if _, ok := questionsSummary[questionId]; !ok {
			var question QuestionEntity
			question.Create(questionId, questionText, questionType)

			questionsSummary[questionId] = question
		}

		var answer AnswerEntity
		answer.Create(answerId, questionId, answerText, answerIsCorrect)

		// Вставлення варіантів відповіді в питання, оскільки структура всередені таблиці не адресується використовуєм цикл
		if question, ok := questionsSummary[questionId]; ok {
			question.AnswerOptions = append(question.AnswerOptions, answer)
			questionsSummary[questionId] = question
		}
	}

	// Конвертація таблиці в результуючий масив
	for _, v := range questionsSummary {
		resultQuestions = append(resultQuestions, v)
	}

	return resultQuestions, nil
}

// Видалити питання
func (q *QuestionEntity) Delete(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE from questions WHERE id_q = ?", q.ID)
	return err
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

// Оновлення існуючого питання
func (q *QuestionEntity) update(tx *sql.Tx, testId int64, creatorId int) error {
	_, err := tx.Exec("UPDATE questions SET text = ?, type = ? WHERE id_q = ?", q.Label, q.Type, q.ID)

	if err != nil {
		return err
	}

	for _, answer := range q.AnswerOptions {
		answer.Save(tx, int64(q.ID), testId, creatorId)
	}

	return nil
}

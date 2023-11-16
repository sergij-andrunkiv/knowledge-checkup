package model

import "database/sql"

type AnswerEntity struct {
	ID         int
	QuestionID int
	Label      string
	IsCorrect  bool
}

// Ініціалізація відповіді
func (a *AnswerEntity) Create(id int, questionId int, label string, isCorrect bool) {
	a.ID = id
	a.QuestionID = questionId
	a.Label = label
	a.IsCorrect = isCorrect
}

// Зберегти питання
func (a *AnswerEntity) Save(tx *sql.Tx, questionId int64, testId int64, creatorId int) error {
	// Якщо коректного ід немає, значить це нове питання
	if a.ID == -1 {
		return a.createNew(tx, questionId, testId, creatorId)
	}

	return a.update(tx)
}

// Зберегти відповідь в базу
func (a *AnswerEntity) createNew(tx *sql.Tx, questionId int64, testId int64, creatorId int) error {
	insert, err := tx.Exec("INSERT INTO answers(id_q, text, is_correct) VALUES(?, ?, ?)", questionId, a.Label, a.IsCorrect)

	if err != nil {
		return err
	}

	answerId, err := insert.LastInsertId()

	_, err = tx.Exec("INSERT INTO tests_structure(id_t, id_q, id_a, id_creator) VALUES(?, ?, ?, ?)", testId, questionId, answerId, creatorId)
	if err != nil {
		return err
	}

	return nil
}

// Оновити існуюче питання
func (a *AnswerEntity) update(tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE answers SET text = ?, is_correct = ? WHERE id_a = ?", a.Label, a.IsCorrect, a.ID)

	if err != nil {
		return err
	}

	return nil
}

package model

import (
	"database/sql"
	"knowledge_checkup/backend/dataStorage"
)

type TestEntity struct {
	ID            int
	CreatedAt     string
	UpdatedAt     string
	Title         string
	QuestionCount int
	MaxMark       int
	Tags          string
	CreatorID     int
	Questions     []QuestionEntity
}

// Ініціалізація тесту
func (t *TestEntity) Create(id int, createdAt string, updatedAt string, title string, questionCount int, maxMark int, tags string, creatorId int) {
	t.ID = id
	t.CreatedAt = createdAt
	t.UpdatedAt = updatedAt
	t.Title = title
	t.QuestionCount = questionCount
	t.MaxMark = maxMark
	t.Tags = tags
	t.CreatorID = creatorId
}

// Завантажити тест з бази даних за ід
func (t *TestEntity) LoadById(id int) error {
	db := dataStorage.GetDB()
	defer db.Close()
	db.QueryRow("SELECT id_t, title, created_at, updated_at, count_of_questions, max_mark, tags, creator FROM tests WHERE id_t = ?", id).Scan(&t.ID, &t.Title, &t.CreatedAt, &t.UpdatedAt, &t.QuestionCount, &t.MaxMark, &t.Tags, &t.CreatorID)

	var question QuestionEntity

	questions, err := question.GetByTestID(id)

	if err != nil {
		return err
	}

	t.Questions = questions

	return nil
}

// Додати питання
func (t *TestEntity) AddQuestion(question QuestionEntity) {
	t.Questions = append(t.Questions, question)
}

// Зберегти весь тест (включно з питаннямя та відповідями)
func (t *TestEntity) Save() error {
	// Якщо коректного ід немає, значить це новий тест
	if t.ID == -1 {
		return t.createNew()
	}

	return t.update()
}

// Отримати список тестів за авторством вчителя
func (t *TestEntity) GetListForTeacher(teacherId int) ([]TestEntity, error) {
	db := dataStorage.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT id_t, created_at, updated_at, title, count_of_questions, max_mark, tags, creator FROM tests WHERE creator = ?", teacherId)

	if err != nil {
		return nil, err
	}

	return t.getList(rows)
}

// Отримати список всіх тестів
func (t *TestEntity) GetList() ([]TestEntity, error) {
	db := dataStorage.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tests")

	if err != nil {
		return nil, err
	}

	return t.getList(rows)
}

// Згенерувати спсиок тестів на основі виконаного запиту
func (t *TestEntity) getList(rows *sql.Rows) ([]TestEntity, error) {
	var testList []TestEntity

	for rows.Next() {
		var testItem TestEntity
		if err := rows.Scan(&testItem.ID, &testItem.CreatedAt, &testItem.UpdatedAt, &testItem.Title, &testItem.QuestionCount, &testItem.MaxMark, &testItem.Tags, &testItem.CreatorID); err != nil {
			return nil, err
		}

		testList = append(testList, testItem)
	}

	return testList, nil
}

// Видалення питань та відповедей з тесту
func (t *TestEntity) HandleQuestionAndAnswerDeletion(questionIds []int, answerIds []int) error {
	if len(answerIds) == 0 && len(questionIds) == 0 {
		return nil
	}

	db := dataStorage.GetDB()
	defer db.Close()
	// Початок транзації
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	for _, qId := range questionIds {
		tx.Exec("DELETE FROM questions WHERE id_q = ?", qId)
	}

	for _, aId := range answerIds {
		tx.Exec("DELETE FROM answers WHERE id_a = ?", aId)
	}

	return tx.Commit()
}

// Видалити тест
func (t *TestEntity) Delete() error {
	db := dataStorage.GetDB()
	defer db.Close()

	// Початок транзації
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	for _, question := range t.Questions {
		err = question.Delete(tx)

		if err != nil {
			return err
		}
	}

	_, err = tx.Exec("DELETE FROM tests WHERE id_t = ?", t.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

// Перевірити тест
func (t *TestEntity) Submit(submission []SubmitTestJSONPayload) (float32, float32, error) {
	var result float32 = 0

	for _, questionSubmission := range submission {
		question := t.findQuestion(questionSubmission.QuestionId)
		result += question.CheckAnswers(questionSubmission.SelectedAnswersId)
	}

	return result, result * t.getQuestionValue(), nil
}

// Отримати к-сть балів, яке нараховуєтсья за одне питання
func (t *TestEntity) getQuestionValue() float32 {
	return float32(t.MaxMark / len(t.Questions))
}

// Знайти питання
func (t *TestEntity) findQuestion(qId int) QuestionEntity {
	for _, question := range t.Questions {
		if question.ID == qId {
			return question
		}
	}

	return QuestionEntity{}
}

// Створення нового тесту
func (t *TestEntity) createNew() error {
	db := dataStorage.GetDB()
	defer db.Close()

	// Початок транзації
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	insert, err := tx.Exec("INSERT INTO tests(title, count_of_questions, max_mark, tags, creator) VALUES(?, ?, ?, ?, ?)", t.Title, t.QuestionCount, t.MaxMark, t.Tags, t.CreatorID)

	if err != nil {
		return err
	}

	// Отримання ід тесту
	lastTestId, err := insert.LastInsertId()

	if err != nil {
		return err
	}

	// Збереження питань
	for _, question := range t.Questions {
		question.Save(tx, t.CreatorID, lastTestId)
	}

	// Збережння всього тесту, питань і відповідей однією транзакцією для забезпечення цілістності даних
	return tx.Commit()
}

// Оновити існуючий тест
func (t *TestEntity) update() error {
	db := dataStorage.GetDB()
	defer db.Close()

	// Початок транзації
	tx, err := db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE tests SET title = ?, count_of_questions = ?, max_mark = ?, tags = ? WHERE id_t = ?", t.Title, t.QuestionCount, t.MaxMark, t.Tags, t.ID)

	if err != nil {
		return err
	}

	for _, question := range t.Questions {
		question.Save(tx, t.CreatorID, int64(t.ID))
	}

	return tx.Commit()
}

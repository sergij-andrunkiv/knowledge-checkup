package model

import "knowledge_checkup/backend/dataStorage"

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

// Отримати список тестів
func (t *TestEntity) GetListForTeacher(teacherId int) ([]TestEntity, error) {
	var testList []TestEntity
	db := dataStorage.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT id_t, title, count_of_questions, max_mark, tags, creator FROM tests WHERE creator = ?", teacherId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var testItem TestEntity
		if err := rows.Scan(&testItem.ID, &testItem.Title, &testItem.QuestionCount, &testItem.MaxMark, &testItem.Tags, &testItem.CreatorID); err != nil {
			return nil, err
		}

		testList = append(testList, testItem)
	}

	return testList, nil
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
	tx.Commit()

	return nil
}

func (t *TestEntity) update() error {
	return nil
}

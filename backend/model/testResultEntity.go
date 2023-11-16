package model

import (
	"knowledge_checkup/backend/dataStorage"
	"net/http"
	"time"
)

type TestResultEntity struct {
	ID             int
	UserID         int
	TestID         int
	Mark           float32
	CorrectAnswers float32
	TimeTakenS     int64
	Test           TestEntity
}

// Ініціалізація структури
func (tr *TestResultEntity) Create(id int, userId int, testId int, mark float32, correctAnswers float32, timeTakenS int64) {
	tr.ID = id
	tr.UserID = userId
	tr.TestID = testId
	tr.Mark = mark
	tr.CorrectAnswers = correctAnswers
	tr.TimeTakenS = timeTakenS
}

// Зберегти інформацію про спробу проходження тесту
func (tr *TestResultEntity) Save(r *http.Request, w http.ResponseWriter) error {
	db := dataStorage.GetDB()
	defer db.Close()

	tr.calculateTimeTaken(r)

	insert, err := db.Query("INSERT INTO marks (user, test, mark, count_of_correct_answers, time_taken_s) VALUES (?, ?, ?, ?, ?)", tr.UserID, tr.TestID, tr.Mark, tr.CorrectAnswers, tr.TimeTakenS)

	defer insert.Close()
	return err
}

// Завнтажити всі результати корситувача
func (tr *TestResultEntity) GetUserResults(userId int) ([]TestResultEntity, error) {
	var results []TestResultEntity

	db := dataStorage.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM marks WHERE user = ? ORDER BY id_m DESC", userId)

	if err != nil {
		return results, err
	}

	for rows.Next() {
		var testResultItem TestResultEntity
		rows.Scan(&testResultItem.ID, &testResultItem.UserID, &testResultItem.TestID, &testResultItem.Mark, &testResultItem.CorrectAnswers, &testResultItem.TimeTakenS)
		testResultItem.loadTestInfo()
		results = append(results, testResultItem)
	}

	return results, nil
}

// Зберегти в сесії час початку спроби проходження тесту
func (tr TestResultEntity) SetStartAttemptTime(r *http.Request, w http.ResponseWriter) {
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")
	session.Values["startAttemptTime"] = time.Now().Unix()
	session.Save(r, w)
}

// Отримати час, який пройшов з початку проходження тесту
func (tr *TestResultEntity) calculateTimeTaken(r *http.Request) {
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")
	startTime := session.Values["startAttemptTime"].(int64)
	delete(session.Values, "startAttemptTime")

	tr.TimeTakenS = int64(time.Now().Sub(time.Unix(startTime, 0)).Seconds())
}

// Завантажити інформацію про тест, по відношенню до якого була зроблена спроба
func (tr *TestResultEntity) loadTestInfo() {
	var testData TestEntity
	testData.LoadById(tr.TestID)
	tr.Test = testData
}

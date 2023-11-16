package handlers

import (
	"encoding/json"
	"fmt"
	"knowledge_checkup/backend/model"
	"knowledge_checkup/backend/services"
	"net/http"
	"strconv"
)

// функція збереження питань і варіантів відповідей для тесту
func SaveTestQuestionsAnswersChanges(w http.ResponseWriter, r *http.Request) {
	var data []model.JSONPayload //пустий масив структур ПитанняВідповіді
	var msgMn services.MessageManager

	// обробка отриманих питань/відповідей у форматі JSON і збереженння у data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || len(data) == 0 {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалося зберегти тест")
		return
	}

	var testEntity model.TestEntity
	var currentUser model.Account

	currentUser.LoadFromSession(r)

	// Створення тесту
	testEntity.Create(-1, "", "", data[0].TestTitle, data[0].CountOfQuestions, data[0].MaxMark, data[0].Tags, currentUser.Id)

	// Наповнення тесту питаннями
	for _, question := range data {
		var currentQuestion model.QuestionEntity

		currentQuestion.Create(-1, question.Question, model.ANSWER_TYPE(question.QuestionType))

		// Створення питання відповідями
		for _, answer := range question.Answers {
			var currentAnswer model.AnswerEntity

			currentAnswer.Create(-1, -1, answer.Answer, answer.IsCorrect != 0)
			currentQuestion.AddAnswer(currentAnswer)
		}

		testEntity.AddQuestion(currentQuestion)
	}

	if testEntity.Save() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалося зберегти тест")
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Готово", "Тест збережено")
	w.WriteHeader(http.StatusOK)
}

// функція інформації про тест до клієнта (сторінка для вчителя)
func SendTestsInformationToClient(w http.ResponseWriter, r *http.Request) {
	var currentUser model.Account
	var testEntity model.TestEntity

	currentUser.LoadFromSession(r)

	tests, err := testEntity.GetListForTeacher(currentUser.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// передавання об'єкту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tests)
}

// функція інформації про тест до клієнта (весь спсиок)
func GetTestList(w http.ResponseWriter, r *http.Request) {
	var testEntity model.TestEntity

	tests, err := testEntity.GetList()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// передавання об'єкту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tests)
}

// Завантажити дані про тест для подальшого редагування
func GetTestToEdit(w http.ResponseWriter, r *http.Request) {
	var test model.TestEntity
	testId, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	test.LoadById(testId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(test)
}

// Зберегти відредагований тест
func SaveTest(w http.ResponseWriter, r *http.Request) {
	var msgMn services.MessageManager
	var editData model.EditJSONPayload

	err := json.NewDecoder(r.Body).Decode(&editData)

	if err != nil {
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Відбулась помилка при збереженні тесту")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Спочатку відбувається видалення питань та відповідей
	err = editData.Test.HandleQuestionAndAnswerDeletion(editData.QuestionsToDelete, editData.AnswersToDelete)

	if err != nil {
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Відбулась помилка при збереженні тесту")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Потім зберігаються зміни
	err = editData.Test.Save()

	if err != nil {
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Відбулась помилка при збереженні тесту")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Готово", "Тест збережено")
	w.WriteHeader(http.StatusOK)
}

// Видалити тест
func DeleteTest(w http.ResponseWriter, r *http.Request) {
	var msgMn services.MessageManager
	var test model.TestEntity
	testId, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалось видалити тест")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	test.LoadById(testId)

	err = test.Delete()

	if err != nil {
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалось видалити тест")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Готово", "Тест видалено")
	w.WriteHeader(http.StatusOK)
}

// Перевірити тест
func SubmitTest(w http.ResponseWriter, r *http.Request) {
	var msgMn services.MessageManager
	var test model.TestEntity
	var submission []model.SubmitTestJSONPayload
	var markResult model.TestResultEntity
	var userAccount model.Account

	userAccount.LoadFromSession(r)

	testId, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		fmt.Print(err)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "При перевірці тесту виникла помилка")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&submission)

	if err != nil {
		fmt.Print(err)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "При перевірці тесту виникла помилка")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	test.LoadById(testId)

	correctAnswers, resultMark, err := test.Submit(submission)

	if err != nil {
		fmt.Print(err)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "При перевірці тесту виникла помилка")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	markResult.Create(-1, userAccount.Id, testId, resultMark, correctAnswers, 0)
	err = markResult.Save(r, w)

	if err != nil {
		fmt.Print(err)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "При перевірці тесту виникла помилка")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Тест пройдено", "Тест успішно пройдено!")
	w.WriteHeader(http.StatusOK)
}

// Отримати список результатів
func GetTestResults(w http.ResponseWriter, r *http.Request) {
	var userAccount model.Account
	var testResults model.TestResultEntity

	userAccount.LoadFromSession(r)
	results, err := testResults.GetUserResults(userAccount.Id)

	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(results)
}

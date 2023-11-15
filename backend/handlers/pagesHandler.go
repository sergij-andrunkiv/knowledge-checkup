package handlers

import (
	"fmt"
	"knowledge_checkup/backend/dataStorage"
	"knowledge_checkup/backend/model"
	"knowledge_checkup/backend/view"
	"net/http"
	"strconv"
)

// Функція для обробки та відображення головної сторінки
func IndexPage(w http.ResponseWriter, r *http.Request) {
	var userAccount model.Account

	// Перевірка, чи користувач авторизований (перевірка наявності ключа "Last_name" у сесії)
	if userAccount.LoadFromSession(r) {
		if userAccount.Teacher_status == 1 {
			view.GetTpl().ExecuteTemplate(w, "index_teacher_page.html", userAccount)
		} else {
			view.GetTpl().ExecuteTemplate(w, "index_user_page.html", userAccount)
		}
	} else {
		view.GetTpl().ExecuteTemplate(w, "index_guest_page.html", nil)
	}
}

// Функція для виведення сторінки реєстрації
func RegistrationPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "registration_page.html", nil)
}

// Функція для виведення сторінки авторизації
func AuthorizationPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "authorization_page.html", nil)
}

// Функція для виведення сторінки тестів
func MyTestsPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "mytests_page.html", nil)
}

// Функція для виведення сторінки конструктора
func TestConstructorPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "testconstructor_page.html", nil)
}

// Функція для виведення сторінки помилки
func ErrorRedirectPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "error_redirect_page.html", nil)
}

// Функція для виведення сторінки аккаунта користувача
func AccountPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "account_page.html", nil)

	// Отримати дані користувача з сесії
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")
	last_name, ok := session.Values["last_name"].(string)
	if !ok {
		// Немає даних про користувача у сесії, редирект або обробка помилки
		http.Redirect(w, r, "/error_redirect", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "Ласкаво просимо, %s!", last_name)
}

// Функція для виведення сторінки списка тестів
func TestsListPage(w http.ResponseWriter, r *http.Request) {
	var userAccount model.Account
	userAccount.LoadFromSession(r)

	view.GetTpl().ExecuteTemplate(w, "testslist_page.html", userAccount)
}

// Сторінка редагування тесту
func EditTestPage(w http.ResponseWriter, r *http.Request) {
	view.GetTpl().ExecuteTemplate(w, "test_edit_page.html", nil)
}

// Сторінка проходження тесту
func TestCompletionPage(w http.ResponseWriter, r *http.Request) {
	var userAccount model.Account
	var test model.TestEntity
	var testAttempt model.TestResultEntity

	userAccount.LoadFromSession(r)
	testId, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		view.ErrorPage(w, "Не вдалося знайти тест")
		return
	}

	err = test.LoadById(testId)

	if err != nil {
		view.ErrorPage(w, "Не вдалося знайти тест")
		return
	}

	helperData := model.TestWithAccountHelper{
		UserAccount: userAccount,
		TestEntity:  test,
	}

	testAttempt.SetStartAttemptTime(r, w)

	view.GetTpl().ExecuteTemplate(w, "test_completion.html", helperData)
}

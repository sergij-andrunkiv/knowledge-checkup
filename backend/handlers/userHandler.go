package handlers

import (
	"fmt"
	"knowledge_checkup/backend/dataStorage"
	"knowledge_checkup/backend/model"
	"knowledge_checkup/backend/view"
	"net/http"
	"strconv"
)

// Функція для обробки запиту на реєстрацію
func HandleRegistration(w http.ResponseWriter, r *http.Request) {
	var userAccount model.Account

	// Отримуємо дані з форми
	last_name := r.FormValue("last_name")
	first_name := r.FormValue("first_name")
	middle_name := r.FormValue("middle_name")
	year_of_birth := r.FormValue("year_of_birth")
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repeat_password := r.FormValue("repeat_password")
	approved, _ := strconv.Atoi(r.FormValue("approved"))

	userAccount.Create(-1, first_name, last_name, middle_name, year_of_birth, nickname, email, password, approved, "", "", 0)

	if userAccount.IsAlreadyRegistered() {
		http.Error(w, "Цей email вже зареєстрований", http.StatusConflict)
		return
	}

	if userAccount.PasswordsDoNotMatch(repeat_password) {
		fmt.Fprintf(w, "Паролі не збігаються.")
		return
	}

	// Встановлення даних
	err := userAccount.Save()
	if err != nil {
		panic(err)
	}

	// При успішній реєстрації перенаправлення на сторінку авторизації
	http.Redirect(w, r, "/authorization", http.StatusSeeOther)
}

// Обробник виходу з акаунту
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := dataStorage.GetStore().Get(r, "user-data-session") // отримання токена сесії
	// Видалення інформації про користувача з сесії
	delete(session.Values, "id")
	delete(session.Values, "last_name")
	delete(session.Values, "first_name")
	delete(session.Values, "middle_name")
	delete(session.Values, "year_of_birth")
	delete(session.Values, "nickname")
	delete(session.Values, "email")
	delete(session.Values, "password")
	delete(session.Values, "approved")
	delete(session.Values, "gender")
	delete(session.Values, "educational_institutional")
	delete(session.Values, "teacher_status")
	session.Save(r, w)                            // збереження стану сесії
	http.Redirect(w, r, "/", http.StatusSeeOther) // перенаправлення на головну сторінку гостя
}

// Функція для обробки запиту на авторизацію
func HandleAuthorization(w http.ResponseWriter, r *http.Request) {
	var userAccount model.Account

	// Отримуємо дані з форми
	loginEmail := r.FormValue("loginEmail")
	loginPassword := r.FormValue("loginPassword")

	err := userAccount.LoadByAuth(loginEmail, loginPassword)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		view.ErrorPage(w, "Не вдалось авторизуватись. Перевірте правильність введених даних.")
		return
	}

	// Зберігаємо дані користувача у сесії
	if userAccount.SaveToSession(w, r) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		view.ErrorPage(w, "Не вдалось авторизуватись. Перевірте правильність введених даних.")
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther) // Перенаправлення на сторінку привітання після успішної авторизації
}

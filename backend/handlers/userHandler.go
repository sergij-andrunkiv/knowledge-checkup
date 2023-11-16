package handlers

import (
	"encoding/json"
	"fmt"
	"knowledge_checkup/backend/dataStorage"
	"knowledge_checkup/backend/model"
	"knowledge_checkup/backend/services"
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

// Оновити акаунт
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var currentUser model.Account
	var updatedUser model.Account
	var msgMn services.MessageManager

	currentUser.LoadFromSession(r)

	err := json.NewDecoder(r.Body).Decode(&updatedUser)

	if err != nil {
		fmt.Println(err)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалось зберегти зміни")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = currentUser.ChangeGeneralData(&updatedUser, w, r)

	if err != nil {
		fmt.Println(err)
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалось зберегти зміни")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Збережено", "Ваші дані успішно оновлено!")
	w.WriteHeader(http.StatusOK)
}

// Надіслати запит на підвищення повноважень
func SendPromotionRequest(w http.ResponseWriter, r *http.Request) {
	var teacherUser model.Account
	var currentUser model.Account
	var msgMn services.MessageManager
	currentUser.LoadFromSession(r)

	err := json.NewDecoder(r.Body).Decode(&teacherUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	if r.URL.Scheme == "" {
		r.URL.Scheme = "http"
	}

	confirmationLink := fmt.Sprintf("%s://%s/account/promotion_request/confirm?userId=%d", r.URL.Scheme, r.Host, currentUser.Id)

	messageBody := fmt.Sprintf("Користувач %s %s надіслав вам запит на підвищення повноважень. Перейдіть за <a href='%s' target='_blank'>посиланням</a>(%s) щоб підтвердити.", currentUser.First_name, currentUser.Last_name, confirmationLink, confirmationLink)

	err = services.SendEmail([]string{teacherUser.Email}, "Запит на підвищення повноважень", messageBody)

	if err != nil {
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", "Не вдалось надіслати запит")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Надіслано", "Запит на підвищення рівня доступу надіслано обраному вчителю!")
	w.WriteHeader(http.StatusOK)
}

// Підвищити користувача до вчителя
func PromoteUser(w http.ResponseWriter, r *http.Request) {
	var msgMn services.MessageManager
	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		view.ErrorPage(w, "Некоректний запит.")
		return
	}

	var promotedUser model.Account
	err = promotedUser.LoadById(userId)

	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		view.ErrorPage(w, "Відбулась внутрішня помилка.")
		return
	}

	promotedUser.Teacher_status = 1

	err = promotedUser.Save()

	if err != nil {
		fmt.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		view.ErrorPage(w, "Відбулась внутрішня помилка.")
		return
	}

	services.SendEmail([]string{promotedUser.Email}, "Ваш запит схвалено", "Ваш запит на підвищення прав доступу схвалено.")

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Опрацьовано", fmt.Sprintf("Рівень доступу користувача %s підвищено!", promotedUser.Nickname))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Змінити пароль
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var msgMn services.MessageManager
	var passwordData model.PasswordChangeJSONPayload
	var currentUser model.Account

	err := json.NewDecoder(r.Body).Decode(&passwordData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	currentUser.LoadFromSession(r)
	currentUser.LoadById(currentUser.Id)
	err, msg := currentUser.ChangePassword(&passwordData)

	if err != nil {
		fmt.Printf(err.Error())
		msgMn.Push(r, w, model.MESSAGE_STATUS_FAILURE, "Помилка", msg)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msgMn.Push(r, w, model.MESSAGE_STATUS_SUCCESS, "Готовно", "Ваш пароль успішно змінено!")
	w.WriteHeader(http.StatusOK)
}

// Отримати повідомлення
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var msgManager services.MessageManager
	w.Write([]byte(msgManager.Flush(r, w)))
}

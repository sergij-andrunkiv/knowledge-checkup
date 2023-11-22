package router

import (
	"fmt"
	"knowledge_checkup/backend/handlers"
	"knowledge_checkup/backend/model"
	"knowledge_checkup/backend/view"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// Middleware функція для перевірки HTTP методу та рівня доступу користувача перед запуском обробника
func configurableHadnler(handler HandlerFunc, method string, accessRestrictions int8) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userAccount model.Account
		var hasAccess bool
		methodAllowed := r.Method == method
		isAuth := userAccount.LoadFromSession(r)

		switch accessRestrictions {
		case model.GUEST:
			hasAccess = true
			break
		case model.USER:
			hasAccess = isAuth
			break
		case model.TEACHER:
			hasAccess = isAuth && userAccount.IsTeacher()
		}

		if methodAllowed && hasAccess {
			handler(w, r)
		} else if !methodAllowed {
			w.WriteHeader(http.StatusMethodNotAllowed)
			view.ErrorPage(w, fmt.Sprintf("Некоректний запит: %s не дозволено", r.Method))
		} else if !hasAccess {
			w.WriteHeader(http.StatusForbidden)
			view.ErrorPage(w, "Ви не вповноважені на виконання цієї операції")
		}
	}
}

// Налаштовуємо обробники (handlers) для різних URL-шляхів
func SetupRoutes() {
	// Завантаження сторінок
	http.HandleFunc("/", configurableHadnler(handlers.IndexPage, "GET", model.GUEST))                            // Головна сторінка
	http.HandleFunc("/registration", configurableHadnler(handlers.RegistrationPage, "GET", model.GUEST))         // Сторінка реєстрації
	http.HandleFunc("/authorization", configurableHadnler(handlers.AuthorizationPage, "GET", model.GUEST))       // Сторінка авторизації
	http.HandleFunc("/mytests", configurableHadnler(handlers.MyTestsPage, "GET", model.TEACHER))                 // Сторінка моїх тестів (вчителя)
	http.HandleFunc("/error_redirect", handlers.ErrorRedirectPage)                                               // Сторінка помилки
	http.HandleFunc("/account", configurableHadnler(handlers.AccountPage, "GET", model.USER))                    // Сторінка користувача
	http.HandleFunc("/testslist", configurableHadnler(handlers.TestsListPage, "GET", model.GUEST))               // Сторінка тестів
	http.HandleFunc("/testconstructor", configurableHadnler(handlers.TestConstructorPage, "GET", model.TEACHER)) // Сторінка конструктору тесту
	http.HandleFunc("/test/edit", configurableHadnler(handlers.EditTestPage, "GET", model.TEACHER))              // Сторінка редагування тесту
	http.HandleFunc("/test/start", configurableHadnler(handlers.TestCompletionPage, "GET", model.GUEST))         // Сторінка проходження тесту
	http.HandleFunc("/account/settings", configurableHadnler(handlers.AccountSettingsPage, "GET", model.USER))   // Сторінка налаштування профілю

	// Робота з акаунтом
	http.HandleFunc("/handleRegistration", configurableHadnler(handlers.HandleRegistration, "POST", model.GUEST))          // Обробник реєстрації
	http.HandleFunc("/handleLogout", configurableHadnler(handlers.HandleLogout, "GET", model.GUEST))                       // Обробник розлогінення
	http.HandleFunc("/handleAuthorization", configurableHadnler(handlers.HandleAuthorization, "POST", model.GUEST))        // Обробник авторизації
	http.HandleFunc("/account/update", configurableHadnler(handlers.UpdateAccount, "PATCH", model.GUEST))                  // Зберегти зміни в налаштуваннях профіля
	http.HandleFunc("/account/promotion_request", configurableHadnler(handlers.SendPromotionRequest, "POST", model.USER))  // Надіслати запит на підвищення повноважень
	http.HandleFunc("/account/promotion_request/confirm", configurableHadnler(handlers.PromoteUser, "GET", model.TEACHER)) // Підвищити користувача
	http.HandleFunc("/account/change_password", configurableHadnler(handlers.ChangePassword, "PATCH", model.USER))         // Змінити пароль
	http.HandleFunc("/account/get_messages", configurableHadnler(handlers.GetMessages, "GET", model.GUEST))                // Отримати системні повідомлення

	// Робота з тестами
	http.HandleFunc("/saveTestQuestionsAnswersChanges", configurableHadnler(handlers.SaveTestQuestionsAnswersChanges, "POST", model.TEACHER)) // Обробник збереження питань і варіантів відповідей в БД
	http.HandleFunc("/sendTestsInformationToClient", configurableHadnler(handlers.SendTestsInformationToClient, "GET", model.TEACHER))        // Отримання інформації про тести
	http.HandleFunc("/getTestToEdit", configurableHadnler(handlers.GetTestToEdit, "GET", model.TEACHER))                                      // Отримання інформації про тест для редагування
	http.HandleFunc("/saveEditedTest", configurableHadnler(handlers.SaveTest, "PUT", model.TEACHER))                                          // Зберегти тест після редагування
	http.HandleFunc("/test/delete", configurableHadnler(handlers.DeleteTest, "DELETE", model.TEACHER))                                        // Видалити тест
	http.HandleFunc("/test/list", configurableHadnler(handlers.GetTestList, "GET", model.GUEST))                                              // Отримати список всіх тестів
	http.HandleFunc("/test/submit", configurableHadnler(handlers.SubmitTest, "POST", model.USER))                                             // Перевірити результат проходження тесту
	http.HandleFunc("/test/results", configurableHadnler(handlers.GetTestResults, "GET", model.GUEST))                                        // Отримати результати тестів

	// Налаштування файлового серверу для статичних ресурсів
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server started on :4444") // Вивід повідомлення про запуск сервера на порту 4444
	http.ListenAndServe(":4444", nil)      // Запуск веб-сервера на порту 4444
}

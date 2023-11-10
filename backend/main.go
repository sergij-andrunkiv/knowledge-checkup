package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Account struct {
	Id int
	Last_name string
	First_name string
	Middle_name string
	Year_of_birth int
	Nickname string
	Email string
	Password string
	Approved int
	Gender string
	Educational_institution string
	Teacher_status int
}

type Test struct {
	Id_t int `json:"id_t"`
	Title string `json:"title"`
	Count_of_questions int `json:"count_of_questions"`
	Max_mark int `json:"max_mark"`
	Tags string `json:"tags"`
	Creator int `json:"creator"`
}

type FullTest struct {
	Id_t int `json:"id_t"`
	Title string `json:"title"`
	Count_of_questions int `json:"count_of_questions"`
	Max_mark int `json:"max_mark"`
	Tags string `json:"tags"`
	Creator int `json:"creator"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Id_q int `json:"id_q"`
	Text string `json:"text"`
	Id_creator int `json:"id_creator"`
	Type string `json:"type"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Id_a int `json:"id_a"`
	Id_q int `json:"id_q"`
	Text string `json:"text"`
	Is_correct int `json:"is_correct"`
}

// Створення об'єкту сховища сесій для зберігання сесій у формі куків на стороні клієнта
var store = sessions.NewCookieStore([]byte(":a#uX}h1.W91r~w:4YGU6B?`T4~>:>"))
// Глобальна змінна для зберігання підключення до бази даних
var db *sql.DB
var fullTest FullTest // глобаьлний екземпляр структури FullTest


// Функція для обробки та відображення головної сторінки
func indexPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-data-session")

	// Перевірка, чи користувач авторизований (перевірка наявності ключа "Last_name" у сесії)
	if _, ok := session.Values["email"]; ok {
		account := Account{
			session.Values["id"].(int),
			session.Values["last_name"].(string),
			session.Values["first_name"].(string),
			session.Values["middle_name"].(string),
			session.Values["year_of_birth"].(int),
			session.Values["nickname"].(string),
			session.Values["email"].(string),
			session.Values["password"].(string),
			session.Values["approved"].(int),
			"",
			"",
			session.Values["teacher_status"].(int),
		}
		// Перевірка, чи авторизований користувач є вчителем
		if session.Values["teacher_status"] == 1 {
			t, err := template.ParseFiles("templates/index_teacher_page.html")
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			t.Execute(w, account)
		} else {
			t, err := template.ParseFiles("templates/index_user_page.html")
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			t.Execute(w, account)
		}

	} else {
		t, err := template.ParseFiles("templates/index_guest_page.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		t.Execute(w, nil)
	}
}



// Функція для виведення сторінки реєстрації
func registrationPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/registration_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
}

// Функція для перевірки, чи існує email або nickname в базі даних
func emailExists(db *sql.DB, email_nickname string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM accounts WHERE email = ?", email_nickname).Scan(&count)
	if err != nil {
		err := db.QueryRow("SELECT COUNT(*) FROM accounts WHERE nickname = ?", email_nickname).Scan(&count)
		if err != nil {
			return false, err
		}
	}
	return count > 0, nil
}

// Функція для перевірки підтвердження паролю
func passwordsDoNotMatch(password string, repeat_password string) (bool) {
	if password != repeat_password {
		return true
	}
	return false
}


// Функція для обробки запиту на реєстрацію
func handleRegistration(w http.ResponseWriter, r *http.Request) {
	// Перенаправлення на сторінку помилки, якщо не POST-запит
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/error_redirect", http.StatusSeeOther)
		return
	}

	// Отримуємо дані з форми
	last_name := r.FormValue("last_name")
	first_name := r.FormValue("first_name")
	middle_name := r.FormValue("middle_name")
	year_of_birth := r.FormValue("year_of_birth")
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	repeat_password := r.FormValue("repeat_password")
	approved := r.FormValue("approved")

	// Отримання підключення до бази даних
	db, err := getDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Перевірка, чи існує email в базі даних
	emailExists, err := emailExists(db, email)
	if err != nil {
		panic(err) // обробка помилки
	}
	if emailExists {
		// Якщо email вже існує, ви можете виконати необхідні дії, наприклад, показати помилку користувачу і завершити функцію.
		http.Error(w, "Цей email вже зареєстрований", http.StatusConflict)
		return
	}

	passwordsDoNotMatch := passwordsDoNotMatch(password, repeat_password)
	if passwordsDoNotMatch {
		fmt.Fprintf(w, "Паролі не збігаються.")
		return
	}

	// Встановлення даних
	insert, err := db.Query(fmt.Sprintf("INSERT INTO accounts (last_name, first_name, middle_name, year_of_birth, nickname, email, password, approved, gender, educational_institution, teacher_status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', 'N/A', 'N/A', 0)", last_name, first_name, middle_name, year_of_birth, nickname, email, password, approved))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	// При успішній реєстрації перенаправлення на сторінку авторизації
	http.Redirect(w, r, "/authorization", http.StatusSeeOther)
}


// Функція для виведення сторінки авторизації
func authorizationPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/authorization_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
}


// Функція для обробки запиту на авторизацію
func handleAuthorization(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/error_redirect", http.StatusSeeOther)
		return
	}

	// Отримуємо дані з форми
	loginEmail := r.FormValue("loginEmail")
	loginPassword := r.FormValue("loginPassword")

	// Отримання підключення до бази даних
	db, err := getDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Перевірка, чи існує користувач з такими даними
	var storedPassword string
	err = db.QueryRow("SELECT password FROM accounts WHERE email = ?", loginEmail).Scan(&storedPassword)
	if err != nil {
		fmt.Println("Error:", err.Error()) // Виводимо помилку у випадку помилки запиту до бази даних
		http.Error(w, "Login failed", http.StatusUnauthorized) // Відправляємо помилку користувачу
		return
	}
	if loginPassword != storedPassword {
		http.Error(w, "Incorrect password", http.StatusUnauthorized) // Відправляємо помилку у випадку неправильного пароля
		return
	}

	// Отримуємо дані про користувача з бази даних
	var account Account
	err = db.QueryRow("SELECT id, last_name, first_name, middle_name, year_of_birth, nickname, email, password, approved, gender, educational_institution, teacher_status FROM accounts WHERE email = ?", loginEmail).Scan(&account.Id, &account.Last_name, &account.First_name, &account.Middle_name, &account.Year_of_birth, &account.Nickname, &account.Email, &account.Password, &account.Approved, &account.Gender, &account.Educational_institution, &account.Teacher_status)
	if err != nil {
		fmt.Println("Error:", err.Error())
		http.Error(w, "Login failed123", http.StatusUnauthorized)
		return
	}

	// Зберігаємо дані користувача у сесії
	session, _ := store.Get(r, "user-data-session")
	session.Values["id"] = account.Id
	session.Values["last_name"] = account.Last_name
	session.Values["first_name"] = account.First_name
	session.Values["middle_name"] = account.Middle_name
	session.Values["year_of_birth"] = account.Year_of_birth
	session.Values["nickname"] = account.Nickname
	session.Values["email"] = account.Email
	session.Values["password"] = account.Password
	session.Values["approved"] = account.Approved
	session.Values["gender"] = account.Gender
	session.Values["educational_institutional"] = account.Educational_institution
	session.Values["teacher_status"] = account.Teacher_status
	er := session.Save(r, w)

	if er != nil {
		http.Error(w, "Дані не додані в сесію", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther) // Перенаправлення на сторінку привітання після успішної авторизації
}


func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-data-session") // отримання токена сесії
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
	session.Save(r, w) // збереження стану сесії
	http.Redirect(w, r, "/", http.StatusSeeOther) // перенаправлення на головну сторінку гостя
}


func mytestsPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/mytests_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
}


func testconstructorPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/testconstructor_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
}


func errorRedirectPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/error_redirect_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
}


func accountPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/account_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)

	// Отримати дані користувача з сесії
	session, _ := store.Get(r, "user-data-session")
	last_name, ok := session.Values["last_name"].(string)
	if !ok {
		// Немає даних про користувача у сесії, редирект або обробка помилки
		http.Redirect(w, r, "/error_redirect", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "Ласкаво просимо, %s!", last_name)
}


func testslistPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/testslist_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, nil)
}


// функція збереження питань і варіантів відповідей для тесту
func saveTestQuestionsAnswersChanges(w http.ResponseWriter, r *http.Request) {
	// структура для збереження питання і відповідей
	type QuestionAnswer struct {
		TestTitle string `json:"testTitle"`
		CountOfQuestions int `json:"countOfQuestions"`
		MaxMark int `json:"maxMark"`
		Tags string `json:"tags"`
		Question string `json:"question"`
		QuestionType string `json:"questionType"`
		Answers  []struct {
			Answer    string `json:"answer"`
			IsCorrect int    `json:"isCorrect"`
		} `json:"answers"`
	}

	if r.Method == "POST" {
		var data []QuestionAnswer //пустий масив структур ПитанняВідповіді

		// обробка отриманих питань/відповідей у форматі JSON і збереженння у data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// підключення до бази даних
		db, err := getDBConnection()
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		// отримання потрібних даних користувача із сесії
		session, _ := store.Get(r, "user-data-session")
		val := session.Values["id"]
		id_user, ok := val.(int)
		if !ok {
			fmt.Println("Значення 'id' не є типу int")
			return
		}

		// Початок транзакції
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		// Вставка інформації про тест у таблицю tests
		insert, err := tx.Exec("INSERT INTO tests(title, count_of_questions, max_mark, tags, creator) VALUES(?, ?, ?, ?, ?)", data[0].TestTitle, data[0].CountOfQuestions, data[0].MaxMark, data[0].Tags, id_user)
		if err != nil {
			log.Fatal(err)
		}
		// Отримання автоінкрементованого ID останнього доданого тесту
		id_test, err := insert.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		// Коміт транзакції
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

		for _, qa := range data {
			// Початок транзакції
			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			// Вставка інформації про питання у таблицю questions
			insert, err := tx.Exec("INSERT INTO questions(text, id_creator, type) VALUES(?, ?, ?)", qa.Question, id_user, qa.QuestionType)
			if err != nil {
				log.Fatal(err)
			}
			// Отримання автоінкрементованого ID останнього доданого тесту
			id_question, err := insert.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}
			// Коміт транзакції
			err = tx.Commit()
			if err != nil {
				log.Fatal(err)
			}

			for _, ans := range qa.Answers {
				// Транзакція зберження варіанту відповіді
				tx, err := db.Begin()
				if err != nil {
					log.Fatal(err)
				}
				// Вставка інформації про варіанти відповідей у таблицю answers
				insert, err := tx.Exec("INSERT INTO answers(id_q, text, is_correct) VALUES(?, ?, ?)", id_question, ans.Answer, ans.IsCorrect)
				if err != nil {
					log.Fatal(err)
				}
				// Отримання автоінкрементованого ID останнього доданого тесту
				id_answer, err := insert.LastInsertId()
				if err != nil {
					log.Fatal(err)
				}
				err = tx.Commit()
				if err != nil {
					log.Fatal(err)
				}

				// Транзакція збереження запису структури тесту
				tx, err = db.Begin()
				if err != nil {
					log.Fatal(err)
				}
				// Вставка інформації про варіанти відповідей у таблицю answers
				_, err = tx.Exec("INSERT INTO tests_structure(id_t, id_q, id_a, id_creator) VALUES(?, ?, ?, ?)", id_test, id_question, id_answer, id_user)
				if err != nil {
					log.Fatal(err)
				}
				err = tx.Commit()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		w.WriteHeader(http.StatusOK) // надсилання клієнту статусу успішного виконання запиту

	} else {
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
	}
}



// функція інформації про тест до клієнта (сторінка для вчителя)
func sendTestsInformationToClient(w http.ResponseWriter, r *http.Request) {
	// Отримання підключення до бази даних
	db, err := getDBConnection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Отримуємо ID поточного юзера (вчителя) з сесії
	session, _ := store.Get(r, "user-data-session")
	creator_id := session.Values["id"]

	// Отримуємо дані про тест з таблиці tests
	rows, err := db.Query("SELECT id_t, title, count_of_questions, max_mark, tags, creator FROM tests WHERE creator = ?", creator_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tests []Test
	for rows.Next() {
		var test Test
		if err := rows.Scan(&test.Id_t, &test.Title, &test.Count_of_questions, &test.Max_mark, &test.Tags, &test.Creator); err != nil {
			log.Fatal(err)
		}
		newTest := Test {
			Id_t: test.Id_t,
			Title: test.Title,
			Count_of_questions: test.Count_of_questions,
			Max_mark: test.Max_mark,
			Tags: test.Tags,
			Creator: test.Creator,
		}
		tests = append(tests, newTest)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println("", test.Id_t)

	//// Зберігаємо дані користувача у сесії
	//test_data_session, _ := store.Get(r, "test-data-session")
	//test_data_session.Values["id_t"] = test.Id_t
	//test_data_session.Values["title"] = test.Title
	//test_data_session.Values["count_of_questions"] = test.Count_of_questions
	//test_data_session.Values["max_mark"] = test.Max_mark
	//test_data_session.Values["tags"] = test.Tags
	//test_data_session.Values["creator"] = test.Creator
	//er := test_data_session.Save(r, w)
	//if er != nil {
	//	http.Error(w, "Дані не додані в сесію", http.StatusUnauthorized)
	//	return
	//}
	// створити об'єкт структури Test для передавання на клієнт
	//data := Test {
	//	Id_t: test.Id_t,
	//	Title: test.Title,
	//	Count_of_questions: test.Count_of_questions,
	//	Max_mark: test.Max_mark,
	//	Tags: test.Tags,
	//	Creator: test.Creator,
	//}
	// передавання об'єкту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tests)
}



func existingtestconstructor(w http.ResponseWriter, r *http.Request) {
	// відображаємо сторінку конструктора існуючого тесту
	t, err := template.ParseFiles("templates/existingtestconstructor_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.Execute(w, map[string]interface{}{"Test": fullTest})

	// локальна структура для зберігання отриманих id тесту та власника від клієнта
	type TestIdAndCreator struct {
		TestId int `json:"id_test"`
		TestCreator  int    `json:"id_creator"`
	}

	// якщо від клієнта метод POST
	if r.Method == "POST" {
		// читаємо отримані дані, створюємо екземпляр структури та зберігаємо в неї отримані від клієнта дані (id тесту та власника)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			//http.Error(w, "Неможливо прочитати тіло запиту", http.StatusBadRequest)
			return
		}
		var testIdAndCreator TestIdAndCreator
		err = json.Unmarshal(body, &testIdAndCreator)
		if err != nil {
			http.Error(w, "Помилка розшифровки JSON", http.StatusBadRequest)
			return
		}

		// тимчасово виводимо отримані дані
		fmt.Println("Дані отримані на сервері:")
		fmt.Println(testIdAndCreator.TestId) // Виведення отриманих даних у консоль
		fmt.Println(testIdAndCreator.TestCreator)

		// Тут ви можете обробити дані, що отримані з клієнта
		// Тепер отримуємо дані про тест із бази даних і виводимо їх на екран для початку

		// Отримання підключення до бази даних
		db, err := getDBConnection()
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()


		// переініціалізація екземпляру структури FullTest і зберігання в неї результату запиту в БД (запитаний тест)
		fullTest = FullTest{}
		err = db.QueryRow("SELECT id_t, title, count_of_questions, max_mark, tags, creator FROM tests WHERE id_t = ? AND creator = ?", testIdAndCreator.TestId, testIdAndCreator.TestCreator).Scan(&fullTest.Id_t, &fullTest.Title, &fullTest.Count_of_questions, &fullTest.Max_mark, &fullTest.Tags, &fullTest.Creator)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID вибраного тесту", fullTest.Id_t)
		fmt.Println("Назва вибраного тесту", fullTest.Title)
		fmt.Println("Кількість питань вибраного тесту", fullTest.Count_of_questions)
		fmt.Println("Максимальна оцінка вибраного тесту", fullTest.Max_mark)
		fmt.Println("Теги вибраного тесту", fullTest.Tags)
		fmt.Println("Власник вибраного тесту", fullTest.Creator)

		idQuestionsRows, err := db.Query("SELECT id_q FROM tests_structure WHERE id_t = ?", fullTest.Id_t)
		if err != nil {
			log.Fatal(err)
		}
		defer idQuestionsRows.Close()
		var questionCounter = 0
		var countOfQuestions = fullTest.Count_of_questions // змінна для заглушки
		for idQuestionsRows.Next() {
			// заглушка для одноразового проходження по поточному питанню
			if(countOfQuestions < fullTest.Count_of_questions) {
				if(countOfQuestions == 0) {
					countOfQuestions = fullTest.Count_of_questions
				} else {
					countOfQuestions--
					continue
				}
			}
			var idQuestion int
			if err := idQuestionsRows.Scan(&idQuestion); err != nil {
				log.Fatal(err)
			}
			question := Question{}
			err = db.QueryRow("SELECT id_q, text, id_creator, type FROM questions WHERE id_q = ?", idQuestion).Scan(&question.Id_q, &question.Text, &question.Id_creator, &question.Type)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\tID питання: ", question.Id_q)
			fmt.Println("\tТекст: ", question.Text)
			fmt.Println("\tID креатора: ", question.Id_creator)
			fmt.Println("\tТип питання: ", question.Type)
			// Додати питання до зрізу питань у FullTest
			if questionCounter < len(fullTest.Questions) {
				fullTest.Questions[questionCounter] = question
			} else {
				// Якщо зріз не достатньо великий, використовуйте append
				fullTest.Questions = append(fullTest.Questions, question)
			}

			idAnswersRows, err := db.Query("SELECT id_a FROM tests_structure WHERE id_q = ?", question.Id_q)
			if err != nil {
				log.Fatal(err)
			}
			defer idAnswersRows.Close()
			var answerCounter = 0
			for idAnswersRows.Next() {
				var idAnswer int
				if err := idAnswersRows.Scan(&idAnswer); err != nil {
					log.Fatal(err)
				}
				answer := Answer{}
				err = db.QueryRow("SELECT id_a, id_q, text, is_correct FROM answers WHERE id_a = ?", idAnswer).Scan(&answer.Id_a, &answer.Id_q, &answer.Text, &answer.Is_correct)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("\t\tID варіанту відповіді: ", answer.Id_a)
				fmt.Println("\t\tID питання, до якого належить відповідь: ", answer.Id_q)
				fmt.Println("\t\tТекст варіанту відповіді: ", answer.Text)
				fmt.Println("\t\tПравильна відповідь: ", answer.Is_correct)
				// Додати варіант відповіді до зрізу варіантів відповідей у FullTest
				if answerCounter < len(fullTest.Questions[questionCounter].Answers) {
					fullTest.Questions[questionCounter].Answers[answerCounter] = answer
				} else {
					// Якщо зріз не достатньо великий, використовуйте append
					fullTest.Questions[questionCounter].Answers = append(fullTest.Questions[questionCounter].Answers, answer)
				}
				answerCounter++
			}
			countOfQuestions--
			questionCounter++
		}
		if err = idQuestionsRows.Err(); err != nil {
			log.Fatal(err)
		}

		// Відправка структури FullTest на клієнт у вигляді JSON
		fullTestJSON, err := json.Marshal(fullTest)
		if err != nil {
			http.Error(w, "Помилка кодування JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(fullTestJSON)
		return


	}
}




func getDBConnection() (*sql.DB, error) {
	// Підключення до бази даних MySQL
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/knowledge_checkup")
	if err != nil {
		panic(err.Error())
	}
	return db, err
}
func handleFunc() {
	// Налаштовуємо обробники (handlers) для різних URL-шляхів
	http.HandleFunc("/", indexPage) // Головна сторінка
	http.HandleFunc("/registration", registrationPage) // Сторінка реєстрації
	http.HandleFunc("/handleRegistration", handleRegistration) // Обробник реєстрації
	http.HandleFunc("/authorization", authorizationPage) // Сторінка авторизації
	http.HandleFunc("/handleAuthorization", handleAuthorization) // Обробник авторизації
	http.HandleFunc("/handleLogout", handleLogout) // Обробник розлогінення
	http.HandleFunc("/mytests", mytestsPage) // Сторінка моїх тестів (вчителя)
	http.HandleFunc("/testconstructor", testconstructorPage) // Сторінка конструктору тесту
	http.HandleFunc("/error_redirect", errorRedirectPage) // Сторінка помилки
	http.HandleFunc("/account", accountPage) // Сторінка користувача
	http.HandleFunc("/testslist", testslistPage)
	http.HandleFunc("/saveTestQuestionsAnswersChanges", saveTestQuestionsAnswersChanges) // Обробник збереження питань і варіантів відповідей в БД
	http.HandleFunc("/sendTestsInformationToClient", sendTestsInformationToClient)
	http.HandleFunc("/existingtestconstructor", existingtestconstructor)

	fmt.Println("Server started on :4444") // Вивід повідомлення про запуск сервера на порту 4444
	http.ListenAndServe(":4444", nil) // Запуск веб-сервера на порту 4444
}


func main() {
	getDBConnection()
	handleFunc()
}
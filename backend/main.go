package main

import (
	"knowledge_checkup/backend/router"
	"knowledge_checkup/backend/view"

	_ "github.com/go-sql-driver/mysql"
)

// var fullTest model.FullTest // глобаьлний екземпляр структури FullTest

// func existingtestconstructor(w http.ResponseWriter, r *http.Request) {
// 	// відображаємо сторінку конструктора існуючого тесту
// 	view.GetTpl().ExecuteTemplate(w, "existingtestconstructor_page.html", map[string]interface{}{"Test": fullTest})

// 	// локальна структура для зберігання отриманих id тесту та власника від клієнта
// 	type TestIdAndCreator struct {
// 		TestId      int `json:"id_test"`
// 		TestCreator int `json:"id_creator"`
// 	}

// 	// якщо від клієнта метод POST
// 	if r.Method == "POST" {
// 		// читаємо отримані дані, створюємо екземпляр структури та зберігаємо в неї отримані від клієнта дані (id тесту та власника)
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			//http.Error(w, "Неможливо прочитати тіло запиту", http.StatusBadRequest)
// 			return
// 		}
// 		var testIdAndCreator TestIdAndCreator
// 		err = json.Unmarshal(body, &testIdAndCreator)
// 		if err != nil {
// 			http.Error(w, "Помилка розшифровки JSON", http.StatusBadRequest)
// 			return
// 		}

// 		// тимчасово виводимо отримані дані
// 		fmt.Println("Дані отримані на сервері:")
// 		fmt.Println(testIdAndCreator.TestId) // Виведення отриманих даних у консоль
// 		fmt.Println(testIdAndCreator.TestCreator)

// 		// Тут ви можете обробити дані, що отримані з клієнта
// 		// Тепер отримуємо дані про тест із бази даних і виводимо їх на екран для початку

// 		// Отримання підключення до бази даних
// 		db := dataStorage.GetDB()
// 		defer db.Close()

// 		// переініціалізація екземпляру структури FullTest і зберігання в неї результату запиту в БД (запитаний тест)
// 		fullTest = model.FullTest{}
// 		err = db.QueryRow("SELECT id_t, title, count_of_questions, max_mark, tags, creator FROM tests WHERE id_t = ? AND creator = ?", testIdAndCreator.TestId, testIdAndCreator.TestCreator).Scan(&fullTest.Id_t, &fullTest.Title, &fullTest.Count_of_questions, &fullTest.Max_mark, &fullTest.Tags, &fullTest.Creator)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("ID вибраного тесту", fullTest.Id_t)
// 		fmt.Println("Назва вибраного тесту", fullTest.Title)
// 		fmt.Println("Кількість питань вибраного тесту", fullTest.Count_of_questions)
// 		fmt.Println("Максимальна оцінка вибраного тесту", fullTest.Max_mark)
// 		fmt.Println("Теги вибраного тесту", fullTest.Tags)
// 		fmt.Println("Власник вибраного тесту", fullTest.Creator)

// 		idQuestionsRows, err := db.Query("SELECT id_q FROM tests_structure WHERE id_t = ?", fullTest.Id_t)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer idQuestionsRows.Close()
// 		var questionCounter = 0
// 		var countOfQuestions = fullTest.Count_of_questions // змінна для заглушки
// 		for idQuestionsRows.Next() {
// 			// заглушка для одноразового проходження по поточному питанню
// 			if countOfQuestions < fullTest.Count_of_questions {
// 				if countOfQuestions == 0 {
// 					countOfQuestions = fullTest.Count_of_questions
// 				} else {
// 					countOfQuestions--
// 					continue
// 				}
// 			}
// 			var idQuestion int
// 			if err := idQuestionsRows.Scan(&idQuestion); err != nil {
// 				log.Fatal(err)
// 			}
// 			question := model.Question{}
// 			err = db.QueryRow("SELECT id_q, text, id_creator, type FROM questions WHERE id_q = ?", idQuestion).Scan(&question.Id_q, &question.Text, &question.Id_creator, &question.Type)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Println("\tID питання: ", question.Id_q)
// 			fmt.Println("\tТекст: ", question.Text)
// 			fmt.Println("\tID креатора: ", question.Id_creator)
// 			fmt.Println("\tТип питання: ", question.Type)
// 			// Додати питання до зрізу питань у FullTest
// 			if questionCounter < len(fullTest.Questions) {
// 				fullTest.Questions[questionCounter] = question
// 			} else {
// 				// Якщо зріз не достатньо великий, використовуйте append
// 				fullTest.Questions = append(fullTest.Questions, question)
// 			}

// 			idAnswersRows, err := db.Query("SELECT id_a FROM tests_structure WHERE id_q = ?", question.Id_q)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			defer idAnswersRows.Close()
// 			var answerCounter = 0
// 			for idAnswersRows.Next() {
// 				var idAnswer int
// 				if err := idAnswersRows.Scan(&idAnswer); err != nil {
// 					log.Fatal(err)
// 				}
// 				answer := model.Answer{}
// 				err = db.QueryRow("SELECT id_a, id_q, text, is_correct FROM answers WHERE id_a = ?", idAnswer).Scan(&answer.Id_a, &answer.Id_q, &answer.Text, &answer.Is_correct)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				fmt.Println("\t\tID варіанту відповіді: ", answer.Id_a)
// 				fmt.Println("\t\tID питання, до якого належить відповідь: ", answer.Id_q)
// 				fmt.Println("\t\tТекст варіанту відповіді: ", answer.Text)
// 				fmt.Println("\t\tПравильна відповідь: ", answer.Is_correct)
// 				// Додати варіант відповіді до зрізу варіантів відповідей у FullTest
// 				if answerCounter < len(fullTest.Questions[questionCounter].Answers) {
// 					fullTest.Questions[questionCounter].Answers[answerCounter] = answer
// 				} else {
// 					// Якщо зріз не достатньо великий, використовуйте append
// 					fullTest.Questions[questionCounter].Answers = append(fullTest.Questions[questionCounter].Answers, answer)
// 				}
// 				answerCounter++
// 			}
// 			countOfQuestions--
// 			questionCounter++
// 		}
// 		if err = idQuestionsRows.Err(); err != nil {
// 			log.Fatal(err)
// 		}

// 		// Відправка структури FullTest на клієнт у вигляді JSON
// 		fullTestJSON, err := json.Marshal(fullTest)
// 		if err != nil {
// 			http.Error(w, "Помилка кодування JSON", http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(fullTestJSON)
// 		return

// 	}
// }

func main() {
	view.GetTpl()        // Парсинг всіх шаблонів на початку роботи серверу
	router.SetupRoutes() // Налаштування маршрутів
}

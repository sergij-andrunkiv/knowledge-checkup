(() => {
	const queryString = window.location.search;
	const urlParams = new URLSearchParams(queryString);

	const testId = urlParams.get("id")

	const deletedQuestionIds = []
	const deletedAnswersIds = []

	if (!testId) {
		window.location.href = '/mytests'
	}

	let testTitleInput = null;
	let maxMarkInput = null;
	let tagsInput = null;
	let mainContainer = null;

	let originalState = null;
	let currentState = null;

	const SINGLE = 'single'
	const MULTIPLE = 'multiple'

	// Знайти питання
	const getQuestion = (uid) => {
		for (question of currentState.Questions) {
			if (question.tempUid == uid) {
				return question;
			}
		}
	}

	// Знайти відповідь
	const getAnswer = (uid) => {
		for (question of currentState.Questions) {
			for (answer of question.AnswerOptions) {
				if (answer.tempUid == uid) return answer;
			}
		}
	}

	// Функція для генерації унікального ідентифікатора
	const uuidv4 = () => {
		return "10000000-1000-4000-8000-100000000000".replace(/[018]/g, c =>
		  (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
		);
	  }

	// Додати відповідь
	const addAnswer = (questionUid, questionId) => {
		const answer = {
			ID: -1,
			QuestionId: parseInt(questionId),
			Label: "",
			IsCorrect: false,
			tempUid: uuidv4()
		}

		getQuestion(questionUid).AnswerOptions.push(answer);
		render(currentState);
	}

	// Змінити текст відповіді
	const changeAnswerText = (uid, newText) => {
		getAnswer(uid).Label = newText
	}

	// Видалити відповідь
	const deleteAnswer = (uid, questionUid, answerId) => {
		const question = getQuestion(questionUid);
		question.AnswerOptions = question.AnswerOptions.filter(option => option.tempUid != uid);

		if (answerId) {
			deletedAnswersIds.push(parseInt(answerId))
		}

		render(currentState);
	}

	// Обробка змін в правильних відповідях
	const processCorrectAnswer = (answerUid, questionUid, type, checked) => {
		if (type == "checkbox") {
			getAnswer(answerUid).IsCorrect = checked
		} else {
			getQuestion(questionUid).AnswerOptions.forEach(option => option.IsCorrect = false)
			getAnswer(answerUid).IsCorrect = checked
		}

		render(currentState);
	}

	// Додати питання
	const addQuestion = (type) => {
		const question = {
			ID: -1, // Нове питання
			Label: "",
			Type: type,
			AnswerOptions: [],
			tempUid: uuidv4()
		};

		currentState.Questions.push(question);

		render(currentState);
	}

	// Змінити текст питання
	const changeQuestionText = (uid, newText) => {
		getQuestion(uid).Label = newText
	}

	// Видалити питання
	const deleteQuestion = (uid, id) => {
		currentState.Questions = currentState.Questions.filter(question => question.tempUid != uid)

		if (id) {
			deletedQuestionIds.push(parseInt(id))
		}

		render(currentState)
	}

	// Change question type
	const changeQuestionType = (uid, type) => {
		const newType = type == SINGLE ? MULTIPLE : SINGLE
		getQuestion(uid).Type = newType
		render(currentState)
	}

	// Відобразити тест для редагування
	const render = (test) => {
		testTitleInput.value = test.Title;
		maxMarkInput.value = test.MaxMark;
		tagsInput.value = test.Tags;

		mainContainer.innerHTML = ""

		// Поки так, при потребі потім перепишем
		test.Questions && test.Questions.map(question => {
			let answers = '';

			question.AnswerOptions.map(answer => {
				const checked = answer.IsCorrect ? "checked" : ''
				answers += `
					<div class="answer">
						<label>Відповідь: </label>
						<input type="text" class='answer-text' data-answer-uid='${answer.tempUid}' value='${answer.Label}'>
						${
							question.Type == SINGLE 
							? `<input type="radio" class='answer-select' data-answer-uid=${answer.tempUid} data-question-uid=${question.tempUid} name="question${question.ID}" ${checked}>`
							: `<input type="checkbox" class='answer-select' data-answer-uid=${answer.tempUid} data-question-uid=${question.tempUid} name="question${question.ID}" ${checked}>`
						}
						<button data-answer-uid=${answer.tempUid} data-answer-id=${answer.ID} data-question-uid=${question.tempUid} class='delete-answer'>-</button>
					</div>
				`
			})

			mainContainer.innerHTML += `
			<div id="questionNumber${question.ID}" class="question" data-question-type="${question.Type}">
				<label>Питання: </label>
				<input type="text" class='question-text' data-question-uid=${question.tempUid} value="${question.Label}">
				<br>
				<button data-question-uid=${question.tempUid} data-question-id=${question.ID} class='add-answer' data-question-type="${question.Type}">
					Додати відповідь
				</button>
				<button class='delete-question' data-question-id=${question.ID} data-question-uid=${question.tempUid}>
					Видалити питання
				</button>
				<button class='change-question-type' data-question-type='${question.Type}' data-question-uid=${question.tempUid}>
					Змінити тип
				</button>
				${answers}
			</div>
			`
		})
	}

	const loadData = () => {
		// Завантаження даних про тест з сервера
		fetch(`/getTestToEdit?id=${testId}`)
		.then((response) => {
			if (response.status != 200) {
				alert("Критична помилка. Див. консоль для деталей");
				console.log(response);
				return;
			}
		return response.json()
		})
		.then((test) => {
			test.Questions.sort((a, b) => a.ID - b.ID) // Сортуємо питання по ід за зростанням
			originalState = test;
			currentState = JSON.parse(JSON.stringify(originalState));
			
			currentState.Questions.forEach(question => {
				question['tempUid'] = uuidv4();

				question.AnswerOptions.forEach(answer => {
					answer['tempUid'] = uuidv4();
				})
			});

			render(currentState)
		});
	}

	// Надіслати дані на сервер
	const sendData = () => {
		fetch('/saveEditedTest', {
			method: "PUT",
			cache: "no-cache",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify({
				QuestionsToDelete:deletedQuestionIds,
				AnswersToDelete:deletedAnswersIds,
				Test: currentState
			})
		}).then(response => {
			console.log(response) // відповідь
			loadData()
		})
	}

	document.addEventListener("DOMContentLoaded", () => {
		testTitleInput = document.querySelector("#testTitle");
		maxMarkInput = document.querySelector("#maxMark");
		tagsInput = document.querySelector("#tags");
		mainContainer = document.querySelector("#questions")

		loadData();
	})

	// Налаштування обробки подій для динамічно згенерованих елементів
	document.addEventListener("click", e => {
		if (e.target.classList.contains("add-answer")) {
			addAnswer(e.target.dataset.questionUid, e.target.dataset.questionId);
		}

		if (e.target.classList.contains("add-question-editor")) {
			addQuestion(SINGLE)
		}

		if (e.target.classList.contains("add-multi-answer-question-editor")) {
			addQuestion(MULTIPLE)
		}

		if (e.target.classList.contains("delete-question")) {
			deleteQuestion(e.target.dataset.questionUid, e.target.dataset.questionId);
		}

		if (e.target.classList.contains("change-question-type")) {
			changeQuestionType(e.target.dataset.questionUid, e.target.dataset.questionType)
		}

		if (e.target.classList.contains("delete-answer")) {
			deleteAnswer(e.target.dataset.answerUid,e.target.dataset.questionUid, e.target.dataset.answerId)
		}

		if (e.target.classList.contains("answer-select")) {
			processCorrectAnswer(e.target.dataset.answerUid, e.target.dataset.questionUid, e.target.type, e.target.checked)
		}

		if (e.target.classList.contains("submit-edited-changes")) {
			currentState.QuestionCount = currentState.Questions.length;
			sendData()
		}
	})

	document.addEventListener("keyup", e => {
		if (e.target.classList.contains('question-text')) {
			changeQuestionText(e.target.dataset.questionUid, e.target.value)
		}

		if (e.target.classList.contains('answer-text')) {
			changeAnswerText(e.target.dataset.answerUid, e.target.value)
		}

		if (e.target.name == "testTitle") {
			currentState.Title = e.target.value
		}

		if (e.target.name == "maxMark") {
			currentState.MaxMark = e.target.value
		}

		if (e.target.name == "tags") {
			currentState.Tags = e.target.value
		}
	})
})()
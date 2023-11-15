// перевірка незбережених змін при перезавантаженні
let unsavedChanges = false;

let questionCounter = 0; // лічильник питань

// функції додавання питань з одною і декількома варіантами відповідей
function addQuestion() {
    addNewQuestion(false);
}
function addMultiAnswerQuestion() {
    addNewQuestion(true);
}

// загальна функція додавання нового питання з варіантами відповідей
function addNewQuestion(hasMultipleAnswers) {
    unsavedChanges = true;

    const questionsDiv = document.getElementById('questions'); //отримання елементу контейнера всіх питань

    // створення контейнера для питання з відповідями
    const questionDiv = document.createElement('div');
    questionDiv.id = `questionNumber${questionCounter}`; // id контейнера з інкрементацією
    questionDiv.classList.add('question'); // додавання класу

    const questionLabel = document.createElement('label'); //створення мітки питання
    questionLabel.textContent = `Питання: `; //встановлення вмісту мітки
    const questionInput = document.createElement('input'); // створення поля для введення тексту питання
    questionInput.type = 'text'; //встановлення типу поля "text"
    questionInput.dataset['questionId'] = questionCounter; //створення атрибуту data-question-id у елемента questionInput і присвоєння йому номера питання
    if(hasMultipleAnswers) {  //створення атрибуту data-question-type у елемента questionDiv і присвоєння йому типу питання
        questionDiv.dataset['questionType'] = "multiple";
    }
    else {
        questionDiv.dataset['questionType'] = "single";
    }

    const deleteButton = document.createElement('button'); //створення елементу кнопки для видалення питання
    deleteButton.textContent = 'Видалити питання'; //встановлення вмісту кнопки
    deleteButton.onclick = function() { //функція опрацювання при натисненні на кнопку видалення питання
        //видалення всіх дочірніх елементів всередині поточного контейнера для якого натискається кнопка (контейнер з питанням/відповідями і лінія розмежування)
        questionsDiv.removeChild(questionDiv);
        questionsDiv.removeChild(hrElement);
    };

    const addAnswerButton = document.createElement('button'); //створення кнопки для додавання варіантів відповідей
    addAnswerButton.textContent = 'Додати відповідь'; //встановлення вмісту кнопки
    addAnswerButton.dataset['questionId'] = questionCounter; //створення атрибуту data-question-id у елемента addAnswerButton і присвоєння йому номера питання, до якого належить дана відповідь
    addAnswerButton.onclick = function(event) { //функція опрацювання при натисненні на кнопку додавання варіанту відповіді
        const answerDiv = document.createElement('div'); //створення контейнера для варіанту відповіді
        answerDiv.classList.add('answer'); //додавання атрибуту класу answer

        const answerLabel = document.createElement('label'); //створення мітки для варіанту віідповіді
        answerLabel.textContent = `Відповідь: `; //встановлення вмісту мітки
        const answerInput = document.createElement('input'); //створення поля для введення варіанту відповіді
        answerInput.type = 'text'; //встановлення типу text
        const answerCheckbox = document.createElement('input'); //створення елементу для відмічання правильної відповіді
        answerCheckbox.type = hasMultipleAnswers ? 'checkbox' : 'radio'; //встановлення типу checkbox/radio в залежності від вибраного типу питання
        const currentQuestionId = event.currentTarget.dataset['questionId'] //збереження вмісту атрибуту data-question-id кнопки addAnswerButton у змінній currentQuestionId
        answerCheckbox.name = `question${currentQuestionId}`; //встановлення атрибуту імені для створеного чекбоксу відповіді
        answerCheckbox.checked = true; //автоматичне встановлення активності чекбокса/радіокнопки

        const deleteAnswerButton = document.createElement('button'); //створення кнопки видалення відповіді
        deleteAnswerButton.textContent = 'Видалити відповідь'; //встановлення вмісту кнопки
        deleteAnswerButton.onclick = function() { //функція опрацювання кнопки видалення відповіді
            questionDiv.removeChild(answerDiv); //видалити поточний контейнер з відповіддю
        };

        answerDiv.appendChild(answerLabel); //відобразити на сторінці мітку відповіді
        answerDiv.appendChild(answerInput); //відобразити на сторінці поле для введення відповіді
        answerDiv.appendChild(answerCheckbox); //відобразити на сторінці checkbox/radio відповіді
        answerDiv.appendChild(deleteAnswerButton); //відобразити на сторінці кнопку видалення відповіді
        questionDiv.appendChild(answerDiv); //відобразити на сторінці контейнер для варіанту відповіді
    };

    questionDiv.appendChild(questionLabel);//відобразити на сторінці мітку питання
    questionDiv.appendChild(questionInput); //відобразити на сторінці поле для введення питання
    questionDiv.appendChild(document.createElement('br')); //відобразити перехід на новий рядок
    questionDiv.appendChild(addAnswerButton); //відобразити на сторінці кнопку додавання відповіді
    questionDiv.appendChild(deleteButton); //відобразити на сторінці кнопку видалення питання

    questionsDiv.appendChild(questionDiv); //відобразити на сторінці контейнер з питанням і відповідями
    const hrElement = document.createElement('hr');
    questionsDiv.appendChild(hrElement); //відобразити на сторінці лінію-розділювач між питаннями
    questionCounter++; //інкремент номера питання
}

function sendDataToServer() {
    const questions = document.querySelectorAll('.question'); //вибрати всі елементи з класом question (контейнери для кожного з питань)
    const data = []; //пустий масив даних для питань

    // отримання інформації про тест (назва тесту, макс. оцінка і теги)
    const testTitleComponent = document.getElementById("testTitle");
    const testTitle = testTitleComponent.value;
    const countOfQuestionsComponent = document.querySelectorAll('[id^="questionNumber"]');
    const countOfQuestions = countOfQuestionsComponent.length;
    const maxMarkComponent = document.getElementById("maxMark");
    const maxMark = parseInt(maxMarkComponent.value);
    const tagsComponent = document.getElementById("tags");
    const tags = tagsComponent.value;

    questions.forEach(questionDiv => { //пройтися по кожному елементу питання в масиві і виконати для нього код нижче
        const questionTextElement = questionDiv.querySelector('input[type="text"]'); //вибрати елемент поля питання із введеними даними
        if (questionTextElement) { //якщо цей елемент існує
            const questionText = questionTextElement.value; //зберегти значення з поля питання у змінній questionText
            const questionType = questionDiv.dataset.questionType;
            const answers = questionDiv.querySelectorAll('.answer'); //вибрати всі елементи з класом answer у контейнері questionDiv

            const answerData = []; //пустий масив даних для варіантів відповідей

            answers.forEach(answerDiv => { //пройтися по кожному елементу відповіді в масиві і виконати для нього код нижче
                const answerTextElement = answerDiv.querySelector('input[type="text"]'); //вибрати елемент поля відповіді із введеними даними
                const answerCheckbox = answerDiv.querySelector('input[type="checkbox"]'); //вибрати елемент checkbox
                const answerRadio = answerDiv.querySelector('input[type="radio"]'); //вибрати елемент radio

                const answerControl = (answerCheckbox || answerRadio); //змінна answerControl отримає одне із двох значень в залежності від типу питання

                if (answerTextElement && answerControl) { //якщо змінні не пусті
                    //ДАНІ ПРО ПИТАННЯ ДЛЯ ВІДПРАВКИ НА СЕРВЕР
                    //зберегти текст відповіді
                    const answerText = answerTextElement.value;
                    //зберегти наявність вибору відповіді (true/false)
                    let isCorrect;
                    if(answerControl.checked) {
                        isCorrect = 1;
                    }
                    else{
                        isCorrect = 0;
                    }

                    //помістити дані в масив
                    answerData.push({
                        answer: answerText,
                        isCorrect: isCorrect
                    });
                }
            });

            //помістити дані в масив
            data.push({
                testTitle: testTitle,
                countOfQuestions: countOfQuestions,
                maxMark: maxMark,
                tags: tags,
                question: questionText,
                questionType: questionType,
                answers: answerData
            });
        }
    });

    // Відправка питань і варіантів відповідей на сервер за допомогою AJAX
    fetch('/saveTestQuestionsAnswersChanges', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            // Обробка відповіді від сервера
            if (response.ok) {
                location.href = "/mytests" // редірект на сторінку всіх тестів
            } else {
                // Дії в разі помилки
            }
        })
        .catch(error => {
            // Обробка помилок під час відправки
        });
}

document.addEventListener("DOMContentLoaded", () => {
    document.querySelector(".add-question") && document.querySelector(".add-question").addEventListener("click", addQuestion);
    document.querySelector(".add-multi-answer-question") && document.querySelector(".add-multi-answer-question").addEventListener("click", addMultiAnswerQuestion);
    document.querySelector(".submit-changes") && document.querySelector(".submit-changes").addEventListener("click", sendDataToServer);
})
(() => {
    const queryString = window.location.search;
	const urlParams = new URLSearchParams(queryString);

	const testId = urlParams.get("id")

	if (!testId) {
		window.location.href = '/testslist'
	}

    // Надіслати результат проходження тесту
    const submitTest = () => {
        const result = []
        const questionElements = document.querySelectorAll(".question");

        for(question of questionElements) {
            const questionId = question.dataset.questionId

            const questionResult = {
                QuestionId: parseInt(questionId),
                SelectedAnswersId: []
            }

            const questionAnswers = document.querySelectorAll(`[name='question-${questionId}-answer']`)

            for(answer of questionAnswers) {
                if (answer.checked) {
                    questionResult.SelectedAnswersId.push(parseInt(answer.dataset.answerId))
                }
            }

            result.push(questionResult)
        }

        fetch(`/test/submit?id=${testId}`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(result)
        })
        .then(response => {
            if (response.ok) {
                window.location.href='/testslist'
            } else {
                alert("Критична помилка")
            }
        })
    }

    document.querySelector(".submit-test-button").addEventListener("click", submitTest)

}) ();
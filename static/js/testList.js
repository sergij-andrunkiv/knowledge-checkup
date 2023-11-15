( ()=> {
// Назви для JSON полів 
const namingMap = {
    "ID": "ID",
    "CreatedAt": "Дата створення",
    "UpdatedAt": "Дата оновлення",
    "Title": "Назва",
    "QuestionCount": "Кількість питань",
    "MaxMark": "Максимальний бал",
    "Tags": "Теги"
}

const fieldOrder = {
    "ID": 2,
    "CreatedAt": 6,
    "UpdatedAt": 7,
    "Title": 1,
    "QuestionCount": 3,
    "MaxMark": 4,
    "Tags": 5
}

// Видалення тесту
const deleteTest = id => {
    fetch(`/test/delete?id=${id}`, {
        method: "DELETE"
    }).then(response => {
        if (response.ok) {
            loadData();
        }
        console.log(response) // TODO: нормально обробити відповідь
    })
}

// Відображення списку тестів
renderData = (state) => {
    const container = document.querySelector("main");
    container.innerHTML = "";

    state && state.map(testItem => {
        const testInfoContainer = document.createElement('div');
        const constructorButton = document.createElement('button');
        const deleteButton = document.createElement('button');
        deleteButton.addEventListener("click", () => {
            deleteTest(testItem.ID)
        })
        deleteButton.textContent = "Видалити"
        const constructorLink = document.createElement('a');
        constructorButton.textContent = "Перейти в конструктор";
        constructorLink.append(constructorButton)
        constructorLink.href=`/test/edit?id=${testItem.ID}`
        testInfoContainer.classList.add("test-card-container")

        testInfoContainer.id = "test-container-" + testItem.ID;

        let testInfoList = document.createElement("ul")

        let keys = Object.keys(testItem).sort((a, b) => fieldOrder[a] - fieldOrder[b]);

        for (let key of keys) {
            if (!namingMap[key]) {
                continue;
            }

            const listItem = document.createElement("li")
            listItem.innerHTML = `<b>${namingMap[key]}</b>: ${testItem[key]}`
            testInfoList.appendChild(listItem)
        }

        testInfoContainer.appendChild(testInfoList)
        testInfoContainer.appendChild(constructorLink)
        testInfoContainer.appendChild(deleteButton)
        container.append(testInfoContainer)
    })
}

const loadData = function() {
    fetch('/sendTestsInformationToClient')
    .then((response) => {
        if (response.status != 200) {
            alert ("Критична помилка. Див. консоль для деталей")
            console.log(response)
            return;
        }

        return response.json();
    })
    .then((state) => {
        renderData(state)
    });
}

window.addEventListener("DOMContentLoaded", loadData);

}) ();
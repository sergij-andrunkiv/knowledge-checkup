const namingMap = {
    "ID": "ID",
    "CreatedAt": "Дата створення",
    "UpdatedAt": "Дата оновлення",
    "Title": "Назва",
    "QuestionCount": "Кількість питань",
    "MaxMark": "Максимальний бал",
    "Tags": "Теги"
}

renderData = (state) => {
    const container = document.querySelector("main");
    container.innerHTM = "";

    state.map(testItem => {
        const testInfoContainer = document.createElement('div');
        const constructorButton = document.createElement('button');
        const constructorLink = document.createElement('a');
        constructorButton.textContent = "Перейти в конструктор";
        constructorLink.append(constructorButton)
        constructorLink.href=`/test/edit?id=${testItem.ID}`
        testInfoContainer.classList.add("test-card-container")

        testInfoContainer.id = "test-container-" + testItem.ID;

        let testInfoList = document.createElement("ul")

        for (let key in testItem) {
            if (!namingMap[key]) {
                continue;
            }

            const listItem = document.createElement("li")
            listItem.innerHTML = `<b>${namingMap[key]}</b>: ${testItem[key]}`
            testInfoList.appendChild(listItem)
        }

        testInfoContainer.appendChild(testInfoList)
        testInfoContainer.appendChild(constructorLink)
        container.append(testInfoContainer)
    })
}

window.addEventListener("DOMContentLoaded", function() {
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
});
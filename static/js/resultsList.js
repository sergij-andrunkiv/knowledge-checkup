(()=>{
    const conatiner = document.querySelector(".test-results-container table");

    const render = data => {
        data && data.map(resultItem => {
            const tableRow = `
                <tr>
                    <td>${resultItem.Test.Title}</td>
                    <td>${(resultItem.TimeTakenS / 60).toFixed(2)} хв</td>
                    <td>${resultItem.Mark}</td>
                </tr>
            `

            conatiner.innerHTML += tableRow;
        })
    }

    fetch("/test/results")
    .then(response => {
        if (response.ok) {
            return response.json()
        }

        alert("Відбулась помилка")
        console.log(response)
    })
    .then(data => {
        render(data)
    })
}) ()
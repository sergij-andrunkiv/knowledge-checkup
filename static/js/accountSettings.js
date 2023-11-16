(() => {
    const getData = (id) => {
        return document.querySelector(`#${id}`).value
    }

    // Зберегти зміни
    const updateGeneralData = () => {
        const accountData = {
            Last_name: getData("last-name"),
            First_name: getData("first-name"),
            Middle_name: getData("middle-name"),
            Year_of_birth: getData("year-of-birth"),
            Nickname: getData("nickname"),
            Gender: getData("gender"),
            Educational_institution: getData("education-institution")
        }

        fetch("/account/update", {
            method: "PATCH",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(accountData)
        }).then(response => {
            if (!response.ok) {
                alert("Виникла помилка")
                console.log(response)
            }
        })
    }

    // Надіслати запит про підвищення прав
    const sendPromotionRequest = () => {
        const teacher = {
            Email: getData("teacher")
        }

        fetch("/account/promotion_request", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(teacher)
        }).then(response => {
            if (!response.ok) {
                alert("Виникла помилка")
                console.log(response)
            }
        })
    }

    // Змінити пароль
    const changePassword = () => {
        const passwordData = {
            OldPassword: getData("old-password"),
            NewPassword: getData("new-password"),
            NewPasswordRepeat: getData("new-password-repeat"),
        }

        fetch("/account/change_password", {
            method: "PATCH",
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(passwordData)
        }).then(response => {
            if (!response.ok) {
                alert("Виникла помилка")
                console.log(response)
            }
        })
    }

    document.querySelector(".save-general-settings").addEventListener("click", updateGeneralData)
    document.querySelector(".send-request").addEventListener("click", sendPromotionRequest)
    document.querySelector(".change-password").addEventListener("click", changePassword)
}) ()
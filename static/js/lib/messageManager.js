const checkMessages = () => {
    fetch("/account/get_messages")
    .then(response => {
        if (response.ok) {
            return response.json()
        }
    })
    .then(messages => {
        (messages[0]) && messages.map(message => {
            const messageContainer = document.createElement('div')
            messageContainer.classList.add(`message-${message.Status}`)
            messageContainer.innerHTML = `<h4>${message.Title}</h4><p>${message.Text}</p>`

            document.querySelector("header").appendChild(messageContainer)

            setTimeout(()=>{
                messageContainer.remove()
            }, 10000)
        })
    })
}

document.addEventListener("DOMContentLoaded", checkMessages)
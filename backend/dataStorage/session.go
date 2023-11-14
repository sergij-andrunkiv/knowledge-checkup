package dataStorage

import "github.com/gorilla/sessions"

var store *sessions.CookieStore = nil

const SESSION_KEY = ":a#uX}h1.W91r~w:4YGU6B?`T4~>:>"

// Створення об'єкту сховища сесій для зберігання сесій у формі куків на стороні клієнта
func GetStore() *sessions.CookieStore {
	if store == nil {
		store = sessions.NewCookieStore([]byte(SESSION_KEY))
	}

	return store
}

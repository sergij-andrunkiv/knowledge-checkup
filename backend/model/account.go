package model

import (
	"fmt"
	"knowledge_checkup/backend/dataStorage"
	"net/http"
)

const GUEST = 0
const USER = 1
const TEACHER = 2

type Account struct {
	Id                      int
	Last_name               string
	First_name              string
	Middle_name             string
	Year_of_birth           string
	Nickname                string
	Email                   string
	Password                string
	Approved                int
	Gender                  string
	Educational_institution string
	Teacher_status          int
}

// Створити екземпляр акаунта користувача
func (a *Account) Create(id int, firstname string, lastname string, middlename string, yob string, nickname string, email string, password string, approved int, gender string, institution string, isTeacher int) {
	a.Id = id
	a.First_name = firstname
	a.Last_name = lastname
	a.Middle_name = middlename
	a.Year_of_birth = yob
	a.Nickname = nickname
	a.Email = email
	a.Password = password
	a.Approved = approved
	a.Gender = gender
	a.Educational_institution = institution
	a.Teacher_status = isTeacher
}

// Перевірити чи користувач з таким емейлом ще не зареєстрований
func (a *Account) IsAlreadyRegistered() bool {
	db := dataStorage.GetDB()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM accounts WHERE email = ?", a.Email).Scan(&count)
	if err != nil {
		err := db.QueryRow("SELECT COUNT(*) FROM accounts WHERE nickname = ?", a.Nickname).Scan(&count)
		if err != nil {
			return false
		}
	}
	return count > 0
}

// Перевірити коректність повтору введерного паролю
func (a *Account) PasswordsDoNotMatch(repeatPassword string) bool {
	return a.Password != repeatPassword
}

// Зберегти корситувача в базу даних
func (a *Account) Save() error {
	db := dataStorage.GetDB()
	defer db.Close()
	insert, err := db.Query(fmt.Sprintf("INSERT INTO accounts (last_name, first_name, middle_name, year_of_birth, nickname, email, password, approved, gender, educational_institution, teacher_status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d', 'N/A', 'N/A', 0)", a.Last_name, a.First_name, a.Middle_name, a.Year_of_birth, a.Nickname, a.Email, a.preperePassword(a.Password), a.Approved))
	defer insert.Close()

	return err
}

// Завантажити дані користувача з сесії
func (a *Account) LoadFromSession(r *http.Request) bool {
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")

	if _, ok := session.Values["email"]; ok {
		a.Id = session.Values["id"].(int)
		a.Last_name = session.Values["last_name"].(string)
		a.First_name = session.Values["first_name"].(string)
		a.Middle_name = session.Values["middle_name"].(string)
		a.Year_of_birth = session.Values["year_of_birth"].(string)
		a.Nickname = session.Values["nickname"].(string)
		a.Email = session.Values["email"].(string)
		a.Password = session.Values["password"].(string)
		a.Approved = session.Values["approved"].(int)
		a.Gender = ""
		a.Educational_institution = ""
		a.Teacher_status = session.Values["teacher_status"].(int)
		return true
	} else {
		return false
	}
}

// Авторизуватись в акаунт користувача
func (a *Account) LoadByAuth(email string, password string) error {
	db := dataStorage.GetDB()
	defer db.Close()

	return db.QueryRow("SELECT id, last_name, first_name, middle_name, year_of_birth, nickname, email, approved, gender, educational_institution, teacher_status FROM accounts WHERE email = ? AND password = ?", email, a.preperePassword(password)).Scan(&a.Id, &a.Last_name, &a.First_name, &a.Middle_name, &a.Year_of_birth, &a.Nickname, &a.Email, &a.Approved, &a.Gender, &a.Educational_institution, &a.Teacher_status)
}

// Зберігаємо дані користувача у сесії
func (a *Account) SaveToSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := dataStorage.GetStore().Get(r, "user-data-session")

	session.Values["id"] = a.Id
	session.Values["last_name"] = a.Last_name
	session.Values["first_name"] = a.First_name
	session.Values["middle_name"] = a.Middle_name
	session.Values["year_of_birth"] = a.Year_of_birth
	session.Values["nickname"] = a.Nickname
	session.Values["email"] = a.Email
	session.Values["password"] = a.Password
	session.Values["approved"] = a.Approved
	session.Values["gender"] = a.Gender
	session.Values["educational_institutional"] = a.Educational_institution
	session.Values["teacher_status"] = a.Teacher_status

	return session.Save(r, w)
}

// Перевірити, чи є в користувача права вчителя
func (a *Account) IsTeacher() bool {
	return a.Teacher_status == 1
}

// TODO: хешування паролю?
func (a *Account) preperePassword(password string) string {
	return password
}

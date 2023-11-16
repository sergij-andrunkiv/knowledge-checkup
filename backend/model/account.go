package model

import (
	"errors"
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

	if a.Id == -1 {
		insert, err := db.Query(fmt.Sprintf("INSERT INTO accounts (last_name, first_name, middle_name, year_of_birth, nickname, email, password, approved, gender, educational_institution, teacher_status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d', 'N/A', 'N/A', 0)", a.Last_name, a.First_name, a.Middle_name, a.Year_of_birth, a.Nickname, a.Email, a.preperePassword(a.Password), a.Approved))
		defer insert.Close()
		return err
	}

	update, err := db.Query("UPDATE accounts SET last_name = ?, first_name = ?, middle_name = ?, year_of_birth = ?, nickname = ?, email = ?, password = ?, approved = ?, gender = ?, educational_institution = ?, teacher_status = ? WHERE id = ?", a.Last_name, a.First_name, a.Middle_name, a.Year_of_birth, a.Nickname, a.Email, a.Password, a.Approved, a.Gender, a.Educational_institution, a.Teacher_status, a.Id)

	if err != nil {
		return err
	}

	update.Close()
	return nil
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
		a.Gender = session.Values["gender"].(string)
		a.Educational_institution = session.Values["educational_institutional"].(string)
		a.Teacher_status = session.Values["teacher_status"].(int)
		return true
	} else {
		return false
	}
}

// Змінити пароль
func (a *Account) ChangePassword(passwordData *PasswordChangeJSONPayload) (error, string) {

	if passwordData.NewPassword == "" || passwordData.OldPassword == "" || passwordData.NewPasswordRepeat == "" {
		return errors.New("Заповніть всі поля"), "Заповніть всі поля"
	}

	if a.PasswordsDoNotMatch(a.preperePassword(passwordData.OldPassword)) {
		return errors.New("Неправильний старий пароль"), "Неправильний старий пароль"
	}

	a.Password = a.preperePassword(passwordData.NewPassword)

	if a.PasswordsDoNotMatch(a.preperePassword(passwordData.NewPasswordRepeat)) {
		return errors.New("Новий пароль та повтор нового паролю не збігаються"), "Новий пароль та повтор нового паролю не збігаються"
	}

	return a.Save(), ""
}

// Авторизуватись в акаунт користувача
func (a *Account) LoadByAuth(email string, password string) error {
	db := dataStorage.GetDB()
	defer db.Close()

	return db.QueryRow("SELECT id, last_name, first_name, middle_name, year_of_birth, nickname, email, approved, gender, educational_institution, teacher_status FROM accounts WHERE email = ? AND password = ?", email, a.preperePassword(password)).Scan(&a.Id, &a.Last_name, &a.First_name, &a.Middle_name, &a.Year_of_birth, &a.Nickname, &a.Email, &a.Approved, &a.Gender, &a.Educational_institution, &a.Teacher_status)
}

// Отримати користувача за ID
func (a *Account) LoadById(id int) error {
	db := dataStorage.GetDB()
	defer db.Close()

	return db.QueryRow("SELECT id, last_name, first_name, middle_name, year_of_birth, nickname, email, approved, gender, educational_institution, teacher_status, password FROM accounts WHERE id = ?", id).Scan(&a.Id, &a.Last_name, &a.First_name, &a.Middle_name, &a.Year_of_birth, &a.Nickname, &a.Email, &a.Approved, &a.Gender, &a.Educational_institution, &a.Teacher_status, &a.Password)
}

// Оновити дані
func (a *Account) ChangeGeneralData(diff *Account, w http.ResponseWriter, r *http.Request) error {
	if diff.Last_name != a.Last_name {
		a.Last_name = diff.Last_name
	}

	if diff.First_name != a.First_name {
		a.First_name = diff.First_name
	}

	if diff.Middle_name != a.Middle_name {
		a.Middle_name = diff.Middle_name
	}

	if diff.Year_of_birth != a.Year_of_birth {
		a.Year_of_birth = diff.Year_of_birth
	}

	if diff.Nickname != a.Nickname {
		a.Nickname = diff.Nickname
	}

	if diff.Gender != a.Gender {
		a.Gender = diff.Gender
	}

	if diff.Educational_institution != a.Educational_institution {
		a.Educational_institution = diff.Educational_institution
	}

	err := a.Save()

	if err != nil {
		return err
	}

	a.SaveToSession(w, r)

	return nil
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

// Отримати список вчителів
func (a Account) GetTeachers() ([]Account, error) {
	var teachers []Account

	db := dataStorage.GetDB()
	defer db.Close()

	rows, err := db.Query("SELECT first_name, last_name, email FROM accounts WHERE teacher_status = ?", 1)

	if err != nil {
		return teachers, err
	}

	for rows.Next() {
		var teacher Account
		err := rows.Scan(&teacher.First_name, &teacher.Last_name, &teacher.Email)

		if err != nil {
			return teachers, err
		}

		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

// TODO: хешування паролю?
func (a *Account) preperePassword(password string) string {
	return password
}

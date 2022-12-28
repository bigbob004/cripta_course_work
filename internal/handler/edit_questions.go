package handler

import (
	"cripta_course_work/internal/model"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

type ViewDataEditQuestions struct {
	EditingQuestionID int
	ModelQuestions    []model.Question
	Err               string
	IsThereError      bool
}

func (h *Handler) EditQuestionsView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		var isThereError bool
		errDeleting := r.FormValue("err")
		if len(errDeleting) != 0 {
			isThereError = true
		}
		editingQuestionID, _ := strconv.Atoi(r.FormValue("question_id"))
		funcMap := template.FuncMap{
			// The name "inc" is what the function will be called in the template text.
			"inc": func(i int) int {
				return i + 1
			},
		}
		name := path.Base("./template/question_redactor_new.html")
		tmpl, err := template.New(name).Funcs(funcMap).ParseFiles("./template/question_redactor_new.html")
		if err != nil {
			logrus.Error(err)
		}

		if h.cache.ChoosenUser == nil {
			err = tmpl.Execute(w, ViewDataEditQuestions{ModelQuestions: h.cache.ModelQuestion, EditingQuestionID: editingQuestionID, Err: errDeleting, IsThereError: isThereError})
		} else {
			err = tmpl.Execute(w, ViewDataEditQuestions{ModelQuestions: h.cache.ChoosenUser.Questions, EditingQuestionID: editingQuestionID, Err: errDeleting, IsThereError: isThereError})
		}
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (h *Handler) EditQuestion(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		//TODO сделать отдельную функцию сбора данных
		//TODO: маппинг и валидация
		editingQuestionID, _ := strconv.Atoi(r.FormValue("question_id"))
		funcMap := template.FuncMap{
			// The name "inc" is what the function will be called in the template text.
			"inc": func(i int) int {
				return i + 1
			},
		}
		name := path.Base("./template/question_redactor_new.html")
		tmpl, err := template.New(name).Funcs(funcMap).ParseFiles("./template/question_redactor_new.html")
		if err != nil {
			logrus.Error(err)
		}
		if h.cache.ChoosenUser == nil {
			err = tmpl.Execute(w, ViewDataEditQuestions{ModelQuestions: h.cache.ModelQuestion, EditingQuestionID: editingQuestionID})
		} else {
			err = tmpl.Execute(w, ViewDataEditQuestions{ModelQuestions: h.cache.ChoosenUser.Questions, EditingQuestionID: editingQuestionID})
		}
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (h *Handler) EditQuestionWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			logrus.Error("handler/EditQuestionWithID: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		//TODO сделать отдельную функцию сбора данных
		//TODO: маппинг и валидация
		var cacheOfCurrentUser []model.Question
		if h.cache.ChoosenUser == nil {
			cacheOfCurrentUser = h.cache.ModelQuestion
		} else {
			cacheOfCurrentUser = h.cache.ChoosenUser.Questions
		}

		editingQuestionID, _ := strconv.Atoi(r.FormValue("question_id"))
		newTitele := r.FormValue("question")
		newAnswer := r.FormValue("answer")
		newData := model.Question{
			QuestionID: editingQuestionID,
			Title:      newTitele,
			Answer:     newAnswer,
		}
		updateCache(cacheOfCurrentUser, newData)
		err := h.services.UpdateQuestion(newData)
		if err != nil {
			logrus.Error("handler/EditQuestionWithID: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}

		http.Redirect(w, r, "/edit_questions", http.StatusFound)
	}
}

func (h *Handler) DeleteQuestionWithID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			logrus.Error("handler/DeleteQuestionWithID: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}

		deletingQuestionID, _ := strconv.Atoi(r.FormValue("question_id"))
		var errDeleting string
		if h.cache.ChoosenUser == nil {
			if len(h.cache.ModelQuestion) == 1 {
				errDeleting = "Количество вопросов не может быть меньше 1"
			} else if h.cache.CountOfRequiredQuestions == len(h.cache.ModelQuestion) {
				errDeleting = "Общее количество вопросов не можеть быть меньше количество обязательных вопросов"
			} else {
				h.cache.ModelQuestion = deleteFromCache(h.cache.ModelQuestion, deletingQuestionID)
			}
		} else {
			if len(h.cache.ChoosenUser.Questions) == 1 {
				errDeleting = "Количество вопросов не может быть меньше 1"
			} else if h.cache.ChoosenUser.User.CountOfRequiredQuestions == len(h.cache.ChoosenUser.Questions) {
				errDeleting = "Общее количество вопросов не можеть быть меньше количество обязательных вопросов"
			} else {
				h.cache.ChoosenUser.Questions = deleteFromCache(h.cache.ChoosenUser.Questions, deletingQuestionID)
			}
		}
		if errDeleting != "" {
			http.Redirect(w, r, fmt.Sprintf("/edit_questions?err=%s", errDeleting), http.StatusFound)
			return
		}
		err := h.services.DropQuestion(deletingQuestionID)
		if err != nil {
			logrus.Error("handler/DeleteQuestionWithID: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}

		http.Redirect(w, r, "/edit_questions", http.StatusFound)
	}
}

func (h *Handler) AddQuestion(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			logrus.Error("handler/AddQuestion: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if h.cache.ChoosenUser == nil {
			userID := h.cache.UserID
			newTitele := r.FormValue("question")
			newAnswer := r.FormValue("answer")
			newQuestion := model.Question{
				UserID: userID,
				Title:  newTitele,
				Answer: newAnswer,
			}
			questionID, err := h.services.AddQuestion(newQuestion)
			if err != nil {
				logrus.Error("handler/AddQuestion: ", err)
				w.Write([]byte("что-то пошло не так"))
				return
			}
			newQuestion.QuestionID = questionID
			h.cache.ModelQuestion = append(h.cache.ModelQuestion, newQuestion)
		} else {
			userID := h.cache.ChoosenUser.User.UserID
			newTitele := r.FormValue("question")
			newAnswer := r.FormValue("answer")
			newQuestion := model.Question{
				UserID: userID,
				Title:  newTitele,
				Answer: newAnswer,
			}
			questionID, err := h.services.AddQuestion(newQuestion)
			if err != nil {
				logrus.Error("handler/AddQuestion: ", err)
				w.Write([]byte("что-то пошло не так"))
				return
			}
			newQuestion.QuestionID = questionID
			h.cache.ChoosenUser.Questions = append(h.cache.ChoosenUser.Questions, newQuestion)
		}

		http.Redirect(w, r, "/edit_questions", http.StatusFound)
	}
}

func updateCache(oldQuestions []model.Question, newQuestion model.Question) {
	var indexOfEditingQuestion int
	for index, item := range oldQuestions {
		if item.QuestionID == newQuestion.QuestionID {
			indexOfEditingQuestion = index
			break
		}
	}
	oldQuestions[indexOfEditingQuestion] = newQuestion
}

func deleteFromCache(oldQuestions []model.Question, deletingID int) []model.Question {
	var indexOfDeletingQuestion int
	for index, item := range oldQuestions {
		if item.QuestionID == deletingID {
			indexOfDeletingQuestion = index
			break
		}
	}
	return Remove(oldQuestions, indexOfDeletingQuestion)
}

func Remove(slice []model.Question, s int) []model.Question {
	return append(slice[:s], slice[s+1:]...)
}

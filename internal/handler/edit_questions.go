package handler

import (
	"cripta_course_work/internal/model"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

type ViewDataEditQuestions struct {
	EditingQuestionID int
	ModelQuestions    []model.Question
}

func (h *Handler) EditQuestionsView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
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
		err = tmpl.Execute(w, ViewDataEditQuestions{ModelQuestions: h.cache.ModelQuestion, EditingQuestionID: editingQuestionID})
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
		err = tmpl.Execute(w, ViewDataEditQuestions{ModelQuestions: h.cache.ModelQuestion, EditingQuestionID: editingQuestionID})
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
		if h.cache.UserName == "admin" {
			cacheOfCurrentUser = h.cache.ModelQuestion
		} else {
			cacheOfCurrentUser = h.cache.QuestionsOfChoosenUser
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

		http.Redirect(w, r, "/edit_questions", http.StatusMovedPermanently)
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
		if h.cache.UserName == "admin" {
			h.cache.ModelQuestion = deleteFromCache(h.cache.ModelQuestion, deletingQuestionID)
		} else {
			h.cache.QuestionsOfChoosenUser = deleteFromCache(h.cache.QuestionsOfChoosenUser, deletingQuestionID)
		}

		err := h.services.DropQuestion(deletingQuestionID)
		if err != nil {
			logrus.Error("handler/DeleteQuestionWithID: ", err)
			w.Write([]byte("что-то пошло не так"))
			return
		}

		http.Redirect(w, r, "/edit_questions", http.StatusMovedPermanently)
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

		http.Redirect(w, r, "/edit_questions", http.StatusMovedPermanently)
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

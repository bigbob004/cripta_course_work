package handler

import (
	"cripta_course_work/internal/service"
	tools "cripta_course_work/pkg"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"strconv"
)

type ViewDataForAuth struct {
	IsThereError           bool
	Questions              []string
	RemainingCountAttempts int
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
		}
		//Получаем ответы пользователя на shuffle вопросы
		var userAnswers []string
		for key, values := range r.PostForm {
			if key == "answer" {
				userAnswers = values
			}
		}
		isEqual := h.services.ValidateQuestions(userAnswers, h.cache)
		if isEqual {
			http.Redirect(w, r, "/account", http.StatusFound)
		} else {
			if h.cache.RemainingCountAttempts-1 == 0 {
				http.Redirect(w, r, "/exit", http.StatusFound)
			}
			h.cache.RemainingCountAttempts--
			data := ViewDataForAuth{
				IsThereError:           true,
				Questions:              h.cache.ShuffleQuestions,
				RemainingCountAttempts: h.cache.RemainingCountAttempts,
			}
			tmpl, _ := template.ParseFiles("./template/auth_form.html")
			tmpl.Execute(w, data)
		}
	}
}

func (h *Handler) AuthView(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := r.ParseForm(); err != nil {
			w.Write([]byte("что-то пошло не так"))
			return
		}
		if h.cache.ShuffleQuestions == nil {
			userIDStr := r.URL.Query().Get("user_id")
			if userIDStr == "" {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Something bad happened!"))
				return
			}
			userID, _ := strconv.Atoi(userIDStr)
			questionsAndAnswers, err := h.services.GetQuestionsByUserID(userID)
			h.cache.ModelQuestion = questionsAndAnswers
			if err != nil {
				logrus.Error("internal/handler/auth: ", err)
				w.Write([]byte("что-то пошло не так"))
				return
			}

			//TODO КОСТЫЛИ ЕБУЧИЕ, ЗА ЭТО ТЫ ПОПАДЁШЬ В АД!
			questions := make([]string, 0, len(questionsAndAnswers))
			questionToAnswer := make(map[string]string, len(questionsAndAnswers))
			for _, item := range questionsAndAnswers {
				questions = append(questions, item.Title)
				questionToAnswer[item.Title] = item.Answer
			}
			//выбираем подмн-во обязательных вопросов из общего кол-ва вопросов данного пользователя
			subSetQuestions := tools.UniqueSubSet(questions, h.cache.CountOfRequiredQuestions)
			subSetQuestionToAnswer := make(map[string]string, h.cache.CountOfRequiredQuestions)
			for _, item := range subSetQuestions {
				subSetQuestionToAnswer[item] = questionToAnswer[item]
			}
			h.cache.Questions = subSetQuestions
			h.cache.AuthQuestionToAnswer = subSetQuestionToAnswer

			subSetOptionalQuestions := tools.UniqueSubSet(service.OptionalQuestions, service.CountOfOptionalQuestionsWhileAuth)

			authQuestions := make([]string, 0, len(subSetQuestions)+len(subSetOptionalQuestions))
			authQuestions = append(authQuestions, subSetQuestions...)
			authQuestions = append(authQuestions, subSetOptionalQuestions...)
			tools.ShuffleSlice(authQuestions)
			h.cache.ShuffleQuestions = authQuestions
		}
		data := ViewDataForAuth{
			IsThereError: false,
			Questions:    h.cache.ShuffleQuestions,
		}
		tmpl, _ := template.ParseFiles("./template/auth_form.html")
		tmpl.Execute(w, data)
	}
}

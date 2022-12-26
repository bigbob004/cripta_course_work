package main

import (
	"cripta_course_work/internal/handler"
	"cripta_course_work/internal/repository"
	"cripta_course_work/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

//TODO валидации на пользовательский ввод (кол-во обязательных <= кол-ов вопросов)

//TODO баг по нажатию кнопки назад при добавлении нового пользователя(чекай мидлвару)

//TODO делаем 3 таблицы: редактирование вопросов конкретного пользователя (осталось сделать для обычного пользователя), редактирование пользователя (При удалении пользователя, должны удаляться с ним связанные вопросы, и нас разу должно выкидывать на предыдущий экран)
//TODO повесить кнопку "Назад" на редакторы
//TODO Админа удалять нельзя
//TODO все вопросы пользователя удалить нельзя (мин кол-во вопросов 1)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)

	client, err := repository.NewSqlite()
	if err != nil {
		logrus.Fatalf("faied to initialize db: %s", err.Error())
	}
	dataBase := repository.NewRepository(client)

	services := service.NewService(dataBase)
	handlers := handler.NewHandler(services)

	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.SignIn).Methods("POST")
	router.HandleFunc("/login", handlers.SignInView).Methods("GET")

	router.HandleFunc("/auth", handlers.AuthView).Methods("GET")
	router.HandleFunc("/auth", handlers.Auth).Methods("POST")

	router.HandleFunc("/send_new_user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/send_new_user", handlers.CreateUserView).Methods("GET")

	router.Handle("/account", handlers.Middleware(http.HandlerFunc(handlers.AccountView)))

	router.HandleFunc("/edit_questions", handlers.EditQuestionsView).Methods("GET")
	router.HandleFunc("/edit_question", handlers.EditQuestion).Methods("GET")
	router.HandleFunc("/edit_question_with_id", handlers.EditQuestionWithID).Methods("POST")
	router.HandleFunc("/delete_question", handlers.DeleteQuestionWithID).Methods("POST")
	router.HandleFunc("/add_question", handlers.AddQuestion).Methods("POST")

	router.HandleFunc("/signout", handlers.SignOut).Methods("POST")

	router.HandleFunc("/edit_user", handlers.EditUserView).Methods("GET")
	router.HandleFunc("/block_user", handlers.BlockUser).Methods("POST")

	router.HandleFunc("/exit", handlers.Exit).Methods("POST")

	fileServer := http.FileServer(http.Dir("./template"))
	router.PathPrefix("/res/").Handler(http.StripPrefix("/res/", fileServer))
	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

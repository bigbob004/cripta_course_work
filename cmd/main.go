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

//TODO в БД должно быть поле "кол-во обязательных вопросов" и "общее кол-во вопросов" обязательные - есть подмножество общего кол-ва вопросов

//TODO научиться различать админа и обычного пользователя - done
//TODO по нажатию на кнопку "Выйти" из аккаунта - разлогинить пользователя - done
//TODO можно также сделать мидлвару, которая будет проверять, аутентифицирован пользователь или нет (проверять на nil поле стр-ры questionOfCurrentUser) - done
//TODO хэшировать id пользователя, чтобы он не отображался с параметрах get-запроса
//TODO Gracefull shutdown (при превышении макc. кол-ва попыток входа) - done (не Gracefull shutdown, но завершается, перед завершением можно показывать формочку)
//TODO валидация на форме создания вопросов для пользователя (кол-во обязательных вопросов <= кол-во введённых вопросов)
//TODO рандомайзер вопросов - кажется, что done (нужно ещё потестить)
//TODO сделать форму для редактирования вопросов конкретного пользователя (тут идея закостылить с id-шками вопросов, отрисовывая их прозрачным цветои, чтобы не видел пользователь, но при это я смошу получить эти id)
//TODO баг по нажатию кнопки назад при добавлении нового пользователя
//TODO баг по нажатию кнопки редактор контрольных вопросов
//TODO прикрутить список пользователей к личному кабинету администратора и допилить работу с формой аккаунта юзера
//TODO валидации на пользовательский ввод
// TODO кнопка выхода у обучного юзера - баг
// TODO проверка на существования пользователя с таким же username

type ViewData struct {
	isThereError bool
}

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
	router.HandleFunc("/send_new_user", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "template/add_user.html")
	}).Methods("GET")

	router.Handle("/account", handlers.Middleware(http.HandlerFunc(handlers.AccountView)))

	router.Handle("/edit_questions", http.HandlerFunc(handlers.EditQuestionsView)).Methods("GET")

	router.HandleFunc("/signout", handlers.SignOut).Methods("POST")

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

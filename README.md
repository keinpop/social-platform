# Социальная платформа МАИ

### О проекте:

Приложение «Социальная платформа МАИ» представляет собой нечто вроде «поисковика» студентов и учителей по всему 8 Университету МАИ. У пользователя приложения будет доступен личный кабинет, куда будет возможно добавить: ФИО, программу обучения (например, «ПМИ»), текущий курс, стек технологий (например, «C++», «Golang»), компании, в которых работал пользователь, роли в команде (например, «Backend-разработчик»), развёрнутое поле «о себе», где также можно будет подробнее описать о том, чем человек занимался в том или ином проекте, оставить личные данные, ссылки на GitHub или LeetCode и т.д.

### Стек технологий:

Используемый язык для написания backend составляющей - Go. Нами был выбран фреймворк Gin и Gorm для работы с базой данных Postres. Проект написан в стиле архитектуры REST. Весь проект запускается в качестве пакета docker-контейнеров. Помимо backend-части в docker также лежат Prometheus и Grafana, для сбора различных метрик и их визуализации.

### Краткое описание содержания 'backend' части:

* cmd: Здесь находятся main файлы двух сервисов приложения - auth (сервис аутентификации пользователей) и app (сервис с самим приложением),
* config: Папка содержит конфиги наших сервисов,
* docs: Содержит документацию Swagger,
* internal:
  * api: В ней находятся различные API (Company, Programm, Role, Technology, User),
  * app: Содержит работу сервиса App,
  * auth: Содержит работу сервиса Auth,
  * clients: Содержит работу с базами данных.
* db: Хранит данные для отправки в базу данных,
* frontend: Содержит реализацию frontend-составляющей приложения,
* grafana: Описание контейнера grafana,
* prometheus: Описание контейнера prometheus,
* compose.yaml: Содержит данные Docker-контейнеров.

### Запуск

В корневой папке проекта пишем 'make docker-compose'. После этого проект запустится.

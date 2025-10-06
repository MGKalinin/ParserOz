# What is it about? 
## Примерный алгоритм:   
- сайт ozon -> в поле поиска наименование -> срок доставки: ? (до 7 дней) -> тип: часы наручные -> сортировка цены: по возрастанию
### Вехи:
- парсинг цены в Ozon по наменованию
- результат в бд
- графическое отображение динамики движения цены?
- по изменению цены уведомление в телеграм?
- деплой в Docker

```
ozon-price-tracker/
├── docker-compose.yml
├── .env
├── go.mod
├── pkg/
│   ├── parser/              # Сервис парсинга
│   │   ├── Dockerfile
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── parser/
│   │   │   │   └── ozon_parser.go
│   │   │   ├── kafka/
│   │   │   │   └── producer.go
│   │   │   └── config/
│   │   │       └── config.go
│   │   └── go.mod
│   ├── storage/             # Сервис сохранения в БД
│   │   ├── Dockerfile
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── repository/
│   │   │   │   └── price_repo.go
│   │   │   ├── kafka/
│   │   │   │   └── consumer.go
│   │   │   └── config/
│   │   │       └── config.go
│   │   └── go.mod
│   └── common/              # Общие утилиты
│       ├── types/
│       │   └── price.go
│       └── go.mod
└── scripts/
    └── init-db.sql
```
- [Selenium in Golang: Step-by-Step Tutorial 2025](https://www.zenrows.com/blog/selenium-golang#why-use-selenium-in-go) 
- [Web Scraping in Golang: 2025 Complete Guide](https://www.zenrows.com/blog/web-scraping-golang#build-first-golang-scraper)
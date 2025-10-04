# What is it about? 
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
- для парсинга использовать Colly

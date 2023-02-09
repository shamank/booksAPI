### BOOKS API

пробный REST API сервис на Go

настройка сервера и подключаемой БД:
```path
./configs/config.yml
```

нужно создать .env файл.
пример: .env.example


##### запуск проекта:
```bash
make build
```

##### создание миграций:
```bash
make migrate
```

##### отмена миграций:
```bash
make migrate-down
```
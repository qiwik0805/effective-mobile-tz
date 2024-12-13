## Swagger

Доступен по адресу .../swagger/index.html

## Документация по запуску приложения

Это приложение можно запустить как напрямую (без Docker), так и с использованием Docker.

### Запуск без Docker

1. **Настройка переменных окружения:**
* Создайте или отредактируйте файл `.env` в корневой директории проекта.
* Укажите следующие настройки:
* `PORT`: Порт, на котором будет работать сервер (например, `80`).
* `MUSIC_INFO_BASE_URL`: Базовый URL для сервиса music_info.
* `SONG_REPOSITORY_DSN`: Строка подключения к базе данных PostgreSQL.
* Пример содержимого файла `.env`:
```env
PORT=80
MUSIC_INFO_BASE_URL=http://localhost:35090
SONG_REPOSITORY_DSN=postgresql://user:password@localhost:8431/tzdb?sslmode=disable
```
2. **Запуск сервера:**
* Запустите приложение, выполнив команду в корневой директории проекта:
```bash
go run cmd/server/server.go
```
* Приложение будет доступно по адресу, сконфигурированному через переменные окружения.

### Запуск с использованием Docker

1. **Настройка переменных окружения в `docker-compose.yml`:**
* Откройте файл `docker-compose.yml` в корневой директории проекта.
* Найдите секцию `environment` для сервиса вашего приложения.
* Измените переменные окружения в соответствии с вашими потребностями:
* `PORT`: Порт, на котором будет работать сервер.
* `MUSIC_INFO_BASE_URL`: Базовый URL для сервиса music_info.
* `SONG_REPOSITORY_DSN`: Строка подключения к базе данных PostgreSQL.
* Пример конфигурации `docker-compose.yml`:
```yaml
# ... остальные настройки сервиса
environment:
PORT: 80
MUSIC_INFO_BASE_URL: http://localhost:35090
SONG_REPOSITORY_DSN: postgresql://user:password@effective-mobile-tz-db:5432/tzdb?sslmode=disable
# ... остальные настройки сервиса
```
2. **Запуск Docker Compose:**
* В корневой директории проекта выполните одну из следующих команд:
* `make run` (если у вас настроен `make`)
* `docker-compose up` (если у вас нет `make`, или вы хотите использовать Docker Compose напрямую)
* Приложение будет запущено в Docker-контейнере и будет доступно по адресу, сконфигурированному через переменные окружения.

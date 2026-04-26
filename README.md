# go-blog

**Go Blog** - это современный блоговый движок, написанный на Go с использованием фреймворка Gin. Проект предоставляет полноценный бэкенд с REST API для управления статьями и пользователями, а также фронтенд часть.

## Особенности

- REST API на основе Gin
- Аутентификация с использованием JWT
- PostgreSQL в качестве базы данных
- Docker контейнеризация
- Чистая архитектура с разделенными слоями
- Логгирование с использованием Zap
- CORS поддержка
- Middleware для аутентификации и авторизации

## Архитектура

Проект следует принципам чистой архитектуры и разделен на следующие слои:

```
cmd/            # Точка входа в приложение
internal/       # Основной код приложения
  app/          # Инициализация и запуск приложения
  config/       # Конфигурация
  handler/      # HTTP обработчики
  middleware/   # Middleware
  model/        # Модели данных
  repository/   # Работа с базой данных
  router/       # Маршрутизация
  service/      # Бизнес-логика
  util/         # Утилиты
frontend/      # Фронтенд часть
```

## Технологии

- **Backend**: Go 1.25.7, Gin, GORM
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Logging**: Zap
- **Frontend**: HTML, CSS, JavaScript
- **Containerization**: Docker

## API Эндпоинты

### Аутентификация

| Метод | Эндпоинт | Описание |
|-------|----------|----------|
| POST | `/api/users/register` | Регистрация нового пользователя |
| POST | `/api/users/login` | Авторизация пользователя |

### Статьи

| Метод | Эндпоинт | Описание | Требуется авторизация |
|-------|----------|----------|-----------------------|
| GET | `/api/articles` | Получение всех статей | Нет |
| GET | `/api/articles/:id` | Получение статьи по ID | Нет |
| POST | `/api/admin/articles` | Создание новой статьи | Да |
| PUT | `/api/admin/articles/:id` | Обновление статьи | Да |
| DELETE | `/api/admin/articles/:id` | Удаление статьи | Да |

## Установка и запуск

### Предварительные требования

- Docker и Docker Compose
- Go 1.25.7 (если хотите запускать без Docker)

### 1. Клонирование репозитория

```bash
git clone https://github.com/kavlan-dev/go-blog.git
cd go-blog
```

### 2. Настройка окружения

Создайте файл `.env` на основе примера:

```bash
cp .env.example .env
```

Отредактируйте `.env` файл, указав свои параметры для базы данных:

```bash
ENV=dev #по умолчанию prod
CORS=http://localhost:8000 #по умолчанию *
DB_HOST=db #по умолчанию localhost
DB_USER=myuser
DB_PASSWORD=mypass
DB_NAME=blogdb
JWT_SECRET=your_jwt_secret
```

### 3. Запуск с помощью Docker

```bash
docker-compose up --build
```

Эта команда запустит:
- Backend сервер на порту `8080`
- Frontend сервер на порту `8000`
- PostgreSQL базу данных на порту `5432`

### 4. Запуск без Docker (только бэкенд)

```bash
# Установите зависимости
go mod download

# Запустите приложение
go run cmd/app/main.go
```

## Использование API

### Регистрация пользователя

```bash
curl -X POST http://localhost:8080/api/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpassword"
  }'
```

### Авторизация пользователя

```bash
curl -X POST http://localhost:8080/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpassword"
  }'
```

### Получение всех статей

```bash
curl -X GET http://localhost:8080/api/articles
```

### Создание статьи (требуется JWT токен)

```bash
curl -X POST http://localhost:8080/api/admin/articles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "Моя первая статья",
    "content": "Содержимое статьи..."
  }'
```

## Структура базы данных

Проект использует PostgreSQL с следующими основными таблицами:

- **users**: Хранит информацию о пользователях
- **articles**: Хранит статьи блога

## Аутентификация

Проект использует JWT (JSON Web Tokens) для аутентификации:
- Токены генерируются при успешной авторизации
- Токены должны передаваться в заголовке `Authorization: Bearer <token>`
- Срок действия токена можно настроить в конфигурации

## Конфигурация

Конфигурация загружается из переменных окружения и `.env` файла. Основные параметры:

- `DB_*` - параметры подключения к базе данных
- `JWT_SECRET` - секретный ключ для генерации JWT токенов
- `ENV` - окружение (dev/prod)

## Тестирование

Проект поддерживает различные окружения:
- **dev**: Режим разработки с детальным логгированием
- **prod**: Продакшн режим с оптимизированными настройками

## Фронтенд

Фронтенд часть находится в директории `frontend/` и включает:
- HTML шаблоны для основных страниц
- CSS стили
- JavaScript для взаимодействия с API

Доступные страницы:
- Главная страница (`/`)
- Страница статьи (`/article.html`)
- Дашборд администратора (`/dashboard.html`)
- Страница редактирования статьи (`/edit-article.html`)
- Страница авторизации (`/login.html`)

## Docker

Проект полностью контейнеризирован:
- **backend**: Go приложение
- **frontend**: Nginx сервер для статических файлов
- **db**: PostgreSQL база данных

## Лицензия

Проект распространяется под лицензией MIT. Подробности смотрите в файле [LICENSE](LICENSE).

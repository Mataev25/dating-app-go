Приложение для знакомств. Pet-проект для изучения Go.

Цель проекта

Изучение и демонстрация возможностей Go в создании бэкенд-приложений:
- REST API архитектура
- Работа с базой данных
- Аутентификация и авторизация  
- Тестирование и документация


### Технологический стек
- **Язык:** Go 1.22+
- **База данных:** SQLite3 (для разработки), PostgreSQL (готово к продакшену)
- **HTTP роутинг:** Стандартная библиотека net/http
- **Работа с БД:** SQLx + драйверы SQLite/PostgreSQL
- **Безопасность:** BCrypt для хеширования паролей
- **Идентификаторы:** UUID v4

## 🔧 Детальное описание компонентов

### Модель данных (`internal/user/model.go`)

```go
// User представляет сущность пользователя в системе
type User struct {
    ID          string     `json:"id" db:"id"`
    Email       string     `json:"email" db:"email"`
    Password    string     `json:"-" db:"password"`  // Не сериализуется в JSON
    Name        string     `json:"name" db:"name"`
    Age         int        `json:"age" db:"age"`
    Gender      Gender     `json:"gender" db:"gender"`
    Description string     `json:"description" db:"description"`
    LookingFor  LookingFor `json:"looking_for" db:"looking_for"`
    CreatedAt   time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
```
Особенности:
Используем теги db для маппинга на колонки БД
Пароль исключен из JSON ответов (json:"-")
Типизированные enum'ы для Gender и LookingFor

Repository интерфейс определяет контракты для работы с данными
userRepository реализует Repository интерфейс

Реализованные методы:

CreateUser() - создает пользователя с хешированием пароля

GetUserByEmail() - поиск по email (для проверки дубликатов)

GetUserByID() - поиск по UUID

Service содержит бизнес-логику приложения
    // 1. Валидация входных данных
    // 2. Проверка уникальности email
    // 3. Хеширование пароля
    // 4. Сохранение в БД
    // 5. Возврат UserResponse (без пароля)

Бизнес-правила:
Email должен быть уникальным
Возраст: 18-100 лет
Имя: 2-50 символов
Описание: до 500 символов

Handler обрабатывает HTTP запросы
    // Роутинг на основе пути и метода
    // POST /api/v1/users/register -> Register()
    // GET  /api/v1/users/{id}     -> GetUser()
Индексы БД:

idx_users_email - для быстрого поиска по email

idx_users_age_gender - для будущего поиска по критериям

idx_users_looking_for - для матчинга


# 1. Клонирование
git clone https://github.com/yourusername/dating-app.git
cd dating-app

# 2. Установка зависимостей
go mod tidy

# 3. Запуск
go run cmd/server/main.go

# 4. Тестирование
curl http://localhost:8080/health

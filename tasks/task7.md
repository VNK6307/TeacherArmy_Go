# Задание 7. Добавление интерфейсов и модульность

## Цель:
Научиться отделять бизнес-логику от реализации хранилища через интерфейсы. Это позволит переключаться между хранением в памяти и базой данных без изменений в основной логике.

---

## Что нужно сделать:

### 1. Создай интерфейс `Store`
В файле `internal/race/service.go` или `internal/storage/store.go` определи интерфейс:

```go
type Store interface {
    Add(RaceResult) error
    GetAll() ([]RaceResult, error)
}
```

---

### 2. Обнови структуру `Service`
Сделай так, чтобы сервис зависел от интерфейса:

```go
type Service struct {
    store Store
}

func NewService(s Store) *Service {
    return &Service{store: s}
}
```

---

### 3. Переделай вызовы внутри `Service`
Все операции с данными (добавление, получение) теперь должны вызываться через `s.store.Add(...)` и `s.store.GetAll(...)`.

---

### 4. Реализуй `Store` в памяти
В `internal/storage/memory.go` создай:

```go
type InMemoryStore struct {
    mu   sync.Mutex
    data []RaceResult
}

func (m *InMemoryStore) Add(r RaceResult) error {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.data = append(m.data, r)
    return nil
}

func (m *InMemoryStore) GetAll() ([]RaceResult, error) {
    m.mu.Lock()
    defer m.mu.Unlock()
    return append([]RaceResult(nil), m.data...), nil
}
```

---

### 5. Реализуй `Store` для Postgres
В `internal/storage/postgres.go` создай аналогичную структуру с методами, работающими через `sqlx`.

Пример:

```go
type PostgresStore struct {
    db *sqlx.DB
}

func (p *PostgresStore) Add(r RaceResult) error {
    query := `INSERT INTO race_results (...) VALUES (...)`
    _, err := p.db.Exec(query, ...)
    return err
}

func (p *PostgresStore) GetAll() ([]RaceResult, error) {
    var results []RaceResult
    query := `SELECT * FROM race_results`
    err := p.db.Select(&results, query)
    return results, err
}
```

---

### 6. Обнови `main.go`
Передай нужную реализацию (например, `PostgresStore` или `InMemoryStore`) в `NewService`.

Пример:

```go
store := storage.NewPostgresStore(db)
// или
// store := storage.NewInMemoryStore()

service := race.NewService(store)
handler := http.NewHandler(service)
```

Теперь вся логика "что делать" — в `race`, а "где хранить" — в `storage`.

---

## Что должно получиться:
- Сервис работает с интерфейсом `Store`, не зная, где именно хранятся данные.
- Можно легко менять хранилище: в памяти, в базе, в файле — ничего не меняя в HTTP и бизнес-логике.
- Код стал более гибким и масштабируемым.

---

## Подсказки:
- Интерфейсы в Go — это поведение, а не структуры. Главное — чтобы методы совпадали.
- Можно временно держать оба хранилища (`memory` и `postgres`) и переключаться между ними через флаг или в `main.go`.

---

## Пример:
В `main.go`:

```go
store := storage.NewPostgresStore(db)
// или
// store := storage.NewInMemoryStore()

service := race.NewService(store)
handler := http.NewHandler(service)
```

Теперь вся логика "что делать" — в `race`, а "где хранить" — в `storage`.

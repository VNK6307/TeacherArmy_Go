### Задание 6. Подключение PostgreSQL

**Цель:**  
Научиться подключать базу данных PostgreSQL, сохранять и получать данные о гонках.

---

**Что нужно сделать:**

1. **Установи PostgreSQL** (если ещё не установлен).
   - На Linux: `sudo apt install postgresql`
   - На Mac: `brew install postgresql`
   - Создай базу: `createdb race_db`

2. **Создай таблицу `race_results`**.  
   Используй следующую структуру:

   SQL-запрос:
```
CREATE TABLE race_results (
id SERIAL PRIMARY KEY,
driver TEXT NOT NULL,
position INT NOT NULL,
time TEXT NOT NULL
);
```

3. **Добавь пакет `sqlx` в проект**  
Выполни команду:
```
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
```

4. **Создай файл `internal/storage/postgres.go`**  
В нём должна быть структура `PostgresStore` с методами:

- `Add(result RaceResult) error`
- `GetAll() ([]RaceResult, error)`

Пример подключения:
```
db, err := sqlx.Connect(“postgres”, “user=postgres dbname=race_db sslmode=disable”)
```
5. **Сделай интерфейс `Store`**  
В `race/service.go` или `storage/store.go` определи интерфейс:
```
type Store interface {
Add(RaceResult) error
GetAll() ([]RaceResult, error)
}
```

Теперь ты сможешь использовать как память, так и базу.

6. **Обнови `main.go`**  
Создай `PostgresStore` и передай его в сервис вместо памяти.

---

**Что должно получиться:**

- Программа подключается к PostgreSQL.
- Данные сохраняются в базу при добавлении.
- При запросе `/results` — читаются из базы.
- В памяти больше ничего не хранится.

---

**Подсказки:**

- Тип `RaceResult` должен совпадать с колонками в базе.
- Проверь, чтобы `time` был текстом или временем (на выбор).
- Не забудь обрабатывать ошибки (`if err != nil`).

---

**Пример вывода:**

POST `/results`
```
{
“driver”: “Lewis Hamilton”,
“position”: 1,
“time”: “1:32:45”
}
```

Ответ:
```
{“status”: “ok”}
```


GET `/results`  
Ответ:
```
[
{“driver”: “Lewis Hamilton”, “position”: 1, “time”: “1:32:45”},
…
]
```
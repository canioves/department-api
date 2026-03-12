## Тестовое задание Department API

### Запуск:
1. Скопировать репозиторий
```bash
git clone github.com/canioves/department-api
cd department-api
```
2. Добавить .env файл
``` env
GO_ENV="dev"              #для локального запуска

DB_NAME="db-name"         #имя БД
DB_USER="user"            #имя пользователя 
DB_PASSWORD="pass"        #пароль

DB_HOST_LOCAL="localhost" #локальный хост
DB_HOST_DOCKER="postgres" #хост в docker
DB_PORT="5432"            #порт БД

APP_PORT="8080"           #порт севрера
```
3. Запуск Docker
```bash
docker-compose build
docker-compose up
```

### Эндпоинты:
1. POST: `/departments` - cоздание подразделения

Тело запроса (JSON):
- name: string
- parent_id: int | null - опционально

Пример ответа:
```json
{
    "id": 40,
    "name": "new",
    "created_at": "2026-03-12T10:05:11.836135997Z"
}
```
2. GET: `/departments/{id}` - получение подразделения по ID

Query параметры:
- depth: int - глубина вложенности подразделений, от 1 до 5, опционально, по умолчанию 1
- include_employees: bool - включать ли сотрудников, опционально, по умолчанию true

Пример ответа с параметрами по умолчанию:
```json
{
    "id": 28,
    "name": "Redis Team",
    "created_at": "2026-03-12T09:43:45.002369Z",
    "children": [
        {
            "id": 38,
            "name": "Redis Performance",
            "created_at": "2026-03-12T09:43:45.002369Z",
            "employees": [
                {
                    "id": 77,
                    "department_id": 38,
                    "full_name": "Jackson Green",
                    "position": "Redis Performance Engineer",
                    "hired_at": "2024-01-20T00:00:00Z",
                    "created_at": "2026-03-12T09:43:45.002369Z"
                }
            ]
        },
        {
            "id": 39,
            "name": "Redis Replication",
            "created_at": "2026-03-12T09:43:45.002369Z",
            "employees": [
                {
                    "id": 78,
                    "department_id": 39,
                    "full_name": "Avery Hill",
                    "position": "Replication Specialist",
                    "hired_at": "2024-02-15T00:00:00Z",
                    "created_at": "2026-03-12T09:43:45.002369Z"
                }
            ]
        }
    ],
    "employees": [
        {
            "id": 60,
            "department_id": 28,
            "full_name": "Amy Clark",
            "position": "Redis Specialist",
            "hired_at": "2023-09-15T00:00:00Z",
            "created_at": "2026-03-12T09:43:45.002369Z"
        },
        {
            "id": 61,
            "department_id": 28,
            "full_name": "Kevin Hall",
            "position": "Senior Redis Engineer",
            "hired_at": "2023-04-08T00:00:00Z",
            "created_at": "2026-03-12T09:43:45.002369Z"
        }
    ]
}
```
3. PATCH: `/departments/{id}` - переместить подразделение

Тело запроса (JSON):
- name: string | null - опционально
- parent_id: int | null - опционально

Пример ответа:
```json
{
    "id": 28,
    "name": "updated",
    "parent_id": 1,
    "created_at": "2026-03-12T09:43:45.002369Z"
}
```
4. DELETE: `departments/{id}` - удалить подразделение

Query параметры:
- mode: string, "cascade" | "reassign"
  - cascade - каскадное удаление дочерних подразделений и сотрудников
  - reassign - удаление подразделений и перемещение сотрудников в подразделение с ID = reassign_id
- reassign_id: int - обязательно при mode = "reassign", ID подразделения, в которое перемещать сотрудников

Ответ - 204 (NoContent)

5. POST: `/departments/{id}/employees` - создание сотрудника

Тело запроса (JSON):
- "full_name": string - имя сотрудника
- "position": string - должность
- "hired_at": string | null - дата приема на работу, опционально, формат: dd/mm/yyyy

Пример ответа:
```json
{
    "id": 88,
    "department_id": 28,
    "full_name": "test test",
    "position": "test",
    "hired_at": "2026-03-12T00:00:00Z",
    "created_at": "2026-03-12T10:47:32.653921157Z"
}
```

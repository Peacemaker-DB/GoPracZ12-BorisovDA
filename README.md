# Практическое задание № 12 Борисов Денис Александрович ЭФМО-01-25

Цели:

1.	Освоить основы спецификации OpenAPI (Swagger) для REST API.
2.	Подключить автогенерацию документации к проекту из ПЗ 11 (notes-api).
3.	Научиться публиковать интерактивную документацию (Swagger UI / ReDoc) на эндпоинте GET /docs.
4.	Синхронизировать код и спецификацию (комментарии-аннотации → генерация) и/или «schema-first» (генерация кода из openapi.yaml).
5.	Подготовить процесс обновления документации (Makefile/скрипт).

# Выполнение практического задания

Для выполнения практики за основу был взят проект из практического задания № 12

1. Структура проекта

Для выполнения практической работы была сделана следующая структура проекта

<img width="179" height="450" alt="image" src="https://github.com/user-attachments/assets/9911afb8-b561-4901-9a10-0f3c7bf57b6d" />

Так же были установлены все необходимые расширения для выполнения практики

<img width="790" height="587" alt="Снимок экрана 2025-11-17 035324" src="https://github.com/user-attachments/assets/6428027e-5f90-4b5d-9012-ffaddada774e" />

2. Верхнеуровневые аннотации (в cmd/api/main.go)

Была написана анотация в cmd/api/main.go

<img width="492" height="321" alt="Снимок экрана 2025-11-17 035745" src="https://github.com/user-attachments/assets/9ba281ad-92e3-4062-8a9f-1a0bd6aa3e75" />


3. Аннотации над хэндлерами

После были написаны анотации над функциями в internal/http/handlers/notes.go

Пример аннотации для функции обновления заметки

<img width="616" height="499" alt="image" src="https://github.com/user-attachments/assets/9510213e-f528-405a-84f0-d7b62d9b8c7c" />

Добавляем DTO модели в internal/core

<img width="576" height="460" alt="image" src="https://github.com/user-attachments/assets/475167a5-bb4a-4f84-a812-45cd82ac3738" />

4. Генерация swagger-доков

После было выполнена генерация swagger документа

<img width="1440" height="190" alt="image" src="https://github.com/user-attachments/assets/f9288891-19f3-46f9-b682-de04ff7faf89" />

5. Подключение Swagger UI к серверу (роут /docs)

Для подлючения к Swagger UI был обновленн файл cmd/api/main.go 

<img width="525" height="661" alt="image" src="https://github.com/user-attachments/assets/77d7cb6d-498b-437e-b723-cffca9a36b01" />

Результат подключения

<img width="1811" height="827" alt="image" src="https://github.com/user-attachments/assets/8a7246f2-a854-483c-b67d-9f51845b0a94" />

6. Автоматизация через Makefile

Для автоматизации через Makefile был создан файл Makefile

<img width="382" height="227" alt="image" src="https://github.com/user-attachments/assets/0f0e7275-8bc6-4fb3-b4a7-dbecfdc78f45" />


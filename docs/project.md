# Микро-биржа фрилансеров - описание проекта

## Основные принципы

- монорепо : `backend/`, `web/`, `mobile/`, `packages/shared/` (API-клиент, типы, zod, хуки React Query).
- общий контракт API: `openapi.yaml` -> oapi-codegen (Go/Gin) + orval (TS + React Query).
- модульный монолит, слои handler -> service -> repo.
- эскроу - без реальных платежей, mock-пополнение.

## Стек

| | |
| --- | --- |
| API | Gin, oapi-codegen для клиентов |
| БД | PostgreSQL 16, pgx/v5 |
| Очереди | Redis + asynq |
| WS | coder/websocket, Redis Pub/Sub |
| Файлы | MinIO, presigned URL |
| Веб | Vite, React, TS, React Query, Router, Zustand |
| Мобайл | Expo, Expo Router, SecureStore, NativeWind, Expo Push |
| Инфраструктура | Docker Compose, Caddy/Nginx, GitHub Actions → GHCR → ssh |
| Тесты | testcontainers-go, Playwright |

## Этапы

**0. База**

- Монорепо, линтеры, Makefile
- docker-compose: postgres, redis, minio, api, worker, web; air
- Каркас Go: env-конфиг, slog, graceful shutdown, errgroup, ошибки, middleware
- Каркас веба; CI/CD + VPS + HTTPS

**1. Auth и профили**

- JWT access 15 мин + refresh в БД (ротация), Bearer
- Роли
- Профиль, навыки, аватар, каталог исполнителей
- Seed-данные

**2. Заказы и отклики**

- CRUD заказов, лента: фильтры, FTS (tsvector + GIN), курсорная пагинация
- Отклики
- FSM-ядро + `order_events` + unit-тесты
- Экраны: лента, заказ, публикация, «мои заказы/отклики»

**3. Эскроу**

- Кошелёк, mock-пополнение
- hold / release (− комиссия) / refund
- Идемпотентность, инварианты баланса
- Интеграционные тесты: гонки, двойные запросы
- Споры + мини-админка; экран кошелька

**4. Чат**

- Беседа на заказ; хранение в Postgres, WS только для доставки
- История через REST (курсор); auth WS по тикету; ping/pong
- Redis Pub/Sub; непрочитанные
- Веб: `setQueryData`, реконнект с backoff

**5. Рейтинг, уведомления, фон**

- Отзывы после `completed`; байесовское среднее через задачу asynq
- Outbox → asynq
- Уведомления: in-app + WS, email (Mailpit)
- Cron-задачи

**6. Доработка взаимодействия и UI**

- UX: пустые состояния, ошибки, адаптив
- Проверки владельца (IDOR), rate limit, лимиты загрузок
- `/healthz`, бэкапы `pg_dump`
- Playwright E2E основного сценария

**7. React Native**

- Expo + подключение `packages/shared` (metro.config)
- SecureStore для токенов
- Экраны: вход → лента → заказ/отклик → чат → кошелёк → профиль → отзывы
- Expo Push; сборка EAS / Expo Go

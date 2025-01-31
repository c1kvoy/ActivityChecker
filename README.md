# Site Checker 🌐

Простой и эффективный инструмент для мониторинга доступности веб-сайтов, написанный на Go.

## Возможности ✨

- Параллельная проверка множества сайтов
- Конфигурация через YAML файл
- Измерение времени ответа
- Детальная информация о статусе каждого сайта

## Установка 🚀

1. Установите Go на ваш компьютер.
2. Клонируйте репозиторий:

```bash
git clone https://github.com/your-username/ActivityChecker
cd ActivityChecker
go mod tidy
```

## Использование 📝

1. Настройте список сайтов в `config/urls.yaml`:

```yaml
urls:
  - name: "Google"
    url: "https://google.com"
  - name: "GitHub"
    url: "https://github.com"
```

2. Запустите программу:

```bash
go run main.go
```

## Планы развития 🔮

- [ ] Веб-интерфейс на React
- [ ] Графики времени ответа
- [ ] Уведомления о недоступности
- [ ] API для интеграции
- [ ] Метрики и статистика
- [ ] Docker контейнеризация

## Структура проекта 📁

```
site-checker/
├── config/
│   └── urls.yaml
├── pkg/
│   ├── checker/
│   │   └── checker.go
│   └── config/
│       └── config.go
├── main.go
└── README.md
```

## Лицензия 📄

MIT License


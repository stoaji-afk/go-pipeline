# Device Pipeline

Go-проект, реализующий конвейер обработки устройств: генерация заказов → сборка → прошивка → тестирование → упаковка.

## Структура

```
device-pipeline/
├── cmd/
│   └── main.go              — точка входа: генератор заказов, конфигурация пайплайна, запуск
├── pkg/
│   ├── pipeline/
│   │   ├── pipeline.go      — тип Pipeline: поля, PushOrder, AddStep, Run
│   │   └── worker.go        — обёртка-воркер: читает канал, вызывает stepFn, пишет в следующий канал
│   └── types/
│       └── types.go         — DeviceOrder, Device, Firmware, Report, Package, StepFunc
├── steps/
│   ├── log_step.go          — логирование заказа (мгновенный, ~0 мс)
│   ├── assemble.go          — Заказ → Девайс (400 мс)
│   ├── firmware.go          — Девайс → Девайс (200 мс)
│   ├── testing.go           — Девайс → Девайс (100 мс)
│   └── package.go           — Девайс → Пакет (100 мс)
└── go.mod
```

## Сборка

```bash
go build -o device-pipeline ./cmd/...
```

## Запуск

```bash
# Скомпилированный бинарник
./device-pipeline

# Или через Go
go run ./cmd/...
```

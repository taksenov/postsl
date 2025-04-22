# Пакет логирования

Простой, но мощный пакет логирования для Go-приложений.

## Особенности

- Поддержка разных уровней логирования (DEBUG, INFO, WARNING, ERROR, FATAL)
- Запись логов в файл с возможностью ротации
- Поддержка информации о месте вызова (файл и строка)
- Потокобезопасность
- Гибкая настройка формата вывода

## Использование

### Инициализация

```go
// Логгер с выводом в консоль
log := logger.New(os.Stdout, "", logger.LevelDebug, true)

// Логгер с выводом в файл
log, err := logger.InitFileLogger("logs/app.log", "", logger.LevelInfo, true)
if err != nil {
    panic(err)
}
defer log.Close()
```

### Логирование

```go
log.Debug("Отладочное сообщение")
log.Info("Информационное сообщение")
log.Warning("Предупреждение")
log.Error("Ошибка")
log.Fatal("Фатальная ошибка") // завершает программу
```

### Ротация логов

```go
err := log.Rotate()
if err != nil {
    log.Error("Ошибка ротации: %v", err)
}
```

### Установка уровня логирования

```go
log.SetLevel(logger.LevelWarning) // теперь только WARNING, ERROR, FATAL
```

### Использование контекста

Вспомогательные функции для упрощения работы:

```go
// log.go в вашем основном пакете
package main

import (
    "context"
    "yourproject/logger"
)

func logInfo(ctx context.Context, format string, args ...interface{}) {
    logger.FromContext(ctx).Info(format, args...)
}

func logError(ctx context.Context, format string, args ...interface{}) {
    logger.FromContext(ctx).Error(format, args...)
}

// Аналогично для других уровней...
```

Использование:

```go
func consistentlyRun(ctx context.Context, config *conf.AppConfig) {
    logInfo(ctx, "Consistently Run")
}
```

Преимущества:

- Единый логгер на всю цепочку вызовов
- Легкость тестирования - можно подменить логгер в контексте для тестов
- Чистые сигнатуры функций - не нужно передавать логгер явно в каждую функцию
- Гибкость - можно легко изменить способ логирования для отдельных частей приложения

Важные замечания:

- Всегда проверяйте, что контекст не nil перед использованием

package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// Уровни логирования
const (
	LevelDebug = iota
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

var levelNames = []string{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
	"FATAL",
}

// Logger представляет собой экземпляр логгера
type Logger struct {
	*log.Logger
	level      int
	file       *os.File
	mu         sync.Mutex
	callerInfo bool
}

// contextKey это приватный тип для ключа контекста
type contextKey struct{}

// WithContext добавляет логгер в контекст
func WithContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext извлекает логгер из контекста
func FromContext(ctx context.Context) *Logger {
	if logger, ok := ctx.Value(contextKey{}).(*Logger); ok {
		return logger
	}
	// Возвращаем логгер по умолчанию, если в контексте нет логгера
	return New(os.Stdout, "", LevelInfo, false)
}

// New создает новый экземпляр логгера
func New(out io.Writer, prefix string, level int, callerInfo bool) *Logger {
	return &Logger{
		Logger:     log.New(out, prefix, 0),
		level:      level,
		callerInfo: callerInfo,
	}
}

// InitFileLogger инициализирует логгер с записью в файл
func InitFileLogger(logPath, prefix string, level int, callerInfo bool) (*Logger, error) {
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	return &Logger{
		Logger:     log.New(file, prefix, 0),
		level:      level,
		file:       file,
		callerInfo: callerInfo,
	}, nil
}

// Close закрывает файл лога, если он используется
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

// Rotate выполняет ротацию логов (переименовывает текущий и создает новый)
func (l *Logger) Rotate() error {
	if l.file == nil {
		return nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Закрываем текущий файл
	if err := l.file.Close(); err != nil {
		return err
	}

	// Переименовываем текущий файл
	oldPath := l.file.Name()
	newPath := oldPath + "." + time.Now().Format("2006-01-02_15-04-05")
	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	// Открываем новый файл
	file, err := os.OpenFile(oldPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	l.file = file
	l.Logger.SetOutput(file)
	return nil
}

// logInternal реализует основную логику логирования
func (l *Logger) logInternal(level int, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	// Получаем информацию о вызывающем коде, если нужно
	var caller string
	if l.callerInfo {
		_, file, line, ok := runtime.Caller(2) // Пропускаем 2 фрейма
		if ok {
			caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
		}
	}

	// Формируем сообщение
	msg := fmt.Sprintf(format, args...)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	levelStr := levelNames[level]

	// Формируем итоговую строку лога
	var logLine string
	if l.callerInfo && caller != "" {
		logLine = fmt.Sprintf("%s [%s] %s - %s", timestamp, levelStr, caller, msg)
	} else {
		logLine = fmt.Sprintf("%s [%s] %s", timestamp, levelStr, msg)
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.Logger.Println(logLine)
}

// Debug логирует сообщение уровня DEBUG
func (l *Logger) Debug(format string, args ...interface{}) {
	l.logInternal(LevelDebug, format, args...)
}

// Info логирует сообщение уровня INFO
func (l *Logger) Info(format string, args ...interface{}) {
	l.logInternal(LevelInfo, format, args...)
}

// Warning логирует сообщение уровня WARNING
func (l *Logger) Warning(format string, args ...interface{}) {
	l.logInternal(LevelWarning, format, args...)
}

// Error логирует сообщение уровня ERROR
func (l *Logger) Error(format string, args ...interface{}) {
	l.logInternal(LevelError, format, args...)
}

// Fatal логирует сообщение уровня FATAL и завершает программу
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.logInternal(LevelFatal, format, args...)
	os.Exit(1)
}

// SetLevel устанавливает уровень логирования
func (l *Logger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// GetLevel возвращает текущий уровень логирования
func (l *Logger) GetLevel() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

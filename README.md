# postsl

Раннер запускает, удаленные команды на серверах вашей локальной сети.

## Настройки

Файл `config.example.yaml` переименовать в `config.yaml`:

```sh
cp ./config.example.yaml ./config.yaml
```

Описание настроек конфигурационного файла:

```yaml
# Запускать последовательно или конкурентно
concurency:
  use: false
  # Количество одновременно работающих горутин
  coeff: 5

# Описание серверов
servers:
  # Как коннектиться по ssh: "ADDR" | "NAME"
  use: "ADDR"
  # Список серверов
  list:
    - name: server1
      addr: 10.0.0.1
    - name: server2
      addr: 10.0.0.2

# Выполняемые удаленно команды
# Чтобы работало в цикле, то завершайте каждую запись `command` 
# командой выхода "... ; exit"
remote_command:
    # Команда
  - command: "ls -lah ; exit"
    # Использовать темплейт: `"SERVER_NAME"`, 
    # произойдет подмена в command на имя или адрес сервера
    is_template: false

# Выполняемая команда
command:
  # Предварительно необходимо установить:
  # - переменные окружения `export SSH_PASSWORD="{secret}" && export SSH_USER="{username}"`
  # - программу sshpass
  # Формат команды:
  # `... %s ... %s %s`
  # - Первый %s -- соберутся все параметры для ssh из command.params
  # - Второй %s -- подставится имя или адрес сервера из servers.list.name|addr:
  # - Третий %s -- команда из remote_command.command:
  main: sshpass -p "$SSH_PASSWORD" ssh -tt %s $SSH_USER@%s %s
  # Параметры команды ssh
  params:
    - -o "StrictHostKeyChecking=accept-new"
    - -o "PreferredAuthentications=password"
```

Все параметры конфигурации можно заменить переменными окружения с спрефиксом: `POSTSL_`:
- Например для `command.main` будет `POSTSL_COMMAND_MAIN`

## Запуск

```sh
go run main.go ssh-run
```

## Статистика работы

В режиме посдледовательного запуска, подсчет количества выполненных команд ведется по каждой команде отдельно и выводится в терминал, например:

```sh
COMMAND INDEX= 66
```

В конкуретном режиме, подсчет отработавщих запросов выодится после окончания работы программы:

```sh
All requests done. Count:  56
```

Статистика помогает понять, на каком количестве серверов были выполнены команды.

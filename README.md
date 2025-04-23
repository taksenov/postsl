# postsl

## Внимание

**Отказ от ответственности:** Данное программное обеспечение предоставляется «как есть», без каких-либо гарантий. Разработчик не несёт ответственности за любые прямые или косвенные убытки, возникшие в результате использования или невозможности использования данного ПО.

**Disclaimer:** This software is provided "as is" without any warranties. The developer shall not be liable for any direct or indirect damages arising from the use or inability to use this software.

**Обнаруженные особенности работы:** By design в версии `0.1.0`, в процессе работы, в директории из которой запускается программа командой `postsl ssh-run`, создается директория с логами `./logs/`. Данное поведение в последующих версиях будет изменено, создание логов будет управляться через конфигурирование.

---

## О программе

Раннер запускает, удаленные команды на серверах вашей локальной сети.

## Установка

На примере macOS Apple silicon CPU.

1) Создать директорию и распаковать в нее бинарный файл с программой:

  ```sh
  mkdir -p ~/.local/postsl/bin/
  tar -C ~/.local/postsl/bin -xzf ./dist/postsl-darwin-arm64.tar.gz
  ```

2) Для обновления, предварительно требуется очистить директорию:

  ```sh
  rm -rf ~/.local/postsl/bin/* && tar -C ~/.local/postsl/bin -xzf ./dist/postsl-darwin-arm64.tar.gz
  ```

3) Добавить путь в переменную окружения PATH:

  ```sh
  export PATH=$PATH:'/Users/{YOUR_USERNAME}/.local/postsl/bin'
  ```

  Следует руководствоваться настройками, зависящими от используемого типа shell.

4) Проверка установки:

  ```sh
  postsl version
  ```

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
postsl ssh-run
```

```sh
go run -race main.go ssh-run
```

`postsl ssh-run` ищет файл с именем `config.yaml` в той директории из которой он запущен, это поведение останется неизменным. Все параметры конфигурации можно заменить переменными окружения с спрефиксом: `POSTSL_`.

## Статистика работы

В режиме посдледовательного запуска, подсчет количества выполненных команд ведется по каждой команде отдельно и выводится в терминал, например:

```sh
COMMAND INDEX= 9
```

В конкуретном режиме, подсчет отработавщих запросов выодится после окончания работы программы:

```sh
All requests done. Count:  56
```

В процессе работы создаются log-файлы в формате: `./logs/log-2006-01-02--15-04-05.log`

Статистика помогает понять, какие и на каком количестве серверов команды были выполнены.

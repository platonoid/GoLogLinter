## Сборка бинарника

Перед использованием необходимо собрать исполняемый файл.

### macOS

#### Apple Silicon (M1 / M2):

```bash
cd /path/to/linter
GOOS=darwin GOARCH=arm64 go build -o loggerlint ./cmd/loggerlint
chmod +x ./loggerlint
```

#### Intel Mac:

```bash
cd /path/to/linter
GOOS=darwin GOARCH=amd64 go build -o loggerlint ./cmd/loggerlint
chmod +x ./loggerlint
```

#### Windows:

```bash
cd \path\to\linter
GOOS=windows GOARCH=amd64 go build -o loggerlint.exe ./cmd/loggerlint
```

После сборки бинарник будет в корне проекта (loggerlint на macOS/Linux, loggerlint.exe на Windows).

## Проверка работы бинарника

Запуск без аргументов:

./loggerlint         # macOS/Linux
loggerlint.exe       # Windows

Вывод должен быть примерно таким:
loggerlint is a tool for static analysis of Go programs.

Usage of loggerlint:
        loggerlint unit.cfg     # execute analysis specified by config file
        loggerlint help         # general help, including listing analyzers and flags
        loggerlint help name    # help on specific analyzer and its flags

## Запуск тестов

```bash
cd /path/to/linter
go test ./...
```

## Интеграция с golangci-lint

В .golangci.yml укажите путь к бинарнику:
```YAML
linters-settings:
  loggerlint:
    path: ./loggerlint
```

Golangci-lint сам вызовет loggerlint для всех пакетов проекта.

Не нужно добавлять loggerlint в linters.enable.

Запуск проверки всего проекта: golangci-lint run
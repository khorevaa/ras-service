# ras-service
Microservice control `RAS` over HTTP/gRPC  


[![Release](https://img.shields.io/github/release/khorevaa/ras-service.svg?style=for-the-badge)](https://github.com/khorevaa/ras-service/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Build status](https://img.shields.io/github/workflow/status/khorevaa/ras-service/goreleaser?style=for-the-badge)](https://github.com/khorevaa/ras-service/actions?workflow=goreleaser)
[![Codecov branch](https://img.shields.io/codecov/c/github/khorevaa/ras-service/master.svg?style=for-the-badge)](https://codecov.io/gh/khorevaa/ras-service)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](http://godoc.org/github.com/khorevaa/ras-service)
[![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=for-the-badge)](https://saythanks.io/to/khorevaa)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=for-the-badge)](https://github.com/goreleaser)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)



## Запуск в докере

```shell
docker run -v $(pwd):/tmp/dist ghcr.io/khorevaa/ras-service     
```

## Настройка логов

### Через файл настройки
Создать рядом с приложением файл `logos.yaml` с содержимым

```yaml
appenders:
  console:
    - name: CONSOLE
      target: stdout
      encoder:
        console:

  rolling_file:
    - name: FILE
      file_name: ./logs/ras-service.log
      max_size: 100
      max_age: 10
      encoder:
        json:
loggers:
  root:
    level: info
    appender_refs:
      - CONSOLE
  logger:
    - name: "github.com/khorevaa/ras-service"
      appender_refs:
        - CONSOLE
        - FILE
      level: debug     

```

### Через переменные окружения
```shell
export LOGOS_CONFIG="appenders.rolling_file.0.name=FILE;
appenders.rolling_file.0.file_name=./logs/ras-service.log;
appenders.rolling_file.0.max_size=100;
appenders.rolling_file.0.encoder.json;
loggers.logger.0.level=debug;
loggers.logger.0.name=github.com/khorevaa/ras-service;
loggers.logger.0.appender_refs.0=CONSOLE;
loggers.logger.0.appender_refs.1=FILE;"
```
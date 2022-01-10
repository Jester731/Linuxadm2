## Условия:
- docker, docker-compose
- gitlab runner

## Конфиги:

yaml-файл конфигурации Gitlab CI/CD (SSH_PRIVATE_KEY указан в gitlab variables):

[.gitlab-ci.yml](.gitlab-ci.yml)

docker-compose файл для mysql (bitnami):

[docker-compose.yml](docker-compose.yml)

## Чат с pub/sub в Redis:
- Контейнеры поднялись после pipeline'а gitlab ci/cd:

![linux console](1.png)

- Для реализации pub/sub механизма написана программа на языке go (можно запускать как скрипт, либо скомпилировать и запускать собранный бинарник)

[Исходный код](chat/chat.go)

- Каждые 5 секунд скрипт, запущенный с флагом -pub будет отправлять в Redis сообщение "Hello, world". Все подписанные клиенты будут его получать. Пример работы:

![linux console](2.png)

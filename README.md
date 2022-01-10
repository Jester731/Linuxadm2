## Условия:
- docker, docker-compose
- gitlab runner

## Конфиг:

yaml-файл конфигурации Gitlab CI/CD (SSH_PRIVATE_KEY указан в gitlab variables):
[.gitlab-ci.yml](.gitlab-ci.yml)

docker-compose файл для mysql (bitnami):
[docker-compose.yml](docker-compose.yml)

## Шаги:

- Запуск pipeline и его выполнение:

![linux console](1.png)

![linux console](2.png)

![linux console](3.png)

![linux console](4.png)

![linux console](5.png)

- Проверка статуса контейнеров:

![linux console](6.png)

- Доступ к БД с master:

![linux console](7.png)

- Репликация данных на slave:

![linux console](8.png)

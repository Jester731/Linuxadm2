1. Jenkins установлен в docker-контейнере
2. Gitlab

```
  docker run --detach \
  --publish 443:443 --publish 80:80 --publish 22:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  gitlab/gitlab-ee:latest
```

3. Сделать задания на разворачивания Nginx, при деплое изменений в html страничке эти изменения должны автоматически выкладываться на сервер
  - 192.168.14.132 - сервер с поднятым Nginx. Настраиваем доступ по ssh с машин, на которых запущен jenkins и gitlab-runner.
  - Для Jenkins подключаем репозиторий github с index.html и настраиваем job :

  ![](1.jpg)

  - Для gitlab-ci создаем проект, в него также добавляем index.html и .gitlab-ci.yml

  [Gitlab-ci yaml](.gitlab-ci.yml)

Собрать docker образ: `make build`.

Запустить docker образ: `make run`. Контейнер запускается с флагом --rm.
upd: уже нет. Сделал через docker-compose дабы обеспечить отказоустойчивость.

Скачать docker образ - `docker pull dnevsky/vk-bot-welcome`

`.env` переменные:
```
VK_TOKEN = <token>
```

Так же бота залил к себе на сервер и он доступен по ссылке - `https://vk.com/public220322235`. Напишите что-нибудь в сообщения.

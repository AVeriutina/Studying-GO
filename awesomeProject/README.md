Последние добавленные изменения с использованием базы данных не отдебажены.
http и grpc обрабатывают запросы и хранят аккаунты на сервере хорошо. Вот несколько примеров запуска в клиенте:

Создание аккаунта: `./main -cmd create -name alice -amount 50`

Вывод аккаунта: `./main -cmd get -name alice`

Изменение имени аккаунта: `./main -cmd change_name -name alice -new_name alice123`

Изменение баланса аккаунта: `./main -cmd change_amount -name alice123 -amount 60`

Удаление аккаунта: `./main -cmd delete -name alice123`

Часть изменений была закоммичена до дедлайнов, некоторые добавлялись чуть позже

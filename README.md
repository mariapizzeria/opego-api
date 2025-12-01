# Тестовое задание для OPEGO 


## Описание реализованных API

**GET /api/order/{order_id}/http**

Тестовое апи - получение данных о статусе заказа, с учловием что оно будет вызываться если на фронте будет реализован цикл

**GET /api/order/{order_id}** 

Апи получения статуса заказа. Реализовано с помощью gRPC streaming

**POST /api/order**                           

Добавление заказа

_Тело запроса:_

`````
{
   "passenger_id": 1,
   "address_from": "ул. Братьев Кашириных",
   "address_to": "ТЦ МореМолл",
   "tariff": "comfort_plus",
   "selected_services": ["pet"], 
   "comment": "Буду на парковке"
}
`````

_Тело ответа:_

`````{
"order_id": 46,
"updated_at": "2025-12-01T14:21:34.4866672+05:00",
"created_at": "2025-12-01T14:21:34.4866672+05:00",
"canceled_at": null,
"completed_at": null,
"passenger": 1,
"order_status": "pending",
"arrived_code": "",
"address_from": "ул. Братьев Кашириных",
"address_to": "ТЦ МореМолл",
"tariff": "comfort_plus",
"selected_services": ["pet"],
"comment": "Буду на парковке",
"price": 350
}
`````

**POST /api/order/{order_id}/cancel**  

Отмена заказа по id

**POST /api/order/{order_id}/accept**  

Принятие заказа исполнителем. Предварительно проверяя его статус

**POST /api/driver/status** 

Обновление статуса водителя

_Тело запроса_:

`````
{
    "driver_id": 1,
    "available": true,
    "current_location": {"lat": 59.9343, "ing": 30.3061}
}
`````

_Тело ответа_:

`````
{
    "driver_id": 1,
    "available": true,
    "current_location": {
        "lat": 59.9343,
        "ing": 30.3061
    }
}
`````

**POST /api/order/{order_id}/arrived**  

Получение кода подтверждения. Изменение статуса заказа

Тело ответа:

`````
{
    "order_id": 36,
    "order_status": "waiting_for_confirmation",
    "arrived_code": "1032"
}
`````

**PUT /api/order/{order_id}/status**      

Обновление статуса заказа после того как исполнитель и заказчик сравнили коды

Тело запроса:

`````
{ "order_status": "completed" }
`````

**PUT /api/order/{order_id}/status/search**

Обновление статуса заказа с pending на searching когда пользователь нажал кнопку "Заказать"
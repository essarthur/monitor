# API for dashboard

## POST Input 
1. Установка количества транзакций
2. Установка количества шард

## GET Return
1. Получение TPS
2. Количества транзакций
3. Количества нод
4. Количсетва шард


## Установка количества шард
Method ("POST")
**/api/set/shard/{{count}}**

Return 
{"status":"ok", "id":11234}

## Установка количества транзакций
Method ("POST")
**/api/set/tx/{{count}}**

Return 
{"status":"ok", "id":11235}


# Получение информации от сервера для дашборда
Method ("GET")


**/api/output/**

```json
{
  "id",       123446,
  "tps":      10,
  "shard":    20,
  "txcount":  1000,
  "time":     "2018-12-25 15:34:45",
  "status":   "success"
  }
```

# Subscribe project

## Basically opportunities 
1. Created API 
2. View for preview reports
3. Manage accounts
4. Log critical operation
5. Send notification to telegram
6. Setting system


## API
1. Created notification from internal systems
2. Monitoring activity system
3. Dleted old information
4. Geting information by the period
5. Import log in JSON format
6. Backaping data

|Api|Method|Description|
|--|--|--|
|**/api/add/**|POST|Adding information to system|
|**/api/project/{Pojectname}**|POST|Make new project
|**/api/project/{Pojectname}**|GET|Sending all information by project
|**/api/project/{Pojectname}**|DELETE|Deleting all information by project
|**/api/report/last/{Number}**|GET|Getting information about last {.} operation
|**/api/adm/user/{Name}**|POST|Adding new user in system
|**/api/adm/user/{Name}**|GET|Get information abaut user in system by name
|**/api/adm/user/{all}**|GET|Get information abaut all users in system
|**/api/adm/user/{Name}**|DELETE|Delete user form system


## Reports 

|Path|Description report|
|--|--|
|**/report/log/{num}**|Preview last {**num**} operation all projects sorting by desc time|
|**/report/log/project/{name}/{num}**|Preview last count operation {**num**} in peoject {**name**}|
|**/report/log/project/{name}/{num}**|Preview last count operation {**num**} in peoject {**name**}|


## Objects

### Structure log

```
{
   Id           string // Id unig identificator 
   Operation    strig  // Operation in system
   Project      string // Project name
   Module       string // Module in system name
   Datetime     string // Timestamp Date and time operation
   Status       string // Critical, Info, Warn, Danger
   BlockId      string // Block Id
   AccountId    string // AccountID/ContractID
   CreateTime   string // Block CreationTime
}
```

### Цель: логирование параметров
- Структура проекта:
  корневой пакет   (gess)
  вложенные пакеты (core, miner etc.)
 
### Задача: 
Создать логирование состояния блокчейн в реальном времени с использованем RethinkDB
 
### Описание: 
Параметры должны передаваться на сервер RethinkDB, которы   их отображает

### Требования:
 1. запуск через параметр коммандной строки (./gess -rlog 111.222.333.444 )/111.222.333.444- ip- адрес сервера логирования
 2. Структура лога  {Project, Module, Ststus, BlockId, Block CreationTime, AccountID/ContractID}
 
### Tasks
1. Создать вложенный пакет логирования (rlog)
2. Интегрировать https://github.com/GoRethink/gorethink
3. Реализовать интрефейс rLogger  аналогично essentiaHybrid/log/logger.go - New , Info ,
4. Инициализировать логгер  essentiaHybrid/ess/peer.go     (rlog *rLogger )
5. Добавить вызовы rlog.Info(...) в еужные места по коду

### Statuses
|Status|Description|
|--|--|
|Info|Information|
|Warn|Warning|
|Dang|Danger|
|Critical|Critical|

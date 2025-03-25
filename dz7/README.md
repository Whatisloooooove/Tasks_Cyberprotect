# Задачи с использованием пакета sync и функции defer

## Структура

[goroutineSum](goroutineSum) - считает сумму непрерывного диапазона с помощью горутин 

[protectCounter](protectCounter) - защищенный с помощью `sync.Mutex` счётчик

[atomicCounter](atomicCounter) - защищенный с помощью `atomic` счётчик

[parallelArray](parallelArray) - подсчёт суммы массива с помощью горутин

[connectDB](connectDB) - пример использования `sync.Once` - подключение к базе данных

[condUsage](condUsage) - использование `sync.Cond` на примере взаимодействия двух горутин

[safeFileWrite](safeFileWrite) - безопасная запись в файл с использование функции `defer`

## Запуск

```
go run .
```
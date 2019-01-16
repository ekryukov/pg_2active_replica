# Deploy script

1. Global sequence

global_seq.sh - модификация всех сиквенсов в глобальный тип

2. Update sequence 

update_seq.go - докатка сиквенсов на заданное значение. Требуется для глобальных типов сиквенсов bdr, так как эти типы не могут менять текущее значение счетчика через alter sequence

Для компиляции требуется установленный go


Запуск:
```
./update_seq.bin имя_сиквенса счетчик
```

команда для кросс-платформенной сборки:

``` 
GOOS=linux GOARCH=amd64 go build  -o update_seq.bin update_seq.go
```

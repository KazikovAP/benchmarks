[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)

# benchmarks
# Измерение производительности бенчмарк-тестов

---
## Описание проекта
Проект представляет собой набор бенчмарк-тестов, которые покажут скорость работы разных версий одного и того же метода.

1) Сравнить производительность алгоритмов сортировки на массивах различной длины: стандартная сортировка, сортировка пузырьком

2) Создать функции для подсчёта количества непересекающихся вхождений подстроки в строку через strings.Count, регулярные выражения и посимвольную проверку

3) Сравнение скорости сериализации и десериализации данных по размеру объекта: 1КБ, 1МБ, 10МБ: бинарная сериализация, json, msgpack, jsoniter

4) Сравнить сложение элементов большого массива с использованием канала, mutex'a, WaitGroup'ы, atomic'a

---
## Технологии
* Go 1.23.0
* Tests

---
## Запуск проекта

**1. Клонировать репозиторий:**
```
git clone https://github.com/KazikovAP/benchmarks
```

**2. Запустить бенчмарк-тесты:**
```
go test -bench=. ./benchmarks
```

## Пример полученных результатов
```sh
goos: windows
goarch: amd64
pkg: github.com/KazikovAP/benchmarks/benchmarks
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz 
BenchmarkCount/Strings_6-8                               1378741             877.6 ns/op
BenchmarkCount/Regex_6-8                                  477261              2455 ns/op
BenchmarkCount/Manual_6-8                                  46789             26095 ns/op

BenchmarkCount/Strings_16-8                               116308             11953 ns/op
BenchmarkCount/Regex_16-8                                   2575            495042 ns/op
BenchmarkCount/Manual_16-8                                 20310             60624 ns/op

BenchmarkCount/Strings_26-8                                87836             14831 ns/op
BenchmarkCount/Regex_26-8                                   1381            773304 ns/op
BenchmarkCount/Manual_26-8                                 12843             98582 ns/op


BenchmarkSerialization/Serialize_JSON_1KB-8                32686             33266 ns/op
BenchmarkSerialization/Deserialize_JSON_1KB-8              19768             62559 ns/op
BenchmarkSerialization/Serialize_JSON_1MB-8                   19          58648942 ns/op
BenchmarkSerialization/Deserialize_JSON_1MB-8                 15          79176520 ns/op
BenchmarkSerialization/Serialize_JSON_10MB-8                   2         843351300 ns/op
BenchmarkSerialization/Deserialize_JSON_10MB-8                 2         804836050 ns/op

BenchmarkSerialization/Serialize_Binary_10MB-8                 4         280573750 ns/op
BenchmarkSerialization/Deserialize_Binary_10MB-8               4         298021975 ns/op
BenchmarkSerialization/Serialize_Binary_1KB-8              54769             20102 ns/op
BenchmarkSerialization/Deserialize_Binary_1KB-8            34862             35583 ns/op
BenchmarkSerialization/Serialize_Binary_1MB-8                 42          24623043 ns/op
BenchmarkSerialization/Deserialize_Binary_1MB-8               46          26642798 ns/op

BenchmarkSerialization/Serialize_MsgPack_1KB-8            150574              7613 ns/op
BenchmarkSerialization/Deserialize_MsgPack_1KB-8           75025             16733 ns/op
BenchmarkSerialization/Serialize_MsgPack_1MB-8               150           8573845 ns/op
BenchmarkSerialization/Deserialize_MsgPack_1MB-8              50          24227284 ns/op
BenchmarkSerialization/Serialize_MsgPack_10MB-8                7         153590571 ns/op
BenchmarkSerialization/Deserialize_MsgPack_10MB-8              4         278012575 ns/op

BenchmarkSerialization/Serialize_JSONIter_1KB-8           187365              6594 ns/op
BenchmarkSerialization/Deserialize_JSONIter_1KB-8          35815             36509 ns/op
BenchmarkSerialization/Serialize_JSONIter_1MB-8              165           7500587 ns/op
BenchmarkSerialization/Deserialize_JSONIter_1MB-8             20          54866155 ns/op
BenchmarkSerialization/Serialize_JSONIter_10MB-8               7         181247200 ns/op
BenchmarkSerialization/Deserialize_JSONIter_10MB-8             2         571717300 ns/op


BenchmarkSortingAlgorithms/StandardSort_1000-8             41348             29832 ns/op
BenchmarkSortingAlgorithms/BubbleSort_1000-8                1026           1183583 ns/op
BenchmarkSortingAlgorithms/StandardSort_10000-8             2940            411414 ns/op
BenchmarkSortingAlgorithms/BubbleSort_10000-8                 22          52227405 ns/op
BenchmarkSortingAlgorithms/StandardSort_100000-8             314           3807386 ns/op
BenchmarkSortingAlgorithms/BubbleSort_100000-8                 1        2815067800 ns/op


BenchmarkSumWithChannel/Channel_1000-8                      7992            156518 ns/op
BenchmarkSumWithChannel/Channel_10000-8                      774           1618324 ns/op
BenchmarkSumWithChannel/Channel_100000-8                      78          15880390 ns/op

BenchmarkSumWithMutex/Mutex_1000-8                         40596             29978 ns/op
BenchmarkSumWithMutex/Mutex_10000-8                         3298            344387 ns/op
BenchmarkSumWithMutex/Mutex_100000-8                         333           3612973 ns/op

BenchmarkSumWithWaitGroup/WaitGroup_1000-8                173791              6237 ns/op
BenchmarkSumWithWaitGroup/WaitGroup_10000-8                75895             15844 ns/op
BenchmarkSumWithWaitGroup/WaitGroup_100000-8               10000            101440 ns/op

BenchmarkSumWithAtomic/Atomic_1000-8                       98888             12559 ns/op
BenchmarkSumWithAtomic/Atomic_10000-8                       7963            151402 ns/op
BenchmarkSumWithAtomic/Atomic_100000-8                       789           1532389 ns/op

PASS
ok      github.com/KazikovAP/benchmarks/benchmarks      86.648s

```

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)

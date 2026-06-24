# algo-go

Мои решения алгоритмических задач на Go.

## Структура

```
algo-go/
├── arrays/
├── hash-map/
│   └── two-sum/
│       ├── main.go
│       ├── main_test.go
│       └── README.md
├── sliding-window/
├── linked-list/
│   └── reverse-linked-list/
│       ├── main.go
│       ├── main_test.go
│       └── README.md
├── stack/
│   └── valid-parentheses/
│       ├── main.go
│       ├── main_test.go
│       └── README.md
├── tree/
├── graph/
├── dynamic-programming/
└── binary-search/
    └── binary-search/
        ├── main.go
        ├── main_test.go
        └── README.md
```

## Задачи

### Хеш-таблицы
| Задача | Время | Память |
|--------|-------|--------|
| [Two Sum](./hash-map/two-sum/) | O(n) | O(n) |

### Скользящее окно
| Задача | Время | Память |
|--------|-------|--------|

### Связный список
| Задача | Время | Память |
|--------|-------|--------|
| [Reverse Linked List](./linked-list/reverse-linked-list/) | O(n) | O(1) |

### Стек
| Задача | Время | Память |
|--------|-------|--------|
| [Valid Parentheses](./stack/valid-parentheses/) | O(n) | O(n) |

### Деревья
| Задача | Время | Память |
|--------|-------|--------|

### Графы
| Задача | Время | Память |
|--------|-------|--------|

### Динамическое программирование
| Задача | Время | Память |
|--------|-------|--------|

### Двоичный поиск
| Задача | Время | Память |
|--------|-------|--------|
| [Binary Search](./binary-search/binary-search/) | O(log n) | O(1) |

## Запуск

```bash
# тесты конкретной задачи
cd hash-map/two-sum
go test -v

# тесты + бенчмарк
go test -v -bench=. -benchmem

# все тесты сразу
go test ./...
```
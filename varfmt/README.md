# varfmt

Реализуйте функцию `varfmt.Sprintf`. Функция принимает формат строку и переменное число аргументов.

Синтаксис формат-строки похож на формат-строки питона:
- `{number}` - ссылается на аргумент с индексом `number`
- `{}` - задаёт ссылку на аргумент с индексом, равным позиции `{}` среди всех ссылок (как `{}`, так и `{number}`) внутри паттерна 

Например, `varfmt.Sprintf("{1} {0}", "Hello", "World")` должен вернуть строку `World Hello`, а `varfmt.Sprintf("{0} {}", "Hello", "World")` должен вернуть строку `Hello World`.

Аргументы функции могут быть произвольными типами. Вам нужно форматировать их так же, как это
делает функция `fmt.Sprint`. Вызывать `fmt.Sprint` для форматирования отдельного аргумента
не запрещается.

Ваше решение будет сравниваться с baseline-решением на бенчмарке. Сравнение будет
проходить независимо по трем метрикам. 
  - `time/op` - время на одну итерацию бенчмарка
  - `alloc/op` - число выделенных байт на одну итерацию бенчмарка
  - `allocs/op` - число выделенных объектов на одну итерацию бенчмарка

Ваш код должен быть не более чем в два раза хуже чем baseline.

```
goos: linux
goarch: amd64
pkg: gitlab.com/slon/shad-go/varfmt
BenchmarkFormat/small_int-4         	 4744729	       263 ns/op	      64 B/op	       4 allocs/op
BenchmarkFormat/small_string-4      	 2388128	       484 ns/op	     168 B/op	       8 allocs/op
BenchmarkFormat/big-4               	    8997	    127827 ns/op	  194656 B/op	      41 allocs/op
BenchmarkSprintf/small-4            	13330094	        85.7 ns/op	       2 B/op	       1 allocs/op
BenchmarkSprintf/small_string-4     	 9351295	       123 ns/op	      16 B/op	       1 allocs/op
BenchmarkSprintf/big-4              	   12006	    108144 ns/op	   16392 B/op	       1 allocs/op
PASS
```

Для поиска лишних аллокаций используйте [`pprof`](../docs/allocation_profiling.md).

### Примеры

Как запустить все тесты и бенчмарки:
```
go test -v -bench=. ./varfmt/...
```

Как запустить только бенчмарки:
```
go test -v -run=^a -bench=. ./varfmt/...
```
Здесь `^a` - регулярное выражение, задающее тесты для запуска,
а `.` - задаёт бенчмарки.

Как запустить только big бенчмарки:
```
go test -v -run=^a -bench=/big ./varfmt/...
```

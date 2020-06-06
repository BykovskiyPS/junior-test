# Урок №2: Maps
## Ссылки для изучения
	
* https://medium.com/rungo/the-anatomy-of-maps-in-go-79b82836838b
* https://www.callicoder.com/golang-maps/
* https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html

## Изучаем исходники

* https://github.com/golang/go/blob/master/src/runtime/map.go
* https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/map_benchmark_test.go
* https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/map_test.go

## Задача

#### Изоморфные строки

Даны две строки **s** и **t**, необходимо определить изоморфны ли они.

Две строки считаются изоморфными если символы в **s** можно заменить чтобы получить **t**

Вы можете предположить, что длина **s** и **t** одинаковая.

##### Пример 1

```
Input: s = "egg", t = "add"
Output: true
```

##### Пример 2

```
Input: s = "foo", t = "bar"
Output: false
```

##### Пример 3

```
Input: s = "paper", t = "title"
Output: true
```


##### Шаблон

```Go
func isIsomorphic(s string, t string) bool {
    
}
```
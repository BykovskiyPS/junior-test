# Урок №1: Arrays and slices
## Ссылки для изучения
	
* https://www.sohamkamani.com/blog/golang/arrays-vs-slices/
* https://programming.guide/go/nil-slice-vs-empty-slice.html
* https://medium.com/rungo/the-anatomy-of-slices-in-go-6450e3bb2b94
* https://www.callicoder.com/golang-slices/
* https://yadi.sk/d/z9aQ0QGfo2c1sg
* https://github.com/golang/go/wiki/SliceTricks
* https://go101.org/article/memory-leaking.html

## Изучаем исходники

* https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/slice.go
* https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/slice_test.go

#### Cклонировать себе репозиторий и вне обучения изучить все ридми из корня репозитория

* https://github.com/quii/learn-go-with-tests/blob/master/arrays-and-slices.md
* https://github.com/quii/learn-go-with-tests/tree/master/arrays

## Задача

#### Две суммы

В полученном массиве целых чисел, вернуть **индексы** таких двух чисел, сумма которых равна заданному значению target.

Вы можете предположить, что у каждого варианта входных параметров есть **только одно** решение и вы не можете использовать один и тот же элемент дважды.

##### Пример

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

##### Шаблон

```Go
func twoSum(nums []int, target int) []int {
    
}
```
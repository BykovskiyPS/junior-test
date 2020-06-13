
# Урок №9: Изучаем работу сборщика мусора (garbage collector)
## Ссылки для изучения
* https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html
* https://blog.golang.org/ismmkeynote
* https://www.facebook.com/watch/?v=871606259683390
* https://habr.com/ru/post/265833/
* https://github.com/golang/go/blob/50bd1c4d4eb4fac8ddeb5f063c099daccfb71b26/src/runtime/debug/garbage.go
* https://github.com/golang/go/blob/7955ecebfc85851d43913f9358fa5f6a7bbb7c59/src/runtime/extern.go

## Исходники
* https://github.com/golang/go/blob/05511a5c0ae238325c10b0e4e44c3f66f928e4cf/src/runtime/mgc.go
  
## Задача
Переверните односвязный список.
* Связанный список может быть перевернут либо итеративно, либо рекурсивно. Можете реализовать оба варианта?

#### Пример 1 
```
 Input: 1->2->3->4->5->NULL
 Output: 5->4->3->2->1->NULL
```


## Шаблон 
``` Go
/** 
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 **/
func reverseList(head *ListNode) *ListNode {
    
}
```
## Задача 
Изучаем паттерн observer(наблюдатель) придумать и написать свою реализацию паттерна observer с тестом
https://github.com/AlexanderGrom/go-patterns/tree/master/Behavioral/Observer
  
# Урок №3: Channels
## Ссылки для изучения
	
* https://go101.org/article/101.html
* https://habr.com/ru/post/308070/
* https://github.com/quii/learn-go-with-tests/tree/master/select
* https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html

## Изучаем исходники

* https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/chan.go

* https://github.com/golang/go/blob/440f7d64048cd94cba669e16fe92137ce6b84073/src/runtime/chan_test.go
 
## Изучаем package oriented design
* https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html

* https://golang.org/pkg/context/

* https://habr.com/ru/company/nixys/blog/461723/
## Задача



Написать программу worker pool с использованием каналов, входные параметры - количество рабочих
 У нас есть рабочие, которые начинают работу параллельно (время работы - рандомное число от 1 до 3 секунд), затем они завершают работу и сдают свои результаты на проверку прорабу, который их принимает сделать
  бенчмарк тесты для решения https://habr.com/ru/post/268585/ 




```

#!/bin/bash

go build

 ./dev06 сын text.txt  > s21_grep.txt
 grep сын text.txt  > grep.txt
 diff -s s21_grep.txt grep.txt

 ./dev06 int task.go > s21_grep.txt
 grep int task.go > grep.txt
 diff -s s21_grep.txt grep.txt

 ./dev06 -c int task.go > s21_grep.txt
 grep -c int task.go > grep.txt
 diff -s s21_grep.txt grep.txt


 ./dev06 -v int task.go > s21_grep.txt
 grep -v int task.go > grep.txt
 diff -s s21_grep.txt grep.txt


./dev06 -A 2 Яна text.txt > s21_grep.txt
grep -A 2 Яна text.txt  > grep.txt
diff -s s21_grep.txt grep.txt


./dev06 -C 2 Яна text.txt   > s21_grep.txt
grep -C 2 Яна text.txt    > grep.txt
diff -s s21_grep.txt grep.txt

./dev06 -B  2 Яна text.txt > s21_grep.txt
grep  -B 2 Яна text.txt > grep.txt
diff -s s21_grep.txt grep.txt

./dev06 -A 2 func task.go > s21_grep.txt
grep -A 2 func task.go  > grep.txt
diff -s s21_grep.txt grep.txt


./dev06 -C 2 func task.go   > s21_grep.txt
grep -C 2 func task.go    > grep.txt
diff -s s21_grep.txt grep.txt

./dev06 -B 2 func task.go   > s21_grep.txt
grep -B 2 func task.go    > grep.txt
diff -s s21_grep.txt grep.txt


./dev06 -i яна text.txt > s21_grep.txt
grep -i яна text.txt  > grep.txt
diff -s s21_grep.txt grep.txt


./dev06 -n int task.go > s21_grep.txt
grep -n  int task.go  > grep.txt
diff -s s21_grep.txt grep.txt

./dev06 -F } task.go > s21_grep.txt
grep -F  } task.go  > grep.txt
diff -s s21_grep.txt grep.txt

rm s21_grep.txt grep.txt
rm -rf dev06
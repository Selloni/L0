#!/bin/bash

go build

echo "test - 1"
 ./dev06 сын text.txt  > s21_grep.txt
 grep сын text.txt  > grep.txt
 diff -s s21_grep.txt grep.txt

echo "test - 2"
 ./dev06 int task.go > s21_grep.txt
 grep int task.go > grep.txt
 diff -s s21_grep.txt grep.txt

echo "test - 3"
 ./dev06 -c int task.go > s21_grep.txt
 grep -c int task.go > grep.txt
 diff -s s21_grep.txt grep.txt

echo "test - 4"
 ./dev06 -v int task.go > s21_grep.txt
 grep -v int task.go > grep.txt
 diff -s s21_grep.txt grep.txt

echo "test - 5"
./dev06 -A 2 Яна text.txt > s21_grep.txt
grep -A 2 Яна text.txt  > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 6"
./dev06 -C 2 художнику text.txt   > s21_grep.txt
grep -C 2 художнику text.txt    > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 7"
./dev06 -B  2 Яна text.txt > s21_grep.txt
grep  -B 2 Яна text.txt > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 8"
./dev06 -A 2 func task.go > s21_grep.txt
grep -A 2 func task.go  > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 9"
./dev06 -C 2 import task.go   > s21_grep.txt
grep -C 2 import task.go    > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 10"
./dev06 -B 2 func task.go   > s21_grep.txt
grep -B 2 func task.go    > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 11"
./dev06 -i яна text.txt > s21_grep.txt
grep -i яна text.txt  > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 12"
./dev06 -n int task.go > s21_grep.txt
grep -n  int task.go  > grep.txt
diff -s s21_grep.txt grep.txt

echo "test - 13"
./dev06 -F } task.go > s21_grep.txt
grep -F  } task.go  > grep.txt
diff -s s21_grep.txt grep.txt

rm s21_grep.txt grep.txt
rm -rf dev06
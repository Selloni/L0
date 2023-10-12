#!/bin/bash

go build

 ./dev06 сын text.txt  > s21_grep.txt
 grep сын text.txt  > grep.txt
 diff -s s21_grep.txt grep.txt
 rm s21_grep.txt grep.txt

 ./dev06 int task.go > s21_grep.txt
 grep int task.go > grep.txt
 diff -s s21_grep.txt grep.txt
 rm s21_grep.txt grep.txt

 ./dev06 -c int task.go > s21_grep.txt
 grep -c int task.go > grep.txt
 diff -s s21_grep.txt grep.txt
 rm s21_grep.txt grep.txt


 ./dev06 -v int task.go > s21_grep.txt
 grep -v int task.go > grep.txt
 diff -s s21_grep.txt grep.txt
 rm s21_grep.txt grep.txt

rm -rf dev06
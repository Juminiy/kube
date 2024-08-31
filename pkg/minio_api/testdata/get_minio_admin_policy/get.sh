#!/bin/bash

cat $1 | grep -v "^$" | grep "admin:" > $1_modify.txt
sed -i 's/\(.*\):\(.*\)/Admin\2= "\1:\2"/' $1_modify.txt
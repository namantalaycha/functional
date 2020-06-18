#!/bin/bash

num=`grep "FAIL" -w  temp/text.txt `
echo $num
if [ "$num" !=  "" ]
then
        echo "Test FAIL"
        rm -r temp/text.txt
        exit 1
fi
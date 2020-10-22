#!/bin/sh
cd `awk -F " " -v prefix=$1 '$1 == prefix { print $2 }' < paths.txt`

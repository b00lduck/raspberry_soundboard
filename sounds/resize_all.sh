#!/bin/bash
for i in `find . -name "*.jpg" -o -name "*.png"`; do convert $i -resize 500x100 $i; done;

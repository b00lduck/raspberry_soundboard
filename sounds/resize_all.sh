#!/bin/bash
for i in `ls | grep ".jpg"`; do convert $i -resize 500x100 $i; done;
for i in `ls | grep ".png"`; do convert $i -resize 500x100 $i; done;
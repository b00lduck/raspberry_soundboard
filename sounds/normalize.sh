#!/bin/bash

normalize-mp3 --mp3encode="lame --quiet %w %m" --mp3decode="mpg123 -q -w %w %m" -a -18dBFS $1

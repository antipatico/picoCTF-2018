#!/bin/bash
rm comp.o flag.o -f
nasm -f elf comp.nasm
gcc -Wall flag.c comp.o -o flag

# GCC-less version.
#ld comp.o -lc -I /lib/ld-linux.so.2 -o comp

# Using non std write.
#gcc -Wall -c write.c
#ld comp.o write.o -lc -I /lib/ld-linux.so.2 -o comp

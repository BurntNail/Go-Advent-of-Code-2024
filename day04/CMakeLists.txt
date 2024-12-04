cmake_minimum_required(VERSION 3.30)
project(day04 C)

set(CMAKE_C_STANDARD 11)

add_executable(day04 main.c
        Board.c
        Board.h
        challenges.h
        PartOne.c)

#include "Board.h"

#include <stdio.h>
#include <assert.h>
#include <stdlib.h>

int createBoardFromFile(Board *toFillIn, const char* filename) {
    char ch;
    FILE* filePtr = fopen(filename, "r");
    if (filePtr == NULL) {
        return 1;
    }

    Letter* data = malloc(sizeof(Letter) * WIDTH * HEIGHT);
    if (data == NULL) {
        return 1;
    }

    int i = 0;
    while ((ch = fgetc(filePtr)) != EOF) {
        switch (ch) {
            case 'X':
                data[i] = X;
            break;
            case 'M':
                data[i] = M;
            break;
            case 'A':
                data[i] = A;
            break;
            case 'S':
                data[i] = S;
            break;
            default:
                continue;
        }
        i += 1;
    }

    toFillIn->data = data;
    toFillIn->size = WIDTH;

    fclose(filePtr);
    return 0;
}

Letter getIndex(const Board* board, const int row, const int col) {
    const int trueIndex = row * board->size + col;
    assert(trueIndex >= 0 && trueIndex < (WIDTH * HEIGHT));
    return board->data[trueIndex];
}

void cleanupBoard (const Board b) {
    free(b.data);
}
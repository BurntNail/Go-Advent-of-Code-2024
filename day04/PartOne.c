//
// Created by jack on 12/4/24.
//

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#include "Board.h"
#include "challenges.h"

int searchHoriz(const Board* b);
int searchVertical(const Board* b);
int searchDiagonal(const Board* b);

int partOne (const Board* b) {
    return searchHoriz(b) + searchVertical(b) + searchDiagonal(b);
}

int searchHoriz (const Board* b) {
    int found = 0;

    for (int row = 0; row < b->size; row++) {
        //search forwards
        for (int col = 0; col < (b->size - 3); col++) {
            Letter i = getIndex(b, row, col);
            Letter j = getIndex(b, row, col + 1);
            Letter k = getIndex(b, row, col + 2);
            Letter l = getIndex(b, row, col + 3);

            if (i == X && j == M && k == A && l == S) {
                found += 1;
            }
        }

        //search backwards
        for (int col = b->size - 1; col > 2; col--) {
            Letter i = getIndex(b, row, col);
            Letter j = getIndex(b, row, col - 1);
            Letter k = getIndex(b, row, col - 2);
            Letter l = getIndex(b, row, col - 3);

            if (i == S && j == A && k == M && l == X) {
                found += 1;
            }
        }
    }

    return found;
}

int searchVertical (const Board* b) {
    int found = 1;
    for (int col = 0; col < b->size; col++) {
        for (int row = 0; row < (b->size-3); row++) {
            Letter i = getIndex(b, row, col);
            Letter j = getIndex(b, row + 1, col);
            Letter k = getIndex(b, row + 2, col);
            Letter l = getIndex(b, row + 3, col);

            if (i == X && j == M && k == A && l == S) {
                found += 1;
            }
        }

        for (int row = b->size - 1; row > 2; row--) {
            Letter i = getIndex(b, row, col);
            Letter j = getIndex(b, row - 1, col);
            Letter k = getIndex(b, row - 2, col);
            Letter l = getIndex(b, row - 3, col);

            if (i == S && j == A && k == M && l == X) {
                found += 1;
            }
        }
    }

    return found;
}

int searchRopes (Letter** ropes, const int* individualLens, const int individualLensLen) {
    int found = 0;

    for (int ropeIndex = 0; ropeIndex < individualLensLen; ropeIndex++) {
        const int len = individualLens[ropeIndex];
        const Letter * rope = ropes[ropeIndex];

        for (int col = 0; col < (len - 3); col++) {
            Letter i = rope[col];
            Letter j = rope[col + 1];
            Letter k = rope[col + 2];
            Letter l = rope[col + 3];

            if (i == X && j == M && k == A && l == S) {
                found += 1;
            }
        }

        //search backwards
        for (int col = len - 1; col > 2; col--) {
            Letter i = rope[col];
            Letter j = rope[col - 1];
            Letter k = rope[col - 2];
            Letter l = rope[col - 3];

            if (i == S && j == A && k == M && l == X) {
                found += 1;
            }
        }
    }

    return found;
}

int searchDiagonal (const Board* b) {
    int found = 0;

    Letter** diagonalVersion = malloc(sizeof(Letter*) * b->size * b->size);
    int* individualLens = malloc(sizeof(int) * (b->size * 2));

    if (diagonalVersion == NULL || individualLens == NULL) {
        return -1;
    }

    int thisRopeLen = 0;
    bool goingDown = false;
    for (int k = 0; k < (b->size * 2); k++) {
        if (thisRopeLen < (b->size) && !goingDown) {
            thisRopeLen++;
        } else {
            goingDown = true;
            thisRopeLen--;
        }

        Letter* thisRope = malloc(sizeof(Letter) * thisRopeLen);
        int thisRopeIndex = 0;
        //thx https://stackoverflow.com/questions/20420065/loop-diagonally-through-two-dimensional-array
        for (int j = 0; j <= k; j++) {
            int i = k - j;
            if (i < b->size && j < b->size) {
                thisRope[thisRopeIndex] = getIndex(b, i, j);
                thisRopeIndex += 1;
            }
        }

        if (thisRopeIndex != thisRopeLen) {
            printf("Expected len %d, found %d\n", thisRopeLen, thisRopeIndex);
        }

        diagonalVersion[k] = thisRope;
        individualLens[k] = thisRopeIndex;
    }
    found += searchRopes(diagonalVersion, individualLens, b->size * 2);

    for (int i = 0; i < (b->size * 2); i++) {
        free(diagonalVersion[i]);
    }
    free(diagonalVersion);
    free(individualLens);


    diagonalVersion = malloc(sizeof(Letter*) * b->size * b->size);
    individualLens = malloc(sizeof(int) * (b->size * 2));

    if (diagonalVersion == NULL || individualLens == NULL) {
        return -1;
    }

    thisRopeLen = 0;
    goingDown = false;
    for (int k = 0; k < (b->size * 2); k++) {
        if (thisRopeLen < (b->size) && !goingDown) {
            thisRopeLen++;
        } else {
            goingDown = true;
            thisRopeLen--;
        }

        Letter* thisRope = malloc(sizeof(Letter) * thisRopeLen);
        int thisRopeIndex = 0;
        //thx https://stackoverflow.com/questions/20420065/loop-diagonally-through-two-dimensional-array
        for (int j = 0; j <= k; j++) {
            int i = k - j;
            if (i < b->size && j < b->size) {
                thisRope[thisRopeIndex] = getIndex(b, i, b->size - j - 1);
                thisRopeIndex += 1;
            }
        }

        if (thisRopeIndex != thisRopeLen) {
            printf("Expected len %d, found %d\n", thisRopeLen, thisRopeIndex);
        }

        diagonalVersion[k] = thisRope;
        individualLens[k] = thisRopeIndex;
    }
    found += searchRopes(diagonalVersion, individualLens, b->size * 2);

    for (int i = 0; i < (b->size * 2); i++) {
        free(diagonalVersion[i]);
    }
    free(diagonalVersion);
    free(individualLens);



    return found;
}

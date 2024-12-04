#include <stdio.h>

#include "Board.h"
#include "challenges.h"

int main(void) {
    Board b;
    if (createBoardFromFile(&b, "../input.txt") != 0) {
        printf("oops");
        return 1;
    }


    int pone = partOne(&b);
    printf("%d\n", pone);

    cleanupBoard(b);
    return 0;
}

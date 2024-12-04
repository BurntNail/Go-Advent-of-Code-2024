#ifndef BOARD_H
#define BOARD_H

//yay no easy dynamic stuff!!!
#define WIDTH 10
#define HEIGHT WIDTH

typedef enum Letter {
    X,M,A,S
} Letter;

typedef struct Board {
    Letter* data;
    int size;
} Board;

int createBoardFromFile (Board *toFillIn, const char* filename);
Letter getIndex(const Board* board, const int row, const int col);
void cleanupBoard (Board b);

#endif //BOARD_H

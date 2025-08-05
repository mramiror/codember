#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define MAX_LINE  1024

enum State { START, DIGIT, CHAR };

bool is_valid(const char *pwd) {
    enum State state = START;
    char last_digit = '0';
    char last_char = 'a';

    for (size_t i = 0; pwd[i] != '\0'; i++) {
        char c = pwd[i];

        if (c >= '0' && c <= '9') {
            if (state == CHAR) {
                return false;
            }
            state = DIGIT;
            if (c < last_digit) {
                return false;
            }
            last_digit = c;
        } else if (c >= 'a' && c <= 'z') {
            state = CHAR;
            if (c < last_char) {
                return false;
            }
            last_char = c;
        } else {
            return false;
        }
    }
    return true;
}

int main(void) {
    FILE *fp = fopen("log.txt", "r");
    if (!fp) {
        perror("Failed to open log.txt");
        return EXIT_FAILURE;
    }

    char line[MAX_LINE];
    int valid = 0, invalid = 0;

    while (fgets(line, sizeof(line), fp)) {
        // Remove newline character if present
        line[strcspn(line, "\r\n")] = '\0';
        if (line[0] == '\0') {
            continue;
        }
        if (is_valid(line)) {
            valid++;
        } else {
            invalid++;
        }
    }

    fclose(fp);

    printf("submit %dtrue%dfalse\n", valid, invalid);
    return EXIT_SUCCESS;
}
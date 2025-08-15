#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define MAX_LINE  1024

int steps_to_exit(int *nums, int len) {
    int steps = 0;
    int pos = 0;

    while (pos >= 0 && pos < len) {
        int jump = nums[pos];
        nums[pos]++;
        pos += jump;
        steps++;
    }

    return steps;
}

int main(void) {
    FILE *fp = fopen("trace.txt", "r");
    if (!fp) {
        perror("Failed to open trace.txt");
        return EXIT_FAILURE;
    }

    char line[MAX_LINE];
    int total_steps = 0, steps_last_line = 0;

    while (fgets(line, sizeof(line), fp)) {
        line[strcspn(line, "\r\n")] = '\0';
        if (line[0] == '\0') {
            continue;
        }

        int nums[MAX_LINE];
        int count = 0;

        char *token = strtok(line, " ");
        while (token != NULL) {
            nums[count++] = atoi(token);
            token = strtok(NULL, " ");
        }

        int temp[MAX_LINE];
        memcpy(temp, nums, count * sizeof(int));

        int steps = steps_to_exit(temp, count);
        total_steps += steps;
        steps_last_line = steps;
    }

    fclose(fp);

    printf("submit %d-%d\n", total_steps, steps_last_line);
    return EXIT_SUCCESS;
}
// 単純交換（バブル）ソート

#include <stdio.h>
#include <stdlib.h>

#define swap(type, x, y) do { type t = x; x = y; y = t; } while (0)

void bubble(int a[], int n)
{
    int i, j;

    for (i = 0; i < n - 1; i++) {
        int exchange = 0;
        for (j = n; j > i; j--) {
            if (a[j - 1] > a[j]) swap(int, a[j - 1], a[j]);
            exchange++;
        }
        if (exchange == 0) break;
    }
}

int main(void)
{
    int i, nx;
    int *x;

    puts("単純交換ソート");
    printf("要素数 : ");
    scanf("%d", &nx);
    x = calloc(nx, sizeof(int));
    for (i = 0; i < nx; i++) {
        printf("x[%d] : ", i);
        scanf("%d", &x[i]);
    }

    bubble(x, nx);

    puts("昇順ソートしました.");
    for (i = 0; i < nx; i++) {
        printf("x[%d] : %d\n", i, x[i]);
    }

    free(x);

    return 0;
}
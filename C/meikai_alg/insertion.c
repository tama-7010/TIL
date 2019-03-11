#include <stdio.h>
#include <stdlib.h>

void insertion(int a[], int n) {
    int i, j;

    for (i = 1; i < n; i++) {
        int tmp = a[i]; // 着目する要素をtmpに代入する
        for (j = i; j > 0 && a[j - 1] > tmp; j--) 
            // 配列の最後より進んでいない && 着目する要素の隣の値が着目する要素より大きいあいだ
            // 要素をコピーする
            a[j] = a[j - 1];
        // コピーし終わったらtmpを代入する
        a[j] = tmp;
    }
}

int main(void) {
    int i, nx;
    int *x;

    puts("単純挿入ソート");
    printf("要素数 : ");
    scanf("%d", &nx);
    x = calloc(nx, sizeof(int));
    for (i = 0; i < nx; i++) {
        printf("x[%d] : ", i);
        scanf("%d", &x[i]);
    }

    insertion(x, nx);

    puts("昇順ソートしました.");
    for (i = 0; i < nx; i++) {
        printf("x[%d] : %d\n", i, x[i]);
    }

    free(x);

    return 0;
}
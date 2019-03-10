/*
2分探索(binary search)
要素がキーの昇順または降順にソートされている配列から効率よく探索を行う。
配列の中央の要素を決定し、等しい場合は探索完了となる。
決定した値より探索している値が小さければ左側の要素を再度探索、大きければ右側の要素を再度探索する。
*/

#include <stdio.h>
#include <stdlib.h>

// 要素数nの配列aからkeyと一致する要素を2分探索
int bin_search(const int a[], int n, int key)
{
    int pl = 0; // 探索範囲の先頭
    int pr = n - 1; // 探索範囲の末尾
    int pc; // 探索範囲の中央

    do {
        pc = (pl + pr) / 2;
        if (a[pc] == key) return pc;
        else if (a[pc] < key) pl = pc + 1;
        else pr = pc - 1;
    } while (pl <= pr);

    return -1;
}

int main(void){
    int i, nx, ky, idx;
    int *x;

    puts("2分探索");
    printf("要素数 : ");
    scanf("%d", &nx);
    x = calloc(nx, sizeof(int));

    printf("昇順に入力する。");
    printf("x[0] : ");
    scanf("%d", &x[0]);

    for (i = 1; i < nx; i++) {
        do {
            printf("x[%d] : ", i);
            scanf("%d", &x[i]);
        } while (x[i] < x[i - 1]);
    }

    printf("探す値 : ");
    scanf("%d", &ky);

    idx = bin_search(x, nx, ky);

    if (idx == -1) puts("探索失敗");
    else printf("%dはx[%d]にあります\n", ky, idx);
    free(x);
 
    return 0;
}
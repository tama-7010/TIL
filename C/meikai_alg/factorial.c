// 階乗を再帰によって求める

#include <stdio.h>

int factorial(int n)
{
    if (n > 0) return n * factorial(n - 1);
    else return 1;
}

int main(void)
{
   int x;

   printf("整数を入力 : ");
   scanf("%d", &x);

   printf("%dの階乗は%dです.\n", x, factorial(x));

   return 0;
}
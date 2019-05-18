#include <stdio.h>

int main() {

  int n, pessoas, cont = 0;
  scanf("%d\n", &n);
  for(int i = 0; i < n; i++) {
    scanf("%d", &pessoas);
    if(pessoas == 0) {
      cont++;
    }
  }

  if(cont > (n - cont)) {
     printf("Y\n");
  } else {
    printf("N\n");
  }

  return 0;
}

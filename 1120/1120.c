#include <stdio.h>
#include <string.h>

int main() {
  char n1;
  char n2[9999], n3[9999];
  scanf("%s %s", &n1, n2);
  while(n1 != '0' && n2 != "0") {
    int i = 0, j = 0;
    while(i < strlen(n2)) {
      if(n2[i] != n1) {
        n3[j] = n2[i];
        n3[j + 1] = '\0';
        j++;
        i++;
      } else {
        while (i != strlen(n2) && (n2[i] == '0' || n2[i] == n1)) {
          i++;
        }
      }
    }

    if(j == 0) {
      n3[0] = '0';
      n3[j + 1] = '\0';
      j++;
    }
    printf("%s\n", n3);
    scanf("%s %s", &n1, n2);
  }

  return 0;
}

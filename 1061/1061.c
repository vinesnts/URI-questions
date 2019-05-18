#include <stdio.h>

int main() {

  int diaInicial[4], diaFinal[4], tempoInicial, tempoFinal, tempoTotal;
  char dia;

  scanf("%s %d", &dia, &diaInicial[0]);
  scanf("%d : %d : %d", &diaInicial[1], &diaInicial[2], &diaInicial[3]);

  scanf("%s %d", &dia, &diaFinal[0]);
  scanf("%d : %d : %d", &diaFinal[1], &diaFinal[2], &diaFinal[3]);

  tempoInicial = diaInicial[0]*24*60*60
                + diaInicial[1]*60*60
                + diaInicial[2]*60
                + diaInicial[3];
  tempoFinal = diaFinal[0]*24*60*60
                + diaFinal[1]*60*60
                + diaFinal[2]*60
                + diaFinal[3];

  tempoTotal = tempoFinal - tempoInicial;

  printf("%d dia(s)\n", tempoTotal/(24*60*60));
  tempoTotal %= (24*60*60);
  printf("%d hora(s)\n", tempoTotal/(60*60));
  tempoTotal %= (60*60);
  printf("%d minuto(s)\n", tempoTotal/60);
  tempoTotal %= 60;
  printf("%d segundo(s)\n", tempoTotal);

  return 0;

}

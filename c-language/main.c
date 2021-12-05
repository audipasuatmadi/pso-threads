#include <stdio.h>
#include <sched.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>

int nantiFunctionIniBerjalanDiThreadLain()
{
  printf("[THREAD LAIN] halo bos dari thread lain\n");
  sleep(2);
  printf("[THREAD LAIN] goodbye bos dari thread lain\n");
  return 0;
}

int main()
{
  void *pchildStack = malloc(1024 * 1024);
  if (pchildStack == NULL)
  {
    printf("gagal alokasi memori\n");
    return 1;
  }

  int pid = clone(&nantiFunctionIniBerjalanDiThreadLain, pchildStack + (1024 * 1024), SIGCHLD);
  if (pid < 0)
  {
    printf("gagal membuat thread baru\n");
    return 1;
  }
  printf("[THREAD MAIN] thread dijalankan\n");
  sleep(1);
  printf("[THREAD MAIN] gass\n");

  sleep(1);
  free(pchildStack);
  printf("[THREAD MAIN] done\n");
  return 0;
}
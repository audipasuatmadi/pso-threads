package com.audipasuatmadi.pso.main;

public class UserImplementRunnable implements Runnable {
  private String name;
  private int age;
  private long durasiTunggu;

  public UserImplementRunnable(String name, int age, long durasiTunggu) {
    this.name = name;
    this.age = age;
    this.durasiTunggu = durasiTunggu;
  }

  @Override
  public void run() {
    System.out.printf("[THREAD %s] Namaku %s\n", this.name, this.name);
    try {
      Thread.sleep(this.durasiTunggu);
    } catch (InterruptedException e) {
      System.out.printf("[THREAD %s] terjadi kesalahan dalam memanggil Thread.sleep\n", this.name);
    }
    System.out.printf("[THREAD %s] Umurku %d\n", this.name, this.age);
  }
}

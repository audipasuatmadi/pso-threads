package com.audipasuatmadi.pso.main;

public class Main {
  public static void main(String[] args) throws InterruptedException {
    System.out.println("[MAIN THREAD]: hello ser! program mulai jalan");
    UserExtendThread userThread1 = new UserExtendThread("Seseorang A", 2, 400);
    UserExtendThread userThread2 = new UserExtendThread("Seseorang B", 3, 600);
    userThread1.start();
    userThread2.start();
    System.out.println("[MAIN THREAD]: main thread menunggu 0.5 detik");
    Thread.sleep(500);
    System.out.println("[MAIN THREAD]: main thread menunggu 0.5 detik lagi");
    Thread.sleep(500);
    System.out.println("[MAIN THREAD]: main thread mulai hidup lagi");
    System.out.printf("[MAIN THREAD]: program selesai jalan ser!");
  }
}

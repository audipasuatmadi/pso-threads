import threading
import time

def print_dengan_nama_fungsi(nama, msg):
  print("[{}]: {}".format(nama, msg))

def fungsi_pada_thread(nama, milisecond):
  print_dengan_nama_fungsi(nama, "memulai fungsi pada thread")
  print_dengan_nama_fungsi(nama, "mulai time.sleep untuk menunggu")
  time.sleep(milisecond/1000)
  print_dengan_nama_fungsi(nama, "fungsi selesai")

thread1 = threading.Thread(target=fungsi_pada_thread, args=("Thread 1", 500))
thread1.start()
thread1 = threading.Thread(target=fungsi_pada_thread, args=("Thread 2", 400))
thread1.start()
async function wait(milisecond) {
  return new Promise((res, rej) => setTimeout(res, milisecond));
} 

function printDenganNamaFunction(namaFunction, msg) {
  console.log("[" + namaFunction + "]: " + msg)
}

async function contohAsyncFunction(namaAsyncFunction, milisecond) {
  printDenganNamaFunction(namaAsyncFunction, "async function mulai berjalan");
  printDenganNamaFunction(namaAsyncFunction, "async function mulai menunggu");
  await wait(milisecond);
  printDenganNamaFunction(namaAsyncFunction, "async function seleasaii menunggu");
  printDenganNamaFunction(namaAsyncFunction, "async function seleasai");
}

contohAsyncFunction("Async 1", 500);
contohAsyncFunction("Async 2", 400);

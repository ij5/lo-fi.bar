const cacheName = "app-" + "5b854139d57743ce2a03ecb2fe87a157fbe1656d";

self.addEventListener("install", event => {
  console.log("installing app worker 5b854139d57743ce2a03ecb2fe87a157fbe1656d");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/",
          "/app.css",
          "/app.js",
          "/manifest.webmanifest",
          "/wasm_exec.js",
          "/web/app.css",
          "/web/app.wasm",
          "https://filedn.com/ljcU5RLfox2Sbr04VfCcCX5/lo-fi.png",
          "https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
      })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 5b854139d57743ce2a03ecb2fe87a157fbe1656d is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

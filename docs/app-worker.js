const cacheName = "app-" + "f132d9882fe7df5d255205180cd6beb616a3a7f1";

self.addEventListener("install", event => {
  console.log("installing app worker f132d9882fe7df5d255205180cd6beb616a3a7f1");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/",
          "/app.css",
          "/app.js",
          "/docs/web/app.css",
          "/manifest.webmanifest",
          "/wasm_exec.js",
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
  console.log("app worker f132d9882fe7df5d255205180cd6beb616a3a7f1 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

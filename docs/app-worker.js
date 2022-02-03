const cacheName = "app-" + "181db87b94a9361c3b465f07b2fa8412da70e2c0";

self.addEventListener("install", event => {
  console.log("installing app worker 181db87b94a9361c3b465f07b2fa8412da70e2c0");

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
  console.log("app worker 181db87b94a9361c3b465f07b2fa8412da70e2c0 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});

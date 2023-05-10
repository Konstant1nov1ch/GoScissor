# GoScissor

GoScissor is a simple web API service for trimming links (currently in development).

Example: http://localhost:8080/admin/tokens

![Example screenshot](https://github.com/Konstant1nov1ch/GoScissor/assets/105445251/7dfbfcef-aa0a-4475-b8c7-fabd67c15250.png)

At the moment, the service supports 3 types of requests:

- `GET /admin/tokens` - list tokens from database
- `GET /:short_url` - redirect by short link
- `POST /sci` - generate short_url from full_url
- 'Body:{"full_url": "https://loooong/url/for/example.html"}'


I added a two-level algorithm caches 2Q on a screenshot can be seen that the difference redirect when you call from the cache is 31 times faster (I think it's a success)!

Pruf: 

![image](https://github.com/Konstant1nov1ch/GoScissor/assets/105445251/e50d6bd4-7553-4943-98b4-c7e581ae7a91)

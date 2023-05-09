# GoScissor

GoScissor is a simple web API service for trimming links (currently in development).

Example: http://localhost:8080/admin/tokens

![Example screenshot](https://github.com/Konstant1nov1ch/GoScissor/assets/105445251/7dfbfcef-aa0a-4475-b8c7-fabd67c15250.png)

At the moment, the service supports 3 types of requests:

- `GET /admin/tokens` - list tokens from database
- `GET /:short_url` - redirect by short link
- `POST /admin/tokens` - generate short_url from full_url
- 'Body:{"full_url": "https://t.me/Algoru_bot"}'
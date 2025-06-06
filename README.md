# disweather

## Discord Weather bot

<img src="./docs/pictures/profile.jpg" width=50% height=50%>

> [image source](https://www.pixiv.net/en/artworks/129763395)

### Docker

Copy `.env.example` to `.env`

```
docker build -t disweather:latest .
```

```
docker run -d --name disweather --restart=always disweather:latest
```

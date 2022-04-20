docker build -t oj:v1 .  --no-cache
docker run -p 8000:8000 -d --name myoj oj:v1
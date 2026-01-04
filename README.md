# bas-backend

```

 # Билд имеджа для вм
 
 
docker buildx build \
  --platform linux/amd64 \
  -t bas-backend:prod-amd64 \
  --load \
  .
  
  
  
   Сохраняем как тар файл имедж
   
   docker save -o bas-backend-prod.tar bas-backend:prod-amd64
   
   
   разархивируем имедж
   
   docker load -i bas-backend-prod.tar
   
   правим докер композ на имедж и запускаем
  
```
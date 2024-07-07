
## デプロイメモ

### ENV
```.bash_profile
export DB_USERNAME=
export DB_PASSWORD=
export DB_HOST=
export DB_NAME=

export GMAIL_APP_PASSWORD=
```

### バックエンド
```
cd backend
go build main.go
nohup ./main &
```

### フロントエンド
```
$ cd frontend
$ npm install
$ npm run build
$ sudo rm -rf /usr/share/nginx/html/*
$ sudo cp -r /home/ec2-user/QRmark/frontend/build/* /usr/share/nginx/html/
$ sudo systemctl restart nginx
```

### Nginx
/etc/nginx/conf.d/default.conf
```
server {
    listen       80;
    listen  [::]:80;

    location /api/ {
        proxy_pass http://localhost:8080/;
    }

    location / {
        root   /usr/share/nginx/html;
        try_files $uri $uri/ /index.html;
    }
}
```

# QRコード型ベルマーク「QRmark」

<a href="https://ibukiqrmark.com">QRmark</a>

## もしベルマークがQRコードだったら...って思い作りました。

## 機能一覧
| トップ画面                                                                                           | ログイン後のトップ画面                                                               |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| ![トップ画面](https://github.com/v420v/QRmark/assets/106643445/a809098f-d513-4913-93e1-06767c6d1436) | ![ログイン後のトップ画面](https://github.com/v420v/QRmark/assets/106643445/12f5e644-3e3b-4352-8251-529f905d7ce8)
| ログイン画面、新規登録画面に遷移できる様になっています。                                                     | QRマーク履歴が表示されます。「QRマークをスキャンする」ボタンでカメラが開きます |


| 新規登録画面                                                                                           | ログイン画面                                                                     |
| --------------------------------------------------------------------------------------------------   | ------------------------------------------------------------------------------ |
| ![新規登録画面](https://github.com/v420v/QRmark/assets/106643445/06f19a26-e4f0-40d4-baa2-733750be675a) | ![ログイン画面](https://github.com/v420v/QRmark/assets/106643445/dd7c1e2f-22fd-4cc0-ad7f-39421eea24d8)
| ユーザーの新規登録                                                                                      | ユーザーのログイン                                |

| 学校のポイント集計結果画面                                                                                           | 学校検索画面                                                          |
| --------------------------------------------------------------------------------------------------   | ------------------------------------------------------------------------------ |
| ![学校のポイント集計結果画面](https://github.com/v420v/QRmark/assets/106643445/d9255a69-5727-41ae-b6c4-bca9f8145a8c) | ![学校検索画面](https://github.com/v420v/QRmark/assets/106643445/79aaa3d1-c0ff-4732-a011-ad60acdaf495)
| 学校のその月のポイント集計結果が表示されます                                                                                     | 学校を検索                                            |


## ER図
![sql](https://github.com/v420v/QRmark/assets/106643445/1fd13029-0e08-4927-8d16-f5fd0dd98745)

## AWS 構成図
<img width="771" alt="Screenshot 2024-07-04 at 15 48 30" src="https://github.com/v420v/QRmark/assets/106643445/0ed84a76-cda9-4af1-ae63-2fb79f8a82d5">



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

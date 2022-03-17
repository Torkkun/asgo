#asgo re:
時間あるとき進める

##後で状況を思い出す用のめも
* docker container
    * discordbot
    * apiserver
        * firebase authentication(user management)
    * selenium
        * `sudo docker pull selenium/standalone-firefox` にてseleniumのdockerimageを拾ってくる
        * `docker run -d -p 4444:4444 -p 7900:7900 --shm-size="2g" selenium/standalone-firefox` でrunする
        * コード内でremoteアクセスすれば良い神 実行状況を見る場合は、http://localhost:4444 にアクセスする
        詳しくはhttps://github.com/SeleniumHQ/docker-selenium

    * (DB postgreSQL)

後にユーザー毎にコンフィグを設定できるようにしたい現状はログインページのみとする

適当
　一日またはすべてのガチャが轢き終わったらその結果をまたはキャッシュして残すreddis
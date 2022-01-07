#asgo re:
時間あるとき進める

##後で状況を思い出す用のめも
* docker container
    * discordbot
    * apiserver
        * firebase authentication(user management)
    * selenium
        * `sudo docker pull selenium/standalone-firefox` にてseleniumのdockerimageを拾ってくる
        * コード内でremoteアクセスすれば良い神 実行状況を見る場合は、http://localhost:4444 にアクセスする
        詳しくはhttps://github.com/SeleniumHQ/docker-selenium

    * (DB postgreSQL)

後にユーザー毎にコンフィグを設定できるようにしたい現状はログインページのみとする

基本的な流れ
ユーザー登録
サイト自体のログインの後
クローリングするサイトのパスワードとメールアドレスが必要パスワードを可逆暗号で暗号化してDBに保存

discordで自分のデータを取得するコマンドを入力するとデータが帰ってくる
実行可能であれば通知、それ以外はデータだけ
実行可能だった場合は現状任意でユーザーに実行させる

毎日決まった時間にデイリーが実行され結果がdiscordに帰ってくる

clean architecture によりDBとseleniumライブラリを独立させたい！

tables

gatyadata
{
    daily_point int
    sum_point int
    bonus_ticket int
    until_bonus int
    exchange_ticket int
    enquete bool
    user_id string
    updated_at time.Time
}

user
{
    user_id string
    username string
    password string
    firebase_uuid string
    client string
}
#asgo re:
時間あるとき進める

##めも
* docker container
    * discordbot
    * apiserver
        * firebase authentication(user management)
    * selenium
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

tables

gatyadata
{
    daily_point
    sum_point
    bonus_ticket
    bonus_weekDue
    user_id
    updated_at
}

user
{

}
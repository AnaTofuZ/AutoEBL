# AutoEBL

 Auto Extension of Boll Lending tool
 自動で図書貸出やってくれる君

# SYNOPSIS

## Download

```sh
$ go get github.com/AnaTofuZ/AutoEBL
```



```sh
$ brew tap anatofuz
$ brew install anatofuz/autoebl
```

## Usage

``` sh
$ autoebl -u e155730 -p XXXXXX
or
$ AUEBL_USERNAME=e155730 AUEBL_PASSWORD="XXX" autoebl
or 
writing configure
$ autoebl
```

# DESCRIPTION

基本的には全自動でやってくれますが，usernameとpasswordは何かしらの方法で設定する必要があります．
AutoEBLは,設定ファイル,環境変数,コマンドラインの順でusernameなどを確認していきます.
最後に書かれている設定の方が優先されます．

chromeが勝手に立ち上がりますが，頬って置けばなんとかなります

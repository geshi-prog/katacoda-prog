## ubuntuのイメージをPULLする。

`docker pull ubuntu:latest`{{execute}}<br>

## PULLしてきたイメージを使用してjuliaの環境を構築してみる

`docker images`{{execute}}<br>
`docker run --name julia-ubuntu -u 0 -itd f63181f19b2f /bin/sh`{{execute}}<br>
`docker ps`{{execute}}<br>
`docker exec -it b2a1099ea753 /bin/sh`{{execute}}<br>
`apt update`{{execute}}<br>
`apt -y install build-essential wget`{{execute}}<br>
`wget https://julialang-s3.julialang.org/bin/linux/x64/1.0/julia-1.0.5-linux-x86_64.tar.gz`{{execute}}<br>
`ls julia*`{{execute}}<br>
`tar xvfz julia-1.0.5-linux-x86_64.tar.gz`{{execute}}<br>
`ls julia*`{{execute}}<br>
`ln -s /julia-1.0.5/bin/julia /usr/bin/julia`{{execute}}<br>
`julia`{{execute}}<br>
`@which sin(1.0)`{{execute}}<br>
## 動画を流してみる

<iframe width="560" height="315" src="https://www.youtube.com/embed/tAlggq5dT3M" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

## 画像を表示してみる

![Sammy Campaign](./assets/001.jpg)

## クイズを実装する
>>Q1: テストと正しく入力してください<<
=== テスト

>>Q2: juliaで「@which cos(1.0)」を実行した際の出力を入力してください<<
=== cos(x::T) where T<:Union{Float32, Float64} in Base.Math at special/trig.jl:100

>>Q3: テストを含む文字を入力してください<<
=~= テスト

>>Q4: 正しい物を以下から全て選んでください<<
[*] 正しい
[*] 正しい
[ ] 間違い

>>Q5: 正しい方のどちらか一つを選んでください<<
(*) 正しい
( ) 間違い

解答後は下のボタンを押してください！
全問に正解した場合のみ次のステップに進めます。
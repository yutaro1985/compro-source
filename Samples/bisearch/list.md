# 二分探索のメモ

- [二分探索のメモ](#二分探索のメモ)
  - [メモ](#メモ)
  - [ABC](#abc)
    - [ABC 146 C](#abc-146-c)
    - [ABC 172 C](#abc-172-c)
    - [ABC 174 E](#abc-174-e)

## メモ

golangでは`sort.Search`が使える。  
[sort.Search](https://xn--go-hh0g6u.com/pkg/sort/#Search)  
sortパッケージだが、別にスライスが必要なわけではない。  
sort.Searchを使うものをメインにまとめているが、そうでないものは基本的に[めぐる式二分探索](https://twitter.com/meguru_comp/status/697008509376835584)を使っている。  
sort.SearchではOKが大きい側の場合（条件を満たす最小の値）が求まるので、  
小さい側（条件を満たす最大値）を求めたいときはN+1まで見ないといけないことがある、かも（要検証）

## ABC

### ABC 146 C

[C - Buy an Integer](https://atcoder.jp/contests/abc146/tasks/abc146_c)  
[コード](../../AtCoder/ABC/abc146/c/C_sort_search.go)  
※N+1までを見ないとだめな例  
[参考：sort.Searchを使わないコード](../../AtCoder/ABC/abc146/c/C.go)

### ABC 172 C

[C - Tsundoku](https://atcoder.jp/contests/abc172/tasks/abc172_c)  
[コード](./sort_search_binary.go)  
[参考：sort.Searchを使わない二分探索](../../AtCoder/ABC/abc172/c/C.go)  
[参考：二分探索でない解き方](../../AtCoder/ABC/abc172/c/C_editorial.go)

### ABC 174 E

[E - Logs](https://atcoder.jp/contests/abc174/tasks/abc174_e)  
[コード](../../AtCoder/ABC/abc174/e/E.go)  
※i==0のときにゼロ除算になってしまうが、条件から0になることはありえないのでfalseにすることに注意

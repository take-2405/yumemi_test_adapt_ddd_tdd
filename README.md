# yumemi-coding-test-practice
ゆめみさんのコーディングテストを解いてみる

### リポジトリの説明  
ゆめみさんのコーディンテストを解きつつ，最近学んだTDDとDDDを意識して取り組む  
なお、今回はドメインがはっきりしていないため、アーキテクチャを意識するという意味でDDDに取り組む

### アーキテクチャ
- オニオンアーキテクチャ
```
.
├── application
│   └── rankingData.go
├── domain
│   ├── playData.go
│   ├── rankingData.go
│   └── repository
│       ├── playData_repository.go
│       └── rankingData_repository.go
├── infrastructure
│   ├── playData.go
│   └── rankingData.go
└── presentation
    ├── hander
    │   └── ranking.go
    └── response
        └── ranking.go

```

##### 依存関係  
```
presentation  
↓  
application  
↓  
domain 
```
```
infrastructure
↓  
domain 
```



### 取り組んだ問題
https://www.yumemi.co.jp/serverside_recruit

## 実行方法
```cassandraql
go run main.go test.csv
```
test.csv=使用するcsvファイル

## 今後の課題
- 他の方の回答と比較 
- パフォーマンス改善  
- DDDを初めてやったため、責務分けがきちんとできているかを確認  

[比較対象](https://zenn.dev/foxtail88/scraps/17e94c540e0771)  
[確認はこの本を読むことで行う](https://www.amazon.co.jp/dp/B082WXZVPC/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1)  

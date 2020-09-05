# genji getting started

組み込み可能なドキュメントストアであるgenjiを動かしてみるテスト。

[genjidb/genji: Document-oriented, embedded SQL database, works with Bolt, Badger and memory](https://github.com/genjidb/genji)

## 感想

めっちゃ便利。  
クエリ書いてプレースホルダに構造体入れるとかちょっと「うん？」ってところはあるけど、  
なんでもかんでも `map[string]interface{}` で扱わないといけないのに比べると最高。

パフォーマンスはわからんけど、とりあえず基本的なのは全部扱えそう。  
日付型も扱えて、タイムゾーンまでもキープできるなんて！
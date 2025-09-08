# これは何？

[slim-sprig](https://github.com/go-task/slim-sprig)は、[Task](https://taskfile.dev/#/)内部で利用されているテンプレート向け共通ユーティリティです。

元々、[sprig](https://github.com/Masterminds/sprig) というライブラリがあり、それを [Task](https://taskfile.dev/#/) チームがフォークしてスリム版としてリリースしている模様。

使い方は、本家の [sprig](https://github.com/Masterminds/sprig) と同じで、[Template.Funcs](https://pkg.go.dev/text/template@go1.25.1#Template.Funcs) にそれぞれのライブラリの関数マップを渡すだけです。

```go
import "github.com/Masterminds/sprig/v3"

template.New("xxx").Funcs(sprig.FuncMap())
```

```go
import "github.com/go-task/slim-sprig/v3"

template.New("xxx").Funcs(sprig.FuncMap())
```

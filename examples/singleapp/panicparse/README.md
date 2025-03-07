# これは何？

[maruel/panicparse](https://github.com/maruel/panicparse) というパニック時の出力をヒューマンフレンドリーにしてくれる 便利なライブラリのサンプルが置かれています。

## 実行

```sh
$ task
task: [clean] rm -f ./app
task: [build] go build -o app main.go
task: [run] ./app
0
1
2
3
panic: send on closed channel

goroutine 6 [running]:
main.run.func1(0xc000096070)
        /workspace/try-golang-extlib/examples/singleapp/panicparse/main.go:29 +0x5f
created by main.run in goroutine 1
        /workspace/try-golang-extlib/examples/singleapp/panicparse/main.go:26 +0xd1

goroutine 1 [runnable]:
sync.runtime_SemacquireWaitGroup(0xc00000e0f0?)
        /home/gitpod/go/src/runtime/sema.go:110 +0x25
sync.(*WaitGroup).Wait(0x0?)
        /home/gitpod/go/src/sync/waitgroup.go:118 +0x48
main.run()
        /workspace/try-golang-extlib/examples/singleapp/panicparse/main.go:45 +0x1f1
main.main()
        /workspace/try-golang-extlib/examples/singleapp/panicparse/main.go:15 +0x13

task: [run] ./app |& pp
0
1
2
3
panic: send on closed channel

1: running [Created by main.run in goroutine 1 @ main.go:26]
    main main.go:29       run.func1(0xc000096070)
1: runnable
    sync sema.go:110      runtime_SemacquireWaitGroup(*uint32(#1))
    sync waitgroup.go:118 (*WaitGroup).Wait(*WaitGroup(0x0))
    main main.go:45       run()
    main main.go:15       main()
```

## 参考情報

- [maruel/panicparse](https://github.com/maruel/panicparse)
- [Tips to debug hanging Go programs](https://michael.stapelberg.ch/posts/2025-02-27-debug-hanging-go-programs/)

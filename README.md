## solana-go crash with go 1.18rc1

```sh
go1.18rc1 run main.go
```

Will crash with error:

```
unexpected fault address 0xb01dfacedebac1e
fatal error: fault
[signal SIGSEGV: segmentation violation code=0x1 addr=0xb01dfacedebac1e pc=0x105fe9f]

goroutine 1 [running]:
runtime.throw({0x13632ca?, 0xc0002df438?})
```

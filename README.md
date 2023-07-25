The repository demonstrates a crash of [`gocritic`](https://github.com/go-critic/go-critic) on a code using generics. The problem was reported to the project at https://github.com/go-critic/go-critic/issues/1354.

To reproduce the issue run:

```
make check`
```

or invoke `gocritic` directly via:

```
gocritic check -enable rangeValCopy .
```

The error I get is as follows:


```
gocritic check -enable rangeValCopy .
panic: /opt/homebrew/Cellar/go/1.20.5/libexec/src/go/types/sizes.go:82: assertion failed [recovered]
	panic: /opt/homebrew/Cellar/go/1.20.5/libexec/src/go/types/sizes.go:82: assertion failed

goroutine 955 [running]:
main.(*program).checkFile.func1.1()
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/cmd/gocritic/check.go:168 +0x114
panic({0x104776c80, 0x1400078b420})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/runtime/panic.go:884 +0x204
go/types.assert(0xc8?)
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/types/errors.go:28 +0x60
go/types.(*StdSizes).Alignof(0x1400002a4b0, {0x1047f2dc8?, 0x140009edb60})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/types/sizes.go:82 +0xfc
go/types.(*StdSizes).Offsetsof(0x1047f2d78?, {0x1400078ae10, 0x2, 0x1400161ab90?})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/types/sizes.go:123 +0xc8
go/types.(*StdSizes).Sizeof(0x1400002a4b0, {0x1047f2d78?, 0x140009edd70})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/types/sizes.go:176 +0x1c8
github.com/go-critic/go-critic/linter.(*CheckerContext).SizeOf(0x140003f87d0?, {0x1047f2d78?, 0x140009edd70?})
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/linter/linter.go:353 +0xac
github.com/go-critic/go-critic/checkers.(*rangeValCopyChecker).VisitStmt(0x14000777340, {0x1047f38c8?, 0x1400176a1e0?})
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/checkers/rangeValCopy_checker.go:68 +0x98
github.com/go-critic/go-critic/checkers/internal/astwalk.(*stmtWalker).WalkFile.func1({0x1047f29b8?, 0x1400176a1e0?})
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/checkers/internal/astwalk/stmt_walker.go:23 +0x60
go/ast.inspector.Visit(0x1400078b3f0, {0x1047f29b8?, 0x1400176a1e0?})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/ast/walk.go:386 +0x38
go/ast.Walk({0x1047f1bd8?, 0x1400078b3f0?}, {0x1047f29b8?, 0x1400176a1e0?})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/ast/walk.go:51 +0x4c
go/ast.walkStmtList({0x1047f1bd8, 0x1400078b3f0}, {0x14000776b20?, 0x2, 0x1400045ee08?})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/ast/walk.go:32 +0x64
go/ast.Walk({0x1047f1bd8?, 0x1400078b3f0?}, {0x1047f2440?, 0x140009ed620?})
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/ast/walk.go:234 +0xce0
go/ast.Inspect(...)
	/opt/homebrew/Cellar/go/1.20.5/libexec/src/go/ast/walk.go:397
github.com/go-critic/go-critic/checkers/internal/astwalk.(*stmtWalker).WalkFile(0x1400078b3d0, 0x14000172ab0)
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/checkers/internal/astwalk/stmt_walker.go:21 +0x138
github.com/go-critic/go-critic/linter.(*Checker).Check(...)
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/linter/linter.go:159
main.(*program).checkFile.func1()
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/cmd/gocritic/check.go:172 +0x94
created by main.(*program).checkFile
	/Users/dmitri/go/pkg/mod/github.com/go-critic/go-critic@v0.8.1/cmd/gocritic/check.go:152 +0x94
make: *** [Makefile:2: check] Error 2
```
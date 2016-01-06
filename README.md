# convert
convert is a text file encoding converter

such as you can convert text file's encoding from gbk to utf-8

support encoding : utf-8, gbk, big5, latin-1, UTF-16, ASCII,, gb18030, SJIS, EUC-JP

**convert use golang to develop, you can use "go get" command to compile executable program that is suitable for your computer.**

```bash
go get github.com/Comdex/convert
```

### usage

if you want to convert a.txt encoding from gbk to utf-8(--src and --scode are necessary), like this:
```bash
convert f --src a.txt --scode gbk
```
if you want to convert all file in a directory like ok ,do this:
```bash
convert d --src ok --scode gbk
```

more help maybe you can use
```bash
convert -h
```

```bash
convert f -h
```

```bash
convert d -h
```

### executable download link

[win32 executable download](http://pan.baidu.com/s/1c1m4qre)

[linux amd64 executable download](http://pan.baidu.com/s/1dDWTiu9)

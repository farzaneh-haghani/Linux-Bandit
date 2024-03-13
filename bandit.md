#### Level 0

```sh
- ssh bandit0@bandit.labs.overthewire.org -p 2220
- cat readme
```

#### Level 1

```sh
- cat ./-
```

#### Level 2

```sh
- cat spaces\ in\ this\ filename
```

#### Level 3

```sh
- cd inhere/
- ls -a
- cat .hidden
```

#### Level 4

```sh
- find inhere/ -readable -exec file {} \;
- cat inhere/-file07
```

#### Level 5

```sh
- find inhere/ -readable ! -executable -size 1033c
- cat inhere/maybehere07/.file2
```

#### Level 6

```sh
- find / -user bandit7 -group bandit6 -size 33c
- cat /var/lib/dpkg/info/bandit7.password
```

#### Level 7

```sh
- grep "millionth" data.txt
```

#### Level 8

```sh
- sort data.txt | uniq -u
```

#### Level 9

```sh
- strings data.txt | grep "="
```

#### Level 10

```sh
- base64 -d data.txt
```

#### Level 11

```sh
- cat data.txt | tr '[A-Za-z]' '[N-ZA-Mn-za-m]'
```

#### Level 12

```sh
mkdir /tmp/farzaneh
cp data.txt /tmp/farzaneh
cd /tmp/farzaneh
xxd -r data.txt > pass
mv pass pass.gz
gzip -d pass.gz
mv pass pass.bz2
bzip2 -d pass.bz2
mv pass pass.gz
gzip -d pass.gz
mv pass pass.tar
tar -xf pass.tar
mv data5.bin data5.bin.tar
tar -xf data5.bin.tar
mv data6.bin data6.bin.bz2
bzip2 -d data6.bin.bz2
mv data6.bin data6.bin.tar
tar -xf data6.bin.tar
mv data8.bin data8.bin.gz
gzip -d data8.bin.gz
cat data8.bin
```

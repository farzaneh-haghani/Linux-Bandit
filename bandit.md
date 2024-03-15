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

#### Level 13

```sh
ssh -i sshkey.private bandit14@bandit.labs.overthewire.org -p 2220

```

#### Level 14

```sh
cat /etc/bandit_pass/bandit14
nc localhost 30000
fGrHPx402xGC7U7rXKDaxiWFTOiF0ENq
```

#### Level 15

```sh
openssl s_client -connect localhost:30001
jN2kgmIXJ6fShzhT2avhotn4Zcka6tnt
```

#### Level 16

```sh
nmap -p 31000-32000 localhost
openssl s_client -connect localhost:31790
JQttfApK4SeyHwDlI9SXGR50qclOAil1
```

#### Level 17

```sh
nano private.pem   //past the private key
chmod 400 private.pem
ssh -i private.pem bandit17@bandit.labs.overthewire.org -p 2220
diff passwords.old passwords.new

```

#### Level 0

- ssh bandit0@bandit.labs.overthewire.org -p 2220
- cat readme

#### Level 1

- cat ./-

#### Level 2

- cat spaces\ in\ this\ filename

#### Level 3

- cd inhere/
- ls -a
- cat .hidden

#### Level 4

- find inhere/ -readable -exec file {} \;
- cat inhere/-file07

#### Level 5

- find inhere/ -readable ! -executable -size 1033c
- cat inhere/maybehere07/.file2

#### Level 6

- find / -user bandit7 -group bandit6 -size 33c
- cat /var/lib/dpkg/info/bandit7.password

#### Level 7

- grep "millionth" data.txt

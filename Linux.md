#### To learn a command from the manual page

- "q" to exit
- "/keyword" to search

```sh
man <command>
```

#### To learn a command which is shell built-in and doesn't have manual page

```sh
help <command>
```

#### To connect to a remote server by Secure Shell Protocol (SSH) which is providing a secure channel over an unsecured network

```sh
ssh [username]@[host] -p [port]
```

#### To show the list of all files in a directory including hidden files

```sh
ls -a
```

#### To show the list with details "long format"

```sh
ls -l
```

#### To show the names of all files in double quotation marks "" to see if there is any extra spaces in the names

```sh
ls -Q
```

#### To read a file (Adding the path is required for files that start with special character)

```sh
cat <file_name>
cat /path/file_name
```

#### To find all files with .txt

- `find /`: searches for files and directories starting from the root or highest-level directory
- `find .`: searches for files and directories starting from the current directory
- `find path/`: searches for files and directories starting from specific path

```sh
find path/ -name "*.txt"
```

#### To find a file with specific size

- c for byte
- k for kilobyte = 1024 bytes
- m for megabyte = 1024 kilobytes
- g for gigabyte = 1024 megabytes

```sh
find path/ -size 123c
```

#### To see all files that owned by a user or owned by group or

```sh
find path/ -user <user_name>
find path/ -group <group_name>
```

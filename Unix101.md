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

#### To find the name of the current directory (Print Working Directory)

```sh
pwd
```

#### To see the list of files and directories in this directory (List)

```sh
ls
```

#### To see the list with permissions, dates, ownership and other information (long listing)

```sh
ls -l
```

#### `/` (root directory) To see all directories and files are contained under it.

```sh
ls -l /
```

#### To change directories

```sh
cd
```

#### To find a file on your filesystem

- `find /`: searches for files and directories starting from the root or highest-level directory
- `find .`: searches for files and directories starting from the current directory
- `find path/`: searches for files and directories starting from specific path

```sh
find /home/opsschool -type f -name file3.txt
```

#### To show the list of all files in a directory including hidden files

```sh
ls -a
```

#### To show the names of all files in double quotation marks "" to see if there is any extra spaces in the names

```sh
ls -Q
```

#### To connect to a remote server by Secure Shell Protocol (SSH) which is providing a secure channel over an unsecured network

```sh
ssh [username]@[host] -p [port]
```

#### To read a file (Adding the path is required for files that start with special character)

```sh
cat <file_name>
cat /path/file_name
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

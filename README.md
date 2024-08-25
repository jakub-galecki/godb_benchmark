This repository contains banchmark for my implementation of Log Structured Merge Tree (https://github.com/jakub-galecki/godb/).

### Installing 

To install all required dependencies run command `go mod tidy`.

### Tested on Macbook M1
```
Software:
    System Software Overview:
      System Version: macOS 13.1 (22C65)
      Kernel Version: Darwin 22.2.0
Hardware:
    Hardware Overview:
      Model Name: MacBook Air
      Model Identifier: MacBookAir10,1
      Chip: Apple M1
      Total Number of Cores: 8 (4 performance and 4 efficiency)
      Memory: 16 GB
```

### Sqlite

Comparing inserting 100 000 keys to both my implementation and sqlite3 gives us following results:
```
1. Inserting 100000 keys to my implementation took: 1.450964333s
> Stats for generated files in my implementation 
file: /tmp/thesis_test/MANIFEST,  size: 211 B
file: /tmp/thesis_test/sst,  size: 128 B
file: /tmp/thesis_test/sst/0.db,  size: 2.4 MB
file: /tmp/thesis_test/sst/1.db,  size: 2.4 MB
file: /tmp/thesis_test/wal,  size: 160 B
file: /tmp/thesis_test/wal/000000000.log,  size: 1.4 MB
file: /tmp/thesis_test/wal/000000001.log,  size: 1.4 MB
file: /tmp/thesis_test/wal/000000002.log,  size: 631 kB
Creating sqlite file in: /var/folders/y4/0k0_krgj5sd6wf8wrkk1qqtc0000gn/T/test1005467574
2. Inserting 100000 keys to my sqlite took: 20.484513416s
> Stats for generated files in sqlite 
file: /var/folders/y4/0k0_krgj5sd6wf8wrkk1qqtc0000gn/T/test1005467574/test.db,  size: 2.7 MB
```

# Toolkit Manual

This package contains a set of tools for TRON, the followings are the documentation for each tool.

## Build The Toolkit

First, you can build the toolkit by executing the following command:
```shell script
# clone the tron-docker
git clone https://github.com/tronprotocol/tron-docker.git
# enter gradlew directory
cd tron-docker/tools/gradlew
# build the toolkit
./gradlew :toolkit:build
# execute the command
java -jar ../toolkit/build/libs/Toolkit.jar db -h
```
The most commonly used db commands are:
- `help`: Displays help information about the specified command
- `mv, move`: Move db to pre-set new path . For example HDD,reduce storage
expenses.
- `archive`: A helper to rewrite leveldb manifest.
- `convert`: Covert leveldb to rocksdb.
- `lite`: Split lite data for java-tron.
- `cp, copy`: Quick copy leveldb or rocksdb data.
- `root`: compute merkle root for tiny db. NOTE: large db may GC overhead
limit exceeded.
- `fork`: Modify the database of java-tron for shadow fork testing.

## DB Archive

DB archive provides the ability to reformat the manifest according to the current `database`, parameters are compatible with the previous `ArchiveManifest`.

### Available parameters:

- `-b | --batch-size`: Specify the batch manifest size, default: 80000.
- `-d | --database-directory`: Specify the database directory to be processed, default: output-directory/database.
- `-m | --manifest-size`: Specify the minimum required manifest file size, unit: M, default: 0.
- `-h | --help`: Provide the help info.

### Examples:

```shell script
# full command
  java -jar Toolkit.jar db archive [-h] [-b=<maxBatchSize>] [-d=<databaseDirectory>] [-m=<maxManifestSize>]
# examples
   java -jar Toolkit.jar db archive #1. use default settings
   java -jar Toolkit.jar db archive -d /tmp/db/database #2. specify the database directory as /tmp/db/database
   java -jar Toolkit.jar db archive -b 64000 #3. specify the batch size to 64000 when optimizing manifest
   java -jar Toolkit.jar db archive -m 128 #4. specify optimization only when Manifest exceeds 128M
```


## DB Convert

DB convert provides a helper which can convert LevelDB data to RocksDB data, parameters are compatible with previous `DBConvert`.

### Available parameters:

- `<src>`: Input path for leveldb, default: output-directory/database.
- `<dest>`: Output path for rocksdb, default: output-directory-dst/database.
- `--safe`: In safe mode, read data from leveldb then put into rocksdb, it's a very time-consuming procedure. If not, just change engine.properties from leveldb to rocksdb, rocksdb
  is compatible with leveldb for the current version. This may not be the case in the future, default: false.
- `-h | --help`: Provide the help info.

### Examples:

```shell script
# full command
  java -jar Toolkit.jar db convert [-h] [--safe] <src> <dest>
# examples
  java -jar Toolkit.jar db convert  output-directory/database /tmp/database
```

## DB Copy

DB copy provides a helper which can copy LevelDB or RocksDB data quickly on the same file systems by creating hard links.

### Available parameters:

- `<src>`: Source path for database. Default: output-directory/database
- `<dest>`: Output path for database. Default: output-directory-cp/database
- `-h | --help`: provide the help info

### Examples:

```shell script
# full command
  java -jar Toolkit.jar db cp [-h] <src> <dest>
# examples
  java -jar Toolkit.jar db cp  output-directory/database /tmp/databse
```

## DB Lite

DB lite provides lite database, parameters are compatible with previous `LiteFullNodeTool`.

### Available parameters:

- `-o | --operate`: [split,merge], default: split.
- `-t | --type`: Only used with operate=split: [snapshot,history], default: snapshot.
- `-fn | --fn-data-path`: The database path to be split or merged.
- `-ds | --dataset-path`: When operation is `split`,`dataset-path` is the path that store the `snapshot` or `history`, when
  operation is `split`, `dataset-path` is the `history` data path.
- `-h | --help`: Provide the help info.

### Examples:

```shell script
# full command
  java -jar Toolkit.jar db lite [-h] -ds=<datasetPath> -fn=<fnDataPath> [-o=<operate>] [-t=<type>]
# examples
  #split and get a snapshot dataset
  java -jar Toolkit.jar db lite -o split -t snapshot --fn-data-path output-directory/database --dataset-path /tmp
  #split and get a history dataset
  java -jar Toolkit.jar db lite -o split -t history --fn-data-path output-directory/database --dataset-path /tmp
  #merge history dataset and snapshot dataset
  java -jar Toolkit.jar db lite -o merge --fn-data-path /tmp/snapshot --dataset-path /tmp/history
```

## DB Move

DB move provides a helper to move some dbs to a pre-set new path. For example move `block`, `transactionRetStore` or `transactionHistoryStore` to HDD for reducing storage expenses.

### Available parameters:

- `-c | --config`: config file. Default: config.conf.
- `-d | --database-directory`: database directory path. Default: output-directory.
- `-h | --help`: provide the help info

### Examples:

Take the example of moving `block` and `trans`.


Set path for `block` and `trans`.

```conf
storage {
 ......
  properties = [
    {
     name = "block",
     path = "/data1/tron",
    },
    {
     name = "trans",
     path = "/data1/tron",
   }
  ]
 ......
}
```
Execute move command.
```shell script
# full command
  java -jar Toolkit.jar db mv [-h] [-c=<config>] [-d=<database>]
# examples
  java -jar Toolkit.jar db mv -c main_net_config.conf -d /data/tron/output-directory
```

## DB Root

DB root provides a helper which can compute merkle root for tiny db.

NOTE: large db may GC overhead limit exceeded.

### Available parameters:

- `<src>`: Source path for database. Default: output-directory/database
- `--db`: db name.
- `-h | --help`: provide the help info


## DB Fork
DB fork tool can help launch a private java-tron FullNode or network based on the state of public chain database to support shadow fork testing.

### Available parameters:
- `-c, --config=<config>`: config the new witnesses, balances, etc for shadow
fork. Default: fork.conf
- `-d, --database-directory=<database>`: java-tron database directory path. Default:
output-directory
- `-r, --retain-witnesses`: retain the previous witnesses and active witnesses.
- `-h, --help`: provide the help info

Please refer [DBFork](DBFork.md) guidance for more details.

## DB Query
The [ListWitnesses](https://developers.tron.network/reference/listwitnesses) and [getReward](https://developers.tron.network/reference/wallet-getreward)
APIs provide vote and reward data only until the end of the last maintenance period.
To obtain the latest information, please query the database directly using the DB query tool.

### Available parameters:
- `-c, --config=<config>`: config the vote and reward options. Default: query.conf
- `-d, --database-directory=<database>`: java-tron database directory path. Default: output-directory
- `-h, --help`: provide the help info

### Examples:
The example of `query.conf`:
```conf
vote = {
  allWitnesses = true
  witnessList = [
    "TPDa1KRmFpLA42WffiXBwEykef9exHNPKV",
    "TXnqPcy9zWHYyCiQ7Kmvt3aCx5jm9Qnq3e",
  ]
}

reward = [
  "TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZYH",
  "TJvaAeFb8Lykt9RQcVyyTFN2iDvGMuyD4M"
]
```
- `vote.allWitnesses`: configure whether to query the vote information of all witnesses.
- `vote.witnessList`: configure to query the vote information from specified witness list.
                      The option is valid only when `vote.allWitnesses = false`.
- `reward`: configure the address list you need to query the reward.

Execute query command.
```shell script
# full command
  java -jar Toolkit.jar db query [-h] [-c=<config>] [-d=<database>]
# examples
  java -jar Toolkit.jar db query -c query.conf -d output-directory
```

## Apache Log Analyzer (ALA)

The Apache Log Analyzer (ALA) is a straightforward tool for analyzing Apache logs. It can filter and enhance log data,
then generate reports based on the aggregated information. The flexible design of ALA enables the quick creation and
addition of new reports. This tool processes and analyzes the data from the file at the same time as it is reading from
the file, without loading the entire data into memory. This can be useful for handling large data sets that cannot be
stored in memory.

### Installation

Clone the repository and run the following command in the terminal to install the required dependencies:

```
go get -v -t ./...
```

### Usage

You can run the `ala` tool using the following command:

```
go run cmd/ala/main.go -l <log_file_path> -d <geo_db_path>
```

Or you can build an executable file:

```
go build -o ala ./cmd/ala/main.go
./ala -l <log_file_path> -d <geo_db_path>
```

* `log_file_path` is the path to the Apache log file you want to analyze.
* `geo_db_path` is the path to the GeoLite2 City database.
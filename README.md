# goenum
Enum generator for golang.  

# Install
## golang
```
$ go get -u github.com/ken-aio/goenum
```

## Linux or Mac
You can confirm available version in [releases](https://github.com/ken-aio/execql/releases).  
```
$ wget https://github.com/ken-aio/execql/releases/download/v{version}/execql_v{version}_{OS}_amd64.tar.gz
$ tar xvf execql_v{version}_{OS}_amd64.tar.gz
$ mv execql /usr/local/bin/
```

# How to use
## 1. Initialize goenum
Initialize goenum with below command.  
```
$ goenum init
Create new config file   : .goenum.yml
create new template file : goenum/template/enum.tmpl
```

### .goenum.yml
This file is goenum master setting.  
Default is below.  

| parameter    | setting         | contents                                                            |
|--------------|-----------------|---------------------------------------------------------------------|
| template.dir | goenum/template | goenum template file. enum golang file is generated from this file. |
| yaml.dir     | goenum/yaml     | directory for enum parameter setting yaml files.                    |
| gofile.dir   | app/enum        | generated golang file(.go) output dir.                              |

## 2. Generate new enum setting yaml
Generate new enum setting yaml file using below command.  
Modify this file if needed.  
```
$ goenum new weekday sunday:sun monday:mon tuesday:tue wednesday:wed thursday:thu friday:fri saturday:sat
create new enum file: goenum/yaml/weekday.yml
$ cat goenum/yaml/weekday.yml
# This file is goenum enum setting file auto generated. if you see more detail, https://github.com
/ken-aio/goenume
name: Weekday
description: |-
  This is Weekday enums.
values:
  Sunday: sun
  Monday: mon
  Tuesday: tue
  Wednesday: wed
  Thursday: thu
  Friday: fri
  Saturday: sat
```

| parameter    | contents                          |
|--------------|-----------------------------------|
| name         | enum name. camel case recommended |
| description  | enum command                      |
| values.key   | enum variable                     |
| values.value | enum value                        |

## Generate enum files
Generate enums from all yaml settings with blow command.  
```
$ goenum g
create new enum file: app/enum/weekday.go
```

## Using generated enum
Using example code.
```
// using enum
fmt.Printf("sunday = %s\n", enum.Sunday)

// convert enum from string
monday := enum.NewWeekday("mon")
fmt.Printf("monday = %s\n", monday)

// get all enums
for i, w := range enum.WeekdayList() {
	fmt.Printf("[%d] : %s\n", i, w)
}

// get all enum strings
for i, w := range enum.WeekdayNames() {
	fmt.Printf("[%d] : %s\n", i, w)
}
```
Executing.
```
sunday = sun
monday = mon
[0] : unknown
[1] : sun
[2] : mon
[3] : tue
[4] : wed
[5] : thu
[6] : fri
[7] : sat
[0] : unknown
[1] : sun
[2] : mon
[3] : tue
[4] : wed
[5] : thu
[6] : fri
[7] : sat
```

## Usage
```
$ goenum
goenum is golang enum generator from yaml settings.

Usage:
  goenum [command]

Available Commands:
  g           Initialize goenum. create template yml files in goenum dir
  generate    Initialize goenum. create template yml files in goenum dir
  help        Help about any command
  init        Initialize goenum. generate enum template file.
  n           Generate new enum file
  new         Generate new enum file
  version     All software has versions. This is Goenum's.

Flags:
      --config string   config file (default is .goenum.yml) (default ".goenum.yml")
  -h, --help            help for goenum

Use "goenum [command] --help" for more information about a command.
```

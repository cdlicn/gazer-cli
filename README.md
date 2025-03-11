<img src="README.assets/gazer-17417039661842.png" alt="gazer" style="zoom:38%;float:left;" />

# gazer-cli

Command line interface of "gazer" project

| command | Remarks                    |
| ------- | -------------------------- |
| add     | Add a topic and path       |
| upd     | Modify the path of a topic |
| del     | Delete a topic             |
| list    | View all topics and paths  |

### How to use?

- add:

  ​	gazer add [topic] [path]

  e.g.

  ​	gazer add 'mysql' '/var/log/mysqld.log'

- upd:

  ​	gazer upd [topic] [path]

  e.g.

  ​	gazer upd 'mysql' '/var/log/new_mysqld.log'

- del:

  ​	gazer del [topic]

  e.g.

  ​	gazer del 'mysql'

- list:

  ​	gazer list

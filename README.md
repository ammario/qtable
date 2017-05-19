# qtable

A simple Go package which converts the results of an SQL query to a textual table via [table writer](https://github.com/olekukonko/tablewriter).

![GoDoc](https://godoc.org/github.com/ammario/qtable)

A query such as `SELECT name FROM users` could be presented as:

```
+------+
| NAME |
+------+
| bob  |
| cob  |
| dob  |
| hey  |
| job  |
| slob |
+------+
```
## Patterns

[Link](patterns.md)

## Cut

**Examples:**

```bash
   echo -e "column1\column2\column3" | ./cut -f 1,3
```

```bash
   echo -e "column1,column2,column3" | ./cut -f 1,2 -d ','
```

```bash
    echo -e "column1\tcolumn2\ncolumn3" | ./cut -f 1 -s
```
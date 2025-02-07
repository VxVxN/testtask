## Patterns

[Link](patterns.md)

## Grep

**Examples:**

```bash
  echo -e "line1\nline2\nline3" | ./grep -A 1 line1
```

```bash
  echo -e "line1\nline2\nline3" | ./grep -B 1 line3
```

```bash
  echo -e "line1\nline2\nline3\nline4\nline5" | ./grep -C 1 line3
```

```bash
  echo -e "line1\nline2\nline3" | ./grep line1[12]
```

```bash
  echo -e "line1\nLine2\nLINE3" | ./grep -i line
```

```bash
  echo -e "line1\nLine2\nLINE3" | ./grep -v Line
```

```bash
  echo -e "line1\nhello\nline3" |./grep -c line
```

```bash
  echo -e "line1?\nline2\nline3" | ./grep -F line1?
```

```bash
  echo -e "line1\nline2\nline3" | ./grep -n line[23]
```

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
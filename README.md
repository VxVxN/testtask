## Patterns

[Link](patterns.md)

## Grep

**Flags:**

- -A - Print N lines of trailing context after matching lines
- -B - Print N lines of leading context before matching lines
- -C - Print N lines of output context
- -c - Suppress normal output; instead print a count of matching lines for each input file
- -i - Ignore case distinctions in patterns and input data
- -v - Select non-matching lines
- -F - Interpret PATTERN as a fixed string, not a regular expression
- -n - Print line numbers with output lines

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

**Flags:**
- -f - Select the fields
- -d - Use a different separator
- -s - Delimited lines only

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

## Calendar App

**Examples:**

```bash
curl --request POST -d "user_id=3&date=2025-02-07" http://localhost:8081/event
```

```bash
curl --request PUT -d "user_id=3&date=2025-02-07" http://localhost:8081/event
```

```bash
curl --request GET http://localhost:8081/events?period=month
```

```bash
curl --request DELETE -d "user_id=3&date=2025-02-07" http://localhost:8081/event
```
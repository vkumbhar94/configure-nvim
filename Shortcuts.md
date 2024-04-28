# Shortcuts

- Delete all lines till next non empty lines
  `d/.` + enter - d means delete / means until . means any character
- Delete until next character
    `d/char` + enter - d means delete / means until char
- Delete inner paragraph
    `dip` - d means delete i means inner p means paragraph

- Adding some var in front of lines
 ```go
 a = 10
 b = 20
 c = 30
 ```
 How: Press Ctrl + v and select lines (like vertical multi cursor) and then press Shift + i and type the text and then press Esc

 - Similarly for appending semi colon at end of lines
 How: Press Ctrl + v and select lines (like vertical multi cursor) and then press Shift + a and type the text and then press Esc

 - Markdown preview: `:MarkdownPreview` - This will open the markdown preview in browser


 - if you want to copy same line with increamented number like following:
 ```go
a = append(a, 0)
a = append(a, 0)
a = append(a, 0)
a = append(a, 0)
a = append(a, 0)
 ```

To 
 ```go
a = append(a, 1)
a = append(a, 2)
a = append(a, 3)
a = append(a, 4)
a = append(a, 5)
 ```

 How: enter visual mode, select all lines and press g and press control + a

 The above will do for the first digit in the line but if you have multiple digits then you can use ctrl + v and select single column then do the same g, control + a
 h


 - To sort lines
 select lines and press `:sort`
 to filter on second field `:sort -k2`
 to reverse sort `:sort!`

# ANSI escape sequences

[ANSI escape sequences](https://en.wikipedia.org/wiki/ANSI_escape_code) can be printed to a shell to as instructions.
The below is a list of codes I have used often in my CLI programs and I find myself looking up over and over again.

A great article about it can be found [here](https://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html).

## Content

- [Colors](#colors)
- [Styles](#styles)
- [Cursor](#cursor)
- [Screen](#screen)

## Prefix

The standard prefix is `ESC` which can be expressed as:

- Ctrl-Key: `^[`
- Octal: `\033`
- Unicode: `\u001b`
- Hexadecimal: `\x1B`
- Decimal: `27`

## Colors

### Font color

#### 8 Colors

| code       | description        |
| ---------- | ------------------ |
| `\x1B[30m` | Black font color   |
| `\x1B[31m` | Red font color     |
| `\x1B[32m` | Green font color   |
| `\x1B[33m` | Yellow font color  |
| `\x1B[34m` | Blue font color    |
| `\x1B[35m` | Magenta font color |
| `\x1B[36m` | Cyan font color    |
| `\x1B[37m` | White font color   |

#### 16 Colors

| code         | description               |
| ------------ | ------------------------- |
| `\x1B[30;1m` | Bright Black font color   |
| `\x1B[31;1m` | Bright Red font color     |
| `\x1B[32;1m` | Bright Green font color   |
| `\x1B[33;1m` | Bright Yellow font color  |
| `\x1B[34;1m` | Bright Blue font color    |
| `\x1B[35;1m` | Bright Magenta font color |
| `\x1B[36;1m` | Bright Cyan font color    |
| `\x1B[37;1m` | Bright White font color   |

#### 256 colors

| code                    | description                                                      |
| ----------------------- | ---------------------------------------------------------------- |
| `\x1B[38;5;` + n + `m ` | Standard font color where `n` can be a number between 0-7        |
| `\x1B[38;5;` + n + `m ` | High intensity font color where `n` can be a number between 8-15 |
| `\x1B[38;5;` + n + `m ` | Rainbow font color where `n` can be a number between 16-231      |
| `\x1B[38;5;` + n + `m ` | Gray font color where `n` can be a number between 232-255        |

#### 24bit Truecolor

| code                          | description      |
| ----------------------------- | ---------------- |
| `\x1B[38;2;{r};{g};{b}m`      | Foreground color |

### Background colors

#### 8 Colors

| code       | description              |
| ---------- | ------------------------ |
| `\x1B[40m` | Black background color   |
| `\x1B[41m` | Red background color     |
| `\x1B[42m` | Green background color   |
| `\x1B[43m` | Yellow background color  |
| `\x1B[44m` | Blue background color    |
| `\x1B[45m` | Magenta background color |
| `\x1B[46m` | Cyan background color    |
| `\x1B[47m` | White background color   |

#### 16 Colors

| code         | description                     |
| ------------ | ------------------------------- |
| `\x1B[40;1m` | Bright Black background color   |
| `\x1B[41;1m` | Bright Red background color     |
| `\x1B[42;1m` | Bright Green background color   |
| `\x1B[43;1m` | Bright Yellow background color  |
| `\x1B[44;1m` | Bright Blue background color    |
| `\x1B[45;1m` | Bright Magenta background color |
| `\x1B[46;1m` | Bright Cyan background color    |
| `\x1B[47;1m` | Bright White background color   |

#### 256 colors

| code                    | description                                                            |
| ----------------------- | ---------------------------------------------------------------------- |
| `\x1B[48;5;` + n + `m ` | Standard background color where `n` can be a number between 0-7        |
| `\x1B[48;5;` + n + `m ` | High intensity background color where `n` can be a number between 8-15 |
| `\x1B[48;5;` + n + `m ` | Rainbow background color where `n` can be a number between 16-231      |
| `\x1B[48;5;` + n + `m ` | Gray background color where `n` can be a number between 232-255        |

#### 24bit Truecolor

| code                          | description      |
| ----------------------------- | ---------------- |
| `\x1B[48;2;{r};{g};{b}m`      | Background color |

## Styles

| code      | description      |
| --------- | ---------------- |
| `\x1B[0m` | Reset all styles |
| `\x1B[39m`| Reset font color |
| `\x1B[49m`| Reset bg color   |
| `\x1B[1m` | Bold             |
| `\x1B[4m` | Underline        |
| `\x1B[7m` | Reversed         |

## Cursor

| code                        | description                                      |
| --------------------------- | ------------------------------------------------ |
| `\x1B[?25l`                 | Hide cursor                                      |
| `\x1B[?25h`                 | Show cursor                                      |
| `\x1B[` + n + `A`           | Move Up by `n` rows                              |
| `\x1B[` + n + `B`           | Move Down by `n` rows                            |
| `\x1B[` + n + `C`           | Move Right by `n` columns                        |
| `\x1B[` + n + `D`           | Move Left by `n` columns                         |
| `\x1B[` + n + `E`           | Move cursor to beginning of line, `n` lines down |
| `\x1B[` + n + `F`           | Move cursor to beginning of line, `n` lines up   |
| `\x1B[` + n + `G`           | Move cursor to column `n`                        |
| `\x1B[` + n + `;` + m + `H` | Move cursor to row `n` column `m`                |
| `\x1B[{s}`                  | Save the current cursor position                 |
| `\x1B[{u}`                  | Restore the cursor to the last saved position    |

## Screen

| code      | description                               |
| --------- | ----------------------------------------- |
| `\x1B[0J` | clears from cursor until end of screen    |
| `\x1B[1J` | clears from cursor to beginning of screen |
| `\x1B[2J` | clears entire screen                      |
| `\x1B[0K` | clears from cursor to end of line         |
| `\x1B[1K` | clears from cursor to start of line       |
| `\x1B[2K` | clears entire line                        |

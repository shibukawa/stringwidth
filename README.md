# stringwidth

Go library that calculate display width of string.

## Why this is difficult?

Unicode database has display width information.
For example, 'a' has single, but 'あ' has double width.

But there are ambiguous characters like '¼'.

And emoji has modifier. For example, England Flag(🏴󠁧󠁢󠁥󠁮󠁧󠁿)
seems to be single character width, but it consists of 6 runes.

## API

`func width.Calc(string, ...opt) int`

It returns display length of input string.

```go
w := width.Calc("🏴󠁧󠁢󠁥󠁮󠁧󠁿")
// ↑2
```

If you want to specify ambiguous character width, you should add option:

```go
w := width.Calc("¼", width.Opt{
    IsAmbiguousWide: false,
})
// ↑1
```

## License

Apache2

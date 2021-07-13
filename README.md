# Glad - Glossary Add

One thing that annoyed me while writing my thesis was adding all the skeleton stuff to a `.tex`glossary entry. So `glad` was born. A simple command line tool, that generates one glossary entry at a time without the overhead.

```sh
$ glad "Three Letter Acronym" "This is a description, for those who don't know what's happening"

$ tail glossary.tex

\newglossaryentry{tla}{
    name={TLA},
        first={Three Letter Acronym (TLA)},
    description={This is a description, for those who don't know what's happening}
}
```

## Is it any good?

[Yes.](https://news.ycombinator.com/item?id=3067434)

## But there is this other tool, that does this for ages!

Maybe. But I needed a break from writing. And this way I had a reason to look into the Go templating package.

## Usage

Just provide the glossary entry you want and its longer description as positional arguments. Acronyms are auto-generated from the first letter of each word in the first positional argument. Use the "-a" flag to provide a custom acronym.

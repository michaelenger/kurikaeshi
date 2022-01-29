# Kurikaeshi (繰り返し)

Learn you Japanese good through repetition.

Kurikaeshi is a simple CLI word challenge game written in [Go](https://go.dev/).
It's meant for practising hiragana and katakana knowledge by presenting you with
words in the respective syllabary and expecting you to write the rōmaji version.

## Usage

The program is run by specifying a syllabary, either `hiragana` or `katakana` and
and optional flag to specify the amount of words you want to test yourself on.

```shell
kurikaeshi katakana -n 10
```

It will present you with words and you have to write the correct rōmaji version.
If at any time you want to quit you just need to enter in a blank guess and the
program will exit.

![An example of usage](https://github.com/michaelenger/kurikaeshi/blob/master/example.png?raw=true)

## Credits

* Words taken from ["1000 Japanese basic words" on Wikitionary](https://en.wiktionary.org/wiki/Appendix:1000_Japanese_basic_words)

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"n2k"
	"os"
)

type Option struct {
	Object    string `short:"s" long:"string" description:"変換する文字列を入力します"`
	Num2Kanji bool   `short:"k" long:"n2k" description:"数字を漢数字に変換します"`
	Kanji2Num bool   `short:"n" long:"k2n" description:"漢数字を数字に変換します"`
}

var opts Option

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res string
	if opts.Num2Kanji {
		res, err = n2k.Number2kanji(opts.Object)
		if err != nil {
			return
		}
	} else if opts.Kanji2Num {
		res, err = n2k.Kanji2number(opts.Object)
		if err != nil {
			return
		}
	}
	fmt.Println(res)
}
package main


import  (
	"text/scanner"
	"strings"
	"fmt"
)

func  main () {
	fmt.Println("constant EOF : ",scanner.EOF)
	fmt.Println("constant Ident : ",scanner.Ident)
	fmt.Println("constant Int   : ",scanner.Int)
	fmt.Println("constant Float   : ",scanner.Float)
	fmt.Println("constant Char    : ",scanner.Char)
	fmt.Println("constant Comment ",scanner.Comment)
//	fmt.Println("constant skipComment ",scanner.skipComment)
	fmt.Printf("constant ScanIdents : %d  %08[1]b\n",scanner.ScanIdents)
	fmt.Printf("constant ScanInts   : %d  %08[1]b\n",scanner.ScanInts)
	fmt.Printf("constant ScanComments: %d  %08[1]b\n",scanner.ScanComments)
	fmt.Printf("constant SkipComments: %d  %08[1]b\n",scanner.SkipComments)
	fmt.Printf("constant GoWhitespace`: %d  %08[1]b\n",scanner.GoWhitespace)
	fmt.Printf("constant GoTokens : %d  %08[1]b\n",scanner.GoTokens)
	fmt.Printf("rune %d [%[1]c]\n",'\n')
	fmt.Printf("rune %d [%[1]c]\n",'\t')
	fmt.Printf("rune %d [%[1]c]\n",'\r')
	fmt.Printf("rune %d [%[1]c]\n",' ')
	fmt.Printf("constant GoTokens : %d  %08[1]b\n",scanner.GoTokens)
	const src = `((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964) (Actor (("Maj. T.J. \"King\" Kong" "Slim Pickens") ("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Sterling Hayden"))) (Oscars ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)" "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))`
	var s scanner.Scanner
	fmt.Printf("GoTokens   %d %08[1]b\n", scanner.GoTokens)
	s.Init(strings.NewReader(src))
	fmt.Printf("constant Mode     : %d  %08[1]b\n",s.Mode)
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s  %s\n", scanner.TokenString(tok), s.TokenText())
	}
        fmt.Println()
      // GoTokens       = ScanIdents |         ScanFloats |         ScanChars | ScanStrings | ScanRawStrings |         ScanComments |         SkipComments
	s.Init(strings.NewReader(src))
        s.Mode  = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars |       scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
	fmt.Printf("constant GoTokens : %d  %08[1]b\n",scanner.GoTokens)
	fmt.Printf("constant Mode     : %d  %08[1]b\n",s.Mode)
	fmt.Println("Tok, Tok,  TokString TokenText")
	fmt.Println("------------------------------")
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%d %c %s  %s\n", tok, tok, scanner.TokenString(tok), s.TokenText())
	}
}

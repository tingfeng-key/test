package src
import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"encoding/base64"
	"strconv"
	"fmt"
)
type CNum struct{
	code string
	numF string
	numS string
	numT string
	big bool
	odd bool
}

func (cnum *CNum) test()  {
	fmt.Print("test: "+cnum.code+"\n");
}
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
func ExampleScrape() {
	doc, err := goquery.NewDocument("http://shk3.icaile.com/")
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("td.chart-bg-qh").Each(func(i int, s *goquery.Selection) {
		numF := s.SiblingsFiltered("td.cc").Eq(0).Text();
		numS := s.SiblingsFiltered("td.cc").Eq(1).Text();
		numT := s.SiblingsFiltered("td.cc").Eq(2).Text();
		numFs,err := base64.StdEncoding.DecodeString(Substr(numF, 1, 4));
		if(err != nil){
			log.Fatalln(err.Error(), string([]rune(numF)[1:4]));
		}
		numSs,err := base64.StdEncoding.DecodeString(Substr(numS, 1, 4));
		if(err != nil){
			log.Fatalln(err.Error());
		}
		numTs,err := base64.StdEncoding.DecodeString(Substr(numT, 1, 4));
		if(err != nil){
			log.Fatalln(err.Error());
		}
		numfInt,_ := strconv.Atoi(string(numFs));
		numSInt,_ := strconv.Atoi(string(numSs));
		numTInt,_ := strconv.Atoi(string(numTs));
		count := numfInt+numSInt+numTInt;
		big := count >10;
		odd := ((count % 2) == 0)
		num := CNum{
			code: s.Text(),
			numF: string(numFs),
			numS: string(numSs),
			numT: string(numTs),
			big: big,
			odd: odd,
		}
		fmt.Printf("%v\n", num)
		num.test();
	})
}

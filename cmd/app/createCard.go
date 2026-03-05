package main

type CreaterCard interface{}

type CreateCard struct {
	cNum, cSymbol, cards []string
}

func (cc *CreateCard) SetCreateCard(cNum, cSymnbol, cards []string) {
	cc.cNum = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13"}
	cc.cSymbol = []string{"♥", "♦", "♠", "♣"}
	cc.cards = cards
	//cc.exeCreateCard()

}

func (cc *CreateCard) GetCreateCard() ([]string, []string, []string) {
	return cc.cNum, cc.cSymbol, cc.cards
}

func (cc *CreateCard) exeCreateCard() []string {
	cc.cards = make([]string, 0, len(cc.cNum)*len(cc.cSymbol))
	for i := 0; i < len(cc.cSymbol); i++ {
		for j := 0; j < len(cc.cNum); j++ {
			cc.cards = append(cc.cards, cc.cSymbol[i]+cc.cNum[j])
		}
	}
	return cc.cards

}

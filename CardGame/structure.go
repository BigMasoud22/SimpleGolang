package main

type Player struct {
	Name  string
	Cards deck
}

func (player Player) Hand() deck {
     
	if player.Cards != nil {
		return player.Cards
	} else {
		newhand, err := deal(newDeck(), 13)
		if err == nil && len(newhand) != 0 {
			return newhand
		}
	}
	//If we got here there is something wrong with our code
	return deck{}
}

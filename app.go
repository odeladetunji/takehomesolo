
package main

import (
	"math/rand"
)

type BinName string;
type BagName string;


type Ticket struct {
	BinName string `json:"binName"`
	BagId int `json:"bagId"`
}

//BAG (BagId is also present on each Ticket) a one to one relationship
type Bag struct {
	BagId int `json:"bagId"`
	Size int `json:"size"`
	BagType string `json:"bagType"`
	BagName string `json:"bagName"`
}

//Bin
type Bin struct {
	Size int `json:"size"`
	Compactment map[BagName]Bag `json:"compactment"`
}

//Storage
type Storage map[BinName]Bin;
var storage Storage = make([BinName]Bin);

//Line
type Line Storage;


type BaggageConcierge interface {    
	Store(bag Bag) Ticket    
	Retrieve(ticket Ticket) Bag
}

type BaggageManagment struct {

}

func (bgg *BaggageManagment) CreateNewBin(bag Bag){

	var bin Bin; 
    bin.Size = 100;
	
	var storage Storage;
	var binName string = "bin1";
	storage[binName] = bin;

    return bin;
}

func Store(bag Bag) Ticket, error {
	
	var ticket Ticket;
	var isFullForCarryon bool = true;
	var isFullForCheckout bool = true;
    
	if len(storage) == 0 {
		ticket = CreateNewBin(bag);
		return ticket, nil;
	}

	storeInBin := func(bin, key) Ticket {
		var compactment map[BagName]Bag; 
		compactment[bag.BagName] = bag;

		bin.Size == 0;
		bin.Compactment = compactment;

		ticket.BagId = bag.BagId;
		ticket.BinName = key;
	}

    saveABag := func(bin Bin, key string) Ticket {
		switch bag.BagType {
			case "Carry-ons": 
				if bin.Size != 0 {
                    if bag.Size <= bin.Size {
						ticket = storeInBin(bin, key);
						isFullForCarryon = false;
						break;
					}
				}
				
			case "Checked": 
				if bin.Size == 100 {
					ticket = storeInBin(bin, key);
					isFullForCheckout = false;
					break;
				}
		}
	}

	for key, bin := range storage {
		ticket = saveABag(bin, key);
		if !isFullForCheckout {
			break;
		}
		if !isFullForCarryon {
			break;
		}
	}

	if isFullForCarryon {
		return ticket, errors.New("Its full for two bags")
	}

	if isFullForCheckout {
		return ticket, errors.New("Its full for one bag")
	}

	return ticket, nil;

}

func Retrieve(ticket Ticket) Bag {
     return storage[ticket.BinName][ticket.BagId];  //storage is basically a map
}

func generateBag(bagType string, bagName string) Bag {
	var bag Bag;
	bag.BagId = rand.Int();
	bag.Size = 100;
	bag.BagType = bagType;
	bag.BagName = bagName;

	return bag;
}

func main() {

    var newBag Bag = generateBag("Checked", fmt.Sprintln(rand.Int()));
	
    var ticket, err = Store(newBag);
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(ticket);

	var bag Bag = Retrieve(ticket);
	fmt.Println(bag);
	

}





































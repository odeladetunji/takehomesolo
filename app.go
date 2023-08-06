
package main

import (
	"math/rand"
	"errors"
	"fmt"
)

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
	Compactment map[string]Bag `json:"compactment"`
}

//Storage
type Storage map[string]Bin;
var storage Storage = make(map[string]Bin);

//Line
type Line Storage;

type BaggageConcierge interface {    
	Store(bag Bag) Ticket    
	Retrieve(ticket Ticket) Bag
}

type BaggageManagment struct {

}

func (bgg *BaggageManagment) CreateNewBin(bag Bag) {

	var bin Bin; 
    bin.Size = 100;
	
	var binName string = "bin1";
	storage[binName] = bin;

}

func Store(bag Bag) (Ticket, error) {
	
	var ticket Ticket;
	var isFullForCarryon bool = true;
	var isFullForCheckout bool = true;

	if len(storage) == 0 {
		var baggageManagement BaggageManagment;
		baggageManagement.CreateNewBin(bag);
		return ticket, nil;
	}

	storeInBin := func(bin Bin, key string) Ticket {
		var compactment map[string]Bag; 
		compactment = make(map[string]Bag)
		compactment[bag.BagName] = bag;

		bin.Size = 0;
		bin.Compactment = compactment;

		ticket.BagId = bag.BagId;
		ticket.BinName = key;

		return ticket;
	}

    saveABag := func(bin Bin, key string) Ticket {
		switch bag.BagType {
			case "Carryon": 
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

	    return ticket;
	}

	for key, aBin := range storage {
		ticket = saveABag(aBin, key);
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
     return storage[ticket.BinName].Compactment[fmt.Sprint(ticket.BagId)];  //storage is basically a map
}

func generateBag(bagType string, bagName string, bagSize int) Bag {
	var bag Bag;
	bag.BagId = rand.Int();

	if bagType == "Checked" {
		bag.Size = 100;
	}else { 
		bag.Size = bagSize;
	}
	
	bag.BagType = bagType;
	bag.BagName = bagName;

	return bag;
}

func main() {

    var checkedBag Bag = generateBag("Checked", fmt.Sprintln(rand.Int()), 0);
	var ticket, err = Store(checkedBag);
	if err != nil {
		fmt.Println(err.Error())
		// panic(err.Error())
	}

	fmt.Println(ticket);
	var carryonBag Bag = generateBag("Carryon", fmt.Sprintln(rand.Int()), 50);
	var ticket2, err2 = Store(carryonBag);
	if err2 != nil {
		fmt.Println(err2.Error())
		// panic(err2.Error())
	}
    
	fmt.Println(ticket2);

	var bag Bag = Retrieve(ticket); 
	fmt.Println(bag);

	var line Storage;
	line = storage;

	fmt.Println(line)
	
}

/* SECTION 2  */


//HOW WOULD I WRITE UNIT TESTS

/*  
   >> FIRST TEST CASE 
      I will try and pass 2 large bags in a bin, if it succeeds, the algorithm is falty
      if it fails the algorithm pass this perticular test

   >> SECOND TEST CASE
      I will try and fit a small bag and a large bag in a bin. acceptable result should fail

   >>> I will try and pass 3 small bags to a bin. acceptable result should be a failure.
   
*/


//HOW DOES THE ALGORITHM PERFORM WHEN SCALED TO LARGE SIZE?
/* 
        Scalability was considered when designing this algorithm, HashMap Data structure was used to implememt 
the Storage and each individual Bin. this ensures that a Big O notation is O(1) (constant time) irrespective of 
the Data Size.  
   */

// BETTER DATASTRUCTURES 
/* Prioritiy Queue can actually be used to manage the available space in the Bin



































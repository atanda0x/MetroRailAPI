package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can' find the product")
	ErrCantDEecodeProuct  = errors.New("can't find the product")
	ErrUserIDIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("can't remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to getthe item from the cart")
	ErrCantBuyItemItem    = errors.New("can't update the purhase")
)

func AddProduct() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}

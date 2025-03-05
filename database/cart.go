package controllers

import "errors"

var (

	ErrCantFindProduct = errors.New("Can't find the product")
	ErrCantDecodeProducts = errors.New("Can't find the product")
	ErrUserIdIsNotValid = errors.New("User id is not valid")
	ErrCantUpdateUser = errors.New("Can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("Can't remove this item from the cart")
	ErrCantGetItem = errors.New("Can't get this item from the cart")
	ErrCantBuyCartItem = errors.New("Can't buy this item from the cart")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
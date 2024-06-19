package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("cant find product")
	ErrCantDecodeProduct  = errors.New("cant decode product ")
	ErrUserIdIsNotValid   = errors.New("user id is not valid ")
	ErrCantUpdateUser     = errors.New("cant update user ")
	ErrCantRemoveItemCart = errors.New("cant remove item cart ")
	ErrCantGetItem        = errors.New("cant get item ")
	ErrCantBuyCartItem    = errors.New("cant buy cart item ")
)

func AddToCart()       {}
func RemoteItem()      {}
func GetItemFromCart() {}
func BuyFromCart()     {}
func InstantBuy()      {}

package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashgenerate(data string)string{
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil{
		log.Fatal(err)
	}
	return string(hash)
}

func verifyhashdata(data string,hash string)bool{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data)) == nil
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type URLShortener struct {
	urls map{string}string
}
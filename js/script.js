// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Charlotte Jhu
// Created on: March 2023
// This file contains the JS functions for index.html

'use strict'

const randomNumber = Math.floor(Math.random() * 6) + 1 

function myButtonClicked() {
  // input
  const numberGuessed = parseInt (document.getElementById("inputValue").value)

  // process
  if (numberGuessed == randomNumber) {
    // output
    document.getElementById("answer").innerHTML = "You got it!"
  }
  if (numberGuessed != randomNumber) {
    // output
    document.getElementById("answer").innerHTML = "Sorry, you are incorrect."
  }
}

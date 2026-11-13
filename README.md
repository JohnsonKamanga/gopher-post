# Gopher Post

## Introduction

Gopher Post simple message brocker built with golang. The main purpose of this project is to investigate how messaging queues work under the hood. 


### The Messaging Protocol : Gopher Messaging Protocol

As part of the learning exercise, I will build a custom Application Layer Protocol on top of TCP. The protocol implementation details are found within the protocol package. The default port is `8090`.

#### 1. Overview

#### 1.A Purpose 
The purpose of this protocol is to facilitate communication between the Producer & the Message Brocker, and the Consumer & the Message Brocker. The scope of this implementation will not include authentication for now, because this project is not intended to be used in any production environments. 


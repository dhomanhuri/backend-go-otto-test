# Voucher app

<!-- ABOUT THE PROJECT -->

## 💻 About The Project

Backend Developer task for join opportunity otto digital as backend developer using golang

## 💻 Step to run program
clone this repository
create database mysql
create .env like .env.example


## 💻 Api Doc

<div>
  
| Feature User | Endpoint | Param | Function |
| --- | --- | --- | --- |
| POST | /api/brand  | - | add new brand |
| GET | /api/brand | id | get brand by id |
| POST | /api/voucher | - | add new voucher |
| GET | /api/voucher | id | get voucher by id |
| GET | /api/voucher/brand | id | get list voucher by brand id |
| POST | /api/transaction | - | add new transaction |
| POST | /api/transaction/redemption | - | add voucher in transaction |
| GET | /api/transaction/redemption | transactionId | get detail transaction by transaction id |


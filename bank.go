package main

import (
    "errors"
    "fmt"
    "time"
)

type Account struct {
    AccountNumber string
    HolderName    string
    Balance       float64
    AccountType   string
}

type Transaction struct {
    ID        string
    Type      string
    Amount    float64
    Timestamp time.Time
}

type Bank struct {
    Name              string
    Accounts          map[string]*Account
    TransactionHistory []*Transaction
}

type AccountOperations interface {
    Deposit(amount float64) error
    Withdraw(amount float64) error
    GetBalance() float64
}

func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("amount must be greater than zero")
    }
    a.Balance += amount
    return nil
}

func (a *Account) Withdraw(amount float64) error {
    if amount > a.Balance {
        return errors.New("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

func (a *Account) GetBalance() float64 {
    return a.Balance
}

func displayMenu() {
    fmt.Println("1. Create Account")
    fmt.Println("2. Deposit Money")
    fmt.Println("3. Withdraw Money")
    fmt.Println("4. Check Balance")
    fmt.Println("5. Exit")
}

func getUserInput() string {
    var input string
    fmt.Scanln(&input)
    return input
}

func (bank *Bank) CreateAccount(holderName string, accountType string) *Account {
    accountNumber := generateAccountNumber()
    account := &Account{
        AccountNumber: accountNumber,
        HolderName:    holderName,
        Balance:       0.0,
        AccountType:   accountType,
    }
    bank.Accounts[accountNumber] = account
    fmt.Printf("Account created: %s\n", accountNumber)
    return account
}

func generateAccountNumber() string {
    return fmt.Sprintf("%d", time.Now().UnixNano())
}

func (bank *Bank) AddTransaction(t *Transaction) {
    bank.TransactionHistory = append(bank.TransactionHistory, t)
}

func main() {
    bank := &Bank{
        Name:              "Go Bank",
        Accounts:          make(map[string]*Account),
        TransactionHistory: []*Transaction{},
    }

    for {
        displayMenu()
        choice := getUserInput()

        switch choice {
        case "1":
            fmt.Println("Enter holder name:")
            holderName := getUserInput()
            fmt.Println("Enter account type: Savings or Checking")
            accountType := getUserInput()
            bank.CreateAccount(holderName, accountType)
        case "2":
            fmt.Println("Enter account number:")
            accountNumber := getUserInput()
            account, exists := bank.Accounts[accountNumber]
            if !exists {
                fmt.Println("Account not found.")
                continue
            }
            fmt.Println("Enter deposit amount:")
            var amount float64
            fmt.Scanf("%f", &amount)
            err := account.Deposit(amount)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                bank.AddTransaction(&Transaction{
                    ID:        generateAccountNumber(),
                    Type:      "deposit",
                    Amount:    amount,
                    Timestamp: time.Now(),
                })
                fmt.Println("Deposit successful.")
            }
        case "3":
            fmt.Println("Enter account number:")
            accountNumber := getUserInput()
            account, exists := bank.Accounts[accountNumber]
            if !exists {
                fmt.Println("Account not found.")
                continue
            }
            fmt.Println("Enter withdrawal amount:")
            var amount float64
            fmt.Scanf("%f", &amount)
            err := account.Withdraw(amount)
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                bank.AddTransaction(&Transaction{
                    ID:        generateAccountNumber(),
                    Type:      "withdrawal",
                    Amount:    amount,
                    Timestamp: time.Now(),
                })
                fmt.Println("Withdrawal successful.")
            }
        case "4":
            fmt.Println("Enter account number:")
            accountNumber := getUserInput()
            account, exists := bank.Accounts[accountNumber]
            if !exists {
                fmt.Println("Account not found.")
                continue
            }
            fmt.Printf("Balance: %.2f\n", account.GetBalance())
        case "5":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
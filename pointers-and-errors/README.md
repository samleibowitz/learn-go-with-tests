# Pointers and Errors

goddamn stupid bitcoin.

## Remember pass by value / reference?

God awful flashbacks to my undergrad work learning Borland C++. But the important bit here is that the **arguments to a method are copies, not the originals**, unless you specify a value.

You can get the original location in memory of a function with the address operator:

```go
func (w Wallet) Deposit(amount int) {
    fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
    w.balance += amount
}
```
Output:
```
address of balance in Deposit is 0xc420012268
```

So if you want to pass the real deal, you can define it as taking a *pointer* to a Wallet.

```go
func (w *Wallet) Deposit(amount int) {
  // notice we don't have to explicitly dereference the pointer like we did in C++,
  // because Go isn't that dumb.  
  w.balance += amount
}
```

## Custom Types

You can define new types from existing ones using the syntax `type NewType OldType`.  Cool if you want to implement some custom stuff on your type.

In this example, we created a custom interface called Stringer, and applied it to our new Bitcoin class:

```go
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```

## Errors

So first of all, Go has a `nil`. Neat!

Secondly: Go has its own error class. It's not an exception. Instead, it's similar to stuff you've worked on in Node.js when calling a function that returns two items - the thing you actualy want, and (potentially) an error. So you'll see constructions like:

```go
func main() {  
    f, err := os.Open("/test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(f.Name(), "opened successfully")
}
```

We now use an Err object to inform the developer when trying to withdraw too much money from an account:

```go
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//...

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
```

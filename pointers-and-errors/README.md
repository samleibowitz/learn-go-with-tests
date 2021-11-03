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


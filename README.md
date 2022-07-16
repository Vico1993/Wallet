# Wallet

Small cli application to help myself deal with my Crypto Wallet.
Trying to see based on my history. Which of my trade is good and which one is bad.

Also an excuse to try: [lipgloss](https://github.com/charmbracelet/lipgloss) and Golang. To make fancy CLI

## Getting started

1. Clone the project

```shell
git clone git@github.com:Vico1993/Wallet.git
```

2. Add your asset with Crypto.com csv

```shell
make build && ./wallet add -p=<PATH_TO_MY_CSV>
```

3. List your asset and Enjoy

```shell
make build && ./wallet list -c -u=BTC
```

## Testing

You can run the unit and integration tests:

```shell
make test
```

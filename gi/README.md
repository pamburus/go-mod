# gi [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

Module `gi` provides a set of generic iterator combinators and corresponding helpers.

## Packages

- [gi](./)

    Package [gi](https://pkg.go.dev/github.com/pamburus/go-mod/gi) provides a set of generic iterator combinators over scalar values [iter.Seq](https://pkg.go.dev/iter#Seq).

- [gi2](./gi2/)

    Package [gi2](https://pkg.go.dev/github.com/pamburus/go-mod/gi/gi2) provides a set of generic iterator combinators over pairs of values [iter.Seq2](https://pkg.go.dev/iter#Seq2).

- [giopt](./giopt/)

    Package [giopt](https://pkg.go.dev/github.com/pamburus/go-mod/gi/giopt) provides a set of generic iterator combinators over scalar values [iter.Seq](https://pkg.go.dev/iter#Seq) adapted for [optval.Value](https://pkg.go.dev/github.com/pamburus/go-mod/optional/optval#Value).

- [gi2opt](./gi2opt/)

    Package [gi2opt](https://pkg.go.dev/github.com/pamburus/go-mod/gi/gi2opt) provides a set of generic iterator combinators over pairs of values [iter.Seq](https://pkg.go.dev/iter#Seq2) adapted for [optpair.Pair](https://pkg.go.dev/github.com/pamburus/go-mod/optional/optpair#Pair).

- [giop](./giop/)

    Package [giop](https://pkg.go.dev/github.com/pamburus/go-mod/gi/giop) provides a set of binary operators on scalar values for generic iterator combinators.

- [gi2op](./gi2op/)

    Package [gi2op](https://pkg.go.dev/github.com/pamburus/go-mod/gi/gi2op) provides a set of binary operators on pairs of values for generic iterator combinators.


[doc-img]: https://pkg.go.dev/badge/github.com/pamburus/go-mod/gi
[doc]: https://pkg.go.dev/github.com/pamburus/go-mod/gi
[ci-img]: https://github.com/pamburus/go-mod/actions/workflows/ci.yml/badge.svg
[ci]: https://github.com/pamburus/go-mod/actions/workflows/ci.yml
[cov-img]: https://codecov.io/gh/pamburus/go-mod/graph/badge.svg?flag=gi&token=CC2G17UKAS
[cov]: https://app.codecov.io/gh/pamburus/go-mod/tree/main/gi

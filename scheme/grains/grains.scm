(define-module (grains)
  #:export (square total)
  #:autoload (srfi srfi-1) (iota))

(define (square n)
  (expt 2 (1- n))
)

(define (total)
  18446744073709551615 ; (2^64 - 1)
)

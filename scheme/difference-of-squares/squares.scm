(define-module (squares)
  #:export (sum-of-squares
            square-of-sums
            difference)
  #:autoload (srfi srfi-1) (reduce iota))

(define (square x) (* x x))

(define (sum-of-squares n)
  (reduce + 0 (map square (iota n 1)))
)

(define (square-of-sums n)
  (square (reduce + 0 (iota n 1)))
)

(define (difference n)
  (- (square-of-sums n) (sum-of-squares n))
)

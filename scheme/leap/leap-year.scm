(define-module (leap-year)
  #:export (leap-year?))

(define (leap-year? y)
  (or
    (and (> (remainder y 100) 0) (= (remainder y 4) 0))
    (= (remainder y 400) 0)
  )
)

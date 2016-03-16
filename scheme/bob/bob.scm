(define-module (bob)
  #:export (response-for))

(define (response-for s)
  (let
    ((ss (string-trim-both s)))
    (cond
      ((string-null? ss) "Fine. Be that way!")
      ((and (string-any char-alphabetic? ss) (= (string-count ss char-lower-case?) 0)) "Whoa, chill out!")
      ((string-suffix? "?" ss) "Sure.")
      (else "Whatever.")
    )
  )
)

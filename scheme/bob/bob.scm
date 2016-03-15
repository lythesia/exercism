(define-module (bob)
  #:export (response-for))

(define (response-for s)
  (let*
    ((ss (string-trim-both s)) (as (string-filter char-alphabetic? ss)))
    (cond
      ((string-null? ss) "Fine. Be that way!")
      ((and (not (string-null? as)) (string-every char-upper-case? as)) "Whoa, chill out!")
      ((equal? (string-ref ss (1- (string-length ss))) #\?) "Sure.")
      (else "Whatever.")
    )
  )
)

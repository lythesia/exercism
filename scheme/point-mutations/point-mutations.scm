(define-module (point-mutations)
  #:export (hamming-distance))

(define (hamming-distance s1 s2)
  (define (iter r1 r2)
    (cond
      ((and (null? r1) (null? r2)) 0)
      ((and (not (null? r1)) (not (null? r2)))
        (let ((rest (iter (cdr r1) (cdr r2))))
          (if (not (eqv? (car r1) (car r2)))
            (1+ rest)
            rest
          )
        )
      )
      (else
        (error "unequal length" s1 s2)
      )
    )
  )
  (iter (string->list s1) (string->list s2))
)

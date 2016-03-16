(define-module (raindrops)
  #:export (convert))

(define (convert n)
  (let ((s
    (string-concatenate
      (map
        (lambda (elem)
          (if (zero? (remainder n (car elem))) (cdr elem) "")
        )
        '((3 . "Pling") (5 . "Plang") (7 . "Plong"))
      )
    )))
    (if (string-null? s) (number->string n) s)
  )
)

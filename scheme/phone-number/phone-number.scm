(define-module (phone-number)
  #:export (numbers area-code pprint)
  #:autoload (ice-9 format) (format))

(define (numbers number-string)
  (let*
    ((number (string-filter char-numeric? number-string))
     (n (string-length number)))
    (cond
      ((= n 10) number)
      ((and (= n 11) (char=? (string-ref number 0) #\1)) (string-drop number 1))
      (else "0000000000")
    )
  )
)

(define (area-code number-string)
  (string-take (numbers number-string) 3)
)

(define (pprint number-string)
  (let ((valid (numbers number-string)))
    (format #f "(~a) ~a-~a" (area-code valid) (substring valid 3 6) (string-take-right valid 4))
  )
)

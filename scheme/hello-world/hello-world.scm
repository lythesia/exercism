(define-module (hello-world)
  #:export (hello))

;; define* or lambda* would allow opt-arg
(define* (hello #:optional (name "World"))
  (string-concatenate (list "Hello, " name "!"))
)

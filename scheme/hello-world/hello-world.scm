(define-module (hello-world)
  #:export (hello))

;; define* or lambda* would allow opt-arg
(define* (hello #:optional name)
  (string-concatenate (list "Hello, " (or name "World") "!"))
)

(define-module (dna)
  #:export (to-rna))

(define (to-rna dna)
  (let ((alist '((#\G . #\C) (#\C . #\G) (#\T . #\A) (#\A . #\U))))
    (string-map (lambda (c) (assoc-ref alist c)) dna)
  )
)
